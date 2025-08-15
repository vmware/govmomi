// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package clusters

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/vmware/govmomi/vapi/esx/settings"
	"github.com/vmware/govmomi/vapi/rest"
)

const (
	// BasePath The base endpoint for the clusters API
	BasePath = settings.BasePath + "/clusters"
	// SoftwareDraftsPath The endpoint for the software drafts API
	SoftwareDraftsPath = BasePath + "/%s/software/drafts"
	// SoftwareComponentsPath The endpoint for retrieving the custom components in a software draft
	SoftwareComponentsPath = SoftwareDraftsPath + "/%s/software/components"
	// BaseImagePath The endpoint for retrieving the base image of a software draft
	BaseImagePath = SoftwareDraftsPath + "/%s/software/base-image"
	// SoftwareEnablementPath The endpoint for retrieving the vLCM status (enabled/disabled) of a cluster
	SoftwareEnablementPath = BasePath + "/%s/enablement/software"
)

// Manager extends rest.Client, adding Software Drafts related methods.
type Manager struct {
	*rest.Client
}

// NewManager creates a new Manager instance with the given client.
func NewManager(client *rest.Client) *Manager {
	return &Manager{
		Client: client,
	}
}

// SettingsClustersSoftwareDraftsMetadata is a type mapping for
// https://developer.vmware.com/apis/vsphere-automation/latest/esx/data-structures/Settings/Clusters/Software/Drafts/Metadata/
type SettingsClustersSoftwareDraftsMetadata struct {
	CreationTime string `json:"creation_time"`
	Owner        string `json:"owner"`
	Status       string `json:"status"`
}

// SettingsBaseImageDetails is a type mapping for
// https://developer.vmware.com/apis/vsphere-automation/latest/esx/data-structures/Settings/BaseImageDetails/
type SettingsBaseImageDetails struct {
	DisplayName    string `json:"display_name"`
	DisplayVersion string `json:"display_version"`
	ReleaseDate    string `json:"release_date"`
}

// SettingsBaseImageInfo is a type mapping for
// https://developer.vmware.com/apis/vsphere-automation/latest/esx/data-structures/Settings/BaseImageInfo/
type SettingsBaseImageInfo struct {
	Version string                   `json:"version"`
	Details SettingsBaseImageDetails `json:"details"`
}

// SettingsComponentDetails is a type mapping for
// https://developer.vmware.com/apis/vsphere-automation/latest/esx/data-structures/Settings/ComponentDetails/
type SettingsComponentDetails struct {
	DisplayName string `json:"display_name"`
	Vendor      string `json:"vendor"`
}

// SettingsComponentInfo is a type mapping for
// https://developer.vmware.com/apis/vsphere-automation/latest/esx/data-structures/Settings/ComponentInfo/
type SettingsComponentInfo struct {
	Version string                   `json:"version"`
	Details SettingsComponentDetails `json:"details"`
}

// SettingsSolutionComponentSpec is a type mapping for
// https://developer.vmware.com/apis/vsphere-automation/latest/esx/data-structures/Settings/SolutionComponentSpec/
type SettingsSolutionComponentSpec struct {
	Component string `json:"component"`
}

// SettingsSolutionComponentDetails is a type mapping for
// https://developer.vmware.com/apis/vsphere-automation/latest/esx/data-structures/Settings/SolutionComponentDetails/
type SettingsSolutionComponentDetails struct {
	Component      string `json:"component"`
	DisplayName    string `json:"display_name"`
	Vendor         string `json:"vendor"`
	DisplayVersion string `json:"display_version,omitempty"`
}

// SettingsSolutionDetails is a type mapping for
// https://developer.vmware.com/apis/vsphere-automation/latest/esx/data-structures/Settings/SolutionDetails/
type SettingsSolutionDetails struct {
	DisplayName    string `json:"display_name"`
	DisplayVersion string `json:"display_version"`
}

