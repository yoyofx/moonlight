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

package outPutTitle

import (
	"context"

	"github.com/ping-cloudnative/moonlight-utils/base/servicehub"
	"github.com/ping-cloudnative/moonlight-utils/providers/component-protocol/cpregister/base"
	"github.com/ping-cloudnative/moonlight-utils/providers/component-protocol/cptype"
	"github.com/ping-cloudnative/moonlight-utils/providers/component-protocol/utils/cputil"
)

type ComponentAction struct {
}

func (ca *ComponentAction) Render(ctx context.Context, c *cptype.Component, scenario cptype.Scenario, event cptype.ComponentEvent, gs *cptype.GlobalStateData) error {
	c.Props = map[string]interface{}{
		"title": cputil.I18n(ctx, "sceneOut"),
		"level": 2,
	}
	return nil
}

func init() {
	base.InitProviderWithCreator("auto-test-scenes", "outPutTitle",
		func() servicehub.Provider { return &ComponentAction{} })
}
