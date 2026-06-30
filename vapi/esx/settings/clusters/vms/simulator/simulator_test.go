// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"context"
	"testing"

	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vapi/esx/settings/clusters/vms"
	"github.com/vmware/govmomi/vapi/rest"
	_ "github.com/vmware/govmomi/vapi/simulator"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/types"
)

// testSetup returns a logged-in vms.Manager and a cluster MOR from the simulator.
func testSetup(ctx context.Context, vc *vim25.Client, t *testing.T) (*vms.Manager, types.ManagedObjectReference) {
	t.Helper()

	c := rest.NewClient(vc)
	if err := c.Login(ctx, simulator.DefaultLogin); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = c.Logout(ctx) })

	m := &vms.Manager{Client: c}

	ccr := simulator.Map(ctx).Any("ClusterComputeResource")
	if ccr == nil {
		t.Fatal("no ClusterComputeResource found in simulator")
	}
	return m, ccr.Reference()
}

// minimalSpec returns a SolutionSpec with the minimum fields populated.
func minimalSpec(name string) *vms.SolutionSpec {
	return &vms.SolutionSpec{
		DeploymentType:     vms.EveryHostPinned,
		DisplayName:        name,
		DisplayVersion:     "1.0.0",
		VmNameTemplate:     vms.VmNameTemplate{Prefix: name + "-vm", Suffix: vms.Uuid},
		VmCloneConfig:      vms.NoClones,
		VmStoragePolicy:    vms.Default,
		VmDiskType:         vms.DiskTypeDefault,
		RedeploymentPolicy: vms.ReCreate,
		OvfResource:        vms.OvfResource{LocationType: vms.RemoteFile},
	}
}

// TestSetGetDelete exercises the basic CRUD lifecycle via task-based operations.
func TestSetGetDelete(t *testing.T) {
	simulator.Test(func(ctx context.Context, vc *vim25.Client) {
		m, cluster := testSetup(ctx, vc, t)

		const id = "sol-a"
		spec := minimalSpec(id)

		if err := m.Set(ctx, cluster, id, spec); err != nil {
			t.Fatalf("Set: %v", err)
		}

		info, err := m.Get(ctx, cluster, id)
		if err != nil {
			t.Fatalf("Get: %v", err)
		}
		if info.DisplayName != spec.DisplayName {
			t.Errorf("DisplayName: got %q, want %q", info.DisplayName, spec.DisplayName)
		}

		if err := m.Delete(ctx, cluster, id); err != nil {
			t.Fatalf("Delete: %v", err)
		}

		if _, err := m.Get(ctx, cluster, id); err == nil {
			t.Fatal("Get after Delete: expected error, got nil")
		}
	})
}

// TestList verifies that List returns an accurate count of known solutions.
func TestList(t *testing.T) {
	simulator.Test(func(ctx context.Context, vc *vim25.Client) {
		m, cluster := testSetup(ctx, vc, t)

		result, err := m.List(ctx, cluster)
		if err != nil {
			t.Fatalf("List (empty): %v", err)
		}
		if len(result.Solutions) != 0 {
			t.Fatalf("expected 0 solutions initially, got %d", len(result.Solutions))
		}

		for _, id := range []string{"sol-1", "sol-2"} {
			if err := m.Set(ctx, cluster, id, minimalSpec(id)); err != nil {
				t.Fatalf("Set %s: %v", id, err)
			}
		}

		result, err = m.List(ctx, cluster)
		if err != nil {
			t.Fatalf("List: %v", err)
		}
		if len(result.Solutions) != 2 {
			t.Fatalf("expected 2 solutions, got %d", len(result.Solutions))
		}
	})
}

// TestApply verifies that Apply returns a task ID and the task completes.
func TestApply(t *testing.T) {
	simulator.Test(func(ctx context.Context, vc *vim25.Client) {
		m, cluster := testSetup(ctx, vc, t)

		taskID, err := m.Apply(ctx, cluster, &vms.ApplySpec{})
		if err != nil {
			t.Fatalf("Apply: %v", err)
		}
		if taskID == "" {
			t.Fatal("Apply: empty task ID")
		}
		if _, err := m.ApplyWaitForCompletion(ctx, taskID); err != nil {
			t.Fatalf("ApplyWaitForCompletion: %v", err)
		}
	})
}

