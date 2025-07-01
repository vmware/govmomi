// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package pool

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vim25/types"
)

type change struct {
	*flags.DatacenterFlag
	*ResourceConfigSpecFlag

	name string
}

func init() {
	cli.Register("pool.change", &change{})
}

func (cmd *change) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatacenterFlag, ctx = flags.NewDatacenterFlag(ctx)
	cmd.DatacenterFlag.Register(ctx, f)
	cmd.ResourceConfigSpecFlag = &ResourceConfigSpecFlag{
		ResourceConfigSpec: types.ResourceConfigSpec{
			CpuAllocation: types.ResourceAllocationInfo{
				Shares: new(types.SharesInfo),
			},
			MemoryAllocation: types.ResourceAllocationInfo{
				Shares: new(types.SharesInfo),
			},
		},
	}
	cmd.ResourceConfigSpecFlag.Register(ctx, f)

	f.StringVar(&cmd.name, "name", "", "Resource pool name")
}

func (cmd *change) Process(ctx context.Context) error {
	if err := cmd.DatacenterFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.ResourceConfigSpecFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *change) Usage() string {
	return "POOL..."
}

func (cmd *change) Description() string {
	return "Change the configuration of one or more resource POOLs.\n" + poolNameHelp
}

func (cmd *change) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() == 0 {
		return flag.ErrHelp
	}

	finder, err := cmd.Finder()
	if err != nil {
		return err
	}

	for _, ra := range []*types.ResourceAllocationInfo{&cmd.CpuAllocation, &cmd.MemoryAllocation} {
		if ra.Shares.Level == "" {
			ra.Shares = nil
		}
	}

	for _, arg := range f.Args() {
		pools, err := finder.ResourcePoolListAll(ctx, arg)
		if err != nil {
			return err
		}

		for _, pool := range pools {
			err := pool.UpdateConfig(ctx, cmd.name, &cmd.ResourceConfigSpec)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
