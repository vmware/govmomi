// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/vmware/govmomi/performance"
	"github.com/vmware/govmomi/simulator/esx"
	"github.com/vmware/govmomi/simulator/vpx"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

func testMetricsConsistency(counterInfo []types.PerfCounterInfo, ids [][]types.PerfMetricId) error {

	// Build a lookup table for speed and convenience
	lookup := make(map[int32]bool, len(counterInfo))
	for _, pc := range counterInfo {
		lookup[pc.Key] = true
	}

	// Check metric ids against map
	for _, list := range ids {
		for _, id := range list {
			if _, ok := lookup[id.CounterId]; !ok {
				return fmt.Errorf("Counter with ID %d not found in PerfCounter", id.CounterId)
			}
		}
	}
	return nil
}

func TestMetricsConsistency(t *testing.T) {
	esxIds := [][]types.PerfMetricId{esx.VmMetrics, esx.HostMetrics, esx.ResourcePoolMetrics}
	vpxIds := [][]types.PerfMetricId{vpx.VmMetrics, vpx.HostMetrics, vpx.ClusterMetrics,
		vpx.DatastoreMetrics, vpx.ResourcePoolMetrics}
	if err := testMetricsConsistency(esx.PerfCounter, esxIds); err != nil {
		t.Fatal(err)
	}
	if err := testMetricsConsistency(vpx.PerfCounter, vpxIds); err != nil {
		t.Fatal(err)
	}
}

func checkDuplicates(ids []types.PerfMetricId) error {
	m := make(map[string]bool, len(ids))
	for _, id := range ids {
		k := fmt.Sprintf("%d|%s", id.CounterId, id.Instance)
		if _, ok := m[k]; ok {
			return fmt.Errorf("Duplicate metric key: %s", k)
		}
		m[k] = true
	}
	return nil
}

func TestMetricsDuplicates(t *testing.T) {
	if err := checkDuplicates(esx.VmMetrics); err != nil {
		t.Fatal(err)
	}
	if err := checkDuplicates(esx.HostMetrics); err != nil {
		t.Fatal(err)
	}
	if err := checkDuplicates(vpx.VmMetrics); err != nil {
		t.Fatal(err)
	}
	if err := checkDuplicates(vpx.HostMetrics); err != nil {
		t.Fatal(err)
	}
	if err := checkDuplicates(vpx.ClusterMetrics); err != nil {
		t.Fatal(err)
	}
	if err := checkDuplicates(vpx.DatastoreMetrics); err != nil {
		t.Fatal(err)
	}
}

func TestQueryProviderSummary(t *testing.T) {
	m := VPX()

	err := m.Create()
	if err != nil {
		t.Fatal(err)
	}

	defer m.Remove()

	c := m.Service.client()

	p := performance.NewManager(c)
	ctx := m.Service.Context

	vm := ctx.Map.Any("VirtualMachine").(*VirtualMachine)
	if info, err := p.ProviderSummary(ctx, vm.Reference()); err != nil {
		t.Fatal(err)
	} else {
		if info.RefreshRate != 20 {
			t.Fatalf("VM wefresh rate is %d, should be 20", info.RefreshRate)
		}
	}

	host := ctx.Map.Any("HostSystem").(*HostSystem)
	if info, err := p.ProviderSummary(ctx, host.Reference()); err != nil {
		t.Fatal(err)
	} else {
		if info.RefreshRate != 20 {
			t.Fatalf("Host refresh rate is %d, should be 20", info.RefreshRate)
		}
	}

	pool := ctx.Map.Any("ResourcePool").(*ResourcePool)
	if info, err := p.ProviderSummary(ctx, pool.Reference()); err != nil {
		t.Fatal(err)
	} else {
		if info.RefreshRate != 20 {
			t.Fatalf("ResourcePool refresh rate is %d, should be 20", info.RefreshRate)
		}
	}

	cluster := ctx.Map.Any("ClusterComputeResource").(*ClusterComputeResource)
	if info, err := p.ProviderSummary(ctx, cluster.Reference()); err != nil {
		t.Fatal(err)
	} else {
		if info.RefreshRate != -1 {
			t.Fatalf("Cluster refresh rate is %d, should be -1", info.RefreshRate)
		}
	}

	datastore := ctx.Map.Any("Datastore").(*Datastore)
	if info, err := p.ProviderSummary(ctx, datastore.Reference()); err != nil {
		t.Fatal(err)
	} else {
		if info.RefreshRate != -1 {
			t.Fatalf("Datastore refresh rate is %d, should be -1", info.RefreshRate)
		}
	}

	nonExistent := types.ManagedObjectReference{
		Type:  "Not a valid type",
		Value: "This object doesn't exist",
	}
	if _, err := p.ProviderSummary(ctx, nonExistent); err == nil {
		t.Fatal("This should have failed (nonexistent object)")
	}
}

