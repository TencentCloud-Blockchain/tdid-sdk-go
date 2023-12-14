package tdid

import (
    "encoding/json"
    "fmt"
    "strings"
    "testing"

    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tdid "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tdid/v20210519"
)

//自动生成公私钥创建DID
func TestCreateTDidByHost(t *testing.T) {
    request := tdid.NewCreateTDidByHostRequest()
    request.DAPId = common.Uint64Ptr(25)
    // 通过client对象调用想要访问的接口，需要传入请求对象
    response, err := client.CreateTDidByHost(request)
    // 处理异常
    if _, ok := err.(*errors.TencentCloudSDKError); ok {
        fmt.Printf("An API error has returned: %s", err)
        return
    }
    // 非SDK异常，直接失败。实际代码中可以加入其他的处理。
    if err != nil {
        panic(err)
    }
    // 打印返回的json字符串
    fmt.Printf("%s", response.ToJsonString())
}

//上传本地公钥创建DID
func TestCreateTDidByPukKey(t *testing.T) {
    request := tdid.NewCreateTDidByPubKeyRequest()
    request.DAPId = common.Uint64Ptr(25)
    request.CustomAttribute = common.StringPtr("{\"objectId\": \"test\"}")
    publicKey := "-----BEGIN PUBLIC KEY-----\n" +
        "MFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAEMD2Ieucy2Pwdzb84ytLKoPkJ7PLJ\n" +
        "XrTV1M0hEpMqpFBgCjPrHwS7ENU1vkXIXu8rctNVuXLxYIPXVw4TfyhhQg==\n" +
        "-----END PUBLIC KEY-----"
    request.PublicKey = common.StringPtr(publicKey)
    // 通过client对象调用想要访问的接口，需要传入请求对象
    response, err := client.CreateTDidByPubKey(request)
    // 处理异常
    if _, ok := err.(*errors.TencentCloudSDKError); ok {
        fmt.Printf("An API error has returned: %s", err)
        return
    }
    // 非SDK异常，直接失败。实际代码中可以加入其他的处理。
    if err != nil {
        panic(err)
    }
    // 打印返回的json字符串
    fmt.Printf("%s", response.ToJsonString())
}

//查询DID文档
func TestGetTDidDocument(t *testing.T) {
    request := tdid.NewGetTDidDocumentRequest()
    request.DAPId = common.Uint64Ptr(25)
    request.Did = common.StringPtr("did:tdid:c16:0x15dd70ef48f57196ecbbb32335bb9bf8bfcde191")
    // 通过client对象调用想要访问的接口，需要传入请求对象
    response, err := client.GetTDidDocument(request)
    // 处理异常
    if _, ok := err.(*errors.TencentCloudSDKError); ok {
        fmt.Printf("An API error has returned: %s", err)
        return
    }
    // 非SDK异常，直接失败。实际代码中可以加入其他的处理。
    if err != nil {
        panic(err)
    }
    // 打印返回的json字符串
    fmt.Printf("%s", response.ToJsonString())
}

//查询DID的认证公钥
func TestGetTDidPukKey(t *testing.T) {
    request := tdid.NewGetTDidPubKeyRequest()
    request.DAPId = common.Uint64Ptr(25)
    request.Did = common.StringPtr("did:tdid:c16:0x15dd70ef48f57196ecbbb32335bb9bf8bfcde191")
    // 通过client对象调用想要访问的接口，需要传入请求对象
    response, err := client.GetTDidPubKey(request)
    // 处理异常
    if _, ok := err.(*errors.TencentCloudSDKError); ok {
        fmt.Printf("An API error has returned: %s", err)
        return
    }
    // 非SDK异常，直接失败。实际代码中可以加入其他的处理。
    if err != nil {
        panic(err)
    }
    // 打印返回的json字符串
    fmt.Printf("%s", response.ToJsonString())
}

//根据对象ID查询DID
func TestGetDidByObjectId(t *testing.T) {
    request := &tdid.GetTDidByObjectIdRequest{}
    request.DAPId = common.Uint64Ptr(25)
    request.ObjectId = common.StringPtr("test")
    // 通过client对象调用想要访问的接口，需要传入请求对象
    response, err := client.GetTDidByObjectId(request)
    // 处理异常
    if _, ok := err.(*errors.TencentCloudSDKError); ok {
        fmt.Printf("An API error has returned: %s", err)
        return
    }
    // 非SDK异常，直接失败。实际代码中可以加入其他的处理。
    if err != nil {
        panic(err)
    }
    // 打印返回的json字符串
    fmt.Printf("%s", response.ToJsonString())
}

