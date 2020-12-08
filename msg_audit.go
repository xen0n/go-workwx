package workwx

import (
	"time"
)

// MsgAuditAgreeStatus 会话中外部成员的同意状态
type MsgAuditAgreeStatus string

const (
	// MsgAuditAgreeStatusAgree 同意
	MsgAuditAgreeStatusAgree = "Agree"
	// MsgAuditAgreeStatusDisagree 不同意
	MsgAuditAgreeStatusDisagree = "Disagree"
	// MsgAuditAgreeStatusDefaultAgree 默认同意
	MsgAuditAgreeStatusDefaultAgree = "Default_Agree"
)

// CheckSingleAgreeUserInfo 获取会话同意情况（单聊）内外成员
type CheckSingleAgreeUserInfo struct {
	// UserID 内部成员的userid
	UserID string `json:"userid"`
	// ExternalOpenID 外部成员的externalopenid
	ExternalOpenID string `json:"exteranalopenid"`
}

// CheckSingleAgreeInfo 获取会话同意情况（单聊）同意信息
type CheckSingleAgreeInfo struct {
	CheckSingleAgreeUserInfo
	// AgreeStatus 同意:”Agree”，不同意:”Disagree”，默认同意:”Default_Agree”
	AgreeStatus MsgAuditAgreeStatus
	// StatusChangeTime 同意状态改变的具体时间
	StatusChangeTime time.Time
}

// CheckSingleAgree 获取会话同意情况（单聊）
func (c *WorkwxApp) CheckSingleAgree(infos []CheckSingleAgreeUserInfo) ([]CheckSingleAgreeInfo, error) {
	resp, err := c.execCheckSingleAgree(reqCheckSingleAgree{
		Infos: infos,
	})
	if err != nil {
		return nil, err
	}
	return resp.intoCheckSingleAgreeInfoList(), nil
}

// CheckRoomAgreeInfo 获取会话同意情况（群聊）同意信息
type CheckRoomAgreeInfo struct {
	// StatusChangeTime 同意状态改变的具体时间
	StatusChangeTime time.Time
	// AgreeStatus 同意:”Agree”，不同意:”Disagree”，默认同意:”Default_Agree”
	AgreeStatus MsgAuditAgreeStatus
	// ExternalOpenID 群内外部联系人的externalopenid
	ExternalOpenID string
}

// CheckRoomAgree 获取会话同意情况（群聊）
func (c *WorkwxApp) CheckRoomAgree(roomId string) ([]CheckRoomAgreeInfo, error) {
	resp, err := c.execCheckRoomAgree(reqCheckRoomAgree{
		RoomID: roomId,
	})
	if err != nil {
		return nil, err
	}
	return resp.intoCheckRoomAgreeInfoList(), nil
}

// MsgAuditEdition 会话内容存档版本
type MsgAuditEdition uint8

const (
	// MsgAuditEditionOffice 会话内容存档办公版
	MsgAuditEditionOffice MsgAuditEdition = 1
	// MsgAuditEditionService 会话内容存档服务版
	MsgAuditEditionService MsgAuditEdition = 2
	// MsgAuditEditionEnterprise 会话内容存档企业版
	MsgAuditEditionEnterprise MsgAuditEdition = 3
)

// ListPermitUser 获取会话内容存档开启成员列表
func (c *WorkwxApp) ListPermitUser(msgAuditEdition MsgAuditEdition) ([]string, error) {
	resp, err := c.execListPermitUser(reqListPermitUser{
		MsgAuditEdition: msgAuditEdition,
	})
	if err != nil {
		return nil, err
	}
	return resp.IDs, nil
}

// GroupChatMember 获取会话内容存档内部群成员
type GroupChatMember struct {
	// MemberID roomid群成员的id，userid
	MemberID int
	// JoinTime roomid群成员的入群时间
	JoinTime time.Time
}

// GroupChat 获取会话内容存档内部群信息
type GroupChat struct {
	// Members roomid对应的群成员列表
	Members []GroupChatMember
	// RoomName roomid对应的群名称
	RoomName string
	// Creator roomid对应的群创建者，userid
	Creator string
	// RoomCreateTime roomid对应的群创建时间
	RoomCreateTime time.Time
	// Notice roomid对应的群公告
	Notice string
}

// GetGroupChat 获取会话内容存档内部群信息
func (c *WorkwxApp) GetGroupChat(roomID string) (*GroupChat, error) {
	resp, err := c.execGetGroupChat(reqGetGroupChat{
		RoomID: roomID,
	})
	if err != nil {
		return nil, err
	}
	groupChat := resp.intoGroupChat()
	return &groupChat, nil
}
