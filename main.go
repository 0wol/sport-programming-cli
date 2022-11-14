package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"sport-programming-cli/commands"
)

func main() {
	app := &cli.App{
		Name:                 "sport-programming-cli",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:    "update",
				Aliases: []string{"u"},
				Usage:   "Update the library",
				Action:  commands.Update,
			},
			{
				Name:    "get",
				Aliases: []string{"g"},
				Usage:   "Get function/file by name",
				Action:  commands.Get,
				// BashComplete: ,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
