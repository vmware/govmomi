// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package types_test

import (
	"context"
	"reflect"
	"strings"
	"testing"

	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vim25/xml"
)

func TestPodVMInfo(t *testing.T) {
	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		host := simulator.Map(ctx).Any("HostSystem").(*simulator.HostSystem)

		host.Runtime.PodVMInfo = &types.PodVMInfo{
			HasPageSharingPodVM: true,
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

// TestExtensionCompatibilityConstraint verifies the read-back property
// config.extensionCompatibilityConstraint round-trips through the
// PropertyCollector (as config.managedBy does), and that the request-side
// write/bypass fields on ConfigSpec and RelocateSpec serialize.
func TestExtensionCompatibilityConstraint(t *testing.T) {
	set := &types.VirtualMachineExtensionCompatibilityConstraintSet{
		Constraint: []types.VirtualMachineExtensionCompatibilityConstraint{
			{
				ConstraintName: "pool-invariant",
				ConstraintType: string(types.VirtualMachineExtensionCompatibilityConstraintTypePOOL),
				ConstraintKind: string(types.VirtualMachineExtensionCompatibilityConstraintKindINVARIANT),
			},
		},
	}

	// Read-back path via PropertyCollector.
	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		vm := simulator.Map(ctx).Any("VirtualMachine").(*simulator.VirtualMachine)
		vm.Config.ExtensionCompatibilityConstraint = set

		var props mo.VirtualMachine
		pc := property.DefaultCollector(c)
		if err := pc.RetrieveOne(ctx, vm.Self, []string{"config"}, &props); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(props.Config.ExtensionCompatibilityConstraint, set) {
			t.Errorf("config.extensionCompatibilityConstraint: %#v", props.Config.ExtensionCompatibilityConstraint)
		}
	})

	// Request-side write + bypass fields (vcsim does not persist these); assert
	// they serialize, and in VMODL sequence order: vpxd parses a DataObject
	// with a forward cursor over the property list, so an element emitted
	// after a later-declared property (e.g. vmProfile) is rejected with
	// "Unexpected element tag".
	skip := true
	b, err := xml.Marshal(types.VirtualMachineConfigSpec{
		ManagedBy:                        &types.ManagedByInfo{ExtensionKey: "ext", Type: "t"},
		ExtensionCompatibilityConstraint: set,
		SkipExtensionCompatibilityChecks: &skip,
		MemoryReservationLockedToMax:     &skip,
		VmProfile:                        []types.BaseVirtualMachineProfileSpec{&types.VirtualMachineDefaultProfileSpec{}},
	})
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"pool-invariant", "POOL", "INVARIANT"} {
		if !strings.Contains(string(b), want) {
			t.Errorf("expected %q in marshaled ConfigSpec:\n%s", want, b)
		}
	}
	prev := -1
	for _, name := range []string{
		"managedBy", "extensionCompatibilityConstraint", "skipExtensionCompatibilityChecks",
		"memoryReservationLockedToMax", "vmProfile",
	} {
		i := strings.Index(string(b), "<"+name)
		if i < 0 {
			t.Fatalf("expected %q in marshaled ConfigSpec:\n%s", name, b)
		}
		if i <= prev {
			t.Errorf("%q emitted out of VMODL order in marshaled ConfigSpec:\n%s", name, b)
		}
		prev = i
	}

	if b, err = xml.Marshal(types.VirtualMachineRelocateSpec{SkipExtensionCompatibilityChecks: &skip}); err != nil {
		t.Fatal(err)
	} else if !strings.Contains(string(b), "skipExtensionCompatibilityChecks") {
		t.Errorf("expected skipExtensionCompatibilityChecks in RelocateSpec:\n%s", b)
	}
}
