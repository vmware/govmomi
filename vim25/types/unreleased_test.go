// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package types_test

import (
	"bytes"
	"context"
	"encoding/xml"
	"reflect"
	"testing"

	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

func TestPodVMOverheadInfo(t *testing.T) {
	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		host := simulator.Map(ctx).Any("HostSystem").(*simulator.HostSystem)

		host.Capability.PodVMOverheadInfo = &types.PodVMOverheadInfo{
			CrxPageSharingSupported:         true,
			PodVMOverheadWithoutPageSharing: int32(42),
			PodVMOverheadWithPageSharing:    int32(53),
		}

		var props mo.HostSystem
		pc := property.DefaultCollector(c)
		err := pc.RetrieveOne(ctx, host.Self, []string{"capability"}, &props)
		if err != nil {
			t.Fatal(err)
		}

		if *props.Capability.PodVMOverheadInfo != *host.Capability.PodVMOverheadInfo {
			t.Errorf("%#v", props.Capability.PodVMOverheadInfo)
		}
	})
}

func TestTypeClusterClusterInitialPlacementActionEx(t *testing.T) {
	var ok bool

	// Register the original base type for "ClusterClusterInitialPlacementAction".
	// It simulates the default SDK behavior before override.
	types.Add("ClusterClusterInitialPlacementAction", reflect.TypeOf((*types.ClusterClusterInitialPlacementAction)(nil)).Elem())
	fn := types.TypeFunc()

	// Lookup an unknown type - should return ok==false.
	_, ok = fn("unknown")
	if ok {
		t.Errorf("Expected ok==false")
	}

	// Lookup the registered base type ClusterClusterInitialPlacementAction - should return ok==true.
	actual, ok := fn("ClusterClusterInitialPlacementAction")
	if !ok {
		t.Errorf("Expected ok==true")
	}

	expected := reflect.TypeOf(types.ClusterClusterInitialPlacementAction{})

	// Validate that the type lookup matches the base type.
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %#v, actual: %#v", expected, actual)
	}

	// Override the registered type with our extended struct ClusterClusterInitialPlacementActionEx.
	types.Add("ClusterClusterInitialPlacementAction", reflect.TypeOf((*types.ClusterClusterInitialPlacementActionEx)(nil)).Elem())

	// Lookup the same name again - should now return the extended type.
	actual, ok = fn("ClusterClusterInitialPlacementAction")
	if !ok {
		t.Errorf("Expected ok==true")
	}

	// Expected type is now the extended struct.
	expected = reflect.TypeOf(types.ClusterClusterInitialPlacementActionEx{})
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %#v, actual: %#v", expected, actual)
	}
}

func TestTypeResourcePoolRuntimeInfoEx(t *testing.T) {
	fn := types.TypeFunc()

	// By default, the registry should return the BASE type.
	actual, ok := fn("ResourcePoolRuntimeInfo")
	if !ok {
		t.Fatalf("expected ok for ResourcePoolRuntimeInfo")
	}
	if want := reflect.TypeOf(types.ResourcePoolRuntimeInfo{}); !reflect.DeepEqual(want, actual) {
		t.Fatalf("default mapping should be base; want=%#v got=%#v", want, actual)
	}

	// override: map the name to Ex and verify.
	prev := actual
	types.Add("ResourcePoolRuntimeInfo", reflect.TypeOf(types.ResourcePoolRuntimeInfoEx{}))
	defer types.Add("ResourcePoolRuntimeInfo", prev)

	actual2, ok := fn("ResourcePoolRuntimeInfo")
	if !ok {
		t.Fatalf("expected ok for ResourcePoolRuntimeInfo after override")
	}
	if want := reflect.TypeOf(types.ResourcePoolRuntimeInfoEx{}); !reflect.DeepEqual(want, actual2) {
		t.Fatalf("override mapping failed; want=%#v got=%#v", want, actual2)
	}
}

func TestResourcePoolRuntimeInfoEx_XMLOmitempty(t *testing.T) {
	// Case 1: VmRp is nil → <vmRp> MUST be omitted due to ,omitempty on the wrapper pointer.
	exNil := &types.ResourcePoolRuntimeInfoEx{} // VmRp == nil
	bNil, err := xml.Marshal(exNil)
	if err != nil {
		t.Fatalf("marshal (nil VmRp) failed: %v", err)
	}
	if bytes.Contains(bNil, []byte("<vmRp")) {
		t.Fatalf("expected no <vmRp> element when VmRp is nil, got: %s", string(bNil))
	}

	// Case 2: VmRp has one item → <vmRp> MUST be present with child item(s).
	exOne := &types.ResourcePoolRuntimeInfoEx{
		VmRp: &types.ArrayOfResourcePoolVmResourceProfileUsage{
			ResourcePoolVmResourceProfileUsage: []types.ResourcePoolVmResourceProfileUsage{
				{
					Id:                           "small-vm",
					ReservedForPool:              4,
					ReservationUsedForVms:        0,
					ReservationUsedForChildPools: 0,
				},
			},
		},
	}
	bOne, err := xml.Marshal(exOne)
	if err != nil {
		t.Fatalf("marshal (one VmRp) failed: %v", err)
	}
	if !bytes.Contains(bOne, []byte("<vmRp")) {
		t.Fatalf("expected <vmRp> element when VmRp has items, got: %s", string(bOne))
	}
	if !bytes.Contains(bOne, []byte("<ResourcePoolVmResourceProfileUsage")) {
		t.Fatalf("expected ResourcePoolVmResourceProfileUsage item, got: %s", string(bOne))
	}

	// Case 3: empty wrapper (non-nil, zero items) → <vmRp/> WILL be present.
	exEmpty := &types.ResourcePoolRuntimeInfoEx{
		VmRp: &types.ArrayOfResourcePoolVmResourceProfileUsage{
			ResourcePoolVmResourceProfileUsage: nil, // or: []types.ResourcePoolVmResourceProfileUsage{}
		},
	}
	bEmpty, err := xml.Marshal(exEmpty)
	if err != nil {
		t.Fatalf("marshal (empty wrapper) failed: %v", err)
	}
	if !bytes.Contains(bEmpty, []byte("<vmRp")) {
		t.Fatalf("expected <vmRp> element when wrapper is non-nil (even if empty), got: %s", string(bEmpty))
	}
}
