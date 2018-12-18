package workwx

import (
	"net/url"
)

// urlValuer 可转化为 url.Values 类型的 marker trait
type urlValuer interface {
	// IntoURLValues 转换为 url.Values 类型
	IntoURLValues() url.Values
}
