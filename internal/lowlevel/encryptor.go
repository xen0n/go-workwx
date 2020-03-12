package lowlevel

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"io"
)

type WorkwxPayload struct {
	Msg       []byte
	ReceiveID []byte
}

type WorkwxEncryptor struct {
	aesKey        []byte
	entropySource io.Reader
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
		aesKey:        aesKey,
		entropySource: rand.Reader, // TODO: allow customizing this
	}, nil
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

func (e *WorkwxEncryptor) prepareBufForEncryption(msg []byte) ([]byte, error) {
	// TODO: what about ReceiveID?
	resultMsgLen := 16 + 4 + len(msg)

	// allocate buffer
	buf := make([]byte, 16, resultMsgLen)

	// add random prefix
	_, err := io.ReadFull(e.entropySource, buf) // len(buf) == 16 at this moment
	if err != nil {
		return nil, err
	}

	buf = buf[:cap(buf)] // grow to full capacity
	binary.BigEndian.PutUint32(buf[16:], uint32(len(msg)))
	copy(buf[20:], msg)

	return pkcs7Pad(buf), nil
}

func (e *WorkwxEncryptor) Encrypt(msg []byte) (string, error) {
	buf, err := e.prepareBufForEncryption(msg)
	if err != nil {
		return "", err
	}

	// init cipher
	block, err := aes.NewCipher(e.aesKey)
	if err != nil {
		return "", err
	}

	iv := e.aesKey[:16]
	state := cipher.NewCBCEncrypter(block, iv)

	// encrypt in-place as we own the buffer
	state.CryptBlocks(buf, buf)

	return base64.StdEncoding.EncodeToString(buf), nil
}
