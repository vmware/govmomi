// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package eam

import (
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/soap"
)

const (
	// Namespace is the namespace for EAM SOAP operations.
	Namespace = "eam"

	// Path is the path to the EAM service.
	Path = "/eam/sdk"
)

// Client is a client for the ESX Agent Manager API.
type Client struct {
	*soap.Client
}

// NewClient returns a new EAM client.
func NewClient(c *vim25.Client) *Client {
	return &Client{
		Client: c.Client.NewServiceClient(Path, Namespace),
	}
}
