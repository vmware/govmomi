// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator_test

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/vmware/govmomi/sts"
	stssim "github.com/vmware/govmomi/sts/simulator"
	"github.com/vmware/govmomi/vapi/authentication"
	"github.com/vmware/govmomi/vapi/rest"
	"github.com/vmware/govmomi/vim25"
)

// agent returns a holder-of-key token for the given user, and a vAPI client logged in with it,
// as supervisor-agent obtains them.
func agent(ctx context.Context, t *testing.T, c *vim25.Client, user string) (*sts.Signer, *rest.Client) {
	t.Helper()
	sc, err := sts.NewClient(ctx, c)
	require.NoError(t, err)
	signer, err := sc.Issue(ctx, sts.TokenRequest{
		Userinfo:    url.UserPassword(user, "password"),
		Certificate: certificate(t, "agent"),
		Lifetime:    10 * time.Minute,
		Delegatable: true,
		Renewable:   true,
	})
	require.NoError(t, err)

	rc := rest.NewClient(c)
	require.NoError(t, rc.LoginByToken(rc.WithSigner(ctx, signer)))
	return signer, rc
}

// exchangeSpec is the request supervisor-agent's TES client sends.
func exchangeSpec(token string) authentication.TokenIssueSpec {
	return authentication.TokenIssueSpec{
		Audience:           stssim.AudienceVNS,
		GrantType:          "urn:ietf:params:oauth:grant-type:token-exchange",
		RequestedTokenType: "urn:ietf:params:oauth:token-type:id_token",
		SubjectToken:       base64.StdEncoding.EncodeToString([]byte(token)),
		SubjectTokenType:   "urn:ietf:params:oauth:token-type:saml2",
	}
}

// exchange posts spec to TES with the session of rc, returning the status and body.
func exchange(ctx context.Context, t *testing.T, rc *rest.Client, spec authentication.TokenIssueSpec) (int, string) {
	t.Helper()
	body, err := json.Marshal(map[string]any{"spec": spec})
	require.NoError(t, err)
	u := rc.URL()
	u.Path = "/rest/vcenter/tokenservice/token-exchange"
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u.String(), bytes.NewReader(body))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	if id := rc.SessionID(); id != "" {
		req.Header.Set("vmware-api-session-id", id)
	}

	var status int
	var res []byte
	err = rc.Client.Do(ctx, req, func(r *http.Response) error {
		status = r.StatusCode
		res, err = io.ReadAll(r.Body)
		return err
	})
	require.NoError(t, err)
	return status, string(res)
}

// claims decodes the claims of a JWT, without verifying it.
func claims(t *testing.T, jwt string) map[string]any {
	t.Helper()
	parts := strings.Split(jwt, ".")
	require.Len(t, parts, 3)
	b, err := base64.RawURLEncoding.DecodeString(parts[1])
	require.NoError(t, err)
	var c map[string]any
	require.NoError(t, json.Unmarshal(b, &c))
	return c
}

func TestTokenExchange(t *testing.T) {
	test(t, func(ctx context.Context, c *vim25.Client, h *stssim.Handler) {
		signer, rc := agent(ctx, t, c, "wcp-1234@vsphere.local")

		t.Run("SAML token to JWT", func(t *testing.T) {
			info, err := authentication.NewManager(rc).Issue(ctx, exchangeSpec(signer.Token))
			require.NoError(t, err)
			assert.Equal(t, "urn:ietf:params:oauth:token-type:id_token", info.IssuedTokenType)
			assert.Contains(t, h.Minted(), info.AccessToken)

			c := claims(t, info.AccessToken)
			assert.Equal(t, h.Issuer(), c["iss"])
			assert.Equal(t, "https://"+rc.URL().Host+"/openidconnect/vsphere.local", c["iss"])
			assert.Equal(t, stssim.AudienceVNS, c["aud"])
			assert.Equal(t, "wcp-1234@vsphere.local", c["sub"])
			assert.Equal(t, "wcp-1234", c["username"])
			assert.Equal(t, "vsphere.local", c["domain"])
			assert.Equal(t, []any{"Users@vsphere.local", "Everyone@vsphere.local"}, c["group_names"])
			assert.Greater(t, c["exp"], c["iat"])
		})

		t.Run("the subject's domain is in lower case", func(t *testing.T) {
			admin, _ := agent(ctx, t, c, "Administrator@VSPHERE.LOCAL")
			info, err := authentication.NewManager(rc).Issue(ctx, exchangeSpec(admin.Token))
			require.NoError(t, err)
			c := claims(t, info.AccessToken)
			assert.Equal(t, "Administrator@vsphere.local", c["sub"])
			assert.Contains(t, c["group_names"], "Administrators@vsphere.local")
		})

		t.Run("session required", func(t *testing.T) {
			status, _ := exchange(ctx, t, rest.NewClient(c), exchangeSpec(signer.Token))
			assert.Equal(t, http.StatusUnauthorized, status)
		})

		t.Run("invalid subject token", func(t *testing.T) {
			forged := strings.Replace(signer.Token, ">wcp-1234@", ">root@", 1)
			status, body := exchange(ctx, t, rc, exchangeSpec(forged))
			assert.Equal(t, http.StatusInternalServerError, status)
			assert.Contains(t, body, stssim.InvalidGrant)
			assert.NotContains(t, body, "root@")
		})

		t.Run("invalid request", func(t *testing.T) {
			for name, f := range map[string]func(*authentication.TokenIssueSpec){
				"audience":             func(s *authentication.TokenIssueSpec) { s.Audience = "vmware-tes:vapi" },
				"no audience":          func(s *authentication.TokenIssueSpec) { s.Audience = "" },
				"grant type":           func(s *authentication.TokenIssueSpec) { s.GrantType = "password" },
				"subject token type":   func(s *authentication.TokenIssueSpec) { s.SubjectTokenType = "jwt" },
				"requested token type": func(s *authentication.TokenIssueSpec) { s.RequestedTokenType = "saml2" },
				"subject token":        func(s *authentication.TokenIssueSpec) { s.SubjectToken = "%" },
			} {
				spec := exchangeSpec(signer.Token)
				f(&spec)
				status, body := exchange(ctx, t, rc, spec)
				assert.Equal(t, http.StatusBadRequest, status, name)
				assert.Contains(t, body, "com.vmware.vapi.std.errors.invalid_request", name)
			}
		})

		t.Run("injected InvalidGrant", func(t *testing.T) {
			h.InjectInvalidGrant(1)
			status, body := exchange(ctx, t, rc, exchangeSpec(signer.Token))
			assert.Equal(t, http.StatusInternalServerError, status)
			assert.Contains(t, body, stssim.InvalidGrant)

			status, _ = exchange(ctx, t, rc, exchangeSpec(signer.Token))
			assert.Equal(t, http.StatusOK, status)
		})
	})
}
