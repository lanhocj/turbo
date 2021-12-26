package cmd

import (
	"github.com/urfave/cli/v2"
)

var configFlag = &cli.StringFlag{
	Name:     "config",
	Aliases:  []string{"c"},
	Usage:    "Load config through `FILE`",
	Required: true,
}

func RunApplicationWithArgs(args []string) error {
	authors := []*cli.Author{
		&cli.Author{
			Name:  "Laam Ho",
			Email: "hello@laamho.com",
		},
	}

	commands := []*cli.Command{
		webServerCommand(),
		authControlCommands(),
		secretKeyGeneratorCommand(),
	}

	app := &cli.App{
		Authors:  authors,
		HideHelp: true,
		Usage:    "A Proxy Management Tool",
		Commands: commands,
	}

	return app.Run(args)
}
