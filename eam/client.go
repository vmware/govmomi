/*
Copyright (c) 2021 VMware, Inc. All Rights Reserved.

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
