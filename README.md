# go-workwx

[![Travis Build Status](https://img.shields.io/travis/xen0n/go-workwx.svg)](https://travis-ci.org/xen0n/go-workwx)
[![Go Report Card](https://goreportcard.com/badge/github.com/xen0n/go-workwx)](https://goreportcard.com/report/github.com/xen0n/go-workwx)
[![GoDoc](http://godoc.org/github.com/xen0n/go-workwx?status.svg)](http://godoc.org/github.com/xen0n/go-workwx)

```go
import (
    "github.com/xen0n/go-workwx" // package workwx
)
```

Yet another Work Weixin client for Golang

又一个 Golang 企业微信客户端


> English translation TODO for now, as the service covered here is not available
> outside of China (AFAIK).


## Why another wheel?

工作中需要用 Go 实现一个简单的消息推送，想着找个开源库算了，然而现有唯一的开源企业微信 Golang SDK 代码质量不佳。只好自己写一个。


## Features

* [x] access token 刷新
* [ ] 通讯录管理
    - [ ] 成员管理
        - [ ] 创建成员
        - [ ] 读取成员
        - [ ] 更新成员
        - [ ] 删除成员
        - [ ] 批量删除成员
        - [ ] 获取部门成员
        - [ ] 获取部门成员详情
        - [ ] userid与openid互换
        - [ ] 二次验证
        - [ ] 邀请成员
    - [ ] 部门管理
        - [ ] 创建部门
        - [ ] 更新部门
        - [ ] 删除部门
        - [ ] 获取部门列表
    - [ ] 标签管理
        - [ ] 创建标签
        - [ ] 更新标签名字
        - [ ] 删除标签
        - [ ] 获取标签成员
        - [ ] 增加标签成员
        - [ ] 删除标签成员
        - [ ] 获取标签列表
    - [ ] 异步批量接口
        - [ ] 增量更新成员
        - [ ] 全量覆盖成员
        - [ ] 全量覆盖部门
        - [ ] 获取异步任务结果
    - [ ] 通讯录回调通知
        - [ ] 成员变更通知
        - [ ] 部门变更通知
        - [ ] 标签变更通知
        - [ ] 异步任务完成通知
* [ ] 外部联系人管理
    - [ ] 离职成员的外部联系人再分配
    - [ ] 成员对外信息
    - [ ] 获取外部联系人详情
* [ ] 应用管理
    - [ ] 获取应用
    - [ ] 设置应用
    - [ ] 自定义菜单
        - [ ] 创建菜单
        - [ ] 获取菜单
        - [ ] 删除菜单
* [ ] 消息发送
    - [x] 发送应用消息
    - [ ] 接收消息
    - [x] 发送消息到群聊会话
        - [ ] 创建群聊会话
        - [ ] 修改群聊会话
        - [ ] 获取群聊会话
        - [x] 应用推送消息
    - 消息类型
        - [x] 文本消息
        - [ ] 图片消息
        - [ ] 语音消息
        - [ ] 视频消息
        - [ ] 文件消息
        - [ ] 文本卡片消息
        - [ ] 图文消息
        - [ ] 图文消息（mpnews）
        - [x] markdown消息
* [ ] 素材管理


## License

* [MIT](./LICENSE)