// SettingsSolutionInfo is a type mapping for
// https://developer.vmware.com/apis/vsphere-automation/latest/esx/data-structures/Settings/SolutionInfo/
type SettingsSolutionInfo struct {
	Version    string                          `json:"version"`
	Components []SettingsSolutionComponentSpec `json:"components"`
	Details    SettingsSolutionDetails         `json:"details"`
}

// SettingsAddOnDetails is a type mapping for
// https://developer.vmware.com/apis/vsphere-automation/latest/esx/data-structures/Settings/AddOnDetails/
type SettingsAddOnDetails struct {
	DisplayName    string `json:"display_name"`
	DisplayVersion string `json:"display_version"`
	Vendor         string `json:"vendor"`
}

// SettingsAddOnInfo is a type mapping for
// https://developer.vmware.com/apis/vsphere-automation/latest/esx/data-structures/Settings/AddOnInfo/
type SettingsAddOnInfo struct {
	Name    string               `json:"name"`
	Version string               `json:"version"`
	Details SettingsAddOnDetails `json:"details,omitempty"`
}

// SettingsHardwareSupportPackageInfo is a type mapping for
// https://developer.vmware.com/apis/vsphere-automation/latest/esx/data-structures/Settings/HardwareSupportPackageInfo/
type SettingsHardwareSupportPackageInfo struct {
	Pkg     string `json:"pkg"`
	Version string `json:"version"`
}

// SettingsHardwareSupportInfo is a type mapping for
// https://developer.vmware.com/apis/vsphere-automation/latest/esx/data-structures/Settings/HardwareSupportInfo/
type SettingsHardwareSupportInfo struct {
	Packages map[string]SettingsHardwareSupportPackageInfo `json:"packages"`
}

// SettingsSoftwareInfo is a type mapping for
// https://developer.vmware.com/apis/vsphere-automation/latest/esx/data-structures/Settings/SoftwareInfo/
type SettingsSoftwareInfo struct {
	BaseImage       SettingsBaseImageInfo            `json:"base_image"`
	Components      map[string]SettingsComponentInfo `json:"components"`
	Solutions       map[string]SettingsSolutionInfo  `json:"solutions"`
	AddOn           SettingsAddOnInfo                `json:"add_on,omitempty"`
	HardwareSupport SettingsHardwareSupportInfo      `json:"hardware_support,omitempty"`
}

// SettingsClustersSoftwareDraftsInfo is a type mapping for
// https://developer.vmware.com/apis/vsphere-automation/latest/esx/data-structures/Settings/Clusters/Software/Drafts/Info/
type SettingsClustersSoftwareDraftsInfo struct {
	Metadata SettingsClustersSoftwareDraftsMetadata `json:"metadata"`
	Software SettingsSoftwareInfo                   `json:"software"`
}

// SoftwareComponentsUpdateSpec is a type mapping for
// https://developer.vmware.com/apis/vsphere-automation/latest/esx/data-structures/Settings/Clusters/Software/Drafts/Software/Components/UpdateSpec/
type SoftwareComponentsUpdateSpec struct {
	ComponentsToDelete []string          `json:"components_to_delete,omitempty"`
	ComponentsToSet    map[string]string `json:"components_to_set,omitempty"`
}

// SettingsClustersSoftwareDraftsCommitSpec is a type mapping for
// https://developer.vmware.com/apis/vsphere-automation/latest/esx/data-structures/Settings/Clusters/Software/Drafts/CommitSpec/
type SettingsClustersSoftwareDraftsCommitSpec struct {
	Message string `json:"message,omitempty"`
}

// SettingsBaseImageSpec is a type mapping for
// https://developer.vmware.com/apis/vsphere-automation/latest/esx/data-structures/Settings/BaseImageSpec/
type SettingsBaseImageSpec struct {
	Version string `json:"version"`
}

// EnableSoftwareManagementSpec is a type mapping for
// https://developer.vmware.com/apis/vsphere-automation/latest/esx/data-structures/Settings/Clusters/Enablement/Software/EnableSpec/
type EnableSoftwareManagementSpec struct {
	SkipSoftwareCheck bool `json:"skip_software_check"`
}

