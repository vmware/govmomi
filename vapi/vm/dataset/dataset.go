// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package dataset

import (
	"context"
	"net/http"
	"net/url"
	"path"
	"strconv"

	"github.com/vmware/govmomi/vapi/rest"
	"github.com/vmware/govmomi/vapi/vm/internal"
)

// Manager extends rest.Client, adding data set related methods.
//
// Data sets functionality was introduced in vSphere 8.0,
// and requires the VM to have virtual hardware version 20 or newer.
//
// See the VMware Guest SDK Programming Guide for details on using data sets
// from within the guest OS of a VM.
//
// See https://developer.vmware.com/apis/vsphere-automation/latest/vcenter/vm/data_sets/
type Manager struct {
	*rest.Client
}

// NewManager creates a new Manager instance with the given client.
func NewManager(client *rest.Client) *Manager {
	return &Manager{
		Client: client,
	}
}

// Access permission to the entries of a data set.
type Access string

const (
	AccessNone      = Access("NONE")
	AccessReadOnly  = Access("READ_ONLY")
	AccessReadWrite = Access("READ_WRITE")
)

// Describes a data set to be created.
type CreateSpec struct {
	// Name should take the form "com.company.project" to avoid conflict with other uses.
	// Must not be empty.
	Name string `json:"name"`

	// Description of the data set.
	Description string `json:"description"`

	// Host controls access to the data set entries by the ESXi host and the vCenter.
	// For example, if the host access is set to NONE, the entries of this data set
	// will not be accessible through the vCenter API.
	// Must not be empty.
	Host Access `json:"host"`

	// Guest controls access to the data set entries by the guest OS of the VM (i.e. in-guest APIs).
	// For example, if the guest access is set to READ_ONLY, it will be forbidden
	// to create, delete, and update entries in this data set via the VMware Guest SDK.
	// Must not be empty.
	Guest Access `json:"guest"`

	// OmitFromSnapshotAndClone controls whether the data set is included in snapshots and clones of the VM.
	// When a VM is reverted to a snapshot, any data set with OmitFromSnapshotAndClone=true will be destroyed.
	// Default is false.
	OmitFromSnapshotAndClone *bool `json:"omit_from_snapshot_and_clone,omitempty"`
}

// Describes modifications to a data set.
type UpdateSpec struct {
	Description              *string `json:"description,omitempty"`
	Host                     *Access `json:"host,omitempty"`
	Guest                    *Access `json:"guest,omitempty"`
	OmitFromSnapshotAndClone *bool   `json:"omit_from_snapshot_and_clone,omitempty"`
}

// Data set information.
type Info struct {
	Name                     string `json:"name"`
	Description              string `json:"description"`
	Host                     Access `json:"host"`
	Guest                    Access `json:"guest"`
	Used                     int    `json:"used"`
	OmitFromSnapshotAndClone bool   `json:"omit_from_snapshot_and_clone"`
}

// Brief data set information.
type Summary struct {
	DataSet     string `json:"data_set"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

const dataSetsPathField = "data-sets"

func dataSetPath(vm string, dataSet string) string {
	return path.Join(internal.VCenterVMPath, url.PathEscape(vm), dataSetsPathField, url.PathEscape(dataSet))
}

func dataSetsPath(vm string) string {
	return path.Join(internal.VCenterVMPath, url.PathEscape(vm), dataSetsPathField)
}

const entriesPathField = "entries"

func entryPath(vm string, dataSet string, key string) string {
	return path.Join(internal.VCenterVMPath, url.PathEscape(vm), dataSetsPathField, url.PathEscape(dataSet), entriesPathField, url.PathEscape(key))
}

func entriesPath(vm string, dataSet string) string {
	return path.Join(internal.VCenterVMPath, url.PathEscape(vm), dataSetsPathField, url.PathEscape(dataSet), entriesPathField)
}

// CreateDataSet creates a data set associated with the given virtual machine.
func (c *Manager) CreateDataSet(ctx context.Context, vm string, spec *CreateSpec) (string, error) {
	url := c.Resource(dataSetsPath(vm))
	var res string
	err := c.Do(ctx, url.Request(http.MethodPost, spec), &res)
	return res, err
}

// DeleteDataSet deletes an existing data set from the given virtual machine.
// The operation will fail if the data set is not empty.
// Set the force flag to delete a non-empty data set.
func (c *Manager) DeleteDataSet(ctx context.Context, vm string, dataSet string, force bool) error {
	url := c.Resource(dataSetPath(vm, dataSet))
	if force {
		url.WithParam("force", strconv.FormatBool(force))
	}
	return c.Do(ctx, url.Request(http.MethodDelete), nil)
}

// GetDataSet retrieves information about the given data set.
func (c *Manager) GetDataSet(ctx context.Context, vm string, dataSet string) (*Info, error) {
	url := c.Resource(dataSetPath(vm, dataSet))
	var res Info
	err := c.Do(ctx, url.Request(http.MethodGet), &res)
	return &res, err
}

// UpdateDataSet modifies the given data set.
func (c *Manager) UpdateDataSet(ctx context.Context, vm string, dataSet string, spec *UpdateSpec) error {
	url := c.Resource(dataSetPath(vm, dataSet))
	return c.Do(ctx, url.Request(http.MethodPatch, spec), nil)
}

// ListDataSets returns a list of brief descriptions of the data sets on with the given virtual machine.
func (c *Manager) ListDataSets(ctx context.Context, vm string) ([]Summary, error) {
	url := c.Resource(dataSetsPath(vm))
	var res []Summary
	err := c.Do(ctx, url.Request(http.MethodGet), &res)
	return res, err
}

// SetEntry creates or updates an entry in the given data set.
// If an entry with the given key already exists, it will be overwritten.
// The key can be at most 4096 bytes. The value can be at most 1MB.
func (c *Manager) SetEntry(ctx context.Context, vm string, dataSet string, key string, value string) error {
	url := c.Resource(entryPath(vm, dataSet, key))
	return c.Do(ctx, url.Request(http.MethodPut, value), nil)
}

// GetEntry returns the value of the data set entry with the given key.
func (c *Manager) GetEntry(ctx context.Context, vm string, dataSet string, key string) (string, error) {
	url := c.Resource(entryPath(vm, dataSet, key))
	var res string
	err := c.Do(ctx, url.Request(http.MethodGet), &res)
	return res, err
}

// DeleteEntry removes an existing entry from the given data set.
func (c *Manager) DeleteEntry(ctx context.Context, vm string, dataSet string, key string) error {
	url := c.Resource(entryPath(vm, dataSet, key))
	return c.Do(ctx, url.Request(http.MethodDelete), nil)
}

// ListEntries returns a list of all entry keys in the given data set.
func (c *Manager) ListEntries(ctx context.Context, vm string, dataSet string) ([]string, error) {
	url := c.Resource(entriesPath(vm, dataSet))
	var res []string
	err := c.Do(ctx, url.Request(http.MethodGet), &res)
	return res, err
}
