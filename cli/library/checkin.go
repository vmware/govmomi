// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package library

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/vcenter"
)

type checkin struct {
	*flags.VirtualMachineFlag
	vcenter.CheckIn
}

func init() {
	cli.Register("library.checkin", &checkin{})
}

func (cmd *checkin) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)

	f.StringVar(&cmd.Message, "m", "", "Check in message")
}

func (cmd *checkin) Usage() string {
	return "PATH"
}

func (cmd *checkin) Description() string {
	return `Check in VM to Content Library item PATH.

Note: this command requires vCenter 7.0 or higher.

Examples:
  govc library.checkin -vm my-vm my-content/template-vm-item`
}

func (cmd *checkin) Run(ctx context.Context, f *flag.FlagSet) error {
	path := f.Arg(0)
	vm, err := cmd.VirtualMachine()
	if err != nil {
		return err
	}
	if vm == nil {
		return flag.ErrHelp
	}

	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	l, err := flags.ContentLibraryItem(ctx, c, path)
	if err != nil {
		return err
	}

	version, err := vcenter.NewManager(c).CheckIn(ctx, l.ID, vm, &cmd.CheckIn)
	if err != nil {
		return err
	}

	fmt.Printf("%s (%s) checked in as content version %s", l.Name, l.ID, version)

	return nil
}
