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

type ClusterComputeResource struct {
	ComputeResource
}

func NewClusterComputeResource(c *vim25.Client, ref types.ManagedObjectReference) *ClusterComputeResource {
	return &ClusterComputeResource{
		ComputeResource: *NewComputeResource(c, ref),
	}
}

func (c ClusterComputeResource) Configuration(ctx context.Context) (*types.ClusterConfigInfoEx, error) {
	var obj mo.ClusterComputeResource

	err := c.Properties(ctx, c.Reference(), []string{"configurationEx"}, &obj)
	if err != nil {
		return nil, err
	}

	return obj.ConfigurationEx.(*types.ClusterConfigInfoEx), nil
}

func (c ClusterComputeResource) AddHost(ctx context.Context, spec types.HostConnectSpec, asConnected bool, license *string, resourcePool *types.ManagedObjectReference) (*Task, error) {
	req := types.AddHost_Task{
		This:        c.Reference(),
		Spec:        spec,
		AsConnected: asConnected,
	}

	if license != nil {
		req.License = *license
	}

	if resourcePool != nil {
		req.ResourcePool = resourcePool
	}

	res, err := methods.AddHost_Task(ctx, c.c, &req)
	if err != nil {
		return nil, err
	}

	return NewTask(c.c, res.Returnval), nil
}

func (c ClusterComputeResource) MoveInto(ctx context.Context, hosts ...*HostSystem) (*Task, error) {
	req := types.MoveInto_Task{
		This: c.Reference(),
	}

	hostReferences := make([]types.ManagedObjectReference, len(hosts))
	for i, host := range hosts {
		hostReferences[i] = host.Reference()
	}
	req.Host = hostReferences

	res, err := methods.MoveInto_Task(ctx, c.c, &req)
	if err != nil {
		return nil, err
	}

	return NewTask(c.c, res.Returnval), nil
}

func (c ClusterComputeResource) PlaceVm(ctx context.Context, spec types.PlacementSpec) (*types.PlacementResult, error) {
	req := types.PlaceVm{
		This:          c.Reference(),
		PlacementSpec: spec,
	}

	res, err := methods.PlaceVm(ctx, c.c, &req)
	if err != nil {
		return nil, err
	}

	return &res.Returnval, nil
}
