package vapp

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/types"
)

const removeCmdUsage = "VM..."
const removeCmdDescription = `Remove vApp configuration from a virtual machine.

WARNING: Removal is permanent. Any vApp configuration will need to be
reconstructed from scratch after this operation.`

type remove struct {
	*flags.VirtualMachineFlag
}

func init() {
	cli.Register("vm.vapp.remove", &remove{})
}

func (cmd *remove) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)
}

func (cmd *remove) Usage() string {
	return removeCmdUsage
}

func (cmd *remove) Description() string {
	return infoCmdDescription
}

func (cmd *remove) Process(ctx context.Context) error {
	return cmd.VirtualMachineFlag.Process(ctx)
}

func (cmd *remove) Run(ctx context.Context, f *flag.FlagSet) error {
	vms, err := cmd.VirtualMachines(f.Args())
	if err != nil {
		return err
	}

	for _, vm := range vms {
		if err := cmd.removeVAppConfig(ctx, vm); err != nil {
			return err
		}
	}
	return nil
}

func (cmd *remove) removeVAppConfig(ctx context.Context, vm *object.VirtualMachine) error {
	hasConfig, err := hasVAppConfig(ctx, vm)
	switch {
	case err != nil:
		return err
	case !hasConfig:
		fmt.Fprintf(cmd, "Virtual machine %s has no vApp configuration.\n", vm.Name())
		return nil
	}

	t := true
	spec := types.VirtualMachineConfigSpec{
		VAppConfigRemoved: &t,
	}

	task, err := vm.Reconfigure(ctx, spec)
	if err != nil {
		return err
	}

	if err := task.Wait(ctx); err != nil {
		return err
	}

	fmt.Fprintf(cmd, "vApp configuration successfully removed for virtual machine %s.\n", vm.Name())
	return nil
}
