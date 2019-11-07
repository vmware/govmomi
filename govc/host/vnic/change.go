/*
Copyright (c) 2015 VMware, Inc. All Rights Reserved.

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

package vnic

import (
	"context"
	"errors"
	"flag"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
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
