/*
Copyright (c) 2018-2023 VMware, Inc. All Rights Reserved.

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

package ssoadmin_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	lsim "github.com/vmware/govmomi/lookup/simulator"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/ssoadmin"
	_ "github.com/vmware/govmomi/ssoadmin/simulator"
	"github.com/vmware/govmomi/ssoadmin/types"
	_ "github.com/vmware/govmomi/sts/simulator"
	"github.com/vmware/govmomi/vim25"
)

func TestClient(t *testing.T) {
	t.Run("Happy path using lookup service", func(t *testing.T) {
		simulator.Test(func(ctx context.Context, client *vim25.Client) {
			c, err := ssoadmin.NewClient(ctx, client)
			require.NoError(t, err)

			verifyClient(t, ctx, c)
		})
	})
	t.Run("With Envoy sidecar and a malfunctioning lookup service, ssoadmin client creation should still succeed", func(t *testing.T) {
		model := simulator.VPX()
		model.Create()
		simulator.Test(func(ctx context.Context, client *vim25.Client) {
			// Map Envoy sidecar on the same port as the vcsim client.
			os.Setenv("GOVMOMI_ENVOY_SIDECAR_PORT", client.Client.URL().Port())
			os.Setenv("GOVMOMI_ENVOY_SIDECAR_HOST", client.Client.URL().Hostname())

			lsim.BreakLookupServiceURLs()

			c, err := ssoadmin.NewClient(ctx, client)
			require.NoError(t, err)

			verifyClient(t, ctx, c)
		}, model)
	})
}

func verifyClient(t *testing.T, ctx context.Context, c *ssoadmin.Client) {
	err := c.CreatePersonUser(ctx, "testuser", types.AdminPersonDetails{FirstName: "test", LastName: "user"}, "password")
	require.NoError(t, err)

	user, err := c.FindUser(ctx, "testuser")
	require.NoError(t, err)
	require.Equal(t, &types.AdminUser{Id: types.PrincipalId{Name: "testuser", Domain: "vsphere.local"}, Kind: "person"}, user)

}
