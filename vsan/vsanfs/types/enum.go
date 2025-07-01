// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"reflect"

	"github.com/vmware/govmomi/vim25/types"
)

type VsanFileShareAccessType string

const (
	VsanFileShareAccessTypeREAD_ONLY  = VsanFileShareAccessType("READ_ONLY")
	VsanFileShareAccessTypeREAD_WRITE = VsanFileShareAccessType("READ_WRITE")
	VsanFileShareAccessTypeNO_ACCESS  = VsanFileShareAccessType("NO_ACCESS")
)

func init() {
	types.Add("VsanFileShareAccessType", reflect.TypeOf((*VsanFileShareAccessType)(nil)).Elem())
}
