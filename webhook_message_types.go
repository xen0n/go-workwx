package workwx

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

type WebhookMessage interface {
	ToWebhookMessagePayload() (map[string]any, error)
	Validate() error
}

// See https://developer.work.weixin.qq.com/document/path/99110#%E5%A6%82%E4%BD%95%E4%BD%BF%E7%94%A8%E7%BE%A4%E6%9C%BA%E5%99%A8%E4%BA%BA

// WebhookTextMessage 文本消息
type WebhookTextMessage struct {
	Text struct {
		Content             string   `json:"content" validate:"required"`
		MentionedList       []string `json:"mentioned_list"`
		MentionedMobileList []string `json:"mentioned_mobile_list"`
	} `json:"text"`
}

var _ WebhookMessage = (*WebhookTextMessage)(nil)

func (t *WebhookTextMessage) ToWebhookMessagePayload() (map[string]any, error) {
	var dataMap = make(map[string]any)
	buf, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(buf, &dataMap)

	return dataMap, err
}

func (t *WebhookTextMessage) Validate() error {
	validate := validator.New()
	if err := validate.Struct(t); err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, err.Error())
		}
		return errors.New("validation failed: " + strings.Join(validationErrors, ", "))
	}
	return nil
}

// WebhookMarkdownMessage markdown
type WebhookMarkdownMessage struct {
	Markdown struct {
		Content string `json:"content" validate:"required"`
	} `json:"markdown" validate:"required"`
}

var _ WebhookMessage = (*WebhookMarkdownMessage)(nil)

func (t *WebhookMarkdownMessage) ToWebhookMessagePayload() (map[string]any, error) {
	var dataMap = make(map[string]any)
	buf, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(buf, &dataMap)

	return dataMap, err
}

func (t *WebhookMarkdownMessage) Validate() error {
	validate := validator.New()
	if err := validate.Struct(t); err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, err.Error())
		}
		return errors.New("validation failed: " + strings.Join(validationErrors, ", "))
	}
	return nil
}

// WebhookImageMessage 图片类型
type WebhookImageMessage struct {
	Image struct {
		Base64 string `json:"base64" validate:"required"` //图片内容的base64编码
		Md5    string `json:"md5" validate:"required"`
	} `json:"image" validate:"required"`
}

var _ WebhookMessage = (*WebhookImageMessage)(nil)

func (t *WebhookImageMessage) ToWebhookMessagePayload() (map[string]any, error) {
	var dataMap = make(map[string]any)
	buf, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(buf, &dataMap)

	return dataMap, err
}

func (t *WebhookImageMessage) Validate() error {
	validate := validator.New()
	if err := validate.Struct(t); err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, err.Error())
		}
		return errors.New("validation failed: " + strings.Join(validationErrors, ", "))
	}
	return nil
}

// WebhookVoiceMessage 语音类型
type WebhookVoiceMessage struct {
	Voice struct {
		MediaID string `json:"media_id" validate:"required"` //	语音文件id，通过下文的文件上传接口获取
	} `json:"voice" validate:"required"`
}

var _ WebhookMessage = (*WebhookVoiceMessage)(nil)

func (t *WebhookVoiceMessage) ToWebhookMessagePayload() (map[string]any, error) {
	var dataMap = make(map[string]any)
	buf, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(buf, &dataMap)

	return dataMap, err
}

func (t *WebhookVoiceMessage) Validate() error {
	validate := validator.New()
	if err := validate.Struct(t); err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, err.Error())
		}
		return errors.New("validation failed: " + strings.Join(validationErrors, ", "))
	}
	return nil
}

// WebhookImageArticles 图文类型
type WebhookImageArticles struct {
	News struct {
		Articles []struct {
			Title       string `json:"title" validate:"required"`
			Description string `json:"description" validate:"required"`
			URL         string `json:"url" validate:"required"`
			PicURL      string `json:"picurl"`
		} `json:"articles"`
	} `json:"news"`
}

var _ WebhookMessage = (*WebhookImageArticles)(nil)

func (t *WebhookImageArticles) ToWebhookMessagePayload() (map[string]any, error) {
	var dataMap = make(map[string]any)
	buf, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(buf, &dataMap)

	return dataMap, err
}

func (t *WebhookImageArticles) Validate() error {
	validate := validator.New()
	if err := validate.Struct(t); err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, err.Error())
		}
		return errors.New("validation failed: " + strings.Join(validationErrors, ", "))
	}
	return nil
}

type WebhookFileMessage struct {
	File struct {
		MediaID string `json:"media_id" validate:"required"` //文件id，通过下文的文件上传接口获取
	} `json:"file" validate:"required"`
}

var _ WebhookMessage = (*WebhookFileMessage)(nil)

func (t *WebhookFileMessage) ToWebhookMessagePayload() (map[string]any, error) {
	var dataMap = make(map[string]any)
	buf, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(buf, &dataMap)

	return dataMap, err
}

func (t *WebhookFileMessage) Validate() error {
	validate := validator.New()
	if err := validate.Struct(t); err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, err.Error())
		}
		return errors.New("validation failed: " + strings.Join(validationErrors, ", "))
	}
	return nil
}

type WebhookTemplateCardMessage struct {
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

var _ WebhookMessage = (*WebhookTemplateCardMessage)(nil)

func (t *WebhookTemplateCardMessage) ToWebhookMessagePayload() (map[string]any, error) {
	var dataMap = make(map[string]any)
	buf, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(buf, &dataMap)

	return dataMap, err
}

func (t *WebhookTemplateCardMessage) Validate() error {
	validate := validator.New()
	if err := validate.Struct(t); err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, err.Error())
		}
		return errors.New("validation failed: " + strings.Join(validationErrors, ", "))
	}
	return nil
}
