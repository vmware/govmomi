// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package authz_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vapi/cis/authz"
	"github.com/vmware/govmomi/vapi/rest"
	"github.com/vmware/govmomi/vim25"

	_ "github.com/vmware/govmomi/vapi/cis/authz/simulator"
	_ "github.com/vmware/govmomi/vapi/simulator"
)

func TestPermissionID(t *testing.T) {
	user := authz.Principal{Type: authz.PrincipalUser, UserName: `VSPHERE.LOCAL\wcp`}
	group := authz.Principal{Type: authz.PrincipalGroup, GroupName: `VSPHERE.LOCAL\Administrators`}

	id := authz.PermissionID(authz.GlobalPermissionsDoc, user)
	assert.Equal(t, `username=VSPHERE.LOCAL\wcp;doc=urn:acl:global:permissions`, id)
	doc, p, ok := authz.ParsePermissionID(id)
	assert.True(t, ok)
	assert.Equal(t, authz.GlobalPermissionsDoc, doc)
	assert.Equal(t, user, p)

	id = authz.PermissionID(authz.GlobalPermissionsDoc, group)
	assert.Equal(t, `groupname=VSPHERE.LOCAL\Administrators;doc=urn:acl:global:permissions`, id)
	doc, p, ok = authz.ParsePermissionID(id)
	assert.True(t, ok)
	assert.Equal(t, authz.GlobalPermissionsDoc, doc)
	assert.Equal(t, group, p)

	for _, bad := range []string{"", "username=", "username=x", "username=;doc=d", "username=x;doc=", "other=x;doc=d"} {
		_, _, ok = authz.ParsePermissionID(bad)
		assert.False(t, ok, bad)
	}
}

func TestGlobalPermission(t *testing.T) {
	simulator.Test(func(ctx context.Context, vc *vim25.Client) {
		rc := rest.NewClient(vc)
		require.NoError(t, rc.Login(ctx, simulator.DefaultLogin))

		m := authz.NewManager(rc)
		user := authz.Principal{Type: authz.PrincipalUser, UserName: `VSPHERE.LOCAL\wcp`}

		ids, err := m.ListPermissions(ctx)
		require.NoError(t, err)
		assert.Empty(t, ids)

		id, err := m.CreateGlobalPermission(ctx, user, []string{"-1"}, true)
		require.NoError(t, err)
		assert.Equal(t, authz.PermissionID(authz.GlobalPermissionsDoc, user), id)

		info, err := m.GetPermission(ctx, id)
		require.NoError(t, err)
		assert.Equal(t, authz.Info{
			ID:         id,
			ResourceID: authz.GlobalPermission,
			Principal:  user,
			RoleID:     []string{"-1"},
			Propagate:  true,
		}, *info)

		ids, err = m.ListPermissions(ctx)
		require.NoError(t, err)
		assert.Equal(t, []string{id}, ids)

		infos, err := m.ListPermissionsDetail(ctx)
		require.NoError(t, err)
		assert.Equal(t, []authz.Info{*info}, infos)

		_, err = m.CreateGlobalPermission(ctx, user, []string{"-2"}, false)
		require.Error(t, err)
		assert.Contains(t, err.Error(), authz.AlreadyExistsGlobal)

		require.NoError(t, m.DeletePermission(ctx, id))

		_, err = m.GetPermission(ctx, id)
		assert.True(t, rest.IsStatusError(err, 404), err)

		err = m.DeletePermission(ctx, id)
		assert.True(t, rest.IsStatusError(err, 404), err)
	})
}
