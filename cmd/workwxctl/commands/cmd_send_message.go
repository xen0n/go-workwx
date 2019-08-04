package commands

import (
	"fmt"

	"github.com/xen0n/go-workwx"
	"gopkg.in/urfave/cli.v2"
)

func cmdSendMessage(c *cli.Context) error {
	cfg := mustGetConfig(c)
	isSafe := c.Bool(flagSafe)
	toUsers := c.StringSlice(flagToUser)
	toParties := c.StringSlice(flagToParty)
	toTags := c.StringSlice(flagToTag)
	toChat := c.String(flagToChat)
	content := c.Args().Get(0)
	msgtype := c.String(flagMessageType)

	app := cfg.MakeWorkwxApp()

	recipient := workwx.Recipient{
		UserIDs:  toUsers,
		PartyIDs: toParties,
		TagIDs:   toTags,
		ChatID:   toChat,
	}

	if msgtype == "" {
		// default to text
		msgtype = "text"
	}

	var err error
	switch msgtype {
	case "text":
		err = app.SendTextMessage(&recipient, content, isSafe)
	default:
		fmt.Printf("unrecognized message type: %s\n", msgtype)
		panic("unrecognized message type")
	}

	return err
}
