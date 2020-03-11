package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"
	"net/url"
	"sort"
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

func makeDevMsgSignature(paramValues ...string) string {
	tmp := make([]string, len(paramValues))
	for i, x := range paramValues {
		tmp[i] = x
	}

	sort.Strings(tmp)

	state := sha1.New()
	for _, x := range tmp {
		_, _ = state.Write([]byte(x))
	}

	result := state.Sum(nil)
	return fmt.Sprintf("%x", result)
}

type WorkwxPayload struct {
	Msg       []byte
	ReceiveID []byte
}

type WorkwxEncryptor struct {
	aesKey []byte
}

var errMalformedEncodingAESKey = errors.New("malformed EncodingAESKey")

func NewWorkwxEncryptor(encodingAESKey string) (*WorkwxEncryptor, error) {
	aesKey, err := base64.StdEncoding.DecodeString(encodingAESKey + "=")
	if err != nil {
		return nil, err
	}

	if len(aesKey) != 32 {
		return nil, errMalformedEncodingAESKey
	}

	return &WorkwxEncryptor{
		aesKey: aesKey,
	}, nil
}

func pkcs7Unpad(x []byte) []byte {
	// last byte is number of suffix bytes to remove
	n := int(x[len(x)-1])
	return x[:len(x)-n]
}

func (e *WorkwxEncryptor) Decrypt(base64Msg []byte) (WorkwxPayload, error) {
	// base64 decode
	buflen := base64.StdEncoding.DecodedLen(len(base64Msg))
	buf := make([]byte, buflen)
	n, err := base64.StdEncoding.Decode(buf, base64Msg)
	if err != nil {
		return WorkwxPayload{}, err
	}
	buf = buf[:n]

	// init cipher
	block, err := aes.NewCipher(e.aesKey)
	if err != nil {
		return WorkwxPayload{}, err
	}

	iv := e.aesKey[:16]
	state := cipher.NewCBCDecrypter(block, iv)

	// decrypt in-place in the allocated temp buffer
	state.CryptBlocks(buf, buf)
	buf = pkcs7Unpad(buf)

	// assemble decrypted payload
	// drop the 16-byte random prefix
	msglen := binary.BigEndian.Uint32(buf[16:20])
	msg := buf[20 : 20+msglen]
	receiveID := buf[20+msglen:]

	return WorkwxPayload{
		Msg:       msg,
		ReceiveID: receiveID,
	}, nil
}

func main() {
	s := "http://test.example.com/?echostr=6KmUQuPVu7UhjyVqRdbo5SfcRqaHvbUlKSHFvBV2ZuR6TIlKsygcfeSd1GDplg1C5KSKr6UPHCaC%2FnIX3ZNt9w%3D%3D&msg_signature=1ba3cb09c0d2c2b3ed6900d37f91a6efae6cb011&timestamp=1583940690&nonce=VHh7ymSeb0jc4lSb"
	token := "kjr2TKI8umCBfVF3wAHk8JiPwma5VBme"
	encodingaeskey := "4Ma3YBrSBbX2aez8MJpXGBne5LSDwgGqHbhM9WPYIws"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	a := URLValuesEchoTestAdapter(u.Query())

	p, err := a.ParseEchoTestAPIArgs()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", p)

	sign := makeDevMsgSignature(token, p.EchoStr, p.Nonce, strconv.FormatInt(p.Timestamp, 10))
	fmt.Printf("sign=%s\n", sign)

	enc, err := NewWorkwxEncryptor(encodingaeskey)
	if err != nil {
		panic(err)
	}

	x, err := enc.Decrypt([]byte(p.EchoStr))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n%s\n%s\n", x, string(x.Msg), string(x.ReceiveID))
}
