// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package mo

import (
	"github.com/vmware/govmomi/eam/types"
)

// Agent is the vSphere ESX Agent Manager managed object responsible
// for deploying an Agency on a single host. The Agent maintains the state
// of the current deployment in its runtime information
type Agent struct {
	EamObject `yaml:",inline"`

	Config  types.AgentConfigInfo  `json:"config,omitempty"`
	Runtime types.AgentRuntimeInfo `json:"runtime,omitempty"`
}
