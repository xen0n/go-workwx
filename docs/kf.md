# 客服

## Models

### `KfAccount` 客服账号

 Name              | JSON                         | Type     | Doc                                                      
:------------------|:-----------------------------|:---------|:---------------------------------------------------------
 `OpenKfID`        | `open_kfid`                  | `string` | 客服账号ID                                                   
 `Name`            | `name`                       | `string` | 客服名称                                                     
 `Avatar`          | `avatar`                     | `string` | 客服头像URL                                                  
 `ManagePrivilege` | `manage_privilege,omitempty` | `bool`   | 当前调用接口的应用身份，是否有该客服账号的管理权限（编辑客服账号信息、分配会话和收发消息）。组件应用不返回此字段 

### `KfServicer` 客服接待人员

 Name           | JSON                      | Type     | Doc                                         
:---------------|:--------------------------|:---------|:--------------------------------------------
 `UserID`       | `userid,omitempty`        | `string` | 接待人员的userid。第三方应用获取到的为密文userid，即open_userid 
 `Status`       | `status`                  | `int`    | 接待人员的接待状态。0:接待中,1:停止接待。                     
 `StopType`     | `stop_type`               | `int`    | 接待人员的接待状态为「停止接待」的子类型。0:停止接待,1:暂时挂起          
 `DepartmentID` | `department_id,omitempty` | `int64`  | 接待人员部门的id                                   

### `KfServicerResult` 接待人员数据

 Name           | JSON                      | Type     | Doc         
:---------------|:--------------------------|:---------|:------------
 `UserID`       | `userid,omitempty`        | `string` | 接待人员的userid 
 `DepartmentID` | `department_id,omitempty` | `int64`  | 接待人员部门的id   
 `ErrCode`      | `errcode`                 | `int64`  | 该条记录的结果     
 `ErrMsg`       | `errmsg`                  | `string` | 结果信息

### `KfMsg` 客服消息数据

 Name           | JSON                        | Type          | Doc         
:---------------|:----------------------------|:--------------|:------------
`MsgID`      | `msgid,omitempty`           | `string`      | 消息ID
`OpenKfID`| `open_kfid,omitempty`       | `string`      |客服账号ID（msgtype为event，该字段不返回）
`ExternalUserID`| `external_userid,omitempty` | `string`      |客客户UserID（msgtype为event，该字段不返回）
`SendTime`      | `send_time,omitempty`       | `int64`       | 消息发送时间   
`Origin`      | `origin,omitempty`          | `int`         | 消息来源。3-微信客户发送的消息 4-系统推送的事件消息 5-接待人员在企业微信客户端发送的消息
`ServicerUserID`| `servicer_userid,omitempty` | `string`      |从企业微信给客户发消息的接待人员userid（即仅origin为5才返回；msgtype为event，该字段不返回）
`MsgType`      | `msgtype`                   | `MessageType` | 消息类型   
`Text`       | `text,omitempty`            | `Text`   | 文本消息
`Image`       | `image,omitempty`           | `Image`   | 图片消息
`Link`       | `link,omitempty`            | `Link`   | 链接消息
`MiniProgram`       | `mini_program,omitempty`    | `MiniProgram`   | 小程序消息
`Event`       | `event,omitempty`           | `KfEvent`     | 事件类型

### `KfEvent` 客服会话事件

