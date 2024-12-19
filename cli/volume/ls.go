/*
Copyright (c) 2020-2023 VMware, Inc. All Rights Reserved.

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
	"flag"
	"fmt"
	"io"
	"log"
	"strings"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/cns/types"
	"github.com/vmware/govmomi/units"
	vim "github.com/vmware/govmomi/vim25/types"
)

type ls struct {
	*flags.ClientFlag
	*flags.DatastoreFlag
	*flags.OutputFlag

	types.CnsQueryFilter

	long bool
	id   bool
	disk bool
	back bool
}

func init() {
	cli.Register("volume.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.DatastoreFlag, ctx = flags.NewDatastoreFlag(ctx)
	cmd.DatastoreFlag.Register(ctx, f)

	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)

	f.BoolVar(&cmd.long, "l", false, "Long listing format")
	f.BoolVar(&cmd.id, "i", false, "List volume ID only")
	f.BoolVar(&cmd.disk, "L", false, "List volume disk or file backing ID only")
	f.BoolVar(&cmd.back, "b", false, "List file backing path")
}

func (cmd *ls) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.DatastoreFlag.Process(ctx); err != nil {
		return err
	}
	return cmd.OutputFlag.Process(ctx)
}

func (cmd *ls) Usage() string {
	return "[ID...]"
}

func (cmd *ls) Description() string {
	return `List CNS volumes.

Examples:
  govc volume.ls
  govc volume.ls -l
  govc volume.ls -ds vsanDatastore
  govc volume.ls df86393b-5ae0-4fca-87d0-b692dbc67d45
  govc volume.ls -json $id | jq -r .volume[].backingObjectDetails.backingDiskPath
  govc volume.ls -b $id # verify backingDiskPath exists
  govc disk.ls -l $(govc volume.ls -L pvc-9744a4ff-07f4-43c4-b8ed-48ea7a528734)`
}

type lsWriter struct {
	Volume []types.CnsVolume                    `json:"volume"`
	Info   []types.BaseCnsVolumeOperationResult `json:"info,omitempty"`
	cmd    *ls
}

func (r *lsWriter) Dump() any {
	if len(r.Info) != 0 {
		return r.Info
	}
	return r.Volume
}

func (r *lsWriter) Write(w io.Writer) error {
	if r.cmd.id {
		for _, volume := range r.Volume {
			fmt.Fprintln(r.cmd.Out, volume.VolumeId.Id)
		}
		return nil
	}

	if r.cmd.disk {
		for _, volume := range r.Volume {
			var id string
			switch backing := volume.BackingObjectDetails.(type) {
			case *types.CnsBlockBackingDetails:
				id = backing.BackingDiskId
			case *types.CnsFileBackingDetails:
				id = backing.BackingFileId
			case *types.CnsVsanFileShareBackingDetails:
				id = backing.Name
			default:
				log.Printf("%s unknown backing type: %T", volume.VolumeId.Id, backing)
			}
			fmt.Fprintln(r.cmd.Out, id)

		}
		return nil
	}

	tw := tabwriter.NewWriter(r.cmd.Out, 2, 0, 2, ' ', 0)

	for _, volume := range r.Volume {
		fmt.Printf("%s\t%s", volume.VolumeId.Id, volume.Name)
		if r.cmd.back {
			fmt.Printf("\t%s", r.backing(volume.VolumeId))
		}
		if r.cmd.long {
			capacity := volume.BackingObjectDetails.GetCnsBackingObjectDetails().CapacityInMb
			c := volume.Metadata.ContainerCluster
			fmt.Printf("\t%s\t%s\t%s", units.ByteSize(capacity*1024*1024), c.ClusterType, c.ClusterId)
		}
		fmt.Println()
	}

	return tw.Flush()
}

func (r *lsWriter) backing(id types.CnsVolumeId) string {
	for _, info := range r.Info {
		res, ok := info.(*types.CnsQueryVolumeInfoResult)
		if !ok {
			continue
		}

		switch vol := res.VolumeInfo.(type) {
		case *types.CnsBlockVolumeInfo:
			if vol.VStorageObject.Config.Id.Id == id.Id {
				switch backing := vol.VStorageObject.Config.BaseConfigInfo.Backing.(type) {
				case *vim.BaseConfigInfoDiskFileBackingInfo:
					return backing.FilePath
				}
			}
		}

		if fault := res.Fault; fault != nil {
			if f, ok := fault.Fault.(types.CnsFault); ok {
				if strings.Contains(f.Reason, id.Id) {
					return f.Reason
				}
			}
		}
	}
	return "???"
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {
	ds, err := cmd.DatastoreIfSpecified()
	if err != nil {
		return err
	}

	if ds != nil {
		cmd.Datastores = []vim.ManagedObjectReference{ds.Reference()}
	}

	c, err := cmd.CnsClient()
	if err != nil {
		return err
	}

	for _, arg := range f.Args() {
		cmd.VolumeIds = append(cmd.VolumeIds, types.CnsVolumeId{Id: arg})
	}

	var volumes []types.CnsVolume
	var info []types.BaseCnsVolumeOperationResult

	for {
		res, err := c.QueryVolume(ctx, cmd.CnsQueryFilter)
		if err != nil {
			return err
		}

		volumes = append(volumes, res.Volumes...)

		if res.Cursor.Offset == res.Cursor.TotalRecords || len(res.Volumes) == 0 {
			break
		}

		cmd.Cursor = &res.Cursor
	}

	if cmd.back {
		ids := make([]types.CnsVolumeId, len(volumes))
		for i := range volumes {
			ids[i] = volumes[i].VolumeId
		}

		task, err := c.QueryVolumeInfo(ctx, ids)
		if err != nil {
			return err
		}

		res, err := task.WaitForResult(ctx, nil)
		if err != nil {
			return err
		}

		if batch, ok := res.Result.(types.CnsVolumeOperationBatchResult); ok {
			info = batch.VolumeResults
		}
	}

	return cmd.WriteResult(&lsWriter{volumes, info, cmd})
}