//设置DID文档的属性
func TestSetDidAttribute(t *testing.T) {
    request := tdid.NewSetTDidAttributeRequest()
    request.DAPId = common.Uint64Ptr(25)
    request.Did = common.StringPtr("did:tdid:c16:0x15dd70ef48f57196ecbbb32335bb9bf8bfcde191")
    request.Attributes = append(request.Attributes, &tdid.DidAttribute{common.StringPtr("a"), common.StringPtr("1")})
    // 通过client对象调用想要访问的接口，需要传入请求对象
    response, err := client.SetTDidAttribute(request)
    // 处理异常
    if _, ok := err.(*errors.TencentCloudSDKError); ok {
        fmt.Printf("An API error has returned: %s", err)
        return
    }
    // 非SDK异常，直接失败。实际代码中可以加入其他的处理。
    if err != nil {
        panic(err)
    }
    // 打印返回的json字符串
    fmt.Printf("%s", response.ToJsonString())
}

//获取权威机构DID
func TestGetAuthorityDid(t *testing.T) {
    request := tdid.NewQueryAuthorityInfoRequest()
    request.DAPId = common.Uint64Ptr(25)
    request.Did = common.StringPtr("did:tdid:c16:0x15dd70ef48f57196ecbbb32335bb9bf8bfcde191")
    // 通过client对象调用想要访问的接口，需要传入请求对象
    response, err := client.QueryAuthorityInfo(request)
    // 处理异常
    if _, ok := err.(*errors.TencentCloudSDKError); ok {
        fmt.Printf("An API error has returned: %s", err)
        return
    }
    // 非SDK异常，直接失败。实际代码中可以加入其他的处理。
    if err != nil {
        panic(err)
    }
    // 打印返回的json字符串
    fmt.Printf("%s", response.ToJsonString())
}

//轮换DID的认证公钥
func TestRotateTDidAuthPubKey(t *testing.T) {
    crdlReq := tdid.NewIssueCredentialRequest()
    crdlReq.DAPId = common.Uint64Ptr(25)
    did := "did:tdid:f34:0x817a3a6d136e8aa8e2c52fe2663afca473d39a12"
    crdlArg := &tdid.CRDLArg{}
    crdlArg.Issuer = &did
    crdlArg.CPTId = common.Uint64Ptr(1)
    crdlArg.ExpirationDate = common.StringPtr("2023-12-01 10:00:00")
    claim := map[string]interface{}{"action": "rotateAuthPublicKey"}
    claim["publickey"] = common.StringPtr("-----BEGIN PUBLIC KEY-----MFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAEMD2Ieucy2Pwdzb84ytLKoPkJ7PLJXrTV1M0hEpMqpFBgCjPrHwS7ENU1vkXIXu8rctNVuXLxYIPXVw4TfyhhQg==-----END PUBLIC KEY-----")
    claim["did"] = common.StringPtr("did:tdid:c780:0x94324460c4785c417a276f762bbf1e39a137e594")
    byteClaim, _ := json.Marshal(claim)
    strClaim := string(byteClaim)
    crdlArg.ClaimJson = &strClaim
    crdlReq.CRDLArg = crdlArg
    // 通过client对象调用想要访问的接口，需要传入请求对象
    crdlResp, err := client.IssueCredential(crdlReq)

    request := tdid.NewRotateTDidAuthPubKeyRequest()
    request.SetScheme("http")
    // 设置访问地址的service
    request.WithApiInfo(strings.Split(cpf.HttpProfile.Endpoint, ".")[0], request.GetVersion(), request.GetAction())
    request.DAPId = common.Uint64Ptr(144)
    request.Did = common.StringPtr("did:tdid:f34:0x817a3a6d136e8aa8e2c52fe2663afca473d39a12")
    request.OperateCredential = crdlResp.Response.CredentialData
    // 通过client对象调用想要访问的接口，需要传入请求对象
    response, err := client.RotateTDidAuthPubKey(request)
    // 处理异常
    if _, ok := err.(*errors.TencentCloudSDKError); ok {
        fmt.Printf("An API error has returned: %s", err)
        return
    }
    // 非SDK异常，直接失败。实际代码中可以加入其他的处理。
    if err != nil {
        panic(err)
    }
    // 打印返回的json字符串
    fmt.Printf("%s", response.ToJsonString())
}

