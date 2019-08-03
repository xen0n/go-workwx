package workwx

import (
	"net/http"
)

// DefaultQYAPIHost 默认企业微信 API Host
const DefaultQYAPIHost = "https://qyapi.weixin.qq.com"

type options struct {
	QYAPIHost string
	HTTP      *http.Client
}

type ctorOption interface {
	ApplyTo(*options)
}

// impl Default for options
func defaultOptions() options {
	return options{
		QYAPIHost: DefaultQYAPIHost,
		HTTP:      &http.Client{},
	}
}

//
//
//

type withQYAPIHost struct {
	x string
}

// WithQYAPIHost 覆盖默认企业微信 API 域名
func WithQYAPIHost(host string) ctorOption {
	return &withQYAPIHost{x: host}
}

var _ ctorOption = (*withQYAPIHost)(nil)

func (x *withQYAPIHost) ApplyTo(y *options) {
	y.QYAPIHost = x.x
}

//
//
//

type withHTTPClient struct {
	x *http.Client
}

// WithHTTPClient 使用给定的 http.Client 作为 HTTP 客户端
func WithHTTPClient(client *http.Client) ctorOption {
	return &withHTTPClient{x: client}
}

var _ ctorOption = (*withHTTPClient)(nil)

func (x *withHTTPClient) ApplyTo(y *options) {
	y.HTTP = x.x
}
