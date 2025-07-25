// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package enablement

import (
	"context"
	"fmt"
	"net/http"

	"github.com/vmware/govmomi/vapi/esx/settings"
	"github.com/vmware/govmomi/vapi/rest"
)

const (
	// BasePath The base endpoint for the clusters enablement configuration API
	BasePath      = settings.BasePath + "/%s/enablement"
	Configuration = BasePath + "/configuration"
	Transition    = Configuration + "/transition"
)

type FileSpec struct {
	Config   string `json:"config"`
	Filename string `json:"filename"`
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
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/api/esx/settings/clusters/cluster/enablement/configuration/transition
func (c *Manager) EnableClusterConfiguration(clusterId string) (string, error) {
	path := c.getBaseTransitionUrl(clusterId, "enable")
	req := path.Request(http.MethodPost, nil)
	var res string
	return res, c.Do(context.Background(), req, &res)
}

// ImportFromReferenceHost imports the configuration of an existing ESXi host
// Returns a task identifier and an error
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/api/esx/settings/clusters/cluster/enablement/configuration/transition
func (c *Manager) ImportFromReferenceHost(clusterId, hostId string) (string, error) {
	path := c.getBaseTransitionUrl(clusterId, "importFromHost")
	req := path.Request(http.MethodPost, hostId)
	var res string
	return res, c.Do(context.Background(), req, &res)
}

// ImportFromFile imports the configuration in the provided json string
// Returns a task identifier and an error
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/api/esx/settings/clusters/cluster/enablement/configuration/transition
func (c *Manager) ImportFromFile(clusterId string, spec FileSpec) (string, error) {
	path := c.getBaseTransitionUrl(clusterId, "importFromFile")
	req := path.Request(http.MethodPost, spec)
	var res string
	return res, c.Do(context.Background(), req, &res)
}

// ValidateConfiguration performs server-side validation of the pending cluster configuration
// Returns a task identifier and an error
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/api/esx/settings/clusters/cluster/enablement/configuration/transition
func (c *Manager) ValidateConfiguration(clusterId string) (string, error) {
	path := c.getBaseTransitionUrl(clusterId, "validateConfig")
	req := path.Request(http.MethodPost, nil)
	var res string
	return res, c.Do(context.Background(), req, &res)
}

// CheckEligibility performs server-side validation of whether the cluster is eligible for configuration management via profiles
// Returns a task identifier and an error
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/api/esx/settings/clusters/cluster/enablement/configuration/transition
func (c *Manager) CheckEligibility(clusterId string) (string, error) {
	path := c.getBaseTransitionUrl(clusterId, "checkEligibility")
	req := path.Request(http.MethodPost, nil)
	var res string
	return res, c.Do(context.Background(), req, &res)
}

// RunPrecheck performs server-side pre-checks of the pending cluster configuration
// Returns a task identifier and an error
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/api/esx/settings/clusters/cluster/enablement/configuration/transition
func (c *Manager) RunPrecheck(clusterId string) (string, error) {
	path := c.getBaseTransitionUrl(clusterId, "precheck")
	req := path.Request(http.MethodPost, nil)
	var res string
	return res, c.Do(context.Background(), req, &res)
}

// Cancel deletes the pending cluster configuration
// Returns a task identifier and an error
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/api/esx/settings/clusters/cluster/enablement/configuration/transition
func (c *Manager) Cancel(clusterId string) (string, error) {
	path := c.Resource(fmt.Sprintf(Transition, clusterId)).WithParam("action", "cancel")
	req := path.Request(http.MethodPost, nil)
	var res string
	return res, c.Do(context.Background(), req, &res)
}

// GetClusterConfigurationStatus returns the status of the current pending cluster configuration
// Returns the config status and an error
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/api/esx/settings/clusters/cluster/enablement/configuration/transition
func (c *Manager) GetClusterConfigurationStatus(clusterId string) (string, error) {
	path := c.Resource(fmt.Sprintf(Transition, clusterId))
	req := path.Request(http.MethodGet)
	var res string
	return res, c.Do(context.Background(), req, &res)
}

func (c *Manager) getBaseTransitionUrl(clusterId, action string) *rest.Resource {
	return c.Resource(fmt.Sprintf(Transition, clusterId)).WithParam("action", action).WithParam("vmw-task", "true")
}
