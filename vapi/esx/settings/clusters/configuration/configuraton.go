// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package configuration

import (
	"context"
	"fmt"
	"net/http"

	"github.com/vmware/govmomi/vapi/esx/settings/clusters"
	"github.com/vmware/govmomi/vapi/rest"
)

const (
	// BasePath the base endpoint for the clusters configuration API
	BasePath = clusters.BasePath + "/%s/configuration"
	// RecentTasksPath the endpoint for retrieving the most recent tasks for each draft-related operation
	RecentTasksPath = BasePath + "/reports/recent-tasks"
	// SchemaPath the endpoint for retrieving the configuration schema of a cluster
	SchemaPath = BasePath + "/schema"
)

// Manager extends rest.Client, adding cluster configuration related methods.
type Manager struct {
	*rest.Client
}

// NewManager creates a new Manager instance with the given client.
func NewManager(client *rest.Client) *Manager {
	return &Manager{
		Client: client,
	}
}

// Info configuration information
type Info struct {
	Config   string   `json:"config"`
	Metadata Metadata `json:"metadata"`
}

// Metadata configuration metadata
type Metadata struct {
	Id string `json:"id"`
}

// ExportResult the result of a configuration export operation
type ExportResult struct {
	Config string `json:"config"`
}

// ImportSpec a specification for importing a cluster configuration
type ImportSpec struct {
	Config      string  `json:"config"`
	Description *string `json:"description"`
}

// ApplySpec a specification for applying a cluster configuration
type ApplySpec struct {
	Hosts           *[]string        `json:"hosts"`
	ApplyPolicySpec *ApplyPolicySpec `json:"apply_policy_spec"`
}

// ApplyPolicySpec contains options for customizing how a draft is applied on a cluster
type ApplyPolicySpec struct {
	FailureAction             *FailureAction             `json:"failure_action,omitempty"`
	PreRemediationPowerAction *string                    `json:"pre_remediation_power_action,omitempty"`
	EnableQuickBoot           *bool                      `json:"enable_quick_boot,omitempty"`
	DisableDpm                *bool                      `json:"disable_dpm,omitempty"`
	DisableHac                *bool                      `json:"disable_hac,omitempty"`
	EvacuateOfflineVms        *bool                      `json:"evacuate_offline_vms,omitempty"`
	EnforceHclValidation      *bool                      `json:"enforce_hcl_validation,omitempty"`
	EnforceQuickPatch         *bool                      `json:"enforce_quick_patch,omitempty"`
	ParallelRemediationAction *ParallelRemediationAction `json:"parallel_remediation_action,omitempty"`
	ConfigManagerPolicySpec   *ConfigManagerPolicySpec   `json:"config_manager_policy_spec,omitempty"`
}

// FailureAction contains options for scheduling actions that will be trigger if a draft application fails
type FailureAction struct {
	Action     string `json:"action"`
	RetryDelay int64  `json:"retry_delay"`
	RetryCount int64  `json:"retry_count"`
}

// ParallelRemediationAction contains options for configuring parallel host remediation
type ParallelRemediationAction struct {
	Enabled  bool  `json:"enabled"`
	MaxHosts int64 `json:"max_hosts"`
}

// ConfigManagerPolicySpec a structure for enabling or disabling serial remediation
type ConfigManagerPolicySpec struct {
	SerialRemediation bool `json:"serial_remediation"`
}

// RecentTasksInfo a structure that contains the task identifiers of the most recent draft-related tasks
type RecentTasksInfo struct {
	CheckCompliance string `json:"check_compliance"`
	Precheck        string `json:"precheck"`
	Apply           string `json:"apply"`
	DraftTasks      []any  `json:"draft_tasks"`
}

// SchemaInfo contains data for a configuration schema
type SchemaInfo struct {
	Source string `json:"source"`
	Schema string `json:"schema"`
}

