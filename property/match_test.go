// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package property_test

import (
	"testing"

	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/vim25/types"
)

func TestMatchProperty(t *testing.T) {
	for _, test := range []struct {
		key  string
		val  types.AnyType
		pass types.AnyType
		fail types.AnyType
	}{
		{"string", "bar", "bar", "foo"},
		{"match", "foo.bar", "foo.*", "foobarbaz"},
		{"moref", types.ManagedObjectReference{Type: "HostSystem", Value: "foo"}, "HostSystem:foo", "bar"}, // implements fmt.Stringer
		{"morefm", types.ManagedObjectReference{Type: "HostSystem", Value: "foo"}, "*foo", "bar"},
		{"morefs", types.ArrayOfManagedObjectReference{ManagedObjectReference: []types.ManagedObjectReference{{Type: "HostSystem", Value: "foo"}}}, "*foo", "bar"},
		{"enum", types.VirtualMachinePowerStatePoweredOn, "poweredOn", "poweredOff"},
		{"int16", int32(16), int32(16), int32(42)},
		{"int32", int32(32), int32(32), int32(42)},
		{"int32s", int32(32), "32", "42"},
		{"int64", int64(64), int64(64), int64(42)},
		{"int64s", int64(64), "64", "42"},
		{"float32", float32(32.32), float32(32.32), float32(42.0)},
		{"float32s", float32(32.32), "32.32", "42.0"},
		{"float64", float64(64.64), float64(64.64), float64(42.0)},
		{"float64s", float64(64.64), "64.64", "42.0"},
		{"matchFunc", "bar", func(s any) bool { return s.(string) == "bar" }, func(s any) bool { return s.(string) == "foo" }},
	} {
		p := types.DynamicProperty{Name: test.key, Val: test.val}

		for match, value := range map[bool]types.AnyType{true: test.pass, false: test.fail} {
			result := property.Match{test.key: value}.Property(p)

			if result != match {
				t.Errorf("%s: %t", test.key, result)
			}
		}
	}
}
