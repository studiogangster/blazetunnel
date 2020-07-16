package client

import (
	"blazetunnel/common"
	"context"
	"crypto/tls"
	"errors"
	"log"
	"net"
	"sync"
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
	token       string
	sync.RWMutex
}

// NewClient is used to create a new client
func NewClient(tunnel, local string, idleTimeout uint, token string) *Client {
	return &Client{
		tunnel:      tunnel,
		local:       local,
		idleTimeout: idleTimeout,
		token:       token,
	}
}

// Start starts the peer connection to the tunnel server
func (c *Client) Start() error {

	log.Println("[DEBUG]", "Starting", "Start()")

	defer func() {
		log.Println("[DEBUG]", "Closing", "Start()")

	}()

	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"quic-echo-example"},
	}

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	session, err := quic.DialAddrContext(ctx, c.tunnel, tlsConf, &quic.Config{
		MaxIdleTimeout: time.Second * time.Duration(c.idleTimeout),
	})

	if err != nil {
		return err
	}

	defer func() {
		log.Println("Defer", "session.close()")
		// session.ConnectionState()
		if ctx.Err() == nil {

			log.Println("[DEBUG]", "Session not closed yet", "Attempt to close")
			session.CloseWithError(500, "Closing")
			log.Println("[DEBUG]", "Session Closed")
		} else {
			log.Println("[DEBUG]", "Session already closed")
		}

	}()

	ctlStream, err := session.OpenStreamSync(ctx)
	if err != nil {
		return err
	}

	defer func() {
		log.Println("[DEBUG]", "Attempt to close controlStream")

		if ctx.Err() == nil {

			log.Println("[DEBUG]", "controlStream not closed yet", "Attempt to close")
			ctlStream.Close()
			log.Println("[DEBUG]", "controlStream Closed")
		} else {
			log.Println("[DEBUG]", "controlStream already closed")
		}
		ctlStream.Close()
	}()

	err = newmsg(common.CommandNewClient, c.token).EncodeTo(ctlStream)
	if err != nil {
		return err
	}

	m, err := newmsg("", "").DecodeFrom(ctlStream)
	if err != nil {
		return err
	}

	log.Printf("Message received: %s(%s)\n", m.Command, m.Context)

	// ctx := context.Background()
	// Create a new context, with its cancellation function
	// from the original context

	var wg sync.WaitGroup
	wg.Add(1)
	go c.handleCtlStream(ctlStream, cancel, &wg)
	// i := 0
	for {

		stream, err := session.AcceptStream(ctx)
		if err != nil {
			log.Printf("[client:tunnelConnection] unable to open a stream: %s\n", err)
			break

		}
		wg.Add(1)
		go c.handleStream(ctx, common.NewCompressedStream(stream), &wg)
	}

	log.Println("Waiting for all goroutines to finish that were started via control stream")
	defer wg.Wait()
	return errors.New("[DEBUG]:" + "Clodes complete quic session and clearing left over garbage")

}

func (c *Client) handleStream(ctx context.Context, stream quic.Stream, _wg *sync.WaitGroup) {

	defer _wg.Done()
	log.Println("[DEBUG]", "handleStream(")
	defer func() {
		log.Println("[DEBUG]", "~handleStream()")
	}()

	defer func() {

		stream.Close()
	}()

	localhost := ""
	c.RLock()
	localhost = c.local
	c.RUnlock()
	dest, err := net.Dial("tcp", localhost)

	if err != nil {
		log.Printf("[client:localConnection] unable to open local connection: %s\n", err)
		return
	}
	defer dest.Close()

	// Listen for context closure, and close the goroutine

	completion := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		wg.Done()
		log.Println("[DEBUG]", "contextobserver()")

		defer func() {
			log.Println("[DEBUG]", "~contextobserver()")
		}()
		select {
		case <-ctx.Done():
			log.Println("[DEBUG]", "contextobserver('CONTEXT_DONE')")

			dest.Close()
			stream.Close()

			break
		case <-completion:
			log.Println("[DEBUG]", "contextobserver('COMPLETION')")
			return

		}
	}()

	defer wg.Wait()

	go zerocopy.Transfer(dest, stream)
	if _, err := zerocopy.Transfer(stream, dest); err != nil {
		log.Printf("[client:localConnection] unable to open local connection: %s\n", err)

	}

	close(completion)
}

func (c *Client) handleCtlStream(ctlStream quic.Stream, cancel context.CancelFunc, wg *sync.WaitGroup) {

	defer wg.Done()

	log.Println("Starting handleCtlStream")
	defer func() {
		log.Println("Closing handleCtlStream")
	}()

	ctlStream.SetReadDeadline(time.Now().Add(time.Second * 5))
	err := newmsg(common.CommandPingPeer, "").EncodeTo(ctlStream)
	if err != nil {
		log.Printf("[server:pong] unable to decode from msgpack: %s\n", err)
		cancel()
		return
	}
	ctlStream.SetReadDeadline(time.Time{})

	for {
		// default:
		// log.Println("PingPong Timeout", "Closing Connection")
		// <-time.After(5 * time.Second)

		ctlStream.SetReadDeadline(time.Now().Add(time.Second * 5))
		m, err := newmsg("", "").DecodeFrom(ctlStream)

		if err != nil {
			// Coudn't read message from control stream | Probably due to timeout | Close the session and reinitiate
			log.Println("Error reading timeout")
			cancel()
			return
		}
		ctlStream.SetReadDeadline(time.Time{})

		// Check which operation was message about
		switch m.Command {
		case common.CommandPongPeer:
			log.Printf("[client:message] Got pong from %s\n", c.tunnel)
			<-time.After(3 * time.Second)
			ctlStream.SetWriteDeadline(time.Now().Add(time.Second * 5))

			err = newmsg(common.CommandPingPeer, "").EncodeTo(ctlStream)

			if err != nil {
				log.Printf("[client:ping] unable to encode to msgpack: %s\n", err)
				cancel()
				return

			}
			ctlStream.SetWriteDeadline(time.Time{})
			break

		default:
			log.Println("[client:message] Unknwon type of message recieved")
			break

		}

	}

}