func TestQueryAvailablePerfMetric(t *testing.T) {
	m := VPX()

	err := m.Create()
	if err != nil {
		t.Fatal(err)
	}

	defer m.Remove()

	c := m.Service.client()
	p := performance.NewManager(c)
	ctx := m.Service.Context

	vm := ctx.Map.Any("VirtualMachine").(*VirtualMachine)
	if info, err := p.AvailableMetric(ctx, vm.Reference(), 20); err != nil {
		t.Fatal(err)
	} else {
		if len(info) == 0 {
			t.Fatal("Expected non-empty list of vm")
		}
	}

	vm.Datastore = nil // e.g. vCLS VMs have no Datastore
	if info, err := p.AvailableMetric(ctx, vm.Reference(), 20); err != nil {
		t.Fatal(err)
	} else {
		if len(info) == 0 {
			t.Fatal("Expected non-empty list of vm")
		}
	}

	host := ctx.Map.Any("HostSystem").(*HostSystem)
	if info, err := p.AvailableMetric(ctx, host.Reference(), 20); err != nil {
		t.Fatal(err)
	} else {
		if len(info) == 0 {
			t.Fatal("Expected non-empty list of host")
		}
		var ids []int32
		for i := range info {
			ids = append(ids, info[i].CounterId)
		}
		perf, err := p.QueryCounter(ctx, ids)
		if err != nil {
			t.Fatal(err)
		}
		if len(perf) != len(ids) {
			t.Errorf("%d counters", len(perf))
		}
	}

	pool := ctx.Map.Any("ResourcePool").(*ResourcePool)
	if info, err := p.AvailableMetric(ctx, pool.Reference(), 20); err != nil {
		t.Fatal(err)
	} else {
		if len(info) == 0 {
			t.Fatal("Expected non-empty list of resource pool")
		}
	}

	cluster := ctx.Map.Any("ClusterComputeResource").(*ClusterComputeResource)
	if info, err := p.AvailableMetric(ctx, cluster.Reference(), 300); err != nil {
		t.Fatal(err)
	} else {
		if len(info) == 0 {
			t.Fatal("Expected non-empty list of clusters")
		}
	}

	if info, err := p.AvailableMetric(ctx, cluster.Reference(), 20); err != nil {
		t.Fatal(err)
	} else {
		if len(info) != 0 {
			t.Fatal("Expected empty list of clusters")
		}
	}

	ds := ctx.Map.Any("Datastore").(*Datastore)
	if info, err := p.AvailableMetric(ctx, ds.Reference(), 300); err != nil {
		t.Fatal(err)
	} else {
		if len(info) == 0 {
			t.Fatal("Expected non-empty list of datastores")
		}
	}

	if info, err := p.AvailableMetric(ctx, ds.Reference(), 20); err != nil {
		t.Fatal(err)
	} else {
		if len(info) != 0 {
			t.Fatal("Expected empty list of datastores")
		}
	}

	dc := ctx.Map.Any("Datacenter").(*Datacenter)
	if info, err := p.AvailableMetric(ctx, dc.Reference(), 300); err != nil {
		t.Fatal(err)
	} else {
		if len(info) == 0 {
			t.Fatal("Expected non-empty list of datacenters")
		}
	}

	if info, err := p.AvailableMetric(ctx, dc.Reference(), 20); err != nil {
		t.Fatal(err)
	} else {
		if len(info) != 0 {
			t.Fatal("Expected empty list of datacenters")
		}
	}

}

