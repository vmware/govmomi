// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package enablement

import (
	"context"
	"fmt"
	"net/http"

	"github.com/vmware/govmomi/vapi/esx/settings/clusters"
	"github.com/vmware/govmomi/vapi/rest"
)

const (
	// BasePath the base endpoint for the clusters enablement configuration API
	BasePath = clusters.BasePath + "/%s/enablement"
	// ConfigurationPath the endpoint for the enablement configuration API
	ConfigurationPath = BasePath + "/configuration"
	// TransitionPath the endpoint for the transition API
	TransitionPath = ConfigurationPath + "/transition"
)

// FileSpec a specification used for importing cluster configuration form a file
type FileSpec struct {
	Config   string `json:"config"`
	Filename string `json:"filename"`
}

// DraftImportResult contains draft information, a result of importing a configuration from a file or a reference host
type DraftImportResult struct {
	Status string `json:"status"`
	Draft  string `json:"draft"`
}

// Manager extends rest.Client, adding cluster configuration enablement related methods.
type Manager struct {
	*rest.Client
}

// NewManager creates a new Manager instance with the given client.
func NewManager(client *rest.Client) *Manager {
	return &Manager{
		Client: client,
	}
}

// EnableClusterConfiguration enables cluster configuration profiles
// Returns a task identifier and an error
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/api/esx/settings/clusters/cluster/enablement/configuration/transition__action=enable__vmw-task=true/post/
func (c *Manager) EnableClusterConfiguration(clusterId string) (string, error) {
	path := c.getUrlWithActionAndTask(clusterId, "enable")
	req := path.Request(http.MethodPost, nil)
	var res string
	return res, c.Do(context.Background(), req, &res)
}

// ImportFromReferenceHost imports the configuration of an existing ESXi host
// Returns a task identifier and an error
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/api/esx/settings/clusters/cluster/enablement/configuration/transition__action=importFromHost__vmw-task=true/post/
func (c *Manager) ImportFromReferenceHost(clusterId, hostId string) (string, error) {
	path := c.getUrlWithActionAndTask(clusterId, "importFromHost")
	req := path.Request(http.MethodPost, hostId)
	var res string
	return res, c.Do(context.Background(), req, &res)
}

// ImportFromFile imports the configuration in the provided json string
// Returns a task identifier and an error
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/api/esx/settings/clusters/cluster/enablement/configuration/transition__action=importFromFile/post/
func (c *Manager) ImportFromFile(clusterId string, spec FileSpec) (DraftImportResult, error) {
	path := c.getUrlWithActionAndTask(clusterId, "importFromFile")
	req := path.Request(http.MethodPost, spec)
	var res DraftImportResult
	return res, c.Do(context.Background(), req, &res)
}

// ValidateConfiguration performs server-side validation of the pending cluster configuration
// Returns a task identifier and an error
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/api/esx/settings/clusters/cluster/enablement/configuration/transition__action=validateConfig__vmw-task=true/post/
func (c *Manager) ValidateConfiguration(clusterId string) (string, error) {
	path := c.getUrlWithActionAndTask(clusterId, "validateConfig")
	req := path.Request(http.MethodPost, nil)
	var res string
	return res, c.Do(context.Background(), req, &res)
}

// CheckEligibility performs server-side validation of whether the cluster is eligible for configuration management via profiles
// Returns a task identifier and an error
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/api/esx/settings/clusters/cluster/enablement/configuration/transition__action=checkEligibility__vmw-task=true/post/
func (c *Manager) CheckEligibility(clusterId string) (string, error) {
	path := c.getUrlWithActionAndTask(clusterId, "checkEligibility")
	req := path.Request(http.MethodPost, nil)
	var res string
	return res, c.Do(context.Background(), req, &res)
}

// RunPrecheck performs server-side pre-checks of the pending cluster configuration
// Returns a task identifier and an error
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/api/esx/settings/clusters/cluster/enablement/configuration/transition__action=precheck__vmw-task=true/post/
func (c *Manager) RunPrecheck(clusterId string) (string, error) {
	path := c.getUrlWithActionAndTask(clusterId, "precheck")
	req := path.Request(http.MethodPost, nil)
	var res string
	return res, c.Do(context.Background(), req, &res)
}

// Cancel deletes the pending cluster configuration
// Returns a task identifier and an error
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/api/esx/settings/clusters/cluster/enablement/configuration/transition__action=cancel/post/
func (c *Manager) Cancel(clusterId string) (string, error) {
	path := c.Resource(fmt.Sprintf(TransitionPath, clusterId)).WithParam("action", "cancel")
	req := path.Request(http.MethodPost, nil)
	var res string
	return res, c.Do(context.Background(), req, &res)
}

// GetClusterConfigurationStatus returns the status of the current pending cluster configuration
// Returns the config status and an error
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/api/esx/settings/clusters/cluster/enablement/configuration/transition/get/
func (c *Manager) GetClusterConfigurationStatus(clusterId string) (string, error) {
	path := c.Resource(fmt.Sprintf(TransitionPath, clusterId))
	req := path.Request(http.MethodGet)
	var res string
	return res, c.Do(context.Background(), req, &res)
}

func (c *Manager) getUrlWithActionAndTask(clusterId, action string) *rest.Resource {
	return c.Resource(fmt.Sprintf(TransitionPath, clusterId)).WithParam("action", action).WithParam("vmw-task", "true")
}
