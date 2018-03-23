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
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/simulator/esx"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/methods"
	"time"
	"math/rand"
)

type PerformanceManager struct {
	mo.PerformanceManager
}

func NewPerformanceManager(ref types.ManagedObjectReference) object.Reference {
	m := &PerformanceManager{}
	m.Self = ref
	m.PerfCounter = esx.PerfCounter
	Map.Put(m).Reference()
	return m
}

func (m *PerformanceManager) QueryPerfCounter(ref *types.QueryPerfCounter) soap.HasFault {
	// just return everything, for now
	return &methods.QueryPerfCounterBody{
		Res: &types.QueryPerfCounterResponse{
			Returnval: esx.PerfCounter,
		},
	}
}

func (m *PerformanceManager) QueryPerfProviderSummary(ref *types.QueryPerfProviderSummary) soap.HasFault {
	return &methods.QueryPerfProviderSummaryBody{
		Res: &types.QueryPerfProviderSummaryResponse{
			Returnval: types.PerfProviderSummary{
				Entity:           ref.Entity,
				CurrentSupported: true,
				SummarySupported: true,
				RefreshRate:      20,
			},
		},
	}
}

func (m *PerformanceManager) QueryAvailablePerfMetric(ref *types.QueryAvailablePerfMetric) soap.HasFault {
	return &methods.QueryAvailablePerfMetricBody{
		Res: &types.QueryAvailablePerfMetricResponse{
			Returnval: esx.AvailablePerfMetric,
		},
	}
}

func (m *PerformanceManager) QueryPerf(ref *types.QueryPerf) soap.HasFault {
	// convert to map for easy fetching
	var perfMap = make(map[int32]types.PerfMetricId)
	for i := 0; i < len(esx.AvailablePerfMetric); i++ {
		perfMap[esx.AvailablePerfMetric[i].CounterId] = esx.AvailablePerfMetric[i]
	}
	var ret = []types.BasePerfEntityMetricBase{}

	// fetch perf metrics with matching IDs
	for i := 0; i < len(ref.QuerySpec); i++ {
		// for each entity
		var added = &types.PerfEntityMetric{
			SampleInfo: []types.PerfSampleInfo{{
				Interval:  0,
				Timestamp: time.Now(),
			},},
			Value: []types.BasePerfMetricSeries{},
			PerfEntityMetricBase: types.PerfEntityMetricBase{
				Entity: ref.QuerySpec[i].Entity,
				DynamicData: ref.QuerySpec[i].DynamicData,
			},
		}
		ret = append(ret, added)

		for j := 0; j < len(ref.QuerySpec[i].MetricId); j++ {
			id, present := perfMap[ref.QuerySpec[i].MetricId[j].CounterId]
			if present {
				metrics := []int64{}
				count := rand.Intn(100)
				for i := 0 ; i < count ; i++ {
					metrics = append(metrics, rand.Int63())
				}
				added.Value = append(added.Value, &types.PerfMetricIntSeries{
					Value: metrics,
					PerfMetricSeries: types.PerfMetricSeries {
						Id: id,
					},
				})
			}
		}
	}

	return &methods.QueryPerfBody{
		Res: &types.QueryPerfResponse{
			Returnval: ret,
		},
	}
}
