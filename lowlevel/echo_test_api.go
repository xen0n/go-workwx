package lowlevel

import (
	"errors"
	"net/url"
	"strconv"
)

type EchoTestAPIArgsAdapter interface {
	ParseEchoTestAPIArgs() (EchoTestAPIArgs, error)
}

type EchoTestAPIArgs struct {
	MsgSignature string
	Timestamp    int64
	Nonce        string
	EchoStr      string
}

type URLValuesEchoTestAdapter url.Values

var _ EchoTestAPIArgsAdapter = URLValuesEchoTestAdapter{}

var errMalformedArgs = errors.New("malformed arguments for echo test API")

func (x URLValuesEchoTestAdapter) ParseEchoTestAPIArgs() (EchoTestAPIArgs, error) {
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
