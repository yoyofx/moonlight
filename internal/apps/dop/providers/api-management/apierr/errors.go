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

package apierr

import (
	"github.com/ping-cloudnative/moonlight/pkg/http/httpserver/errorresp"
)

var (
	CreateExportRecord = err("CreateExportRecord", "添加导出记录失败")
	ListExportRecords  = err("ListExportRecords", "查询导出记录失败")
	DeleteExportRecord = err("DeleteExportRecord", "删除导出记录失败")
)

func err(template, defaultValue string) *errorresp.APIError {
	return errorresp.New(errorresp.WithTemplateMessage(template, defaultValue))
}
