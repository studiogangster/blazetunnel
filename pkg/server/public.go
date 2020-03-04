package server

import (
	"blazetunnel/common"
	"bufio"
	"bytes"
	"errors"
	"log"
	"net"
	"strings"
	"time"

	"acln.ro/zerocopy"
)

// Constants
const hostPrefix, crlf, nextLine = "Host: ", "\r\n", byte('\n')

//

func (s *Server) initPublic() error {
	cfg := generateTLSConfig()
	log.Println("Allowed protos: ", cfg.NextProtos)
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
			log.Printf("[server:publicListener] unable to accept connection: %s\n", err)
			continue
		}

		go s.handlePublic(conn)
	}
}

const timeoutDuration = 3 * time.Second

// Finding host over TCP connection
func findHost(conn net.Conn, buf *bufio.Reader) (err error, Host string, buffer bytes.Buffer) {

	// var buf bytes.Buffer
	// tee := io.TeeReader(conn, &buf)
	err = errors.New("Host header not found")
	Host = ""

	for {

		conn.SetReadDeadline(time.Now().Add(timeoutDuration))

		// TODO: Set readtimeout accordingly
		// Need to make it configuratble as well
		// TODO: Add maximum header size, max number of headers too, to avoid DOS attacks
		// will listen for message to process ending in newline (\n)
		var message string
		message, err = buf.ReadString(nextLine)
		// Copy message to header
		buffer.Write([]byte(message))
		if err != nil {
			log.Println("Error", err)
			buffer.Reset()
			return
		}

		if message == crlf {
			// log.Println("Request ")
			// buffer.Write([]byte(CRLF))
			// Request Headers ended
			return
		}

		if strings.HasPrefix(message, hostPrefix) {

			message = message[6:]
			Host = strings.Split(message, ":")[0]
			log.Println("HOST Request %s", Host)
			err = nil
			// Host header found in request header
			break
		}
	}
	return
}

func (s *Server) handlePublic(conn net.Conn) {

	defer conn.Close()

	originaReader := bufio.NewReader(conn)
	err, ServerName, preRequest := findHost(conn, originaReader)
	// ServerName := "quick.server"
	// var err error

	log.Println("host found", ServerName)
	if err != nil {
		log.Println("Error occured while finding host", err)
		return
	}

	conn.SetReadDeadline(time.Time{})

	rwc, err := s.hostmap.NewStreamFor(ServerName)
	if err != nil {
		log.Printf("[server:publicListener] unable to open a client stream: %s\n", err)
		return
	}

	defer rwc.Close()
	crwc := common.NewCompressedStream(rwc)
	// Not doing in go routine...Can be improved
	_, err = crwc.Write(preRequest.Bytes())

	if err != nil {
		log.Printf("[Error while writting header response to tunnel")
	}

	go func() {
		zerocopy.Transfer(crwc, originaReader)
		log.Println("copying data")
	}()
	if _, err := zerocopy.Transfer(conn, crwc); err != nil {
		log.Printf("[server:publicListener] unable to open a client stream: %s\n", err)
		return
	}
}
