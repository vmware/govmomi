// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package drafts

import (
	"context"
	"fmt"
	"net/http"

	"github.com/vmware/govmomi/vapi/esx/settings/clusters/configuration"
	"github.com/vmware/govmomi/vapi/rest"
)

const (
	// BasePath the base endpoint for the clusters configuration API
	BasePath = configuration.BasePath + "/drafts"
	// DraftPath the base endpoint for operations on individual configuration drafts
	DraftPath = BasePath + "/%s"
)

// Manager extends rest.Client, adding cluster configuration drafts related methods.
type Manager struct {
	*rest.Client
}

// NewManager creates a new Manager instance with the given client.
func NewManager(client *rest.Client) *Manager {
	return &Manager{
		Client: client,
	}
}

// Draft a cluster configuration draft
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/data-structures/Esx%20Settings%20Clusters%20Configuration%20Drafts%20Info/
type Draft struct {
	ID    string `json:"id"`
	State string `json:"state"`
}

// CreateSpec a specification for creating configuration drafts
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/data-structures/Esx%20Settings%20Clusters%20Configuration%20Drafts%20CreateSpec/
type CreateSpec struct {
	Config        string `json:"config"`
	ReferenceHost string `json:"reference_host"`
}

// ApplySpec a specification for committing a pending configuration draft
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/data-structures/Esx%20Settings%20Clusters%20Configuration%20Drafts%20ApplySpec/
type ApplySpec struct {
	PolicySpec ApplyPolicySpec `json:"apply_policy_spec"`
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

// UpdateSpec a specification for updating a pending configuration draft
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/data-structures/Esx%20Settings%20Clusters%20Configuration%20Drafts%20UpdateSpec/
type UpdateSpec struct {
	Config string `json:"config"`
}

// ImportSpec a specification for importing a configuration from a reference host into a draft
type ImportSpec struct {
	Host string `json:"host"`
}

// ApplyResult the result of a draft application
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/data-structures/Esx%20Settings%20Clusters%20Configuration%20Drafts%20ApplyResult/
type ApplyResult struct {
	Commit    string `json:"commit"`
	ApplyTask string `json:"apply_task"`
}

// ListDrafts returns all active drafts
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/api/esx/settings/clusters/cluster/configuration/drafts
func (c *Manager) ListDrafts(clusterId string) (map[string]Draft, error) {
	path := c.Resource(fmt.Sprintf(BasePath, clusterId))
	req := path.Request(http.MethodGet)
	var res map[string]Draft
	return res, c.Do(context.Background(), req, &res)
}

// GetDraft returns a draft by its ID
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/api/esx/settings/clusters/cluster/configuration/drafts/draft/get/
func (c *Manager) GetDraft(clusterId, draftId string) (Draft, error) {
	path := c.Resource(fmt.Sprintf(DraftPath, clusterId, draftId))
	req := path.Request(http.MethodGet)
	var res Draft
	return res, c.Do(context.Background(), req, &res)
}

// DeleteDraft deletes a draft
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/api/esx/settings/clusters/cluster/configuration/drafts/draft/delete/
func (c *Manager) DeleteDraft(clusterId, draftId string) error {
	path := c.Resource(fmt.Sprintf(DraftPath, clusterId, draftId))
	req := path.Request(http.MethodDelete)
	return c.Do(context.Background(), req, nil)
}

// CreateDraft creates a draft with the provided configuration
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/api/esx/settings/clusters/cluster/configuration/drafts/post/
func (c *Manager) CreateDraft(clusterId string, spec CreateSpec) (string, error) {
	path := c.Resource(fmt.Sprintf(BasePath, clusterId))
	req := path.Request(http.MethodPost, spec)
	var res string
	return res, c.Do(context.Background(), req, &res)
}

// ApplyDraft commits the draft with the specified ID
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/api/esx/settings/clusters/cluster/configuration/drafts/draft__action=apply/post/
func (c *Manager) ApplyDraft(clusterId, draftId string) (ApplyResult, error) {
	path := c.Resource(fmt.Sprintf(DraftPath, clusterId, draftId)).WithParam("action", "apply")
	req := path.Request(http.MethodPost)
	var res ApplyResult
	return res, c.Do(context.Background(), req, &res)
}

// UpdateDraft updates the configuration of the draft with the specified ID
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/api/esx/settings/clusters/cluster/configuration/drafts/draft__action=update/post/
func (c *Manager) UpdateDraft(clusterId, draftId string, spec UpdateSpec) error {
	path := c.Resource(fmt.Sprintf(DraftPath, clusterId, draftId)).WithParam("action", "update")
	req := path.Request(http.MethodPost, spec)
	return c.Do(context.Background(), req, nil)
}

// ImportFromHost sets a reference host to use as the source for the draft configuration
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/api/esx/settings/clusters/cluster/configuration/drafts/draft__action=importFromHost__vmw-task=true/post/
func (c *Manager) ImportFromHost(clusterId, draftId string, spec ImportSpec) (string, error) {
	path := c.Resource(fmt.Sprintf(DraftPath, clusterId, draftId)).WithParam("action", "importFromHost").WithParam("vmw-task", "true")
	req := path.Request(http.MethodPost, spec)
	var res string
	return res, c.Do(context.Background(), req, &res)
}

// Precheck runs pre-checks for the provided draft on the specified cluster
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/api/esx/settings/clusters/cluster/configuration/drafts/draft__action=precheck__vmw-task=true/post/
func (c *Manager) Precheck(clusterId, draftId string) (string, error) {
	path := c.Resource(fmt.Sprintf(DraftPath, clusterId, draftId)).WithParam("action", "precheck").WithParam("vmw-task", "true")
	req := path.Request(http.MethodPost)
	var res string
	return res, c.Do(context.Background(), req, &res)
}
