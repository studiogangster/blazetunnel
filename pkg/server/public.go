package server

import (
	"blazetunnel/common"
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
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

// Finding host over TCP connection
func findHost(conn net.Conn) (err error, Host string, buffer bytes.Buffer) {

	err = errors.New("Host header not found")

	var buf = bufio.NewReader(conn)

	Host = ""

	CRLF := "\r\n"

	for {
		// will listen for message to process ending in newline (\n)

		var message string
		message, err = buf.ReadString('\n')

		if err != nil {
			log.Println("Error", err)
			buffer.Reset()
			return
		}

		// Copy message to header
		buffer.Write([]byte(message))

		if message == CRLF {
			log.Println("End")
			// buffer.Write([]byte(CRLF))
			// Request Headers ended
			return
		}

		if strings.HasPrefix(message, "Host: ") {
			message = strings.Split(message, "Host: ")[1]
			Host = strings.Split(message, ":")[0]
			err = nil
			// Host header found in request header
			break

		}

	}

	return
}

func (s *Server) handlePublic(conn net.Conn) {

	defer conn.Close()

	err, ServerName, reqHeaderConn := findHost(conn)

	if err != nil {
		log.Println("Error occured while finding host", err)
		return
	}

	log.Println("Host found", ServerName)

	// ServerName := "quic.meddler.xyz"

	fmt.Println("Connecting to : ", ServerName) //, conn.RemoteAddr(), conn.LocalAddr())
	rwc, err := s.hostmap.NewStreamFor(ServerName)
	if err != nil {
		fmt.Printf("[server:publicListener] unable to open a client stream: %s\n", err)
		return
	}

	defer rwc.Close()

	crwc := common.NewCompressedStream(rwc)

	// Not doing in go routine...Can be improved
	_, err = rwc.Write(reqHeaderConn.Bytes())

	if err != nil {
		fmt.Printf("[Error while writting header response to tunnel")
	}

	go io.Copy(crwc, conn)
	if _, err := io.Copy(conn, crwc); err != nil {
		fmt.Printf("[server:publicListener] unable to open a client stream: %s\n", err)
		return
	}
}
