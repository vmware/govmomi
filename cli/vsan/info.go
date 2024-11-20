/*
Copyright (c) 2021-2024 VMware, Inc. All Rights Reserved.

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

package vsan

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vsan"
	"github.com/vmware/govmomi/vsan/types"
)

type info struct {
	*flags.DatacenterFlag
}

func init() {
	cli.Register("vsan.info", &info{})
}

func (cmd *info) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatacenterFlag, ctx = flags.NewDatacenterFlag(ctx)
	cmd.DatacenterFlag.Register(ctx, f)
}

func (cmd *info) Usage() string {
	return "CLUSTER..."
}

func (cmd *info) Description() string {
	return `Display vSAN configuration.

Examples:
  govc vsan.info
  govc vsan.info ClusterA
  govc vsan.info -json`
}

func (cmd *info) Run(ctx context.Context, f *flag.FlagSet) error {
	vc, err := cmd.Client()
	if err != nil {
		return err
	}

	c, err := vsan.NewClient(ctx, vc)
	if err != nil {
		return err
	}

	finder, err := cmd.Finder()
	if err != nil {
		return err
	}

	args := f.Args()
	if len(args) == 0 {
		args = []string{"*"}
	}

	var res []Cluster

	for _, arg := range args {
		clusters, err := finder.ClusterComputeResourceList(ctx, arg)
		if err != nil {
			return err
		}

		for _, cluster := range clusters {
			info, err := c.VsanClusterGetConfig(ctx, cluster.Reference())
			if err != nil {
				return err
			}
			res = append(res, Cluster{cluster.InventoryPath, info})
		}
	}

	return cmd.WriteResult(&infoResult{res})
}

type Cluster struct {
	Path string                  `json:"path"`
	Info *types.VsanConfigInfoEx `json:"info"`
}

type infoResult struct {
	Clusters []Cluster `json:"clusters"`
}

func (r *infoResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(os.Stdout, 2, 0, 2, ' ', 0)

	for _, cluster := range r.Clusters {
		fmt.Fprintf(tw, "Path:\t%s\n", cluster.Path)
		fmt.Fprintf(tw, "  Enabled:\t%t\n", *cluster.Info.Enabled)
		if unmap := cluster.Info.UnmapConfig; unmap != nil {
			fmt.Fprintf(tw, "  Unmap Enabled:\t%t\n", unmap.Enable)
		}
		if fs := cluster.Info.FileServiceConfig; fs != nil {
			fmt.Fprintf(tw, "  FileService Enabled:\t%t\n", fs.Enabled)
		}
	}

	return tw.Flush()
}
