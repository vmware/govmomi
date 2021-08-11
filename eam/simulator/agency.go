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
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"

	"github.com/vmware/govmomi/eam/internal"
	"github.com/vmware/govmomi/eam/methods"
	"github.com/vmware/govmomi/eam/mo"
	"github.com/vmware/govmomi/eam/types"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25/soap"
	vim "github.com/vmware/govmomi/vim25/types"
)

// Agency handles the deployment of a single type of agent virtual
// machine and any associated VIB bundle, on a set of compute resources.
type Agency struct {
	EamObject
	mo.Agency
}

// NewAgency returns a new Agency as if CreateAgency were called on the
// EsxAgentManager object.
func NewAgency(
	ctx *simulator.Context,
	agencyConfig types.AgencyConfigInfo,
	initialGoalState string) (*Agency, vim.BaseMethodFault) {

	if agencyConfig.AgentName == "" {
		agencyConfig.AgentName = agencyConfig.AgencyName
	}

	// Define a new Agency object.
	agency := &Agency{
		EamObject: EamObject{
			Self: vim.ManagedObjectReference{
				Type:  internal.Agency,
				Value: uuid.New().String(),
			},
		},
		Agency: mo.Agency{
			Config: agencyConfig,
			Runtime: types.EamObjectRuntimeInfo{
				GoalState: initialGoalState,
			},
		},
	}

	// Register the agency with the registry in order for the agency to
	// start receiving API calls from clients.
	ctx.Map.Put(agency)

	// Define a random numbrer generator to help select resources for the
	// agent VMs.
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Alias the registry that contains the vim25 objects.
	vimMap := simulator.Map

	// Create the agents.
	for i, agentConfig := range agencyConfig.AgentConfig {

		// vmName follows the defined pattern for naming agent VMs
		vmName := fmt.Sprintf("%s (%d)", agencyConfig.AgentName, i+1)

		// vmPlacement contains MoRefs to the resources required to create and
		// place the VM inside of the inventory.
		vmPlacement, err := getAgentVMPlacementOptions(
			ctx,
			vimMap,
			rng,
			i,
			agencyConfig)
		if err != nil {
			return nil, &vim.MethodFault{
				FaultCause: &vim.LocalizedMethodFault{
					LocalizedMessage: err.Error(),
				},
			}
		}

		if _, fault := NewAgent(
			ctx,
			agency.Self,
			agentConfig,
			vmName,
			vmPlacement); fault != nil {

			return nil, fault
		}
	}

	return agency, nil
}

func (m *Agency) AgencyQueryRuntime(
	ctx *simulator.Context,
	req *types.AgencyQueryRuntime) soap.HasFault {

	// Copy the agency's issues into its runtime object upon return.
	m.Runtime.Issue = make([]types.BaseIssue, len(m.Issue))
	i := 0
	for _, issue := range m.Issue {
		m.Runtime.Issue[i] = issue
		i++
	}

	return &methods.AgencyQueryRuntimeBody{
		Res: &types.AgencyQueryRuntimeResponse{
			Returnval: &m.Runtime,
		},
	}
}

func (m *Agency) DestroyAgency(
	ctx *simulator.Context,
	req *types.DestroyAgency) soap.HasFault {

	// Remove any agents associated with this agency.
	agentObjs := ctx.Map.AllReference(internal.Agent)
	for _, obj := range agentObjs {
		agent := obj.(*Agent)
		if *agent.Runtime.Agency == m.Self {
			ctx.Map.Remove(ctx, agent.Self)
		}
	}

	ctx.Map.Remove(ctx, m.Self)
	return &methods.DestroyAgencyBody{
		Res: &types.DestroyAgencyResponse{},
	}
}

func (m *Agency) Disable(
	ctx *simulator.Context,
	req *types.Disable) soap.HasFault {

	m.Runtime.GoalState = string(types.EamObjectRuntimeInfoGoalStateDisabled)

	return &methods.DisableBody{
		Res: &types.DisableResponse{},
	}
}

func (m *Agency) Enable(
	ctx *simulator.Context,
	req *types.Enable) soap.HasFault {

	m.Runtime.GoalState = string(types.EamObjectRuntimeInfoGoalStateEnabled)

	return &methods.EnableBody{
		Res: &types.EnableResponse{},
	}
}

func (m *Agency) QueryAgent(
	ctx *simulator.Context,
	req *types.QueryAgent) soap.HasFault {

	objs := ctx.Map.AllReference(internal.Agent)
	moRefs := make([]vim.ManagedObjectReference, len(objs))
	i := 0
	for _, ref := range objs {
		moRefs[i] = ref.Reference()
		i++
	}
	return &methods.QueryAgentBody{
		Res: &types.QueryAgentResponse{
			Returnval: moRefs,
		},
	}
}

func (m *Agency) QueryConfig(
	ctx *simulator.Context,
	req *types.QueryConfig) soap.HasFault {

	return &methods.QueryConfigBody{
		Res: &types.QueryConfigResponse{
			Returnval: &m.Config,
		},
	}
}

func (m *Agency) RegisterAgentVm(
	ctx *simulator.Context,
	req *types.RegisterAgentVm) soap.HasFault {

	return &methods.RegisterAgentVmBody{
		Res: &types.RegisterAgentVmResponse{
			Returnval: vim.ManagedObjectReference{},
		},
	}
}

func (m *Agency) Uninstall(
	ctx *simulator.Context,
	req *types.Uninstall) soap.HasFault {

	m.Runtime.GoalState = string(types.EamObjectRuntimeInfoGoalStateUninstalled)

	return &methods.UninstallBody{
		Res: &types.UninstallResponse{},
	}
}

func (m *Agency) UnregisterAgentVm(
	ctx *simulator.Context,
	req *types.UnregisterAgentVm) soap.HasFault {

	return &methods.UnregisterAgentVmBody{
		Res: &types.UnregisterAgentVmResponse{},
	}
}

func (m *Agency) Update(
	ctx *simulator.Context,
	req *types.Update) soap.HasFault {

	m.Config = *req.Config.GetAgencyConfigInfo()

	return &methods.UpdateBody{
		Res: &types.UpdateResponse{},
	}
}
