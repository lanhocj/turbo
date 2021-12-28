package cmd

import (
	"bufio"
	"fmt"
	"github.com/laamho/turbo/common"
	"github.com/laamho/turbo/common/lpc"
	"github.com/laamho/turbo/common/orm"
	"github.com/laamho/turbo/common/util"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

// Impl admin command, including add or remove
// example:
//    turbo admin add
//    turbo admin remove

var l = &lpc.RPC{}

func authControlCommands() *cli.Command {
	return &cli.Command{
		Name:  "admin",
		Usage: "Administrator account management",
		Before: func(context *cli.Context) error {
			return l.Dial()
		},
		Subcommands: []*cli.Command{
			&cli.Command{
				Name:   "add",
				Usage:  "Add a account to administrator group",
				Action: addSuperUserActionHandler,
			},
			&cli.Command{
				Name:   "remove",
				Usage:  "Remove a account for administrator group",
				Action: removeSuperActionHandler,
			},
		},
	}
}

func removeSuperActionHandler(_ *cli.Context) error {
	render := bufio.NewReader(os.Stdin)
	var email, confirm string

	for !util.MustNotEmpty(email) {
		if util.Empty(email) {
		email:
			fmt.Printf("Email: ")
			email, _ = render.ReadString('\n')
			email = strings.Replace(email, "\n", "", -1)
			if util.Empty(email) {
				goto email
			}
		}

	c:
		switch confirm {
		case "y", "yes":
			break
		case "n", "no":
			os.Exit(0)
		default:
			fmt.Printf("\nConfirm delete? (y|yes) or (n|no): ")
			confirm, _ = render.ReadString('\n')
			confirm = strings.Replace(confirm, "\n", "", -1)
			goto c
		}
	}

	reply, err := l.Send(lpc.RemoveUserService, func() interface{} {
		return lpc.UserLPCServiceObject{
			Email: email,
		}
	})

	if err != nil {
		return err
	}

	fmt.Println(reply)

	return nil
}

func addSuperUserActionHandler(_ *cli.Context) error {
	render := bufio.NewReader(os.Stdin)
	var email, password, confirm string

	for !util.MustNotEmpty(email, password) {
		if util.Empty(email) {
		email:
			fmt.Printf("Email: ")
			email, _ = render.ReadString('\n')
			email = strings.Replace(email, "\n", "", -1)
			if util.Empty(email) {
				goto email
			}
		}

		if util.Empty(password) {
		password:
			fmt.Printf("Password: ")
			password, _ = render.ReadString('\n')
			password = strings.Replace(password, "\n", "", -1)
			if util.Empty(password) {
				goto password
			}
		}
	c:
		switch confirm {
		case "y", "yes":
			break
		case "n", "no":
			os.Exit(0)
		default:
			fmt.Printf("\nConfirm create? (y|yes) or (n|no): ")
			confirm, _ = render.ReadString('\n')
			confirm = strings.Replace(confirm, "\n", "", -1)
			goto c
		}
	}

	wait := make(chan string)
	go func() {
		reply, err := l.Send(lpc.CreateUserService, func() interface{} {
			return lpc.UserLPCServiceObject{
				Email:    email,
				Password: password,
				Role:     orm.LEVEL_USER_ADMIN,
			}
		})

		wait <- reply
		common.Silent(err)
	}()

	r := <-wait
	fmt.Printf(r)

	return nil
}
