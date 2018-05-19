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

type propertyAddEdit struct {
	*flags.VirtualMachineFlag

	types.VAppPropertyInfo

	operation types.ArrayUpdateOperation
}

const propertyAddEditUsage = "KEY"
const propertyAddEditescription = `%s a virtual machine property key.

Empty values are ignored. To blank out attributes, delete the key with
vm.vapp.property.delete first, and then re-add the key with only the necessary
values set.

Examples:

vm.vapp.property.add  -vm=foobar -label=Hostname -type=string -userconfigurable=true guestinfo.hostname
vm.vapp.property.edit -vm=foobar -userconfigurable=false -default=foobar.local guestinfo.hostname
vm.vapp.property.edit -vm=foobar -userconfigurable=true -value=foobar.local guestinfo.hostname`

func init() {
	cli.Register("vm.vapp.property.add", &propertyAddEdit{
		operation: types.ArrayUpdateOperationAdd,
	})
	cli.Register("vm.vapp.property.edit", &propertyAddEdit{
		operation: types.ArrayUpdateOperationEdit,
	})
}

func (cmd *propertyAddEdit) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)

	f.StringVar(&cmd.VAppPropertyInfo.Category, "category", "", "Property category")
	f.StringVar(&cmd.VAppPropertyInfo.Label, "label", "", "Property label")
	f.StringVar(&cmd.VAppPropertyInfo.ClassId, "classid", "", "Class ID")
	f.StringVar(&cmd.VAppPropertyInfo.InstanceId, "instanceid", "", "Instance ID")
	f.StringVar(&cmd.VAppPropertyInfo.Description, "description", "", "Property description")
	f.StringVar(&cmd.VAppPropertyInfo.Type, "type", "", "Property type")
	f.StringVar(&cmd.VAppPropertyInfo.TypeReference, "typereference", "", "Additional type reference data")
	f.Var(flags.NewOptionalBool(&cmd.VAppPropertyInfo.UserConfigurable), "userconfigurable", "User configurable flag")
	f.StringVar(&cmd.VAppPropertyInfo.Value, "default", "", "Default value of property")
	f.StringVar(&cmd.VAppPropertyInfo.Value, "value", "", "Current value of property")
}

func (cmd *propertyAddEdit) Usage() string {
	return propertyAddEditUsage
}

func (cmd *propertyAddEdit) Description() string {
	return fmt.Sprintf(propertyAddEditescription, strings.Title(string(cmd.operation)))
}

func (cmd *propertyAddEdit) Run(ctx context.Context, f *flag.FlagSet) error {
	vm, err := cmd.VirtualMachine()
	if err != nil {
		return err
	}
	if vm == nil {
		return errors.New("please specify a virtual machine")
	}

	args := f.Args()
	switch {
	case len(args) > 1:
		return fmt.Errorf("please specify a single property")
	case len(args) < 1:
		return fmt.Errorf("please specify a property")
	}
	propertyID := args[0]

	cfg, err := retrieveVAppConfig(ctx, vm)
	if err != nil {
		return err
	}
	if cfg == nil {
		return fmt.Errorf("%s has no vApp configuration, please set one first with vm.vapp.change", vm.Name())
	}

	op := types.VAppPropertySpec{
		ArrayUpdateSpec: types.ArrayUpdateSpec{
			Operation: cmd.operation,
		},
		Info: &cmd.VAppPropertyInfo,
	}

	op.Info.Id = propertyID

	err = nil
	switch cmd.operation {
	case types.ArrayUpdateOperationAdd:
		op.Info.Key, err = validatePropertyForAdd(cfg.Property, propertyID)
	case types.ArrayUpdateOperationEdit:
		op.Info.Key, err = validatePropertyForEdit(cfg.Property, propertyID)
	}
	if err != nil {
		return err
	}

	spec := types.VirtualMachineConfigSpec{
		VAppConfig: &types.VmConfigSpec{
			Property: []types.VAppPropertySpec{op},
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
		fmt.Sprintf("%sing key %q on virtual machine %s...\n", cmd.operation, propertyID, vm.Name()),
	)
}

func validatePropertyForAdd(ps []types.VAppPropertyInfo, key string) (int32, error) {
	if hasProperty(ps, key) {
		return 0, fmt.Errorf("property already exists: %s", key)
	}
	return int32(len(ps)), nil
}

func validatePropertyForEdit(ps []types.VAppPropertyInfo, key string) (int32, error) {
	p, err := findProperty(ps, key)
	if err != nil {
		return 0, err
	}
	return p.Key, nil
}
