package server

import (
	"blazetunnel/common"
	"errors"

	"github.com/urfave/cli/v2"
)

const (
	errDomain = "a single allocation domain is required for the server to function"
)

// Init initializes the commandline option flag for
// server mode
func Init() *cli.Command {
	return &(cli.Command{
		Name: "server",

		Usage:  "Run a server instance",
		Action: createServer,

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "domain",
				Usage:    "Domain at which the new host allocations had to be done",
				Required: true,
			},
			&cli.UintFlag{
				Name:    "idle-timeout",
				Aliases: []string{"i"},
				Usage:   "Idle timeout for the quic sessions (in seconds)",
				Value:   1800,
			},

			&cli.StringFlag{
				Name:        "secret",
				Aliases:     []string{"s"},
				DefaultText: "Secret Key",
				Usage:       "Secret key to be used for encryption",
				Value:       "01234567890123456789012345678912",
			},
		},
	})
}

func createServer(ctx *cli.Context) error {
	domain := ctx.String("domain")
	if domain == "" {
		return errors.New(errDomain)
	}

	secret := ctx.String("secret")
	if secret == "" {
		return errors.New(errDomain)
	}

	common.SetSecretKey(secret)
	NewServer(domain, ctx.Uint("i")).Start()
	return nil
}
