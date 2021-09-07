package commands

import (
	"fmt"
	"strconv"

	"github.com/61qt/go-workwx"

	"github.com/urfave/cli/v2"
)

func cmdDeptList(c *cli.Context) error {
	cfg := mustGetConfig(c)
	haveID := c.Args().Len() > 0
	deptID := int64(0)
	if haveID {
		var err error
		deptIDStr := c.Args().Get(0)
		deptID, err = strconv.ParseInt(deptIDStr, 10, 64)
		if err != nil {
			fmt.Printf("invalid department ID: %+v\n", err)
			return err
		}
	}

	app := cfg.MakeWorkwxApp()
	// TODO: failed requests currently panics
	var info []*workwx.DeptInfo
	var err error
	if deptID == 0 {
		info, err = app.ListAllDepts()
	} else {
		info, err = app.ListDepts(deptID)
	}

	if err != nil {
		fmt.Printf("error = %+v\n", err)
	} else {
		fmt.Printf("departments:\n\n")
		for _, dept := range info {
			fmt.Printf("    %+v\n", dept)
		}
	}

	return err
}
