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

	"github.com/davecgh/go-spew/spew"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/simulator/esx"
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
	vmMetrics   []types.PerfMetricId
	hostMetrics []types.PerfMetricId
}

func NewPerformanceManager(ref types.ManagedObjectReference) object.Reference {
	m := &PerformanceManager{}
	m.Self = ref
	m.PerfCounter = esx.PerfCounter[:]
	m.hostMetrics = esx.HostMetrics[:]
	m.vmMetrics = esx.VmMetrics[:]
	return m
}

func (p *PerformanceManager) QueryPerfProviderSummary(ctx *Context, req *types.QueryPerfProviderSummary) soap.HasFault {
	body := new(methods.QueryPerfProviderSummaryBody)
	body.Req = req
	body.Res = new(types.QueryPerfProviderSummaryResponse)
	switch req.Entity.Type {
	case "VirtualMachine":
		body.Res.Returnval = realtimeProviderSummary
	case "HostSystem":
		body.Res.Returnval = realtimeProviderSummary
	default:
		body.Res.Returnval = historicProviderSummary
	}
	body.Res.Returnval.Entity = req.Entity
	return body
}

func (p *PerformanceManager) buildPerfMetricsQueryResponse(ids []types.PerfMetricId, numCPU int, datastoreURL string) *types.QueryAvailablePerfMetricResponse {
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

func (p *PerformanceManager) QueryAvailablePerfMetric(ctx *Context, req *types.QueryAvailablePerfMetric) soap.HasFault {
	body := new(methods.QueryAvailablePerfMetricBody)
	body.Req = req
	body.Res = new(types.QueryAvailablePerfMetricResponse)

	switch req.Entity.Type {
	case "VirtualMachine":
		vm := Map.Get(req.Entity).(*VirtualMachine)
		body.Res = p.buildPerfMetricsQueryResponse(p.vmMetrics, int(vm.Summary.Config.NumCpu), vm.Datastore[0].Value)
	case "HostSystem":
		host := Map.Get(req.Entity).(*HostSystem)
		spew.Dump(host)
		body.Res = p.buildPerfMetricsQueryResponse(p.hostMetrics, int(host.Hardware.CpuInfo.NumCpuThreads), host.Datastore[0].Value)
		/*case "ClusterComputeResource":
			if req.IntervalId != 20 {
				body.Res = p.buildPerfMetricsQueryResponse(clusterMetrics)
			}
		case "Datastore":
			if req.IntervalId != 20 {
				body.Res = p.buildPerfMetricsQueryResponse(clusterMetrics)
			}*/
	}
	spew.Dump(body)
	return body
}
