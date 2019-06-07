# Foo bar 模型生成测试

## Models

### `FooTy` type for foo

Name|Type|Doc
:---|:---|:--
`ID`|`int64`|部门 ID
`Name`|`string`|部门名称
`ParentID`|`int64`|父亲部门id。根部门为1
`Order`|`uint32`|在父部门中的次序值。order值大的排序靠前。值范围是[0, 2^32)

### `BarTy` type for bar

Name|Type|Doc|JSON
:---|:---|:--|:---
`ID`|`int64`|ID|`id`
`Foos`|`[]*FooTy`|foos|`foos`

```go
type respBar = BarTy
```

```go
type reqBar struct{}
```

## API calls

Name|Request Type|Response Type|Access Token|URL
:---|------------|-------------|------------|:--
`execUserGet`|`reqUserGet`|`respUserGet`|Y|`GET /cgi-bin/user/get`
`getAccessToken`|`reqAccessToken`|`respAccessToken`|N|`GET /cgi-bin/accesstoken`
