// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package sts

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net/url"
	"time"

	internalhelpers "github.com/vmware/govmomi/internal"
	"github.com/vmware/govmomi/lookup"
	"github.com/vmware/govmomi/lookup/types"
	"github.com/vmware/govmomi/sts/internal"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/soap"
	vim25types "github.com/vmware/govmomi/vim25/types"
)

const (
	Namespace  = "oasis:names:tc:SAML:2.0:assertion"
	basePath   = "/sts"
	Path       = basePath + "/STSService"
	SystemPath = basePath + "/system-STSService/sdk"
)

// Client is a soap.Client targeting the STS (Secure Token Service) API endpoint.
type Client struct {
	*soap.Client

	RoundTripper soap.RoundTripper
}

func getEndpointURL(ctx context.Context, c *vim25.Client) string {
	// Services running on vCenter can bypass lookup service using the
	// system-STSService path. This avoids the need to lookup the system domain.
	if usingSidecar := internalhelpers.UsingEnvoySidecar(c); usingSidecar {
		return fmt.Sprintf("http://%s%s", c.URL().Host, SystemPath)
	}
	return getEndpointURLFromLookupService(ctx, c)
}

func getEndpointURLFromLookupService(ctx context.Context, c *vim25.Client) string {
	filter := &types.LookupServiceRegistrationFilter{
		ServiceType: &types.LookupServiceRegistrationServiceType{
			Product: "com.vmware.cis",
			Type:    "cs.identity",
		},
		EndpointType: &types.LookupServiceRegistrationEndpointType{
			Protocol: "wsTrust",
			Type:     "com.vmware.cis.cs.identity.sso",
		},
	}

	return lookup.EndpointURL(ctx, c, Path, filter)
}

// NewClient returns a client targeting the STS API endpoint.
// The Client.URL will be set to that of the Lookup Service's endpoint registration,
// as the SSO endpoint can be external to vCenter.  If the Lookup Service is not available,
// URL defaults to Path on the vim25.Client.URL.Host.
func NewClient(ctx context.Context, c *vim25.Client) (*Client, error) {

	url := getEndpointURL(ctx, c)
	sc := c.Client.NewServiceClient(url, Namespace)

	return &Client{sc, sc}, nil
}

// RoundTrip dispatches to the RoundTripper field.
func (c *Client) RoundTrip(ctx context.Context, req, res soap.HasFault) error {
	// Drop any operationID header, not used by STS
	ctx = context.WithValue(ctx, vim25types.ID{}, "")
	return c.RoundTripper.RoundTrip(ctx, req, res)
}

// TokenRequest parameters for issuing a SAML token.
// At least one of Userinfo or Certificate must be specified.
// When `TokenRequest.Certificate` is set, the `tls.Certificate.PrivateKey` field must be set as it is required to sign the request.
// When the `tls.Certificate.Certificate` field is not set, the request Assertion header is set to that of the TokenRequest.Token.
// Otherwise `tls.Certificate.Certificate` is used as the BinarySecurityToken in the request.
type TokenRequest struct {
	Userinfo    *url.Userinfo    // Userinfo when set issues a Bearer token
	Certificate *tls.Certificate // Certificate when set issues a HoK token
	Lifetime    time.Duration    // Lifetime is the token's lifetime, defaults to 10m
	Renewable   bool             // Renewable allows the issued token to be renewed
	Delegatable bool             // Delegatable allows the issued token to be delegated (e.g. for use with ActAs)
	ActAs       bool             // ActAs allows to request an ActAs token based on the passed Token.
	Token       string           // Token for Renew request or Issue request ActAs identity or to be exchanged.
	KeyType     string           // KeyType for requested token (if not set will be decucted from Userinfo and Certificate options)
	KeyID       string           // KeyID used for signing the requests
}

