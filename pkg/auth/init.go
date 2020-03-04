package auth

import (
	"errors"

	"github.com/urfave/cli/v2"
)

// Init function initializes the client command
// commandline functionality and retursn the cli.Command
func Init() *cli.Command {

	return &(cli.Command{
		Name: "auth",

		Usage:  "Authenticate with server to reserve a subdomain for your service",
		Action: createClient,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "server",
				Usage:    "Auth server address",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "username",
				Usage:    "Username",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "password",
				Usage:    "Password",
				Required: true,
			},
		},
	})
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

	return errors.New("Authentication not implemented yet!")
	// return NewClient(tunnel, local, ctx.Uint("i")).Start()
}
