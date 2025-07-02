// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package pci

import (
	"context"
	"errors"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vim25/types"
)

type add struct {
	*flags.VirtualMachineFlag
}

func init() {
	cli.Register("device.pci.add", &add{})
}

func (cmd *add) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)
}

func (cmd *add) Description() string {
	return `Add PCI Passthrough device to VM.

Examples:
  govc device.pci.ls -vm $vm
  govc device.pci.add -vm $vm $pci_address
  govc device.info -vm $vm

Assuming vm name is helloworld, list command has below output

$ govc device.pci.ls -vm helloworld
System ID                             Address       Vendor Name Device Name
5b087ce4-ce46-72c0-c7c2-28ac9e22c3c2  0000:60:00.0  Pensando    Ethernet Controller 1
5b087ce4-ce46-72c0-c7c2-28ac9e22c3c2  0000:61:00.0  Pensando    Ethernet Controller 2

To add only 'Ethernet Controller 1', command should be as below. No output upon success.

$ govc device.pci.add -vm helloworld 0000:60:00.0

To add both 'Ethernet Controller 1' and 'Ethernet Controller 2', command should be as below.
No output upon success.

$ govc device.pci.add -vm helloworld 0000:60:00.0 0000:61:00.0

$ govc device.info -vm helloworld
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
  Unit number:      19`
}

func (cmd *add) Usage() string {
	return "PCI_ADDRESS..."
}

func (cmd *add) Run(ctx context.Context, f *flag.FlagSet) error {
	if len(f.Args()) == 0 {
		return flag.ErrHelp
	}

	reqDevices := map[string]*types.VirtualMachinePciPassthroughInfo{}
	for _, n := range f.Args() {
		reqDevices[n] = nil
	}

	vm, err := cmd.VirtualMachine()
	if err != nil {
		return err
	}
	if vm == nil {
		return flag.ErrHelp
	}

	vmConfigOptions, err := queryConfigTarget(ctx, vm)
	if err != nil {
		return err
	}

	for _, d := range vmConfigOptions.PciPassthrough {
		info := d.GetVirtualMachinePciPassthroughInfo()
		if info == nil {
			return errors.New("received invalid pci passthrough info")
		}

		_, ok := reqDevices[info.PciDevice.Id]
		if !ok {
			continue
		}
		reqDevices[info.PciDevice.Id] = info
	}

	newDevices := []types.BaseVirtualDevice{}
	for id, d := range reqDevices {
		if d == nil {
			return fmt.Errorf("%s is not found in allowed PCI passthrough device list", id)
		}
		device := &types.VirtualPCIPassthrough{}
		device.Backing = &types.VirtualPCIPassthroughDeviceBackingInfo{
			Id: d.PciDevice.Id, DeviceId: fmt.Sprintf("%x", d.PciDevice.DeviceId),
			SystemId: d.SystemId, VendorId: d.PciDevice.VendorId,
		}
		device.Connectable = &types.VirtualDeviceConnectInfo{StartConnected: true, Connected: true}
		newDevices = append(newDevices, device)
	}

	return vm.AddDevice(ctx, newDevices...)
}
