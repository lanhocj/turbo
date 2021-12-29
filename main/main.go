package main

import (
	"github.com/laamho/turbo/common"
	"github.com/laamho/turbo/main/cmd"
	"os"
)

var release bool

func init() {
	//gin.SetMode(gin.ReleaseMode)
}

func main() {
	common.Must(cmd.RunApplicationWithArgs(os.Args))
}
