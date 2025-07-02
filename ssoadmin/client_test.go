// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package ssoadmin_test

import (
	"context"
	"fmt"
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
	t.Run("With DNS errors, lookup client should rewrite URLs to VC's Host", func(t *testing.T) {
		simulator.Test(func(ctx context.Context, client *vim25.Client) {
			lsim.UnresolveLookupServiceURLs(ctx)

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

			lsim.BreakLookupServiceURLs(ctx)

			c, err := ssoadmin.NewClient(ctx, client)
			require.NoError(t, err)

			verifyClient(t, ctx, c)
		}, model)
	})
	t.Run("System.Anonymous methods", func(t *testing.T) {
		simulator.Test(func(ctx context.Context, client *vim25.Client) {
			c, err := ssoadmin.NewClient(ctx, client)
			require.NoError(t, err)

			c.Jar = nil // session cookie will not be sent

			_, err = c.FindUser(ctx, "testuser")
			require.Error(t, err) // NotAuthenticated

			certs, err := c.GetTrustedCertificates(ctx)
			require.NoError(t, err)
			fmt.Println(certs[0])
			require.NotEmpty(t, certs)
		})
	})
}

func verifyClient(t *testing.T, ctx context.Context, c *ssoadmin.Client) {
	err := c.CreatePersonUser(ctx, "testuser", types.AdminPersonDetails{FirstName: "test", LastName: "user"}, "password")
	require.NoError(t, err)

	user, err := c.FindUser(ctx, "testuser")
	require.NoError(t, err)
	require.Equal(t, &types.AdminUser{Id: types.PrincipalId{Name: "testuser", Domain: "vsphere.local"}, Kind: "person"}, user)
}
