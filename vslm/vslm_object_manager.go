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

type VslmObjectManager struct {
	vim.ManagedObjectReference
	c *Client
}

// NewVslmObjectManager returns an ObjectManager referecing the vslm VcenterVStorageObjectManager endpoint.
// This endpoint is always connected to vpxd and utilizes the global catalog to locate objects and does
// not require a datastore.  To connect to the VStorageObjectManager on the host or in vpxd use the vslm.ObjectManager type.
func NewVslmObjectManager(client *Client) *VslmObjectManager {
	mref := client.ServiceContent.VStorageObjectManager

	m := VslmObjectManager{
		ManagedObjectReference: mref,
		c:                      client,
	}

	return &m
}

func (m *VslmObjectManager) CreateDisk(ctx context.Context, spec vim.VslmCreateSpec) (*VslmTask, error) {
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

func (m *VslmObjectManager) CreateDiskFromSnapshot(ctx context.Context, id vim.ID, snapshotId vim.ID, name string,
	profile []vim.VirtualMachineProfileSpec, crypto *vim.CryptoSpec,
	path string) (*VslmTask, error) {
	req := types.VslmCreateDiskFromSnapshot_Task{
		This:       m.Reference(),
		Id:         id,
		SnapshotId: snapshotId,
		Name:       name,
		Profile:    profile,
		Crypto:     crypto,
		Path:       path,
	}

	res, err := methods.VslmCreateDiskFromSnapshot_Task(ctx, m.c, &req)
	if err != nil {
		return nil, err
	}

	return NewVslmTask(m.c, res.Returnval), nil
}


func (m *VslmObjectManager) CreateSnapshot(ctx context.Context, id vim.ID, description string) (*VslmTask, error) {
	req := types.VslmCreateSnapshot_Task{
		This:        m.Reference(),
		Id:          id,
		Description: description,
	}

	res, err := methods.VslmCreateSnapshot_Task(ctx, m.c, &req)
	if err != nil {
		return nil, err
	}

	return NewVslmTask(m.c, res.Returnval), nil
}

func (m *VslmObjectManager) DeleteSnapshot(ctx context.Context, id vim.ID, description string) (*VslmTask, error) {
	req := types.VslmCreateSnapshot_Task{
		This:        m.Reference(),
		Id:          id,
		Description: description,
	}

	res, err := methods.VslmCreateSnapshot_Task(ctx, m.c, &req)
	if err != nil {
		return nil, err
	}

	return NewVslmTask(m.c, res.Returnval), nil
}

func (m *VslmObjectManager) CloneVStorageObject(ctx context.Context, id vim.ID, spec vim.VslmCloneSpec) (*VslmTask, error) {
	req := types.VslmCloneVStorageObject_Task{
		This: m.Reference(),
		Id:   id,
		Spec: spec,
	}

	res, err := methods.VslmCloneVStorageObject_Task(ctx, m.c, &req)
	if err != nil {
		return nil, err
	}

	return NewVslmTask(m.c, res.Returnval), nil
}

func (m *VslmObjectManager) ListVStorageObjectForSpec(ctx context.Context, query []types.VslmVsoVStorageObjectQuerySpec, maxResult int32) (*types.VslmVsoVStorageObjectQueryResult, error) {
	req := types.VslmListVStorageObjectForSpec{
		This:      m.Reference(),
		Query:     query,
		MaxResult: maxResult,
	}

	res, err := methods.VslmListVStorageObjectForSpec(ctx, m.c, &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}

func (m *VslmObjectManager) AttachTagToVStorageObject(ctx context.Context, id vim.ID, category string, tag string) error {
		req := types.VslmAttachTagToVStorageObject{
		This:      m.Reference(),
		Id:     id,
		Category: category,
		Tag: tag,
	}

	_, err := methods.VslmAttachTagToVStorageObject(ctx, m.c, &req)

	return err
}

func (m *VslmObjectManager) DetachTagFromVStorageObject(ctx context.Context, id vim.ID, category string, tag string) error {
		req := types.VslmDetachTagFromVStorageObject{
		This:      m.Reference(),
		Id:     id,
		Category: category,
		Tag: tag,
	}

	_, err := methods.VslmDetachTagFromVStorageObject(ctx, m.c, &req)

	return err
}

func (m *VslmObjectManager) ListTagsAttachedToVStorageObject(ctx context.Context, id vim.ID) ([]vim.VslmTagEntry, error) {
	req := types.VslmListTagsAttachedToVStorageObject{
		This:      m.Reference(),
		Id:     id,
	}

	res, err := methods.VslmListTagsAttachedToVStorageObject(ctx, m.c, &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}

func (m *VslmObjectManager) InflateDisk(ctx context.Context, id vim.ID) (*VslmTask, error) {
		req := types.VslmInflateDisk_Task{
		This: m.Reference(),
		Id: id,
	}

	res, err := methods.VslmInflateDisk_Task(ctx, m.c, &req)
	if err != nil {
		return nil, err
	}

	return NewVslmTask(m.c, res.Returnval), nil
}

func (m *VslmObjectManager) QueryChangedDiskAreas(ctx context.Context, id vim.ID, snapshotId vim.ID, startOffset int64,
	changeId string) (*vim.DiskChangeInfo, error) {
	req := types.VslmQueryChangedDiskAreas{
		This:      m.Reference(),
		Id:     id,
		SnapshotId: snapshotId,
		StartOffset: startOffset,
	}

	res, err := methods.VslmQueryChangedDiskAreas(ctx, m.c, &req)
	if err != nil {
		return nil, err
	}

	return &res.Returnval, nil
}

