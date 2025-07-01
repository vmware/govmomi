// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package disk

import (
	"context"
	"flag"
	"fmt"
	"strings"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vim25/types"
)

type attach struct {
	*flags.DatastoreFlag
	*flags.VirtualMachineFlag
	*flags.StorageProfileFlag

	persist    bool
	link       bool
	disk       string
	controller string
	mode       string
	sharing    string
}

func init() {
	cli.Register("vm.disk.attach", &attach{})
}

func (cmd *attach) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatastoreFlag, ctx = flags.NewDatastoreFlag(ctx)
	cmd.DatastoreFlag.Register(ctx, f)
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)
	cmd.StorageProfileFlag, ctx = flags.NewStorageProfileFlag(ctx)
	cmd.StorageProfileFlag.Register(ctx, f)

	f.BoolVar(&cmd.persist, "persist", true, "Persist attached disk")
	f.BoolVar(&cmd.link, "link", true, "Link specified disk")
	f.StringVar(&cmd.controller, "controller", "", "Disk controller")
	f.StringVar(&cmd.disk, "disk", "", "Disk path name")
	f.StringVar(&cmd.mode, "mode", "", fmt.Sprintf("Disk mode (%s)", strings.Join(vdmTypes, "|")))
	f.StringVar(&cmd.sharing, "sharing", "", fmt.Sprintf("Sharing (%s)", strings.Join(sharing, "|")))
}

func (cmd *attach) Process(ctx context.Context) error {
	if err := cmd.DatastoreFlag.Process(ctx); err != nil {
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

func (cmd *attach) Description() string {
	return `Attach existing disk to VM.

A delta disk is created by default, where changes are persisted. Specifying '-link=false' will persist to the same disk.

Examples:
  govc vm.disk.attach -vm $name -disk $name/disk1.vmdk
  govc device.info -vm $name disk-* # 'File' field is where changes are persisted. 'Parent' field is set when '-link=true'
  govc vm.disk.attach -vm $name -disk $name/shared.vmdk -link=false -sharing sharingMultiWriter
  govc device.remove -vm $name -keep disk-* # detach disk(s)`
}

func (cmd *attach) Run(ctx context.Context, f *flag.FlagSet) error {
	vm, err := cmd.VirtualMachine()
	if err != nil {
		return err
	}

	if vm == nil {
		return flag.ErrHelp
	}

	ds, err := cmd.Datastore()
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

	disk := devices.CreateDisk(controller, ds.Reference(), ds.Path(cmd.disk))
	backing := disk.Backing.(*types.VirtualDiskFlatVer2BackingInfo)
	backing.Sharing = cmd.sharing

	if cmd.link {
		if cmd.persist {
			backing.DiskMode = string(types.VirtualDiskModeIndependent_persistent)
		} else {
			backing.DiskMode = string(types.VirtualDiskModeIndependent_nonpersistent)
		}

		disk = devices.ChildDisk(disk)
		return vm.AddDevice(ctx, disk)
	} else {
		if cmd.persist {
			backing.DiskMode = string(types.VirtualDiskModePersistent)
		} else {
			backing.DiskMode = string(types.VirtualDiskModeNonpersistent)
		}
	}

	if len(cmd.mode) != 0 {
		backing.DiskMode = cmd.mode
	}

	profile, err := cmd.StorageProfileSpec(ctx)
	if err != nil {
		return err
	}

	return vm.AddDeviceWithProfile(ctx, profile, disk)
}
