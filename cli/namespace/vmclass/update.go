// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vmclass

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/namespace"
)

type update struct {
	*flags.ClientFlag

	spec namespace.VirtualMachineClassUpdateSpec
}

func init() {
	cli.Register("namespace.vmclass.update", &update{})
}

func (cmd *update) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.Int64Var(&cmd.spec.CpuCount, "cpus", 0, "The number of CPUs.")
	f.Int64Var(&cmd.spec.MemoryMb, "memory", 0, "The amount of memory (in MB).")
}

func (cmd *update) Process(ctx context.Context) error {
	return cmd.ClientFlag.Process(ctx)
}

func (cmd *update) Usage() string {
	return "NAME"
}

func (cmd *update) Description() string {
	return `Modifies an existing virtual machine class.

Examples:
  govc namespace.vmclass.update -cpus=8 -memory=8192 test-class`
}

func (cmd *update) Run(ctx context.Context, f *flag.FlagSet) error {
	rc, err := cmd.RestClient()

	if err != nil {
		return err
	}

	cmd.spec.Id = f.Arg(0)

	nm := namespace.NewManager(rc)

	return nm.UpdateVmClass(ctx, cmd.spec.Id, cmd.spec)
}
