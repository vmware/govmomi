/*
Copyright (c) 2014 VMware, Inc. All Rights Reserved.

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
	"fmt"
	"reflect"
	"regexp"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

var dsPathRegexp = regexp.MustCompile(`^\[.*\] (?:.*/)?([^/]+)\.vmdk$`)

func FindController(mvm mo.VirtualMachine) (int, error) {
	for _, dev := range mvm.Config.Hardware.Device {
		switch disk := dev.(type) {
		case *types.VirtualLsiLogicController:
			vdev := disk.GetVirtualDevice()
			return vdev.Key, nil
		}
	}

	return -1, nil
}

func FindDisk(dname string, mvm mo.VirtualMachine) (*types.VirtualDisk, error) {
	for _, dev := range mvm.Config.Hardware.Device {
		switch disk := dev.(type) {
		case *types.VirtualDisk:
			switch backing := disk.Backing.(type) {
			case *types.VirtualDiskFlatVer2BackingInfo:
				m := dsPathRegexp.FindStringSubmatch(backing.FileName)
				if len(m) >= 2 && m[1] == dname {
					return disk, nil
				}
			default:
				name := reflect.TypeOf(disk.Backing).String()
				panic(fmt.Sprintf("unsupported backing: %s", name))
			}
		}
	}

	return nil, nil
}

func CreateDisk(name string, bytes ByteValue, controllerKey int, vm *govmomi.VirtualMachine) error {
	var err error

	disk := &types.VirtualDisk{
		VirtualDevice: types.VirtualDevice{
			Key: -1,
			Backing: &types.VirtualDiskFlatVer2BackingInfo{
				VirtualDeviceFileBackingInfo: types.VirtualDeviceFileBackingInfo{
					FileName: name + ".vmdk",
				},
				DiskMode:        string(types.VirtualDiskModePersistent),
				ThinProvisioned: true,
			},
			ControllerKey: controllerKey,
			UnitNumber:    -1,
		},
		CapacityInKB: bytes.Bytes / 1024,
	}

	diskAddOp := &types.VirtualDeviceConfigSpec{
		Device:        disk,
		FileOperation: types.VirtualDeviceConfigSpecFileOperationCreate,
		Operation:     types.VirtualDeviceConfigSpecOperationAdd,
	}

	spec := new(configSpec)
	spec.AddChange(diskAddOp)

	task, err := vm.Reconfigure(spec.ToSpec())
	if err != nil {
		return err
	}

	return task.Wait()
}

func AttachDisk(disk *types.VirtualDisk, vm *govmomi.VirtualMachine, ds *govmomi.Datastore, controllerKey int, link bool, persist bool) error {
	var err error

	disk.VirtualDevice.ControllerKey = controllerKey

	diskAddOp := &types.VirtualDeviceConfigSpec{
		Device:    disk,
		Operation: types.VirtualDeviceConfigSpecOperationAdd,
	}

	if link {
		LinkDisk(disk, ds, persist)
		diskAddOp.FileOperation = types.VirtualDeviceConfigSpecFileOperationCreate
	} else {
		ConfigureDisk(disk, persist)
	}

	spec := new(configSpec)
	spec.AddChange(diskAddOp)

	task, err := vm.Reconfigure(spec.ToSpec())
	if err != nil {
		return err
	}

	return task.Wait()
}

func ConfigureDisk(disk *types.VirtualDisk, persist bool) error {
	diskMode := string(types.VirtualDiskModeNonpersistent)
	if persist {
		diskMode = string(types.VirtualDiskModePersistent)
	}
	disk.Backing.(*types.VirtualDiskFlatVer2BackingInfo).DiskMode = diskMode

	return nil
}

func LinkDisk(disk *types.VirtualDisk, ds *govmomi.Datastore, persist bool) error {
	datastore := fmt.Sprintf("[%s]", ds.Name())
	parent := disk.Backing.(*types.VirtualDiskFlatVer2BackingInfo)

	diskMode := string(types.VirtualDiskModeIndependent_nonpersistent)
	if persist {
		diskMode = string(types.VirtualDiskModeIndependent_persistent)
	}

	disk.Backing = &types.VirtualDiskFlatVer2BackingInfo{
		VirtualDeviceFileBackingInfo: types.VirtualDeviceFileBackingInfo{
			FileName: datastore,
		},
		Parent:          parent,
		DiskMode:        diskMode,
		ThinProvisioned: true,
	}
	return nil
}
