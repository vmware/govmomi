// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package lookup

import (
	"context"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net/url"

	"github.com/vmware/govmomi/internal"
	"github.com/vmware/govmomi/lookup/methods"
	"github.com/vmware/govmomi/lookup/types"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/soap"
	vim "github.com/vmware/govmomi/vim25/types"
)

const (
	Namespace = "lookup"
	Version   = "2.0"
	Path      = "/lookupservice" + vim25.Path
)

var (
	ServiceInstance = vim.ManagedObjectReference{
		Type:  "LookupServiceInstance",
		Value: "ServiceInstance",
	}
)

// Client is a soap.Client targeting the SSO Lookup Service API endpoint.
type Client struct {
	*soap.Client

	RoundTripper soap.RoundTripper

	ServiceContent types.LookupServiceContent

	// Rewrite when true changes EndpointURL Host to the VC connection's Host
	Rewrite bool
}

// NewClient returns a client targeting the SSO Lookup Service API endpoint.
func NewClient(ctx context.Context, c *vim25.Client) (*Client, error) {
	path := &url.URL{Path: Path}
	// PSC may be external, attempt to derive from sts.uri if not using envoy sidecar
	if !internal.UsingEnvoySidecar(c) && c.ServiceContent.Setting != nil {
		m := object.NewOptionManager(c, *c.ServiceContent.Setting)
		opts, err := m.Query(ctx, "config.vpxd.sso.sts.uri")
		if err == nil && len(opts) == 1 {
			u, err := url.Parse(opts[0].GetOptionValue().Value.(string))
			if err == nil {
				path.Scheme = u.Scheme
				path.Host = u.Host
			}
		}
	}

	// 1st try: use the URL from OptionManager as-is, continue to 2nd try on DNS error
	// 2nd try: use the URL from OptionManager, changing Host to vim25.Client's Host
	var attempts []error

	for _, rewrite := range []bool{false, true} {
		if rewrite {
			path.Host = c.URL().Host
		}

		sc := c.Client.NewServiceClient(path.String(), Namespace)
		sc.Version = Version
		client := &Client{Client: sc, RoundTripper: sc, Rewrite: rewrite}

		req := types.RetrieveServiceContent{
			This: ServiceInstance,
		}

		res, err := methods.RetrieveServiceContent(ctx, client, &req)
		if err != nil {
			attempts = append(attempts, err)
			continue
		}

		client.ServiceContent = res.Returnval

		return client, nil
	}

	return nil, errors.Join(attempts...)
}

// RoundTrip dispatches to the RoundTripper field.
func (c *Client) RoundTrip(ctx context.Context, req, res soap.HasFault) error {
	// Drop any operationID header, not used by lookup service
	ctx = context.WithValue(ctx, vim.ID{}, "")
	return c.RoundTripper.RoundTrip(ctx, req, res)
}

func (c *Client) List(ctx context.Context, filter *types.LookupServiceRegistrationFilter) ([]types.LookupServiceRegistrationInfo, error) {
	req := types.List{
		This:           *c.ServiceContent.ServiceRegistration,
		FilterCriteria: filter,
	}

	res, err := methods.List(ctx, c, &req)
	if err != nil {
		return nil, err
	}
	return res.Returnval, nil
}

func (c *Client) SiteID(ctx context.Context) (string, error) {
	req := types.GetSiteId{
		This: *c.ServiceContent.ServiceRegistration,
	}

	res, err := methods.GetSiteId(ctx, c, &req)
	if err != nil {
		return "", err
	}
	return res.Returnval, nil
}

// EndpointURL uses the Lookup Service to find the endpoint URL and thumbprint for the given filter.
// If the endpoint is found, its TLS certificate is also added to the vim25.Client's trusted host thumbprints.
// If the Lookup Service is not available, the given path is returned as the default.
func EndpointURL(ctx context.Context, c *vim25.Client, path string, filter *types.LookupServiceRegistrationFilter) string {
	// Services running on vCenter can bypass lookup service.
	if useSidecar := internal.UsingEnvoySidecar(c); useSidecar {
		return fmt.Sprintf("http://%s%s", c.URL().Host, path)
	}
	if lu, err := NewClient(ctx, c); err == nil {
		info, _ := lu.List(ctx, filter)
		if len(info) != 0 && len(info[0].ServiceEndpoints) != 0 {
			endpoint := &info[0].ServiceEndpoints[0]
			path = endpoint.Url

			if u, err := url.Parse(path); err == nil {
				if lu.Rewrite {
					u.Host = c.URL().Host
					path = u.String()
				} else {
					// Set thumbprint only for endpoints on hosts outside this vCenter.
					// Platform Services may live on multiple hosts.
					if c.URL().Host != u.Host && c.Thumbprint(u.Host) == "" {
						c.SetThumbprint(u.Host, endpointThumbprint(endpoint))
					}
				}
			}
		}
	}
	return path
}

// endpointThumbprint converts the base64 encoded endpoint certificate to a SHA1 thumbprint.
func endpointThumbprint(endpoint *types.LookupServiceRegistrationEndpoint) string {
	if len(endpoint.SslTrust) == 0 {
		return ""
	}
	enc := endpoint.SslTrust[0]

	b, err := base64.StdEncoding.DecodeString(enc)
	if err != nil {
		log.Printf("base64.Decode(%q): %s", enc, err)
		return ""
	}

	cert, err := x509.ParseCertificate(b)
	if err != nil {
		log.Printf("x509.ParseCertificate(%q): %s", enc, err)
		return ""
	}

	return soap.ThumbprintSHA1(cert)
}
