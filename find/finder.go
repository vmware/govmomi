/*
Copyright (c) 2014-2015 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package find

import (
	"errors"
	"path"

	"github.com/vmware/govmomi/list"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	"golang.org/x/net/context"
)

// Finder finds things in vspeher
type Finder struct {
	client   *vim25.Client
	recurser list.Recurser

	dc      *object.Datacenter
	folders *object.DatacenterFolders
}

// NewFinder creates a new finder
func NewFinder(client *vim25.Client, all bool) *Finder {
	f := &Finder{
		client: client,
		recurser: list.Recurser{
			Collector: property.DefaultCollector(client),
			All:       all,
		},
	}

	return f
}

// SetDatacenter sets the datacenter
func (f *Finder) SetDatacenter(dc *object.Datacenter) *Finder {
	f.dc = dc
	f.folders = nil
	return f
}

type findRelativeFunc func(ctx context.Context) (object.Reference, error)

func (f *Finder) find(ctx context.Context, fn findRelativeFunc, tl bool, arg string) ([]list.Element, error) {
	root := list.Element{
		Path:   "/",
		Object: object.NewRootFolder(f.client),
	}

	parts := list.ToParts(arg)

	if len(parts) > 0 {
		switch parts[0] {
		case "..": // Not supported; many edge case, little value
			return nil, errors.New("cannot traverse up a tree")
		case ".": // Relative to whatever
			pivot, err := fn(ctx)
			if err != nil {
				return nil, err
			}

			mes, err := mo.Ancestors(ctx, f.client, f.client.ServiceContent.PropertyCollector, pivot.Reference())
			if err != nil {
				return nil, err
			}

			for _, me := range mes {
				// Skip root entity in building inventory path.
				if me.Parent == nil {
					continue
				}
				root.Path = path.Join(root.Path, me.Name)
			}

			root.Object = pivot
			parts = parts[1:]
		}
	}

	f.recurser.TraverseLeafs = tl
	es, err := f.recurser.Recurse(ctx, root, parts)
	if err != nil {
		return nil, err
	}

	return es, nil
}

func (f *Finder) datacenter() (*object.Datacenter, error) {
	if f.dc == nil {
		return nil, errors.New("please specify a datacenter")
	}

	return f.dc, nil
}

func (f *Finder) dcFolders(ctx context.Context) (*object.DatacenterFolders, error) {
	if f.folders != nil {
		return f.folders, nil
	}

	dc, err := f.datacenter()
	if err != nil {
		return nil, err
	}

	folders, err := dc.Folders(ctx)
	if err != nil {
		return nil, err
	}

	f.folders = folders

	return f.folders, nil
}

func (f *Finder) dcReference(_ context.Context) (object.Reference, error) {
	dc, err := f.datacenter()
	if err != nil {
		return nil, err
	}

	return dc, nil
}

func (f *Finder) vmFolder(ctx context.Context) (object.Reference, error) {
	folders, err := f.dcFolders(ctx)
	if err != nil {
		return nil, err
	}

	return folders.VMFolder, nil
}

func (f *Finder) hostFolder(ctx context.Context) (object.Reference, error) {
	folders, err := f.dcFolders(ctx)
	if err != nil {
		return nil, err
	}

	return folders.HostFolder, nil
}

func (f *Finder) datastoreFolder(ctx context.Context) (object.Reference, error) {
	folders, err := f.dcFolders(ctx)
	if err != nil {
		return nil, err
	}

	return folders.DatastoreFolder, nil
}

func (f *Finder) networkFolder(ctx context.Context) (object.Reference, error) {
	folders, err := f.dcFolders(ctx)
	if err != nil {
		return nil, err
	}

	return folders.NetworkFolder, nil
}

func (f *Finder) rootFolder(_ context.Context) (object.Reference, error) {
	return object.NewRootFolder(f.client), nil
}

func (f *Finder) managedObjectList(ctx context.Context, path string, tl bool) ([]list.Element, error) {
	fn := f.rootFolder

	if f.dc != nil {
		fn = f.dcReference
	}

	if len(path) == 0 {
		path = "."
	}

	return f.find(ctx, fn, tl, path)
}

// ManagedObjectList a managed object list
func (f *Finder) ManagedObjectList(ctx context.Context, path string) ([]list.Element, error) {
	return f.managedObjectList(ctx, path, false)
}

// ManagedObjectListChildren the children of a managed object list
func (f *Finder) ManagedObjectListChildren(ctx context.Context, path string) ([]list.Element, error) {
	return f.managedObjectList(ctx, path, true)
}

// DatacenterList list the datacenters
func (f *Finder) DatacenterList(ctx context.Context, path string) ([]*object.Datacenter, error) {
	es, err := f.find(ctx, f.rootFolder, false, path)
	if err != nil {
		return nil, err
	}

	var dcs []*object.Datacenter
	for _, e := range es {
		ref := e.Object.Reference()
		if ref.Type == "Datacenter" {
			dcs = append(dcs, object.NewDatacenter(f.client, ref))
		}
	}

	if len(dcs) == 0 {
		return nil, &NotFoundError{"datacenter", path}
	}

	return dcs, nil
}

// Datacenter finds a datacenter
func (f *Finder) Datacenter(ctx context.Context, path string) (*object.Datacenter, error) {
	dcs, err := f.DatacenterList(ctx, path)
	if err != nil {
		return nil, err
	}

	if len(dcs) > 1 {
		return nil, &MultipleFoundError{"datacenter", path}
	}

	return dcs[0], nil
}

// DefaultDatacenter finds the default data center
func (f *Finder) DefaultDatacenter(ctx context.Context) (*object.Datacenter, error) {
	dc, err := f.Datacenter(ctx, "*")
	if err != nil {
		return nil, toDefaultError(err)
	}

	return dc, nil
}

// DatastoreList lists the datastores
func (f *Finder) DatastoreList(ctx context.Context, path string) ([]*object.Datastore, error) {
	es, err := f.find(ctx, f.datastoreFolder, false, path)
	if err != nil {
		return nil, err
	}

	var dss []*object.Datastore
	for _, e := range es {
		ref := e.Object.Reference()
		if ref.Type == "Datastore" {
			ds := object.NewDatastore(f.client, ref)
			ds.InventoryPath = e.Path

			dss = append(dss, ds)
		}
	}

	if len(dss) == 0 {
		return nil, &NotFoundError{"datastore", path}
	}

	return dss, nil
}

// Datastore finds a data store
func (f *Finder) Datastore(ctx context.Context, path string) (*object.Datastore, error) {
	dss, err := f.DatastoreList(ctx, path)
	if err != nil {
		return nil, err
	}

	if len(dss) > 1 {
		return nil, &MultipleFoundError{"datastore", path}
	}

	return dss[0], nil
}

// DefaultDatastore finds the default datastore
func (f *Finder) DefaultDatastore(ctx context.Context) (*object.Datastore, error) {
	ds, err := f.Datastore(ctx, "*")
	if err != nil {
		return nil, toDefaultError(err)
	}

	return ds, nil
}

// ComputeResourceList finds the computer resources
func (f *Finder) ComputeResourceList(ctx context.Context, path string) ([]*object.ComputeResource, error) {
	es, err := f.find(ctx, f.hostFolder, false, path)
	if err != nil {
		return nil, err
	}

	var crs []*object.ComputeResource
	for _, e := range es {
		var cr *object.ComputeResource

		switch o := e.Object.(type) {
		case mo.ComputeResource, mo.ClusterComputeResource:
			cr = object.NewComputeResource(f.client, o.Reference())
		default:
			continue
		}

		cr.InventoryPath = e.Path
		crs = append(crs, cr)
	}

	if len(crs) == 0 {
		return nil, &NotFoundError{"compute resource", path}
	}

	return crs, nil
}

// ComputeResource represents a compute resource
func (f *Finder) ComputeResource(ctx context.Context, path string) (*object.ComputeResource, error) {
	crs, err := f.ComputeResourceList(ctx, path)
	if err != nil {
		return nil, err
	}

	if len(crs) > 1 {
		return nil, &MultipleFoundError{"compute resource", path}
	}

	return crs[0], nil
}

// DefaultComputeResource finds the default compute resource
func (f *Finder) DefaultComputeResource(ctx context.Context) (*object.ComputeResource, error) {
	cr, err := f.ComputeResource(ctx, "*")
	if err != nil {
		return nil, toDefaultError(err)
	}

	return cr, nil
}

// ClusterComputeResourceList finds the compute resource list for a cluster
func (f *Finder) ClusterComputeResourceList(ctx context.Context, path string) ([]*object.ClusterComputeResource, error) {
	es, err := f.find(ctx, f.hostFolder, false, path)
	if err != nil {
		return nil, err
	}

	var ccrs []*object.ClusterComputeResource
	for _, e := range es {
		var ccr *object.ClusterComputeResource

		switch o := e.Object.(type) {
		case mo.ClusterComputeResource:
			ccr = object.NewClusterComputeResource(f.client, o.Reference())
		default:
			continue
		}

		ccr.InventoryPath = e.Path
		ccrs = append(ccrs, ccr)
	}

	if len(ccrs) == 0 {
		return nil, &NotFoundError{"cluster", path}
	}

	return ccrs, nil
}

// ClusterComputeResource finds a compute resource in a cluster
func (f *Finder) ClusterComputeResource(ctx context.Context, path string) (*object.ClusterComputeResource, error) {
	ccrs, err := f.ClusterComputeResourceList(ctx, path)
	if err != nil {
		return nil, err
	}

	if len(ccrs) > 1 {
		return nil, &MultipleFoundError{"cluster", path}
	}

	return ccrs[0], nil
}

// HostSystemList represents a host system list
func (f *Finder) HostSystemList(ctx context.Context, path string) ([]*object.HostSystem, error) {
	es, err := f.find(ctx, f.hostFolder, false, path)
	if err != nil {
		return nil, err
	}

	var hss []*object.HostSystem
	for _, e := range es {
		var hs *object.HostSystem

		switch o := e.Object.(type) {
		case mo.HostSystem:
			hs = object.NewHostSystem(f.client, o.Reference())

			hs.InventoryPath = e.Path
			hss = append(hss, hs)
		case mo.ComputeResource, mo.ClusterComputeResource:
			cr := object.NewComputeResource(f.client, o.Reference())

			cr.InventoryPath = e.Path

			hosts, err := cr.Hosts(ctx)
			if err != nil {
				return nil, err
			}

			hss = append(hss, hosts...)
		}
	}

	if len(hss) == 0 {
		return nil, &NotFoundError{"host", path}
	}

	return hss, nil
}

// HostSystem represents a host system
func (f *Finder) HostSystem(ctx context.Context, path string) (*object.HostSystem, error) {
	hss, err := f.HostSystemList(ctx, path)
	if err != nil {
		return nil, err
	}

	if len(hss) > 1 {
		return nil, &MultipleFoundError{"host", path}
	}

	return hss[0], nil
}

// DefaultHostSystem finds the default host system
func (f *Finder) DefaultHostSystem(ctx context.Context) (*object.HostSystem, error) {
	hs, err := f.HostSystem(ctx, "*/*")
	if err != nil {
		return nil, toDefaultError(err)
	}

	return hs, nil
}

