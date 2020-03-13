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
	rw http.ResponseWriter,
	r *http.Request,
) {
	switch r.Method {
	case http.MethodGet:
		// 测试回调模式请求
		h.echoTestHandler(rw, r)

	case http.MethodPost:
		// TODO
		rw.WriteHeader(http.StatusNotImplemented)

	default:
		// unhandled request method
		rw.WriteHeader(http.StatusNotImplemented)
	}
}

func (h *LowlevelHandler) echoTestHandler(
	rw http.ResponseWriter,
	r *http.Request,
) {
	statusCode, body := doEchoTest(r.URL, h.token, h.encryptor)
	rw.WriteHeader(statusCode)
	if body != nil {
		rw.Write(body)
	}
}
