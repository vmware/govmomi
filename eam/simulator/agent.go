// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"

	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25"
	vimmethods "github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/soap"
	vim "github.com/vmware/govmomi/vim25/types"

	"github.com/vmware/govmomi/eam/internal"
	"github.com/vmware/govmomi/eam/methods"
	"github.com/vmware/govmomi/eam/mo"
	"github.com/vmware/govmomi/eam/types"
)

// Agenct is the vSphere ESX Agent Manager managed object responsible
// fordeploying an Agency on a single host. The Agent maintains the state
// of the current deployment in its runtime information
type Agent struct {
	EamObject
	mo.Agent
}

type AgentVMPlacementOptions struct {
	computeResource vim.ManagedObjectReference
	datacenter      vim.ManagedObjectReference
	datastore       vim.ManagedObjectReference
	folder          vim.ManagedObjectReference
	host            vim.ManagedObjectReference
	network         vim.ManagedObjectReference
	pool            vim.ManagedObjectReference
}

// NewAgent returns a new Agent as if CreateAgency were called on the
// EsxAgentManager object.
func NewAgent(
	ctx *simulator.Context,
	agency vim.ManagedObjectReference,
	config types.AgentConfigInfo,
	vmName string,
	vmPlacement AgentVMPlacementOptions) (*Agent, vim.BaseMethodFault) {
	vimCtx := ctx.For(vim25.Path)
	vimMap := vimCtx.Map

	agent := &Agent{
		EamObject: EamObject{
			Self: vim.ManagedObjectReference{
				Type:  internal.Agent,
				Value: uuid.New().String(),
			},
		},
		Agent: mo.Agent{
			Config: config,
			Runtime: types.AgentRuntimeInfo{
				Agency:               &agency,
				VmName:               vmName,
				Host:                 &vmPlacement.host,
				EsxAgentFolder:       &vmPlacement.folder,
				EsxAgentResourcePool: &vmPlacement.pool,
			},
		},
	}

	// Register the agent with the registry in order for the agent to start
	// receiving API calls from clients.
	ctx.Map.Put(agent)

	createVm := func() (vim.ManagedObjectReference, *vim.LocalizedMethodFault) {
		var vmRef vim.ManagedObjectReference

		// vmExtraConfig is used when creating the VM for this agent.
		vmExtraConfig := []vim.BaseOptionValue{}

		// If config.OvfPackageUrl is non-empty and does not appear to point to
		// a local file or an HTTP URI, then assume it is a container.
		if url := config.OvfPackageUrl; url != "" && !fsOrHTTPRx.MatchString(url) {
			vmExtraConfig = append(
				vmExtraConfig,
				&vim.OptionValue{
					Key:   "RUN.container",
					Value: url,
				})
		}

		// Copy the OVF environment properties into the VM's ExtraConfig property.
		if ovfEnv := config.OvfEnvironment; ovfEnv != nil {
			for _, ovfProp := range ovfEnv.OvfProperty {
				vmExtraConfig = append(
					vmExtraConfig,
					&vim.OptionValue{
						Key:   ovfProp.Key,
						Value: ovfProp.Value,
					})
			}
		}

		datastore := vimMap.Get(vmPlacement.datastore).(*simulator.Datastore)
		vmPathName := fmt.Sprintf("[%[1]s] %[2]s/%[2]s.vmx", datastore.Name, vmName)
		vmConfigSpec := vim.VirtualMachineConfigSpec{
			Name:        vmName,
			ExtraConfig: vmExtraConfig,
			Files: &vim.VirtualMachineFileInfo{
				VmPathName: vmPathName,
			},
		}

		// Create the VM for this agent.
		vmFolder := vimMap.Get(vmPlacement.folder).(*simulator.Folder)
		createVmTaskRef := vmFolder.CreateVMTask(vimCtx, &vim.CreateVM_Task{
			This:   vmFolder.Self,
			Config: vmConfigSpec,
			Pool:   vmPlacement.pool,
			Host:   &vmPlacement.host,
		}).(*vimmethods.CreateVM_TaskBody).Res.Returnval
		createVmTask := vimMap.Get(createVmTaskRef).(*simulator.Task)

		// Wait for the task to complete and see if there is an error.
		createVmTask.Wait()
		if createVmTask.Info.Error != nil {
			return vmRef, createVmTask.Info.Error
		}

		vmRef = createVmTask.Info.Result.(vim.ManagedObjectReference)
		vm := vimMap.Get(vmRef).(*simulator.VirtualMachine)
		log.Printf("created agent vm: MoRef=%v, Name=%s", vm.Self, vm.Name)

		// Link the agent to this VM.
		agent.Runtime.Vm = &vm.Self

		return vm.Self, nil
	}

	vmRef, err := createVm()
	if err != nil {
		return nil, &vim.RuntimeFault{
			MethodFault: vim.MethodFault{
				FaultCause: err,
			},
		}
	}

	// Start watching this VM and updating the agent's information about the VM.
	go func(ctx *simulator.Context, eamReg, vimReg *simulator.Registry) {
		var (
			ticker = time.NewTicker(1 * time.Second)
			vmName string
		)
		for range ticker.C {
			eamReg.WithLock(ctx, agent.Self, func() {
				agentObj := eamReg.Get(agent.Self)
				if agentObj == nil {
					log.Printf("not found: %v", agent.Self)
					// If the agent no longer exists then stop watching it.
					ticker.Stop()
					return
				}

				updateAgent := func(vm *simulator.VirtualMachine) {
					if vmName == "" {
						vmName = vm.Config.Name
					}

					// Update the agent's properties from the VM.
					agent := agentObj.(*Agent)
					agent.Runtime.VmPowerState = vm.Runtime.PowerState
					if guest := vm.Summary.Guest; guest == nil {
						agent.Runtime.VmIp = ""
					} else {
						agent.Runtime.VmIp = guest.IpAddress
					}
				}

				vimReg.WithLock(ctx, vmRef, func() {
					if vmObj := vimReg.Get(vmRef); vmObj != nil {
						updateAgent(vmObj.(*simulator.VirtualMachine))
					} else {
						// If the VM no longer exists then create a new agent VM.
						log.Printf(
							"creating new agent vm: %v, %v, vmName=%s",
							agent.Self, vmRef, vmName)

						newVmRef, err := createVm()
						if err != nil {
							log.Printf(
								"failed to create new agent vm: %v, %v, vmName=%s, err=%v",
								agent.Self, vmRef, vmName, *err)
							ticker.Stop()
							return
						}

						// Make sure the vmRef variable is assigned to the new
						// VM's reference for the next time through this loop.
						vmRef = newVmRef

						// Get a lock for the *new* VM.
						vimReg.WithLock(ctx, vmRef, func() {
							vmObj = vimReg.Get(vmRef)
							if vmObj == nil {
								log.Printf("not found: %v", vmRef)
								ticker.Stop()
								return
							}
							updateAgent(vmObj.(*simulator.VirtualMachine))
						})
					}

				})
			})
		}
	}(vimCtx, ctx.Map, vimMap)

	return agent, nil
}

func (m *Agent) AgentQueryConfig(
	ctx *simulator.Context,
	req *types.AgentQueryConfig) soap.HasFault {

	return &methods.AgentQueryConfigBody{
		Res: &types.AgentQueryConfigResponse{
			Returnval: m.Config,
		},
	}
}

func (m *Agent) AgentQueryRuntime(
	ctx *simulator.Context,
	req *types.AgentQueryRuntime) soap.HasFault {

	return &methods.AgentQueryRuntimeBody{
		Res: &types.AgentQueryRuntimeResponse{
			Returnval: m.Runtime,
		},
	}
}

func (m *Agent) MarkAsAvailable(
	ctx *simulator.Context,
	req *types.MarkAsAvailable) soap.HasFault {

	return &methods.MarkAsAvailableBody{
		Res: &types.MarkAsAvailableResponse{},
	}
}
