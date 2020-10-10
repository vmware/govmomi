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

	"github.com/vmware/govmomi/cns"
	"github.com/vmware/govmomi/cns/types"
	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
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

Examples:
  govc volume.rm f75989dc-95b9-4db7-af96-8583f24bc59d
  govc volume.rm $(govc volume.ls -i pvc-de368f19-a997-4d5d-9eae-4496f10f429a)`
}

func (cmd *rm) Run(ctx context.Context, f *flag.FlagSet) error {
	vc, err := cmd.Client()
	if err != nil {
		return err
	}
	_ = vc.UseServiceVersion("vsan")

	c, err := cns.NewClient(ctx, vc)
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