func testPerfQuery(ctx context.Context, m *Model, e mo.Entity, interval int32, maxSample int32) error {
	c := m.Service.client()

	p := performance.NewManager(c)

	// Single metric, single VM
	//
	qs := []types.PerfQuerySpec{
		{
			MaxSample:  maxSample,
			IntervalId: interval,
			MetricId:   []types.PerfMetricId{{CounterId: 1, Instance: ""}},
			Entity:     e.Reference(),
		},
	}
	result, err := p.Query(ctx, qs)
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return errors.New("Empty result set")
	}
	ms, err := p.ToMetricSeries(ctx, result)
	if err != nil {
		return err
	}
	if len(ms) == 0 {
		return errors.New("Empty metric series")
	}
	for _, em := range ms {
		if len(em.SampleInfo) == 0 {
			return errors.New("Empty SampleInfo")
		}
	}

	return nil
}

func testPerfQueryCSV(ctx context.Context, m *Model, e mo.Entity, interval int32, maxSample int32) error {
	c := m.Service.client()

	p := performance.NewManager(c)

	// Single metric, single VM
	//
	qs := []types.PerfQuerySpec{
		{
			MaxSample:  maxSample,
			IntervalId: interval,
			MetricId:   []types.PerfMetricId{{CounterId: 1, Instance: ""}},
			Entity:     e.Reference(),
			Format:     string(types.PerfFormatCsv),
		},
	}
	series, err := p.Query(ctx, qs)
	if err != nil {
		return err
	}
	if len(series) == 0 {
		return errors.New("Empty result set")
	}
	for i := range series {
		s, ok := series[i].(*types.PerfEntityMetricCSV)
		if !ok {
			panic(fmt.Errorf("expected type %T, got: %T", s, series[i]))
		}
		if len(s.SampleInfoCSV) == 0 {
			return errors.New("Empty SampleInfoCSV")
		}
		if len(strings.Split(s.SampleInfoCSV, ",")) == 0 {
			return errors.New("SampleInfoCSV not in CSV format")
		}
		for _, v := range s.Value {
			if len(v.Value) == 0 {
				return errors.New("Empty PerfEntityMetricCSV.Value")
			}
			if len(strings.Split(v.Value, ",")) == 0 {
				return errors.New("PerfEntityMetricCSV.Value not in CSV format")
			}
		}
	}

	return nil
}

