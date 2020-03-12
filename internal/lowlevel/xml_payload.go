package lowlevel

import (
	"encoding/xml"
	"errors"
	"net/url"
)

type xmlRxPayload struct {
	ToUserName string `xml:"ToUserName"`
	AgentID    string `xml:"AgentID"`
	Encrypt    string `xml:"Encrypt"`
}

type PayloadProcessor struct {
	token     string
	encryptor *WorkwxEncryptor
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
		token:     token,
		encryptor: enc,
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
