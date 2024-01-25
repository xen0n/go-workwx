// Code generated by sdkcodegen; DO NOT EDIT.

package workwx

// execGetAccessToken 获取access_token
func (c *WorkwxApp) execGetAccessToken(req reqAccessToken) (respAccessToken, error) {
	var resp respAccessToken
	err := c.executeQyapiGet("/cgi-bin/gettoken", req, &resp, false)
	if err != nil {
		return respAccessToken{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respAccessToken{}, bizErr
	}

	return resp, nil
}

// execGetJSAPITicket 获取企业的jsapi_ticket
func (c *WorkwxApp) execGetJSAPITicket(req reqJSAPITicket) (respJSAPITicket, error) {
	var resp respJSAPITicket
	err := c.executeQyapiGet("/cgi-bin/get_jsapi_ticket", req, &resp, true)
	if err != nil {
		return respJSAPITicket{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respJSAPITicket{}, bizErr
	}

	return resp, nil
}

// execGetJSAPITicketAgentConfig 获取应用的jsapi_ticket
func (c *WorkwxApp) execGetJSAPITicketAgentConfig(req reqJSAPITicketAgentConfig) (respJSAPITicket, error) {
	var resp respJSAPITicket
	err := c.executeQyapiGet("/cgi-bin/ticket/get", req, &resp, true)
	if err != nil {
		return respJSAPITicket{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respJSAPITicket{}, bizErr
	}

	return resp, nil
}

// execJSCode2Session 临时登录凭证校验code2Session
func (c *WorkwxApp) execJSCode2Session(req reqJSCode2Session) (respJSCode2Session, error) {
	var resp respJSCode2Session
	err := c.executeQyapiGet("/cgi-bin/miniprogram/jscode2session", req, &resp, true)
	if err != nil {
		return respJSCode2Session{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respJSCode2Session{}, bizErr
	}

	return resp, nil
}

// execUserGet 读取成员
func (c *WorkwxApp) execUserGet(req reqUserGet) (respUserGet, error) {
	var resp respUserGet
	err := c.executeQyapiGet("/cgi-bin/user/get", req, &resp, true)
	if err != nil {
		return respUserGet{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respUserGet{}, bizErr
	}

	return resp, nil
}

// execUserUpdate 更新成员
func (c *WorkwxApp) execUserUpdate(req reqUserUpdate) (respUserUpdate, error) {
	var resp respUserUpdate
	err := c.executeQyapiJSONPost("/cgi-bin/user/update", req, &resp, true)
	if err != nil {
		return respUserUpdate{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respUserUpdate{}, bizErr
	}

	return resp, nil
}

// execUserList 获取部门成员详情
func (c *WorkwxApp) execUserList(req reqUserList) (respUserList, error) {
	var resp respUserList
	err := c.executeQyapiGet("/cgi-bin/user/list", req, &resp, true)
	if err != nil {
		return respUserList{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respUserList{}, bizErr
	}

	return resp, nil
}

// execConvertUserIDToOpenID userid转openid
func (c *WorkwxApp) execConvertUserIDToOpenID(req reqConvertUserIDToOpenID) (respConvertUserIDToOpenID, error) {
	var resp respConvertUserIDToOpenID
	err := c.executeQyapiJSONPost("/cgi-bin/user/convert_to_openid", req, &resp, true)
	if err != nil {
		return respConvertUserIDToOpenID{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respConvertUserIDToOpenID{}, bizErr
	}

	return resp, nil
}

// execConvertOpenIDToUserID openid转userid
func (c *WorkwxApp) execConvertOpenIDToUserID(req reqConvertOpenIDToUserID) (respConvertOpenIDToUserID, error) {
	var resp respConvertOpenIDToUserID
	err := c.executeQyapiJSONPost("/cgi-bin/user/convert_to_userid", req, &resp, true)
	if err != nil {
		return respConvertOpenIDToUserID{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respConvertOpenIDToUserID{}, bizErr
	}

	return resp, nil
}

// execUserJoinQrcode 获取加入企业二维码
func (c *WorkwxApp) execUserJoinQrcode(req reqUserJoinQrcode) (respUserJoinQrcode, error) {
	var resp respUserJoinQrcode
	err := c.executeQyapiGet("/cgi-bin/corp/get_join_qrcode", req, &resp, true)
	if err != nil {
		return respUserJoinQrcode{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respUserJoinQrcode{}, bizErr
	}

	return resp, nil
}

// execUserIDByMobile 手机号获取userid
func (c *WorkwxApp) execUserIDByMobile(req reqUserIDByMobile) (respUserIDByMobile, error) {
	var resp respUserIDByMobile
	err := c.executeQyapiJSONPost("/cgi-bin/user/getuserid", req, &resp, true)
	if err != nil {
		return respUserIDByMobile{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respUserIDByMobile{}, bizErr
	}

	return resp, nil
}

// execUserIDByEmail 邮箱获取userid
func (c *WorkwxApp) execUserIDByEmail(req reqUserIDByEmail) (respUserIDByEmail, error) {
	var resp respUserIDByEmail
	err := c.executeQyapiJSONPost("/cgi-bin/user/get_userid_by_email", req, &resp, true)
	if err != nil {
		return respUserIDByEmail{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respUserIDByEmail{}, bizErr
	}

	return resp, nil
}

// execDeptCreate 创建部门
func (c *WorkwxApp) execDeptCreate(req reqDeptCreate) (respDeptCreate, error) {
	var resp respDeptCreate
	err := c.executeQyapiJSONPost("/cgi-bin/department/create", req, &resp, true)
	if err != nil {
		return respDeptCreate{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respDeptCreate{}, bizErr
	}

	return resp, nil
}

// execDeptList 获取部门列表
func (c *WorkwxApp) execDeptList(req reqDeptList) (respDeptList, error) {
	var resp respDeptList
	err := c.executeQyapiGet("/cgi-bin/department/list", req, &resp, true)
	if err != nil {
		return respDeptList{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respDeptList{}, bizErr
	}

	return resp, nil
}

// execDeptSimpleList 获取子部门ID列表
func (c *WorkwxApp) execDeptSimpleList(req reqDeptSimpleList) (respDeptSimpleList, error) {
	var resp respDeptSimpleList
	err := c.executeQyapiGet("/cgi-bin/department/simplelist", req, &resp, true)
	if err != nil {
		return respDeptSimpleList{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respDeptSimpleList{}, bizErr
	}

	return resp, nil
}

// execUserInfoGet 获取访问用户身份
func (c *WorkwxApp) execUserInfoGet(req reqUserInfoGet) (respUserInfoGet, error) {
	var resp respUserInfoGet
	err := c.executeQyapiGet("/cgi-bin/user/getuserinfo", req, &resp, true)
	if err != nil {
		return respUserInfoGet{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respUserInfoGet{}, bizErr
	}

	return resp, nil
}

// execExternalContactList 获取客户列表
func (c *WorkwxApp) execExternalContactList(req reqExternalContactList) (respExternalContactList, error) {
	var resp respExternalContactList
	err := c.executeQyapiGet("/cgi-bin/externalcontact/list", req, &resp, true)
	if err != nil {
		return respExternalContactList{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respExternalContactList{}, bizErr
	}

	return resp, nil
}

// execExternalContactGet 获取客户详情
func (c *WorkwxApp) execExternalContactGet(req reqExternalContactGet) (respExternalContactGet, error) {
	var resp respExternalContactGet
	err := c.executeQyapiGet("/cgi-bin/externalcontact/get", req, &resp, true)
	if err != nil {
		return respExternalContactGet{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respExternalContactGet{}, bizErr
	}

	return resp, nil
}

// execExternalContactBatchList 批量获取客户详情
func (c *WorkwxApp) execExternalContactBatchList(req reqExternalContactBatchList) (respExternalContactBatchList, error) {
	var resp respExternalContactBatchList
	err := c.executeQyapiJSONPost("/cgi-bin/externalcontact/batch/get_by_user", req, &resp, true)
	if err != nil {
		return respExternalContactBatchList{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respExternalContactBatchList{}, bizErr
	}

	return resp, nil
}

// execExternalContactRemark 修改客户备注信息
func (c *WorkwxApp) execExternalContactRemark(req reqExternalContactRemark) (respExternalContactRemark, error) {
	var resp respExternalContactRemark
	err := c.executeQyapiJSONPost("/cgi-bin/externalcontact/remark", req, &resp, true)
	if err != nil {
		return respExternalContactRemark{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respExternalContactRemark{}, bizErr
	}

	return resp, nil
}

// execExternalContactListCorpTags 获取企业标签库
func (c *WorkwxApp) execExternalContactListCorpTags(req reqExternalContactListCorpTags) (respExternalContactListCorpTags, error) {
	var resp respExternalContactListCorpTags
	err := c.executeQyapiJSONPost("/cgi-bin/externalcontact/get_corp_tag_list", req, &resp, true)
	if err != nil {
		return respExternalContactListCorpTags{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respExternalContactListCorpTags{}, bizErr
	}

	return resp, nil
}

// execExternalContactAddCorpTag 添加企业客户标签
func (c *WorkwxApp) execExternalContactAddCorpTag(req reqExternalContactAddCorpTagGroup) (respExternalContactAddCorpTag, error) {
	var resp respExternalContactAddCorpTag
	err := c.executeQyapiJSONPost("/cgi-bin/externalcontact/add_corp_tag", req, &resp, true)
	if err != nil {
		return respExternalContactAddCorpTag{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respExternalContactAddCorpTag{}, bizErr
	}

	return resp, nil
}

// execExternalContactEditCorpTag 编辑企业客户标签
func (c *WorkwxApp) execExternalContactEditCorpTag(req reqExternalContactEditCorpTag) (respExternalContactEditCorpTag, error) {
	var resp respExternalContactEditCorpTag
	err := c.executeQyapiJSONPost("/cgi-bin/externalcontact/edit_corp_tag", req, &resp, true)
	if err != nil {
		return respExternalContactEditCorpTag{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respExternalContactEditCorpTag{}, bizErr
	}

	return resp, nil
}

// execExternalContactDelCorpTag 删除企业客户标签
func (c *WorkwxApp) execExternalContactDelCorpTag(req reqExternalContactDelCorpTag) (respExternalContactDelCorpTag, error) {
	var resp respExternalContactDelCorpTag
	err := c.executeQyapiJSONPost("/cgi-bin/externalcontact/del_corp_tag", req, &resp, true)
	if err != nil {
		return respExternalContactDelCorpTag{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respExternalContactDelCorpTag{}, bizErr
	}

	return resp, nil
}

// execExternalContactMarkTag 标记客户企业标签
func (c *WorkwxApp) execExternalContactMarkTag(req reqExternalContactMarkTag) (respExternalContactMarkTag, error) {
	var resp respExternalContactMarkTag
	err := c.executeQyapiJSONPost("/cgi-bin/externalcontact/mark_tag", req, &resp, true)
	if err != nil {
		return respExternalContactMarkTag{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respExternalContactMarkTag{}, bizErr
	}

	return resp, nil
}

// execListUnassignedExternalContact 获取离职成员的客户列表
func (c *WorkwxApp) execListUnassignedExternalContact(req reqListUnassignedExternalContact) (respListUnassignedExternalContact, error) {
	var resp respListUnassignedExternalContact
	err := c.executeQyapiJSONPost("/cgi-bin/externalcontact/get_unassigned_list", req, &resp, true)
	if err != nil {
		return respListUnassignedExternalContact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respListUnassignedExternalContact{}, bizErr
	}

	return resp, nil
}

// execTransferExternalContact 分配成员的客户
func (c *WorkwxApp) execTransferExternalContact(req reqTransferExternalContact) (respTransferExternalContact, error) {
	var resp respTransferExternalContact
	err := c.executeQyapiJSONPost("/cgi-bin/externalcontact/transfer", req, &resp, true)
	if err != nil {
		return respTransferExternalContact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respTransferExternalContact{}, bizErr
	}

	return resp, nil
}

// execGetTransferExternalContactResult 查询客户接替结果
func (c *WorkwxApp) execGetTransferExternalContactResult(req reqGetTransferExternalContactResult) (respGetTransferExternalContactResult, error) {
	var resp respGetTransferExternalContactResult
	err := c.executeQyapiJSONPost("/cgi-bin/externalcontact/get_transfer_result", req, &resp, true)
	if err != nil {
		return respGetTransferExternalContactResult{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respGetTransferExternalContactResult{}, bizErr
	}

	return resp, nil
}

// execTransferGroupChatExternalContact 离职成员的群再分配
func (c *WorkwxApp) execTransferGroupChatExternalContact(req reqTransferGroupChatExternalContact) (respTransferGroupChatExternalContact, error) {
	var resp respTransferGroupChatExternalContact
	err := c.executeQyapiJSONPost("/cgi-bin/externalcontact/groupchat/transfer", req, &resp, true)
	if err != nil {
		return respTransferGroupChatExternalContact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respTransferGroupChatExternalContact{}, bizErr
	}

	return resp, nil
}

// execAppchatCreate 创建群聊会话
func (c *WorkwxApp) execAppchatCreate(req reqAppchatCreate) (respAppchatCreate, error) {
	var resp respAppchatCreate
	err := c.executeQyapiJSONPost("/cgi-bin/appchat/create", req, &resp, true)
	if err != nil {
		return respAppchatCreate{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respAppchatCreate{}, bizErr
	}

	return resp, nil
}

// execAppchatUpdate 修改群聊会话
func (c *WorkwxApp) execAppchatUpdate(req reqAppchatUpdate) (respAppchatUpdate, error) {
	var resp respAppchatUpdate
	err := c.executeQyapiJSONPost("/cgi-bin/appchat/update", req, &resp, true)
	if err != nil {
		return respAppchatUpdate{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respAppchatUpdate{}, bizErr
	}

	return resp, nil
}

// execAppchatGet 获取群聊会话
func (c *WorkwxApp) execAppchatGet(req reqAppchatGet) (respAppchatGet, error) {
	var resp respAppchatGet
	err := c.executeQyapiGet("/cgi-bin/appchat/get", req, &resp, true)
	if err != nil {
		return respAppchatGet{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respAppchatGet{}, bizErr
	}

	return resp, nil
}

// execMessageSend 发送应用消息
func (c *WorkwxApp) execMessageSend(req reqMessage) (respMessageSend, error) {
	var resp respMessageSend
	err := c.executeQyapiJSONPost("/cgi-bin/message/send", req, &resp, true)
	if err != nil {
		return respMessageSend{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respMessageSend{}, bizErr
	}

	return resp, nil
}

// execAppchatSend 应用推送消息
func (c *WorkwxApp) execAppchatSend(req reqMessage) (respMessageSend, error) {
	var resp respMessageSend
	err := c.executeQyapiJSONPost("/cgi-bin/appchat/send", req, &resp, true)
	if err != nil {
		return respMessageSend{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respMessageSend{}, bizErr
	}

	return resp, nil
}

// execMediaUpload 上传临时素材
func (c *WorkwxApp) execMediaUpload(req reqMediaUpload) (respMediaUpload, error) {
	var resp respMediaUpload
	err := c.executeQyapiMediaUpload("/cgi-bin/media/upload", req, &resp, true)
	if err != nil {
		return respMediaUpload{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respMediaUpload{}, bizErr
	}

	return resp, nil
}

// execMediaUploadImg 上传永久图片
func (c *WorkwxApp) execMediaUploadImg(req reqMediaUploadImg) (respMediaUploadImg, error) {
	var resp respMediaUploadImg
	err := c.executeQyapiMediaUpload("/cgi-bin/media/uploadimg", req, &resp, true)
	if err != nil {
		return respMediaUploadImg{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respMediaUploadImg{}, bizErr
	}

	return resp, nil
}

// execOAGetTemplateDetail 获取审批模板详情
func (c *WorkwxApp) execOAGetTemplateDetail(req reqOAGetTemplateDetail) (respOAGetTemplateDetail, error) {
	var resp respOAGetTemplateDetail
	err := c.executeQyapiJSONPost("/cgi-bin/oa/gettemplatedetail", req, &resp, true)
	if err != nil {
		return respOAGetTemplateDetail{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respOAGetTemplateDetail{}, bizErr
	}

	return resp, nil
}

// execOAApplyEvent 提交审批申请
func (c *WorkwxApp) execOAApplyEvent(req reqOAApplyEvent) (respOAApplyEvent, error) {
	var resp respOAApplyEvent
	err := c.executeQyapiJSONPost("/cgi-bin/oa/applyevent", req, &resp, true)
	if err != nil {
		return respOAApplyEvent{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respOAApplyEvent{}, bizErr
	}

	return resp, nil
}

// execOAGetApprovalInfo 批量获取审批单号
func (c *WorkwxApp) execOAGetApprovalInfo(req reqOAGetApprovalInfo) (respOAGetApprovalInfo, error) {
	var resp respOAGetApprovalInfo
	err := c.executeQyapiJSONPost("/cgi-bin/oa/getapprovalinfo", req, &resp, true)
	if err != nil {
		return respOAGetApprovalInfo{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respOAGetApprovalInfo{}, bizErr
	}

	return resp, nil
}

// execOAGetApprovalDetail 获取审批申请详情
func (c *WorkwxApp) execOAGetApprovalDetail(req reqOAGetApprovalDetail) (respOAGetApprovalDetail, error) {
	var resp respOAGetApprovalDetail
	err := c.executeQyapiJSONPost("/cgi-bin/oa/getapprovaldetail", req, &resp, true)
	if err != nil {
		return respOAGetApprovalDetail{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respOAGetApprovalDetail{}, bizErr
	}

	return resp, nil
}

// execMsgAuditListPermitUser 获取会话内容存档开启成员列表
func (c *WorkwxApp) execMsgAuditListPermitUser(req reqMsgAuditListPermitUser) (respMsgAuditListPermitUser, error) {
	var resp respMsgAuditListPermitUser
	err := c.executeQyapiJSONPost("/cgi-bin/msgaudit/get_permit_user_list", req, &resp, true)
	if err != nil {
		return respMsgAuditListPermitUser{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respMsgAuditListPermitUser{}, bizErr
	}

	return resp, nil
}

// execMsgAuditCheckSingleAgree 获取会话同意情况（单聊）
func (c *WorkwxApp) execMsgAuditCheckSingleAgree(req reqMsgAuditCheckSingleAgree) (respMsgAuditCheckSingleAgree, error) {
	var resp respMsgAuditCheckSingleAgree
	err := c.executeQyapiJSONPost("/cgi-bin/msgaudit/check_single_agree", req, &resp, true)
	if err != nil {
		return respMsgAuditCheckSingleAgree{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respMsgAuditCheckSingleAgree{}, bizErr
	}

	return resp, nil
}

// execMsgAuditCheckRoomAgree 获取会话同意情况（群聊）
func (c *WorkwxApp) execMsgAuditCheckRoomAgree(req reqMsgAuditCheckRoomAgree) (respMsgAuditCheckRoomAgree, error) {
	var resp respMsgAuditCheckRoomAgree
	err := c.executeQyapiJSONPost("/cgi-bin/msgaudit/check_room_agree", req, &resp, true)
	if err != nil {
		return respMsgAuditCheckRoomAgree{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respMsgAuditCheckRoomAgree{}, bizErr
	}

	return resp, nil
}

// execMsgAuditGetGroupChat 获取会话内容存档内部群信息
func (c *WorkwxApp) execMsgAuditGetGroupChat(req reqMsgAuditGetGroupChat) (respMsgAuditGetGroupChat, error) {
	var resp respMsgAuditGetGroupChat
	err := c.executeQyapiJSONPost("/cgi-bin/msgaudit/groupchat/get", req, &resp, true)
	if err != nil {
		return respMsgAuditGetGroupChat{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respMsgAuditGetGroupChat{}, bizErr
	}

	return resp, nil
}

// execListFollowUserExternalContact 获取配置了客户联系功能的成员列表
func (c *WorkwxApp) execListFollowUserExternalContact(req reqListFollowUserExternalContact) (respListFollowUserExternalContact, error) {
	var resp respListFollowUserExternalContact
	err := c.executeQyapiGet("/cgi-bin/externalcontact/get_follow_user_list", req, &resp, true)
	if err != nil {
		return respListFollowUserExternalContact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respListFollowUserExternalContact{}, bizErr
	}

	return resp, nil
}

// execAddContactExternalContact 配置客户联系「联系我」方式
func (c *WorkwxApp) execAddContactExternalContact(req reqAddContactExternalContact) (respAddContactExternalContact, error) {
	var resp respAddContactExternalContact
	err := c.executeQyapiJSONPost("/cgi-bin/externalcontact/add_contact_way", req, &resp, true)
	if err != nil {
		return respAddContactExternalContact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respAddContactExternalContact{}, bizErr
	}

	return resp, nil
}

// execGetContactWayExternalContact 获取企业已配置的「联系我」方式
func (c *WorkwxApp) execGetContactWayExternalContact(req reqGetContactWayExternalContact) (respGetContactWayExternalContact, error) {
	var resp respGetContactWayExternalContact
	err := c.executeQyapiJSONPost("/cgi-bin/externalcontact/get_contact_way", req, &resp, true)
	if err != nil {
		return respGetContactWayExternalContact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respGetContactWayExternalContact{}, bizErr
	}

	return resp, nil
}

// execListContactWayChatExternalContact 获取企业已配置的「联系我」列表
func (c *WorkwxApp) execListContactWayChatExternalContact(req reqListContactWayExternalContact) (respListContactWayChatExternalContact, error) {
	var resp respListContactWayChatExternalContact
	err := c.executeQyapiJSONPost("/cgi-bin/externalcontact/list_contact_way", req, &resp, true)
	if err != nil {
		return respListContactWayChatExternalContact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respListContactWayChatExternalContact{}, bizErr
	}

	return resp, nil
}

// execUpdateContactWayExternalContact 更新企业已配置的「联系我」成员配置
func (c *WorkwxApp) execUpdateContactWayExternalContact(req reqUpdateContactWayExternalContact) (respUpdateContactWayExternalContact, error) {
	var resp respUpdateContactWayExternalContact
	err := c.executeQyapiJSONPost("/cgi-bin/externalcontact/update_contact_way", req, &resp, true)
	if err != nil {
		return respUpdateContactWayExternalContact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respUpdateContactWayExternalContact{}, bizErr
	}

	return resp, nil
}

// execDelContactWayExternalContact 删除企业已配置的「联系我」方式
func (c *WorkwxApp) execDelContactWayExternalContact(req reqDelContactWayExternalContact) (respDelContactWayExternalContact, error) {
	var resp respDelContactWayExternalContact
	err := c.executeQyapiJSONPost("/cgi-bin/externalcontact/del_contact_way", req, &resp, true)
	if err != nil {
		return respDelContactWayExternalContact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respDelContactWayExternalContact{}, bizErr
	}

	return resp, nil
}

// execCloseTempChatExternalContact 结束临时会话
func (c *WorkwxApp) execCloseTempChatExternalContact(req reqCloseTempChatExternalContact) (respCloseTempChatExternalContact, error) {
	var resp respCloseTempChatExternalContact
	err := c.executeQyapiJSONPost("/cgi-bin/externalcontact/close_temp_chat", req, &resp, true)
	if err != nil {
		return respCloseTempChatExternalContact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respCloseTempChatExternalContact{}, bizErr
	}

	return resp, nil
}

// execAddGroupChatJoinWayExternalContact 配置客户群「加入群聊」方式
func (c *WorkwxApp) execAddGroupChatJoinWayExternalContact(req reqAddGroupChatJoinWayExternalContact) (respAddGroupChatJoinWayExternalContact, error) {
	var resp respAddGroupChatJoinWayExternalContact
	err := c.executeQyapiJSONPost("/cgi-bin/externalcontact/groupchat/add_join_way", req, &resp, true)
	if err != nil {
		return respAddGroupChatJoinWayExternalContact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respAddGroupChatJoinWayExternalContact{}, bizErr
	}

	return resp, nil
}

// execGetGroupChatJoinWayExternalContact 获取企业已配置的客户群「加入群聊」方式
func (c *WorkwxApp) execGetGroupChatJoinWayExternalContact(req reqGetGroupChatJoinWayExternalContact) (respGetGroupChatJoinWayExternalContact, error) {
	var resp respGetGroupChatJoinWayExternalContact
	err := c.executeQyapiJSONPost("/cgi-bin/externalcontact/groupchat/get_join_way", req, &resp, true)
	if err != nil {
		return respGetGroupChatJoinWayExternalContact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respGetGroupChatJoinWayExternalContact{}, bizErr
	}

	return resp, nil
}

// execUpdateGroupChatJoinWayExternalContact 更新企业已配置的客户群「加入群聊」方式
func (c *WorkwxApp) execUpdateGroupChatJoinWayExternalContact(req reqUpdateGroupChatJoinWayExternalContact) (respUpdateGroupChatJoinWayExternalContact, error) {
	var resp respUpdateGroupChatJoinWayExternalContact
	err := c.executeQyapiJSONPost("/cgi-bin/externalcontact/groupchat/update_join_way", req, &resp, true)
	if err != nil {
		return respUpdateGroupChatJoinWayExternalContact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respUpdateGroupChatJoinWayExternalContact{}, bizErr
	}

	return resp, nil
}

// execDelGroupChatJoinWayExternalContact 删除企业已配置的客户群「加入群聊」方式
func (c *WorkwxApp) execDelGroupChatJoinWayExternalContact(req reqDelGroupChatJoinWayExternalContact) (respDelGroupChatJoinWayExternalContact, error) {
	var resp respDelGroupChatJoinWayExternalContact
	err := c.executeQyapiJSONPost("/cgi-bin/externalcontact/groupchat/del_join_way", req, &resp, true)
	if err != nil {
		return respDelGroupChatJoinWayExternalContact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respDelGroupChatJoinWayExternalContact{}, bizErr
	}

	return resp, nil
}

// execGroupChatListGet 获取客户群列表
func (c *WorkwxApp) execGroupChatListGet(req reqGroupChatList) (respGroupChatList, error) {
	var resp respGroupChatList
	err := c.executeQyapiJSONPost("/cgi-bin/externalcontact/groupchat/list", req, &resp, true)
	if err != nil {
		return respGroupChatList{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respGroupChatList{}, bizErr
	}

	return resp, nil
}

// execGroupChatInfoGet 获取客户群详细
func (c *WorkwxApp) execGroupChatInfoGet(req reqGroupChatInfo) (respGroupChatInfo, error) {
	var resp respGroupChatInfo
	err := c.executeQyapiJSONPost("/cgi-bin/externalcontact/groupchat/get", req, &resp, true)
	if err != nil {
		return respGroupChatInfo{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respGroupChatInfo{}, bizErr
	}

	return resp, nil
}

// execConvertOpenGIDToChatID 客户群opengid转换
func (c *WorkwxApp) execConvertOpenGIDToChatID(req reqConvertOpenGIDToChatID) (respConvertOpenGIDToChatID, error) {
	var resp respConvertOpenGIDToChatID
	err := c.executeQyapiJSONPost("/cgi-bin/externalcontact/opengid_to_chatid", req, &resp, true)
	if err != nil {
		return respConvertOpenGIDToChatID{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respConvertOpenGIDToChatID{}, bizErr
	}

	return resp, nil
}

// execTransferCustomer 在职继承 分配在职成员的客户
func (c *WorkwxApp) execTransferCustomer(req reqTransferCustomer) (respTransferCustomer, error) {
	var resp respTransferCustomer
	err := c.executeQyapiJSONPost("/cgi-bin/externalcontact/transfer_customer", req, &resp, true)
	if err != nil {
		return respTransferCustomer{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respTransferCustomer{}, bizErr
	}

	return resp, nil
}

// execGetTransferCustomerResult 在职继承 查询客户接替状态
func (c *WorkwxApp) execGetTransferCustomerResult(req reqGetTransferCustomerResult) (respGetTransferCustomerResult, error) {
	var resp respGetTransferCustomerResult
	err := c.executeQyapiJSONPost("/cgi-bin/externalcontact/transfer_result", req, &resp, true)
	if err != nil {
		return respGetTransferCustomerResult{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respGetTransferCustomerResult{}, bizErr
	}

	return resp, nil
}

// execTransferResignedCustomer 离职继承 分配离职成员的客户
func (c *WorkwxApp) execTransferResignedCustomer(req reqTransferCustomer) (respTransferCustomer, error) {
	var resp respTransferCustomer
	err := c.executeQyapiJSONPost("/cgi-bin/externalcontact/resigned/transfer_customer", req, &resp, true)
	if err != nil {
		return respTransferCustomer{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respTransferCustomer{}, bizErr
	}

	return resp, nil
}

// execGetTransferResignedCustomerResult 离职继承 查询客户接替状态
func (c *WorkwxApp) execGetTransferResignedCustomerResult(req reqGetTransferCustomerResult) (respGetTransferCustomerResult, error) {
	var resp respGetTransferCustomerResult
	err := c.executeQyapiJSONPost("/cgi-bin/externalcontact/resigned/transfer_result", req, &resp, true)
	if err != nil {
		return respGetTransferCustomerResult{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respGetTransferCustomerResult{}, bizErr
	}

	return resp, nil
}

// execAddMsgTemplate 创建企业群发
func (c *WorkwxApp) execAddMsgTemplate(req reqAddMsgTemplateExternalContact) (respAddMsgTemplateExternalContact, error) {
	var resp respAddMsgTemplateExternalContact
	err := c.executeQyapiJSONPost("/cgi-bin/externalcontact/add_msg_template", req, &resp, true)
	if err != nil {
		return respAddMsgTemplateExternalContact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respAddMsgTemplateExternalContact{}, bizErr
	}

	return resp, nil
}

// execSendWelcomeMsg 发送新客户欢迎语
func (c *WorkwxApp) execSendWelcomeMsg(req reqSendWelcomeMsgExternalContact) (respSendWelcomeMsgExternalContact, error) {
	var resp respSendWelcomeMsgExternalContact
	err := c.executeQyapiJSONPost("/cgi-bin/externalcontact/send_welcome_msg", req, &resp, true)
	if err != nil {
		return respSendWelcomeMsgExternalContact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respSendWelcomeMsgExternalContact{}, bizErr
	}

	return resp, nil
}

// execKfAccountCreate 添加客服账号
func (c *WorkwxApp) execKfAccountCreate(req reqKfAccountCreate) (respKfAccountCreate, error) {
	var resp respKfAccountCreate
	err := c.executeQyapiJSONPost("/cgi-bin/kf/account/add", req, &resp, true)
	if err != nil {
		return respKfAccountCreate{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respKfAccountCreate{}, bizErr
	}

	return resp, nil
}

// execKfAccountUpdate 修改客服账号
func (c *WorkwxApp) execKfAccountUpdate(req reqKfAccountUpdate) (respKfAccountUpdate, error) {
	var resp respKfAccountUpdate
	err := c.executeQyapiJSONPost("/cgi-bin/kf/account/update", req, &resp, true)
	if err != nil {
		return respKfAccountUpdate{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respKfAccountUpdate{}, bizErr
	}

	return resp, nil
}

// execKfAccountDelete 删除客服账号
func (c *WorkwxApp) execKfAccountDelete(req reqKfAccountDelete) (respKfAccountDelete, error) {
	var resp respKfAccountDelete
	err := c.executeQyapiJSONPost("/cgi-bin/kf/account/del", req, &resp, true)
	if err != nil {
		return respKfAccountDelete{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respKfAccountDelete{}, bizErr
	}

	return resp, nil
}

// execKfAccountList 获取客服账号列表
func (c *WorkwxApp) execKfAccountList(req reqKfAccountList) (respKfAccountList, error) {
	var resp respKfAccountList
	err := c.executeQyapiGet("/cgi-bin/kf/account/list", req, &resp, true)
	if err != nil {
		return respKfAccountList{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respKfAccountList{}, bizErr
	}

	return resp, nil
}

// execAddKfContact 获取客服账号链接
func (c *WorkwxApp) execAddKfContact(req reqAddKfContact) (respAddKfContact, error) {
	var resp respAddKfContact
	err := c.executeQyapiJSONPost("/cgi-bin/kf/add_contact_way", req, &resp, true)
	if err != nil {
		return respAddKfContact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respAddKfContact{}, bizErr
	}

	return resp, nil
}

// execKfServicerCreate 添加接待人员
func (c *WorkwxApp) execKfServicerCreate(req reqKfServicerCreate) (respKfServicerCreate, error) {
	var resp respKfServicerCreate
	err := c.executeQyapiJSONPost("/cgi-bin/kf/servicer/add", req, &resp, true)
	if err != nil {
		return respKfServicerCreate{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respKfServicerCreate{}, bizErr
	}

	return resp, nil
}

// execKfServicerDelete 删除接待人员
func (c *WorkwxApp) execKfServicerDelete(req reqKfServicerDelete) (respKfServicerDelete, error) {
	var resp respKfServicerDelete
	err := c.executeQyapiJSONPost("/cgi-bin/kf/servicer/del", req, &resp, true)
	if err != nil {
		return respKfServicerDelete{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respKfServicerDelete{}, bizErr
	}

	return resp, nil
}

// execKfServicerList 获取接待人员列表
func (c *WorkwxApp) execKfServicerList(req reqKfServicerList) (respKfServicerList, error) {
	var resp respKfServicerList
	err := c.executeQyapiGet("/cgi-bin/kf/servicer/list", req, &resp, true)
	if err != nil {
		return respKfServicerList{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respKfServicerList{}, bizErr
	}

	return resp, nil
}

// execKfServiceStateGet 获取会话状态
func (c *WorkwxApp) execKfServiceStateGet(req reqKfServiceStateGet) (respKfServiceStateGet, error) {
	var resp respKfServiceStateGet
	err := c.executeQyapiJSONPost("/cgi-bin/kf/service_state/get", req, &resp, true)
	if err != nil {
		return respKfServiceStateGet{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respKfServiceStateGet{}, bizErr
	}

	return resp, nil
}

// execKfServiceStateTrans 变更会话状态
func (c *WorkwxApp) execKfServiceStateTrans(req reqKfServiceStateTrans) (respKfServiceStateTrans, error) {
	var resp respKfServiceStateTrans
	err := c.executeQyapiJSONPost("/cgi-bin/kf/service_state/trans", req, &resp, true)
	if err != nil {
		return respKfServiceStateTrans{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respKfServiceStateTrans{}, bizErr
	}

	return resp, nil
}

// execKfSyncMsg 读取消息
func (c *WorkwxApp) execKfSyncMsg(req reqKfSyncMsg) (respKfSyncMsg, error) {
	var resp respKfSyncMsg
	err := c.executeQyapiJSONPost("/cgi-bin/kf/sync_msg", req, &resp, true)
	if err != nil {
		return respKfSyncMsg{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respKfSyncMsg{}, bizErr
	}

	return resp, nil
}

// execKfSend 发送消息
func (c *WorkwxApp) execKfSend(req reqMessage) (respMessageSend, error) {
	var resp respMessageSend
	err := c.executeQyapiJSONPost("/cgi-bin/kf/send_msg", req, &resp, true)
	if err != nil {
		return respMessageSend{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respMessageSend{}, bizErr
	}

	return resp, nil
}

// execKfOnEventSend 发送欢迎语等事件响应消息
func (c *WorkwxApp) execKfOnEventSend(req reqMessage) (respMessageSend, error) {
	var resp respMessageSend
	err := c.executeQyapiJSONPost("/cgi-bin/kf/send_msg_on_event", req, &resp, true)
	if err != nil {
		return respMessageSend{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respMessageSend{}, bizErr
	}

	return resp, nil
}