func TestQueryPerf(t *testing.T) {
	m := VPX()

	err := m.Create()
	if err != nil {
		t.Fatal(err)
	}

	defer m.Remove()

	ctx := m.Service.Context

	for _, maxSample := range []int32{4, 0} {
		if err := testPerfQuery(ctx, m, ctx.Map.Any("VirtualMachine"), 20, maxSample); err != nil {
			t.Fatal(err)
		}
		if err := testPerfQuery(ctx, m, ctx.Map.Any("HostSystem"), 20, maxSample); err != nil {
			t.Fatal(err)
		}
		if err := testPerfQuery(ctx, m, ctx.Map.Any("ClusterComputeResource"), 300, maxSample); err != nil {
			t.Fatal(err)
		}
		if err := testPerfQuery(ctx, m, ctx.Map.Any("Datastore"), 300, maxSample); err != nil {
			t.Fatal(err)
		}
		if err := testPerfQuery(ctx, m, ctx.Map.Any("Datacenter"), 300, maxSample); err != nil {
			t.Fatal(err)
		}
		if err := testPerfQuery(ctx, m, ctx.Map.Any("ResourcePool"), 300, maxSample); err != nil {
			t.Fatal(err)
		}

		//csv format
		if err := testPerfQueryCSV(ctx, m, ctx.Map.Any("VirtualMachine"), 20, maxSample); err != nil {
			t.Fatal(err)
		}
		if err := testPerfQueryCSV(ctx, m, ctx.Map.Any("HostSystem"), 20, maxSample); err != nil {
			t.Fatal(err)
		}
		if err := testPerfQueryCSV(ctx, m, ctx.Map.Any("ClusterComputeResource"), 300, maxSample); err != nil {
			t.Fatal(err)
		}
		if err := testPerfQueryCSV(ctx, m, ctx.Map.Any("Datastore"), 300, maxSample); err != nil {
			t.Fatal(err)
		}
		if err := testPerfQueryCSV(ctx, m, ctx.Map.Any("Datacenter"), 300, maxSample); err != nil {
			t.Fatal(err)
		}
		if err := testPerfQueryCSV(ctx, m, ctx.Map.Any("ResourcePool"), 300, maxSample); err != nil {
			t.Fatal(err)
		}
	}
}

// TestQueryPerfSampleInfoOrder asserts that QueryPerf returns PerfSampleInfo in
// ascending (oldest first) chronological order, which is what a real vCenter does.
// Returning them newest-first breaks clients that follow the vSphere convention of
// treating SampleInfo[len-1] as the most recent sample.
func TestQueryPerfSampleInfoOrder(t *testing.T) {
	m := VPX()

	if err := m.Create(); err != nil {
		t.Fatal(err)
	}
	defer m.Remove()

	ctx := m.Service.Context
	p := performance.NewManager(m.Service.client())

	for _, test := range []struct {
		kind     string
		interval int32
	}{
		{"VirtualMachine", 20},
		{"VirtualMachine", 300},
		{"HostSystem", 20},
		{"ClusterComputeResource", 300},
	} {
		entity := ctx.Map.Any(test.kind)

		result, err := p.Query(ctx, []types.PerfQuerySpec{{
			IntervalId: test.interval,
			MetricId:   []types.PerfMetricId{{CounterId: 1}},
			Entity:     entity.Reference(),
		}})
		if err != nil {
			t.Fatalf("%s interval=%d: %s", test.kind, test.interval, err)
		}

		series, err := p.ToMetricSeries(ctx, result)
		if err != nil {
			t.Fatalf("%s interval=%d: %s", test.kind, test.interval, err)
		}
		if len(series) == 0 {
			t.Fatalf("%s interval=%d: empty metric series", test.kind, test.interval)
		}

		for _, em := range series {
			if len(em.SampleInfo) < 2 {
				t.Fatalf("%s interval=%d: need >1 sample to check ordering, got %d",
					test.kind, test.interval, len(em.SampleInfo))
			}

			for i := 1; i < len(em.SampleInfo); i++ {
				prev, cur := em.SampleInfo[i-1].Timestamp, em.SampleInfo[i].Timestamp
				if !cur.After(prev) {
					t.Errorf("%s interval=%d: SampleInfo not ascending at index %d: %s then %s",
						test.kind, test.interval, i, prev, cur)
				}
			}

			// The interval spacing should match what was requested.
			gap := em.SampleInfo[1].Timestamp.Sub(em.SampleInfo[0].Timestamp)
			if want := time.Duration(test.interval) * time.Second; gap != want {
				t.Errorf("%s interval=%d: sample spacing = %s, want %s",
					test.kind, test.interval, gap, want)
			}

			// Values must line up 1:1 with the sample timestamps.
			for _, v := range em.Value {
				if len(v.Value) != len(em.SampleInfo) {
					t.Errorf("%s interval=%d: %d values for %d samples",
						test.kind, test.interval, len(v.Value), len(em.SampleInfo))
				}
			}
		}
	}
}
