package workwx

// MentionAll 表示提醒所有人（“@所有人”）的特殊标记
const MentionAll = "@all"

// Mentions 群机器人消息的提醒设置
type Mentions struct {
	// UserIDs userid 的列表，提醒群中的指定成员（@某个成员），MentionAll 表示提醒所有人
	//
	// 如果开发者获取不到 userid，可使用 Mobiles
	UserIDs []string
	// Mobiles 手机号列表，提醒手机号对应的群成员（@某个成员），MentionAll 表示提醒所有人
	//
	// 如果开发者获取不到 userid，可使用该列表，否则可使用 UserIDs
	Mobiles []string
}

// SendTextMessage 发送文本消息
func (c *WebhookClient) SendTextMessage(
	content string,
	mentions *Mentions,
) error {
	params := map[string]any{
		"content": content,
	}

	if mentions != nil {
		if len(mentions.UserIDs) > 0 {
			params["mentioned_list"] = mentions.UserIDs
		}

		if len(mentions.Mobiles) > 0 {
			params["mentioned_mobile_list"] = mentions.Mobiles
		}
	}

	return c.sendMessage("text", params)
}

// SendMarkdownMessage 发送 Markdown 消息
//
// NOTE: 使用群机器人接口发送 Markdown 消息时，不能传递 Mentions 结构体，而需要使用
// `<@userid>` 的特殊扩展语法来表示 at 给定的 userid。
func (c *WebhookClient) SendMarkdownMessage(
	content string,
) error {
	params := map[string]any{
		"content": content,
	}

	return c.sendMessage("markdown", params)
}

// sendMessage 发送消息底层接口
func (c *WebhookClient) sendMessage(
	msgtype string,
	content map[string]any,
) error {
	req := map[string]any{
		"msgtype": msgtype,
		msgtype:   content,
	}

	err := c.executeQyapiJSONPost("/cgi-bin/webhook/send", req, nil)
	if err != nil {
		return err
	}

	return nil
}

// SendMessage 机器人支持文本（text）、markdown（markdown）、图片（image）、图文（news）、文件（file）、语音（voice）、模板卡片（template_card）七种消息类型
func (c *WebhookClient) SendMessage(msg WebHookMessage) error {

	if err := msg.Validate(); err != nil {
		return err
	}
	req, err := msg.Struct2Map()
	if err != nil {
		return err
	}
	switch msg.(type) {
	case *TextMessage:
		req["msgtype"] = "text"
	case *MarkdownMessage:
		req["msgtype"] = "markdown"

	case *ImageMessage:
		req["msgtype"] = "image"
	case *ImageArticles:
		req["msgtype"] = "news"

	case *FileMessage:
		req["msgtype"] = "file"
	case *VoiceMessage:
		req["msgtype"] = "voice"
	case *TemplateCardMessage:
		req["msgtype"] = "template_card"
	}
	return c.executeQyapiJSONPost("/cgi-bin/webhook/send", req, nil)

}
