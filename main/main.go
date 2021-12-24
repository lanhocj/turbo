package main

import (
	"github.com/laamho/turbo/main/cmd"
	"os"
)

func main() {
	// Run the application.
	if err := cmd.RunApplicationWithArgs(os.Args); err != nil {
		os.Exit(0)
	}
}
