package server

import (
	"blazetunnel/common"
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"log"
	"net"
	"strings"

	"acln.ro/zerocopy"
)

// Constants
const HostPrefix = "Host: "
const Crlf = "\r\n"
const NextLine = byte('\n')

//

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
func findHost(buf *bufio.Reader) (err error, Host string, buffer bytes.Buffer) {

	// var buf bytes.Buffer
	// tee := io.TeeReader(conn, &buf)
	err = errors.New("Host header not found")
	Host = ""

	for {
		// TODO: Set readtimeout accordingly
		// Need to make it configuratble as well
		// TODO: Add maximum header size, max number of headers too, to avoid DOS attacks
		// will listen for message to process ending in newline (\n)
		var message string
		message, err = buf.ReadString(NextLine)
		// Copy message to header
		buffer.Write([]byte(message))
		if err != nil {
			log.Println("Error", err)
			buffer.Reset()
			return
		}

		if message == Crlf {
			// log.Println("Request ")
			// buffer.Write([]byte(CRLF))
			// Request Headers ended
			return
		}

		if strings.HasPrefix(message, HostPrefix) {

			message = message[6:]
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

	originaReader := bufio.NewReader(conn)
	err, ServerName, preRequest := findHost(originaReader)
	// ServerName := "quick.server"
	// var err error

	log.Println("host found", ServerName)
	if err != nil {
		log.Println("Error occured while finding host", err)
		return
	}

	rwc, err := s.hostmap.NewStreamFor(ServerName)
	if err != nil {
		fmt.Printf("[server:publicListener] unable to open a client stream: %s\n", err)
		return
	}

	defer rwc.Close()
	crwc := common.NewCompressedStream(rwc)
	// Not doing in go routine...Can be improved
	_, err = crwc.Write(preRequest.Bytes())

	if err != nil {
		fmt.Printf("[Error while writting header response to tunnel")
	}

	go func() {
		wreitten, err := zerocopy.Transfer(crwc, originaReader)
		log.Println("copying data", wreitten, err)
	}()
	if _, err := zerocopy.Transfer(conn, crwc); err != nil {
		fmt.Printf("[server:publicListener] unable to open a client stream: %s\n", err)
		return
	}
}
