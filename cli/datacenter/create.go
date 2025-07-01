// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package datacenter

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
)

type create struct {
	*flags.FolderFlag
}

func init() {
	cli.Register("datacenter.create", &create{})
}

func (cmd *create) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.FolderFlag, ctx = flags.NewFolderFlag(ctx)
	cmd.FolderFlag.Register(ctx, f)
}

func (cmd *create) Usage() string {
	return "NAME..."
}

func (cmd *create) Description() string {
	return `Create datacenter NAME.

Examples:
  govc datacenter.create MyDC # create
  govc object.destroy /MyDC   # delete`
}

func (cmd *create) Process(ctx context.Context) error {
	if err := cmd.FolderFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *create) Run(ctx context.Context, f *flag.FlagSet) error {
	folder, err := cmd.FolderOrDefault("/")
	if err != nil {
		return err
	}

	if f.NArg() == 0 {
		return flag.ErrHelp
	}

	for _, name := range f.Args() {
		_, err := folder.CreateDatacenter(ctx, name)
		if err != nil {
			return err
		}
	}

	return nil
}
