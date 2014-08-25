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
	"flag"
	"fmt"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type create struct {
	*flags.ClientFlag
	*flags.DatacenterFlag
	*flags.DatastoreFlag
	*flags.ResourcePoolFlag
	*flags.HostSystemFlag
	*flags.DiskFlag
	*flags.NetworkFlag

	memory  int
	cpus    int
	guestID string
	link    bool
	on      bool

	Client       *govmomi.Client
	Datacenter   *govmomi.Datacenter
	Datastore    *govmomi.Datastore
	ResourcePool *govmomi.ResourcePool
	HostSystem   *govmomi.HostSystem
}

func init() {
	cli.Register(&create{})
}

func (cmd *create) Register(f *flag.FlagSet) {
	f.IntVar(&cmd.memory, "m", 128, "Size in MB of memory")
	f.IntVar(&cmd.cpus, "c", 1, "Number of CPUs")
	f.StringVar(&cmd.guestID, "g", "otherGuest", "Guest OS")
	f.BoolVar(&cmd.link, "link", true, "Link specified disk")
	f.BoolVar(&cmd.on, "on", true, "Power on VM. Default is true if -disk argument is given.")
}

func (cmd *create) Process() error { return nil }

func (cmd *create) Run(f *flag.FlagSet) error {
	var err error

	if len(f.Args()) != 1 {
		return flag.ErrHelp
	}

	cmd.Client, err = cmd.ClientFlag.Client()
	if err != nil {
		return err
	}

	cmd.Datacenter, err = cmd.DatacenterFlag.Datacenter()
	if err != nil {
		return err
	}

	cmd.Datastore, err = cmd.DatastoreFlag.Datastore()
	if err != nil {
		return err
	}

	cmd.HostSystem, err = cmd.HostSystemFlag.HostSystem()
	if err != nil {
		return err
	}

	if cmd.HostSystem == nil { // -host is optional
		if cmd.ResourcePool, err = cmd.ResourcePoolFlag.ResourcePool(); err != nil {
			return err
		}
	} else {
		if cmd.ResourcePool, err = cmd.HostSystemFlag.ResourcePool(); err != nil {
			return err
		}
	}

	task, err := cmd.CreateVM(f.Arg(0))
	if err != nil {
		return err
	}

	info, err := task.WaitForResult(nil)
	if err != nil {
		return err
	}

	vm := govmomi.NewVirtualMachine(info.Result.(types.ManagedObjectReference))

	if cmd.on {
		task, err := vm.PowerOn(cmd.Client)
		if err != nil {
			return err
		}

		_, err = task.WaitForResult(nil)
		if err != nil {
			return err
		}
	}

	return nil
}

func (cmd *create) CreateVM(name string) (*govmomi.Task, error) {
	spec := &configSpec{
		Name:     name,
		GuestId:  cmd.guestID,
		Files:    &types.VirtualMachineFileInfo{VmPathName: fmt.Sprintf("[%s]", cmd.Datastore.Name())},
		NumCPUs:  cmd.cpus,
		MemoryMB: int64(cmd.memory),
	}

	if cmd.DiskFlag.IsSet() {
		controller, err := cmd.DiskFlag.Controller()
		if err != nil {
			return nil, err
		}

		spec.AddDevice(controller)

		disk, err := cmd.DiskFlag.Disk()
		if err != nil {
			return nil, err
		}

		diskAddOp := &types.VirtualDeviceConfigSpec{
			Operation: types.VirtualDeviceConfigSpecOperationAdd,
			Device:    disk,
		}

		if cmd.link {
			parent := disk.Backing.(*types.VirtualDiskFlatVer2BackingInfo)

			// Use specified disk as parent backing to a new disk.
			disk.Backing = &types.VirtualDiskFlatVer2BackingInfo{
				VirtualDeviceFileBackingInfo: types.VirtualDeviceFileBackingInfo{
					FileName: fmt.Sprintf("[%s]", cmd.Datastore.Name()),
				},
				Parent:          parent,
				DiskMode:        string(types.VirtualDiskModePersistent),
				ThinProvisioned: true,
			}

			// Create the new disk (won't happen without this flag).
			diskAddOp.FileOperation = types.VirtualDeviceConfigSpecFileOperationCreate
		}

		spec.AddChange(diskAddOp)
	}

	netdev, err := cmd.NetworkFlag.Device()
	if err != nil {
		return nil, err
	}

	spec.AddDevice(netdev)

	folders, err := cmd.Datacenter.Folders(cmd.Client)
	if err != nil {
		return nil, err
	}

	return folders.VmFolder.CreateVM(cmd.Client, spec.ToSpec(), cmd.ResourcePool, cmd.HostSystem)
}

// Cleanup tries to clean up the specified VM. As it is called from error
// handling paths, it is a best effort function. It attempts to remove disks
// from the VM prior to destroying it, to prevent unintended purging of disks.
func (cmd *create) Cleanup(vm *govmomi.VirtualMachine) {
	var mvm mo.VirtualMachine

	// TODO(PN): Use `config.hardware` here, see issue #44.
	err := cmd.Client.Properties(vm.Reference(), []string{"config"}, &mvm)
	if err != nil {
		return
	}

	spec := new(configSpec)
	spec.RemoveDisks(&mvm)

	task, err := vm.Reconfigure(cmd.Client, spec.ToSpec())
	if err != nil {
		return
	}

	err = task.Wait()
	if err != nil {
		return
	}

	task, err = vm.Destroy(cmd.Client)
	if err != nil {
		return
	}

	err = task.Wait()
	if err != nil {
		return
	}
}

type configSpec types.VirtualMachineConfigSpec

func (c *configSpec) ToSpec() types.VirtualMachineConfigSpec {
	return types.VirtualMachineConfigSpec(*c)
}

func (c *configSpec) AddChange(d types.BaseVirtualDeviceConfigSpec) {
	c.DeviceChange = append(c.DeviceChange, d)
}

func (c *configSpec) AddDevice(d types.BaseVirtualDevice) {
	op := &types.VirtualDeviceConfigSpec{
		Operation: types.VirtualDeviceConfigSpecOperationAdd,
		Device:    d,
	}

	c.AddChange(op)
}

func (c *configSpec) RemoveDisks(vm *mo.VirtualMachine) {
	for _, d := range vm.Config.Hardware.Device {
		switch device := d.(type) {
		case *types.VirtualDisk:
			removeOp := &types.VirtualDeviceConfigSpec{
				Operation: types.VirtualDeviceConfigSpecOperationRemove,
				Device:    device,
			}

			c.AddChange(removeOp)
		}
	}
}
