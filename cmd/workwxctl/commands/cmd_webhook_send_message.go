package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/xen0n/go-workwx/v2"
)

func cmdWebhookSendMessage(c *cli.Context) error {
	cfg := mustGetConfig(c)
	mentionUsers := c.StringSlice(flagMentionUser)
	mentionMobiles := c.StringSlice(flagMentionMobile)
	content := c.Args().Get(0)
	msgtype := c.String(flagMessageType)

	wh := cfg.makeWebhookClient()

	if msgtype == "" {
		// default to text
		msgtype = "text"
	}

	mentions := workwx.Mentions{
		UserIDs: mentionUsers,
		Mobiles: mentionMobiles,
	}

	var err error
	switch msgtype {
	case "text":
		err = wh.SendTextMessage(content, &mentions)

	case "markdown":
		if len(mentionUsers) > 0 || len(mentionMobiles) > 0 {
			panic("cannot specify mention parameters when sending markdown messages with webhook")
		}

		err = wh.SendMarkdownMessage(content)

	default:
		fmt.Printf("unrecognized message type: %s\n", msgtype)
		panic("unrecognized message type")
	}

	return err
}
