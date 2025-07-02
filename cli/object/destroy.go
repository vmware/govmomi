// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package object

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/object"
)

type destroy struct {
	*flags.DatacenterFlag
}

func init() {
	cli.Register("object.destroy", &destroy{})
}

func (cmd *destroy) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatacenterFlag, ctx = flags.NewDatacenterFlag(ctx)
	cmd.DatacenterFlag.Register(ctx, f)
}

func (cmd *destroy) Usage() string {
	return "PATH..."
}

func (cmd *destroy) Description() string {
	return `Destroy managed objects.

Examples:
  govc object.destroy /dc1/network/dvs /dc1/host/cluster`
}

func (cmd *destroy) Process(ctx context.Context) error {
	if err := cmd.DatacenterFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *destroy) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() == 0 {
		return flag.ErrHelp
	}

	c, err := cmd.Client()
	if err != nil {
		return err
	}

	objs, err := cmd.ManagedObjects(ctx, f.Args())
	if err != nil {
		return err
	}

	for _, obj := range objs {
		task, err := object.NewCommon(c, obj).Destroy(ctx)
		if err != nil {
			return err
		}

		logger := cmd.ProgressLogger(fmt.Sprintf("destroying %s... ", obj))
		_, err = task.WaitForResult(ctx, logger)
		logger.Wait()
		if err != nil {
			return err
		}
	}

	return nil
}
