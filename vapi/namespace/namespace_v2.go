// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package namespace

import (
	"context"
	"net/http"

	"github.com/vmware/govmomi/vapi/namespace/internal"
)

// NamespaceInstanceSummaryV2
// https://developer.broadcom.com/xapis/vsphere-automation-api/9.0/data-structures/Vcenter%20Namespaces%20Instances%20SummaryV2
// Since 8.0.0.1
type NamespaceInstanceSummaryV2 struct {
	Supervisor           string `json:"supervisor"`
	Namespace            string `json:"namespace"`
	Description          string `json:"description"`
	ConfigStatus         string `json:"config_status"`
	Stats                Stats  `json:"stats"`
	SelfServiceNamespace bool   `json:"self_service_namespace"`
}

// NamespaceInstanceInfoV2
// https://developer.broadcom.com/xapis/vsphere-automation-api/9.0/data-structures/Vcenter%20Namespaces%20Instances%20InfoV2
// Since 8.0.0.1
type NamespaceInstanceInfoV2 struct {
	Supervisor           string             `json:"supervisor"`
	ConfigStatus         string             `json:"config_status"`
	Stats                Stats              `json:"stats"`
	Description          string             `json:"description"`
	StorageSpecs         []StorageSpec      `json:"storage_specs"`
	VmServiceSpec        VmServiceSpec      `json:"vm_service_spec"`
	ContentLibraries     []ContentLibraryV2 `json:"content_libraries"`
	SelfServiceNamespace bool               `json:"self_service_namespace"`
	NetworkSpec          *NetworkConfigInfo `json:"network_spec,omitempty"`
}

// NamespaceInstanceCreateSpecV2
// https://developer.broadcom.com/xapis/vsphere-automation-api/9.0/data-structures/Vcenter%20Namespaces%20Instances%20CreateSpecV2
// Since 8.0.0.1
type NamespaceInstanceCreateSpecV2 struct {
	Namespace            string                      `json:"namespace"`
	Supervisor           string                      `json:"supervisor"`
	Description          *string                     `json:"description,omitempty"`
	StorageSpecs         *[]StorageSpec              `json:"storage_specs,omitempty"`
	VmServiceSpec        *VmServiceSpec              `json:"vm_service_spec,omitempty"`
	ContentLibraries     *[]ContentLibraryV2         `json:"content_libraries,omitempty"`
	SelfServiceNamespace *bool                       `json:"self_service_namespace,omitempty"`
	NetworkSpec          *NetworkConfigCreateSpec    `json:"network_spec,omitempty"`
	NamespaceNetwork     *NamespaceNetworkCreateSpec `json:"namespace_network,omitempty"`
}

type Stats struct {
	CPUUsed     int `json:"cpu_used"`
	MemoryUsed  int `json:"memory_used"`
	StorageUsed int `json:"storage_used"`
}

type ContentLibraryV2 struct {
	ContentLibrary         string `json:"content_library"`
	Writable               bool   `json:"writable"`
	AllowImport            bool   `json:"allow_import"`
	ResourceNamingStrategy string `json:"resource_naming_strategy"`
}

// NamespaceNetworkCreateSpec represents the complete namespace_network field for creation
type NamespaceNetworkCreateSpec struct {
	NetworkProvider string                  `json:"network_provider"`
	Network         *NamespaceNetworkConfig `json:"network,omitempty"`
}

// NamespaceNetworkConfig represents the network configuration within namespace_network
type NamespaceNetworkConfig struct {
	NamespaceNetworkCidrs []CidrBlock `json:"namespace_network_cidrs"`
	IngressCidrs          []CidrBlock `json:"ingress_cidrs"`
	EgressCidrs           []CidrBlock `json:"egress_cidrs"`
	NsxTier0Gateway       string      `json:"nsx_tier0_gateway"`
	SubnetPrefixLength    int         `json:"subnet_prefix_length"`
	RoutedMode            bool        `json:"routed_mode"`
	LoadBalancerSize      string      `json:"load_balancer_size"`
}

