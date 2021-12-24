package cmd

import (
	"fmt"
	"github.com/laamho/turbo/common/config"
	"github.com/laamho/turbo/common/http"
	"github.com/laamho/turbo/common/orm"
	"github.com/urfave/cli/v2"
	"log"
)

func supportServeCommand() *cli.Command {
	return &cli.Command{ // turbo run
		Name:  "run",
		Usage: "Start Turbo Web Service",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "config",
				Aliases:  []string{"c"},
				Usage:    "set config file `path`",
				Required: true,
			},
		},
		Action: serveActionHandler,
	}
}

func serveActionHandler(c *cli.Context) error {
	conf := config.New()

	if err := conf.LoadFile(c.String("config")); err != nil {
		log.Fatal(err)
	}

	orm.Init()

	fmt.Printf("DB:init:%v", orm.DB)

	return http.Start()
}
