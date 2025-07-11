// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"strings"

	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/simulator/esx"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

type AuthorizationManager struct {
	mo.AuthorizationManager

	permissions map[types.ManagedObjectReference][]types.Permission
	privileges  map[string]struct{}
	system      []string
	nextID      int32
}

func (m *AuthorizationManager) init(r *Registry) {
	if len(m.RoleList) == 0 {
		m.RoleList = make([]types.AuthorizationRole, len(esx.RoleList))
		copy(m.RoleList, esx.RoleList)
	}

	m.permissions = make(map[types.ManagedObjectReference][]types.Permission)

	l := object.AuthorizationRoleList(m.RoleList)
	m.system = l.ByName("ReadOnly").Privilege
	admin := l.ByName("Admin")
	m.privileges = make(map[string]struct{}, len(admin.Privilege))

	for _, id := range admin.Privilege {
		m.privileges[id] = struct{}{}
	}

	root := r.content().RootFolder

	for _, u := range DefaultUserGroup {
		m.permissions[root] = append(m.permissions[root], types.Permission{
			Entity:    &root,
			Principal: u.Principal,
			Group:     u.Group,
			RoleId:    admin.RoleId,
			Propagate: true,
		})
	}
}

func (m *AuthorizationManager) RetrieveEntityPermissions(ctx *Context, req *types.RetrieveEntityPermissions) soap.HasFault {
	e := ctx.Map.Get(req.Entity).(mo.Entity)

	p := m.permissions[e.Reference()]

	if req.Inherited {
		for {
			parent := e.Entity().Parent
			if parent == nil {
				break
			}

			e = ctx.Map.Get(parent.Reference()).(mo.Entity)

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

func (m *AuthorizationManager) HasPrivilegeOnEntities(req *types.HasPrivilegeOnEntities) soap.HasFault {
	var p []types.EntityPrivilege

	for _, e := range req.Entity {
		priv := types.EntityPrivilege{Entity: e}

		for _, id := range req.PrivId {
			priv.PrivAvailability = append(priv.PrivAvailability, types.PrivilegeAvailability{
				PrivId:    id,
				IsGranted: true,
			})
		}

		p = append(p, priv)
	}

	return &methods.HasPrivilegeOnEntitiesBody{
		Res: &types.HasPrivilegeOnEntitiesResponse{
			Returnval: p,
		},
	}
}

func (m *AuthorizationManager) HasPrivilegeOnEntity(req *types.HasPrivilegeOnEntity) soap.HasFault {
	p := make([]bool, len(req.PrivId))

	for i := range req.PrivId {
		p[i] = true
	}

	return &methods.HasPrivilegeOnEntityBody{
		Res: &types.HasPrivilegeOnEntityResponse{
			Returnval: p,
		},
	}
}

func (m *AuthorizationManager) HasUserPrivilegeOnEntities(req *types.HasUserPrivilegeOnEntities) soap.HasFault {
	var p []types.EntityPrivilege

	for _, e := range req.Entities {
		priv := types.EntityPrivilege{Entity: e}

		for _, id := range req.PrivId {
			priv.PrivAvailability = append(priv.PrivAvailability, types.PrivilegeAvailability{
				PrivId:    id,
				IsGranted: true,
			})
		}

		p = append(p, priv)
	}

	return &methods.HasUserPrivilegeOnEntitiesBody{
		Res: &types.HasUserPrivilegeOnEntitiesResponse{
			Returnval: p,
		},
	}
}

func (m *AuthorizationManager) FetchUserPrivilegeOnEntities(req *types.FetchUserPrivilegeOnEntities) soap.HasFault {
	admin := object.AuthorizationRoleList(m.RoleList).ByName("Admin").Privilege

	var p []types.UserPrivilegeResult

	for _, e := range req.Entities {
		p = append(p, types.UserPrivilegeResult{
			Entity:     e,
			Privileges: admin,
		})
	}

	return &methods.FetchUserPrivilegeOnEntitiesBody{
		Res: &types.FetchUserPrivilegeOnEntitiesResponse{
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

	ids, err := m.privIDs(req.PrivIds)
	if err != nil {
		body.Fault_ = err
		return body
	}

	m.RoleList = append(m.RoleList, types.AuthorizationRole{
		Info: &types.Description{
			Label:   req.Name,
			Summary: req.Name,
		},
		RoleId:    m.nextID,
		Privilege: ids,
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
			if len(req.PrivIds) != 0 {
				ids, err := m.privIDs(req.PrivIds)
				if err != nil {
					body.Fault_ = err
					return body
				}
				m.RoleList[i].Privilege = ids
			}

			m.RoleList[i].Name = req.NewName

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

func (m *AuthorizationManager) privIDs(ids []string) ([]string, *soap.Fault) {
	system := make(map[string]struct{}, len(m.system))

	for _, id := range ids {
		if _, ok := m.privileges[id]; !ok {
			return nil, Fault("", &types.InvalidArgument{InvalidProperty: "privIds"})
		}

		if strings.HasPrefix(id, "System.") {
			system[id] = struct{}{}
		}
	}

	for _, id := range m.system {
		if _, ok := system[id]; ok {
			continue
		}

		ids = append(ids, id)
	}

	return ids, nil
}
