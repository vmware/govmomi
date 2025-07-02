// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package metadata

import (
	"context"
	"flag"
	"strings"
	"time"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vslm"
)

type update struct {
	*flags.ClientFlag

	remove flags.StringList
}

func init() {
	cli.Register("disk.metadata.update", &update{})
}

func (cmd *update) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.Var(&cmd.remove, "d", "Delete keys")
}

func (cmd *update) Usage() string {
	return "ID"
}

func (cmd *update) Description() string {
	return `Update metadata for disk ID.

Examples:
  govc disk.metadata.update $id foo=bar biz=baz
  govc disk.metadata.update -d foo -d biz $id`
}

func (cmd *update) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.Client()
	if err != nil {
		return err
	}

	vc, err := vslm.NewClient(ctx, c)
	if err != nil {
		return err
	}

	m := vslm.NewGlobalObjectManager(vc)

	id := types.ID{Id: f.Arg(0)}

	var update []types.KeyValue

	for _, arg := range f.Args()[1:] {
		kv := strings.SplitN(arg, "=", 2)
		if len(kv) == 1 {
			kv = append(kv, "")
		}
		update = append(update, types.KeyValue{Key: kv[0], Value: kv[1]})
	}

	task, err := m.UpdateMetadata(ctx, id, update, cmd.remove)
	if err != nil {
		return err
	}

	_, err = task.Wait(ctx, time.Hour)
	return err
}
