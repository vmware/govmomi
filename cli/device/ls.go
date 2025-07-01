// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package device

import (
	"context"
	"flag"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/object"
)

type ls struct {
	*flags.VirtualMachineFlag

	boot bool
}

func init() {
	cli.Register("device.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)

	f.BoolVar(&cmd.boot, "boot", false, "List devices configured in the VM's boot options")
}

func (cmd *ls) Process(ctx context.Context) error {
	if err := cmd.VirtualMachineFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *ls) Description() string {
	return `List devices for VM.

Examples:
  govc device.ls -vm $name
  govc device.ls -vm $name disk-*
  govc device.ls -vm $name -json | jq '.devices[].name'`
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {
	vm, err := cmd.VirtualMachine()
	if err != nil {
		return err
	}

	if vm == nil {
		return flag.ErrHelp
	}

	devices, err := vm.Device(ctx)
	if err != nil {
		return err
	}

	if f.NArg() != 0 {
		var matches object.VirtualDeviceList
		for _, name := range f.Args() {
			device := match(name, devices)
			if len(device) == 0 {
				return fmt.Errorf("device '%s' not found", name)
			}
			matches = append(matches, device...)
		}
		devices = matches
	}

	if cmd.boot {
		options, err := vm.BootOptions(ctx)
		if err != nil {
			return err
		}

		devices = devices.SelectBootOrder(options.BootOrder)
	}

	res := lsResult{toLsList(devices), devices}
	return cmd.WriteResult(&res)
}

type lsDevice struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Summary string `json:"summary"`
}

func toLsList(devices object.VirtualDeviceList) []lsDevice {
	var res []lsDevice

	for _, device := range devices {
		res = append(res, lsDevice{
			Name:    devices.Name(device),
			Type:    devices.TypeName(device),
			Summary: device.GetVirtualDevice().DeviceInfo.GetDescription().Summary,
		})
	}

	return res
}

type lsResult struct {
	Devices []lsDevice `json:"devices"`
	list    object.VirtualDeviceList
}

func (r *lsResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 3, 0, 2, ' ', 0)

	for _, device := range r.list {
		fmt.Fprintf(tw, "%s\t%s\t%s\n", r.list.Name(device), r.list.TypeName(device),
			device.GetVirtualDevice().DeviceInfo.GetDescription().Summary)
	}

	return tw.Flush()
}
