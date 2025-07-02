// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package sts_test

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"encoding/pem"
	"log"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	lsim "github.com/vmware/govmomi/lookup/simulator"
	"github.com/vmware/govmomi/session"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/ssoadmin"
	_ "github.com/vmware/govmomi/ssoadmin/simulator"
	"github.com/vmware/govmomi/ssoadmin/types"
	"github.com/vmware/govmomi/sts"
	_ "github.com/vmware/govmomi/sts/simulator"
	"github.com/vmware/govmomi/vapi/rest"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/soap"
)

// The following can help debug signature mismatch:
// % vi /usr/lib/vmware-sso/vmware-sts/conf/logging.properties
// # turn up logging for dsig:
// org.jcp.xml.dsig.internal.level = FINE
// com.sun.org.apache.xml.internal.security.level = FINE
// # restart the STS service:
// % service-control --stop vmware-stsd
// % service-control --start vmware-stsd
// % tail -f /var/log/vmware/sso/vmware-rest-idm-http.log

// solutionUserCreate ensures that solution user "govmomi-test" exists for uses with the tests that follow.
func solutionUserCreate(ctx context.Context, info *url.Userinfo, stsClient *sts.Client, vc *vim25.Client) error {
	s, err := stsClient.Issue(ctx, sts.TokenRequest{Userinfo: info})
	if err != nil {
		return err
	}

	admin, err := ssoadmin.NewClient(ctx, vc)
	if err != nil {
		return err
	}

	header := soap.Header{Security: s}
	if err = admin.Login(stsClient.WithHeader(ctx, header)); err != nil {
		return err
	}

	defer func() {
		if err := admin.Logout(ctx); err != nil {
			log.Printf("user logout error: %v", err)
		}
	}()

	id := types.PrincipalId{
		Name:   "govmomi-test",
		Domain: admin.Domain,
	}

	user, err := admin.FindSolutionUser(ctx, id.Name)
	if err != nil {
		return err
	}

	if user == nil {
		block, _ := pem.Decode([]byte(sts.LocalhostCert))
		details := types.AdminSolutionDetails{
			Certificate: base64.StdEncoding.EncodeToString(block.Bytes),
			Description: "govmomi test solution user",
		}

		if err = admin.CreateSolutionUser(ctx, id.Name, details); err != nil {
			return err
		}
	}

	if _, err = admin.GrantWSTrustRole(ctx, id, types.RoleActAsUser); err != nil {
		return err
	}

	_, err = admin.SetRole(ctx, id, types.RoleAdministrator)
	return err
}

func solutionUserCert() *tls.Certificate {
	cert, err := tls.X509KeyPair(sts.LocalhostCert, sts.LocalhostKey)
	if err != nil {
		panic(err)
	}
	return &cert
}

