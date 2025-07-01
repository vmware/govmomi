// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vnic

import (
	"context"
	"flag"
	"fmt"
	"strings"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vim25/types"
)

type service struct {
	*flags.HostSystemFlag

	Enable bool
}

func init() {
	cli.Register("host.vnic.service", &service{})
}

func (cmd *service) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.HostSystemFlag, ctx = flags.NewHostSystemFlag(ctx)
	cmd.HostSystemFlag.Register(ctx, f)

	f.BoolVar(&cmd.Enable, "enable", true, "Enable service")
}

func (cmd *service) Process(ctx context.Context) error {
	if err := cmd.HostSystemFlag.Process(ctx); err != nil {
		return err
	}

	return nil
}

func (cmd *service) Usage() string {
	return "SERVICE DEVICE"
}

func (cmd *service) Description() string {
	nicTypes := types.HostVirtualNicManagerNicType("").Strings()

	return fmt.Sprintf(`
Enable or disable service on a virtual nic device.

Where SERVICE is one of: %s
Where DEVICE is one of: %s

Examples:
  govc host.vnic.service -host hostname -enable vsan vmk0
  govc host.vnic.service -host hostname -enable=false vmotion vmk1`,
		strings.Join(nicTypes, "|"),
		strings.Join([]string{"vmk0", "vmk1", "..."}, "|"))
}

func (cmd *service) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 2 {
		return flag.ErrHelp
	}

	service := f.Arg(0)
	device := f.Arg(1)

	host, err := cmd.HostSystem()
	if err != nil {
		return err
	}

	m, err := host.ConfigManager().VirtualNicManager(ctx)
	if err != nil {
		return err
	}

	var method func(context.Context, string, string) error

	if cmd.Enable {
		method = m.SelectVnic
	} else {
		method = m.DeselectVnic
	}

	if method == nil {
		return flag.ErrHelp
	}

	return method(ctx, service, device)
}
