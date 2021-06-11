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

package cns

import (
	"context"

	"github.com/vmware/govmomi/cns/methods"
	cnstypes "github.com/vmware/govmomi/cns/types"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/soap"
	vimtypes "github.com/vmware/govmomi/vim25/types"
)

// Namespace and Path constants
const (
	Namespace = "vsan"
	Path      = "/vsanHealth"
)

const (
	ReleaseVSAN67u3 = "vSAN 6.7U3"
	ReleaseVSAN70   = "7.0"
	ReleaseVSAN70u1 = "vSAN 7.0U1"
)

var (
	CnsVolumeManagerInstance = vimtypes.ManagedObjectReference{
		Type:  "CnsVolumeManager",
		Value: "cns-volume-manager",
	}
)

type Client struct {
	*soap.Client

	RoundTripper soap.RoundTripper

	vim25Client *vim25.Client
}

// NewClient creates a new CNS client
func NewClient(ctx context.Context, c *vim25.Client) (*Client, error) {
	sc := c.Client.NewServiceClient(Path, Namespace)
	sc.Namespace = c.Namespace
	sc.Version = c.Version
	return &Client{sc, sc, c}, nil
}

// RoundTrip dispatches to the RoundTripper field.
func (c *Client) RoundTrip(ctx context.Context, req, res soap.HasFault) error {
	return c.RoundTripper.RoundTrip(ctx, req, res)
}

// CreateVolume calls the CNS create API.
func (c *Client) CreateVolume(ctx context.Context, createSpecList []cnstypes.CnsVolumeCreateSpec) (*object.Task, error) {
	createSpecList = dropUnknownCreateSpecElements(c, createSpecList)
	req := cnstypes.CnsCreateVolume{
		This:        CnsVolumeManagerInstance,
		CreateSpecs: createSpecList,
	}
	res, err := methods.CnsCreateVolume(ctx, c, &req)
	if err != nil {
		return nil, err
	}
	return object.NewTask(c.vim25Client, res.Returnval), nil
}

// UpdateVolumeMetadata calls the CNS CnsUpdateVolumeMetadata API with UpdateSpecs specified in the argument
func (c *Client) UpdateVolumeMetadata(ctx context.Context, updateSpecList []cnstypes.CnsVolumeMetadataUpdateSpec) (*object.Task, error) {
	updateSpecList = dropUnknownVolumeMetadataUpdateSpecElements(c, updateSpecList)
	req := cnstypes.CnsUpdateVolumeMetadata{
		This:        CnsVolumeManagerInstance,
		UpdateSpecs: updateSpecList,
	}
	res, err := methods.CnsUpdateVolumeMetadata(ctx, c, &req)
	if err != nil {
		return nil, err
	}
	return object.NewTask(c.vim25Client, res.Returnval), nil
}

// DeleteVolume calls the CNS delete API.
func (c *Client) DeleteVolume(ctx context.Context, volumeIDList []cnstypes.CnsVolumeId, deleteDisk bool) (*object.Task, error) {
	req := cnstypes.CnsDeleteVolume{
		This:       CnsVolumeManagerInstance,
		VolumeIds:  volumeIDList,
		DeleteDisk: deleteDisk,
	}
	res, err := methods.CnsDeleteVolume(ctx, c, &req)
	if err != nil {
		return nil, err
	}
	return object.NewTask(c.vim25Client, res.Returnval), nil
}

// ExtendVolume calls the CNS Extend API.
func (c *Client) ExtendVolume(ctx context.Context, extendSpecList []cnstypes.CnsVolumeExtendSpec) (*object.Task, error) {
	req := cnstypes.CnsExtendVolume{
		This:        CnsVolumeManagerInstance,
		ExtendSpecs: extendSpecList,
	}
	res, err := methods.CnsExtendVolume(ctx, c, &req)
	if err != nil {
		return nil, err
	}
	return object.NewTask(c.vim25Client, res.Returnval), nil
}

// AttachVolume calls the CNS Attach API.
func (c *Client) AttachVolume(ctx context.Context, attachSpecList []cnstypes.CnsVolumeAttachDetachSpec) (*object.Task, error) {
	req := cnstypes.CnsAttachVolume{
		This:        CnsVolumeManagerInstance,
		AttachSpecs: attachSpecList,
	}
	res, err := methods.CnsAttachVolume(ctx, c, &req)
	if err != nil {
		return nil, err
	}
	return object.NewTask(c.vim25Client, res.Returnval), nil
}

// DetachVolume calls the CNS Detach API.
func (c *Client) DetachVolume(ctx context.Context, detachSpecList []cnstypes.CnsVolumeAttachDetachSpec) (*object.Task, error) {
	req := cnstypes.CnsDetachVolume{
		This:        CnsVolumeManagerInstance,
		DetachSpecs: detachSpecList,
	}
	res, err := methods.CnsDetachVolume(ctx, c, &req)
	if err != nil {
		return nil, err
	}
	return object.NewTask(c.vim25Client, res.Returnval), nil
}

