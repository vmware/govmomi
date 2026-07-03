// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"strings"
	"testing"

	"github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vim25/xml"
)

// TestCnsVolumePolicyReconfigSpecActiveClusters verifies that the
// ActiveClusters field serializes into the CNS ReconfigPolicy request, so CNS
// can auto-relocate a volume whose target profile differs. It mirrors the
// existing CnsVolumeCreateSpec.ActiveClusters serialization.
func TestCnsVolumePolicyReconfigSpecActiveClusters(t *testing.T) {
	spec := CnsVolumePolicyReconfigSpec{
		VolumeId: CnsVolumeId{Id: "vol-1"},
		Profile: []types.BaseVirtualMachineProfileSpec{
			&types.VirtualMachineDefinedProfileSpec{ProfileId: "policy-1"},
		},
		ActiveClusters: []types.ManagedObjectReference{
			{Type: "ClusterComputeResource", Value: "domain-c1"},
		},
	}

	b, err := xml.Marshal(spec)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	out := string(b)
	if !strings.Contains(out, "activeClusters") {
		t.Errorf("expected activeClusters element in:\n%s", out)
	}
	if !strings.Contains(out, "domain-c1") {
		t.Errorf("expected cluster moref value domain-c1 in:\n%s", out)
	}

	// omitempty: no ActiveClusters -> no element (an in-place reconfigure).
	spec.ActiveClusters = nil
	b, err = xml.Marshal(spec)
	if err != nil {
		t.Fatalf("marshal (empty): %v", err)
	}
	if strings.Contains(string(b), "activeClusters") {
		t.Errorf("expected no activeClusters element when empty, got:\n%s", string(b))
	}
}
