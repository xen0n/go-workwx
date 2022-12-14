# 群聊

## Models

### `ChatInfo` 群聊信息

Name|JSON|Type|Doc
:---|:---|:---|:--
`ChatID`|`chatid`|`string`| 群聊唯一标志
`Name`|`name`|`string`|群聊名
`OwnerUserID`|`owner`|`string`|群主id
`MemberUserIDs`|`userlist`|`[]string`|群成员id列表

### `ReqChatListOwnerFilter` 群主过滤

Name|JSON|Type|Doc
:---|:---|:---|:--
`UserIDList`|`userid_list`|`[]string`| 用户ID列表。最多100个

### `ReqChatList` 获取客户群列表参数

Name|JSON|Type|Doc
:---|:---|:---|:--
`StatusFilter`|`status_filter`|`int64`| 客户群跟进状态过滤
`OwnerFilter`|`owner_filter`|`ReqChatListOwnerFilter`| 群主过滤
`Cursor`|`cursor`|`string`| 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用不填
`Limit`|`limit`|`int64`| 分页，预期请求的数据量，取值范围 1 ~ 1000

### `RespGroupChatList` 客户群列表数据

Name|JSON|Type|Doc
:---|:---|:---|:--
`ChatID`|`chat_id`|`string`| 客户群ID
`Status`|`status`|`int64`| 客户群跟进状态 0 - 跟进人正常 1 - 跟进人离职 2 - 离职继承中 3 - 离职继承完成

### `RespAppchatList` 客户群列表结果

Name|JSON|Type|Doc
:---|:---|:---|:--
`GroupChatList`|`group_chat_list`|`[]RespGroupChatList`| 客户群列表
`NextCursor`|`next_cursor`|`string`| 分页游标

### `ChatMemberList` 客户群成员列表

Name|JSON|Type|Doc
:---|:---|:---|:--
`UserID`|`userid`|`string`| 群成员ID
`Type`|`type`|`int64`| 群成员类型 1 - 企业成员  2 - 外部联系人
`UnionID`|`unionid`|`string`| 微信unionid
`JoinTime`|`join_time`|`int64`| 入群时间
`JoinScene`|`join_scene`|`int64`| 入群方式。1 - 由群成员邀请入群（直接邀请入群）2 - 由群成员邀请入群（通过邀请链接入群）3 - 通过扫描群二维码入群
`Invitor`|`invitor`|`ChatMemberListInvitor`| 邀请者。目前仅当是由本企业内部成员邀请入群时会返回该值
`GroupNickname`|`group_nickname`|`string`| 在群里的昵称
`Name`|`name`|`string`| 在群里名字

### `ChatMemberListInvitor` 入群邀请者

Name|JSON|Type|Doc
:---|:---|:---|:--
`UserID`|`userid`|`string`| 邀请者ID

### `ChatAdminList` 客户群管理员列表

Name|JSON|Type|Doc
:---|:---|:---|:--
`UserID`|`userid`|`string`| 管理员ID

### `RespAppChatInfo` 客户群详情
Name|JSON|Type|Doc
:---|:---|:---|:--
`ChatID`|`chat_id`|`string`| 客户群ID
`Name`|`name`|`string`| 客户群名称
`Owner`|`owner`|`string`| 群主ID
`CreateTime`|`create_time`|`int64`| 群创建时间
`Notice`|`notice`|`string`| 群公告
`MemberList`|`member_list`|`[]*ChatMemberList`| 群成员列表
`AdminList`|`admin_list`|`[]*ChatAdminList`| 群管理员列表

```go
// ChatNeedName 是否需要返回群成员的名字 0-不返回；1-返回。默认不返回
const ChatNeedName int64 = 1
```
