// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

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
