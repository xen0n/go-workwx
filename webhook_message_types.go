package workwx

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

type WebhookMessage interface {
	ToWebhookMessagePayload() (map[string]any, error)
	Validate() error
}

// See https://developer.work.weixin.qq.com/document/path/99110#%E5%A6%82%E4%BD%95%E4%BD%BF%E7%94%A8%E7%BE%A4%E6%9C%BA%E5%99%A8%E4%BA%BA

// TextMessage 文本消息
type TextMessage struct {
	Text struct {
		Content             string   `json:"content" validate:"required"`
		MentionedList       []string `json:"mentioned_list"`
		MentionedMobileList []string `json:"mentioned_mobile_list"`
	} `json:"text"`
}

var _ WebhookMessage = (*TextMessage)(nil)

func (t *TextMessage) ToWebhookMessagePayload() (map[string]any, error) {
	var dataMap = make(map[string]any)
	buf, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(buf, &dataMap)

	return dataMap, err
}

func (t *TextMessage) Validate() error {
	return validate.Struct(t)
}

// MarkdownMessage markdown
type MarkdownMessage struct {
	Markdown struct {
		Content string `json:"content" validate:"required"`
	} `json:"markdown" validate:"required"`
}

var _ WebhookMessage = (*MarkdownMessage)(nil)

func (t *MarkdownMessage) ToWebhookMessagePayload() (map[string]any, error) {
	var dataMap = make(map[string]any)
	buf, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(buf, &dataMap)

	return dataMap, err
}

func (t *MarkdownMessage) Validate() error {
	return validate.Struct(t)
}

// ImageMessage 图片类型
type ImageMessage struct {
	Image struct {
		Base64 string `json:"base64"   validate:"required"` //图片内容的base64编码
		Md5    string `json:"md5"  validate:"required"`
	} `json:"image" validate:"required"`
}

var _ WebhookMessage = (*ImageMessage)(nil)

func (t *ImageMessage) ToWebhookMessagePayload() (map[string]any, error) {
	var dataMap = make(map[string]any)
	buf, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(buf, &dataMap)

	return dataMap, err
}

func (t *ImageMessage) Validate() error {
	return validate.Struct(t)
}

// VoiceMessage 语音类型
type VoiceMessage struct {
	Voice struct {
		MediaID string `json:"media_id" validate:"required" ` //	语音文件id，通过下文的文件上传接口获取
	} `json:"voice" validate:"required"`
}

var _ WebhookMessage = (*VoiceMessage)(nil)

func (t *VoiceMessage) ToWebhookMessagePayload() (map[string]any, error) {
	var dataMap = make(map[string]any)
	buf, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(buf, &dataMap)

	return dataMap, err
}

func (t *VoiceMessage) Validate() error {
	return validate.Struct(t)
}

// ImageArticles 图文类型
type ImageArticles struct {
	News struct {
		Articles []struct {
			Title       string `json:"title" validate:"required"`
			Description string `json:"description" validate:"required"`
			URL         string `json:"url" validate:"required"`
			PicURL      string `json:"picurl"`
		} `json:"articles"`
	} `json:"news"`
}

var _ WebhookMessage = (*ImageArticles)(nil)

func (t *ImageArticles) ToWebhookMessagePayload() (map[string]any, error) {
	var dataMap = make(map[string]any)
	buf, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(buf, &dataMap)

	return dataMap, err
}

func (t *ImageArticles) Validate() error {
	return validate.Struct(t)
}

type FileMessage struct {
	File struct {
		MediaID string `json:"media_id" validate:"required"` //文件id，通过下文的文件上传接口获取
	} `json:"file" validate:"required"`
}

var _ WebhookMessage = (*FileMessage)(nil)

func (t *FileMessage) ToWebhookMessagePayload() (map[string]any, error) {
	var dataMap = make(map[string]any)
	buf, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(buf, &dataMap)

	return dataMap, err
}

func (t *FileMessage) Validate() error {
	return validate.Struct(t)
}

type TemplateCardMessage struct {
	TemplateCard struct {
		CardType string `json:"card_type" validate:"required"`
		Source   struct {
			IconURL   string `json:"icon_url"`
			Desc      string `json:"desc"`
			DescColor int    `json:"desc_color"`
		} `json:"source"`
		MainTitle struct {
			Title string `json:"title"`
			Desc  string `json:"desc"`
		} `json:"main_title" validate:"required"`
		EmphasisContent struct {
			Title string `json:"title"`
			Desc  string `json:"desc"`
		} `json:"emphasis_content"`
		QuoteArea struct {
			Type      int    `json:"type"`
			URL       string `json:"url"`
			Appid     string `json:"appid"`
			Pagepath  string `json:"pagepath"`
			Title     string `json:"title"`
			QuoteText string `json:"quote_text"`
		} `json:"quote_area"`
		SubTitleText          string `json:"sub_title_text"`
		HorizontalContentList []struct {
			Keyname string `json:"keyname" validate:"required"`
			Value   string `json:"value"`
			Type    int    `json:"type,omitempty"`
			URL     string `json:"url,omitempty"`
			MediaID string `json:"media_id,omitempty"`
		} `json:"horizontal_content_list"`
		JumpList []struct {
			Type     int    `json:"type"`
			URL      string `json:"url,omitempty"`
			Title    string `json:"title" validate:"required"`
			Appid    string `json:"appid,omitempty"`
			Pagepath string `json:"pagepath,omitempty"`
		} `json:"jump_list"`
		CardAction struct {
			Type     int    `json:"type" validate:"required"`
			URL      string `json:"url"`
			Appid    string `json:"appid"`
			Pagepath string `json:"pagepath"`
		} `json:"card_action" validate:"required"`
	} `json:"template_card"`
}

var _ WebhookMessage = (*TemplateCardMessage)(nil)

func (t *TemplateCardMessage) ToWebhookMessagePayload() (map[string]any, error) {
	var dataMap = make(map[string]any)
	buf, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(buf, &dataMap)

	return dataMap, err
}

func (t *TemplateCardMessage) Validate() error {
	return validate.Struct(t)
}
