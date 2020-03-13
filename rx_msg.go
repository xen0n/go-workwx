package workwx

import (
	"encoding/xml"
	"fmt"
	"strings"
	"time"
)

// RxMessage 一条接收到的消息
type RxMessage struct {
	// FromUserID 发送者的 UserID
	FromUserID string
	// SendTime 消息发送时间
	SendTime time.Time
	// MsgType 消息类型
	MsgType MessageType
	// MsgID 消息 ID
	MsgID int64
	// AgentID 企业应用 ID，可在应用的设置页面查看
	AgentID int64

	extras messageKind
}

func fromEnvelope(body []byte) (*RxMessage, error) {
	// extract common part
	var common rxMessageCommon
	err := xml.Unmarshal(body, &common)
	if err != nil {
		return nil, err
	}

	// deal with polymorphic message types
	extras, err := extractMessageExtras(common.MsgType, body)
	if err != nil {
		return nil, err
	}

	// assemble message object
	var obj RxMessage
	{
		// let's force people to think about timezones okay?
		// -- let's not
		sendTime := time.Unix(common.CreateTime, 0) // in time.Local

		obj = RxMessage{
			FromUserID: common.FromUserName,
			SendTime:   sendTime,
			MsgType:    common.MsgType,
			MsgID:      common.MsgID,
			AgentID:    common.AgentID,

			extras: extras,
		}
	}

	return &obj, nil
}

func (m *RxMessage) String() string {
	var sb strings.Builder

	_, _ = fmt.Fprintf(
		&sb,
		"RxMessage { FromUserID: %#v, SendTime: %d, MsgType: %#v, MsgID: %d, AgentID: %d, ",
		m.FromUserID,
		m.SendTime.UnixNano(),
		m.MsgType,
		m.MsgID,
		m.AgentID,
	)

	m.extras.formatInto(&sb)

	sb.WriteString(" }")

	return sb.String()
}

// Text 如果消息为文本类型，则拿出相应的消息参数，否则返回 nil, false
func (m *RxMessage) Text() (TextMessageExtras, bool) {
	y, ok := m.extras.(TextMessageExtras)
	return y, ok
}

// Image 如果消息为图片类型，则拿出相应的消息参数，否则返回 nil, false
func (m *RxMessage) Image() (ImageMessageExtras, bool) {
	y, ok := m.extras.(ImageMessageExtras)
	return y, ok
}

// Voice 如果消息为语音类型，则拿出相应的消息参数，否则返回 nil, false
func (m *RxMessage) Voice() (VoiceMessageExtras, bool) {
	y, ok := m.extras.(VoiceMessageExtras)
	return y, ok
}

// Video 如果消息为视频类型，则拿出相应的消息参数，否则返回 nil, false
func (m *RxMessage) Video() (VideoMessageExtras, bool) {
	y, ok := m.extras.(VideoMessageExtras)
	return y, ok
}

// Location 如果消息为位置类型，则拿出相应的消息参数，否则返回 nil, false
func (m *RxMessage) Location() (LocationMessageExtras, bool) {
	y, ok := m.extras.(LocationMessageExtras)
	return y, ok
}

// Link 如果消息为链接类型，则拿出相应的消息参数，否则返回 nil, false
func (m *RxMessage) Link() (LinkMessageExtras, bool) {
	y, ok := m.extras.(LinkMessageExtras)
	return y, ok
}
