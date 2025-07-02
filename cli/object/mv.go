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
)

type mv struct {
	*flags.DatacenterFlag
}

func init() {
	cli.Register("object.mv", &mv{})
}

func (cmd *mv) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatacenterFlag, ctx = flags.NewDatacenterFlag(ctx)
	cmd.DatacenterFlag.Register(ctx, f)
}

func (cmd *mv) Usage() string {
	return "PATH... FOLDER"
}

func (cmd *mv) Description() string {
	return `Move managed entities to FOLDER.

Examples:
  govc folder.create /dc1/host/example
  govc object.mv /dc2/host/*.example.com /dc1/host/example`
}

func (cmd *mv) Process(ctx context.Context) error {
	if err := cmd.DatacenterFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *mv) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() < 2 {
		return flag.ErrHelp
	}

	finder, err := cmd.Finder()
	if err != nil {
		return err
	}

	n := f.NArg() - 1

	folder, err := finder.Folder(ctx, f.Arg(n))
	if err != nil {
		return err
	}

	objs, err := cmd.ManagedObjects(ctx, f.Args()[:n])
	if err != nil {
		return err
	}

	task, err := folder.MoveInto(ctx, objs)
	if err != nil {
		return err
	}

	logger := cmd.ProgressLogger(fmt.Sprintf("moving %d objects to %s... ", len(objs), folder.InventoryPath))
	_, err = task.WaitForResult(ctx, logger)
	logger.Wait()

	return err
}
