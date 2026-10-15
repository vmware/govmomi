// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vapp

import (
	"context"
	"flag"
	"fmt"
	"strings"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type property struct {
	id    string
	value string
}

type propertyFlag []property

func (p *propertyFlag) String() string {
	return fmt.Sprintf("%v", *p)
}

func (p *propertyFlag) Set(v string) error {
	r := strings.SplitN(v, "=", 2)
	if len(r) != 2 {
		return fmt.Errorf("failed to parse property: %s", v)
	}

	*p = append(*p, property{id: r[0], value: r[1]})

	return nil
}

type set struct {
	*flags.VirtualMachineFlag

	property propertyFlag
}

func init() {
	cli.Register("vm.vapp.set", &set{})
}

func (cmd *set) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)

	f.Var(&cmd.property, "property", "vApp property. <id>=<value>")
}

func (cmd *set) Process(ctx context.Context) error {
	return cmd.VirtualMachineFlag.Process(ctx)
}

func (cmd *set) Description() string {
	return `Set VM vApp property (OVF property) values.

The properties must already exist on the VM, typically defined by the OVF/OVA
template the VM was deployed or cloned from. Use 'vm.vapp.info' to list the
available property ids.

Examples:
  govc vm.vapp.info -vm $vm
  govc vm.vapp.set -vm $vm -property guestinfo.hostname=my-vm
  govc vm.vapp.set -vm $vm -property guestinfo.hostname=my-vm -property guestinfo.ip=10.0.0.5`
}

func (cmd *set) Run(ctx context.Context, f *flag.FlagSet) error {
	vm, err := cmd.VirtualMachine()
	if err != nil {
		return err
	}

	if vm == nil {
		return flag.ErrHelp
	}

	if len(cmd.property) == 0 {
		return flag.ErrHelp
	}

	var mvm mo.VirtualMachine
	if err := vm.Properties(ctx, vm.Reference(), []string{"config.vAppConfig"}, &mvm); err != nil {
		return err
	}

	existing := make(map[string]types.VAppPropertyInfo)
	if mvm.Config != nil && mvm.Config.VAppConfig != nil {
		for _, p := range mvm.Config.VAppConfig.GetVmConfigInfo().Property {
			existing[p.Id] = p
		}
	}

	specs := make([]types.VAppPropertySpec, 0, len(cmd.property))
	for _, p := range cmd.property {
		info, ok := existing[p.id]
		if !ok {
			return fmt.Errorf("vApp property %q does not exist on %s", p.id, vm.InventoryPath)
		}

		// Preserve the existing property's metadata (label, type, etc.) -
		// the API replaces the entire VAppPropertyInfo for "edit" operations.
		info.Value = p.value

		specs = append(specs, types.VAppPropertySpec{
			ArrayUpdateSpec: types.ArrayUpdateSpec{
				Operation: types.ArrayUpdateOperationEdit,
			},
			Info: &info,
		})
	}

	spec := types.VirtualMachineConfigSpec{
		VAppConfig: &types.VmConfigSpec{
			Property: specs,
		},
	}

	task, err := vm.Reconfigure(ctx, spec)
	if err != nil {
		return err
	}

	return task.Wait(ctx)
}
