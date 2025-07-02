// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package check

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
)

type compat struct {
	checkFlag
}

func init() {
	cli.Register("vm.check.compat", &compat{}, true)
}

func (cmd *compat) Description() string {
	return `Check if VM can be placed on the given HOST in the given resource POOL.

Examples:
  govc vm.check.compat -vm my-vm -host $host -pool $pool`
}

func (cmd *compat) Run(ctx context.Context, f *flag.FlagSet) error {
	vm, err := cmd.VirtualMachine()
	if err != nil {
		return err
	}

	if vm == nil {
		return flag.ErrHelp
	}

	checker, err := cmd.compatChecker()
	if err != nil {
		return err
	}

	res, err := checker.CheckCompatibility(ctx, vm.Reference(), cmd.Host, cmd.Pool, cmd.testTypes...)
	if err != nil {
		return err
	}

	return cmd.result(ctx, res)
}
