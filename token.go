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

// syncAccessToken 同步该客户端实例的 access token
//
// 会拿 `mutex` 写锁
func (c *WorkwxApp) syncAccessToken() error {
	return c.accessToken.syncToken()
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

// GetJsApiTicket 获取 jsapi_ticket
func (c *WorkwxApp) GetJsApiTicket() (string, error) {
	return c.jsApiTicket.getToken(), nil
}

// getJsApiTicket 获取 jsapi_ticket
func (c *WorkwxApp) getJsApiTicket() (tokenInfo, error) {
	get, err := c.execGetJsApiTicket(reqJsApiTicket{})
	if err != nil {
		return tokenInfo{}, err
	}
	return tokenInfo{token: get.Ticket, expiresIn: time.Duration(get.ExpiresInSecs)}, nil
}

// syncJsApiTicket 同步该客户端实例的 jsapi_ticket
//
// 会拿 `mutex` 写锁
func (c *WorkwxApp) syncJsApiTicket() error {
	return c.jsApiTicket.syncToken()
}

// SpawnJsApiTicketRefresher 启动该 app 的 jsapi_ticket 刷新 goroutine
//
// NOTE: 该 goroutine 本身没有 keep-alive 逻辑，需要自助保活
func (c *WorkwxApp) SpawnJsApiTicketRefresher() {
	ctx := context.Background()
	c.SpawnJsApiTicketRefresherWithContext(ctx)
}

// SpawnJsApiTicketRefresherWithContext 启动该 app 的 jsapi_ticket 刷新 goroutine
// 可以通过 context cancellation 停止此 goroutine
//
// NOTE: 该 goroutine 本身没有 keep-alive 逻辑，需要自助保活
func (c *WorkwxApp) SpawnJsApiTicketRefresherWithContext(ctx context.Context) {
	go c.jsApiTicket.tokenRefresher(ctx)
}

// GetJsApiTicketAgentConfig 获取 jsapi_ticket_agent_config
func (c *WorkwxApp) GetJsApiTicketAgentConfig() (string, error) {
	return c.jsApiTicketAgentConfig.getToken(), nil
}

// getJsApiTicketAgentConfig 获取 jsapi_ticket_agent_config
func (c *WorkwxApp) getJsApiTicketAgentConfig() (tokenInfo, error) {
	get, err := c.execGetJsApiTicketAgentConfig(reqJsApiTicketAgentConfig{})
	if err != nil {
		return tokenInfo{}, err
	}
	return tokenInfo{token: get.Ticket, expiresIn: time.Duration(get.ExpiresInSecs)}, nil
}

// syncJsApiTicket 同步该客户端实例的 jsapi_ticket
//
// 会拿 `mutex` 写锁
func (c *WorkwxApp) syncJsApiTicketAgentConfig() error {
	return c.jsApiTicketAgentConfig.syncToken()
}

// SpawnJsApiTicketAgentConfigRefresher 启动该 app 的 jsapi_ticket_agent_config 刷新 goroutine
//
// NOTE: 该 goroutine 本身没有 keep-alive 逻辑，需要自助保活
func (c *WorkwxApp) SpawnJsApiTicketAgentConfigRefresher() {
	ctx := context.Background()
	c.SpawnJsApiTicketAgentConfigRefresherWithContext(ctx)
}

// SpawnJsApiTicketAgentConfigRefresherWithContext 启动该 app 的 jsapi_ticket_agent_config 刷新 goroutine
// 可以通过 context cancellation 停止此 goroutine
//
// NOTE: 该 goroutine 本身没有 keep-alive 逻辑，需要自助保活
func (c *WorkwxApp) SpawnJsApiTicketAgentConfigRefresherWithContext(ctx context.Context) {
	go c.jsApiTicketAgentConfig.tokenRefresher(ctx)
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
