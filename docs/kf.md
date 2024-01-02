# 客服

## Models

### `KfAccount` 客服账号

 Name              | JSON                         | Type     | Doc                                                      
:------------------|:-----------------------------|:---------|:---------------------------------------------------------
 `OpenKfID`        | `open_kfid`                  | `string` | 客服账号ID                                                   
 `Name`            | `name`                       | `string` | 客服名称                                                     
 `Avatar`          | `avatar`                     | `string` | 客服头像URL                                                  
 `ManagePrivilege` | `manage_privilege,omitempty` | `bool`   | 当前调用接口的应用身份，是否有该客服账号的管理权限（编辑客服账号信息、分配会话和收发消息）。组件应用不返回此字段 
