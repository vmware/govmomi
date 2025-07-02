// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package mo

import (
	vim "github.com/vmware/govmomi/vim25/types"
)

// EsxAgentManager is the main entry point for a solution to create
// agencies in the vSphere ESX Agent Manager server.
type EsxAgentManager struct {
	EamObject `yaml:",inline"`

	Agency []vim.ManagedObjectReference `json:"agency,omitempty"`
}
