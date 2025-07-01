// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package object_test

import (
	"context"
	"testing"

	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/session"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/simulator/esx"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/types"
)

func TestAuthorizationManagerPrivilege(t *testing.T) {
	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		m := object.NewAuthorizationManager(c)
		objs := []types.ManagedObjectReference{c.ServiceContent.RootFolder}
		plist := esx.RoleList[0].Privilege[0:2]
		s, _ := session.NewManager(c).UserSession(ctx)

		p, err := m.HasPrivilegeOnEntity(ctx, objs[0], s.Key, plist)
		if err != nil {
			t.Fatal(err)
		}
		if len(p) != len(plist) {
			t.Errorf("HasPrivilegeOnEntity=%v", p)
		}

		ep, err := m.HasUserPrivilegeOnEntities(ctx, objs, s.UserName, plist)
		if err != nil {
			t.Fatal(err)
		}
		if len(ep[0].PrivAvailability) != len(plist) {
			t.Errorf("HasUserPrivilegeOnEntities=%v", ep)
		}

		up, err := m.FetchUserPrivilegeOnEntities(ctx, objs, s.UserName)
		if err != nil {
			t.Fatal(err)
		}
		if len(up) != 1 {
			t.Errorf("FetchUserPrivilegeOnEntities=%v", up)
		}
	})
}
