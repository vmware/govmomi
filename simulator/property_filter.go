// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"reflect"
	"strings"
	"sync/atomic"

	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

type PropertyFilter struct {
	mo.PropertyFilter

	pc   *PropertyCollector
	refs map[types.ManagedObjectReference]struct{}
	// sync is atomic so UpdateObject can flag a re-sync without taking this
	// filter's object lock. UpdateObject runs from applyHandlers while the
	// caller already holds the traversal object's lock (e.g. ModifyListView
	// holds the ListView lock); taking the filter lock there inverts the
	// filter->object order used by update/collect and deadlocks.
	sync atomic.Bool
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
			f.sync.Store(true)
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

	ctx.Map.RemoveHandler(f)
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
	if !f.sync.CompareAndSwap(true, false) {
		return
	}
	ctx.WithLock(f, func() {
		clear(f.refs)
		_, _ = f.collect(ctx)
	})
}

// matches returns the PropertyChanges the filter reports for change: one for
// each PathSet entry that change matches. A change to a parent property, as
// PropertyDiff reports "guest" for a change to guest.net, matches every
// requested path under it, and each is reported. It returns nil if none match.
func (f *PropertyFilter) matches(ctx *Context, ref types.ManagedObjectReference, change types.PropertyChange) []types.PropertyChange {
	var kind reflect.Type
	var out []types.PropertyChange

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
			return []types.PropertyChange{change}
		}

		for _, name := range p.PathSet {
			c := change

			if name == change.Name {
				out = append(out, c)
				continue
			}

			var field mo.Field
			if field.FromString(name) && field.Item != "" {
				// "field[key].item" -> "field[key]"
				item := field.Item
				field.Item = ""
				if field.String() == change.Name {
					c.Name = name
					c.Val, _ = fieldValue(reflect.ValueOf(c.Val), item)
					out = append(out, c)
					continue
				}
			}

			if field.FromString(change.Name) && field.Key != nil {
				continue // case below does not apply to property index
			}

			// strings.HasPrefix("runtime.powerState", "runtime") == parent field matches.
			// strings.HasPrefix("guest.hostName", "guest.") == the change is to a
			// parent of the field, as PropertyDiff reports, which may change it.
			if strings.HasPrefix(change.Name, name) || strings.HasPrefix(name, change.Name+".") {
				if obj := ctx.Map.Get(ref); obj != nil { // object may have since been deleted
					c.Name = name
					c.Val, _ = fieldValue(reflect.ValueOf(obj), name)
				}
				out = append(out, c)
			}
		}
	}

	return out
}

// apply the PropertyFilter.Spec to the given ObjectUpdate
func (f *PropertyFilter) apply(ctx *Context, change types.ObjectUpdate) types.ObjectUpdate {
	seen := make(map[string]bool)
	set := change.ChangeSet
	change.ChangeSet = nil

	for _, p := range set {
		for _, c := range f.matches(ctx, change.Obj, p) {
			if seen[c.Name] {
				continue // only report 1 instance of each path
			}
			seen[c.Name] = true
			change.ChangeSet = append(change.ChangeSet, c)
		}
	}

	return change
}
