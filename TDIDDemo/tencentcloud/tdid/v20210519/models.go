// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20210519

import (
	"encoding/json"

	tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type CRDLArg struct {
	// CPT ID
	CPTId *uint64 `json:"CPTId,omitnil" name:"CPTId"`

	// 签发者 did
	Issuer *string `json:"Issuer,omitnil" name:"Issuer"`

	// 过期时间
	ExpirationDate *string `json:"ExpirationDate,omitnil" name:"ExpirationDate"`

	// 声明
	ClaimJson *string `json:"ClaimJson,omitnil" name:"ClaimJson"`

	// 凭证类型
	Type []*string `json:"Type,omitnil" name:"Type"`

	// 多方签名的用户did
	Parties []*string `json:"Parties,omitnil" name:"Parties"`
}

type ChainTransaction struct {
	// 交易哈希
	TransactionHash *string `json:"TransactionHash,omitnil" name:"TransactionHash"`
}

type CommitmentClaim struct {
	// 致盲因子
	Opening *string `json:"Opening,omitnil" name:"Opening"`

	// 原始值声明字段名称
	Origin *string `json:"Origin,omitnil" name:"Origin"`

	// 承诺声明字段名称
	Commitment *string `json:"Commitment,omitnil" name:"Commitment"`
}

// Predefined struct for user
type CreateCPTRequestParams struct {
	// CPT模板的JSON字符串
	CPTJson *string `json:"CPTJson,omitnil" name:"CPTJson"`

	// CPT ID
	CPTId *uint64 `json:"CPTId,omitnil" name:"CPTId"`

	// 用户应用ID
	UAPId *uint64 `json:"UAPId,omitnil" name:"UAPId"`

	// DID应用ID
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`
}

type CreateCPTRequest struct {
	*tchttp.BaseRequest

	// CPT模板的JSON字符串
	CPTJson *string `json:"CPTJson,omitnil" name:"CPTJson"`

	// CPT ID
	CPTId *uint64 `json:"CPTId,omitnil" name:"CPTId"`

	// 用户应用ID
	UAPId *uint64 `json:"UAPId,omitnil" name:"UAPId"`

	// DID应用ID
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`
}

func (r *CreateCPTRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCPTRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CPTJson")
	delete(f, "CPTId")
	delete(f, "UAPId")
	delete(f, "DAPId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCPTRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCPTResponseParams struct {
	// 上链交易信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Transaction *ChainTransaction `json:"Transaction,omitnil" name:"Transaction"`

	// 凭证模板索引ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 凭证模板ID
	CPTId *uint64 `json:"CPTId,omitnil" name:"CPTId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateCPTResponse struct {
	*tchttp.BaseResponse
	Response *CreateCPTResponseParams `json:"Response"`
}

func (r *CreateCPTResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCPTResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDisclosedCredentialRequestParams struct {
	// DID应用ID
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// 披露策略id，PolicyJson和PolicyId任选其一
	PolicyId *uint64 `json:"PolicyId,omitnil" name:"PolicyId"`

	// 凭证文本内容，FunctionArg和CredentialText任选其一
	CredentialData *string `json:"CredentialData,omitnil" name:"CredentialData"`

	// 披露策略文本
	PolicyJson *string `json:"PolicyJson,omitnil" name:"PolicyJson"`
}

type CreateDisclosedCredentialRequest struct {
	*tchttp.BaseRequest

	// DID应用ID
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// 披露策略id，PolicyJson和PolicyId任选其一
	PolicyId *uint64 `json:"PolicyId,omitnil" name:"PolicyId"`

	// 凭证文本内容，FunctionArg和CredentialText任选其一
	CredentialData *string `json:"CredentialData,omitnil" name:"CredentialData"`

	// 披露策略文本
	PolicyJson *string `json:"PolicyJson,omitnil" name:"PolicyJson"`
}

func (r *CreateDisclosedCredentialRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDisclosedCredentialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DAPId")
	delete(f, "PolicyId")
	delete(f, "CredentialData")
	delete(f, "PolicyJson")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDisclosedCredentialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDisclosedCredentialResponseParams struct {
	// 凭证字符串
	CredentialData *string `json:"CredentialData,omitnil" name:"CredentialData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateDisclosedCredentialResponse struct {
	*tchttp.BaseResponse
	Response *CreateDisclosedCredentialResponseParams `json:"Response"`
}

func (r *CreateDisclosedCredentialResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDisclosedCredentialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePresentationRequestParams struct {
	// DID应用id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// VP持有人的DID标识
	Did *string `json:"Did,omitnil" name:"Did"`

	// 选择性披露策略
	PolicyJson *string `json:"PolicyJson,omitnil" name:"PolicyJson"`

	// 凭证列表
	Credentials []*string `json:"Credentials,omitnil" name:"Credentials"`

	// VP随机验证码
	VerifyCode *string `json:"VerifyCode,omitnil" name:"VerifyCode"`

	// 是否签名，ture时signatureValue为待签名内容由调用端自行签名，false时signatureValue为平台自动已签名的内容。默认false
	Unsigned *bool `json:"Unsigned,omitnil" name:"Unsigned"`

	// 可验证凭证证明列表
	CredentialList []*CredentialProof `json:"CredentialList,omitnil" name:"CredentialList"`
}

type CreatePresentationRequest struct {
	*tchttp.BaseRequest

	// DID应用id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// VP持有人的DID标识
	Did *string `json:"Did,omitnil" name:"Did"`

	// 选择性披露策略
	PolicyJson *string `json:"PolicyJson,omitnil" name:"PolicyJson"`

	// 凭证列表
	Credentials []*string `json:"Credentials,omitnil" name:"Credentials"`

	// VP随机验证码
	VerifyCode *string `json:"VerifyCode,omitnil" name:"VerifyCode"`

	// 是否签名，ture时signatureValue为待签名内容由调用端自行签名，false时signatureValue为平台自动已签名的内容。默认false
	Unsigned *bool `json:"Unsigned,omitnil" name:"Unsigned"`

	// 可验证凭证证明列表
	CredentialList []*CredentialProof `json:"CredentialList,omitnil" name:"CredentialList"`
}

func (r *CreatePresentationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePresentationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DAPId")
	delete(f, "Did")
	delete(f, "PolicyJson")
	delete(f, "Credentials")
	delete(f, "VerifyCode")
	delete(f, "Unsigned")
	delete(f, "CredentialList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePresentationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePresentationResponseParams struct {
	// 可验证表达内容
	PresentationData *string `json:"PresentationData,omitnil" name:"PresentationData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreatePresentationResponse struct {
	*tchttp.BaseResponse
	Response *CreatePresentationResponseParams `json:"Response"`
}

func (r *CreatePresentationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePresentationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTDidByHostRequestParams struct {
	// DID应用ID
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// 自定义DID文档json属性
	CustomAttribute *string `json:"CustomAttribute,omitnil" name:"CustomAttribute"`

	// 用户应用ID
	UAPId *uint64 `json:"UAPId,omitnil" name:"UAPId"`
}

type CreateTDidByHostRequest struct {
	*tchttp.BaseRequest

	// DID应用ID
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// 自定义DID文档json属性
	CustomAttribute *string `json:"CustomAttribute,omitnil" name:"CustomAttribute"`

	// 用户应用ID
	UAPId *uint64 `json:"UAPId,omitnil" name:"UAPId"`
}

func (r *CreateTDidByHostRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTDidByHostRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DAPId")
	delete(f, "CustomAttribute")
	delete(f, "UAPId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTDidByHostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTDidByHostResponseParams struct {
	// TDID标识
	Did *string `json:"Did,omitnil" name:"Did"`

	// 链上返回交易信息
	Transaction *ChainTransaction `json:"Transaction,omitnil" name:"Transaction"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateTDidByHostResponse struct {
	*tchttp.BaseResponse
	Response *CreateTDidByHostResponseParams `json:"Response"`
}

func (r *CreateTDidByHostResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTDidByHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTDidByPubKeyRequestParams struct {
	// pem格式的认证公钥
	PublicKey *string `json:"PublicKey,omitnil" name:"PublicKey"`

	// DID应用id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// pem格式的加密公钥
	EncryptPubKey *string `json:"EncryptPubKey,omitnil" name:"EncryptPubKey"`

	// 自定义DID初始化属性json字符串
	CustomAttribute *string `json:"CustomAttribute,omitnil" name:"CustomAttribute"`

	// 0:did存在返回错误，1:did存在返回该did，默认:0
	IgnoreExisted *uint64 `json:"IgnoreExisted,omitnil" name:"IgnoreExisted"`

	// 用户应用id
	UAPId *uint64 `json:"UAPId,omitnil" name:"UAPId"`
}

type CreateTDidByPubKeyRequest struct {
	*tchttp.BaseRequest

	// pem格式的认证公钥
	PublicKey *string `json:"PublicKey,omitnil" name:"PublicKey"`

	// DID应用id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// pem格式的加密公钥
	EncryptPubKey *string `json:"EncryptPubKey,omitnil" name:"EncryptPubKey"`

	// 自定义DID初始化属性json字符串
	CustomAttribute *string `json:"CustomAttribute,omitnil" name:"CustomAttribute"`

	// 0:did存在返回错误，1:did存在返回该did，默认:0
	IgnoreExisted *uint64 `json:"IgnoreExisted,omitnil" name:"IgnoreExisted"`

	// 用户应用id
	UAPId *uint64 `json:"UAPId,omitnil" name:"UAPId"`
}

func (r *CreateTDidByPubKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTDidByPubKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PublicKey")
	delete(f, "DAPId")
	delete(f, "EncryptPubKey")
	delete(f, "CustomAttribute")
	delete(f, "IgnoreExisted")
	delete(f, "UAPId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTDidByPubKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTDidByPubKeyResponseParams struct {
	// did具体信息
	Did *string `json:"Did,omitnil" name:"Did"`

	// 链上返回交易信息
	Transaction *ChainTransaction `json:"Transaction,omitnil" name:"Transaction"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateTDidByPubKeyResponse struct {
	*tchttp.BaseResponse
	Response *CreateTDidByPubKeyResponseParams `json:"Response"`
}

func (r *CreateTDidByPubKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTDidByPubKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CredentialProof struct {
	// 可验证凭证内容
	Credential *string `json:"Credential,omitnil" name:"Credential"`

	// 零知识范围证明阈值信息
	ProofThreshold []*ProofThreshold `json:"ProofThreshold,omitnil" name:"ProofThreshold"`
}

type CredentialState struct {
	// 凭证唯一id
	Id *string `json:"Id,omitnil" name:"Id"`

	// 凭证状态（0：吊销；1：有效）
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 凭证颁发者Did
	Issuer *string `json:"Issuer,omitnil" name:"Issuer"`

	// VC摘要，对应凭证Proof的vcDigest字段
	VCDigest *string `json:"VCDigest,omitnil" name:"VCDigest"`

	// 交易摘要，对应凭证Proof的txDigest字段
	TXDigest *string `json:"TXDigest,omitnil" name:"TXDigest"`

	// 颁布凭证的UTC时间戳
	IssueTime *uint64 `json:"IssueTime,omitnil" name:"IssueTime"`

	// 凭证过期的UTC时间戳
	ExpireTime *uint64 `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// 凭证模板id
	CPTId *uint64 `json:"CPTId,omitnil" name:"CPTId"`

	// 凭证签名
	Signature *string `json:"Signature,omitnil" name:"Signature"`

	// 元数据摘要
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetaDigest *string `json:"MetaDigest,omitnil" name:"MetaDigest"`
}

// Predefined struct for user
type DeactivateTDidRequestParams struct {
	// DID标识符
	Did *string `json:"Did,omitnil" name:"Did"`

	// 设置DID禁用状态的临时凭证内容，通过创建凭证接口(CreateCredential)生成并签名，凭证类型为：OperateCredential, 为安全起见凭证过期时间不适合太长，建议设置为1分钟内
	OperateCredential *string `json:"OperateCredential,omitnil" name:"OperateCredential"`

	// DID应用id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`
}

type DeactivateTDidRequest struct {
	*tchttp.BaseRequest

	// DID标识符
	Did *string `json:"Did,omitnil" name:"Did"`

	// 设置DID禁用状态的临时凭证内容，通过创建凭证接口(CreateCredential)生成并签名，凭证类型为：OperateCredential, 为安全起见凭证过期时间不适合太长，建议设置为1分钟内
	OperateCredential *string `json:"OperateCredential,omitnil" name:"OperateCredential"`

	// DID应用id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`
}

func (r *DeactivateTDidRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeactivateTDidRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Did")
	delete(f, "OperateCredential")
	delete(f, "DAPId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeactivateTDidRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeactivateTDidResponseParams struct {
	// 链上返回交易信息
	Transaction *ChainTransaction `json:"Transaction,omitnil" name:"Transaction"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeactivateTDidResponse struct {
	*tchttp.BaseResponse
	Response *DeactivateTDidResponseParams `json:"Response"`
}

func (r *DeactivateTDidResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeactivateTDidResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DidAttribute struct {
	// 键名
	Key *string `json:"Key,omitnil" name:"Key"`

	// 键值
	Val *string `json:"Val,omitnil" name:"Val"`
}

// Predefined struct for user
type GetAppSummaryRequestParams struct {
	// DID应用id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`
}

type GetAppSummaryRequest struct {
	*tchttp.BaseRequest

	// DID应用id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`
}

func (r *GetAppSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAppSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DAPId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetAppSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAppSummaryResponseParams struct {
	// 用户参与应用的统计指标
	AppCounter *ResourceCounterData `json:"AppCounter,omitnil" name:"AppCounter"`

	// 用户创建资源的统计指标
	UserCounter *ResourceCounterData `json:"UserCounter,omitnil" name:"UserCounter"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetAppSummaryResponse struct {
	*tchttp.BaseResponse
	Response *GetAppSummaryResponseParams `json:"Response"`
}

func (r *GetAppSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAppSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCredentialStateRequestParams struct {
	// 凭证唯一ID
	CredentialId *string `json:"CredentialId,omitnil" name:"CredentialId"`

	// 用户应用id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// DID应用id
	UAPId *uint64 `json:"UAPId,omitnil" name:"UAPId"`
}

type GetCredentialStateRequest struct {
	*tchttp.BaseRequest

	// 凭证唯一ID
	CredentialId *string `json:"CredentialId,omitnil" name:"CredentialId"`

	// 用户应用id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// DID应用id
	UAPId *uint64 `json:"UAPId,omitnil" name:"UAPId"`
}

func (r *GetCredentialStateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCredentialStateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CredentialId")
	delete(f, "DAPId")
	delete(f, "UAPId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetCredentialStateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCredentialStateResponseParams struct {
	// 凭证状态信息
	CredentialState *CredentialState `json:"CredentialState,omitnil" name:"CredentialState"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetCredentialStateResponse struct {
	*tchttp.BaseResponse
	Response *GetCredentialStateResponseParams `json:"Response"`
}

func (r *GetCredentialStateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCredentialStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetOverSummaryRequestParams struct {
}

type GetOverSummaryRequest struct {
	*tchttp.BaseRequest
}

func (r *GetOverSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetOverSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}

	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetOverSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetOverSummaryResponseParams struct {
	// 用户参与应用的统计指标
	AppCounter *ResourceCounterData `json:"AppCounter,omitnil" name:"AppCounter"`

	// 用户部署应用的统计指标
	UserCounter *ResourceCounterData `json:"UserCounter,omitnil" name:"UserCounter"`

	// 用户参与的应用总数
	AppCnt *uint64 `json:"AppCnt,omitnil" name:"AppCnt"`

	// 用户部署的应用总数
	DeployCnt *uint64 `json:"DeployCnt,omitnil" name:"DeployCnt"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetOverSummaryResponse struct {
	*tchttp.BaseResponse
	Response *GetOverSummaryResponseParams `json:"Response"`
}

func (r *GetOverSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetOverSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetProposalBaseRequestParams struct {
	// 提案Id
	ProposalId *uint64 `json:"ProposalId,omitnil" name:"ProposalId"`
}

type GetProposalBaseRequest struct {
	*tchttp.BaseRequest

	// 提案Id
	ProposalId *uint64 `json:"ProposalId,omitnil" name:"ProposalId"`
}

func (r *GetProposalBaseRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetProposalBaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProposalId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetProposalBaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetProposalBaseResponseParams struct {
	// 提案基础信息
	ProposalBase *ProposalBase `json:"ProposalBase,omitnil" name:"ProposalBase"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetProposalBaseResponse struct {
	*tchttp.BaseResponse
	Response *GetProposalBaseResponseParams `json:"Response"`
}

func (r *GetProposalBaseResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetProposalBaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetProposalResultRequestParams struct {
	// 提案Id
	ProposalId *uint64 `json:"ProposalId,omitnil" name:"ProposalId"`
}

type GetProposalResultRequest struct {
	*tchttp.BaseRequest

	// 提案Id
	ProposalId *uint64 `json:"ProposalId,omitnil" name:"ProposalId"`
}

func (r *GetProposalResultRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetProposalResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProposalId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetProposalResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetProposalResultResponseParams struct {
	// 当前支持权重/投票数
	AgreeWeight *uint64 `json:"AgreeWeight,omitnil" name:"AgreeWeight"`

	// 当前反对权重/投票数
	RefuseWeight *uint64 `json:"RefuseWeight,omitnil" name:"RefuseWeight"`

	// 当前弃权权重/投票数
	AbandonWeight *uint64 `json:"AbandonWeight,omitnil" name:"AbandonWeight"`

	// 当前控制台用户投票选择，-1代表暂未投票，0代表反对，1代表支持，2代表弃权
	VoterChoice *int64 `json:"VoterChoice,omitnil" name:"VoterChoice"`

	// 当前控制台用户是否参与提案
	IsParticipate *bool `json:"IsParticipate,omitnil" name:"IsParticipate"`

	// 总权重/投票数
	TotalWeight *uint64 `json:"TotalWeight,omitnil" name:"TotalWeight"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetProposalResultResponse struct {
	*tchttp.BaseResponse
	Response *GetProposalResultResponseParams `json:"Response"`
}

func (r *GetProposalResultResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetProposalResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTDidByObjectIdRequestParams struct {
	// 业务层用户ID
	ObjectId *string `json:"ObjectId,omitnil" name:"ObjectId"`

	// DID应用ID
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`
}

type GetTDidByObjectIdRequest struct {
	*tchttp.BaseRequest

	// 业务层用户ID
	ObjectId *string `json:"ObjectId,omitnil" name:"ObjectId"`

	// DID应用ID
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`
}

func (r *GetTDidByObjectIdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTDidByObjectIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ObjectId")
	delete(f, "DAPId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTDidByObjectIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTDidByObjectIdResponseParams struct {
	// DID标识
	Did *string `json:"Did,omitnil" name:"Did"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetTDidByObjectIdResponse struct {
	*tchttp.BaseResponse
	Response *GetTDidByObjectIdResponseParams `json:"Response"`
}

func (r *GetTDidByObjectIdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTDidByObjectIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTDidDocumentRequestParams struct {
	// DID标识
	Did *string `json:"Did,omitnil" name:"Did"`

	// DID应用ID
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// 用户应用ID
	UAPId *string `json:"UAPId,omitnil" name:"UAPId"`
}

type GetTDidDocumentRequest struct {
	*tchttp.BaseRequest

	// DID标识
	Did *string `json:"Did,omitnil" name:"Did"`

	// DID应用ID
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// 用户应用ID
	UAPId *string `json:"UAPId,omitnil" name:"UAPId"`
}

func (r *GetTDidDocumentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTDidDocumentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Did")
	delete(f, "DAPId")
	delete(f, "UAPId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTDidDocumentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTDidDocumentResponseParams struct {
	// DID文档
	Document *string `json:"Document,omitnil" name:"Document"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetTDidDocumentResponse struct {
	*tchttp.BaseResponse
	Response *GetTDidDocumentResponseParams `json:"Response"`
}

func (r *GetTDidDocumentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTDidDocumentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTDidPubKeyRequestParams struct {
	// DID标识
	Did *string `json:"Did,omitnil" name:"Did"`

	// 用户应用ID
	UAPId *uint64 `json:"UAPId,omitnil" name:"UAPId"`

	// DID应用ID
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`
}

type GetTDidPubKeyRequest struct {
	*tchttp.BaseRequest

	// DID标识
	Did *string `json:"Did,omitnil" name:"Did"`

	// 用户应用ID
	UAPId *uint64 `json:"UAPId,omitnil" name:"UAPId"`

	// DID应用ID
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`
}

func (r *GetTDidPubKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTDidPubKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Did")
	delete(f, "UAPId")
	delete(f, "DAPId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTDidPubKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTDidPubKeyResponseParams struct {
	// DID公钥数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthPublicKeyList []*string `json:"AuthPublicKeyList,omitnil" name:"AuthPublicKeyList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetTDidPubKeyResponse struct {
	*tchttp.BaseResponse
	Response *GetTDidPubKeyResponseParams `json:"Response"`
}

func (r *GetTDidPubKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTDidPubKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTraceListRequestParams struct {
	// 应用ID
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// 资源类型，目前支持：DID、CPT
	ResourceType *string `json:"ResourceType,omitnil" name:"ResourceType"`

	// 资源Id
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 搜索起始时间，必须与EndTime成对出现
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 搜索截止时间，必须与StartTime成对出现
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`
}

type GetTraceListRequest struct {
	*tchttp.BaseRequest

	// 应用ID
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// 资源类型，目前支持：DID、CPT
	ResourceType *string `json:"ResourceType,omitnil" name:"ResourceType"`

	// 资源Id
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 搜索起始时间，必须与EndTime成对出现
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 搜索截止时间，必须与StartTime成对出现
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`
}

func (r *GetTraceListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTraceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DAPId")
	delete(f, "ResourceType")
	delete(f, "ResourceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTraceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTraceListResponseParams struct {
	// DID区块链网络信息
	TraceList []*Trace `json:"TraceList,omitnil" name:"TraceList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetTraceListResponse struct {
	*tchttp.BaseResponse
	Response *GetTraceListResponseParams `json:"Response"`
}

func (r *GetTraceListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTraceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IssueCredentialRequestParams struct {
	// 参数集合，详见示例
	CRDLArg *CRDLArg `json:"CRDLArg,omitnil" name:"CRDLArg"`

	// DID应用id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// 凭证版本
	CredentialVersion *string `json:"CredentialVersion,omitnil" name:"CredentialVersion"`

	// 是否未签名
	UnSigned *bool `json:"UnSigned,omitnil" name:"UnSigned"`

	// 零知识范围证明承诺信息
	CommitmentClaims []*CommitmentClaim `json:"CommitmentClaims,omitnil" name:"CommitmentClaims"`
}

type IssueCredentialRequest struct {
	*tchttp.BaseRequest

	// 参数集合，详见示例
	CRDLArg *CRDLArg `json:"CRDLArg,omitnil" name:"CRDLArg"`

	// DID应用id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// 凭证版本
	CredentialVersion *string `json:"CredentialVersion,omitnil" name:"CredentialVersion"`

	// 是否未签名
	UnSigned *bool `json:"UnSigned,omitnil" name:"UnSigned"`

	// 零知识范围证明承诺信息
	CommitmentClaims []*CommitmentClaim `json:"CommitmentClaims,omitnil" name:"CommitmentClaims"`
}

func (r *IssueCredentialRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IssueCredentialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CRDLArg")
	delete(f, "DAPId")
	delete(f, "CredentialVersion")
	delete(f, "UnSigned")
	delete(f, "CommitmentClaims")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IssueCredentialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IssueCredentialResponseParams struct {
	// 凭证的json内容
	CredentialData *string `json:"CredentialData,omitnil" name:"CredentialData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type IssueCredentialResponse struct {
	*tchttp.BaseResponse
	Response *IssueCredentialResponseParams `json:"Response"`
}

func (r *IssueCredentialResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IssueCredentialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListVoterRecordRequestParams struct {
	// 提案Id
	ProposalId *uint64 `json:"ProposalId,omitnil" name:"ProposalId"`

	// 加载更多的起始Id，第一次填0，后续根据接口返回值填写
	StartId *uint64 `json:"StartId,omitnil" name:"StartId"`

	// 每次加载条数
	Count *uint64 `json:"Count,omitnil" name:"Count"`

	// DID 标识的文本，精确搜索
	SearchKey *string `json:"SearchKey,omitnil" name:"SearchKey"`

	// 是否获取全部提案成员
	IsAllVoter *bool `json:"IsAllVoter,omitnil" name:"IsAllVoter"`
}

type ListVoterRecordRequest struct {
	*tchttp.BaseRequest

	// 提案Id
	ProposalId *uint64 `json:"ProposalId,omitnil" name:"ProposalId"`

	// 加载更多的起始Id，第一次填0，后续根据接口返回值填写
	StartId *uint64 `json:"StartId,omitnil" name:"StartId"`

	// 每次加载条数
	Count *uint64 `json:"Count,omitnil" name:"Count"`

	// DID 标识的文本，精确搜索
	SearchKey *string `json:"SearchKey,omitnil" name:"SearchKey"`

	// 是否获取全部提案成员
	IsAllVoter *bool `json:"IsAllVoter,omitnil" name:"IsAllVoter"`
}

func (r *ListVoterRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListVoterRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProposalId")
	delete(f, "StartId")
	delete(f, "Count")
	delete(f, "SearchKey")
	delete(f, "IsAllVoter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListVoterRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListVoterRecordResponseParams struct {
	// 提案投票记录列表
	VoterRecordList []*VoteRecordItem `json:"VoterRecordList,omitnil" name:"VoterRecordList"`

	// 是否已加载结束
	IsEnded *bool `json:"IsEnded,omitnil" name:"IsEnded"`

	// 下载加载的起始Id
	NextStartId *uint64 `json:"NextStartId,omitnil" name:"NextStartId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ListVoterRecordResponse struct {
	*tchttp.BaseResponse
	Response *ListVoterRecordResponseParams `json:"Response"`
}

func (r *ListVoterRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListVoterRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProofThreshold struct {
	// 致盲因子
	Opening *string `json:"Opening,omitnil" name:"Opening"`

	// 承诺声明字段名称
	Commitment *string `json:"Commitment,omitnil" name:"Commitment"`

	// 范围证明的目标阈值:，证明命题为: Commitment>=Threshold
	Threshold *uint64 `json:"Threshold,omitnil" name:"Threshold"`

	// 原始值声明字段名称，DiffVal存在时，此字段忽略
	Origin *string `json:"Origin,omitnil" name:"Origin"`

	// 原始值与阈值的差值，此字段存在时，origin字段忽略
	DiffVal *uint64 `json:"DiffVal,omitnil" name:"DiffVal"`
}

type ProposalBase struct {
	// 提案Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 提案发布者DID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatorDid *string `json:"CreatorDid,omitnil" name:"CreatorDid"`

	// 提案名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 提案状态，0 代表已否决，1 代表已通过，2 代表待开始，3 代表投票中
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 当前控制台用户是否关注提案
	// 注意：此字段可能返回 null，表示取不到有效值。
	Followed *bool `json:"Followed,omitnil" name:"Followed"`

	// 提案发布时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 投票开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	VotingStartTime *string `json:"VotingStartTime,omitnil" name:"VotingStartTime"`

	// 投票截止时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	VotingDeadline *string `json:"VotingDeadline,omitnil" name:"VotingDeadline"`

	// 提案类型：0代表普通提案，1代表DAO提案
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProposalType *uint64 `json:"ProposalType,omitnil" name:"ProposalType"`

	// 提案摘要
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 提案发布者 DID 归属应用的 UAPId
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatorUAPIds []*uint64 `json:"CreatorUAPIds,omitnil" name:"CreatorUAPIds"`

	// 提案描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Detail *string `json:"Detail,omitnil" name:"Detail"`

	// 提案合约地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContractAddress *string `json:"ContractAddress,omitnil" name:"ContractAddress"`

	// 提案实际结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	VotingEndTime *string `json:"VotingEndTime,omitnil" name:"VotingEndTime"`
}

// Predefined struct for user
type QueryAuthorityInfoRequestParams struct {
	// DID标识
	Did *string `json:"Did,omitnil" name:"Did"`

	// DID应用id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`
}

type QueryAuthorityInfoRequest struct {
	*tchttp.BaseRequest

	// DID标识
	Did *string `json:"Did,omitnil" name:"Did"`

	// DID应用id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`
}

func (r *QueryAuthorityInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryAuthorityInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Did")
	delete(f, "DAPId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryAuthorityInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryAuthorityInfoResponseParams struct {
	// 名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 权威机构did
	Did *string `json:"Did,omitnil" name:"Did"`

	// 状态：1为已认证，2为未认证
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 机构备注信息
	Description *string `json:"Description,omitnil" name:"Description"`

	// 认证时间
	RecognizeTime *string `json:"RecognizeTime,omitnil" name:"RecognizeTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type QueryAuthorityInfoResponse struct {
	*tchttp.BaseResponse
	Response *QueryAuthorityInfoResponseParams `json:"Response"`
}

func (r *QueryAuthorityInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryAuthorityInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryCPTRequestParams struct {
	// 凭证模板id
	CPTId *int64 `json:"CPTId,omitnil" name:"CPTId"`

	// DID应用id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// 用户应用id
	UAPId *uint64 `json:"UAPId,omitnil" name:"UAPId"`
}

type QueryCPTRequest struct {
	*tchttp.BaseRequest

	// 凭证模板id
	CPTId *int64 `json:"CPTId,omitnil" name:"CPTId"`

	// DID应用id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// 用户应用id
	UAPId *uint64 `json:"UAPId,omitnil" name:"UAPId"`
}

func (r *QueryCPTRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCPTRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CPTId")
	delete(f, "DAPId")
	delete(f, "UAPId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryCPTRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryCPTResponseParams struct {
	// 凭证模板内容
	CPTJson *string `json:"CPTJson,omitnil" name:"CPTJson"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type QueryCPTResponse struct {
	*tchttp.BaseResponse
	Response *QueryCPTResponseParams `json:"Response"`
}

func (r *QueryCPTResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCPTResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceCounterData struct {
	// DID总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DidCnt *uint64 `json:"DidCnt,omitnil" name:"DidCnt"`

	// VC总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	VCCnt *uint64 `json:"VCCnt,omitnil" name:"VCCnt"`

	// CPT总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CPTCnt *uint64 `json:"CPTCnt,omitnil" name:"CPTCnt"`

	//  VC验证总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	VerifyCnt *uint64 `json:"VerifyCnt,omitnil" name:"VerifyCnt"`
}

// Predefined struct for user
type RotateTDidAuthPubKeyRequestParams struct {
	// DID标识
	Did *string `json:"Did,omitnil" name:"Did"`

	// 轮换DID公钥的临时凭证内容，通过创建凭证接口(CreateCredential)生成并签名，凭证类型为：OperateCredential, 为安全起见凭证过期时间不适合太长，建议设置为1分钟内
	OperateCredential *string `json:"OperateCredential,omitnil" name:"OperateCredential"`

	// DID应用id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`
}

type RotateTDidAuthPubKeyRequest struct {
	*tchttp.BaseRequest

	// DID标识
	Did *string `json:"Did,omitnil" name:"Did"`

	// 轮换DID公钥的临时凭证内容，通过创建凭证接口(CreateCredential)生成并签名，凭证类型为：OperateCredential, 为安全起见凭证过期时间不适合太长，建议设置为1分钟内
	OperateCredential *string `json:"OperateCredential,omitnil" name:"OperateCredential"`

	// DID应用id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`
}

func (r *RotateTDidAuthPubKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RotateTDidAuthPubKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Did")
	delete(f, "OperateCredential")
	delete(f, "DAPId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RotateTDidAuthPubKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RotateTDidAuthPubKeyResponseParams struct {
	// 链上返回交易信息
	Transaction *ChainTransaction `json:"Transaction,omitnil" name:"Transaction"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RotateTDidAuthPubKeyResponse struct {
	*tchttp.BaseResponse
	Response *RotateTDidAuthPubKeyResponseParams `json:"Response"`
}

func (r *RotateTDidAuthPubKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RotateTDidAuthPubKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetTDidAttributeRequestParams struct {
	// did标识符
	Did *string `json:"Did,omitnil" name:"Did"`

	// 属性名值对数组
	Attributes []*DidAttribute `json:"Attributes,omitnil" name:"Attributes"`

	// DID应用id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`
}

type SetTDidAttributeRequest struct {
	*tchttp.BaseRequest

	// did标识符
	Did *string `json:"Did,omitnil" name:"Did"`

	// 属性名值对数组
	Attributes []*DidAttribute `json:"Attributes,omitnil" name:"Attributes"`

	// DID应用id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`
}

func (r *SetTDidAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetTDidAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Did")
	delete(f, "Attributes")
	delete(f, "DAPId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetTDidAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetTDidAttributeResponseParams struct {
	// 链上返回交易信息
	Transaction *ChainTransaction `json:"Transaction,omitnil" name:"Transaction"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SetTDidAttributeResponse struct {
	*tchttp.BaseResponse
	Response *SetTDidAttributeResponseParams `json:"Response"`
}

func (r *SetTDidAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetTDidAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Trace struct {
	// 区块高度
	// 注意：此字段可能返回 null，表示取不到有效值。
	BlockHeight *uint64 `json:"BlockHeight,omitnil" name:"BlockHeight"`

	// 交易Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TxId *string `json:"TxId,omitnil" name:"TxId"`

	// 调用的合约类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContractType *string `json:"ContractType,omitnil" name:"ContractType"`

	// 调用的合约方法
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContractMethod *string `json:"ContractMethod,omitnil" name:"ContractMethod"`

	// 调用的合约参数，以JSON形式展示
	// 注意：此字段可能返回 null，表示取不到有效值。
	Parameters *string `json:"Parameters,omitnil" name:"Parameters"`

	// 合约执行过程中抛出的事件，以JSON形式展示
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventData *string `json:"EventData,omitnil" name:"EventData"`

	// 交易时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamp *string `json:"Timestamp,omitnil" name:"Timestamp"`

	// 资源类型，目前包含两类资源：DID、CPT
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceType *string `json:"ResourceType,omitnil" name:"ResourceType"`

	// 资源Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`
}

// Predefined struct for user
type UpdateCredentialStateRequestParams struct {
	// DID应用id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// 更新VC状态的临时凭证内容，通过创建凭证接口(CreateCredential)生成并签名，凭证类型为：OperateCredential, 为安全起见凭证过期时间不适合太长，建议设置为1分钟内
	OperateCredential *string `json:"OperateCredential,omitnil" name:"OperateCredential"`
}

type UpdateCredentialStateRequest struct {
	*tchttp.BaseRequest

	// DID应用id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// 更新VC状态的临时凭证内容，通过创建凭证接口(CreateCredential)生成并签名，凭证类型为：OperateCredential, 为安全起见凭证过期时间不适合太长，建议设置为1分钟内
	OperateCredential *string `json:"OperateCredential,omitnil" name:"OperateCredential"`
}

func (r *UpdateCredentialStateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCredentialStateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DAPId")
	delete(f, "OperateCredential")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateCredentialStateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCredentialStateResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UpdateCredentialStateResponse struct {
	*tchttp.BaseResponse
	Response *UpdateCredentialStateResponseParams `json:"Response"`
}

func (r *UpdateCredentialStateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCredentialStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyCredentialsRequestParams struct {
	// DID应用id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// 0:仅验证签名，1:验证签名和链上状态，2, 仅验证链上状态，默认为0
	VerifyType *uint64 `json:"VerifyType,omitnil" name:"VerifyType"`

	// 凭证内容
	CredentialData *string `json:"CredentialData,omitnil" name:"CredentialData"`
}

type VerifyCredentialsRequest struct {
	*tchttp.BaseRequest

	// DID应用id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// 0:仅验证签名，1:验证签名和链上状态，2, 仅验证链上状态，默认为0
	VerifyType *uint64 `json:"VerifyType,omitnil" name:"VerifyType"`

	// 凭证内容
	CredentialData *string `json:"CredentialData,omitnil" name:"CredentialData"`
}

func (r *VerifyCredentialsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyCredentialsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DAPId")
	delete(f, "VerifyType")
	delete(f, "CredentialData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VerifyCredentialsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyCredentialsResponseParams struct {
	// 是否验证成功
	Result *bool `json:"Result,omitnil" name:"Result"`

	// 验证返回码
	VerifyCode *uint64 `json:"VerifyCode,omitnil" name:"VerifyCode"`

	// 验证消息
	VerifyMessage *string `json:"VerifyMessage,omitnil" name:"VerifyMessage"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type VerifyCredentialsResponse struct {
	*tchttp.BaseResponse
	Response *VerifyCredentialsResponseParams `json:"Response"`
}

func (r *VerifyCredentialsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyCredentialsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyPresentationRequestParams struct {
	// VP持有人的did标识
	Did *string `json:"Did,omitnil" name:"Did"`

	// 可验证表达内容
	PresentationData *string `json:"PresentationData,omitnil" name:"PresentationData"`

	// DID应用id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// 随机验证码
	VerifyCode *string `json:"VerifyCode,omitnil" name:"VerifyCode"`
}

type VerifyPresentationRequest struct {
	*tchttp.BaseRequest

	// VP持有人的did标识
	Did *string `json:"Did,omitnil" name:"Did"`

	// 可验证表达内容
	PresentationData *string `json:"PresentationData,omitnil" name:"PresentationData"`

	// DID应用id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// 随机验证码
	VerifyCode *string `json:"VerifyCode,omitnil" name:"VerifyCode"`
}

func (r *VerifyPresentationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyPresentationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Did")
	delete(f, "PresentationData")
	delete(f, "DAPId")
	delete(f, "VerifyCode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VerifyPresentationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyPresentationResponseParams struct {
	// 是否验证成功
	Result *bool `json:"Result,omitnil" name:"Result"`

	// 验证返回码
	VerifyCode *uint64 `json:"VerifyCode,omitnil" name:"VerifyCode"`

	// 验证消息
	VerifyMessage *string `json:"VerifyMessage,omitnil" name:"VerifyMessage"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type VerifyPresentationResponse struct {
	*tchttp.BaseResponse
	Response *VerifyPresentationResponseParams `json:"Response"`
}

func (r *VerifyPresentationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyPresentationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VoteProposalRequestParams struct {
	// 提案Id
	ProposalId *uint64 `json:"ProposalId,omitnil" name:"ProposalId"`

	// 应用Id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// 投票凭证
	VoteCredential *string `json:"VoteCredential,omitnil" name:"VoteCredential"`
}

type VoteProposalRequest struct {
	*tchttp.BaseRequest

	// 提案Id
	ProposalId *uint64 `json:"ProposalId,omitnil" name:"ProposalId"`

	// 应用Id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// 投票凭证
	VoteCredential *string `json:"VoteCredential,omitnil" name:"VoteCredential"`
}

func (r *VoteProposalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VoteProposalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProposalId")
	delete(f, "DAPId")
	delete(f, "VoteCredential")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VoteProposalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VoteProposalResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type VoteProposalResponse struct {
	*tchttp.BaseResponse
	Response *VoteProposalResponseParams `json:"Response"`
}

func (r *VoteProposalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VoteProposalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VoteRecordItem struct {
	// 投票记录Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 投票者DID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Voter *string `json:"Voter,omitnil" name:"Voter"`

	// 投票权重
	// 注意：此字段可能返回 null，表示取不到有效值。
	Weight *uint64 `json:"Weight,omitnil" name:"Weight"`

	// 投票选择
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 投票时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	VoteTime *string `json:"VoteTime,omitnil" name:"VoteTime"`
}
