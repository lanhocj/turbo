package main

import (
	"github.com/gin-gonic/gin"
	"github.com/laamho/turbo/common"
	"github.com/laamho/turbo/main/cmd"
	"os"
)

var release string

func init() {
	if release == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
}

func main() {
	common.Must(cmd.RunApplicationWithArgs(os.Args))
}
