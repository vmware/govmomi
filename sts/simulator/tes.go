// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1" // #nosec G505 -- a certificate thumbprint, not a security control
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/vmware/govmomi/vapi/authentication"
	"github.com/vmware/govmomi/vapi/rest"
)

// Token Exchange Service (TES) constants, per com.vmware.vcenter.tokenservice.
const (
	tesPath = rest.Path + "/vcenter/tokenservice/token-exchange"

	grantTokenExchange = "urn:ietf:params:oauth:grant-type:token-exchange" // #nosec G101 -- not a credential
	tokenTypeSAML2     = "urn:ietf:params:oauth:token-type:saml2"          // #nosec G101 -- not a credential
	tokenTypeIDToken   = "urn:ietf:params:oauth:token-type:id_token"       // #nosec G101 -- not a credential

	// AudienceVNS is the audience of a JWT for the Supervisor's Kubernetes API server.
	AudienceVNS = "vmware-tes:vc:vns:k8s"

	// InvalidGrant is the error class TES responds with when a token is invalid.
	// It is returned as a 500 status, with the class in the response body.
	InvalidGrant = "com.vmware.vcenter.tokenservice.exceptions.InvalidGrant"

	errInternal       = "com.vmware.vapi.std.errors.internal_server_error"
	errInvalidRequest = "com.vmware.vapi.std.errors.invalid_request"

	// vnsLifetime is the lifetime of a JWT for the vns audience.
	vnsLifetime = 10 * time.Hour
)

// InjectInvalidGrant makes the next n token exchange requests fail with InvalidGrant,
// as vCenter does when the token that authenticates the session is no longer valid.
func (s *Handler) InjectInvalidGrant(n int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.invalidGrant = n
}

func (s *Handler) takeInvalidGrant() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.invalidGrant > 0 {
		s.invalidGrant--
		return true
	}
	return false
}

// tesError responds with a "/rest" style vAPI error. The message never includes the request.
func tesError(w http.ResponseWriter, status int, kind, id, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	type value struct {
		Messages []rest.LocalizableMessage `json:"messages"`
	}

	_ = json.NewEncoder(w).Encode(struct {
		Type  string `json:"type"`
		Value value  `json:"value"`
	}{kind, value{[]rest.LocalizableMessage{{Args: []string{}, DefaultMessage: message, ID: id}}}})
}

func invalidRequest(w http.ResponseWriter, message string) {
	tesError(w, http.StatusBadRequest, errInvalidRequest, errInvalidRequest, message)
}

func invalidGrant(w http.ResponseWriter, message string) {
	tesError(w, http.StatusInternalServerError, errInternal, InvalidGrant, InvalidGrant+": "+message)
}

// tokenExchange implements TokenExchange.exchange for a SAML token to a JWT with the vns audience.
// The request is registered with the vapi simulator, which requires a session.
// Neither the request nor the response is logged, as both carry credentials.
func (s *Handler) tokenExchange(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Spec authentication.TokenIssueSpec `json:"spec"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		invalidRequest(w, "Invalid request body")
		return
	}
	spec := &req.Spec

	if s.takeInvalidGrant() {
		invalidGrant(w, "Invalid CALLER token: tokenType=SAML2")
		return
	}

	switch {
	case spec.GrantType != grantTokenExchange:
		invalidRequest(w, "Unsupported grant type")
		return
	case spec.SubjectTokenType != tokenTypeSAML2:
		invalidRequest(w, "Unsupported subject token type")
		return
	case spec.RequestedTokenType != tokenTypeIDToken:
		invalidRequest(w, "Unsupported requested token type")
		return
	case strings.TrimSpace(spec.Audience) == "":
		invalidRequest(w, "Audience is required to request a JWT.")
		return
	case !strings.EqualFold(spec.Audience, AudienceVNS):
		invalidRequest(w, "Unknown audience value")
		return
	}

	assertion, err := base64.StdEncoding.DecodeString(spec.SubjectToken)
	if err != nil {
		invalidRequest(w, "Failed to parse SUBJECT token: tokenType=SAML2")
		return
	}
	t, err := s.verify(string(assertion))
	if err != nil {
		invalidGrant(w, "Invalid SUBJECT token: tokenType=SAML2")
		return
	}

	jwt, err := s.jwt(t, strings.ToLower(spec.Audience))
	if err != nil {
		tesError(w, http.StatusInternalServerError, errInternal, errInternal, "Failed to issue token")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(struct {
		Value authentication.TokenInfo `json:"value"`
	}{authentication.TokenInfo{
		AccessToken:     jwt,
		TokenType:       "N_A",
		ExpiresIn:       int(vnsLifetime.Seconds()),
		IssuedTokenType: tokenTypeIDToken,
		Scope:           "openid",
	}})
}

// upn returns name@domain with the domain in lower case, as TES writes principals.
func upn(name, domain string) string {
	return name + "@" + strings.ToLower(domain)
}

// Issuer returns the OpenID Connect issuer of the JWTs the simulator mints: https://<host>/openidconnect/<domain>.
func (s *Handler) Issuer() string {
	return s.URL.JoinPath("openidconnect", s.Domain).String()
}

// keyID returns the id of the signing key: the hex SHA-1 thumbprint of its certificate.
func keyID(cert *x509.Certificate) string {
	sum := sha1.Sum(cert.Raw) // #nosec G401 -- a certificate thumbprint, not a security control
	return strings.ToUpper(hex.EncodeToString(sum[:]))
}

// jwt returns an RS256 JWT for the subject of t, with the claims TES writes for the vns audience.
func (s *Handler) jwt(t *token, audience string) (string, error) {
	key, cert, err := s.signer()
	if err != nil {
		return "", err
	}

	name, domain, _ := strings.Cut(t.Subject, "@")
	groups := []string{}
	for _, g := range t.Groups {
		if d, n, ok := strings.Cut(g, `\`); ok {
			groups = append(groups, upn(n, d))
		}
	}

	now := time.Now()
	claims := map[string]any{
		"iss":         s.Issuer(),
		"sub":         upn(name, domain),
		"username":    name,
		"domain":      strings.ToLower(domain),
		"aud":         audience,
		"exp":         now.Add(vnsLifetime).Unix(),
		"iat":         now.Unix(),
		"group_names": groups,
		"jti":         uuid.NewString(),
	}
	header := map[string]string{"alg": "RS256", "kid": keyID(cert)}

	enc := func(v any) (string, error) {
		b, err := json.Marshal(v)
		return base64.RawURLEncoding.EncodeToString(b), err
	}
	h, err := enc(header)
	if err != nil {
		return "", err
	}
	c, err := enc(claims)
	if err != nil {
		return "", err
	}

	input := h + "." + c
	sum := sha256.Sum256([]byte(input))
	sig, err := rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA256, sum[:])
	if err != nil {
		return "", err
	}
	jwt := input + "." + base64.RawURLEncoding.EncodeToString(sig)
	s.mint(jwt)
	return jwt, nil
}
