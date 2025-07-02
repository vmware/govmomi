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
