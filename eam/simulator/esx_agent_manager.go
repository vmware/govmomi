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

package simulator

import (
	"github.com/vmware/govmomi/eam/internal"
	"github.com/vmware/govmomi/eam/methods"
	"github.com/vmware/govmomi/eam/types"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25/soap"
	vim "github.com/vmware/govmomi/vim25/types"
)

// EsxAgentManager is the main entry point for a solution to create
// agencies in the vSphere ESX Agent Manager server.
type EsxAgentManager struct {
	EamObject
}

func (m *EsxAgentManager) CreateAgency(
	ctx *simulator.Context,
	req *types.CreateAgency) soap.HasFault {

	var res methods.CreateAgencyBody

	if agency, err := NewAgency(
		ctx,
		req.AgencyConfigInfo,
		req.InitialGoalState); err != nil {

		res.Fault_ = simulator.Fault("", err)

	} else {
		res.Res = &types.CreateAgencyResponse{
			Returnval: agency.Self,
		}
	}

	return &res
}

func (m *EsxAgentManager) QueryAgency(
	ctx *simulator.Context,
	req *types.QueryAgency) soap.HasFault {

	objs := ctx.Map.AllReference(internal.Agency)
	moRefs := make([]vim.ManagedObjectReference, len(objs))
	i := 0
	for _, ref := range objs {
		moRefs[i] = ref.Reference()
		i++
	}
	return &methods.QueryAgencyBody{
		Res: &types.QueryAgencyResponse{
			Returnval: moRefs,
		},
	}
}

func (m *EsxAgentManager) ScanForUnknownAgentVm(
	ctx *simulator.Context,
	req *types.ScanForUnknownAgentVm) soap.HasFault {

	return &methods.ScanForUnknownAgentVmBody{
		Res: &types.ScanForUnknownAgentVmResponse{},
	}
}
