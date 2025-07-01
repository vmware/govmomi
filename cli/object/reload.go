// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package object

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/types"
)

type reload struct {
	*flags.DatacenterFlag
}

func init() {
	cli.Register("object.reload", &reload{})
}

func (cmd *reload) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatacenterFlag, ctx = flags.NewDatacenterFlag(ctx)
	cmd.DatacenterFlag.Register(ctx, f)
}

func (cmd *reload) Usage() string {
	return "PATH..."
}

func (cmd *reload) Description() string {
	return `Reload managed object state.

Examples:
  govc datastore.upload $vm.vmx $vm/$vm.vmx
  govc object.reload /dc1/vm/$vm`
}

func (cmd *reload) Process(ctx context.Context) error {
	if err := cmd.DatacenterFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *reload) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() == 0 {
		return flag.ErrHelp
	}

	c, err := cmd.Client()
	if err != nil {
		return err
	}

	objs, err := cmd.ManagedObjects(ctx, f.Args())
	if err != nil {
		return err
	}

	for _, obj := range objs {
		req := types.Reload{
			This: obj,
		}

		_, err = methods.Reload(ctx, c, &req)
		if err != nil {
			return err
		}
	}

	return nil
}
