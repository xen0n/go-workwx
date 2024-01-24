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

### `KfServicerResult` 客户群列表数据

 Name           | JSON                      | Type     | Doc         
:---------------|:--------------------------|:---------|:------------
 `UserID`       | `userid,omitempty`        | `string` | 接待人员的userid 
 `DepartmentID` | `department_id,omitempty` | `int64`  | 接待人员部门的id   
 `ErrCode`      | `errcode`                 | `int64`  | 该条记录的结果     
 `ErrMsg`       | `errmsg`                  | `string` | 结果信息

```go
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