Name|JSON|Type|Doc
:---|:--|:---|:--
`EventType`|`event_type`|`KfEventType`|事件类型
`OpenKfID`|`open_kfid`|`string`|客服账号ID
`ExternalUserID`|`external_userid,omitempty`|`string`|客户UserID，注意不是企业成员的帐号
`ServicerUserID`|`servicer_userid,omitempty`|`string`|接待人员userid
`Scene`|`scene,omitempty`|`string`|用户进入会话事件特有。进入会话的场景值，获取客服账号链接开发者自定义的场景值
`SceneParam`|`scene_param,omitempty`|`string`|用户进入会话事件特有。进入会话的自定义参数，获取客服账号链接返回的url，开发者按规范拼接的scene_param参数
`WelcomeCode`|`welcome_code,omitempty`|`string`|用户进入会话事件特有。如果满足发送欢迎语条件（条件为：用户在过去48小时里未收过欢迎语，且未向客服发过消息），会返回该字段。可用该welcome_code调用发送事件响应消息接口给客户发送欢迎语。
`WechatChannels`|`wechat_channels,omitempty`|`KfWechatChannels`|用户进入会话事件特有。进入会话的视频号信息，从视频号进入会话才有值
`FailMsgID`|`fail_msgid,omitempty`|`string`|消息发送失败事件特有。发送失败的消息msgid
`FailType`|`fail_type,omitempty`|`int`|消息发送失败事件特有。失败类型。0-未知原因 1-客服账号已删除 2-应用已关闭 4-会话已过期，超过48小时 5-会话已关闭 6-超过5条限制 8-主体未验证 10-用户拒收 11-企业未有成员登录企业微信App（排查方法：企业至少一个成员通过手机号验证/微信授权登录企业微信App即可）12-发送的消息为客服组件禁发的消息类型
`Status`|`status,omitempty`|`int`|接待人员接待状态变更事件特有。状态类型。1-接待中 2-停止接待
`StopType`|`stop_type,omitempty`|`int`|接待人员接待状态变更事件特有。接待人员的状态为「停止接待」的子类型。0:停止接待,1:暂时挂起
`ChangeType`|`change_type,omitempty`|`KfServiceState`|会话状态变更事件特有。变更类型，均为接待人员在企业微信客户端操作触发。1-从接待池接入会话 2-转接会话 3-结束会话 4-重新接入已结束/已转接会话
`OldServicerUserID`|`old_servicer_userid,omitempty`|`string`|会话状态变更事件特有。老的接待人员userid。仅change_type为2、3和4有值
`NewServicerUserid`|`new_servicer_userid,omitempty`|`string`|会话状态变更事件特有。新的接待人员userid。仅change_type为1、2和4有值
`MsgCode`|`msg_code,omitempty`|`string`|会话状态变更事件特有。用于发送事件响应消息的code，仅change_type为1和3时，会返回该字段。可用该msg_code调用发送事件响应消息接口给客户发送回复语或结束语。
`RecallMsgID`|`recall_msgid,omitempty`|`string`|撤回消息事件特有。 撤回的消息msgid
`RejectSwitch`|`reject_switch,omitempty`|`int`|拒收客户消息变更事件特有。 拒收客户消息，1表示接待人员拒收了客户消息，0表示接待人员取消拒收客户消息

### `KfWechatChannels` 进入会话的视频号信息，从视频号进入会话才有值

 Name           | JSON                      | Type     | Doc         
:---------------|:--------------------------|:---------|:------------
 `NickName`       | `nickname,omitempty`        | `string` | 视频号名称，视频号场景值为1、2、3时返回此项
 `ShopNickName` | `shop_nickname,omitempty` | `string`  | 视频号小店名称，视频号场景值为4、5时返回此项
 `Scene`      | `scene`                 | `int64`  | 视频号场景值。1：视频号主页，2：视频号直播间商品列表页，3：视频号商品橱窗页，4：视频号小店商品详情页，5：视频号小店订单页

```go
// KfEventType 事件类型
type KfEventType string

const (
    // KfEventTypeEnterSession 用户进入会话事件
    KfEventTypeEnterSession KfEventType = "enter_session"
    // KfEventTypeMsgSendFail 消息发送失败事件
	KfEventTypeMsgSendFail KfEventType = "msg_send_fail"
    // KfEventTypeServicerStatusChange 接待人员接待状态变更事件
	KfEventTypeServicerStatusChange KfEventType = "servicer_status_change"
    // KfEventTypeSessionStatusChange 会话状态变更事件
	KfEventTypeSessionStatusChange KfEventType = "session_status_change"
    // KfEventTypeUserRecallMsg 用户撤回消息事件
	KfEventTypeUserRecallMsg KfEventType = "user_recall_msg"
    // KfEventTypeServicerRecallMsg 接待人员撤回消息事件
    KfEventTypeServicerRecallMsg KfEventType = "servicer_recall_msg" 
    // KfEventTypeRejectCustomerMsgSwitchChange 拒收客户消息变更事件
    KfEventTypeRejectCustomerMsgSwitchChange KfEventType = "reject_customer_msg_switch_change"
)

// KfServiceState 客服会话状态
//
// 0 未处理 新会话接入
// 1 由智能助手接待
// 2 待接入池排队中
// 3 由人工接待
// 4 已结束/未开始
type KfServiceState int

const (
	// KfServiceStateUntreated 未处理 新会话接入
    KfServiceStateUntreated KfServiceState = iota
	// KfServiceStateRobotReception 由智能助手接待
    KfServiceStateRobotReception
    // KfServiceStateInQueue 待接入池排队中
    KfServiceStateInQueue
	// KfServiceStateManualReception 由人工接待
	KfServiceStateManualReception
	// KfServiceStateFinished 已结束/未开始
	KfServiceStateFinished
)
```
