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

package utils

import (
	"context"

	"github.com/ping-cloudnative/moonlight-utils/pkg/transport"
)

func NewContextWithHeader(ctx context.Context) context.Context {
	header := transport.Header{}
	for k, vals := range transport.ContextHeader(ctx) {
		header.Append(k, vals...)
	}
	return transport.WithHeader(context.Background(), header)
}
