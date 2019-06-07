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
`Status`|`UserStatus`|成员激活状态
`QRCodeURL`|`string`|员工个人二维码；第三方仅通讯录应用可获取<br />扫描可添加为外部联系人

<!-- TODO: extattr external_profile external_position -->

### `UserDeptInfo` 用户部门信息

Name|Type|Doc
:---|:---|:--
`DeptID`|`int64`|部门 ID
`Order`|`uint32`|部门内的排序值，默认为0，数值越大排序越前面
`IsLeader`|`bool`|在所在的部门内是否为上级
