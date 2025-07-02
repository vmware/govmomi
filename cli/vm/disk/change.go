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
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/units"
	"github.com/vmware/govmomi/vim25/types"
)

type change struct {
	*flags.VirtualMachineFlag

	name     string
	key      int
	label    string
	filePath string
	sharing  string

	bytes units.ByteSize
	mode  string

	// 	SIOC
	limit *int64
}

func init() {
	cli.Register("vm.disk.change", &change{})
}

func (cmd *change) Description() string {
	return `Change some properties of a VM's DISK

In particular, you can change the DISK mode, and the size (as long as it is bigger)

Examples:
  govc vm.disk.change -vm VM -disk.key 2001 -size 10G
  govc vm.disk.change -vm VM -disk.label "BDD disk" -size 10G
  govc vm.disk.change -vm VM -disk.name "hard-1000-0" -size 12G
  govc vm.disk.change -vm VM -disk.filePath "[DS] VM/VM-1.vmdk" -mode nonpersistent`
}

func (cmd *change) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)
	err := (&cmd.bytes).Set("0G")
	if err != nil {
		panic(err)
	}
	f.Var(&cmd.bytes, "size", "New disk size")
	f.StringVar(&cmd.name, "disk.name", "", "Disk name")
	f.StringVar(&cmd.label, "disk.label", "", "Disk label")
	f.StringVar(&cmd.filePath, "disk.filePath", "", "Disk file name")
	f.IntVar(&cmd.key, "disk.key", 0, "Disk unique key")
	f.StringVar(&cmd.mode, "mode", "", fmt.Sprintf("Disk mode (%s)", strings.Join(vdmTypes, "|")))
	f.StringVar(&cmd.sharing, "sharing", "", fmt.Sprintf("Sharing (%s)", strings.Join(sharing, "|")))
	f.Var(flags.NewOptionalInt64(&cmd.limit), "disk.io.limit", "Disk storage IO per seconds limit (-1 for unlimited)")
}

func (cmd *change) Process(ctx context.Context) error {
	if err := cmd.VirtualMachineFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *change) FindDisk(ctx context.Context, list object.VirtualDeviceList) (*types.VirtualDisk, error) {
	var disks []*types.VirtualDisk
	for _, device := range list {
		switch md := device.(type) {
		case *types.VirtualDisk:
			if cmd.CheckDiskProperties(ctx, list.Name(device), md) {
				disks = append(disks, md)
			}
		default:
			continue
		}
	}

	switch len(disks) {
	case 0:
		return nil, errors.New("no disk found using the given values")
	case 1:
		return disks[0], nil
	}
	return nil, errors.New("the given disk values match multiple disks")
}

func (cmd *change) CheckDiskProperties(ctx context.Context, name string, disk *types.VirtualDisk) bool {
	switch {
	case cmd.key != 0 && disk.Key != int32(cmd.key):
		fallthrough
	case cmd.name != "" && name != cmd.name:
		fallthrough
	case cmd.label != "" && disk.DeviceInfo.GetDescription().Label != cmd.label:
		return false
	case cmd.filePath != "":
		if b, ok := disk.Backing.(types.BaseVirtualDeviceFileBackingInfo); ok {
			if b.GetVirtualDeviceFileBackingInfo().FileName != cmd.filePath {
				return false
			}
		}
	}
	return true
}

func (cmd *change) setMode(mode *string) {
	if cmd.mode != "" {
		*mode = cmd.mode
	}
}

func (cmd *change) Run(ctx context.Context, f *flag.FlagSet) error {
	vm, err := cmd.VirtualMachine()
	if err != nil {
		return err
	}

	if vm == nil {
		return flag.ErrHelp
	}

	devices, err := vm.Device(ctx)
	if err != nil {
		return err
	}

	editdisk, err := cmd.FindDisk(ctx, devices)
	if err != nil {
		return err
	}

	if int64(cmd.bytes) != 0 {
		editdisk.CapacityInBytes = int64(cmd.bytes)
	}

	if editdisk.StorageIOAllocation == nil {
		editdisk.StorageIOAllocation = new(types.StorageIOAllocationInfo)
	}
	editdisk.StorageIOAllocation.Limit = cmd.limit

	switch backing := editdisk.Backing.(type) {
	case *types.VirtualDiskFlatVer2BackingInfo:
		backing.Sharing = cmd.sharing
		cmd.setMode(&backing.DiskMode)
	case *types.VirtualDiskSeSparseBackingInfo:
		cmd.setMode(&backing.DiskMode)
	}

	spec := types.VirtualMachineConfigSpec{}

	config := &types.VirtualDeviceConfigSpec{
		Device:    editdisk,
		Operation: types.VirtualDeviceConfigSpecOperationEdit,
	}

	config.FileOperation = ""

	spec.DeviceChange = append(spec.DeviceChange, config)

	task, err := vm.Reconfigure(ctx, spec)
	if err != nil {
		return err
	}

	err = task.Wait(ctx)
	if err != nil {
		return fmt.Errorf("error resizing main disk\nLogged Item:  %s", err)
	}
	return nil
}
