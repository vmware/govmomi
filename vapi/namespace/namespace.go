/*
Copyright (c) 2020-2024 VMware, Inc. All Rights Reserved.

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
	"fmt"
	"net/http"
	"os"
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

// EnableClusterSpec defines a Tanzu Supervisor Cluster for creation.
// See: https://developer.vmware.com/apis/vsphere-automation/latest/vcenter/data-structures/NamespaceManagement/Clusters/EnableSpec/
// Since 7.0.0:-
type EnableClusterSpec struct {
	MasterDNSSearchDomains []string               `json:"master_DNS_search_domains,omitempty"`
	ImageStorage           ImageStorageSpec       `json:"image_storage"`
	NcpClusterNetworkSpec  *NcpClusterNetworkSpec `json:"ncp_cluster_network_spec,omitempty"`
	// Note: NcpClusterNetworkSpec is replaced by WorkloadNetworksSpec in vSphere 7.0u2+
	// Since 7.0u1:-
	WorkloadNetworksSpec    *WorkloadNetworksEnableSpec `json:"workload_networks_spec,omitempty"`
	MasterManagementNetwork *MasterManagementNetwork    `json:"master_management_network"`
	MasterDNSNames          []string                    `json:"Master_DNS_names,omitempty"`
	MasterNTPServers        []string                    `json:"master_NTP_servers,omitempty"`
	EphemeralStoragePolicy  string                      `json:"ephemeral_storage_policy,omitempty"`
	DefaultImageRepository  string                      `json:"default_image_repository,omitempty"`
	ServiceCidr             *Cidr                       `json:"service_cidr"`
	LoginBanner             string                      `json:"login_banner,omitempty"`
	// Was string until #2860:-
	SizeHint             *SizingHint           `json:"size_hint"`
	WorkerDNS            []string              `json:"worker_DNS,omitempty"`
	DefaultImageRegistry *DefaultImageRegistry `json:"default_image_registry,omitempty"`
	MasterDNS            []string              `json:"master_DNS,omitempty"`
	// Was string until #2860:-
	NetworkProvider                        *NetworkProvider        `json:"network_provider"`
	MasterStoragePolicy                    string                  `json:"master_storage_policy,omitempty"`
	DefaultKubernetesServiceContentLibrary string                  `json:"default_kubernetes_service_content_library,omitempty"`
	WorkloadNTPServers                     []string                `json:"workload_ntp_servers,omitempty"`
	LoadBalancerConfigSpec                 *LoadBalancerConfigSpec `json:"load_balancer_config_spec,omitempty"`
}

// SizingHint determines the size of the Tanzu Kubernetes Grid
// Supervisor cluster's kubeapi instances.
// Note: Only use TinySizingHint in non-production environments.
// Note: This is a secure coding pattern to avoid Stringly typed fields.
// See https://developer.vmware.com/apis/vsphere-automation/latest/vcenter/data-structures/NamespaceManagement/SizingHint/
// Since 7.0.0
type SizingHint struct {
	slug string
}

var (
	UndefinedSizingHint = SizingHint{""}
	TinySizingHint      = SizingHint{"TINY"}
	SmallSizingHint     = SizingHint{"SMALL"}
	MediumSizingHint    = SizingHint{"MEDIUM"}
	LargeSizingHint     = SizingHint{"LARGE"}
)

func (v SizingHint) String() string {
	return v.slug
}

func (v *SizingHint) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); nil != err {
		return err
	}
	v.FromString(s)
	return nil
}

func (v SizingHint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.slug)
}

func (v *SizingHint) FromString(s string) {
	v.slug = SizingHintFromString(s).slug
}

func SizingHintFromString(s string) SizingHint {
	if "TINY" == s {
		return TinySizingHint
	}
	if "SMALL" == s {
		return SmallSizingHint
	}
	if "MEDIUM" == s {
		return MediumSizingHint
	}
	if "LARGE" == s {
		return LargeSizingHint
	}
	return UndefinedSizingHint
}

// ImageStorageSpec defines the storage policy ID (not name) assigned to
// a Tanzu Kubernetes Grid cluster (supervisor or workload clusters)
// See https://developer.vmware.com/apis/vsphere-automation/latest/vcenter/data-structures/NamespaceManagement/Clusters/ImageStorageSpec/
// Since 7.0.0:-
type ImageStorageSpec struct {
	StoragePolicy string `json:"storage_policy"`
}

// Cidr defines an IPv4 CIDR range for a subnet.
// See https://developer.vmware.com/apis/vsphere-automation/latest/vcenter/data-structures/Namespaces/Instances/Ipv4Cidr/
// TODO decide whether to rename this in the Go API to match the vSphere API.
// Since 7.0.0:-
type Cidr struct {
	Address string `json:"address"`
	Prefix  int    `json:"prefix"`
}

// NcpClusterNetworkSpec defines an NSX-T network for a Tanzu Kubernetes
// Grid workload cluster in vSphere 7.0.0 until 7.0u1.
// Note: NcpClusterNetworkSpec is replaced by WorkloadNetworksSpec in 7.0u2+.
// See https://developer.vmware.com/apis/vsphere-automation/latest/vcenter/data-structures/NamespaceManagement/Clusters/NCPClusterNetworkEnableSpec/
// TODO decide whether to rename this in the Go API to match the vSphere API.
// Since 7.0.0:-
type NcpClusterNetworkSpec struct {
	NsxEdgeCluster           string `json:"nsx_edge_cluster,omitempty"`
	PodCidrs                 []Cidr `json:"pod_cidrs"`
	EgressCidrs              []Cidr `json:"egress_cidrs"`
	ClusterDistributedSwitch string `json:"cluster_distributed_switch,omitempty"`
	IngressCidrs             []Cidr `json:"ingress_cidrs"`
}

// NsxNetwork defines a supervisor or workload NSX-T network for use with
// a Tanzu Kubernetes Cluster.
// See https://developer.vmware.com/apis/vsphere-automation/latest/vcenter/data-structures/NamespaceManagement/Networks/NsxNetworkCreateSpec/
// TODO decide whether to rename this in the Go API to match the vSphere API.
// Since 7.0u3:-
type NsxNetwork struct {
	EgressCidrs           []Cidr `json:"egress_cidrs"`
	IngressCidrs          []Cidr `json:"ingress_cidrs"`
	LoadBalancerSize      string `json:"load_balancer_size"`
	NamespaceNetworkCidrs []Cidr `json:"namespace_network_cidrs"`
	NsxTier0Gateway       string `json:"nsx_tier0_gateway"`
	RoutedMode            bool   `json:"routed_mode"`
	SubnetPrefixLength    int    `json:"subnet_prefix_length"`
}

// IpRange specifies a contiguous set of IPv4 Addresses
// See https://developer.vmware.com/apis/vsphere-automation/latest/vcenter/data-structures/NamespaceManagement/IPRange/
// TODO decide whether to rename this in the Go API to match the vSphere API.
// Note: omitempty allows AddressRanges: []IpRange to become json []
// Since 7.0u1:-
type IpRange struct {
	Address string `json:"address,omitempty"`
	Count   int    `json:"count,omitempty"`
}

// IpAssignmentMode specifies whether DHCP or a static range assignment
// method is used. This is used for both Supervisor Cluster and Workload
// Cluster networks in 7.0u3 and above.
// See https://developer.vmware.com/apis/vsphere-automation/latest/vcenter/data-structures/NamespaceManagement/Networks/IPAssignmentMode/
// TODO decide whether to rename this in the Go API to match the vSphere API.
// Since 7.0u3:-
type IpAssignmentMode struct {
	slug string
}

var (
	UndefinedIpAssignmentMode   = IpAssignmentMode{""}
	DhcpIpAssignmentMode        = IpAssignmentMode{"DHCP"}
	StaticRangeIpAssignmentMode = IpAssignmentMode{"STATICRANGE"}
	// NOTE: Add new types at the END of this const to preserve previous values
)

func (v IpAssignmentMode) String() string {
	return v.slug
}
func (v *IpAssignmentMode) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); nil != err {
		return err
	}
	v.FromString(s)
	return nil
}

func (v IpAssignmentMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.slug)
}

func (v *IpAssignmentMode) FromString(s string) {
	v.slug = IpAssignmentModeFromString(s).slug
}

func IpAssignmentModeFromString(s string) IpAssignmentMode {
	if "DHCP" == s {
		return DhcpIpAssignmentMode
	}
	if "STATICRANGE" == s {
		return StaticRangeIpAssignmentMode
	}
	return UndefinedIpAssignmentMode
}

// See https://developer.vmware.com/apis/vsphere-automation/latest/vcenter/data-structures/NamespaceManagement/Networks/VsphereDVPGNetworkCreateSpec/
// Since 7.0u1:-
type VsphereDVPGNetworkCreateSpec struct {
	AddressRanges []IpRange `json:"address_ranges"`
	Gateway       string    `json:"gateway"`
	PortGroup     string    `json:"portgroup"`
	SubnetMask    string    `json:"subnet_mask"`
	// Since 7.0u3:-
	IpAssignmentMode *IpAssignmentMode `json:"ip_assignment_mode,omitempty"`
}

// NetworkProvider defines which type of
// See https://developer.vmware.com/apis/vsphere-automation/latest/vcenter/data-structures/NamespaceManagement/Clusters/NetworkProvider/
// Since 7.0.0:-
type NetworkProvider struct {
	slug string
}

var (
	UndefinedNetworkProvider           = NetworkProvider{""}
	NsxtContainerPluginNetworkProvider = NetworkProvider{"NSXT_CONTAINER_PLUGIN"}
	// Since 7.0u1:-
	VSphereNetworkProvider = NetworkProvider{"VSPHERE_NETWORK"}
	// TODO vSphere (as in product), Vsphere (as in tag), or VSphere (as in camel case)???
	// E.g. see from 7.0u3: https://developer.vmware.com/apis/vsphere-automation/latest/vcenter/data-structures/NamespaceManagement/NSXTier0Gateway/Summary/
)

func (v NetworkProvider) String() string {
	return v.slug
}

func (v *NetworkProvider) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); nil != err {
		return err
	}
	v.FromString(s)
	return nil
}

func (v NetworkProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.slug)
}

func (v *NetworkProvider) FromString(s string) {
	v.slug = ClusterNetworkProviderFromString(s).slug
}

func ClusterNetworkProviderFromString(s string) NetworkProvider {
	if "NSXT_CONTAINER_PLUGIN" == s {
		return NsxtContainerPluginNetworkProvider
	}
	if "VSPHERE_NETWORK" == s {
		return VSphereNetworkProvider
	}
	return UndefinedNetworkProvider
}

// NetworksCreateSpec specifies a Tanzu Kubernetes Grid Supervisor
// or Workload network that should be created.
// See https://developer.vmware.com/apis/vsphere-automation/latest/vcenter/data-structures/NamespaceManagement/Networks/CreateSpec/
// Since 7.0u1:-
type NetworksCreateSpec struct {
	Network         string                        `json:"network"`
	NetworkProvider *NetworkProvider              `json:"network_provider"`
	VSphereNetwork  *VsphereDVPGNetworkCreateSpec `json:"vsphere_network,omitempty"`
	// Since 7.0u3:-
	NsxNetwork *NsxNetwork `json:"nsx_network,omitempty"`
}

// WorkloadNetworksEnableSpec defines the primary workload network for a new
// Tanzu Kubernetes Grid supervisor cluster. This may be used by namespaces
// for workloads too.
// See https://developer.vmware.com/apis/vsphere-automation/latest/vcenter/data-structures/NamespaceManagement/Clusters/WorkloadNetworksEnableSpec/
// TODO decide whether to rename this in the Go API to match the vSphere API.
// Since 7.0u1:-
type WorkloadNetworksEnableSpec struct {
	SupervisorPrimaryWorkloadNetwork *NetworksCreateSpec `json:"supervisor_primary_workload_network"`
	// TODO also support other workload networks in network_list
	//      (These are not used for EnableCluster, and so left out for now)
}

// LoadBalancersServer defines an Avi or HA Proxy load balancer location.
// Host can be an IP Address (normally an Avi Management Virtual IP for
// the Avi Controller(s)) or a hostname.
// See https://developer.vmware.com/apis/vsphere-automation/latest/vcenter/data-structures/NamespaceManagement/LoadBalancers/Server/
// Since 7.0u1:-
type LoadBalancersServer struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

// AviConfigCreateSpec defines full information for the linking of
// a Tanzu Kubernetes Grid enabled vSphere cluster to an NSX
// Advanced Load Balancer (formerly Avi Load Balancer)
// See https://developer.vmware.com/apis/vsphere-automation/latest/vcenter/data-structures/NamespaceManagement/LoadBalancers/AviConfigCreateSpec/
// Since 7.0u2:-
type AviConfigCreateSpec struct {
	CertificateAuthorityChain string               `json:"certificate_authority_chain"`
	Password                  string               `json:"password"`
	Server                    *LoadBalancersServer `json:"server"`
	Username                  string               `json:"username"`
}

// HAProxyConfigCreateSpec defines full information for the linking of
// a Tanzu Kubernetes Grid enabled vSphere cluster to a HA Proxy
// Load Balancer.
// Note: HA Proxy is not supported in vSphere 7.0u3 and above. Use Avi
//
//	with vSphere networking, or NSX-T networking, instead.
//
// See https://developer.vmware.com/apis/vsphere-automation/latest/vcenter/data-structures/NamespaceManagement/LoadBalancers/HAProxyConfigCreateSpec/
// Since 7.0u1:-
// Deprecated: HA Proxy is being deprecated in vSphere 9.0. Use
// Avi with vSphere networking, or NSX-T networking, instead.
type HAProxyConfigCreateSpec struct {
	CertificateAuthorityChain string                `json:"certificate_authority_chain"`
	Password                  string                `json:"password"`
	Servers                   []LoadBalancersServer `json:"servers"`
	Username                  string                `json:"username"`
}

// A LoadBalancerProvider is an enum type that defines
// the Load Balancer technology in use in a Tanzu Kubernetes Grid
// cluster.
// Note: If invalid or undefined (E.g. if a newer/older vSphere
//
//	version is used whose option isn't listed) then the
//	UndefinedLoadBalancerProvider value shall be set.
//	This translates to an empty string, removing its element
//	from the produces JSON.
//
// See https://developer.vmware.com/apis/vsphere-automation/latest/vcenter/data-structures/NamespaceManagement/LoadBalancers/Provider/
type LoadBalancerProvider struct {
	slug string
}

var (
	UndefinedLoadBalancerProvider = LoadBalancerProvider{""}
	// Deprecated: HA Proxy is being deprecated in vSphere 9.0. Use
	// Avi vSphere networking, or NSX-T networking, instead.
	HAProxyLoadBalancerProvider = LoadBalancerProvider{"HA_PROXY"}
	AviLoadBalancerProvider     = LoadBalancerProvider{"AVI"}
)

func (v LoadBalancerProvider) String() string {
	return v.slug
}

func (v *LoadBalancerProvider) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); nil != err {
		return err
	}
	v.FromString(s)
	return nil
}

func (v LoadBalancerProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.slug)
}

func (v *LoadBalancerProvider) FromString(s string) {
	v.slug = LoadBalancerFromString(s).slug
}

func LoadBalancerFromString(s string) LoadBalancerProvider {
	if "HA_PROXY" == s {
		return HAProxyLoadBalancerProvider
	}
	if "AVI" == s {
		return AviLoadBalancerProvider
	}
	return UndefinedLoadBalancerProvider
}

// LoadBalancerConfigSpec defines LoadBalancer options for Tanzu
// Kubernetes Grid, both for the Supervisor Cluster and for
// Workload Cluster kubeapi endpoints, and services of type
// LoadBalancer
// See https://developer.vmware.com/apis/vsphere-automation/latest/vcenter/data-structures/NamespaceManagement/LoadBalancers/ConfigSpec/
// Since 7.0u1:-
type LoadBalancerConfigSpec struct {
	// AddressRanges removed since 7.0u2:- (Now in workload network spec)
	AddressRanges           []IpRange                `json:"address_ranges,omitempty"` // omitempty to prevent null being the value
	HAProxyConfigCreateSpec *HAProxyConfigCreateSpec `json:"ha_proxy_config_create_spec,omitempty"`
	// Optional for create:-
	Id       string                `json:"id"`
	Provider *LoadBalancerProvider `json:"provider"`
	// Since 7.0u2:-
	AviConfigCreateSpec *AviConfigCreateSpec `json:"avi_config_create_spec,omitempty"`
}

// Since 7.0.0:-
type AddressRange struct {
	SubnetMask      string `json:"subnet_mask,omitempty"`
	StartingAddress string `json:"starting_address"`
	Gateway         string `json:"gateway"`
	AddressCount    int    `json:"address_count,omitempty"`
}

// Since 7.0.0:-
type MasterManagementNetwork struct {
	Mode         *IpAssignmentMode `json:"mode"`
	FloatingIP   string            `json:"floating_IP,omitempty"`
	AddressRange *AddressRange     `json:"address_range,omitempty"`
	Network      string            `json:"network"`
}

// Since 7.0.0:-
type DefaultImageRegistry struct {
	Hostname string `json:"hostname"`
	Port     int    `json:"port,omitempty"`
}

// EnableCluster enables vSphere Namespaces on the specified cluster, using the given spec.
// See https://developer.vmware.com/apis/vsphere-automation/latest/vcenter/api/vcenter/namespace-management/clusters/clusteractionenable/post/
func (c *Manager) EnableCluster(ctx context.Context, id string, spec *EnableClusterSpec) error {
	var response interface{}
	url := c.Resource(path.Join(internal.NamespaceClusterPath, id)).WithParam("action", "enable")
	fmt.Fprint(os.Stdout, spec)
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

type KubernetesStatus struct {
	slug string
}

var (
	UndefinedKubernetesStatus = KubernetesStatus{""}
	ReadyKubernetesStatus     = KubernetesStatus{"READY"}
	WarningKubernetesStatus   = KubernetesStatus{"WARNING"}
	ErrorKubernetesStatus     = KubernetesStatus{"ERROR"}
	// NOTE: Add new types at the END of this const to preserve previous values
)

func (v KubernetesStatus) String() string {
	return v.slug
}

func (v *KubernetesStatus) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); nil != err {
		return err
	}
	v.FromString(s)
	return nil
}

func (v KubernetesStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.slug)
}

func (v *KubernetesStatus) FromString(s string) {
	v.slug = KubernetesStatusFromString(s).slug
}

func KubernetesStatusFromString(s string) KubernetesStatus {
	if "READY" == s {
		return ReadyKubernetesStatus
	}
	if "WARNING" == s {
		return WarningKubernetesStatus
	}
	if "ERROR" == s {
		return ErrorKubernetesStatus
	}
	return UndefinedKubernetesStatus
}

type ConfigStatus struct {
	slug string
}

var (
	UndefinedConfigStatus   = ConfigStatus{""}
	ConfiguringConfigStatus = ConfigStatus{"CONFIGURING"}
	RemovingConfigStatus    = ConfigStatus{"REMOVING"}
	RunningConfigStatus     = ConfigStatus{"RUNNING"}
	ErrorConfigStatus       = ConfigStatus{"ERROR"}
	// NOTE: Add new types at the END of this const to preserve previous values
)

func (v ConfigStatus) String() string {
	return v.slug
}

func (v *ConfigStatus) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); nil != err {
		return err
	}
	v.FromString(s)
	return nil
}

func (v ConfigStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.slug)
}

func (v *ConfigStatus) FromString(s string) {
	v.slug = ConfigStatusFromString(s).slug
}

func ConfigStatusFromString(s string) ConfigStatus {
	if "CONFIGURING" == s {
		return ConfiguringConfigStatus
	}
	if "REMOVING" == s {
		return RemovingConfigStatus
	}
	if "RUNNING" == s {
		return RunningConfigStatus
	}
	if "ERROR" == s {
		return ErrorConfigStatus
	}
	return UndefinedConfigStatus
}

// TODO CHANGE ENUMS TO: https://stackoverflow.com/questions/53569573/parsing-string-to-enum-from-json-in-golang
// Also: https://eagain.net/articles/go-json-kind/

// ClusterSummary for a cluster with vSphere Namespaces enabled.
// See https://developer.vmware.com/apis/vsphere-automation/latest/vcenter/data-structures/NamespaceManagement/Clusters/Summary/
// TODO plural vs singular - consistency with REST API above vs Go
// Since 7.0.0:-
type ClusterSummary struct {
	ID   string `json:"cluster"`
	Name string `json:"cluster_name"`
	// Was string until #2860:-
	KubernetesStatus *KubernetesStatus `json:"kubernetes_status"`
	// Was string until #2860:-
	ConfigStatus *ConfigStatus `json:"config_status"`
}

// TODO whether to replace the below with a Go GUID (json to string) reference type? (I.e. replace ClusterSummary.ID string with ID ManagedObjectID)
// Reference implements the mo.Reference interface
func (c *ClusterSummary) Reference() types.ManagedObjectReference {
	return types.ManagedObjectReference{
		Type:  "ClusterComputeResource",
		Value: c.ID, // TODO replace with c.ID.(string) when ID changes from String to ManagedObjectID
	}
}

// ListClusters returns a summary of all clusters with vSphere Namespaces enabled.
func (c *Manager) ListClusters(ctx context.Context) ([]ClusterSummary, error) {
	var res []ClusterSummary
	url := c.Resource(internal.NamespaceClusterPath)
	return res, c.Do(ctx, url.Request(http.MethodGet), &res)
}

// SupportBundleToken information about the token required in the HTTP GET request to generate the support bundle.
// Since 7.0.0:-
type SupportBundleToken struct {
	Expiry string `json:"expiry"`
	Token  string `json:"token"`
}

// SupportBundleLocation contains the URL to download the per-cluster support bundle from, as well as a token required.
// Since 7.0.0:-
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

// Since 7.0.0:-
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

// Since 7.0.0:-
type EdgeClusterCompatibilitySummary struct {
	Compatible  bool   `json:"compatible"`
	EdgeCluster string `json:"edge_cluster"`
	DisplayName string `json:"display_name"`
}

func (c *Manager) ListCompatibleEdgeClusters(ctx context.Context, clusterId string, switchId string) (result []EdgeClusterCompatibilitySummary, err error) {
	listUrl := c.Resource(internal.NamespaceEdgeClusterCompatibility).
		WithParam("cluster", clusterId).
		WithParam("compatible", "true").
		WithPathEncodedParam("distributed_switch", switchId)
	return result, c.Do(ctx, listUrl.Request(http.MethodGet), &result)
}

// NamespacesInstanceStats https://developer.vmware.com/apis/vsphere-automation/v7.0U3/vcenter/data-structures/Namespaces/Instances/Stats/
type NamespacesInstanceStats struct {
	CpuUsed     int64 `json:"cpu_used"`
	MemoryUsed  int64 `json:"memory_used"`
	StorageUsed int64 `json:"storage_used"`
}

// NamespacesInstanceSummary https://developer.vmware.com/apis/vsphere-automation/v7.0U3/vcenter/data-structures/Namespaces/Instances/Summary/
type NamespacesInstanceSummary struct {
	ClusterId            string                  `json:"cluster"`
	Namespace            string                  `json:"namespace"`
	ConfigStatus         string                  `json:"config_status"`
	Description          string                  `json:"description"`
	Stats                NamespacesInstanceStats `json:"stats"`
	SelfServiceNamespace bool                    `json:"self_service_namespace,omitempty"`
}

type LocalizableMessage struct {
	Details  interface{} `json:"details"`
	Severity string      `json:"severity"`
}

// NamespacesInstanceInfo https://developer.vmware.com/apis/vsphere-automation/v7.0U3/vcenter/data-structures/Namespaces/Instances/Info/
type NamespacesInstanceInfo struct {
	ClusterId            string                  `json:"cluster"`
	ConfigStatus         string                  `json:"config_status"`
	Description          string                  `json:"description"`
	Stats                NamespacesInstanceStats `json:"stats"`
	SelfServiceNamespace bool                    `json:"self_service_namespace,omitempty"`
	Messages             []LocalizableMessage    `json:"message"`
	VmServiceSpec        VmServiceSpec           `json:"vm_service_spec,omitempty"`
	StorageSpecs         []StorageSpec           `json:"storage_specs,omitempty"`
}

// NamespacesInstanceCreateSpec https://developer.vmware.com/apis/vsphere-automation/v7.0U3/vcenter/data-structures/Namespaces/Instances/CreateSpec/
type NamespacesInstanceCreateSpec struct {
	Cluster       string        `json:"cluster"`
	Namespace     string        `json:"namespace"`
	VmServiceSpec VmServiceSpec `json:"vm_service_spec,omitempty"`
	StorageSpecs  []StorageSpec `json:"storage_specs,omitempty"`
}

// VmServiceSpec https://developer.vmware.com/apis/vsphere-automation/v7.0U3/vcenter/data-structures/Namespaces/Instances/VMServiceSpec/
type VmServiceSpec struct {
	ContentLibraries []string `json:"content_libraries,omitempty"`
	VmClasses        []string `json:"vm_classes,omitempty"`
}

// StorageSpec https://developer.broadcom.com/apis/vsphere-automation-api/v7.0U3/vcenter/data-structures/Namespaces_Instances_StorageSpec/
type StorageSpec struct {
	Policy string `json:"policy"`
	Limit  int64  `json:"limit,omitempty"`
}

// NamespacesInstanceUpdateSpec https://developer.vmware.com/apis/vsphere-automation/v7.0U3/vcenter/data-structures/Namespaces/Instances/UpdateSpec/
type NamespacesInstanceUpdateSpec struct {
	VmServiceSpec VmServiceSpec `json:"vm_service_spec,omitempty"`
	StorageSpecs  []StorageSpec `json:"storage_specs,omitempty"`
}

// ListNamespaces https://developer.vmware.com/apis/vsphere-automation/v7.0U3/vcenter/api/vcenter/namespaces/instances/get/
func (c *Manager) ListNamespaces(ctx context.Context) ([]NamespacesInstanceSummary, error) {
	resource := c.Resource(internal.NamespacesPath)
	request := resource.Request(http.MethodGet)
	var result []NamespacesInstanceSummary
	return result, c.Do(ctx, request, &result)
}

// GetNamespace https://developer.vmware.com/apis/vsphere-automation/v7.0U3/vcenter/api/vcenter/namespaces/instances/namespace/get/
func (c *Manager) GetNamespace(ctx context.Context, namespace string) (NamespacesInstanceInfo, error) {
	resource := c.Resource(internal.NamespacesPath).WithSubpath(namespace)
	request := resource.Request(http.MethodGet)
	var result NamespacesInstanceInfo
	return result, c.Do(ctx, request, &result)
}

// CreateNamespace https://developer.vmware.com/apis/vsphere-automation/v7.0U3/vcenter/api/vcenter/namespaces/instances/post/
func (c *Manager) CreateNamespace(ctx context.Context, spec NamespacesInstanceCreateSpec) error {
	resource := c.Resource(internal.NamespacesPath)
	request := resource.Request(http.MethodPost, spec)
	return c.Do(ctx, request, nil)
}

// UpdateNamespace https://developer.vmware.com/apis/vsphere-automation/v7.0U3/vcenter/api/vcenter/namespaces/instances/namespace/patch/
func (c *Manager) UpdateNamespace(ctx context.Context, namespace string, spec NamespacesInstanceUpdateSpec) error {
	resource := c.Resource(internal.NamespacesPath).WithSubpath(namespace)
	request := resource.Request(http.MethodPatch, spec)
	return c.Do(ctx, request, nil)
}

// DeleteNamespace https://developer.vmware.com/apis/vsphere-automation/v7.0U3/vcenter/api/vcenter/namespaces/instances/namespace/delete/
func (c *Manager) DeleteNamespace(ctx context.Context, namespace string) error {
	resource := c.Resource(internal.NamespacesPath).WithSubpath(namespace)
	request := resource.Request(http.MethodDelete)
	return c.Do(ctx, request, nil)
}

// RegisterVM https://developer.broadcom.com/xapis/vsphere-automation-api/latest/vcenter/api/vcenter/namespaces/instances/namespace/registervm/post/
func (c *Manager) RegisterVM(ctx context.Context, namespace string, spec RegisterVMSpec) (string, error) {
	resource := c.Resource(internal.NamespacesPath).WithSubpath(namespace).WithSubpath("registervm")
	request := resource.Request(http.MethodPost, spec)
	var task string
	return task, c.Do(ctx, request, &task)
}

// VirtualMachineClassInfo https://developer.vmware.com/apis/vsphere-automation/v7.0U3/vcenter/data-structures/NamespaceManagement/VirtualMachineClasses/Info/
type VirtualMachineClassInfo struct {
	ConfigStatus      string               `json:"config_status"`
	Description       string               `json:"description"`
	Id                string               `json:"id"`
	CpuCount          int64                `json:"cpu_count"`
	MemoryMb          int64                `json:"memory_mb"`
	Messages          []LocalizableMessage `json:"messages"`
	Namespaces        []string             `json:"namespaces"`
	Vms               []string             `json:"vms"`
	Devices           VirtualDevices       `json:"devices"`
	CpuReservation    int64                `json:"cpu_reservation,omitempty"`
	MemoryReservation int64                `json:"memory_reservation,omitempty"`
}

// VirtualMachineClassCreateSpec https://developer.vmware.com/apis/vsphere-automation/v7.0U3/vcenter/data-structures/NamespaceManagement/VirtualMachineClasses/CreateSpec/
type VirtualMachineClassCreateSpec struct {
	Id                string         `json:"id"`
	CpuCount          int64          `json:"cpu_count"`
	MemoryMb          int64          `json:"memory_MB"`
	CpuReservation    int64          `json:"cpu_reservation,omitempty"`
	MemoryReservation int64          `json:"memory_reservation,omitempty"`
	Devices           VirtualDevices `json:"devices"`
}

// VirtualMachineClassUpdateSpec https://developer.vmware.com/apis/vsphere-automation/v7.0U3/vcenter/data-structures/NamespaceManagement/VirtualMachineClasses/UpdateSpec/
type VirtualMachineClassUpdateSpec struct {
	Id                string         `json:"id"`
	CpuCount          int64          `json:"cpu_count"`
	MemoryMb          int64          `json:"memory_MB"`
	CpuReservation    int64          `json:"cpu_reservation,omitempty"`
	MemoryReservation int64          `json:"memory_reservation,omitempty"`
	Devices           VirtualDevices `json:"devices"`
}

// DirectPathIoDevice https://developer.vmware.com/apis/vsphere-automation/v7.0U3/vcenter/data-structures/NamespaceManagement/VirtualMachineClasses/DynamicDirectPathIODevice/
type DirectPathIoDevice struct {
	CustomLabel string `json:"custom_label,omitempty"`
	DeviceId    int64  `json:"device_id"`
	VendorId    int64  `json:"vendor_id"`
}

// VgpuDevice https://developer.vmware.com/apis/vsphere-automation/v7.0U3/vcenter/data-structures/NamespaceManagement/VirtualMachineClasses/VGPUDevice/
type VgpuDevice struct {
	ProfileName string `json:"profile_name"`
}

// VirtualDevices https://developer.vmware.com/apis/vsphere-automation/v7.0U3/vcenter/data-structures/NamespaceManagement/VirtualMachineClasses/VirtualDevices/
type VirtualDevices struct {
	DirectPathIoDevices []DirectPathIoDevice `json:"direct_path_io_devices,omitempty"`
	VgpuDevices         []VgpuDevice         `json:"vgpu_devices,omitempty"`
}

// RegisterVMSpec https://developer.broadcom.com/xapis/vsphere-automation-api/latest/vcenter/data-structures/Namespaces_Instances_RegisterVMSpec/
type RegisterVMSpec struct {
	VM string `json:"vm"`
}

// ListVmClasses https://developer.vmware.com/apis/vsphere-automation/v7.0U3/vcenter/api/vcenter/namespace-management/virtual-machine-classes/get/
func (c *Manager) ListVmClasses(ctx context.Context) ([]VirtualMachineClassInfo, error) {
	resource := c.Resource(internal.VmClassesPath)
	request := resource.Request(http.MethodGet)
	var result []VirtualMachineClassInfo
	return result, c.Do(ctx, request, &result)
}

// GetVmClass https://developer.vmware.com/apis/vsphere-automation/v7.0U3/vcenter/api/vcenter/namespace-management/virtual-machine-classes/vm_class/get/
func (c *Manager) GetVmClass(ctx context.Context, vmClass string) (VirtualMachineClassInfo, error) {
	resource := c.Resource(internal.VmClassesPath).WithSubpath(vmClass)
	request := resource.Request(http.MethodGet)
	var result VirtualMachineClassInfo
	return result, c.Do(ctx, request, &result)
}

// CreateVmClass https://developer.vmware.com/apis/vsphere-automation/v7.0U3/vcenter/api/vcenter/namespace-management/virtual-machine-classes/post/
func (c *Manager) CreateVmClass(ctx context.Context, spec VirtualMachineClassCreateSpec) error {
	resource := c.Resource(internal.VmClassesPath)
	request := resource.Request(http.MethodPost, spec)
	return c.Do(ctx, request, nil)
}

// DeleteVmClass https://developer.vmware.com/apis/vsphere-automation/v7.0U3/vcenter/api/vcenter/namespace-management/virtual-machine-classes/vm_class/delete/
func (c *Manager) DeleteVmClass(ctx context.Context, vmClass string) error {
	resource := c.Resource(internal.VmClassesPath).WithSubpath(vmClass)
	request := resource.Request(http.MethodDelete)
	return c.Do(ctx, request, nil)
}

// UpdateVmClass https://developer.vmware.com/apis/vsphere-automation/v7.0U3/vcenter/api/vcenter/namespace-management/virtual-machine-classes/vm_class/patch/
func (c *Manager) UpdateVmClass(ctx context.Context, vmClass string, spec VirtualMachineClassUpdateSpec) error {
	resource := c.Resource(internal.VmClassesPath).WithSubpath(vmClass)
	request := resource.Request(http.MethodPatch, spec)
	return c.Do(ctx, request, nil)
}
