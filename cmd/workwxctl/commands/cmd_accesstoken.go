package commands

import (
	"fmt"

	"github.com/urfave/cli"
)

func cmdAccesstoken(c *cli.Context) error {
	cfg := mustGetConfig(c)
	fmt.Printf("accesstoken: cfg=%+v\n", cfg)
	return nil
}
