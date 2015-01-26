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

package govmomi

import (
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type VirtualMachine struct {
	types.ManagedObjectReference

	InventoryPath string

	c *Client
}

func NewVirtualMachine(c *Client, ref types.ManagedObjectReference) *VirtualMachine {
	return &VirtualMachine{
		ManagedObjectReference: ref,
		c: c,
	}
}

func (v VirtualMachine) Reference() types.ManagedObjectReference {
	return v.ManagedObjectReference
}

func (v VirtualMachine) PowerOn() (*Task, error) {
	req := types.PowerOnVM_Task{
		This: v.Reference(),
	}

	res, err := methods.PowerOnVM_Task(v.c, &req)
	if err != nil {
		return nil, err
	}

	return NewTask(v.c, res.Returnval), nil
}

func (v VirtualMachine) PowerOff() (*Task, error) {
	req := types.PowerOffVM_Task{
		This: v.Reference(),
	}

	res, err := methods.PowerOffVM_Task(v.c, &req)
	if err != nil {
		return nil, err
	}

	return NewTask(v.c, res.Returnval), nil
}

func (v VirtualMachine) Reset() (*Task, error) {
	req := types.ResetVM_Task{
		This: v.Reference(),
	}

	res, err := methods.ResetVM_Task(v.c, &req)
	if err != nil {
		return nil, err
	}

	return NewTask(v.c, res.Returnval), nil
}

func (v VirtualMachine) Suspend() (*Task, error) {
	req := types.SuspendVM_Task{
		This: v.Reference(),
	}

	res, err := methods.SuspendVM_Task(v.c, &req)
	if err != nil {
		return nil, err
	}

	return NewTask(v.c, res.Returnval), nil
}

func (v VirtualMachine) ShutdownGuest() error {
	req := types.ShutdownGuest{
		This: v.Reference(),
	}

	_, err := methods.ShutdownGuest(v.c, &req)
	return err
}

func (v VirtualMachine) RebootGuest() error {
	req := types.RebootGuest{
		This: v.Reference(),
	}

	_, err := methods.RebootGuest(v.c, &req)
	return err
}

func (v VirtualMachine) Destroy() (*Task, error) {
	req := types.Destroy_Task{
		This: v.Reference(),
	}

	res, err := methods.Destroy_Task(v.c, &req)
	if err != nil {
		return nil, err
	}

	return NewTask(v.c, res.Returnval), nil
}

func (v VirtualMachine) Clone(folder *Folder, name string, config types.VirtualMachineCloneSpec) (*Task, error) {
	req := types.CloneVM_Task{
		This:   v.Reference(),
		Folder: folder.Reference(),
		Name:   name,
		Spec:   config,
	}

	res, err := methods.CloneVM_Task(v.c, &req)
	if err != nil {
		return nil, err
	}

	return NewTask(v.c, res.Returnval), nil
}

func (v VirtualMachine) Reconfigure(config types.VirtualMachineConfigSpec) (*Task, error) {
	req := types.ReconfigVM_Task{
		This: v.Reference(),
		Spec: config,
	}

	res, err := methods.ReconfigVM_Task(v.c, &req)
	if err != nil {
		return nil, err
	}

	return NewTask(v.c, res.Returnval), nil
}

func (v VirtualMachine) WaitForIP() (string, error) {
	var ip string

	err := v.c.WaitForProperties(v.Reference(), []string{"guest.ipAddress"}, func(pc []types.PropertyChange) bool {
		for _, c := range pc {
			if c.Name != "guest.ipAddress" {
				continue
			}
			if c.Op != types.PropertyChangeOpAssign {
				continue
			}
			if c.Val == nil {
				continue
			}

			ip = c.Val.(string)
			return true
		}

		return false
	})

	if err != nil {
		return "", err
	}

	return ip, nil
}

// Device returns the VirtualMachine's config.hardware.device property.
func (v VirtualMachine) Device() (VirtualDeviceList, error) {
	var o mo.VirtualMachine

	err := v.c.Properties(v.Reference(), []string{"config.hardware.device"}, &o)
	if err != nil {
		return nil, err
	}

	return VirtualDeviceList(o.Config.Hardware.Device), nil
}

func (v VirtualMachine) configureDevice(op types.VirtualDeviceConfigSpecOperation, fop types.VirtualDeviceConfigSpecFileOperation, devices ...types.BaseVirtualDevice) error {
	spec := types.VirtualMachineConfigSpec{}

	for _, device := range devices {
		config := &types.VirtualDeviceConfigSpec{
			Device:    device,
			Operation: op,
		}

		if disk, ok := device.(*types.VirtualDisk); ok {
			config.FileOperation = fop

			// Special case to attach an existing disk
			if op == types.VirtualDeviceConfigSpecOperationAdd && disk.CapacityInKB == 0 {
				childDisk := false
				if b, ok := disk.Backing.(*types.VirtualDiskFlatVer2BackingInfo); ok {
					childDisk = b.Parent != nil
				}

				if !childDisk {
					config.FileOperation = "" // existing disk
				}
			}
		}

		spec.DeviceChange = append(spec.DeviceChange, config)
	}

	task, err := v.Reconfigure(spec)
	if err != nil {
		return err
	}

	return task.Wait()
}

// AddDevice adds the given devices to the VirtualMachine
func (v VirtualMachine) AddDevice(device ...types.BaseVirtualDevice) error {
	return v.configureDevice(types.VirtualDeviceConfigSpecOperationAdd, types.VirtualDeviceConfigSpecFileOperationCreate, device...)
}

// EditDevice edits the given (existing) devices on the VirtualMachine
func (v VirtualMachine) EditDevice(device ...types.BaseVirtualDevice) error {
	return v.configureDevice(types.VirtualDeviceConfigSpecOperationEdit, types.VirtualDeviceConfigSpecFileOperationReplace, device...)
}

// RemoveDevice removes the given devices on the VirtualMachine
func (v VirtualMachine) RemoveDevice(device ...types.BaseVirtualDevice) error {
	return v.configureDevice(types.VirtualDeviceConfigSpecOperationRemove, types.VirtualDeviceConfigSpecFileOperationDestroy, device...)
}

// BootOptions returns the VirtualMachine's config.bootOptions property.
func (v VirtualMachine) BootOptions() (*types.VirtualMachineBootOptions, error) {
	var o mo.VirtualMachine

	err := v.c.Properties(v.Reference(), []string{"config.bootOptions"}, &o)
	if err != nil {
		return nil, err
	}

	return o.Config.BootOptions, nil
}

// SetBootOptions reconfigures the VirtualMachine with the given options.
func (v VirtualMachine) SetBootOptions(options *types.VirtualMachineBootOptions) error {
	spec := types.VirtualMachineConfigSpec{}

	spec.BootOptions = options

	task, err := v.Reconfigure(spec)
	if err != nil {
		return err
	}

	return task.Wait()
}

// Answer answers a pending question.
func (v VirtualMachine) Answer(id, answer string) error {
	req := types.AnswerVM{
		This:         v.Reference(),
		QuestionId:   id,
		AnswerChoice: answer,
	}

	_, err := methods.AnswerVM(v.c, &req)
	if err != nil {
		return err
	}

	return nil
}

func (v VirtualMachine) MarkAsTemplate() error {
	req := types.MarkAsTemplate{
		This: v.Reference(),
	}

	_, err := methods.MarkAsTemplate(v.c, &req)
	if err != nil {
		return err
	}

	return nil
}

func (v VirtualMachine) MarkAsVirtualMachine(pool ResourcePool, host *HostSystem) error {
	req := types.MarkAsVirtualMachine{
		This: v.Reference(),
		Pool: pool.Reference(),
	}

	if host != nil {
		ref := host.Reference()
		req.Host = &ref
	}

	_, err := methods.MarkAsVirtualMachine(v.c, &req)
	if err != nil {
		return err
	}

	return nil
}
