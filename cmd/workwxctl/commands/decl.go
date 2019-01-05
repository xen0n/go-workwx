package commands

import (
	"github.com/urfave/cli"
)

// InitApp defines the workwxctl CLI.
func InitApp() *cli.App {
	return &cli.App{
		Name:  "workwxctl",
		Usage: "企业微信命令行客户端 powered by go-workwx",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    flagCorpID,
				Usage:   "使用 `CORPID` 作为企业 ID",
				EnvVars: []string{"WORKWXCTL_CORPID"},
			},
			&cli.StringFlag{
				Name:    flagCorpSecret,
				Usage:   "使用 `SECRET` 作为应用凭证密钥",
				EnvVars: []string{"WORKWXCTL_CORPSECRET"},
			},
			&cli.Int64Flag{
				Name:    flagAgentID,
				Usage:   "使用 `AGENTID` 作为企业应用 ID",
				EnvVars: []string{"WORKWXCTL_AGENTID"},
			},
			&cli.StringFlag{
				Name:    flagQyapiHostOverride,
				Usage:   "使用 `HOST` 覆盖默认企业微信 API 地址",
				EnvVars: []string{"WORKWXCTL_QYAPI_HOST_OVERRIDE"},
			},
		},
		Commands: []*cli.Command{
			{
				Name:   "user-get",
				Usage:  "读取成员",
				Action: cmdUserGet,
			},
			{
				Name:   "send-text-message",
				Usage:  "发送纯文本消息",
				Action: cmdSendTextMessage,
				Flags: []cli.Flag{
					&cli.StringSliceFlag{
						Name:    flagToUser,
						Aliases: []string{flagToUserShort},
						Usage:   "收信用户 ID (可指定多次)",
					},
					&cli.StringSliceFlag{
						Name:    flagToParty,
						Aliases: []string{flagToPartyShort},
						Usage:   "收信部门 ID (可指定多次)",
					},
					&cli.StringSliceFlag{
						Name:    flagToTag,
						Aliases: []string{flagToTagShort},
						Usage:   "收信标签 ID (可指定多次)",
					},
					&cli.StringFlag{
						Name:    flagToChat,
						Aliases: []string{flagToChatShort},
						Usage:   "收信群聊 chatid (不可与其他收信人选项同时指定)",
					},
					&cli.BoolFlag{
						Name:  flagSafe,
						Usage: "作为保密消息发送",
					},
				},
			},
		},
	}
}
