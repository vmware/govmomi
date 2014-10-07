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
	"errors"
	"flag"
	"fmt"
	"reflect"
	"regexp"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type add struct {
	*flags.VirtualMachineFlag

	Name  string
	Bytes ByteValue

	Client         *govmomi.Client
	VirtualMachine *govmomi.VirtualMachine
}

func init() {
	cli.Register("vm.disk.add", &add{})
}

func (cmd *add) Register(f *flag.FlagSet) {
	err := (&cmd.Bytes).Set("10G")
	if err != nil {
		panic(err)
	}

	f.StringVar(&cmd.Name, "name", "", "Name for new disk")
	f.Var(&cmd.Bytes, "size", "Size of new disk")
}

func (cmd *add) Process() error { return nil }

func (cmd *add) Run(f *flag.FlagSet) error {
	var err error

	cmd.Client, err = cmd.ClientFlag.Client()
	if err != nil {
		return err
	}

	cmd.VirtualMachine, err = cmd.VirtualMachineFlag.VirtualMachine()
	if err != nil {
		return err
	}

	if cmd.VirtualMachine == nil {
		return errors.New("please specify a vm")
	}

	var mvm mo.VirtualMachine

	err = cmd.Client.Properties(cmd.VirtualMachine.Reference(), []string{"config.hardware"}, &mvm)
	if err != nil {
		return err
	}

	dev, err := cmd.FindDisk(mvm)
	if err != nil {
		return err
	}

	if dev == nil {
		cmd.Log("Creating disk\n")
		err = cmd.CreateDisk(mvm)
		if err != nil {
			return err
		}
	} else {
		cmd.Log("Disk already present\n")
	}

	return nil
}

var dsPathRegexp = regexp.MustCompile(`^\[.*\] (?:.*/)?([^/]+)\.vmdk$`)

func (cmd *add) FindDisk(mvm mo.VirtualMachine) (*types.VirtualDisk, error) {
	for _, dev := range mvm.Config.Hardware.Device {
		switch disk := dev.(type) {
		case *types.VirtualDisk:
			switch backing := disk.Backing.(type) {
			case *types.VirtualDiskFlatVer2BackingInfo:
				m := dsPathRegexp.FindStringSubmatch(backing.FileName)
				if len(m) >= 2 && m[1] == cmd.Name {
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

func (cmd *add) FindController(mvm mo.VirtualMachine) (int, error) {
	for _, dev := range mvm.Config.Hardware.Device {
		switch disk := dev.(type) {
		case *types.VirtualLsiLogicController:
			vdev := disk.GetVirtualDevice()
			return vdev.Key, nil
		}
	}

	return -1, nil
}

func (cmd *add) CreateDisk(mvm mo.VirtualMachine) error {
	controllerKey, err := cmd.FindController(mvm)
	if err != nil {
		return err
	}

	disk := &types.VirtualDisk{
		VirtualDevice: types.VirtualDevice{
			Key: -1,
			Backing: &types.VirtualDiskFlatVer2BackingInfo{
				VirtualDeviceFileBackingInfo: types.VirtualDeviceFileBackingInfo{
					FileName: cmd.Name + ".vmdk",
				},
				DiskMode:        string(types.VirtualDiskModePersistent),
				ThinProvisioned: true,
			},
			ControllerKey: controllerKey,
			UnitNumber:    -1,
		},
		CapacityInKB: cmd.Bytes.Bytes / 1024,
	}

	diskAddOp := &types.VirtualDeviceConfigSpec{
		Device:        disk,
		FileOperation: types.VirtualDeviceConfigSpecFileOperationCreate,
		Operation:     types.VirtualDeviceConfigSpecOperationAdd,
	}

	spec := new(configSpec)
	spec.AddChange(diskAddOp)

	task, err := cmd.VirtualMachine.Reconfigure(spec.ToSpec())
	if err != nil {
		return err
	}

	return task.Wait()
}
