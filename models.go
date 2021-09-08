package workwx

import (
	"encoding/json"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type reqAccessToken struct {
	CorpID     string
	CorpSecret string
}

var _ urlValuer = reqAccessToken{}

func (x reqAccessToken) intoURLValues() url.Values {
	return url.Values{
		"corpid":     {x.CorpID},
		"corpsecret": {x.CorpSecret},
	}
}

type respCommon struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// IsOK 响应体是否为一次成功请求的响应
//
// 实现依据: https://work.weixin.qq.com/api/doc#10013
//
// > 企业微信所有接口，返回包里都有errcode、errmsg。
// > 开发者需根据errcode是否为0判断是否调用成功(errcode意义请见全局错误码)。
// > 而errmsg仅作参考，后续可能会有变动，因此不可作为是否调用成功的判据。
func (x *respCommon) IsOK() bool {
	return x.ErrCode == 0
}

func (x *respCommon) TryIntoErr() error {
	if x.IsOK() {
		return nil
	}

	return &WorkwxClientError{
		Code: x.ErrCode,
		Msg:  x.ErrMsg,
	}
}

type respAccessToken struct {
	respCommon

	AccessToken   string `json:"access_token"`
	ExpiresInSecs int64  `json:"expires_in"`
}

type reqJSAPITicketAgentConfig struct{}

var _ urlValuer = reqJSAPITicketAgentConfig{}

func (x reqJSAPITicketAgentConfig) intoURLValues() url.Values {
	return url.Values{
		"type": {"agent_config"},
	}
}

type reqJSAPITicket struct{}

var _ urlValuer = reqJSAPITicket{}

func (x reqJSAPITicket) intoURLValues() url.Values {
	return url.Values{}
}

type respJSAPITicket struct {
	respCommon

	Ticket        string `json:"ticket"`
	ExpiresInSecs int64  `json:"expires_in"`
}

// reqMessage 消息发送请求
type reqMessage struct {
	ToUser  []string
	ToParty []string
	ToTag   []string
	ChatID  string
	AgentID int64
	MsgType string
	Content map[string]interface{}
	IsSafe  bool
}

var _ bodyer = reqMessage{}

func (x reqMessage) intoBody() ([]byte, error) {
	// fuck
	safeInt := 0
	if x.IsSafe {
		safeInt = 1
	}

	obj := map[string]interface{}{
		"msgtype": x.MsgType,
		"agentid": x.AgentID,
		"safe":    safeInt,
	}

	// msgtype polymorphism
	obj[x.MsgType] = x.Content

	// 复用这个结构体，因为是 package-private 的所以这么做没风险
	if x.ChatID != "" {
		obj["chatid"] = x.ChatID
	} else {
		obj["touser"] = strings.Join(x.ToUser, "|")
		obj["toparty"] = strings.Join(x.ToParty, "|")
		obj["totag"] = strings.Join(x.ToTag, "|")
	}

	result, err := json.Marshal(obj)
	if err != nil {
		// should never happen unless OOM or similar bad things
		// TODO: error_chain
		return nil, err
	}

	return result, nil
}

// respMessageSend 消息发送响应
type respMessageSend struct {
	respCommon

	InvalidUsers   string `json:"invaliduser"`
	InvalidParties string `json:"invalidparty"`
	InvalidTags    string `json:"invalidtag"`
}

type reqUserGet struct {
	UserID string
}

var _ urlValuer = reqUserGet{}

func (x reqUserGet) intoURLValues() url.Values {
	return url.Values{
		"userid": {x.UserID},
	}
}

// respUserDetail 成员详细信息的公共字段
type respUserDetail struct {
	UserID         string   `json:"userid"`
	Name           string   `json:"name"`
	DeptIDs        []int64  `json:"department"`
	DeptOrder      []uint32 `json:"order"`
	Position       string   `json:"position"`
	Mobile         string   `json:"mobile"`
	Gender         string   `json:"gender"`
	Email          string   `json:"email"`
	IsLeaderInDept []int    `json:"is_leader_in_dept"`
	AvatarURL      string   `json:"avatar"`
	Telephone      string   `json:"telephone"`
	IsEnabled      int      `json:"enable"`
	Alias          string   `json:"alias"`
	Status         int      `json:"status"`
	QRCodeURL      string   `json:"qr_code"`
	// TODO: extattr external_profile external_position
}

// respUserGet 读取成员响应
type respUserGet struct {
	respCommon

	respUserDetail
}

// reqUserList 部门成员请求
type reqUserList struct {
	DeptID     int64
	FetchChild bool
}

var _ urlValuer = reqUserList{}

func (x reqUserList) intoURLValues() url.Values {
	var fetchChild int64
	if x.FetchChild {
		fetchChild = 1
	}

	return url.Values{
		"department_id": {strconv.FormatInt(x.DeptID, 10)},
		"fetch_child":   {strconv.FormatInt(fetchChild, 10)},
	}
}

// respUsersByDeptID 部门成员详情响应
type respUserList struct {
	respCommon

	Users []*respUserDetail `json:"userlist"`
}

// reqUserIDByMobile 手机号获取 userid 请求
type reqUserIDByMobile struct {
	Mobile string `json:"mobile"`
}

var _ bodyer = reqUserIDByMobile{}

func (x reqUserIDByMobile) intoBody() ([]byte, error) {
	body, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// respUserIDByMobile 手机号获取 userid 响应
type respUserIDByMobile struct {
	respCommon

	UserID string `json:"userid"`
}

type reqDeptList struct {
	HaveID bool
	ID     int64
}

var _ urlValuer = reqDeptList{}

func (x reqDeptList) intoURLValues() url.Values {
	if !x.HaveID {
		return url.Values{}
	}

	return url.Values{
		"id": {strconv.FormatInt(x.ID, 10)},
	}
}

// respDeptList 部门列表响应
type respDeptList struct {
	respCommon

	// TODO: 不要懒惰，把 API 层的类型写好
	Department []*DeptInfo `json:"department"`
}

// reqAppchatGet 获取群聊会话请求
type reqAppchatGet struct {
	ChatID string
}

var _ urlValuer = reqAppchatGet{}

func (x reqAppchatGet) intoURLValues() url.Values {
	return url.Values{
		"chatid": {x.ChatID},
	}
}

// respAppchatGet 获取群聊会话响应
type respAppchatGet struct {
	respCommon

	ChatInfo *ChatInfo `json:"chat_info"`
}

// reqAppchatCreate 创建群聊会话请求
type reqAppchatCreate struct {
	ChatInfo *ChatInfo
}

var _ bodyer = reqAppchatCreate{}

func (x reqAppchatCreate) intoBody() ([]byte, error) {
	result, err := json.Marshal(x.ChatInfo)
	if err != nil {
		// should never happen unless OOM or similar bad things
		// TODO: error_chain
		return nil, err
	}

	return result, nil
}

// respAppchatCreate 创建群聊会话响应
type respAppchatCreate struct {
	respCommon

	ChatID string `json:"chatid"`
}

// reqMediaUpload 临时素材上传请求
type reqMediaUpload struct {
	Type  string
	Media *Media
}

var _ urlValuer = reqMediaUpload{}
var _ mediaUploader = reqMediaUpload{}

func (x reqMediaUpload) intoURLValues() url.Values {
	return url.Values{
		"type": {x.Type},
	}
}

func (x reqMediaUpload) getMedia() *Media {
	return x.Media
}

// respMediaUpload 临时素材上传响应
type respMediaUpload struct {
	respCommon

	Type      string `json:"type"`
	MediaID   string `json:"media_id"`
	CreatedAt string `json:"created_at"`
}

// reqMediaUploadImg 永久图片素材上传请求
type reqMediaUploadImg struct {
	Media *Media
}

var _ urlValuer = reqMediaUploadImg{}
var _ mediaUploader = reqMediaUploadImg{}

func (x reqMediaUploadImg) intoURLValues() url.Values {
	return url.Values{}
}

func (x reqMediaUploadImg) getMedia() *Media {
	return x.Media
}

// respMediaUploadImg 永久图片素材上传响应
type respMediaUploadImg struct {
	respCommon

	URL string `json:"url"`
}

// reqExternalContactList 获取客户列表
type reqExternalContactList struct {
	UserID string `json:"userid"`
}

var _ urlValuer = reqExternalContactList{}

func (x reqExternalContactList) intoURLValues() url.Values {
	return url.Values{
		"userid": {x.UserID},
	}
}

// respExternalContactList 获取客户列表
type respExternalContactList struct {
	respCommon

	ExternalUserID []string `json:"external_userid"`
}

// reqExternalContactGet 获取客户详情
type reqExternalContactGet struct {
	ExternalUserID string `json:"external_userid"`
}

var _ urlValuer = reqExternalContactGet{}

func (x reqExternalContactGet) intoURLValues() url.Values {
	return url.Values{
		"external_userid": {x.ExternalUserID},
	}
}

// respExternalContactGet 获取客户详情
type respExternalContactGet struct {
	respCommon
	ExternalContactInfo
}

// ExternalContactInfo 外部联系人信息
type ExternalContactInfo struct {
	ExternalContact ExternalContact `json:"external_contact"`
	FollowUser      []FollowUser    `json:"follow_user"`
}

// ExternalContactBatchInfo 外部联系人信息
type ExternalContactBatchInfo struct {
	ExternalContact ExternalContact `json:"external_contact"`
	FollowInfo      FollowInfo      `json:"follow_info"`
}

// BatchListExternalContactsResp 外部联系人信息
type BatchListExternalContactsResp struct {
	Result     []ExternalContactBatchInfo
	NextCursor string
}

// reqExternalContactBatchList 批量获取客户详情
type reqExternalContactBatchList struct {
	UserID string `json:"userid"`
	Cursor string `json:"cursor"`
	Limit  int    `json:"limit"`
}

var _ bodyer = reqExternalContactBatchList{}

func (x reqExternalContactBatchList) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		// should never happen unless OOM or similar bad things
		// TODO: error_chain
		return nil, err
	}

	return result, nil
}

