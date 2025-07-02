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

type reset struct {
	*PerformanceFlag
}

func init() {
	cli.Register("metric.reset", &reset{})
}

func (cmd *reset) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.PerformanceFlag, ctx = NewPerformanceFlag(ctx)
	cmd.PerformanceFlag.Register(ctx, f)
}

func (cmd *reset) Usage() string {
	return "NAME..."
}

func (cmd *reset) Description() string {
	return `Reset counter NAME to the default level of data collection.

Examples:
  govc metric.reset net.bytesRx.average net.bytesTx.average`
}

func (cmd *reset) Process(ctx context.Context) error {
	if err := cmd.PerformanceFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *reset) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() == 0 {
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

	var ids []int32

	for _, name := range f.Args() {
		counter, ok := counters[name]
		if !ok {
			return cmd.ErrNotFound(name)
		}

		ids = append(ids, counter.Key)
	}

	_, err = methods.ResetCounterLevelMapping(ctx, m.Client(), &types.ResetCounterLevelMapping{
		This:     m.Reference(),
		Counters: ids,
	})

	return err
}
