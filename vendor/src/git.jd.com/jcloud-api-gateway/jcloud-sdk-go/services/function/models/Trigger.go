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


type Trigger struct {

    /* 触发器Id (Optional) */
    TriggerId string `json:"triggerId"`

    /* 触发器所属的函数名称 (Optional) */
    FunctionName string `json:"functionName"`

    /* 触发器所属的函数版本名称 (Optional) */
    VersionName string `json:"versionName"`

    /* 触发器对应的事件源类型，目前有oss和apigateway (Optional) */
    EventSource string `json:"eventSource"`

    /* 触发器对应的事件源Id (Optional) */
    EventSourceId string `json:"eventSourceId"`

    /* jqs队列名称 (Optional) */
    JqsName string `json:"jqsName"`

    /* jqs批处理大小 (Optional) */
    JqsBatchSize int `json:"jqsBatchSize"`

    /* 触发器创建时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 触发器最后修改时间 (Optional) */
    UpdateTime string `json:"updateTime"`

    /* function资源信息 (Optional) */
    FunctionSource FunctionSource `json:"functionSource"`
}
