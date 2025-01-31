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

package linters_test

import (
	"testing"

	"github.com/ping-cloudnative/moonlight/pkg/database/sqllint"
)

const foreignKeyLinterConfig = `
- name: ForeignKeyLinter
  switchOn: true
  white:
    patterns:
      - ".*-base$"
      - "^fdp.*"`

func TestNewForeignKeyLinter(t *testing.T) {
	cfg, err := sqllint.LoadConfig([]byte(foreignKeyLinterConfig))
	if err != nil {
		t.Fatal("failed to LoadConfig", err)
	}
	var s = script{
		Name: "stmt",
		Content: `
ALTER TABLE students
ADD CONSTRAINT fk_class_id
FOREIGN KEY (class_id)
REFERENCES classes (id);
`,
	}
	linter := sqllint.New(cfg)
	if err := linter.Input("", s.Name, s.GetContent()); err != nil {
		t.Error(err)
	}
	lints := linter.Errors()[s.Name].Lints
	t.Logf("lints: %v", lints)
	if len(lints) == 0 {
		t.Fatal("failed")
	}
}
