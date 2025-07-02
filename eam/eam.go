// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package eam

import (
	"github.com/vmware/govmomi/eam/internal"
	"github.com/vmware/govmomi/vim25/types"
)

var EsxAgentManager = types.ManagedObjectReference{
	Type:  internal.EsxAgentManager,
	Value: internal.EsxAgentManager,
}
