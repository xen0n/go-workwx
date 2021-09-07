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

func (c *WorkwxApp) SendSchoolTextMessage(recipient *Recipient, content string) error {
	return c.sendSchoolMessage(recipient, "text", map[string]interface{}{"content": content})
}

func (c *WorkwxApp) SendSchoolImageMessage(recipient *Recipient, mediaID string) error {
	return c.sendSchoolMessage(recipient, "image", map[string]interface{}{"media_id": mediaID})
}

func (c *WorkwxApp) SendSchoolVoiceMessage(recipient *Recipient, mediaID string) error {
	return c.sendSchoolMessage(recipient, "voice", map[string]interface{}{"media_id": mediaID})
}

func (c *WorkwxApp) SendSchoolVideoMessage(
	recipient *Recipient,
	mediaID string,
	description string,
	title string,
) error {
	return c.sendSchoolMessage(recipient, "voice", map[string]interface{}{
		"media_id":    mediaID,
		"description": description,
		"title":       title,
	})
}

func (c *WorkwxApp) SendSchoolFileMessage(recipient *Recipient, mediaID string) error {
	return c.sendSchoolMessage(recipient, "voice", map[string]interface{}{"media_id": mediaID})
}

func (c *WorkwxApp) SendSchoolArticleMessage(recipient *Recipient, articles []SchoolArticleMessage) error {
	return c.sendSchoolMessage(recipient, "voice", map[string]interface{}{"articles": articles})
}

func (c *WorkwxApp) SendSchoolMPNewMessage(recipient *Recipient, articles []SchoolMPNewsMessage) error {
	return c.sendSchoolMessage(recipient, "voice", map[string]interface{}{"articles": articles})
}

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
