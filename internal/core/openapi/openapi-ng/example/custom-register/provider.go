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

package static

import (
	"net/http"

	"github.com/ping-cloudnative/moonlight-utils/base/servicehub"
	"github.com/ping-cloudnative/moonlight/internal/core/openapi/openapi-ng"
)

type provider struct {
	Router openapi.Interface `autowired:"openapi-router"`
}

func (p *provider) Init(ctx servicehub.Context) (err error) {
	p.Router.Add("GET", "/api/example/custom-api", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("register custom api"))
	})
	return nil
}

func init() {
	servicehub.Register("openapi-example-custom", &servicehub.Spec{
		Creator: func() servicehub.Provider { return &provider{} },
	})
}
