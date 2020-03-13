package httpapi

import (
	"net/http"

	"github.com/xen0n/go-workwx/internal/lowlevel/encryptor"
)

type LowlevelHandler struct {
	token     string
	encryptor *encryptor.WorkwxEncryptor
}

var _ http.Handler = (*LowlevelHandler)(nil)

func NewLowlevelHandler(
	token string,
	encodingAESKey string,
) (*LowlevelHandler, error) {
	enc, err := encryptor.NewWorkwxEncryptor(encodingAESKey)
	if err != nil {
		return nil, err
	}

	return &LowlevelHandler{
		token:     token,
		encryptor: enc,
	}, nil
}

func (h *LowlevelHandler) ServeHTTP(
	wr http.ResponseWriter,
	r *http.Request,
) {
	switch r.Method {
	case http.MethodGet:
		// 测试回调模式请求
		h.echoTestHandler(wr, r)

	case http.MethodPost:
		// TODO
		wr.WriteHeader(http.StatusNotImplemented)

	default:
		// unhandled request method
		wr.WriteHeader(http.StatusNotImplemented)
	}
}

func (h *LowlevelHandler) echoTestHandler(
	wr http.ResponseWriter,
	r *http.Request,
) {
	statusCode, body := doEchoTest(r.URL, h.token, h.encryptor)
	wr.WriteHeader(statusCode)
	if body != nil {
		wr.Write(body)
	}
}