//设置 DID 停用的状态
func TestDeactivateTDid(t *testing.T) {
    crdlReq := tdid.NewIssueCredentialRequest()
    crdlReq.DAPId = common.Uint64Ptr(25)
    did := "did:tdid:f34:0x94324460c4785c417a276f762bbf1e39a137e594"
    crdlArg := &tdid.CRDLArg{}
    crdlArg.Issuer = &did
    crdlArg.CPTId = common.Uint64Ptr(1)
    crdlArg.ExpirationDate = common.StringPtr("2023-12-01 10:00:00")
    claim := map[string]interface{}{"action": "deactiveDid"}
    claim["deactivated"] = common.StringPtr("false")
    claim["did"] = &did
    byteClaim, _ := json.Marshal(claim)
    strClaim := string(byteClaim)
    crdlArg.ClaimJson = &strClaim
    crdlReq.CRDLArg = crdlArg
    // 通过client对象调用想要访问的接口，需要传入请求对象
    crdlResp, err := client.IssueCredential(crdlReq)

    request := tdid.NewDeactivateTDidRequest()
    request.SetScheme("http")
    request.DAPId = common.Uint64Ptr(144)
    request.Did = common.StringPtr("did:tdid:f34:0x94324460c4785c417a276f762bbf1e39a137e594")
    request.OperateCredential = crdlResp.Response.CredentialData
    // 通过client对象调用想要访问的接口，需要传入请求对象
    response, err := client.DeactivateTDid(request)
    // 处理异常
    if _, ok := err.(*errors.TencentCloudSDKError); ok {
        fmt.Printf("An API error has returned: %s", err)
        return
    }
    // 非SDK异常，直接失败。实际代码中可以加入其他的处理。
    if err != nil {
        panic(err)
    }
    // 打印返回的json字符串
    fmt.Printf("%s", response.ToJsonString())
}

func TestCreateCPT(t *testing.T) {
    request := tdid.NewCreateCPTRequest()
    request.DAPId = common.Uint64Ptr(25)
    cptJson := "{\n  \"cptType\": \"original\", \"$schema\": \"http://json-schema.org/draft-04/schema#\", \"description\": \"operate credential\", \"title\": \"operation\", \"type\": \"object\",  \"properties\": {    \"action\": {     \"description\": \"operate action\",      \"type\": \"string\"   } }, \"required\": [    \"action\"  ]}"
    request.CPTJson = common.StringPtr(cptJson)
    // 通过client对象调用想要访问的接口，需要传入请求对象
    response, err := client.CreateCPT(request)
    // 处理异常
    if _, ok := err.(*errors.TencentCloudSDKError); ok {
        fmt.Printf("An API error has returned: %s", err)
        return
    }
    // 非SDK异常，直接失败。实际代码中可以加入其他的处理。
    if err != nil {
        panic(err)
    }
    // 打印返回的json字符串
    fmt.Printf("%s", response.ToJsonString())
}

//创建可验证凭证
func TestIssueCredential(t *testing.T) {
    request := tdid.NewIssueCredentialRequest()
    request.DAPId = common.Uint64Ptr(25)
    did := "did:tdid:c16:0x15dd70ef48f57196ecbbb32335bb9bf8bfcde191"
    crdlArg := &tdid.CRDLArg{}
    crdlArg.Issuer = &did
    crdlArg.CPTId = common.Uint64Ptr(1)
    crdlArg.ExpirationDate = common.StringPtr("2024-12-01 10:00:00")
    claim := map[string]interface{}{"did": did, "action": "test"}
    byteClaim, _ := json.Marshal(claim)
    strClaim := string(byteClaim)
    crdlArg.ClaimJson = &strClaim
    crdlArg.Type = common.StringPtrs([]string{"TestCredential"})
    request.CRDLArg = crdlArg
    // 通过client对象调用想要访问的接口，需要传入请求对象
    response, err := client.IssueCredential(request)
    // 处理异常
    if _, ok := err.(*errors.TencentCloudSDKError); ok {
        fmt.Printf("An API error has returned: %s", err)
        return
    }
    // 非SDK异常，直接失败。实际代码中可以加入其他的处理。
    if err != nil {
        panic(err)
    }
    // 打印返回的json字符串
    fmt.Printf("%s", response.ToJsonString())
}

