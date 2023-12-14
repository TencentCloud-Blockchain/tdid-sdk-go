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
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-05-19"

type Client struct {
    common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    cpf := profile.NewClientProfile()
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
    return
}

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewCreateCPTRequest() (request *CreateCPTRequest) {
    request = &CreateCPTRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "CreateCPT")
    
    
    return
}

func NewCreateCPTResponse() (response *CreateCPTResponse) {
    response = &CreateCPTResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCPT
// 创建凭证模板
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
func (c *Client) CreateCPT(request *CreateCPTRequest) (response *CreateCPTResponse, err error) {
    return c.CreateCPTWithContext(context.Background(), request)
}

// CreateCPT
// 创建凭证模板
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
func (c *Client) CreateCPTWithContext(ctx context.Context, request *CreateCPTRequest) (response *CreateCPTResponse, err error) {
    if request == nil {
        request = NewCreateCPTRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCPT require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCPTResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDisclosedCredentialRequest() (request *CreateDisclosedCredentialRequest) {
    request = &CreateDisclosedCredentialRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "CreateDisclosedCredential")
    
    
    return
}

func NewCreateDisclosedCredentialResponse() (response *CreateDisclosedCredentialResponse) {
    response = &CreateDisclosedCredentialResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDisclosedCredential
// 根据披露策略创建选择性披露凭证
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDisclosedCredential(request *CreateDisclosedCredentialRequest) (response *CreateDisclosedCredentialResponse, err error) {
    return c.CreateDisclosedCredentialWithContext(context.Background(), request)
}

// CreateDisclosedCredential
// 根据披露策略创建选择性披露凭证
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDisclosedCredentialWithContext(ctx context.Context, request *CreateDisclosedCredentialRequest) (response *CreateDisclosedCredentialResponse, err error) {
    if request == nil {
        request = NewCreateDisclosedCredentialRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDisclosedCredential require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDisclosedCredentialResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePresentationRequest() (request *CreatePresentationRequest) {
    request = &CreatePresentationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "CreatePresentation")
    
    
    return
}

func NewCreatePresentationResponse() (response *CreatePresentationResponse) {
    response = &CreatePresentationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePresentation
// 创建凭证持有人的可验证表达
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreatePresentation(request *CreatePresentationRequest) (response *CreatePresentationResponse, err error) {
    return c.CreatePresentationWithContext(context.Background(), request)
}

// CreatePresentation
// 创建凭证持有人的可验证表达
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreatePresentationWithContext(ctx context.Context, request *CreatePresentationRequest) (response *CreatePresentationResponse, err error) {
    if request == nil {
        request = NewCreatePresentationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePresentation require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePresentationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTDidByHostRequest() (request *CreateTDidByHostRequest) {
    request = &CreateTDidByHostRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "CreateTDidByHost")
    
    
    return
}

func NewCreateTDidByHostResponse() (response *CreateTDidByHostResponse) {
    response = &CreateTDidByHostResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTDidByHost
// 自动生成公私钥对托管在DID平台，并创建DID标识
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
func (c *Client) CreateTDidByHost(request *CreateTDidByHostRequest) (response *CreateTDidByHostResponse, err error) {
    return c.CreateTDidByHostWithContext(context.Background(), request)
}

// CreateTDidByHost
// 自动生成公私钥对托管在DID平台，并创建DID标识
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
func (c *Client) CreateTDidByHostWithContext(ctx context.Context, request *CreateTDidByHostRequest) (response *CreateTDidByHostResponse, err error) {
    if request == nil {
        request = NewCreateTDidByHostRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTDidByHost require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTDidByHostResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTDidByPubKeyRequest() (request *CreateTDidByPubKeyRequest) {
    request = &CreateTDidByPubKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "CreateTDidByPubKey")
    
    
    return
}

func NewCreateTDidByPubKeyResponse() (response *CreateTDidByPubKeyResponse) {
    response = &CreateTDidByPubKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTDidByPubKey
// 通过导入本地的公钥文件注册DID标识
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateTDidByPubKey(request *CreateTDidByPubKeyRequest) (response *CreateTDidByPubKeyResponse, err error) {
    return c.CreateTDidByPubKeyWithContext(context.Background(), request)
}

// CreateTDidByPubKey
// 通过导入本地的公钥文件注册DID标识
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateTDidByPubKeyWithContext(ctx context.Context, request *CreateTDidByPubKeyRequest) (response *CreateTDidByPubKeyResponse, err error) {
    if request == nil {
        request = NewCreateTDidByPubKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTDidByPubKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTDidByPubKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDeactivateTDidRequest() (request *DeactivateTDidRequest) {
    request = &DeactivateTDidRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "DeactivateTDid")
    
    
    return
}

func NewDeactivateTDidResponse() (response *DeactivateTDidResponse) {
    response = &DeactivateTDidResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeactivateTDid
// 更新DID标识的禁用状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeactivateTDid(request *DeactivateTDidRequest) (response *DeactivateTDidResponse, err error) {
    return c.DeactivateTDidWithContext(context.Background(), request)
}

// DeactivateTDid
// 更新DID标识的禁用状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeactivateTDidWithContext(ctx context.Context, request *DeactivateTDidRequest) (response *DeactivateTDidResponse, err error) {
    if request == nil {
        request = NewDeactivateTDidRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeactivateTDid require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeactivateTDidResponse()
    err = c.Send(request, response)
    return
}

func NewGetAppSummaryRequest() (request *GetAppSummaryRequest) {
    request = &GetAppSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "GetAppSummary")
    
    
    return
}

func NewGetAppSummaryResponse() (response *GetAppSummaryResponse) {
    response = &GetAppSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetAppSummary
// 获取某个应用关键指标统计数据
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetAppSummary(request *GetAppSummaryRequest) (response *GetAppSummaryResponse, err error) {
    return c.GetAppSummaryWithContext(context.Background(), request)
}

// GetAppSummary
// 获取某个应用关键指标统计数据
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetAppSummaryWithContext(ctx context.Context, request *GetAppSummaryRequest) (response *GetAppSummaryResponse, err error) {
    if request == nil {
        request = NewGetAppSummaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetAppSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetAppSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewGetCredentialStateRequest() (request *GetCredentialStateRequest) {
    request = &GetCredentialStateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "GetCredentialState")
    
    
    return
}

func NewGetCredentialStateResponse() (response *GetCredentialStateResponse) {
    response = &GetCredentialStateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetCredentialState
// 获取凭证链上状态信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetCredentialState(request *GetCredentialStateRequest) (response *GetCredentialStateResponse, err error) {
    return c.GetCredentialStateWithContext(context.Background(), request)
}

// GetCredentialState
// 获取凭证链上状态信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetCredentialStateWithContext(ctx context.Context, request *GetCredentialStateRequest) (response *GetCredentialStateResponse, err error) {
    if request == nil {
        request = NewGetCredentialStateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetCredentialState require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetCredentialStateResponse()
    err = c.Send(request, response)
    return
}

func NewGetOverSummaryRequest() (request *GetOverSummaryRequest) {
    request = &GetOverSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "GetOverSummary")
    
    
    return
}

func NewGetOverSummaryResponse() (response *GetOverSummaryResponse) {
    response = &GetOverSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetOverSummary
// 获取某个应用关键指标统计数据
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetOverSummary(request *GetOverSummaryRequest) (response *GetOverSummaryResponse, err error) {
    return c.GetOverSummaryWithContext(context.Background(), request)
}

// GetOverSummary
// 获取某个应用关键指标统计数据
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetOverSummaryWithContext(ctx context.Context, request *GetOverSummaryRequest) (response *GetOverSummaryResponse, err error) {
    if request == nil {
        request = NewGetOverSummaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetOverSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetOverSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewGetProposalBaseRequest() (request *GetProposalBaseRequest) {
    request = &GetProposalBaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "GetProposalBase")
    
    
    return
}

func NewGetProposalBaseResponse() (response *GetProposalBaseResponse) {
    response = &GetProposalBaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetProposalBase
// 获取提案基础信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetProposalBase(request *GetProposalBaseRequest) (response *GetProposalBaseResponse, err error) {
    return c.GetProposalBaseWithContext(context.Background(), request)
}

// GetProposalBase
// 获取提案基础信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetProposalBaseWithContext(ctx context.Context, request *GetProposalBaseRequest) (response *GetProposalBaseResponse, err error) {
    if request == nil {
        request = NewGetProposalBaseRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetProposalBase require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetProposalBaseResponse()
    err = c.Send(request, response)
    return
}

func NewGetProposalResultRequest() (request *GetProposalResultRequest) {
    request = &GetProposalResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "GetProposalResult")
    
    
    return
}

func NewGetProposalResultResponse() (response *GetProposalResultResponse) {
    response = &GetProposalResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetProposalResult
// 获取提案投票权重情况
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetProposalResult(request *GetProposalResultRequest) (response *GetProposalResultResponse, err error) {
    return c.GetProposalResultWithContext(context.Background(), request)
}

// GetProposalResult
// 获取提案投票权重情况
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetProposalResultWithContext(ctx context.Context, request *GetProposalResultRequest) (response *GetProposalResultResponse, err error) {
    if request == nil {
        request = NewGetProposalResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetProposalResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetProposalResultResponse()
    err = c.Send(request, response)
    return
}

func NewGetTDidByObjectIdRequest() (request *GetTDidByObjectIdRequest) {
    request = &GetTDidByObjectIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "GetTDidByObjectId")
    
    
    return
}

func NewGetTDidByObjectIdResponse() (response *GetTDidByObjectIdResponse) {
    response = &GetTDidByObjectIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetTDidByObjectId
// 通过业务层对象ID获取DID标识
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetTDidByObjectId(request *GetTDidByObjectIdRequest) (response *GetTDidByObjectIdResponse, err error) {
    return c.GetTDidByObjectIdWithContext(context.Background(), request)
}

// GetTDidByObjectId
// 通过业务层对象ID获取DID标识
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetTDidByObjectIdWithContext(ctx context.Context, request *GetTDidByObjectIdRequest) (response *GetTDidByObjectIdResponse, err error) {
    if request == nil {
        request = NewGetTDidByObjectIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTDidByObjectId require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTDidByObjectIdResponse()
    err = c.Send(request, response)
    return
}

func NewGetTDidDocumentRequest() (request *GetTDidDocumentRequest) {
    request = &GetTDidDocumentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "GetTDidDocument")
    
    
    return
}

func NewGetTDidDocumentResponse() (response *GetTDidDocumentResponse) {
    response = &GetTDidDocumentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetTDidDocument
// 查看DID标识的文档
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetTDidDocument(request *GetTDidDocumentRequest) (response *GetTDidDocumentResponse, err error) {
    return c.GetTDidDocumentWithContext(context.Background(), request)
}

// GetTDidDocument
// 查看DID标识的文档
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetTDidDocumentWithContext(ctx context.Context, request *GetTDidDocumentRequest) (response *GetTDidDocumentResponse, err error) {
    if request == nil {
        request = NewGetTDidDocumentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTDidDocument require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTDidDocumentResponse()
    err = c.Send(request, response)
    return
}

func NewGetTDidPubKeyRequest() (request *GetTDidPubKeyRequest) {
    request = &GetTDidPubKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "GetTDidPubKey")
    
    
    return
}

func NewGetTDidPubKeyResponse() (response *GetTDidPubKeyResponse) {
    response = &GetTDidPubKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetTDidPubKey
// 查询DID标识的认证公钥
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DIDFAILEDOPERATION_DIDNOTEXISTED = "DidFailedOperation.DidNotExisted"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
func (c *Client) GetTDidPubKey(request *GetTDidPubKeyRequest) (response *GetTDidPubKeyResponse, err error) {
    return c.GetTDidPubKeyWithContext(context.Background(), request)
}

// GetTDidPubKey
// 查询DID标识的认证公钥
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DIDFAILEDOPERATION_DIDNOTEXISTED = "DidFailedOperation.DidNotExisted"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
func (c *Client) GetTDidPubKeyWithContext(ctx context.Context, request *GetTDidPubKeyRequest) (response *GetTDidPubKeyResponse, err error) {
    if request == nil {
        request = NewGetTDidPubKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTDidPubKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTDidPubKeyResponse()
    err = c.Send(request, response)
    return
}

func NewGetTraceListRequest() (request *GetTraceListRequest) {
    request = &GetTraceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "GetTraceList")
    
    
    return
}

func NewGetTraceListResponse() (response *GetTraceListResponse) {
    response = &GetTraceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetTraceList
// 获取TDID资源回溯记录
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
func (c *Client) GetTraceList(request *GetTraceListRequest) (response *GetTraceListResponse, err error) {
    return c.GetTraceListWithContext(context.Background(), request)
}

// GetTraceList
// 获取TDID资源回溯记录
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
func (c *Client) GetTraceListWithContext(ctx context.Context, request *GetTraceListRequest) (response *GetTraceListResponse, err error) {
    if request == nil {
        request = NewGetTraceListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTraceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTraceListResponse()
    err = c.Send(request, response)
    return
}

func NewIssueCredentialRequest() (request *IssueCredentialRequest) {
    request = &IssueCredentialRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "IssueCredential")
    
    
    return
}

func NewIssueCredentialResponse() (response *IssueCredentialResponse) {
    response = &IssueCredentialResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// IssueCredential
// 颁发可验证凭证
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) IssueCredential(request *IssueCredentialRequest) (response *IssueCredentialResponse, err error) {
    return c.IssueCredentialWithContext(context.Background(), request)
}

// IssueCredential
// 颁发可验证凭证
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) IssueCredentialWithContext(ctx context.Context, request *IssueCredentialRequest) (response *IssueCredentialResponse, err error) {
    if request == nil {
        request = NewIssueCredentialRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("IssueCredential require credential")
    }

    request.SetContext(ctx)
    
    response = NewIssueCredentialResponse()
    err = c.Send(request, response)
    return
}

func NewListVoterRecordRequest() (request *ListVoterRecordRequest) {
    request = &ListVoterRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "ListVoterRecord")
    
    
    return
}

func NewListVoterRecordResponse() (response *ListVoterRecordResponse) {
    response = &ListVoterRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListVoterRecord
// 获取指定提案投票记录
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ListVoterRecord(request *ListVoterRecordRequest) (response *ListVoterRecordResponse, err error) {
    return c.ListVoterRecordWithContext(context.Background(), request)
}

// ListVoterRecord
// 获取指定提案投票记录
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ListVoterRecordWithContext(ctx context.Context, request *ListVoterRecordRequest) (response *ListVoterRecordResponse, err error) {
    if request == nil {
        request = NewListVoterRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListVoterRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewListVoterRecordResponse()
    err = c.Send(request, response)
    return
}

func NewQueryAuthorityInfoRequest() (request *QueryAuthorityInfoRequest) {
    request = &QueryAuthorityInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "QueryAuthorityInfo")
    
    
    return
}

func NewQueryAuthorityInfoResponse() (response *QueryAuthorityInfoResponse) {
    response = &QueryAuthorityInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryAuthorityInfo
// 查询权威机构信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) QueryAuthorityInfo(request *QueryAuthorityInfoRequest) (response *QueryAuthorityInfoResponse, err error) {
    return c.QueryAuthorityInfoWithContext(context.Background(), request)
}

// QueryAuthorityInfo
// 查询权威机构信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) QueryAuthorityInfoWithContext(ctx context.Context, request *QueryAuthorityInfoRequest) (response *QueryAuthorityInfoResponse, err error) {
    if request == nil {
        request = NewQueryAuthorityInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryAuthorityInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryAuthorityInfoResponse()
    err = c.Send(request, response)
    return
}

func NewQueryCPTRequest() (request *QueryCPTRequest) {
    request = &QueryCPTRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "QueryCPT")
    
    
    return
}

func NewQueryCPTResponse() (response *QueryCPTResponse) {
    response = &QueryCPTResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryCPT
// 查询凭证模版内容
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) QueryCPT(request *QueryCPTRequest) (response *QueryCPTResponse, err error) {
    return c.QueryCPTWithContext(context.Background(), request)
}

// QueryCPT
// 查询凭证模版内容
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) QueryCPTWithContext(ctx context.Context, request *QueryCPTRequest) (response *QueryCPTResponse, err error) {
    if request == nil {
        request = NewQueryCPTRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryCPT require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryCPTResponse()
    err = c.Send(request, response)
    return
}

func NewRotateTDidAuthPubKeyRequest() (request *RotateTDidAuthPubKeyRequest) {
    request = &RotateTDidAuthPubKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "RotateTDidAuthPubKey")
    
    
    return
}

func NewRotateTDidAuthPubKeyResponse() (response *RotateTDidAuthPubKeyResponse) {
    response = &RotateTDidAuthPubKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RotateTDidAuthPubKey
// 轮换DID标识的身份认证公钥
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RotateTDidAuthPubKey(request *RotateTDidAuthPubKeyRequest) (response *RotateTDidAuthPubKeyResponse, err error) {
    return c.RotateTDidAuthPubKeyWithContext(context.Background(), request)
}

// RotateTDidAuthPubKey
// 轮换DID标识的身份认证公钥
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RotateTDidAuthPubKeyWithContext(ctx context.Context, request *RotateTDidAuthPubKeyRequest) (response *RotateTDidAuthPubKeyResponse, err error) {
    if request == nil {
        request = NewRotateTDidAuthPubKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RotateTDidAuthPubKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewRotateTDidAuthPubKeyResponse()
    err = c.Send(request, response)
    return
}

func NewSetTDidAttributeRequest() (request *SetTDidAttributeRequest) {
    request = &SetTDidAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "SetTDidAttribute")
    
    
    return
}

func NewSetTDidAttributeResponse() (response *SetTDidAttributeResponse) {
    response = &SetTDidAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetTDidAttribute
// 设置Did文档的自定义属性
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SetTDidAttribute(request *SetTDidAttributeRequest) (response *SetTDidAttributeResponse, err error) {
    return c.SetTDidAttributeWithContext(context.Background(), request)
}

// SetTDidAttribute
// 设置Did文档的自定义属性
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SetTDidAttributeWithContext(ctx context.Context, request *SetTDidAttributeRequest) (response *SetTDidAttributeResponse, err error) {
    if request == nil {
        request = NewSetTDidAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetTDidAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetTDidAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateCredentialStateRequest() (request *UpdateCredentialStateRequest) {
    request = &UpdateCredentialStateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "UpdateCredentialState")
    
    
    return
}

func NewUpdateCredentialStateResponse() (response *UpdateCredentialStateResponse) {
    response = &UpdateCredentialStateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateCredentialState
// 更新凭证的链上状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateCredentialState(request *UpdateCredentialStateRequest) (response *UpdateCredentialStateResponse, err error) {
    return c.UpdateCredentialStateWithContext(context.Background(), request)
}

// UpdateCredentialState
// 更新凭证的链上状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateCredentialStateWithContext(ctx context.Context, request *UpdateCredentialStateRequest) (response *UpdateCredentialStateResponse, err error) {
    if request == nil {
        request = NewUpdateCredentialStateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateCredentialState require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateCredentialStateResponse()
    err = c.Send(request, response)
    return
}

func NewVerifyCredentialsRequest() (request *VerifyCredentialsRequest) {
    request = &VerifyCredentialsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "VerifyCredentials")
    
    
    return
}

func NewVerifyCredentialsResponse() (response *VerifyCredentialsResponse) {
    response = &VerifyCredentialsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// VerifyCredentials
// 验证已签名的凭证内容
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) VerifyCredentials(request *VerifyCredentialsRequest) (response *VerifyCredentialsResponse, err error) {
    return c.VerifyCredentialsWithContext(context.Background(), request)
}

// VerifyCredentials
// 验证已签名的凭证内容
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) VerifyCredentialsWithContext(ctx context.Context, request *VerifyCredentialsRequest) (response *VerifyCredentialsResponse, err error) {
    if request == nil {
        request = NewVerifyCredentialsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("VerifyCredentials require credential")
    }

    request.SetContext(ctx)
    
    response = NewVerifyCredentialsResponse()
    err = c.Send(request, response)
    return
}

func NewVerifyPresentationRequest() (request *VerifyPresentationRequest) {
    request = &VerifyPresentationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "VerifyPresentation")
    
    
    return
}

func NewVerifyPresentationResponse() (response *VerifyPresentationResponse) {
    response = &VerifyPresentationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// VerifyPresentation
// 验证可验证表达的内容
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) VerifyPresentation(request *VerifyPresentationRequest) (response *VerifyPresentationResponse, err error) {
    return c.VerifyPresentationWithContext(context.Background(), request)
}

// VerifyPresentation
// 验证可验证表达的内容
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) VerifyPresentationWithContext(ctx context.Context, request *VerifyPresentationRequest) (response *VerifyPresentationResponse, err error) {
    if request == nil {
        request = NewVerifyPresentationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("VerifyPresentation require credential")
    }

    request.SetContext(ctx)
    
    response = NewVerifyPresentationResponse()
    err = c.Send(request, response)
    return
}

func NewVoteProposalRequest() (request *VoteProposalRequest) {
    request = &VoteProposalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "VoteProposal")
    
    
    return
}

func NewVoteProposalResponse() (response *VoteProposalResponse) {
    response = &VoteProposalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// VoteProposal
// 通过凭证投票提案
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) VoteProposal(request *VoteProposalRequest) (response *VoteProposalResponse, err error) {
    return c.VoteProposalWithContext(context.Background(), request)
}

// VoteProposal
// 通过凭证投票提案
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) VoteProposalWithContext(ctx context.Context, request *VoteProposalRequest) (response *VoteProposalResponse, err error) {
    if request == nil {
        request = NewVoteProposalRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("VoteProposal require credential")
    }

    request.SetContext(ctx)
    
    response = NewVoteProposalResponse()
    err = c.Send(request, response)
    return
}