// NetworkList returns the network list
func (f *Finder) NetworkList(ctx context.Context, path string) ([]object.NetworkReference, error) {
	es, err := f.find(ctx, f.networkFolder, false, path)
	if err != nil {
		return nil, err
	}

	var ns []object.NetworkReference
	for _, e := range es {
		ref := e.Object.Reference()
		switch ref.Type {
		case "Network":
			r := object.NewNetwork(f.client, ref)
			r.InventoryPath = e.Path
			ns = append(ns, r)
		case "DistributedVirtualPortgroup":
			r := object.NewDistributedVirtualPortgroup(f.client, ref)
			r.InventoryPath = e.Path
			ns = append(ns, r)
		case "DistributedVirtualSwitch", "VmwareDistributedVirtualSwitch":
			r := object.NewDistributedVirtualSwitch(f.client, ref)
			r.InventoryPath = e.Path
			ns = append(ns, r)
		}
	}

	if len(ns) == 0 {
		return nil, &NotFoundError{"network", path}
	}

	return ns, nil
}

// Network finds a network
func (f *Finder) Network(ctx context.Context, path string) (object.NetworkReference, error) {
	networks, err := f.NetworkList(ctx, path)
	if err != nil {
		return nil, err
	}

	if len(networks) > 1 {
		return nil, &MultipleFoundError{"network", path}
	}

	return networks[0], nil
}

