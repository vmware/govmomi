// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package depots

import (
	"context"
	"fmt"
	"net/http"

	"github.com/vmware/govmomi/vapi/esx/settings"
	"github.com/vmware/govmomi/vapi/rest"
)

const (
	// DepotsOfflinePath The endpoint for the offline depots API
	DepotsOfflinePath = settings.BasePath + "/depots/offline"
	// DepotsOfflineContentPath The endpoint for retrieving the components in a depot
	DepotsOfflineContentPath = DepotsOfflinePath + "/%s/content"
	// BaseImagesPath The endpoint for retrieving the list of base ESXi images
	BaseImagesPath = settings.BasePath + "/depot-content/base-images"
)

// Manager extends rest.Client, adding vLCM related methods.
type Manager struct {
	*rest.Client
}

// NewManager creates a new Manager instance with the given client.
func NewManager(client *rest.Client) *Manager {
	return &Manager{
		Client: client,
	}
}

type SourceType string

const (
	SourceTypePush = SourceType("PUSH")
	SourceTypePull = SourceType("PULL")
)

// SettingsDepotsOfflineSummary is a type mapping for
// https://developer.vmware.com/apis/vsphere-automation/latest/esx/data-structures/Settings/Depots/Offline/Summary/
type SettingsDepotsOfflineSummary struct {
	Description string `json:"description"`
	SourceType  string `json:"source_type"`
	FileId      string `json:"file_id,omitempty"`
	Location    string `json:"location,omitempty"`
	Owner       string `json:"owner,omitempty"`
	OwnerData   string `json:"owner_data,omitempty"`
}

// SettingsDepotsOfflineInfo is a type mapping for
// https://developer.vmware.com/apis/vsphere-automation/latest/esx/data-structures/Settings/Depots/Offline/Info/
type SettingsDepotsOfflineInfo struct {
	CreateTime  string `json:"create_time"`
	Description string `json:"description"`
	SourceType  string `json:"source_type"`
	FileId      string `json:"file_id,omitempty"`
	Location    string `json:"location,omitempty"`
	Owner       string `json:"owner,omitempty"`
	OwnerData   string `json:"owner_data,omitempty"`
}

// SettingsDepotsOfflineCreateSpec is a type mapping for
// https://developer.vmware.com/apis/vsphere-automation/latest/esx/data-structures/Settings/Depots/Offline/CreateSpec/
type SettingsDepotsOfflineCreateSpec struct {
	Description string `json:"description,omitempty"`
	SourceType  string `json:"source_type"`
	FileId      string `json:"file_id,omitempty"`
	Location    string `json:"location,omitempty"`
	OwnerData   string `json:"owner_data,omitempty"`
}

// SettingsDepotsComponentSummary is a type mapping for
// https://developer.vmware.com/apis/vsphere-automation/latest/esx/data-structures/Settings/Depots/ComponentSummary/
type SettingsDepotsComponentSummary struct {
	DisplayName string             `json:"display_name"`
	Versions    []ComponentVersion `json:"versions"`
}

// ComponentVersion is a type mapping for
// https://developer.vmware.com/apis/vsphere-automation/latest/esx/data-structures/Settings/Depots/ComponentVersion/
type ComponentVersion struct {
	DisplayVersion string `json:"display_version"`
	Version        string `json:"version"`
}

// SettingsDepotsMetadataInfo is a partial type mapping for
// https://developer.vmware.com/apis/vsphere-automation/latest/esx/data-structures/Settings/Depots/MetadataInfo/
type SettingsDepotsMetadataInfo struct {
	Addons                map[string]any                            `json:"addons,omitempty"`
	BaseImages            []any                                     `json:"base_images,omitempty"`
	FileName              string                                    `json:"file_name"`
	HardwareSupport       map[string]any                            `json:"hardware_support,omitempty"`
	IndependentComponents map[string]SettingsDepotsComponentSummary `json:"independent_components,omitempty"`
	Solutions             map[string]any                            `json:"solutions,omitempty"`
	Updates               map[string]any                            `json:"updates,omitempty"`
}

