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
	"reflect"

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
		if cmd.ResourcePool, err = cmd.HostSystemFlag.HostResourcePool(); err != nil {
			return err
		}
	}

	vm, err := cmd.CreateVM(f.Arg(0))
	if err != nil {
		return err
	}

	if cmd.DiskFlag.IsSet() && cmd.link {
		err = cmd.Link(vm)
		if err != nil {
			// Unable to link the disk. Make an attempt to clean up the VM, so that
			// it doesn't hold a reference to this disk.
			cmd.Cleanup(vm)
			return err
		}
	}

	if cmd.on {
		err = vm.PowerOn(cmd.Client)
		if err != nil {
			return err
		}
	}

	return nil
}

func (cmd *create) CreateVM(name string) (*govmomi.VirtualMachine, error) {
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

		spec.AddDevice(disk)
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

func (cmd *create) Link(vm *govmomi.VirtualMachine) error {
	var mvm mo.VirtualMachine

	// TODO(PN): Use `config.hardware` here, see issue #44.
	err := cmd.Client.Properties(vm.Reference(), []string{"config"}, &mvm)
	if err != nil {
		return err
	}

	spec := new(configSpec)

	for _, d := range mvm.Config.Hardware.Device {
		switch device := d.(type) {
		case *types.VirtualDisk:
			var addBacking types.BaseVirtualDeviceBackingInfo

			switch b := device.Backing.(type) {
			case *types.VirtualDiskFlatVer2BackingInfo:
				bcopy := *b // Make copy before modifying it
				bcopy.Parent = b
				bcopy.FileName = fmt.Sprintf("[%s]", cmd.Datastore.Name())
				addBacking = &bcopy
			case *types.VirtualDiskSparseVer2BackingInfo:
				bcopy := *b // Make copy before modifying it
				bcopy.Parent = b
				bcopy.FileName = fmt.Sprintf("[%s]", cmd.Datastore.Name())
				addBacking = &bcopy
			default:
				panic("backing not implemented: " + reflect.TypeOf(device.Backing).String())
			}

			removeDevice := *device
			removeOp := &types.VirtualDeviceConfigSpec{
				Operation: types.VirtualDeviceConfigSpecOperationRemove,
				Device:    &removeDevice,
			}

			spec.AddChange(removeOp)

			addDevice := *device
			addDevice.Backing = addBacking
			addDevice.UnitNumber = -1
			addOp := &types.VirtualDeviceConfigSpec{
				Operation:     types.VirtualDeviceConfigSpecOperationAdd,
				FileOperation: types.VirtualDeviceConfigSpecFileOperationCreate,
				Device:        &addDevice,
			}

			spec.AddChange(addOp)
		}
	}

	return vm.Reconfigure(cmd.Client, spec.ToSpec())
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
	err = vm.Reconfigure(cmd.Client, spec.ToSpec())
	if err != nil {
		return
	}

	err = vm.Destroy(cmd.Client)
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

func (c *configSpec) AddDisk(ds *govmomi.Datastore, path string) {
	controller := &types.VirtualLsiLogicController{
		types.VirtualSCSIController{
			SharedBus: types.VirtualSCSISharingNoSharing,
			VirtualController: types.VirtualController{
				BusNumber: 0,
				VirtualDevice: types.VirtualDevice{
					Key: -1,
				},
			},
		},
	}

	c.AddDevice(controller)

	disk := &types.VirtualDisk{
		VirtualDevice: types.VirtualDevice{
			Key:           -1,
			ControllerKey: -1,
			UnitNumber:    -1,
			Backing: &types.VirtualDiskFlatVer2BackingInfo{
				VirtualDeviceFileBackingInfo: types.VirtualDeviceFileBackingInfo{
					FileName: ds.Path(path),
				},
				DiskMode:        string(types.VirtualDiskModePersistent),
				ThinProvisioned: true,
			},
		},
	}

	c.AddDevice(disk)
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
