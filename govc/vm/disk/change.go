/*
Copyright (c) 2014-2015 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package disk

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"strings"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
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

	bytes units.ByteSize
	mode  string
}

func init() {
	cli.Register("vm.disk.change", &change{})
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
			switch {
			case cmd.key != 0 && md.Key == int32(cmd.key):
				fallthrough
			case cmd.name != "" && list.Name(device) == cmd.name:
				fallthrough
			case cmd.label != "" && md.DeviceInfo.GetDescription().Label == cmd.label:
				disks = append(disks, md)
			case cmd.filePath != "":
				if b, ok := md.Backing.(types.BaseVirtualDeviceFileBackingInfo); ok {
					if b.GetVirtualDeviceFileBackingInfo().FileName == cmd.filePath {
						disks = append(disks, md)
					}
				}
			}
		default:
			continue
		}
	}

	switch {
	case len(disks) == 1:
		return disks[0], nil
	case len(disks) == 0:
		return nil, errors.New("No disk found using the given values")
	case len(disks) > 1:
		return nil, errors.New("The given disk values match multiple disks")
	}
	return nil, nil
}

func (cmd *change) CheckDiskProperties(ctx context.Context, name string, disk *types.VirtualDisk) error {
	switch {
	case cmd.key != 0 && disk.Key != int32(cmd.key):
		fallthrough
	case cmd.name != "" && name != cmd.name:
		fallthrough
	case cmd.label != "" && disk.DeviceInfo.GetDescription().Label != cmd.label:
		return errors.New("No disk found using the given values")
	case cmd.filePath != "":
		if b, ok := disk.Backing.(types.BaseVirtualDeviceFileBackingInfo); ok {
			if b.GetVirtualDeviceFileBackingInfo().FileName != cmd.filePath {
				return errors.New("No disk found using the given values")
			}
		}
	}
	return nil
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

	err = cmd.CheckDiskProperties(ctx, devices.Name(editdisk), editdisk)
	if err != nil {
		return err
	}
	if int64(cmd.bytes) != 0 {
		editdisk.CapacityInKB = int64(cmd.bytes) / 1024
	}

	backing := editdisk.Backing.(*types.VirtualDiskFlatVer2BackingInfo)

	if len(cmd.mode) != 0 {
		backing.DiskMode = cmd.mode
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
		return errors.New(fmt.Sprintf("Error resizing main disk\nLogged Item:  %s", err))
	}
	return nil
}
