// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vmclass

import (
	"context"
	"encoding/json"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/namespace"
)

type update struct {
	*flags.ClientFlag

	spec namespace.VirtualMachineClassUpdateSpec

	configSpec string
}

func init() {
	cli.Register("namespace.vmclass.update", &update{})
}

func (cmd *update) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.Int64Var(&cmd.spec.CpuCount, "cpus", 0, "The number of CPUs.")
	f.Int64Var(&cmd.spec.MemoryMb, "memory", 0, "The amount of memory (in MB).")
	f.StringVar(&cmd.configSpec, "spec", "", "VirtualMachineConfigSpec json")
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
	cmd.spec.Id = f.Arg(0)
	if f.NArg() != 1 {
		return flag.ErrHelp
	}

	if cmd.configSpec != "" {
		cmd.spec.ConfigSpec = json.RawMessage(cmd.configSpec)

		err := mergeConfigSpec(&cmd.spec)
		if err != nil {
			return err
		}
	}

	rc, err := cmd.RestClient()
	if err != nil {
		return err
	}

	nm := namespace.NewManager(rc)

	return nm.UpdateVmClass(ctx, cmd.spec.Id, cmd.spec)
}
