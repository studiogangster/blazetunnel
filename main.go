package main

import (
	"blazetunnel/pkg/client"
	"blazetunnel/pkg/server"
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "blazetunnel"
	app.Version = "v0.0.1"
	app.Usage = "Expose your local application ports to the internet"
	app.Copyright = "Akilan Elango 2019"
	app.Commands = []cli.Command{
		server.Init(),
		client.Init(),
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(2)
	}
}
