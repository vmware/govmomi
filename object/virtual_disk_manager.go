// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package object

import (
	"context"

	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/types"
)

type VirtualDiskManager struct {
	Common
}

func NewVirtualDiskManager(c *vim25.Client) *VirtualDiskManager {
	m := VirtualDiskManager{
		Common: NewCommon(c, *c.ServiceContent.VirtualDiskManager),
	}

	return &m
}

// CopyVirtualDisk copies a virtual disk, performing conversions as specified in the spec.
func (m VirtualDiskManager) CopyVirtualDisk(
	ctx context.Context,
	sourceName string, sourceDatacenter *Datacenter,
	destName string, destDatacenter *Datacenter,
	destSpec types.BaseVirtualDiskSpec, force bool) (*Task, error) {

	req := types.CopyVirtualDisk_Task{
		This:       m.Reference(),
		SourceName: sourceName,
		DestName:   destName,
		DestSpec:   destSpec,
		Force:      types.NewBool(force),
	}

	if sourceDatacenter != nil {
		ref := sourceDatacenter.Reference()
		req.SourceDatacenter = &ref
	}

	if destDatacenter != nil {
		ref := destDatacenter.Reference()
		req.DestDatacenter = &ref
	}

	res, err := methods.CopyVirtualDisk_Task(ctx, m.c, &req)
	if err != nil {
		return nil, err
	}

	return NewTask(m.c, res.Returnval), nil
}

// CreateVirtualDisk creates a new virtual disk.
func (m VirtualDiskManager) CreateVirtualDisk(
	ctx context.Context,
	name string, datacenter *Datacenter,
	spec types.BaseVirtualDiskSpec) (*Task, error) {

	req := types.CreateVirtualDisk_Task{
		This: m.Reference(),
		Name: name,
		Spec: spec,
	}

	if datacenter != nil {
		ref := datacenter.Reference()
		req.Datacenter = &ref
	}

	res, err := methods.CreateVirtualDisk_Task(ctx, m.c, &req)
	if err != nil {
		return nil, err
	}

	return NewTask(m.c, res.Returnval), nil
}

// ExtendVirtualDisk extends an existing virtual disk.
func (m VirtualDiskManager) ExtendVirtualDisk(
	ctx context.Context,
	name string, datacenter *Datacenter,
	capacityKb int64,
	eagerZero *bool) (*Task, error) {

	req := types.ExtendVirtualDisk_Task{
		This:          m.Reference(),
		Name:          name,
		NewCapacityKb: capacityKb,
		EagerZero:     eagerZero,
	}

	if datacenter != nil {
		ref := datacenter.Reference()
		req.Datacenter = &ref
	}

	res, err := methods.ExtendVirtualDisk_Task(ctx, m.c, &req)
	if err != nil {
		return nil, err
	}

	return NewTask(m.c, res.Returnval), nil
}

// MoveVirtualDisk moves a virtual disk.
func (m VirtualDiskManager) MoveVirtualDisk(
	ctx context.Context,
	sourceName string, sourceDatacenter *Datacenter,
	destName string, destDatacenter *Datacenter,
	force bool) (*Task, error) {
	req := types.MoveVirtualDisk_Task{
		This:       m.Reference(),
		SourceName: sourceName,
		DestName:   destName,
		Force:      types.NewBool(force),
	}

	if sourceDatacenter != nil {
		ref := sourceDatacenter.Reference()
		req.SourceDatacenter = &ref
	}

	if destDatacenter != nil {
		ref := destDatacenter.Reference()
		req.DestDatacenter = &ref
	}

	res, err := methods.MoveVirtualDisk_Task(ctx, m.c, &req)
	if err != nil {
		return nil, err
	}

	return NewTask(m.c, res.Returnval), nil
}

// DeleteVirtualDisk deletes a virtual disk.
func (m VirtualDiskManager) DeleteVirtualDisk(ctx context.Context, name string, dc *Datacenter) (*Task, error) {
	req := types.DeleteVirtualDisk_Task{
		This: m.Reference(),
		Name: name,
	}

	if dc != nil {
		ref := dc.Reference()
		req.Datacenter = &ref
	}

	res, err := methods.DeleteVirtualDisk_Task(ctx, m.c, &req)
	if err != nil {
		return nil, err
	}

	return NewTask(m.c, res.Returnval), nil
}

// InflateVirtualDisk inflates a virtual disk.
func (m VirtualDiskManager) InflateVirtualDisk(ctx context.Context, name string, dc *Datacenter) (*Task, error) {
	req := types.InflateVirtualDisk_Task{
		This: m.Reference(),
		Name: name,
	}

	if dc != nil {
		ref := dc.Reference()
		req.Datacenter = &ref
	}

	res, err := methods.InflateVirtualDisk_Task(ctx, m.c, &req)
	if err != nil {
		return nil, err
	}

	return NewTask(m.c, res.Returnval), nil
}

// ShrinkVirtualDisk shrinks a virtual disk.
func (m VirtualDiskManager) ShrinkVirtualDisk(ctx context.Context, name string, dc *Datacenter, copy *bool) (*Task, error) {
	req := types.ShrinkVirtualDisk_Task{
		This: m.Reference(),
		Name: name,
		Copy: copy,
	}

	if dc != nil {
		ref := dc.Reference()
		req.Datacenter = &ref
	}

	res, err := methods.ShrinkVirtualDisk_Task(ctx, m.c, &req)
	if err != nil {
		return nil, err
	}

	return NewTask(m.c, res.Returnval), nil
}

// Queries virtual disk uuid
func (m VirtualDiskManager) QueryVirtualDiskUuid(ctx context.Context, name string, dc *Datacenter) (string, error) {
	req := types.QueryVirtualDiskUuid{
		This: m.Reference(),
		Name: name,
	}

	if dc != nil {
		ref := dc.Reference()
		req.Datacenter = &ref
	}

	res, err := methods.QueryVirtualDiskUuid(ctx, m.c, &req)
	if err != nil {
		return "", err
	}

	if res == nil {
		return "", nil
	}

	return res.Returnval, nil
}

func (m VirtualDiskManager) SetVirtualDiskUuid(ctx context.Context, name string, dc *Datacenter, uuid string) error {
	req := types.SetVirtualDiskUuid{
		This: m.Reference(),
		Name: name,
		Uuid: uuid,
	}

	if dc != nil {
		ref := dc.Reference()
		req.Datacenter = &ref
	}

	_, err := methods.SetVirtualDiskUuid(ctx, m.c, &req)
	return err
}
