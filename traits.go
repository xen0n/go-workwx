package workwx

import (
	"net/url"
)

// urlValuer 可转化为 url.Values 类型的 trait
type urlValuer interface {
	// IntoURLValues 转换为 url.Values 类型
	IntoURLValues() url.Values
}

// bodyer 可转化为 API 请求体的 trait
type bodyer interface {
	// IntoBody 转换为请求体的 []byte 类型
	IntoBody() ([]byte, error)
}

// mediaUploader 携带 *Media 对象，可转化为 multipart 文件上传请求体的 trait
type mediaUploader interface {
	getMedia() *Media
}