// respExternalContactBatchList 批量获取客户详情
type respExternalContactBatchList struct {
	respCommon
	NextCursor          string                     `json:"next_cursor"`
	ExternalContactList []ExternalContactBatchInfo `json:"external_contact_list"`
}

// reqExternalContactRemark 获取客户详情
type reqExternalContactRemark struct {
	Remark *ExternalContactRemark
}

var _ bodyer = reqExternalContactRemark{}

func (x reqExternalContactRemark) intoBody() ([]byte, error) {
	result, err := json.Marshal(x.Remark)
	if err != nil {
		// should never happen unless OOM or similar bad things
		// TODO: error_chain
		return nil, err
	}

	return result, nil
}

// respExternalContactRemark 获取客户详情
type respExternalContactRemark struct {
	respCommon
}

// reqUserInfoGet 获取访问用户身份
type reqUserInfoGet struct {
	// 通过成员授权获取到的code，最大为512字节。每次成员授权带上的code将不一样，code只能使用一次，5分钟未被使用自动过期。
	Code string
}

var _ urlValuer = reqUserInfoGet{}

func (x reqUserInfoGet) intoURLValues() url.Values {
	return url.Values{
		"code": {x.Code},
	}
}

// respUserInfoGet 部门列表响应
type respUserInfoGet struct {
	respCommon
	UserIdentityInfo
}