func (c *Client) newRequest(req TokenRequest, kind string, s *Signer) (internal.RequestSecurityToken, error) {
	if req.Lifetime == 0 {
		req.Lifetime = 5 * time.Minute
	}

	created := time.Now().UTC()
	rst := internal.RequestSecurityToken{
		TokenType:          c.Namespace,
		RequestType:        "http://docs.oasis-open.org/ws-sx/ws-trust/200512/" + kind,
		SignatureAlgorithm: internal.SHA256,
		Lifetime: &internal.Lifetime{
			Created: created.Format(internal.Time),
			Expires: created.Add(req.Lifetime).Format(internal.Time),
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
		KeyType:     req.KeyType,
	}

	if req.KeyType == "" {
		// Deduce KeyType based on Certificate nad Userinfo.
		if req.Certificate == nil {
			if req.Userinfo == nil {
				return rst, errors.New("one of TokenRequest Certificate or Userinfo is required")
			}
			rst.KeyType = "http://docs.oasis-open.org/ws-sx/ws-trust/200512/Bearer"
		} else {
			rst.KeyType = "http://docs.oasis-open.org/ws-sx/ws-trust/200512/PublicKey"
			// For HOK KeyID is required.
			if req.KeyID == "" {
				req.KeyID = newID()
			}
		}
	}

	if req.KeyID != "" {
		rst.UseKey = &internal.UseKey{Sig: req.KeyID}
		s.keyID = rst.UseKey.Sig
	}

	return rst, nil
}

func (s *Signer) setLifetime(lifetime *internal.Lifetime) error {
	var err error
	if lifetime != nil {
		s.Lifetime.Created, err = time.Parse(internal.Time, lifetime.Created)
		if err == nil {
			s.Lifetime.Expires, err = time.Parse(internal.Time, lifetime.Expires)
		}
	}
	return err
}

// Issue is used to request a security token.
// The returned Signer can be used to sign SOAP requests, such as the SessionManager LoginByToken method and the RequestSecurityToken method itself.
// One of TokenRequest Certificate or Userinfo is required, with Certificate taking precedence.
// When Certificate is set, a Holder-of-Key token will be requested.  Otherwise, a Bearer token is requested with the Userinfo credentials.
// See: http://docs.oasis-open.org/ws-sx/ws-trust/v1.4/errata01/os/ws-trust-1.4-errata01-os-complete.html#_Toc325658937
func (c *Client) Issue(ctx context.Context, req TokenRequest) (*Signer, error) {
	s := &Signer{
		Certificate: req.Certificate,
		keyID:       req.KeyID,
		Token:       req.Token,
		user:        req.Userinfo,
	}

	rst, err := c.newRequest(req, "Issue", s)
	if err != nil {
		return nil, err
	}

	if req.ActAs {
		rst.ActAs = &internal.Target{
			Token: req.Token,
		}
	}

	header := soap.Header{
		Security: s,
		Action:   fmt.Sprintf(`"%s"`, rst.Action()),
	}

	res, err := internal.Issue(c.WithHeader(ctx, header), c, &rst)
	if err != nil {
		return nil, err
	}

	s.Token = res.RequestSecurityTokenResponse.RequestedSecurityToken.Assertion

	return s, s.setLifetime(res.RequestSecurityTokenResponse.Lifetime)
}

// Renew is used to request a security token renewal.
func (c *Client) Renew(ctx context.Context, req TokenRequest) (*Signer, error) {
	s := &Signer{
		Certificate: req.Certificate,
	}

	rst, err := c.newRequest(req, "Renew", s)
	if err != nil {
		return nil, err
	}

	if req.Token == "" {
		return nil, errors.New("TokenRequest Token is required")
	}

	rst.RenewTarget = &internal.Target{Token: req.Token}

	header := soap.Header{
		Security: s,
		Action:   rst.Action(),
	}

	res, err := internal.Renew(c.WithHeader(ctx, header), c, &rst)
	if err != nil {
		return nil, err
	}

	s.Token = res.RequestedSecurityToken.Assertion

	return s, s.setLifetime(res.Lifetime)
}
