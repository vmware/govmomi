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

type rename struct {
	*flags.DatacenterFlag
}

func init() {
	cli.Register("object.rename", &rename{})
}

func (cmd *rename) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatacenterFlag, ctx = flags.NewDatacenterFlag(ctx)
	cmd.DatacenterFlag.Register(ctx, f)
}

func (cmd *rename) Usage() string {
	return "PATH NAME"
}

func (cmd *rename) Description() string {
	return `Rename managed objects.

Examples:
  govc object.rename /dc1/network/dvs1 Switch1`
}

func (cmd *rename) Process(ctx context.Context) error {
	if err := cmd.DatacenterFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *rename) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 2 {
		return flag.ErrHelp
	}

	c, err := cmd.Client()
	if err != nil {
		return err
	}

	objs, err := cmd.ManagedObjects(ctx, f.Args()[:1])
	if err != nil {
		return err
	}

	task, err := object.NewCommon(c, objs[0]).Rename(ctx, f.Arg(1))
	if err != nil {
		return err
	}

	logger := cmd.ProgressLogger(fmt.Sprintf("renaming %s... ", objs[0]))
	_, err = task.WaitForResult(ctx, logger)
	logger.Wait()

	return err
}
