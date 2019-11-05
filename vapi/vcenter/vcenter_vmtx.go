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

package vcenter

import (
	"context"
	"net/http"
	"path"

	"github.com/vmware/govmomi/vapi/internal"
	"github.com/vmware/govmomi/vim25/types"
)

// vcenter vm template
// The vcenter.vm_template API provides structures and services that will let its client manage VMTX template in Content Library.
// http://vmware.github.io/vsphere-automation-sdk-rest/6.7.1/index.html#SVC_com.vmware.vcenter.vm_template.library_items

// Template create spec
type Template struct {
	Description          string                `json:"description,omitempty"`
	DiskStorage          *DiskStorage          `json:"disk_storage,omitempty"`
	DiskStorageOverrides []DiskStorageOverride `json:"disk_storage_overrides,omitempty"`
	Library              string                `json:"library,omitempty"`
	Name                 string                `json:"name,omitempty"`
	Placement            *Placement            `json:"placement,omitempty"`
	SourceVM             string                `json:"source_vm,omitempty"`
	VMHomeStorage        *DiskStorage          `json:"vm_home_storage,omitempty"`
}

// TemplateInfo for a VM template contained in an existing library item
type TemplateInfo struct {
	GuestOS string `json:"guest_OS,omitempty"`
	// TODO...
}

// Placement information used to place the virtual machine template
type Placement struct {
	ResourcePool string `json:"resource_pool,omitempty"`
	Host         string `json:"host,omitempty"`
	Folder       string `json:"folder,omitempty"`
	Cluster      string `json:"cluster,omitempty"`
}

// StoragePolicy for DiskStorage
type StoragePolicy struct {
	Policy string `json:"policy,omitempty"`
	Type   string `json:"type"`
}

// DiskStorage defines the storage specification for VM files
type DiskStorage struct {
	Datastore     string         `json:"datastore,omitempty"`
	StoragePolicy *StoragePolicy `json:"storage_policy,omitempty"`
}

// DiskStorageOverride storage specification for individual disks in the virtual machine template
type DiskStorageOverride struct {
	Key   string      `json:"key"`
	Value DiskStorage `json:"value"`
}

// GuestCustomization spec to apply to the deployed VM
type GuestCustomization struct {
	Name string `json:"name,omitempty"`
}

// HardwareCustomization spec which specifies updates to the deployed VM
type HardwareCustomization struct {
	// TODO
}

// DeployTemplate specification of how a library VM template clone should be deployed.
type DeployTemplate struct {
	Description           string                 `json:"description,omitempty"`
	DiskStorage           *DiskStorage           `json:"disk_storage,omitempty"`
	DiskStorageOverrides  []DiskStorageOverride  `json:"disk_storage_overrides,omitempty"`
	GuestCustomization    *GuestCustomization    `json:"guest_customization,omitempty"`
	HardwareCustomization *HardwareCustomization `json:"hardware_customization,omitempty"`
	Name                  string                 `json:"name,omitempty"`
	Placement             *Placement             `json:"placement,omitempty"`
	PoweredOn             bool                   `json:"powered_on"`
	VMHomeStorage         *DiskStorage           `json:"vm_home_storage,omitempty"`
}

// CreateTemplate creates a library item in content library from an existing VM
func (c *Manager) CreateTemplate(ctx context.Context, vmtx Template) (string, error) {
	url := c.Resource(internal.VCenterVMTXLibraryItem)
	var res string
	spec := struct {
		Template `json:"spec"`
	}{vmtx}
	return res, c.Do(ctx, url.Request(http.MethodPost, spec), &res)
}

// DeployTemplateLibraryItem deploys a VM as a copy of the source VM template contained in the given library item
func (c *Manager) DeployTemplateLibraryItem(ctx context.Context, libraryItemID string, deploy DeployTemplate) (*types.ManagedObjectReference, error) {
	url := c.Resource(path.Join(internal.VCenterVMTXLibraryItem, libraryItemID)).WithParam("action", "deploy")
	var res string
	spec := struct {
		DeployTemplate `json:"spec"`
	}{deploy}
	err := c.Do(ctx, url.Request(http.MethodPost, spec), &res)
	if err != nil {
		return nil, err
	}
	return &types.ManagedObjectReference{Type: "VirtualMachine", Value: res}, nil
}
