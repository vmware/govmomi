// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package esx

import "github.com/vmware/govmomi/vim25/types"

// HostCapability captured via `govc object.collect -dump $host capability`
var HostCapability = &types.HostCapability{
	MaxSupportedVmMemory: 25149440, // 25TB since 7.0U1
}