//验证已签名凭证
func TestVerifyCredential(t *testing.T) {
    request := tdid.NewVerifyCredentialsRequest()
    request.DAPId = common.Uint64Ptr(25)
    crdl := "{\"cptId\":1,\"issuer\":\"did:tdid:c16:0x15dd70ef48f57196ecbbb32335bb9bf8bfcde191\",\"expirationDate\":\"2024-12-01T10:00:00+08:00\",\"issuanceDate\":\"2023-12-14T10:59:05+08:00\",\"context\":\"https://github.com/TencentCloud-Blockchain/TDID/blob/main/context/v1\",\"id\":\"3d5863968bf56eaa0986e8f05e70c10a\",\"type\":[\"VerifiableCredential\",\"TestCredential\"],\"credentialSubject\":{\"action\":\"test\",\"did\":\"did:tdid:c16:0x15dd70ef48f57196ecbbb32335bb9bf8bfcde191\"},\"proof\":{\"created\":\"2023-12-14T10:59:05+08:00\",\"creator\":\"did:tdid:c16:0x15dd70ef48f57196ecbbb32335bb9bf8bfcde191#keys-0\",\"signatureValue\":\"MEYCIQDINYl7wsAIrasOWtXkekTbXczGDZsRQyyHLtZwjy6BcgIhAPBPmuH/Zb/C7Lw/45k+uzg8YgHu3In6422OAOGPj0ro\",\"type\":\"Secp256r1\",\"metaDigest\":\"ef6b2660c1748441c26a311ded5b61f1d116da06965f31f4b48877406f5a0556\",\"vcDigest\":\"d0f9699052b42dce80188bd96b63836eec52135607a0126463fdefef7b050e2b\",\"privacy\":\"Public\",\"salt\":{\"action\":\"3qstH\",\"did\":\"ZQMzx\"}}}"
    request.CredentialData = &crdl
    // 通过client对象调用想要访问的接口，需要传入请求对象
    response, err := client.VerifyCredentials(request)
    // 处理异常
    if _, ok := err.(*errors.TencentCloudSDKError); ok {
        fmt.Printf("An API error has returned: %s", err)
        return
    }
    // 非SDK异常，直接失败。实际代码中可以加入其他的处理。
    if err != nil {
        panic(err)
    }
    // 打印返回的json字符串
    fmt.Printf("%s", response.ToJsonString())
}

//创建选择性披露凭证
func TestCreateDisclosedCredential(t *testing.T) {
    request := tdid.NewCreateDisclosedCredentialRequest()
    request.DAPId = common.Uint64Ptr(25)
    request.CredentialData = common.StringPtr("{\"cptId\":1,\"issuer\":\"did:tdid:c16:0x15dd70ef48f57196ecbbb32335bb9bf8bfcde191\",\"expirationDate\":\"2024-12-01T10:00:00+08:00\",\"issuanceDate\":\"2023-12-14T10:59:05+08:00\",\"context\":\"https://github.com/TencentCloud-Blockchain/TDID/blob/main/context/v1\",\"id\":\"3d5863968bf56eaa0986e8f05e70c10a\",\"type\":[\"VerifiableCredential\",\"TestCredential\"],\"credentialSubject\":{\"action\":\"test\",\"did\":\"did:tdid:c16:0x15dd70ef48f57196ecbbb32335bb9bf8bfcde191\"},\"proof\":{\"created\":\"2023-12-14T10:59:05+08:00\",\"creator\":\"did:tdid:c16:0x15dd70ef48f57196ecbbb32335bb9bf8bfcde191#keys-0\",\"signatureValue\":\"MEYCIQDINYl7wsAIrasOWtXkekTbXczGDZsRQyyHLtZwjy6BcgIhAPBPmuH/Zb/C7Lw/45k+uzg8YgHu3In6422OAOGPj0ro\",\"type\":\"Secp256r1\",\"metaDigest\":\"ef6b2660c1748441c26a311ded5b61f1d116da06965f31f4b48877406f5a0556\",\"vcDigest\":\"d0f9699052b42dce80188bd96b63836eec52135607a0126463fdefef7b050e2b\",\"privacy\":\"Public\",\"salt\":{\"action\":\"3qstH\",\"did\":\"ZQMzx\"}}}")
    request.PolicyJson = common.StringPtr("{\"action\":0}")
    // 通过client对象调用想要访问的接口，需要传入请求对象
    response, err := client.CreateDisclosedCredential(request)
    // 处理异常
    if _, ok := err.(*errors.TencentCloudSDKError); ok {
        fmt.Printf("An API error has returned: %s", err)
        return
    }
    // 非SDK异常，直接失败。实际代码中可以加入其他的处理。
    if err != nil {
        panic(err)
    }
    // 打印返回的json字符串
    fmt.Printf("%s", response.ToJsonString())
}

