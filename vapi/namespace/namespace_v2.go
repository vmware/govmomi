// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package namespace

import (
	"context"
	"net/http"

	"github.com/vmware/govmomi/vapi/namespace/internal"
)

// NamespaceInstanceConfigStatus
// https://developer.broadcom.com/xapis/vsphere-automation-api/9.1.1/data-structures/Vcenter%20Namespaces%20Instances%20ConfigStatus/
// Since 7.0.0
type NamespaceInstanceConfigStatus string

const (
	NamespaceInstanceConfigStatusConfiguring NamespaceInstanceConfigStatus = "CONFIGURING"
	NamespaceInstanceConfigStatusRemoving    NamespaceInstanceConfigStatus = "REMOVING"
	NamespaceInstanceConfigStatusRunning     NamespaceInstanceConfigStatus = "RUNNING"
	NamespaceInstanceConfigStatusError       NamespaceInstanceConfigStatus = "ERROR"
)

// NamespaceInstanceSummaryV2
// https://developer.broadcom.com/xapis/vsphere-automation-api/9.0/data-structures/Vcenter%20Namespaces%20Instances%20SummaryV2
// Since 8.0.0.1
type NamespaceInstanceSummaryV2 struct {
	Supervisor           string                        `json:"supervisor"`
	Namespace            string                        `json:"namespace"`
	Description          string                        `json:"description"`
	ConfigStatus         NamespaceInstanceConfigStatus `json:"config_status"`
	Stats                Stats                         `json:"stats"`
	SelfServiceNamespace bool                          `json:"self_service_namespace"`
}

// NamespaceInstanceMessage
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/data-structures/Vcenter%20Namespaces%20Instances%20Message/
type NamespaceInstanceMessage struct {
	Severity string                          `json:"severity"`
	Details  NamespaceInstanceMessageDetails `json:"details"`
}

// NamespaceInstanceMessageDetails
type NamespaceInstanceMessageDetails struct {
	ID             string `json:"id"`
	DefaultMessage string `json:"default_message"`
	Localized      string `json:"localized"`
}

// NamespaceInstanceInfoV2
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/data-structures/Vcenter%20Namespaces%20Instances%20InfoV2/
// Since 8.0.0.1
type NamespaceInstanceInfoV2 struct {
	Supervisor           string                        `json:"supervisor"`
	ConfigStatus         NamespaceInstanceConfigStatus `json:"config_status"`
	Stats                Stats                         `json:"stats"`
	Description          string                        `json:"description"`
	StorageSpecs         []StorageSpec                 `json:"storage_specs"`
	VmServiceSpec        VmServiceSpec                 `json:"vm_service_spec"`
	ContentLibraries     []ContentLibraryV2            `json:"content_libraries"`
	SelfServiceNamespace bool                          `json:"self_service_namespace"`
	NetworkSpec          *NetworkConfigInfo            `json:"network_spec,omitempty"`
	Edges                []NamespaceEdgeSpec           `json:"edges,omitempty"`
	Zones                []ZoneInfo                    `json:"zones,omitempty"`
	AccessList           []Access                      `json:"access_list"`
	Messages             []NamespaceInstanceMessage    `json:"messages"`
}

// ZoneInfo
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/data-structures/Vcenter%20Namespaces%20Instances%20ZoneInfo/
// Since 9.0.0.0
type ZoneInfo struct {
	Zone             string   `json:"zone"`
	Type             string   `json:"type"`
	MarkedForRemoval bool     `json:"marked_for_removal"`
	Clusters         []string `json:"clusters,omitempty"`
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
	Edges                *[]NamespaceEdgeSpec        `json:"edges,omitempty"`
	Zones                *[]ZoneSpec                 `json:"zones,omitempty"`
	AccessList           *[]Access                   `json:"access_list,omitempty"`
	Networks             *[]string                   `json:"networks,omitempty"`
}

// ZoneSpec
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/data-structures/Vcenter%20Namespaces%20Instances%20ZoneSpec/
// Since 9.0.0.0
type ZoneSpec struct {
	Name                              string                               `json:"name"`
	VmReservations                    *[]VirtualMachineClassAllocationSpec `json:"vm_reservations,omitempty"`
	CpuLimitMHz                       *int64                               `json:"cpu_limit_MHz,omitempty"`
	CpuReservationMHz                 *int64                               `json:"cpu_reservation_MHz,omitempty"`
	MemoryLimitMiB                    *int64                               `json:"memory_limit_MiB,omitempty"`
	MemoryReservationMiB              *int64                               `json:"memory_reservation_MiB,omitempty"`
	DisallowUnreservedDirectPathUsage *bool                                `json:"disallow_unreserved_direct_path_usage,omitempty"`
}

// VirtualMachineClassAllocationSpec
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/data-structures/Vcenter%20Namespaces%20Instances%20VirtualMachineClassAllocationSpec/
// Since 9.0.0.0
type VirtualMachineClassAllocationSpec struct {
	ReservedVmClass string `json:"reserved_vm_class"`
	Count           int64  `json:"count"`
}

// Access
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/data-structures/Vcenter%20Namespaces%20Instances%20Access/
// Since 7.0.0.0
type Access struct {
	SubjectType      AccessSubjectType `json:"subject_type"`
	Subject          string            `json:"subject"`
	Domain           string            `json:"domain"`
	Role             AccessRole        `json:"role"`
	IdentityProvider *string           `json:"identity_provider,omitempty"`
}

// AccessSubjectType
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/data-structures/Vcenter%20Namespaces%20Access%20SubjectType/
// Since 7.0.0.0
type AccessSubjectType string

