# 外部联系人

## Models

### `ExternalContact` 外部联系人

Name|JSON|Type|Doc
:---|:---|:---|:--
`ExternalUserid`|`external_userid`|`string`| 外部联系人的userid
`Name`|`name`|`string`| 外部联系人的名称，如果外部联系人为微信用户，则返回外部联系人的名称为其微信昵称；如果外部联系人为企业微信用户，则会按照以下优先级顺序返回：此外部联系人或管理员设置的昵称、认证的实名和账号名称。
`Position`|`position`|`string`| 外部联系人的职位，如果外部企业或用户选择隐藏职位，则不返回，仅当联系人类型是企业微信用户时有此字段
`Avatar`|`avatar`|`string`| 外部联系人头像，第三方不可获取
`CorpName`|`corp_name`|`string`| 外部联系人所在企业的简称，仅当联系人类型是企业微信用户时有此字段
`Type`|`type`|`ExternalUserType`| 外部联系人的类型，1表示该外部联系人是微信用户，2表示该外部联系人是企业微信用户
`Gender`|`gender`|`UserGender`| 外部联系人性别 0-未知 1-男性 2-女性
`Unionid`|`unionid`|`string`| 外部联系人在微信开放平台的唯一身份标识（微信unionid），通过此字段企业可将外部联系人与公众号/小程序用户关联起来。仅当联系人类型是微信用户，且企业或第三方服务商绑定了微信开发者ID有此字段。[查看绑定方法](https://work.weixin.qq.com/api/doc/90000/90135/92114#%E5%A6%82%E4%BD%95%E7%BB%91%E5%AE%9A%E5%BE%AE%E4%BF%A1%E5%BC%80%E5%8F%91%E8%80%85ID) 关于返回的unionid，如果是第三方应用调用该接口，则返回的unionid是该第三方服务商所关联的微信开放者帐号下的unionid。也就是说，同一个企业客户，企业自己调用，与第三方服务商调用，所返回的unionid不同；不同的服务商调用，所返回的unionid也不同。
`ExternalProfile`|`external_profile`|`ExternalProfile`| 成员对外信息

### `ExternalProfile` 成员对外信息

Name|JSON|Type|Doc
:---|:---|:---|:--
`ExternalCorpName`|`external_corp_name`|`string`| 企业简称
`ExternalAttr`|`external_attr`|`[]ExternalAttr`| 属性列表，目前支持文本、网页、小程序三种类型

### `ExternalAttr` 属性列表，目前支持文本、网页、小程序三种类型

Name|JSON|Type|Doc
:---|:---|:--|:---
`Type`|`type`|`int`|属性类型: 0-文本 1-网页 2-小程序
`Name`|`name`|`string`|属性名称： 需要先确保在管理端有创建该属性，否则会忽略
`Text`|`text`|`ExternalAttrText`|文本类型的属性 ，type为0时必填
`Web`|`web`|`ExternalAttrWeb`|网页类型的属性，url和title字段要么同时为空表示清除该属性，要么同时不为空 ，type为1时必填
`Miniprogram`|`miniprogram`|`ExternalAttrMiniprogram`|小程序类型的属性，appid和title字段要么同时为空表示清除改属性，要么同时不为空 ，type为2时必填

### `ExternalAttrText` 文本类型的属性

Name|JSON|Type|Doc
:---|:---|:--|:---
`Value`|`value`|`string`|文本属性内容,长度限制12个UTF8字符

### `ExternalAttrWeb` 网页类型的属性，url和title字段要么同时为空表示清除该属性，要么同时不为空 ，type为1时必填

Name|JSON|Type|Doc
:---|:---|:--|:---
`Url`|`url`|`string`|网页的url,必须包含http或者https头
`Title`|`title`|`string`|网页的展示标题,长度限制12个UTF8字符

### `ExternalAttrMiniprogram` 小程序类型的属性，appid和title字段要么同时为空表示清除改属性，要么同时不为空 ，type为2时必填

Name|JSON|Type|Doc
:---|:---|:--|:---
`Appid`|`appid`|`string`|小程序appid，必须是有在本企业安装授权的小程序，否则会被忽略
`Pagepath`|`pagepath`|`string`|小程序的页面路径
`Title`|`title`|`string`|企业对外简称，需从已认证的企业简称中选填。可在“我的企业”页中查看企业简称认证状态。

```go
// ExternalUserType 外部联系人的类型
//
// 1表示该外部联系人是微信用户
// 2表示该外部联系人是企业微信用户
type ExternalUserType int

const (
	// ExternalUserTypeWeChat 微信用户
	ExternalUserTypeWeChat ExternalUserType = 1
	// ExternalUserTypeWorkWeChat 企业微信用户
	ExternalUserTypeWorkWeChat ExternalUserType = 2
)
```

### `FollowUser` 添加了外部联系人的企业成员

Name|JSON|Type|Doc
:---|:---|:---|:--
``|``|`FollowUserInfo`| 添加了外部联系人的企业成员
`Tags`|`tags`|`[]FollowUserTag`| 该成员添加此外部联系人所打标签

### `FollowInfo` 企业成员客户跟进信息，可以参考获取客户详情，但标签信息只会返回企业标签的tag_id，个人标签将不再返回

Name|JSON|Type|Doc
:---|:---|:---|:--
``|``|`FollowUserInfo`| 添加了外部联系人的企业成员
`TagID`|`tag_id`|`[]string`| 该成员添加此外部联系人所打标签

### `FollowUserInfo` 添加了外部联系人的企业成员

Name|JSON|Type|Doc
:---|:---|:---|:--
`UserID`|`userid`|`string`| 外部联系人的userid
`Remark`|`remark`|`string`| 该成员对此外部联系人的备注
`Description`|`description`|`string`| 该成员对此外部联系人的描述
`Createtime`|`createtime`|`int`| 该成员添加此外部联系人的时间
`RemarkCorpName`|`remark_corp_name`|`string`| 该成员对此客户备注的企业名称
`RemarkMobiles`|`remark_mobiles`|`[]string`| 该成员对此客户备注的手机号码，第三方不可获取
`AddWay`|`add_way`|`FollowUserAddWay`| 该成员添加此客户的来源
`OperUserID`|`oper_userid`|`string`| 发起添加的userid，如果成员主动添加，为成员的userid；如果是客户主动添加，则为客户的外部联系人userid；如果是内部成员共享/管理员分配，则为对应的成员/管理员userid
`State`|`state`|`string`| 企业自定义的state参数，用于区分客户具体是通过哪个「联系我」添加，由企业通过[创建「联系我」方式](https://work.weixin.qq.com/api/doc/90000/90135/92114#15645/%E9%85%8D%E7%BD%AE%E5%AE%A2%E6%88%B7%E8%81%94%E7%B3%BB%E3%80%8C%E8%81%94%E7%B3%BB%E6%88%91%E3%80%8D%E6%96%B9%E5%BC%8F)指定

### `FollowUserTag` 该成员添加此外部联系人所打标签

Name|JSON|Type|Doc
:---|:---|:---|:--
`GroupName`|`group_name`|`string`| 该成员添加此外部联系人所打标签的分组名称（标签功能需要企业微信升级到2.7.5及以上版本）
`TagName`|`tag_name`|`string`| 该成员添加此外部联系人所打标签名称
`Type`|`type`|`FollowUserTagType`| 该成员添加此外部联系人所打标签类型, 1-企业设置, 2-用户自定义

```go
// FollowUserTagType 该成员添加此外部联系人所打标签类型
//
// 1-企业设置
// 2-用户自定义
type FollowUserTagType int

const (
	// 企业设置
	FollowUserTagTypeWork FollowUserTagType = 1
	// 用户自定义
	FollowUserTagTypeUser FollowUserTagType = 2
)

// FollowUserAddWay 该成员添加此客户的来源
//
// 具体含义详见[来源定义](https://work.weixin.qq.com/api/doc/90000/90135/92114#13878/%E6%9D%A5%E6%BA%90%E5%AE%9A%E4%B9%89)
type FollowUserAddWay int

const (
	// 未知来源
	FollowUserAddWayUnknown FollowUserAddWay = 0
	// 扫描二维码
	FollowUserAddWayQRCode FollowUserAddWay = 1
	// 搜索手机号
	FollowUserAddWayMobile FollowUserAddWay = 2
	// 名片分享
	FollowUserAddWayCard FollowUserAddWay = 3
	// 群聊
	FollowUserAddWayGroupChat FollowUserAddWay = 4
	// 手机通讯录
	FollowUserAddWayAddressBook FollowUserAddWay = 5
	// 微信联系人
	FollowUserAddWayWeChatContact FollowUserAddWay = 6
	// 来自微信的添加好友申请
	FollowUserAddWayWeChatFriendApply FollowUserAddWay = 7
	// 安装第三方应用时自动添加的客服人员
	FollowUserAddWayThirdParty FollowUserAddWay = 8
	// 搜索邮箱
	FollowUserAddWayEmail FollowUserAddWay = 9
	// 内部成员共享
	FollowUserAddWayInternalShare FollowUserAddWay = 201
	// 管理员/负责人分配
	FollowUserAddWayAdmin FollowUserAddWay = 202
)
```

### `ExternalContactRemark` 客户备注信息

Name|JSON|Type|Doc
:---|:---|:---|:--
`Userid`|`userid`|`string`| 企业成员的userid
`ExternalUserid`|`external_userid`|`string`| 外部联系人userid
`Remark`|`remark`|`string`| 此用户对外部联系人的备注，最多20个字符，remark，description，remark_company，remark_mobiles和remark_pic_mediaid不可同时为空。
`Description`|`description`|`string`| 此用户对外部联系人的描述，最多150个字符
`RemarkCompany`|`remark_company`|`string`| 此用户对外部联系人备注的所属公司名称，最多20个字符，remark_company只在此外部联系人为微信用户时有效。
`RemarkMobiles`|`remark_mobiles`|`[]string`| 此用户对外部联系人备注的手机号，如果填写了remark_mobiles，将会覆盖旧的备注手机号。如果要清除所有备注手机号,请在remark_mobiles填写一个空字符串(“”)。
`RemarkPicMediaid`|`remark_pic_mediaid`|`string`| 备注图片的mediaid，remark_pic_mediaid可以通过素材管理接口获得。

### `ExternalContactCorpTag` 企业客户标签

Name|JSON|Type|Doc
:---|:---|:---|:--
`ID`|`id`|`string`| 标签id
`Name`|`name`|`string`| 标签名称
`CreateTime`|`create_time`|`int`| 标签创建时间
`Order`|`order`|`uint32`| 标签排序的次序值，order值大的排序靠前。有效的值范围是[0, 2^32)
`Deleted`|`deleted`|`bool`| 标签是否已经被删除，只在指定tag_id进行查询时返回

### `ExternalContactCorpTagGroup` 企业客户标签

Name|JSON|Type|Doc
:---|:---|:---|:--
`GroupID`|`group_id`|`string`| 标签组id
`GroupName`|`group_name`|`string`| 标签组名称
`CreateTime`|`create_time`|`int`| 标签组创建时间
`Order`|`order`|`uint32`| 标签组排序的次序值，order值大的排序靠前。有效的值范围是[0, 2^32)
`Deleted`|`deleted`|`bool`| 标签组是否已经被删除，只在指定tag_id进行查询时返回
`Tag`|`tag`|`[]ExternalContactCorpTag`| 标签组内的标签列表

### `ExternalContactMarkTag` 企业标记客户标签

Name|JSON|Type|Doc
:---|:---|:---|:--
`UserID`|`userid`|`string`| 添加外部联系人的userid
`ExternalUserID`|`external_userid`|`string`| 外部联系人userid
`AddTag`|`add_tag`|`[]string`| 要标记的标签列表
`RemoveTag`|`remove_tag`|`[]string`| 要移除的标签列表

### `ExternalContactUnassignedList` 离职成员的客户列表

Name|JSON|Type|Doc
:---|:---|:---|:--
`Info`|`info`|`[]ExternalContactUnassigned`| 离职成员的客户
`IsLast`|`is_last`|`bool`| 是否是最后一条记录
`NextCursor`|`next_cursor`|`string`| 分页查询游标,已经查完则返回空("")

```go
// ExternalContactTransferStatus 客户接替结果状态
type ExternalContactTransferStatus uint8

const (
	// ExternalContactTransferStatusSuccess 1-接替完毕
	ExternalContactTransferStatusSuccess ExternalContactTransferStatus = 1
	// ExternalContactTransferStatusWait 2-等待接替
	ExternalContactTransferStatusWait ExternalContactTransferStatus = 2
	// ExternalContactTransferStatusRefused 3-客户拒绝
	ExternalContactTransferStatusRefused ExternalContactTransferStatus = 3
	// ExternalContactTransferStatusExhausted 4-接替成员客户达到上限
	ExternalContactTransferStatusExhausted ExternalContactTransferStatus = 4
	// ExternalContactTransferStatusNoData 5-无接替记录
	ExternalContactTransferStatusNoData ExternalContactTransferStatus = 5
)
```

### `ExternalContactGroupChatTransferFailed` 离职成员的群再分配失败

Name|JSON|Type|Doc
:---|:---|:---|:--
`ChatID`|`chat_id`|`string`| 没能成功继承的群ID
`ErrCode`|`errcode`|`int`| 没能成功继承的群，错误码
`ErrMsg`|`errmsg`|`string`| 没能成功继承的群，错误描述

### `ExternalContactFollowUserList` 配置了客户联系功能的成员列表

Name|JSON|Type|Doc
:---|:---|:---|:--
`FollowUser`|`follow_user`|`[]string`| 配置了客户联系功能的成员userid列表

### `ExternalContactWay` 配置客户联系「联系我」方式

Name|JSON|Type|Doc
:---|:---|:---|:--
`Type`|`type`|`int`| 联系方式类型,1-单人, 2-多人
`Scene`|`scene`|`int`| 场景，1-在小程序中联系，2-通过二维码联系
`Style`|`style`|`int`| 在小程序中联系时使用的控件样式，详见附表
`Remark`|`remark`|`string`| 联系方式的备注信息，用于助记，不超过30个字符
`SkipVerify`|`skip_verify`|`bool`| 外部客户添加时是否无需验证，默认为true
`State`|`state`|`string`| 企业自定义的state参数，用于区分不同的添加渠道，在调用“获取外部联系人详情”时会返回该参数值，不超过30个字符 <https://developer.work.weixin.qq.com/document/path/92114>
`User`|`user`|`[]string`| 使用该联系方式的用户userID列表，在type为1时为必填，且只能有一个
`Party`|`party`|`[]int`| 使用该联系方式的部门id列表，只在type为2时有效
`IsTemp`|`is_temp`|`bool`| 是否临时会话模式，true表示使用临时会话模式，默认为false
`ExpiresIn`|`expires_in`|`int`| 临时会话二维码有效期，以秒为单位。该参数仅在is_temp为true时有效，默认7天，最多为14天
`ChatExpiresIn`|`chat_expires_in`|`int`| 临时会话有效期，以秒为单位。该参数仅在is_temp为true时有效，默认为添加好友后24小时，最多为14天
`UnionID`|`unionid`|`string`| 可进行临时会话的客户UnionID，该参数仅在is_temp为true时有效，如不指定则不进行限制
`Conclusions`|`conclusions`|`Conclusions`| 结束语，会话结束时自动发送给客户，可参考“结束语定义”，仅在is_temp为true时有效,<https://developer.work.weixin.qq.com/document/path/92572#%E7%BB%93%E6%9D%9F%E8%AF%AD%E5%AE%9A%E4%B9%89>

### `Conclusions` 结束语，会话结束时自动发送给客户

Name|JSON|Type|Doc
:---|:---|:---|:--
`Text`|`text`|`Text`| 文本消息
`Image`|`image`|`Image`| 图片
`Link`|`link`|`Link`| 链接
`MiniProgram`|`miniprogram`|`MiniProgram`| 小程序

### `Text` 结束语，会话结束时自动发送给客户

Name|JSON|Type|Doc
:---|:---|:---|:--
`Content`|`content`|`string`| 消息文本内容,最长为4000字节

### `Image` 结束语，会话结束时自动发送给客户

Name|JSON|Type|Doc
:---|:---|:---|:--
`MediaID`|`media_id`|`string`| 图片的media_id
`PicURL`|`pic_url`|`string`| 图片的url

### `Link` 结束语，会话结束时自动发送给客户

Name|JSON|Type|Doc
:---|:---|:---|:--
`Title`|`title`|`string`| 图文消息标题，最长为128字节
`Picurl`|`picurl`|`string`| 图文消息封面的url
`Desc`|`desc`|`string`| 图文消息的描述，最长为512字节
`URL`|`url`|`string`| 图文消息的链接

### `MiniProgram` 结束语，会话结束时自动发送给客户

Name|JSON|Type|Doc
:---|:---|:---|:--
`Title`|`title`|`string`| 小程序消息标题，最长为64字节
`PicMediaID`|`pic_media_id`|`string`| 小程序消息封面的mediaid，封面图建议尺寸为520*416
`AppID`|`appid`|`string`| 小程序appid，必须是关联到企业的小程序应用
`Page`|`page`|`string`| 小程序page路径

### `reqListContactWayExternalContact` 获取企业已配置的「联系我」列表请求参数

Name|JSON|Type|Doc
:---|:---|:---|:--
`StartTime`|`start_time`|`int`| 「联系我」创建起始时间戳, 默认为90天前
`EndTime`|`end_time`|`int`| 「联系我」创建结束时间戳, 默认为当前时间
`Cursor`|`cursor`|`string`| 分页查询使用的游标，为上次请求返回的 next_cursor
`Limit`|`limit`|`int`| 每次查询的分页大小，默认为100条，最多支持1000条

### `reqUpdateContactWayExternalContact` 更新企业已配置的「联系我」方式请求参数

Name|JSON|Type|Doc
:---|:---|:---|:--
`ConfigID`|`config_id`|`string`| 企业联系方式的配置id
`Remark`|`remark`|`string`| 联系方式的备注信息，不超过30个字符，将覆盖之前的备注
`SkipVerify`|`skip_verify`|`bool`| 外部客户添加时是否无需验证
`Style`|`style`|`int`| 样式，只针对“在小程序中联系”的配置生效
`State`|`state`|`string`| 企业自定义的state参数，用于区分不同的添加渠道，在调用“获取外部联系人详情”时会返回该参数值，不超过30个字符 <https://developer.work.weixin.qq.com/document/path/92114>
`User`|`user`|`[]string`| 使用该联系方式的用户userID列表，在type为1时为必填，且只能有一个
`Party`|`party`|`[]int`| 使用该联系方式的部门id列表，只在type为2时有效
`ExpiresIn`|`expires_in`|`int`| 临时会话二维码有效期，以秒为单位。该参数仅在is_temp为true时有效，默认7天，最多为14天
`ChatExpiresIn`|`chat_expires_in`|`int`| 临时会话有效期，以秒为单位。该参数仅在is_temp为true时有效，默认为添加好友后24小时，最多为14天
`UnionID`|`unionid`|`string`| 可进行临时会话的客户UnionID，该参数仅在is_temp为true时有效，如不指定则不进行限制
`Conclusions`|`conclusions`|`Conclusions`| 结束语，会话结束时自动发送给客户，可参考“结束语定义”，仅在is_temp为true时有效,<https://developer.work.weixin.qq.com/document/path/92572#%E7%BB%93%E6%9D%9F%E8%AF%AD%E5%AE%9A%E4%B9%89>
