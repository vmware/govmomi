// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package mo

import (
	"github.com/vmware/govmomi/vim25/types"
)

type CnsVolumeManager struct {
	Self types.ManagedObjectReference
}

func (m CnsVolumeManager) Reference() types.ManagedObjectReference {
	return m.Self
}
