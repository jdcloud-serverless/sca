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


type SearchLogContextSpec struct {

    /* 查询anchor,基于该值偏移进行上下文检索  */
    Anchor []interface{} `json:"anchor"`

    /* 搜索方向,默认both,可取值:up,down,both (Optional) */
    Direction string `json:"direction"`

    /* 日志记录ID  */
    Id string `json:"id"`

    /* 查看上下文行数大小，最大支持200  */
    LineSize int64 `json:"lineSize"`

    /* 查询日志时返回的时间戳  */
    Time int64 `json:"time"`
}
