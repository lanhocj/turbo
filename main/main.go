package main

import (
	"github.com/laamho/turbo/common"
	"github.com/laamho/turbo/main/cmd"
	"os"
)

func init() {
	
}

func main() {
	common.Must(cmd.RunApplicationWithArgs(os.Args))
}
