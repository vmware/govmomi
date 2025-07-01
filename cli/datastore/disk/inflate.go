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

type inflate struct {
	*flags.DatastoreFlag
}

func init() {
	cli.Register("datastore.disk.inflate", &inflate{})
}

func (cmd *inflate) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatastoreFlag, ctx = flags.NewDatastoreFlag(ctx)
	cmd.DatastoreFlag.Register(ctx, f)
}

func (cmd *inflate) Process(ctx context.Context) error {
	return cmd.DatastoreFlag.Process(ctx)
}

func (cmd *inflate) Usage() string {
	return "VMDK"
}

func (cmd *inflate) Description() string {
	return `Inflate VMDK on DS.

Examples:
  govc datastore.disk.inflate disks/disk1.vmdk`
}

func (cmd *inflate) Run(ctx context.Context, f *flag.FlagSet) error {
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
	task, err := m.InflateVirtualDisk(ctx, path, dc)
	if err != nil {
		return err
	}

	logger := cmd.ProgressLogger(fmt.Sprintf("Inflating %s...", path))
	defer logger.Wait()

	_, err = task.WaitForResult(ctx, logger)
	return err
}
