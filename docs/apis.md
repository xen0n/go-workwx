# Access token 获取

## API calls

Name|Request Type|Response Type|Access Token|URL
:---|------------|-------------|------------|:--
`execGetAccessToken`|`reqAccessToken`|`respAccessToken`|-|`GET /cgi-bin/gettoken`

# 成员管理

## API calls

Name|Request Type|Response Type|Access Token|URL
:---|------------|-------------|------------|:--
`execUserCreate`|TODO|TODO|+|`POST /cgi-bin/user/create`
`execUserGet`|`reqUserGet`|`respUserGet`|+|`GET /cgi-bin/user/get`
`execUserUpdate`|TODO|TODO|+|`POST /cgi/bin/user/update`
`execUserDelete`|TODO|TODO|+|`GET /cgi/bin/user/delete`
`execUserBatchDelete`|TODO|TODO|+|`POST /cgi/bin/user/batchdelete`
`execUserSimpleList`|TODO|TODO|+|`GET /cgi-bin/user/simplelist`
`execUserList`|TODO|TODO|+|`GET /cgi-bin/user/list`
`execUserConvertToOpenID`|TODO|TODO|+|`POST /cgi-bin/user/convert_to_openid`
`execUserAuthSucc`|TODO|TODO|+|`GET /cgi-bin/user/authsucc`
`execUserBatchInvite`|TODO|TODO|+|`POST /cgi-bin/batch/invite`

# 部门管理

## API calls

Name|Request Type|Response Type|Access Token|URL
:---|------------|-------------|------------|:--
`execDeptCreate`|TODO|TODO|+|`POST /cgi-bin/department/create`
`execDeptUpdate`|TODO|TODO|+|`POST /cgi-bin/department/update`
`execDeptDelete`|TODO|TODO|+|`GET /cgi/bin/department/delete`
`execDeptList`|`reqDeptList`|`respDeptList`|+|`GET /cgi-bin/department/list`

# 标签管理

## API calls

Name|Request Type|Response Type|Access Token|URL
:---|------------|-------------|------------|:--
`execTagCreate`|TODO|TODO|+|`POST /cgi-bin/tag/create`
`execTagUpdate`|TODO|TODO|+|`POST /cgi-bin/tag/update`
`execTagDelete`|TODO|TODO|+|`GET /cgi/bin/tag/delete`
`execTagListUsers`|TODO|TODO|+|`GET /cgi/bin/tag/get`
`execTagAddUsers`|TODO|TODO|+|`POST /cgi/bin/tag/addtagusers`
`execTagDeleteUsers`|TODO|TODO|+|`POST /cgi/bin/tag/deltagusers`
`execTagList`|TODO|TODO|+|`GET /cgi/bin/tag/list`

# 消息推送

## API calls

Name|Request Type|Response Type|Access Token|URL
:---|------------|-------------|------------|:--
`execMessageSend`|`reqMessage`|`respMessageSend`|+|`POST /cgi-bin/message/send`
`execAppchatSend`|`reqMessage`|`respMessageSend`|+|`POST /cgi-bin/appchat/send`
