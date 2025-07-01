// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package disk

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"strings"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/units"
	"github.com/vmware/govmomi/vim25/types"
)

type create struct {
	*flags.DatastoreFlag
	*flags.OutputFlag
	*flags.VirtualMachineFlag
	*flags.StorageProfileFlag

	controller string
	Name       string
	Bytes      units.ByteSize
	Thick      bool
	Eager      bool
	DiskMode   string
	Sharing    string
}

var vdmTypes = types.VirtualDiskMode("").Strings()

var sharing = types.VirtualDiskSharing("").Strings()

func init() {
	cli.Register("vm.disk.create", &create{})
}

func (cmd *create) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatastoreFlag, ctx = flags.NewDatastoreFlag(ctx)
	cmd.DatastoreFlag.Register(ctx, f)
	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)
	cmd.StorageProfileFlag, ctx = flags.NewStorageProfileFlag(ctx)
	cmd.StorageProfileFlag.Register(ctx, f)

	err := (&cmd.Bytes).Set("10G")
	if err != nil {
		panic(err)
	}

	f.StringVar(&cmd.controller, "controller", "", "Disk controller")
	f.StringVar(&cmd.Name, "name", "", "Name for new disk")
	f.Var(&cmd.Bytes, "size", "Size of new disk")
	f.BoolVar(&cmd.Thick, "thick", false, "Thick provision new disk")
	f.BoolVar(&cmd.Eager, "eager", false, "Eagerly scrub new disk")
	f.StringVar(&cmd.DiskMode, "mode", vdmTypes[0], fmt.Sprintf("Disk mode (%s)", strings.Join(vdmTypes, "|")))
	f.StringVar(&cmd.Sharing, "sharing", "", fmt.Sprintf("Sharing (%s)", strings.Join(sharing, "|")))
}

func (cmd *create) Process(ctx context.Context) error {
	if err := cmd.DatastoreFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.VirtualMachineFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.StorageProfileFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *create) Description() string {
	return `Create disk and attach to VM.

Examples:
  govc vm.disk.create -vm $name -name $name/disk1 -size 10G
  govc vm.disk.create -vm $name -name $name/disk2 -size 10G -eager -thick -sharing sharingMultiWriter`
}

func (cmd *create) Run(ctx context.Context, f *flag.FlagSet) error {
	if len(cmd.Name) == 0 {
		return errors.New("please specify a disk name")
	}

	vm, err := cmd.VirtualMachine()
	if err != nil {
		return err
	}
	if vm == nil {
		return errors.New("please specify a vm")
	}

	ds, err := cmd.Datastore()
	if err != nil {
		return err
	}

	profile, err := cmd.StorageProfileSpec(ctx)
	if err != nil {
		return err
	}

	devices, err := vm.Device(ctx)
	if err != nil {
		return err
	}

	controller, err := devices.FindDiskController(cmd.controller)
	if err != nil {
		return err
	}

	vdmMatch := false
	for _, vdm := range vdmTypes {
		if cmd.DiskMode == vdm {
			vdmMatch = true
		}
	}

	if !vdmMatch {
		return errors.New("please specify a valid disk mode")
	}

	disk := devices.CreateDisk(controller, ds.Reference(), ds.Path(cmd.Name))

	existing := devices.SelectByBackingInfo(disk.Backing)

	if len(existing) > 0 {
		_, _ = cmd.Log("Disk already present\n")
		return nil
	}

	backing := disk.Backing.(*types.VirtualDiskFlatVer2BackingInfo)

	if cmd.Thick {
		backing.ThinProvisioned = types.NewBool(false)
		backing.EagerlyScrub = types.NewBool(cmd.Eager)
	}

	backing.DiskMode = cmd.DiskMode
	backing.Sharing = cmd.Sharing

	_, _ = cmd.Log("Creating disk\n")
	disk.CapacityInKB = int64(cmd.Bytes) / 1024
	return vm.AddDeviceWithProfile(ctx, profile, disk)
}
