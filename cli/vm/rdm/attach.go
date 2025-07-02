// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package rdm

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
	*flags.VirtualMachineFlag

	device string
}

func init() {
	cli.Register("vm.rdm.attach", &attach{})
}

func (cmd *attach) Register(ctx context.Context, f *flag.FlagSet) {

	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)

	f.StringVar(&cmd.device, "device", "", "Device Name")
}

func (cmd *attach) Description() string {
	return `Attach DEVICE to VM with RDM.

Examples:
  govc vm.rdm.attach -vm VM -device /vmfs/devices/disks/naa.000000000000000000000000000000000`
}

func (cmd *attach) Process(ctx context.Context) error {
	if err := cmd.VirtualMachineFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

// This piece of code was developed mainly thanks to the project govmax on github.com
// This file in particular https://github.com/codedellemc/govmax/blob/master/api/v1/vmomi.go
func (cmd *attach) Run(ctx context.Context, f *flag.FlagSet) error {
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

	controller, err := devices.FindSCSIController("")
	if err != nil {
		return err
	}

	vmConfigOptions, err := queryConfigTarget(ctx, vm)
	if err != nil {
		return err
	}

	for _, scsiDisk := range vmConfigOptions.ScsiDisk {
		if !strings.Contains(scsiDisk.Disk.CanonicalName, cmd.device) {
			continue
		}
		var backing types.VirtualDiskRawDiskMappingVer1BackingInfo
		backing.CompatibilityMode = string(types.VirtualDiskCompatibilityModePhysicalMode)
		backing.DeviceName = scsiDisk.Disk.DeviceName
		for _, descriptor := range scsiDisk.Disk.Descriptor {
			if strings.HasPrefix(descriptor.Id, "vml.") {
				backing.LunUuid = descriptor.Id
				break
			}
		}
		var device types.VirtualDisk
		device.Backing = &backing
		device.ControllerKey = controller.VirtualController.Key

		var unitNumber *int32
		scsiCtrlUnitNumber := controller.VirtualController.UnitNumber
		var u int32
		for u = 0; u < 16; u++ {
			free := true
			for _, d := range devices {
				if d.GetVirtualDevice().ControllerKey == device.GetVirtualDevice().ControllerKey {
					if u == *(d.GetVirtualDevice().UnitNumber) || u == *scsiCtrlUnitNumber {
						free = false
					}
				}
			}
			if free && u != 7 {
				unitNumber = &u
				break
			}
		}
		device.UnitNumber = unitNumber

		spec := types.VirtualMachineConfigSpec{}

		config := &types.VirtualDeviceConfigSpec{
			Device:    &device,
			Operation: types.VirtualDeviceConfigSpecOperationAdd,
		}

		config.FileOperation = types.VirtualDeviceConfigSpecFileOperationCreate

		spec.DeviceChange = append(spec.DeviceChange, config)

		task, err := vm.Reconfigure(ctx, spec)
		if err != nil {
			return err
		}

		err = task.Wait(ctx)
		if err != nil {
			return fmt.Errorf("error adding device %+v \n with backing %+v \nLogged Item:  %s", device, backing, err)
		}
		return nil

	}
	return fmt.Errorf("error: No LUN with device name containing %s found", cmd.device)
}