// DefaultNetwork finds the default network
func (f *Finder) DefaultNetwork(ctx context.Context) (object.NetworkReference, error) {
	network, err := f.Network(ctx, "*")
	if err != nil {
		return nil, toDefaultError(err)
	}

	return network, nil
}

// ResourcePoolList returns the resource pools
func (f *Finder) ResourcePoolList(ctx context.Context, path string) ([]*object.ResourcePool, error) {
	es, err := f.find(ctx, f.hostFolder, true, path)
	if err != nil {
		return nil, err
	}

	var rps []*object.ResourcePool
	for _, e := range es {
		var rp *object.ResourcePool

		switch o := e.Object.(type) {
		case mo.ResourcePool:
			rp = object.NewResourcePool(f.client, o.Reference())
			rp.InventoryPath = e.Path
			rps = append(rps, rp)
		}
	}

	if len(rps) == 0 {
		return nil, &NotFoundError{"resource pool", path}
	}

	return rps, nil
}

// ResourcePool finds a resource pool for the given path
func (f *Finder) ResourcePool(ctx context.Context, path string) (*object.ResourcePool, error) {
	rps, err := f.ResourcePoolList(ctx, path)
	if err != nil {
		return nil, err
	}

	if len(rps) > 1 {
		return nil, &MultipleFoundError{"resource pool", path}
	}

	return rps[0], nil
}

