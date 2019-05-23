package vslm

import (
	"context"

	//"github.com/vmware/govmomi/vim25"
	//	vim "github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/object"
	vim "github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vslm/methods"
	"github.com/vmware/govmomi/vslm/types"
)

type VslmTask struct {
	c *Client
	object.Common
	vim.ManagedObjectReference
}

func NewVslmTask(client *Client, mref vim.ManagedObjectReference) *VslmTask {
	m := VslmTask{
		ManagedObjectReference: mref,
		c:                      client,
	}
	return &m
}

func (t *VslmTask) QueryResult(ctx context.Context) (vim.AnyType, error) {
	req := types.VslmQueryTaskResult{}
	res, err := methods.VslmQueryTaskResult(ctx, t.c, &req)
	if err != nil {
		return nil, err
	}
	return &res.Returnval, nil
}

func (t *VslmTask) QueryInfo(ctx context.Context) (*types.VslmTaskInfo, error) {
	req := types.VslmQueryInfo{}
	res, err := methods.VslmQueryInfo(ctx, t.c, &req)
	if err != nil {
		return nil, err
	}
	return &res.Returnval, nil
}

func (t *VslmTask) Cancel(ctx context.Context) error {
	req := types.VslmCancelTask{}
	_, err := methods.VslmCancelTask(ctx, t.c, &req)
	return err
}

type VSLMObjectManager struct {
	vim.ManagedObjectReference
	c *Client
}

// NewVSLMObjectManager returns an ObjectManager referecing the vslm VcenterVStorageObjectManager endpoint.
// This endpoint is always connected to vpxd and utilizes the global catalog to locate objects and does
// not require a datastore.  To connect to the VStorageObjectManager on the host or in vpxd use the vslm.ObjectManager type.
func NewVSLMObjectManager(client *Client) *VSLMObjectManager {
	mref := client.ServiceContent.VStorageObjectManager

	m := VSLMObjectManager{
		ManagedObjectReference: mref,
		c:                      client,
	}

	return &m
}

func (m *VSLMObjectManager) CreateDisk(ctx context.Context, spec vim.VslmCreateSpec) (*VslmTask, error) {
	req := types.VslmCreateDisk_Task{
		This: m.Reference(),
		Spec: spec,
	}

	res, err := methods.VslmCreateDisk_Task(ctx, m.c, &req)
	if err != nil {
		return nil, err
	}

	return NewVslmTask(m.c, res.Returnval), nil
}

func (m *VSLMObjectManager) ListVStorageObjectForSpec(ctx context.Context, query []types.VslmVsoVStorageObjectQuerySpec, maxResult int32) (*types.VslmVsoVStorageObjectQueryResult, error) {
	req := types.VslmListVStorageObjectForSpec{
		This: m.Reference(),
		Query: query,
		MaxResult: maxResult,
	}

	res, err := methods.VslmListVStorageObjectForSpec(ctx, m.c, &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}
