package lowlevel

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/binary"
	"errors"
)

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
