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

package orykratos

import (
	"context"
	"fmt"
	"math"
	"net"
	"net/http"
	"strconv"

	"github.com/ping-cloudnative/moonlight/apistructs"
	openapiauth "github.com/ping-cloudnative/moonlight/internal/core/openapi/openapi-ng/auth"
	"github.com/ping-cloudnative/moonlight/pkg/common/apis"
	"github.com/ping-cloudnative/moonlight/pkg/discover"
	orgpb "github.com/ping-cloudnative/moonlight/proto-go/core/org/pb"
)

type loginChecker struct {
	p *provider
}

func (a *loginChecker) Weight() int64 { return a.p.Cfg.Weight }

func (a *loginChecker) Match(r *http.Request, opts openapiauth.Options) (bool, interface{}) {
	check, _ := opts.Get("CheckLogin").(bool)
	if check {
		session := a.p.getSession(r)
		if len(session) > 0 {
			return true, session
		}
	}
	return false, nil
}

func (a *loginChecker) Check(r *http.Request, data interface{}, opts openapiauth.Options) (bool, *http.Request, error) {
	sessionID := data.(string)
	user, err := a.p.getUserInfo(sessionID)
	if err != nil {
		return false, r, err
	}
	userID := string(user.ID)
	orgID, err := a.p.getScope(r, userID)
	if err != nil {
		return false, r, err
	}
	setUserInfoHeaders(r, userID, orgID)
	return true, r, nil
}

type tryLoginChecker struct {
	p *provider
}

func (a *tryLoginChecker) Weight() int64 { return math.MinInt64 }

func (a *tryLoginChecker) Match(r *http.Request, opts openapiauth.Options) (bool, interface{}) {
	check, _ := opts.Get("TryCheckLogin").(bool)
	if check {
		session := a.p.getSession(r)
		return true, session
	}
	return false, nil
}

func (a *tryLoginChecker) Check(r *http.Request, data interface{}, opts openapiauth.Options) (bool, *http.Request, error) {
	sessionID := data.(string)
	user, err := a.p.getUserInfo(sessionID)
	if err != nil {
		return true, r, nil
	}
	userID := string(user.ID)
	orgID, err := a.p.getScope(r, userID)
	if err != nil {
		return true, r, nil
	}
	setUserInfoHeaders(r, userID, orgID)
	return true, r, nil
}

func setUserInfoHeaders(req *http.Request, userID string, orgID uint64) {
	req.Header.Set("User-ID", userID)
	if orgID != 0 {
		req.Header.Set("Org-ID", strconv.FormatUint(orgID, 10))
	}
}

func (p *provider) getScope(r *http.Request, userID string) (uint64, error) {
	orgName := r.Header.Get("ORG")
	var orgID uint64
	if orgName != "" && orgName != "-" {
		org, err := p.GetOrg(orgName)
		if err != nil {
			return 0, err
		}
		orgID = org.ID
	} else {
		domain := r.Host
		if host, _, err := net.SplitHostPort(domain); err == nil {
			domain = host
		}
		orgResp, err := p.Org.GetOrgByDomain(apis.WithUserIDContext(apis.WithInternalClientContext(context.Background(), discover.SvcOpenapi), userID), &orgpb.GetOrgByDomainRequest{
			Domain: domain,
		})
		if err != nil {
			return 0, err
		}
		org := orgResp.Data
		if org == nil {
			return 0, nil
		}
		orgID = org.ID
	}
	role, err := p.bundle.ScopeRoleAccess(userID, &apistructs.ScopeRoleAccessRequest{
		Scope: apistructs.Scope{
			Type: apistructs.OrgScope,
			ID:   strconv.FormatUint(orgID, 10),
		},
	})
	if err != nil {
		return 0, err
	}
	if !role.Access {
		return 0, fmt.Errorf("permission denied for user:%s org:%d", userID, orgID)
	}
	return orgID, nil
}

func (p *provider) GetOrg(IdOrName string) (*orgpb.Org, error) {
	if IdOrName == "" {
		return nil, fmt.Errorf("the IdOrName is empty")
	}
	orgResp, err := p.Org.GetOrg(apis.WithInternalClientContext(context.Background(), discover.SvcOpenapi), &orgpb.GetOrgRequest{
		IdOrName: IdOrName,
	})
	if err != nil {
		return nil, err
	}
	return orgResp.Data, nil
}
