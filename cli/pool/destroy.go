// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package pool

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/find"
)

type destroy struct {
	*flags.DatacenterFlag

	children bool
}

func init() {
	cli.Register("pool.destroy", &destroy{})
}

func (cmd *destroy) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatacenterFlag, ctx = flags.NewDatacenterFlag(ctx)
	cmd.DatacenterFlag.Register(ctx, f)

	f.BoolVar(&cmd.children, "children", false, "Remove all children pools")
}

func (cmd *destroy) Process(ctx context.Context) error {
	if err := cmd.DatacenterFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *destroy) Usage() string {
	return "POOL..."
}

func (cmd *destroy) Description() string {
	return "Destroy one or more resource POOLs.\n" + poolNameHelp
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
		pools, err := finder.ResourcePoolList(ctx, arg)
		if err != nil {
			if _, ok := err.(*find.NotFoundError); ok {
				// Ignore if pool cannot be found
				continue
			}

			return err
		}

		for _, pool := range pools {
			if cmd.children {
				err = pool.DestroyChildren(ctx)
				if err != nil {
					return err
				}
			} else {
				task, err := pool.Destroy(ctx)
				if err != nil {
					return err
				}
				err = task.Wait(ctx)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
