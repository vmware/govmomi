// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package snapshot

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/disk"
	"github.com/vmware/govmomi/cli/flags"
)

type rm struct {
	*flags.DatastoreFlag
}

func init() {
	cli.Register("disk.snapshot.rm", &rm{})
}

func (cmd *rm) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatastoreFlag, ctx = flags.NewDatastoreFlag(ctx)
	cmd.DatastoreFlag.Register(ctx, f)
}

func (cmd *rm) Usage() string {
	return "ID SID"
}

func (cmd *rm) Description() string {
	return `Remove disk ID snapshot ID on DS.

Examples:
  govc disk.snapshot.rm ffe6a398-eb8e-4eaa-9118-e1f16b8b8e3c ecbca542-0a25-4127-a585-82e4047750d6`
}

func (cmd *rm) Run(ctx context.Context, f *flag.FlagSet) error {
	m, err := disk.NewManagerFromFlag(ctx, cmd.DatastoreFlag)
	if err != nil {
		return err
	}

	id := f.Arg(0)
	sid := f.Arg(1)

	return m.DeleteSnapshot(ctx, id, sid)
}
