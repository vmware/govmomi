// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package volume

import (
	"context"
	"errors"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/cns/types"
	"github.com/vmware/govmomi/vim25/soap"
)

type rm struct {
	*flags.ClientFlag

	keep bool
}

func init() {
	cli.Register("volume.rm", &rm{})
}

func (cmd *rm) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.BoolVar(&cmd.keep, "keep", false, "Keep backing disk")
}

func (cmd *rm) Usage() string {
	return "ID"
}

func (cmd *rm) Description() string {
	return `Remove CNS volume.

Note: if volume.rm returns not found errors,
consider using 'govc disk.ls -R' to reconcile the datastore inventory.

Examples:
  govc volume.rm f75989dc-95b9-4db7-af96-8583f24bc59d`
}

func (cmd *rm) Run(ctx context.Context, f *flag.FlagSet) error {
	// Despite the method signature, CnsDeleteVolume can only delete 1 at a time.
	if f.NArg() != 1 {
		return flag.ErrHelp
	}

	c, err := cmd.CnsClient()
	if err != nil {
		return err
	}

	ids := []types.CnsVolumeId{{Id: f.Arg(0)}}

	task, err := c.DeleteVolume(ctx, ids, !cmd.keep)
	if err != nil {
		return err
	}

	info, err := task.WaitForResult(ctx, nil)
	if err != nil {
		return err
	}

	if res, ok := info.Result.(types.CnsVolumeOperationBatchResult); ok {
		for _, r := range res.VolumeResults {
			fault := r.GetCnsVolumeOperationResult().Fault

			if fault != nil {
				if fault.Fault != nil {
					return soap.WrapVimFault(fault.Fault)
				}
				return errors.New(fault.LocalizedMessage)
			}
		}
	}

	return nil
}
