// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package mo

import (
	"reflect"
	"testing"

	"github.com/vmware/govmomi/vim25/types"
)

func TestLoadAll(*testing.T) {
	for _, typ := range t {
		newTypeInfo(typ)
	}
}

func TestApplyPropertyChange(t *testing.T) {
	changes := []types.PropertyChange{
		{Name: "snapshot.currentSnapshot", Val: nil},
		{Name: "snapshot.currentSnapshot", Val: types.ManagedObjectReference{}},
		{Name: "snapshot.currentSnapshot", Val: &types.ManagedObjectReference{}},
	}
	var vm VirtualMachine
	vm.Self = types.ManagedObjectReference{Type: "VirtualMachine"}
	vm.Snapshot = new(types.VirtualMachineSnapshotInfo)
	ApplyPropertyChange(vm, changes)
}

// The virtual machine managed object has about 500 nested properties.
// It's likely to be indicative of the function's performance in general.
func BenchmarkLoadVirtualMachine(b *testing.B) {
	vmtyp := reflect.TypeOf((*VirtualMachine)(nil)).Elem()
	for i := 0; i < b.N; i++ {
		newTypeInfo(vmtyp)
	}
}

func TestPropertyPathFromString(t *testing.T) {
	tests := []struct {
		path   string
		expect *Field
	}{
		{`foo.bar`, &Field{Path: "foo.bar"}},
		{`foo.bar["biz"]`, &Field{Path: "foo.bar", Key: "biz"}},
		{`foo.bar["biz"].baz`, &Field{Path: "foo.bar", Key: "biz", Item: "baz"}},
		{`foo.bar[0]`, &Field{Path: "foo.bar", Key: int32(0)}},
		{`foo.bar[1].baz`, &Field{Path: "foo.bar", Key: int32(1), Item: "baz"}},
		{`foo.bar[1].baz.buz`, &Field{Path: "foo.bar", Key: int32(1), Item: "baz.buz"}},
	}

	for i, tp := range tests {
		var field Field
		if field.FromString(tp.path) {
			if field.String() != tp.expect.String() {
				t.Errorf("%d: %s != %s", i, field, tp.expect)
			}
			if field != *tp.expect {
				t.Errorf("%d: %#v != %#v", i, field, *tp.expect)
			}
		} else {
			t.Error(tp.path)
		}
	}
}
