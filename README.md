# 微信登录公共服务

## 基本信息：

+ 入口
  + 生产：`https://weixin.bingyan-tech.hustonline.net/weixin-login/`
  + 接口文档：`https://api.bingyan.net/weixin-login/`
+ 域名
  + 生产：`weixin-login.hustonline.net`
+ 镜像
  + 生产：`registry.cn-hangzhou.aliyuncs.com/bingyan/weixin-login:1.0.0`
+ 开发环境
  + 语言：`golang 1.10`
  + 框架：`echo`
+ 项目部署位置
  + server：`阿里云-53 /mnt/var/www/tofar/weixin-login`
+ 开发人员
  + 后台：赵楠

## 项目介绍

### 简介

微信登录公共服务

#### 流程：

+ 向公共服务发送请求获取 redirect_uri

  要求query:

  ```
  {
      "index_url": "微信登录之后跳转回的链接，一般为首页"
  }
  ```

+ 前端根据获取的 redirect_uri，进行跳转

+ 用户确认登录（前端不用管）

+ 微信官方后台跳转到我的后台接口(带着code 和 state )，我的后台根据 code 获取用户信息

+ 获取用户信息之后，存储在cookie里面，同时设置 cookie 为 httpOnly: true，过期时间为七天，设置完相应的cookie之后进行 302 跳转到最开始 前端要求的 index_url

+ 后续可由项目的后台直接根据 cookie 取出用户信息传给前端 (httpOnly 前端无法读取cookie)

  > 本微信公共服务已提供接口获取用户信息给前端，但是不包含标识用户敏感信息，如：openid, unionid

### 开发简介

+ API 文档

  使用 apidoc 从代码生成可视化 API 文档，文档链接: https://api.bingyan.net/weixin-login/

  包含:

  + route
  + version
  + title
  + description
  + param
  + param example
  + success data
  + success example
  + error param
  + error example

+ 日志

  logrus，生产版本使用 json 格式log，测试或者本地版本使用 text 格式

  格式：

  eg:

  ```
  {
      "package": "model",
      "file": "ticket.go",
      "function": "GetTicketsByShopName",
      "err": error,
      "msg": "err_msg"
  }
  ```

+ response 格式

  ```go
  // ErrorRes ErrorResponse
  type ErrorRes struct {
  	St atus int    `json:"status"`  // http.Status
  	ErrMsg string `json:"err_msg"`
  }

  // DataRes DataResponse
  type DataRes struct {
  	Status int         `json:"status"` // 200
  	Data   interface{} `json:"data"`
  }
  ```
  status

  状态码采用[标准 http 状态码](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status) (中文：[http 响应代码](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Status))

### 文件简介

+ client 前端文件

+ deploy 项目部署

  + docker-compose
  + nginx

+ docs 文档

  + API
  + iteration 迭代文档
  + record 记录

+ dist 静态文件

  + apidoc

    可视化 API 文档

+ src 源码

+ README.md

+ restart.sh 重启项目

## 附录

### docker-compose 说明

+ 环境变量
  + 环境变量 APP_ADDR=:6458 项目启动地址
  + WEIXIN_APPID=WEIXIN_APPID 微信 appid
  + WEIXIN_APPSECRET=WEIXIN_APPSECRET 微信 appsecret

### nginx 说明

概述：通过公有的微信域名作为对外开放的入口，所以需要在  `https://weixin.bingyan-tech.hustonline.net` 域名下添加相关跳转配置

微信域名下配置如下：

```
    location ~ /weixin-login/(.*)? {
        proxy_set_header Host weixin-login.hustonline.net;
        proxy_pass https://120.78.203.153/$1$is_args$args;
        proxy_set_header X-Real-IP $remote_addr;
    }
```
