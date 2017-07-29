/*
Copyright (c) 2017 VMware, Inc. All Rights Reserved.

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

package object

import (
	"context"
	"fmt"
	"path"
	"reflect"
	"regexp"

	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

// VirtualDisk vmdk
type VirtualDisk struct {
	Client       *vim25.Client
	Datacenter   *Datacenter
	Datastore    *Datastore
	ResourcePool *ResourcePool
}

// NewVmdk Object
func NewVmdk(c *vim25.Client, ds *Datastore, dc *Datacenter, pool *ResourcePool) (*VirtualDisk, error) {
	return &VirtualDisk{
		Client:       c,
		Datacenter:   dc,
		Datastore:    ds,
		ResourcePool: pool,
	}, nil
}

// Import Import vmdk to the path given
func (v *VirtualDisk) Import(remotePath string) error {
	err := v.checkVMDKExist(remotePath)
	if err != nil {
		return err
	}

	err = v.copy(remotePath)
	if err != nil {
		// clean up
		return err
	}

	return nil
}

func (v *VirtualDisk) copy(remotePath string) error {
	// src Folder and dst Folder for store temp VM created
	srcName := remotePath[:len(path.Base(remotePath))-len(path.Ext(remotePath))] + "-srcvm"
	dstName := remotePath[:len(path.Base(remotePath))-len(path.Ext(remotePath))] + "-dstvm"

	spec := &configSpec{
		Name:    srcName,
		GuestId: "otherGuest",
		Files: &types.VirtualMachineFileInfo{
			VmPathName: fmt.Sprintf("[%s]", v.Datastore.Name()),
		},
	}

	spec.addDisk(v.Datastore, remotePath)

	src, err := v.createVM(spec)
	if err != nil {
		return err
	}

	dst, err := v.cloneVM(src, dstName)
	if err != nil {
		return err
	}

	err = v.destroyVM(src)
	if err != nil {
		return err
	}

	vmdk, err := v.detachDisk(dst)
	if err != nil {
		return err
	}

	// remove the old vmdk
	err = v.deleteDisk(remotePath)
	if err != nil {
		return err
	}

	err = v.moveDisk(vmdk, remotePath)
	if err != nil {
		return err
	}

	err = v.destroyVM(dst)
	if err != nil {
		return err
	}

	return nil
}

func (v *VirtualDisk) checkVMDKExist(remotePath string) error {
	ctx := context.TODO()

	res, err := v.Datastore.Stat(ctx, remotePath)

	if err != nil {
		switch err.(type) {
		case DatastoreNoSuchDirectoryError:
			// The base path doesn't exist. Create it.
			dsPath := v.Datastore.Path(path.Dir(remotePath))
			m := NewFileManager(v.Client)
			return m.MakeDirectory(ctx, dsPath, v.Datacenter, true)
		case DatastoreNoSuchFileError:
			// Destination path doesn't exist; all good to continue with import.
			return nil
		}

		return err
	}

	// Check that the returned entry has the right type.
	switch res.(type) {
	case *types.VmDiskFileInfo:
	default:
		expected := "VmDiskFileInfo"
		actual := reflect.TypeOf(res)
		return fmt.Errorf("Expected: %s, actual: %s", expected, actual)
	}

	return nil
}

func (v *VirtualDisk) moveDisk(src, dst string) error {
	ctx := context.TODO()
	dsSrc := v.Datastore.Path(src)
	dsDst := v.Datastore.Path(dst)
	vdm := NewVirtualDiskManager(v.Client)
	task, err := vdm.MoveVirtualDisk(ctx, dsSrc, v.Datacenter, dsDst, v.Datacenter, true)
	if err != nil {
		return err
	}

	return task.Wait(ctx)
}

func (v *VirtualDisk) deleteDisk(path string) error {
	ctx := context.TODO()
	vdm := NewVirtualDiskManager(v.Client)
	task, err := vdm.DeleteVirtualDisk(ctx, v.Datastore.Path(path), v.Datacenter)
	if err != nil {
		return err
	}

	return task.Wait(ctx)
}

func (v *VirtualDisk) detachDisk(vm *VirtualMachine) (string, error) {
	ctx := context.TODO()
	var mvm mo.VirtualMachine

	pc := property.DefaultCollector(v.Client)
	err := pc.RetrieveOne(ctx, vm.Reference(), []string{"config.hardware"}, &mvm)
	if err != nil {
		return "", err
	}

	spec := new(configSpec)
	dsFile := spec.removeDisk(&mvm)

	task, err := vm.Reconfigure(ctx, spec.toSpec())
	if err != nil {
		return "", err
	}

	err = task.Wait(ctx)
	if err != nil {
		return "", err
	}

	return dsFile, nil
}

func (v *VirtualDisk) createVM(spec *configSpec) (*VirtualMachine, error) {
	ctx := context.TODO()
	folders, err := v.Datacenter.Folders(ctx)
	if err != nil {
		return nil, err
	}

	task, err := folders.VmFolder.CreateVM(ctx, spec.toSpec(), v.ResourcePool, nil)
	if err != nil {
		return nil, err
	}

	info, err := task.WaitForResult(ctx, nil)
	if err != nil {
		return nil, err
	}

	return NewVirtualMachine(v.Client, info.Result.(types.ManagedObjectReference)), nil
}

func (v *VirtualDisk) cloneVM(vm *VirtualMachine, name string) (*VirtualMachine, error) {
	ctx := context.TODO()
	folders, err := v.Datacenter.Folders(ctx)
	if err != nil {
		return nil, err
	}

	spec := types.VirtualMachineCloneSpec{
		Config:   &types.VirtualMachineConfigSpec{},
		Location: types.VirtualMachineRelocateSpec{},
	}

	task, err := vm.Clone(ctx, folders.VmFolder, name, spec)
	if err != nil {
		return nil, err
	}

	info, err := task.WaitForResult(ctx, nil)
	if err != nil {
		return nil, err
	}

	return NewVirtualMachine(v.Client, info.Result.(types.ManagedObjectReference)), nil
}

func (v *VirtualDisk) destroyVM(vm *VirtualMachine) error {
	ctx := context.TODO()
	_, err := v.detachDisk(vm)
	if err != nil {
		return err
	}

	task, err := vm.Destroy(ctx)
	if err != nil {
		return err
	}

	err = task.Wait(ctx)
	if err != nil {
		return err
	}

	return nil
}

type configSpec types.VirtualMachineConfigSpec

func (c *configSpec) toSpec() types.VirtualMachineConfigSpec {
	return types.VirtualMachineConfigSpec(*c)
}

func (c *configSpec) addChange(d types.BaseVirtualDeviceConfigSpec) {
	c.DeviceChange = append(c.DeviceChange, d)
}

func (c *configSpec) addDisk(ds *Datastore, path string) {
	var devices VirtualDeviceList

	controller, err := devices.CreateSCSIController("")
	if err != nil {
		panic(err)
	}
	devices = append(devices, controller)

	disk := devices.CreateDisk(controller.(types.BaseVirtualController), ds.Reference(), ds.Path(path))
	devices = append(devices, disk)

	spec, err := devices.ConfigSpec(types.VirtualDeviceConfigSpecOperationAdd)
	if err != nil {
		panic(err)
	}

	c.DeviceChange = append(c.DeviceChange, spec...)
}

var dsPathRegexp = regexp.MustCompile(`^\[.*\] (.*)$`)

func (c *configSpec) removeDisk(vm *mo.VirtualMachine) string {
	var file string

	for _, d := range vm.Config.Hardware.Device {
		switch device := d.(type) {
		case *types.VirtualDisk:
			if file != "" {
				panic("expected VM to have only one disk")
			}

			switch backing := device.Backing.(type) {
			case *types.VirtualDiskFlatVer1BackingInfo:
				file = backing.FileName
			case *types.VirtualDiskFlatVer2BackingInfo:
				file = backing.FileName
			case *types.VirtualDiskSeSparseBackingInfo:
				file = backing.FileName
			case *types.VirtualDiskSparseVer1BackingInfo:
				file = backing.FileName
			case *types.VirtualDiskSparseVer2BackingInfo:
				file = backing.FileName
			default:
				name := reflect.TypeOf(device.Backing).String()
				panic(fmt.Sprintf("unexpected backing type: %s", name))
			}

			// Remove [datastore] prefix
			m := dsPathRegexp.FindStringSubmatch(file)
			if len(m) != 2 {
				panic(fmt.Sprintf("expected regexp match for %#v", file))
			}
			file = m[1]

			removeOp := &types.VirtualDeviceConfigSpec{
				Operation: types.VirtualDeviceConfigSpecOperationRemove,
				Device:    device,
			}

			c.addChange(removeOp)
		}
	}

	return file
}
