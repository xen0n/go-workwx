package commands

import (
	"github.com/xen0n/go-workwx"
	"gopkg.in/urfave/cli.v2"
)

const (
	flagCorpID            = "corpid"
	flagCorpSecret        = "corpsecret"
	flagAgentID           = "agentid"
	flagQyapiHostOverride = "qyapi-host-override"

	flagSafe         = "safe"
	flagToUser       = "to-user"
	flagToUserShort  = "u"
	flagToParty      = "to-party"
	flagToPartyShort = "p"
	flagToTag        = "to-tag"
	flagToTagShort   = "t"
	flagToChat       = "to-chat"
	flagToChatShort  = "c"

	flagMediaType = "media-type"
)

type cliOptions struct {
	CorpID     string
	CorpSecret string
	AgentID    int64

	QYAPIHostOverride string
}

func mustGetConfig(c *cli.Context) *cliOptions {
	if !c.IsSet(flagCorpID) {
		panic("corpid must be set")
	}

	if !c.IsSet(flagCorpSecret) {
		panic("corpsecret must be set")
	}

	if !c.IsSet(flagAgentID) {
		panic("agentid must be set (for now; may later lift the restriction)")
	}

	return &cliOptions{
		CorpID:     c.String(flagCorpID),
		CorpSecret: c.String(flagCorpSecret),
		AgentID:    c.Int64(flagAgentID),

		QYAPIHostOverride: c.String(flagQyapiHostOverride),
	}
}

//
// impl cliOptions
//

func (c *cliOptions) makeWorkwxClient() *workwx.Workwx {
	if c.QYAPIHostOverride != "" {
		// wtf think of a way to change this
		return workwx.New(c.CorpID, workwx.WithQYAPIHost(c.QYAPIHostOverride))
	}
	return workwx.New(c.CorpID)
}

func (c *cliOptions) MakeWorkwxApp() *workwx.WorkwxApp {
	return c.makeWorkwxClient().WithApp(c.CorpSecret, c.AgentID)
}
