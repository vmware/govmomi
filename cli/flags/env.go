// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package flags

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/mo"
)

type EnvBrowser struct {
	*ClusterFlag
	*HostSystemFlag
	*VirtualMachineFlag
}

func (cmd *EnvBrowser) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClusterFlag, ctx = NewClusterFlag(ctx)
	cmd.ClusterFlag.Register(ctx, f)

	cmd.HostSystemFlag, ctx = NewHostSystemFlag(ctx)
	cmd.HostSystemFlag.Register(ctx, f)

	cmd.VirtualMachineFlag, ctx = NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)
}

func (cmd *EnvBrowser) Process(ctx context.Context) error {
	if err := cmd.ClusterFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.HostSystemFlag.Process(ctx); err != nil {
		return err
	}
	return cmd.VirtualMachineFlag.Process(ctx)
}

func (cmd *EnvBrowser) Browser(ctx context.Context) (*object.EnvironmentBrowser, error) {
	c, err := cmd.VirtualMachineFlag.Client()
	if err != nil {
		return nil, err
	}

	vm, err := cmd.VirtualMachine()
	if err != nil {
		return nil, err
	}
	if vm != nil {
		return vm.EnvironmentBrowser(ctx)
	}

	host, err := cmd.HostSystemIfSpecified()
	if err != nil {
		return nil, err
	}

	if host != nil {
		var h mo.HostSystem
		err = host.Properties(ctx, host.Reference(), []string{"parent"}, &h)
		if err != nil {
			return nil, err
		}

		return object.NewComputeResource(c, *h.Parent).EnvironmentBrowser(ctx)
	}

	finder, ferr := cmd.ClusterFlag.Finder()
	if ferr != nil {
		return nil, ferr
	}

	cr, ferr := finder.ComputeResourceOrDefault(ctx, cmd.ClusterFlag.Name)
	if ferr != nil {
		return nil, ferr
	}

	return cr.EnvironmentBrowser(ctx)
}