// SoftwareManagementInfo is a type mapping for
// https://developer.vmware.com/apis/vsphere-automation/latest/esx/data-structures/Settings/Clusters/Enablement/Software/Info/
type SoftwareManagementInfo struct {
	Enabled bool `json:"enabled"`
}

// ListSoftwareDrafts retrieves the software drafts for a cluster
// https://developer.vmware.com/apis/vsphere-automation/latest/esx/api/esx/settings/clusters/cluster/software/drafts/get/
func (c *Manager) ListSoftwareDrafts(clusterId string, owners *[]string) (map[string]SettingsClustersSoftwareDraftsMetadata, error) {
	path := c.Resource(fmt.Sprintf(SoftwareDraftsPath, clusterId))

	if owners != nil && len(*owners) > 0 {
		path = path.WithParam("owners", strings.Join(*owners, ","))
	}

	req := path.Request(http.MethodGet)
	var res map[string]SettingsClustersSoftwareDraftsMetadata
	return res, c.Do(context.Background(), req, &res)
}

// CreateSoftwareDraft creates a software draft on the provided cluster
// https://developer.vmware.com/apis/vsphere-automation/latest/esx/api/esx/settings/clusters/cluster/software/drafts/post/
func (c *Manager) CreateSoftwareDraft(clusterId string) (string, error) {
	path := c.Resource(fmt.Sprintf(SoftwareDraftsPath, clusterId))
	req := path.Request(http.MethodPost)
	var res string
	return res, c.Do(context.Background(), req, &res)
}

// DeleteSoftwareDraft removes the specified draft
// https://developer.vmware.com/apis/vsphere-automation/latest/esx/api/esx/settings/clusters/cluster/software/drafts/draft/delete/
func (c *Manager) DeleteSoftwareDraft(clusterId, draftId string) error {
	path := c.Resource(fmt.Sprintf(SoftwareDraftsPath, clusterId)).WithSubpath(draftId)
	req := path.Request(http.MethodDelete)
	return c.Do(context.Background(), req, nil)
}

// GetSoftwareDraft returns the set of components in the specified draft
// https://developer.vmware.com/apis/vsphere-automation/latest/esx/api/esx/settings/clusters/cluster/software/drafts/draft/get/
func (c *Manager) GetSoftwareDraft(clusterId, draftId string) (SettingsClustersSoftwareDraftsInfo, error) {
	path := c.Resource(fmt.Sprintf(SoftwareDraftsPath, clusterId)).WithSubpath(draftId)
	req := path.Request(http.MethodGet)
	var res SettingsClustersSoftwareDraftsInfo
	return res, c.Do(context.Background(), req, &res)
}

// CommitSoftwareDraft closes and applies the specified draft
// https://developer.vmware.com/apis/vsphere-automation/latest/esx/api/esx/settings/clusters/cluster/software/drafts/draftactioncommitvmw-tasktrue/post/
func (c *Manager) CommitSoftwareDraft(clusterId, draftId string, spec SettingsClustersSoftwareDraftsCommitSpec) (string, error) {
	path := c.Resource(fmt.Sprintf(SoftwareDraftsPath, clusterId)).WithSubpath(draftId).WithParam("action", "commit").WithParam("vmw-task", "true")
	req := path.Request(http.MethodPost, spec)
	var res string
	return res, c.Do(context.Background(), req, &res)
}

// ListSoftwareDraftComponents returns all components in the specified draft
// https://developer.vmware.com/apis/vsphere-automation/latest/esx/api/esx/settings/clusters/cluster/software/drafts/draft/software/components/get/
func (c *Manager) ListSoftwareDraftComponents(clusterId, draftId string) (map[string]SettingsComponentInfo, error) {
	path := c.Resource(fmt.Sprintf(SoftwareComponentsPath, clusterId, draftId))
	req := path.Request(http.MethodGet)
	var res map[string]SettingsComponentInfo
	return res, c.Do(context.Background(), req, &res)
}