func TestIssueHOK(t *testing.T) {
	ctx := context.Background()
	url := os.Getenv("GOVC_TEST_URL")
	if url == "" {
		t.SkipNow()
	}

	u, err := soap.ParseURL(url)
	if err != nil {
		t.Fatal(err)
	}

	c, err := vim25.NewClient(ctx, soap.NewClient(u, true))
	if err != nil {
		log.Fatal(err)
	}
	_ = c.UseServiceVersion()

	if !c.IsVC() {
		t.SkipNow()
	}

	stsClient, err := sts.NewClient(ctx, c)
	if err != nil {
		t.Fatal(err)
	}

	if err = solutionUserCreate(ctx, u.User, stsClient, c); err != nil {
		t.Fatal(err)
	}

	req := sts.TokenRequest{
		Certificate: solutionUserCert(),
		Delegatable: true,
	}

	s, err := stsClient.Issue(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	header := soap.Header{Security: s}

	err = session.NewManager(c).LoginByToken(c.WithHeader(ctx, header))
	if err != nil {
		t.Fatal(err)
	}

	now, err := methods.GetCurrentTime(ctx, c)
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("current time=%s", now)

	rc := rest.NewClient(c)
	err = rc.LoginByToken(rc.WithSigner(ctx, s))
	if err != nil {
		t.Fatal(err)
	}
}

func TestIssueTokenByToken(t *testing.T) {
	ctx := context.Background()
	url := os.Getenv("GOVC_TEST_URL")
	if url == "" {
		t.SkipNow()
	}

	u, err := soap.ParseURL(url)
	if err != nil {
		t.Fatal(err)
	}

	vc1, err := vim25.NewClient(ctx, soap.NewClient(u, true))
	if err != nil {
		log.Fatal(err)
	}
	_ = vc1.UseServiceVersion()

	if !vc1.IsVC() {
		t.SkipNow()
	}

	vc2, err := vim25.NewClient(ctx, soap.NewClient(u, true))
	if err != nil {
		log.Fatal(err)
	}
	_ = vc2.UseServiceVersion()

	sts1, err := sts.NewClient(ctx, vc1)
	if err != nil {
		t.Fatal(err)
	}

	if err = solutionUserCreate(ctx, u.User, sts1, vc1); err != nil {
		t.Fatal(err)
	}

	sts2, err := sts.NewClient(ctx, vc2)
	if err != nil {
		t.Fatal(err)
	}

	req1 := sts.TokenRequest{
		Certificate: solutionUserCert(),
		Delegatable: true,
	}

	signer1, err := sts1.Issue(ctx, req1)
	if err != nil {
		t.Fatal(err)
	}

	req2 := signer1.NewRequest()
	// use Assertion header instead of BinarySecurityToken
	req2.Certificate = &tls.Certificate{PrivateKey: req1.Certificate.PrivateKey}

	signer2, err := sts2.Issue(ctx, req2)
	if err != nil {
		t.Fatal(err)
	}

	header := soap.Header{Security: signer2}

	err = session.NewManager(vc2).LoginByToken(vc2.WithHeader(ctx, header))
	if err != nil {
		t.Fatal(err)
	}
}

func TestIssueBearer(t *testing.T) {
	ctx := context.Background()
	url := os.Getenv("GOVC_TEST_URL")
	if url == "" {
		t.SkipNow()
	}

	u, err := soap.ParseURL(url)
	if err != nil {
		t.Fatal(err)
	}

	c, err := vim25.NewClient(ctx, soap.NewClient(u, true))
	if err != nil {
		log.Fatal(err)
	}
	_ = c.UseServiceVersion()

	if !c.IsVC() {
		t.SkipNow()
	}

	stsClient, err := sts.NewClient(ctx, c)
	if err != nil {
		t.Fatal(err)
	}

	// Test that either Certificate or Userinfo is set.
	_, err = stsClient.Issue(ctx, sts.TokenRequest{})
	if err == nil {
		t.Error("expected error")
	}

	req := sts.TokenRequest{
		Userinfo: u.User,
	}

	s, err := stsClient.Issue(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	header := soap.Header{Security: s}

	err = session.NewManager(c).LoginByToken(c.WithHeader(ctx, header))
	if err != nil {
		t.Fatal(err)
	}

	now, err := methods.GetCurrentTime(ctx, c)
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("current time=%s", now)
}

func TestIssueActAs(t *testing.T) {
	ctx := context.Background()
	url := os.Getenv("GOVC_TEST_URL")
	if url == "" {
		t.SkipNow()
	}

	u, err := soap.ParseURL(url)
	if err != nil {
		t.Fatal(err)
	}

	c, err := vim25.NewClient(ctx, soap.NewClient(u, true))
	if err != nil {
		log.Fatal(err)
	}
	_ = c.UseServiceVersion()

	if !c.IsVC() {
		t.SkipNow()
	}

	stsClient, err := sts.NewClient(ctx, c)
	if err != nil {
		t.Fatal(err)
	}

	if err = solutionUserCreate(ctx, u.User, stsClient, c); err != nil {
		t.Fatal(err)
	}

	req := sts.TokenRequest{
		Delegatable: true,
		Renewable:   true,
		Userinfo:    u.User,
	}

	s, err := stsClient.Issue(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	req = sts.TokenRequest{
		Lifetime:    24 * time.Hour,
		Token:       s.Token,
		ActAs:       true,
		Delegatable: true,
		Renewable:   true,
		Certificate: solutionUserCert(),
	}

	s, err = stsClient.Issue(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	header := soap.Header{Security: s}

	err = session.NewManager(c).LoginByToken(c.WithHeader(ctx, header))
	if err != nil {
		t.Fatal(err)
	}

	now, err := methods.GetCurrentTime(ctx, c)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("current time=%s", now)

	duration := s.Lifetime.Expires.Sub(s.Lifetime.Created)
	if duration < req.Lifetime {
		req.Lifetime = 24 * time.Hour
		req.Token = s.Token
		log.Printf("extending lifetime from %s", s.Lifetime.Expires.Sub(s.Lifetime.Created))
		s, err = stsClient.Renew(ctx, req)
		if err != nil {
			t.Fatal(err)
		}
	} else {
		t.Errorf("duration=%s", duration)
	}

	t.Logf("expires in %s", s.Lifetime.Expires.Sub(s.Lifetime.Created))
}

func TestNewClient(t *testing.T) {
	t.Run("Happy path client creation", func(t *testing.T) {
		simulator.Test(func(ctx context.Context, client *vim25.Client) {
			_, err := sts.NewClient(ctx, client)
			require.NoError(t, err)
		})
	})

	t.Run("STS client should work with Envoy sidecar even when lookup service is down", func(t *testing.T) {
		model := simulator.VPX()

		model.Create()
		simulator.Test(func(ctx context.Context, client *vim25.Client) {
			lsim.BreakLookupServiceURLs(ctx)
			// Map Envoy sidecar on the same port as the vcsim client.
			os.Setenv("GOVMOMI_ENVOY_SIDECAR_PORT", client.Client.URL().Port())
			os.Setenv("GOVMOMI_ENVOY_SIDECAR_HOST", client.Client.URL().Hostname())

			c, err := sts.NewClient(ctx, client)
			require.NoError(t, err)

			req := sts.TokenRequest{
				Userinfo: url.UserPassword("foo", "bar"),
			}

			_, err = c.Issue(ctx, req)
			require.NoError(t, err)
		}, model)
	})
}
