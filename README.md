# 阿里云邮件推送SDK
阿里云官方未封装GO语言SDK，故开发本SDK。
## 说明
1. 本SDK只负责执行请求，得到请求结果。由于请求参数错误(如传入错误的regionId导致的后端错误等等)，请自行判断返回值进行处理。
2. BatchSendEmail需要测试时请替换成您自己的TemplateName等参数。
## 功能
+ [x] SingleSendEmail
+ [x] BatchSendEmail

## 单元测试
1. 配置以下环境变量
    + ACCESS_KEY_ID 阿里云accessKeyId
    + ACCESS_SECRET 阿里云accessKeySecret
    + ACCOUNT_NAME  阿里云accountName
    + FROM_ALIAS    阿里云fromAlias
2. `go test`

