package workwx

import (
	"context"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
)

type tokenInfo struct {
	token     string
	expiresIn time.Duration
}

type token struct {
	mutex *sync.RWMutex
	tokenInfo
	lastRefresh  time.Time
	getTokenFunc func() (tokenInfo, error)
}

// getAccessToken 获取 access token
func (c *WorkwxApp) getAccessToken() (tokenInfo, error) {
	get, err := c.execGetAccessToken(reqAccessToken{
		CorpID:     c.CorpID,
		CorpSecret: c.CorpSecret,
	})
	if err != nil {
		return tokenInfo{}, err
	}
	return tokenInfo{token: get.AccessToken, expiresIn: time.Duration(get.ExpiresInSecs)}, nil
}

// SpawnAccessTokenRefresher 启动该 app 的 access token 刷新 goroutine
//
// NOTE: 该 goroutine 本身没有 keep-alive 逻辑，需要自助保活
func (c *WorkwxApp) SpawnAccessTokenRefresher() {
	ctx := context.Background()
	c.SpawnAccessTokenRefresherWithContext(ctx)
}

// SpawnAccessTokenRefresherWithContext 启动该 app 的 access token 刷新 goroutine
// 可以通过 context cancellation 停止此 goroutine
//
// NOTE: 该 goroutine 本身没有 keep-alive 逻辑，需要自助保活
func (c *WorkwxApp) SpawnAccessTokenRefresherWithContext(ctx context.Context) {
	go c.accessToken.tokenRefresher(ctx)
}

// GetJSAPITicket 获取 JSAPI_ticket
func (c *WorkwxApp) GetJSAPITicket() (string, error) {
	return c.jsapiTicket.getToken(), nil
}

// getJSAPITicket 获取 JSAPI_ticket
func (c *WorkwxApp) getJSAPITicket() (tokenInfo, error) {
	get, err := c.execGetJSAPITicket(reqJSAPITicket{})
	if err != nil {
		return tokenInfo{}, err
	}
	return tokenInfo{token: get.Ticket, expiresIn: time.Duration(get.ExpiresInSecs)}, nil
}

// SpawnJSAPITicketRefresher 启动该 app 的 JSAPI_ticket 刷新 goroutine
//
// NOTE: 该 goroutine 本身没有 keep-alive 逻辑，需要自助保活
func (c *WorkwxApp) SpawnJSAPITicketRefresher() {
	ctx := context.Background()
	c.SpawnJSAPITicketRefresherWithContext(ctx)
}

// SpawnJSAPITicketRefresherWithContext 启动该 app 的 JSAPI_ticket 刷新 goroutine
// 可以通过 context cancellation 停止此 goroutine
//
// NOTE: 该 goroutine 本身没有 keep-alive 逻辑，需要自助保活
func (c *WorkwxApp) SpawnJSAPITicketRefresherWithContext(ctx context.Context) {
	go c.jsapiTicket.tokenRefresher(ctx)
}

// GetJSAPITicketAgentConfig 获取 JSAPI_ticket_agent_config
func (c *WorkwxApp) GetJSAPITicketAgentConfig() (string, error) {
	return c.jsapiTicketAgentConfig.getToken(), nil
}

// getJSAPITicketAgentConfig 获取 JSAPI_ticket_agent_config
func (c *WorkwxApp) getJSAPITicketAgentConfig() (tokenInfo, error) {
	get, err := c.execGetJSAPITicketAgentConfig(reqJSAPITicketAgentConfig{})
	if err != nil {
		return tokenInfo{}, err
	}
	return tokenInfo{token: get.Ticket, expiresIn: time.Duration(get.ExpiresInSecs)}, nil
}

// SpawnJSAPITicketAgentConfigRefresher 启动该 app 的 JSAPI_ticket_agent_config 刷新 goroutine
//
// NOTE: 该 goroutine 本身没有 keep-alive 逻辑，需要自助保活
func (c *WorkwxApp) SpawnJSAPITicketAgentConfigRefresher() {
	ctx := context.Background()
	c.SpawnJSAPITicketAgentConfigRefresherWithContext(ctx)
}

// SpawnJSAPITicketAgentConfigRefresherWithContext 启动该 app 的 JSAPI_ticket_agent_config 刷新 goroutine
// 可以通过 context cancellation 停止此 goroutine
//
// NOTE: 该 goroutine 本身没有 keep-alive 逻辑，需要自助保活
func (c *WorkwxApp) SpawnJSAPITicketAgentConfigRefresherWithContext(ctx context.Context) {
	go c.jsapiTicketAgentConfig.tokenRefresher(ctx)
}

func (t *token) setGetTokenFunc(f func() (tokenInfo, error)) {
	t.getTokenFunc = f
}

func (t *token) getToken() string {
	// intensive mutex juggling action
	t.mutex.RLock()
	if t.token == "" {
		t.mutex.RUnlock() // RWMutex doesn't like recursive locking
		// TODO: what to do with the possible error?
		_ = t.syncToken()
		t.mutex.RLock()
	}
	tokenToUse := t.token
	t.mutex.RUnlock()
	return tokenToUse
}

func (t *token) syncToken() error {
	get, err := t.getTokenFunc()
	if err != nil {
		return err
	}
	t.mutex.Lock()
	defer t.mutex.Unlock()

	t.token = get.token
	t.expiresIn = get.expiresIn * time.Second
	t.lastRefresh = time.Now()
	return nil
}

func (t *token) tokenRefresher(ctx context.Context) {
	const refreshTimeWindow = 30 * time.Minute
	const minRefreshDuration = 5 * time.Second

	var waitDuration time.Duration = 0
	for {
		select {
		case <-time.After(waitDuration):
			retryer := backoff.WithContext(backoff.NewExponentialBackOff(), ctx)
			if err := backoff.Retry(t.syncToken, retryer); err != nil {
				// TODO: logging
				_ = err
			}

			waitUntilTime := t.lastRefresh.Add(t.expiresIn).Add(-refreshTimeWindow)
			waitDuration = time.Until(waitUntilTime)
			if waitDuration < minRefreshDuration {
				waitDuration = minRefreshDuration
			}
		case <-ctx.Done():
			return
		}
	}
}

// JSCode2Session 临时登录凭证校验
func (c *WorkwxApp) JSCode2Session(jscode string) (*JSCodeSession, error) {
	resp, err := c.execJSCode2Session(reqJSCode2Session{JSCode: jscode})
	if err != nil {
		return nil, err
	}
	return &resp.JSCodeSession, nil
}
