# go-mdy

Go明道云API,封装明道云的API

## v0.1.0

- 获取应用信息
- 获取工作表结构信息
- 获取列表
- 新建行记录

## 快速开始

```shell
go get github.com/gerardhb/go-mdy
```

```go

import (
"mdy github.com/gerardhb/go-mdy"
)

mdyClient := mdy.New("appKey", "sign", "secret option")
// 开启debug
// mdyClient.WithDebug()

req := mdyClient.WorkSheetReq()

appReq := mdyClient.AppReq()

// 通过request调用明道云api接口

```
