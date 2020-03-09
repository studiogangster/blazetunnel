package auth

import (
	"errors"

	"github.com/urfave/cli/v2"
)

// Init function initializes the auth command
// commandline functionality and return the cli.Command
func Init() *cli.Command {

	return &(cli.Command{
		Name:   "auth",
		Usage:  "Authenticate with server to reserve a subdomain for your service",
		Action: createAuth,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "server",
				Aliases:  []string{"s"},
				Usage:    "Auth server's address",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "username",
				Aliases:  []string{"u"},
				Usage:    "Username",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "password",
				Aliases:  []string{"p"},
				Usage:    "Password",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "service",
				Aliases:  []string{"S"},
				Usage:    "Service name",
				Required: true,
			},

			&cli.Int64Flag{
				Name:     "port",
				Aliases:  []string{"P"},
				Usage:    "Local Server's port ",
				Required: true,
			},
		},
	})
}

func createAuth(ctx *cli.Context) error {
	server := ctx.String("server")
	if server == "" {
		return errors.New("server cannot be empty")
	}

	username := ctx.String("username")
	if username == "" {
		return errors.New("username cannot be empty")
	}

	password := ctx.String("password")
	if password == "" {
		return errors.New("password cannot be empty")
	}

	service := ctx.String("service")
	if service == "" {
		return errors.New("service cannot be empty")
	}

	port := ctx.Int64("port")
	if port == 0 {
		return errors.New("Port cannot be empty")
	}

	return NewAuth(username, password, server, service, port).Start()
}
