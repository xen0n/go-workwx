package commands

import (
	"fmt"

	"github.com/urfave/cli"
)

func cmdSendTextMessage(c *cli.Context) error {
	cfg := mustGetConfig(c)
	fmt.Printf("send-text-message: cfg=%+v\n", cfg)
	return nil
}
