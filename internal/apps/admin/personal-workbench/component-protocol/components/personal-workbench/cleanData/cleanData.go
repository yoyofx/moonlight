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

package cleanData

import (
	"context"

	"github.com/ping-cloudnative/moonlight-utils/base/servicehub"
	"github.com/ping-cloudnative/moonlight-utils/providers/component-protocol/cpregister/base"
	"github.com/ping-cloudnative/moonlight-utils/providers/component-protocol/cptype"
	"github.com/ping-cloudnative/moonlight/internal/apps/admin/personal-workbench/component-protocol/components/personal-workbench/common"
)

type Clean struct {
	State map[string]interface{} `json:"state"`
}

func (c2 *Clean) Render(ctx context.Context, c *cptype.Component, scenario cptype.Scenario, event cptype.ComponentEvent, gs *cptype.GlobalStateData) error {
	//*gs = make(cptype.GlobalStateData)
	delete(*gs, common.TabData)
	return nil
}

func init() {
	base.InitProviderWithCreator(common.ScenarioKey, "cleanData", func() servicehub.Provider {
		return &Clean{State: map[string]interface{}{}}
	})
}
