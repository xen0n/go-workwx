package workwx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/url"
	"sync"
	"time"
)

// Workwx 企业微信客户端
type Workwx struct {
	opts options

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

// New 构造一个 Workwx 客户端对象，需要提供企业 ID
func New(corpID string, opts ...CtorOption) *Workwx {
	optionsObj := defaultOptions()

	for _, o := range opts {
		o.applyTo(&optionsObj)
	}

	return &Workwx{
		opts: optionsObj,

		CorpID: corpID,
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

func (c *WorkwxApp) composeQyapiURL(path string, req interface{}) *url.URL {
	values := url.Values{}
	if valuer, ok := req.(urlValuer); ok {
		values = valuer.intoURLValues()
	}

	// TODO: refactor
	base, err := url.Parse(c.opts.QYAPIHost)
	if err != nil {
		// TODO: error_chain
		panic(fmt.Sprintf("qyapiHost invalid: host=%s err=%+v", c.opts.QYAPIHost, err))
	}

	base.Path = path
	base.RawQuery = values.Encode()

	return base
}

func (c *WorkwxApp) composeQyapiURLWithToken(path string, req interface{}, withAccessToken bool) *url.URL {
	url := c.composeQyapiURL(path, req)

	if !withAccessToken {
		return url
	}

	// intensive mutex juggling action
	c.tokenMu.RLock()
	if c.accessToken == "" {
		c.tokenMu.RUnlock() // RWMutex doesn't like recursive locking
		// TODO: what to do with the possible error?
		_ = c.syncAccessToken()
		c.tokenMu.RLock()
	}
	tokenToUse := c.accessToken
	c.tokenMu.RUnlock()

	q := url.Query()
	q.Set("access_token", tokenToUse)
	url.RawQuery = q.Encode()

	return url
}

func (c *WorkwxApp) executeQyapiGet(path string, req urlValuer, respObj interface{}, withAccessToken bool) error {
	url := c.composeQyapiURLWithToken(path, req, withAccessToken)
	urlStr := url.String()

	resp, err := c.opts.HTTP.Get(urlStr)
	if err != nil {
		// TODO: error_chain
		return err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(respObj)
	if err != nil {
		// TODO: error_chain
		return err
	}

	return nil
}

func (c *WorkwxApp) executeQyapiJSONPost(path string, req bodyer, respObj interface{}, withAccessToken bool) error {
	url := c.composeQyapiURLWithToken(path, req, withAccessToken)
	urlStr := url.String()

	body, err := req.intoBody()
	if err != nil {
		// TODO: error_chain
		return err
	}

	resp, err := c.opts.HTTP.Post(urlStr, "application/json", bytes.NewReader(body))
	if err != nil {
		// TODO: error_chain
		return err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(respObj)
	if err != nil {
		// TODO: error_chain
		return err
	}

	return nil
}

func (c *WorkwxApp) executeQyapiMediaUpload(
	path string,
	req mediaUploader,
	respObj interface{},
	withAccessToken bool,
) error {
	url := c.composeQyapiURLWithToken(path, req, withAccessToken)
	urlStr := url.String()

	m := req.getMedia()

	// FIXME: use streaming upload to conserve memory!
	buf := bytes.Buffer{}
	mw := multipart.NewWriter(&buf)

	err := m.writeTo(mw)
	if err != nil {
		return err
	}

	err = mw.Close()
	if err != nil {
		return err
	}

	resp, err := c.opts.HTTP.Post(urlStr, mw.FormDataContentType(), &buf)
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
