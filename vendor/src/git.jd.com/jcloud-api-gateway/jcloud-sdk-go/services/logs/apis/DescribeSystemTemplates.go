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
    logs "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/logs/models"
)

type DescribeSystemTemplatesRequest struct {

    core.JDCloudRequest

    /* 当前所在页，默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 页面大小，默认为20；取值范围[1, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* serviceCode (Optional) */
    ServiceCode *string `json:"serviceCode"`

    /* enable (Optional) */
    Enable *bool `json:"enable"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeSystemTemplatesRequest(
) *DescribeSystemTemplatesRequest {

	return &DescribeSystemTemplatesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/teamplatesx",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param pageNumber: 当前所在页，默认为1 (Optional)
 * param pageSize: 页面大小，默认为20；取值范围[1, 100] (Optional)
 * param serviceCode: serviceCode (Optional)
 * param enable: enable (Optional)
 */
func NewDescribeSystemTemplatesRequestWithAllParams(
    pageNumber *int,
    pageSize *int,
    serviceCode *string,
    enable *bool,
) *DescribeSystemTemplatesRequest {

    return &DescribeSystemTemplatesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/teamplatesx",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        PageNumber: pageNumber,
        PageSize: pageSize,
        ServiceCode: serviceCode,
        Enable: enable,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeSystemTemplatesRequestWithoutParam() *DescribeSystemTemplatesRequest {

    return &DescribeSystemTemplatesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/teamplatesx",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param pageNumber: 当前所在页，默认为1(Optional) */
func (r *DescribeSystemTemplatesRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 页面大小，默认为20；取值范围[1, 100](Optional) */
func (r *DescribeSystemTemplatesRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param serviceCode: serviceCode(Optional) */
func (r *DescribeSystemTemplatesRequest) SetServiceCode(serviceCode string) {
    r.ServiceCode = &serviceCode
}

/* param enable: enable(Optional) */
func (r *DescribeSystemTemplatesRequest) SetEnable(enable bool) {
    r.Enable = &enable
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeSystemTemplatesRequest) GetRegionId() string {
    return ""
}

type DescribeSystemTemplatesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeSystemTemplatesResult `json:"result"`
}

type DescribeSystemTemplatesResult struct {
    NumberPages int64 `json:"numberPages"`
    NumberRecords int64 `json:"numberRecords"`
    PageNumber int64 `json:"pageNumber"`
    PageSize int64 `json:"pageSize"`
    Users []logs.SysTemplateEnd `json:"users"`
}