// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

// Package authz provides a client for the vAPI com.vmware.cis.authz.permission
// service, which manages vCenter Global Permissions over the legacy /rest endpoint.
//
// The service is deprecated as of vSphere 9.2 in favor of /api/vcenter/authorization/permissions.
package authz

import (
	"context"
	"net/http"
	"strings"

	"github.com/vmware/govmomi/vapi/rest"
)

const (
	// PermissionPath is the endpoint for the permission service.
	PermissionPath = "/com/vmware/cis/authz/permission"

	// GlobalPermissionType is the DynamicID.Type of the Global Permissions resource.
	GlobalPermissionType = "PermissionFolder"
	// GlobalPermissionID is the DynamicID.ID of the Global Permissions resource.
	GlobalPermissionID = "global-permission"
	// GlobalPermissionsDoc is the access control document that holds Global Permissions.
	GlobalPermissionsDoc = "urn:acl:global:permissions"

	// AlreadyExistsGlobal is the message id of the error returned when a principal
	// already holds a Global Permission.
	AlreadyExistsGlobal = "com.vmware.vcenter.authorization.permissions.already_exists_global"

	// permissionIDUserPrefix, permissionIDGroupPrefix and permissionIDDocSeparator
	// form a permission identifier: "username=<principal>;doc=<document>".
	permissionIDUserPrefix   = "username="
	permissionIDGroupPrefix  = "groupname="
	permissionIDDocSeparator = ";doc="
)

// GlobalPermission is the resource id of the Global Permissions resource.
var GlobalPermission = DynamicID{Type: GlobalPermissionType, ID: GlobalPermissionID}

// PrincipalType is the kind of principal that holds a permission.
type PrincipalType string

const (
	// PrincipalUser is an individual user.
	PrincipalUser = PrincipalType("USER")
	// PrincipalGroup is a group of users.
	PrincipalGroup = PrincipalType("GROUP")
)

// DynamicID identifies the resource a permission applies to.
type DynamicID struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

// Principal is the user or group that holds a permission.
// UserName is set when Type is PrincipalUser, GroupName when Type is PrincipalGroup.
// Names are in "DOMAIN\name" form.
type Principal struct {
	Type      PrincipalType `json:"type"`
	UserName  string        `json:"user_name,omitempty"`
	GroupName string        `json:"group_name,omitempty"`
}

// Name returns the user or group name of the principal, according to its Type.
func (p Principal) Name() string {
	if p.Type == PrincipalGroup {
		return p.GroupName
	}
	return p.UserName
}

// CreateSpec is the specification of a new permission.
type CreateSpec struct {
	ResourceID DynamicID `json:"resource_id"`
	Principal  Principal `json:"principal"`
	RoleID     []string  `json:"role_id"`
	Propagate  bool      `json:"propagate"`
}

// Info describes an existing permission.
type Info struct {
	ID         string    `json:"id"`
	ResourceID DynamicID `json:"resource_id"`
	Principal  Principal `json:"principal"`
	RoleID     []string  `json:"role_id"`
	Propagate  bool      `json:"propagate"`
}

// PermissionID returns the identifier of the permission principal holds in the given access control document,
// such as GlobalPermissionsDoc.
func PermissionID(doc string, principal Principal) string {
	prefix := permissionIDUserPrefix
	if principal.Type == PrincipalGroup {
		prefix = permissionIDGroupPrefix
	}
	return prefix + principal.Name() + permissionIDDocSeparator + doc
}

// ParsePermissionID is the inverse of PermissionID.
// It returns false if id is not in the "username=<name>;doc=<doc>" or "groupname=<name>;doc=<doc>" form.
func ParsePermissionID(id string) (string, Principal, bool) {
	var principal Principal

	tail, ok := strings.CutPrefix(id, permissionIDUserPrefix)
	if ok {
		principal.Type = PrincipalUser
	} else if tail, ok = strings.CutPrefix(id, permissionIDGroupPrefix); ok {
		principal.Type = PrincipalGroup
	} else {
		return "", principal, false
	}

	name, doc, ok := strings.Cut(tail, permissionIDDocSeparator)
	if !ok || name == "" || doc == "" {
		return "", principal, false
	}

	if principal.Type == PrincipalGroup {
		principal.GroupName = name
	} else {
		principal.UserName = name
	}
	return doc, principal, true
}

// Manager extends rest.Client, adding permission related methods.
type Manager struct {
	*rest.Client
}

// NewManager creates a new Manager instance with the given client.
func NewManager(client *rest.Client) *Manager {
	return &Manager{
		Client: client,
	}
}

// CreatePermission creates a permission and returns its identifier.
func (c *Manager) CreatePermission(ctx context.Context, spec CreateSpec) (string, error) {
	url := c.Resource(PermissionPath)
	body := struct {
		Spec CreateSpec `json:"create_spec"`
	}{spec}
	var res string
	return res, c.Do(ctx, url.Request(http.MethodPost, body), &res)
}

// CreateGlobalPermission grants the roles to principal as a Global Permission and returns its identifier.
func (c *Manager) CreateGlobalPermission(ctx context.Context, principal Principal, roleIDs []string, propagate bool) (string, error) {
	return c.CreatePermission(ctx, CreateSpec{
		ResourceID: GlobalPermission,
		Principal:  principal,
		RoleID:     roleIDs,
		Propagate:  propagate,
	})
}

// DeletePermission deletes the permission with the given identifier.
func (c *Manager) DeletePermission(ctx context.Context, id string) error {
	url := c.Resource(PermissionPath).WithID(id)
	return c.Do(ctx, url.Request(http.MethodDelete), nil)
}

// GetPermission returns the permission with the given identifier.
func (c *Manager) GetPermission(ctx context.Context, id string) (*Info, error) {
	url := c.Resource(PermissionPath).WithID(id)
	var res Info
	if err := c.Do(ctx, url.Request(http.MethodGet), &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// ListPermissions returns the identifiers of all permissions.
func (c *Manager) ListPermissions(ctx context.Context) ([]string, error) {
	url := c.Resource(PermissionPath)
	var res []string
	return res, c.Do(ctx, url.Request(http.MethodGet), &res)
}

// ListPermissionsDetail returns all permissions.
func (c *Manager) ListPermissionsDetail(ctx context.Context) ([]Info, error) {
	url := c.Resource(PermissionPath).WithAction("list-detail")
	var res []Info
	return res, c.Do(ctx, url.Request(http.MethodPost), &res)
}
