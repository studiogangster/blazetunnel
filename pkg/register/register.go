package register

import (
	"blazetunnel/common"
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/lucas-clemente/quic-go"
)

var newmsg = common.NewMessage

// Client contains the configuration
// of the client to connect
type App struct {
	appname  string
	password string
	server   string
	port     int
}

// NewClient is used to create a new client
func NewApp(appname string, password string, server string, port int) *App {
	return &App{
		appname:  appname,
		password: password,
		server:   server,
		port:     port,
	}
}

// Start starts the peer connection to the tunnel server
func (a *App) Start() error {

	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"quic-echo-example"},
	}

	log.Println("Connecting to ", a.server+":"+strconv.Itoa(a.port))
	session, err := quic.DialAddr(a.server+":"+strconv.Itoa(a.port), tlsConf, &quic.Config{
		IdleTimeout: time.Duration(time.Minute * 1),
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

	registerCredentials := a.appname + ":" + a.password

	err = newmsg(common.CommandRegisterClient, registerCredentials).EncodeTo(ctlStream)
	if err != nil {
		return err
	}

	fmt.Println("Registering with blazerecon server")
	ctlStream.SetReadDeadline(time.Now().Add(time.Second * 5))

	m, err := newmsg("", "").DecodeFrom(ctlStream)
	if err != nil {
		return err
	}

	log.Println(m.Context)
	return nil

}
