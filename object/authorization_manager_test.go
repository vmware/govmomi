/*
Copyright (c) 2020 VMware, Inc. All Rights Reserved.

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
