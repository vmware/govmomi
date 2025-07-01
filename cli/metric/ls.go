// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package metric

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/performance"
	"github.com/vmware/govmomi/vim25/types"
)

type ls struct {
	*PerformanceFlag

	group    string
	long     bool
	longlong bool
}

func init() {
	cli.Register("metric.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.PerformanceFlag, ctx = NewPerformanceFlag(ctx)
	cmd.PerformanceFlag.Register(ctx, f)

	f.StringVar(&cmd.group, "g", "", "List a specific Group")
	f.BoolVar(&cmd.long, "l", false, "Long listing format")
	f.BoolVar(&cmd.longlong, "L", false, "Longer listing format")
}

func (cmd *ls) Usage() string {
	return "PATH"
}

func (cmd *ls) Description() string {
	return `List available metrics for PATH.

The default output format is the metric name.
The '-l' flag includes the metric description.
The '-L' flag includes the metric units, instance count (if any) and description.
The instance count is prefixed with a single '@'.
If no aggregate is provided for the metric, instance count is prefixed with two '@@'.

Examples:
  govc metric.ls /dc1/host/cluster1
  govc metric.ls datastore/*
  govc metric.ls -L -g CPU /dc1/host/cluster1/host1
  govc metric.ls vm/* | grep mem. | xargs govc metric.sample vm/*`
}

func (cmd *ls) Process(ctx context.Context) error {
	if err := cmd.PerformanceFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

type lsResult struct {
	cmd      *ls
	counters map[int32]*types.PerfCounterInfo
	performance.MetricList
}

func (r *lsResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)

	type count struct {
		aggregate bool
		instances int
	}
	seen := make(map[int32]*count)
	var res []types.PerfMetricId

	for _, id := range r.MetricList {
		if r.cmd.group != "" {
			info, ok := r.counters[id.CounterId]
			if !ok || info.GroupInfo.GetElementDescription().Label != r.cmd.group {
				continue
			}
		}

		c := seen[id.CounterId]
		if c == nil {
			c = new(count)
			seen[id.CounterId] = c
			res = append(res, id)
		}

		if id.Instance == "" {
			c.aggregate = true
		} else {
			c.instances++
		}
	}

	for _, id := range res {
		info, ok := r.counters[id.CounterId]
		if !ok {
			continue
		}

		switch {
		case r.cmd.long:
			fmt.Fprintf(tw, "%s\t%s\n", info.Name(),
				info.NameInfo.GetElementDescription().Label)
		case r.cmd.longlong:
			i := ""
			c := seen[id.CounterId]
			if !c.aggregate {
				i = "@"
			}
			if c.instances > 0 {
				i += fmt.Sprintf("@%d", c.instances)
			}
			fmt.Fprintf(tw, "%s\t%s\t%s\t%s\n", info.Name(),
				i,
				info.UnitInfo.GetElementDescription().Label,
				info.NameInfo.GetElementDescription().Label)
		default:
			fmt.Fprintln(w, info.Name())
		}
	}

	return tw.Flush()
}

func (r *lsResult) MarshalJSON() ([]byte, error) {
	m := make(map[string]*types.PerfCounterInfo)

	for _, id := range r.MetricList {
		if info, ok := r.counters[id.CounterId]; ok {
			m[info.Name()] = info
		}
	}

	return json.Marshal(m)
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}

	objs, err := cmd.ManagedObjects(ctx, f.Args())
	if err != nil {
		return err
	}

	m, err := cmd.Manager(ctx)
	if err != nil {
		return err
	}

	s, err := m.ProviderSummary(ctx, objs[0])
	if err != nil {
		return err
	}

	mids, err := m.AvailableMetric(ctx, objs[0], cmd.Interval(s.RefreshRate))
	if err != nil {
		return err
	}

	counters, err := m.CounterInfoByKey(ctx)
	if err != nil {
		return err
	}

	return cmd.WriteResult(&lsResult{cmd, counters, mids})
}
