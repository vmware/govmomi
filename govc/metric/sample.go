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
	"io"
	"text/tabwriter"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/performance"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type sample struct {
	*PerformanceFlag

	d        int
	n        int
	t        bool
	instance string
}

func init() {
	cli.Register("metric.sample", &sample{})
}

func (cmd *sample) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.PerformanceFlag, ctx = NewPerformanceFlag(ctx)
	cmd.PerformanceFlag.Register(ctx, f)

	f.IntVar(&cmd.d, "d", 30, "Limit object display name to D chars")
	f.IntVar(&cmd.n, "n", 6, "Max number of samples")
	f.BoolVar(&cmd.t, "t", false, "Include sample times")
	f.StringVar(&cmd.instance, "instance", "*", "Instance")
}

func (cmd *sample) Usage() string {
	return "PATH... NAME..."
}

func (cmd *sample) Description() string {
	return `Sample for object PATH of metric NAME.

Interval ID defaults to 20 (realtime) if supported, otherwise 300 (5m interval).

Examples:
  govc metric.sample host/cluster1/* cpu.usage.average
  govc metric.sample vm/* net.bytesTx.average net.bytesTx.average`
}

func (cmd *sample) Process(ctx context.Context) error {
	if err := cmd.PerformanceFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

type sampleResult struct {
	cmd    *sample
	m      *performance.Manager
	Sample []performance.EntityMetric
}

func (r *sampleResult) Write(w io.Writer) error {
	ctx := context.Background()
	cmd := r.cmd
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)

	counters, err := r.m.CounterInfoByName(ctx)
	if err != nil {
		return err
	}

	for i := range r.Sample {
		metric := r.Sample[i]

		var me mo.ManagedEntity
		_ = r.m.Properties(ctx, metric.Entity, []string{"name"}, &me)

		name := me.Name
		if cmd.d > 0 && len(name) > cmd.d {
			name = name[:cmd.d] + "*"
		}

		for _, v := range metric.Value {
			counter := counters[v.Name]
			units := counter.UnitInfo.GetElementDescription().Label

			instance := v.Instance
			if instance == "" {
				instance = "-"
			}

			t := ""
			if cmd.t {
				t = metric.SampleInfoCSV()
			}

			fmt.Fprintf(tw, "%s\t%s\t%s\t%v\t%s\t%s\n",
				name, instance, v.Name, t, v.ValueCSV(), units)
		}
	}

	return tw.Flush()
}

func (cmd *sample) Run(ctx context.Context, f *flag.FlagSet) error {
	m, err := cmd.Manager(ctx)
	if err != nil {
		return err
	}

	var paths []string
	var names []string

	byName, err := m.CounterInfoByName(ctx)
	if err != nil {
		return err
	}

	for _, arg := range f.Args() {
		if _, ok := byName[arg]; ok {
			names = append(names, arg)
		} else {
			paths = append(paths, arg)
		}
	}

	if len(paths) == 0 || len(names) == 0 {
		return flag.ErrHelp
	}

	objs, err := cmd.ManagedObjects(ctx, paths)
	if err != nil {
		return err
	}

	s, err := m.ProviderSummary(ctx, objs[0])
	if err != nil {
		return err
	}

	spec := types.PerfQuerySpec{
		Format:     string(types.PerfFormatNormal),
		MaxSample:  int32(cmd.n),
		MetricId:   []types.PerfMetricId{{Instance: cmd.instance}},
		IntervalId: cmd.Interval(s.RefreshRate),
	}

	sample, err := m.SampleByName(ctx, spec, names, objs)
	if err != nil {
		return err
	}

	result, err := m.ToMetricSeries(ctx, sample)
	if err != nil {
		return err
	}

	return cmd.WriteResult(&sampleResult{cmd, m, result})
}
