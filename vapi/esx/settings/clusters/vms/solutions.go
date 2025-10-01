// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vms

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vmware/govmomi/vapi/cis/tasks"
	"github.com/vmware/govmomi/vapi/esx/settings/clusters"
	"github.com/vmware/govmomi/vapi/rest"
	"github.com/vmware/govmomi/vim25/types"
)

type clusterSolutionPath types.ManagedObjectReference

const (
	SolutionsPath = clusters.BasePath + "/%s/vms/solutions"
)

func (c clusterSolutionPath) String() string {
	cid := types.ManagedObjectReference(c).Value
	return fmt.Sprintf(SolutionsPath, cid)
}

// VmPlacementPolicy defines the DRS placement policies applied on the VMs.
type VmPlacementPolicy string

const (
	// VmVmAntiAffinity defines VMs are anti-affined to each other.
	VmVmAntiAffinity VmPlacementPolicy = "VM_VM_ANTI_AFFINITY"
)

// RemediationPolicy defines the remediation policies applied to entities.
type RemediationPolicy string

const (
	// Parallel is the default remediation policy. Entities are remediated in parallel.
	Parallel RemediationPolicy = "PARALLEL"

	// Sequential policy is where entities are remediated sequentially, one at a time.
	Sequential RemediationPolicy = "SEQUENTIAL"
)

// AlternativeVmSpec contains to describe alternative VM configuration to be applied on
// VMs that matches a given selection criteria defined as VmSelectionSpec.
type AlternativeVmSpec struct {
	// SelectionCriteria is the criteria to match System VMs for which to apply this alternative VM spec.
	SelectionCriteria VmSelectionSpec `json:"selection_criteria"`

	// Devices of the VMs not defined in the OVF descriptor. This
	// takes precedence over ClusterSolutionSpec#Devices.
	//
	// If ClusterSolutionSpec#VmDatastores is not set, the devices of
	// the VMs not defined in the OVF descriptor should be provided to
	// Devices and not as part of a VM lifecycle hook
	// (VM reconfiguration). Otherwise, the compatibility of the devices with the
	// selected host and datastore where the VM is deployed needs to be ensured
	// by the client.
	//
	// 1. For VM initial placement the devices are added to the VM configuration.
	// 2. For the reconfiguration it is checked what devices need to be added,
	// removed, and edited on the existing VMs. NOTE: No VM relocation is
	// executed before the VM reconfiguration.
	//
	// The supported property of vim.vm.ConfigSpec is
	// vim.vm.ConfigSpec.deviceChange. The supported
	// vim.vm.device.VirtualDeviceSpec.operation is Operation#add. For
	// vim.vm.device.VirtualEthernetCard the unique identifier is
	// vim.vm.device.VirtualDevice#unitNumber.
	// ClusterSolutionSpec#vmNetworks and devices are mutually
	// exclusive.
	//
	// If unset, ClusterSolutionSpec#Devices is used.
	Devices json.RawMessage `json:"devices,omitempty"`
}

// ClusterSolutionSpec contains fields that describe solution configuration
// only applicable for solutions with deployment type DeploymentType#CLUSTER_VM_SET}.
type ClusterSolutionSpec struct {
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
	Devices json.RawMessage `json:"devices,omitempty"`

	// RemediationPolicy to be configured for the deployment units.
	RemediationPolicy RemediationPolicy `json:"remediation_policy"`

	// AlternativeVmSpecs to be applied on the System VMs.
	// If unset no AlternativeVmSpecs applied to the System VMs.
	AlternativeVmSpecs []AlternativeVmSpec `json:"alternative_vm_specs,omitempty"`
}

type DeploymentType string

const (
	EveryHostPinned DeploymentType = "EVERY_HOST_PINNED"
	ClusterVmSet    DeploymentType = "CLUSTER_VM_SET"
)

// SuffixFormat  defines the types of VM name suffixes.
type SuffixFormat string

const (
	// UUID suffix format.
	Uuid SuffixFormat = "UUID"

	// Suffix in the format "(counter)" where "counter" is monotonically
	// growing integer.
	Counter SuffixFormat = "COUNTER"
)

// VmNameTemplate  contains  that describe a template for VM names.
type VmNameTemplate struct {
	// Prefix is the  VM name prefix.
	Prefix string `json:"prefix"`

	// Suffix is VM name suffix format.
	Suffix SuffixFormat `json:"suffix"`
}

type StoragePolicy string

const (
	Default StoragePolicy = "DEFAULT"
	Profile StoragePolicy = "PROFILE"
)

type DiskType string

