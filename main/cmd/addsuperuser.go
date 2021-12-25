package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/laamho/turbo/common/config"
	"github.com/laamho/turbo/common/orm"
	"github.com/laamho/turbo/common/util"
	"github.com/urfave/cli/v2"
)

func supportAddSuperUser() *cli.Command {
	return &cli.Command{ // turbo run
		Name:  "addsuperuser",
		Usage: "Create a Super User",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "config",
				Aliases:  []string{"c"},
				Usage:    "set config file `path`",
				Required: true,
			},
		},
		Action: addSuperUserActionHandler,
	}
}

func addSuperUserActionHandler(c *cli.Context) error {
	render := bufio.NewReader(os.Stdin)
	var email string
	var password = "helo"

	for {
		fmt.Printf("Email: ")
		email, _ = render.ReadString('\n')
		email = strings.Replace(email, "\n", "", -1)

		if email != "" {
			fmt.Printf("\b\rEmail: \033[1m%s\033[0m", email)
			break
		}
	}

	os.Exit(0)

	conf := config.New()

	if err := conf.LoadFile(c.String("config")); err != nil {
		log.Fatal(err)
	}

	orm.Init()

	var user orm.User

	if orm.DB.First(&user, "email=?", email); user.ID != 0 {
		log.Printf("Email %s Already exists", email)
		return nil
	}

	user = orm.User{
		Email:    email,
		Password: util.Hash(email, password),
		Role:     orm.LEVEL_USER_ADMIN,
	}

	orm.DB.Create(&user)

	return nil
}
