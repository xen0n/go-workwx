package lowlevel

import (
	"crypto/rand"
	"encoding/xml"
	"errors"
	"io"
	"math/big"
	"net/url"
	"strconv"
	"time"
)

type xmlRxPayload struct {
	ToUserName string `xml:"ToUserName"`
	AgentID    string `xml:"AgentID"`
	Encrypt    string `xml:"Encrypt"`
}

type cdataNode struct {
	CData string `xml:",cdata"`
}

type xmlTxPayload struct {
	XMLName      xml.Name  `xml:"xml"`
	Encrypt      cdataNode `xml:"Encrypt"`
	MsgSignature cdataNode `xml:"MsgSignature"`
	Timestamp    int64     `xml:"Timestamp"`
	Nonce        cdataNode `xml:"Nonce"`
}

type TimeSource interface {
	GetCurrentTimestamp() time.Time
}

type DefaultTimeSource struct{}

var _ TimeSource = DefaultTimeSource{}

func (DefaultTimeSource) GetCurrentTimestamp() time.Time {
	return time.Now()
}

type PayloadProcessor struct {
	token         string
	encryptor     *WorkwxEncryptor
	entropySource io.Reader
	timeSource    TimeSource
}

type MessagePayload struct {
	ToUserName string
	AgentID    string
	Msg        []byte
	ReceiveID  []byte
}

func NewPayloadProcessor(
	token string,
	encodingAESKey string,
) (*PayloadProcessor, error) {
	enc, err := NewWorkwxEncryptor(encodingAESKey)
	if err != nil {
		return nil, err
	}

	return &PayloadProcessor{
		token:         token,
		encryptor:     enc,
		entropySource: rand.Reader,         // TODO: support customization
		timeSource:    DefaultTimeSource{}, // TODO: ditto
	}, nil
}

var errInvalidSignature = errors.New("invalid signature")

func (p *PayloadProcessor) HandleIncomingMsg(
	url *url.URL,
	body []byte,
) (MessagePayload, error) {
	// xml unmarshal
	var x xmlRxPayload
	err := xml.Unmarshal(body, &x)
	if err != nil {
		return MessagePayload{}, err
	}

	// check signature
	if !VerifyHTTPRequestSignature(p.token, url, x.Encrypt) {
		return MessagePayload{}, errInvalidSignature
	}

	// decrypt message
	msg, err := p.encryptor.Decrypt([]byte(x.Encrypt))
	if err != nil {
		return MessagePayload{}, err
	}

	// assemble payload to return
	return MessagePayload{
		ToUserName: x.ToUserName,
		AgentID:    x.AgentID,
		Msg:        msg.Msg,
		ReceiveID:  msg.ReceiveID,
	}, nil
}

func (p *PayloadProcessor) MakeOutgoingMessage(msg []byte) ([]byte, error) {
	workwxPayload := WorkwxPayload{
		Msg:       msg,
		ReceiveID: nil,
	}
	encryptedMsg, err := p.encryptor.Encrypt(&workwxPayload)
	if err != nil {
		return nil, err
	}

	ts := p.timeSource.GetCurrentTimestamp().Unix()
	nonce, err := makeNonce(p.entropySource)
	if err != nil {
		return nil, err
	}

	msgSignature := makeDevMsgSignature(
		p.token,
		strconv.FormatInt(ts, 10),
		nonce,
		encryptedMsg,
	)

	payload := xmlTxPayload{
		XMLName: xml.Name{},
		Encrypt: cdataNode{
			CData: encryptedMsg,
		},
		MsgSignature: cdataNode{
			CData: msgSignature,
		},
		Timestamp: ts,
		Nonce: cdataNode{
			CData: nonce,
		},
	}

	result, err := xml.Marshal(payload)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func makeNonce(entropySource io.Reader) (string, error) {
	limit := big.NewInt(1)
	limit = limit.Lsh(limit, 64)
	n, err := rand.Int(entropySource, limit)
	if err != nil {
		return "", err
	}
	return n.String(), nil
}
