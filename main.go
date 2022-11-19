package main

import (
	"log"
	"os"

	"sport-programming-cli/commands"

	"github.com/urfave/cli/v2"
)

const Version = "0.2.2"

func main() {
	app := &cli.App{
		Name:           "sport-programming-cli",
		Usage:          "Simple CLI Application for copying from github.com/0wol/sport-programming-library",
		Version:        Version,
		DefaultCommand: "get",
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
			},
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "Get full directory",
				Action:  commands.List,
			},
		},
		EnableBashCompletion: true,
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "0w0l Team",
				Email: "github.com/0wol",
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
