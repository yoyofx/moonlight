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

package notificationContentInfo

import (
	"context"

	"github.com/ping-cloudnative/moonlight-utils/providers/component-protocol/cptype"
)

type ComponentNotifyInfo struct {
	sdk *cptype.SDK `json:"-"`
	ctx context.Context

	Type  string            `json:"type,omitempty"`
	Data  map[string]string `json:"data,omitempty"`
	State State             `json:"state,omitempty"`
}

type State struct {
}
