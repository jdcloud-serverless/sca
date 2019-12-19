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

package models


type ShipperTaskEnd struct {

    /* UID (Optional) */
    UID string `json:"uID"`

    /* 是否允许重试， true，false (Optional) */
    AllowRetry bool `json:"allowRetry"`

    /* 创建时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 结束时间 (Optional) */
    EndTime int64 `json:"endTime"`

    /* 日志集uuid (Optional) */
    LogsetUID string `json:"logsetUID"`

    /* 日志主题uuid (Optional) */
    LogtopicUID string `json:"logtopicUID"`

    /* 地域信息 (Optional) */
    Region string `json:"region"`

    /* 日志批次任务截止时间 (Optional) */
    ShipperEndTime int64 `json:"shipperEndTime"`

    /* 日志批次任务起始时间 (Optional) */
    ShipperStartTime int64 `json:"shipperStartTime"`

    /* shipperUID (Optional) */
    ShipperUID string `json:"shipperUID"`

    /* 开始时间 (Optional) */
    StartTime int64 `json:"startTime"`

    /* 该批次转储任务状态： 1 发送中 2 成功，3 失败 (Optional) */
    Status int64 `json:"status"`

    /* 更新时间 (Optional) */
    UpdateTime string `json:"updateTime"`
}
