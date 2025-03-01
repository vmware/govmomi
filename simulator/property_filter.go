// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"reflect"
	"strings"

	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

type PropertyFilter struct {
	mo.PropertyFilter

	pc   *PropertyCollector
	refs map[types.ManagedObjectReference]struct{}
	sync bool
}

func (f *PropertyFilter) UpdateObject(ctx *Context, o mo.Reference, changes []types.PropertyChange) {
	// A PropertyFilter's traversal spec is "applied" on the initial call to WaitForUpdates,
	// with matching objects tracked in the `refs` field.
	// New and deleted objects matching the filter are accounted for within PropertyCollector.
	// But when an object used for the traversal itself is updated (e.g. ListView),
	// we need to update the tracked `refs` on the next call to WaitForUpdates.
	ref := o.Reference()

	for _, set := range f.Spec.ObjectSet {
		if set.Obj == ref && len(set.SelectSet) != 0 {
			ctx.WithLock(f, func() { f.sync = true })
			break
		}
	}
}

func (_ *PropertyFilter) PutObject(_ *Context, _ mo.Reference) {}

func (_ *PropertyFilter) RemoveObject(_ *Context, _ types.ManagedObjectReference) {}

func (f *PropertyFilter) DestroyPropertyFilter(ctx *Context, c *types.DestroyPropertyFilter) soap.HasFault {
	body := &methods.DestroyPropertyFilterBody{}

	ctx.WithLock(f.pc, func() {
		RemoveReference(&f.pc.Filter, c.This)
	})

	ctx.Session.Remove(ctx, c.This)

	body.Res = &types.DestroyPropertyFilterResponse{}

	return body
}

func (f *PropertyFilter) collect(ctx *Context) (*types.RetrieveResult, types.BaseMethodFault) {
	req := &types.RetrievePropertiesEx{
		SpecSet: []types.PropertyFilterSpec{f.Spec},
	}
	return collect(ctx, req)
}

func (f *PropertyFilter) update(ctx *Context) {
	ctx.WithLock(f, func() {
		if f.sync {
			f.sync = false
			clear(f.refs)
			_, _ = f.collect(ctx)
		}
	})
}

// matches returns true if the change matches one of the filter Spec.PropSet
func (f *PropertyFilter) matches(ctx *Context, ref types.ManagedObjectReference, change *types.PropertyChange) bool {
	var kind reflect.Type

	for _, p := range f.Spec.PropSet {
		if p.Type != ref.Type {
			if kind == nil {
				obj := ctx.Map.Get(ref)
				if obj == nil { // object may have since been deleted
					continue
				}
				kind = getManagedObject(obj).Type()
			}
			// e.g. ManagedEntity, ComputeResource
			field, ok := kind.FieldByName(p.Type)
			if !(ok && field.Anonymous) {
				continue
			}
		}

		if isTrue(p.All) {
			return true
		}

		for _, name := range p.PathSet {
			if name == change.Name {
				return true
			}

			var field mo.Field
			if field.FromString(name) && field.Item != "" {
				// "field[key].item" -> "field[key]"
				item := field.Item
				field.Item = ""
				if field.String() == change.Name {
					change.Name = name
					change.Val, _ = fieldValue(reflect.ValueOf(change.Val), item)
					return true
				}
			}

			if field.FromString(change.Name) && field.Key != nil {
				continue // case below does not apply to property index
			}

			// strings.HasPrefix("runtime.powerState", "runtime") == parent field matches
			if strings.HasPrefix(change.Name, name) {
				if obj := ctx.Map.Get(ref); obj != nil { // object may have since been deleted
					change.Name = name
					change.Val, _ = fieldValue(reflect.ValueOf(obj), name)
				}

				return true
			}
		}
	}

	return false
}

// apply the PropertyFilter.Spec to the given ObjectUpdate
func (f *PropertyFilter) apply(ctx *Context, change types.ObjectUpdate) types.ObjectUpdate {
	parents := make(map[string]bool)
	set := change.ChangeSet
	change.ChangeSet = nil

	for i, p := range set {
		if f.matches(ctx, change.Obj, &p) {
			if p.Name != set[i].Name {
				// update matches a parent field from the spec.
				if parents[p.Name] {
					continue // only return 1 instance of the parent
				}
				parents[p.Name] = true
			}
			change.ChangeSet = append(change.ChangeSet, p)
		}
	}

	return change
}
