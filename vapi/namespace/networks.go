// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package namespace

// Package: vapi/namespace/networks.go
//
// Implements all vSphere Namespaces Networks API operations for a Supervisor Cluster:
//
//   GET    /vcenter/namespace-management/clusters/{cluster}/networks
//   POST   /vcenter/namespace-management/clusters/{cluster}/networks
//   GET    /vcenter/namespace-management/clusters/{cluster}/networks/{network}
//   PATCH  /vcenter/namespace-management/clusters/{cluster}/networks/{network}
//   PUT    /vcenter/namespace-management/clusters/{cluster}/networks/{network}
//   DELETE /vcenter/namespace-management/clusters/{cluster}/networks/{network}
//
// API reference:
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/vcenter-namespace-management/vcenter-namespace-management-networks/

import (
	"context"
	"fmt"
	"net/http"

	"github.com/vmware/govmomi/vapi/namespace/internal"
)

// LoadBalancerSize enumerates the NSX load balancer sizes.
type LoadBalancerSize string

const (
	LoadBalancerSizeSmall  LoadBalancerSize = "SMALL"
	LoadBalancerSizeMedium LoadBalancerSize = "MEDIUM"
	LoadBalancerSizeLarge  LoadBalancerSize = "LARGE"
)

// IPAssignmentMode enumerates IP assignment modes for vSphere networks.
type IPAssignmentMode string

const (
	IPAssignmentModeDHCP        IPAssignmentMode = "DHCP"
	IPAssignmentModeStaticRange IPAssignmentMode = "STATICRANGE"
)

// -------- NSX-T sub-types --------

// NsxNetworkInfo describes the runtime state of an NSX-T-backed network.
// Returned by GET (list and get single).
// See: https://developer.broadcom.com/xapis/vsphere-automation-api/latest/data-structures/Vcenter%20NamespaceManagement%20Networks%20NsxNetworkInfo/
type NsxNetworkInfo struct {
	EgressCidrs           []CidrBlock      `json:"egress_cidrs,omitempty"`
	IngressCidrs          []CidrBlock      `json:"ingress_cidrs,omitempty"`
	LoadBalancerSize      LoadBalancerSize `json:"load_balancer_size,omitempty"`
	NamespaceNetworkCidrs []CidrBlock      `json:"namespace_network_cidrs,omitempty"`
	NsxTier0Gateway       string           `json:"nsx_tier0_gateway,omitempty"`
	RoutedMode            *bool            `json:"routed_mode,omitempty"`
	SubnetPrefixLength    int              `json:"subnet_prefix_length,omitempty"`
}

// NsxNetworkCreateSpec is the create spec for an NSX-T-backed network.
// Used in NetworkCreateSpec when NetworkProvider == NSXT_CONTAINER_PLUGIN.
// See: https://developer.broadcom.com/xapis/vsphere-automation-api/latest/data-structures/Vcenter%20NamespaceManagement%20Networks%20NsxNetworkCreateSpec/
type NsxNetworkCreateSpec struct {
	EgressCidrs           []CidrBlock      `json:"egress_cidrs,omitempty"`
	IngressCidrs          []CidrBlock      `json:"ingress_cidrs,omitempty"`
	LoadBalancerSize      LoadBalancerSize `json:"load_balancer_size,omitempty"`
	NamespaceNetworkCidrs []CidrBlock      `json:"namespace_network_cidrs,omitempty"`
	NsxTier0Gateway       string           `json:"nsx_tier0_gateway,omitempty"`
	RoutedMode            *bool            `json:"routed_mode,omitempty"`
	SubnetPrefixLength    int              `json:"subnet_prefix_length,omitempty"`
}

// NsxNetworkUpdateSpec is the partial-update spec for an NSX-T-backed network.
// Used in NetworkUpdateSpec (PATCH).
// See: https://developer.broadcom.com/xapis/vsphere-automation-api/latest/data-structures/Vcenter%20NamespaceManagement%20Networks%20NsxNetworkUpdateSpec/
type NsxNetworkUpdateSpec struct {
	EgressCidrs           []CidrBlock `json:"egress_cidrs,omitempty"`
	IngressCidrs          []CidrBlock `json:"ingress_cidrs,omitempty"`
	NamespaceNetworkCidrs []CidrBlock `json:"namespace_network_cidrs,omitempty"`
}

