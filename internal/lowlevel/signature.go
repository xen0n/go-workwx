package lowlevel

import (
	"crypto/sha1"
	"crypto/subtle"
	"fmt"
	"net/url"
	"sort"
)

func makeDevMsgSignature(paramValues ...string) string {
	tmp := make([]string, len(paramValues))
	copy(tmp, paramValues)

	sort.Strings(tmp)

	state := sha1.New()
	for _, x := range tmp {
		_, _ = state.Write([]byte(x))
	}

	result := state.Sum(nil)
	return fmt.Sprintf("%x", result)
}

// ToMsgSignature 适配企业微信请求参数签名的 interface
type ToMsgSignature interface {
	// GetMsgSignature 取请求上携带的签名串
	GetMsgSignature() (string, bool)
	// GetParamValues 取所有请求参数值（不必有序）
	GetParamValues() ([]string, bool)
}

// VerifySignature 校验一个 ToMsgSignature 的签名是否完好
//
// NOTE: Go 没有 default method for interface，因此无法以 `foo.VerifySignature()`
// 的形式实现。
func VerifySignature(token string, x ToMsgSignature) bool {
	msgSignature, ok := x.GetMsgSignature()
	if !ok {
		return false
	}

	paramValues, ok := x.GetParamValues()
	if !ok {
		return false
	}

	devMsgSignature := makeDevMsgSignature(append(paramValues, token)...)
	eq := subtle.ConstantTimeCompare([]byte(msgSignature), []byte(devMsgSignature))
	return eq != 0
}

// VerifyURLSignature 校验一个 URL 的签名是否完好
//
// 这是 VerifySignature 的简单包装。
func VerifyURLSignature(token string, x *url.URL) bool {
	// XXX seems this is a memcpy...
	wrapped := URLWithSignature(*x)
	return VerifySignature(token, &wrapped)
}

// URLWithSignature 为 url.URL 类型适配签名校验逻辑
type URLWithSignature url.URL

var _ ToMsgSignature = (*URLWithSignature)(nil)

// GetMsgSignature 取请求上携带的签名串
func (u *URLWithSignature) GetMsgSignature() (string, bool) {
	inner := (*url.URL)(u)
	l := inner.Query()["msg_signature"]
	if len(l) != 1 {
		return "", false
	}

	return l[0], true
}

// GetParamValues 取所有请求参数值（不必有序）
func (u *URLWithSignature) GetParamValues() ([]string, bool) {
	inner := (*url.URL)(u)
	result := make([]string, 0)
	for k, l := range inner.Query() {
		if k == "msg_signature" {
			continue
		}
		result = append(result, l...)
	}
	return result, true
}
