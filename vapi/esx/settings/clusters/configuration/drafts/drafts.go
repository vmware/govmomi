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
	// BasePath The base endpoint for the clusters configuration API
	BasePath  = configuration.BasePath + "/drafts"
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

type Draft struct {
	ID    string `json:"id"`
	State string `json:"state"`
}

type CreateSpec struct {
	Config        string `json:"config"`
	ReferenceHost string `json:"reference_host"`
}

type ApplySpec struct {
	PolicySpec ApplyPolicySpec `json:"apply_policy_spec"`
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

type UpdateSpec struct {
	Config string `json:"config"`
}

type ImportSpec struct {
	Host string `json:"host"`
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
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/api/esx/settings/clusters/cluster/configuration/drafts
func (c *Manager) GetDraft(clusterId, draftId string) (Draft, error) {
	path := c.Resource(fmt.Sprintf(DraftPath, clusterId, draftId))
	req := path.Request(http.MethodGet)
	var res Draft
	return res, c.Do(context.Background(), req, &res)
}

// DeleteDraft deletes a draft
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/api/esx/settings/clusters/cluster/configuration/drafts
func (c *Manager) DeleteDraft(clusterId, draftId string) (Draft, error) {
	path := c.Resource(fmt.Sprintf(DraftPath, clusterId, draftId))
	req := path.Request(http.MethodDelete)
	var res Draft
	return res, c.Do(context.Background(), req, &res)
}

// CreateDraft creates a draft with the provided configuration
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/api/esx/settings/clusters/cluster/configuration/drafts
func (c *Manager) CreateDraft(clusterId string, spec CreateSpec) error {
	path := c.Resource(fmt.Sprintf(BasePath, clusterId))
	req := path.Request(http.MethodPost, spec)
	return c.Do(context.Background(), req, nil)
}

// ApplyDraft commits the draft with the specified ID
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/api/esx/settings/clusters/cluster/configuration/drafts
func (c *Manager) ApplyDraft(clusterId, draftId string) error {
	path := c.Resource(fmt.Sprintf(DraftPath, clusterId, draftId)).WithAction("apply")
	req := path.Request(http.MethodPost)
	return c.Do(context.Background(), req, nil)
}

// UpdateDraft updates the configuration of the draft with the specified ID
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/api/esx/settings/clusters/cluster/configuration/drafts
func (c *Manager) UpdateDraft(clusterId, draftId string, spec UpdateSpec) error {
	path := c.Resource(fmt.Sprintf(DraftPath, clusterId, draftId)).WithAction("update")
	req := path.Request(http.MethodPost, spec)
	return c.Do(context.Background(), req, nil)
}

// ImportFromHost sets a reference host to use as the source for the draft configuration
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/api/esx/settings/clusters/cluster/configuration/drafts
func (c *Manager) ImportFromHost(clusterId, draftId string, spec ImportSpec) (string, error) {
	path := c.Resource(fmt.Sprintf(DraftPath, clusterId, draftId)).WithAction("importFromHost").WithParam("vmw-task", "true")
	req := path.Request(http.MethodPost, spec)
	var res string
	return res, c.Do(context.Background(), req, &res)
}

// Precheck runs pre-checks for the provided draft on the specified cluster
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/api/esx/settings/clusters/cluster/configuration/drafts
func (c *Manager) Precheck(clusterId, draftId string) (string, error) {
	path := c.Resource(fmt.Sprintf(DraftPath, clusterId, draftId)).WithAction("precheck").WithParam("vmw-task", "true")
	req := path.Request(http.MethodPost)
	var res string
	return res, c.Do(context.Background(), req, &res)
}
