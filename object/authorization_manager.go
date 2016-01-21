/*
Copyright (c) 2015 VMware, Inc. All Rights Reserved.

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

package object

import (
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
	"golang.org/x/net/context"
)

// AuthorizationManager represents an authorization manager
type AuthorizationManager struct {
	Common
}

// NewAuthorizationManager creates an authorization manager
func NewAuthorizationManager(c *vim25.Client) *AuthorizationManager {
	m := AuthorizationManager{
		Common: NewCommon(c, *c.ServiceContent.AuthorizationManager),
	}

	return &m
}

// AuthorizationRoleList respresents a list of authorization roles
type AuthorizationRoleList []types.AuthorizationRole

// ByID finds the role by id
func (l AuthorizationRoleList) ByID(id int) *types.AuthorizationRole {
	for _, role := range l {
		if role.RoleId == id {
			return &role
		}
	}

	return nil
}

// ByName finds the role by name
func (l AuthorizationRoleList) ByName(name string) *types.AuthorizationRole {
	for _, role := range l {
		if role.Name == name {
			return &role
		}
	}

	return nil
}

// RoleList lists the known roles
func (m AuthorizationManager) RoleList(ctx context.Context) (AuthorizationRoleList, error) {
	var am mo.AuthorizationManager

	err := m.Properties(ctx, m.Reference(), []string{"roleList"}, &am)
	if err != nil {
		return nil, err
	}

	return AuthorizationRoleList(am.RoleList), nil
}

// RetrieveEntityPermissions retrieves entity permissions
func (m AuthorizationManager) RetrieveEntityPermissions(ctx context.Context, entity types.ManagedObjectReference, inherited bool) ([]types.Permission, error) {
	req := types.RetrieveEntityPermissions{
		This:      m.Reference(),
		Entity:    entity,
		Inherited: inherited,
	}

	res, err := methods.RetrieveEntityPermissions(ctx, m.Client(), &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}

// RemoveEntityPermission removes an entity permission
func (m AuthorizationManager) RemoveEntityPermission(ctx context.Context, entity types.ManagedObjectReference, user string, isGroup bool) error {
	req := types.RemoveEntityPermission{
		This:    m.Reference(),
		Entity:  entity,
		User:    user,
		IsGroup: isGroup,
	}

	_, err := methods.RemoveEntityPermission(ctx, m.Client(), &req)
	return err
}

// SetEntityPermissions set entity permissions
func (m AuthorizationManager) SetEntityPermissions(ctx context.Context, entity types.ManagedObjectReference, permission []types.Permission) error {
	req := types.SetEntityPermissions{
		This:       m.Reference(),
		Entity:     entity,
		Permission: permission,
	}

	_, err := methods.SetEntityPermissions(ctx, m.Client(), &req)
	return err
}
