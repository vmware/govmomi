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

type EnableClusterSpec struct {
	MasterDNSSearchDomains                 []string                 `json:"master_DNS_search_domains,omitempty"`
	ImageStorage                           ImageStorage             `json:"image_storage"`
	NcpClusterNetworkSpec                  *NcpClusterNetworkSpec   `json:"ncp_cluster_network_spec"`
	MasterManagementNetwork                *MasterManagementNetwork `json:"master_management_network"`
	MasterDNSNames                         []string                 `json:"Master_DNS_names,omitempty"`
	MasterNTPServers                       []string                 `json:"master_NTP_servers,omitempty"`
	EphemeralStoragePolicy                 string                   `json:"ephemeral_storage_policy,omitempty"`
	DefaultImageRepository                 string                   `json:"default_image_repository,omitempty"`
	ServiceCidr                            *Cidr                    `json:"service_cidr"`
	LoginBanner                            string                   `json:"login_banner,omitempty"`
	SizeHint                               string                   `json:"size_hint"`
	WorkerDNS                              []string                 `json:"worker_DNS,omitempty"`
	DefaultImageRegistry                   *DefaultImageRegistry    `json:"default_image_registry,omitempty"`
	MasterDNS                              []string                 `json:"master_DNS,omitempty"`
	NetworkProvider                        string                   `json:"network_provider"`
	MasterStoragePolicy                    string                   `json:"master_storage_policy,omitempty"`
	DefaultKubernetesServiceContentLibrary string                   `json:"default_kubernetes_service_content_library,omitempty"`
}

type ImageStorage struct {
	StoragePolicy string `json:"storage_policy"`
}

type Cidr struct {
	Address string `json:"address"`
	Prefix  int    `json:"prefix"`
}

type NcpClusterNetworkSpec struct {
	NsxEdgeCluster           string `json:"nsx_edge_cluster,omitempty"`
	PodCidrs                 []Cidr `json:"pod_cidrs"`
	EgressCidrs              []Cidr `json:"egress_cidrs"`
	ClusterDistributedSwitch string `json:"cluster_distributed_switch,omitempty"`
	IngressCidrs             []Cidr `json:"ingress_cidrs"`
}

type AddressRange struct {
	SubnetMask      string `json:"subnet_mask,omitempty"`
	StartingAddress string `json:"starting_address"`
	Gateway         string `json:"gateway"`
	AddressCount    int    `json:"address_count,omitempty"`
}

type MasterManagementNetwork struct {
	Mode         string        `json:"mode"`
	FloatingIP   string        `json:"floating_IP,omitempty"`
	AddressRange *AddressRange `json:"address_range,omitempty"`
	Network      string        `json:"network"`
}

type DefaultImageRegistry struct {
	Hostname string `json:"hostname"`
	Port     int    `json:"port,omitempty"`
}

// EnableCluster enables vSphere Namespaces on the specified cluster, using the given spec.
func (c *Manager) EnableCluster(ctx context.Context, id string, spec *EnableClusterSpec) error {
	var response interface{}
	url := c.Resource(path.Join(internal.NamespaceClusterPath, id)).WithParam("action", "enable")
	err := c.Do(ctx, url.Request(http.MethodPost, spec), response)
	return err
}

// EnableCluster enables vSphere Namespaces on the specified cluster, using the given spec.
func (c *Manager) DisableCluster(ctx context.Context, id string) error {
	var response interface{}
	url := c.Resource(path.Join(internal.NamespaceClusterPath, id)).WithParam("action", "disable")
	err := c.Do(ctx, url.Request(http.MethodPost), response)
	return err
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

type DistributedSwitchCompatibilitySummary struct {
	Compatible        bool   `json:"compatible"`
	DistributedSwitch string `json:"distributed_switch"`
}

func (c *Manager) ListCompatibleDistributedSwitches(ctx context.Context, clusterId string) (result []DistributedSwitchCompatibilitySummary, err error) {
	listUrl := c.Resource(internal.NamespaceDistributedSwitchCompatibility).
		WithParam("cluster", clusterId).
		WithParam("compatible", "true")
	return result, c.Do(ctx, listUrl.Request(http.MethodGet), &result)
}

type EdgeClusterCompatibilitySummary struct {
	Compatible  bool   `json:"compatible"`
	EdgeCluster string `json:"edge_cluster"`
	DisplayName string `json:"display_name"`
}

func (c *Manager) ListCompatibleEdgeClusters(ctx context.Context, clusterId string, switchId string) (result []EdgeClusterCompatibilitySummary, err error) {
	listUrl := c.Resource(internal.NamespaceEdgeClusterCompatibility).
		WithParam("cluster", clusterId).
		WithParam("distributed_switch", switchId).
		WithParam("compatible", "true")
	return result, c.Do(ctx, listUrl.Request(http.MethodGet), &result)
}
