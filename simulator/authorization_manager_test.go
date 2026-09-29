// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"context"
	"testing"

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
