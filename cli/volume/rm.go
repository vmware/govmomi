/*
Copyright (c) 2020 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
}

func init() {
	cli.Register("volume.rm", &rm{})
}

func (cmd *rm) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
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
	c, err := cmd.CnsClient()
	if err != nil {
		return err
	}

	ids := []types.CnsVolumeId{{Id: f.Arg(0)}}

	task, err := c.DeleteVolume(ctx, ids, true)
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
