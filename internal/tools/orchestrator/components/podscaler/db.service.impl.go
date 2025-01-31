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

package podscaler

import (
	"github.com/jinzhu/gorm"

	"github.com/ping-cloudnative/moonlight/internal/tools/orchestrator/dbclient"
	"github.com/ping-cloudnative/moonlight/internal/tools/orchestrator/spec"
	"github.com/ping-cloudnative/moonlight/pkg/database/dbengine"
)

type dbServiceImpl struct {
	db *dbclient.DBClient
}

func (d *dbServiceImpl) CreateHPARule(runtimeHPA *dbclient.RuntimeHPA) error {
	return d.db.CreateRuntimeHPA(runtimeHPA)
}

func (d *dbServiceImpl) UpdateHPARule(runtimeHPA *dbclient.RuntimeHPA) error {
	return d.db.UpdateRuntimeHPA(runtimeHPA)
}

func (d *dbServiceImpl) GetRuntimeHPARulesByServices(id spec.RuntimeUniqueId, services []string) ([]dbclient.RuntimeHPA, error) {
	return d.db.GetRuntimeHPAByServices(id, services)
}

func (d *dbServiceImpl) DeleteRuntimeHPARulesByRuleId(ruleId string) error {
	if err := d.db.DeleteRuntimeHPAByRuleId(ruleId); err != nil {
		return err
	}
	return nil
}

func (d *dbServiceImpl) GetRuntimeHPARuleByRuleId(ruleId string) (dbclient.RuntimeHPA, error) {
	runtimeHPA, err := d.db.GetRuntimeHPARuleByRuleId(ruleId)
	if err != nil {
		return dbclient.RuntimeHPA{}, err
	}
	return *runtimeHPA, nil
}

func (d *dbServiceImpl) GetRuntimeHPARulesByRuntimeId(runtimeID uint64) ([]dbclient.RuntimeHPA, error) {
	runtimeHPAs, err := d.db.GetRuntimeHPARulesByRuntimeId(runtimeID)
	if err != nil {
		return []dbclient.RuntimeHPA{}, err
	}
	return runtimeHPAs, nil
}

func (d *dbServiceImpl) GetRuntime(id uint64) (*dbclient.Runtime, error) {
	return d.db.GetRuntime(id)
}

func (d *dbServiceImpl) GetRuntimeByUniqueID(id spec.RuntimeUniqueId) (*dbclient.Runtime, error) {
	return d.db.FindRuntime(id)
}

func (d *dbServiceImpl) UpdateRuntime(runtime *dbclient.Runtime) error {
	return d.db.UpdateRuntime(runtime)
}

func (d *dbServiceImpl) GetInstanceRouting(id string) (*dbclient.AddonInstanceRouting, error) {
	return d.db.GetInstanceRouting(id)
}

func (d *dbServiceImpl) UpdateAttachment(addonAttachment *dbclient.AddonAttachment) error {
	return d.db.UpdateAttachment(addonAttachment)
}

func (d *dbServiceImpl) UpdatePreDeployment(pre *dbclient.PreDeployment) error {
	return d.db.UpdatePreDeployment(pre)
}

func (d *dbServiceImpl) FindRuntimesByIds(ids []uint64) ([]dbclient.Runtime, error) {
	return d.db.FindRuntimesByIds(ids)
}

func (d *dbServiceImpl) GetPreDeployment(uniqueId spec.RuntimeUniqueId) (*dbclient.PreDeployment, error) {
	return d.db.FindPreDeployment(uniqueId)
}

func (d *dbServiceImpl) GetRuntimeHPAEventsByServices(runtimeId uint64, services []string) ([]dbclient.HPAEventInfo, error) {
	return d.db.GetRuntimeHPAEventsByServices(runtimeId, services)
}

func (d *dbServiceImpl) DeleteRuntimeHPAEventsByRuleId(ruleId string) error {
	if err := d.db.DeleteRuntimeHPAEventsByRuleId(ruleId); err != nil {
		return err
	}
	return nil
}

func (d *dbServiceImpl) GetUnDeletableAttachMentsByRuntimeID(runtimeID uint64) (*[]dbclient.AddonAttachment, error) {
	return d.db.GetUnDeletableAttachMentsByRuntimeID(runtimeID)
}

func (d *dbServiceImpl) CreateVPARule(runtimeVPA *dbclient.RuntimeVPA) error {
	return d.db.CreateRuntimeVPA(runtimeVPA)
}

func (d *dbServiceImpl) UpdateVPARule(runtimeVPA *dbclient.RuntimeVPA) error {
	return d.db.UpdateRuntimeVPA(runtimeVPA)
}

func (d *dbServiceImpl) GetRuntimeVPARulesByServices(id spec.RuntimeUniqueId, services []string) ([]dbclient.RuntimeVPA, error) {
	return d.db.GetRuntimeVPAByServices(id, services)
}

func (d *dbServiceImpl) DeleteRuntimeVPARulesByRuleId(ruleId string) error {
	if err := d.db.DeleteRuntimeVPAByRuleId(ruleId); err != nil {
		return err
	}
	return nil
}

func (d *dbServiceImpl) GetRuntimeVPARuleByRuleId(ruleId string) (dbclient.RuntimeVPA, error) {
	runtimeVPA, err := d.db.GetRuntimeVPARuleByRuleId(ruleId)
	if err != nil {
		return dbclient.RuntimeVPA{}, err
	}
	return *runtimeVPA, nil
}

func (d *dbServiceImpl) GetRuntimeVPARulesByRuntimeId(runtimeID uint64) ([]dbclient.RuntimeVPA, error) {
	runtimeVPAs, err := d.db.GetRuntimeVPARulesByRuntimeId(runtimeID)
	if err != nil {
		return []dbclient.RuntimeVPA{}, err
	}
	return runtimeVPAs, nil
}

func (d *dbServiceImpl) GetRuntimeVPARecommendationsByServices(runtimeId uint64, services []string) ([]dbclient.RuntimeVPAContainerRecommendation, error) {
	return d.db.GetRuntimeVPARecommendationsByServices(runtimeId, services)
}

func newDBService(db *dbclient.DBClient) DBService {
	return &dbServiceImpl{db: db}
}

func NewDBService(orm *gorm.DB) DBService {
	return newDBService(&dbclient.DBClient{
		DBEngine: &dbengine.DBEngine{
			DB: orm,
		},
	})
}
