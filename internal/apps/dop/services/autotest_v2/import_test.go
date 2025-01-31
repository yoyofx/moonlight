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

package autotestv2

import (
	"mime/multipart"
	"net/http"
	"os"
	"reflect"
	"testing"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"

	"github.com/ping-cloudnative/moonlight/apistructs"
	"github.com/ping-cloudnative/moonlight/bundle"
	"github.com/ping-cloudnative/moonlight/internal/core/file/filetypes"
	"github.com/ping-cloudnative/moonlight/pkg/i18n"
	"github.com/ping-cloudnative/moonlight/proto-go/core/file/pb"
)

func TestImportSceneSet(t *testing.T) {
	bdl := bundle.New(bundle.WithI18nLoader(&i18n.LocaleResourceLoader{}))
	m := monkey.PatchInstanceMethod(reflect.TypeOf(bdl), "GetTestSpace",
		func(bdl *bundle.Bundle, id uint64) (*apistructs.AutoTestSpace, error) {
			return &apistructs.AutoTestSpace{ID: 1}, nil
		})
	defer m.Unpatch()

	r := &http.Request{}
	m1 := monkey.PatchInstanceMethod(reflect.TypeOf(r), "FormFile",
		func(r *http.Request, key string) (multipart.File, *multipart.FileHeader, error) {
			return &os.File{}, &multipart.FileHeader{Filename: "autotest-scene-set.xlsx"}, nil
		})
	defer m1.Unpatch()

	m2 := monkey.PatchInstanceMethod(reflect.TypeOf(bdl), "UploadFile",
		func(bdl *bundle.Bundle, req filetypes.FileUploadRequest, clientTimeout ...int64) (*pb.File, error) {
			return &pb.File{UUID: "123"}, nil
		})
	defer m2.Unpatch()

	autotestSvc := New()
	autotestSvc.CreateFileRecord = func(req apistructs.TestFileRecordRequest) (uint64, error) {
		return 1, nil
	}
	autotestSvc.bdl = bdl
	_, err := autotestSvc.ImportSceneSet(apistructs.AutoTestSceneSetImportRequest{
		FileType: apistructs.TestSceneSetFileTypeExcel,
		SpaceID:  1,
	}, r)
	assert.NoError(t, err)
}
