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
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

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
