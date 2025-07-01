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

type create struct {
	*flags.ClientFlag

	spec namespace.VirtualMachineClassCreateSpec

	configSpec string
}

func init() {
	cli.Register("namespace.vmclass.create", &create{})
}

func (cmd *create) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.Int64Var(&cmd.spec.CpuCount, "cpus", 0, "The number of CPUs.")
	f.Int64Var(&cmd.spec.MemoryMb, "memory", 0, "The amount of memory (in MB).")
	f.StringVar(&cmd.configSpec, "spec", "", "VirtualMachineConfigSpec json")
}

func (*create) Usage() string {
	return "NAME"
}

func (cmd *create) Description() string {
	return `Creates a new virtual machine class.

The name of the virtual machine class has DNS_LABEL restrictions
as specified in "https://tools.ietf.org/html/rfc1123". It
must be an alphanumeric (a-z and 0-9) string and with maximum length
of 63 characters and with the '-' character allowed anywhere except
the first or last character. This name is unique in this vCenter server.

Examples:
  govc namespace.vmclass.create -cpus=8 -memory=8192 test-class-01
  govc namespace.vmclass.create -spec '{"numCPUs":8,"memoryMB":8192}' test-class-02`
}

func (cmd *create) Run(ctx context.Context, f *flag.FlagSet) error {
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

	return nm.CreateVmClass(ctx, cmd.spec)
}

func mergeConfigSpec(in *namespace.VirtualMachineClassCreateSpec) error {
	if in.ConfigSpec == nil {
		return nil
	}

	spec, err := configSpecFromJSON(json.RawMessage(in.ConfigSpec))
	if err != nil {
		return err
	}

	// If these fields are set in configSpec, they must have the same value.
	// Otherwise results in 400 Bad Request + "error_type": "INVALID_ARGUMENT"
	if in.CpuCount == 0 {
		in.CpuCount = int64(spec.NumCPUs)
	}

	if in.MemoryMb == 0 {
		in.MemoryMb = spec.MemoryMB
	}

	in.ConfigSpec, err = configSpecToJSON(spec)

	return err
}
