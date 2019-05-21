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

/*
This example program shows how to collect performance metrics from virtual machines using govmomi.
*/

package main

import (
	"context"
	"fmt"

	"github.com/vmware/govmomi/examples"
	"github.com/vmware/govmomi/performance"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/types"
)

func main() {
	examples.Run(func(ctx context.Context, c *vim25.Client) error {
		// Get virtual machines references
		m := view.NewManager(c)

		v, err := m.CreateContainerView(ctx, c.ServiceContent.RootFolder, nil, true)
		if err != nil {
			return err
		}

		defer v.Destroy(ctx)

		vmsRefs, err := v.Find(ctx, []string{"VirtualMachine"}, nil)
		if err != nil {
			return err
		}

		// Create a PerfManager
		perfManager := performance.NewManager(c)

		// Retrieve counters name list
		counters, err := perfManager.CounterInfoByName(ctx)
		if err != nil {
			return err
		}

		var names []string
		for name := range counters {
			names = append(names, name)
		}

		// Create PerfQuerySpec
		spec := types.PerfQuerySpec{
			MaxSample:  1,
			MetricId:   []types.PerfMetricId{{Instance: "*"}},
			IntervalId: 300,
		}

		// Query metrics
		sample, err := perfManager.SampleByName(ctx, spec, names, vmsRefs)
		if err != nil {
			return err
		}

		result, err := perfManager.ToMetricSeries(ctx, sample)
		if err != nil {
			return err
		}

		// Read result
		for _, metric := range result {
			name := metric.Entity

			for _, v := range metric.Value {
				counter := counters[v.Name]
				units := counter.UnitInfo.GetElementDescription().Label

				instance := v.Instance
				if instance == "" {
					instance = "-"
				}

				if len(v.Value) != 0 {
					fmt.Printf("%s\t%s\t%s\t%s\t%s\n",
						name, instance, v.Name, v.ValueCSV(), units)
				}
			}
		}
		return nil
	})
}
