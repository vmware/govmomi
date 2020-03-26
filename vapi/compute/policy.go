/*
Copyright (c) 2019-2020 VMware, Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0.
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package compute

import (
	"context"
	"net/http"
	"path"

	"github.com/vmware/govmomi/vapi/compute/internal"
	"github.com/vmware/govmomi/vapi/rest"
)

// PolicyManager provides convenience methods to operate Compute Policy
type PolicyManager struct {
	*rest.Client
}

// NewPolicyManager creates a new PolicyManager with the given client
func NewPolicyManager(client *rest.Client) *PolicyManager {
	return &PolicyManager{
		Client: client,
	}
}

// Policy represents the structure a compute policy.
type Policy struct {
	Name        string `json:"name,omitempty"`
	Capability  string `json:"capability,omitempty"`
	Description string `json:"description,omitempty"`
	Policy      string `json:"policy,omitempty"`
	HostTag     string `json:"host_tag,omitempty"`
	VMTag       string `json:"vm_tag,omitempty"`
}

// Capability contains information about a compute policy capability.
type Capability struct {
	Capability  string `json:"capability,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// Create creates a new Compute Policy and returns the id.
func (c *PolicyManager) Create(ctx context.Context, policy Policy) (string, error) {
	spec := struct {
		Policy `json:"spec"`
	}{policy}

	r := c.Resource(internal.PolicyPath)
	var res string
	return res, c.Do(ctx, r.Request(http.MethodPost, spec), &res)
}

// Delete deletes a Compute Policy.
func (c *PolicyManager) Delete(ctx context.Context, id string) error {
	r := c.Resource(path.Join(internal.PolicyPath, id))
	return c.Do(ctx, r.Request(http.MethodDelete), nil)
}

// List returns information about the compute policies available in this vCenter server.
func (c *PolicyManager) List(ctx context.Context) ([]Policy, error) {
	r := c.Resource(internal.PolicyPath)
	var res []Policy
	return res, c.Do(ctx, r.Request(http.MethodGet), &res)
}

// Get returns information about a specific compute policy
func (c *PolicyManager) Get(ctx context.Context, id string) (*Policy, error) {
	r := c.Resource(path.Join(internal.PolicyPath, id))
	res := Policy{Policy: id}
	return &res, c.Do(ctx, r.Request(http.MethodGet), &res)
}

// ListCapability returns information about the compute policy capabilities available in this vCenter server.
func (c *PolicyManager) ListCapability(ctx context.Context) ([]Capability, error) {
	r := c.Resource(internal.PolicyCapabilitiesPath)
	var res []Capability
	return res, c.Do(ctx, r.Request(http.MethodGet), &res)
}

// GetCapability returns information about the compute policy capability id.
func (c *PolicyManager) GetCapability(ctx context.Context, id string) (*Capability, error) {
	r := c.Resource(path.Join(internal.PolicyCapabilitiesPath, id))
	res := Capability{Capability: id}
	return &res, c.Do(ctx, r.Request(http.MethodGet), &res)
}
