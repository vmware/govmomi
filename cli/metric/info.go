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
	"strings"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/vim25/types"
)

type info struct {
	*PerformanceFlag
	group string
}

func init() {
	cli.Register("metric.info", &info{})
}

func (cmd *info) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.PerformanceFlag, ctx = NewPerformanceFlag(ctx)
	cmd.PerformanceFlag.Register(ctx, f)

	f.StringVar(&cmd.group, "g", "", "Show info for a specific Group")
}

func (cmd *info) Usage() string {
	return "PATH [NAME]..."
}

func (cmd *info) Description() string {
	return `Metric info for NAME.

If PATH is a value other than '-', provider summary and instance list are included
for the given object type.

If NAME is not specified, all available metrics for the given INTERVAL are listed.
An object PATH must be provided in this case.

Examples:
  govc metric.info vm/my-vm
  govc metric.info -i 300 vm/my-vm
  govc metric.info - cpu.usage.average
  govc metric.info /dc1/host/cluster cpu.usage.average`
}

func (cmd *info) Process(ctx context.Context) error {
	if err := cmd.PerformanceFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

type EntityDetail struct {
	Realtime   bool     `json:"realtime"`
	Historical bool     `json:"historical"`
	Instance   []string `json:"instance"`
}

type MetricInfo struct {
	Counter          *types.PerfCounterInfo `json:"counter"`
	Enabled          []string               `json:"enabled"`
	PerDeviceEnabled []string               `json:"perDeviceEnabled"`
	Detail           *EntityDetail          `json:"detail"`
}

type infoResult struct {
	Info []*MetricInfo `json:"info"`
	cmd  *info
}

func (r *infoResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)

	for _, info := range r.Info {
		counter := info.Counter

		fmt.Fprintf(tw, "Name:\t%s\n", counter.Name())
		fmt.Fprintf(tw, "  Label:\t%s\n", counter.NameInfo.GetElementDescription().Label)
		fmt.Fprintf(tw, "  Summary:\t%s\n", counter.NameInfo.GetElementDescription().Summary)
		fmt.Fprintf(tw, "  Group:\t%s\n", counter.GroupInfo.GetElementDescription().Label)
		fmt.Fprintf(tw, "  Unit:\t%s\n", counter.UnitInfo.GetElementDescription().Label)
		fmt.Fprintf(tw, "  Rollup type:\t%s\n", counter.RollupType)
		fmt.Fprintf(tw, "  Stats type:\t%s\n", counter.StatsType)
		fmt.Fprintf(tw, "  Level:\t%d\n", counter.Level)
		fmt.Fprintf(tw, "    Intervals:\t%s\n", strings.Join(info.Enabled, ","))
		fmt.Fprintf(tw, "  Per-device level:\t%d\n", counter.PerDeviceLevel)
		fmt.Fprintf(tw, "    Intervals:\t%s\n", strings.Join(info.PerDeviceEnabled, ","))

		summary := info.Detail
		if summary == nil {
			continue
		}

		fmt.Fprintf(tw, "  Realtime:\t%t\n", summary.Realtime)
		fmt.Fprintf(tw, "  Historical:\t%t\n", summary.Historical)
		fmt.Fprintf(tw, "  Instances:\t%s\n", strings.Join(summary.Instance, ","))
	}

	return tw.Flush()
}

func (r *infoResult) MarshalJSON() ([]byte, error) {
	m := make(map[string]*MetricInfo)

	for _, info := range r.Info {
		m[info.Counter.Name()] = info
	}

	return json.Marshal(m)
}

func (cmd *info) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() == 0 {
		return flag.ErrHelp
	}

	names := f.Args()[1:]

	m, err := cmd.Manager(ctx)
	if err != nil {
		return err
	}

	counters, err := m.CounterInfoByName(ctx)
	if err != nil {
		return err
	}

	intervals, err := m.HistoricalInterval(ctx)
	if err != nil {
		return err
	}
	enabled := intervals.Enabled()

	var summary *types.PerfProviderSummary
	var mids map[int32][]*types.PerfMetricId

	if f.Arg(0) == "-" {
		if len(names) == 0 {
			return flag.ErrHelp
		}
	} else {
		objs, err := cmd.ManagedObjects(ctx, f.Args()[:1])
		if err != nil {
			return err
		}

		summary, err = m.ProviderSummary(ctx, objs[0])
		if err != nil {
			return err
		}

		all, err := m.AvailableMetric(ctx, objs[0], cmd.Interval(summary.RefreshRate))
		if err != nil {
			return err
		}

		mids = all.ByKey()

		if len(names) == 0 {
			nc, _ := m.CounterInfoByKey(ctx)

			seen := make(map[int32]bool)
			for i := range all {
				id := &all[i]
				info, ok := nc[id.CounterId]
				if !ok || seen[id.CounterId] {
					continue
				}
				seen[id.CounterId] = true

				names = append(names, info.Name())
			}
		}
	}

	var metrics []*MetricInfo

	for _, name := range names {
		counter, ok := counters[name]
		if !ok {
			return cmd.ErrNotFound(name)
		}

		if cmd.group != "" {
			if counter.GroupInfo.GetElementDescription().Label != cmd.group {
				continue
			}
		}

		info := &MetricInfo{
			Counter:          counter,
			Enabled:          enabled[counter.Level],
			PerDeviceEnabled: enabled[counter.PerDeviceLevel],
		}

		metrics = append(metrics, info)

		if summary == nil {
			continue
		}

		var instances []string

		for _, id := range mids[counter.Key] {
			if id.Instance != "" {
				instances = append(instances, id.Instance)
			}
		}

		info.Detail = &EntityDetail{
			Realtime:   summary.CurrentSupported,
			Historical: summary.SummarySupported,
			Instance:   instances,
		}

	}

	return cmd.WriteResult(&infoResult{metrics, cmd})
}
