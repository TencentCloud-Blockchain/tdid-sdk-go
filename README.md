<p align="center">
<a href="https://cloud.tencent.com/"><img src="https://imgcache.qq.com/qcloud/tcloud_dtc/static/static_source_business/eec00e38-a178-479f-83d4-853a18575ac4.png" height="100" ></a>
</p>
<h1 align="center">TDID Tencent Cloud SDK DEMO for Go</h1>

# 目录
1. [简介](#简介)
2. [安装](#获取安装)
3. [快速开始](#快速开始)
4. [相关配置](#相关配置)
5. [功能介绍](#功能介绍)
   - [凭证管理](#凭证管理)
   - [错误处理](#错误处理)
   - [Common Client](#Common-Client)
   - [自定义Header](#自定义-header)
   - [http代理](#http-代理)
   - [请求重试](#请求重试)
   - [空数组和omitempty](#空数组和omitempty)


# 简介

欢迎使用腾讯云TDID API接口DEMO用例。本示例使用腾讯云开发者工具套件（SDK），此 SDK 是云 API 3.0 平台的配套开发工具。
源码地址：https://github.com/tencentcloud/tencentcloud-sdk-go

# 获取安装

## 依赖环境

1. Go 1.9 版本及以上（如使用 go mod 需要 Go 1.14）。
2. 在腾讯云控制台 [访问管理](https://console.cloud.tencent.com/cam/capi) 页面获取密钥 SecretID 和 SecretKey，请务必妥善保管，或者使用更安全的临时安全凭证。


## 通过go get安装（推荐）

推荐使用腾讯云镜像加速下载：

1. Linux 或 MacOS:

    ```bash
    export GOPROXY=https://mirrors.tencent.com/go/
    ```

2. Windows:

    ```cmd
    set GOPROXY=https://mirrors.tencent.com/go/
    ```

## 按需安装（推荐）

注意：此安装方式仅支持使用 **Go Modules** 模式进行依赖管理，即环境变量 `GO111MODULE=auto`或者`GO111MODULE=on`, 并且在您的项目中执行了 `go mod init xxx`.

如果您使用 GOPATH, 请参考下节： 全部安装

v1.0.170后可以按照产品下载，您只需下载基础包和对应的产品包(如cvm)即可，不需要下载全部的产品，从而加快您构建镜像或者编译的速度：

1. 安装公共基础包：

    ```bash
    go get -v -u github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common
    ```

2. 安装对应的产品包(如cvm):

    ```bash
    go get -v -u github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm
    ```

注意：为了支持 go mod，SDK 版本号从 v3.x 降到了 v1.x。并于2021.05.10移除了所有`v3.0.*`和`3.0.*`的tag，如需追溯以前的tag，请参考项目根目录下的 `commit2tag` 文件。

## 通过源码安装

前往代码托管地址 [Github](https://github.com/tencentcloud/tencentcloud-sdk-go) 或者 [Gitee](https://gitee.com/tencentcloud/tencentcloud-sdk-go) 下载最新代码，解压后安装到 $GOPATH/src/github.com/tencentcloud 目录下。

# 快速开始

每个接口都有一个对应的 Request 结构和一个 Response 结构。例如DID标识注册接口 CreateTDidByHost 有对应的请求结构体 CreateTDidByHostRequest 和返回结构体 CreateTDidByHostResponse 。

TDIDDemo目录下是完整的TDID API接口调用单元测试用例，业务可参考相关代码进行编码。