// GetSoftwareDraftComponent returns a component from the specified draft
// https://developer.vmware.com/apis/vsphere-automation/latest/esx/api/esx/settings/clusters/cluster/software/drafts/draft/software/components/component/get/
func (c *Manager) GetSoftwareDraftComponent(clusterId, draftId, component string) (SettingsComponentInfo, error) {
	path := c.Resource(fmt.Sprintf(SoftwareComponentsPath, clusterId, draftId)).WithSubpath(component)
	req := path.Request(http.MethodGet)
	var res SettingsComponentInfo
	return res, c.Do(context.Background(), req, &res)
}

// UpdateSoftwareDraftComponents updates the set of components in the specified draft
// https://developer.vmware.com/apis/vsphere-automation/latest/esx/api/esx/settings/clusters/cluster/software/drafts/draft/software/components/patch/
func (c *Manager) UpdateSoftwareDraftComponents(clusterId, draftId string, spec SoftwareComponentsUpdateSpec) error {
	path := c.Resource(fmt.Sprintf(SoftwareComponentsPath, clusterId, draftId))
	req := path.Request(http.MethodPatch, spec)
	return c.Do(context.Background(), req, nil)
}

// RemoveSoftwareDraftComponents removes a component from the specified draft
// https://developer.vmware.com/apis/vsphere-automation/latest/esx/api/esx/settings/clusters/cluster/software/drafts/draft/software/components/component/delete/
func (c *Manager) RemoveSoftwareDraftComponents(clusterId, draftId, component string) error {
	path := c.Resource(fmt.Sprintf(SoftwareComponentsPath, clusterId, draftId)).WithSubpath(component)
	req := path.Request(http.MethodDelete)
	return c.Do(context.Background(), req, nil)
}

// GetSoftwareDraftBaseImage retrieves the ESXi image version on the specified draft
// https://developer.vmware.com/apis/vsphere-automation/latest/esx/api/esx/settings/clusters/cluster/software/drafts/draft/software/base-image/get
func (c *Manager) GetSoftwareDraftBaseImage(clusterId, draftId string) (SettingsBaseImageInfo, error) {
	path := c.Resource(fmt.Sprintf(BaseImagePath, clusterId, draftId))
	req := path.Request(http.MethodGet)
	var res SettingsBaseImageInfo
	return res, c.Do(context.Background(), req, &res)
}

// SetSoftwareDraftBaseImage sets the ESXi image version on the specified draft
// https://developer.vmware.com/apis/vsphere-automation/latest/esx/settings/clusters.software.drafts.software.base_image/put
func (c *Manager) SetSoftwareDraftBaseImage(clusterId, draftId, version string) error {
	path := c.Resource(fmt.Sprintf(BaseImagePath, clusterId, draftId))
	req := path.Request(http.MethodPut, SettingsBaseImageSpec{Version: version})
	return c.Do(context.Background(), req, nil)
}

// EnableSoftwareManagement enables vLCM on the cluster
// https://developer.vmware.com/apis/vsphere-automation/latest/esx/settings/clusters.enablement.software/put
func (c *Manager) EnableSoftwareManagement(clusterId string, skipCheck bool) (string, error) {
	path := c.Resource(fmt.Sprintf(SoftwareEnablementPath, clusterId)).WithParam("vmw-task", "true")
	req := path.Request(http.MethodPut, EnableSoftwareManagementSpec{SkipSoftwareCheck: skipCheck})
	var res string
	return res, c.Do(context.Background(), req, &res)
}

// GetSoftwareManagement checks whether vLCM is enabled on the cluster
// https://developer.vmware.com/apis/vsphere-automation/latest/esx/api/esx/settings/clusters/cluster/enablement/software/get/
func (c *Manager) GetSoftwareManagement(clusterId string) (SoftwareManagementInfo, error) {
	path := c.Resource(fmt.Sprintf(SoftwareEnablementPath, clusterId))
	req := path.Request(http.MethodGet)
	var res SoftwareManagementInfo
	return res, c.Do(context.Background(), req, &res)
}
