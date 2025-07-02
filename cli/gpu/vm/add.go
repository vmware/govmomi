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

type add struct {
	*flags.ClientFlag
	*flags.VirtualMachineFlag
	profile string
}

func init() {
	cli.Register("gpu.vm.add", &add{})
}

func (cmd *add) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)

	f.StringVar(&cmd.profile, "profile", "", "vGPU profile")
}

func (cmd *add) Description() string {
	return `Add vGPU to VM.

Examples:
  govc gpu.vm.add -vm $vm -profile nvidia_a40-1b`
}

func (cmd *add) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.VirtualMachineFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *add) Run(ctx context.Context, f *flag.FlagSet) error {
	vm, err := cmd.VirtualMachineFlag.VirtualMachine()
	if err != nil {
		return err
	}

	if vm == nil {
		return flag.ErrHelp
	}

	if cmd.profile == "" {
		return fmt.Errorf("profile argument must be specified")
	}

	// Check power state
	var o mo.VirtualMachine
	pc := property.DefaultCollector(vm.Client())
	err = pc.RetrieveOne(ctx, vm.Reference(), []string{"runtime.powerState"}, &o)
	if err != nil {
		return err
	}

	if o.Runtime.PowerState == types.VirtualMachinePowerStatePoweredOn {
		return fmt.Errorf("VM must be powered off to add vGPU")
	}

	device := &types.VirtualPCIPassthrough{
		VirtualDevice: types.VirtualDevice{
			Key: -100,
			Backing: &types.VirtualPCIPassthroughVmiopBackingInfo{
				Vgpu: cmd.profile,
			},
		},
	}

	vmConfigSpec := types.VirtualMachineConfigSpec{
		DeviceChange: []types.BaseVirtualDeviceConfigSpec{
			&types.VirtualDeviceConfigSpec{
				Operation: types.VirtualDeviceConfigSpecOperationAdd,
				Device:    device,
			},
		},
	}

	task, err := vm.Reconfigure(ctx, vmConfigSpec)
	if err != nil {
		return err
	}

	return task.Wait(ctx)
}
