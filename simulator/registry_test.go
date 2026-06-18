// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"fmt"
	"testing"

	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

func TestRegistry(t *testing.T) {
	r := NewRegistry()

	ref := types.ManagedObjectReference{Type: "Test", Value: "Test"}
	f := &mo.Folder{}
	f.Self = ref
	r.PutEntity(nil, f)

	e := r.Get(ref)

	if e.Reference() != ref {
		t.Fail()
	}

	r.Remove(NewContext(), ref)

	if r.Get(ref) != nil {
		t.Fail()
	}

	r.Put(e)
	e = r.Get(ref)

	if e.Reference() != ref {
		t.Fail()
	}
}

func TestRemoveReference(t *testing.T) {
	var refs []types.ManagedObjectReference

	for i := 0; i < 5; i++ {
		refs = append(refs, types.ManagedObjectReference{Type: "any", Value: fmt.Sprintf("%d", i)})
	}

	n := len(refs)

	RemoveReference(&refs, refs[2])

	if len(refs) != n-1 {
		t.Errorf("%d", len(refs))
	}
}

func TestAlignCounter(t *testing.T) {
	r := NewRegistry()

	// 1. Empty registry AlignCounter should result in counter 0
	if err := r.AlignCounter(); err != nil {
		t.Fatalf("AlignCounter failed: %v", err)
	}
	if r.counter != 0 {
		t.Errorf("expected counter 0, got %d", r.counter)
	}

	// 2. Add some entities with specific types and references
	vm1 := &mo.VirtualMachine{}
	vm1.Self = types.ManagedObjectReference{Type: "VirtualMachine", Value: "vm-42"}
	r.Put(vm1)

	host1 := &mo.HostSystem{}
	host1.Self = types.ManagedObjectReference{Type: "HostSystem", Value: "host-10"}
	r.Put(host1)

	rp1 := &mo.ResourcePool{}
	rp1.Self = types.ManagedObjectReference{Type: "ResourcePool", Value: "resgroup-250"}
	r.Put(rp1)

	// Non-conforming/invalid suffix formats should be skipped
	vmInvalid := &mo.VirtualMachine{}
	vmInvalid.Self = types.ManagedObjectReference{Type: "VirtualMachine", Value: "vm-abc"}
	r.Put(vmInvalid)

	vmNonNumeric := &mo.VirtualMachine{}
	vmNonNumeric.Self = types.ManagedObjectReference{Type: "VirtualMachine", Value: "vm-v50"}
	r.Put(vmNonNumeric)

	// Another valid one but smaller than max
	vmSmall := &mo.VirtualMachine{}
	vmSmall.Self = types.ManagedObjectReference{Type: "VirtualMachine", Value: "vm-100"}
	r.Put(vmSmall)

	if err := r.AlignCounter(); err != nil {
		t.Fatalf("AlignCounter failed: %v", err)
	}

	// Max numeric suffix should be 250 (from resgroup-250)
	if r.counter != 250 {
		t.Errorf("expected counter 250, got %d", r.counter)
	}

	// 3. Verify next allocated reference uses the aligned counter + 1 (251)
	newVM := &mo.VirtualMachine{}
	ref := r.reference(newVM)
	expectedValue := "vm-251"
	if ref.Value != expectedValue {
		t.Errorf("expected new reference value %q, got %q", expectedValue, ref.Value)
	}
}
