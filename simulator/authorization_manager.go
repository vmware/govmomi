/*
Copyright (c) 2017 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package simulator

import (
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/simulator/esx"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

type AuthorizationManager struct {
	mo.AuthorizationManager
}

func NewAuthorizationManager(ref types.ManagedObjectReference) object.Reference {
	s := &AuthorizationManager{}
	s.Self = ref
	s.RoleList = esx.RoleList
	return s
}

func (m *AuthorizationManager) RetrieveEntityPermissions(req *types.RetrieveEntityPermissions) soap.HasFault {
	var p []types.Permission

	for _, u := range DefaultUserGroup {
		p = append(p, types.Permission{
			Entity:    &req.Entity,
			Principal: u.Principal,
			Group:     u.Group,
			RoleId:    -1, // "Admin"
		})
	}

	return &methods.RetrieveEntityPermissionsBody{
		Res: &types.RetrieveEntityPermissionsResponse{
			Returnval: p,
		},
	}
}

func (m *AuthorizationManager) RetrieveAllPermissions(req *types.RetrieveAllPermissions) soap.HasFault {
	var p []types.Permission

	root := Map.content().RootFolder

	for _, u := range DefaultUserGroup {
		p = append(p, types.Permission{
			Entity:    &root,
			Principal: u.Principal,
			Group:     u.Group,
			RoleId:    -1, // "Admin"
		})
	}

	return &methods.RetrieveAllPermissionsBody{
		Res: &types.RetrieveAllPermissionsResponse{
			Returnval: p,
		},
	}
}

func (m *AuthorizationManager) RetrieveRolePermissions(req *types.RetrieveRolePermissions) soap.HasFault {
	var p []types.Permission

	root := Map.content().RootFolder

	for _, u := range DefaultUserGroup {
		p = append(p, types.Permission{
			Entity:    &root,
			Principal: u.Principal,
			Group:     u.Group,
			RoleId:    req.RoleId,
		})
	}

	return &methods.RetrieveRolePermissionsBody{
		Res: &types.RetrieveRolePermissionsResponse{
			Returnval: p,
		},
	}
}
