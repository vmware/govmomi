// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package performance_test

import (
	"context"
	"fmt"

	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/performance"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/types"
)

func ExampleManager_ToMetricSeries() {
	simulator.Run(func(ctx context.Context, c *vim25.Client) error {
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
			vm := object.NewVirtualMachine(c, metric.Entity)
			name, err := vm.ObjectName(ctx)
			if err != nil {
				return err
			}

			for _, v := range metric.Value {
				counter := counters[v.Name]
				units := counter.UnitInfo.GetElementDescription().Label

				instance := v.Instance
				if instance == "" {
					instance = "-"
				}

				if len(v.Value) != 0 && v.Name == "sys.uptime.latest" {
					fmt.Printf("%s\t%s\t%s\t%s\n", name, instance, v.Name, units)
					break
				}
			}
		}
		return nil
	})

	// Output:
	// DC0_H0_VM0	*	sys.uptime.latest	s
	// DC0_H0_VM1	*	sys.uptime.latest	s
	// DC0_C0_RP0_VM0	*	sys.uptime.latest	s
	// DC0_C0_RP0_VM1	*	sys.uptime.latest	s
}
