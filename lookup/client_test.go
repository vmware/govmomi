// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package lookup_test

// lookup/simulator/simulator_test.go has more tests

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/vmware/govmomi/lookup"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/ssoadmin"
	"github.com/vmware/govmomi/sts"
	"github.com/vmware/govmomi/vim25"

	lsim "github.com/vmware/govmomi/lookup/simulator"
	_ "github.com/vmware/govmomi/ssoadmin/simulator"
	_ "github.com/vmware/govmomi/sts/simulator"
)

// test lookup.EndpointURL usage by the ssoadmin and sts clients
func TestEndpointURL(t *testing.T) {
	// these client calls should fail since we'll break the URL paths
	simulator.Test(func(ctx context.Context, vc *vim25.Client) {
		lsim.BreakLookupServiceURLs(ctx)

		{
			_, err := ssoadmin.NewClient(ctx, vc)
			if err == nil {
				t.Error("expected error")
			}
			if !strings.Contains(err.Error(), http.StatusText(404)) {
				t.Error(err)
			}
		}

		{
			c, err := sts.NewClient(ctx, vc)
			if err != nil {
				t.Fatal(err)
			}

			req := sts.TokenRequest{
				Userinfo: url.UserPassword("Administrator@VSPHERE.LOCAL", "password"),
			}
			_, err = c.Issue(ctx, req)
			if err == nil {
				t.Error("expected error")
			}
			if !strings.Contains(err.Error(), http.StatusText(404)) {
				t.Error(err)
			}
		}
	})
	// these client calls should not fail
	simulator.Test(func(ctx context.Context, vc *vim25.Client) {
		{
			// NewClient calls ServiceInstance methods
			_, err := ssoadmin.NewClient(ctx, vc)
			if err != nil {
				t.Fatal(err)
			}
		}

		{
			c, err := sts.NewClient(ctx, vc)
			if err != nil {
				t.Fatal(err)
			}

			req := sts.TokenRequest{
				Userinfo: url.UserPassword("Administrator@VSPHERE.LOCAL", "password"),
			}

			_, err = c.Issue(ctx, req)
			if err != nil {
				t.Fatal(err)
			}
		}
	})

	t.Run("With Envoy sidecar and a malfunctioning lookup service, endpoint url should still succeed", func(t *testing.T) {
		model := simulator.VPX()
		model.Create()
		simulator.Test(func(ctx context.Context, vc *vim25.Client) {
			lsim.BreakLookupServiceURLs(ctx)
			// Map Envoy sidecar on the same port as the vcsim client.
			os.Setenv("GOVMOMI_ENVOY_SIDECAR_PORT", vc.Client.URL().Port())
			os.Setenv("GOVMOMI_ENVOY_SIDECAR_HOST", vc.Client.URL().Hostname())
			testPath := "/fake/path"
			expectedUrl := fmt.Sprintf("http://%s%s", vc.Client.URL().Host, testPath)
			url := lookup.EndpointURL(ctx, vc, testPath, nil)
			require.Equal(t, expectedUrl, url)
		}, model)
	})
}
