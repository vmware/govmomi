// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator_test

import (
	"bytes"
	"context"
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"net/url"
	"os"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/simulator"
	stssim "github.com/vmware/govmomi/sts/simulator"
	"github.com/vmware/govmomi/vapi/authentication"
)

// verifier checks a JWT as a Kubernetes API server configured with a JWT authenticator does:
// discovery and JWKS are fetched from the issuer over TLS, trusting only the given CA.
type verifier struct {
	issuer   string
	audience string
	client   *http.Client
}

func (v *verifier) get(ctx context.Context, u string, val any) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return err
	}
	res, err := v.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("GET %s: %s", u, res.Status)
	}
	return json.NewDecoder(res.Body).Decode(val)
}

// verify returns the claims of jwt if it is valid.
func (v *verifier) verify(ctx context.Context, jwt string) (map[string]any, error) {
	var discovery struct {
		Issuer  string   `json:"issuer"`
		JWKS    string   `json:"jwks_uri"`
		Signing []string `json:"id_token_signing_alg_values_supported"`
	}
	if err := v.get(ctx, v.issuer+"/.well-known/openid-configuration", &discovery); err != nil {
		return nil, err
	}
	if discovery.Issuer != v.issuer {
		return nil, fmt.Errorf("discovery issuer %q, want %q", discovery.Issuer, v.issuer)
	}
	if !slices.Contains(discovery.Signing, "RS256") {
		return nil, errors.New("RS256 not supported")
	}

	var jwks struct {
		Keys []struct {
			Kty string `json:"kty"`
			Kid string `json:"kid"`
			N   string `json:"n"`
			E   string `json:"e"`
		} `json:"keys"`
	}
	if err := v.get(ctx, discovery.JWKS, &jwks); err != nil {
		return nil, err
	}

	parts := strings.Split(jwt, ".")
	if len(parts) != 3 {
		return nil, errors.New("malformed JWT")
	}
	var header struct {
		Alg string `json:"alg"`
		Kid string `json:"kid"`
	}
	b, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(b, &header); err != nil {
		return nil, err
	}
	if header.Alg != "RS256" {
		return nil, fmt.Errorf("alg %q", header.Alg)
	}

	var key *rsa.PublicKey
	for _, k := range jwks.Keys {
		if k.Kty == "RSA" && k.Kid == header.Kid {
			n, err := base64.RawURLEncoding.DecodeString(k.N)
			if err != nil {
				return nil, err
			}
			e, err := base64.RawURLEncoding.DecodeString(k.E)
			if err != nil {
				return nil, err
			}
			key = &rsa.PublicKey{N: new(big.Int).SetBytes(n), E: int(new(big.Int).SetBytes(e).Int64())}
		}
	}
	if key == nil {
		return nil, fmt.Errorf("no key %q in JWKS", header.Kid)
	}

	sig, err := base64.RawURLEncoding.DecodeString(parts[2])
	if err != nil {
		return nil, err
	}
	sum := sha256.Sum256([]byte(parts[0] + "." + parts[1]))
	if err = rsa.VerifyPKCS1v15(key, crypto.SHA256, sum[:], sig); err != nil {
		return nil, err
	}

	var claims map[string]any
	if b, err = base64.RawURLEncoding.DecodeString(parts[1]); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(b, &claims); err != nil {
		return nil, err
	}
	if claims["iss"] != v.issuer {
		return nil, fmt.Errorf("iss %q", claims["iss"])
	}
	if claims["aud"] != v.audience {
		return nil, fmt.Errorf("aud %q", claims["aud"])
	}
	if exp, ok := claims["exp"].(float64); !ok || time.Unix(int64(exp), 0).Before(time.Now()) {
		return nil, errors.New("expired")
	}
	return claims, nil
}

// TestTokenChain follows supervisor-agent's credential from its SAML token, through the JWT that TES issues for
// it, to the Supervisor's API server verifying that JWT against the discovery and JWKS vcsim serves.
func TestTokenChain(t *testing.T) {
	ctx := context.Background()

	// The simulator logs with the standard logger; none of its credentials may appear there.
	var logs bytes.Buffer
	log.SetOutput(&logs)
	defer log.SetOutput(os.Stderr)

	model := simulator.VPX()
	defer model.Remove()
	require.NoError(t, model.Create())
	model.Service.TLS = new(tls.Config)
	model.Service.RegisterEndpoints = true
	srv := model.Service.NewServer()
	defer srv.Close()

	c, err := govmomi.NewClient(ctx, srv.URL, true)
	require.NoError(t, err)
	h := stssim.Lookup(model.Service)
	require.NotNil(t, h)

	// The agent's HoK token, and a vAPI session for it.
	signer, rc := agent(ctx, t, c.Client, "wcp-1234@vsphere.local")

	// The agent exchanges its token for a JWT.
	info, err := authentication.NewManager(rc).Issue(ctx, exchangeSpec(signer.Token))
	require.NoError(t, err)

	// The API server's issuer, as supervisor's bootstrapper derives it from the vCenter endpoint and SSO domain,
	// trusting the vcsim certificate as its CA bundle.
	endpoint := &url.URL{Scheme: "https", Host: srv.URL.Host}
	roots := x509.NewCertPool()
	roots.AddCert(srv.Certificate())
	v := &verifier{
		issuer:   endpoint.JoinPath("openidconnect", "vsphere.local").String(),
		audience: stssim.AudienceVNS,
		client:   &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{RootCAs: roots}}},
	}

	claims, err := v.verify(ctx, info.AccessToken)
	require.NoError(t, err)
	// The API server maps the claims to user "sso:wcp-1234@vsphere.local", the name bootstrapper binds.
	assert.Equal(t, "wcp-1234@vsphere.local", claims["sub"])
	assert.Contains(t, claims["group_names"], "Everyone@vsphere.local")

	t.Run("altered JWT", func(t *testing.T) {
		parts := strings.Split(info.AccessToken, ".")
		b, err := json.Marshal(map[string]any{"iss": v.issuer, "aud": v.audience, "sub": "root@vsphere.local", "exp": time.Now().Add(time.Hour).Unix()})
		require.NoError(t, err)
		forged := parts[0] + "." + base64.RawURLEncoding.EncodeToString(b) + "." + parts[2]
		_, err = v.verify(ctx, forged)
		assert.Error(t, err)
	})

	t.Run("canaries", func(t *testing.T) {
		minted := h.Minted()
		assert.Contains(t, minted, signer.Token)
		assert.Contains(t, minted, info.AccessToken)

		// A failed exchange is logged by neither the simulator nor, via its response, the client.
		status, body := exchange(ctx, t, rc, exchangeSpec(strings.Replace(signer.Token, "wcp-1234", "root", 1)))
		assert.Equal(t, http.StatusInternalServerError, status)
		minted = append(minted, base64.StdEncoding.EncodeToString([]byte(signer.Token)))
		for _, canary := range minted {
			assert.NotContains(t, logs.String(), canary)
			assert.NotContains(t, body, canary)
		}
	})

	t.Run("JWKS for the default domain", func(t *testing.T) {
		var jwks map[string]any
		require.NoError(t, v.get(ctx, endpoint.JoinPath("openidconnect", "jwks").String(), &jwks))
		assert.Len(t, jwks["keys"], 1)

		err := v.get(ctx, endpoint.JoinPath("openidconnect", "example.com", ".well-known", "openid-configuration").String(), &jwks)
		assert.ErrorContains(t, err, "404")
	})
}