// reqExternalContactListCorpTags 获取企业标签库
type reqExternalContactListCorpTags struct {
	// 要查询的标签id，如果不填则获取该企业的所有客户标签，目前暂不支持标签组id
	TagIDs []string `json:"tag_id"`
}

var _ bodyer = reqExternalContactListCorpTags{}

func (x reqExternalContactListCorpTags) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		// should never happen unless OOM or similar bad things
		// TODO: error_chain
		return nil, err
	}

	return result, nil
}

// respExternalContactListCorpTags 获取企业标签库
type respExternalContactListCorpTags struct {
	respCommon
	// 标签组列表
	TagGroup []ExternalContactCorpTagGroup `json:"tag_group"`
}

// reqExternalContactAddCorpTag 添加企业客户标签
type reqExternalContactAddCorpTag struct {
	ExternalContactCorpTagGroup
}

var _ bodyer = reqExternalContactAddCorpTag{}

func (x reqExternalContactAddCorpTag) intoBody() ([]byte, error) {
	result, err := json.Marshal(x.ExternalContactCorpTagGroup)
	if err != nil {
		// should never happen unless OOM or similar bad things
		// TODO: error_chain
		return nil, err
	}

	return result, nil
}

// respExternalContactAddCorpTag 添加企业客户标签
type respExternalContactAddCorpTag struct {
	respCommon
	// 标签组列表
	TagGroup []ExternalContactCorpTagGroup `json:"tag_group"`
}

// reqExternalContactEditCorpTag 编辑企业客户标签
type reqExternalContactEditCorpTag struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Order uint32 `json:"order"`
}

