// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package metric

import (
	"context"
	"flag"
	"fmt"
	"math"
	"strconv"

	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/performance"
)

type PerformanceFlag struct {
	*flags.DatacenterFlag
	*flags.OutputFlag

	m *performance.Manager

	interval string
}

func NewPerformanceFlag(ctx context.Context) (*PerformanceFlag, context.Context) {
	f := &PerformanceFlag{}
	f.DatacenterFlag, ctx = flags.NewDatacenterFlag(ctx)
	f.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	return f, ctx
}

func (f *PerformanceFlag) Register(ctx context.Context, fs *flag.FlagSet) {
	f.DatacenterFlag.Register(ctx, fs)
	f.OutputFlag.Register(ctx, fs)

	fs.StringVar(&f.interval, "i", "real", "Interval ID (real|day|week|month|year)")
}

func (f *PerformanceFlag) Process(ctx context.Context) error {
	if err := f.DatacenterFlag.Process(ctx); err != nil {
		return err
	}
	if err := f.OutputFlag.Process(ctx); err != nil {
		return err
	}

	return nil
}

func (f *PerformanceFlag) Manager(ctx context.Context) (*performance.Manager, error) {
	if f.m != nil {
		return f.m, nil
	}

	c, err := f.Client()
	if err != nil {
		return nil, err
	}

	f.m = performance.NewManager(c)

	f.m.Sort = true

	return f.m, err
}

func (f *PerformanceFlag) Interval(val int32) int32 {
	var interval int32

	if f.interval != "" {
		if i, ok := performance.Intervals[f.interval]; ok {
			interval = i
		} else {
			n, err := strconv.ParseUint(f.interval, 10, 32)
			if err != nil {
				panic(err)
			}

			if n > math.MaxInt32 {
				panic(fmt.Errorf("value out of range for int32: %d", n))
			} else {
				interval = int32(n)
			}
		}
	}

	if interval == 0 {
		if val == -1 {
			// realtime not supported
			return 300
		}

		return val
	}

	return interval
}

func (f *PerformanceFlag) ErrNotFound(name string) error {
	return fmt.Errorf("counter %q not found", name)
}
