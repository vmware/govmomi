package vapp

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vim25/types"
)

const cmdDescription = `Remove vApp configuration from a virtual machine.

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

func (cmd *remove) Description() string {
	return cmdDescription
}

func (cmd *remove) Process(ctx context.Context) error {
	return cmd.VirtualMachineFlag.Process(ctx)
}

func (cmd *remove) Run(ctx context.Context, f *flag.FlagSet) error {
	vm, err := cmd.VirtualMachine()
	if err != nil {
		return err
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

	fmt.Printf("vApp properties removed successfully for %s.\n", vm.InventoryPath)
	return nil
}
