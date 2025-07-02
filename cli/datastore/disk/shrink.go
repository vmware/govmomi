// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package disk

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/object"
)

type shrink struct {
	*flags.DatastoreFlag

	copy *bool
}

func init() {
	cli.Register("datastore.disk.shrink", &shrink{})
}

func (cmd *shrink) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatastoreFlag, ctx = flags.NewDatastoreFlag(ctx)
	cmd.DatastoreFlag.Register(ctx, f)

	f.Var(flags.NewOptionalBool(&cmd.copy), "copy", "Perform shrink in-place mode if false, copy-shrink mode otherwise")
}

func (cmd *shrink) Process(ctx context.Context) error {
	return cmd.DatastoreFlag.Process(ctx)
}

func (cmd *shrink) Usage() string {
	return "VMDK"
}

func (cmd *shrink) Description() string {
	return `Shrink VMDK on DS.

Examples:
  govc datastore.disk.shrink disks/disk1.vmdk`
}

func (cmd *shrink) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() == 0 {
		return flag.ErrHelp
	}

	dc, err := cmd.Datacenter()
	if err != nil {
		return err
	}

	ds, err := cmd.Datastore()
	if err != nil {
		return err
	}

	m := object.NewVirtualDiskManager(ds.Client())
	path := ds.Path(f.Arg(0))
	task, err := m.ShrinkVirtualDisk(ctx, path, dc, cmd.copy)
	if err != nil {
		return err
	}

	logger := cmd.ProgressLogger(fmt.Sprintf("Shrinking %s...", path))
	defer logger.Wait()

	_, err = task.WaitForResult(ctx, logger)
	return err
}
