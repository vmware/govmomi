// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vmclass

import (
	"context"
	"flag"
	"fmt"
	"io"

	"github.com/vmware/govmomi/vapi/namespace"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
)

type lsResult []namespace.VirtualMachineClassInfo

func (r lsResult) Write(w io.Writer) error {
	for _, e := range r {
		fmt.Fprintln(w, e.Id)
	}
	return nil
}

type ls struct {
	*flags.ClientFlag
	*flags.OutputFlag
}

func init() {
	cli.Register("namespace.vmclass.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)
}

func (cmd *ls) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}

	return nil
}

func (cmd *ls) Description() string {
	return `Displays the list of virtual machine classes.

Examples:
  govc namespace.vmclass.ls`
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {
	rc, err := cmd.RestClient()
	if err != nil {
		return err
	}

	nm := namespace.NewManager(rc)

	d, err := nm.ListVmClasses(ctx)
	if err != nil {
		return err
	}

	return cmd.WriteResult(lsResult(d))
}
