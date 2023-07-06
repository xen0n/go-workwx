package commands

import (
	"github.com/urfave/cli/v2"

	"github.com/xen0n/go-workwx/v2"
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
				Name:    flagWebhookKey,
				Usage:   "使用 `KEY` 作为群机器人 webhook key",
				EnvVars: []string{"WORKWXCTL_WEBHOOK_KEY"},
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
				Name:   "user-list-by-dept",
				Usage:  "获取部门成员详情",
				Action: cmdUserListByDept,
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

					// 发消息参数
					&cli.StringFlag{
						Name:  flagMediaID,
						Usage: "图片媒体文件id，可以调用上传临时素材接口获取",
					},
					&cli.StringFlag{
						Name:  flagThumbMediaID,
						Usage: "图文消息缩略图的media_id, 可以通过素材管理接口获得。",
					},
					&cli.StringFlag{
						Name:  flagAuthor,
						Usage: "图文消息的作者，不超过64个字节",
					},
					&cli.StringFlag{
						Name:  flagDescription,
						Usage: "描述，不超过512个字节，超过会自动截断",
					},
					&cli.StringFlag{
						Name:  flagTitle,
						Usage: "标题，不超过128个字节，超过会自动截断",
					},
					&cli.StringFlag{
						Name:  flagURL,
						Usage: "点击后跳转的链接。",
					},
					&cli.StringFlag{
						Name:  flagPicURL,
						Usage: "图文消息的图片链接，支持JPG、PNG格式，较好的效果为大图1068*455，小图150*150。",
					},
					&cli.StringFlag{
						Name:  flagButtonText,
						Usage: "按钮文字。 默认为“详情”， 不超过4个文字，超过自动截断。",
					},
					&cli.StringFlag{
						Name:  flagSourceContentURL,
						Usage: "图文消息点击“阅读原文”之后的页面链接",
					},
					&cli.StringFlag{
						Name:  flagDigest,
						Usage: "图文消息的描述，不超过512个字节，超过会自动截断",
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
			{
				Name:   "webhook-send-message",
				Usage:  "使用群机器人接口发送消息",
				Action: cmdWebhookSendMessage,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  flagMessageType,
						Usage: "发送消息的类型: text, markdown",
					},
					&cli.StringSliceFlag{
						Name:    flagMentionUser,
						Aliases: []string{flagToUserShort},
						Usage:   "需要被提醒的用户 ID (可指定多次), 特殊值 '" + workwx.MentionAll + "' 表示提醒所有人",
					},
					&cli.StringSliceFlag{
						Name:    flagMentionMobile,
						Aliases: []string{flagMentionMobileShort},
						Usage:   "需要被提醒的用户手机号 (可指定多次), 特殊值 '" + workwx.MentionAll + "' 表示提醒所有人",
					},
				},
			},
		},
	}
}
