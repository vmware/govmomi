// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"reflect"

	"github.com/vmware/govmomi/vim25/types"
)

func (b *VslmFault) GetVslmFault() *VslmFault { return b }

type BaseVslmFault interface {
	GetVslmFault() *VslmFault
}

func init() {
	types.Add("BaseVslmFault", reflect.TypeOf((*VslmFault)(nil)).Elem())
}

func (b *VslmTaskReason) GetVslmTaskReason() *VslmTaskReason { return b }

type BaseVslmTaskReason interface {
	GetVslmTaskReason() *VslmTaskReason
}

func init() {
	types.Add("BaseVslmTaskReason", reflect.TypeOf((*VslmTaskReason)(nil)).Elem())
}
