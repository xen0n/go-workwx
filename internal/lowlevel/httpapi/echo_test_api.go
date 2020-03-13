package httpapi

import (
	"errors"
	"net/http"
	"net/url"
	"strconv"

	"github.com/xen0n/go-workwx/internal/lowlevel/encryptor"
	"github.com/xen0n/go-workwx/internal/lowlevel/signature"
)

type ToEchoTestAPIArgs interface {
	ToEchoTestAPIArgs() (EchoTestAPIArgs, error)
}

type EchoTestAPIArgs struct {
	MsgSignature string
	Timestamp    int64
	Nonce        string
	EchoStr      string
}

type URLValuesForEchoTestAPI url.Values

var _ ToEchoTestAPIArgs = URLValuesForEchoTestAPI{}

var errMalformedArgs = errors.New("malformed arguments for echo test API")

func (x URLValuesForEchoTestAPI) ToEchoTestAPIArgs() (EchoTestAPIArgs, error) {
	var msgSignature string
	{
		l := x["msg_signature"]
		if len(l) != 1 {
			return EchoTestAPIArgs{}, errMalformedArgs
		}
		msgSignature = l[0]
	}

	var timestamp int64
	{
		l := x["timestamp"]
		if len(l) != 1 {
			return EchoTestAPIArgs{}, errMalformedArgs
		}
		timestampStr := l[0]

		timestampInt, err := strconv.ParseInt(timestampStr, 10, 64)
		if err != nil {
			return EchoTestAPIArgs{}, errMalformedArgs
		}

		timestamp = timestampInt
	}

	var nonce string
	{
		l := x["nonce"]
		if len(l) != 1 {
			return EchoTestAPIArgs{}, errMalformedArgs
		}
		nonce = l[0]
	}

	var echoStr string
	{
		l := x["echostr"]
		if len(l) != 1 {
			return EchoTestAPIArgs{}, errMalformedArgs
		}
		echoStr = l[0]
	}

	return EchoTestAPIArgs{
		MsgSignature: msgSignature,
		Timestamp:    timestamp,
		Nonce:        nonce,
		EchoStr:      echoStr,
	}, nil
}

type HTTPEchoTestAPIHandler struct {
	token     string
	encryptor *encryptor.WorkwxEncryptor
}

var _ http.Handler = (*HTTPEchoTestAPIHandler)(nil)

func NewHTTPEchoTestAPIHandler(
	token string,
	encodingAESKey string,
) (*HTTPEchoTestAPIHandler, error) {
	enc, err := encryptor.NewWorkwxEncryptor(encodingAESKey)
	if err != nil {
		return nil, err
	}

	return &HTTPEchoTestAPIHandler{
		token:     token,
		encryptor: enc,
	}, nil
}

func (h *HTTPEchoTestAPIHandler) ServeHTTP(
	wr http.ResponseWriter,
	r *http.Request,
) {
	if !signature.VerifyHTTPRequestSignature(h.token, r.URL, "") {
		wr.WriteHeader(http.StatusBadRequest)
		return
	}

	adapter := URLValuesForEchoTestAPI(r.URL.Query())
	args, err := adapter.ToEchoTestAPIArgs()
	if err != nil {
		wr.WriteHeader(http.StatusBadRequest)
		return
	}

	payload, err := h.encryptor.Decrypt([]byte(args.EchoStr))
	if err != nil {
		wr.WriteHeader(http.StatusBadRequest)
		return
	}

	wr.WriteHeader(http.StatusOK)
	wr.Write(payload.Msg)
}