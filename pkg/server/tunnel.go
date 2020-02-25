package server

import (
	"blazetunnel/common"
	"context"
	"fmt"
	"time"

	"github.com/lucas-clemente/quic-go"
)

var newmsg = common.NewMessage

func (s *Server) initTunnel() error {
	cfg := generateTLSConfig()
	fmt.Println("Allowed protos: ", cfg.NextProtos)
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
			fmt.Printf("[server:tunnelListener] unable to open a client session : %s\n", err)
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
		fmt.Printf("[server:tunnelListener] unable to accept a client stream : %s\n", err)
		return
	}

	close := func() {
		ctlStream.Close()
		session.Close()
	}

	m, err := newmsg("", "").DecodeFrom(ctlStream)
	if err != nil {
		fmt.Printf("[server:tunnelListener] unable to decode msgpack: %s\n", err)
		close()
		return
	}
	if m.Command != common.CommandNewClient {
		fmt.Printf("[server:tunnelListener] expected NewClient command, got: %s\n", m.Command)
		close()
		return
	}

	exposedDomain := genRandomString() + "." + s.domain
	err = newmsg(common.CommandSetConfig, exposedDomain).EncodeTo(ctlStream)
	if err != nil {
		fmt.Printf("[server:tunnelListener] unable to encode to msgpack: %s\n", err)
		close()
		return
	}

	// Add rw lock.
	ok := s.hostmap.Put(exposedDomain, &TunnelState{
		session:   session,
		ctlStream: ctlStream,
	})
	if !ok {
		fmt.Printf("[server:tunnelListener] server host config already found\n")
		close()
	}

	getOut := false

	for !getOut {
		m, err := newmsg("", "").DecodeFrom(ctlStream)
		if err != nil {
			fmt.Printf("[server:pong] unable to decode from msgpack: %s\n", err)
			close()
			return
		}
		switch m.Command {
		case common.CommandPingPeer:
			fmt.Printf("[server:message] Got ping from %s\n", session.RemoteAddr())
			err = newmsg(common.CommandPongPeer, "").EncodeTo(ctlStream)
			if err != nil {
				fmt.Printf("[server:pong] unable to encode to msgpack: %s\n", err)
				close()
				getOut = true
				break
			}
		}
	}

	return
}
