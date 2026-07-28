// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vms_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vapi/esx/settings/clusters/vms"
	"github.com/vmware/govmomi/vapi/rest"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/types"

	_ "github.com/vmware/govmomi/vapi/esx/settings/clusters/vms/simulator"
	_ "github.com/vmware/govmomi/vapi/simulator"
)

// transitionTestSetup returns a logged-in vms.Manager and a cluster MOR from the simulator.
func transitionTestSetup(ctx context.Context, vc *vim25.Client, t *testing.T) (*vms.Manager, types.ManagedObjectReference) {
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

// minimalSolutionSpec returns a SolutionSpec with the minimum fields populated.
func minimalSolutionSpec(name string) *vms.SolutionSpec {
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

// TestDeleteSolutionOnly is the primary test for the DeleteSolutionOnly API.
// It sets a solution, confirms it exists, calls DeleteSolutionOnly, then verifies
// the solution is absent.
func TestDeleteSolutionOnly(t *testing.T) {
	simulator.Test(func(ctx context.Context, vc *vim25.Client) {
		m, cluster := transitionTestSetup(ctx, vc, t)

		const solutionID = "my-solution"

		// Set the solution so it exists in the simulator state.
		if err := m.Set(ctx, cluster, solutionID, minimalSolutionSpec(solutionID)); err != nil {
			t.Fatalf("Set: %v", err)
		}

		// Confirm the solution is visible before calling DeleteSolutionOnly.
		if _, err := m.Get(ctx, cluster, solutionID); err != nil {
			t.Fatalf("Get (before DeleteSolutionOnly): %v", err)
		}

		// DeleteSolutionOnly is the method under test.
		if err := m.DeleteSolutionOnly(ctx, cluster, solutionID); err != nil {
			t.Fatalf("DeleteSolutionOnly: %v", err)
		}

		// The solution must no longer exist after DeleteSolutionOnly.
		_, err := m.Get(ctx, cluster, solutionID)
		if err == nil {
			t.Fatal("Get after DeleteSolutionOnly: expected error, got nil")
		}
		if !rest.IsStatusError(err, http.StatusNotFound) {
			t.Fatalf("Get after DeleteSolutionOnly: expected 404, got %v", err)
		}
	})
}

// TestDeleteSolutionOnlyNotFound verifies that DeleteSolutionOnly returns a 404
// when the solution does not exist.
func TestDeleteSolutionOnlyNotFound(t *testing.T) {
	simulator.Test(func(ctx context.Context, vc *vim25.Client) {
		m, cluster := transitionTestSetup(ctx, vc, t)

		err := m.DeleteSolutionOnly(ctx, cluster, "nonexistent")
		if err == nil {
			t.Fatal("expected 404 error for nonexistent solution, got nil")
		}
		if !rest.IsStatusError(err, http.StatusNotFound) {
			t.Fatalf("expected 404, got %v", err)
		}
	})
}
