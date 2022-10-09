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
`UseridList`|`userid_list`|`[]string`| 用户ID列表。最多100个



### `ReqChatList` 群聊列表获取参数

Name|JSON|Type|Doc
:---|:---|:---|:--
`StatusFilter`|`status_filter`|`int64`| 客户群跟进状态过滤
`OwnerFilter`|`owner_filter`|`ReqChatListOwnerFilter`| 群主过滤
`Cursor`|`cursor`|`string`| 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用不填
`Limit`|`limit`|`int64`| 分页，预期请求的数据量，取值范围 1 ~ 1000


### `RespGroupChatList` 客户群列表数据

Name|JSON|Type|Doc
:---|:---|:---|:--
`ChatId`|`chat_id`|`string`| 客户群ID
`Status`|`status`|`int64`| 客户群跟进状态 0 - 跟进人正常 1 - 跟进人离职 2 - 离职继承中 3 - 离职继承完成




### `RespAppchatList` 群聊列表结果

Name|JSON|Type|Doc
:---|:---|:---|:--
`GroupChatList`|`group_chat_list`|`[]RespGroupChatList`| 客户群列表
`NextCursor`|`next_cursor`|`string`| 分页游标