// QueryVolume calls the CNS QueryVolume API.
func (c *Client) QueryVolume(ctx context.Context, queryFilter cnstypes.CnsQueryFilter) (*cnstypes.CnsQueryResult, error) {
	req := cnstypes.CnsQueryVolume{
		This:   CnsVolumeManagerInstance,
		Filter: queryFilter,
	}
	res, err := methods.CnsQueryVolume(ctx, c, &req)
	if err != nil {
		return nil, err
	}
	return &res.Returnval, nil
}

// QueryVolumeInfo calls the CNS QueryVolumeInfo API and return a task, from which we can extract VolumeInfo
// containing VStorageObject
func (c *Client) QueryVolumeInfo(ctx context.Context, volumeIDList []cnstypes.CnsVolumeId) (*object.Task, error) {
	req := cnstypes.CnsQueryVolumeInfo{
		This:      CnsVolumeManagerInstance,
		VolumeIds: volumeIDList,
	}
	res, err := methods.CnsQueryVolumeInfo(ctx, c, &req)
	if err != nil {
		return nil, err
	}
	return object.NewTask(c.vim25Client, res.Returnval), nil
}

// QueryAllVolume calls the CNS QueryAllVolume API.
func (c *Client) QueryAllVolume(ctx context.Context, queryFilter cnstypes.CnsQueryFilter, querySelection cnstypes.CnsQuerySelection) (*cnstypes.CnsQueryResult, error) {
	req := cnstypes.CnsQueryAllVolume{
		This:      CnsVolumeManagerInstance,
		Filter:    queryFilter,
		Selection: querySelection,
	}
	res, err := methods.CnsQueryAllVolume(ctx, c, &req)
	if err != nil {
		return nil, err
	}
	return &res.Returnval, nil
}

// QueryVolumeAsync calls the CNS QueryAsync API and return a task, from which we can extract CnsQueryResult
func (c *Client) QueryVolumeAsync(ctx context.Context, queryFilter cnstypes.CnsQueryFilter, querySelection cnstypes.CnsQuerySelection) (*object.Task, error) {
	req := cnstypes.CnsQueryAsync{
		This:      CnsVolumeManagerInstance,
		Filter:    queryFilter,
		Selection: &querySelection,
	}
	res, err := methods.CnsQueryAsync(ctx, c, &req)
	if err != nil {
		return nil, err
	}
	return object.NewTask(c.vim25Client, res.Returnval), nil
}

// RelocateVolume calls the CNS Relocate API.
func (c *Client) RelocateVolume(ctx context.Context, relocateSpecs ...cnstypes.BaseCnsVolumeRelocateSpec) (*object.Task, error) {
	req := cnstypes.CnsRelocateVolume{
		This:          CnsVolumeManagerInstance,
		RelocateSpecs: relocateSpecs,
	}
	res, err := methods.CnsRelocateVolume(ctx, c, &req)
	if err != nil {
		return nil, err
	}
	return object.NewTask(c.vim25Client, res.Returnval), nil
}

// ConfigureVolumeACLs calls the CNS Configure ACL API.
func (c *Client) ConfigureVolumeACLs(ctx context.Context, aclConfigSpecs ...cnstypes.CnsVolumeACLConfigureSpec) (*object.Task, error) {
	req := cnstypes.CnsConfigureVolumeACLs{
		This:           CnsVolumeManagerInstance,
		ACLConfigSpecs: aclConfigSpecs,
	}
	res, err := methods.CnsConfigureVolumeACLs(ctx, c, &req)
	if err != nil {
		return nil, err
	}
	return object.NewTask(c.vim25Client, res.Returnval), nil
}

// CreateSnapshots calls the CNS CreateSnapshots API

func (c *Client) CreateSnapshots(ctx context.Context, snapshotCreateSpecList []cnstypes.CnsSnapshotCreateSpec) (*object.Task, error) {
	req := cnstypes.CnsCreateSnapshots{
		This:          CnsVolumeManagerInstance,
		SnapshotSpecs: snapshotCreateSpecList,
	}
	res, err := methods.CnsCreateSnapshots(ctx, c, &req)
	if err != nil {
		return nil, err
	}

	return object.NewTask(c.vim25Client, res.Returnval), nil
}

// DeleteSnapshots calls the CNS DeleteSnapshots API
func (c *Client) DeleteSnapshots(ctx context.Context, snapshotDeleteSpecList []cnstypes.CnsSnapshotDeleteSpec) (*object.Task, error) {
	req := cnstypes.CnsDeleteSnapshots{
		This:                CnsVolumeManagerInstance,
		SnapshotDeleteSpecs: snapshotDeleteSpecList,
	}
	res, err := methods.CnsDeleteSnapshots(ctx, c, &req)
	if err != nil {
		return nil, err
	}
	return object.NewTask(c.vim25Client, res.Returnval), nil
}

// QuerySnapshots calls the CNS QuerySnapshots API
func (c *Client) QuerySnapshots(ctx context.Context, snapshotQueryFilter cnstypes.CnsSnapshotQueryFilter) (*object.Task, error) {
	req := cnstypes.CnsQuerySnapshots{
		This:                CnsVolumeManagerInstance,
		SnapshotQueryFilter: snapshotQueryFilter,
	}
	res, err := methods.CnsQuerySnapshots(ctx, c, &req)
	if err != nil {
		return nil, err
	}
	return object.NewTask(c.vim25Client, res.Returnval), nil
}