var _ bodyer = reqExternalContactEditCorpTag{}

func (x reqExternalContactEditCorpTag) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		// should never happen unless OOM or similar bad things
		// TODO: error_chain
		return nil, err
	}

	return result, nil
}

// respExternalContactEditCorpTag 编辑企业客户标签
type respExternalContactEditCorpTag struct {
	respCommon
}

// reqExternalContactDelCorpTag 删除企业客户标签
type reqExternalContactDelCorpTag struct {
	TagID   []string `json:"tag_id"`
	GroupID []string `json:"group_id"`
}

var _ bodyer = reqExternalContactDelCorpTag{}

func (x reqExternalContactDelCorpTag) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		// should never happen unless OOM or similar bad things
		// TODO: error_chain
		return nil, err
	}

	return result, nil
}

// respExternalContactDelCorpTag 删除企业客户标签
type respExternalContactDelCorpTag struct {
	respCommon
}

// reqExternalContactMarkTag 编辑企业客户标签
type reqExternalContactMarkTag struct {
	UserID         string   `json:"userid"`
	ExternalUserID string   `json:"external_userid"`
	AddTag         []string `json:"add_tag"`
	RemoveTag      []string `json:"remove_tag"`
}

var _ bodyer = reqExternalContactMarkTag{}

func (x reqExternalContactMarkTag) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		// should never happen unless OOM or similar bad things
		// TODO: error_chain
		return nil, err
	}

	return result, nil
}

// respExternalContactMarkTag 编辑企业客户标签
type respExternalContactMarkTag struct {
	respCommon
}

// reqJSCode2Session 临时登录凭证校验
type reqJSCode2Session struct {
	JSCode string
}

var _ urlValuer = reqJSCode2Session{}

func (x reqJSCode2Session) intoURLValues() url.Values {
	return url.Values{
		"js_code":    {x.JSCode},
		"grant_type": {"authorization_code"},
	}
}

// respJSCode2Session 临时登录凭证校验
type respJSCode2Session struct {
	respCommon
	JSCodeSession
}

// JSCodeSession 临时登录凭证
type JSCodeSession struct {
	CorpID     string `json:"corpid"`
	UserID     string `json:"userid"`
	SessionKey string `json:"session_key"`
}

type reqMsgAuditListPermitUser struct {
	MsgAuditEdition MsgAuditEdition `json:"type"`
}

var _ bodyer = reqMsgAuditListPermitUser{}

func (x reqMsgAuditListPermitUser) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type respMsgAuditListPermitUser struct {
	respCommon
	IDs []string `json:"ids"`
}

type reqMsgAuditCheckSingleAgree struct {
	Infos []CheckMsgAuditSingleAgreeUserInfo `json:"info"`
}

var _ bodyer = reqMsgAuditCheckSingleAgree{}

func (x reqMsgAuditCheckSingleAgree) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type respMsgAuditCheckSingleAgree struct {
	respCommon
	AgreeInfo []struct {
		UserID           string              `json:"userid"`
		ExternalOpenID   string              `json:"exteranalopenid"`
		AgreeStatus      MsgAuditAgreeStatus `json:"agree_status"`
		StatusChangeTime int                 `json:"status_change_time"`
	} `json:"agreeinfo"`
}

func (x respMsgAuditCheckSingleAgree) intoCheckSingleAgreeInfoList() (resp []CheckMsgAuditSingleAgreeInfo) {
	for _, agreeInfo := range x.AgreeInfo {
		resp = append(resp, CheckMsgAuditSingleAgreeInfo{
			CheckMsgAuditSingleAgreeUserInfo: CheckMsgAuditSingleAgreeUserInfo{
				UserID:         agreeInfo.UserID,
				ExternalOpenID: agreeInfo.ExternalOpenID,
			},
			AgreeStatus:      agreeInfo.AgreeStatus,
			StatusChangeTime: time.Unix(int64(agreeInfo.StatusChangeTime), 0),
		})
	}
	return resp
}

type reqMsgAuditCheckRoomAgree struct {
	RoomID string `json:"roomid"`
}

var _ bodyer = reqMsgAuditCheckRoomAgree{}

func (x reqMsgAuditCheckRoomAgree) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type respMsgAuditCheckRoomAgree struct {
	respCommon
	AgreeInfo []struct {
		StatusChangeTime int                 `json:"status_change_time"`
		AgreeStatus      MsgAuditAgreeStatus `json:"agree_status"`
		ExternalOpenID   string              `json:"exteranalopenid"`
	} `json:"agreeinfo"`
}

