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

type create struct {
	*flags.ClientFlag

	spec namespace.VirtualMachineClassCreateSpec
}

func init() {
	cli.Register("namespace.vmclass.create", &create{})
}

func (cmd *create) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.Int64Var(&cmd.spec.CpuCount, "cpus", 0, "The number of CPUs.")
	f.Int64Var(&cmd.spec.MemoryMb, "memory", 0, "The amount of memory (in MB).")
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
  govc namespace.vmclass.create -cpus=8 -memory=8192 test-class-01`
}

func (cmd *create) Run(ctx context.Context, f *flag.FlagSet) error {
	cmd.spec.Id = f.Arg(0)
	if f.NArg() != 1 {
		return flag.ErrHelp
	}

	rc, err := cmd.RestClient()
	if err != nil {
		return err
	}

	nm := namespace.NewManager(rc)

	return nm.CreateVmClass(ctx, cmd.spec)
}
