# 上传临时素材

## Models

### `MediaUploadResult` 临时素材上传结果

Name|Type|Doc
:---|:---|:--
`Type`|`string`|媒体文件类型，分别有图片（image）、语音（voice）、视频（video），普通文件(file)
`MediaID`|`string`|媒体文件上传后获取的唯一标识，3天内有效
`CreatedAt`|`time.Time`|媒体文件上传时间戳

```go
// XXX: 由于 sdkcodegen 目前不支持生成 `import` 语句，这个模型不能用 sdkcodegen 生成
// import "time"
```
