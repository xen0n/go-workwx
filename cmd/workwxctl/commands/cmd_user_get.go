package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func cmdUserGet(c *cli.Context) error {
	cfg := mustGetConfig(c)
	userid := c.Args().Get(0)

	app := cfg.MakeWorkwxApp()
	// TODO: failed requests currently panics
	info, err := app.GetUser(userid)

	if err != nil {
		fmt.Printf("error = %+v\n", err)
	} else {
		fmt.Printf("user = %+v\n", info)
	}

	return err
}
