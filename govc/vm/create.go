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

package vm

import (
	"errors"
	"flag"
	"fmt"
	"path/filepath"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vim25/types"
)

type create struct {
	*flags.ResourcePoolFlag
	*flags.HostSystemFlag
	*flags.DatastoreFlag
	*flags.VmFolderFlag
	*flags.NetworkFlag
	*flags.DiskFlag

	memory  int
	cpus    int
	guestID string
	on      bool
}

func init() {
	cli.Register(&create{})
}

func (cmd *create) Register(f *flag.FlagSet) {
	f.IntVar(&cmd.memory, "m", 128, "Size in MB of memory")
	f.IntVar(&cmd.cpus, "c", 1, "Number of CPUs")
	f.StringVar(&cmd.guestID, "g", "otherGuest", "Guest OS")
	f.BoolVar(&cmd.on, "on", true, "Power on VM. Default is true if -disk argument is given.")
}

func (cmd *create) Process() error { return nil }

func (cmd *create) Run(f *flag.FlagSet) error {
	if len(f.Args()) != 1 {
		return flag.ErrHelp
	}

	var pool *govmomi.ResourcePool

	host, err := cmd.HostSystem()
	if err != nil {
		return err
	}

	if host == nil { // -host is optional
		if pool, err = cmd.ResourcePool(); err != nil {
			return err
		}
	} else {
		if pool, err = cmd.HostResourcePool(); err != nil {
			return err
		}
	}

	ds, err := cmd.DatastoreName()
	if err != nil {
		return err
	}

	name := f.Arg(0)

	spec := types.VirtualMachineConfigSpec{
		Name:     name,
		GuestId:  cmd.guestID,
		Files:    &types.VirtualMachineFileInfo{VmPathName: fmt.Sprintf("[%s]", ds)},
		NumCPUs:  cmd.cpus,
		MemoryMB: int64(cmd.memory),
	}

	if err = cmd.addDisk(&spec); err != nil {
		return err
	}

	if err = cmd.addNetwork(&spec); err != nil {
		return err
	}

	c, err := cmd.DatastoreFlag.Client()
	if err != nil {
		return err
	}

	folder, err := cmd.VmFolder()
	if err != nil {
		return err
	}

	vm, err := folder.CreateVM(c, spec, pool, host)
	if err != nil {
		return err
	}

	if cmd.DiskFlag.IsSet() && cmd.on {
		return vm.PowerOn(c)
	}

	return nil
}

func (cmd *create) addDevice(spec *types.VirtualMachineConfigSpec, device types.BaseVirtualDevice) {
	spec.DeviceChange = append(spec.DeviceChange, &types.VirtualDeviceConfigSpec{
		Operation: types.VirtualDeviceConfigSpecOperationAdd,
		Device:    device,
	})
}

func (cmd *create) addNetwork(spec *types.VirtualMachineConfigSpec) error {
	network, err := cmd.NetworkFlag.Device()
	if err != nil {
		return err
	}

	cmd.addDevice(spec, network)

	return nil
}

func (cmd *create) addDisk(spec *types.VirtualMachineConfigSpec) error {
	if !cmd.DiskFlag.IsSet() {
		return nil
	}

	diskPath, err := cmd.DiskFlag.Path()
	if err != nil {
		return err
	}

	switch filepath.Ext(diskPath) {
	case ".vmdk":
		device, err := cmd.DiskFlag.Controller()
		if err != nil {
			return err
		}
		cmd.addDevice(spec, device)

		device, err = cmd.DiskFlag.Copy(fmt.Sprintf("%s/%s.vmdk", spec.Name, spec.Name))
		if err != nil {
			return err
		}
		cmd.addDevice(spec, device)
	case ".iso":
		device, err := cmd.DiskFlag.Cdrom(diskPath)
		if err != nil {
			return err
		}
		cmd.addDevice(spec, device)
	default:
		return errors.New("unsupported disk type")
	}

	return nil
}
