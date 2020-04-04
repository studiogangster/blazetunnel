package main

import (
	"blazetunnel/pkg/auth"
	"blazetunnel/pkg/client"
	"blazetunnel/pkg/register"

	"blazetunnel/pkg/server"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name:                 "blazetunnel",
		Version:              "1.0.5",
		Usage:                "Expose your local application / docker container to the internet",
		EnableBashCompletion: true,

		Commands: []*cli.Command{
			server.Init(),
			client.Init(),
			auth.Init(),
			register.Init(),
		},
	}

	// app := cli.NewApp()
	// app.Name = "blazetunnel"
	// app.Version = "v0.0.1"
	// app.Usage = "Expose your local applications ports to the internet"
	// app.Copyright = "Akilan Elango 2019"
	// app.Commands = []cli.Command{
	server.Init()
	// 	client.Init(),
	// }
	if err := app.Run(os.Args); err != nil {
		log.Println(err)
		os.Exit(2)
	}
}
