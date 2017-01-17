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

package rdm

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"strings"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
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

	f.StringVar(&cmd.device, "deviceName", "", "Device Name")
}

func (cmd *attach) Process(ctx context.Context) error {
	if err := cmd.VirtualMachineFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

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
	var VM_withProp mo.VirtualMachine
	err = vm.Properties(ctx, vm.Reference(), []string{"environmentBrowser"}, &VM_withProp)
	if err != nil {
		return err
	}

	//Query VM To Find Devices avilable for attaching to VM
	var queryConfigRequest types.QueryConfigTarget
	queryConfigRequest.This = VM_withProp.EnvironmentBrowser
	cl, err := cmd.Client()
	queryConfigResp, err := methods.QueryConfigTarget(ctx, cl, &queryConfigRequest)
	if err != nil {
		return err
	}
	vmConfigOptions := *queryConfigResp.Returnval
	for _, ScsiDisk := range vmConfigOptions.ScsiDisk {
		if !strings.Contains(ScsiDisk.Disk.CanonicalName, cmd.device) {
			continue
		}
		var backing types.VirtualDiskRawDiskMappingVer1BackingInfo
		backing.CompatibilityMode = "physicalMode"
		backing.DeviceName = ScsiDisk.Disk.DeviceName
		for _, descriptor := range ScsiDisk.Disk.Descriptor {
			if string([]rune(descriptor.Id)[:4]) == "vml." {
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
			for _, device := range devices {
				if device.GetVirtualDevice().ControllerKey == device.GetVirtualDevice().ControllerKey {
					if u == *(device.GetVirtualDevice().UnitNumber) || u == *scsiCtrlUnitNumber {
						free = false
					}
				}
			}
			if free {
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
			return errors.New(fmt.Sprintf("Error adding device %+v \n with backing %+v \nLogged Item:  %s", device, backing, err))
		}
		return nil

	}
	return errors.New(fmt.Sprintf("Error: No LUN with device name containing %s found.", cmd.device))
}
