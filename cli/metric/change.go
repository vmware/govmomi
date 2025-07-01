// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package metric

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/types"
)

type change struct {
	*PerformanceFlag

	level  int
	device int
}

func init() {
	cli.Register("metric.change", &change{})
}

func (cmd *change) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.PerformanceFlag, ctx = NewPerformanceFlag(ctx)
	cmd.PerformanceFlag.Register(ctx, f)

	f.IntVar(&cmd.level, "level", 0, "Level for the aggregate counter")
	f.IntVar(&cmd.device, "device-level", 0, "Level for the per device counter")
}

func (cmd *change) Usage() string {
	return "NAME..."
}

func (cmd *change) Description() string {
	return `Change counter NAME levels.

Examples:
  govc metric.change -level 1 net.bytesRx.average net.bytesTx.average`
}

func (cmd *change) Process(ctx context.Context) error {
	if err := cmd.PerformanceFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *change) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() == 0 || (cmd.level == 0 && cmd.device == 0) {
		return flag.ErrHelp
	}

	m, err := cmd.Manager(ctx)
	if err != nil {
		return err
	}

	counters, err := m.CounterInfoByName(ctx)
	if err != nil {
		return err
	}

	var mapping []types.PerformanceManagerCounterLevelMapping

	for _, name := range f.Args() {
		counter, ok := counters[name]
		if !ok {
			return cmd.ErrNotFound(name)
		}

		mapping = append(mapping, types.PerformanceManagerCounterLevelMapping{
			CounterId:      counter.Key,
			AggregateLevel: int32(cmd.level),
			PerDeviceLevel: int32(cmd.device),
		})
	}

	_, err = methods.UpdateCounterLevelMapping(ctx, m.Client(), &types.UpdateCounterLevelMapping{
		This:            m.Reference(),
		CounterLevelMap: mapping,
	})

	return err
}
