# Access token 获取

## API calls

Name|Request Type|Response Type|Access Token|URL|Doc
:---|------------|-------------|------------|:--|:--
`execGetAccessToken`|`reqAccessToken`|`respAccessToken`|-|`GET /cgi-bin/gettoken`|[获取access_token](https://work.weixin.qq.com/api/doc#90000/90135/91039)

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
`execCorpGetOpenApprovalData`|TODO|TODO|+|`POST /cgi-bin/corp/getopenapprovaldata`|[查询自建应用审批单当前状态](https://work.weixin.qq.com/api/doc#90000/90135/90269)

# 企业支付

## API calls

Name|Request Type|Response Type|Access Token|URL|Doc
:---|------------|-------------|------------|:--|:--

# 电子发票

## API calls

Name|Request Type|Response Type|Access Token|URL|Doc
:---|------------|-------------|------------|:--|:--
