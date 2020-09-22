# 接收消息格式

## Models

### `rxMessageCommon` 接收消息的公共部分

Name|XML|Type|Doc
:---|:--|:---|:--
`ToUserName`|`ToUserName`|`string`|企业微信CorpID
`FromUserName`|`FromUserName`|`string`|成员UserID
`CreateTime`|`CreateTime`|`int64`|消息创建时间（整型）
`MsgType`|`MsgType`|`MessageType`|消息类型
`MsgID`|`MsgId`|`int64`|消息id，64位整型
`AgentID`|`AgentID`|`int64`|企业应用的id，整型。可在应用的设置页面查看
`Event`|`Event`|`EventType`|事件类型 MsgType为event存在
`ChangeType`|`ChangeType`|`ChangeType`|变更类型 Event为change_external_contact存在
```go
// MessageType 消息类型
type MessageType string

// MessageTypeText 文本消息
const MessageTypeText MessageType = "text"

// MessageTypeImage 图片消息
const MessageTypeImage MessageType = "image"

// MessageTypeVoice 语音消息
const MessageTypeVoice MessageType = "voice"

// MessageTypeVideo 视频消息
const MessageTypeVideo MessageType = "video"

// MessageTypeLocation 位置消息
const MessageTypeLocation MessageType = "location"

// MessageTypeLink 链接消息
const MessageTypeLink MessageType = "link"

// MessageTypeEvent 事件消息
const MessageTypeEvent MessageType = "event"

// EventType 事件类型
type EventType string

// EventTypeChangeExternalContact 企业客户事件
const EventTypeChangeExternalContact EventType = "change_external_contact"

// EventTypeChangeExternalChat 客户群变更事件
const EventTypeChangeExternalChat EventType = "change_external_chat"

// ChangeType 变更类型
type ChangeType string

// ChangeTypeAddExternalContact 添加企业客户事件
const ChangeTypeAddExternalContact ChangeType = "add_external_contact"

// ChangeTypeEditExternalContact 编辑企业客户事件
const ChangeTypeEditExternalContact ChangeType = "edit_external_contact"

// ChangeTypeAddHalfExternalContact 外部联系人免验证添加成员事件
const ChangeTypeAddHalfExternalContact ChangeType = "add_half_external_contact"

// ChangeTypeDelExternalContact 删除企业客户事件
const ChangeTypeDelExternalContact ChangeType = "del_external_contact"

// ChangeTypeDelFollowUser 删除跟进成员事件
const ChangeTypeDelFollowUser ChangeType = "del_follow_user"

// ChangeTypeTransferFail 客户接替失败事件
const ChangeTypeTransferFail ChangeType = "transfer_fail"

```

### `rxTextMessageSpecifics` 接收的文本消息，特有字段

Name|XML|Type|Doc
:---|:--|:---|:--
`Content`|`Content`|`string`|文本消息内容

### `rxImageMessageSpecifics` 接收的图片消息，特有字段

Name|XML|Type|Doc
:---|:--|:---|:--
`PicURL`|`PicUrl`|`string`|图片链接
`MediaID`|`MediaId`|`string`|图片媒体文件id，可以调用获取媒体文件接口拉取，仅三天内有效

### `rxVoiceMessageSpecifics` 接收的语音消息，特有字段

Name|XML|Type|Doc
:---|:--|:---|:--
`MediaID`|`MediaId`|`string`|语音媒体文件id，可以调用获取媒体文件接口拉取数据，仅三天内有效
`Format`|`Format`|`string`|语音格式，如amr，speex等

### `rxVideoMessageSpecifics` 接收的视频消息，特有字段

Name|XML|Type|Doc
:---|:--|:---|:--
`MediaID`|`MediaId`|`string`|视频媒体文件id，可以调用获取媒体文件接口拉取数据，仅三天内有效
`ThumbMediaID`|`ThumbMediaId`|`string`|视频消息缩略图的媒体id，可以调用获取媒体文件接口拉取数据，仅三天内有效

### `rxLocationMessageSpecifics` 接收的位置消息，特有字段

