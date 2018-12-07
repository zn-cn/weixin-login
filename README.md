# 微信登录公共服务

## 基本信息：

+ 入口
  + 生产：`https://<weixin host>/weixin-login/`
  + 接口文档：`https://<api host>/weixin-login/`
+ 域名
  + 生产：`weixin-login.<host>`
+ 镜像
  + 生产：`<name space>/weixin-login:1.0.0`
+ 开发环境
  + 语言：`golang 1.10`
  + 框架：`echo`

## 项目介绍

### 简介

微信登录公共服务

#### 流程：

+ 向公共服务发送请求获取 redirect_uri

  要求query:

  ```
  {
      "index_url": "微信登录之后跳转回的链接，一般为首页"
      "cb_api": "https://test.<host>/api/v1/userInfo" // 回调后台接口，会将用户信息发送到这个接口
  }
  ```
  用户信息如下：

  | 字段       | 示例                                                    |
  | ---------- | ---------------------------------------------------- |
  | country    | China |
  | headimgurl | http://thirdwx.qlogo.cn/mmopen/vi_32/gkmFC8cZJ5Ulhg9sgUHVn5W7VnIhMybvG7b4okHMWRpdgHhh5FSgDYJzXBvULVQaG1We4kaibJicf6J1rJ8UctlA/132 |
  | nickname   | MU   |
  | openid     | <openid>  |
  | province   | Hubei      |
  | sex        | 1   |
  | unionid    | <unionid>        |
  | city       | Wuhan            |

+ 前端根据获取的 redirect_uri，进行跳转

+ 用户确认登录（前端不用管）

+ 微信官方后台跳转到我的后台接口(带着code 和 state )，我的后台根据 code 获取用户信息

+ 获取用户信息之后，设置cookie：openid和unionid，同时设置 cookie 为 httpOnly: true，过期时间为七天，设置完相应的cookie之后进行 302 跳转到最开始 前端要求的 index_url

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

概述：通过公有的微信域名作为对外开放的入口，所以需要在  `https://<weixin host>` 域名下添加相关跳转配置

微信域名下配置如下：

```
    location ~ /weixin-login/(.*)? {
        proxy_set_header Host weixin-login.<host>;
        proxy_pass https://<server ip>/$1$is_args$args;
        proxy_set_header X-Real-IP $remote_addr;
    }
```
