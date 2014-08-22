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

package datastore

import (
	"errors"
	"flag"
	"fmt"
	"path"
	"strings"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type import_ struct {
	*flags.DatastoreFlag
	*flags.ResourcePoolFlag

	Client       *govmomi.Client
	Datacenter   *govmomi.Datacenter
	Datastore    *govmomi.Datastore
	ResourcePool *govmomi.ResourcePool
}

func init() {
	cli.Register(&import_{})
}

func (cmd *import_) Register(f *flag.FlagSet) {}

func (cmd *import_) Process() error { return nil }

func (cmd *import_) Run(f *flag.FlagSet) error {
	var err error

	args := f.Args()
	if len(args) != 1 {
		return errors.New("no file to import")
	}

	file := importable(f.Arg(0))
	switch file.Ext() {
	case ".vmdk":
	default:
		return fmt.Errorf(`unknown type: %s`, file)
	}

	cmd.Client, err = cmd.DatastoreFlag.Client()
	if err != nil {
		return err
	}

	cmd.Datacenter, err = cmd.DatastoreFlag.Datacenter()
	if err != nil {
		return err
	}

	cmd.Datastore, err = cmd.DatastoreFlag.Datastore()
	if err != nil {
		return err
	}

	cmd.ResourcePool, err = cmd.ResourcePoolFlag.ResourcePool()
	if err != nil {
		return err
	}

	u, err := cmd.DatastoreURL(file.RemoteSrc())
	if err != nil {
		return err
	}

	err = cmd.Client.Client.UploadFile(string(file), u)
	if err != nil {
		return err
	}

	err = cmd.Copy(file)
	if err != nil {
		return err
	}

	fm := cmd.Client.FileManager()
	err = fm.DeleteDatastoreFile(cmd.Datastore.Path(file.RemoteSrc()), cmd.Datacenter)
	if err != nil {
		return err
	}

	return nil
}

func (cmd *import_) Copy(i importable) error {
	var err error

	dstName := path.Dir(i.RemoteDst())
	srcName := dstName + "-src"

	spec := &configSpec{
		Name:    srcName,
		GuestId: "otherGuest",
		Files: &types.VirtualMachineFileInfo{
			VmPathName: fmt.Sprintf("[%s]", cmd.Datastore.Name()),
		},
	}

	spec.AddDisk(cmd.Datastore, i.RemoteSrc())

	src, err := cmd.CreateVM(spec)
	if err != nil {
		return err
	}

	dst, err := cmd.CloneVM(src, dstName)
	if err != nil {
		return err
	}

	err = cmd.DestroyVM(src)
	if err != nil {
		return err
	}

	err = cmd.DestroyVM(dst)
	if err != nil {
		return err
	}

	// TODO(PN): Don't use hardcoded file suffixes.
	// TODO(PN): Support arbitrary destination paths, not just top level.
	for _, s := range []string{"", "-flat"} {
		src := i.RemoteDstWithSuffix(s)
		dst := path.Base(src)
		err = cmd.Move(src, dst)
		if err != nil {
			return err
		}
	}

	// Delete source virtual machine directory.
	err = cmd.Delete(dstName)
	if err != nil {
		return err
	}

	return nil
}

func (cmd *import_) Move(src, dst string) error {
	fm := cmd.Client.FileManager()
	dsSrc := cmd.Datastore.Path(src)
	dsDst := cmd.Datastore.Path(dst)
	return fm.MoveDatastoreFile(dsSrc, cmd.Datacenter, dsDst, cmd.Datacenter, true)
}

func (cmd *import_) Delete(path string) error {
	fm := cmd.Client.FileManager()
	dsPath := cmd.Datastore.Path(path)
	return fm.DeleteDatastoreFile(dsPath, cmd.Datacenter)
}

func (cmd *import_) CreateVM(spec *configSpec) (*govmomi.VirtualMachine, error) {
	folders, err := cmd.Datacenter.Folders(cmd.Client)
	if err != nil {
		return nil, err
	}

	vm, err := folders.VmFolder.CreateVM(cmd.Client, spec.ToSpec(), cmd.ResourcePool, nil)
	if err != nil {
		return nil, err
	}

	return vm, nil
}

func (cmd *import_) CloneVM(vm *govmomi.VirtualMachine, name string) (*govmomi.VirtualMachine, error) {
	folders, err := cmd.Datacenter.Folders(cmd.Client)
	if err != nil {
		return nil, err
	}

	spec := types.VirtualMachineCloneSpec{
		Config:   &types.VirtualMachineConfigSpec{},
		Location: types.VirtualMachineRelocateSpec{},
	}

	return vm.Clone(cmd.Client, folders.VmFolder, name, spec)
}

func (cmd *import_) DestroyVM(vm *govmomi.VirtualMachine) error {
	var mvm mo.VirtualMachine

	// TODO(PN): Use `config.hardware` here, see issue #44.
	err := cmd.Client.Properties(vm.Reference(), []string{"config"}, &mvm)
	if err != nil {
		return err
	}

	spec := new(configSpec)
	spec.RemoveDisks(&mvm)
	err = vm.Reconfigure(cmd.Client, spec.ToSpec())
	if err != nil {
		return err
	}

	// Best effort
	_ = vm.Destroy(cmd.Client)
	return nil
}

type importable string

func (i importable) Ext() string {
	return strings.ToLower(path.Ext(string(i)))
}

func (i importable) Base() string {
	return path.Base(string(i))
}

func (i importable) BaseClean() string {
	b := path.Base(string(i))
	e := path.Ext(string(i))
	return b[:len(b)-len(e)]
}

func (i importable) RemoteSrc() string {
	bc := i.BaseClean()
	return fmt.Sprintf(".%s.vmdk", bc)
}

func (i importable) RemoteDst() string {
	bc := i.BaseClean()
	return fmt.Sprintf("%s/%s.vmdk", bc, bc)
}

func (i importable) RemoteDstWithSuffix(s string) string {
	bc := i.BaseClean()
	return fmt.Sprintf("%s/%s%s.vmdk", bc, bc, s)
}

type configSpec types.VirtualMachineConfigSpec

func (c *configSpec) ToSpec() types.VirtualMachineConfigSpec {
	return types.VirtualMachineConfigSpec(*c)
}

func (c *configSpec) AddChange(d types.BaseVirtualDeviceConfigSpec) {
	c.DeviceChange = append(c.DeviceChange, d)
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

	controllerSpec := &types.VirtualDeviceConfigSpec{
		Device:    controller,
		Operation: types.VirtualDeviceConfigSpecOperationAdd,
	}

	c.AddChange(controllerSpec)

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

	diskSpec := &types.VirtualDeviceConfigSpec{
		Device:    disk,
		Operation: types.VirtualDeviceConfigSpecOperationAdd,
	}

	c.AddChange(diskSpec)
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