func (x respMsgAuditCheckRoomAgree) intoCheckRoomAgreeInfoList() (resp []CheckMsgAuditRoomAgreeInfo) {
	for _, agreeInfo := range x.AgreeInfo {
		resp = append(resp, CheckMsgAuditRoomAgreeInfo{
			StatusChangeTime: time.Unix(int64(agreeInfo.StatusChangeTime), 0),
			AgreeStatus:      agreeInfo.AgreeStatus,
			ExternalOpenID:   agreeInfo.ExternalOpenID,
		})
	}
	return resp
}

type reqMsgAuditGetGroupChat struct {
	RoomID string `json:"roomid"`
}

var _ bodyer = reqMsgAuditGetGroupChat{}

func (x reqMsgAuditGetGroupChat) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type respMsgAuditGetGroupChat struct {
	respCommon
	Members []struct {
		MemberID int `json:"memberid"`
		JoinTime int `json:"jointime"`
	} `json:"members"`
	RoomName       string `json:"roomname"`
	Creator        string `json:"creator"`
	RoomCreateTime int    `json:"room_create_time"`
	Notice         string `json:"notice"`
}

func (x respMsgAuditGetGroupChat) intoGroupChat() (resp MsgAuditGroupChat) {
	resp.Creator = x.Creator
	resp.Notice = x.Notice
	resp.RoomName = x.RoomName
	resp.RoomCreateTime = time.Unix(int64(x.RoomCreateTime), 0)
	for _, member := range x.Members {
		resp.Members = append(resp.Members, MsgAuditGroupChatMember{
			MemberID: member.MemberID,
			JoinTime: time.Unix(int64(member.JoinTime), 0),
		})
	}
	return resp
}

type reqListUnassignedExternalContact struct {
	// PageID 分页查询，要查询页号，从0开始
	PageID uint32 `json:"page_id"`
	// PageSize 每次返回的最大记录数，默认为1000，最大值为1000
	PageSize uint32 `json:"page_size"`
	// Cursor 分页查询游标，字符串类型，适用于数据量较大的情况，如果使用该参数则无需填写page_id，该参数由上一次调用返回
	Cursor string `json:"cursor"`
}

var _ bodyer = reqListUnassignedExternalContact{}

