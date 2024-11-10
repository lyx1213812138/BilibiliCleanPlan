# Bilbili Clean Plan
全栈web项目，前端基于 vue，arcodesign，后端基于 go， mongodb。通过API获取Up主，合集，及视频列表，存储到Mongodb中，筛选Up主和合集后自动推荐视频。

## 本地部署
- 下载Mongodb

- 进入 `BackEnd` 目录

- 新建并设置 `myCookie.txt`， `myVmid.txt`。cookie就在b站请求里找一个长一点的。vmid可以看b站自己空间页面的地址。

- 运行 `go run server.go test.go`

- 访问 https://localhost:12121

## 鸣谢
API文档 https://github.com/SocialSisterYi/bilibili-API-collect

