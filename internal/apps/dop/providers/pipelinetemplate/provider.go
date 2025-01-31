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

package pipelinetemplate

import (
	"context"

	"github.com/ping-cloudnative/moonlight-utils/base/logs"
	"github.com/ping-cloudnative/moonlight-utils/base/servicehub"
	"github.com/ping-cloudnative/moonlight-utils/pkg/transport"
	"github.com/ping-cloudnative/moonlight-utils/providers/mysql"
	dbclient "github.com/ping-cloudnative/moonlight/internal/apps/dop/providers/pipelinetemplate/db"
	"github.com/ping-cloudnative/moonlight/pkg/common/apis"
	"github.com/ping-cloudnative/moonlight/pkg/database/dbengine"
	"github.com/ping-cloudnative/moonlight/proto-go/dop/pipelinetemplate/pb"
)

type config struct{}

type provider struct {
	Log      logs.Logger
	Cfg      *config
	Register transport.Register
	MySQL    mysql.Interface

	service *ServiceImpl
}

func (p *provider) Init(ctx servicehub.Context) error {
	p.service = &ServiceImpl{
		log: p.Log,
		db: &dbclient.DBClient{DBEngine: &dbengine.DBEngine{
			DB: p.MySQL.DB(),
		}},
	}
	if p.Register != nil {
		pb.RegisterTemplateServiceImp(p.Register, p.service, apis.Options())
	}
	return nil
}

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	switch {
	case ctx.Service() == "erda.dop.pipelinetemplate.TemplateService" || ctx.Type() == pb.TemplateServiceServerType() || ctx.Type() == pb.TemplateServiceHandlerType():
		return p.service
	}
	return p
}

func (q *provider) Run(ctx context.Context) error {
	return nil
}

func init() {
	servicehub.Register("erda.dop.pipelinetemplate", &servicehub.Spec{
		Services:             pb.ServiceNames(),
		Types:                append(pb.Types()),
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
