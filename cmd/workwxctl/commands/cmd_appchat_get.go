package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func cmdAppchatGet(c *cli.Context) error {
	cfg := mustGetConfig(c)
	chatid := c.Args().Get(0)

	app := cfg.MakeWorkwxApp()
	// TODO: failed requests currently panics
	info, err := app.GetAppchat(chatid)

	if err != nil {
		fmt.Printf("error = %+v\n", err)
	} else {
		fmt.Printf("appchat = %+v\n", info)
	}

	return err
}
