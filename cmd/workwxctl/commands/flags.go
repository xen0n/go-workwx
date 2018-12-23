package commands

import (
	"github.com/urfave/cli"
)

const (
	flagCorpID            = "corpid"
	flagCorpSecret        = "corpsecret"
	flagAgentID           = "agentid"
	flagQyapiHostOverride = "qyapi-host-override"
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
