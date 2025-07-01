// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vm

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type remove struct {
	*flags.ClientFlag
	*flags.VirtualMachineFlag
}

func init() {
	cli.Register("gpu.vm.remove", &remove{})
}

func (cmd *remove) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)
}

func (cmd *remove) Description() string {
	return `Remove all vGPUs from VM.

Examples:
  govc gpu.vm.remove -vm $vm`
}

func (cmd *remove) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.VirtualMachineFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *remove) Run(ctx context.Context, f *flag.FlagSet) error {
	vm, err := cmd.VirtualMachineFlag.VirtualMachine()
	if err != nil {
		return err
	}

	if vm == nil {
		return flag.ErrHelp
	}

	// Check power state and get device list
	var o mo.VirtualMachine
	pc := property.DefaultCollector(vm.Client())
	err = pc.RetrieveOne(ctx, vm.Reference(), []string{"runtime.powerState", "config.hardware"}, &o)
	if err != nil {
		return err
	}

	if o.Runtime.PowerState == types.VirtualMachinePowerStatePoweredOn {
		return fmt.Errorf("VM must be powered off to remove vGPU")
	}

	if o.Config == nil {
		return fmt.Errorf("failed to get VM configuration")
	}

	// Find vGPU devices
	var deviceChanges []types.BaseVirtualDeviceConfigSpec
	for _, device := range o.Config.Hardware.Device {
		if pci, ok := device.(*types.VirtualPCIPassthrough); ok {
			if _, ok := pci.Backing.(*types.VirtualPCIPassthroughVmiopBackingInfo); ok {
				spec := &types.VirtualDeviceConfigSpec{
					Operation: types.VirtualDeviceConfigSpecOperationRemove,
					Device:    pci,
				}
				deviceChanges = append(deviceChanges, spec)
			}
		}
	}

	if len(deviceChanges) == 0 {
		return fmt.Errorf("no vGPU devices found")
	}

	vmConfigSpec := types.VirtualMachineConfigSpec{
		DeviceChange: deviceChanges,
	}

	task, err := vm.Reconfigure(ctx, vmConfigSpec)
	if err != nil {
		return err
	}

	return task.Wait(ctx)
}
