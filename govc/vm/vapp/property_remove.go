package vapp

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"strings"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vim25/types"
)

const propertyRemoveCmdUsage = "KEYS..."
const propertyRemoveCmdDescription = "Remove specified vApp property keys from a virtual machine."

type propertyRemove struct {
	*flags.VirtualMachineFlag
}

func init() {
	cli.Register("vm.vapp.property.remove", &propertyRemove{})
}

func (cmd *propertyRemove) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)
}

func (cmd *propertyRemove) Usage() string {
	return propertyRemoveCmdUsage
}

func (cmd *propertyRemove) Description() string {
	return propertyRemoveCmdDescription
}

func (cmd *propertyRemove) Process(ctx context.Context) error {
	return cmd.VirtualMachineFlag.Process(ctx)
}

func (cmd *propertyRemove) Run(ctx context.Context, f *flag.FlagSet) error {
	vm, err := cmd.VirtualMachine()
	if err != nil {
		return err
	}

	keys := f.Args()
	if len(keys) < 1 {
		return errors.New("no keys were specified")
	}

	cfg, err := retrieveVAppConfig(ctx, vm)
	if err != nil {
		return err
	}
	if cfg == nil {
		return fmt.Errorf("%s has no vApp configuration", vm.Name())
	}

	var ops []types.VAppPropertySpec

	for _, k := range keys {
		pi, err := findProperty(cfg.Property, k)
		if err != nil {
			return err
		}
		op := types.VAppPropertySpec{
			ArrayUpdateSpec: types.ArrayUpdateSpec{
				Operation: types.ArrayUpdateOperationRemove,
				RemoveKey: pi.Key,
			},
		}
		ops = append(ops, op)
	}

	spec := types.VirtualMachineConfigSpec{
		VAppConfig: &types.VmConfigSpec{
			Property: ops,
		},
	}

	task, err := vm.Reconfigure(ctx, spec)
	if err != nil {
		return err
	}

	return waitLog(
		ctx,
		cmd.VirtualMachineFlag.DatacenterFlag.OutputFlag,
		task,
		fmt.Sprintf(fmt.Sprintf("Removing keys on virtual machine %s: [%s]...\n", vm.Name(), strings.Join(keys, ","))),
	)
}
