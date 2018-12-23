package workwx

import (
	"errors"
)

const messageSendEndpoint = "/cgi-bin/message/send"
const apichatSendEndpoint = "/cgi-bin/appchat/send"

// SendTextMessage 发送文本消息
//
// 收件人参数如果仅设置了 `ChatID` 字段，则为【发送消息到群聊会话】接口调用；
// 否则为单纯的【发送应用消息】接口调用。
func (c *WorkwxApp) SendTextMessage(
	recipient *Recipient,
	content string,
	isSafe bool,
) error {
	return c.sendMessage(recipient, "text", map[string]interface{}{"content": content}, isSafe)
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
	return c.sendMessage(recipient, "markdown", map[string]interface{}{"content": content}, isSafe)
}

// sendMessage 发送消息底层接口
//
// 收件人参数如果仅设置了 `ChatID` 字段，则为【发送消息到群聊会话】接口调用；
// 否则为单纯的【发送应用消息】接口调用。
func (c *WorkwxApp) sendMessage(
	recipient *Recipient,
	msgtype string,
	content map[string]interface{},
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

	apiPath := messageSendEndpoint
	if isApichatSendRequest {
		apiPath = apichatSendEndpoint
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
	err := c.executeQyapiJSONPost(apiPath, req, &resp, true)
	if err != nil {
		// TODO: error_chain
		return err
	}

	// TODO: what to do with resp?
	return nil
}
