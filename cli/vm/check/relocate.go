// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package check

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/vim25/types"
)

type relocate struct {
	checkFlag
}

func init() {
	cli.Register("vm.check.relocate", &relocate{}, true)
}

func (cmd *relocate) Description() string {
	return `Check if VM can be relocated.

Examples:
  govc vm.migrate -spec my-vm | govc vm.check.relocate -vm my-vm`
}

func (cmd *relocate) Run(ctx context.Context, f *flag.FlagSet) error {
	vm, err := cmd.VirtualMachine()
	if err != nil {
		return err
	}

	if vm == nil {
		return flag.ErrHelp
	}

	var spec types.VirtualMachineRelocateSpec
	if err := cmd.Spec(&spec); err != nil {
		return err
	}

	checker, err := cmd.provChecker()
	if err != nil {
		return err
	}

	res, err := checker.CheckRelocate(ctx, vm.Reference(), spec, cmd.testTypes...)
	if err != nil {
		return err
	}

	return cmd.result(ctx, res)
}
