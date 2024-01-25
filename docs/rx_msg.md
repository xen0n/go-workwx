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

// EventTypeSysApprovalChange 审批申请状态变化回调通知
const EventTypeSysApprovalChange EventType = "sys_approval_change"

// EventTypeChangeContact 通讯录回调通知
const EventTypeChangeContact EventType = "change_contact"

// EventTypeKfMsgOrEvent 客服回调通知
const EventTypeKfMsgOrEvent EventType = "kf_msg_or_event"

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

// ChangeTypeCreateUser 新增成员事件
const ChangeTypeCreateUser ChangeType = "create_user"

// ChangeTypeUpdateUser 更新成员事件
const ChangeTypeUpdateUser ChangeType = "update_user"

// EventTypeAppMenuClick 点击菜单
const EventTypeAppMenuClick = "click"

// EventTypeAppMenuCView 打开菜单链接
const EventTypeAppMenuView = "view"

// EventTypeAppMenuScanCodePush 扫码上传
const EventTypeAppMenuScanCodePush = "scancode_push"

// EventTypeAppMenuScanCodeWaitMsg 扫码等待消息
const EventTypeAppMenuScanCodeWaitMsg = "scancode_waitmsg"

// EventTypeAppMenuPicSysPhoto 弹出系统拍照发图
const EventTypeAppMenuPicSysPhoto = "pic_sysphoto"

// EventTypeAppMenuPicPhotoOrAlbum 弹出系统拍照发图
const EventTypeAppMenuPicPhotoOrAlbum = "pic_photo_or_album"

// EventTypeAppMenuPicWeixin 弹出微信相册发图器
const EventTypeAppMenuPicWeixin = "pic_weixin"

// EventTypeAppMenuLocationSelect 弹出微信位置选择器
const EventTypeAppMenuLocationSelect = "location_select"

// EventTypeAppSubscribe 应用订阅
const EventTypeAppSubscribe = "subscribe"

// EventTypeAppUnsubscribe 应用订阅取消
const EventTypeAppUnsubscribe = "unsubscribe"
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

### `rxEventSysApprovalChange` 接收的事件消息，审批申请状态变化回调通知

Name|XML|Type|Doc
:---|:--|:---|:--
`ApprovalInfo`|`ApprovalInfo`|`OAApprovalInfo`|审批信息、

### `rxEventChangeTypeCreateUser` 接受的事件消息，新增成员事件

Name|XML|Type|Doc
:---|:--|:---|:--
`UserID`|`UserID`|`string`|成员UserID
`Name`|`Name`|`string`|成员名称
`Department`|`Department`|`string`|成员部门列表，仅返回该应用有查看权限的部门id
`IsLeaderInDept`|`IsLeaderInDept`|`string`|表示所在部门是否为上级，0-否，1-是，顺序与Department字段的部门逐一对应
`Mobile`|`Mobile`|`string`|手机号
`Position`|`Position`|`string`|职位信息。长度为0~64个字节
`Gender`|`Gender`|`int`|性别，1表示男性，2表示女性
`Email`|`Email`|`string`|邮箱
`Status`|`Status`|`int`|激活状态：1=已激活 2=已禁用 4=未激活 已激活代表已激活企业微信或已关注微工作台（原企业号）5=成员退出
`Avatar`|`Avatar`|`string`|头像url。注：如果要获取小图将url最后的”/0”改成”/100”即可。
`Alias`|`Alias`|`string`|成员别名
`Telephone`|`Telephone`|`string`|座机
`Address`|`Address`|`string`|地址
`ExtAttr`|`ExtAttr`|`string`|扩展属性
`Type`|`Type`|`string`|扩展属性类型: 0-本文 1-网页
`Text`|`Text`|`string`|文本属性类型，扩展属性类型为0时填写
`Value`|`Value`|`string`|文本属性内容
`Web`|`Web`|`string`|网页类型属性，扩展属性类型为1时填写
`Title`|`Title`|`string`|网页的展示标题
`Url`|`Url`|`string`|网页的url

### `rxEventChangeTypeUpdateUser` 接受的事件消息，更新成员事件

Name|XML|Type|Doc
:---|:--|:---|:--
`UserID`|`UserID`|`string`|成员UserID
`NewUserID`|`NewUserID`|`string`|新的UserID，变更时推送（userid由系统生成时可更改一次）
`Name`|`Name`|`string`|成员名称
`Department`|`Department`|`string`|成员部门列表，仅返回该应用有查看权限的部门id
`IsLeaderInDept`|`IsLeaderInDept`|`string`|表示所在部门是否为上级，0-否，1-是，顺序与Department字段的部门逐一对应
`Mobile`|`Mobile`|`string`|手机号
`Position`|`Position`|`string`|职位信息。长度为0~64个字节
`Gender`|`Gender`|`int`|性别，1表示男性，2表示女性
`Email`|`Email`|`string`|邮箱
`Status`|`Status`|`int`|激活状态：1=已激活 2=已禁用 4=未激活 已激活代表已激活企业微信或已关注微工作台（原企业号）5=成员退出
`Avatar`|`Avatar`|`string`|头像url。注：如果要获取小图将url最后的”/0”改成”/100”即可。
`Alias`|`Alias`|`string`|成员别名
`Telephone`|`Telephone`|`string`|座机
`Address`|`Address`|`string`|地址
`ExtAttr`|`ExtAttr`|`string`|扩展属性
`Type`|`Type`|`string`|扩展属性类型: 0-本文 1-网页
`Text`|`Text`|`string`|文本属性类型，扩展属性类型为0时填写
`Value`|`Value`|`string`|文本属性内容
`Web`|`Web`|`string`|网页类型属性，扩展属性类型为1时填写
`Title`|`Title`|`string`|网页的展示标题
`Url`|`Url`|`string`|网页的url

### `rxEventAppMenuClick` 接受的事件消息，应用菜单点击事件

Name|XML|Type|Doc
:---|:--|:---|:--
`EventKey`|`EventKey`|`string`|事件key

### `rxEventAppMenuView ` 接受的事件消息，应用菜单点击链接事件

Name|XML|Type|Doc
:---|:--|:---|:--
`EventKey`|`EventKey`|`string`|事件key

### `rxEventAppSubscribe` 接受的事件消息，用户订阅事件

Name|XML|Type|Doc
:---|:--|:---|:--
`EventKey`|`EventKey`|`string`|事件key

### `rxEventAppUnsubscribe` 接受的事件消息，用户取消订阅事件

Name|XML|Type|Doc
:---|:--|:---|:--
`EventKey`|`EventKey`|`string`|事件key

### `rxEventKfMsgOrEvent` 接受的事件消息，客服接收消息和事件

Name|XML|Type|Doc
:---|:--|:---|:--
`OpenKfID`|`OpenKfId`|`string`|有新消息的客服账号。可通过sync_msg接口指定open_kfid获取此客服账号的消息
`Token`|`Token`|`string`|调用拉取消息接口时，需要传此token，用于校验请求的合法性

### `rxEventUnknown` 接受的事件消息，未定义的事件类型

Name|XML|Type|Doc
:---|:--|:---|:--
`EventType`|`-`|`string`|事件类型
`Raw`|`-`|`string`|原始的消息体
