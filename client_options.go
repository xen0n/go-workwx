package workwx

import (
	"net/http"
)

type options struct {
	QYAPIHost string
	HTTP      *http.Client
}

type ctorOption interface {
	ApplyTo(*options)
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

// impl ctorOption for withQYAPIHost
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

// impl ctorOption for withHTTPClient
func (x *withHTTPClient) ApplyTo(y *options) {
	y.HTTP = x.x
}
