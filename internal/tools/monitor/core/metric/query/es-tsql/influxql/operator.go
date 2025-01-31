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

package esinfluxql

import (
	"github.com/influxdata/influxql"

	tsql "github.com/ping-cloudnative/moonlight/internal/tools/monitor/core/metric/query/es-tsql"
)

func reverseOperator(op influxql.Token) influxql.Token {
	switch op {
	case influxql.LT:
		return influxql.GT
	case influxql.LTE:
		return influxql.GTE
	case influxql.GT:
		return influxql.LT
	case influxql.GTE:
		return influxql.LTE
	}
	return op
}

func toOperator(op influxql.Token) tsql.Operator {
	return tsql.Operator(int(op) - int(influxql.ADD) + 1)
}
