// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"encoding/base64"
	"encoding/json"
	"math/big"
	"net/http"
	"strings"
)

// OpenID Connect endpoints, per the vCenter SSO server.
const (
	oidcPath     = "/openidconnect"
	oidcMetadata = "/.well-known/openid-configuration"
	oidcJWKS     = "/jwks"
)

// endpoint returns the URL of an OpenID Connect endpoint for the SSO domain: https://<host>/openidconnect/<name>/<domain>.
func (s *Handler) endpoint(name string) string {
	return s.URL.JoinPath("openidconnect", name, s.Domain).String()
}

// openIDConnect serves OpenID Connect discovery and the JWKS for the SSO domain, so that a JWT issued by the
// Token Exchange Service can be verified, by a Kubernetes API server for example. Only these two endpoints are
// served, and neither requires a session.
func (s *Handler) openIDConnect(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	switch strings.TrimPrefix(r.URL.Path, oidcPath) {
	case oidcMetadata, "/" + s.Domain + oidcMetadata:
		s.metadata(w)
	case oidcJWKS, oidcJWKS + "/" + s.Domain:
		s.jwks(w)
	default:
		http.NotFound(w, r)
	}
}

func (s *Handler) metadata(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{
		"issuer":                                s.Issuer(),
		"authorization_endpoint":                s.endpoint("oidc/authorize"),
		"token_endpoint":                        s.endpoint("token"),
		"end_session_endpoint":                  s.endpoint("logout"),
		"jwks_uri":                              s.endpoint("jwks"),
		"response_types_supported":              []string{"code", "id_token", "token id_token"},
		"subject_types_supported":               []string{"public"},
		"id_token_signing_alg_values_supported": []string{"RS256"},
	})
}

func (s *Handler) jwks(w http.ResponseWriter) {
	key, cert, err := s.signer()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	enc := base64.RawURLEncoding.EncodeToString
	pub := &key.PublicKey

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{
		"keys": []map[string]any{{
			"kty": "RSA",
			"use": "sig",
			"alg": "RS256",
			"kid": keyID(cert),
			"n":   enc(pub.N.Bytes()),
			"e":   enc(big.NewInt(int64(pub.E)).Bytes()),
			"x5c": []string{base64.StdEncoding.EncodeToString(cert.Raw)},
		}},
	})
}
