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
    function "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/function/models"
)

type ListTriggersRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 函数名称  */
    FunctionName string `json:"functionName"`

    /* function资源信息 (Optional) */
    FunctionSource *function.FunctionSource `json:"functionSource"`
}

/*
 * param regionId: Region ID (Required)
 * param functionName: 函数名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewListTriggersRequest(
    regionId string,
    functionName string,
) *ListTriggersRequest {

	return &ListTriggersRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/functions/{functionName}:listtrigger",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        FunctionName: functionName,
	}
}

/*
 * param regionId: Region ID (Required)
 * param functionName: 函数名称 (Required)
 * param functionSource: function资源信息 (Optional)
 */
func NewListTriggersRequestWithAllParams(
    regionId string,
    functionName string,
    functionSource *function.FunctionSource,
) *ListTriggersRequest {

    return &ListTriggersRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/functions/{functionName}:listtrigger",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        FunctionName: functionName,
        FunctionSource: functionSource,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewListTriggersRequestWithoutParam() *ListTriggersRequest {

    return &ListTriggersRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/functions/{functionName}:listtrigger",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *ListTriggersRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param functionName: 函数名称(Required) */
func (r *ListTriggersRequest) SetFunctionName(functionName string) {
    r.FunctionName = functionName
}

/* param functionSource: function资源信息(Optional) */
func (r *ListTriggersRequest) SetFunctionSource(functionSource *function.FunctionSource) {
    r.FunctionSource = functionSource
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ListTriggersRequest) GetRegionId() string {
    return r.RegionId
}

type ListTriggersResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ListTriggersResult `json:"result"`
}

type ListTriggersResult struct {
    Data function.ListTriggerData `json:"data"`
}