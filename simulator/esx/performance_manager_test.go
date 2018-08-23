/*
Copyright (c) 2018 VMware, Inc. All Rights Reserved.

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
package esx

import (
	"fmt"
	"testing"

	"github.com/vmware/govmomi/vim25/types"
)

func checkMetricConsistency(lookup map[int32]bool, ids []types.PerfMetricId) error {
	for _, id := range ids {
		if _, ok := lookup[id.CounterId]; !ok {
			return fmt.Errorf("Counter with ID %d not found in PerfCounter", id.CounterId)
		}
	}
	return nil
}

func TestMetricsConsistency(t *testing.T) {

	// Build a lookup table for speed and convenience
	lookup := make(map[int32]bool, len(PerfCounter))
	for _, pc := range PerfCounter {
		lookup[pc.Key] = true
	}

	// Check metric ids against map
	if err := checkMetricConsistency(lookup, VmMetrics); err != nil {
		t.Fatal(err)
	}
	if err := checkMetricConsistency(lookup, HostMetrics); err != nil {
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
	if err := checkDuplicates(VmMetrics); err != nil {
		t.Fatal(err)
	}
	if err := checkDuplicates(HostMetrics); err != nil {
		t.Fatal(err)
	}
}
