/*
Copyright (c) 2014 VMware, Inc. All Rights Reserved.

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

package govmomi

import (
	"net/url"

	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/session"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
	"golang.org/x/net/context"
)

type Client struct {
	*vim25.Client

	SessionManager *session.Manager
}

// NewClient creates a new client from a URL. The client authenticates with the
// server before returning if the URL contains user information.
func NewClient(ctx context.Context, u *url.URL, insecure bool) (*Client, error) {
	soapClient := soap.NewClient(u, insecure)
	vimClient, err := vim25.NewClient(ctx, soapClient)
	if err != nil {
		return nil, err
	}

	c := &Client{
		Client:         vimClient,
		SessionManager: session.NewManager(vimClient),
	}

	// Only login if the URL contains user information.
	if u.User != nil {
		err = c.Login(ctx, u.User)
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

// Login dispatches to the SessionManager.
func (c *Client) Login(ctx context.Context, u *url.Userinfo) error {
	return c.SessionManager.Login(ctx, u)
}

// Logout dispatches to the SessionManager.
func (c *Client) Logout(ctx context.Context) error {
	// Close any idle connections after logging out.
	defer c.Client.CloseIdleConnections()
	return c.SessionManager.Logout(ctx)
}

// PropertyCollector returns the session's default property collector.
func (c *Client) PropertyCollector() *property.Collector {
	return property.DefaultCollector(c.Client)
}

// RetrieveOne dispatches to the Retrieve function on the default property collector.
func (c *Client) RetrieveOne(ctx context.Context, obj types.ManagedObjectReference, p []string, dst interface{}) error {
	return c.PropertyCollector().RetrieveOne(ctx, obj, p, dst)
}

// Retrieve dispatches to the Retrieve function on the default property collector.
func (c *Client) Retrieve(ctx context.Context, objs []types.ManagedObjectReference, p []string, dst interface{}) error {
	return c.PropertyCollector().Retrieve(ctx, objs, p, dst)
}

// Wait dispatches to property.Wait.
func (c *Client) Wait(ctx context.Context, obj types.ManagedObjectReference, ps []string, f func([]types.PropertyChange) bool) error {
	return property.Wait(ctx, c.PropertyCollector(), obj, ps, f)
}
