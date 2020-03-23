package server

import (
	"blazetunnel/common"
	"context"
	"log"
	"time"

	"github.com/lucas-clemente/quic-go"
)

const handshakeTimeout = 5

var newmsg = common.NewMessage

func (s *Server) initTunnel() error {
	cfg := generateTLSConfig()
	log.Println("Allowed protos: ", cfg.NextProtos)
	cfg.NextProtos = []string{"h2", "http/1.1", "acme-tls/1", "quic-echo-example"}
	ln, err := quic.ListenAddr(":2723", cfg, &quic.Config{
		IdleTimeout: time.Second * time.Duration(s.idleTimeout),
	})
	if err != nil {
		return err
	}

	s.tunnelListener = ln
	return nil
}

func (s *Server) startTunnel() {
	for {
		session, err := s.tunnelListener.Accept(context.Background())
		if err != nil {
			log.Printf("[server:tunnelListener] unable to open a client session : %s\n", err)
			continue
		}

		go s.handleTunnelSession(session)
	}
}

// Session handles all the incoming streams of the QUIC connections.
// Generally the client opens a master stream which acts as the command buffer
// for the communication between the server and the client. When a new TCP connection
// is made to the public handler, the server sends a command to client to open a new
// client stream with a specific ID which can be used to relay/proxy the incoming TCP with
// the tunneling connection
func (s *Server) handleTunnelSession(session quic.Session) {
	ctlStream, err := session.AcceptStream(context.Background())
	if err != nil {
		log.Printf("[server:tunnelListener] unable to accept a client stream : %s\n", err)
		return
	}

	close := func() {
		ctlStream.Close()
		session.Close()
	}

	ctlStream.SetReadDeadline(time.Now().Add(time.Duration(handshakeTimeout) * time.Second))
	m, err := newmsg("", "").DecodeFrom(ctlStream)
	ctlStream.SetDeadline(time.Time{})
	if err != nil {
		log.Printf("[server:tunnelListener] unable to decode msgpack: %s\n", err)
		close()
		return
	}
	if m.Command == common.CommandAuthClient {
		log.Printf("[server:tunnelListener] Authenticating: %s %s\n", m.Command, m.Context)

		ctlStream.SetWriteDeadline(time.Now().Add(time.Duration(handshakeTimeout) * time.Second))
		err := newmsg(common.CommandAuthServer, m.Context).EnryptTo(ctlStream)
		ctlStream.SetWriteDeadline(time.Time{})

		log.Println("Error", err)
		close()
		return
	}

	if m.Command != common.CommandNewClient {

		log.Println("[server:tunnelListener] expected NewClient command, got:", m.Command, err)
		close()
		return
	}

	// Authenticate server

	err = m.Authenticate()

	if err != nil {
		close()
		log.Println("Authenticated failed", m.Context)
		return
	}

	serviceName := m.Context

	exposedDomain := serviceName + "." + s.domain

	ctlStream.SetWriteDeadline(time.Now().Add(time.Duration(handshakeTimeout) * time.Second))

	err = newmsg(common.CommandSetConfig, exposedDomain).EncodeTo(ctlStream)
	ctlStream.SetWriteDeadline(time.Time{})

	if err != nil {
		log.Printf("[server:tunnelListener] unable to encode to msgpack: %s\n", err)
		close()
		return
	}

	// Add rw lock.
	ok := s.hostmap.Put(exposedDomain, &TunnelState{
		session:   session,
		ctlStream: ctlStream,
	})

	if !ok {
		log.Printf("[server:tunnelListener] server host config already found.\n")
		log.Printf("[server:tunnelListener] trying to cose older and connect again.\n")

		tunnelRef, ifExists := s.hostmap.Get(exposedDomain)
		if ifExists {
			tunnelRef.Close()
			s.hostmap.Delete(exposedDomain)
		}

		log.Printf("[server:tunnelListener] Older connection closed & deleted\n")

		log.Printf("[server:tunnelListener] Trying to update hostmap again\n")
		// Try to put again.
		ok = s.hostmap.Put(exposedDomain, &TunnelState{
			session:   session,
			ctlStream: ctlStream,
		})

		// close()
	}

	if !ok {
		log.Printf("[server:tunnelListener] Failed again..Closing this connection\n")
		close()

	}

	getOut := false

	for !getOut {
		m, err := newmsg("", "").DecodeFrom(ctlStream)
		if err != nil {
			log.Printf("[server:pong] unable to decode from msgpack: %s\n", err)
			close()
			return
		}
		switch m.Command {
		case common.CommandPingPeer:
			log.Printf("[server:message] Got ping from %s\n", session.RemoteAddr())
			err = newmsg(common.CommandPongPeer, "").EncodeTo(ctlStream)
			if err != nil {
				log.Printf("[server:pong] unable to encode to msgpack: %s\n", err)
				close()
				getOut = true
				break
			}
		}
	}

	return
}
