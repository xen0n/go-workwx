package workwx

import (
	"errors"
)

// SendTextMessage 发送文本消息
//
// 收件人参数如果仅设置了 `ChatID` 字段，则为【发送消息到群聊会话】接口调用；
// 否则为单纯的【发送应用消息】接口调用。
func (c *WorkwxApp) SendTextMessage(
	recipient *Recipient,
	content string,
	isSafe bool,
) error {
	return c.sendMessage(recipient, "text", map[string]any{"content": content}, isSafe)
}

// SendImageMessage 发送图片消息
//
// 收件人参数如果仅设置了 `ChatID` 字段，则为【发送消息到群聊会话】接口调用；
// 否则为单纯的【发送应用消息】接口调用。
func (c *WorkwxApp) SendImageMessage(
	recipient *Recipient,
	mediaID string,
	isSafe bool,
) error {
	return c.sendMessage(
		recipient,
		"image",
		map[string]any{
			"media_id": mediaID,
		}, isSafe,
	)
}

// SendVoiceMessage 发送语音消息
//
// 收件人参数如果仅设置了 `ChatID` 字段，则为【发送消息到群聊会话】接口调用；
// 否则为单纯的【发送应用消息】接口调用。
func (c *WorkwxApp) SendVoiceMessage(
	recipient *Recipient,
	mediaID string,
	isSafe bool,
) error {
	return c.sendMessage(
		recipient,
		"voice",
		map[string]any{
			"media_id": mediaID,
		}, isSafe,
	)
}

// SendVideoMessage 发送视频消息
//
// 收件人参数如果仅设置了 `ChatID` 字段，则为【发送消息到群聊会话】接口调用；
// 否则为单纯的【发送应用消息】接口调用。
func (c *WorkwxApp) SendVideoMessage(
	recipient *Recipient,
	mediaID string,
	description string,
	title string,
	isSafe bool,
) error {
	return c.sendMessage(
		recipient,
		"video",
		map[string]any{
			"media_id":    mediaID,
			"description": description, // TODO: 零值
			"title":       title,       // TODO: 零值
		}, isSafe,
	)
}

// SendFileMessage 发送文件消息
//
// 收件人参数如果仅设置了 `ChatID` 字段，则为【发送消息到群聊会话】接口调用；
// 否则为单纯的【发送应用消息】接口调用。
func (c *WorkwxApp) SendFileMessage(
	recipient *Recipient,
	mediaID string,
	isSafe bool,
) error {
	return c.sendMessage(
		recipient,
		"file",
		map[string]any{
			"media_id": mediaID,
		}, isSafe,
	)
}

// SendTextCardMessage 发送文本卡片消息
//
// 收件人参数如果仅设置了 `ChatID` 字段，则为【发送消息到群聊会话】接口调用；
// 否则为单纯的【发送应用消息】接口调用。
func (c *WorkwxApp) SendTextCardMessage(
	recipient *Recipient,
	title string,
	description string,
	url string,
	buttonText string,
	isSafe bool,
) error {
	return c.sendMessage(
		recipient,
		"textcard",
		map[string]any{
			"title":       title,
			"description": description,
			"url":         url,
			"btntxt":      buttonText, // TODO: 零值
		}, isSafe,
	)
}

// SendNewsMessage 发送图文消息
//
// 收件人参数如果仅设置了 `ChatID` 字段，则为【发送消息到群聊会话】接口调用；
// 否则为单纯的【发送应用消息】接口调用。
func (c *WorkwxApp) SendNewsMessage(
	recipient *Recipient,
	title string,
	description string,
	url string,
	picURL string,
	isSafe bool,
) error {
	return c.sendMessage(
		recipient,
		"news",
		map[string]any{
			"title":       title,
			"description": description, // TODO: 零值
			"url":         url,
			"picurl":      picURL, // TODO: 零值
		}, isSafe,
	)
}

// SendMPNewsMessage 发送 mpnews 类型的图文消息
//
// 收件人参数如果仅设置了 `ChatID` 字段，则为【发送消息到群聊会话】接口调用；
// 否则为单纯的【发送应用消息】接口调用。
func (c *WorkwxApp) SendMPNewsMessage(
	recipient *Recipient,
	title string,
	thumbMediaID string,
	author string,
	sourceContentURL string,
	content string,
	digest string,
	isSafe bool,
) error {
	return c.sendMessage(
		recipient,
		"mpnews",
		map[string]any{
			// TODO: 支持发送多条图文
			"articles": []any{
				map[string]any{
					"title":              title,
					"thumb_media_id":     thumbMediaID,
					"author":             author,           // TODO: 零值
					"content_source_url": sourceContentURL, // TODO: 零值
					"content":            content,
					"digest":             digest,
				},
			},
		}, isSafe,
	)
}

// SendMarkdownMessage 发送 Markdown 消息
//
// 仅支持 Markdown 的子集，详见[官方文档](https://work.weixin.qq.com/api/doc#90002/90151/90854/%E6%94%AF%E6%8C%81%E7%9A%84markdown%E8%AF%AD%E6%B3%95)。
//
// 收件人参数如果仅设置了 `ChatID` 字段，则为【发送消息到群聊会话】接口调用；
// 否则为单纯的【发送应用消息】接口调用。
func (c *WorkwxApp) SendMarkdownMessage(
	recipient *Recipient,
	content string,
	isSafe bool,
) error {
	return c.sendMessage(recipient, "markdown", map[string]any{"content": content}, isSafe)
}

// SendTaskCardMessage 发送 任务卡片 消息
func (c *WorkwxApp) SendTaskCardMessage(
	recipient *Recipient,
	title string,
	description string,
	url string,
	taskid string,
	btn []TaskCardBtn,
	isSafe bool,
) error {
	return c.sendMessage(
		recipient,
		"taskcard",
		map[string]any{
			"title":       title,
			"description": description,
			"url":         url,
			"task_id":     taskid,
			"btn":         btn,
		}, isSafe,
	)
}

// sendMessage 发送消息底层接口
//
// 收件人参数如果仅设置了 `ChatID` 字段，则为【发送消息到群聊会话】接口调用；
// 否则为单纯的【发送应用消息】接口调用。
func (c *WorkwxApp) sendMessage(
	recipient *Recipient,
	msgtype string,
	content map[string]any,
	isSafe bool,
) error {
	isApichatSendRequest := false
	if !recipient.isValidForMessageSend() {
		if !recipient.isValidForAppchatSend() {
			// TODO: better error
			return errors.New("recipient invalid for message sending")
		}

		// 发送给群聊
		isApichatSendRequest = true
	}

	req := reqMessage{
		ToUser:  recipient.UserIDs,
		ToParty: recipient.PartyIDs,
		ToTag:   recipient.TagIDs,
		ChatID:  recipient.ChatID,
		AgentID: c.AgentID,
		MsgType: msgtype,
		Content: content,
		IsSafe:  isSafe,
	}

	var resp respMessageSend
	var err error
	if isApichatSendRequest {
		resp, err = c.execAppchatSend(req)
	} else {
		resp, err = c.execMessageSend(req)
	}

	if err != nil {
		return err
	}

	// TODO: what to do with resp?
	_ = resp
	return nil
}
