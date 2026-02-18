// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package mo

import (
	"reflect"

	"github.com/vmware/govmomi/vim25/types"
)

var t = map[string]reflect.Type{}

// TODO: 9.0+ MOs below, not included in the generated mo/mo.go, since the generator still uses older rbvmomi vmodl.db

type DirectPathProfileManager struct {
	Self types.ManagedObjectReference `json:"self"`
}

func (m DirectPathProfileManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["DirectPathProfileManager"] = reflect.TypeOf((*DirectPathProfileManager)(nil)).Elem()
}

type TransitGateway struct {
	ManagedEntity

	Config types.TransitGatewayConfigInfo `json:"config"`
}

func (m TransitGateway) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["TransitGateway"] = reflect.TypeOf((*TransitGateway)(nil)).Elem()
}
