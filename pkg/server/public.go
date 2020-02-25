package server

import (
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"blazetunnel/common"
)

func (s *Server) initPublic() error {
	cfg := generateTLSConfig()
	fmt.Println("Allowed protos: ", cfg.NextProtos)
	cfg.NextProtos = []string{"http/1.1", "acme-tls/1", "quic-echo-example"}
	ln, err := tls.Listen("tcp", ":443", cfg)
	if err != nil {
		return err
	}

	s.publicListener = ln
	return nil
}

func (s *Server) startPublic() {
	for {
		conn, err := s.publicListener.Accept()
		if err != nil {
			fmt.Printf("[server:publicListener] unable to accept connection: %s\n", err)
			continue
		}

		go s.handlePublic(conn)
	}
}

func (s *Server) handlePublic(conn net.Conn) {
	defer conn.Close()
	tlsconn, ok := conn.(*tls.Conn)
	if !ok {
		return
	}

	err := tlsconn.Handshake()
	if err != nil {
		fmt.Printf("[server:publicListener] unable to process handshake: %s\n", err)
		return
	}

	connState := tlsconn.ConnectionState()

	if connState.ServerName == "" {
		fmt.Printf("[server:publicListener] unable to process handshake: No SNI found\n")
		return
	}

	fmt.Println("Connecting to : ", connState.ServerName)
	rwc, err := s.hostmap.NewStreamFor(connState.ServerName)
	if err != nil {
		fmt.Printf("[server:publicListener] unable to open a client stream: %s\n", err)
		return
	}

	defer rwc.Close()

	crwc := common.NewCompressedStream(rwc)

	go io.Copy(crwc, conn)
	if _, err := io.Copy(conn, crwc); err != nil {
		fmt.Printf("[server:publicListener] unable to open a client stream: %s\n", err)
		return
	}
}
