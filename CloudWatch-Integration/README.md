# CloudWatch-Integration

## 简介

通过 Lambda 函数实现监听 CloudWatchLogs，特定事件发生时发送飞书消息



## 使用方法

### 编译 GO 文件

```bash
go mod download
GOOS=linux GOARCH=amd64 go build -o main .
zip main.zip main
```



### 上传 Lambda 函数

将刚刚打包出的 main.zip 上传至 AWS Lambda 中。



### 设置环境变量

需要设置一个名为 `WEBHOOK_KEY` 的环境变量，其值为自己的飞书机器人 webhook url 的最后哈希串。

例如某个机器人 webhook 为：`https://open.feishu.cn/open-apis/bot/v2/hook/abcdabcd-aaaa-bbbb-cccc-c4a8aa6ed91c`，那么需要设置一个值为 `abcdabcd-aaaa-bbbb-cccc-c4a8aa6ed91c` 的环境变量 `WEBHOOK_KEY`



### 设置 CloudWatch 关联

在 AWS 中将 CloudWatchLogs 作为 Lambda 的触发器