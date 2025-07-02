// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vapp

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/object"
)

type power struct {
	*flags.SearchFlag

	On      bool
	Off     bool
	Suspend bool
	Force   bool
}

func init() {
	cli.Register("vapp.power", &power{})
}

func (cmd *power) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.SearchFlag, ctx = flags.NewSearchFlag(ctx, flags.SearchVirtualApps)
	cmd.SearchFlag.Register(ctx, f)

	f.BoolVar(&cmd.On, "on", false, "Power on")
	f.BoolVar(&cmd.Off, "off", false, "Power off")
	f.BoolVar(&cmd.Suspend, "suspend", false, "Power suspend")
	f.BoolVar(&cmd.Force, "force", false, "Force (If force is false, the shutdown order in the vApp is executed. If force is true, all virtual machines are powered-off (regardless of shutdown order))")
}

func (cmd *power) Process(ctx context.Context) error {
	if err := cmd.SearchFlag.Process(ctx); err != nil {
		return err
	}
	opts := []bool{cmd.On, cmd.Off, cmd.Suspend}
	selected := false

	for _, opt := range opts {
		if opt {
			if selected {
				return flag.ErrHelp
			}
			selected = opt
		}
	}

	if !selected {
		return flag.ErrHelp
	}

	return nil
}

func (cmd *power) Run(ctx context.Context, f *flag.FlagSet) error {
	vapps, err := cmd.VirtualApps(f.Args())
	if err != nil {
		return err
	}

	for _, vapp := range vapps {
		var task *object.Task

		switch {
		case cmd.On:
			fmt.Fprintf(cmd, "Powering on %s... ", vapp.Reference())
			task, err = vapp.PowerOn(ctx)
		case cmd.Off:
			fmt.Fprintf(cmd, "Powering off %s... ", vapp.Reference())
			task, err = vapp.PowerOff(ctx, cmd.Force)
		case cmd.Suspend:
			fmt.Fprintf(cmd, "Suspend %s... ", vapp.Reference())
			task, err = vapp.Suspend(ctx)
		}

		if err != nil {
			return err
		}

		if task != nil {
			err = task.Wait(ctx)
		}
		if err == nil {
			fmt.Fprintf(cmd, "OK\n")
			continue
		}

		return err
	}

	return nil
}
