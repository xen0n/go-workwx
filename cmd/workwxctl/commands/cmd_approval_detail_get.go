package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func cmdApprovalDetailGet(c *cli.Context) error {
	cfg := mustGetConfig(c)
	spNo := c.Args().Get(0)

	app := cfg.MakeWorkwxApp()
	info, err := app.GetOAApprovalDetail(spNo)

	if err != nil {
		fmt.Printf("error = %+v\n", err)
	} else {
		fmt.Printf("approval detail = %+v\n", info)
	}

	return err
}
