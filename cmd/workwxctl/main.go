package main

import (
	"os"

	"github.com/xen0n/go-workwx/cmd/workwxctl/commands"
)

func main() {
	app := commands.InitApp()
	// ignore errors
	_ = app.Run(os.Args)
}
