// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package object

import (
	"context"

	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type Folder struct {
	Common
}

func NewFolder(c *vim25.Client, ref types.ManagedObjectReference) *Folder {
	return &Folder{
		Common: NewCommon(c, ref),
	}
}

func NewRootFolder(c *vim25.Client) *Folder {
	f := NewFolder(c, c.ServiceContent.RootFolder)
	f.InventoryPath = "/"
	return f
}

func (f Folder) Children(ctx context.Context) ([]Reference, error) {
	var mf mo.Folder

	err := f.Properties(ctx, f.Reference(), []string{"childEntity"}, &mf)
	if err != nil {
		return nil, err
	}

	var rs []Reference
	for _, e := range mf.ChildEntity {
		if r := NewReference(f.c, e); r != nil {
			rs = append(rs, r)
		}
	}

	return rs, nil
}

func (f Folder) CreateDatacenter(ctx context.Context, datacenter string) (*Datacenter, error) {
	req := types.CreateDatacenter{
		This: f.Reference(),
		Name: datacenter,
	}

	res, err := methods.CreateDatacenter(ctx, f.c, &req)
	if err != nil {
		return nil, err
	}

	// Response will be nil if this is an ESX host that does not belong to a vCenter
	if res == nil {
		return nil, nil
	}

	return NewDatacenter(f.c, res.Returnval), nil
}

func (f Folder) CreateCluster(ctx context.Context, cluster string, spec types.ClusterConfigSpecEx) (*ClusterComputeResource, error) {
	req := types.CreateClusterEx{
		This: f.Reference(),
		Name: cluster,
		Spec: spec,
	}

	res, err := methods.CreateClusterEx(ctx, f.c, &req)
	if err != nil {
		return nil, err
	}

	// Response will be nil if this is an ESX host that does not belong to a vCenter
	if res == nil {
		return nil, nil
	}

	return NewClusterComputeResource(f.c, res.Returnval), nil
}

func (f Folder) CreateFolder(ctx context.Context, name string) (*Folder, error) {
	req := types.CreateFolder{
		This: f.Reference(),
		Name: name,
	}

	res, err := methods.CreateFolder(ctx, f.c, &req)
	if err != nil {
		return nil, err
	}

	return NewFolder(f.c, res.Returnval), err
}

func (f Folder) CreateStoragePod(ctx context.Context, name string) (*StoragePod, error) {
	req := types.CreateStoragePod{
		This: f.Reference(),
		Name: name,
	}

	res, err := methods.CreateStoragePod(ctx, f.c, &req)
	if err != nil {
		return nil, err
	}

	return NewStoragePod(f.c, res.Returnval), err
}

func (f Folder) AddStandaloneHost(ctx context.Context, spec types.HostConnectSpec, addConnected bool, license *string, compResSpec *types.BaseComputeResourceConfigSpec) (*Task, error) {
	req := types.AddStandaloneHost_Task{
		This:         f.Reference(),
		Spec:         spec,
		AddConnected: addConnected,
	}

	if license != nil {
		req.License = *license
	}

	if compResSpec != nil {
		req.CompResSpec = *compResSpec
	}

	res, err := methods.AddStandaloneHost_Task(ctx, f.c, &req)
	if err != nil {
		return nil, err
	}

	return NewTask(f.c, res.Returnval), nil
}

func (f Folder) CreateVM(ctx context.Context, config types.VirtualMachineConfigSpec, pool *ResourcePool, host *HostSystem) (*Task, error) {
	req := types.CreateVM_Task{
		This:   f.Reference(),
		Config: config,
		Pool:   pool.Reference(),
	}

	if host != nil {
		ref := host.Reference()
		req.Host = &ref
	}

	res, err := methods.CreateVM_Task(ctx, f.c, &req)
	if err != nil {
		return nil, err
	}

	return NewTask(f.c, res.Returnval), nil
}

func (f Folder) RegisterVM(ctx context.Context, path string, name string, asTemplate bool, pool *ResourcePool, host *HostSystem) (*Task, error) {
	req := types.RegisterVM_Task{
		This:       f.Reference(),
		Path:       path,
		AsTemplate: asTemplate,
	}

	if name != "" {
		req.Name = name
	}

	if host != nil {
		ref := host.Reference()
		req.Host = &ref
	}

	if pool != nil {
		ref := pool.Reference()
		req.Pool = &ref
	}

	res, err := methods.RegisterVM_Task(ctx, f.c, &req)
	if err != nil {
		return nil, err
	}

	return NewTask(f.c, res.Returnval), nil
}

func (f Folder) CreateDVS(ctx context.Context, spec types.DVSCreateSpec) (*Task, error) {
	req := types.CreateDVS_Task{
		This: f.Reference(),
		Spec: spec,
	}

	res, err := methods.CreateDVS_Task(ctx, f.c, &req)
	if err != nil {
		return nil, err
	}

	return NewTask(f.c, res.Returnval), nil
}

func (f Folder) MoveInto(ctx context.Context, list []types.ManagedObjectReference) (*Task, error) {
	req := types.MoveIntoFolder_Task{
		This: f.Reference(),
		List: list,
	}

	res, err := methods.MoveIntoFolder_Task(ctx, f.c, &req)
	if err != nil {
		return nil, err
	}

	return NewTask(f.c, res.Returnval), nil
}

func (f Folder) PlaceVmsXCluster(ctx context.Context, spec types.PlaceVmsXClusterSpec) (*types.PlaceVmsXClusterResult, error) {
	req := types.PlaceVmsXCluster{
		This:          f.Reference(),
		PlacementSpec: spec,
	}

	res, err := methods.PlaceVmsXCluster(ctx, f.c, &req)
	if err != nil {
		return nil, err
	}

	return &res.Returnval, nil
}
