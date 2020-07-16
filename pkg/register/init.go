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
				EnvVars:  []string{"server"},
			},
			&cli.StringFlag{
				Name:     "Username",
				Aliases:  []string{"u"},
				Usage:    "Username",
				Required: true,
				EnvVars:  []string{"APPNAME"},
			},
			&cli.StringFlag{
				Name:     "Password",
				Aliases:  []string{"p"},
				Usage:    "Password",
				Required: true,
				EnvVars:  []string{"PASSWORD"},
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
	Username := ctx.String("Username")
	if Username == "" {
		return errors.New("Username can not be empty")
	}

	Password := ctx.String("Password")
	if Password == "" {
		return errors.New("Password can not be empty")
	}

	port := ctx.Int("port")
	if port == 0 {
		return errors.New("Port cannot be empty")
	}

	return NewApp(Username, Password, server, port).Start()
}
