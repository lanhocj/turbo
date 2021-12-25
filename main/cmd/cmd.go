package cmd

import "github.com/urfave/cli/v2"

func RunApplicationWithArgs(args []string) error {
	authors := []*cli.Author{
		&cli.Author{
			Name:  "Laam Ho",
			Email: "hello@laamho.com",
		},
	}

	commands := []*cli.Command{
		supportServeCommand(),
		supportAddSuperUser(),
	}

	app := &cli.App{
		Authors:  authors,
		HideHelp: true,
		Usage:    "A Proxy Management Tool",
		Commands: commands,
	}

	return app.Run(args)
}
