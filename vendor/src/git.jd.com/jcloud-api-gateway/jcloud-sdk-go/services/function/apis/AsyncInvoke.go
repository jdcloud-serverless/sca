// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package apis

import (
    "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/core"
)

type AsyncInvokeRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 函数名称  */
    FunctionName string `json:"functionName"`

    /* 版本名称  */
    VersionName string `json:"versionName"`

    /* 异步执行函数的输入事件  */
    Event string `json:"event"`
}

/*
 * param regionId: Region ID (Required)
 * param functionName: 函数名称 (Required)
 * param versionName: 版本名称 (Required)
 * param event: 异步执行函数的输入事件 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAsyncInvokeRequest(
    regionId string,
    functionName string,
    versionName string,
    event string,
) *AsyncInvokeRequest {

	return &AsyncInvokeRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/functions/{functionName}/versions/{versionName}:asyncinvoke",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        FunctionName: functionName,
        VersionName: versionName,
        Event: event,
	}
}

/*
 * param regionId: Region ID (Required)
 * param functionName: 函数名称 (Required)
 * param versionName: 版本名称 (Required)
 * param event: 异步执行函数的输入事件 (Required)
 */
func NewAsyncInvokeRequestWithAllParams(
    regionId string,
    functionName string,
    versionName string,
    event string,
) *AsyncInvokeRequest {

    return &AsyncInvokeRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/functions/{functionName}/versions/{versionName}:asyncinvoke",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        FunctionName: functionName,
        VersionName: versionName,
        Event: event,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAsyncInvokeRequestWithoutParam() *AsyncInvokeRequest {

    return &AsyncInvokeRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/functions/{functionName}/versions/{versionName}:asyncinvoke",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *AsyncInvokeRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param functionName: 函数名称(Required) */
func (r *AsyncInvokeRequest) SetFunctionName(functionName string) {
    r.FunctionName = functionName
}

/* param versionName: 版本名称(Required) */
func (r *AsyncInvokeRequest) SetVersionName(versionName string) {
    r.VersionName = versionName
}

/* param event: 异步执行函数的输入事件(Required) */
func (r *AsyncInvokeRequest) SetEvent(event string) {
    r.Event = event
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AsyncInvokeRequest) GetRegionId() string {
    return r.RegionId
}

type AsyncInvokeResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AsyncInvokeResult `json:"result"`
}

type AsyncInvokeResult struct {
}