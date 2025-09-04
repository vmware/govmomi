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

	// Expected type is the base struct (before override).
	expected := reflect.TypeOf(types.ClusterClusterInitialPlacementAction{})

	// Validate that the type lookup matches the base type.
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %#v, actual: %#v", expected, actual)
	}

	// Override the registered type with our extended struct ClusterClusterInitialPlacementActionEx.
	types.Add("ClusterClusterInitialPlacementAction", reflect.TypeOf((*types.ClusterClusterInitialPlacementActionEx)(nil)).Elem())

	var actual2 reflect.Type
	// Lookup the same name again - should now return the extended type.
	// Now `actual` refers to the reflect.Type of the extended struct.
	actual2, ok = fn("ClusterClusterInitialPlacementAction")
	if !ok {
		t.Errorf("Expected ok==true")
	}

	// Expected type is now the extended struct (after override).
	expected2 := reflect.TypeOf(types.ClusterClusterInitialPlacementActionEx{})
	if !reflect.DeepEqual(expected2, actual2) {
		t.Errorf("Expected: %#v, actual: %#v", expected2, actual2)
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

func TestResourcePoolRuntimeInfoEx_XML_vmRp(t *testing.T) {
	// Case 1: VmRp is nil – do NOT assert absence of <vmRp>, just ensure no items.
	exNil := &types.ResourcePoolRuntimeInfoEx{}
	bNil, err := xml.Marshal(exNil)
	if err != nil {
		t.Fatalf("marshal (nil VmRp) failed: %v", err)
	}
	if bytes.Contains(bNil, []byte("<ResourcePoolVmResourceProfileUsage")) {
		t.Fatalf("unexpected item when VmRp is nil: %s", string(bNil))
	}

	// Case 2: VmRp empty slice – same: allow <vmRp>, but no items.
	exEmpty := &types.ResourcePoolRuntimeInfoEx{VmRp: []types.ResourcePoolVmResourceProfileUsage{}}
	bEmpty, err := xml.Marshal(exEmpty)
	if err != nil {
		t.Fatalf("marshal (empty VmRp) failed: %v", err)
	}
	if bytes.Contains(bEmpty, []byte("<ResourcePoolVmResourceProfileUsage")) {
		t.Fatalf("unexpected item when VmRp is empty: %s", string(bEmpty))
	}

	// Case 3: VmRp with one item – expect both <vmRp> and an item.
	exOne := &types.ResourcePoolRuntimeInfoEx{
		VmRp: []types.ResourcePoolVmResourceProfileUsage{
			{Id: "small-vm", ReservedForPool: 2, ReservationUsedForVms: 0, ReservationUsedForChildPools: 0},
		},
	}
	bOne, err := xml.Marshal(exOne)
	if err != nil {
		t.Fatalf("marshal (one VmRp) failed: %v", err)
	}
	// <vmRp> must be present in encoded xml.
	if !bytes.Contains(bOne, []byte("<vmRp")) {
		t.Fatalf("expected <vmRp> when VmRp has items, got: %s", string(bOne))
	}
	// <ResourcePoolVmResourceProfileUsage in encoded xml.
	if !bytes.Contains(bOne, []byte("<ResourcePoolVmResourceProfileUsage")) {
		t.Fatalf("expected ResourcePoolVmResourceProfileUsage item, got: %s", string(bOne))
	}
}
