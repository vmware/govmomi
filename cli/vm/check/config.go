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

type config struct {
	checkFlag
}

func init() {
	cli.Register("vm.check.config", &config{}, true)
}

func (cmd *config) Description() string {
	return `Check if VM config spec can be applied.

Examples:
  govc vm.create -spec ... | govc vm.check.config -pool $pool`
}

func (cmd *config) Run(ctx context.Context, f *flag.FlagSet) error {
	var spec types.VirtualMachineConfigSpec

	if err := cmd.Spec(&spec); err != nil {
		return err
	}

	checker, err := cmd.compatChecker()
	if err != nil {
		return err
	}

	res, err := checker.CheckVmConfig(ctx, spec, cmd.Machine, cmd.Host, cmd.Pool, cmd.testTypes...)
	if err != nil {
		return err
	}

	return cmd.result(ctx, res)
}
