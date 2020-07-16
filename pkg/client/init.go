package client

import (
	"errors"
	"log"
	"time"

	"github.com/urfave/cli/v2"
)

const reconnectDelay time.Duration = 4

// Init function initializes the client command
// commandline functionality and retursn the cli.Command
func Init() *cli.Command {

	flags := []cli.Flag{
		&cli.StringFlag{
			Name:     "tunnel",
			Usage:    "Remote public tunnel address to connect to",
			Required: true,
			EnvVars:  []string{"tunnel"},
		},

		&cli.Int64Flag{
			Name:     "tunnelport",
			Usage:    "Remote public tunnel's port",
			Required: false,
			Value:    2723,
		},
		&cli.StringFlag{
			Name:     "local",
			Usage:    "Local TCP server to proxy the connections to",
			Required: true,
		},
		&cli.StringFlag{
			Name:  "token",
			Usage: "auth-token ",
			// FilePath: ".blazetoken",
			Required: true,
			EnvVars:  []string{"token"},
		},
		&cli.UintFlag{
			Name:  "i,idle-timeout",
			Usage: "Idle timeout for the quic sessions (in seconds)",
			Value: 1800,
		},
	}

	return &(cli.Command{
		Name: "client",

		Usage:  "Run a client instance",
		Action: createSelfConnectingClient,
		Flags:  flags,
	})
}

func createSelfConnectingClient(ctx *cli.Context) error {

	log.Println("createSelfConnectingClient")
	for {
		err := createClient(ctx)
		log.Println("[DEBUG]", "Blazetunnel client was closed ", err)
		log.Println("[DEBUG]", "Creating new cient in ", reconnectDelay, "seconds")
		<-time.After(reconnectDelay * time.Second)

	}

}

func createClient(ctx *cli.Context) error {

	tunnel := ctx.String("tunnel")
	if tunnel == "" {
		return errors.New("Tunnel address cannot be empty")
	}

	local := ctx.String("local")
	if local == "" {
		return errors.New("Local address cannot be empty")
	}

	tunnelport := ctx.String("tunnelport")

	token := ctx.String("token")
	if token == "" {
		return errors.New("Token can not be empty")
	}

	tunnel = tunnel + ":" + tunnelport
	log.Println("Connecting to", tunnel, local)
	return NewClient(tunnel, local, ctx.Uint("i"), token).Start()
}
