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

	"github.com/vmware/govmomi/vapi/internal"
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
	Class       string `json:"@class,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	VMTag       string `json:"vm_tag,omitempty"`
	Capability  string `json:"capability,omitempty"`
	Policy      string `json:"policy,omitempty"`
}

// Create creates a new Compute Policy and returns the id.
func (c *PolicyManager) Create(ctx context.Context, policy Policy) (string, error) {
	spec := struct {
		ComputePolicy Policy `json:"spec"`
	}{
		ComputePolicy: policy,
	}

	r := c.Resource(internal.ComputePolicyPath)
	var res string
	return res, c.Do(ctx, r.Request(http.MethodPost, spec), &res)
}

// List returns information about the compute policies available in this vCenter server.
func (c *PolicyManager) List(ctx context.Context) ([]Policy, error) {
	r := c.Resource(internal.ComputePolicyPath)
	var res []Policy
	return res, c.Do(ctx, r.Request(http.MethodGet), &res)
}
