/*
Copyright (c) 2019 VMware, Inc. All Rights Reserved.

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

package logging

import (
	"context"
	"net/http"

	"github.com/vmware/govmomi/vapi/rest"
)

const ApplianceLoggingForwardingPath = "/appliance/logging/forwarding"

// Forwarding provides convenience methods to configure appliance logging forwarding.
type Forwarding struct {
	*rest.Client
}

// NewForwarding creates a new Forwarding with the given client
func NewForwarding(client *rest.Client) *Forwarding {
	return &Forwarding{
		Client: client,
	}
}

// Config represents configuration for log message forwarding.
type Config struct {
	Hostname string `json:"hostname,omitempty"`
	Port     int    `json:"port,omitempty"`
	Protocol string `json:"protocol,omitempty"`
}

func (f *Forwarding) getForwardingResource() *rest.Resource {
	return f.Resource(ApplianceLoggingForwardingPath)
}

// Config returns all logging forwarding config.
func (f *Forwarding) Config(ctx context.Context) ([]Config, error) {
	r := f.getForwardingResource()
	var res []Config
	return res, f.Do(ctx, r.Request(http.MethodGet), &res)
}
