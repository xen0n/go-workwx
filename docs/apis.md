# Access token 获取

## API calls

Name|Request Type|Response Type|Access Token|URL|Doc
:---|------------|-------------|------------|:--|:--
`execGetAccessToken`|`reqAccessToken`|`respAccessToken`|-|`GET /cgi-bin/gettoken`|[获取access_token](https://work.weixin.qq.com/api/doc#90000/90135/91039)
`execGetJSAPITicket`|`reqJSAPITicket`|`respJSAPITicket`|+|`GET /cgi-bin/get_jsapi_ticket`|[获取企业的jsapi_ticket](https://open.work.weixin.qq.com/api/doc/90000/90136/90506)
`execGetJSAPITicketAgentConfig`|`reqJSAPITicketAgentConfig`|`respJSAPITicket`|+|`GET /cgi-bin/ticket/get`|[获取应用的jsapi_ticket](https://open.work.weixin.qq.com/api/doc/90000/90136/90506)
`execJSCode2Session`|`reqJSCode2Session`|`respJSCode2Session`|+|`GET /cgi-bin/miniprogram/jscode2session`|[临时登录凭证校验code2Session](https://open.work.weixin.qq.com/api/doc/90000/90136/91507)

# 成员管理

## API calls

Name|Request Type|Response Type|Access Token|URL|Doc
:---|------------|-------------|------------|:--|:--
`execUserCreate`|TODO|TODO|+|`POST /cgi-bin/user/create`|[创建成员](https://work.weixin.qq.com/api/doc#90000/90135/90195)
`execUserGet`|`reqUserGet`|`respUserGet`|+|`GET /cgi-bin/user/get`|[读取成员](https://work.weixin.qq.com/api/doc#90000/90135/90196)
`execUserUpdate`|TODO|TODO|+|`POST /cgi/bin/user/update`|[更新成员](https://work.weixin.qq.com/api/doc#90000/90135/90197)
`execUserDelete`|TODO|TODO|+|`GET /cgi/bin/user/delete`|[删除成员](https://work.weixin.qq.com/api/doc#90000/90135/90198)
`execUserBatchDelete`|TODO|TODO|+|`POST /cgi/bin/user/batchdelete`|[批量删除成员](https://work.weixin.qq.com/api/doc#90000/90135/90199)
`execUserSimpleList`|TODO|TODO|+|`GET /cgi-bin/user/simplelist`|[获取部门成员](https://work.weixin.qq.com/api/doc#90000/90135/90200)
`execUserList`|`reqUserList`|`respUserList`|+|`GET /cgi-bin/user/list`|[获取部门成员详情](https://work.weixin.qq.com/api/doc#90000/90135/90201)
`execUserConvertToOpenID`|TODO|TODO|+|`POST /cgi-bin/user/convert_to_openid`|[userid与openid互换](https://work.weixin.qq.com/api/doc#90000/90135/90202)
`execUserAuthSucc`|TODO|TODO|+|`GET /cgi-bin/user/authsucc`|[二次验证](https://work.weixin.qq.com/api/doc#90000/90135/90203)
`execUserBatchInvite`|TODO|TODO|+|`POST /cgi-bin/batch/invite`|[邀请成员](https://work.weixin.qq.com/api/doc#90000/90135/90975)
`execUserIDByMobile`|`reqUserIDByMobile`|`respUserIDByMobile`|+|`POST /cgi-bin/user/getuserid`|[手机号获取userid](https://work.weixin.qq.com/api/doc/90001/90143/91693)

# 部门管理

## API calls

Name|Request Type|Response Type|Access Token|URL|Doc
:---|------------|-------------|------------|:--|:--
`execDeptCreate`|TODO|TODO|+|`POST /cgi-bin/department/create`|[创建部门](https://work.weixin.qq.com/api/doc#90000/90135/90205)
`execDeptUpdate`|TODO|TODO|+|`POST /cgi-bin/department/update`|[更新部门](https://work.weixin.qq.com/api/doc#90000/90135/90206)
`execDeptDelete`|TODO|TODO|+|`GET /cgi/bin/department/delete`|[删除部门](https://work.weixin.qq.com/api/doc#90000/90135/90207)
`execDeptList`|`reqDeptList`|`respDeptList`|+|`GET /cgi-bin/department/list`|[获取部门列表](https://work.weixin.qq.com/api/doc#90000/90135/90208)

# 标签管理

## API calls

Name|Request Type|Response Type|Access Token|URL|Doc
:---|------------|-------------|------------|:--|:--
`execTagCreate`|TODO|TODO|+|`POST /cgi-bin/tag/create`|[创建标签](https://work.weixin.qq.com/api/doc#90000/90135/90210)
`execTagUpdate`|TODO|TODO|+|`POST /cgi-bin/tag/update`|[更新标签名字](https://work.weixin.qq.com/api/doc#90000/90135/90211)
`execTagDelete`|TODO|TODO|+|`GET /cgi/bin/tag/delete`|[删除标签](https://work.weixin.qq.com/api/doc#90000/90135/90212)
`execTagListUsers`|TODO|TODO|+|`GET /cgi/bin/tag/get`|[获取标签成员](https://work.weixin.qq.com/api/doc#90000/90135/90213)
`execTagAddUsers`|TODO|TODO|+|`POST /cgi/bin/tag/addtagusers`|[增加标签成员](https://work.weixin.qq.com/api/doc#90000/90135/90214)
`execTagDeleteUsers`|TODO|TODO|+|`POST /cgi/bin/tag/deltagusers`|[删除标签成员](https://work.weixin.qq.com/api/doc#90000/90135/90215)
`execTagList`|TODO|TODO|+|`GET /cgi/bin/tag/list`|[获取标签列表](https://work.weixin.qq.com/api/doc#90000/90135/90216)

# 异步批量接口

## API calls

Name|Request Type|Response Type|Access Token|URL|Doc
:---|------------|-------------|------------|:--|:--

# 身份验证

## API calls

Name|Request Type|Response Type|Access Token|URL|Doc
:---|------------|-------------|------------|:--|:--
`execUserInfoGet`|`reqUserInfoGet`|`respUserInfoGet`|+|`GET /cgi-bin/user/getuserinfo`|[获取访问用户身份](https://work.weixin.qq.com/api/doc/90000/90135/91023)

# 外部联系人管理 - 客户管理

## API calls

Name|Request Type|Response Type|Access Token|URL|Doc
:---|------------|-------------|------------|:--|:--
`execExternalContactList`|`reqExternalContactList`|`respExternalContactList`|+|`GET /cgi-bin/externalcontact/list`|[获取客户列表](https://work.weixin.qq.com/api/doc/90000/90135/92113)
`execExternalContactGet`|`reqExternalContactGet`|`respExternalContactGet`|+|`GET /cgi-bin/externalcontact/get`|[获取客户详情](https://work.weixin.qq.com/api/doc/90000/90135/92114)
`execExternalContactBatchList`|`reqExternalContactBatchList`|`respExternalContactBatchList`|+|`POST /cgi-bin/externalcontact/batch/get_by_user`|[批量获取客户详情](https://work.weixin.qq.com/api/doc/90000/90135/92994)
`execExternalContactRemark`|`reqExternalContactRemark`|`respExternalContactRemark`|+|`POST /cgi-bin/externalcontact/remark`|[修改客户备注信息](https://work.weixin.qq.com/api/doc/90000/90135/92115)

# 外部联系人管理 - 客户标签管理

## API calls

Name|Request Type|Response Type|Access Token|URL|Doc
:---|------------|-------------|------------|:--|:--
`execExternalContactListCorpTags`|`reqExternalContactListCorpTags`|`respExternalContactListCorpTags`|+|`POST /cgi-bin/externalcontact/get_corp_tag_list`|[获取企业标签库](https://work.weixin.qq.com/api/doc/90000/90135/92117)
`execExternalContactAddCorpTag`|`reqExternalContactAddCorpTag`|`respExternalContactAddCorpTag`|+|`POST /cgi-bin/externalcontact/add_corp_tag`|[添加企业客户标签](https://work.weixin.qq.com/api/doc/90000/90135/92117)
`execExternalContactEditCorpTag`|`reqExternalContactEditCorpTag`|`respExternalContactEditCorpTag`|+|`POST /cgi-bin/externalcontact/edit_corp_tag`|[编辑企业客户标签](https://work.weixin.qq.com/api/doc/90000/90135/92117)
`execExternalContactDelCorpTag`|`reqExternalContactDelCorpTag`|`respExternalContactDelCorpTag`|+|`POST /cgi-bin/externalcontact/del_corp_tag`|[删除企业客户标签](https://work.weixin.qq.com/api/doc/90000/90135/92117)
`execExternalContactMarkTag`|`reqExternalContactMarkTag`|`respExternalContactMarkTag`|+|`POST /cgi-bin/externalcontact/mark_tag`|[标记客户企业标签](https://work.weixin.qq.com/api/doc/90000/90135/92118)

# 外部联系人管理 - 客户分配

## API calls

Name|Request Type|Response Type|Access Token|URL|Doc
:---|------------|-------------|------------|:--|:--
`execListUnassignedExternalContact`|`reqListUnassignedExternalContact`|`respListUnassignedExternalContact`|+|`POST /cgi-bin/externalcontact/get_unassigned_list`|[获取离职成员的客户列表](https://work.weixin.qq.com/api/doc/90000/90135/92124)
`execTransferExternalContact`|`reqTransferExternalContact`|`respTransferExternalContact`|+|`POST /cgi-bin/externalcontact/transfer`|[分配成员的客户](https://work.weixin.qq.com/api/doc/90000/90135/92125)
`execGetTransferExternalContactResult`|`reqGetTransferExternalContactResult`|`respGetTransferExternalContactResult`|+|`POST /cgi-bin/externalcontact/get_transfer_result`|[查询客户接替结果](https://work.weixin.qq.com/api/doc/90000/90135/92973)
`execTransferGroupChatExternalContact`|`reqTransferGroupChatExternalContact`|`respTransferGroupChatExternalContact`|+|`POST /cgi-bin/externalcontact/groupchat/transfer`|[离职成员的群再分配](https://work.weixin.qq.com/api/doc/90000/90135/92127)
`execExternalContractGroupChatGet`|`reqGroupChatExternalContact`|`respGetExternalContractGroupChatResult`|+|`POST /cgi-bin/externalcontact/groupchat/get`|[获取客户群详情](https://work.weixin.qq.com/api/doc/90000/90135/92122)

# 应用管理

## API calls

Name|Request Type|Response Type|Access Token|URL|Doc
:---|------------|-------------|------------|:--|:--
`execAgentGet`|TODO|TODO|+|`GET /cgi-bin/agent/get`|[获取指定的应用详情](https://work.weixin.qq.com/api/doc#90000/90135/90227)
`execAgentList`|TODO|TODO|+|`GET /cgi-bin/agent/list`|[获取access_token对应的应用列表](https://work.weixin.qq.com/api/doc#90000/90135/90227)
`execAgentSet`|TODO|TODO|+|`POST /cgi-bin/agent/set`|[设置应用](https://work.weixin.qq.com/api/doc#90000/90135/90228)

# 应用管理 - 自定义菜单

## API calls

Name|Request Type|Response Type|Access Token|URL|Doc
:---|------------|-------------|------------|:--|:--
`execMenuCreate`|TODO|TODO|+|`POST /cgi-bin/menu/create`|[创建菜单](https://work.weixin.qq.com/api/doc#90000/90135/90231)
`execMenuGet`|TODO|TODO|+|`GET /cgi-bin/menu/get`|[获取菜单](https://work.weixin.qq.com/api/doc#90000/90135/90232)
`execMenuDelete`|TODO|TODO|+|`GET /cgi-bin/menu/delete`|[删除菜单](https://work.weixin.qq.com/api/doc#90000/90135/90233)

# 消息推送

## API calls

Name|Request Type|Response Type|Access Token|URL|Doc
:---|------------|-------------|------------|:--|:--
`execAppchatCreate`|`reqAppchatCreate`|`respAppchatCreate`|+|`POST /cgi-bin/appchat/create`|[创建群聊会话](https://work.weixin.qq.com/api/doc#90000/90135/90245)
`execAppchatUpdate`|TODO|TODO|+|`POST /cgi-bin/appchat/update`|[修改群聊会话](https://work.weixin.qq.com/api/doc#90000/90135/90246)
`execAppchatGet`|`reqAppchatGet`|`respAppchatGet`|+|`GET /cgi-bin/appchat/get`|[获取群聊会话](https://work.weixin.qq.com/api/doc#90000/90135/90247)
`execMessageSend`|`reqMessage`|`respMessageSend`|+|`POST /cgi-bin/message/send`|[发送应用消息](https://work.weixin.qq.com/api/doc#90000/90135/90236)
`execAppchatSend`|`reqMessage`|`respMessageSend`|+|`POST /cgi-bin/appchat/send`|[应用推送消息](https://work.weixin.qq.com/api/doc#90000/90135/90248)

# 素材管理

## API calls

Name|Request Type|Response Type|Access Token|URL|Doc
:---|------------|-------------|------------|:--|:--
`execMediaUpload`|`reqMediaUpload`|`respMediaUpload`|+|`POST(media) /cgi-bin/media/upload`|[上传临时素材](https://work.weixin.qq.com/api/doc#90000/90135/90253)
`execMediaUploadImg`|`reqMediaUploadImg`|`respMediaUploadImg`|+|`POST(media) /cgi-bin/media/uploadimg`|[上传永久图片](https://work.weixin.qq.com/api/doc#90000/90135/90256)
`execMediaGet`|TODO|TODO|+|`GET /cgi-bin/media/get`|[获取临时素材](https://work.weixin.qq.com/api/doc#90000/90135/90254)
`execMediaGetJSSDK`|TODO|TODO|+|`GET /cgi-bin/media/get/jssdk`|[获取高清语音素材](https://work.weixin.qq.com/api/doc#90000/90135/90255)

# OA 数据接口

## API calls

Name|Request Type|Response Type|Access Token|URL|Doc
:---|------------|-------------|------------|:--|:--
`execOAGetTemplateDetail`|`reqOAGetTemplateDetail`|`respOAGetTemplateDetail`|+|`POST /cgi-bin/oa/gettemplatedetail`|[获取审批模板详情](https://work.weixin.qq.com/api/doc/90000/90135/91982)
`execOAApplyEvent`|`reqOAApplyEvent`|`respOAApplyEvent`|+|`POST /cgi-bin/oa/applyevent`|[提交审批申请](https://work.weixin.qq.com/api/doc/90000/90135/91853)
`execOAGetApprovalInfo`|`reqOAGetApprovalInfo`|`respOAGetApprovalInfo`|+|`POST /cgi-bin/oa/getapprovalinfo`|[批量获取审批单号](https://work.weixin.qq.com/api/doc/90000/90135/91816)
`execOAGetApprovalDetail`|`reqOAGetApprovalDetail`|`respOAGetApprovalDetail`|+|`POST /cgi-bin/oa/getapprovaldetail`|[获取审批申请详情](https://work.weixin.qq.com/api/doc/90000/90135/91983)

# 企业支付

## API calls

Name|Request Type|Response Type|Access Token|URL|Doc
:---|------------|-------------|------------|:--|:--

# 电子发票

## API calls

Name|Request Type|Response Type|Access Token|URL|Doc
:---|------------|-------------|------------|:--|:--


# 会话内容存档

## API calls

Name|Request Type|Response Type|Access Token|URL|Doc
:---|------------|-------------|------------|:--|:--
`execMsgAuditListPermitUser`|`reqMsgAuditListPermitUser`|`respMsgAuditListPermitUser`|+|`POST /cgi-bin/msgaudit/get_permit_user_list`|[获取会话内容存档开启成员列表](https://work.weixin.qq.com/api/doc/90000/90135/91614)
`execMsgAuditCheckSingleAgree`|`reqMsgAuditCheckSingleAgree`|`respMsgAuditCheckSingleAgree`|+|`POST /cgi-bin/msgaudit/check_single_agree`|[获取会话同意情况（单聊）](https://work.weixin.qq.com/api/doc/90000/90135/91782)
`execMsgAuditCheckRoomAgree`|`reqMsgAuditCheckRoomAgree`|`respMsgAuditCheckRoomAgree`|+|`POST /cgi-bin/msgaudit/check_room_agree`|[获取会话同意情况（群聊）](https://work.weixin.qq.com/api/doc/90000/90135/91782)
`execMsgAuditGetGroupChat`|`reqMsgAuditGetGroupChat`|`respMsgAuditGetGroupChat`|+|`POST /cgi-bin/msgaudit/groupchat/get`|[获取会话内容存档内部群信息](https://work.weixin.qq.com/api/doc/90000/90135/92951)
