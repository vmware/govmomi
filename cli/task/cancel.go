// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package task

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/types"
)

type cancel struct {
	*flags.ClientFlag
}

func init() {
	cli.Register("task.cancel", &cancel{})
}

func (cmd *cancel) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
}

func (cmd *cancel) Description() string {
	return `Cancel tasks.

Examples:
  govc task.cancel task-759`
}

func (cmd *cancel) Usage() string {
	return "ID..."
}

func (cmd *cancel) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *cancel) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.Client()
	if err != nil {
		return err
	}

	for _, id := range f.Args() {
		var ref types.ManagedObjectReference

		if !ref.FromString(id) {
			ref.Type = "Task"
			ref.Value = id
		}

		err = object.NewTask(c, ref).Cancel(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
