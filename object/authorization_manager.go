// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package object

import (
	"context"

	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type AuthorizationManager struct {
	Common
}

func NewAuthorizationManager(c *vim25.Client) *AuthorizationManager {
	m := AuthorizationManager{
		Common: NewCommon(c, *c.ServiceContent.AuthorizationManager),
	}

	return &m
}

type AuthorizationRoleList []types.AuthorizationRole

func (l AuthorizationRoleList) ById(id int32) *types.AuthorizationRole {
	for _, role := range l {
		if role.RoleId == id {
			return &role
		}
	}

	return nil
}

func (l AuthorizationRoleList) ByName(name string) *types.AuthorizationRole {
	for _, role := range l {
		if role.Name == name {
			return &role
		}
	}

	return nil
}

func (m AuthorizationManager) RoleList(ctx context.Context) (AuthorizationRoleList, error) {
	var am mo.AuthorizationManager

	err := m.Properties(ctx, m.Reference(), []string{"roleList"}, &am)
	if err != nil {
		return nil, err
	}

	return AuthorizationRoleList(am.RoleList), nil
}

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

func (m AuthorizationManager) SetEntityPermissions(ctx context.Context, entity types.ManagedObjectReference, permission []types.Permission) error {
	req := types.SetEntityPermissions{
		This:       m.Reference(),
		Entity:     entity,
		Permission: permission,
	}

	_, err := methods.SetEntityPermissions(ctx, m.Client(), &req)
	return err
}

func (m AuthorizationManager) RetrieveRolePermissions(ctx context.Context, id int32) ([]types.Permission, error) {
	req := types.RetrieveRolePermissions{
		This:   m.Reference(),
		RoleId: id,
	}

	res, err := methods.RetrieveRolePermissions(ctx, m.Client(), &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}

func (m AuthorizationManager) RetrieveAllPermissions(ctx context.Context) ([]types.Permission, error) {
	req := types.RetrieveAllPermissions{
		This: m.Reference(),
	}

	res, err := methods.RetrieveAllPermissions(ctx, m.Client(), &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}

func (m AuthorizationManager) AddRole(ctx context.Context, name string, ids []string) (int32, error) {
	req := types.AddAuthorizationRole{
		This:    m.Reference(),
		Name:    name,
		PrivIds: ids,
	}

	res, err := methods.AddAuthorizationRole(ctx, m.Client(), &req)
	if err != nil {
		return -1, err
	}

	return res.Returnval, nil
}

func (m AuthorizationManager) RemoveRole(ctx context.Context, id int32, failIfUsed bool) error {
	req := types.RemoveAuthorizationRole{
		This:       m.Reference(),
		RoleId:     id,
		FailIfUsed: failIfUsed,
	}

	_, err := methods.RemoveAuthorizationRole(ctx, m.Client(), &req)
	return err
}

func (m AuthorizationManager) UpdateRole(ctx context.Context, id int32, name string, ids []string) error {
	req := types.UpdateAuthorizationRole{
		This:    m.Reference(),
		RoleId:  id,
		NewName: name,
		PrivIds: ids,
	}

	_, err := methods.UpdateAuthorizationRole(ctx, m.Client(), &req)
	return err
}

func (m AuthorizationManager) HasUserPrivilegeOnEntities(ctx context.Context, entities []types.ManagedObjectReference, userName string, privID []string) ([]types.EntityPrivilege, error) {
	req := types.HasUserPrivilegeOnEntities{
		This:     m.Reference(),
		Entities: entities,
		UserName: userName,
		PrivId:   privID,
	}

	res, err := methods.HasUserPrivilegeOnEntities(ctx, m.Client(), &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}

func (m AuthorizationManager) HasPrivilegeOnEntity(ctx context.Context, entity types.ManagedObjectReference, sessionID string, privID []string) ([]bool, error) {
	req := types.HasPrivilegeOnEntity{
		This:      m.Reference(),
		Entity:    entity,
		SessionId: sessionID,
		PrivId:    privID,
	}

	res, err := methods.HasPrivilegeOnEntity(ctx, m.Client(), &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}

func (m AuthorizationManager) FetchUserPrivilegeOnEntities(ctx context.Context, entities []types.ManagedObjectReference, userName string) ([]types.UserPrivilegeResult, error) {
	req := types.FetchUserPrivilegeOnEntities{
		This:     m.Reference(),
		Entities: entities,
		UserName: userName,
	}

	res, err := methods.FetchUserPrivilegeOnEntities(ctx, m.Client(), &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}
