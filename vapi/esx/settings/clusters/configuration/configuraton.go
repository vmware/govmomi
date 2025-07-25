// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package configuration

import (
	"context"
	"fmt"
	"net/http"

	"github.com/vmware/govmomi/vapi/esx/settings"
	"github.com/vmware/govmomi/vapi/rest"
)

const (
	// BasePath The base endpoint for the clusters configuration API
	BasePath = settings.BasePath + "/%s/configuration"
)

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

type Info struct {
	Config   string   `json:"config"`
	Metadata Metadata `json:"metadata"`
}

type Metadata struct {
	Id string `json:"id"`
}

type ExportResult struct {
	Config string `json:"config"`
}

type ImportSpec struct {
	Config      string  `json:"config"`
	Description *string `json:"description"`
}

type ApplySpec struct {
	Hosts           *[]string        `json:"hosts"`
	ApplyPolicySpec *ApplyPolicySpec `json:"apply_policy_spec"`
}

type ApplyPolicySpec struct {
	FailureAction             *FailureAction             `json:"failure_action"`
	PreRemediationPowerAction *string                    `json:"pre_remediation_power_action"`
	EnableQuickBoot           *bool                      `json:"enable_quick_boot"`
	DisableDpm                *bool                      `json:"disable_dpm"`
	DisableHac                *bool                      `json:"disable_hac"`
	EvacuateOfflineVms        *bool                      `json:"evacuate_offline_vms"`
	EnforceHclValidation      *bool                      `json:"enforce_hcl_validation"`
	EnforceQuickPatch         *bool                      `json:"enforce_quick_patch"`
	ParallelRemediationAction *ParallelRemediationAction `json:"parallel_remediation_action"`
	ConfigManagerPolicySpec   *ConfigManagerPolicySpec   `json:"config_manager_policy_spec"`
}

type FailureAction struct {
	Action     string `json:"action"`
	RetryDelay int    `json:"retry_delay"`
	RetryCount int    `json:"retry_count"`
}

type ParallelRemediationAction struct {
	Enabled  bool `json:"enabled"`
	MaxHosts int  `json:"max_hosts"`
}

type ConfigManagerPolicySpec struct {
	SerialRemediation bool `json:"serial_remediation"`
}

// GetConfiguration returns the cluster configuration
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/api/esx/settings/clusters/cluster/configuration
func (c *Manager) GetConfiguration(clusterId string) (Info, error) {
	path := c.Resource(fmt.Sprintf(BasePath, clusterId))
	req := path.Request(http.MethodGet)
	var res Info
	return res, c.Do(context.Background(), req, &res)
}

// ExportConfiguration returns the cluster configuration
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/api/esx/settings/clusters/cluster/configuration
func (c *Manager) ExportConfiguration(clusterId string) (ExportResult, error) {
	path := c.Resource(fmt.Sprintf(BasePath, clusterId)).WithAction("exportConfig")
	req := path.Request(http.MethodPost)
	var res ExportResult
	return res, c.Do(context.Background(), req, &res)
}

// ApplyConfiguration applies the current configuration to the provided hosts in the cluster
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/api/esx/settings/clusters/cluster/configuration
func (c *Manager) ApplyConfiguration(clusterId string, spec ApplySpec) (string, error) {
	path := c.getBaseTransitionUrl(clusterId, "apply")
	req := path.Request(http.MethodPost, spec)
	var res string
	return res, c.Do(context.Background(), req, &res)
}

// CheckCompliance initiates a compliance check on the cluster
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/api/esx/settings/clusters/cluster/configuration
func (c *Manager) CheckCompliance(clusterId string) (string, error) {
	path := c.getBaseTransitionUrl(clusterId, "checkCompliance")
	req := path.Request(http.MethodPost)
	var res string
	return res, c.Do(context.Background(), req, &res)
}

// Validate initiates a validation check on the pending cluster configuration
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/api/esx/settings/clusters/cluster/configuration
func (c *Manager) Validate(clusterId string) (string, error) {
	path := c.getBaseTransitionUrl(clusterId, "validate")
	req := path.Request(http.MethodPost)
	var res string
	return res, c.Do(context.Background(), req, &res)
}

// Precheck initiates a precheck on the pending cluster configuration
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/api/esx/settings/clusters/cluster/configuration
func (c *Manager) Precheck(clusterId string) (string, error) {
	path := c.getBaseTransitionUrl(clusterId, "precheck")
	req := path.Request(http.MethodPost)
	var res string
	return res, c.Do(context.Background(), req, &res)
}

// Import imports the provided configuration
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/api/esx/settings/clusters/cluster/configuration
func (c *Manager) Import(clusterId string, spec ImportSpec) (string, error) {
	path := c.getBaseTransitionUrl(clusterId, "importConfig")
	req := path.Request(http.MethodPost, spec)
	var res string
	return res, c.Do(context.Background(), req, &res)
}

func (c *Manager) getBaseTransitionUrl(clusterId, action string) *rest.Resource {
	return c.Resource(fmt.Sprintf(BasePath, clusterId)).WithParam("action", action).WithParam("vmw-task", "true")
}
