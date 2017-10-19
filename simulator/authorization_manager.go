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

var systemPrivileges = []string{
	"System.Anonymous",
	"System.View",
	"System.Read",
}

type AuthorizationManager struct {
	mo.AuthorizationManager

	permissions map[types.ManagedObjectReference][]types.Permission

	nextID int32
}

func NewAuthorizationManager(ref types.ManagedObjectReference) object.Reference {
	m := &AuthorizationManager{}
	m.Self = ref
	m.RoleList = esx.RoleList
	m.permissions = make(map[types.ManagedObjectReference][]types.Permission)

	root := Map.content().RootFolder

	for _, u := range DefaultUserGroup {
		m.permissions[root] = append(m.permissions[root], types.Permission{
			Entity:    &root,
			Principal: u.Principal,
			Group:     u.Group,
			RoleId:    -1, // "Admin"
			Propagate: true,
		})
	}

	return m
}

func (m *AuthorizationManager) RetrieveEntityPermissions(req *types.RetrieveEntityPermissions) soap.HasFault {
	e := Map.Get(req.Entity).(mo.Entity)

	p := m.permissions[e.Reference()]

	if req.Inherited {
		for {
			parent := e.Entity().Parent
			if parent == nil {
				break
			}

			e = Map.Get(parent.Reference()).(mo.Entity)

			p = append(p, m.permissions[e.Reference()]...)
		}
	}

	return &methods.RetrieveEntityPermissionsBody{
		Res: &types.RetrieveEntityPermissionsResponse{
			Returnval: p,
		},
	}
}

func (m *AuthorizationManager) RetrieveAllPermissions(req *types.RetrieveAllPermissions) soap.HasFault {
	var p []types.Permission

	for _, v := range m.permissions {
		p = append(p, v...)
	}

	return &methods.RetrieveAllPermissionsBody{
		Res: &types.RetrieveAllPermissionsResponse{
			Returnval: p,
		},
	}
}

func (m *AuthorizationManager) RemoveEntityPermission(req *types.RemoveEntityPermission) soap.HasFault {
	var p []types.Permission

	for _, v := range m.permissions[req.Entity] {
		if v.Group == req.IsGroup && v.Principal == req.User {
			continue
		}
		p = append(p, v)
	}

	m.permissions[req.Entity] = p

	return &methods.RemoveEntityPermissionBody{
		Res: &types.RemoveEntityPermissionResponse{},
	}
}

func (m *AuthorizationManager) SetEntityPermissions(req *types.SetEntityPermissions) soap.HasFault {
	m.permissions[req.Entity] = req.Permission

	return &methods.SetEntityPermissionsBody{
		Res: &types.SetEntityPermissionsResponse{},
	}
}

func (m *AuthorizationManager) RetrieveRolePermissions(req *types.RetrieveRolePermissions) soap.HasFault {
	var p []types.Permission

	for _, set := range m.permissions {
		for _, v := range set {
			if v.RoleId == req.RoleId {
				p = append(p, v)
			}
		}
	}

	return &methods.RetrieveRolePermissionsBody{
		Res: &types.RetrieveRolePermissionsResponse{
			Returnval: p,
		},
	}
}

func (m *AuthorizationManager) AddAuthorizationRole(req *types.AddAuthorizationRole) soap.HasFault {
	body := &methods.AddAuthorizationRoleBody{}

	for _, role := range m.RoleList {
		if role.Name == req.Name {
			body.Fault_ = Fault("", &types.AlreadyExists{})
			return body
		}
	}

	m.RoleList = append(m.RoleList, types.AuthorizationRole{
		Info: &types.Description{
			Label:   req.Name,
			Summary: req.Name,
		},
		RoleId:    m.nextID,
		Privilege: updateSystemPrivileges(req.PrivIds),
		Name:      req.Name,
		System:    false,
	})

	m.nextID++

	body.Res = &types.AddAuthorizationRoleResponse{}

	return body
}

func (m *AuthorizationManager) UpdateAuthorizationRole(req *types.UpdateAuthorizationRole) soap.HasFault {
	body := &methods.UpdateAuthorizationRoleBody{}

	for _, role := range m.RoleList {
		if role.Name == req.NewName && role.RoleId != req.RoleId {
			body.Fault_ = Fault("", &types.AlreadyExists{})
			return body
		}
	}

	for i, role := range m.RoleList {
		if role.RoleId == req.RoleId {
			m.RoleList[i].Name = req.NewName

			if req.PrivIds != nil {
				m.RoleList[i].Privilege = updateSystemPrivileges(req.PrivIds)
			}

			body.Res = &types.UpdateAuthorizationRoleResponse{}
			return body
		}
	}

	body.Fault_ = Fault("", &types.NotFound{})

	return body
}

func (m *AuthorizationManager) RemoveAuthorizationRole(req *types.RemoveAuthorizationRole) soap.HasFault {
	body := &methods.RemoveAuthorizationRoleBody{}

	for i, role := range m.RoleList {
		if role.RoleId == req.RoleId {
			m.RoleList = append(m.RoleList[:i], m.RoleList[i+1:]...)

			body.Res = &types.RemoveAuthorizationRoleResponse{}
			return body
		}
	}

	body.Fault_ = Fault("", &types.NotFound{})

	return body
}

func updateSystemPrivileges(privileges []string) []string {
OUT:
	for _, spr := range systemPrivileges {
		for _, pr := range privileges {
			if spr == pr {
				continue OUT
			}
		}
		privileges = append(privileges, spr)
	}
	return privileges
}
