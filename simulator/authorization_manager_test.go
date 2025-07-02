// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"testing"

	"github.com/vmware/govmomi/simulator/vpx"
	"github.com/vmware/govmomi/vim25/types"
)

func TestAuthorizationManager(t *testing.T) {
	for i := 0; i < 2; i++ {
		model := VPX()
		ctx := NewContext()
		_ = New(NewServiceInstance(ctx, model.ServiceContent, model.RootFolder)) // 2nd pass panics w/o copying RoleList

		authz := ctx.Map.Get(*vpx.ServiceContent.AuthorizationManager).(*AuthorizationManager)
		authz.RemoveAuthorizationRole(&types.RemoveAuthorizationRole{
			RoleId: -2, // ReadOnly
		})
	}
}
