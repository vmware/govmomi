// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package volume

import (
	"context"
	"errors"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/cns"
	"github.com/vmware/govmomi/cns/types"
	"github.com/vmware/govmomi/units"
)

type extend struct {
	*flags.ClientFlag

	size units.ByteSize
}

func init() {
	cli.Register("volume.extend", &extend{})
}

func (cmd *extend) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.Var(&cmd.size, "size", "New size of new volume")
}

func (cmd *extend) Usage() string {
	return "ID"
}

func (cmd *extend) Description() string {
	return `Extend CNS volume.

Examples:
  govc volume.extend -size 10GB f75989dc-95b9-4db7-af96-8583f24bc59d`
}

func (cmd *extend) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}

	c, err := cmd.CnsClient()
	if err != nil {
		return err
	}

	spec := []types.CnsVolumeExtendSpec{{
		VolumeId: types.CnsVolumeId{
			Id: f.Arg(0),
		},
		CapacityInMb: int64(cmd.size) / units.MB,
	}}

	task, err := c.ExtendVolume(ctx, spec)
	if err != nil {
		return err
	}

	info, err := cns.GetTaskInfo(ctx, task)
	if err != nil {
		return err
	}

	res, err := cns.GetTaskResult(ctx, info)
	if err != nil {
		return err
	}

	fault := res.GetCnsVolumeOperationResult().Fault
	if fault != nil {
		return errors.New(fault.LocalizedMessage)
	}

	return nil
}
