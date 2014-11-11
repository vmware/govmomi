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
)

type change struct {
	*flags.VirtualMachineFlag
}

func init() {
	cli.Register("vm.network.change", &change{})
}

func (cmd *change) Register(f *flag.FlagSet) {}

func (cmd *change) Process() error { return nil }

func (cmd *change) Run(f *flag.FlagSet) error {
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

	finder, err := cmd.Finder()
	if err != nil {
		return err
	}

	network, err := finder.Network(name)
	if err != nil {
		return err
	}

	backing, err := network.EthernetCardBackingInfo()
	if err != nil {
		return err
	}

	devices, err := vm.Device()
	if err != nil {
		return err
	}

	net := devices.FindByBackingInfo(backing)

	if net == nil {
		return fmt.Errorf("vm network device '%s' not found", name)
	}

	network, err = finder.Network(f.Arg(1))
	if err != nil {
		return err
	}

	backing, err = network.EthernetCardBackingInfo()
	if err != nil {
		return err
	}

	net.GetVirtualDevice().Backing = backing

	return vm.EditDevice(net)
}