// SettingsDepotsOfflineContentInfo is a type mapping for
// https://developer.vmware.com/apis/vsphere-automation/latest/esx/data-structures/Settings/Depots/Offline/Content/Info/
type SettingsDepotsOfflineContentInfo struct {
	MetadataBundles map[string][]SettingsDepotsMetadataInfo `json:"metadata_bundles"`
}

// BaseImagesSummary is a type mapping for
// https://developer.vmware.com/apis/vsphere-automation/latest/esx/data-structures/Settings/DepotContent/BaseImages/Summary/
type BaseImagesSummary struct {
	DisplayName    string `json:"display_name"`
	DisplayVersion string `json:"display_version"`
	Kb             string `json:"kb"`
	ReleaseDate    string `json:"release_date"`
	Summary        string `json:"summary"`
	Version        string `json:"version"`
}

// GetOfflineDepot retrieves an offline depot by its identifier
// https://developer.vmware.com/apis/vsphere-automation/latest/esx/api/esx/settings/depots/offline/depot/get/
func (c *Manager) GetOfflineDepot(depotId string) (SettingsDepotsOfflineSummary, error) {
	path := c.Resource(DepotsOfflinePath).WithSubpath(depotId)
	req := path.Request(http.MethodGet)
	var res SettingsDepotsOfflineSummary
	return res, c.Do(context.Background(), req, &res)
}

// GetOfflineDepots retrieves all offline depots
// https://developer.vmware.com/apis/vsphere-automation/latest/esx/api/esx/settings/depots/offline/get/
func (c *Manager) GetOfflineDepots() (map[string]SettingsDepotsOfflineInfo, error) {
	path := c.Resource(DepotsOfflinePath)
	req := path.Request(http.MethodGet)
	var res map[string]SettingsDepotsOfflineInfo
	return res, c.Do(context.Background(), req, &res)
}

// DeleteOfflineDepot triggers a task to delete an offline depot by its identifier
// https://developer.vmware.com/apis/vsphere-automation/latest/esx/api/esx/settings/depots/offline/depotvmw-tasktrue/delete/
func (c *Manager) DeleteOfflineDepot(depotId string) (string, error) {
	path := c.Resource(DepotsOfflinePath).WithSubpath(depotId).WithParam("vmw-task", "true")
	req := path.Request(http.MethodDelete)
	var res string
	return res, c.Do(context.Background(), req, &res)
}

// CreateOfflineDepot triggers a task to create an offline depot
// https://developer.vmware.com/apis/vsphere-automation/latest/esx/api/esx/settings/depots/offlinevmw-tasktrue/post/
func (c *Manager) CreateOfflineDepot(spec SettingsDepotsOfflineCreateSpec) (string, error) {
	path := c.Resource(DepotsOfflinePath).WithParam("vmw-task", "true")
	req := path.Request(http.MethodPost, spec)
	var res string
	return res, c.Do(context.Background(), req, &res)
}

// GetOfflineDepotContent retrieves the contents of a depot
// https://developer.vmware.com/apis/vsphere-automation/latest/esx/api/esx/settings/depots/offline/depot/content/get/
func (c *Manager) GetOfflineDepotContent(depotId string) (SettingsDepotsOfflineContentInfo, error) {
	path := c.Resource(fmt.Sprintf(DepotsOfflineContentPath, depotId))
	req := path.Request(http.MethodGet)
	var res SettingsDepotsOfflineContentInfo
	return res, c.Do(context.Background(), req, &res)
}

// ListBaseImages retrieves the available ESXi versions
// https://developer.vmware.com/apis/vsphere-automation/latest/esx/api/esx/settings/depot-content/base-images/get/
func (c *Manager) ListBaseImages() ([]BaseImagesSummary, error) {
	path := c.Resource(BaseImagesPath)
	req := path.Request(http.MethodGet)
	var res []BaseImagesSummary
	return res, c.Do(context.Background(), req, &res)
}
