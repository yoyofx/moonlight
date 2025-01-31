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
	"context"
	"net/http"

	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/services/apierrors"
	"github.com/ping-cloudnative/moonlight/pkg/http/httpserver"
)

func (e *Endpoints) healthCheck(ctx context.Context, r *http.Request, vars map[string]string) (
	httpserver.Responser, error) {
	_, err := e.dbClient.Exec("select 1")
	if err != nil {
		return apierrors.ErrPipelineHealthCheck.InternalError(err).ToResp(), nil
	}
	return httpserver.OkResp("success")
}