// DefaultResourcePool finds the default resource pool
func (f *Finder) DefaultResourcePool(ctx context.Context) (*object.ResourcePool, error) {
	rp, err := f.ResourcePool(ctx, "*/Resources")
	if err != nil {
		return nil, toDefaultError(err)
	}

	return rp, nil
}

// VirtualMachineList returns the list of virtual machines
func (f *Finder) VirtualMachineList(ctx context.Context, path string) ([]*object.VirtualMachine, error) {
	es, err := f.find(ctx, f.vmFolder, false, path)
	if err != nil {
		return nil, err
	}

	var vms []*object.VirtualMachine
	for _, e := range es {
		switch o := e.Object.(type) {
		case mo.VirtualMachine:
			vm := object.NewVirtualMachine(f.client, o.Reference())
			vm.InventoryPath = e.Path
			vms = append(vms, vm)
		}
	}

	if len(vms) == 0 {
		return nil, &NotFoundError{"vm", path}
	}

	return vms, nil
}

// VirtualMachine finds a virtual machine for a given path
func (f *Finder) VirtualMachine(ctx context.Context, path string) (*object.VirtualMachine, error) {
	vms, err := f.VirtualMachineList(ctx, path)
	if err != nil {
		return nil, err
	}

	if len(vms) > 1 {
		return nil, &MultipleFoundError{"vm", path}
	}

	return vms[0], nil
}

// VirtualAppList lists the virtual applications
func (f *Finder) VirtualAppList(ctx context.Context, path string) ([]*object.VirtualApp, error) {
	es, err := f.find(ctx, f.vmFolder, false, path)
	if err != nil {
		return nil, err
	}

	var apps []*object.VirtualApp
	for _, e := range es {
		switch o := e.Object.(type) {
		case mo.VirtualApp:
			app := object.NewVirtualApp(f.client, o.Reference())
			app.InventoryPath = e.Path
			apps = append(apps, app)
		}
	}

	if len(apps) == 0 {
		return nil, &NotFoundError{"app", path}
	}

	return apps, nil
}

// VirtualApp creates a virtual application
func (f *Finder) VirtualApp(ctx context.Context, path string) (*object.VirtualApp, error) {
	apps, err := f.VirtualAppList(ctx, path)
	if err != nil {
		return nil, err
	}

	if len(apps) > 1 {
		return nil, &MultipleFoundError{"app", path}
	}

	return apps[0], nil
}

// Folder finds a folder
func (f *Finder) Folder(ctx context.Context, path string) (*object.Folder, error) {
	mo, err := f.ManagedObjectList(ctx, path)
	if err != nil {
		return nil, err
	}

	if len(mo) == 0 {
		return nil, &NotFoundError{"folder", path}
	}

	if len(mo) > 1 {
		return nil, &MultipleFoundError{"folder", path}
	}

	ref := mo[0].Object.Reference()
	if ref.Type != "Folder" {
		return nil, &NotFoundError{"folder", path}
	}

	folder := object.NewFolder(f.client, ref)

	folder.InventoryPath = mo[0].Path

	return folder, nil
}
