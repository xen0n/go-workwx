package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/xen0n/go-workwx"
)

func cmdAppchatCreate(c *cli.Context) error {
	cfg := mustGetConfig(c)
	chatID := c.String(flagChatID)
	name := c.String(flagName)
	ownerID := c.String(flagOwner)
	userIDs := c.StringSlice(flagUser)

	app := cfg.MakeWorkwxApp()
	// TODO: failed requests currently panics
	req := workwx.ChatInfo{
		ChatID:        chatID,
		Name:          name,
		OwnerUserID:   ownerID,
		MemberUserIDs: userIDs,
	}
	fmt.Printf("about to create appchat %+v\n", req)

	newChatID, err := app.CreateAppchat(&req)
	if err != nil {
		fmt.Printf("failed to create appchat: error = %+v\n", err)
	} else {
		fmt.Printf("created appchat: chatid = \"%s\"\n", newChatID)
	}

	return err
}
