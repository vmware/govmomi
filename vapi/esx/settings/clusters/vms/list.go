// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vms

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/vmware/govmomi/vim25/types"
)

// ClusterSolutionInfo contains fields that describe solution configuration
// only applicable for solutions with deployment type DeploymentType#CLUSTER_VM_SET}.
type ClusterSolutionInfo struct {
	// The number of instances of the specified VM to be deployed across the
	// cluster.
	VmCount int `json:"vm_count"`

	// VmPlacementPolicies  are the VM placement policies to be configured on the VMs.
	VmPlacementPolicies []VmPlacementPolicy `json:"vm_placement_policies"`

	// VmNetworks to be configured on the VMs. The map keys represent the
	// logical network names defined in the OVF descriptor while the map values
	// represent the VM network identifiers.
	//
	// If no VM networks are configured.
	VmNetworks map[string]string `json:"vm_networks"`

	// VmDatastores to be configured as a storage of the VMs. The first datastore
	// from the list available in the cluster is used.
	//
	// If unset the system automatically selects the datastore. The selection
	// takes into account the other properties of the desired state
	// specification (the provided VM storage policies and VM devices) and
	// the runtime state of the datastores in the cluster. It is required
	// DRS to be enabled on the cluster.
	VmDatastores *[]string `json:"vm_datastores,omitempty"`

	// Devices of the VMs not defined in the OVF descriptor. If VmDatastores is
	// not set, the devices of the VMs not defined in the OVF descriptor should
	// be provided and not as part of a VM lifecycle hook (VM reconfiguration).
	// Otherwise, the compatibility of the devices with the selected host and
	// datastore where the VM is deployed needs to be ensured by the client.
	//
	// 1. For VM initial placement the devices are added to the VM
	// configuration. 2. For the reconfiguration it is checked what devices
	// need to be added, removed, and edited on the existing VMs. NOTE: No VM
	// relocation is executed before the VM reconfiguration.
	//
	// The supported property of vim.vm.ConfigSpec is
	// vim.vm.ConfigSpec.deviceChange. The supported
	// vim.vm.device.VirtualDeviceSpec.operation is Operation#add. For
	// vim.vm.device.VirtualEthernetCard the unique identifier is
	// vim.vm.device.VirtualDevice#unitNumber. VmNetworks and Devices are
	// mutually exclusive.
	//
	// If unset no additional devices will be added to
	// the VMs.
	// Optional<DynamicStructure> devices;

	// RemediationPolicy to be configured for the deployment units.
	RemediationPolicy RemediationPolicy `json:"remediation_policy"`

	// AlternativeVmSpecs to be applied on the System VMs.
	//
	// If unset no AlternativeVmSpecs applied to the System VMs.
	AlternativeVmSpecs []AlternativeVmSpec `json:"alternative_vm_specs,omitempty"`

	// Devices of the VMs not defined in the OVF descriptor. If VmDatastores is
	// not set, the devices of the VMs not defined in the OVF descriptor should
	// be provided and not as part of a VM lifecycle hook (VM reconfiguration).
	// Otherwise, the compatibility of the devices with the selected host and
	// datastore where the VM is deployed needs to be ensured by the client.
	Devices json.RawMessage `json:"devices,omitempty"`
}

type SolutionInfo struct {
	// DeploymentType of the solution
	DeploymentType DeploymentType `json:"deployment_type"`

	// DisplayName is the display name of the solution.
	DisplayName string `json:"display_name"`

	// DisplayVersion is the display version of the solution.
	DisplayVersion string `json:"display_version"`

	// VmNameTemplate is the VM name template.
	VmNameTemplate VmNameTemplate `json:"vm_name_template"`

	// ClusterSolutionInfo is the configuration that is only applicable for
	// solutions with deployment type ClusterVmSET.
	ClusterSolutionInfo ClusterSolutionInfo `json:"cluster_solution_info"`

	// HookConfigurations keys represent LifecycleStates while the map values
	// represent their configurations.
	HookConfigurations map[LifecycleState]LifecycleHookConfig `json:"hook_configurations"`

	// OvfResource is information about the OVF resource that to be used for
	// the VM deployments.
	OvfResource OvfResource `json:"ovf_resource"`

	//  OvfDescriptorProperties are the OVF properties that to be assigned to
	//  the VMs' OVF properties when powered on. The keys of the map must not
	//  include any white-space characters. The map keys represent the names of
	//  properties while the map values represent the values of those
	//  properties.
	OvfDescriptorProperties map[string]string `json:"ovf_descriptor_properties"`

	// VmCloneConfig is the VM cloning configuration.
	VmCloneConfig VmCloneConfig `json:"vm_clone_config"`

	// Storage policies to be configured on the VMs.
	VmStoragePolicy StoragePolicy `json:"vm_storage_policy"`

	// Storage policy profiles to be configured on the VMs. The profiles are
	// passed to vim.vm.ConfigSpec#vmProfile without any interpretation.
	VmStorageProfiles []string `json:"vm_storage_profiles"`

	VmDiskType DiskType `json:"vm_disk_type"`

	VmResourcePool string `json:"vm_resource_pool"`

	VmFolder string `json:"vm_folder"`

	// VmResourceSpec is the VMs resource configuration.
	//
	// If unset the default resource configuration specified in the OVF
	// descriptor is used.
	VmResourceSpec *VmResourceSpec `json:"vm_resource_spec,omitempty"`

	// Specifies System VMs redeployment policy.
	RedeploymentPolicy RedeploymentPolicy `json:"redeployment_policy"`
}

// ListResult contains fields that describe the desired specification of the
// solutions in a given cluster specified by the FilterSpec of the
// corresponding list operation.
type ListResult struct {
	Solutions map[string]SolutionInfo `json:"solutions"`
}

func (m *Manager) List(ctx context.Context, cluster types.ManagedObjectReference) (*ListResult, error) {
	var r ListResult
	p := clusterSolutionPath(cluster).String()
	url := m.Resource(p)
	return &r, m.Do(ctx, url.Request(http.MethodGet), &r)
}
