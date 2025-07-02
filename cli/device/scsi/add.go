// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package scsi

import (
	"context"
	"flag"
	"fmt"
	"strings"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/types"
)

type add struct {
	*flags.VirtualMachineFlag

	controller   string
	sharedBus    string
	hotAddRemove bool
}

func init() {
	cli.Register("device.scsi.add", &add{})
}

func (cmd *add) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)

	var ctypes []string
	ct := object.SCSIControllerTypes()
	for _, t := range ct {
		ctypes = append(ctypes, ct.Type(t))
	}
	f.StringVar(&cmd.controller, "type", ct.Type(ct[0]),
		fmt.Sprintf("SCSI controller type (%s)", strings.Join(ctypes, "|")))
	f.StringVar(&cmd.sharedBus, "sharing", string(types.VirtualSCSISharingNoSharing), "SCSI sharing")
	f.BoolVar(&cmd.hotAddRemove, "hot", false, "Enable hot-add/remove")
}

func (cmd *add) Description() string {
	return `Add SCSI controller to VM.

Examples:
  govc device.scsi.add -vm $vm
  govc device.scsi.add -vm $vm -type pvscsi
  govc device.info -vm $vm {lsi,pv}*`
}

func (cmd *add) Process(ctx context.Context) error {
	if err := cmd.VirtualMachineFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *add) Run(ctx context.Context, f *flag.FlagSet) error {
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

	d, err := devices.CreateSCSIController(cmd.controller)
	if err != nil {
		return err
	}

	c := d.(types.BaseVirtualSCSIController).GetVirtualSCSIController()
	c.HotAddRemove = &cmd.hotAddRemove
	c.SharedBus = types.VirtualSCSISharing(cmd.sharedBus)

	err = vm.AddDevice(ctx, d)
	if err != nil {
		return err
	}

	// output name of device we just created
	devices, err = vm.Device(ctx)
	if err != nil {
		return err
	}

	devices = devices.SelectByType(d)

	name := devices.Name(devices[len(devices)-1])

	fmt.Println(name)

	return nil
}
