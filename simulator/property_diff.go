// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"reflect"

	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

// PropertyDiff compares two states of a managed object and returns the PropertyChanges
// representing the differences. The old and new parameters should be pointers to the
// same type of managed object (e.g., *mo.VirtualMachine).
//
// This is useful for generating granular property change notifications after modifying
// an object. The typical pattern is:
//
//	old := new(mo.VirtualMachine)
//	deepCopy(vm, old)
//	// ... make changes to vm ...
//	changes := PropertyDiff(old, vm)
//	ctx.Update(vm, changes)
func PropertyDiff(old, new mo.Reference) []types.PropertyChange {
	var changes []types.PropertyChange

	oldVal := getManagedObject(old)
	newVal := getManagedObject(new)

	diffFields("", oldVal, newVal, oldVal.Type(), &changes)

	return changes
}

// diffFields recursively compares struct fields and appends PropertyChanges for differences.
func diffFields(prefix string, oldVal, newVal reflect.Value, rtype reflect.Type, changes *[]types.PropertyChange) {
	for i := 0; i < rtype.NumField(); i++ {
		f := rtype.Field(i)

		// Skip the Self reference field
		if f.Name == "Self" {
			continue
		}

		oldField := oldVal.Field(i)
		newField := newVal.Field(i)

		// Build the property path
		name := lcFirst(f.Name)
		path := name
		if prefix != "" {
			path = prefix + "." + name
		}

		// Handle embedded/anonymous fields by recursing without adding to path
		if f.Anonymous {
			diffFields(prefix, oldField, newField, f.Type, changes)
			continue
		}

		// Compare the field values
		if !fieldsEqual(oldField, newField) {
			change := types.PropertyChange{
				Name: path,
				Op:   determineChangeOp(oldField, newField),
			}

			// Set the value based on the operation
			if change.Op != types.PropertyChangeOpRemove {
				change.Val = fieldValueInterface(f, newField)
			}

			*changes = append(*changes, change)
		}
	}
}

// fieldsEqual reports whether oldField and newField hold equal values for the
// purpose of change detection. A nil slice and a non-nil, zero-length slice
// of the same element type are treated as equal: neither carries observable
// content, so reflect.DeepEqual's usual nil-vs-empty distinction (which would
// otherwise report a change here) would produce a spurious PropertyChange for
// a transition that changed nothing a consumer can observe.
func fieldsEqual(oldField, newField reflect.Value) bool {
	if oldField.Kind() == reflect.Slice && newField.Kind() == reflect.Slice {
		if oldField.Len() == 0 && newField.Len() == 0 {
			return true
		}
	}
	return reflect.DeepEqual(oldField.Interface(), newField.Interface())
}

// determineChangeOp determines the appropriate PropertyChangeOp based on old and new values.
func determineChangeOp(oldVal, newVal reflect.Value) types.PropertyChangeOp {
	oldEmpty := isEmpty(oldVal)
	newEmpty := isEmpty(newVal)

	switch {
	case oldEmpty && !newEmpty:
		return types.PropertyChangeOpAdd
	case !oldEmpty && newEmpty:
		return types.PropertyChangeOpRemove
	default:
		return types.PropertyChangeOpAssign
	}
}

// Checkpoint creates a deep copy of a managed object's mo state that can later
// be used with PropertyDiff to generate property changes.
//
// obj may be either a pure mo-package type (e.g. *mo.VirtualMachine) or a
// simulator wrapper type (e.g. *simulator.VirtualMachine). In the latter case
// the embedded mo type is extracted first so that deepCopy only operates on
// known vSphere types; the returned snapshot is always a *mo.T value.
//
// Example:
//
//	checkpoint := Checkpoint(vm) // vm can be *simulator.VirtualMachine or *mo.VirtualMachine
//	// ... make changes to vm ...
//	changes := PropertyDiff(checkpoint, vm)
//	ctx.Update(vm, changes)
func Checkpoint(obj mo.Reference) mo.Reference {
	moVal := getManagedObject(obj)
	newPtr := reflect.New(moVal.Type())
	dst := newPtr.Interface().(mo.Reference)
	deepCopy(moVal.Addr().Interface(), dst)
	return dst
}
