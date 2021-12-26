package cmd

import (
	"bufio"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/laamho/turbo/common/util"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
)

func secretKeyGeneratorCommand() *cli.Command {
	return &cli.Command{
		Name:   "secret",
		Action: secretKeyGeneratorActionHandler,
	}
}

func secretKeyGeneratorActionHandler(c *cli.Context) error {
	render := bufio.NewReader(os.Stdin)
	var password, confirm string

	for !util.MustNotEmpty(password) {
		if util.Empty(password) {
		password:
			fmt.Printf("Email: ")
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
			fmt.Printf("\nConfirm delete? (y|yes) or (n|no): ")
			confirm, _ = render.ReadString('\n')
			confirm = strings.Replace(confirm, "\n", "", -1)
			goto c
		}
	}

	h := hmac.New(sha256.New, nil)
	p := []byte(password)
	h.Write(p)
	enc := base64.StdEncoding.EncodeToString(h.Sum(nil))

	fmt.Printf("\033[1;38m%s\033[0m\n", enc)

	return nil
}
