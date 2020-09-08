package workwx

import (
	"context"
	"time"

	"github.com/cenkalti/backoff/v4"
)

type tokenType string

const accessToken = tokenType("accessToken")
const jsApiTicket = tokenType("jsApiTicket")
const jsApiTicketAgentConfig = tokenType("jsApiTicketAgentConfig")

// getAccessToken 获取 access token
func (c *WorkwxApp) getAccessToken() (respAccessToken, error) {
	return c.execGetAccessToken(reqAccessToken{
		CorpID:     c.CorpID,
		CorpSecret: c.CorpSecret,
	})
}

// syncAccessToken 同步该客户端实例的 access token
//
// 会拿 `mutex` 写锁
func (c *WorkwxApp) syncAccessToken() error {
	return c.syncToken(accessToken)
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
	go c.tokenRefresher(ctx, accessToken)
}

// GetJsApiTicket 获取 jsapi_ticket
func (c *WorkwxApp) GetJsApiTicket() (string, error) {
	return c.getToken(jsApiTicket), nil
}

// getJsApiTicket 获取 jsapi_ticket
func (c *WorkwxApp) getJsApiTicket() (respJsApiTicket, error) {
	return c.execGetJsApiTicket(reqJsApiTicket{})
}

// syncJsApiTicket 同步该客户端实例的 jsapi_ticket
//
// 会拿 `mutex` 写锁
func (c *WorkwxApp) syncJsApiTicket() error {
	return c.syncToken(jsApiTicket)
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
	go c.tokenRefresher(ctx, jsApiTicket)
}

// GetJsApiTicketAgentConfig 获取 jsapi_ticket_agent_config
func (c *WorkwxApp) GetJsApiTicketAgentConfig() (string, error) {
	return c.getToken(jsApiTicketAgentConfig), nil
}

// getJsApiTicketAgentConfig 获取 jsapi_ticket_agent_config
func (c *WorkwxApp) getJsApiTicketAgentConfig() (respJsApiTicket, error) {
	return c.execGetJsApiTicketAgentConfig(reqJsApiTicketAgentConfig{})
}

// syncJsApiTicket 同步该客户端实例的 jsapi_ticket
//
// 会拿 `mutex` 写锁
func (c *WorkwxApp) syncJsApiTicketAgentConfig() error {
	return c.syncToken(jsApiTicketAgentConfig)
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
	go c.tokenRefresher(ctx, jsApiTicketAgentConfig)
}

func (c *WorkwxApp) getToken(tokenType tokenType) string {
	token := &token{}
	switch tokenType {
	case accessToken:
		token = c.accessToken
	case jsApiTicket:
		token = c.jsApiTicket
	case jsApiTicketAgentConfig:
		token = c.jsApiTicketAgentConfig
	}
	// intensive mutex juggling action
	token.mutex.RLock()
	if token.value == "" {
		token.mutex.RUnlock() // RWMutex doesn't like recursive locking
		// TODO: what to do with the possible error?
		_ = c.syncToken(tokenType)
		token.mutex.RLock()
	}
	tokenToUse := token.value
	token.mutex.RUnlock()
	return tokenToUse
}

func (c *WorkwxApp) syncToken(tokenType tokenType) error {
	var (
		token            = &token{}
		tok              string
		tokExpiresInSecs int64
	)
	switch tokenType {
	case accessToken:
		getAccessToken, err := c.getAccessToken()
		if err != nil {
			return err
		}
		tok = getAccessToken.AccessToken
		tokExpiresInSecs = getAccessToken.ExpiresInSecs
		token = c.accessToken
	case jsApiTicket:
		getJsApiTicket, err := c.getJsApiTicket()
		if err != nil {
			return err
		}
		tok = getJsApiTicket.Ticket
		tokExpiresInSecs = getJsApiTicket.ExpiresInSecs
		token = c.jsApiTicket
	case jsApiTicketAgentConfig:
		getJsApiTicketAgentConfig, err := c.getJsApiTicketAgentConfig()
		if err != nil {
			return err
		}
		tok = getJsApiTicketAgentConfig.Ticket
		tokExpiresInSecs = getJsApiTicketAgentConfig.ExpiresInSecs
		token = c.jsApiTicketAgentConfig
	}

	token.mutex.Lock()
	defer token.mutex.Unlock()

	token.value = tok
	token.expiresIn = time.Duration(tokExpiresInSecs) * time.Second
	token.lastRefresh = time.Now()
	return nil
}

func (c *WorkwxApp) tokenRefresher(ctx context.Context, tokenType tokenType) {
	const refreshTimeWindow = 30 * time.Minute
	const minRefreshDuration = 5 * time.Second

	var waitDuration time.Duration = 0
	for {
		select {
		case <-time.After(waitDuration):
			retryer := backoff.WithContext(backoff.NewExponentialBackOff(), ctx)
			switch tokenType {
			case accessToken:
				if err := backoff.Retry(c.syncAccessToken, retryer); err != nil {
					// TODO: logging
					_ = err
				}

				waitUntilTime := c.accessToken.lastRefresh.Add(c.accessToken.expiresIn).Add(-refreshTimeWindow)
				waitDuration = time.Until(waitUntilTime)
			case jsApiTicket:
				if err := backoff.Retry(c.syncJsApiTicket, retryer); err != nil {
					// TODO: logging
					_ = err
				}

				waitUntilTime := c.jsApiTicket.lastRefresh.Add(c.jsApiTicket.expiresIn).Add(-refreshTimeWindow)
				waitDuration = time.Until(waitUntilTime)
			case jsApiTicketAgentConfig:
				if err := backoff.Retry(c.syncJsApiTicketAgentConfig, retryer); err != nil {
					// TODO: logging
					_ = err
				}

				waitUntilTime := c.jsApiTicket.lastRefresh.Add(c.jsApiTicket.expiresIn).Add(-refreshTimeWindow)
				waitDuration = time.Until(waitUntilTime)
			}
			if waitDuration < minRefreshDuration {
				waitDuration = minRefreshDuration
			}
		case <-ctx.Done():
			return
		}
	}
}
