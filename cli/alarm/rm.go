// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package alarm

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/types"
)

type rm struct {
	*flags.ClientFlag
}

func init() {
	cli.Register("alarm.rm", &rm{}, true)
}

func (cmd *rm) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
}

func (cmd *rm) Usage() string {
	return "ID"
}

func (cmd *rm) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.Client()
	if err != nil {
		return err
	}

	id := f.Arg(0)
	var ref types.ManagedObjectReference

	if !ref.FromString(id) {
		ref.Type = "Alarm"
		ref.Value = id
	}

	_, err = methods.RemoveAlarm(ctx, c, &types.RemoveAlarm{
		This: ref,
	})

	return err
}
