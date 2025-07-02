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
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type disk struct {
	*flags.DatastoreFlag
	*flags.ResourcePoolFlag
	*flags.StoragePodFlag
	*flags.StorageProfileFlag

	size units.ByteSize
	keep *bool
}

func init() {
	cli.Register("disk.create", &disk{})
}

func (cmd *disk) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatastoreFlag, ctx = flags.NewDatastoreFlag(ctx)
	cmd.DatastoreFlag.Register(ctx, f)

	cmd.StoragePodFlag, ctx = flags.NewStoragePodFlag(ctx)
	cmd.StoragePodFlag.Register(ctx, f)

	cmd.StorageProfileFlag, ctx = flags.NewStorageProfileFlag(ctx)
	cmd.StorageProfileFlag.Register(ctx, f)

	cmd.ResourcePoolFlag, ctx = flags.NewResourcePoolFlag(ctx)
	cmd.ResourcePoolFlag.Register(ctx, f)

	_ = cmd.size.Set("10G")
	f.Var(&cmd.size, "size", "Size of new disk")
	f.Var(flags.NewOptionalBool(&cmd.keep), "keep", "Keep disk after VM is deleted")
}

func (cmd *disk) Process(ctx context.Context) error {
	if err := cmd.DatastoreFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.StoragePodFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.StorageProfileFlag.Process(ctx); err != nil {
		return err
	}
	return cmd.ResourcePoolFlag.Process(ctx)
}

func (cmd *disk) Usage() string {
	return "NAME"
}

func (cmd *disk) Description() string {
	return `Create disk NAME on DS.

Examples:
  govc disk.create -size 24G my-disk`
}

func (cmd *disk) Run(ctx context.Context, f *flag.FlagSet) error {
	name := f.Arg(0)
	if name == "" {
		return flag.ErrHelp
	}

	m, err := NewManagerFromFlag(ctx, cmd.DatastoreFlag)
	if err != nil {
		return err
	}

	var pool *object.ResourcePool
	var ds mo.Reference
	if cmd.StoragePodFlag.Isset() {
		ds, err = cmd.StoragePod()
		if err != nil {
			return err
		}
		pool, err = cmd.ResourcePool()
		if err != nil {
			return err
		}
	} else {
		ds, err = cmd.Datastore()
		if err != nil {
			return err
		}
	}

	profile, err := cmd.StorageProfileSpec(ctx)
	if err != nil {
		return err
	}

	spec := types.VslmCreateSpec{
		Name:              name,
		CapacityInMB:      int64(cmd.size) / units.MB,
		KeepAfterDeleteVm: cmd.keep,
		Profile:           profile,
		BackingSpec: &types.VslmCreateSpecDiskFileBackingSpec{
			VslmCreateSpecBackingSpec: types.VslmCreateSpecBackingSpec{
				Datastore: ds.Reference(),
			},
			ProvisioningType: string(types.BaseConfigInfoDiskFileBackingInfoProvisioningTypeThin),
		},
	}

	if cmd.StoragePodFlag.Isset() {
		if err = m.ObjectManager.PlaceDisk(ctx, &spec, pool.Reference()); err != nil {
			return err
		}
	}

	obj, err := m.CreateDisk(ctx, spec)
	if err != nil {
		return err
	}

	fmt.Println(obj.Config.Id.Id)

	return nil
}