// NsxNetworkSetSpec is the full-replacement spec for an NSX-T-backed network.
// Used in NetworkSetSpec (PUT).
// See: https://developer.broadcom.com/xapis/vsphere-automation-api/latest/data-structures/Vcenter%20NamespaceManagement%20Networks%20NsxNetworkSetSpec/
type NsxNetworkSetSpec struct {
	EgressCidrs           []CidrBlock `json:"egress_cidrs,omitempty"`
	IngressCidrs          []CidrBlock `json:"ingress_cidrs,omitempty"`
	NamespaceNetworkCidrs []CidrBlock `json:"namespace_network_cidrs,omitempty"`
}

// -------- vSphere DVPG sub-types --------

// VsphereNetworkInfo describes the runtime state of a DVPG-backed network.
// Returned by GET (list and get single).
// See: https://developer.broadcom.com/xapis/vsphere-automation-api/latest/data-structures/Vcenter%20Namespaces%20Instances%20VsphereNetworkConfigInfo/
type VsphereNetworkInfo struct {
	AddressRanges    []IpRange        `json:"address_ranges,omitempty"`
	Gateway          string           `json:"gateway,omitempty"`
	IPAssignmentMode IPAssignmentMode `json:"ip_assignment_mode,omitempty"`
	Portgroup        string           `json:"portgroup,omitempty"`
	SubnetMask       string           `json:"subnet_mask,omitempty"`
}

// VsphereNetworkCreateSpec is the create spec for a DVPG-backed network.
// Used in NetworkCreateSpec when NetworkProvider == VSPHERE_NETWORK.
// See: https://developer.broadcom.com/xapis/vsphere-automation-api/latest/data-structures/Vcenter%20Namespaces%20Instances%20VsphereNetworkConfigCreateSpec/
type VsphereNetworkCreateSpec struct {
	AddressRanges    []IpRange        `json:"address_ranges,omitempty"`
	Gateway          string           `json:"gateway,omitempty"`
	IPAssignmentMode IPAssignmentMode `json:"ip_assignment_mode,omitempty"`
	Portgroup        string           `json:"portgroup"`
	SubnetMask       string           `json:"subnet_mask,omitempty"`
}

// VsphereDVPGNetworkSetSpec is the full-replacement spec for a DVPG-backed network.
// Used in NetworkSetSpec (PUT) and also serves as partial update for PATCH.
// See: https://developer.broadcom.com/xapis/vsphere-automation-api/latest/data-structures/Vcenter%20NamespaceManagement%20Networks%20VsphereDVPGNetworkSetSpec/
type VsphereDVPGNetworkSetSpec struct {
	AddressRanges []IpRange `json:"address_ranges,omitempty"`
	Gateway       string    `json:"gateway,omitempty"`
	Portgroup     string    `json:"portgroup,omitempty"`
	SubnetMask    string    `json:"subnet_mask,omitempty"`
}

// -------- Top-level request / response types --------

// NetworkInfo is the detailed view returned by the Get (GET single) endpoint.
// See: https://developer.broadcom.com/xapis/vsphere-automation-api/latest/data-structures/Vcenter%20Namespaces%20Instances%20NetworkInfo/
type NetworkInfo struct {
	Network         string              `json:"network"`
	NetworkProvider NetworkProvider     `json:"network_provider"`
	Namespaces      []string            `json:"namespaces,omitempty"`
	NsxNetwork      *NsxNetworkInfo     `json:"nsx_network,omitempty"`
	VsphereNetwork  *VsphereNetworkInfo `json:"vsphere_network,omitempty"`
}

// NetworkCreateSpec is the request body for POST (create).
// See: https://developer.broadcom.com/xapis/vsphere-automation-api/latest/data-structures/Vcenter%20Namespaces%20Instances%20NetworkCreateSpec/
type NetworkCreateSpec struct {
	// Network is the desired identifier (DNS_LABEL, max 63 chars, alphanumeric + '-').
	Network         string                    `json:"network"`
	NetworkProvider NetworkProvider           `json:"network_provider"`
	NsxNetwork      *NsxNetworkCreateSpec     `json:"nsx_network,omitempty"`
	VsphereNetwork  *VsphereNetworkCreateSpec `json:"vsphere_network,omitempty"`
}

