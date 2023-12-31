# SNS-Integration
## 简介

通过 `AWS Lambda` 函数将 `AWS SNS` 消息发送给飞书。



## 使用方法

### 编译 GO 文件

 ```bash
 git clone https://github.com/xiabee/AWS-Lark-Bot.git && cd AWS-Lark-Bot
 go mod download
 
 GOOS=linux GOARCH=amd64 go build -o main .
 zip main.zip main
 ```



### 上传 Lambda 函数

将刚刚打包出的 main.zip 上传至 AWS Lambda 中。



### 设置环境变量

| 环境变量    | 类型    | 值                                  | 作用                       |
| ----------- | ------- | ----------------------------------- | -------------------------- |
| WEBHOOK_KEY | string  | 飞书机器人 webhook url 的最后哈希串 | 触发飞书 webhook 发送消息  |
| ALERT_LEVEL | float64 | 安全事件等级触发告警的阈值          | 超出该值的安全事件触发告警 |



* 需要设置一个名为 `WEBHOOK_KEY` 的环境变量，其值为自己的飞书机器人 webhook url 的最后哈希串。
  * 例如某个机器人 webhook 为：`https://open.feishu.cn/open-apis/bot/v2/hook/abcdabcd-aaaa-bbbb-cccc-c4a8aa6ed91c`，那么需要设置一个值为 `abcdabcd-aaaa-bbbb-cccc-c4a8aa6ed91c` 的环境变量 `WEBHOOK_KEY`

* 需要设置一个名为 `ALERT_LEVEL` 的环境变量，其值为告警等级的阈值（float64类型），超出该值的事件才出发告警

### 设置 SNS 关联

在 AWS 中将 SNS 消息于刚刚的 Lambda 函数关联
