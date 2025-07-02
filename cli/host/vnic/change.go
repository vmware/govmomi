// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vnic

import (
	"context"
	"errors"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vim25/mo"
)

type change struct {
	*flags.HostSystemFlag

	mtu int32
}

func init() {
	cli.Register("host.vnic.change", &change{})
}

func (cmd *change) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.HostSystemFlag, ctx = flags.NewHostSystemFlag(ctx)
	cmd.HostSystemFlag.Register(ctx, f)

	f.Var(flags.NewInt32(&cmd.mtu), "mtu", "vmk MTU")
}

func (cmd *change) Process(ctx context.Context) error {
	if err := cmd.HostSystemFlag.Process(ctx); err != nil {
		return err
	}

	return nil
}

func (cmd *change) Usage() string {
	return "DEVICE"
}

func (cmd *change) Description() string {
	return `Change a virtual nic DEVICE.

Examples:
  govc host.vnic.change -host hostname -mtu 9000 vmk1`
}

func (cmd *change) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}

	device := f.Arg(0)

	ns, err := cmd.HostNetworkSystem()
	if err != nil {
		return err
	}

	var mns mo.HostNetworkSystem

	err = ns.Properties(ctx, ns.Reference(), []string{"networkInfo"}, &mns)
	if err != nil {
		return err
	}

	for _, nic := range mns.NetworkInfo.Vnic {
		if nic.Device == device {
			nic.Spec.Mtu = cmd.mtu
			return ns.UpdateVirtualNic(ctx, device, nic.Spec)
		}
	}

	return errors.New(device + " not found")
}
