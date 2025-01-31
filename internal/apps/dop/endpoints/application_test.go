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

package endpoints

import (
	"testing"

	"github.com/pkg/errors"

	"github.com/ping-cloudnative/moonlight/apistructs"
)

func Test_checkApplicationCreateParam(t *testing.T) {
	tests := []struct {
		name string
		req  apistructs.ApplicationCreateRequest
		want error
	}{
		{
			name: "invalid_name",
			req:  apistructs.ApplicationCreateRequest{},
			want: errors.Errorf("invalid request, name is empty"),
		},
		{
			name: "invalid_projectID",
			req: apistructs.ApplicationCreateRequest{
				Name: "demo",
			},
			want: errors.Errorf("invalid request, projectId is empty"),
		},
		{
			name: "invalid_mode",
			req: apistructs.ApplicationCreateRequest{
				Name:      "demo",
				ProjectID: uint64(1),
				Mode:      "app",
			},
			want: errors.New("invalid mode"),
		},
	}
	for _, tt := range tests {
		if err := checkApplicationCreateParam(tt.req); err.Error() != tt.want.Error() {
			t.Errorf("checkApplicationCreateParam() = %v, want %v", err, tt.want)
		}
	}
}
