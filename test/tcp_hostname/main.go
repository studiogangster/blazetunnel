package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"log"
	"net"
	"strings"

	"acln.ro/zerocopy"
)

func main() {

	l, err := net.Listen("tcp4", ":4040")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handle(c)
	}
}

func handle(conn net.Conn) {
	conn.Write([]byte("starting rtesponse"))

	err, host, headerStream := findHost(conn)

	conn.Write([]byte("starting rtesponse"))

	go zerocopy.Transfer (conn, conn)
	// conn.Write(buffer.Bytes())
	if err != nil {
		conn.Close()
		log.Println("Error:", err)
	} else {
		log.Println("Host:", host)

		requestHeader := headerStream.Bytes()
		conn.Write(requestHeader)
		// log.Println(string(srequestHeader))
	}
}

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

		if strings.HasPrefix(message, "Host:") {
			Host = message
			err = nil
			// Host header found in request header
			break

		}

	}

	return
}
