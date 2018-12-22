package workwx

import (
	"encoding/json"
	"net/url"
	"strings"
)

type reqAccessToken struct {
	CorpID     string
	CorpSecret string
}

// IntoURLValues 转换为 url.Values 类型
//
// impl urlValuer for reqAccessToken
func (x reqAccessToken) IntoURLValues() url.Values {
	return url.Values{
		"corpid":     {x.CorpID},
		"corpsecret": {x.CorpSecret},
	}
}

type respCommon struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// IsOK 响应体是否为一次成功请求的响应
//
// 实现依据: https://work.weixin.qq.com/api/doc#10013
//
// > 企业微信所有接口，返回包里都有errcode、errmsg。
// > 开发者需根据errcode是否为0判断是否调用成功(errcode意义请见全局错误码)。
// > 而errmsg仅作参考，后续可能会有变动，因此不可作为是否调用成功的判据。
func (x *respCommon) IsOK() bool {
	return x.ErrCode == 0
}

type respAccessToken struct {
	respCommon

	AccessToken   string `json:"access_token"`
	ExpiresInSecs int64  `json:"expires_in"`
}

// reqMessage 消息发送请求
type reqMessage struct {
	ToUser  []string
	ToParty []string
	ToTag   []string
	ChatID  string
	AgentID int64
	MsgType string
	Content map[string]interface{}
	IsSafe  bool
}

// IntoBody 转换为请求体的 []byte 类型
//
// impl bodyer for reqMessage
func (x reqMessage) IntoBody() ([]byte, error) {
	// fuck
	safeInt := 0
	if x.IsSafe {
		safeInt = 1
	}

	obj := map[string]interface{}{
		"msgtype": x.MsgType,
		"agentid": x.AgentID,
		"safe":    safeInt,
	}

	// msgtype polymorphism
	obj[x.MsgType] = x.Content

	// 复用这个结构体，因为是 package-private 的所以这么做没风险
	if x.ChatID != "" {
		obj["chatid"] = x.ChatID
	} else {
		obj["touser"] = strings.Join(x.ToUser, "|")
		obj["toparty"] = strings.Join(x.ToParty, "|")
		obj["totag"] = strings.Join(x.ToTag, "|")
	}

	result, err := json.Marshal(obj)
	if err != nil {
		// should never happen unless OOM or similar bad things
		// TODO: error_chain
		return nil, err
	}

	return result, nil
}

// respMessageSend 消息发送响应
type respMessageSend struct {
	respCommon

	InvalidUsers   string `json:"invaliduser"`
	InvalidParties string `json:"invalidparty"`
	InvalidTags    string `json:"invalidtag"`
}
