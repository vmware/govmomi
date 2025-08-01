// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package object

import (
	"context"
	"fmt"
	"path"

	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type ComputeResource struct {
	Common
}

func NewComputeResource(c *vim25.Client, ref types.ManagedObjectReference) *ComputeResource {
	return &ComputeResource{
		Common: NewCommon(c, ref),
	}
}

func (c ComputeResource) Hosts(ctx context.Context) ([]*HostSystem, error) {
	var cr mo.ComputeResource

	err := c.Properties(ctx, c.Reference(), []string{"host"}, &cr)
	if err != nil {
		return nil, err
	}

	if len(cr.Host) == 0 {
		return nil, nil
	}

	var hs []mo.HostSystem
	pc := property.DefaultCollector(c.Client())
	err = pc.Retrieve(ctx, cr.Host, []string{"name"}, &hs)
	if err != nil {
		return nil, err
	}

	var hosts []*HostSystem

	for _, h := range hs {
		host := NewHostSystem(c.Client(), h.Reference())
		host.InventoryPath = path.Join(c.InventoryPath, h.Name)
		hosts = append(hosts, host)
	}

	return hosts, nil
}

func (c ComputeResource) Datastores(ctx context.Context) ([]*Datastore, error) {
	var cr mo.ComputeResource

	err := c.Properties(ctx, c.Reference(), []string{"datastore"}, &cr)
	if err != nil {
		return nil, err
	}

	var dss []*Datastore
	for _, ref := range cr.Datastore {
		ds := NewDatastore(c.c, ref)
		dss = append(dss, ds)
	}

	return dss, nil
}

func (c ComputeResource) EnvironmentBrowser(ctx context.Context) (*EnvironmentBrowser, error) {
	var cr mo.ComputeResource

	err := c.Properties(ctx, c.Reference(), []string{"environmentBrowser"}, &cr)
	if err != nil {
		return nil, err
	}

	if cr.EnvironmentBrowser == nil {
		return nil, fmt.Errorf("%s: nil environmentBrowser", c.Reference())
	}

	return NewEnvironmentBrowser(c.c, *cr.EnvironmentBrowser), nil
}

func (c ComputeResource) ResourcePool(ctx context.Context) (*ResourcePool, error) {
	var cr mo.ComputeResource

	err := c.Properties(ctx, c.Reference(), []string{"resourcePool"}, &cr)
	if err != nil {
		return nil, err
	}

	return NewResourcePool(c.c, *cr.ResourcePool), nil
}

func (c ComputeResource) Reconfigure(ctx context.Context, spec types.BaseComputeResourceConfigSpec, modify bool) (*Task, error) {
	req := types.ReconfigureComputeResource_Task{
		This:   c.Reference(),
		Spec:   spec,
		Modify: modify,
	}

	res, err := methods.ReconfigureComputeResource_Task(ctx, c.c, &req)
	if err != nil {
		return nil, err
	}

	return NewTask(c.c, res.Returnval), nil
}
