/*
Copyright (c) 2017 VMware, Inc. All Rights Reserved.

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

package metric

import (
	"context"
	"flag"
	"fmt"
	"text/tabwriter"

	"github.com/vmware/govmomi/govc/cli"
)

type ls struct {
	*PerformanceFlag

	long bool
}

func init() {
	cli.Register("metric.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.PerformanceFlag, ctx = NewPerformanceFlag(ctx)
	cmd.PerformanceFlag.Register(ctx, f)

	f.BoolVar(&cmd.long, "l", false, "Long listing format")
}

func (cmd *ls) Usage() string {
	return "PATH"
}

func (cmd *ls) Description() string {
	return `List available metrics for PATH.

Examples:
  govc metric.ls /dc1/host/cluster1
  govc metric.ls datastore/*
  govc metric.ls vm/* | grep mem. | xargs govc metric.sample vm/*`
}

func (cmd *ls) Process(ctx context.Context) error {
	if err := cmd.PerformanceFlag.Process(ctx); err != nil {
		return err
	}
	return nil
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

	counters, err := m.CounterInfoByKey(ctx)
	if err != nil {
		return err
	}

	mids, err := m.AvailableMetric(ctx, objs[0], cmd.Interval(s.RefreshRate))
	if err != nil {
		return err
	}

	tw := tabwriter.NewWriter(cmd.Out, 2, 0, 2, ' ', 0)
	cmd.Out = tw

	for _, id := range mids {
		if id.Instance != "" {
			continue
		}

		info := counters[id.CounterId]

		if cmd.long {
			fmt.Fprintf(cmd.Out, "%s\t%s\n", info.Name(),
				info.NameInfo.GetElementDescription().Label)
			continue
		}

		fmt.Fprintln(cmd.Out, info.Name())
	}

	return tw.Flush()
}
