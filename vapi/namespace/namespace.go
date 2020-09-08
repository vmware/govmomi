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
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"path"

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

// SupportBundleToken information about the token required in the HTTP GET request to generate the support bundle.
type SupportBundleToken struct {
	Expiry string `json:"expiry"`
	Token  string `json:"token"`
}

// SupportBundleLocation contains the URL to download the per-cluster support bundle from, as well as a token required.
type SupportBundleLocation struct {
	Token SupportBundleToken `json:"wcp_support_bundle_token"`
	URL   string             `json:"url"`
}

// CreateSupportBundle retrieves the cluster's Namespaces-related support bundle.
func (c *Manager) CreateSupportBundle(ctx context.Context, id string) (*SupportBundleLocation, error) {
	var res SupportBundleLocation
	url := c.Resource(path.Join(internal.NamespaceClusterPath, id, "support-bundle"))
	return &res, c.Do(ctx, url.Request(http.MethodPost), &res)
}

// SupportBundleRequest returns an http.Request which can be used to download the given support bundle.
func (c *Manager) SupportBundleRequest(ctx context.Context, bundle *SupportBundleLocation) (*http.Request, error) {
	token := internal.SupportBundleToken{Value: bundle.Token.Token}

	var b bytes.Buffer
	err := json.NewEncoder(&b).Encode(token)
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPost, bundle.URL, &b)
}
