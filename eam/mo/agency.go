// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package mo

import (
	"github.com/vmware/govmomi/eam/types"
	vim "github.com/vmware/govmomi/vim25/types"
)

// Agency handles the deployment of a single type of agent virtual
// machine and any associated VIB bundle, on a set of compute resources.
type Agency struct {
	EamObject `yaml:",inline"`

	Agent      []vim.ManagedObjectReference `json:"agent,omitempty"`
	Config     types.BaseAgencyConfigInfo   `json:"config"`
	Runtime    types.EamObjectRuntimeInfo   `json:"runtime"`
	SolutionId string                       `json:"solutionId,omitempty"`
}
