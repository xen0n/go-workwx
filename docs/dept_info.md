# 部门管理

## Models

### `DeptInfo` 部门信息

Name|Type|Doc
:---|:---|:--
`ID`|`int64`|部门 ID
`Name`|`string`|部门名称
`ParentID`|`int64`|父亲部门id。根部门为1
`Order`|`uint32`|在父部门中的次序值。order值大的排序靠前。值范围是[0, 2^32)
