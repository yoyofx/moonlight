// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmp

import (
	"github.com/ping-cloudnative/moonlight/apistructs"
	"github.com/ping-cloudnative/moonlight/internal/core/openapi/legacy/api/apis"
)

var CMP_CLOUD_RESOURCE_MYSQL_DETAIL_INFO = apis.ApiSpec{
	Path:         "/api/cloud-mysql/<id or name>",
	BackendPath:  "/api/cloud-mysql/<id or name>",
	Host:         "cmp.marathon.l4lb.thisdcos.directory:9027",
	Scheme:       "http",
	Method:       "GET",
	CheckLogin:   true,
	RequestType:  apistructs.CloudResourceMysqlDetailInfoRequest{},
	ResponseType: apistructs.CloudResourceMysqlDetailInfoResponse{},
	Doc:          "获取 mysql instance 详细信息",
}