//获取凭证模板信息
func TestQueryCPT(t *testing.T) {
    request := tdid.NewQueryCPTRequest()
    request.DAPId = common.Uint64Ptr(25)
    request.CPTId = common.Int64Ptr(1)
    // 通过client对象调用想要访问的接口，需要传入请求对象
    response, err := client.QueryCPT(request)
    // 处理异常
    if _, ok := err.(*errors.TencentCloudSDKError); ok {
        fmt.Printf("An API error has returned: %s", err)
        return
    }
    // 非SDK异常，直接失败。实际代码中可以加入其他的处理。
    if err != nil {
        panic(err)
    }
    // 打印返回的json字符串
    fmt.Printf("%s", response.ToJsonString())
}

//更新凭证链上状态
func TestUpdateCredentialStatus(t *testing.T) {
    crdlReq := tdid.NewIssueCredentialRequest()
    crdlReq.DAPId = common.Uint64Ptr(25)
    did := "did:tdid:f34:0x817a3a6d136e8aa8e2c52fe2663afca473d39a12"
    crdlArg := &tdid.CRDLArg{}
    crdlArg.Issuer = &did
    crdlArg.CPTId = common.Uint64Ptr(1)
    crdlArg.ExpirationDate = common.StringPtr("2023-12-01 10:00:00")
    claim := map[string]interface{}{"action": "updateCredentialState"}
    claim["CredentialStatus"] = map[string]interface{}{"id": "c372873f65ac012d445c1caa974fa2fe", "issuer": "did:tdid:c780:0x94324460c4785c417a276f762bbf1e39a137e594", "status": 1}
    claim["orignCredential"] = "{\"cptId\":1,\"issuer\":\"did:tdid:c780:0x94324460c4785c417a276f762bbf1e39a137e594\",\"expirationDate\":\"2023-12-01T10:00:00+08:00\",\"issuanceDate\":\"2023-03-13T17:24:32+08:00\",\"context\":\"https://github.com/TencentCloud-Blockchain/TDID/blob/main/context/v1\",\"id\":\"c372873f65ac012d445c1caa974fa2fe\",\"type\":[\"VerifiableCredential\",\"TestCredential\"],\"credentialSubject\":{\"action\":\"test\",\"did\":\"did:tdid:c780:0x94324460c4785c417a276f762bbf1e39a137e594\"},\"proof\":{\"created\":\"2023-03-13T17:24:32+08:00\",\"creator\":\"did:tdid:c780:0x94324460c4785c417a276f762bbf1e39a137e594#keys-0\",\"signatureValue\":\"MEQCIB1a184Aa8UA1mEvNmEKZrZ2iKzX4JDr7OOubnnP1ZSSAiB/+vshYgr9QHoMTKLJAek08oGD2i7NJ4sIqp3oO0vy2Q==\",\"type\":\"Sm2p256v1\",\"txDegist\":\"6648db222e53bb315e813145ac4d9ee917f28f01367d7dab8683d2acc4ac9684\",\"vcDegist\":\"0ccd42f2995c0ff296c18b645839b0ef020f8ccf9ce532fb4768bc7ed6a9755c\",\"method\":\"Public\",\"salt\":{\"action\":\"6MOqr\",\"did\":\"qMy2t\"}}}"
    byteClaim, _ := json.Marshal(claim)
    strClaim := string(byteClaim)
    crdlArg.ClaimJson = &strClaim
    crdlReq.CRDLArg = crdlArg
    // 通过client对象调用想要访问的接口，需要传入请求对象
    crdlResp, err := client.IssueCredential(crdlReq)

    request := tdid.NewUpdateCredentialStateRequest()
    request.SetScheme("http")
    // 设置访问地址的service
    request.WithApiInfo(strings.Split(cpf.HttpProfile.Endpoint, ".")[0], request.GetVersion(), request.GetAction())
    request.DAPId = common.Uint64Ptr(25)
    request.OperateCredential = crdlResp.Response.CredentialData
    // 通过client对象调用想要访问的接口，需要传入请求对象
    response, err := client.UpdateCredentialState(request)
    // 处理异常
    if _, ok := err.(*errors.TencentCloudSDKError); ok {
        fmt.Printf("An API error has returned: %s", err)
        return
    }
    // 非SDK异常，直接失败。实际代码中可以加入其他的处理。
    if err != nil {
        panic(err)
    }
    // 打印返回的json字符串
    fmt.Printf("%s", response.ToJsonString())
}

