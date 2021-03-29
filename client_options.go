package workwx

import (
	"net/http"
)

// DefaultQYAPIHost 默认企业微信 API Host
const DefaultQYAPIHost = "https://qyapi.weixin.qq.com"

type options struct {
	QYAPIHost string
	HTTP      *http.Client
	Logger    Logger
}

// CtorOption 客户端对象构造参数
type CtorOption interface {
	applyTo(*options)
}

// impl Default for options
func defaultOptions() options {
	return options{
		QYAPIHost: DefaultQYAPIHost,
		HTTP:      &http.Client{},
		Logger:    newDefaultLogger(),
	}
}

//
//
//

type withQYAPIHost struct {
	x string
}

// WithQYAPIHost 覆盖默认企业微信 API 域名
func WithQYAPIHost(host string) CtorOption {
	return &withQYAPIHost{x: host}
}

var _ CtorOption = (*withQYAPIHost)(nil)

func (x *withQYAPIHost) applyTo(y *options) {
	y.QYAPIHost = x.x
}

//
//
//

type withHTTPClient struct {
	x *http.Client
}

// WithHTTPClient 使用给定的 http.Client 作为 HTTP 客户端
func WithHTTPClient(client *http.Client) CtorOption {
	return &withHTTPClient{x: client}
}

var _ CtorOption = (*withHTTPClient)(nil)

func (x *withHTTPClient) applyTo(y *options) {
	y.HTTP = x.x
}

//
//
//

type withLogger struct {
	x Logger
}

// WithLogger 使用给定的 logger 作为日志打印的方式。如果不使用 WithLogger 自定义 logger，
// 则默认 logger 会将 Info 级别的日志打印在 stdout 中，将 Error 级别的日志打印在 stderr 中。
func WithLogger(logger Logger) CtorOption {
	return &withLogger{x: logger}
}

var _ CtorOption = (*withLogger)(nil)

func (x *withLogger) applyTo(y *options) {
	y.Logger = x.x
}
