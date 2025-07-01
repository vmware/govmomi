// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vapp

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/fault"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/vim25/types"
)

type destroy struct {
	*flags.DatacenterFlag
}

func init() {
	cli.Register("vapp.destroy", &destroy{})
}

func (cmd *destroy) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatacenterFlag, ctx = flags.NewDatacenterFlag(ctx)
	cmd.DatacenterFlag.Register(ctx, f)
}

func (cmd *destroy) Process(ctx context.Context) error {
	if err := cmd.DatacenterFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *destroy) Usage() string {
	return "VAPP..."
}

func (cmd *destroy) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() == 0 {
		return flag.ErrHelp
	}

	finder, err := cmd.Finder()
	if err != nil {
		return err
	}

	for _, arg := range f.Args() {
		vapps, err := finder.VirtualAppList(ctx, arg)
		if err != nil {
			if _, ok := err.(*find.NotFoundError); ok {
				// Ignore if vapp cannot be found
				continue
			}

			return err
		}

		for _, vapp := range vapps {
			powerOff := func() error {
				task, err := vapp.PowerOff(ctx, false)
				if err != nil {
					return err
				}
				err = task.Wait(ctx)
				if err != nil {
					// it's safe to ignore if the vapp is already powered off
					if fault.Is(err, &types.InvalidPowerState{}) {
						return nil
					}
					return err
				}
				return nil
			}
			if err := powerOff(); err != nil {
				return err
			}

			destroy := func() error {
				task, err := vapp.Destroy(ctx)
				if err != nil {
					return err
				}
				err = task.Wait(ctx)
				if err != nil {
					return err
				}
				return nil
			}
			if err := destroy(); err != nil {
				return err
			}
		}
	}

	return nil
}
