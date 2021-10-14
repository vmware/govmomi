/*
Copyright (c) 2021 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package object

import (
	"context"

	"github.com/vmware/govmomi/eam"
	"github.com/vmware/govmomi/eam/methods"
	"github.com/vmware/govmomi/eam/types"
	vim "github.com/vmware/govmomi/vim25/types"
)

type EsxAgentManager struct {
	EamObject
}

// NewEsxAgentManager returns a wrapper for an EsxAgentManager managed object.
func NewEsxAgentManager(c *eam.Client, ref vim.ManagedObjectReference) EsxAgentManager {
	return EsxAgentManager{
		EamObject: EamObject{
			c: c,
			r: ref,
		},
	}
}

func (m EsxAgentManager) CreateAgency(
	ctx context.Context,
	config types.BaseAgencyConfigInfo,
	initialGoalState string) (Agency, error) {

	var agency Agency
	resp, err := methods.CreateAgency(ctx, m.c, &types.CreateAgency{
		This:             m.r,
		AgencyConfigInfo: config,
		InitialGoalState: initialGoalState,
	})
	if err != nil {
		return agency, err
	}
	agency.c = m.c
	agency.r = resp.Returnval
	return agency, nil
}

func (m EsxAgentManager) Agencies(ctx context.Context) ([]Agency, error) {
	resp, err := methods.QueryAgency(ctx, m.c, &types.QueryAgency{
		This: m.r,
	})
	if err != nil {
		return nil, err
	}
	objs := make([]Agency, len(resp.Returnval))
	for i := range resp.Returnval {
		objs[i].c = m.c
		objs[i].r = resp.Returnval[i]
	}
	return objs, nil
}

func (m EsxAgentManager) ScanForUnknownAgentVm(ctx context.Context) error {
	_, err := methods.ScanForUnknownAgentVm(ctx, m.c, &types.ScanForUnknownAgentVm{
		This: m.r,
	})
	if err != nil {
		return err
	}
	return nil
}
