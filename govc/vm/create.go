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
}

func init() {
	cli.Register(&create{})
}

func (c *create) Register(f *flag.FlagSet) {
	f.IntVar(&c.memory, "m", 128, "Size in MB of memory")
	f.IntVar(&c.cpus, "c", 1, "Number of CPUs")
	f.StringVar(&c.guestID, "g", "otherGuest", "Guest OS")
}

func (c *create) Process() error { return nil }

func (c *create) Run(f *flag.FlagSet) error {
	if len(f.Args()) != 1 {
		return flag.ErrHelp
	}

	var pool *govmomi.ResourcePool

	host, err := c.HostSystem()
	if err != nil {
		return err
	}

	if host == nil { // -host is optional
		if pool, err = c.ResourcePool(); err != nil {
			return err
		}
	} else {
		if pool, err = c.HostResourcePool(); err != nil {
			return err
		}
	}

	ds, err := c.DatastoreName()
	if err != nil {
		return err
	}

	name := f.Arg(0)

	spec := types.VirtualMachineConfigSpec{
		Name:     name,
		GuestId:  c.guestID,
		Files:    &types.VirtualMachineFileInfo{VmPathName: fmt.Sprintf("[%s]", ds)},
		NumCPUs:  c.cpus,
		MemoryMB: int64(c.memory),
	}

	if err = c.addDisk(&spec); err != nil {
		return err
	}

	if err = c.addNetwork(&spec); err != nil {
		return err
	}

	client, err := c.DatastoreFlag.Client()
	if err != nil {
		return err
	}

	folder, err := c.VmFolder()
	if err != nil {
		return err
	}

	_, err = folder.CreateVM(client, spec, pool, host)

	return err
}

func (c *create) addDevice(spec *types.VirtualMachineConfigSpec, device types.BaseVirtualDevice) {
	spec.DeviceChange = append(spec.DeviceChange, &types.VirtualDeviceConfigSpec{
		Operation: types.VirtualDeviceConfigSpecOperationAdd,
		Device:    device,
	})
}

func (c *create) addNetwork(spec *types.VirtualMachineConfigSpec) error {
	network, err := c.NetworkFlag.Device()
	if err != nil {
		return err
	}

	c.addDevice(spec, network)

	return nil
}

func (c *create) addDisk(spec *types.VirtualMachineConfigSpec) error {
	if !c.DiskFlag.IsSet() {
		return nil
	}

	diskPath, err := c.DiskFlag.Path()
	if err != nil {
		return err
	}

	switch filepath.Ext(diskPath) {
	case ".vmdk":
		device, err := c.DiskFlag.Controller()
		if err != nil {
			return err
		}
		c.addDevice(spec, device)

		device, err = c.DiskFlag.Copy(fmt.Sprintf("%s/%s.vmdk", spec.Name, spec.Name))
		if err != nil {
			return err
		}
		c.addDevice(spec, device)
	case ".iso":
		return errors.New("TODO: .iso not supported yet")
	default:
		return errors.New("unsupported disk type")
	}

	return nil
}
