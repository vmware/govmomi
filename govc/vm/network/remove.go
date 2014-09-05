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

package network

import (
	"errors"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type remove struct {
	*flags.VirtualMachineFlag
}

func init() {
	cli.Register("vm.network.remove", &remove{})
}

func (cmd *remove) Register(f *flag.FlagSet) {}

func (cmd *remove) Process() error { return nil }

func (cmd *remove) Run(f *flag.FlagSet) error {
	c, err := cmd.ClientFlag.Client()
	if err != nil {
		return err
	}

	vm, err := cmd.VirtualMachineFlag.VirtualMachine()
	if err != nil {
		return err
	}

	if vm == nil {
		return errors.New("please specify a vm")
	}

	name := f.Arg(0)

	if name == "" {
		return errors.New("please specify a network")
	}

	var mvm mo.VirtualMachine

	err = c.Properties(vm.Reference(), []string{"config.hardware"}, &mvm)
	if err != nil {
		return err
	}

	var net *types.VirtualDevice

	for _, dev := range mvm.Config.Hardware.Device {
		if eth, ok := dev.GetVirtualDevice().Backing.(*types.VirtualEthernetCardNetworkBackingInfo); ok {
			if eth.DeviceName == name {
				net = dev.GetVirtualDevice()
				break
			}
		}
	}

	if net == nil {
		return fmt.Errorf("vm network device '%s' not found", name)
	}

	config := &types.VirtualDeviceConfigSpec{
		Device:    net,
		Operation: types.VirtualDeviceConfigSpecOperationRemove,
	}

	spec := types.VirtualMachineConfigSpec{}
	spec.DeviceChange = append(spec.DeviceChange, config)

	task, err := vm.Reconfigure(c, spec)
	if err != nil {
		return err
	}

	return task.Wait()
}
