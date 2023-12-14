package tdid

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
    tdid "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tdid/v20210519"
)

var client *tdid.Client
var cpf *profile.ClientProfile

func init() {
    // 实例化一个认证对象，入参需要传入腾讯云账户密钥对secretId，secretKey。
    credential := common.NewCredential("secretId", "secretKey")
    // 实例化一个客户端配置对象，可以指定超时时间等配置
    cpf = profile.NewClientProfile()
    // 设置签名算法，要求为TC3
    cpf.SignMethod = "TC3-HMAC-SHA256"
    // 设置访问地址
    cpf.HttpProfile.Endpoint = "tdid.tencentcloudapi.com"
    client, _ = tdid.NewClient(credential, "ap-beijing", cpf)
}
