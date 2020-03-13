package envelope

import (
	"crypto/rand"
	"encoding/xml"
	"errors"
	"io"
	"math/big"
	"net/url"
	"strconv"

	"github.com/xen0n/go-workwx/internal/lowlevel/encryptor"
	"github.com/xen0n/go-workwx/internal/lowlevel/signature"
)

type EnvelopeProcessor struct {
	token         string
	encryptor     *encryptor.WorkwxEncryptor
	entropySource io.Reader
	timeSource    TimeSource
}

func NewEnvelopeProcessor(
	token string,
	encodingAESKey string,
) (*EnvelopeProcessor, error) {
	enc, err := encryptor.NewWorkwxEncryptor(encodingAESKey)
	if err != nil {
		return nil, err
	}

	return &EnvelopeProcessor{
		token:         token,
		encryptor:     enc,
		entropySource: rand.Reader,         // TODO: support customization
		timeSource:    DefaultTimeSource{}, // TODO: ditto
	}, nil
}

var errInvalidSignature = errors.New("invalid signature")

func (p *EnvelopeProcessor) HandleIncomingMsg(
	url *url.URL,
	body []byte,
) (Envelope, error) {
	// xml unmarshal
	var x xmlRxEnvelope
	err := xml.Unmarshal(body, &x)
	if err != nil {
		return Envelope{}, err
	}

	// check signature
	if !signature.VerifyHTTPRequestSignature(p.token, url, x.Encrypt) {
		return Envelope{}, errInvalidSignature
	}

	// decrypt message
	msg, err := p.encryptor.Decrypt([]byte(x.Encrypt))
	if err != nil {
		return Envelope{}, err
	}

	// assemble envelope to return
	return Envelope{
		ToUserName: x.ToUserName,
		AgentID:    x.AgentID,
		Msg:        msg.Msg,
		ReceiveID:  msg.ReceiveID,
	}, nil
}

func (p *EnvelopeProcessor) MakeOutgoingEnvelope(msg []byte) ([]byte, error) {
	workwxPayload := encryptor.WorkwxPayload{
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

	msgSignature := signature.MakeDevMsgSignature(
		p.token,
		strconv.FormatInt(ts, 10),
		nonce,
		encryptedMsg,
	)

	envelope := xmlTxEnvelope{
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

	result, err := xml.Marshal(envelope)
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
