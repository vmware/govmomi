// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package interval

import (
	"context"
	"flag"
	"fmt"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/metric"
)

type info struct {
	*metric.PerformanceFlag
}

func init() {
	cli.Register("metric.interval.info", &info{})
}

func (cmd *info) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.PerformanceFlag, ctx = metric.NewPerformanceFlag(ctx)
	cmd.PerformanceFlag.Register(ctx, f)
}

func (cmd *info) Description() string {
	return `List historical metric intervals.

Examples:
  govc metric.interval.info
  govc metric.interval.info -i 300`
}

func (cmd *info) Process(ctx context.Context) error {
	if err := cmd.PerformanceFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *info) Run(ctx context.Context, f *flag.FlagSet) error {
	m, err := cmd.Manager(ctx)
	if err != nil {
		return err
	}

	intervals, err := m.HistoricalInterval(ctx)
	if err != nil {
		return err
	}

	tw := tabwriter.NewWriter(cmd.Out, 2, 0, 2, ' ', 0)
	cmd.Out = tw

	interval := cmd.Interval(0)

	for _, i := range intervals {
		if interval != 0 && i.SamplingPeriod != interval {
			continue
		}

		period := (time.Duration(i.SamplingPeriod) * time.Second).String()
		period = strings.TrimSuffix(period, "0s")
		if strings.Contains(period, "h") {
			period = strings.TrimSuffix(period, "0m")
		}
		samples := i.Length / i.SamplingPeriod

		fmt.Fprintf(cmd.Out, "ID:\t%d\n", i.SamplingPeriod)
		fmt.Fprintf(cmd.Out, "  Enabled:\t%t\n", i.Enabled)
		fmt.Fprintf(cmd.Out, "  Interval:\t%s\n", period)
		fmt.Fprintf(cmd.Out, "  Available Samples:\t%d\n", samples)
		fmt.Fprintf(cmd.Out, "  Name:\t%s\n", i.Name)
		fmt.Fprintf(cmd.Out, "  Level:\t%d\n", i.Level)
	}

	return tw.Flush()
}
