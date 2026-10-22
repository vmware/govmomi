// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator_test

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"math/big"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	_ "github.com/vmware/govmomi/lookup/simulator"
	"github.com/vmware/govmomi/session"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/sts"
	stssim "github.com/vmware/govmomi/sts/simulator"
	"github.com/vmware/govmomi/vapi/rest"
	_ "github.com/vmware/govmomi/vapi/simulator"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/soap"
)

// certificate returns a self-signed certificate and key with the given common name.
func certificate(t *testing.T, cn string) *tls.Certificate {
	t.Helper()
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)
	tmpl := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject:      pkix.Name{CommonName: cn},
		NotBefore:    time.Now().Add(-time.Hour),
		NotAfter:     time.Now().Add(time.Hour),
	}
	der, err := x509.CreateCertificate(rand.Reader, &tmpl, &tmpl, &key.PublicKey, key)
	require.NoError(t, err)
	return &tls.Certificate{Certificate: [][]byte{der}, PrivateKey: key}
}

// test runs f with a VPX simulator and its STS simulator.
func test(t *testing.T, f func(context.Context, *vim25.Client, *stssim.Handler)) {
	t.Helper()
	model := simulator.VPX()
	defer model.Remove()
	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		h := stssim.Lookup(model.Service)
		require.NotNil(t, h)
		f(ctx, c, h)
	}, model)
}

// restUser logs in to the vAPI endpoint with signer and returns the session user.
func restUser(ctx context.Context, c *vim25.Client, signer rest.Signer) (string, error) {
	rc := rest.NewClient(c)
	if err := rc.LoginByToken(rc.WithSigner(ctx, signer)); err != nil {
		return "", err
	}
	s, err := rc.Session(ctx)
	if err != nil {
		return "", err
	}
	return s.User, nil
}

// soapUser logs in to the vim25 endpoint with signer and returns the session user.
func soapUser(ctx context.Context, c *vim25.Client, signer *sts.Signer) (string, error) {
	vc, err := vim25.NewClient(ctx, soap.NewClient(c.URL(), true))
	if err != nil {
		return "", err
	}
	m := session.NewManager(vc)
	if err = m.LoginByToken(vc.WithHeader(ctx, soap.Header{Security: signer})); err != nil {
		return "", err
	}
	s, err := m.UserSession(ctx)
	if err != nil {
		return "", err
	}
	return s.UserName, nil
}

func TestIssue(t *testing.T) {
	test(t, func(ctx context.Context, c *vim25.Client, h *stssim.Handler) {
		sc, err := sts.NewClient(ctx, c)
		require.NoError(t, err)

		t.Run("bearer token for a user", func(t *testing.T) {
			s, err := sc.Issue(ctx, sts.TokenRequest{
				Userinfo: url.UserPassword("user", "pass"),
				Lifetime: 10 * time.Minute,
			})
			require.NoError(t, err)
			assert.Contains(t, s.Token, ">user@vsphere.local</saml2:NameID>")
			assert.Contains(t, s.Token, "urn:oasis:names:tc:SAML:2.0:cm:bearer")
			assert.Equal(t, 10*time.Minute, s.Lifetime.Expires.Sub(s.Lifetime.Created))
			assert.Contains(t, h.Minted(), s.Token)

			user, err := restUser(ctx, c, s)
			require.NoError(t, err)
			assert.Equal(t, "user@vsphere.local", user)

			user, err = soapUser(ctx, c, s)
			require.NoError(t, err)
			assert.Equal(t, "user@vsphere.local", user)
		})

		t.Run("holder-of-key token for a user", func(t *testing.T) {
			s, err := sc.Issue(ctx, sts.TokenRequest{
				Userinfo:    url.UserPassword("wcp-1234@vsphere.local", "pass"),
				Certificate: certificate(t, "ignored"),
			})
			require.NoError(t, err)
			assert.Contains(t, s.Token, ">wcp-1234@vsphere.local</saml2:NameID>")
			assert.Contains(t, s.Token, "urn:oasis:names:tc:SAML:2.0:cm:holder-of-key")

			user, err := restUser(ctx, c, s)
			require.NoError(t, err)
			assert.Equal(t, "wcp-1234@vsphere.local", user)

			// The request must be signed with the key the token confirms.
			other := &sts.Signer{Token: s.Token, Certificate: certificate(t, "other")}
			_, err = restUser(ctx, c, other)
			assert.ErrorContains(t, err, "401")

			// A holder-of-key token is not accepted without a request signature.
			bearer := &sts.Signer{Token: s.Token}
			_, err = restUser(ctx, c, bearer)
			assert.ErrorContains(t, err, "401")
		})

		t.Run("holder-of-key token for a solution user", func(t *testing.T) {
			s, err := sc.Issue(ctx, sts.TokenRequest{Certificate: certificate(t, "solution")})
			require.NoError(t, err)
			assert.Contains(t, s.Token, ">solution@vsphere.local</saml2:NameID>")
			assert.Contains(t, s.Token, `vsphere.local\SolutionUsers`)

			renewed, err := sc.Renew(ctx, sts.TokenRequest{Certificate: s.Certificate, Token: s.Token})
			require.NoError(t, err)
			assert.NotEqual(t, s.Token, renewed.Token)
			assert.Contains(t, renewed.Token, ">solution@vsphere.local</saml2:NameID>")

			user, err := soapUser(ctx, c, renewed)
			require.NoError(t, err)
			assert.Equal(t, "solution@vsphere.local", user)
		})

		t.Run("token not issued by the simulator", func(t *testing.T) {
			s, err := sc.Issue(ctx, sts.TokenRequest{Userinfo: url.UserPassword("user", "pass")})
			require.NoError(t, err)

			for name, token := range map[string]string{
				"altered subject": strings.Replace(s.Token, ">user@", ">root@", 1),
				"no signature":    s.Token[:strings.Index(s.Token, "<ds:Signature")] + s.Token[strings.Index(s.Token, "</ds:Signature>")+len("</ds:Signature>"):],
				"not a token":     "<saml2:Assertion/>",
			} {
				forged := &sts.Signer{Token: token}

				_, err = restUser(ctx, c, forged)
				assert.ErrorContains(t, err, "401", name)

				_, err = soapUser(ctx, c, forged)
				assert.Error(t, err, name)

				_, err = sc.Renew(ctx, sts.TokenRequest{Token: token, Userinfo: url.UserPassword("user", "pass")})
				assert.Error(t, err, name)
			}
		})
	})
}
