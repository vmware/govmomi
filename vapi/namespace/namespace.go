/*
Copyright (c) 2020 VMware, Inc. All Rights Reserved.

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

package namespace

import (
	"context"
	"net/http"

	"github.com/vmware/govmomi/vapi/namespace/internal"
	"github.com/vmware/govmomi/vapi/rest"
	"github.com/vmware/govmomi/vim25/types"
)

// Manager extends rest.Client, adding namespace related methods.
type Manager struct {
	*rest.Client
}

// NewManager creates a new Manager instance with the given client.
func NewManager(client *rest.Client) *Manager {
	return &Manager{
		Client: client,
	}
}

// ClusterSummary for a cluster with vSphere Namespaces enabled.
type ClusterSummary struct {
	ID               string `json:"cluster"`
	Name             string `json:"cluster_name"`
	KubernetesStatus string `json:"kubernetes_status"`
	ConfigStatus     string `json:"config_status"`
}

// Reference implements the mo.Reference interface
func (c *ClusterSummary) Reference() types.ManagedObjectReference {
	return types.ManagedObjectReference{
		Type:  "ClusterComputeResource",
		Value: c.ID,
	}
}

// ListClusters returns a summary of all clusters with vSphere Namespaces enabled.
func (c *Manager) ListClusters(ctx context.Context) ([]ClusterSummary, error) {
	var res []ClusterSummary
	url := c.Resource(internal.NamespaceClusterPath)
	return res, c.Do(ctx, url.Request(http.MethodGet), &res)
}
