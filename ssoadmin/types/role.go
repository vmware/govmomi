// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"reflect"

	"github.com/vmware/govmomi/vim25/types"
)

// Types here are not included in the wsdl

const (
	RoleActAsUser     = "ActAsUser"
	RoleRegularUser   = "RegularUser"
	RoleAdministrator = "Administrator"
)

type GrantWSTrustRole GrantWSTrustRoleRequestType

func init() {
	types.Add("sso:GrantWSTrustRole", reflect.TypeOf((*GrantWSTrustRole)(nil)).Elem())
}

type GrantWSTrustRoleRequestType struct {
	This   types.ManagedObjectReference `xml:"_this"`
	UserId PrincipalId                  `xml:"userId"`
	Role   string                       `xml:"role"`
}

func init() {
	types.Add("sso:GrantWSTrustRoleRequestType", reflect.TypeOf((*GrantWSTrustRoleRequestType)(nil)).Elem())
}

type GrantWSTrustRoleResponse struct {
	Returnval bool `xml:"returnval"`
}

type RevokeWSTrustRole RevokeWSTrustRoleRequestType

func init() {
	types.Add("sso:RevokeWSTrustRole", reflect.TypeOf((*RevokeWSTrustRole)(nil)).Elem())
}

type RevokeWSTrustRoleRequestType struct {
	This   types.ManagedObjectReference `xml:"_this"`
	UserId PrincipalId                  `xml:"userId"`
	Role   string                       `xml:"role"`
}

func init() {
	types.Add("sso:RevokeWSTrustRoleRequestType", reflect.TypeOf((*RevokeWSTrustRoleRequestType)(nil)).Elem())
}

type RevokeWSTrustRoleResponse struct {
	Returnval bool `xml:"returnval"`
}
