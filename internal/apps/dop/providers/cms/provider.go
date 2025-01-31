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

package cms

import (
	"github.com/ping-cloudnative/moonlight-utils/base/logs"
	"github.com/ping-cloudnative/moonlight-utils/base/servicehub"
	"github.com/ping-cloudnative/moonlight-utils/pkg/transport"
	. "github.com/ping-cloudnative/moonlight/bundle"
	"github.com/ping-cloudnative/moonlight/internal/pkg/audit"
	"github.com/ping-cloudnative/moonlight/pkg/common/apis"
	cmspb "github.com/ping-cloudnative/moonlight/proto-go/core/pipeline/cms/pb"
	"github.com/ping-cloudnative/moonlight/proto-go/dop/cms/pb"
)

type config struct {
}

// +provider
type provider struct {
	Cfg            *config
	Log            logs.Logger
	Register       transport.Register
	cicdCmsService *CICDCmsService
	audit          audit.Auditor

	PipelineCms cmspb.CmsServiceServer `autowired:"erda.core.pipeline.cms.CmsService" optional:"true"`
}

func (p *provider) Init(ctx servicehub.Context) error {
	p.cicdCmsService = &CICDCmsService{p, New(WithErdaServer()), nil}
	p.audit = audit.GetAuditor(ctx)
	if p.Register != nil {
		pb.RegisterCICDCmsServiceImp(p.Register, p.cicdCmsService, apis.Options())
	}
	return nil
}

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	switch {
	case ctx.Service() == "erda.dop.cms.CICDCmsService" || ctx.Type() == pb.CICDCmsServiceServerType() || ctx.Type() == pb.CICDCmsServiceHandlerType():
		return p.cicdCmsService
	}
	return p
}

func init() {
	servicehub.Register("erda.dop.cms", &servicehub.Spec{
		Services:             pb.ServiceNames(),
		Types:                pb.Types(),
		OptionalDependencies: []string{"service-register"},
		Description:          "",
		ConfigFunc: func() interface{} {
			return &config{}
		},
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
