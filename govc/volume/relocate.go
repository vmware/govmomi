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
	"fmt"
	"io"
	"text/tabwriter"

	cnstypes "github.com/vmware/govmomi/cns/types"
	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vim25/soap"
)

type relocate struct {
	*flags.ClientFlag
	*flags.DatastoreFlag
	*flags.OutputFlag
}

func init() {
	cli.Register("volume.relocate", &relocate{})
}

func (cmd *relocate) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.DatastoreFlag, ctx = flags.NewDatastoreFlag(ctx)
	cmd.DatastoreFlag.Register(ctx, f)

	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)
}

func (cmd *relocate) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.DatastoreFlag.Process(ctx); err != nil {
		return err
	}
	return cmd.OutputFlag.Process(ctx)
}

func (cmd *relocate) Usage() string {
	return "ID [ID...]"
}

func (cmd *relocate) Description() string {
	return `Relocate one or more CNS volumes to the target datastore.

All IDs are submitted in a single batch RelocateVolume call.
Per-volume results are printed; the command exits non-zero if any relocation failed.

Examples:
  govc volume.relocate -ds vsanDatastore f75989dc-95b9-4db7-af96-8583f24bc59d
  govc volume.relocate -ds vsanDatastore id1 id2 id3
  govc volume.relocate -ds vsanDatastore -json id1 id2 | jq .`
}

// relocateResult holds the per-volume outcomes of a batch relocation and
// implements flags.OutputWriter for both text and JSON rendering.
type relocateResult struct {
	Results []relocateVolumeResult `json:"results"`
}

type relocateVolumeResult struct {
	VolumeID string `json:"volumeId"`
	Status   string `json:"status"`
	Error    string `json:"error,omitempty"`
}

func (r *relocateResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)
	for _, res := range r.Results {
		if res.Error != "" {
			fmt.Fprintf(tw, "%s\tFAILED\t%s\n", res.VolumeID, res.Error)
		} else {
			fmt.Fprintf(tw, "%s\tOK\n", res.VolumeID)
		}
	}
	return tw.Flush()
}

func (cmd *relocate) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() == 0 {
		return flag.ErrHelp
	}

	ds, err := cmd.Datastore()
	if err != nil {
		return err
	}

	c, err := cmd.CnsClient()
	if err != nil {
		return err
	}

	out := new(relocateResult)
	var firstErr error

	for _, id := range f.Args() {
		spec := cnstypes.CnsBlockVolumeRelocateSpec{
			CnsVolumeRelocateSpec: cnstypes.CnsVolumeRelocateSpec{
				VolumeId:  cnstypes.CnsVolumeId{Id: id},
				Datastore: ds.Reference(),
			},
		}

		entry := relocateVolumeResult{VolumeID: id, Status: "OK"}

		task, err := c.RelocateVolume(ctx, spec)
		if err != nil {
			entry.Status = "FAILED"
			entry.Error = err.Error()
			if firstErr == nil {
				firstErr = err
			}
			out.Results = append(out.Results, entry)
			continue
		}

		info, err := task.WaitForResult(ctx, nil)
		if err != nil {
			entry.Status = "FAILED"
			entry.Error = err.Error()
			if firstErr == nil {
				firstErr = err
			}
			out.Results = append(out.Results, entry)
			continue
		}

		if batchRes, ok := info.Result.(cnstypes.CnsVolumeOperationBatchResult); ok {
			for _, r := range batchRes.VolumeResults {
				opRes := r.GetCnsVolumeOperationResult()
				if opRes.Fault != nil {
					entry.Status = "FAILED"
					entry.Error = opRes.Fault.LocalizedMessage
					if firstErr == nil {
						if opRes.Fault.Fault != nil {
							firstErr = soap.WrapVimFault(opRes.Fault.Fault)
						} else {
							firstErr = errors.New(opRes.Fault.LocalizedMessage)
						}
					}
				}
			}
		}

		out.Results = append(out.Results, entry)
	}

	if wErr := cmd.WriteResult(out); wErr != nil {
		return wErr
	}

	return firstErr
}
