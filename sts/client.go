/*
Copyright (c) 2018 VMware, Inc. All Rights Reserved.

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

package sts

import (
	"context"
	"crypto/tls"
	"errors"
	"net/url"
	"time"

	"github.com/vmware/govmomi/sts/internal"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/soap"
)

// Client is a soap.Client targeting the STS (Secure Token Service) API endpoint.
type Client struct {
	*soap.Client
}

// NewClient returns a client targeting the STS API endpoint.
func NewClient(ctx context.Context, c *vim25.Client) (*Client, error) {
	sc := c.Client.NewServiceClient("/sts/STSService", "urn:oasis:names:tc:SAML:2.0:assertion")

	return &Client{sc}, nil
}

// TokenRequest parameters for issuing a SAML token.
// At least one of Userinfo or Certificate must be specified.
type TokenRequest struct {
	Userinfo    *url.Userinfo    // Userinfo when set issues a Bearer token
	Certificate *tls.Certificate // Certificate when set issues a HoK token
	ActAs       string           // ActAs identity when set is delegated to the caller
	Lifetime    time.Duration    // Lifetime is the token's lifetime, defaults to 10m
	Renewable   bool             // Renewable allows the issued token to be renewed
	Delegatable bool             // Delegatable allows the issued token to be delegated (e.g. for use with ActAs)
}

// Issue is used to request a security token.
// The returned Signer can be used to sign SOAP requests, such as the SessionManager LoginByToken method and the RequestSecurityToken method itself.
// One of TokenRequest Certificate or Userinfo is required, with Certificate taking precedence.
// When Certificate is set, a Holder-of-Key token will be requested.  Otherwise, a Bearer token is requested with the Userinfo credentials.
// See: http://docs.oasis-open.org/ws-sx/ws-trust/v1.4/errata01/os/ws-trust-1.4-errata01-os-complete.html#_Toc325658937
func (c *Client) Issue(ctx context.Context, req TokenRequest) (*Signer, error) {
	if req.Certificate == nil && req.Userinfo == nil {
		return nil, errors.New("one of TokenRequest Certificate or Userinfo is required")
	}

	s := &Signer{
		Certificate: req.Certificate,
		user:        req.Userinfo,
	}

	if req.Lifetime == 0 {
		req.Lifetime = time.Minute * 10
	}

	created := time.Now().UTC()
	rst := internal.RequestSecurityToken{
		TokenType:          c.Namespace,
		RequestType:        "http://docs.oasis-open.org/ws-sx/ws-trust/200512/Issue",
		SignatureAlgorithm: internal.SHA256,
		Lifetime: &internal.Lifetime{
			Created: internal.Time{Time: created},
			Expires: internal.Time{Time: created.Add(req.Lifetime)},
		},
		Renewing: &internal.Renewing{
			Allow: req.Renewable,
			// /wst:RequestSecurityToken/wst:Renewing/@OK
			// "It NOT RECOMMENDED to use this as it can leave you open to certain types of security attacks.
			// Issuers MAY restrict the period after expiration during which time the token can be renewed.
			// This window is governed by the issuer's policy."
			OK: false,
		},
		Delegatable: req.Delegatable,
	}

	if req.ActAs != "" {
		rst.ActAs = &internal.ActAs{
			Assertion: req.ActAs,
		}
	}

	if s.Certificate == nil {
		rst.KeyType = "http://docs.oasis-open.org/ws-sx/ws-trust/200512/Bearer"
	} else {
		rst.KeyType = "http://docs.oasis-open.org/ws-sx/ws-trust/200512/PublicKey"

		s.keyID = newID()
		rst.UseKey = &internal.UseKey{Sig: s.keyID}
	}

	header := soap.Header{
		Security: s,
		Action:   rst.Action(),
	}

	res, err := internal.Issue(c.WithHeader(ctx, header), c, &rst)
	if err != nil {
		return nil, err
	}

	// TODO: consider checking res.RequestSecurityTokenResponse.Lifetime
	// /wst:RequestSecurityToken/wst:Lifetime
	// "The issuer is not obligated to honor this range – they MAY return a more (or less) restrictive interval.
	// It is RECOMMENDED that the issuer return this element with issued tokens (in the RSTR) so the requestor
	// knows the actual validity period without having to parse the returned token."

	s.Token = res.RequestSecurityTokenResponse.RequestedSecurityToken.Assertion

	return s, nil
}
