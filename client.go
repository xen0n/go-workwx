package workwx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/url"
	"strings"
)

// Workwx 企业微信客户端
type Workwx struct {
	opts options

	// CorpID 企业 ID，必填
	CorpID string
}

// WorkwxApp 企业微信客户端（分应用）
//
//nolint:revive // The (stuttering) name is part of public API, so cannot be fixed without a v2 bump
type WorkwxApp struct {
	*Workwx

	// CorpSecret 应用的凭证密钥，必填
	CorpSecret string
	// AgentID 应用 ID，必填
	AgentID int64

	accessToken            *token
	jsapiTicket            *token
	jsapiTicketAgentConfig *token
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
	app := WorkwxApp{
		Workwx: c,

		CorpSecret: corpSecret,
		AgentID:    agentID,
	}

	app.accessToken = newToken(c.opts.AccessTokenProvider, app.getAccessToken)
	app.jsapiTicket = newToken(c.opts.JSAPITicketProvider, app.getJSAPITicket)
	app.jsapiTicketAgentConfig = newToken(c.opts.JSAPITicketAgentConfigProvider, app.getJSAPITicketAgentConfig)

	return &app
}

func (c *WorkwxApp) composeQyapiURL(path string, req any) (*url.URL, error) {
	values := url.Values{}
	if valuer, ok := req.(urlValuer); ok {
		values = valuer.intoURLValues()
	}

	// TODO: refactor
	base, err := url.Parse(c.opts.QYAPIHost)
	if err != nil {
		return nil, fmt.Errorf("qyapiHost invalid: host=%s err=%w", c.opts.QYAPIHost, err)
	}

	base.Path = path
	base.RawQuery = values.Encode()

	return base, nil
}

func (c *WorkwxApp) composeQyapiURLWithToken(path string, req any, withAccessToken bool) (*url.URL, error) {
	url, err := c.composeQyapiURL(path, req)
	if err != nil {
		return nil, err
	}

	if !withAccessToken {
		return url, nil
	}

	tok, err := c.accessToken.getToken()
	if err != nil {
		return nil, err
	}

	q := url.Query()
	q.Set("access_token", tok)
	url.RawQuery = q.Encode()

	return url, nil
}

func executeQyapiGet[T urlValuer, U tryIntoErr](
	c *WorkwxApp,
	path string,
	req T,
	respObj U,
	withAccessToken bool,
) error {
	url, err := c.composeQyapiURLWithToken(path, req, withAccessToken)
	if err != nil {
		return err
	}
	urlStr := url.String()

	resp, err := c.opts.HTTP.Get(urlStr)
	if err != nil {
		return makeRequestErr(err)
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(respObj)
	if err != nil {
		return makeRespUnmarshalErr(err)
	}

	if bizErr := respObj.TryIntoErr(); bizErr != nil {
		return bizErr
	}

	return nil
}

// executeQyapiGetBinary 执行 GET 请求，返回二进制响应体（用于下载文件等）
// 如果响应是 JSON 错误，则返回错误
// executeQyapiGetStream 执行 GET 请求，返回数据流。
// 注意：调用者必须负责关闭返回的 io.ReadCloser
func executeQyapiGetBinary[T urlValuer](
	c *WorkwxApp,
	path string,
	req T,
	withAccessToken bool,
) (io.ReadCloser, error) {
	urlObj, err := c.composeQyapiURLWithToken(path, req, withAccessToken)
	if err != nil {
		return nil, err
	}

	resp, err := c.opts.HTTP.Get(urlObj.String())
	if err != nil {
		return nil, makeRequestErr(err)
	}

	// 重点逻辑：先检查这是否是一个 JSON 报错
	// 企微 API 坑点：下载文件失败时，会返回 application/json 和错误码
	contentType := resp.Header.Get("Content-Type")
	if strings.Contains(contentType, "application/json") {
		// 既然是 JSON，说明不是我们要的文件流，必须读出来检查错误
		defer resp.Body.Close() // 读完错误信息要关闭

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, makeRespUnmarshalErr(err)
		}

		var errResp respCommon
		if err := json.Unmarshal(body, &errResp); err == nil {
			if errResp.ErrCode != 0 {
				return nil, &WorkwxClientError{
					Code: errResp.ErrCode,
					Msg:  errResp.ErrMsg,
				}
			}
		}
		// 极少情况：Content-Type 是 json 但内容没报错，或者无法解析，
		// 这种情况视作业务逻辑不明，返回错误比较安全
		return nil, fmt.Errorf("unexpected json response for download api")
	}

	// 检查 HTTP 状态码 (非 200 也是错误)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		defer resp.Body.Close()
		return nil, fmt.Errorf("http request failed: %s", resp.Status)
	}

	// 成功！返回 resp.Body 这个“水管”给上层去读
	// 千万不要在这里 defer resp.Body.Close()，否则上层拿到的是关掉的流
	return resp.Body, nil
}

func executeQyapiJSONPost[T bodyer, U tryIntoErr](
	c *WorkwxApp,
	path string,
	req T,
	respObj U,
	withAccessToken bool,
) error {
	url, err := c.composeQyapiURLWithToken(path, req, withAccessToken)
	if err != nil {
		return err
	}
	urlStr := url.String()

	body, err := req.intoBody()
	if err != nil {
		return makeReqMarshalErr(err)
	}

	resp, err := c.opts.HTTP.Post(urlStr, "application/json", bytes.NewReader(body))
	if err != nil {
		return makeRequestErr(err)
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(respObj)
	if err != nil {
		return makeRespUnmarshalErr(err)
	}

	if bizErr := respObj.TryIntoErr(); bizErr != nil {
		return bizErr
	}

	return nil
}

func executeQyapiMediaUpload[T mediaUploader, U tryIntoErr](
	c *WorkwxApp,
	path string,
	req T,
	respObj U,
	withAccessToken bool,
) error {
	url, err := c.composeQyapiURLWithToken(path, req, withAccessToken)
	if err != nil {
		return err
	}
	urlStr := url.String()

	m := req.getMedia()

	// FIXME: use streaming upload to conserve memory!
	buf := bytes.Buffer{}
	mw := multipart.NewWriter(&buf)

	err = m.writeTo(mw)
	if err != nil {
		return err
	}

	err = mw.Close()
	if err != nil {
		return err
	}

	resp, err := c.opts.HTTP.Post(urlStr, mw.FormDataContentType(), &buf)
	if err != nil {
		return makeRequestErr(err)
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(respObj)
	if err != nil {
		return makeRespUnmarshalErr(err)
	}

	if bizErr := respObj.TryIntoErr(); bizErr != nil {
		return bizErr
	}

	return nil
}