// GetConfiguration returns the cluster configuration
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/api/esx/settings/clusters/cluster/configuration/get/
func (c *Manager) GetConfiguration(clusterId string) (Info, error) {
	path := c.Resource(fmt.Sprintf(BasePath, clusterId))
	req := path.Request(http.MethodGet)
	var res Info
	return res, c.Do(context.Background(), req, &res)
}

// ExportConfiguration returns the cluster configuration
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/api/esx/settings/clusters/cluster/configuration__action=exportConfig/post/
func (c *Manager) ExportConfiguration(clusterId string) (ExportResult, error) {
	path := c.Resource(fmt.Sprintf(BasePath, clusterId)).WithParam("action", "exportConfig")
	req := path.Request(http.MethodPost)
	var res ExportResult
	return res, c.Do(context.Background(), req, &res)
}

// ApplyConfiguration applies the current configuration to the provided hosts in the cluster
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/api/esx/settings/clusters/cluster/configuration__action=apply__vmw-task=true/post/
func (c *Manager) ApplyConfiguration(clusterId string, spec ApplySpec) (string, error) {
	path := c.getUrlWithActionAndTask(clusterId, "apply")
	req := path.Request(http.MethodPost, spec)
	var res string
	return res, c.Do(context.Background(), req, &res)
}

// CheckCompliance initiates a compliance check on the cluster
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/api/esx/settings/clusters/cluster/configuration__action=checkCompliance__vmw-task=true/post/
func (c *Manager) CheckCompliance(clusterId string) (string, error) {
	path := c.getUrlWithActionAndTask(clusterId, "checkCompliance")
	req := path.Request(http.MethodPost)
	var res string
	return res, c.Do(context.Background(), req, &res)
}

// Validate initiates a validation check on the pending cluster configuration
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/api/esx/settings/clusters/cluster/configuration__action=validate__vmw-task=true/post/
func (c *Manager) Validate(clusterId string) (string, error) {
	path := c.getUrlWithActionAndTask(clusterId, "validate")
	req := path.Request(http.MethodPost)
	var res string
	return res, c.Do(context.Background(), req, &res)
}

// Precheck initiates a precheck on the pending cluster configuration
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/api/esx/settings/clusters/cluster/configuration__action=precheck__vmw-task=true/post/
func (c *Manager) Precheck(clusterId string) (string, error) {
	path := c.getUrlWithActionAndTask(clusterId, "precheck")
	req := path.Request(http.MethodPost)
	var res string
	return res, c.Do(context.Background(), req, &res)
}

// Import imports the provided configuration
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/api/esx/settings/clusters/cluster/configuration__action=importConfig__vmw-task=true/post/
func (c *Manager) Import(clusterId string, spec ImportSpec) (string, error) {
	path := c.getUrlWithActionAndTask(clusterId, "importConfig")
	req := path.Request(http.MethodPost, spec)
	var res string
	return res, c.Do(context.Background(), req, &res)
}

// GetRecentTasks returns the task identifiers for the latest configuration operations of each type
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/api/esx/settings/clusters/cluster/configuration/reports/recent-tasks/get/
func (c *Manager) GetRecentTasks(clusterId string) (RecentTasksInfo, error) {
	path := c.Resource(fmt.Sprintf(RecentTasksPath, clusterId))
	req := path.Request(http.MethodGet)
	var res RecentTasksInfo
	return res, c.Do(context.Background(), req, &res)
}

// GetSchema returns the configuration schema for the cluster
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/api/esx/settings/clusters/cluster/configuration/schema/get/
func (c *Manager) GetSchema(clusterId string) (SchemaInfo, error) {
	path := c.Resource(fmt.Sprintf(SchemaPath, clusterId))
	req := path.Request(http.MethodGet)
	var res SchemaInfo
	return res, c.Do(context.Background(), req, &res)
}

func (c *Manager) getUrlWithActionAndTask(clusterId, action string) *rest.Resource {
	return c.Resource(fmt.Sprintf(BasePath, clusterId)).WithParam("action", action).WithParam("vmw-task", "true")
}
