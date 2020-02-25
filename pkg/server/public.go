package server

import (
	"blazetunnel/common"
	"fmt"
	"io"
	"net"
)

func (s *Server) initPublic() error {
	cfg := generateTLSConfig()
	fmt.Println("Allowed protos: ", cfg.NextProtos)
	cfg.NextProtos = []string{"http/1.1", "acme-tls/1", "quic-echo-example"}
	ln, err := net.Listen("tcp", ":443")
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

	ServerName := "quic.meddler.xyz"

	fmt.Println("Connecting to : ", ServerName)
	rwc, err := s.hostmap.NewStreamFor(ServerName)
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
