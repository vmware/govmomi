// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package snapshot

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/disk"
	"github.com/vmware/govmomi/cli/flags"
)

type create struct {
	*flags.DatastoreFlag
}

func init() {
	cli.Register("disk.snapshot.create", &create{})
}

func (cmd *create) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatastoreFlag, ctx = flags.NewDatastoreFlag(ctx)
	cmd.DatastoreFlag.Register(ctx, f)
}

func (cmd *create) Usage() string {
	return "ID DESC"
}

func (cmd *create) Description() string {
	return `Create snapshot of ID on DS.

Examples:
  govc disk.snapshot.create b9fe5f17-3b87-4a03-9739-09a82ddcc6b0 my-disk-snapshot`
}

func (cmd *create) Run(ctx context.Context, f *flag.FlagSet) error {
	m, err := disk.NewManagerFromFlag(ctx, cmd.DatastoreFlag)
	if err != nil {
		return err
	}

	id := f.Arg(0)
	desc := f.Arg(1)

	res, err := m.CreateSnapshot(ctx, id, desc)
	if err != nil {
		return err
	}

	fmt.Println(res.Id)

	return nil
}
