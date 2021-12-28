package cmd

import (
	"github.com/laamho/turbo/app/service/http"
	"github.com/laamho/turbo/common"
	"github.com/laamho/turbo/common/config"
	"github.com/laamho/turbo/common/orm"
	"github.com/urfave/cli/v2"
	"log"
	"path/filepath"
)

func webServerCommand() *cli.Command {
	return &cli.Command{ // turbo run
		Name:  "run",
		Usage: "Start Turbo Web Service",
		Flags: []cli.Flag{
			configFlag,
		},

		Action: serveActionHandler,
	}
}

func serveActionHandler(c *cli.Context) error {
	conf := config.New()
	configFilePath := c.Path("config")

	switch ext := filepath.Ext(configFilePath); ext {
	case ".yml", ".yaml":
		break
	default:
		log.Fatalf("config file: %s does supported!", ext)
	}

	if err := conf.LoadFile(configFilePath); err != nil {
		log.Fatal(err)
	}

	orm.Initialize()

	common.Must(orm.AutoMigrate(&orm.User{}, &orm.Node{}, &orm.Token{}))

	log.Println("Database initialization & connected successfully")

	go func() {
		if err := l.Listen(); err != nil {
			panic(err)
		}
	}()

	return http.StartWebApplication()
}
