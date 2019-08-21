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

var (
	CnsVolumeManagerInstance = vimtypes.ManagedObjectReference{
		Type:  "CnsVolumeManager",
		Value: "cns-volume-manager",
	}
)

type Client struct {
	vim25Client   *vim25.Client
	serviceClient *soap.Client
}

// NewClient creates a new CNS client
func NewClient(ctx context.Context, c *vim25.Client) (*Client, error) {
	sc := c.Client.NewServiceClient(Path, Namespace)
	return &Client{c, sc}, nil
}

// CreateVolume calls the CNS create API.
func (c *Client) CreateVolume(ctx context.Context, createSpecList []cnstypes.CnsVolumeCreateSpec) (*object.Task, error) {
	req := cnstypes.CnsCreateVolume{
		This:        CnsVolumeManagerInstance,
		CreateSpecs: createSpecList,
	}
	res, err := methods.CnsCreateVolume(ctx, c.serviceClient, &req)
	if err != nil {
		return nil, err
	}
	return object.NewTask(c.vim25Client, res.Returnval), nil
}

// UpdateVolumeMetadata calls the CNS CnsUpdateVolumeMetadata API with UpdateSpecs specified in the argument
func (c *Client) UpdateVolumeMetadata(ctx context.Context, updateSpecList []cnstypes.CnsVolumeMetadataUpdateSpec) (*object.Task, error) {
	req := cnstypes.CnsUpdateVolumeMetadata{
		This:        CnsVolumeManagerInstance,
		UpdateSpecs: updateSpecList,
	}
	res, err := methods.CnsUpdateVolumeMetadata(ctx, c.serviceClient, &req)
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
	res, err := methods.CnsDeleteVolume(ctx, c.serviceClient, &req)
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
	res, err := methods.CnsAttachVolume(ctx, c.serviceClient, &req)
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
	res, err := methods.CnsDetachVolume(ctx, c.serviceClient, &req)
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
	res, err := methods.CnsQueryVolume(ctx, c.serviceClient, &req)
	if err != nil {
		return nil, err
	}
	return &res.Returnval, nil
}

// QueryVolume calls the CNS QueryAllVolume API.
func (c *Client) QueryAllVolume(ctx context.Context, queryFilter cnstypes.CnsQueryFilter, querySelection cnstypes.CnsQuerySelection) (*cnstypes.CnsQueryResult, error) {
	req := cnstypes.CnsQueryAllVolume{
		This:      CnsVolumeManagerInstance,
		Filter:    queryFilter,
		Selection: querySelection,
	}
	res, err := methods.CnsQueryAllVolume(ctx, c.serviceClient, &req)
	if err != nil {
		return nil, err
	}
	return &res.Returnval, nil
}
