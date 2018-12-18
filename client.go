package workwx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"
)

// DefaultQYAPIHost 默认企业微信 API Host
const DefaultQYAPIHost = "https://qyapi.weixin.qq.com"

// Workwx 企业微信客户端
type Workwx struct {
	qyapiHost string
	http      *http.Client

	// CorpID 企业 ID，必填
	CorpID string
}

// WorkwxApp 企业微信客户端（分应用）
type WorkwxApp struct {
	*Workwx

	// CorpSecret 应用的凭证密钥，必填
	CorpSecret string
	// AgentID 应用 ID，必填
	AgentID int64

	tokenMu        *sync.RWMutex
	accessToken    string
	tokenExpiresIn time.Duration
	lastRefresh    time.Time
}

// Default 构造一个默认配置的 `Workwx`，自带独立 `http.Client`，需要自行设置 `CorpID`
//
// impl Default for Workwx
func Default() *Workwx {
	return WithHTTPClient(&http.Client{})
}

// WithHTTPClient 用给定的 `http.Client` 构造一个 `Workwx`，需要自行设置 `CorpID`
func WithHTTPClient(client *http.Client) *Workwx {
	return &Workwx{
		qyapiHost: DefaultQYAPIHost,
		http:      client,
	}
}

// WithApp 构造本企业下某自建 app 的客户端
func (c *Workwx) WithApp(corpSecret string, agentID int64) *WorkwxApp {
	return &WorkwxApp{
		Workwx: c,

		CorpSecret: corpSecret,
		AgentID:    agentID,

		tokenMu:     &sync.RWMutex{},
		accessToken: "",
		lastRefresh: time.Time{},
	}
}

//
// impl WorkwxApp
//

func (c *WorkwxApp) composeQyapiURL(path string, req interface{}) string {
	values := url.Values{}
	if valuer, ok := req.(urlValuer); ok {
		values = valuer.IntoURLValues()
	}

	// TODO: refactor
	base, err := url.Parse(c.qyapiHost)
	if err != nil {
		// TODO: error_chain
		panic(fmt.Sprintf("qyapiHost invalid: host=%s err=%+v", c.qyapiHost, err))
	}

	base.Path = path
	base.RawQuery = values.Encode()

	return base.String()
}

func (c *WorkwxApp) executeQyapiGet(path string, req urlValuer, respObj interface{}) error {
	url := c.composeQyapiURL(path, req)

	resp, err := c.http.Get(url)
	defer resp.Body.Close()

	if err != nil {
		// TODO: error_chain
		return err
	}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(respObj)
	if err != nil {
		// TODO: error_chain
		return err
	}

	return nil
}

func (c *WorkwxApp) executeQyapiJSONPost(path string, req bodyer, respObj interface{}) error {
	url := c.composeQyapiURL(path, req)
	body := req.IntoBody()

	resp, err := c.http.Post(url, "application/json", bytes.NewReader(body))
	defer resp.Body.Close()

	if err != nil {
		// TODO: error_chain
		return err
	}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(respObj)
	if err != nil {
		// TODO: error_chain
		return err
	}

	return nil
}
