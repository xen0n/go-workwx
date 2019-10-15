# 用户管理

## Models

### `UserInfo` 用户信息

Name|Type|Doc
:---|:---|:--
`UserID`|`string`|成员UserID<br />对应管理端的账号，企业内必须唯一。不区分大小写，长度为1~64个字节
`Name`|`string`|成员名称
`Position`|`string`|职务信息；第三方仅通讯录应用可获取
`Departments`|`[]UserDeptInfo`|成员所属部门信息
`Mobile`|`string`|手机号码；第三方仅通讯录应用可获取
`Gender`|`UserGender`|性别
`Email`|`string`|邮箱；第三方仅通讯录应用可获取
`AvatarURL`|`string`|头像 URL；第三方仅通讯录应用可获取<br />NOTE：如果要获取小图将url最后的”/0”改成”/100”即可。
`Telephone`|`string`|座机；第三方仅通讯录应用可获取
`IsEnabled`|`bool`|成员的启用状态
`Alias`|`string`|别名；第三方仅通讯录应用可获取
`ExtAttr`|TODO|扩展属性，第三方仅通讯录应用可获取
`Status`|`UserStatus`|成员激活状态
`QRCodeURL`|`string`|员工个人二维码；第三方仅通讯录应用可获取<br />扫描可添加为外部联系人
`ExternalProfile`|TODO|成员对外属性，字段详情见对外属性；第三方仅通讯录应用可获取
`ExternalPosition`|TODO|对外职务，如果设置了该值，则以此作为对外展示的职务，否则以position来展示。

```go
// UserGender 用户性别
type UserGender int

const (
	// UserGenderUnspecified 性别未定义
	UserGenderUnspecified UserGender = 0
	// UserGenderMale 男性
	UserGenderMale UserGender = 1
	// UserGenderFemale 女性
	UserGenderFemale UserGender = 2
)

// UserStatus 用户激活信息
//
// 已激活代表已激活企业微信或已关注微工作台（原企业号）。
// 未激活代表既未激活企业微信又未关注微工作台（原企业号）。
type UserStatus int

const (
	// UserStatusActivated 已激活
	UserStatusActivated UserStatus = 1
	// UserStatusDeactivated 已禁用
	UserStatusDeactivated UserStatus = 2
	// UserStatusUnactivated 未激活
	UserStatusUnactivated UserStatus = 4
)
```

### `UserDeptInfo` 用户部门信息

Name|Type|Doc
:---|:---|:--
`DeptID`|`int64`|部门 ID
`Order`|`uint32`|部门内的排序值，默认为0，数值越大排序越前面
`IsLeader`|`bool`|在所在的部门内是否为上级