const (
	DiskTypeDefault DiskType = "DEFAULT"
	DiskTypeThin    DiskType = "THIN"
	DiskTypeThick   DiskType = "THICK"
)

// RedeploymentPolicy defines the different remediation policies which require redeployment of the System VMs.
type RedeploymentPolicy string

const (
	// RECREATE is the default policy used by vLCM for System VM redeployment.
	// System VMs are redeployed as follows: Once the new replica is
	// provisioned, the old replica is powered off and deleted. Then the new
	// replica is powered on and it's setup is completed to have the System VM
	// fully operational.
	//
	// This policy causes a downtime.
	ReCreate RedeploymentPolicy = "RECREATE"

	// BlueGreen Follows a standard blue-green strategy. System VMs are
	// redeployed as follows: Once the new replica is provisioned, it is
	// powered on. Then the new replica setup is completed to have the System
	// VM fully operational. Then the old replica is powered off and deleted.
	// This policy provides zero-downtime.
	BlueGreen RedeploymentPolicy = "BLUE_GREEN"
)

// VmResourceSpec describes the VM resource configurations.
type VmResourceSpec struct {

	// OvfDeploymentOption corresponds to the Configuration element of the
	// DeploymentOptionSection in the OVF descriptor (e.g. "small", "medium",
	// "large"). If unset the default deployment options as specified in the
	// OVF descriptor is used.
	OvfDeploymentOption *string `json:"ovf_deployment_option,omitempty"`
}

type VmCloneConfig string

const (
	// The system creates a snapshot of the first deployed VM and after that uses
	// one of the available VM clone methods to deploy others.

	AllClones VmCloneConfig = "ALL_CLONES"

	// The system creates a snapshot of the first deployed VM and after that uses
	// full VM clone method to deploy others.
	FullClones VmCloneConfig = "FULL_CLONES_ONLY"

	// The system does not use VM clone methods to deploy VMs.
	NoClones VmCloneConfig = "NO_CLONES"
)

type SolutionSpec struct {
	// DeploymentType of the solution
	DeploymentType DeploymentType `json:"deployment_type"`

	// DisplayName is the display name of the solution.
	DisplayName string `json:"display_name"`

	// DisplayVersion is the display version of the solution.
	DisplayVersion string `json:"display_version"`

	// VmNameTemplate is the VM name template.
	VmNameTemplate VmNameTemplate `json:"vm_name_template"`

	// ClusterSolutionSpec is the configuration that is only applicable for
	// solutions with deployment type ClusterVmSET.
	ClusterSolutionSpec ClusterSolutionSpec `json:"cluster_solution_spec,omitempty"`

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

// Manager extends rest.Client, adding cluster related methods.
type Manager struct {
	*rest.Client
}

// Set sets and overrides the current desired specification for a given solution
// and cluster. The provided desired specification is validated before that.
//
// Parameters:
//   - cluster: Identifier of the cluster.
//   - solution: Identifier of the solution.
//   - spec: Solution specification.
//
// Returns an error if:
//   - There is an unknown internal error. The accompanying error message will give more details about the failure.
//   - The validation of the solution specification fails.
//   - The cluster is not managed by vLCM.
//   - There is no cluster associated with the given identifier.
//   - The service is not available.
//   - The caller is not authenticated.
//   - The user doesn't have the required privileges.
func (m *Manager) Set(ctx context.Context, cluster types.ManagedObjectReference, solution string, spec *SolutionSpec) error {
	p := clusterSolutionPath(cluster).String()
	url := m.Resource(p).WithSubpath(solution).WithParam("vmw-task", "true")
	var taskId string

	if err := m.Do(ctx, url.Request(http.MethodPut, spec), &taskId); err != nil {
		return err
	}

	_, err := tasks.NewManager(m.Client).WaitForCompletion(ctx, taskId)
	return err
}

func (m *Manager) Get(ctx context.Context, cluster types.ManagedObjectReference, solution string) (*SolutionInfo, error) {
	p := clusterSolutionPath(cluster).String()
	url := m.Resource(p).WithSubpath(solution)
	var res SolutionInfo
	if err := m.Do(ctx, url.Request(http.MethodGet), &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (m *Manager) Delete(ctx context.Context, cluster types.ManagedObjectReference, solution string) error {
	p := clusterSolutionPath(cluster).String()
	url := m.Resource(p).WithSubpath(solution).WithParam("vmw-task", "true")
	var taskId string

	if err := m.Do(ctx, url.Request(http.MethodDelete), &taskId); err != nil {
		return err
	}

	_, err := tasks.NewManager(m.Client).WaitForCompletion(ctx, taskId)
	return err
}
