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

package simulator

import (
	"strconv"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/simulator/esx"
	"github.com/vmware/govmomi/simulator/vpx"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

var realtimeProviderSummary = types.PerfProviderSummary{
	CurrentSupported: true,
	SummarySupported: true,
	RefreshRate:      20,
}

var historicProviderSummary = types.PerfProviderSummary{
	CurrentSupported: false,
	SummarySupported: true,
	RefreshRate:      -1,
}

type PerformanceManager struct {
	mo.PerformanceManager
	vmMetrics        []types.PerfMetricId
	hostMetrics      []types.PerfMetricId
	rpMetrics        []types.PerfMetricId
	clusterMetrics   []types.PerfMetricId
	datastoreMetrics []types.PerfMetricId
	perfCounterIndex map[int32]types.PerfCounterInfo
}

func NewPerformanceManager(ref types.ManagedObjectReference) object.Reference {
	m := &PerformanceManager{}
	m.Self = ref
	if Map.IsESX() {
		m.PerfCounter = esx.PerfCounter[:]
		m.hostMetrics = esx.HostMetrics[:]
		m.vmMetrics = esx.VmMetrics[:]
		m.rpMetrics = esx.ResourcePoolMetrics[:]
	} else {
		m.PerfCounter = vpx.PerfCounter[:]
		m.hostMetrics = vpx.HostMetrics[:]
		m.vmMetrics = vpx.VmMetrics[:]
		m.rpMetrics = vpx.ResourcePoolMetrics[:]
		m.clusterMetrics = vpx.ClusterMetrics[:]
		m.datastoreMetrics = vpx.DatastoreMetrics[:]
	}
	m.perfCounterIndex = make(map[int32]types.PerfCounterInfo, len(m.PerfCounter))
	for _, p := range m.PerfCounter {
		m.perfCounterIndex[p.Key] = p
	}
	return m
}

func (p *PerformanceManager) QueryPerfCounter(ctx *Context, req *types.QueryPerfCounter) soap.HasFault {
	body := new(methods.QueryPerfCounterBody)
	body.Req = req
	body.Res.Returnval = make([]types.PerfCounterInfo, len(req.CounterId))
	for i, id := range req.CounterId {
		if info, ok := p.perfCounterIndex[id]; !ok {
			body.Fault_ = Fault("", &types.InvalidArgument{
				InvalidProperty: "CounterId",
			})
			return body
		} else {
			body.Res.Returnval[i] = info
		}
	}
	return body
}

func (p *PerformanceManager) QueryPerfProviderSummary(ctx *Context, req *types.QueryPerfProviderSummary) soap.HasFault {
	body := new(methods.QueryPerfProviderSummaryBody)
	body.Req = req
	body.Res = new(types.QueryPerfProviderSummaryResponse)

	// The entity must exist
	if Map.Get(req.Entity) == nil {
		body.Fault_ = Fault("", &types.InvalidArgument{
			InvalidProperty: "Entity",
		})
		return body
	}

	switch req.Entity.Type {
	case "VirtualMachine":
		fallthrough
	case "HostSystem":
		fallthrough
	case "ResourcePool":
		body.Res.Returnval = realtimeProviderSummary
	default:
		body.Res.Returnval = historicProviderSummary
	}
	body.Res.Returnval.Entity = req.Entity
	return body
}

func (p *PerformanceManager) buildAvailablePerfMetricsQueryResponse(ids []types.PerfMetricId, numCPU int, datastoreURL string) *types.QueryAvailablePerfMetricResponse {
	r := new(types.QueryAvailablePerfMetricResponse)
	r.Returnval = make([]types.PerfMetricId, 0, len(ids))
	for _, id := range ids {
		switch id.Instance {
		case "$cpu":
			for i := 0; i < numCPU; i++ {
				r.Returnval = append(r.Returnval, types.PerfMetricId{CounterId: id.CounterId, Instance: strconv.Itoa(i)})
			}
		case "$physDisk":
			r.Returnval = append(r.Returnval, types.PerfMetricId{CounterId: id.CounterId, Instance: datastoreURL})
		default:
			r.Returnval = append(r.Returnval, types.PerfMetricId{CounterId: id.CounterId, Instance: id.Instance})
		}
	}
	return r
}

func (p *PerformanceManager) queryAvailablePerfMetric(entity types.ManagedObjectReference, interval int32) *types.QueryAvailablePerfMetricResponse {
	var res *types.QueryAvailablePerfMetricResponse
	switch entity.Type {
	case "VirtualMachine":
		vm := Map.Get(entity).(*VirtualMachine)
		return p.buildAvailablePerfMetricsQueryResponse(p.vmMetrics, int(vm.Summary.Config.NumCpu), vm.Datastore[0].Value)
	case "HostSystem":
		host := Map.Get(entity).(*HostSystem)
		return p.buildAvailablePerfMetricsQueryResponse(p.hostMetrics, int(host.Hardware.CpuInfo.NumCpuThreads), host.Datastore[0].Value)
	case "ResourcePool":
		return p.buildAvailablePerfMetricsQueryResponse(p.rpMetrics, 0, "")
	case "ClusterComputeResource":
		if interval != 20 {
			res = p.buildAvailablePerfMetricsQueryResponse(p.clusterMetrics, 0, "")
		}
	case "Datastore":
		if interval != 20 {
			return p.buildAvailablePerfMetricsQueryResponse(p.datastoreMetrics, 0, "")
		}
	}

	// Don't know how to handle this. Return empty response.
	return new(types.QueryAvailablePerfMetricResponse)
}

func (p *PerformanceManager) QueryAvailablePerfMetric(ctx *Context, req *types.QueryAvailablePerfMetric) soap.HasFault {
	body := new(methods.QueryAvailablePerfMetricBody)
	body.Req = req
	body.Res = p.queryAvailablePerfMetric(req.Entity, req.IntervalId)

	spew.Dump(body)
	return body
}

func (p *PerformanceManager) QueryPerf(ctx *Context, req *types.QueryPerf) soap.HasFault {
	body := new(methods.QueryPerfBody)
	body.Req = req
	body.Res = new(types.QueryPerfResponse)
	body.Res.Returnval = make([]types.BasePerfEntityMetricBase, len(req.QuerySpec))

	for i, qs := range req.QuerySpec {
		metrics := new(types.PerfEntityMetric)
		metrics.Entity = qs.Entity

		// Generate metric series. Divide into n buckets of interval seconds
		interval := qs.IntervalId
		if interval == -1 || interval == 0 {
			interval = 20 // TODO: Determine from entity type
		}
		n := int32(qs.EndTime.Sub(*qs.StartTime).Seconds()) / interval
		if n > qs.MaxSample {
			n = qs.MaxSample
		}

		// Loop through each interval "tick"
		metrics.SampleInfo = make([]types.PerfSampleInfo, n)
		metrics.Value = make([]types.BasePerfMetricSeries, n)
		ts := *qs.StartTime
		for tick := int32(0); tick < n; tick++ {

			// Create sample metadata
			metrics.SampleInfo[tick] = types.PerfSampleInfo{Timestamp: ts, Interval: interval}

			// Create list of metrics for this tick
			series := &types.PerfMetricIntSeries{Value: make([]int64, len(qs.MetricId))}
			for j, id := range qs.MetricId {
				series.Value[j] = int64(0)
			}
			metrics.Value[tick] = series

			// Advance to next interval tick
			ts := ts.Add(time.Duration(interval) * time.Second)

		}
		body.Res.Returnval[i] = metrics
	}
}
