// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package mo

import (
	"github.com/vmware/govmomi/eam/types"
	vim "github.com/vmware/govmomi/vim25/types"
)

// EamObject contains the fields common to all EAM objects.
type EamObject struct {
	Self  vim.ManagedObjectReference `json:"self"`
	Issue []types.BaseIssue          `json:"issue,omitempty"`
}

func (m EamObject) String() string {
	return m.Self.String()
}

func (m EamObject) Reference() vim.ManagedObjectReference {
	return m.Self
}
