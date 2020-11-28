package commands

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/urfave/cli/v2"
)

func cmdUserListByDept(c *cli.Context) error {
	cfg := mustGetConfig(c)
	haveID := c.Args().Len() > 0
	if !haveID {
		return errors.New("no department ID given")
	}

	deptIDStr := c.Args().Get(0)
	deptID, err := strconv.ParseInt(deptIDStr, 10, 64)
	if err != nil {
		fmt.Printf("invalid department ID: %+v\n", err)
		return err
	}

	app := cfg.MakeWorkwxApp()
	// TODO: failed requests currently panics
	// TODO: fetchChild
	info, err := app.ListUsersByDeptID(deptID, false)

	if err != nil {
		fmt.Printf("error = %+v\n", err)
	} else {
		fmt.Printf("users in dept ID %d:\n\n", deptID)
		for _, user := range info {
			fmt.Printf("    %+v\n", user)
		}
	}

	return err
}
