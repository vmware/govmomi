// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package object

import "testing"

// TestKindsAliasVmwareDVS guards against "-type w" mapping back to the
// abstract DistributedVirtualSwitch. That name resolves fine against a
// generic hierarchy-aware type filter, but "govc find" applies it as an
// exact string match against each traversed object's own reported type
// (kinds.wanted), which is always the concrete VmwareDistributedVirtualSwitch
// on a real vCenter -- so the abstract name here silently matched nothing.
// It went unnoticed because vcsim used to (incorrectly) register its own
// switch under the abstract type too, until that was fixed to match real
// vCenter.
func TestKindsAliasVmwareDVS(t *testing.T) {
	var k kinds
	got := k.alias("w")
	want := "VmwareDistributedVirtualSwitch"
	if got != want {
		t.Errorf(`alias("w") = %q, want %q`, got, want)
	}
}
