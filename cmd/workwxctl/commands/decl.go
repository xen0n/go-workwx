package commands

import (
	"gopkg.in/urfave/cli.v2"
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
			&cli.StringFlag{
				Name:    flagTLSKeyLogFile,
				Usage:   "将 HTTPS 会话所用密钥写入 `LOGFILE` 以便 Wireshark 等工具读取",
				EnvVars: []string{"WORKWXCTL_TLS_KEY_LOGFILE"},
			},
		},
		Commands: []*cli.Command{
			{
				Name:   "user-get",
				Usage:  "读取成员",
				Action: cmdUserGet,
			},
			{
				Name:   "dept-list",
				Usage:  "获取部门列表",
				Action: cmdDeptList,
			},
			{
				Name:   "appchat-get",
				Usage:  "获取群聊会话",
				Action: cmdAppchatGet,
			},
			{
				Name:   "send-message",
				Usage:  "发送消息",
				Action: cmdSendMessage,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  flagMessageType,
						Usage: "发送消息的类型: text, image, voice, video, file, textcard, news, mpnews, markdown",
					},
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
			{
				Name:   "upload-temp-media",
				Usage:  "上传临时素材",
				Action: cmdUploadTempMedia,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  flagMediaType,
						Usage: "媒体文件类型，分别有图片（image）、语音（voice）、视频（video），普通文件(file)",
					},
				},
			},
		},
	}
}
