// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package fields

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/object"
)

type rename struct {
	*flags.ClientFlag
}

func init() {
	cli.Register("fields.rename", &rename{})
}

func (cmd *rename) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
}

func (cmd *rename) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *rename) Usage() string {
	return "KEY NAME"
}

func (cmd *rename) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 2 {
		return flag.ErrHelp
	}

	c, err := cmd.Client()
	if err != nil {
		return err
	}

	m, err := object.GetCustomFieldsManager(c)
	if err != nil {
		return err
	}

	key, err := m.FindKey(ctx, f.Arg(0))
	if err != nil {
		return err
	}

	name := f.Arg(1)

	return m.Rename(ctx, key, name)
}