// TestCheckCompliance verifies that CheckCompliance completes without error.
func TestCheckCompliance(t *testing.T) {
	simulator.Test(func(ctx context.Context, vc *vim25.Client) {
		m, cluster := testSetup(ctx, vc, t)

		if _, err := m.CheckCompliance(ctx, cluster, &vms.CheckComplianceFilterSpec{}); err != nil {
			t.Fatalf("CheckCompliance: %v", err)
		}
	})
}

// TestListHooks verifies that ListHooks returns a non-nil hooks slice.
func TestListHooks(t *testing.T) {
	simulator.Test(func(ctx context.Context, vc *vim25.Client) {
		m, cluster := testSetup(ctx, vc, t)

		result, err := m.ListHooks(ctx, cluster, "any-solution")
		if err != nil {
			t.Fatalf("ListHooks: %v", err)
		}
		if result.Hooks == nil {
			t.Fatal("ListHooks: expected non-nil hooks slice")
		}
	})
}

// TestMarkAsProcessed verifies that MarkAsProcessed succeeds.
func TestMarkAsProcessed(t *testing.T) {
	simulator.Test(func(ctx context.Context, vc *vim25.Client) {
		m, cluster := testSetup(ctx, vc, t)

		spec := &vms.ProcessedHookSpec{
			Vm:                    "vm-1",
			LifecycleState:        vms.PostPowerOn,
			ProcessedSuccessfully: true,
		}
		if _, err := m.MarkAsProcessed(ctx, cluster, spec); err != nil {
			t.Fatalf("MarkAsProcessed: %v", err)
		}
	})
}

// TestProcessDynamicUpdate verifies that ProcessDynamicUpdate succeeds.
func TestProcessDynamicUpdate(t *testing.T) {
	simulator.Test(func(ctx context.Context, vc *vim25.Client) {
		m, cluster := testSetup(ctx, vc, t)

		spec := &vms.DynamicUpdateSpec{
			Vm:             "vm-1",
			Solution:       "sol-a",
			LifecycleState: vms.PostPowerOn,
		}
		if err := m.ProcessDynamicUpdate(ctx, cluster, spec); err != nil {
			t.Fatalf("ProcessDynamicUpdate: %v", err)
		}
	})
}

// TestEnable verifies that Enable (transition from EAM-managed to vLCM) succeeds.
func TestEnable(t *testing.T) {
	simulator.Test(func(ctx context.Context, vc *vim25.Client) {
		m, cluster := testSetup(ctx, vc, t)

		spec := &vms.EnableSpec{
			EamAgencyID: "agency-1",
			Solution:    minimalSpec("sol-enable"),
		}
		if err := m.Enable(ctx, cluster, "sol-enable", spec); err != nil {
			t.Fatalf("Enable: %v", err)
		}
	})
}

// TestMultiSourceEnable verifies that MultiSourceEnable succeeds.
func TestMultiSourceEnable(t *testing.T) {
	simulator.Test(func(ctx context.Context, vc *vim25.Client) {
		m, cluster := testSetup(ctx, vc, t)

		spec := &vms.MultiSourceEnableSpec{
			EamAgencyIDs: []string{"agency-1", "agency-2"},
			Solution:     minimalSpec("sol-multi"),
		}
		if err := m.MultiSourceEnable(ctx, cluster, "sol-multi", spec); err != nil {
			t.Fatalf("MultiSourceEnable: %v", err)
		}
	})
}

// TestTransition verifies that Transition (move solution to another cluster) succeeds.
func TestTransition(t *testing.T) {
	simulator.Test(func(ctx context.Context, vc *vim25.Client) {
		m, cluster := testSetup(ctx, vc, t)

		spec := &vms.TransitionSpec{
			SourceCluster: "domain-c8",
			Solution:      minimalSpec("sol-transition"),
		}
		if err := m.Transition(ctx, cluster, "sol-transition", spec); err != nil {
			t.Fatalf("Transition: %v", err)
		}
	})
}
