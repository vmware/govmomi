// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"reflect"

	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

type ViewManager struct {
	mo.ViewManager

	entities map[string]bool
}

var entities = []struct {
	Type      reflect.Type
	Container bool
}{
	{reflect.TypeOf((*mo.ManagedEntity)(nil)).Elem(), true},
	{reflect.TypeOf((*mo.Folder)(nil)).Elem(), true},
	{reflect.TypeOf((*mo.StoragePod)(nil)).Elem(), true},
	{reflect.TypeOf((*mo.Datacenter)(nil)).Elem(), true},
	{reflect.TypeOf((*mo.ComputeResource)(nil)).Elem(), true},
	{reflect.TypeOf((*mo.ClusterComputeResource)(nil)).Elem(), true},
	{reflect.TypeOf((*mo.HostSystem)(nil)).Elem(), true},
	{reflect.TypeOf((*mo.ResourcePool)(nil)).Elem(), true},
	{reflect.TypeOf((*mo.VirtualApp)(nil)).Elem(), true},
	{reflect.TypeOf((*mo.VirtualMachine)(nil)).Elem(), false},
	{reflect.TypeOf((*mo.Datastore)(nil)).Elem(), false},
	{reflect.TypeOf((*mo.Network)(nil)).Elem(), false},
	{reflect.TypeOf((*mo.OpaqueNetwork)(nil)).Elem(), false},
	{reflect.TypeOf((*mo.DistributedVirtualPortgroup)(nil)).Elem(), false},
	{reflect.TypeOf((*mo.DistributedVirtualSwitch)(nil)).Elem(), false},
	{reflect.TypeOf((*mo.VmwareDistributedVirtualSwitch)(nil)).Elem(), false},
}

func (m *ViewManager) init(*Registry) {
	m.entities = make(map[string]bool, len(entities))
	for _, e := range entities {
		m.entities[e.Type.Name()] = e.Container
	}
}

func destroyView(ref types.ManagedObjectReference) soap.HasFault {
	return &methods.DestroyViewBody{
		Res: &types.DestroyViewResponse{},
	}
}

func (m *ViewManager) CreateContainerView(ctx *Context, req *types.CreateContainerView) soap.HasFault {
	body := &methods.CreateContainerViewBody{}

	root := ctx.Map.Get(req.Container)
	if root == nil {
		body.Fault_ = Fault("", &types.ManagedObjectNotFound{Obj: req.Container})
		return body
	}

	if !m.entities[root.Reference().Type] {
		body.Fault_ = Fault("", &types.InvalidArgument{InvalidProperty: "container"})
		return body
	}

	container := &ContainerView{
		mo.ContainerView{
			Container: root.Reference(),
			Recursive: req.Recursive,
			Type:      req.Type,
		},
		root,
		make(map[string]bool),
	}

	for _, ctype := range container.Type {
		if _, ok := m.entities[ctype]; !ok {
			body.Fault_ = Fault("", &types.InvalidArgument{InvalidProperty: "type"})
			return body
		}

		container.types[ctype] = true

		for _, e := range entities {
			// Check for embedded types
			if f, ok := e.Type.FieldByName(ctype); ok && f.Anonymous {
				container.types[e.Type.Name()] = true
			}
		}
	}

	ctx.Session.setReference(container)

	body.Res = &types.CreateContainerViewResponse{
		Returnval: container.Self,
	}

	seen := make(map[types.ManagedObjectReference]bool)
	container.add(ctx, root, seen)

	ctx.Session.Registry.Put(container)
	ctx.Map.AddHandler(container)

	return body
}

type ContainerView struct {
	mo.ContainerView

	root  mo.Reference
	types map[string]bool
}

func (v *ContainerView) DestroyView(ctx *Context, c *types.DestroyView) soap.HasFault {
	ctx.Map.RemoveHandler(v)
	ctx.Session.Remove(ctx, c.This)
	return destroyView(c.This)
}

func (v *ContainerView) include(o types.ManagedObjectReference) bool {
	if len(v.types) == 0 {
		return true
	}

	return v.types[o.Type]
}

