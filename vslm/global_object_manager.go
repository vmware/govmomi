/*
Copyright (c) 2019 VMware, Inc. All Rights Reserved.

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

package vslm

import (
	"context"
	"time"

	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	vim "github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vslm/methods"
	"github.com/vmware/govmomi/vslm/types"
)

type Task struct {
	c *Client
	object.Common
	vim.ManagedObjectReference
}

func NewTask(client *Client, mref vim.ManagedObjectReference) *Task {
	m := Task{
		ManagedObjectReference: mref,
		c:                      client,
	}
	return &m
}

func (this *Task) QueryResult(ctx context.Context) (vim.AnyType, error) {
	req := types.VslmQueryTaskResult{
		This: this.ManagedObjectReference,
	}
	res, err := methods.VslmQueryTaskResult(ctx, this.c, &req)
	if err != nil {
		return nil, err
	}
	return res.Returnval, nil
}

func (this *Task) QueryInfo(ctx context.Context) (*types.VslmTaskInfo, error) {
	req := types.VslmQueryInfo{
		This: this.ManagedObjectReference,
	}
	res, err := methods.VslmQueryInfo(ctx, this.c, &req)
	if err != nil {
		return nil, err
	}
	return &res.Returnval, nil
}

func (this *Task) Cancel(ctx context.Context) error {
	req := types.VslmCancelTask{
		This: this.ManagedObjectReference,
	}
	_, err := methods.VslmCancelTask(ctx, this.c, &req)
	return err
}

func (this *Task) Wait(ctx context.Context, timeoutNS time.Duration) (vim.AnyType, error) {
	return this.WaitNonDefault(ctx, timeoutNS, 10000, true, 10000000)
}

func (this *Task) WaitNonDefault(ctx context.Context, timeoutNS time.Duration, startIntervalNS time.Duration,
	exponential bool, maxIntervalNS time.Duration) (vim.AnyType, error) {
	waitIntervalNS := startIntervalNS
	startTimeNS := time.Now()
	for time.Now().Sub(startTimeNS) < timeoutNS {
		info, err := this.QueryInfo(ctx)
		if err != nil {
			return nil, err
		}
		if info.State == types.VslmTaskInfoStateQueued || info.State == types.VslmTaskInfoStateRunning {
			time.Sleep(waitIntervalNS)
			if exponential {
				waitIntervalNS = waitIntervalNS * 2
				if maxIntervalNS > 0 && waitIntervalNS > maxIntervalNS {
					waitIntervalNS = maxIntervalNS
				}
			}
		} else if info.State == types.VslmTaskInfoStateError {
			return nil, soap.WrapVimFault(info.Error.Fault)
		} else {
			break
		}
		// Check context here to see if we should exit
	}
	return this.QueryResult(ctx)
}

type GlobalObjectManager struct {
	vim.ManagedObjectReference
	c *Client
}

// NewGlobalObjectManager returns an ObjectManager referecing the vslm VcenterVStorageObjectManager endpoint.
// This endpoint is always connected to vpxd and utilizes the global catalog to locate objects and does
// not require a datastore.  To connect to the VStorageObjectManager on the host or in vpxd use the
// vslm.ObjectManager type.
func NewGlobalObjectManager(client *Client) *GlobalObjectManager {
	mref := client.ServiceContent.VStorageObjectManager

	m := GlobalObjectManager{
		ManagedObjectReference: mref,
		c:                      client,
	}

	return &m
}

func (this *GlobalObjectManager) CreateDisk(ctx context.Context, spec vim.VslmCreateSpec) (*Task, error) {
	req := types.VslmCreateDisk_Task{
		This: this.Reference(),
		Spec: spec,
	}

	res, err := methods.VslmCreateDisk_Task(ctx, this.c, &req)
	if err != nil {
		return nil, err
	}

	return NewTask(this.c, res.Returnval), nil
}

func (this *GlobalObjectManager) RegisterDisk(ctx context.Context, path string, name string) (*vim.VStorageObject, error) {
	req := types.VslmRegisterDisk{
		This: this.Reference(),
		Path: path,
		Name: name,
	}

	res, err := methods.VslmRegisterDisk(ctx, this.c, &req)
	if err != nil {
		return nil, err
	}

	return &res.Returnval, nil
}

func (this *GlobalObjectManager) ExtendDisk(ctx context.Context, id vim.ID, newCapacityInMB int64) (*Task, error) {
	req := types.VslmExtendDisk_Task{
		This:            this.Reference(),
		Id:              id,
		NewCapacityInMB: newCapacityInMB,
	}

	res, err := methods.VslmExtendDisk_Task(ctx, this.c, &req)
	if err != nil {
		return nil, err
	}

	return NewTask(this.c, res.Returnval), nil
}

func (this *GlobalObjectManager) InflateDisk(ctx context.Context, id vim.ID) (*Task, error) {
	req := types.VslmInflateDisk_Task{
		This: this.Reference(),
		Id:   id,
	}

	res, err := methods.VslmInflateDisk_Task(ctx, this.c, &req)
	if err != nil {
		return nil, err
	}

	return NewTask(this.c, res.Returnval), nil
}

func (this *GlobalObjectManager) Rename(ctx context.Context, id vim.ID, name string) error {
	req := types.VslmRenameVStorageObject{
		This: this.Reference(),
		Id:   id,
		Name: name,
	}

	_, err := methods.VslmRenameVStorageObject(ctx, this.c, &req)

	return err
}

func (this *GlobalObjectManager) UpdatePolicy(ct context.Context, id vim.ID, profile []vim.VirtualMachineProfileSpec) (
	*Task, error) {
	req := types.VslmUpdateVstorageObjectPolicy_Task{
		This:    this.Reference(),
		Id:      id,
		Profile: profile,
	}

	res, err := methods.VslmUpdateVstorageObjectPolicy_Task(ct, this.c, &req)

	if err != nil {
		return nil, err
	}

	return NewTask(this.c, res.Returnval), nil
}

func (this *GlobalObjectManager) UpdateInfrastructurePolicy(ct context.Context,
	spec vim.VslmInfrastructureObjectPolicySpec) (*Task, error) {
	req := types.VslmUpdateVStorageInfrastructureObjectPolicy_Task{
		This: this.Reference(),
		Spec: spec,
	}

	res, err := methods.VslmUpdateVStorageInfrastructureObjectPolicy_Task(ct, this.c, &req)

	if err != nil {
		return nil, err
	}

	return NewTask(this.c, res.Returnval), nil
}

func (this *GlobalObjectManager) RetrieveInfrastructurePolicy(ct context.Context, datastore mo.Reference) (
	[]vim.VslmInfrastructureObjectPolicy, error) {
	req := types.VslmRetrieveVStorageInfrastructureObjectPolicy{
		This:      this.Reference(),
		Datastore: datastore.Reference(),
	}

	res, err := methods.VslmRetrieveVStorageInfrastructureObjectPolicy(ct, this.c, &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}

func (this *GlobalObjectManager) Delete(ctx context.Context, id vim.ID) (*Task, error) {
	req := types.VslmDeleteVStorageObject_Task{
		This: this.Reference(),
		Id:   id,
	}

	res, err := methods.VslmDeleteVStorageObject_Task(ctx, this.c, &req)
	if err != nil {
		return nil, err
	}

	return NewTask(this.c, res.Returnval), nil
}

func (this *GlobalObjectManager) Retrieve(ctx context.Context, id vim.ID) (*vim.VStorageObject, error) {
	req := types.VslmRetrieveVStorageObject{
		This: this.Reference(),
		Id:   id,
	}

	res, err := methods.VslmRetrieveVStorageObject(ctx, this.c, &req)
	if err != nil {
		return nil, err
	}

	return &res.Returnval, nil
}

func (this *GlobalObjectManager) RetrieveState(ct context.Context, id vim.ID) (*vim.VStorageObjectStateInfo, error) {
	req := types.VslmRetrieveVStorageObjectState{
		This: this.Reference(),
		Id:   id,
	}

	res, err := methods.VslmRetrieveVStorageObjectState(ct, this.c, &req)
	if err != nil {
		return nil, err
	}

	return &res.Returnval, nil
}

func (this *GlobalObjectManager) RetrieveAssociations(ct context.Context, ids []vim.ID) (
	[]types.VslmVsoVStorageObjectAssociations, error) {
	req := types.VslmRetrieveVStorageObjectAssociations{
		This: this.Reference(),
		Ids:  ids,
	}

	res, err := methods.VslmRetrieveVStorageObjectAssociations(ct, this.c, &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}

func (this *GlobalObjectManager) ListObjectsForSpec(ctx context.Context, query []types.VslmVsoVStorageObjectQuerySpec,
	maxResult int32) (*types.VslmVsoVStorageObjectQueryResult, error) {
	req := types.VslmListVStorageObjectForSpec{
		This:      this.Reference(),
		Query:     query,
		MaxResult: maxResult,
	}

	res, err := methods.VslmListVStorageObjectForSpec(ctx, this.c, &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}

func (this *GlobalObjectManager) Clone(ctx context.Context, id vim.ID, spec vim.VslmCloneSpec) (*Task, error) {
	req := types.VslmCloneVStorageObject_Task{
		This: this.Reference(),
		Id:   id,
		Spec: spec,
	}

	res, err := methods.VslmCloneVStorageObject_Task(ctx, this.c, &req)
	if err != nil {
		return nil, err
	}

	return NewTask(this.c, res.Returnval), nil
}

func (this *GlobalObjectManager) Relocate(ct context.Context, id vim.ID, spec vim.VslmRelocateSpec) (*Task, error) {
	req := types.VslmRelocateVStorageObject_Task{
		This: this.Reference(),
		Id:   id,
		Spec: spec,
	}

	res, err := methods.VslmRelocateVStorageObject_Task(ct, this.c, &req)

	if err != nil {
		return nil, err
	}

	return NewTask(this.c, res.Returnval), nil
}

func (this *GlobalObjectManager) SetControlFlags(ct context.Context, controlFlags []string) error {
	req := types.VslmSetVStorageObjectControlFlags{
		This:         this.Reference(),
		ControlFlags: controlFlags,
	}

	_, err := methods.VslmSetVStorageObjectControlFlags(ct, this.c, &req)
	if err != nil {
		return err
	}

	return nil
}

func (this *GlobalObjectManager) ClearControlFlags(ct context.Context) error {
	req := types.VslmClearVStorageObjectControlFlags{
		This: this.Reference(),
	}

	_, err := methods.VslmClearVStorageObjectControlFlags(ct, this.c, &req)

	return err
}

func (this *GlobalObjectManager) AttachTag(ctx context.Context, id vim.ID, category string, tag string) error {
	req := types.VslmAttachTagToVStorageObject{
		This:     this.Reference(),
		Id:       id,
		Category: category,
		Tag:      tag,
	}

	_, err := methods.VslmAttachTagToVStorageObject(ctx, this.c, &req)

	return err
}

func (this *GlobalObjectManager) DetachTag(ctx context.Context, id vim.ID, category string, tag string) error {
	req := types.VslmDetachTagFromVStorageObject{
		This:     this.Reference(),
		Id:       id,
		Category: category,
		Tag:      tag,
	}

	_, err := methods.VslmDetachTagFromVStorageObject(ctx, this.c, &req)

	return err
}

func (this *GlobalObjectManager) ListObjectsAttachedToTag(ctx context.Context, id vim.ID, category string, tag string) (
	[]vim.ID, error) {
	req := types.VslmListVStorageObjectsAttachedToTag{
		This:     this.Reference(),
		Category: category,
		Tag:      tag,
	}

	res, err := methods.VslmListVStorageObjectsAttachedToTag(ctx, this.c, &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, err
}

func (this *GlobalObjectManager) ListAttachedTags(ctx context.Context, id vim.ID) ([]vim.VslmTagEntry, error) {
	req := types.VslmListTagsAttachedToVStorageObject{
		This: this.Reference(),
		Id:   id,
	}

	res, err := methods.VslmListTagsAttachedToVStorageObject(ctx, this.c, &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}

func (this *GlobalObjectManager) ReconcileDatastoreInventory(ctx context.Context, datastore mo.Reference) (*Task, error) {
	req := &types.VslmReconcileDatastoreInventory_Task{
		This:      this.Reference(),
		Datastore: datastore.Reference(),
	}

	res, err := methods.VslmReconcileDatastoreInventory_Task(ctx, this.c, req)
	if err != nil {
		return nil, err
	}
	return NewTask(this.c, res.Returnval), nil
}

func (this *GlobalObjectManager) ScheduleReconcileDatastoreInventory(ctx context.Context, datastore mo.Reference) error {
	req := &types.VslmScheduleReconcileDatastoreInventory{
		This:      this.Reference(),
		Datastore: datastore.Reference(),
	}

	_, err := methods.VslmScheduleReconcileDatastoreInventory(ctx, this.c, req)

	return err
}

func (this *GlobalObjectManager) CreateSnapshot(ctx context.Context, id vim.ID, description string) (*Task, error) {
	req := types.VslmCreateSnapshot_Task{
		This:        this.Reference(),
		Id:          id,
		Description: description,
	}

	res, err := methods.VslmCreateSnapshot_Task(ctx, this.c, &req)
	if err != nil {
		return nil, err
	}

	return NewTask(this.c, res.Returnval), nil
}

func (this *GlobalObjectManager) DeleteSnapshot(ctx context.Context, id vim.ID, snapshotID vim.ID) (*Task, error) {
	req := types.VslmDeleteSnapshot_Task{
		This:       this.Reference(),
		Id:         id,
		SnapshotId: snapshotID,
	}

	res, err := methods.VslmDeleteSnapshot_Task(ctx, this.c, &req)
	if err != nil {
		return nil, err
	}

	return NewTask(this.c, res.Returnval), nil
}

func (this *GlobalObjectManager) RetrieveSnapshotInfo(ctx context.Context, id vim.ID) (
	[]vim.VStorageObjectSnapshotInfoVStorageObjectSnapshot, error) {
	req := types.VslmRetrieveSnapshotInfo{
		This: this.Reference(),
		Id:   id,
	}

	res, err := methods.VslmRetrieveSnapshotInfo(ctx, this.c, &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval.Snapshots, nil
}

func (this *GlobalObjectManager) CreateDiskFromSnapshot(ctx context.Context, id vim.ID, snapshotId vim.ID, name string,
	profile []vim.VirtualMachineProfileSpec, crypto *vim.CryptoSpec,
	path string) (*Task, error) {
	req := types.VslmCreateDiskFromSnapshot_Task{
		This:       this.Reference(),
		Id:         id,
		SnapshotId: snapshotId,
		Name:       name,
		Profile:    profile,
		Crypto:     crypto,
		Path:       path,
	}

	res, err := methods.VslmCreateDiskFromSnapshot_Task(ctx, this.c, &req)
	if err != nil {
		return nil, err
	}

	return NewTask(this.c, res.Returnval), nil
}

func (this *GlobalObjectManager) Revert(ctx context.Context, id vim.ID, snapshotID vim.ID) (*Task, error) {
	req := types.VslmRevertVStorageObject_Task{
		This:       this.Reference(),
		Id:         id,
		SnapshotId: snapshotID,
	}

	res, err := methods.VslmRevertVStorageObject_Task(ctx, this.c, &req)
	if err != nil {
		return nil, err
	}

	return NewTask(this.c, res.Returnval), nil
}

func (this *GlobalObjectManager) RetrieveSnapshotDetails(ctx context.Context, id vim.ID, snapshotId vim.ID) (
	*vim.VStorageObjectSnapshotDetails, error) {
	req := types.VslmRetrieveSnapshotDetails{
		This:       this.Reference(),
		Id:         id,
		SnapshotId: snapshotId,
	}

	res, err := methods.VslmRetrieveSnapshotDetails(ctx, this.c, &req)
	if err != nil {
		return nil, err
	}

	return &res.Returnval, nil
}

func (this *GlobalObjectManager) QueryChangedDiskAreas(ctx context.Context, id vim.ID, snapshotId vim.ID, startOffset int64,
	changeId string) (*vim.DiskChangeInfo, error) {
	req := types.VslmQueryChangedDiskAreas{
		This:        this.Reference(),
		Id:          id,
		SnapshotId:  snapshotId,
		StartOffset: startOffset,
		ChangeId:    changeId,
	}

	res, err := methods.VslmQueryChangedDiskAreas(ctx, this.c, &req)
	if err != nil {
		return nil, err
	}

	return &res.Returnval, nil
}

func (this *GlobalObjectManager) QueryGlobalCatalogSyncStatus(ct context.Context) ([]types.VslmDatastoreSyncStatus, error) {
	req := types.VslmQueryGlobalCatalogSyncStatus{
		This: this.Reference(),
	}

	res, err := methods.VslmQueryGlobalCatalogSyncStatus(ct, this.c, &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}

func (this *GlobalObjectManager) QueryGlobalCatalogSyncStatusForDatastore(ct context.Context, datastoreURL string) (
	*types.VslmDatastoreSyncStatus, error) {
	req := types.VslmQueryGlobalCatalogSyncStatusForDatastore{
		This:         this.Reference(),
		DatastoreURL: datastoreURL,
	}

	res, err := methods.VslmQueryGlobalCatalogSyncStatusForDatastore(ct, this.c, &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}

func (this *GlobalObjectManager) UpdateMetadata(ctx context.Context, id vim.ID, metadata []vim.KeyValue,
	deleteKeys []string) (*Task, error) {
	req := &types.VslmUpdateVStorageObjectMetadata_Task{
		This:       this.Reference(),
		Id:         id,
		Metadata:   metadata,
		DeleteKeys: deleteKeys,
	}

	res, err := methods.VslmUpdateVStorageObjectMetadata_Task(ctx, this.c, req)
	if err != nil {
		return nil, err
	}
	return NewTask(this.c, res.Returnval), nil
}

func (this *GlobalObjectManager) RetrieveMetadata(ct context.Context, id vim.ID, snapshotID *vim.ID, prefix string) (
	[]vim.KeyValue, error) {
	req := types.VslmRetrieveVStorageObjectMetadata{
		This:       this.Reference(),
		Id:         id,
		SnapshotId: snapshotID,
		Prefix:     prefix,
	}

	res, err := methods.VslmRetrieveVStorageObjectMetadata(ct, this.c, &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}

func (this *GlobalObjectManager) RetrieveMetadataValue(ct context.Context, id vim.ID, snapshotID *vim.ID, key string) (
	string, error) {
	req := types.VslmRetrieveVStorageObjectMetadataValue{
		This:       this.Reference(),
		Id:         id,
		SnapshotId: snapshotID,
		Key:        key,
	}

	res, err := methods.VslmRetrieveVStorageObjectMetadataValue(ct, this.c, &req)
	if err != nil {
		return "", err
	}

	return res.Returnval, nil
}

func (this *GlobalObjectManager) RetrieveObjects(ct context.Context, ids []vim.ID) ([]types.VslmVsoVStorageObjectResult,
	error) {
	req := types.VslmRetrieveVStorageObjects{
		This: this.Reference(),
		Ids:  ids,
	}

	res, err := methods.VslmRetrieveVStorageObjects(ct, this.c, &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}

func (this *GlobalObjectManager) AttachDisk(ct context.Context, id vim.ID, vm mo.Reference, controllerKey int32,
	unitNumber int32) (*Task, error) {
	req := types.VslmAttachDisk_Task{
		This:          this.Reference(),
		Id:            id,
		Vm:            vm.Reference(),
		ControllerKey: controllerKey,
		UnitNumber:    &unitNumber,
	}

	res, err := methods.VslmAttachDisk_Task(ct, this.c, &req)
	if err != nil {
		return nil, err
	}

	return NewTask(this.c, res.Returnval), nil
}
