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

package linkutil

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ping-cloudnative/moonlight/apistructs"
	"github.com/ping-cloudnative/moonlight/internal/pkg/mock"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/spec"
)

func TestGetPipelineLink(t *testing.T) {

	type invalidCase struct {
		p    spec.Pipeline
		desc string
	}

	invalidCases := []invalidCase{
		{
			p: spec.Pipeline{
				Labels: map[string]string{
					apistructs.LabelOrgID: "a",
				},
			},
			desc: "invalid orgID",
		},
		{
			p: spec.Pipeline{
				Labels: map[string]string{
					apistructs.LabelOrgID:     "1",
					apistructs.LabelProjectID: "a",
				},
			},
			desc: "invalid projectID",
		},
		{
			p: spec.Pipeline{
				Labels: map[string]string{
					apistructs.LabelOrgID:     "1",
					apistructs.LabelProjectID: "1",
					apistructs.LabelAppID:     "a",
				},
			},
			desc: "invalid appID",
		},
	}

	for _, c := range invalidCases {
		valid, link := GetPipelineLink(mock.OrgMock{}, c.p)
		assert.False(t, valid)
		assert.Empty(t, link)
	}
}
