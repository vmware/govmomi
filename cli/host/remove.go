// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package host

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/mo"
)

type remove struct {
	*flags.HostSystemFlag
}

func init() {
	cli.Register("host.remove", &remove{})
}

func (cmd *remove) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.HostSystemFlag, ctx = flags.NewHostSystemFlag(ctx)
	cmd.HostSystemFlag.Register(ctx, f)
}

func (cmd *remove) Process(ctx context.Context) error {
	if err := cmd.HostSystemFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *remove) Usage() string {
	return "HOST..."
}

func (cmd *remove) Description() string {
	return `Remove HOST from vCenter.`
}

func (cmd *remove) Remove(ctx context.Context, host *object.HostSystem) error {
	var h mo.HostSystem
	err := host.Properties(ctx, host.Reference(), []string{"parent"}, &h)
	if err != nil {
		return err
	}

	remove := host.Destroy

	if h.Parent.Type == "ComputeResource" {
		// Standalone host.  From the docs:
		// "Invoking remove on a HostSystem of standalone type throws a NotSupported fault.
		//  A standalone HostSystem can be removeed only by invoking remove on its parent ComputeResource."
		remove = object.NewComputeResource(host.Client(), *h.Parent).Destroy
	}

	task, err := remove(ctx)
	if err != nil {
		return err
	}

	logger := cmd.ProgressLogger(fmt.Sprintf("%s removing... ", host.InventoryPath))
	defer logger.Wait()

	_, err = task.WaitForResult(ctx, logger)
	return err
}

func (cmd *remove) Run(ctx context.Context, f *flag.FlagSet) error {
	hosts, err := cmd.HostSystems(f.Args())
	if err != nil {
		return err
	}

	for _, host := range hosts {
		err = cmd.Remove(ctx, host)
		if err != nil {
			return err
		}
	}

	return nil
}
