package register

import (
	"errors"

	"github.com/urfave/cli/v2"
)

// Init function initializes the auth command
// commandline functionality and return the cli.Command
func Init() *cli.Command {

	return &(cli.Command{
		Name:   "register",
		Usage:  "Register with server to create a new account",
		Action: registerApp,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "server",
				Aliases:  []string{"s"},
				Usage:    "Auth server's address",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "Appname",
				Aliases:  []string{"a"},
				Usage:    "Application's name",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "Password",
				Aliases:  []string{"p"},
				Usage:    "Password",
				Required: true,
			},

			&cli.Int64Flag{
				Name:     "port",
				Aliases:  []string{"P"},
				Usage:    "Local Server's port ",
				Value:    2723,
				Required: false,
			},
		},
	})
}

func registerApp(ctx *cli.Context) error {

	server := ctx.String("server")
	if server == "" {
		return errors.New("server cannot be empty")
	}
	Appname := ctx.String("Appname")
	if Appname == "" {
		return errors.New("Appname can not be empty")
	}

	Password := ctx.String("Password")
	if Password == "" {
		return errors.New("Password can not be empty")
	}

	port := ctx.Int("port")
	if port == 0 {
		return errors.New("Port cannot be empty")
	}

	return NewApp(Appname, Password, server, port).Start()
}
