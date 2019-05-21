package vslm

import (
	"context"

	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vslm/methods"
)

type VslmTask struct {
	Common
}

func NewVslmTask(client *vim25.Client, ref types.ManagedObjectReference) *VslmTask {
	t := Task{
		Common: NewCommon(c, ref),
	}

	return &t
}

type VSLMObjectManager struct {
	c *vim25.Client
}

func (t *VslmTask) QueryResult(ctx context.Context) (AnyType, error) {
	req := types.VslmQueryTaskResult{}
	res, err = methods.VslmQueryTaskResult(ctx, t.Reference(), &req)
	if err != nil {
		return nil, err
	}
	return &res.Returnval, nil
}

func (t *VslmTask) QueryInfo(ctx context.Context) (VslmTaskInfo, error) {
	req := types.VslmQueryInfo{}
	res, err = methods.VslmQueryInfo(ctx, t.Reference(), &req)
	if err != nil {
		return nil, err
	}
	return &res.Returnval, nil
}

func (t* VslmTask) Cancel(ctx context.Context) error {
	req := types.VslmCancelTask{}
	_, err = methods.VslmCancelTask(ctx, t.Reference(), &req)
	return err
}

// NewVSLMObjectManager returns an ObjectManager referecing the vslm VcenterVStorageObjectManager endpoint.
// This endpoint is always connected to vpxd and utilizes the global catalog to locate objects and does
// not require a datastore.  To connect to the VStorageObjectManager on the host or in vpxd use the vslm.ObjectManager type.
func NewVSLMObjectManager(client *vim25.Client) *VSLMObjectManager {

	m := VSLMObjectManager{
		c: client,
	}

	return &m
}

func (m ObjectManager) CreateDisk(ctx context.Context, spec types.VslmCreateSpec) (*types.VslmCreateDisk_TaskResponse, error) {
	req := types.VslmCreateDisk_Task{
		This: m.Reference(),
		Spec: spec,
	}

	res, err := methods.VslmCreateDisk_Task(ctx, m.c, &req)
	if err != nil {
		return nil, err
	}

	return object.NewTask(m.c, res.Returnval), nil

}