func (x reqListUnassignedExternalContact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type respListUnassignedExternalContact struct {
	respCommon
	Info []struct {
		HandoverUserid string `json:"handover_userid"`
		ExternalUserid string `json:"external_userid"`
		DemissionTime  int    `json:"dimission_time"`
	} `json:"info"`
	IsLast     bool   `json:"is_last"`
	NextCursor string `json:"next_cursor"`
}

func (x respListUnassignedExternalContact) intoExternalContactUnassignedList() (resp ExternalContactUnassignedList) {
	list := make([]ExternalContactUnassigned, 0, len(x.Info))
	for _, info := range x.Info {
		list = append(list, ExternalContactUnassigned{
			HandoverUserID: info.HandoverUserid,
			ExternalUserID: info.ExternalUserid,
			DemissionTime:  time.Unix(int64(info.DemissionTime), 0),
		})
	}
	resp.Info = list
	resp.IsLast = x.IsLast
	resp.NextCursor = x.NextCursor
	return resp
}

type reqTransferExternalContact struct {
	// ExternalUserID 外部联系人的userid，注意不是企业成员的帐号
	ExternalUserID string `json:"external_userid"`
	// HandoverUserID 原跟进成员的userid
	HandoverUserID string `json:"handover_userid"`
	// TakeoverUserID 接替成员的userid
	TakeoverUserID string `json:"takeover_userid"`
	// TransferSuccessMsg 转移成功后发给客户的消息，最多200个字符，不填则使用默认文案，目前只对在职成员分配客户的情况生效
	TransferSuccessMsg string `json:"transfer_success_msg"`
}

var _ bodyer = reqTransferExternalContact{}

func (x reqTransferExternalContact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type respTransferExternalContact struct {
	respCommon
}

type reqGetTransferExternalContactResult struct {
	// ExternalUserID 外部联系人的userid，注意不是企业成员的帐号
	ExternalUserID string `json:"external_userid"`
	// HandoverUserID 原跟进成员的userid
	HandoverUserID string `json:"handover_userid"`
	// TakeoverUserID 接替成员的userid
	TakeoverUserID string `json:"takeover_userid"`
}

var _ bodyer = reqGetTransferExternalContactResult{}

func (x reqGetTransferExternalContactResult) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type respGetTransferExternalContactResult struct {
	respCommon
	Status       uint8 `json:"status"`
	TakeoverTime int   `json:"takeover_time"`
}

func (x respGetTransferExternalContactResult) intoExternalContactTransferResult() ExternalContactTransferResult {
	return ExternalContactTransferResult{
		Status:       ExternalContactTransferStatus(x.Status),
		TakeoverTime: time.Unix(int64(x.TakeoverTime), 0),
	}
}

type reqTransferGroupChatExternalContact struct {
	// ChatIDList 需要转群主的客户群ID列表。取值范围： 1 ~ 100
	ChatIDList []string `json:"chat_id_list"`
	// NewOwner 新群主ID
	NewOwner string `json:"new_owner"`
}

var _ bodyer = reqTransferGroupChatExternalContact{}

func (x reqTransferGroupChatExternalContact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type respTransferGroupChatExternalContact struct {
	respCommon
	FailedChatList []ExternalContactGroupChatTransferFailed `json:"failed_chat_list"`
}

type reqOAGetTemplateDetail struct {
	TemplateID string `json:"template_id"`
}

var _ bodyer = reqOAGetTemplateDetail{}

func (x reqOAGetTemplateDetail) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type respOAGetTemplateDetail struct {
	respCommon
	OATemplateDetail
}

type reqOAApplyEvent struct {
	OAApplyEvent
}

var _ bodyer = reqOAApplyEvent{}

func (x reqOAApplyEvent) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type respOAApplyEvent struct {
	respCommon
	// SpNo 表单提交成功后，返回的表单编号
	SpNo string `json:"sp_no"`
}

type reqOAGetApprovalInfo struct {
	StartTime string                 `json:"starttime"`
	EndTime   string                 `json:"endtime"`
	Cursor    int                    `json:"cursor"`
	Size      uint32                 `json:"size"`
	Filters   []OAApprovalInfoFilter `json:"filters"`
}

var _ bodyer = reqOAGetApprovalInfo{}

func (x reqOAGetApprovalInfo) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type respOAGetApprovalInfo struct {
	respCommon
	// SpNoList 审批单号列表，包含满足条件的审批申请
	SpNoList []string `json:"sp_no_list"`
}

type reqOAGetApprovalDetail struct {
	// SpNo 审批单编号。
	SpNo string `json:"sp_no"`
}

var _ bodyer = reqOAGetApprovalDetail{}

func (x reqOAGetApprovalDetail) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type respOAGetApprovalDetail struct {
	respCommon
	// Info 审批申请详情
	Info OAApprovalDetail `json:"info"`
}

// TaskCardBtn 任务卡片消息按钮
type TaskCardBtn struct {
	// Key 按钮key值，用户点击后，会产生任务卡片回调事件，回调事件会带上该key值，只能由数字、字母和“_-@”组成，最长支持128字节
	Key string `json:"key"`
	// Name 按钮名称
	Name string `json:"name"`
	// ReplaceName 点击按钮后显示的名称，默认为“已处理”
	ReplaceName string `json:"replace_name"`
	// Color 按钮字体颜色，可选“red”或者“blue”,默认为“blue”
	Color string `json:"color"`
	// IsBold 按钮字体是否加粗，默认false
	IsBold bool `json:"is_bold"`
}

// 学校通知请求内容
type reqSchoolMessage struct {
	ParentIDs  []string
	StudentIDs []string
	PartyIDs   []string
	MsgType    string
	AgentID    int64
	Content    map[string]interface{}
}

func (x reqSchoolMessage) intoBody() ([]byte, error) {
	data := map[string]interface{}{
		"msgtype":           x.MsgType,
		"agentid":           x.AgentID,
		"to_parent_userid":  x.ParentIDs,
		"to_student_userid": x.StudentIDs,
		"to_party":          x.PartyIDs,
		"toall":             0,
	}

	data[x.MsgType] = x.Content

	result, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	return result, nil
}

type respSchoolMessageSend struct {
	respCommon

	InvalidParents  []string `json:"invalid_parent_userid"`
	InvalidStudents []string `json:"invalid_student_userid"`
	InvalidParts    []string `json:"invalid_party"`
}
