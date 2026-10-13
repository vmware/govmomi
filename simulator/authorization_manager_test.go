// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"context"
	"net/url"
	"testing"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/fault"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/session"
	"github.com/vmware/govmomi/simulator/vpx"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/types"
)

func TestAuthorizationManager(t *testing.T) {
	for range 2 {
		model := VPX()
		ctx := NewContext()
		_ = New(NewServiceInstance(ctx, model.ServiceContent, model.RootFolder)) // 2nd pass panics w/o copying RoleList

		authz := ctx.Map.Get(*vpx.ServiceContent.AuthorizationManager).(*AuthorizationManager)
		authz.RemoveAuthorizationRole(&types.RemoveAuthorizationRole{
			RoleId: -2, // ReadOnly
		})
	}
}

func TestAuthorizationManagerPrivileges(t *testing.T) {
	for i, model := range []*Model{VPX(), ESX()} {
		name := []string{"vpx", "esx"}[i]
		t.Run(name, func(t *testing.T) {
			err := model.Run(func(ctx context.Context, c *vim25.Client) error {
				m := object.NewAuthorizationManager(c)
				root := c.ServiceContent.RootFolder
				children, err := object.NewFolder(c, root).Children(ctx)
				if err != nil {
					return err
				}
				if len(children) == 0 {
					t.Fatal("root folder has no child entities")
				}
				child := children[0].Reference()

				roleA, err := m.AddRole(ctx, "alice-role", []string{"Alarm.Acknowledge"})
				if err != nil {
					return err
				}
				roleB, err := m.AddRole(ctx, "bob-role", []string{"Alarm.Create"})
				if err != nil {
					return err
				}
				if roleA == roleB {
					t.Fatalf("roles have the same ID: %d", roleA)
				}

				if err := m.SetEntityPermissions(ctx, root, []types.Permission{
					{Principal: "alice", RoleId: roleA, Propagate: true},
					{Principal: "bob", RoleId: roleB, Propagate: false},
				}); err != nil {
					return err
				}
				permissions, err := m.RetrieveEntityPermissions(ctx, root, false)
				if err != nil {
					return err
				}
				for _, permission := range permissions {
					if permission.Principal == "alice" && (permission.Entity == nil || *permission.Entity != root) {
						t.Errorf("permission Entity was not normalized: %+v", permission)
					}
				}

				requested := []string{"Alarm.Acknowledge", "Alarm.Create"}
				alice, err := m.HasUserPrivilegeOnEntities(ctx, []types.ManagedObjectReference{child}, "alice", requested)
				if err != nil {
					return err
				}
				if !availability(alice[0], requested[0]) || availability(alice[0], requested[1]) {
					t.Errorf("alice inherited unexpected privileges: %+v", alice[0])
				}
				bob, err := m.HasUserPrivilegeOnEntities(ctx, []types.ManagedObjectReference{root, child}, "bob", requested)
				if err != nil {
					return err
				}
				if !availability(bob[0], requested[1]) || availability(bob[1], requested[1]) {
					t.Errorf("non-propagating permission did not stay on root: %+v", bob)
				}

				fetched, err := m.FetchUserPrivilegeOnEntities(ctx, []types.ManagedObjectReference{child}, "alice")
				if err != nil {
					return err
				}
				if !containsPrivilege(fetched[0].Privileges, requested[0]) || containsPrivilege(fetched[0].Privileges, requested[1]) {
					t.Errorf("FetchUserPrivilegeOnEntities=%v", fetched)
				}
				// Non-entity references are checked against the inventory root.
				fetched, err = m.FetchUserPrivilegeOnEntities(ctx, []types.ManagedObjectReference{*c.ServiceContent.AuthorizationManager}, "alice")
				if err != nil {
					return err
				}
				if !containsPrivilege(fetched[0].Privileges, requested[0]) {
					t.Errorf("non-entity privilege query did not use root: %v", fetched[0].Privileges)
				}

				// Exercise session-based checks as the same assigned user on a
				// separate unauthenticated connection.
				aliceClient, err := govmomi.NewClient(ctx, c.URL(), true)
				if err != nil {
					return err
				}
				defer aliceClient.Logout(ctx)
				sm := session.NewManager(aliceClient.Client)
				if err := sm.Login(ctx, url.UserPassword("alice", "password")); err != nil {
					return err
				}
				aliceSession, err := sm.UserSession(ctx)
				if err != nil {
					return err
				}
				granted, err := m.HasPrivilegeOnEntity(ctx, child, aliceSession.Key, requested)
				if err != nil {
					return err
				}
				// encoding/xml omits false values from []bool fields, so ask for
				// the expected grant here; the entity-privilege APIs cover denials.
				if len(granted) != 1 || !granted[0] {
					t.Errorf("HasPrivilegeOnEntity=%v", granted)
				}
				multi, err := methods.HasPrivilegeOnEntities(ctx, c, &types.HasPrivilegeOnEntities{
					This: *c.ServiceContent.AuthorizationManager, Entity: []types.ManagedObjectReference{child}, SessionId: aliceSession.Key, PrivId: requested,
				})
				if err != nil {
					return err
				}
				if len(multi.Returnval) != 1 || !availability(multi.Returnval[0], requested[0]) || availability(multi.Returnval[0], requested[1]) {
					t.Errorf("HasPrivilegeOnEntities=%+v", multi.Returnval)
				}

				// SetEntityPermissions updates the named principal without erasing others.
				if err := m.SetEntityPermissions(ctx, root, []types.Permission{{Principal: "alice", RoleId: roleB, Propagate: true}}); err != nil {
					return err
				}
				alice, err = m.HasUserPrivilegeOnEntities(ctx, []types.ManagedObjectReference{child}, "alice", requested)
				if err != nil {
					return err
				}
				bob, err = m.HasUserPrivilegeOnEntities(ctx, []types.ManagedObjectReference{root}, "bob", requested)
				if err != nil {
					return err
				}
				if availability(alice[0], requested[0]) || !availability(alice[0], requested[1]) || !availability(bob[0], requested[1]) {
					t.Errorf("permission update replaced or failed to update assignments: alice=%+v bob=%+v", alice, bob)
				}

				// NoAccess explicitly denies otherwise inherited grants.
				if err := m.SetEntityPermissions(ctx, child, []types.Permission{{Principal: "alice", RoleId: -5}}); err != nil {
					return err
				}
				alice, err = m.HasUserPrivilegeOnEntities(ctx, []types.ManagedObjectReference{child}, "alice", requested)
				if err != nil {
					return err
				}
				if availability(alice[0], requested[0]) || availability(alice[0], requested[1]) {
					t.Errorf("NoAccess did not deny privileges: %+v", alice[0])
				}
				if err := m.RemoveEntityPermission(ctx, child, "alice", false); err != nil {
					return err
				}
				if err := m.SetEntityPermissions(ctx, child, []types.Permission{{Principal: "alice", Group: true, RoleId: roleA}}); err != nil {
					return err
				}
				alice, err = m.HasUserPrivilegeOnEntities(ctx, []types.ManagedObjectReference{child}, "alice", requested)
				if err != nil {
					return err
				}
				if !availability(alice[0], requested[0]) || !availability(alice[0], requested[1]) {
					t.Errorf("exact-name group permission did not match: %+v", alice[0])
				}
				if err := m.RemoveEntityPermission(ctx, child, "alice", true); err != nil {
					return err
				}

				if err := m.RemoveRole(ctx, roleB, true); !fault.Is(err, &types.InvalidArgument{}) {
					t.Errorf("RemoveRole(failIfUsed=true) error = %v, want InvalidArgument", err)
				}
				if err := m.RemoveEntityPermission(ctx, root, "bob", false); err != nil {
					return err
				}
				bob, err = m.HasUserPrivilegeOnEntities(ctx, []types.ManagedObjectReference{root}, "bob", requested)
				if err != nil {
					return err
				}
				if availability(bob[0], requested[1]) {
					t.Errorf("removed permission still grants a privilege: %+v", bob[0])
				}
				if err := m.RemoveRole(ctx, roleB, true); !fault.Is(err, &types.InvalidArgument{}) {
					t.Errorf("RemoveRole(failIfUsed=true) error = %v, want InvalidArgument", err)
				}
				if err := m.RemoveEntityPermission(ctx, root, "alice", false); err != nil {
					return err
				}
				if err := m.RemoveRole(ctx, roleB, true); err != nil {
					return err
				}
				return nil
			})
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func availability(privilege types.EntityPrivilege, id string) bool {
	for _, available := range privilege.PrivAvailability {
		if available.PrivId == id {
			return available.IsGranted
		}
	}
	return false
}

func containsPrivilege(privileges []string, id string) bool {
	for _, privilege := range privileges {
		if privilege == id {
			return true
		}
	}
	return false
}