// CidrBlock represents a CIDR block with address and prefix
type CidrBlock struct {
	Address string `json:"address"`
	Prefix  int    `json:"prefix"`
}

// NetworkConfigInfo
// https://developer.broadcom.com/xapis/vsphere-automation-api/9.0/data-structures/Vcenter%20Namespaces%20Instances%20NetworkConfigInfo/
// Since 9.0.0.0
type NetworkConfigInfo struct {
	NetworkProvider string          `json:"network_provider"`
	VpcNetwork      *VpcNetworkInfo `json:"vpc_network,omitempty"`
}

// VpcNetworkInfo
// https://developer.broadcom.com/xapis/vsphere-automation-api/9.0/data-structures/Vcenter%20Namespaces%20Instances%20VpcNetworkInfo/
// Since 9.0.0.0
type VpcNetworkInfo struct {
	VpcConfig         VpcConfig `json:"vpc_config"`
	Vpc               string    `json:"vpc"`
	AutoCreated       bool      `json:"auto_created"`
	DefaultSubnetSize int64     `json:"default_subnet_size"`
}

// VpcConfig
// https://developer.broadcom.com/xapis/vsphere-automation-api/9.0/data-structures/Vcenter%20Namespaces%20Instances%20VpcConfig/
// Since 9.0.0.0
type VpcConfig struct {
	PrivateCidrs []Ipv4Cidr `json:"private_cidrs,omitempty"`
}

// NetworkConfigCreateSpec
// https://developer.broadcom.com/xapis/vsphere-automation-api/9.0/data-structures/Vcenter%20Namespaces%20Instances%20NetworkConfigCreateSpec/
// Since 9.0.0.0
type NetworkConfigCreateSpec struct {
	NetworkProvider string                `json:"network_provider"`
	VpcNetwork      *VpcNetworkCreateSpec `json:"vpc_network,omitempty"`
}

// VpcNetworkCreateSpec
// https://developer.broadcom.com/xapis/vsphere-automation-api/9.0/data-structures/Vcenter%20Namespaces%20Instances%20VpcNetworkCreateSpec/
// Since 9.0.0.0
type VpcNetworkCreateSpec struct {
	VpcConfig         *VpcConfig `json:"vpc_config,omitempty"`
	Vpc               *string    `json:"vpc,omitempty"`
	DefaultSubnetSize *int64     `json:"default_subnet_size,omitempty"`
}

// ListNamespacesV2 https://developer.broadcom.com/xapis/vsphere-automation-api/9.0/api/vcenter/namespaces/instances/v2/get/
func (c *Manager) ListNamespacesV2(ctx context.Context) ([]NamespaceInstanceSummaryV2, error) {
	resource := c.Resource(internal.NamespacesPathV2)
	request := resource.Request(http.MethodGet)
	var result []NamespaceInstanceSummaryV2
	return result, c.Do(ctx, request, &result)
}

// GetNamespaceV2 https://developer.broadcom.com/xapis/vsphere-automation-api/9.0/api/vcenter/namespaces/instances/v2/namespace/get/
func (c *Manager) GetNamespaceV2(ctx context.Context, namespace string) (NamespaceInstanceInfoV2, error) {
	resource := c.Resource(internal.NamespacesPathV2).WithSubpath(namespace)
	request := resource.Request(http.MethodGet)
	var result NamespaceInstanceInfoV2
	return result, c.Do(ctx, request, &result)
}

// CreateNamespaceV2 https://developer.broadcom.com/xapis/vsphere-automation-api/9.0/api/vcenter/namespaces/instances/v2/post/
func (c *Manager) CreateNamespaceV2(ctx context.Context, spec NamespaceInstanceCreateSpecV2) error {
	resource := c.Resource(internal.NamespacesPathV2)
	request := resource.Request(http.MethodPost, spec)
	return c.Do(ctx, request, nil)
}
