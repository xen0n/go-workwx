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

func doEchoTest(
	url *url.URL,
	token string,
	encryptor *encryptor.WorkwxEncryptor,
) (statusCode int, body []byte) {
	if !signature.VerifyHTTPRequestSignature(token, url, "") {
		return http.StatusBadRequest, nil
	}

	adapter := URLValuesForEchoTestAPI(url.Query())
	args, err := adapter.ToEchoTestAPIArgs()
	if err != nil {
		return http.StatusBadRequest, nil
	}

	payload, err := encryptor.Decrypt([]byte(args.EchoStr))
	if err != nil {
		return http.StatusBadRequest, nil
	}

	return http.StatusOK, payload.Msg
}
