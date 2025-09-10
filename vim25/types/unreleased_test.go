// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package types_test

import (
	"context"
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

func TestPodVMInfo(t *testing.T) {
	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		host := simulator.Map(ctx).Any("HostSystem").(*simulator.HostSystem)

		host.Runtime.PodVMInfo = &types.HostRuntimeInfoPodVMInfo{
			HasPodVM: true,
			PodVMOverheadInfo: types.PodVMOverheadInfo{
				PodVMOverheadWithoutPageSharing: int32(50),
				PodVMOverheadWithPageSharing:    int32(25),
			},
		}

		var props mo.HostSystem
		pc := property.DefaultCollector(c)
		err := pc.RetrieveOne(ctx, host.Self, []string{"runtime"}, &props)
		if err != nil {
			t.Fatal(err)
		}

		if *props.Runtime.PodVMInfo != *host.Runtime.PodVMInfo {
			t.Errorf("%#v", props.Runtime.PodVMInfo)
		}
	})
}
