package main

import (
	"github.com/laamho/turbo/common"
	"github.com/laamho/turbo/main/cmd"
	"os"
)

func main() {
	common.Must(cmd.RunApplicationWithArgs(os.Args))
}
