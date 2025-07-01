// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package cluster

import (
	"context"
	"flag"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/units"
	"github.com/vmware/govmomi/vim25/mo"
)

type usage struct {
	*flags.DatacenterFlag

	shared bool
}

func init() {
	cli.Register("cluster.usage", &usage{})
}

func (cmd *usage) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatacenterFlag, ctx = flags.NewDatacenterFlag(ctx)
	cmd.DatacenterFlag.Register(ctx, f)

	f.BoolVar(&cmd.shared, "S", false, "Exclude host local storage")
}

func (cmd *usage) Usage() string {
	return "CLUSTER"
}

func (cmd *usage) Description() string {
	return `Cluster resource usage summary.

Examples:
  govc cluster.usage ClusterName
  govc cluster.usage -S ClusterName # summarize shared storage only
  govc cluster.usage -json ClusterName | jq -r .cpu.summary.usage`
}

func (cmd *usage) Run(ctx context.Context, f *flag.FlagSet) error {
	finder, err := cmd.Finder()
	if err != nil {
		return err
	}

	obj, err := finder.ClusterComputeResource(ctx, f.Arg(0))
	if err != nil {
		return err
	}

	var res Usage
	var cluster mo.ClusterComputeResource
	var hosts []mo.HostSystem
	var datastores []mo.Datastore

	pc := property.DefaultCollector(obj.Client())

	err = pc.RetrieveOne(ctx, obj.Reference(), []string{"datastore", "host"}, &cluster)
	if err != nil {
		return err
	}

	err = pc.Retrieve(ctx, cluster.Host, []string{"summary"}, &hosts)
	if err != nil {
		return err
	}

	for _, host := range hosts {
		res.CPU.Capacity += int64(int32(host.Summary.Hardware.NumCpuCores) * host.Summary.Hardware.CpuMhz)
		res.CPU.Used += int64(host.Summary.QuickStats.OverallCpuUsage)

		res.Memory.Capacity += host.Summary.Hardware.MemorySize
		res.Memory.Used += int64(host.Summary.QuickStats.OverallMemoryUsage) << 20
	}

	err = pc.Retrieve(ctx, cluster.Datastore, []string{"summary"}, &datastores)
	if err != nil {
		return err
	}

	for _, datastore := range datastores {
		shared := datastore.Summary.MultipleHostAccess
		if cmd.shared && shared != nil && *shared == false {
			continue
		}

		res.Storage.Capacity += datastore.Summary.Capacity
		res.Storage.Free += datastore.Summary.FreeSpace
	}

	res.CPU.Free = res.CPU.Capacity - res.CPU.Used
	res.CPU.summarize(ghz)

	res.Memory.Free = res.Memory.Capacity - res.Memory.Used
	res.Memory.summarize(size)

	res.Storage.Used = res.Storage.Capacity - res.Storage.Free
	res.Storage.summarize(size)

	return cmd.WriteResult(&res)
}

type ResourceUsageSummary struct {
	Used     string `json:"used"`
	Free     string `json:"free"`
	Capacity string `json:"capacity"`
	Usage    string `json:"usage"`
}

type ResourceUsage struct {
	Used     int64                `json:"used"`
	Free     int64                `json:"free"`
	Capacity int64                `json:"capacity"`
	Usage    float64              `json:"usage"`
	Summary  ResourceUsageSummary `json:"summary"`
}

func (r *ResourceUsage) summarize(f func(int64) string) {
	r.Usage = 100 * float64(r.Used) / float64(r.Capacity)

	r.Summary.Usage = fmt.Sprintf("%.1f", r.Usage)
	r.Summary.Capacity = f(r.Capacity)
	r.Summary.Used = f(r.Used)
	r.Summary.Free = f(r.Free)
}

func (r *ResourceUsage) write(w io.Writer, label string) {
	fmt.Fprintf(w, "%s usage:\t%s%%\n", label, r.Summary.Usage)
	fmt.Fprintf(w, "%s capacity:\t%s\n", label, r.Summary.Capacity)
	fmt.Fprintf(w, "%s used:\t%s\n", label, r.Summary.Used)
	fmt.Fprintf(w, "%s free:\t%s\n", label, r.Summary.Free)
}

func ghz(val int64) string {
	return fmt.Sprintf("%.1fGHz", float64(val)/1000)
}

func size(val int64) string {
	return units.ByteSize(val).String()
}

type Usage struct {
	Memory  ResourceUsage `json:"memory"`
	CPU     ResourceUsage `json:"cpu"`
	Storage ResourceUsage `json:"storage"`
}

func (r *Usage) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)

	r.CPU.write(tw, "CPU")
	fmt.Fprintf(tw, "\t\n")

	r.Memory.write(tw, "Memory")
	fmt.Fprintf(tw, "\t\n")

	r.Storage.write(tw, "Storage")

	return tw.Flush()
}