// NetworkUpdateSpec is the request body for PATCH (partial update).
// See: https://developer.broadcom.com/xapis/vsphere-automation-api/latest/data-structures/Vcenter%20NamespaceManagement%20Networks%20UpdateSpec/
type NetworkUpdateSpec struct {
	NetworkProvider NetworkProvider            `json:"network_provider"`
	NsxNetwork      *NsxNetworkUpdateSpec      `json:"nsx_network,omitempty"`
	VsphereNetwork  *VsphereDVPGNetworkSetSpec `json:"vsphere_network,omitempty"`
}

// NetworkSetSpec is the request body for PUT (full replacement).
// See: https://developer.broadcom.com/xapis/vsphere-automation-api/latest/data-structures/Vcenter%20NamespaceManagement%20Networks%20SetSpec/
type NetworkSetSpec struct {
	NetworkProvider NetworkProvider            `json:"network_provider"`
	NsxNetwork      *NsxNetworkSetSpec         `json:"nsx_network,omitempty"`
	VsphereNetwork  *VsphereDVPGNetworkSetSpec `json:"vsphere_network,omitempty"`
}

// -------- Manager methods --------

// ListClusterNetworks returns all vSphere Namespaces networks for the given cluster.
//
//	GET /vcenter/namespace-management/clusters/{cluster}/networks
func (c *Manager) ListClusterNetworks(ctx context.Context, clusterID string) ([]NetworkInfo, error) {
	var res []NetworkInfo
	url := c.Resource(fmt.Sprintf(internal.NamespaceNetworkPath, clusterID))
	err := c.Do(ctx, url.Request(http.MethodGet), &res)
	return res, err
}

// CreateClusterNetwork creates a new vSphere Namespaces network object for the given cluster.
//
//	POST /vcenter/namespace-management/clusters/{cluster}/networks
//
// Note: NsxNetworkCreateSpec is not supported via this endpoint (unsupported error); use VSPHERE_NETWORK
// or pre-provisioned NSX networks instead.
func (c *Manager) CreateClusterNetwork(ctx context.Context, clusterID string, spec *NetworkCreateSpec) error {
	url := c.Resource(fmt.Sprintf(internal.NamespaceNetworkPath, clusterID))
	return c.Do(ctx, url.Request(http.MethodPost, spec), nil)
}

// GetClusterNetwork returns detailed information about a specific network.
//
//	GET /vcenter/namespace-management/clusters/{cluster}/networks/{network}
func (c *Manager) GetClusterNetwork(ctx context.Context, clusterID, networkID string) (*NetworkInfo, error) {
	var res NetworkInfo
	url := c.Resource(fmt.Sprintf(internal.NamespaceNetworkPath, clusterID)).WithSubpath(networkID)
	err := c.Do(ctx, url.Request(http.MethodGet), &res)
	return &res, err
}

// UpdateClusterNetwork partially updates an existing network (PATCH).
// Only the fields present in spec are applied; omitempty fields are left unchanged.
//
//	PATCH /vcenter/namespace-management/clusters/{cluster}/networks/{network}
func (c *Manager) UpdateClusterNetwork(ctx context.Context, clusterID, networkID string, spec *NetworkUpdateSpec) error {
	url := c.Resource(fmt.Sprintf(internal.NamespaceNetworkPath, clusterID)).WithSubpath(networkID)
	return c.Do(ctx, url.Request(http.MethodPatch, spec), nil)
}

// SetClusterNetwork fully replaces the configuration of an existing network (PUT).
// All mutable fields must be supplied; any omitted fields are reset to defaults.
//
//	PUT /vcenter/namespace-management/clusters/{cluster}/networks/{network}
func (c *Manager) SetClusterNetwork(ctx context.Context, clusterID, networkID string, spec *NetworkSetSpec) error {
	url := c.Resource(fmt.Sprintf(internal.NamespaceNetworkPath, clusterID)).WithSubpath(networkID)
	return c.Do(ctx, url.Request(http.MethodPut, spec), nil)
}

// DeleteClusterNetwork removes a vSphere Namespaces network object from the given cluster.
// The network must not be referenced by any namespace before deletion.
//
//	DELETE /vcenter/namespace-management/clusters/{cluster}/networks/{network}
func (c *Manager) DeleteClusterNetwork(ctx context.Context, clusterID, networkID string) error {
	url := c.Resource(fmt.Sprintf(internal.NamespaceNetworkPath, clusterID)).WithSubpath(networkID)
	return c.Do(ctx, url.Request(http.MethodDelete), nil)
}
