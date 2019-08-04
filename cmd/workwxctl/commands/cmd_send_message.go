package commands

import (
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

	app := cfg.MakeWorkwxApp()

	recipient := workwx.Recipient{
		UserIDs:  toUsers,
		PartyIDs: toParties,
		TagIDs:   toTags,
		ChatID:   toChat,
	}
	err := app.SendTextMessage(&recipient, c.Args().Get(0), isSafe)

	return err
}
