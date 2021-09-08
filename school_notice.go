package workwx

type SchoolArticleMessage struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Picurl      string `json:"picurl"`
}

type SchoolMPNewsMessage struct {
	Title            string `json:"title"`
	ThumbMediaID     string `json:"thumb_media_id"`
	Author           string `json:"author"`
	ContentSourceUrl string `json:"content_source_url"`
	Content          string `json:"content"`
	Digest           string `json:"digest"`
}

type MiniProgramMessage struct {
	AppID        string `json:"appid"`
	Title        string `json:"title"`
	ThumbMediaID string `json:"thumb_media_id"`
	PagePath     string `json:"pagepath"`
}

// SendSchoolTextMessage 发送文本消息
func (c *WorkwxApp) SendSchoolTextMessage(recipient *Recipient, content string) error {
	return c.sendSchoolMessage(recipient, "text", map[string]interface{}{"content": content})
}

// SendSchoolImageMessage 发送图片消息
func (c *WorkwxApp) SendSchoolImageMessage(recipient *Recipient, mediaID string) error {
	return c.sendSchoolMessage(recipient, "image", map[string]interface{}{"media_id": mediaID})
}

// SendSchoolVoiceMessage 发送语音消息
func (c *WorkwxApp) SendSchoolVoiceMessage(recipient *Recipient, mediaID string) error {
	return c.sendSchoolMessage(recipient, "voice", map[string]interface{}{"media_id": mediaID})
}

// SendSchoolVideoMessage 发送视频消息
func (c *WorkwxApp) SendSchoolVideoMessage(
	recipient *Recipient,
	mediaID string,
	description string,
	title string,
) error {
	return c.sendSchoolMessage(recipient, "video", map[string]interface{}{
		"media_id":    mediaID,
		"description": description,
		"title":       title,
	})
}

// SendSchoolFileMessage 发送文件消息
func (c *WorkwxApp) SendSchoolFileMessage(recipient *Recipient, mediaID string) error {
	return c.sendSchoolMessage(recipient, "file", map[string]interface{}{"media_id": mediaID})
}

// SendSchoolNewsMessage 发送图文消息
func (c *WorkwxApp) SendSchoolNewsMessage(recipient *Recipient, articles []SchoolArticleMessage) error {
	return c.sendSchoolMessage(recipient, "news", map[string]interface{}{"articles": articles})
}

// SendSchoolMPNewsMessage 发送 mpnews 类型的图文消息
func (c *WorkwxApp) SendSchoolMPNewsMessage(recipient *Recipient, articles []SchoolMPNewsMessage) error {
	return c.sendSchoolMessage(recipient, "mpnews", map[string]interface{}{"articles": articles})
}

// SendSchoolMiniProgramMessage 发送小程序类型的消息
func (c *WorkwxApp) SendSchoolMiniProgramMessage(recipient *Recipient, message MiniProgramMessage) error {
	return c.sendSchoolMessage(recipient, "miniprogram", map[string]interface{}{
		"appid":          message.AppID,
		"title":          message.Title,
		"thumb_media_id": message.ThumbMediaID,
		"pagepath":       message.PagePath,
	})
}

// 发送学校通知底层接口
func (c *WorkwxApp) sendSchoolMessage(
	recipient *Recipient,
	msgtype string,
	content map[string]interface{},
) error {
	_, err := c.execSchoolMessageSend(reqSchoolMessage{
		ParentIDs:  recipient.ParentIDs,
		StudentIDs: recipient.StudentIDs,
		PartyIDs:   recipient.PartyIDs,
		AgentID:    c.AgentID,
		MsgType:    msgtype,
		Content:    content,
	})

	return err
}
