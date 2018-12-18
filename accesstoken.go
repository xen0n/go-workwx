package workwx

import (
	"time"

	"github.com/cenkalti/backoff"
)

// getAccessToken 获取 access token
func (c *WorkwxApp) getAccessToken() (respAccessToken, error) {
	req := reqAccessToken{
		CorpID:     c.CorpID,
		CorpSecret: c.CorpSecret,
	}

	var resp respAccessToken
	err := c.executeQyapiGet("/cgi-bin/gettoken", req, &resp)
	if err != nil {
		// TODO: error_chain
		return respAccessToken{}, err
	}

	return resp, nil
}

// syncAccessToken 同步该客户端实例的 access token
//
// 会拿 `tokenMu` 写锁
func (c *WorkwxApp) syncAccessToken() error {
	tok, err := c.getAccessToken()
	if err != nil {
		// TODO: error_chain
		return err
	}

	c.tokenMu.Lock()
	defer c.tokenMu.Unlock()

	c.accessToken = tok.AccessToken
	c.tokenExpiresIn = time.Duration(tok.ExpiresInSecs) * time.Second
	c.lastRefresh = time.Now()

	return nil
}

func (c *WorkwxApp) accessTokenRefresher() {
	const refreshTimeWindow = 30 * time.Minute
	const minRefreshDuration = 5 * time.Second

	// TODO: context cancellation
	for {
		retryer := backoff.NewExponentialBackOff()
		err := backoff.Retry(c.syncAccessToken, retryer)
		if err != nil {
			// wtf
			// TODO: logging
			_ = err
		}

		waitUntilTime := c.lastRefresh.Add(c.tokenExpiresIn).Add(-refreshTimeWindow)
		waitDuration := waitUntilTime.Sub(time.Now())

		if waitDuration < minRefreshDuration {
			waitDuration = minRefreshDuration
		}

		time.Sleep(waitDuration)
	}
}

// SpawnAccessTokenRefresher 启动该 app 的 access token 刷新 goroutine
//
// NOTE: 该 goroutine 本身没有 keep-alive 逻辑，需要自助保活
func (c *WorkwxApp) SpawnAccessTokenRefresher() {
	go c.accessTokenRefresher()
}
