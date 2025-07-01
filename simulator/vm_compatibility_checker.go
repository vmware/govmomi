// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"fmt"
	"math/rand"
	"slices"

	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

type VmCompatibilityChecker struct {
	mo.VirtualMachineCompatibilityChecker
}

func resolveHostsAndPool(ctx *Context, vm, host, pool *types.ManagedObjectReference) (*ResourcePool, []types.ManagedObjectReference) {
	var vmMo *VirtualMachine
	var poolMo *ResourcePool

	switch {
	case pool != nil:
		poolMo = ctx.Map.Get(*pool).(*ResourcePool)
	case vm != nil:
		vmMo = ctx.Map.Get(*vm).(*VirtualMachine)
		poolMo = ctx.Map.Get(*vmMo.ResourcePool).(*ResourcePool)
	case host != nil:
		h := ctx.Map.Get(*host).(*HostSystem)
		parent := hostParent(ctx, &h.HostSystem).ResourcePool
		poolMo = ctx.Map.Get(*parent).(*ResourcePool)
	}

	var hosts []types.ManagedObjectReference

	switch {
	case host != nil:
		hosts = append(hosts, *host)
	case pool != nil:
		hosts = resourcePoolHosts(ctx, poolMo)
	case vm != nil:
		hosts = append(hosts, *vmMo.Runtime.Host)
	}

	return poolMo, hosts
}

func validateHostsAndPool(ctx *Context, pool *ResourcePool, hosts []types.ManagedObjectReference) *types.InvalidArgument {
	allHosts := resourcePoolHosts(ctx, pool)

	for _, host := range hosts {
		if !slices.Contains(allHosts, host) {
			return &types.InvalidArgument{
				InvalidProperty: "spec.pool",
			}
		}
	}

	return nil
}

func (c *VmCompatibilityChecker) checkVmConfigSpec(ctx *Context, check *types.CheckResult, spec types.VirtualMachineConfigSpec, hosts []types.ManagedObjectReference) {
	if check.Host == nil {
		// By default all hosts use the same HostSystem template, so we check against any.
		// But we could choose a host based on the spec, e.g. record + playback of real hosts
		check.Host = &hosts[rand.Intn(len(hosts))]
	}

	host := ctx.Map.Get(*check.Host).(*HostSystem)

	mem := int32(spec.MemoryMB)
	if mem > 0 {
		min := int32(4)
		max := host.Capability.MaxSupportedVmMemory
		if mem > max || mem < min {
			check.Warning = append(check.Warning, types.LocalizedMethodFault{
				Fault: &types.MemorySizeNotSupported{
					MemorySizeMB:    mem,
					MinMemorySizeMB: min,
					MaxMemorySizeMB: max,
				},
				LocalizedMessage: fmt.Sprintf("vm requires %d MB of memory, outside the range of %d to %d", mem, min, max),
			})
		}
	}

	cpu := spec.NumCPUs
	if cpu > 0 {
		max := int32(host.Summary.Hardware.NumCpuCores)
		if cpu > max {
			check.Warning = append(check.Warning, types.LocalizedMethodFault{
				Fault: &types.NotEnoughCpus{
					NumCpuDest: max,
					NumCpuVm:   cpu,
				},
				LocalizedMessage: fmt.Sprintf("vm requires %d CPUs, host has %d", cpu, max),
			})
		}
	}

	if spec.GuestId != "" {
		var guest types.VirtualMachineGuestOsIdentifier
		if !slices.Contains(guest.Strings(), spec.GuestId) {
			check.Warning = append(check.Warning, types.LocalizedMethodFault{
				Fault: &types.UnsupportedGuest{
					UnsupportedGuestOS: spec.GuestId,
				},
				LocalizedMessage: fmt.Sprintf("vm guest os %s not supported", spec.GuestId),
			})
		}
	}
}

func (c *VmCompatibilityChecker) CheckVmConfigTask(
	ctx *Context,
	r *types.CheckVmConfig_Task) soap.HasFault {

	task := CreateTask(c, "checkVmConfig", func(t *Task) (types.AnyType, types.BaseMethodFault) {
		if r.Vm == nil && r.Host == nil && r.Pool == nil {
			return nil, new(types.InvalidArgument)
		}

		poolMo, hosts := resolveHostsAndPool(ctx, r.Vm, r.Host, r.Pool)
		if err := validateHostsAndPool(ctx, poolMo, hosts); err != nil {
			return nil, err
		}

		check := types.CheckResult{
			Vm:   r.Vm,
			Host: r.Host,
		}

		c.checkVmConfigSpec(ctx, &check, r.Spec, hosts)

		return types.ArrayOfCheckResult{
			CheckResult: []types.CheckResult{check},
		}, nil
	})

	return &methods.CheckVmConfig_TaskBody{
		Res: &types.CheckVmConfig_TaskResponse{
			Returnval: task.Run(ctx),
		},
	}
}

func (c *VmCompatibilityChecker) CheckCompatibilityTask(
	ctx *Context,
	r *types.CheckCompatibility_Task) soap.HasFault {

	task := CreateTask(c, "checkCompatibility", func(t *Task) (types.AnyType, types.BaseMethodFault) {
		poolMo, hosts := resolveHostsAndPool(ctx, &r.Vm, r.Host, r.Pool)
		if err := validateHostsAndPool(ctx, poolMo, hosts); err != nil {
			return nil, err
		}

		check := types.CheckResult{
			Vm:   &r.Vm,
			Host: r.Host,
		}

		return types.ArrayOfCheckResult{
			CheckResult: []types.CheckResult{check},
		}, nil
	})

	return &methods.CheckCompatibility_TaskBody{
		Res: &types.CheckCompatibility_TaskResponse{
			Returnval: task.Run(ctx),
		},
	}
}
