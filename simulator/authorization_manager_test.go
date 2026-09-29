// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"context"
	"testing"

	"github.com/vmware/govmomi/fault"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/simulator/vpx"
	"github.com/vmware/govmomi/vim25"
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

func TestAuthorizationManagerAddRole(t *testing.T) {
	Test(func(ctx context.Context, c *vim25.Client) {
		m := object.NewAuthorizationManager(c)

		ids := make(map[int32]string)

		for _, name := range []string{"role-a", "role-b"} {
			id, err := m.AddRole(ctx, name, []string{"VirtualMachine.Interact.PowerOn"})
			if err != nil {
				t.Fatal(err)
			}

			if id <= 0 {
				t.Errorf("AddRole(%s) id=%d, want > 0", name, id)
			}

			if other, ok := ids[id]; ok {
				t.Errorf("AddRole(%s) id=%d, already used by %s", name, id, other)
			}
			ids[id] = name

			roles, err := m.RoleList(ctx)
			if err != nil {
				t.Fatal(err)
			}

			role := roles.ById(id)
			if role == nil || role.Name != name {
				t.Errorf("RoleList ById(%d)=%v, want %s", id, role, name)
			}
		}
	})
}

func TestAuthorizationManagerSetEntityPermissions(t *testing.T) {
	Test(func(ctx context.Context, c *vim25.Client) {
		m := object.NewAuthorizationManager(c)
		root := c.ServiceContent.RootFolder

		defaults, err := m.RetrieveEntityPermissions(ctx, root, false)
		if err != nil {
			t.Fatal(err)
		}
		if len(defaults) == 0 {
			t.Fatal("no default permissions on the root folder")
		}

		roles, err := m.RoleList(ctx)
		if err != nil {
			t.Fatal(err)
		}
		admin := roles.ByName("Admin").RoleId
		readOnly := roles.ByName("ReadOnly").RoleId

		set := func(perms ...types.Permission) {
			t.Helper()
			if err := m.SetEntityPermissions(ctx, root, perms); err != nil {
				t.Fatal(err)
			}
		}

		find := func(principal string, group bool) *types.Permission {
			t.Helper()
			perms, err := m.RetrieveEntityPermissions(ctx, root, false)
			if err != nil {
				t.Fatal(err)
			}
			for i := range perms {
				if perms[i].Principal == principal && perms[i].Group == group {
					return &perms[i]
				}
			}
			return nil
		}

		// Each call adds to the list, it does not replace it.
		set(types.Permission{Principal: "VSPHERE.LOCAL\\account-a", RoleId: readOnly})
		set(types.Permission{Principal: "VSPHERE.LOCAL\\account-b", RoleId: readOnly})

		for _, d := range defaults {
			if find(d.Principal, d.Group) == nil {
				t.Errorf("default permission for %s was removed", d.Principal)
			}
		}

		a := find("VSPHERE.LOCAL\\account-a", false)
		if a == nil {
			t.Fatal("account-a permission was removed")
		}
		if a.Entity == nil || *a.Entity != root {
			t.Errorf("account-a permission entity=%v, want %s", a.Entity, root)
		}
		if find("VSPHERE.LOCAL\\account-b", false) == nil {
			t.Error("account-b permission missing")
		}

		// The same principal and group updates the existing permission.
		set(types.Permission{Principal: "VSPHERE.LOCAL\\account-a", RoleId: admin, Propagate: true})

		a = find("VSPHERE.LOCAL\\account-a", false)
		if a == nil || a.RoleId != admin || !a.Propagate {
			t.Errorf("account-a permission=%+v, want role %d propagated", a, admin)
		}

		// A group with the same name is a different principal.
		set(types.Permission{Principal: "VSPHERE.LOCAL\\account-a", Group: true, RoleId: readOnly})

		if g := find("VSPHERE.LOCAL\\account-a", true); g == nil || g.RoleId != readOnly {
			t.Errorf("account-a group permission=%+v, want role %d", g, readOnly)
		}
		if a = find("VSPHERE.LOCAL\\account-a", false); a == nil || a.RoleId != admin {
			t.Errorf("account-a user permission=%+v, want role %d", a, admin)
		}

		perms, err := m.RetrieveEntityPermissions(ctx, root, false)
		if err != nil {
			t.Fatal(err)
		}
		if n := len(defaults) + 3; len(perms) != n {
			t.Errorf("len(permissions)=%d, want %d", len(perms), n)
		}
	})
}

func TestAuthorizationManagerAddRolePrivileges(t *testing.T) {
	Test(func(ctx context.Context, c *vim25.Client) {
		m := object.NewAuthorizationManager(c)

		// Privileges used by vSphere Supervisor roles, on top of the ESX catalogue.
		privs := []string{
			"Global.LogEvent",
			"ManagementServiceAccessGrants.Configure",
			"ManagementServices.Configure",
		}

		if _, err := m.AddRole(ctx, "supervisor-role", privs); err != nil {
			t.Errorf("AddRole(%v): %s", privs, err)
		}

		_, err := m.AddRole(ctx, "invalid-role", []string{"No.Such.Privilege"})
		if !fault.Is(err, &types.InvalidArgument{}) {
			t.Errorf("AddRole(No.Such.Privilege) err=%v, want InvalidArgument", err)
		}
	})
}
