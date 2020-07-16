package auth

import (
	"blazetunnel/common"
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"time"

	"github.com/lucas-clemente/quic-go"
)

var newmsg = common.NewMessage

// Client contains the configuration
// of the client to connect
type Auth struct {
	username string
	password string
	server   string
	service  string
	port     int64
}

// NewClient is used to create a new client
func NewAuth(username string, password string, server string, service string, port int64) *Auth {
	return &Auth{
		username: username,
		password: password,
		server:   server,
		port:     port,
		service:  service,
	}
}

// Start starts the peer connection to the tunnel server
func (a *Auth) Start() error {

	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"quic-echo-example"},
	}
	session, err := quic.DialAddr(a.server, tlsConf, &quic.Config{
		MaxIdleTimeout: time.Duration(time.Minute * 1),
	})
	if err != nil {
		return err
	}
	defer session.CloseWithError(500, "Closing")

	ctlStream, err := session.OpenStreamSync(context.Background())
	if err != nil {
		return err
	}
	defer ctlStream.Close()

	authenticationCredemtials := a.username + ":" + a.password + ":" + a.service

	err = newmsg(common.CommandAuthClient, authenticationCredemtials).EncodeTo(ctlStream)
	if err != nil {
		return err
	}

	fmt.Println("Authenticating with blazerecon server")
	ctlStream.SetReadDeadline(time.Now().Add(time.Second * 5))

	m, err := newmsg("", "").DecodeFrom(ctlStream)
	if err != nil {
		return err
	}

	return saveToken(m.Context, *a)

}

func saveToken(authtoken string, auth Auth) error {

	// f, err := os.Create(".blazetoken")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }
	// log.Println("Token:", authtoken)
	// _, err = f.WriteString(authtoken)
	// if err != nil {
	// 	fmt.Println(err)
	// 	f.Close()
	// 	return err
	// }

	// err = f.Close()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }

	if authtoken == "" {
		fmt.Println("Invalid Authtoken")

	} else {
		fmt.Println("Authtoken saved in .blazetoken", authtoken)

	}
	fmt.Printf("Use:\n\tgo run main.go client --tunnel %s --local localhost:%d\n\tto connect to the internet\n", auth.server, auth.port)

	if authtoken == "" {

		return errors.New("Invalid Credentials")

	}
	return nil
}
