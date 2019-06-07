# 企业微信 API 定义

## API calls

Name|Request Type|Response Type|Access Token|URL
:---|------------|-------------|------------|:--
`execGetAccessToken`|`reqAccessToken`|`respAccessToken`|-|`GET /cgi-bin/gettoken`
`execUserGet`|`reqUserGet`|`respUserGet`|+|`GET /cgi-bin/user/get`
`execDeptList`|`reqDeptList`|`respDeptList`|+|`GET /cgi-bin/department/list`
`execMessageSend`|`reqMessage`|`respMessageSend`|+|`POST /cgi-bin/message/send`
`execAppchatSend`|`reqMessage`|`respMessageSend`|+|`POST /cgi-bin/appchat/send`