const (
	AccessSubjectTypeUser  AccessSubjectType = "USER"
	AccessSubjectTypeGroup AccessSubjectType = "GROUP"
)

// AccessRole
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/data-structures/Vcenter%20Namespaces%20Access%20Role/
// Since 7.0.0.0
type AccessRole string

const (
	AccessRoleOwner AccessRole = "OWNER"
	AccessRoleEdit  AccessRole = "EDIT"
	AccessRoleView  AccessRole = "VIEW"
)

// SharedSubnet
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/data-structures/Vcenter%20Namespaces%20Instances%20SharedSubnet/
// Since 9.1.0.0
type SharedSubnet struct {
	Path       string  `json:"path"`
	PodDefault *bool   `json:"pod_default,omitempty"`
	VmDefault  *bool   `json:"vm_default,omitempty"`
	Name       *string `json:"name,omitempty"`
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

// NamespaceEdgeSpec The Vcenter Namespaces Instances EdgeCreateSpec schema contains the specification required to configure Edge provider association with a namespace.
// https://developer.broadcom.com/xapis/vsphere-automation-api/9.1.1/data-structures/Vcenter%20Namespaces%20Instances%20EdgeCreateSpec/
// Since 9.1.0.0
type NamespaceEdgeSpec struct {
	// Retrieve ID from api
	// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/api/vcenter/namespace-management/supervisors/supervisor/networks/edges/get/
	ID string `json:"id,omitempty"`
	// Vcenter Namespaces Instances EdgeProvider
	// https://developer.broadcom.com/xapis/vsphere-automation-api/9.1.1/data-structures/Vcenter%20Namespaces%20Instances%20EdgeProvider/
	EdgeProvider string                `json:"edge_provider,omitempty"`
	Avi          *NamespaceEdgeAviSpec `json:"avi,omitempty"`
}

// NamespaceEdgeAviSpec defines the SE-Group used by the namespace to provision LB
type NamespaceEdgeAviSpec struct {
	SeGroupName string `json:"se_group_name,omitempty"`
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
	NetworkProvider string                    `json:"network_provider"`
	VpcNetwork      *VpcNetworkInfo           `json:"vpc_network,omitempty"`
	VsphereNetwork  *VsphereNetworkConfigInfo `json:"vsphere_network,omitempty"`
}

// VsphereNetworkConfigInfo
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/data-structures/Vcenter%20Namespaces%20Instances%20VsphereNetworkConfigInfo/
// Since 9.1.0.0
type VsphereNetworkConfigInfo struct {
	DefaultWorkloadNetwork string                 `json:"default_workload_network"`
	Networks               []NamespaceNetworkInfo `json:"networks"`
}

// NamespaceNetworkInfo
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/data-structures/Vcenter%20Namespaces%20Instances%20NetworkInfo/
// Since 9.1.0.0
type NamespaceNetworkInfo struct {
	Network      string              `json:"network"`
	ConfigStatus NetworkConfigStatus `json:"config_status"`
}

// NetworkConfigStatus
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/data-structures/Vcenter%20Namespaces%20Instances%20NetworkConfigStatus/
// Since 9.1.0.0
type NetworkConfigStatus string

const (
	NetworkConfigStatusActive   NetworkConfigStatus = "ACTIVE"
	NetworkConfigStatusRemoving NetworkConfigStatus = "REMOVING"
)

// VpcNetworkInfo
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/data-structures/Vcenter%20Namespaces%20Instances%20VpcNetworkInfo/
// Since 9.0.0.0
type VpcNetworkInfo struct {
	VpcConfig         VpcConfig      `json:"vpc_config"`
	Vpc               string         `json:"vpc"`
	SharedSubnets     []SharedSubnet `json:"shared_subnets,omitempty"`
	AutoCreated       bool           `json:"auto_created"`
	DefaultSubnetSize int64          `json:"default_subnet_size"`
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
	NetworkProvider string                          `json:"network_provider"`
	VpcNetwork      *VpcNetworkCreateSpec           `json:"vpc_network,omitempty"`
	VsphereNetwork  *VsphereNetworkConfigCreateSpec `json:"vsphere_network,omitempty"`
}

// VsphereNetworkConfigCreateSpec
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/data-structures/Vcenter%20Namespaces%20Instances%20VsphereNetworkConfigCreateSpec/
// Since 9.1.0.0
type VsphereNetworkConfigCreateSpec struct {
	DefaultWorkloadNetwork *string `json:"default_workload_network,omitempty"`
}

// VpcNetworkCreateSpec
// https://developer.broadcom.com/xapis/vsphere-automation-api/9.1.1/data-structures/Vcenter%20Namespaces%20Instances%20VpcNetworkCreateSpec/
// Since 9.0.0.0
type VpcNetworkCreateSpec struct {
	VpcConfig         *VpcConfig      `json:"vpc_config,omitempty"`
	Vpc               *string         `json:"vpc,omitempty"`
	SharedSubnets     *[]SharedSubnet `json:"shared_subnets,omitempty"`
	DefaultSubnetSize *int64          `json:"default_subnet_size,omitempty"`
}

// ListNamespacesV2 https://developer.broadcom.com/xapis/vsphere-automation-api/9.0/api/vcenter/namespaces/instances/v2/get/
func (c *Manager) ListNamespacesV2(ctx context.Context) ([]NamespaceInstanceSummaryV2, error) {
	resource := c.Resource(internal.NamespacesPathV2)
	request := resource.Request(http.MethodGet)
	var result []NamespaceInstanceSummaryV2
	return result, c.Do(ctx, request, &result)
}

// GetNamespaceV2 https://developer.broadcom.com/xapis/vsphere-automation-api/latest/api/vcenter/namespaces/instances/v2/namespace/get/
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
