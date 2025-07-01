// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package interval

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/cli/metric"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/types"
)

type change struct {
	*metric.PerformanceFlag

	enabled *bool
	level   int
}

func init() {
	cli.Register("metric.interval.change", &change{})
}

func (cmd *change) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.PerformanceFlag, ctx = metric.NewPerformanceFlag(ctx)
	cmd.PerformanceFlag.Register(ctx, f)

	f.Var(flags.NewOptionalBool(&cmd.enabled), "enabled", "Enable or disable")
	f.IntVar(&cmd.level, "level", 0, "Level")
}

func (cmd *change) Description() string {
	return `Change historical metric intervals.

Examples:
  govc metric.interval.change -i 300 -level 2
  govc metric.interval.change -i 86400 -enabled=false`
}

func (cmd *change) Process(ctx context.Context) error {
	if err := cmd.PerformanceFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *change) Run(ctx context.Context, f *flag.FlagSet) error {
	m, err := cmd.Manager(ctx)
	if err != nil {
		return err
	}

	intervals, err := m.HistoricalInterval(ctx)
	if err != nil {
		return err
	}

	interval := cmd.Interval(0)
	if interval == 0 {
		return flag.ErrHelp
	}

	var current *types.PerfInterval

	for _, i := range intervals {
		if i.SamplingPeriod == interval {
			current = &i
			break
		}
	}

	if current == nil {
		return fmt.Errorf("%d interval ID not found", interval)
	}

	if cmd.level != 0 {
		if cmd.level > 4 {
			return flag.ErrHelp
		}
		current.Level = int32(cmd.level)
	}

	if cmd.enabled != nil {
		current.Enabled = *cmd.enabled
	}

	_, err = methods.UpdatePerfInterval(ctx, m.Client(), &types.UpdatePerfInterval{
		This:     m.Reference(),
		Interval: *current,
	})

	return err
}
