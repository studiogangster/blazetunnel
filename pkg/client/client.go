package client

import (
	"blazetunnel/common"
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"os"
	"time"

	"acln.ro/zerocopy"
	"github.com/lucas-clemente/quic-go"
)

var newmsg = common.NewMessage

// Client contains the configuration
// of the client to connect
type Client struct {
	tunnel      string
	local       string
	idleTimeout uint
}

// NewClient is used to create a new client
func NewClient(tunnel, local string, idleTimeout uint) *Client {
	return &Client{
		tunnel:      tunnel,
		local:       local,
		idleTimeout: idleTimeout,
	}
}

func GetServiceName() string {

	return os.Getenv("SERVICE_NAME")

}

var l = fmt.Println

// Start starts the peer connection to the tunnel server
func (c *Client) Start() error {

	// c.tunnel = "server:2723"

	fmt.Println(c.tunnel)
	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"quic-echo-example"},
	}
	session, err := quic.DialAddr(c.tunnel, tlsConf, &quic.Config{
		IdleTimeout: time.Second * time.Duration(c.idleTimeout),
	})
	if err != nil {
		return err
	}
	defer session.Close()

	ctlStream, err := session.OpenStreamSync(context.Background())
	if err != nil {
		return err
	}
	defer ctlStream.Close()

	err = newmsg(common.CommandNewClient, GetServiceName()).EncodeTo(ctlStream)
	if err != nil {
		return err
	}

	l("Opened stream")
	m, err := newmsg("", "").DecodeFrom(ctlStream)
	if err != nil {
		return err
	}

	l("Read msgpack message")

	fmt.Printf("Message received: %s(%s)\n", m.Command, m.Context)
	l("accepting connection")
	go c.handleCtlStream(ctlStream)
	i := 0
	for {
		stream, err := session.AcceptStream(context.Background())
		if err != nil {
			fmt.Printf("[client:tunnelConnection] unable to open a stream: %s\n", err)
			return err
		}
		i++
		l("opened stream: ", i)
		go c.handleStream(common.NewCompressedStream(stream))
	}

}

func (c *Client) handleStream(stream quic.Stream) {
	defer func() {
		fmt.Println("Closing")
		stream.Close()
	}()

	dest, err := net.Dial("tcp", c.local)
	if err != nil {
		fmt.Printf("[client:localConnection] unable to open local connection: %s\n", err)
		return
	}
	defer dest.Close()

	go zerocopy.Transfer(dest, stream)
	if _, err := zerocopy.Transfer(stream, dest); err != nil {
		fmt.Printf("[client:localConnection] unable to open local connection: %s\n", err)
		return
	}
}

func (c *Client) handleCtlStream(ctlStream quic.Stream) {
	err := newmsg(common.CommandPingPeer, "").EncodeTo(ctlStream)
	if err != nil {
		fmt.Printf("[server:pong] unable to decode from msgpack: %s\n", err)
		return
	}
	ctlStream.SetReadDeadline(time.Now().Add(time.Minute))

	getOut := false
	for !getOut {
		m, err := newmsg("", "").DecodeFrom(ctlStream)
		if err != nil {
			fmt.Printf("[client:ping] unable to decode from msgpack: %s\n", err)
			return
		}
		<-time.After(3 * time.Second)
		switch m.Command {
		case common.CommandPongPeer:
			fmt.Printf("[client:message] Got pong from %s\n", c.tunnel)
			err = newmsg(common.CommandPingPeer, "").EncodeTo(ctlStream)
			if err != nil {
				fmt.Printf("[client:ping] unable to encode to msgpack: %s\n", err)
				getOut = true
				break
			}
			ctlStream.SetReadDeadline(time.Now().Add(time.Minute))
		}

	}
}