Name|XML|Type|Doc
:---|:--|:---|:--
`Lat`|`Location_X`|`float64`|地理位置纬度
`Lon`|`Location_Y`|`float64`|地理位置经度
`Scale`|`Scale`|`int`|地图缩放大小
`Label`|`Label`|`string`|地理位置信息
`AppType`|`AppType`|`string`|app类型，在企业微信固定返回wxwork，在微信不返回该字段

### `rxLinkMessageSpecifics` 接收的链接消息，特有字段

Name|XML|Type|Doc
:---|:--|:---|:--
`Title`|`Title`|`string`|标题
`Description`|`Description`|`string`|描述
`URL`|`Url`|`string`|链接跳转的url
`PicURL`|`PicUrl`|`string`|封面缩略图的url

### `rxEventAddExternalContact` 接收的事件消息，添加企业客户事件

Name|XML|Type|Doc
:---|:--|:---|:--
`UserID`|`UserID`|`string`|企业服务人员的UserID
`ExternalUserID`|`ExternalUserID`|`string`|外部联系人的userid，注意不是企业成员的帐号
`State`|`State`|`string`|添加此用户的「联系我」方式配置的state参数，可用于识别添加此用户的渠道
`WelcomeCode`|`WelcomeCode`|`string`|欢迎语code，可用于发送欢迎语

### `rxEventEditExternalContact` 接收的事件消息，编辑企业客户事件

Name|XML|Type|Doc
:---|:--|:---|:--
`UserID`|`UserID`|`string`|企业服务人员的UserID
`ExternalUserID`|`ExternalUserID`|`string`|外部联系人的userid，注意不是企业成员的帐号
`State`|`State`|`string`|添加此用户的「联系我」方式配置的state参数，可用于识别添加此用户的渠道

### `rxEventAddHalfExternalContact` 接收的事件消息，外部联系人免验证添加成员事件

Name|XML|Type|Doc
:---|:--|:---|:--
`UserID`|`UserID`|`string`|企业服务人员的UserID
`ExternalUserID`|`ExternalUserID`|`string`|外部联系人的userid，注意不是企业成员的帐号
`State`|`State`|`string`|添加此用户的「联系我」方式配置的state参数，可用于识别添加此用户的渠道
`WelcomeCode`|`WelcomeCode`|`string`|欢迎语code，可用于发送欢迎语

### `rxEventDelExternalContact` 接收的事件消息，删除企业客户事件

Name|XML|Type|Doc
:---|:--|:---|:--
`UserID`|`UserID`|`string`|企业服务人员的UserID
`ExternalUserID`|`ExternalUserID`|`string`|外部联系人的userid，注意不是企业成员的帐号

### `rxEventDelFollowUser` 接收的事件消息，删除跟进成员事件

Name|XML|Type|Doc
:---|:--|:---|:--
`UserID`|`UserID`|`string`|企业服务人员的UserID
`ExternalUserID`|`ExternalUserID`|`string`|外部联系人的userid，注意不是企业成员的帐号

### `rxEventTransferFail` 接收的事件消息，客户接替失败事件

Name|XML|Type|Doc
:---|:--|:---|:--
`FailReason`|`FailReason`|`string`|接替失败的原因, customer_refused-客户拒绝， customer_limit_exceed-接替成员的客户数达到上限
`UserID`|`UserID`|`string`|企业服务人员的UserID
`ExternalUserID`|`ExternalUserID`|`string`|外部联系人的userid，注意不是企业成员的帐号

### `rxEventChangeExternalChat` 接收的事件消息，客户群变更事件

Name|XML|Type|Doc
:---|:--|:---|:--
`ToUserName`|`ToUserName`|`string`|企业微信CorpID
`FromUserName`|`FromUserName`|`string`|此事件该值固定为sys，表示该消息由系统生成
`FailReason`|`FailReason`|`string`|接替失败的原因, customer_refused-客户拒绝， customer_limit_exceed-接替成员的客户数达到上限
`ChatID`|`ChatId`|`string`|群ID
