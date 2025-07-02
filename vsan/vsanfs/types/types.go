// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"reflect"

	"github.com/vmware/govmomi/vim25/types"
)

type VsanFileShareNetPermission struct {
	Ips         string                  `xml:"ips"`
	Permissions VsanFileShareAccessType `xml:"permissions,omitempty,typeattr"`
	AllowRoot   bool                    `xml:"allowRoot,omitempty"`
}

func init() {
	types.Add("VsanFileShareNetPermission", reflect.TypeOf((*VsanFileShareNetPermission)(nil)).Elem())
}
