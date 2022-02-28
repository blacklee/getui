# Golang GeTui RestAPI V2

个推(Getui) RestAPI V2 的 golang 实现

官方文档：[https://docs.getui.com/getui/server/rest_v2/introduction/](https://docs.getui.com/getui/server/rest_v2/introduction/)

0.1版 只实现了一点点「单推到iOS客户端」的功能，够我简单使用了，所以可能会停留一段时间。

## 体验步骤

### 鉴权

1. 按照 `test.json.example` 的例子，创建一个 `test.json` 的文件，先配置里面的`config`部分。
2. 执行 `auth_test.go` 里的测试代码，最终会输出 `auth token` 的具体值。
3. 把步骤2获得的`token`填入`test.json`里的`push`部分，同时填入自己相关的`cid`。

### 单推到iOS

1. 执行 `single_test.go`里的测试代码，iOS客户端应该就能收到推送消息了。

## Versions

### Version 0.1

完成了单推到 iOS 客户端的功能。
