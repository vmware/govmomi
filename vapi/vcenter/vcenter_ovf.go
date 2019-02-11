/*
Copyright (c) 2018 VMware, Inc. All Rights Reserved.

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

	"github.com/vmware/govmomi/vapi/internal"
	"github.com/vmware/govmomi/vapi/rest"
)

// DeploymentSpec is the deployment specification for the deployment
type DeploymentSpec struct {
	Name               string   `json:"name,omitempty"`
	Annotation         string   `json:"annotation,omitempty"`
	AcceptAllEULA      bool     `json:"accept_all_EULA,omitempty"`
	Flags              []string `json:"flags,omitempty"`
	DefaultDatastoreID string   `json:"default_datastore_id,omitempty"`
}

// Target is the target for the deployment
type Target struct {
	ResourcePoolID string `json:"resource_pool_id,omitempty"`
	HostID         string `json:"host_id,omitempty"`
	FolderID       string `json:"folder_id,omitempty"`
}

// Deploy contains the information to start the deployment of a library OVF
type Deploy struct {
	DeploymentSpec `json:"deployment_spec,omitempty"`
	Target         `json:"target,omitempty"`
}

// LocalizableMessage represents a localizable error
type LocalizableMessage struct {
	Args           []string `json:"args,omitempty"`
	DefaultMessage string   `json:"default_message,omitempty"`
	ID             string   `json:"id,omitempty"`
}

// Error is a SERVER error
type Error struct {
	Class    string               `json:"@class,omitempty"`
	Messages []LocalizableMessage `json:"messages,omitempty"`
}

// ParseIssue is a parse issue struct
type ParseIssue struct {
	Category     string             `json:"@classcategory,omitempty"`
	File         string             `json:"file,omitempty"`
	LineNumber   int64              `json:"line_number,omitempty"`
	ColumnNumber int64              `json:"column_number,omitempty"`
	Message      LocalizableMessage `json:"message,omitempty"`
}

// OVFError is a list of errors from create or deploy
type OVFError struct {
	Category string             `json:"category,omitempty"`
	Error    Error              `json:"error,omitempty"`
	Issues   []ParseIssue       `json:"issues,omitempty"`
	Message  LocalizableMessage `json:"message,omitempty"`
}

// DeployedResourceID is a managed object reference for a deployed resource.
type DeployedResourceID struct {
	ID   string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
}

// DeploymentError is an error that occurs when deploying and OVF from
// a library item.
type DeploymentError struct {
	Errors []OVFError `json:"errors,omitempty"`
}

// Deployment is the results from issuing a library OVF deployment
type Deployment struct {
	Succeeded  bool               `json:"succeeded,omitempty"`
	ResourceID DeployedResourceID `json:"resource_id,omitempty"`
	Error      DeploymentError    `json:"error,omitempty"`
}

// FilterRequest contains the information to start a vcenter filter call
type FilterRequest struct {
	Target `json:"target,omitempty"`
}

// FilterResponse returns information from the vcenter filter call
type FilterResponse struct {
	EULAs         []string `json:"EULAs,omitempty"`
	Annotation    string   `json:"Annotation,omitempty"`
	Name          string   `json:"name,omitempty"`
	Networks      []string `json:"Networks,omitempty"`
	StorageGroups []string `json:"storage_groups,omitempty"`
}

// Manager extends rest.Client, adding content library related methods.
type Manager struct {
	*rest.Client
}

// NewManager creates a new Manager instance with the given client.
func NewManager(client *rest.Client) *Manager {
	return &Manager{
		Client: client,
	}
}

// DeployLibraryItem deploys a library OVF
func (c *Manager) DeployLibraryItem(ctx context.Context, libraryItemID string, deploy Deploy) (Deployment, error) {
	url := internal.URL(c, internal.VCenterOVFLibraryItem).WithID(libraryItemID).WithAction("deploy")
	var res Deployment
	return res, c.Do(ctx, url.Request(http.MethodPost, deploy), &res)
}

// FilterLibraryItem deploys a library OVF
func (c *Manager) FilterLibraryItem(ctx context.Context, libraryItemID string, filter FilterRequest) (FilterResponse, error) {
	url := internal.URL(c, internal.VCenterOVFLibraryItem).WithID(libraryItemID).WithAction("filter")
	var res FilterResponse
	return res, c.Do(ctx, url.Request(http.MethodPost, filter), &res)
}