//获取凭证链上状态
func TestGetCredentialStatus(t *testing.T) {
    request := tdid.NewGetCredentialStateRequest()
    request.DAPId = common.Uint64Ptr(25)
    request.CredentialId = common.StringPtr("")
    // 通过client对象调用想要访问的接口，需要传入请求对象
    response, err := client.GetCredentialState(request)
    // 处理异常
    if _, ok := err.(*errors.TencentCloudSDKError); ok {
        fmt.Printf("An API error has returned: %s", err)
        return
    }
    // 非SDK异常，直接失败。实际代码中可以加入其他的处理。
    if err != nil {
        panic(err)
    }
    // 打印返回的json字符串
    fmt.Printf("%s", response.ToJsonString())
}

//创建可验证表达
func TestCreatePresentation(t *testing.T) {
    request := tdid.NewCreatePresentationRequest()
    request.DAPId = common.Uint64Ptr(25)
    request.Credentials = []*string{common.StringPtr("{\"cptId\":1,\"issuer\":\"did:tdid:c780:0x94324460c4785c417a276f762bbf1e39a137e594\",\"expirationDate\":\"2023-12-01T10:00:00+08:00\",\"issuanceDate\":\"2023-03-13T17:24:32+08:00\",\"context\":\"https://github.com/TencentCloud-Blockchain/TDID/blob/main/context/v1\",\"id\":\"c372873f65ac012d445c1caa974fa2fe\",\"type\":[\"VerifiableCredential\",\"TestCredential\"],\"credentialSubject\":{\"action\":\"test\",\"did\":\"did:tdid:c780:0x94324460c4785c417a276f762bbf1e39a137e594\"},\"proof\":{\"created\":\"2023-03-13T17:24:32+08:00\",\"creator\":\"did:tdid:c780:0x94324460c4785c417a276f762bbf1e39a137e594#keys-0\",\"signatureValue\":\"MEQCIB1a184Aa8UA1mEvNmEKZrZ2iKzX4JDr7OOubnnP1ZSSAiB/+vshYgr9QHoMTKLJAek08oGD2i7NJ4sIqp3oO0vy2Q==\",\"type\":\"Sm2p256v1\",\"txDegist\":\"6648db222e53bb315e813145ac4d9ee917f28f01367d7dab8683d2acc4ac9684\",\"vcDegist\":\"0ccd42f2995c0ff296c18b645839b0ef020f8ccf9ce532fb4768bc7ed6a9755c\",\"method\":\"Public\",\"salt\":{\"action\":\"6MOqr\",\"did\":\"qMy2t\"}}}")}
    request.Did = common.StringPtr("did:tdid:c780:0x94324460c4785c417a276f762bbf1e39a137e594")
    request.VerifyCode = common.StringPtr("test")
    // 通过client对象调用想要访问的接口，需要传入请求对象
    response, err := client.CreatePresentation(request)
    // 处理异常
    if _, ok := err.(*errors.TencentCloudSDKError); ok {
        fmt.Printf("An API error has returned: %s", err)
        return
    }
    // 非SDK异常，直接失败。实际代码中可以加入其他的处理。
    if err != nil {
        panic(err)
    }
    // 打印返回的json字符串
    fmt.Printf("%s", response.ToJsonString())
}

