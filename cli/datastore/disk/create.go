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
	"github.com/vmware/govmomi/units"
	"github.com/vmware/govmomi/vim25/types"
)

type spec struct {
	types.FileBackedVirtualDiskSpec
	force bool
	uuid  string
}

func (s *spec) Register(ctx context.Context, f *flag.FlagSet) {
	f.StringVar(&s.AdapterType, "a", string(types.VirtualDiskAdapterTypeLsiLogic), "Disk adapter")
	f.StringVar(&s.DiskType, "d", string(types.VirtualDiskTypeThin), "Disk format")
	f.BoolVar(&s.force, "f", false, "Force")
	f.StringVar(&s.uuid, "uuid", "", "Disk UUID")
}

type create struct {
	*flags.DatastoreFlag

	Bytes units.ByteSize
	spec
}

func init() {
	cli.Register("datastore.disk.create", &create{})
}

func (cmd *create) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatastoreFlag, ctx = flags.NewDatastoreFlag(ctx)
	cmd.DatastoreFlag.Register(ctx, f)

	_ = cmd.Bytes.Set("10G")
	f.Var(&cmd.Bytes, "size", "Size of new disk")

	cmd.spec.Register(ctx, f)
}

func (cmd *create) Usage() string {
	return "VMDK"
}

func (cmd *create) Description() string {
	return `Create VMDK on DS.

Examples:
  govc datastore.mkdir disks
  govc datastore.disk.create -size 24G disks/disk1.vmdk
  govc datastore.disk.create disks/parent.vmdk disk/child.vmdk`
}

func (cmd *create) Run(ctx context.Context, f *flag.FlagSet) error {
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

	var task *object.Task
	var dst string

	if f.NArg() == 1 {
		cmd.spec.CapacityKb = int64(cmd.Bytes) / 1024
		dst = ds.Path(f.Arg(0))
		task, err = m.CreateVirtualDisk(ctx, dst, dc, &cmd.spec.FileBackedVirtualDiskSpec)
	} else {
		dst = ds.Path(f.Arg(0))
		task, err = m.CreateChildDisk(ctx, ds.Path(f.Arg(0)), dc, dst, dc, cmd.force)
	}

	if err != nil {
		return err
	}

	logger := cmd.ProgressLogger(fmt.Sprintf("Creating %s...", dst))
	defer logger.Wait()

	if _, err = task.WaitForResult(ctx, logger); err != nil {
		return err
	}

	if cmd.uuid != "" {
		return m.SetVirtualDiskUuid(ctx, dst, dc, cmd.uuid)
	}

	return nil
}
