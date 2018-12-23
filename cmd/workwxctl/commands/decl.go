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
				Name:   "accesstoken",
				Usage:  "获取 access token",
				Action: cmdAccesstoken,
			},
		},
	}
}