//验证可验证表达
func TestVerifyPresentation(t *testing.T) {
    request := tdid.NewVerifyPresentationRequest()
    request.DAPId = common.Uint64Ptr(25)
    request.Did = common.StringPtr("did:tdid:c780:0x94324460c4785c417a276f762bbf1e39a137e594")
    request.PresentationData = common.StringPtr("{\"context\":[\"https://github.com/TencentCloud-Blockchain/TDID/blob/main/context/v1\"],\"holder\":\"did:tdid:c130:0x7f6e86656286ddfe1d1089062d87397d97022b43\",\"proof\":{\"created\":\"2023-08-11T17:10:48+08:00\",\"nonce\":\"test\",\"signatureValue\":\"MEYCIQCepjJpP5mhQQ9D2GIoelwbzAwEk9JUiPwhzixR0NWIlgIhANtvKnrc9AIEt3C6MtPMT55EyN8liz1g2dIL87Zns/LW\",\"type\":\"Secp256r1\",\"verificationMethod\":\"did:tdid:c130:0x7f6e86656286ddfe1d1089062d87397d97022b43#keys-0\"},\"type\":[\"VerifiablePresentation\"],\"verifiableCredential\":[{\"context\":\"https://github.com/TencentCloud-Blockchain/TDID/blob/main/context/v1\",\"cptId\":1,\"credentialSubject\":{\"action\":\"0x50ded9f7adba12497fcc010105ef02a0b5aaca791298c2d483c08692384ae04f\"},\"expirationDate\":\"2024-06-29T15:25:00+08:00\",\"id\":\"d8c50122aa6f9e35dd5d29539d2c71ba\",\"issuanceDate\":\"2023-08-11T16:56:52+08:00\",\"issuer\":\"did:tdid:c130:0x7f6e86656286ddfe1d1089062d87397d97022b43\",\"proof\":{\"created\":\"2023-08-11T16:56:52+08:00\",\"creator\":\"did:tdid:c130:0x7f6e86656286ddfe1d1089062d87397d97022b43#keys-0\",\"metaDigest\":\"7f4141c6b6afbbc5eba879a99dde396ef77d0589e36b6a3131a05b9e0d61ab24\",\"privacy\":\"Public\",\"salt\":{\"action\":\"0\"},\"signatureValue\":\"MEUCIBArCupGI+aEln8FzSRXmWARz6lv2XteVOStty1lEqXrAiEA6CCVRjbtdrZ9XsdjPi7ClUoNvlfnC+Wbg/c1oxGgLv0=\",\"type\":\"Secp256r1\",\"vcDigest\":\"73a52f0aca2b2ba21b902f23934f5d79b626f024d6d32b53e2d0f143b5b1c5ca\"},\"type\":[\"VerifiableCredential\"]}]}")
    request.VerifyCode = common.StringPtr("test")
    // 通过client对象调用想要访问的接口，需要传入请求对象
    response, err := client.VerifyPresentation(request)
    // 处理异常
    if _, ok := err.(*errors.TencentCloudSDKError); ok {
        fmt.Printf("An API error has returned: %s", err)
        return
    }
    // 非SDK异常，直接失败。实际代码中可以加入其他的处理。
    if err != nil {
        panic(err)
    }
    // 打印返回的json字符串
    fmt.Printf("%s", response.ToJsonString())
}

//获取某个应用关键指标
func TestGetAppSummary(t *testing.T) {
    request := tdid.NewGetAppSummaryRequest()
    request.DAPId = common.Uint64Ptr(25)
    // 通过client对象调用想要访问的接口，需要传入请求对象
    response, err := client.GetAppSummary(request)
    // 处理异常
    if _, ok := err.(*errors.TencentCloudSDKError); ok {
        fmt.Printf("An API error has returned: %s", err)
        return
    }
    // 非SDK异常，直接失败。实际代码中可以加入其他的处理。
    if err != nil {
        panic(err)
    }
    // 打印返回的json字符串
    fmt.Printf("%s", response.ToJsonString())
}

//获取用户应用关键指标
func TestGetOverSummary(t *testing.T) {
    request := tdid.NewGetOverSummaryRequest()
    // 通过client对象调用想要访问的接口，需要传入请求对象
    response, err := client.GetOverSummary(request)
    // 处理异常
    if _, ok := err.(*errors.TencentCloudSDKError); ok {
        fmt.Printf("An API error has returned: %s", err)
        return
    }
    // 非SDK异常，直接失败。实际代码中可以加入其他的处理。
    if err != nil {
        panic(err)
    }
    // 打印返回的json字符串
    fmt.Printf("%s", response.ToJsonString())
}