func walk(root mo.Reference, f func(child types.ManagedObjectReference)) {
	if _, ok := root.(types.ManagedObjectReference); ok || root == nil {
		return
	}

	var children []types.ManagedObjectReference

	switch e := getManagedObject(root).Addr().Interface().(type) {
	case *mo.Datacenter:
		children = []types.ManagedObjectReference{e.VmFolder, e.HostFolder, e.DatastoreFolder, e.NetworkFolder}
	case *mo.Folder:
		children = e.ChildEntity
	case *mo.StoragePod:
		children = e.ChildEntity
	case *mo.ComputeResource:
		children = e.Host
		children = append(children, *e.ResourcePool)
	case *mo.ClusterComputeResource:
		children = e.Host
		children = append(children, *e.ResourcePool)
	case *mo.ResourcePool:
		children = e.ResourcePool
		children = append(children, e.Vm...)
	case *mo.VirtualApp:
		children = e.ResourcePool.ResourcePool
		children = append(children, e.Vm...)
	case *mo.HostSystem:
		children = e.Vm
	}

	for _, child := range children {
		f(child)
	}
}

func (v *ContainerView) add(ctx *Context, root mo.Reference, seen map[types.ManagedObjectReference]bool) {
	walk(root, func(child types.ManagedObjectReference) {
		if v.include(child) {
			if !seen[child] {
				seen[child] = true
				v.View = append(v.View, child)
			}
		}

		if v.Recursive {
			v.add(ctx, ctx.Map.Get(child), seen)
		}
	})
}

func (v *ContainerView) find(ctx *Context, root mo.Reference, ref types.ManagedObjectReference, found *bool) bool {
	walk(root, func(child types.ManagedObjectReference) {
		if *found {
			return
		}
		if child == ref {
			*found = true
			return
		}
		if v.Recursive {
			*found = v.find(ctx, ctx.Map.Get(child), ref, found)
		}
	})

	return *found
}

func (v *ContainerView) PutObject(ctx *Context, obj mo.Reference) {
	ref := obj.Reference()

	ctx.WithLock(v, func() {
		if v.include(ref) && v.find(ctx, v.root, ref, types.NewBool(false)) {
			ctx.Update(v, []types.PropertyChange{{Name: "view", Val: append(v.View, ref)}})
		}
	})
}

func (v *ContainerView) RemoveObject(ctx *Context, obj types.ManagedObjectReference) {
	ctx.WithLock(v, func() {
		ctx.Map.RemoveReference(ctx, v, &v.View, obj)
	})
}

func (*ContainerView) UpdateObject(*Context, mo.Reference, []types.PropertyChange) {}

func (m *ViewManager) CreateListView(ctx *Context, req *types.CreateListView) soap.HasFault {
	body := new(methods.CreateListViewBody)
	list := new(ListView)

	if refs := list.add(ctx, req.Obj); len(refs) != 0 {
		body.Fault_ = Fault("", &types.ManagedObjectNotFound{Obj: refs[0]})
		return body
	}

	ctx.Session.Put(list)

	body.Res = &types.CreateListViewResponse{
		Returnval: list.Self,
	}

	return body
}

type ListView struct {
	mo.ListView
}

func (v *ListView) update(ctx *Context) {
	ctx.Update(v, []types.PropertyChange{{Name: "view", Val: v.View}})
}

func (v *ListView) add(ctx *Context, refs []types.ManagedObjectReference) []types.ManagedObjectReference {
	var unresolved []types.ManagedObjectReference
	for _, ref := range refs {
		obj := ctx.Session.Get(ref)
		if obj == nil {
			unresolved = append(unresolved, ref)
		}
		v.View = append(v.View, ref)
	}
	return unresolved
}

func (v *ListView) DestroyView(ctx *Context, c *types.DestroyView) soap.HasFault {
	ctx.Session.Remove(ctx, c.This)
	return destroyView(c.This)
}

func (v *ListView) ModifyListView(ctx *Context, req *types.ModifyListView) soap.HasFault {
	body := &methods.ModifyListViewBody{
		Res: new(types.ModifyListViewResponse),
	}

	body.Res.Returnval = v.add(ctx, req.Add)

	for _, ref := range req.Remove {
		RemoveReference(&v.View, ref)
	}

	if len(req.Remove) != 0 || len(req.Add) != 0 {
		v.update(ctx)
	}

	return body
}

func (v *ListView) ResetListView(ctx *Context, req *types.ResetListView) soap.HasFault {
	body := &methods.ResetListViewBody{
		Res: new(types.ResetListViewResponse),
	}

	v.View = nil

	body.Res.Returnval = v.add(ctx, req.Obj)

	v.update(ctx)

	return body
}
