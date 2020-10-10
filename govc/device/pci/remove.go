/*
Copyright (c) 2014-2020 VMware, Inc. All Rights Reserved.

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

package pci

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vim25/types"
)

type remove struct {
	*flags.VirtualMachineFlag
}

func init() {
	cli.Register("device.pci.remove", &remove{})
}

func (cmd *remove) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)
}

func (cmd *remove) Usage() string {
	return "<PCI ADDRESS>..."
}

func (cmd *remove) Description() string {
	return `Remove PCI Passthrough device from VM.

Examples:
  govc device.info -vm $vm
  govc device.pci.remove -vm $vm {pci address}*
  govc device.info -vm $vm

Assuming vm name is helloworld, device info command has below output

$ govc device.info --vm helloworld
...
Name:               pcipassthrough-13000
  Type:             VirtualPCIPassthrough
  Label:            PCI device 0
  Summary:
  Key:              13000
  Controller:       pci-100
  Unit number:      18
Name:               pcipassthrough-13001
  Type:             VirtualPCIPassthrough
  Label:            PCI device 1
  Summary:
  Key:              13001
  Controller:       pci-100
  Unit number:      19

To remove only 'pcipassthrough-13000', command should be as below. No output upon success.

$ govc device.pci.remove --vm helloworld pcipassthrough-13000

To remove both 'pcipassthrough-13000' and 'pcipassthrough-13001', command should be as below.
No output upon success.

$ govc device.pci.remove --vm helloworld pcipassthrough-13000 pcipassthrough-13001
`
}

func (cmd *remove) Run(ctx context.Context, f *flag.FlagSet) error {
	if len(f.Args()) == 0 {
		return flag.ErrHelp
	}

	reqDevices := map[string]bool{}
	for _, n := range f.Args() {
		reqDevices[n] = false
	}

	vm, err := cmd.VirtualMachine()
	if err != nil {
		return err
	}
	if vm == nil {
		return flag.ErrHelp
	}

	vmDevices, err := vm.Device(ctx)
	if err != nil {
		return err
	}

	rmDevices := []types.BaseVirtualDevice{}
	for _, d := range vmDevices.SelectByType(&types.VirtualPCIPassthrough{}) {
		name := vmDevices.Name(d)
		_, ok := reqDevices[name]
		if !ok {
			continue
		}
		reqDevices[name] = true
		rmDevices = append(rmDevices, d)
	}

	for id, found := range reqDevices {
		if !found {
			return fmt.Errorf("%s is not found, please check and try again", id)
		}
	}
	return vm.RemoveDevice(ctx, false, rmDevices...)
}
