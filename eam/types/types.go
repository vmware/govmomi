// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"reflect"
	"time"

	"github.com/vmware/govmomi/vim25/types"
)

type AddIssue AddIssueRequestType

func init() {
	types.Add("eam:AddIssue", reflect.TypeOf((*AddIssue)(nil)).Elem())
}

// The parameters of `Agency.AddIssue`.
//
// This structure may be used only with operations rendered under `/eam`.
type AddIssueRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// A new issue.
	Issue BaseIssue `xml:"issue,typeattr" json:"issue"`
}

func init() {
	types.Add("eam:AddIssueRequestType", reflect.TypeOf((*AddIssueRequestType)(nil)).Elem())
}

type AddIssueResponse struct {
	Returnval BaseIssue `xml:"returnval,typeattr" json:"returnval"`
}

// Deprecated as of vSphere 9.0. Please refer to vLCM APIs.
//
// Scope specifies on which compute resources to deploy a solution's agents.
//
// This structure may be used only with operations rendered under `/eam`.
type AgencyComputeResourceScope struct {
	AgencyScope

	// Compute resources on which to deploy the agents.
	//
	// If `AgencyConfigInfoEx.vmPlacementPolicy` is set, the array needs to
	// contain exactly one cluster compute resource.
	//
	// Refers instances of `ComputeResource`.
	ComputeResource []types.ManagedObjectReference `xml:"computeResource,omitempty" json:"computeResource,omitempty"`
}

func init() {
	types.Add("eam:AgencyComputeResourceScope", reflect.TypeOf((*AgencyComputeResourceScope)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. Please refer to vLCM APIs.
//
// This is the configuration of an <code>Agency</code>.
//
// It determines on
// which compute resources to deploy the agents, which VIB to install, which
// OVF package to install, and how to configure these items by setting the
// OVF environment properties.
//
// This structure may be used only with operations rendered under `/eam`.
type AgencyConfigInfo struct {
	types.DynamicData

	// A list of `AgentConfigInfo`s for hosts covered by this
	// <code>Agency</code>.
	//
	// When provisioning a new agent to a host, vSphere
	// ESX Agent Manager tries to find, from left to right in the array, a
	// match for an `AgentConfigInfo` and stops searching at the first
	// one that it finds.
	// If `AgencyConfigInfoEx.vmPlacementPolicy` is set, the array needs to contain only a
	// single agent config. In that case the agent config is not bound to a
	// specific host, but to the whole cluster.
	AgentConfig []AgentConfigInfo `xml:"agentConfig,omitempty" json:"agentConfig,omitempty"`
	// The scope of the <code>Agency</code>.
	Scope BaseAgencyScope `xml:"scope,omitempty,typeattr" json:"scope,omitempty"`
	// If set to <code>true</code>, the client of this agency must manually
	// mark the agent as ready after the agent virtual machine has been
	// provisioned.
	//
	// This is useful if the client of this solution performs
	// some extra reconfiguration of the agent virtual machine before it is
	// powered on.
	//
	// See also `Agent.MarkAsAvailable`.
	ManuallyMarkAgentVmAvailableAfterProvisioning *bool `xml:"manuallyMarkAgentVmAvailableAfterProvisioning" json:"manuallyMarkAgentVmAvailableAfterProvisioning,omitempty"`
	// If set to <code>true</code>, the client of this agency must manually
	// mark the agent as ready after the agent virtual machine has been
	// powered on.
	//
	// In this case, DRS will not regard the agent virtual machine
	// as ready until the client has marked the agent as ready.
	//
	// See also `Agent.MarkAsAvailable`.
	ManuallyMarkAgentVmAvailableAfterPowerOn *bool `xml:"manuallyMarkAgentVmAvailableAfterPowerOn" json:"manuallyMarkAgentVmAvailableAfterPowerOn,omitempty"`
	// If set to <code>true</code>, ESX Agent Manager will use vSphere Linked
	// Clones to speed up the deployment of agent virtual machines.
	//
	// Using
	// linked clones implies that the agent virtual machines cannot use
	// Storage vMotion to move to another vSphere datastore.
	// If set to <code>false</code>, ESX Agent Manager will use Full VM
	// Cloning.
	// If unset default is <code>true<code>.
	OptimizedDeploymentEnabled *bool `xml:"optimizedDeploymentEnabled" json:"optimizedDeploymentEnabled,omitempty"`
	// An optional name to use when naming agent virtual machines.
	//
	// For
	// example, if set to "example-agent", each agent virtual machine will be
	// named "example-agent (1)", "example-agent (2)", and so on. The maximum
	// length of <code>agentName</code> is 70 characters.
	AgentName string `xml:"agentName,omitempty" json:"agentName,omitempty"`
	// Name of the agency.
	//
	// Must be set when creating the agency.
	AgencyName string `xml:"agencyName,omitempty" json:"agencyName,omitempty"`
	// Property <code>agentName</code> is required if this property is set to
	// <code>true</code>.
	//
	// If set to <code>true</code>, ESX Agent Manager will name virtual
	// machines with UUID suffix. For example, "example-agent-UUID".
	// In this case, the maximum length of <code>agentName</code> is 43
	// characters.
	//
	// If not set or is set to <code>false</code>, virtual
	// machines will not contain UUID in their name.
	UseUuidVmName *bool `xml:"useUuidVmName" json:"useUuidVmName,omitempty"`
	// Deprecated use automatically provisioned VMs and register hooks to
	// have control post provisioning and power on.
	//
	// Set to true if agent VMs are manually provisioned.
	//
	// If unset, defaults
	// to false.
	ManuallyProvisioned *bool `xml:"manuallyProvisioned" json:"manuallyProvisioned,omitempty"`
	// Deprecated use automatically provisioned VMs and register hooks to
	// have control post provisioning and power on.
	//
	// Set to true if agent VMs are manually monitored.
	//
	// If unset, defaults to
	// false. This can only be set to true if
	// `AgencyConfigInfo.manuallyProvisioned` is set to true.
	ManuallyMonitored *bool `xml:"manuallyMonitored" json:"manuallyMonitored,omitempty"`
	// Deprecated vUM is no more consulted so this property has no sense
	// anymore.
	//
	// Set to true will install VIBs directly on the hosts even if VMware
	// Update Manager is installed.
	//
	// If unset, defaults to false.
	BypassVumEnabled *bool `xml:"bypassVumEnabled" json:"bypassVumEnabled,omitempty"`
	// Specifies the networks which to be configured on the agent VMs.
	//
	// This property is only applicable for pinned to host VMs - i.e.
	// (`AgencyConfigInfoEx.vmPlacementPolicy`) is not set.
	//
	// If not set or `AgencyConfigInfo.preferHostConfiguration` is set to true, the
	// default host agent VM network (configured through
	// vim.host.EsxAgentHostManager) is used, otherwise the first network from
	// the array that is present on the host is used.
	//
	// At most one of `AgencyConfigInfo.agentVmNetwork` and `AgencyConfigInfoEx.vmNetworkMapping`
	// needs to be set.
	//
	// Refers instances of `Network`.
	AgentVmNetwork []types.ManagedObjectReference `xml:"agentVmNetwork,omitempty" json:"agentVmNetwork,omitempty"`
	// The datastores used to configure the storage on the agent VMs.
	//
	// This property is required if `AgencyConfigInfoEx.vmPlacementPolicy` is set and
	// `AgencyConfigInfoEx.datastoreSelectionPolicy` is not set. In that case the first
	// element from the list is used.
	//
	// If not set or `AgencyConfigInfo.preferHostConfiguration` is set to true and
	// `AgencyConfigInfoEx.vmPlacementPolicy` is not set, the default host agent VM
	// datastore (configured through vim.host.EsxAgentHostManager) is used,
	// otherwise the first datastore from the array that is present on the
	// host is used.
	//
	// If `AgencyConfigInfoEx.vmPlacementPolicy` is set at most one of
	// `AgencyConfigInfo.agentVmDatastore` and `AgencyConfigInfoEx.datastoreSelectionPolicy` needs
	// to be set. If `AgencyConfigInfoEx.vmPlacementPolicy` is not set
	// `AgencyConfigInfoEx.datastoreSelectionPolicy` takes precedence over
	// `AgencyConfigInfo.agentVmDatastore` .
	//
	// Refers instances of `Datastore`.
	AgentVmDatastore []types.ManagedObjectReference `xml:"agentVmDatastore,omitempty" json:"agentVmDatastore,omitempty"`
	// If set to true the default agent VM datastore and network will take
	// precedence over `AgencyConfigInfo.agentVmNetwork` and
	// `AgencyConfigInfo.agentVmDatastore` when configuring the agent
	// VMs.
	//
	// This property is not used if `AgencyConfigInfoEx.vmPlacementPolicy` is set.
	PreferHostConfiguration *bool `xml:"preferHostConfiguration" json:"preferHostConfiguration,omitempty"`
	// Deprecated that is a custom configuration that should be setup by the
	// agency owner. One way is to use
	// `AgencyConfigInfo.manuallyMarkAgentVmAvailableAfterPowerOn` or
	// `AgencyConfigInfo.manuallyMarkAgentVmAvailableAfterProvisioning`
	// hooks.
	//
	// If set, a property with id "ip" and value an IP from the pool is added
	// to the vApp configuration of the deployed VMs.
	IpPool *types.IpPool `xml:"ipPool,omitempty" json:"ipPool,omitempty"`
	// Defines the resource pools where VMs to be deployed.
	//
	// If specified, the VMs for every compute resource in the scope will be
	// deployed to its corresponding resource pool.
	// If not specified, the agent VMs for each compute resource will be
	// deployed under top level nested resource pool created for the agent
	// VMs. If unable to create a nested resource pool, the root resource pool
	// of the compute resource will be used.
	ResourcePools []AgencyVMResourcePool `xml:"resourcePools,omitempty" json:"resourcePools,omitempty"`
	// Defines the folders where VMs to be deployed.
	//
	// If specified, the VMs for every compute resource in the scope will be
	// deployed to its corresponding folder. The link is made between the
	// compute resource parent and the datacenter the folder belongs to
	// `AgencyVMFolder.datacenterId`.
	// If not specified, the agent VMs for each compute resource will be
	// deployed in top level folder created in each datacenter for the agent
	// VMs.
	Folders         []AgencyVMFolder `xml:"folders,omitempty" json:"folders,omitempty"`
	OmitDRSBlocking *bool            `xml:"omitDRSBlocking" json:"omitDRSBlocking,omitempty"`
}

func init() {
	types.Add("eam:AgencyConfigInfo", reflect.TypeOf((*AgencyConfigInfo)(nil)).Elem())
}

// Agency is disabled - one or more ClusterComputeResources from it's scope are
// disabled.
//
// This is not a remediable issue. To remediate, re-enable the cluster.
//
// This structure may be used only with operations rendered under `/eam`.
type AgencyDisabled struct {
	AgencyIssue
}

func init() {
	types.Add("eam:AgencyDisabled", reflect.TypeOf((*AgencyDisabled)(nil)).Elem())
	types.AddMinAPIVersionForType("eam:AgencyDisabled", "7.6")
}

// Base class for all agency issues.
//
// This structure may be used only with operations rendered under `/eam`.
type AgencyIssue struct {
	Issue

	// The agency to which this issue belongs.
	//
	// Refers instance of `Agency`.
	Agency types.ManagedObjectReference `xml:"agency" json:"agency"`
	// The name of the agency.
	AgencyName string `xml:"agencyName" json:"agencyName"`
	// The ID of the solution to which this issue belongs.
	SolutionId string `xml:"solutionId" json:"solutionId"`
	// The name of the solution to which this issue belongs.
	SolutionName string `xml:"solutionName" json:"solutionName"`
}

func init() {
	types.Add("eam:AgencyIssue", reflect.TypeOf((*AgencyIssue)(nil)).Elem())
}

type AgencyQueryRuntime AgencyQueryRuntimeRequestType

func init() {
	types.Add("eam:AgencyQueryRuntime", reflect.TypeOf((*AgencyQueryRuntime)(nil)).Elem())
}

type AgencyQueryRuntimeRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
}

func init() {
	types.Add("eam:AgencyQueryRuntimeRequestType", reflect.TypeOf((*AgencyQueryRuntimeRequestType)(nil)).Elem())
}

type AgencyQueryRuntimeResponse struct {
	Returnval BaseEamObjectRuntimeInfo `xml:"returnval,typeattr" json:"returnval"`
}

// Deprecated as of vSphere 9.0. Please refer to vLCM APIs.
//
// Scope specifies which where to deploy agents.
//
// This structure may be used only with operations rendered under `/eam`.
type AgencyScope struct {
	types.DynamicData
}

func init() {
	types.Add("eam:AgencyScope", reflect.TypeOf((*AgencyScope)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. Please refer to vLCM System VMs APIs.
//
// Represents the mapping of a VM folder to a datacenter.
//
// This structure may be used only with operations rendered under `/eam`.
type AgencyVMFolder struct {
	types.DynamicData

	// Folder identifier.
	//
	// The folder must be present in the corresponding
	// datacenter.
	//
	// Refers instance of `Folder`.
	FolderId types.ManagedObjectReference `xml:"folderId" json:"folderId"`
	// Datacenter identifier.
	//
	// Refers instance of `Datacenter`.
	DatacenterId types.ManagedObjectReference `xml:"datacenterId" json:"datacenterId"`
}

func init() {
	types.Add("eam:AgencyVMFolder", reflect.TypeOf((*AgencyVMFolder)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. Please refer to vLCM System VMs APIs.
//
// Represents the mapping of a VM resource pool to a compute resource.
//
// This structure may be used only with operations rendered under `/eam`.
type AgencyVMResourcePool struct {
	types.DynamicData

	// Resource pool identifier.
	//
	// The resource pool must be present in the
	// corresponding compute resource.
	//
	// Refers instance of `ResourcePool`.
	ResourcePoolId types.ManagedObjectReference `xml:"resourcePoolId" json:"resourcePoolId"`
	// Compute resource identifier.
	//
	// Refers instance of `ComputeResource`.
	ComputeResourceId types.ManagedObjectReference `xml:"computeResourceId" json:"computeResourceId"`
}

func init() {
	types.Add("eam:AgencyVMResourcePool", reflect.TypeOf((*AgencyVMResourcePool)(nil)).Elem())
}

type Agency_Disable Agency_DisableRequestType

func init() {
	types.Add("eam:Agency_Disable", reflect.TypeOf((*Agency_Disable)(nil)).Elem())
}

type Agency_DisableRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
}

func init() {
	types.Add("eam:Agency_DisableRequestType", reflect.TypeOf((*Agency_DisableRequestType)(nil)).Elem())
}

type Agency_DisableResponse struct {
}

type Agency_Enable Agency_EnableRequestType

func init() {
	types.Add("eam:Agency_Enable", reflect.TypeOf((*Agency_Enable)(nil)).Elem())
}

type Agency_EnableRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
}

func init() {
	types.Add("eam:Agency_EnableRequestType", reflect.TypeOf((*Agency_EnableRequestType)(nil)).Elem())
}

type Agency_EnableResponse struct {
}

// Deprecated as of vSphere 9.0. Please refer to vLCM APIs.
//
// Specifies an SSL policy that trusts any SSL certificate.
//
// This structure may be used only with operations rendered under `/eam`.
type AgentAnyCertificate struct {
	AgentSslTrust
}

func init() {
	types.Add("eam:AgentAnyCertificate", reflect.TypeOf((*AgentAnyCertificate)(nil)).Elem())
	types.AddMinAPIVersionForType("eam:AgentAnyCertificate", "8.2")
}

// Deprecated as of vSphere 9.0. Please refer to vLCM APIs.
//
// A description of what should be put on a host.
//
// By setting the
// <code>productLineId</code> and <code>hostVersion</code>, you can specify
// the types of hosts for which to use the <code>ConfigInfo</code>.
//
// This structure may be used only with operations rendered under `/eam`.
type AgentConfigInfo struct {
	types.DynamicData

	// The product line ID of the host.
	//
	// Examples of values are "esx" or
	// "embeddedEsx". If omitted, the host's product line ID is not considered
	// when matching an <code>AgentPackage</code> against a host.
	ProductLineId string `xml:"productLineId,omitempty" json:"productLineId,omitempty"`
	// A dot-separated string of the host version.
	//
	// Examples of values are
	// "4.1.0" "3.5", "4.\*" where \* is a wildcard meaning any version minor
	// version of the major version 4. If omitted, the host version will not
	// be considered when matching an <code>AgentPackage</code> against a
	// host.
	// This property is not used if
	// `AgencyConfigInfoEx.vmPlacementPolicy` is set. It is client's
	// responsibility to trigger an agency upgrade with a new
	// `AgentConfigInfo.ovfPackageUrl`.
	HostVersion string `xml:"hostVersion,omitempty" json:"hostVersion,omitempty"`
	// The URL of the solution's agent OVF package.
	//
	// If not set, no agent
	// virtual machines are installed on the hosts covered by the scope.
	// If `AgencyConfigInfoEx.vmPlacementPolicy` is set, the VM needs to
	// be agnostic to the different host versions inside the cluster.
	OvfPackageUrl string `xml:"ovfPackageUrl,omitempty" json:"ovfPackageUrl,omitempty"`
	// The authentication scheme needed to access the OVF URL.
	//
	// The supported schemes are listed in `AgentConfigInfoAuthenticationScheme_enum`
	// Default value is `NONE` if not specified.
	//
	// If the scheme is not `null` and is not set to `NONE`
	// then `AgentConfigInfo.ovfPackageUrl` must point to the VC
	AuthenticationScheme string `xml:"authenticationScheme,omitempty" json:"authenticationScheme,omitempty" vim:"9.0"`
	// Specifies an SSL trust policy to be use for verification of the
	// server that hosts the `AgentConfigInfo.ovfPackageUrl`.
	//
	// If not set, the server
	// certificate is validated against the trusted root certificates of the
	// OS (Photon) and VECS (TRUSTED\_ROOTS).
	OvfSslTrust BaseAgentSslTrust `xml:"ovfSslTrust,omitempty,typeattr" json:"ovfSslTrust,omitempty" vim:"8.2"`
	// The part of the OVF environment that can be set by the solution.
	//
	// This
	// is where Properties that the agent virtual machine's OVF descriptor
	// specifies are set here. All properties that are specified as
	// user-configurable must be set.
	OvfEnvironment *AgentOvfEnvironmentInfo `xml:"ovfEnvironment,omitempty" json:"ovfEnvironment,omitempty"`
	// An optional URL to an offline bundle.
	//
	// If not set, no VIB is installed
	// on the hosts in the scope. Offline bundles are only supported on 4.0
	// hosts and later.
	//
	// VIB downgrade is not permitted - in case a VIB with the same name, but
	// lower version is installed on a host in the scope the VIB installation
	// on that host will not succeed.
	//
	// If two or more agents have the same VIB with different versions on the
	// same host, the install/uninstall behaviour is undefined (the VIB may
	// remain installed, etc.).
	//
	// The property is not used if `AgencyConfigInfoEx.vmPlacementPolicy`
	// is set.
	VibUrl string `xml:"vibUrl,omitempty" json:"vibUrl,omitempty"`
	// Specifies an SSL trust policy to be use for verification of the
	// server that hosts the `AgentConfigInfo.vibUrl`.
	//
	// If not set, the server
	// certificate is validated against the trusted root certificates of the
	// OS (Photon) and VECS (TRUSTED\_ROOTS).
	VibSslTrust BaseAgentSslTrust `xml:"vibSslTrust,omitempty,typeattr" json:"vibSslTrust,omitempty" vim:"8.2"`
	// Deprecated vIB matching rules are no longer supported by EAM. Same
	// overlaps with VIB dependency requirements which reside in
	// each VIB's metadata.
	//
	// Optional Vib matching rules.
	//
	// If set, the Vib, specified by vibUrl, will
	// be installed either
	//   - if there is installed Vib on the host which name and version match
	//     the regular expressions in the corresponding rule
	//   - or there isn't any installed Vib on the host with name which
	//     matches the Vib name regular expression in the corresponding rule. Vib
	//     matching rules are usually used for controlling VIB upgrades, in which
	//     case the name regular expression matches any previous versions of the
	//     agency Vib and version regular expression determines whether the
	//     existing Vib should be upgraded.
	//
	// For every Vib in the Vib package, only one Vib matching rule can be
	// defined. If specified more than one, it is not determined which one
	// will be used. The Vib name regular expression in the Vib matching rule
	// will be matched against the name of the Vib which will be installed.
	// Only rules for Vibs which are defined in the Vib package metadata will
	// be taken in account.
	VibMatchingRules []AgentVibMatchingRule `xml:"vibMatchingRules,omitempty" json:"vibMatchingRules,omitempty"`
	// Deprecated use VIB metadata to add such dependency.
	//
	// An optional name of a VIB.
	//
	// If set, no VIB is installed on the host. The
	// host is checked if a VIB with vibName is already installed on it. Also
	// the vibUrl must not be set together with the vibUrl.
	VibName string `xml:"vibName,omitempty" json:"vibName,omitempty"`
	// Deprecated that is a custom setup specific for a particular agency.
	// The agency owner should do it using other means, e.g.
	// `AgencyConfigInfo.manuallyMarkAgentVmAvailableAfterPowerOn` or
	// `AgencyConfigInfo.manuallyMarkAgentVmAvailableAfterProvisioning`
	// hooks. Support for this has been removed. Setting this to
	// <code>true</code> will no longer have any effect.
	//
	// If set to <code>true</code>, the hosts in the scope must be configured
	// for DvFilter before VIBs and agent virtual machines are deployed on
	// them.
	//
	// If not set or set to <code>false</code>, no DvFilter
	// configuration is done on the hosts.
	DvFilterEnabled *bool `xml:"dvFilterEnabled" json:"dvFilterEnabled,omitempty"`
	// Deprecated express that requirement in the VIB descriptor with
	// 'live-remove-allowed=false'.
	//
	// An optional boolean flag to specify whether the agent's host is
	// rebooted after the VIB is uninstalled.
	//
	// If not set, the default value is
	// <code>false</code>. If set to <code>true</code>, the agent gets a
	// `VibRequiresHostReboot` issue after a successful
	// uninstallation.
	RebootHostAfterVibUninstall *bool `xml:"rebootHostAfterVibUninstall" json:"rebootHostAfterVibUninstall,omitempty"`
	// If set the virtual machine will be configured with the services and
	// allow VMCI access from the virtual machine to the installed VIB.
	VmciService []string `xml:"vmciService,omitempty" json:"vmciService,omitempty"`
	// AgentVM disk provisioning type.
	//
	// Defaults to `none` if not specified.
	OvfDiskProvisioning string `xml:"ovfDiskProvisioning,omitempty" json:"ovfDiskProvisioning,omitempty"`
	// Defines the storage policies configured on Agent VMs.
	//
	// Storage policies
	// are configured on all VM related objects including disks.
	// NOTE: The property needs to be configured on each update, otherwise
	// vSphere ESX Agent Manager will unset this configuration for all future
	// agent VMs.
	VmStoragePolicies []BaseAgentStoragePolicy `xml:"vmStoragePolicies,omitempty,typeattr" json:"vmStoragePolicies,omitempty"`
	// Specifies the resource configuration to be used when deploying an Agent
	// VM.
	//
	// It must correspond to the Configuration element of the
	// DeploymentOptionSection in the OVF specification.
	VmResourceConfiguration string `xml:"vmResourceConfiguration,omitempty" json:"vmResourceConfiguration,omitempty" vim:"8.3"`
}

func init() {
	types.Add("eam:AgentConfigInfo", reflect.TypeOf((*AgentConfigInfo)(nil)).Elem())
}

// Base class for all agent issues.
//
// This structure may be used only with operations rendered under `/eam`.
type AgentIssue struct {
	AgencyIssue

	// The agent that has this issue.
	//
	// Refers instance of `Agent`.
	Agent types.ManagedObjectReference `xml:"agent" json:"agent"`
	// The name of the agent.
	AgentName string `xml:"agentName" json:"agentName"`
	// The managed object reference to the host on which this agent is located.
	//
	// Refers instance of `HostSystem`.
	Host types.ManagedObjectReference `xml:"host" json:"host"`
	// The name of the host on which this agent is located.
	HostName string `xml:"hostName" json:"hostName"`
}

func init() {
	types.Add("eam:AgentIssue", reflect.TypeOf((*AgentIssue)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. Please refer to vLCM APIs.
//
// The <code>OvfEnvironment</code> is used to assign OVF environment
// properties in the `AgentConfigInfo`.
//
// It specifies the values that
// map to properties in the agent virtual machine's OVF descriptor.
//
// This structure may be used only with operations rendered under `/eam`.
type AgentOvfEnvironmentInfo struct {
	types.DynamicData

	// The OVF properties that are assigned to the agent virtual machine's OVF
	// environment when it is powered on.
	OvfProperty []AgentOvfEnvironmentInfoOvfProperty `xml:"ovfProperty,omitempty" json:"ovfProperty,omitempty"`
}

func init() {
	types.Add("eam:AgentOvfEnvironmentInfo", reflect.TypeOf((*AgentOvfEnvironmentInfo)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. Please refer to vLCM APIs.
//
// One OVF property.
//
// This structure may be used only with operations rendered under `/eam`.
type AgentOvfEnvironmentInfoOvfProperty struct {
	types.DynamicData

	// The name of the property in the OVF descriptor.
	Key string `xml:"key" json:"key"`
	// The value of the property.
	Value string `xml:"value" json:"value"`
}

func init() {
	types.Add("eam:AgentOvfEnvironmentInfoOvfProperty", reflect.TypeOf((*AgentOvfEnvironmentInfoOvfProperty)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. Please refer to vLCM APIs.
//
// Specifies an SSL policy that trusts one specific pinned PEM encoded
// SSL certificate.
//
// This structure may be used only with operations rendered under `/eam`.
type AgentPinnedPemCertificate struct {
	AgentSslTrust

	// PEM encoded pinned SSL certificate of the server that needs to be
	// trusted.
	SslCertificate string `xml:"sslCertificate" json:"sslCertificate"`
}

func init() {
	types.Add("eam:AgentPinnedPemCertificate", reflect.TypeOf((*AgentPinnedPemCertificate)(nil)).Elem())
	types.AddMinAPIVersionForType("eam:AgentPinnedPemCertificate", "8.2")
}

type AgentQueryConfig AgentQueryConfigRequestType

func init() {
	types.Add("eam:AgentQueryConfig", reflect.TypeOf((*AgentQueryConfig)(nil)).Elem())
}

type AgentQueryConfigRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
}

func init() {
	types.Add("eam:AgentQueryConfigRequestType", reflect.TypeOf((*AgentQueryConfigRequestType)(nil)).Elem())
}

type AgentQueryConfigResponse struct {
	Returnval AgentConfigInfo `xml:"returnval" json:"returnval"`
}

type AgentQueryRuntime AgentQueryRuntimeRequestType

func init() {
	types.Add("eam:AgentQueryRuntime", reflect.TypeOf((*AgentQueryRuntime)(nil)).Elem())
}

type AgentQueryRuntimeRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
}

func init() {
	types.Add("eam:AgentQueryRuntimeRequestType", reflect.TypeOf((*AgentQueryRuntimeRequestType)(nil)).Elem())
}

type AgentQueryRuntimeResponse struct {
	Returnval AgentRuntimeInfo `xml:"returnval" json:"returnval"`
}

// Deprecated as of vSphere 9.0. Please refer to vLCM APIs.
//
// Extends <code>RuntimeInfo</code> with information regarding the deployment
// of an agent on a specific host.
//
// This structure may be used only with operations rendered under `/eam`.
type AgentRuntimeInfo struct {
	EamObjectRuntimeInfo

	// Deprecated get that info calling the virtual machine VIM API.
	//
	// The power state of an agent virtual machine.
	VmPowerState types.VirtualMachinePowerState `xml:"vmPowerState" json:"vmPowerState"`
	// Deprecated get that info calling the virtual machine VIM API.
	//
	// True if the vSphere ESX Agent Manager is receiving heartbeats from the
	// agent virtual machine.
	ReceivingHeartBeat bool `xml:"receivingHeartBeat" json:"receivingHeartBeat"`
	// The agent host.
	//
	// Refers instance of `HostSystem`.
	Host *types.ManagedObjectReference `xml:"host,omitempty" json:"host,omitempty"`
	// The agent virtual machine.
	//
	// Refers instance of `VirtualMachine`.
	Vm *types.ManagedObjectReference `xml:"vm,omitempty" json:"vm,omitempty"`
	// Deprecated get that info calling the virtual machine VIM API.
	//
	// The IP address of the agent virtual machine
	VmIp string `xml:"vmIp,omitempty" json:"vmIp,omitempty"`
	// Deprecated get that info calling the virtual machine VIM API.
	//
	// The name of the agent virtual machine.
	VmName string `xml:"vmName" json:"vmName"`
	// Deprecated in order to retrieve agent resource pool use VIM API.
	//
	// The ESX agent resource pool in which the agent virtual machine resides.
	//
	// Refers instance of `ResourcePool`.
	EsxAgentResourcePool *types.ManagedObjectReference `xml:"esxAgentResourcePool,omitempty" json:"esxAgentResourcePool,omitempty"`
	// Deprecated in order to retrieve agent VM folder use VIM API.
	//
	// The ESX agent folder in which the agent virtual machine resides.
	//
	// Refers instance of `Folder`.
	EsxAgentFolder *types.ManagedObjectReference `xml:"esxAgentFolder,omitempty" json:"esxAgentFolder,omitempty"`
	// Deprecated use `AgentRuntimeInfo.installedVibs` instead.
	//
	// An optional array of IDs of installed bulletins for this agent.
	InstalledBulletin []string `xml:"installedBulletin,omitempty" json:"installedBulletin,omitempty"`
	// Information about the installed vibs on the host.
	InstalledVibs []VibVibInfo `xml:"installedVibs,omitempty" json:"installedVibs,omitempty"`
	// The agency this agent belongs to.
	//
	// Refers instance of `Agency`.
	Agency *types.ManagedObjectReference `xml:"agency,omitempty" json:"agency,omitempty"`
	// Active VM hook.
	//
	// If present agent is actively waiting for `Agent.MarkAsAvailable`.
	// See `AgentVmHook`.
	VmHook *AgentVmHook `xml:"vmHook,omitempty" json:"vmHook,omitempty"`
}

func init() {
	types.Add("eam:AgentRuntimeInfo", reflect.TypeOf((*AgentRuntimeInfo)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. Please refer to vLCM APIs.
//
// Specifies an SSL trust policy.
//
// This structure may be used only with operations rendered under `/eam`.
type AgentSslTrust struct {
	types.DynamicData
}

func init() {
	types.Add("eam:AgentSslTrust", reflect.TypeOf((*AgentSslTrust)(nil)).Elem())
	types.AddMinAPIVersionForType("eam:AgentSslTrust", "8.2")
}

// Deprecated as of vSphere 9.0. Please refer to vLCM APIs.
//
// Specifies the storage policies configured on Agent VMs.
//
// This structure may be used only with operations rendered under `/eam`.
type AgentStoragePolicy struct {
	types.DynamicData
}

func init() {
	types.Add("eam:AgentStoragePolicy", reflect.TypeOf((*AgentStoragePolicy)(nil)).Elem())
}

// Deprecated vIB matching rules are no longer supported by EAM. Same
// overlaps with VIB dependency requirements which reside in each
// VIB's metadata.
//
// Specifies regular expressions for Vib name and version.
//
// This structure may be used only with operations rendered under `/eam`.
type AgentVibMatchingRule struct {
	types.DynamicData

	// Vib name regular expression.
	VibNameRegex string `xml:"vibNameRegex" json:"vibNameRegex"`
	// Vib version regular expression.
	VibVersionRegex string `xml:"vibVersionRegex" json:"vibVersionRegex"`
}

func init() {
	types.Add("eam:AgentVibMatchingRule", reflect.TypeOf((*AgentVibMatchingRule)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. Please refer to vLCM APIs.
//
// Represents an active hook of the VM lifecycle which EAM is waiting on to
// be processed by the client.
//
// The supported hooks are defined by
// `AgencyConfigInfo.manuallyMarkAgentVmAvailableAfterProvisioning`
// and
// `AgencyConfigInfo.manuallyMarkAgentVmAvailableAfterProvisioning`.
// See `Agent.MarkAsAvailable`
//
// This structure may be used only with operations rendered under `/eam`.
type AgentVmHook struct {
	types.DynamicData

	// The VM for which lifecycle is this hook.
	//
	// This VM may differ from `AgentRuntimeInfo.vm` while an upgrade
	// of the agent VM is in progress.
	//
	// Refers instance of `VirtualMachine`.
	Vm types.ManagedObjectReference `xml:"vm" json:"vm"`
	// The current VM lifecycle state.
	VmState string `xml:"vmState" json:"vmState"`
}

func init() {
	types.Add("eam:AgentVmHook", reflect.TypeOf((*AgentVmHook)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. Please refer to vLCM APIs.
//
// Specifies vSAN specific storage policy configured on Agent VMs.
//
// This structure may be used only with operations rendered under `/eam`.
type AgentVsanStoragePolicy struct {
	AgentStoragePolicy

	// ID of a storage policy profile created by the user.
	//
	// The type of the
	// profile must be `VirtualMachineDefinedProfileSpec`. The ID must be valid
	// `VirtualMachineDefinedProfileSpec.profileId`.
	ProfileId string `xml:"profileId" json:"profileId"`
}

func init() {
	types.Add("eam:AgentVsanStoragePolicy", reflect.TypeOf((*AgentVsanStoragePolicy)(nil)).Elem())
}

// A boxed array of `AgencyVMFolder`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/eam`.
type ArrayOfAgencyVMFolder struct {
	AgencyVMFolder []AgencyVMFolder `xml:"AgencyVMFolder,omitempty" json:"_value"`
}

func init() {
	types.Add("eam:ArrayOfAgencyVMFolder", reflect.TypeOf((*ArrayOfAgencyVMFolder)(nil)).Elem())
}

// A boxed array of `AgencyVMResourcePool`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/eam`.
type ArrayOfAgencyVMResourcePool struct {
	AgencyVMResourcePool []AgencyVMResourcePool `xml:"AgencyVMResourcePool,omitempty" json:"_value"`
}

func init() {
	types.Add("eam:ArrayOfAgencyVMResourcePool", reflect.TypeOf((*ArrayOfAgencyVMResourcePool)(nil)).Elem())
}

// A boxed array of `AgentConfigInfo`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/eam`.
type ArrayOfAgentConfigInfo struct {
	AgentConfigInfo []AgentConfigInfo `xml:"AgentConfigInfo,omitempty" json:"_value"`
}

func init() {
	types.Add("eam:ArrayOfAgentConfigInfo", reflect.TypeOf((*ArrayOfAgentConfigInfo)(nil)).Elem())
}

// A boxed array of `AgentOvfEnvironmentInfoOvfProperty`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/eam`.
type ArrayOfAgentOvfEnvironmentInfoOvfProperty struct {
	AgentOvfEnvironmentInfoOvfProperty []AgentOvfEnvironmentInfoOvfProperty `xml:"AgentOvfEnvironmentInfoOvfProperty,omitempty" json:"_value"`
}

func init() {
	types.Add("eam:ArrayOfAgentOvfEnvironmentInfoOvfProperty", reflect.TypeOf((*ArrayOfAgentOvfEnvironmentInfoOvfProperty)(nil)).Elem())
}

// A boxed array of `AgentStoragePolicy`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/eam`.
type ArrayOfAgentStoragePolicy struct {
	AgentStoragePolicy []BaseAgentStoragePolicy `xml:"AgentStoragePolicy,omitempty,typeattr" json:"_value"`
}

func init() {
	types.Add("eam:ArrayOfAgentStoragePolicy", reflect.TypeOf((*ArrayOfAgentStoragePolicy)(nil)).Elem())
}

// A boxed array of `AgentVibMatchingRule`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/eam`.
type ArrayOfAgentVibMatchingRule struct {
	AgentVibMatchingRule []AgentVibMatchingRule `xml:"AgentVibMatchingRule,omitempty" json:"_value"`
}

func init() {
	types.Add("eam:ArrayOfAgentVibMatchingRule", reflect.TypeOf((*ArrayOfAgentVibMatchingRule)(nil)).Elem())
}

// A boxed array of `Issue`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/eam`.
type ArrayOfIssue struct {
	Issue []BaseIssue `xml:"Issue,omitempty,typeattr" json:"_value"`
}

func init() {
	types.Add("eam:ArrayOfIssue", reflect.TypeOf((*ArrayOfIssue)(nil)).Elem())
}

type ArrayOfSolutionsAlternativeVmSpec struct {
	SolutionsAlternativeVmSpec []SolutionsAlternativeVmSpec `xml:"SolutionsAlternativeVmSpec,omitempty" json:"_value"`
}

func init() {
	types.Add("eam:ArrayOfSolutionsAlternativeVmSpec", reflect.TypeOf((*ArrayOfSolutionsAlternativeVmSpec)(nil)).Elem())
}

// A boxed array of `SolutionsClusterSolutionComplianceResult`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/eam`.
type ArrayOfSolutionsClusterSolutionComplianceResult struct {
	SolutionsClusterSolutionComplianceResult []SolutionsClusterSolutionComplianceResult `xml:"SolutionsClusterSolutionComplianceResult,omitempty" json:"_value"`
}

func init() {
	types.Add("eam:ArrayOfSolutionsClusterSolutionComplianceResult", reflect.TypeOf((*ArrayOfSolutionsClusterSolutionComplianceResult)(nil)).Elem())
}

// A boxed array of `SolutionsDeploymentUnitComplianceResult`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/eam`.
type ArrayOfSolutionsDeploymentUnitComplianceResult struct {
	SolutionsDeploymentUnitComplianceResult []SolutionsDeploymentUnitComplianceResult `xml:"SolutionsDeploymentUnitComplianceResult,omitempty" json:"_value"`
}

func init() {
	types.Add("eam:ArrayOfSolutionsDeploymentUnitComplianceResult", reflect.TypeOf((*ArrayOfSolutionsDeploymentUnitComplianceResult)(nil)).Elem())
}

// A boxed array of `SolutionsHookConfig`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/eam`.
type ArrayOfSolutionsHookConfig struct {
	SolutionsHookConfig []SolutionsHookConfig `xml:"SolutionsHookConfig,omitempty" json:"_value"`
}

func init() {
	types.Add("eam:ArrayOfSolutionsHookConfig", reflect.TypeOf((*ArrayOfSolutionsHookConfig)(nil)).Elem())
}

// A boxed array of `SolutionsHookInfo`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/eam`.
type ArrayOfSolutionsHookInfo struct {
	SolutionsHookInfo []SolutionsHookInfo `xml:"SolutionsHookInfo,omitempty" json:"_value"`
}

func init() {
	types.Add("eam:ArrayOfSolutionsHookInfo", reflect.TypeOf((*ArrayOfSolutionsHookInfo)(nil)).Elem())
}

// A boxed array of `SolutionsHostComplianceResult`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/eam`.
type ArrayOfSolutionsHostComplianceResult struct {
	SolutionsHostComplianceResult []SolutionsHostComplianceResult `xml:"SolutionsHostComplianceResult,omitempty" json:"_value"`
}

func init() {
	types.Add("eam:ArrayOfSolutionsHostComplianceResult", reflect.TypeOf((*ArrayOfSolutionsHostComplianceResult)(nil)).Elem())
}

// A boxed array of `SolutionsOvfProperty`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/eam`.
type ArrayOfSolutionsOvfProperty struct {
	SolutionsOvfProperty []SolutionsOvfProperty `xml:"SolutionsOvfProperty,omitempty" json:"_value"`
}

func init() {
	types.Add("eam:ArrayOfSolutionsOvfProperty", reflect.TypeOf((*ArrayOfSolutionsOvfProperty)(nil)).Elem())
}

// A boxed array of `SolutionsSolutionComplianceResult`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/eam`.
type ArrayOfSolutionsSolutionComplianceResult struct {
	SolutionsSolutionComplianceResult []SolutionsSolutionComplianceResult `xml:"SolutionsSolutionComplianceResult,omitempty" json:"_value"`
}

func init() {
	types.Add("eam:ArrayOfSolutionsSolutionComplianceResult", reflect.TypeOf((*ArrayOfSolutionsSolutionComplianceResult)(nil)).Elem())
}

// A boxed array of `SolutionsSolutionConfig`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/eam`.
type ArrayOfSolutionsSolutionConfig struct {
	SolutionsSolutionConfig []SolutionsSolutionConfig `xml:"SolutionsSolutionConfig,omitempty" json:"_value"`
}

func init() {
	types.Add("eam:ArrayOfSolutionsSolutionConfig", reflect.TypeOf((*ArrayOfSolutionsSolutionConfig)(nil)).Elem())
}

// A boxed array of `SolutionsSolutionValidationResult`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/eam`.
type ArrayOfSolutionsSolutionValidationResult struct {
	SolutionsSolutionValidationResult []SolutionsSolutionValidationResult `xml:"SolutionsSolutionValidationResult,omitempty" json:"_value"`
}

func init() {
	types.Add("eam:ArrayOfSolutionsSolutionValidationResult", reflect.TypeOf((*ArrayOfSolutionsSolutionValidationResult)(nil)).Elem())
}

// A boxed array of `SolutionsStoragePolicy`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/eam`.
type ArrayOfSolutionsStoragePolicy struct {
	SolutionsStoragePolicy []BaseSolutionsStoragePolicy `xml:"SolutionsStoragePolicy,omitempty,typeattr" json:"_value"`
}

func init() {
	types.Add("eam:ArrayOfSolutionsStoragePolicy", reflect.TypeOf((*ArrayOfSolutionsStoragePolicy)(nil)).Elem())
}

// A boxed array of `SolutionsVMNetworkMapping`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/eam`.
type ArrayOfSolutionsVMNetworkMapping struct {
	SolutionsVMNetworkMapping []SolutionsVMNetworkMapping `xml:"SolutionsVMNetworkMapping,omitempty" json:"_value"`
}

func init() {
	types.Add("eam:ArrayOfSolutionsVMNetworkMapping", reflect.TypeOf((*ArrayOfSolutionsVMNetworkMapping)(nil)).Elem())
}

type ArrayOfSolutionsVmSelectionSpecMapping struct {
	SolutionsVmSelectionSpecMapping []SolutionsVmSelectionSpecMapping `xml:"SolutionsVmSelectionSpecMapping,omitempty" json:"_value"`
}

func init() {
	types.Add("eam:ArrayOfSolutionsVmSelectionSpecMapping", reflect.TypeOf((*ArrayOfSolutionsVmSelectionSpecMapping)(nil)).Elem())
}

// A boxed array of `VibVibInfo`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/eam`.
type ArrayOfVibVibInfo struct {
	VibVibInfo []VibVibInfo `xml:"VibVibInfo,omitempty" json:"_value"`
}

func init() {
	types.Add("eam:ArrayOfVibVibInfo", reflect.TypeOf((*ArrayOfVibVibInfo)(nil)).Elem())
}

// An agent virtual machine is expected to be deployed on a host, but the agent virtual machine
// cannot be deployed because the vSphere ESX Agent Manager is unable to access the OVF
// package for the agent.
//
// This typically happens because the Web server providing the
// OVF package is down. The Web server is often internal to the solution
// that created the Agency.
//
// This is an active remediable issue. To remediate, vSphere ESX Agent Manager redeploys the agent.
//
// This structure may be used only with operations rendered under `/eam`.
type CannotAccessAgentOVF struct {
	VmNotDeployed

	// The URL from which the OVF could not be downloaded.
	DownloadUrl string `xml:"downloadUrl" json:"downloadUrl"`
}

func init() {
	types.Add("eam:CannotAccessAgentOVF", reflect.TypeOf((*CannotAccessAgentOVF)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. Please refer to vLCM Image APIs.
//
// An agent VIB module is expected to be deployed on a host, but the VIM module
// cannot be deployed because the vSphere ESX Agent Manager is unable to access the VIB
// package for the agent.
//
// This typically happens because the Web server providing the
// VIB package is down. The Web server is often internal to the solution
// that created the Agency.
//
// This is an active remediable issue. To remediate, vSphere ESX Agent Manager reinstalls the VIB module.
//
// This structure may be used only with operations rendered under `/eam`.
type CannotAccessAgentVib struct {
	VibNotInstalled

	// The URL from which the VIB package could not be downloaded.
	DownloadUrl string `xml:"downloadUrl" json:"downloadUrl"`
}

func init() {
	types.Add("eam:CannotAccessAgentVib", reflect.TypeOf((*CannotAccessAgentVib)(nil)).Elem())
}

// This exception is thrown if the certificate provided by the
// provider is not trusted.
//
// This structure may be used only with operations rendered under `/sms`.
type CertificateNotTrusted struct {
	AgentIssue

	Url string `xml:"url" json:"url"`
}

func init() {
	types.Add("eam:CertificateNotTrusted", reflect.TypeOf((*CertificateNotTrusted)(nil)).Elem())
}

// An CertificateNotTrusted fault is thrown when an Agency's configuration
// contains OVF package URL or VIB URL for that vSphere ESX Agent Manager is not
// able to make successful SSL trust verification of the server's certificate.
//
// Reasons for this might be that the certificate provided via the API
// `AgentConfigInfo.ovfSslTrust` and `AgentConfigInfo.vibSslTrust`
// or via the script /usr/lib/vmware-eam/bin/eam-utility.py
//   - is invalid.
//   - does not match the server's certificate.
//
// If there is no provided certificate, the fault is thrown when the server's
// certificate is not trusted by the system or is invalid - @see
// `AgentConfigInfo.ovfSslTrust` and
// `AgentConfigInfo.vibSslTrust`.
// To enable Agency creation 1) provide a valid certificate used by the
// server hosting the `AgentConfigInfo.ovfPackageUrl` or
// `AgentConfigInfo.vibUrl` or 2) ensure the server's certificate is
// signed by a CA trusted by the system. Then retry the operation, vSphere
// ESX Agent Manager will retry the SSL trust verification and proceed with
// reaching the desired state.
//
// This structure may be used only with operations rendered under `/eam`.
type CertificateNotTrustedFault struct {
	EamAppFault

	// The URL that failed the SSL trust verification.
	Url string `xml:"url,omitempty" json:"url,omitempty"`
}

func init() {
	types.Add("eam:CertificateNotTrustedFault", reflect.TypeOf((*CertificateNotTrustedFault)(nil)).Elem())
	types.AddMinAPIVersionForType("eam:CertificateNotTrustedFault", "8.2")
}

type CertificateNotTrustedFaultFault CertificateNotTrustedFault

func init() {
	types.Add("eam:CertificateNotTrustedFaultFault", reflect.TypeOf((*CertificateNotTrustedFaultFault)(nil)).Elem())
}

// Base class for all cluster bound agents.
//
// This structure may be used only with operations rendered under `/eam`.
type ClusterAgentAgentIssue struct {
	AgencyIssue

	// The agent that has this issue.
	//
	// Refers instance of `Agent`.
	Agent types.ManagedObjectReference `xml:"agent" json:"agent"`
	// The cluster for which this issue is raised.
	//
	// Might be null if the cluster
	// is missing in vCenter Server inventory.
	//
	// Refers instance of `ComputeResource`.
	Cluster *types.ManagedObjectReference `xml:"cluster,omitempty" json:"cluster,omitempty"`
}

func init() {
	types.Add("eam:ClusterAgentAgentIssue", reflect.TypeOf((*ClusterAgentAgentIssue)(nil)).Elem())
}

// The cluster agent Virtual Machine cannot be deployed, because vSphere ESX
// Agent Manager is not able to make successful SSL trust verification of the
// server's certificate, when establishing connection to the provided
// `AgentConfigInfo.ovfPackageUrl`.
//
// Reasons for this might be that the
// certificate provided via the API `AgentConfigInfo.ovfSslTrust` or via
// the script /usr/lib/vmware-eam/bin/eam-utility.py
//   - is invalid.
//   - does not match the server's certificate.
//
// If there is no provided certificate, the issue is raised when the server's
// certificate is not trusted by the system or invalid - @see
// `AgentConfigInfo.ovfSslTrust`.
// To remediate the cluster agent Virtual Machine deployment 1) provide a valid
// certificate used by the server hosting the
// `AgentConfigInfo.ovfPackageUrl` or 2) ensure the server's certificate
// is signed by a CA trusted by the system. Then resolve this issue, vSphere ESX
// Agent Manager will retry the SSL trust verification and proceed with reaching
// the desired state.
//
// This structure may be used only with operations rendered under `/eam`.
type ClusterAgentCertificateNotTrusted struct {
	ClusterAgentVmNotDeployed

	// The URL that failed the SSL trust verification.
	Url string `xml:"url" json:"url"`
}

func init() {
	types.Add("eam:ClusterAgentCertificateNotTrusted", reflect.TypeOf((*ClusterAgentCertificateNotTrusted)(nil)).Elem())
	types.AddMinAPIVersionForType("eam:ClusterAgentCertificateNotTrusted", "8.2")
}

type ClusterAgentClusterTransitionFailed struct {
	ClusterAgentAgentIssue
}

func init() {
	types.Add("eam:ClusterAgentClusterTransitionFailed", reflect.TypeOf((*ClusterAgentClusterTransitionFailed)(nil)).Elem())
}

// An agent virtual machine operation cannot be executed on host, because the
// host is in maintenance mode that blocks the virtual machine operation.
//
// This is not a remediable issue. To remediate, take the host ouf of
// maintenance mode.
//
// This structure may be used only with operations rendered under `/eam`.
type ClusterAgentHostInMaintenanceMode struct {
	ClusterAgentVmIssue
}

func init() {
	types.Add("eam:ClusterAgentHostInMaintenanceMode", reflect.TypeOf((*ClusterAgentHostInMaintenanceMode)(nil)).Elem())
	types.AddMinAPIVersionForType("eam:ClusterAgentHostInMaintenanceMode", "8.3")
}

// An agent virtual machine operation cannot be executed on host, because the
// host is in partial maintenance mode that blocks the virtual machine
// operation.
//
// This is not a remediable issue. To remediate, take the host ouf of partial
// maintenance mode.
//
// This structure may be used only with operations rendered under `/eam`.
type ClusterAgentHostInPartialMaintenanceMode struct {
	ClusterAgentVmIssue
}

func init() {
	types.Add("eam:ClusterAgentHostInPartialMaintenanceMode", reflect.TypeOf((*ClusterAgentHostInPartialMaintenanceMode)(nil)).Elem())
	types.AddMinAPIVersionForType("eam:ClusterAgentHostInPartialMaintenanceMode", "8.3")
}

// The cluster agent Virtual Machine could not be powered-on, because the
// cluster does not have enough CPU or memory resources.
//
// This is an active remediable issue. To remediate, vSphere ESX Agent Manager
// powers on the agent Virtual Machine. However, the problem is likely to
// persist until enough CPU and memory resources are made available.
//
// This structure may be used only with operations rendered under `/eam`.
type ClusterAgentInsufficientClusterResources struct {
	ClusterAgentVmPoweredOff
}

func init() {
	types.Add("eam:ClusterAgentInsufficientClusterResources", reflect.TypeOf((*ClusterAgentInsufficientClusterResources)(nil)).Elem())
}

// The cluster agent Virtual Machine cannot be deployed, because any of the
// configured datastores does not have enough disk space.
//
// This is an active remediable issue. To remediate, vSphere ESX Agent Manager
// redeploys the agent Virtual Machine. However, the problem is likely to
// persist until enough disk space is freed up on the cluster datastore.
//
// This structure may be used only with operations rendered under `/eam`.
type ClusterAgentInsufficientClusterSpace struct {
	ClusterAgentVmNotDeployed
}

func init() {
	types.Add("eam:ClusterAgentInsufficientClusterSpace", reflect.TypeOf((*ClusterAgentInsufficientClusterSpace)(nil)).Elem())
}

// Invalid configuration is preventing a cluster agent virtual machine
// operation.
//
// Typically the attached error indicates the particular reason why
// vSphere ESX Agent Manager is unable to power on or reconfigure the agent
// virtual machine.
//
// This is a passive remediable issue. To remediate update the virtual machine
// configuration.
//
// This structure may be used only with operations rendered under `/eam`.
type ClusterAgentInvalidConfig struct {
	ClusterAgentVmIssue

	// The error, that caused this issue.
	//
	// It must be either MethodFault or
	// RuntimeFault.
	Error types.AnyType `xml:"error,typeattr" json:"error"`
}

func init() {
	types.Add("eam:ClusterAgentInvalidConfig", reflect.TypeOf((*ClusterAgentInvalidConfig)(nil)).Elem())
}

// The cluster agent Virtual Machine cannot be deployed, because any of the
// configured datastores does not exist on the cluster.
//
// This is an active remediable issue. To remediate, vSphere ESX Agent Manager
// redeploys the agent Virtual Machine. However, the problem is likely to
// persist until required Virtual Machine datastores are configured on the
// cluster.
//
// This structure may be used only with operations rendered under `/eam`.
type ClusterAgentMissingClusterVmDatastore struct {
	ClusterAgentVmNotDeployed

	// A non-empty array of cluster agent VM datastores that are missing on the
	// cluster.
	//
	// Refers instances of `Datastore`.
	MissingDatastores []types.ManagedObjectReference `xml:"missingDatastores,omitempty" json:"missingDatastores,omitempty"`
}

func init() {
	types.Add("eam:ClusterAgentMissingClusterVmDatastore", reflect.TypeOf((*ClusterAgentMissingClusterVmDatastore)(nil)).Elem())
}

// The cluster agent Virtual Machine cannot be deployed, because the configured
// networks do not exist on the cluster.
//
// This is an active remediable issue. To remediate, vSphere ESX Agent Manager
// redeploys the agent Virtual Machine. However, the problem is likely to
// persist until required Virtual Machine networks are configured on the
// cluster.
//
// This structure may be used only with operations rendered under `/eam`.
type ClusterAgentMissingClusterVmNetwork struct {
	ClusterAgentVmNotDeployed

	// A non-empty array of cluster agent VM networks that are required on the
	// cluster.
	//
	// Refers instances of `Network`.
	MissingNetworks []types.ManagedObjectReference `xml:"missingNetworks,omitempty" json:"missingNetworks,omitempty"`
	// The names of the cluster agent VM networks.
	NetworkNames []string `xml:"networkNames,omitempty" json:"networkNames,omitempty"`
}

func init() {
	types.Add("eam:ClusterAgentMissingClusterVmNetwork", reflect.TypeOf((*ClusterAgentMissingClusterVmNetwork)(nil)).Elem())
}

// A cluster agent virtual machine needs to be provisioned, but an OVF property
// is either missing or has an invalid value.
//
// This is a passive remediable issue. To remediate, update the OVF environment
// in the agent configuration used to provision the agent virtual machine.
//
// This structure may be used only with operations rendered under `/eam`.
type ClusterAgentOvfInvalidProperty struct {
	ClusterAgentAgentIssue

	// An optional list of errors that caused this issue.
	//
	// These errors are
	// generated by the vCenter server.
	Error []types.LocalizedMethodFault `xml:"error,omitempty" json:"error,omitempty"`
}

func init() {
	types.Add("eam:ClusterAgentOvfInvalidProperty", reflect.TypeOf((*ClusterAgentOvfInvalidProperty)(nil)).Elem())
}

// A cluster agent failed to be transitioned to a LCCM Solution.
//
// This is an active remediable issue. To remediate, resolve the issue via vLCM
// System VMs API
//
// This structure may be used only with operations rendered under `/eam`.
type ClusterAgentTransitionFailed struct {
	ClusterAgentAgentIssue
}

func init() {
	types.Add("eam:ClusterAgentTransitionFailed", reflect.TypeOf((*ClusterAgentTransitionFailed)(nil)).Elem())
	types.AddMinAPIVersionForType("eam:ClusterAgentTransitionFailed", "9.0")
}

type ClusterAgentVmHookDynamicUpdateFailed struct {
	ClusterAgentVmHookFailed
}

func init() {
	types.Add("eam:ClusterAgentVmHookDynamicUpdateFailed", reflect.TypeOf((*ClusterAgentVmHookDynamicUpdateFailed)(nil)).Elem())
}

// The VM hook remediation failed.
//
// In order to remediate the issue:
// Resolve the issue via apply API and process the hook within the
// timeout configured for the System VM Solution this issue's VM belongs to.
//
// This structure may be used only with operations rendered under `/eam`.
type ClusterAgentVmHookFailed struct {
	ClusterAgentVmIssue
}

func init() {
	types.Add("eam:ClusterAgentVmHookFailed", reflect.TypeOf((*ClusterAgentVmHookFailed)(nil)).Elem())
	types.AddMinAPIVersionForType("eam:ClusterAgentVmHookFailed", "9.0")
}

// The VM hook remediation timed out.
//
// In order to remediate the issue:
// Resolve the issue via apply API and process the hook within the
// timeout configured for the System VM Solution this issue's VM belongs to.
//
// This structure may be used only with operations rendered under `/eam`.
type ClusterAgentVmHookTimedout struct {
	ClusterAgentVmIssue
}

func init() {
	types.Add("eam:ClusterAgentVmHookTimedout", reflect.TypeOf((*ClusterAgentVmHookTimedout)(nil)).Elem())
	types.AddMinAPIVersionForType("eam:ClusterAgentVmHookTimedout", "9.0")
}

// The connection state of the cluster agent Virtual Machine is
// `inaccessible`.
//
// In order to remediate the issue:
//   - Mark the VM for removal using the `EsxAgentManager.EsxAgentManager_MarkForRemoval`
//     API.
//   - Do the necessary changes to ensure that the connection state of the VM is
//     `connected`.
//
// NOTE: When the HA is enabled on the cluster these issues may be transient and
// automatically remediated.
//
// This structure may be used only with operations rendered under `/eam`.
type ClusterAgentVmInaccessible struct {
	ClusterAgentVmIssue
}

func init() {
	types.Add("eam:ClusterAgentVmInaccessible", reflect.TypeOf((*ClusterAgentVmInaccessible)(nil)).Elem())
	types.AddMinAPIVersionForType("eam:ClusterAgentVmInaccessible", "8.3")
}

// Base class for all cluster bound Virtual Machines.
//
// This structure may be used only with operations rendered under `/eam`.
type ClusterAgentVmIssue struct {
	ClusterAgentAgentIssue

	// The Virtual Machine to which this issue is related.
	//
	// Refers instance of `VirtualMachine`.
	Vm types.ManagedObjectReference `xml:"vm" json:"vm"`
}

func init() {
	types.Add("eam:ClusterAgentVmIssue", reflect.TypeOf((*ClusterAgentVmIssue)(nil)).Elem())
}

// A cluster agent Virtual Machine is expected to be deployed on a cluster, but
// the cluster agent Virtual Machine has not been deployed or has been explicitly
// deleted from the cluster.
//
// Typically more specific issue (a subclass of this
// issue) indicates the particular reason why vSphere ESX Agent Manager was
// unable to deploy the cluster agent Virtual Machine.
//
// This is an active remediable issue. To remediate, vSphere ESX Agent Manager
// redeploys the cluster agent Virtual Machine.
//
// This structure may be used only with operations rendered under `/eam`.
type ClusterAgentVmNotDeployed struct {
	ClusterAgentAgentIssue
}

func init() {
	types.Add("eam:ClusterAgentVmNotDeployed", reflect.TypeOf((*ClusterAgentVmNotDeployed)(nil)).Elem())
}

// The cluster agent Virtual Machine can not be removed from a cluster.
//
// Typically the description indicates the particular reason why vSphere ESX
// Agent Manager was unable to remove the cluster agent Virtual Machine.
//
// This is an active remediable issue. To remediate, vSphere ESX Agent Manager
// removes the cluster agent Virtual Machine.
//
// This structure may be used only with operations rendered under `/eam`.
type ClusterAgentVmNotRemoved struct {
	ClusterAgentVmIssue
}

func init() {
	types.Add("eam:ClusterAgentVmNotRemoved", reflect.TypeOf((*ClusterAgentVmNotRemoved)(nil)).Elem())
}

// A cluster agent Virtual Machine is expected to be powered on, but the agent
// Virtual Machine is powered off.
//
// Typically more specific issue (a subclass of
// this issue) indicates the particular reason why vSphere ESX Agent Manager was
// unable to power on the cluster agent Virtual Machine.
//
// This is an active remediable issue. To remediate, vSphere ESX Agent Manager
// powers on the cluster agent Virtual Machine.
//
// This structure may be used only with operations rendered under `/eam`.
type ClusterAgentVmPoweredOff struct {
	ClusterAgentVmIssue
}

func init() {
	types.Add("eam:ClusterAgentVmPoweredOff", reflect.TypeOf((*ClusterAgentVmPoweredOff)(nil)).Elem())
}

// A cluster agent virtual machine is expected to be powered off, but the agent
// virtual machine is powered on.
//
// This is an active remediable issue. To remediate, vSphere ESX Agent Manager
// powers off the agent virtual machine.
//
// This structure may be used only with operations rendered under `/eam`.
type ClusterAgentVmPoweredOn struct {
	ClusterAgentVmIssue
}

func init() {
	types.Add("eam:ClusterAgentVmPoweredOn", reflect.TypeOf((*ClusterAgentVmPoweredOn)(nil)).Elem())
}

// An agent virtual machine is protected from modifications
// (example: HA recovery).
//
// This is an active remediable issue. To remediate, vSphere ESX Agent Manager
// modifies the virtual machine.
//
// This structure may be used only with operations rendered under `/eam`.
type ClusterAgentVmProtected struct {
	ClusterAgentVmIssue
}

func init() {
	types.Add("eam:ClusterAgentVmProtected", reflect.TypeOf((*ClusterAgentVmProtected)(nil)).Elem())
	types.AddMinAPIVersionForType("eam:ClusterAgentVmProtected", "9.0")
}

// A cluster agent Virtual Machine is expected to be powered on, but the agent
// Virtual Machine is suspended.
//
// This is an active remediable issue. To remediate, vSphere ESX Agent Manager
// powers on the cluster agent Virtual Machine.
//
// This structure may be used only with operations rendered under `/eam`.
type ClusterAgentVmSuspended struct {
	ClusterAgentVmIssue
}

func init() {
	types.Add("eam:ClusterAgentVmSuspended", reflect.TypeOf((*ClusterAgentVmSuspended)(nil)).Elem())
}

type CreateAgency CreateAgencyRequestType

func init() {
	types.Add("eam:CreateAgency", reflect.TypeOf((*CreateAgency)(nil)).Elem())
}

// The parameters of `EsxAgentManager.CreateAgency`.
//
// This structure may be used only with operations rendered under `/eam`.
type CreateAgencyRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// The configuration that describes how to deploy the agents in the
	// created agency.
	AgencyConfigInfo BaseAgencyConfigInfo `xml:"agencyConfigInfo,typeattr" json:"agencyConfigInfo"`
	// Deprecated. No sense to create agency in other state than
	// <code>enabled</code>. <code>disabled</code> is deprecated
	// whereas <code>uninstalled</code> is useless.
	// The initial goal state of the agency. See
	// `EamObjectRuntimeInfoGoalState_enum`.
	InitialGoalState string `xml:"initialGoalState" json:"initialGoalState"`
}

func init() {
	types.Add("eam:CreateAgencyRequestType", reflect.TypeOf((*CreateAgencyRequestType)(nil)).Elem())
}

type CreateAgencyResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}

type DestroyAgency DestroyAgencyRequestType

func init() {
	types.Add("eam:DestroyAgency", reflect.TypeOf((*DestroyAgency)(nil)).Elem())
}

type DestroyAgencyRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
}

func init() {
	types.Add("eam:DestroyAgencyRequestType", reflect.TypeOf((*DestroyAgencyRequestType)(nil)).Elem())
}

type DestroyAgencyResponse struct {
}

// Thrown when trying to modify state over disabled clusters.
//
// This structure may be used only with operations rendered under `/eam`.
type DisabledClusterFault struct {
	EamAppFault

	// The MoRefs of the disabled compute resources.
	//
	// Refers instances of `ComputeResource`.
	DisabledComputeResource []types.ManagedObjectReference `xml:"disabledComputeResource,omitempty" json:"disabledComputeResource,omitempty"`
}

func init() {
	types.Add("eam:DisabledClusterFault", reflect.TypeOf((*DisabledClusterFault)(nil)).Elem())
	types.AddMinAPIVersionForType("eam:DisabledClusterFault", "7.6")
}

type DisabledClusterFaultFault DisabledClusterFault

func init() {
	types.Add("eam:DisabledClusterFaultFault", reflect.TypeOf((*DisabledClusterFaultFault)(nil)).Elem())
}

// Application related error
// As opposed to system errors, application ones are always function of the
// input and the current state.
//
// They occur always upon same conditions. In most
// of the cases they are recoverable, i.e. the client can determine what is
// wrong and know how to recover.
// NOTE: Since there is not yet need to distinguish among specific error
// sub-types then we define a common type. Tomorrow, if necessary, we can add an
// additional level of detailed exception types and make this one abstract.
//
// This structure may be used only with operations rendered under `/eam`.
type EamAppFault struct {
	EamRuntimeFault
}

func init() {
	types.Add("eam:EamAppFault", reflect.TypeOf((*EamAppFault)(nil)).Elem())
}

type EamAppFaultFault BaseEamAppFault

func init() {
	types.Add("eam:EamAppFaultFault", reflect.TypeOf((*EamAppFaultFault)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. Please refer to vLCM APIs.
//
// The common base type for all vSphere ESX Agent Manager exceptions.
//
// # TODO migrate to EamRuntimeFault
//
// This structure may be used only with operations rendered under `/eam`.
type EamFault struct {
	types.MethodFault
}

func init() {
	types.Add("eam:EamFault", reflect.TypeOf((*EamFault)(nil)).Elem())
}

type EamFaultFault BaseEamFault

func init() {
	types.Add("eam:EamFaultFault", reflect.TypeOf((*EamFaultFault)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. Please refer to vLCM APIs.
//
// IO error
// NOTE: Since this type is a first of system-type errors we do not introduce a
// common base type for them.
//
// Once add a second system type exception though, it
// should be introduced.
//
// This structure may be used only with operations rendered under `/eam`.
type EamIOFault struct {
	EamRuntimeFault
}

func init() {
	types.Add("eam:EamIOFault", reflect.TypeOf((*EamIOFault)(nil)).Elem())
}

type EamIOFaultFault EamIOFault

func init() {
	types.Add("eam:EamIOFaultFault", reflect.TypeOf((*EamIOFaultFault)(nil)).Elem())
}

// Thrown when a user cannot be authenticated.
//
// This structure may be used only with operations rendered under `/eam`.
type EamInvalidLogin struct {
	EamRuntimeFault
}

func init() {
	types.Add("eam:EamInvalidLogin", reflect.TypeOf((*EamInvalidLogin)(nil)).Elem())
}

type EamInvalidLoginFault EamInvalidLogin

func init() {
	types.Add("eam:EamInvalidLoginFault", reflect.TypeOf((*EamInvalidLoginFault)(nil)).Elem())
}

// Thrown when a user is not allowed to execute an operation.
//
// This structure may be used only with operations rendered under `/eam`.
type EamInvalidState struct {
	EamAppFault
}

func init() {
	types.Add("eam:EamInvalidState", reflect.TypeOf((*EamInvalidState)(nil)).Elem())
}

type EamInvalidStateFault EamInvalidState

func init() {
	types.Add("eam:EamInvalidStateFault", reflect.TypeOf((*EamInvalidStateFault)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. Please refer to vLCM Image APIs.
//
// Indicates for an invalid or unknown Vib package structure and/or format.
//
// This structure may be used only with operations rendered under `/eam`.
type EamInvalidVibPackage struct {
	EamRuntimeFault
}

func init() {
	types.Add("eam:EamInvalidVibPackage", reflect.TypeOf((*EamInvalidVibPackage)(nil)).Elem())
}

type EamInvalidVibPackageFault EamInvalidVibPackage

func init() {
	types.Add("eam:EamInvalidVibPackageFault", reflect.TypeOf((*EamInvalidVibPackageFault)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. Please refer to vLCM APIs.
//
// The <code>RuntimeInfo</code> represents the runtime information of the vSphere ESX Agent
// Manager managed
// objects `Agency` and `Agent`.
//
// The runtime information provides
// two kinds of information, namely, the
// desired goal state of the entity and the status with regards to conforming
// to that goal state.
//
// This structure may be used only with operations rendered under `/eam`.
type EamObjectRuntimeInfo struct {
	types.DynamicData

	// The health of the managed entity.
	//
	// This denotes how well the entity conforms to the
	// goal state.
	//
	// See also `EamObjectRuntimeInfoStatus_enum`.
	Status string `xml:"status" json:"status"`
	// Current issues that have been detected for this entity.
	//
	// Each issue can be remediated
	// by invoking `EamObject.Resolve` or `EamObject.ResolveAll`.
	Issue []BaseIssue `xml:"issue,omitempty,typeattr" json:"issue,omitempty"`
	// The desired state of the entity.
	//
	// See also `EamObjectRuntimeInfoGoalState_enum`.
	GoalState string `xml:"goalState" json:"goalState"`
	// The `Agent` or `Agency` with which this <code>RuntimeInfo</code> object is associated.
	//
	// Refers instance of `EamObject`.
	Entity types.ManagedObjectReference `xml:"entity" json:"entity"`
}

func init() {
	types.Add("eam:EamObjectRuntimeInfo", reflect.TypeOf((*EamObjectRuntimeInfo)(nil)).Elem())
}

// The common base type for all vSphere ESX Agent Manager runtime exceptions.
//
// This structure may be used only with operations rendered under `/eam`.
type EamRuntimeFault struct {
	types.RuntimeFault
}

func init() {
	types.Add("eam:EamRuntimeFault", reflect.TypeOf((*EamRuntimeFault)(nil)).Elem())
}

type EamRuntimeFaultFault BaseEamRuntimeFault

func init() {
	types.Add("eam:EamRuntimeFaultFault", reflect.TypeOf((*EamRuntimeFaultFault)(nil)).Elem())
}

// Thrown when calling vSphere ESX Agent Manager when it is not fully
// initialized.
//
// This structure may be used only with operations rendered under `/eam`.
type EamServiceNotInitialized struct {
	EamRuntimeFault
}

func init() {
	types.Add("eam:EamServiceNotInitialized", reflect.TypeOf((*EamServiceNotInitialized)(nil)).Elem())
}

type EamServiceNotInitializedFault EamServiceNotInitialized

func init() {
	types.Add("eam:EamServiceNotInitializedFault", reflect.TypeOf((*EamServiceNotInitializedFault)(nil)).Elem())
}

// System fault.
//
// This structure may be used only with operations rendered under `/eam`.
type EamSystemFault struct {
	EamRuntimeFault
}

func init() {
	types.Add("eam:EamSystemFault", reflect.TypeOf((*EamSystemFault)(nil)).Elem())
}

type EamSystemFaultFault EamSystemFault

func init() {
	types.Add("eam:EamSystemFaultFault", reflect.TypeOf((*EamSystemFaultFault)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. No longer in use since vSphee 6.5.
//
// Extensible issue class used by solutions to add custom issues to agents.
//
// When resolved, the issue is removed from the agent and an event is generated.
//
// This structure may be used only with operations rendered under `/eam`.
type ExtensibleIssue struct {
	Issue

	// Unique string for this type of issue.
	//
	// The type must match an event registered
	// by the solution as part of its extension.
	TypeId string `xml:"typeId" json:"typeId"`
	// Arguments associated with the typeId.
	Argument []types.KeyAnyValue `xml:"argument,omitempty" json:"argument,omitempty"`
	// A managed object reference to the object this issue is related to.
	//
	// Refers instance of `ManagedEntity`.
	Target *types.ManagedObjectReference `xml:"target,omitempty" json:"target,omitempty"`
	// An optional agent this issue pertains
	//
	// Refers instance of `Agent`.
	Agent *types.ManagedObjectReference `xml:"agent,omitempty" json:"agent,omitempty"`
	// An optional agency this issue pertains
	//
	// Refers instance of `Agency`.
	Agency *types.ManagedObjectReference `xml:"agency,omitempty" json:"agency,omitempty"`
}

func init() {
	types.Add("eam:ExtensibleIssue", reflect.TypeOf((*ExtensibleIssue)(nil)).Elem())
}

type GetMaintenanceModePolicy GetMaintenanceModePolicyRequestType

func init() {
	types.Add("eam:GetMaintenanceModePolicy", reflect.TypeOf((*GetMaintenanceModePolicy)(nil)).Elem())
}

type GetMaintenanceModePolicyRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
}

func init() {
	types.Add("eam:GetMaintenanceModePolicyRequestType", reflect.TypeOf((*GetMaintenanceModePolicyRequestType)(nil)).Elem())
}

type GetMaintenanceModePolicyResponse struct {
	Returnval string `xml:"returnval" json:"returnval"`
}

type HooksDynamicUpdateSpec struct {
	types.DynamicData

	Vm                types.ManagedObjectReference `xml:"vm" json:"vm"`
	HookType          string                       `xml:"hookType" json:"hookType"`
	AlternativeVmSpec *SolutionsAlternativeVmSpec  `xml:"alternativeVmSpec,omitempty" json:"alternativeVmSpec,omitempty"`
}

func init() {
	types.Add("eam:HooksDynamicUpdateSpec", reflect.TypeOf((*HooksDynamicUpdateSpec)(nil)).Elem())
}

// Limits the hooks reported to the user.
//
// This structure may be used only with operations rendered under `/eam`.
type HooksHookListSpec struct {
	types.DynamicData

	// If specified - will report hooks only for agents from the specified
	// solutions, otherwise - will report hooks for agents from all solutions.
	Solutions []string `xml:"solutions,omitempty" json:"solutions,omitempty"`
	// If specified - will report hooks only for agents on the specified
	// hosts, otherwise - will report hooks for agents on all hosts.
	//
	// Refers instances of `HostSystem`.
	Hosts []types.ManagedObjectReference `xml:"hosts,omitempty" json:"hosts,omitempty"`
}

func init() {
	types.Add("eam:HooksHookListSpec", reflect.TypeOf((*HooksHookListSpec)(nil)).Elem())
}

// Specification for marking a raised hook on an agent Virtual Machine as
// processed.
//
// This structure may be used only with operations rendered under `/eam`.
type HooksMarkAsProcessedSpec struct {
	types.DynamicData

	// Virtual Machine to mark a hook as processed.
	//
	// Refers instance of `VirtualMachine`.
	Vm types.ManagedObjectReference `xml:"vm" json:"vm"`
	// Type of hook to be marked as processed `HooksHookType_enum`.
	HookType string `xml:"hookType" json:"hookType"`
	// `True` - if the hook was processed successfully, `False` -
	// if the hook could not be processed.
	Success bool `xml:"success" json:"success"`
}

func init() {
	types.Add("eam:HooksMarkAsProcessedSpec", reflect.TypeOf((*HooksMarkAsProcessedSpec)(nil)).Elem())
}

// An agent virtual machine operation is expected to be initiated on host, but
// the agent virtual machine operation has not been initiated.
//
// The reason is
// that the host is in maintenance mode.
//
// This is an active remediable issue. To remediate, vSphere ESX Agent Manager
// takes the host out of maintenance mode and initiates the agent virtual
// machine operation.
//
// Resolving this issue in vSphere Lifecycle Manager environment will be no-op.
// In those cases user must take the host out of Maintenance Mode manually or
// wait vSphere Lifecycle Manager cluster remediation to complete (if any).
//
// This structure may be used only with operations rendered under `/eam`.
type HostInMaintenanceMode struct {
	VmDeployed
}

func init() {
	types.Add("eam:HostInMaintenanceMode", reflect.TypeOf((*HostInMaintenanceMode)(nil)).Elem())
}

// An agent virtual machine operation cannot be executed on host, because the
// host is in partial maintenance mode that blocks the virtual machine
// operation.
//
// This is not a remediable issue. To remediate, take the host ouf of partial
// maintenance mode.
//
// This structure may be used only with operations rendered under `/eam`.
type HostInPartialMaintenanceMode struct {
	AgentIssue

	// The virtual machine to which this issue is related.
	//
	// Unset if the
	// operation that is blocked by partial maintenance mode is preventing the
	// virtual machine deployment.
	//
	// Refers instance of `VirtualMachine`.
	Vm *types.ManagedObjectReference `xml:"vm,omitempty" json:"vm,omitempty"`
}

func init() {
	types.Add("eam:HostInPartialMaintenanceMode", reflect.TypeOf((*HostInPartialMaintenanceMode)(nil)).Elem())
	types.AddMinAPIVersionForType("eam:HostInPartialMaintenanceMode", "8.3")
}

// An agent virtual machine is expected to be removed from a host, but the agent virtual machine has not
// been removed.
//
// The reason is that the host is in standby mode.
//
// This is an active remediable issue. To remediate, vSphere ESX Agent Manager puts the host in standby mode
// and removes the agent virtual machine.
//
// This structure may be used only with operations rendered under `/eam`.
type HostInStandbyMode struct {
	VmDeployed
}

func init() {
	types.Add("eam:HostInStandbyMode", reflect.TypeOf((*HostInStandbyMode)(nil)).Elem())
}

// Deprecated all host issues were removed.
//
// Base class for all host issues.
//
// This structure may be used only with operations rendered under `/eam`.
type HostIssue struct {
	Issue

	// The host to which the issue is related.
	//
	// Refers instance of `HostSystem`.
	Host types.ManagedObjectReference `xml:"host" json:"host"`
}

func init() {
	types.Add("eam:HostIssue", reflect.TypeOf((*HostIssue)(nil)).Elem())
}

// Deprecated hostPoweredOff will no longer be used, instead
// `ManagedHostNotReachable` will be raised.
//
// An agent virtual machine is expected to be removed from a host, but the agent
// virtual machine has not been removed.
//
// The reason is that the host is powered
// off.
//
// This is not a remediable issue. To remediate, power on the host.
//
// This structure may be used only with operations rendered under `/eam`.
type HostPoweredOff struct {
	VmDeployed
}

func init() {
	types.Add("eam:HostPoweredOff", reflect.TypeOf((*HostPoweredOff)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. Please refer to vLCM Image APIs.
//
// Live VIB operation failed.
//
// An immediate reboot is required to clear live VIB
// operation failure.
//
// This is an active remediable issue. To remediate, vSphere ESX Agent Manager
// puts the host into maintenance mode and reboots it.
//
// This structure may be used only with operations rendered under `/eam`.
type ImmediateHostRebootRequired struct {
	VibIssue
}

func init() {
	types.Add("eam:ImmediateHostRebootRequired", reflect.TypeOf((*ImmediateHostRebootRequired)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. Please refer to vLCM APIs.
//
// An agent virtual machine is expected to be deployed on a host, but the agent could not be
// deployed because it was incompatible with the host.
//
// This is an active remediable issue. To remediate, vSphere ESX Agent Manager redeployes the agent. However,
// the problem is likely to persist until either the host or the solution has been
// upgraded, so that the agent will become compatible with the host.
//
// This structure may be used only with operations rendered under `/eam`.
type IncompatibleHostVersion struct {
	VmNotDeployed
}

func init() {
	types.Add("eam:IncompatibleHostVersion", reflect.TypeOf((*IncompatibleHostVersion)(nil)).Elem())
}

// Deprecated this issue is no longer raised by EAM. It is replaced by
// `InvalidConfig`.
//
// An agent virtual machine is expected to be powered on, but there are no free IP addresses in the
// agent's pool of virtual machine IP addresses.
//
// To remediate, free some IP addresses or add some more to the IP pool and invoke
// <code>resolve</code>.
//
// This structure may be used only with operations rendered under `/eam`.
type InsufficientIpAddresses struct {
	VmPoweredOff

	// The agent virtual machine network.
	//
	// Refers instance of `Network`.
	Network types.ManagedObjectReference `xml:"network" json:"network"`
}

func init() {
	types.Add("eam:InsufficientIpAddresses", reflect.TypeOf((*InsufficientIpAddresses)(nil)).Elem())
}

// An agent virtual machine is expected to be deployed on a host, but the agent virtual machine could not be
// deployed because the host does not have enough free CPU or memory resources.
//
// This is an active remediable issue. To remediate, vSphere ESX Agent Manager redeploys the agent virtual machine. However,
// the problem is likely to persist until enough CPU and memory resources are made available.
//
// This structure may be used only with operations rendered under `/eam`.
type InsufficientResources struct {
	VmNotDeployed
}

func init() {
	types.Add("eam:InsufficientResources", reflect.TypeOf((*InsufficientResources)(nil)).Elem())
}

// An agent virtual machine is expected to be deployed on a host, but the agent virtual machine could not be
// deployed because the host's agent datastore did not have enough free space.
//
// This is an active remediable issue. To remediate, vSphere ESX Agent Manager redeploys the agent virtual machine. However,
// the problem is likely to persist until either space is freed up on the host's agent
// virtual machine datastore or a new agent virtual machine datastore with enough free space is configured.
//
// This structure may be used only with operations rendered under `/eam`.
type InsufficientSpace struct {
	VmNotDeployed
}

func init() {
	types.Add("eam:InsufficientSpace", reflect.TypeOf((*InsufficientSpace)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. Please refer to vLCM Image APIs.
//
// Cannot remove the Baseline associated with an Agency from VUM.
//
// This is an active remediable issue. To remediate, vSphere ESX Agent Manager
// retries the delete operation.
// Note: In future this issue will denote also the removal of the Agency
// software (VIBs) from VUM software depot once VUM provides an API for that.
//
// This structure may be used only with operations rendered under `/eam`.
type IntegrityAgencyCannotDeleteSoftware struct {
	IntegrityAgencyVUMIssue
}

func init() {
	types.Add("eam:IntegrityAgencyCannotDeleteSoftware", reflect.TypeOf((*IntegrityAgencyCannotDeleteSoftware)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. Please refer to vLCM Image APIs.
//
// The software defined by an Agency cannot be staged in VUM.
//
// The staging
// operation consists of the following steps:
//   - Upload the Agency software (VIBs) to the VUM depot.
//   - Create or update a VUM Baseline with the Agency software and scope.
//
// This is an active remediable issue. To remediate, vSphere ESX Agent Manager
// retries the stage operation.
//
// This structure may be used only with operations rendered under `/eam`.
type IntegrityAgencyCannotStageSoftware struct {
	IntegrityAgencyVUMIssue
}

func init() {
	types.Add("eam:IntegrityAgencyCannotStageSoftware", reflect.TypeOf((*IntegrityAgencyCannotStageSoftware)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. Please refer to vLCM Image APIs.
//
// Base class for all issues which occurred during EAM communication with
// vSphere Update Manager (VUM).
//
// This structure may be used only with operations rendered under `/eam`.
type IntegrityAgencyVUMIssue struct {
	AgencyIssue
}

func init() {
	types.Add("eam:IntegrityAgencyVUMIssue", reflect.TypeOf((*IntegrityAgencyVUMIssue)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. Please refer to vLCM Image APIs.
//
// VUM service is not available - its registered SOAP endpoint cannot be
// accessed or it is malfunctioning.
//
// This is an active and passive remediable issue. To remediate, vSphere ESX
// Agent Manager retries to access VUM service.
//
// This structure may be used only with operations rendered under `/eam`.
type IntegrityAgencyVUMUnavailable struct {
	IntegrityAgencyVUMIssue
}

func init() {
	types.Add("eam:IntegrityAgencyVUMUnavailable", reflect.TypeOf((*IntegrityAgencyVUMUnavailable)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. Please refer to vLCM APIs.
//
// An <code>InvalidAgencyScope</code> fault is thrown when the scope in an
// `AgencyConfigInfo` is invalid.
//
// See also `AgencyConfigInfo`.
//
// This structure may be used only with operations rendered under `/eam`.
type InvalidAgencyScope struct {
	EamFault

	// The MoRefs of the unknown compute resources.
	//
	// Refers instances of `ComputeResource`.
	UnknownComputeResource []types.ManagedObjectReference `xml:"unknownComputeResource,omitempty" json:"unknownComputeResource,omitempty"`
}

func init() {
	types.Add("eam:InvalidAgencyScope", reflect.TypeOf((*InvalidAgencyScope)(nil)).Elem())
}

type InvalidAgencyScopeFault InvalidAgencyScope

func init() {
	types.Add("eam:InvalidAgencyScopeFault", reflect.TypeOf((*InvalidAgencyScopeFault)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. Please refer to vLCM APIs.
//
// An <code>InvalidAgentConfiguration</code> fault is thrown when the agent
// configuration of an agency configuration is empty or invalid.
//
// See also `AgencyConfigInfo`.
//
// This structure may be used only with operations rendered under `/eam`.
type InvalidAgentConfiguration struct {
	EamFault

	// An optional invalid agent configuration.
	InvalidAgentConfiguration *AgentConfigInfo `xml:"invalidAgentConfiguration,omitempty" json:"invalidAgentConfiguration,omitempty"`
	// The invalid field.
	InvalidField string `xml:"invalidField,omitempty" json:"invalidField,omitempty"`
}

func init() {
	types.Add("eam:InvalidAgentConfiguration", reflect.TypeOf((*InvalidAgentConfiguration)(nil)).Elem())
}

type InvalidAgentConfigurationFault InvalidAgentConfiguration

func init() {
	types.Add("eam:InvalidAgentConfigurationFault", reflect.TypeOf((*InvalidAgentConfigurationFault)(nil)).Elem())
}

// Invalid configuration is preventing a virtual machine operation.
//
// Typically
// the attached error indicates the particular reason why vSphere ESX Agent
// Manager is unable to power on or reconfigure the agent virtual machine.
//
// This is a passive remediable issue. To remediate update the virtual machine
// configuration.
//
// This structure may be used only with operations rendered under `/eam`.
type InvalidConfig struct {
	VmIssue

	// The error, that caused this issue.
	//
	// It must be either MethodFault or
	// RuntimeFault.
	Error types.AnyType `xml:"error,typeattr" json:"error"`
}

func init() {
	types.Add("eam:InvalidConfig", reflect.TypeOf((*InvalidConfig)(nil)).Elem())
}

// This exception is thrown if `VasaProviderSpec.url` is malformed.
//
// This structure may be used only with operations rendered under `/sms`.
type InvalidUrl struct {
	EamFault

	// Provider `VasaProviderSpec.url`
	Url               string `xml:"url" json:"url"`
	MalformedUrl      bool   `xml:"malformedUrl" json:"malformedUrl"`
	UnknownHost       bool   `xml:"unknownHost" json:"unknownHost"`
	ConnectionRefused bool   `xml:"connectionRefused" json:"connectionRefused"`
	ResponseCode      int32  `xml:"responseCode,omitempty" json:"responseCode,omitempty"`
}

func init() {
	types.Add("eam:InvalidUrl", reflect.TypeOf((*InvalidUrl)(nil)).Elem())
}

type InvalidUrlFault InvalidUrl

func init() {
	types.Add("eam:InvalidUrlFault", reflect.TypeOf((*InvalidUrlFault)(nil)).Elem())
}

// An issue represents a problem encountered while deploying and configuring agents
// in a vCenter installation.
//
// An issue conveys the type of problem and the
// entity on which the problem has been encountered. Most issues are related to agents,
// but they can also relate to an agency or a host.
//
// The set of issues provided by the vSphere ESX Agent Manager describes the discrepancy between
// the _desired_ agent deployment state, as defined by the agency configurations,
// and the _actual_ deployment. The (@link EamObject.RuntimeInfo.Status.status)
// of an agency or agent is green if it has reached its goal state. It is
// marked as yellow if the vSphere ESX Agent Manager is actively working to bring the object
// to its goal state. It is red if there is a discrepancy between the current state and
// the desired state. In the red state, a set of issues are filed on the object that
// describe the reason for the discrepancy between the desired and actual states.
//
// Issues are characterized as either active or passive remediable issues. For an active
// remediable issue, the vSphere ESX Agent Manager can actively try to solve the issue. For
// example, by deploying a new agent, removing an agent, changing its power state, and so
// on. For a passive remediable issue, the vSphere ESX Agent Manager is not able to solve the
// problem directly, and can only report the problem. For example, this could be
// caused by an incomplete host configuration.
//
// When <code>resolve</code> is called for an active remediable issue, the vSphere ESX Agent Manager
// starts performing the appropriate remediation steps for the particular issue. For a passive
// remediable issue, the EAM manager simply checks if the condition
// still exists, and if not it removes the issue.
//
// The vSphere ESX Agent Manager actively monitors most conditions relating to both
// active and passive issues. Thus, it often automatically discovers when an
// issue has been remediated and removes the issue without needing to explicitly
// call <code>resolve</code> on an issue.
//
// The complete Issue hierarchy is shown below:
//   - `Issue`
//   - `AgencyIssue`
//   - `AgentIssue`
//   - `ManagedHostNotReachable`
//   - `VmNotDeployed`
//   - `CannotAccessAgentOVF`
//   - `IncompatibleHostVersion`
//   - `InsufficientResources`
//   - `InsufficientSpace`
//   - `OvfInvalidFormat`
//   - `NoAgentVmDatastore`
//   - `NoAgentVmNetwork`
//   - `VmIssue`
//   - `OvfInvalidProperty`
//   - `VmDeployed`
//   - `HostInMaintenanceMode`
//   - `HostInStandbyMode`
//   - `VmCorrupted`
//   - `VmOrphaned`
//   - `VmPoweredOff`
//   - `InsufficientIpAddresses`
//   - `MissingAgentIpPool`
//   - `VmPoweredOn`
//   - `VmSuspended`
//   - `VibIssue`
//   - `VibCannotPutHostInMaintenanceMode`
//   - `VibNotInstalled`
//   - `CannotAccessAgentVib`
//   - `VibDependenciesNotMetByHost`
//   - `VibInvalidFormat`
//   - `VibRequirementsNotMetByHost`
//   - `VibRequiresHostInMaintenanceMode`
//   - `VibRequiresHostReboot`
//   - `VibRequiresManualInstallation`
//   - `VibRequiresManualUninstallation`
//   - `ImmediateHostRebootRequired`
//   - `OrphanedAgency`
//   - `IntegrityAgencyVUMIssue`
//   - `IntegrityAgencyVUMUnavailable`
//   - `IntegrityAgencyCannotStageSoftware`
//   - `IntegrityAgencyCannotDeleteSoftware`
//   - `ClusterAgentAgentIssue`
//   - `ClusterAgentVmIssue`
//   - `ClusterAgentVmNotRemoved`
//   - `ClusterAgentVmPoweredOff`
//   - `ClusterAgentInsufficientClusterResources`
//   - `ClusterAgentVmNotDeployed`
//   - `ClusterAgentInsufficientClusterSpace`
//   - `ClusterAgentMissingClusterVmDatastore`
//   - `ClusterAgentMissingClusterVmNetwork`
//
// See also `EamObject.Resolve`, `EamObject.ResolveAll`.
//
// This structure may be used only with operations rendered under `/eam`.
type Issue struct {
	types.DynamicData

	// A unique identifier per <code>Issue</code> instance.
	Key int32 `xml:"key" json:"key"`
	// A localized message describing the issue.
	Description string `xml:"description" json:"description"`
	// The point in time when this issue was generated.
	//
	// Note that issues can be
	// regenerated periodically, so this time does not necessarily reflect the
	// first time the issue was detected.
	Time time.Time `xml:"time" json:"time"`
}

func init() {
	types.Add("eam:Issue", reflect.TypeOf((*Issue)(nil)).Elem())
}

// Managed ESXi Server is unreachable from vCenter Server or vSphere ESX Agent
// Manager.
//
// Currently all operations on the affected host are impossible. Reasons
// for this might be :
//   - ESXi Server is not connected from vCenter Server
//   - ESXi Server powered off
//
// This is not a remediable issue. To remediate, connect, power on or reboot the
// host.
//
// This structure may be used only with operations rendered under `/eam`.
type ManagedHostNotReachable struct {
	AgentIssue
}

func init() {
	types.Add("eam:ManagedHostNotReachable", reflect.TypeOf((*ManagedHostNotReachable)(nil)).Elem())
}

type MarkAsAvailable MarkAsAvailableRequestType

func init() {
	types.Add("eam:MarkAsAvailable", reflect.TypeOf((*MarkAsAvailable)(nil)).Elem())
}

type MarkAsAvailableRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
}

func init() {
	types.Add("eam:MarkAsAvailableRequestType", reflect.TypeOf((*MarkAsAvailableRequestType)(nil)).Elem())
}

type MarkAsAvailableResponse struct {
}

// Deprecated this issue is no longer raised by EAM. It is replaced by
// `InvalidConfig`.
//
// An agent virtual machine is expected to be powered on, but the agent virtual machine is powered off because
// there there are no IP addresses defined on the agent's virtual machine network.
//
// To remediate, create an IP pool on the agent's virtual machine network and invoke <code>resolve</code>.
//
// This structure may be used only with operations rendered under `/eam`.
type MissingAgentIpPool struct {
	VmPoweredOff

	// The agent virtual machine network.
	//
	// Refers instance of `Network`.
	Network types.ManagedObjectReference `xml:"network" json:"network"`
}

func init() {
	types.Add("eam:MissingAgentIpPool", reflect.TypeOf((*MissingAgentIpPool)(nil)).Elem())
}

// Deprecated dvFilters are no longer supported by EAM.
//
// The agent is using the dvFilter API on the ESX host, but no dvFilter switch
// has been configured on the host.
//
// This can happen due to host communication
// failures or if the dvSwitch was (presumably accidentally) deleted from the
// host configuration.
//
// This is an active remediable issue. To remediate, vSphere ESX Agent Manager
// recreates the dvFilter switch.
//
// This structure may be used only with operations rendered under `/eam`.
type MissingDvFilterSwitch struct {
	AgentIssue
}

func init() {
	types.Add("eam:MissingDvFilterSwitch", reflect.TypeOf((*MissingDvFilterSwitch)(nil)).Elem())
}

// An agent virtual machine is expected to be deployed on a host, but the agent cannot be
// deployed because the agent datastore has not been configured on the host.
//
// This is a passive remediable issue. The administrator must configure
// the agent virtual machine datastore on the host.
//
// This structure may be used only with operations rendered under `/eam`.
type NoAgentVmDatastore struct {
	VmNotDeployed
}

func init() {
	types.Add("eam:NoAgentVmDatastore", reflect.TypeOf((*NoAgentVmDatastore)(nil)).Elem())
}

// An agent virtual machine is expected to be deployed on a host, but the agent cannot be
// deployed because the agent network has not been configured on the host.
//
// This is a passive remediable issue. The administrator must configure
// the agent virtual machine network on the host.
//
// This structure may be used only with operations rendered under `/eam`.
type NoAgentVmNetwork struct {
	VmNotDeployed
}

func init() {
	types.Add("eam:NoAgentVmNetwork", reflect.TypeOf((*NoAgentVmNetwork)(nil)).Elem())
}

// Thrown when calling vSphere ESX Agent Manager when it is not connected to the vCenter server.
//
// This structure may be used only with operations rendered under `/eam`.
type NoConnectionToVCenter struct {
	EamRuntimeFault
}

func init() {
	types.Add("eam:NoConnectionToVCenter", reflect.TypeOf((*NoConnectionToVCenter)(nil)).Elem())
}

type NoConnectionToVCenterFault NoConnectionToVCenter

func init() {
	types.Add("eam:NoConnectionToVCenterFault", reflect.TypeOf((*NoConnectionToVCenterFault)(nil)).Elem())
}

// An agent virtual machine is expected to be deployed on a host, but the agent cannot be
// deployed because the agent datastore has not been configured on the host.
//
// The host
// needs to be added to one of the datastores listed in customAgentVmDatastore.
//
// This is a passive remediable issue. The administrator must add one of the datastores
// customAgentVmDatastore to the host.
//
// This structure may be used only with operations rendered under `/eam`.
type NoCustomAgentVmDatastore struct {
	NoAgentVmDatastore

	// A non-empty array of agent VM datastores that is required on the host.
	//
	// Refers instances of `Datastore`.
	CustomAgentVmDatastore []types.ManagedObjectReference `xml:"customAgentVmDatastore" json:"customAgentVmDatastore"`
	// The names of the agent VM datastores.
	CustomAgentVmDatastoreName []string `xml:"customAgentVmDatastoreName" json:"customAgentVmDatastoreName"`
}

func init() {
	types.Add("eam:NoCustomAgentVmDatastore", reflect.TypeOf((*NoCustomAgentVmDatastore)(nil)).Elem())
}

// An agent virtual machine is expected to be deployed on a host, but the agent cannot be
// deployed because the agent network has not been configured on the host.
//
// The host
// needs to be added to one of the networks listed in customAgentVmNetwork.
//
// This is a passive remediable issue. The administrator must add one of the networks
// customAgentVmNetwork to the host.
//
// This structure may be used only with operations rendered under `/eam`.
type NoCustomAgentVmNetwork struct {
	NoAgentVmNetwork

	// A non-empty array of agent VM networks that is required on the host.
	//
	// Refers instances of `Network`.
	CustomAgentVmNetwork []types.ManagedObjectReference `xml:"customAgentVmNetwork" json:"customAgentVmNetwork"`
	// The names of the agent VM networks.
	CustomAgentVmNetworkName []string `xml:"customAgentVmNetworkName" json:"customAgentVmNetworkName"`
}

func init() {
	types.Add("eam:NoCustomAgentVmNetwork", reflect.TypeOf((*NoCustomAgentVmNetwork)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. No longer in use since vCLS 2.0.
//
// An agent virtual machine is expected to be deployed on a host, but the
// agent cannot be deployed because the agent VM datastore could not be
// discovered, as per defined selection policy, on the host.
//
// This issue can be remediated passively if the administrator configures
// new datastores on the host.
//
// This structure may be used only with operations rendered under `/eam`.
type NoDiscoverableAgentVmDatastore struct {
	VmNotDeployed
}

func init() {
	types.Add("eam:NoDiscoverableAgentVmDatastore", reflect.TypeOf((*NoDiscoverableAgentVmDatastore)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. No longer in use since vCLS 2.0.
//
// An agent virtual machine is expected to be deployed on a host, but the
// agent cannot be deployed because the agent VM network could not be
// discovered, as per defined selection policy, on the host.
//
// This issue can be remediated passively if the administrator configures
// new networks on the host.
//
// This structure may be used only with operations rendered under `/eam`.
type NoDiscoverableAgentVmNetwork struct {
	VmNotDeployed
}

func init() {
	types.Add("eam:NoDiscoverableAgentVmNetwork", reflect.TypeOf((*NoDiscoverableAgentVmNetwork)(nil)).Elem())
}

// Thrown when an a user cannot be authorized.
//
// This structure may be used only with operations rendered under `/eam`.
type NotAuthorized struct {
	EamRuntimeFault
}

func init() {
	types.Add("eam:NotAuthorized", reflect.TypeOf((*NotAuthorized)(nil)).Elem())
}

type NotAuthorizedFault NotAuthorized

func init() {
	types.Add("eam:NotAuthorizedFault", reflect.TypeOf((*NotAuthorizedFault)(nil)).Elem())
}

// Deprecated eAM no longer raises this issue. If agecny is getting orphaned
// EAM simply destroys it.
//
// The solution that created the agency is no longer registered with the vCenter
// server.
//
// This is an active remediable issue. To remediate, vSphere ESX Agent Manager
// undeploys and removes the agency.
//
// This structure may be used only with operations rendered under `/eam`.
type OrphanedAgency struct {
	AgencyIssue
}

func init() {
	types.Add("eam:OrphanedAgency", reflect.TypeOf((*OrphanedAgency)(nil)).Elem())
}

// Deprecated dvFilters are no longer supported by EAM.
//
// A dvFilter switch exists on a host but no agents on the host depend on
// dvFilter.
//
// This typically happens if a host is disconnected when an agency
// configuration changed.
//
// This is an active remediable issue. To remediate, vSphere ESX Agent Manager
// removes the dvFilterSwitch.
//
// This structure may be used only with operations rendered under `/eam`.
type OrphanedDvFilterSwitch struct {
	HostIssue
}

func init() {
	types.Add("eam:OrphanedDvFilterSwitch", reflect.TypeOf((*OrphanedDvFilterSwitch)(nil)).Elem())
}

// An Agent virtual machine is expected to be provisioned on a host, but it failed to do so
// because the provisioning of the OVF package failed.
//
// The provisioning is unlikely to
// succeed until the solution that provides the OVF package has been upgraded or
// patched to provide a valid OVF package for the agent virtual machine.
//
// This is an active remediable issue. To remediate, vSphere ESX Agent Manager attempts the OVF provisioning again.
//
// This structure may be used only with operations rendered under `/eam`.
type OvfInvalidFormat struct {
	VmNotDeployed

	// An optional list of errors that caused this issue.
	//
	// These errors are generated by the
	// vCenter server.
	Error []types.LocalizedMethodFault `xml:"error,omitempty" json:"error,omitempty"`
}

func init() {
	types.Add("eam:OvfInvalidFormat", reflect.TypeOf((*OvfInvalidFormat)(nil)).Elem())
}

// An agent virtual machine needs to be provisioned or reconfigured, but an OVF
// property is either missing or has an invalid value.
//
// This is a passive remediable issue. To remediate, update the OVF environment
// in the agent configuration used to provision the agent virtual machine.
//
// This structure may be used only with operations rendered under `/eam`.
type OvfInvalidProperty struct {
	AgentIssue

	// An optional list of errors that caused this issue.
	//
	// These errors are
	// generated by the vCenter server.
	Error []types.LocalizedMethodFault `xml:"error,omitempty" json:"error,omitempty"`
}

func init() {
	types.Add("eam:OvfInvalidProperty", reflect.TypeOf((*OvfInvalidProperty)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. Please refer to vLCM Image APIs.
//
// EAM was unable to set its required compute resource configuration in PM.
//
// EAM configuration needs to be updated or PM needs to be repaired manually to
// allow the configuration.
//
// This structure may be used only with operations rendered under `/eam`.
type PersonalityAgencyCannotConfigureSolutions struct {
	PersonalityAgencyPMIssue

	// Compute resource that couldn't be configured
	//
	// Refers instance of `ComputeResource`.
	Cr types.ManagedObjectReference `xml:"cr" json:"cr"`
	// Names of the solutions attempted to be modified
	SolutionsToModify []string `xml:"solutionsToModify,omitempty" json:"solutionsToModify,omitempty"`
	// Names of the solutions attempted to be removed
	SolutionsToRemove []string `xml:"solutionsToRemove,omitempty" json:"solutionsToRemove,omitempty"`
}

func init() {
	types.Add("eam:PersonalityAgencyCannotConfigureSolutions", reflect.TypeOf((*PersonalityAgencyCannotConfigureSolutions)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. Please refer to vLCM Image APIs.
//
// The offline depot could not be uploaded in Personality Manager.
//
// This structure may be used only with operations rendered under `/eam`.
type PersonalityAgencyCannotUploadDepot struct {
	PersonalityAgencyDepotIssue

	// URL EAM hosted the offline bundle as in vCenter.
	LocalDepotUrl string `xml:"localDepotUrl" json:"localDepotUrl"`
}

func init() {
	types.Add("eam:PersonalityAgencyCannotUploadDepot", reflect.TypeOf((*PersonalityAgencyCannotUploadDepot)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. Please refer to vLCM Image APIs.
//
// Base class for all offline depot (VIB) issues while communicating with
// Personality Manager.
//
// This structure may be used only with operations rendered under `/eam`.
type PersonalityAgencyDepotIssue struct {
	PersonalityAgencyPMIssue

	// URL the offline bundle is configured in EAM.
	RemoteDepotUrl string `xml:"remoteDepotUrl" json:"remoteDepotUrl"`
}

func init() {
	types.Add("eam:PersonalityAgencyDepotIssue", reflect.TypeOf((*PersonalityAgencyDepotIssue)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. Please refer to vLCM Image APIs.
//
// The offline depot was not available for download during communicating with
// Personality Manager.
//
// This structure may be used only with operations rendered under `/eam`.
type PersonalityAgencyInaccessibleDepot struct {
	PersonalityAgencyDepotIssue
}

func init() {
	types.Add("eam:PersonalityAgencyInaccessibleDepot", reflect.TypeOf((*PersonalityAgencyInaccessibleDepot)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. Please refer to vLCM Image APIs.
//
// The offline depot has missing or invalid metadata to be usable by
// Personality Manager.
//
// This structure may be used only with operations rendered under `/eam`.
type PersonalityAgencyInvalidDepot struct {
	PersonalityAgencyDepotIssue
}

func init() {
	types.Add("eam:PersonalityAgencyInvalidDepot", reflect.TypeOf((*PersonalityAgencyInvalidDepot)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. Please refer to vLCM Image APIs.
//
// Base class for all issues which occurred during EAM communication with
// Personality Manager.
//
// This structure may be used only with operations rendered under `/eam`.
type PersonalityAgencyPMIssue struct {
	AgencyIssue
}

func init() {
	types.Add("eam:PersonalityAgencyPMIssue", reflect.TypeOf((*PersonalityAgencyPMIssue)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. Please refer to vLCM Image APIs.
//
// PM service is not available - its endpoint cannot be accessed or it is
// malfunctioning.
//
// This is an active and passive remediable issue. To remediate, vSphere ESX
// Agent Manager retries to access PM service.
//
// This structure may be used only with operations rendered under `/eam`.
type PersonalityAgencyPMUnavailable struct {
	PersonalityAgencyPMIssue
}

func init() {
	types.Add("eam:PersonalityAgencyPMUnavailable", reflect.TypeOf((*PersonalityAgencyPMUnavailable)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. Please refer to vLCM Image APIs.
//
// The agent workflow is blocked until its' required solutions are re-mediated
// externally in Personality Manager.
//
// This issue is only passively remediable. The desired state has to be applied
// in PM by an user.
//
// This structure may be used only with operations rendered under `/eam`.
type PersonalityAgentAwaitingPMRemediation struct {
	PersonalityAgentPMIssue
}

func init() {
	types.Add("eam:PersonalityAgentAwaitingPMRemediation", reflect.TypeOf((*PersonalityAgentAwaitingPMRemediation)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. Please refer to vLCM Image APIs.
//
// The agent workflow is blocked by a failed agency operation with
// Personality Manager.
//
// This issue is only passively remediable. The agency's PM related issue has to
// be resolved.
//
// This structure may be used only with operations rendered under `/eam`.
type PersonalityAgentBlockedByAgencyOperation struct {
	PersonalityAgentPMIssue
}

func init() {
	types.Add("eam:PersonalityAgentBlockedByAgencyOperation", reflect.TypeOf((*PersonalityAgentBlockedByAgencyOperation)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. Please refer to vLCM Image APIs.
//
// Base class for all issues which occurred during EAM communication with
// Personality Manager.
//
// This structure may be used only with operations rendered under `/eam`.
type PersonalityAgentPMIssue struct {
	AgentIssue
}

func init() {
	types.Add("eam:PersonalityAgentPMIssue", reflect.TypeOf((*PersonalityAgentPMIssue)(nil)).Elem())
}

type QueryAgency QueryAgencyRequestType

func init() {
	types.Add("eam:QueryAgency", reflect.TypeOf((*QueryAgency)(nil)).Elem())
}

type QueryAgencyRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
}

func init() {
	types.Add("eam:QueryAgencyRequestType", reflect.TypeOf((*QueryAgencyRequestType)(nil)).Elem())
}

type QueryAgencyResponse struct {
	Returnval []types.ManagedObjectReference `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

type QueryAgent QueryAgentRequestType

func init() {
	types.Add("eam:QueryAgent", reflect.TypeOf((*QueryAgent)(nil)).Elem())
}

type QueryAgentRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
}

func init() {
	types.Add("eam:QueryAgentRequestType", reflect.TypeOf((*QueryAgentRequestType)(nil)).Elem())
}

type QueryAgentResponse struct {
	Returnval []types.ManagedObjectReference `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

type QueryConfig QueryConfigRequestType

func init() {
	types.Add("eam:QueryConfig", reflect.TypeOf((*QueryConfig)(nil)).Elem())
}

type QueryConfigRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
}

func init() {
	types.Add("eam:QueryConfigRequestType", reflect.TypeOf((*QueryConfigRequestType)(nil)).Elem())
}

type QueryConfigResponse struct {
	Returnval BaseAgencyConfigInfo `xml:"returnval,typeattr" json:"returnval"`
}

type QueryIssue QueryIssueRequestType

func init() {
	types.Add("eam:QueryIssue", reflect.TypeOf((*QueryIssue)(nil)).Elem())
}

// The parameters of `EamObject.QueryIssue`.
//
// This structure may be used only with operations rendered under `/eam`.
type QueryIssueRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// An optional array of issue keys. If not set, all issues for this
	// entity are returned.
	IssueKey []int32 `xml:"issueKey,omitempty" json:"issueKey,omitempty"`
}

func init() {
	types.Add("eam:QueryIssueRequestType", reflect.TypeOf((*QueryIssueRequestType)(nil)).Elem())
}

type QueryIssueResponse struct {
	Returnval []BaseIssue `xml:"returnval,omitempty,typeattr" json:"returnval,omitempty"`
}

type QuerySolutionId QuerySolutionIdRequestType

func init() {
	types.Add("eam:QuerySolutionId", reflect.TypeOf((*QuerySolutionId)(nil)).Elem())
}

type QuerySolutionIdRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
}

func init() {
	types.Add("eam:QuerySolutionIdRequestType", reflect.TypeOf((*QuerySolutionIdRequestType)(nil)).Elem())
}

type QuerySolutionIdResponse struct {
	Returnval string `xml:"returnval" json:"returnval"`
}

type RegisterAgentVm RegisterAgentVmRequestType

func init() {
	types.Add("eam:RegisterAgentVm", reflect.TypeOf((*RegisterAgentVm)(nil)).Elem())
}

// The parameters of `Agency.RegisterAgentVm`.
//
// This structure may be used only with operations rendered under `/eam`.
type RegisterAgentVmRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// The managed object reference to the agent VM.
	//
	// Refers instance of `VirtualMachine`.
	AgentVm types.ManagedObjectReference `xml:"agentVm" json:"agentVm"`
}

func init() {
	types.Add("eam:RegisterAgentVmRequestType", reflect.TypeOf((*RegisterAgentVmRequestType)(nil)).Elem())
}

type RegisterAgentVmResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}

type Resolve ResolveRequestType

func init() {
	types.Add("eam:Resolve", reflect.TypeOf((*Resolve)(nil)).Elem())
}

type ResolveAll ResolveAllRequestType

func init() {
	types.Add("eam:ResolveAll", reflect.TypeOf((*ResolveAll)(nil)).Elem())
}

type ResolveAllRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
}

func init() {
	types.Add("eam:ResolveAllRequestType", reflect.TypeOf((*ResolveAllRequestType)(nil)).Elem())
}

type ResolveAllResponse struct {
}

// The parameters of `EamObject.Resolve`.
//
// This structure may be used only with operations rendered under `/eam`.
type ResolveRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// A non-empty array of issue keys.
	IssueKey []int32 `xml:"issueKey" json:"issueKey"`
}

func init() {
	types.Add("eam:ResolveRequestType", reflect.TypeOf((*ResolveRequestType)(nil)).Elem())
}

type ResolveResponse struct {
	Returnval []int32 `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

type ScanForUnknownAgentVm ScanForUnknownAgentVmRequestType

func init() {
	types.Add("eam:ScanForUnknownAgentVm", reflect.TypeOf((*ScanForUnknownAgentVm)(nil)).Elem())
}

type ScanForUnknownAgentVmRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
}

func init() {
	types.Add("eam:ScanForUnknownAgentVmRequestType", reflect.TypeOf((*ScanForUnknownAgentVmRequestType)(nil)).Elem())
}

type ScanForUnknownAgentVmResponse struct {
}

type SetMaintenanceModePolicy SetMaintenanceModePolicyRequestType

func init() {
	types.Add("eam:SetMaintenanceModePolicy", reflect.TypeOf((*SetMaintenanceModePolicy)(nil)).Elem())
}

// The parameters of `EsxAgentManager.SetMaintenanceModePolicy`.
//
// This structure may be used only with operations rendered under `/eam`.
type SetMaintenanceModePolicyRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// The policy to use.
	Policy string `xml:"policy" json:"policy"`
}

func init() {
	types.Add("eam:SetMaintenanceModePolicyRequestType", reflect.TypeOf((*SetMaintenanceModePolicyRequestType)(nil)).Elem())
	types.AddMinAPIVersionForType("eam:SetMaintenanceModePolicyRequestType", "7.4")
}

type SetMaintenanceModePolicyResponse struct {
}

type SolutionsAlternativeVmSpec struct {
	types.DynamicData

	SelectionCriteria SolutionsVmSelectionSpec        `xml:"selectionCriteria" json:"selectionCriteria"`
	Devices           *types.VirtualMachineConfigSpec `xml:"devices,omitempty" json:"devices,omitempty"`
}

func init() {
	types.Add("eam:SolutionsAlternativeVmSpec", reflect.TypeOf((*SolutionsAlternativeVmSpec)(nil)).Elem())
}

// Specification describing a desired state to be applied.
//
// This structure may be used only with operations rendered under `/eam`.
type SolutionsApplySpec struct {
	types.DynamicData

	// Complete desired state to be applied on the target entity.
	//
	// the `*solutions*` member limits which parts of this desired state to
	// be applied
	//
	//	  If the `*solutions*` member is omitted.
	//	- Any solution described in this structure will be applied on the
	//	  target entity
	//	- Any solution already existing on the target entity, but missing
	//	  from this structure, will be deleted from the target entity
	DesiredState []SolutionsSolutionConfig `xml:"desiredState,omitempty" json:"desiredState,omitempty"`
	// If provided, limits the parts of the `*desiredState*` structure to
	// be applied on the target entity.
	//   - solutions that are also present in the `*desiredState*`
	//     structure will be applied on the target entity.
	//   - solutions that are missing from the `*desiredState*` structure
	//     will be deleted from the target entity.
	Solutions []string `xml:"solutions,omitempty" json:"solutions,omitempty"`
	// Specifies exact hosts to apply the desired state to, instead of every
	// host in the cluster.
	//
	// Applicable only to solutions with
	// `SolutionsHostBoundSolutionConfig`.
	//
	// Refers instances of `HostSystem`.
	Hosts []types.ManagedObjectReference `xml:"hosts,omitempty" json:"hosts,omitempty"`
	// Deployment units on which solutions that are specified by the this
	// structure need to be applied.
	//
	// Applicable only to solutions with
	// `SolutionsClusterBoundSolutionConfig`.
	//
	// The deployment unit represents a single VM instance deployment. It is
	// returned by the `Solutions.Compliance` operation.
	//
	// This filtering is not supported in case of subsequent
	// `Hooks#processDynamicUpdate` is invoked.
	// If omitted - the configured solutions by `SolutionsApplySpec.solutions` are applied
	// on all of the deployment units in the cluster.
	DeploymentUnits []string `xml:"deploymentUnits,omitempty" json:"deploymentUnits,omitempty"`
}

func init() {
	types.Add("eam:SolutionsApplySpec", reflect.TypeOf((*SolutionsApplySpec)(nil)).Elem())
}

// Specifies cluster-bound solution configuration.
//
// This structure may be used only with operations rendered under `/eam`.
type SolutionsClusterBoundSolutionConfig struct {
	SolutionsTypeSpecificSolutionConfig

	// The number of instances of the specified VM to be deployed across the
	// cluster.
	VmCount int32 `xml:"vmCount" json:"vmCount"`
	// VM placement policies to be configured on the VMs
	// `SolutionsVmPlacementPolicy_enum` If omitted - no VM placement
	// policies are configured.
	VmPlacementPolicies []string `xml:"vmPlacementPolicies,omitempty" json:"vmPlacementPolicies,omitempty"`
	// Networks defined in the OVF to be configured on the VMs.
	//
	// Mutually
	// exclusive with `SolutionsClusterBoundSolutionConfig.devices`. If omitted - no VM networks are
	// configured.
	VmNetworks []SolutionsVMNetworkMapping `xml:"vmNetworks,omitempty" json:"vmNetworks,omitempty"`
	// Datastores to be configured as a storage of the VMs.
	//
	// The first
	// available datastore in the cluster is used. The collection cannot
	// contain duplicate elements.
	// If omitted - the system automatically selects the datastore. The
	// selection takes into account the other properties of the desired state
	// specification (the provided VM storage policies and VM devices) and the
	// runtime state of the datastores in the cluster. It is required DRS to
	// be enabled on the cluster.
	//
	// Refers instances of `Datastore`.
	Datastores []types.ManagedObjectReference `xml:"datastores,omitempty" json:"datastores,omitempty"`
	// Devices of the VMs not defined in the OVF descriptor.
	//
	// Mutually
	// exclusive with `SolutionsClusterBoundSolutionConfig.vmNetworks`.
	//
	// If `SolutionsClusterBoundSolutionConfig.datastores` is not set, the devices of the VMs not defined
	// in the OVF descriptor should be provided to `SolutionsClusterBoundSolutionConfig.devices` and not as
	// part of a VM lifecycle hook (VM reconfiguration). Otherwise, the
	// compatibility of the devices with the selected host and datastore where
	// the VM is deployed needs to be ensured by the client.
	//
	// 1\. For VM initial placement the devices are added to the VM configuration.
	// 2\. For the reconfiguration it is checked what devices need to be added,
	// removed, and edited on the existing VMs. NOTE: No VM relocation is
	// executed before the VM reconfiguration.
	//
	// The supported property of vim.vm.ConfigSpec is
	// vim.vm.ConfigSpec.deviceChange. The supported
	// vim.vm.device.VirtualDeviceSpec.operation is Operation#add. For
	// vim.vm.device.VirtualEthernetCard the unique identifier is
	// vim.vm.device.VirtualDevice#unitNumber.
	//
	// If omitted - no additional devices will be added to the VMs.
	Devices            *types.VirtualMachineConfigSpec `xml:"devices,omitempty" json:"devices,omitempty"`
	RemediationPolicy  BaseSolutionsRemediationPolicy  `xml:"remediationPolicy,omitempty,typeattr" json:"remediationPolicy,omitempty"`
	AlternativeVmSpecs []SolutionsAlternativeVmSpec    `xml:"alternativeVmSpecs,omitempty" json:"alternativeVmSpecs,omitempty"`
}

func init() {
	types.Add("eam:SolutionsClusterBoundSolutionConfig", reflect.TypeOf((*SolutionsClusterBoundSolutionConfig)(nil)).Elem())
}

// Result of a compliance check of a desired state for a solution with
// `SolutionsClusterBoundSolutionConfig`.
//
// This structure may be used only with operations rendered under `/eam`.
type SolutionsClusterSolutionComplianceResult struct {
	types.DynamicData

	// The solution being checked for compliance.
	Solution string `xml:"solution" json:"solution"`
	// `True` if the solution is compliant with the described desired
	// state, `False` - otherwise.
	Compliant bool `xml:"compliant" json:"compliant"`
	// Detailed per deployment-unit compliance result.
	DeploymentUnits []SolutionsDeploymentUnitComplianceResult `xml:"deploymentUnits,omitempty" json:"deploymentUnits,omitempty"`
}

func init() {
	types.Add("eam:SolutionsClusterSolutionComplianceResult", reflect.TypeOf((*SolutionsClusterSolutionComplianceResult)(nil)).Elem())
}

type SolutionsClusterTransitionSpec struct {
	types.DynamicData

	Solution      string                       `xml:"solution" json:"solution"`
	SourceCluster types.ManagedObjectReference `xml:"sourceCluster" json:"sourceCluster"`
}

func init() {
	types.Add("eam:SolutionsClusterTransitionSpec", reflect.TypeOf((*SolutionsClusterTransitionSpec)(nil)).Elem())
}

// Result of a compliance check of a desired state on a compute resource.
//
// This structure may be used only with operations rendered under `/eam`.
type SolutionsComplianceResult struct {
	types.DynamicData

	// `True` if the compute resource is compliant with the described
	// desired state, `False` - otherwise.
	Compliant bool `xml:"compliant" json:"compliant"`
	// Detailed per-host compliance result of the compute resource for
	// solutions with `SolutionsHostBoundSolutionConfig`.
	Hosts []SolutionsHostComplianceResult `xml:"hosts,omitempty" json:"hosts,omitempty"`
	// Detailed per-solution unit compliance result of the compute resource
	// for solutions with `SolutionsClusterBoundSolutionConfig`.
	ClusterSolutionsCompliance []SolutionsClusterSolutionComplianceResult `xml:"clusterSolutionsCompliance,omitempty" json:"clusterSolutionsCompliance,omitempty"`
}

func init() {
	types.Add("eam:SolutionsComplianceResult", reflect.TypeOf((*SolutionsComplianceResult)(nil)).Elem())
}

// Specification describing how to calculate compliance of a compute resource
// against a desired state.
//
// This structure may be used only with operations rendered under `/eam`.
type SolutionsComplianceSpec struct {
	types.DynamicData

	// Desired state to be checked for compliance.
	//
	// May be incomplete if exact
	// solutions to be checked are provided. Empty desired state means all
	// present solutions must be removed.
	DesiredState []SolutionsSolutionConfig `xml:"desiredState,omitempty" json:"desiredState,omitempty"`
	// Specifies exact solutions to be checked for compliance instead of the
	// complete desired state.
	Solutions []string `xml:"solutions,omitempty" json:"solutions,omitempty"`
	// Specifies exact hosts to be checked for compliance of all solutions
	// with `SolutionsHostBoundSolutionConfig`.
	//
	// If omitted - the compliance is checked for all hosts in the cluster.
	//
	// Refers instances of `HostSystem`.
	Hosts []types.ManagedObjectReference `xml:"hosts,omitempty" json:"hosts,omitempty"`
	// Identifiers of the deployment units that to be checked for compliance
	// of all solutions with `SolutionsClusterBoundSolutionConfig`.
	//
	// The deployment unit represents a single VM instance deployment.
	//
	// If omitted - the compliance is checked for all deployment units in the
	// cluster.
	DeploymentUnits []string `xml:"deploymentUnits,omitempty" json:"deploymentUnits,omitempty"`
}

func init() {
	types.Add("eam:SolutionsComplianceSpec", reflect.TypeOf((*SolutionsComplianceSpec)(nil)).Elem())
}

// Result of a compliance check of a deployment unit.
//
// This structure may be used only with operations rendered under `/eam`.
type SolutionsDeploymentUnitComplianceResult struct {
	types.DynamicData

	// The deployment unit being checked for compliance.
	DeploymentUnit string `xml:"deploymentUnit" json:"deploymentUnit"`
	// `True` if the deployment unit is compliant with the described
	// desired state, `False` - otherwise.
	Compliant bool `xml:"compliant" json:"compliant"`
	// Detailed compliance result of the deployment unit.
	Compliance *SolutionsSolutionComplianceResult `xml:"compliance,omitempty" json:"compliance,omitempty"`
}

func init() {
	types.Add("eam:SolutionsDeploymentUnitComplianceResult", reflect.TypeOf((*SolutionsDeploymentUnitComplianceResult)(nil)).Elem())
}

// Specifies the acknowledgement type of a configured System Virtual
// Machine's lifecycle hook.
//
// This structure may be used only with operations rendered under `/eam`.
type SolutionsHookAcknowledgeConfig struct {
	types.DynamicData
}

func init() {
	types.Add("eam:SolutionsHookAcknowledgeConfig", reflect.TypeOf((*SolutionsHookAcknowledgeConfig)(nil)).Elem())
}

// Configuration for System Virtual Machine's lifecycle hooks.
//
// This structure may be used only with operations rendered under `/eam`.
type SolutionsHookConfig struct {
	types.DynamicData

	// Type of the configured hook, possible values - `HooksHookType_enum`.
	Type string `xml:"type" json:"type"`
	// Type of acknowledgement of the configured hook.
	Acknowledgement BaseSolutionsHookAcknowledgeConfig `xml:"acknowledgement,typeattr" json:"acknowledgement"`
	// The maximum time in seconds to wait for a hook to be processed.
	//
	// An
	// issue is raised if the time elapsed and the hook is still not
	// processed.
	// If omitted - defaults to 10 hours.
	Timeout int64 `xml:"timeout,omitempty" json:"timeout,omitempty"`
}

func init() {
	types.Add("eam:SolutionsHookConfig", reflect.TypeOf((*SolutionsHookConfig)(nil)).Elem())
}

// Contains information for a raised hook.
//
// This structure may be used only with operations rendered under `/eam`.
type SolutionsHookInfo struct {
	types.DynamicData

	// Virtual Machine, the hook was raised for.
	//
	// Refers instance of `VirtualMachine`.
	Vm types.ManagedObjectReference `xml:"vm" json:"vm"`
	// Solution the Virtual Machine belongs to.
	Solution string `xml:"solution" json:"solution"`
	// Configuration of the hook.
	Config SolutionsHookConfig `xml:"config" json:"config"`
	// Time the hook was raised.
	RaisedAt time.Time `xml:"raisedAt" json:"raisedAt"`
	// True if `Hooks#processDynamicUpdate` method invocation completed
	// successfully for this hook.
	//
	// Otherwise defaults to False.
	DynamicUpdateProcessed bool `xml:"dynamicUpdateProcessed" json:"dynamicUpdateProcessed"`
}

func init() {
	types.Add("eam:SolutionsHookInfo", reflect.TypeOf((*SolutionsHookInfo)(nil)).Elem())
}

// Specifies host-bound solution configuration.
//
// This structure may be used only with operations rendered under `/eam`.
type SolutionsHostBoundSolutionConfig struct {
	SolutionsTypeSpecificSolutionConfig

	OmitDRSBlocking *bool `xml:"omitDRSBlocking" json:"omitDRSBlocking,omitempty"`
	// If set to true - default network and datastore configured on host will
	// take precedence over
	// `SolutionsHostBoundSolutionConfig.datastores` and
	// `SolutionsHostBoundSolutionConfig.networks`.
	PreferHostConfiguration *bool `xml:"preferHostConfiguration" json:"preferHostConfiguration,omitempty"`
	// networks to satisfy system Virtual Machine network adapter
	// requirements.
	//
	// If omitted - default configured network on the host will
	// be used.
	//
	// Refers instances of `Network`.
	Networks []types.ManagedObjectReference `xml:"networks,omitempty" json:"networks,omitempty"`
	// Datastores to be configured as a storage of the VMs.
	//
	// The first
	// available datastore on the host is used. The collection cannot contain
	// duplicate elements. If omitted - default configured datastore on the
	// host will be used.
	//
	// Refers instances of `Datastore`.
	Datastores []types.ManagedObjectReference `xml:"datastores,omitempty" json:"datastores,omitempty"`
	// VMCI to be allowed access from the system Virtual Machine.
	Vmci []string `xml:"vmci,omitempty" json:"vmci,omitempty"`
}

func init() {
	types.Add("eam:SolutionsHostBoundSolutionConfig", reflect.TypeOf((*SolutionsHostBoundSolutionConfig)(nil)).Elem())
}

// Result of a compliance check of a desired state on a host.
//
// This structure may be used only with operations rendered under `/eam`.
type SolutionsHostComplianceResult struct {
	types.DynamicData

	// The host being checked for compliance.
	//
	// Refers instance of `HostSystem`.
	Host types.ManagedObjectReference `xml:"host" json:"host"`
	// `True` if the compute host is compliant with the described
	// desired state, `False` - otherwise.
	Compliant bool `xml:"compliant" json:"compliant"`
	// Detailed per-solution compliance result of the host.
	Solutions []SolutionsSolutionComplianceResult `xml:"solutions,omitempty" json:"solutions,omitempty"`
}

func init() {
	types.Add("eam:SolutionsHostComplianceResult", reflect.TypeOf((*SolutionsHostComplianceResult)(nil)).Elem())
}

// The user will have to (manually) invoke an API
// (`Hooks.MarkAsProcessed`) to acknowledge, the user operations for
// this lifecycle hook have been completed.
//
// This structure may be used only with operations rendered under `/eam`.
type SolutionsInteractiveHookAcknowledgeConfig struct {
	SolutionsHookAcknowledgeConfig
}

func init() {
	types.Add("eam:SolutionsInteractiveHookAcknowledgeConfig", reflect.TypeOf((*SolutionsInteractiveHookAcknowledgeConfig)(nil)).Elem())
}

type SolutionsMultiSourceTransitionSpec struct {
	types.DynamicData

	Solution               string                            `xml:"solution" json:"solution"`
	AgencyIds              []string                          `xml:"agencyIds,omitempty" json:"agencyIds,omitempty"`
	SourceVmSelectionSpecs []SolutionsVmSelectionSpecMapping `xml:"sourceVmSelectionSpecs,omitempty" json:"sourceVmSelectionSpecs,omitempty"`
	Module                 string                            `xml:"module,omitempty" json:"module,omitempty"`
}

func init() {
	types.Add("eam:SolutionsMultiSourceTransitionSpec", reflect.TypeOf((*SolutionsMultiSourceTransitionSpec)(nil)).Elem())
}

// One OVF Property.
//
// This structure may be used only with operations rendered under `/eam`.
type SolutionsOvfProperty struct {
	types.DynamicData

	// The name of the property in the OVF descriptor.
	Key string `xml:"key" json:"key"`
	// The value of the property.
	Value string `xml:"value" json:"value"`
}

func init() {
	types.Add("eam:SolutionsOvfProperty", reflect.TypeOf((*SolutionsOvfProperty)(nil)).Elem())
}

type SolutionsParallelRemediationPolicy struct {
	SolutionsRemediationPolicy
}

func init() {
	types.Add("eam:SolutionsParallelRemediationPolicy", reflect.TypeOf((*SolutionsParallelRemediationPolicy)(nil)).Elem())
}

// Specifies a user defined profile ID to be applied during Virtual Machine
// creation.
//
// This structure may be used only with operations rendered under `/eam`.
type SolutionsProfileIdStoragePolicy struct {
	SolutionsStoragePolicy

	// ID of a storage policy profile created by the user.
	//
	// The type of the
	// profile must be `VirtualMachineDefinedProfileSpec`. The ID must be valid
	// `VirtualMachineDefinedProfileSpec.profileId`.
	ProfileId string `xml:"profileId" json:"profileId"`
}

func init() {
	types.Add("eam:SolutionsProfileIdStoragePolicy", reflect.TypeOf((*SolutionsProfileIdStoragePolicy)(nil)).Elem())
}

type SolutionsRemediationPolicy struct {
	types.DynamicData
}

func init() {
	types.Add("eam:SolutionsRemediationPolicy", reflect.TypeOf((*SolutionsRemediationPolicy)(nil)).Elem())
}

type SolutionsSequentialRemediationPolicy struct {
	SolutionsRemediationPolicy
}

func init() {
	types.Add("eam:SolutionsSequentialRemediationPolicy", reflect.TypeOf((*SolutionsSequentialRemediationPolicy)(nil)).Elem())
}

// Result of a compliance check of a desired state for a solution(on a host).
//
// This structure may be used only with operations rendered under `/eam`.
type SolutionsSolutionComplianceResult struct {
	types.DynamicData

	// Solution checked for compliance.
	Solution string `xml:"solution" json:"solution"`
	// `True` if the compute solution is compliant with the described
	// desired state, `False` - otherwise.
	Compliant bool `xml:"compliant" json:"compliant"`
	// Reason the solution is non-compliant
	// `SolutionsNonComplianceReason_enum`.
	NonComplianceReason string `xml:"nonComplianceReason,omitempty" json:"nonComplianceReason,omitempty"`
	// system Virtual Machine created for the solution.
	//
	// Refers instance of `VirtualMachine`.
	Vm *types.ManagedObjectReference `xml:"vm,omitempty" json:"vm,omitempty"`
	// system Virtual Machine created for upgrading the obsoleted system
	// Virtual Machine.
	//
	// Refers instance of `VirtualMachine`.
	UpgradingVm *types.ManagedObjectReference `xml:"upgradingVm,omitempty" json:"upgradingVm,omitempty"`
	// Hook, ESX Agent Manager is awaiting to be processed for this solution.
	Hook *SolutionsHookInfo `xml:"hook,omitempty" json:"hook,omitempty"`
	// Issues, ESX Agent Manager has encountered while attempting to achieve
	// the solution's requested desired state.
	Issues []BaseIssue `xml:"issues,omitempty,typeattr" json:"issues,omitempty"`
	// Last desired state for the solution, requested from ESX Agent Manager,
	// for application.
	SolutionConfig *SolutionsSolutionConfig `xml:"solutionConfig,omitempty" json:"solutionConfig,omitempty"`
}

func init() {
	types.Add("eam:SolutionsSolutionComplianceResult", reflect.TypeOf((*SolutionsSolutionComplianceResult)(nil)).Elem())
}

// Configuration for a solution's required system Virtual Machine.
//
// This structure may be used only with operations rendered under `/eam`.
type SolutionsSolutionConfig struct {
	types.DynamicData

	// Solution, this configuration belongs to.
	Solution string `xml:"solution" json:"solution"`
	// Display name of the solution.
	DisplayName string `xml:"displayName" json:"displayName"`
	// Display version of the solution.
	DisplayVersion string `xml:"displayVersion" json:"displayVersion"`
	// Source of the system Virtual Machine files.
	VmSource BaseSolutionsVMSource `xml:"vmSource,typeattr" json:"vmSource"`
	// VM name prefix.
	PrefixVmName string `xml:"prefixVmName" json:"prefixVmName"`
	// If set to `True` - will insert an UUID in the system Virtual
	// Machines' names created for the solution, otherwise - no additional
	// UUID will be inserted in the system Virtual Machines' names.
	UuidVmName bool `xml:"uuidVmName" json:"uuidVmName"`
	// Resource pool to place the system Virtual Machine in.
	//
	// If omitted a
	// default resource pool will be used.
	//
	// Refers instance of `ResourcePool`.
	ResourcePool *types.ManagedObjectReference `xml:"resourcePool,omitempty" json:"resourcePool,omitempty"`
	// Folder to place the system Virtual Machine in.
	//
	// If omitted a default
	// folder will be used.
	//
	// Refers instance of `Folder`.
	Folder *types.ManagedObjectReference `xml:"folder,omitempty" json:"folder,omitempty"`
	// User configurable OVF properties to be assigned during system Virtual
	// Machine creation.
	OvfProperties []SolutionsOvfProperty `xml:"ovfProperties,omitempty" json:"ovfProperties,omitempty"`
	// Storage policies to be applied during system Virtual Machine creation.
	StoragePolicies []BaseSolutionsStoragePolicy `xml:"storagePolicies,omitempty,typeattr" json:"storagePolicies,omitempty"`
	// Provisioning type for the system Virtual Machines
	// `SolutionsVMDiskProvisioning_enum`.
	//
	// Default provisioning will be used
	// if not specified.
	VmDiskProvisioning string `xml:"vmDiskProvisioning,omitempty" json:"vmDiskProvisioning,omitempty"`
	// Optimization strategy for deploying Virtual Machines
	// `SolutionsVMDeploymentOptimization_enum`.
	//
	// Default optimization will
	// be selected if not specified.
	VmDeploymentOptimization string `xml:"vmDeploymentOptimization,omitempty" json:"vmDeploymentOptimization,omitempty"`
	// Solution type-specific configuration.
	TypeSpecificConfig BaseSolutionsTypeSpecificSolutionConfig `xml:"typeSpecificConfig,typeattr" json:"typeSpecificConfig"`
	// Lifecycle hooks for the solution's virtual machines.
	Hooks []SolutionsHookConfig `xml:"hooks,omitempty" json:"hooks,omitempty"`
	// VMs resource configuration.
	//
	// If omitted - the default resource
	// configuration specified in the OVF descriptor is used.
	VmResourceSpec     *SolutionsVmResourceSpec `xml:"vmResourceSpec,omitempty" json:"vmResourceSpec,omitempty"`
	RedeploymentPolicy string                   `xml:"redeploymentPolicy,omitempty" json:"redeploymentPolicy,omitempty"`
}

func init() {
	types.Add("eam:SolutionsSolutionConfig", reflect.TypeOf((*SolutionsSolutionConfig)(nil)).Elem())
}

// Result of validation, of a solution, for application.
//
// This structure may be used only with operations rendered under `/eam`.
type SolutionsSolutionValidationResult struct {
	types.DynamicData

	// Validated solution.
	Solution string `xml:"solution" json:"solution"`
	// `True` - if the solution is valid for application, `False`
	// \- otherwise.
	Valid bool `xml:"valid" json:"valid"`
	// Populated with the reason the solution is not valid for application
	// `SolutionsInvalidReason_enum`.
	InvalidReason string `xml:"invalidReason,omitempty" json:"invalidReason,omitempty"`
}

func init() {
	types.Add("eam:SolutionsSolutionValidationResult", reflect.TypeOf((*SolutionsSolutionValidationResult)(nil)).Elem())
}

// Storage policy to be applied during system Virtual Machine creation.
//
// This structure may be used only with operations rendered under `/eam`.
type SolutionsStoragePolicy struct {
	types.DynamicData
}

func init() {
	types.Add("eam:SolutionsStoragePolicy", reflect.TypeOf((*SolutionsStoragePolicy)(nil)).Elem())
}

// Specification necessary to transition a solution from an existing legacy
// agency.
//
// This structure may be used only with operations rendered under `/eam`.
type SolutionsTransitionSpec struct {
	types.DynamicData

	// Solution to transition from an old legacy agency.
	Solution string `xml:"solution" json:"solution"`
	// Old legacy agency ID to transition from.
	AgencyId string `xml:"agencyId" json:"agencyId"`
}

func init() {
	types.Add("eam:SolutionsTransitionSpec", reflect.TypeOf((*SolutionsTransitionSpec)(nil)).Elem())
}

// Specifies the specific solution configuration based on its type.
//
// This structure may be used only with operations rendered under `/eam`.
type SolutionsTypeSpecificSolutionConfig struct {
	types.DynamicData
}

func init() {
	types.Add("eam:SolutionsTypeSpecificSolutionConfig", reflect.TypeOf((*SolutionsTypeSpecificSolutionConfig)(nil)).Elem())
}

// Specified the system Virtual Machine sources are to be obtained from an
// URL.
//
// This structure may be used only with operations rendered under `/eam`.
type SolutionsUrlVMSource struct {
	SolutionsVMSource

	// URL to the solution's system Virtual Machine OVF.
	OvfUrl string `xml:"ovfUrl" json:"ovfUrl"`
	// Overrides the OVF URL certificate validation.
	//
	// If `True` or
	// `<unset>` - the certificate will be subject to standard trust
	// validation, if `False` - any certificate will be considered
	// trusted.
	CertificateValidation *bool `xml:"certificateValidation" json:"certificateValidation,omitempty"`
	// PEM encoded certificate to use to trust the URL.
	//
	// If omitted - URL will
	// be trusted using well known methods.
	CertificatePEM       string `xml:"certificatePEM,omitempty" json:"certificatePEM,omitempty"`
	AuthenticationScheme string `xml:"authenticationScheme,omitempty" json:"authenticationScheme,omitempty"`
}

func init() {
	types.Add("eam:SolutionsUrlVMSource", reflect.TypeOf((*SolutionsUrlVMSource)(nil)).Elem())
}

// Represents the mapping of the logical network defined in the OVF
// descriptor to the Virtual Infrastructure (VI) network.
//
// This structure may be used only with operations rendered under `/eam`.
type SolutionsVMNetworkMapping struct {
	types.DynamicData

	// Logical network name defined in the OVF descriptor.
	Name string `xml:"name" json:"name"`
	// VM network identifier.
	//
	// Refers instance of `Network`.
	Id types.ManagedObjectReference `xml:"id" json:"id"`
}

func init() {
	types.Add("eam:SolutionsVMNetworkMapping", reflect.TypeOf((*SolutionsVMNetworkMapping)(nil)).Elem())
}

// Specifies how to find the files of the system Virtual Machine to be
// created.
//
// This structure may be used only with operations rendered under `/eam`.
type SolutionsVMSource struct {
	types.DynamicData
}

func init() {
	types.Add("eam:SolutionsVMSource", reflect.TypeOf((*SolutionsVMSource)(nil)).Elem())
}

// Specification describing a desired state to be validated for application.
//
// This structure may be used only with operations rendered under `/eam`.
type SolutionsValidateSpec struct {
	types.DynamicData

	// Desired state to be validated.
	DesiredState []SolutionsSolutionConfig `xml:"desiredState" json:"desiredState"`
	// Transition specification to be validated.
	//
	// Mutually exclusive with `#multiSourceTransitionSpec` and
	// `#clusterTransitionSpec`.
	TransitionSpec            *SolutionsTransitionSpec            `xml:"transitionSpec,omitempty" json:"transitionSpec,omitempty" vim:"9.0"`
	MultiSourceTransitionSpec *SolutionsMultiSourceTransitionSpec `xml:"multiSourceTransitionSpec,omitempty" json:"multiSourceTransitionSpec,omitempty"`
	ClusterTransitionSpec     *SolutionsClusterTransitionSpec     `xml:"clusterTransitionSpec,omitempty" json:"clusterTransitionSpec,omitempty"`
}

func init() {
	types.Add("eam:SolutionsValidateSpec", reflect.TypeOf((*SolutionsValidateSpec)(nil)).Elem())
}

// Result of validation, of a desired state, for application.
//
// This structure may be used only with operations rendered under `/eam`.
type SolutionsValidationResult struct {
	types.DynamicData

	// `True` - if the desired state is valid for application,
	// `False` - otherwise.
	Valid bool `xml:"valid" json:"valid"`
	// Detailed per-solution result of the validation.
	SolutionResult []SolutionsSolutionValidationResult `xml:"solutionResult,omitempty" json:"solutionResult,omitempty"`
}

func init() {
	types.Add("eam:SolutionsValidationResult", reflect.TypeOf((*SolutionsValidationResult)(nil)).Elem())
}

// Specifies the VM resource configurations.
//
// This structure may be used only with operations rendered under `/eam`.
type SolutionsVmResourceSpec struct {
	types.DynamicData

	// The VM deployment option that corresponds to the Configuration element
	// of the DeploymentOptionSection in the OVF descriptor (e.g.
	//
	// "small",
	// "medium", "large").
	// If omitted - the default deployment option as specified in the OVF
	// descriptor is used.
	OvfDeploymentOption string `xml:"ovfDeploymentOption,omitempty" json:"ovfDeploymentOption,omitempty"`
}

func init() {
	types.Add("eam:SolutionsVmResourceSpec", reflect.TypeOf((*SolutionsVmResourceSpec)(nil)).Elem())
}

type SolutionsVmSelectionSpec struct {
	types.DynamicData

	SelectionType    string `xml:"selectionType" json:"selectionType"`
	ExtraConfigValue string `xml:"extraConfigValue,omitempty" json:"extraConfigValue,omitempty"`
}

func init() {
	types.Add("eam:SolutionsVmSelectionSpec", reflect.TypeOf((*SolutionsVmSelectionSpec)(nil)).Elem())
}

type SolutionsVmSelectionSpecMapping struct {
	types.DynamicData

	VmId            types.ManagedObjectReference `xml:"vmId" json:"vmId"`
	VmSelectionSpec SolutionsVmSelectionSpec     `xml:"vmSelectionSpec" json:"vmSelectionSpec"`
}

func init() {
	types.Add("eam:SolutionsVmSelectionSpecMapping", reflect.TypeOf((*SolutionsVmSelectionSpecMapping)(nil)).Elem())
}

// An agent failed to be transitioned to a LCCM Solution.
//
// This is an active remediable issue. To remediate, resolve the issue via vLCM
// System VMs API
//
// This structure may be used only with operations rendered under `/eam`.
type TransitionFailed struct {
	AgentIssue
}

func init() {
	types.Add("eam:TransitionFailed", reflect.TypeOf((*TransitionFailed)(nil)).Elem())
	types.AddMinAPIVersionForType("eam:TransitionFailed", "9.0")
}

type Uninstall UninstallRequestType

func init() {
	types.Add("eam:Uninstall", reflect.TypeOf((*Uninstall)(nil)).Elem())
}

type UninstallRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
}

func init() {
	types.Add("eam:UninstallRequestType", reflect.TypeOf((*UninstallRequestType)(nil)).Elem())
}

type UninstallResponse struct {
}

// Deprecated presence of unknown VMs is no more acceptable.
//
// An agent virtual machine has been found in the vCenter inventory that does
// not belong to any agency in this vSphere ESX Agent Manager server instance.
//
// This is an active remediable issue. To remediate, vSphere ESX Agent Manager
// powers off (if powered on) and deletes the agent virtual machine.
//
// This structure may be used only with operations rendered under `/eam`.
type UnknownAgentVm struct {
	HostIssue

	// The unknown agent virtual machine.
	//
	// Refers instance of `VirtualMachine`.
	Vm types.ManagedObjectReference `xml:"vm" json:"vm"`
}

func init() {
	types.Add("eam:UnknownAgentVm", reflect.TypeOf((*UnknownAgentVm)(nil)).Elem())
}

type UnregisterAgentVm UnregisterAgentVmRequestType

func init() {
	types.Add("eam:UnregisterAgentVm", reflect.TypeOf((*UnregisterAgentVm)(nil)).Elem())
}

// The parameters of `Agency.UnregisterAgentVm`.
//
// This structure may be used only with operations rendered under `/eam`.
type UnregisterAgentVmRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// The managed object reference to the agent VM.
	//
	// Refers instance of `VirtualMachine`.
	AgentVm types.ManagedObjectReference `xml:"agentVm" json:"agentVm"`
}

func init() {
	types.Add("eam:UnregisterAgentVmRequestType", reflect.TypeOf((*UnregisterAgentVmRequestType)(nil)).Elem())
}

type UnregisterAgentVmResponse struct {
}

type Update UpdateRequestType

func init() {
	types.Add("eam:Update", reflect.TypeOf((*Update)(nil)).Elem())
}

// The parameters of `Agency.Update`.
//
// This structure may be used only with operations rendered under `/eam`.
type UpdateRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// The new configuration for this <code>Agency</code>
	Config BaseAgencyConfigInfo `xml:"config,typeattr" json:"config"`
}

func init() {
	types.Add("eam:UpdateRequestType", reflect.TypeOf((*UpdateRequestType)(nil)).Elem())
}

type UpdateResponse struct {
}

// Deprecated as of vSphere 9.0. Please refer to vLCM Image APIs.
//
// A VIB module requires the host to be in maintenance mode, but the vSphere ESX Agent Manager
// is unable toput the host in maintenance mode.
//
// This can happen if there are virtual machines running on the host that cannot
// be moved and must be stopped before the host can enter maintenance mode.
//
// This is an active remediable issue. To remediate, the vSphere ESX Agent Manager will try again
// to put the host into maintenance mode. However, the vSphere ESX Agent Manager will not power
// off or move any virtual machines to put the host into maintenance mode. This must be
// done by the client.
//
// This structure may be used only with operations rendered under `/eam`.
type VibCannotPutHostInMaintenanceMode struct {
	VibIssue
}

func init() {
	types.Add("eam:VibCannotPutHostInMaintenanceMode", reflect.TypeOf((*VibCannotPutHostInMaintenanceMode)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. Please refer to vLCM Image APIs.
//
// ESXi host is in Maintenance Mode.
//
// This prevents powering on and
// re-configuring Agent Virtual Machines. Also if the host's entering in
// Maintenance Mode was initiated by vSphere Esx Agent Manager, the same is
// responsible to initiate exit Maintenance Mode.
//
// This is an active remediable issue. To remediate, vSphere ESX Agent Manager
// puts the host out of Maintenance Mode.
//
// This structure may be used only with operations rendered under `/eam`.
type VibCannotPutHostOutOfMaintenanceMode struct {
	VibIssue
}

func init() {
	types.Add("eam:VibCannotPutHostOutOfMaintenanceMode", reflect.TypeOf((*VibCannotPutHostOutOfMaintenanceMode)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. Please refer to vLCM Image APIs.
//
// A VIB module is expected to be installed on a host, but the dependencies,
// describred within the module, were not satisfied by the host.
//
// This is an active remediable issue. To remediate, vSphere ESX Agent Manager
// attempts the VIB installation again.
//
// This structure may be used only with operations rendered under `/eam`.
type VibDependenciesNotMetByHost struct {
	VibNotInstalled
}

func init() {
	types.Add("eam:VibDependenciesNotMetByHost", reflect.TypeOf((*VibDependenciesNotMetByHost)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. Please refer to vLCM Image APIs.
//
// A VIB module is expected to be installed on a host, but it failed to install
// since the VIB package is in an invalid format.
//
// The installation is unlikely to
// succeed until the solution providing the bundle has been upgraded or patched to
// provide a valid VIB package.
//
// This is an active remediable issue. To remediate, vSphere ESX Agent Manager attempts the VIB installation again.
//
// This structure may be used only with operations rendered under `/eam`.
type VibInvalidFormat struct {
	VibNotInstalled
}

func init() {
	types.Add("eam:VibInvalidFormat", reflect.TypeOf((*VibInvalidFormat)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. Please refer to vLCM Image APIs.
//
// Base class for all issues related to the VIB modules that belong to an agent.
//
// This structure may be used only with operations rendered under `/eam`.
type VibIssue struct {
	AgentIssue
}

func init() {
	types.Add("eam:VibIssue", reflect.TypeOf((*VibIssue)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. Please refer to vLCM Image APIs.
//
// A VIB module is expected to be installed/removed on a host, but it has not
// been installed/removed.
//
// Typically, a more specific issue (a subclass of this
// issue) indicates the particular reason why the VIB module operation failed.
//
// This is an active remediable issue. To remediate, vSphere ESX Agent Manager
// attempts the VIB operation again.
// In case of unreachable host vSphere ESX Agent Manager will remediate the
// issue automatically when the host becomes reachable.
//
// This structure may be used only with operations rendered under `/eam`.
type VibNotInstalled struct {
	VibIssue
}

func init() {
	types.Add("eam:VibNotInstalled", reflect.TypeOf((*VibNotInstalled)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. Please refer to vLCM Image APIs.
//
// A VIB module is expected to be installed on a host, but the system
// requirements, describred within the module, were not satisfied by the host.
//
// This is an active remediable issue. To remediate, vSphere ESX Agent Manager
// attempts the VIB installation again.
//
// This structure may be used only with operations rendered under `/eam`.
type VibRequirementsNotMetByHost struct {
	VibNotInstalled
}

func init() {
	types.Add("eam:VibRequirementsNotMetByHost", reflect.TypeOf((*VibRequirementsNotMetByHost)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. Please refer to vLCM Image APIs.
//
// A VIB module has been uploaded to the host, but will not be fully installed
// until the host has been put in maintenance mode.
//
// This is an active remediable issue. To remediate, vSphere ESX Agent Manager puts the host into maintenance
// mode.
//
// This structure may be used only with operations rendered under `/eam`.
type VibRequiresHostInMaintenanceMode struct {
	VibIssue
}

func init() {
	types.Add("eam:VibRequiresHostInMaintenanceMode", reflect.TypeOf((*VibRequiresHostInMaintenanceMode)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. Please refer to vLCM Image APIs.
//
// A VIB module has been uploaded to the host, but will not be activated
// until the host is rebooted.
//
// This is an active remediable issue. To remediate, vSphere ESX Agent Manager puts the host into maintenance
// mode and reboots it.
//
// This structure may be used only with operations rendered under `/eam`.
type VibRequiresHostReboot struct {
	VibIssue
}

func init() {
	types.Add("eam:VibRequiresHostReboot", reflect.TypeOf((*VibRequiresHostReboot)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. Please refer to vLCM Image APIs.
//
// A VIB module failed to install, but failed to do so because automatic installation
// by vSphere ESX Agent Manager is not allowed on the host.
//
// This is a passive remediable issue. To remediate, go to VMware Update Manager
// and install the required bulletins on the host or add the bulletins to the
// host's image profile.
//
// This structure may be used only with operations rendered under `/eam`.
type VibRequiresManualInstallation struct {
	VibIssue

	// A non-empty array of bulletins required to be installed on the host.
	Bulletin []string `xml:"bulletin" json:"bulletin"`
}

func init() {
	types.Add("eam:VibRequiresManualInstallation", reflect.TypeOf((*VibRequiresManualInstallation)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. Please refer to vLCM Image APIs.
//
// A VIB module failed to uninstall, but failed to do so because automatic uninstallation
// by vSphere ESX Agent Manager is not allowed on the host.
//
// This is a passive remediable issue. To remediate, go to VMware Update Manager
// and uninstall the required bulletins on the host or remove the bulletins from the
// host's image profile.
//
// This structure may be used only with operations rendered under `/eam`.
type VibRequiresManualUninstallation struct {
	VibIssue

	// A non-empty array of bulletins required to be uninstalled on the host.
	Bulletin []string `xml:"bulletin" json:"bulletin"`
}

func init() {
	types.Add("eam:VibRequiresManualUninstallation", reflect.TypeOf((*VibRequiresManualUninstallation)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. Please refer to vLCM Image APIs.
//
// A data entity providing information about a VIB.
//
// This abstraction contains only those of the VIB attributes which convey
// important information for the client to identify, preview and select VIBs.
//
// This structure may be used only with operations rendered under `/eam`.
type VibVibInfo struct {
	types.DynamicData

	Id           string                  `xml:"id" json:"id"`
	Name         string                  `xml:"name" json:"name"`
	Version      string                  `xml:"version" json:"version"`
	Vendor       string                  `xml:"vendor" json:"vendor"`
	Summary      string                  `xml:"summary" json:"summary"`
	SoftwareTags *VibVibInfoSoftwareTags `xml:"softwareTags,omitempty" json:"softwareTags,omitempty"`
	ReleaseDate  time.Time               `xml:"releaseDate" json:"releaseDate"`
}

func init() {
	types.Add("eam:VibVibInfo", reflect.TypeOf((*VibVibInfo)(nil)).Elem())
}

// A data entity providing information about software tags of a VIB
//
// This structure may be used only with operations rendered under `/eam`.
type VibVibInfoSoftwareTags struct {
	types.DynamicData

	Tags []string `xml:"tags,omitempty" json:"tags,omitempty"`
}

func init() {
	types.Add("eam:VibVibInfoSoftwareTags", reflect.TypeOf((*VibVibInfoSoftwareTags)(nil)).Elem())
}

// Specifies an SSL policy that trusts any SSL certificate.
//
// This structure may be used only with operations rendered under `/eam`.
type VibVibServicesAnyCertificate struct {
	VibVibServicesSslTrust
}

func init() {
	types.Add("eam:VibVibServicesAnyCertificate", reflect.TypeOf((*VibVibServicesAnyCertificate)(nil)).Elem())
	types.AddMinAPIVersionForType("eam:VibVibServicesAnyCertificate", "8.2")
}

// Specifies an SSL policy that trusts one specific pinned PEM encoded SSL
// certificate.
//
// This structure may be used only with operations rendered under `/eam`.
type VibVibServicesPinnedPemCertificate struct {
	VibVibServicesSslTrust

	// PEM encoded pinned SSL certificate of the server that needs to be
	// trusted.
	SslCertificate string `xml:"sslCertificate" json:"sslCertificate"`
}

func init() {
	types.Add("eam:VibVibServicesPinnedPemCertificate", reflect.TypeOf((*VibVibServicesPinnedPemCertificate)(nil)).Elem())
	types.AddMinAPIVersionForType("eam:VibVibServicesPinnedPemCertificate", "8.2")
}

// This structure may be used only with operations rendered under `/eam`.
type VibVibServicesSslTrust struct {
	types.DynamicData
}

func init() {
	types.Add("eam:VibVibServicesSslTrust", reflect.TypeOf((*VibVibServicesSslTrust)(nil)).Elem())
	types.AddMinAPIVersionForType("eam:VibVibServicesSslTrust", "8.2")
}

// An agent virtual machine is corrupted.
//
// This is an active remediable issue. To remediate, vSphere ESX Agent Manager deletes and
// reprovisions the agent virtual machine. To remediate manually, fix the missing file issue and power on the
// agent virtual machine.
//
// This structure may be used only with operations rendered under `/eam`.
type VmCorrupted struct {
	VmIssue

	// An optional path for a missing file.
	MissingFile string `xml:"missingFile,omitempty" json:"missingFile,omitempty"`
}

func init() {
	types.Add("eam:VmCorrupted", reflect.TypeOf((*VmCorrupted)(nil)).Elem())
}

// An agent virtual machine is expected to be removed from a host, but the agent virtual machine has not
// been removed.
//
// Typically, a more specific issue (a subclass of this issue)
// indicates the particular reason why vSphere ESX Agent Manager was unable to remove the
// agent virtual machine, such as the host is in maintenance mode, powered off or in standby
// mode.
//
// This is an active remediable issue. To remediate, vSphere ESX Agent Manager redeploys the agent.
//
// This structure may be used only with operations rendered under `/eam`.
type VmDeployed struct {
	VmIssue
}

func init() {
	types.Add("eam:VmDeployed", reflect.TypeOf((*VmDeployed)(nil)).Elem())
}

// The VM hook remediation failed.
//
// In order to remediate the issue:
// Resolve the issue via vLCM System VMs API and process the hook within the
// timeout configured for the System VM Solution this issue's VM belongs to.
//
// This structure may be used only with operations rendered under `/eam`.
type VmHookFailed struct {
	VmIssue
}

func init() {
	types.Add("eam:VmHookFailed", reflect.TypeOf((*VmHookFailed)(nil)).Elem())
	types.AddMinAPIVersionForType("eam:VmHookFailed", "9.0")
}

// The VM hook remediation timed out.
//
// In order to remediate the issue:
// Resolve the issue via vLCM System VMs API and process the hook within the
// timeout configured for the System VM Solution this issue's VM belongs to.
//
// This structure may be used only with operations rendered under `/eam`.
type VmHookTimedout struct {
	VmIssue
}

func init() {
	types.Add("eam:VmHookTimedout", reflect.TypeOf((*VmHookTimedout)(nil)).Elem())
	types.AddMinAPIVersionForType("eam:VmHookTimedout", "9.0")
}

// The connection state of the agent Virtual Machine is
// `inaccessible`.
//
// In order to remediate the issue:
//   - Mark the VM for removal using the `EsxAgentManager.EsxAgentManager_MarkForRemoval`
//     API.
//   - Do the necessary changes to ensure that the connection state of the VM is
//     `connected`.
//
// NOTE: When the HA is enabled on the cluster these issues may be transient and
// automatically remediated.
//
// This structure may be used only with operations rendered under `/eam`.
type VmInaccessible struct {
	VmIssue
}

func init() {
	types.Add("eam:VmInaccessible", reflect.TypeOf((*VmInaccessible)(nil)).Elem())
	types.AddMinAPIVersionForType("eam:VmInaccessible", "8.3")
}

// Base class for all issues related to the deployed virtual machine for a
// particular agent.
//
// This structure may be used only with operations rendered under `/eam`.
type VmIssue struct {
	AgentIssue

	// The virtual machine to which this issue is related.
	//
	// Refers instance of `VirtualMachine`.
	Vm types.ManagedObjectReference `xml:"vm" json:"vm"`
}

func init() {
	types.Add("eam:VmIssue", reflect.TypeOf((*VmIssue)(nil)).Elem())
}

// Deprecated template agent VMs are not used anymore by VM deployment and
// monitoring.
//
// An agent virtual machine is a virtual machine template.
//
// This is an active remediable issue. To remediate, vSphere ESX Agent Manager
// converts the agent virtual machine template to a virtual machine.
//
// This structure may be used only with operations rendered under `/eam`.
type VmMarkedAsTemplate struct {
	VmIssue
}

func init() {
	types.Add("eam:VmMarkedAsTemplate", reflect.TypeOf((*VmMarkedAsTemplate)(nil)).Elem())
}

// An agent virtual machine is expected to be deployed on a host, but the agent virtual machine has not
// been deployed.
//
// Typically, a more specific issue (a subclass of this issue)
// indicates the particular reason why vSphere ESX Agent Manager was unable to deploy the
// agent, such as being unable to access the OVF package for the agent or a missing host
// configuration. This issue can also happen if the agent virtual machine is explicitly deleted
// from the host.
//
// This is an active remediable issue. To remediate, vSphere ESX Agent Manager redeploys the agent virtual machine.
//
// This structure may be used only with operations rendered under `/eam`.
type VmNotDeployed struct {
	AgentIssue
}

func init() {
	types.Add("eam:VmNotDeployed", reflect.TypeOf((*VmNotDeployed)(nil)).Elem())
}

// An agent virtual machine exists on a host, but the host is no longer part of scope for the
// agency.
//
// This typically happens if a host is disconnected when the agency
// configuration is changed.
//
// This is an active remediable issue. To remediate, vSphere ESX Agent Manager deletes the agent virtual machine.
//
// This structure may be used only with operations rendered under `/eam`.
type VmOrphaned struct {
	VmIssue
}

func init() {
	types.Add("eam:VmOrphaned", reflect.TypeOf((*VmOrphaned)(nil)).Elem())
}

// An agent virtual machine is expected to be powered on, but the agent virtual machine is powered off.
//
// This is an active remediable issue. To remediate, vSphere ESX Agent Manager
// powers on the agent virtual machine.
//
// This structure may be used only with operations rendered under `/eam`.
type VmPoweredOff struct {
	VmIssue
}

func init() {
	types.Add("eam:VmPoweredOff", reflect.TypeOf((*VmPoweredOff)(nil)).Elem())
}

// An agent virtual machine is expected to be powered off, but the agent virtual machine is powered on.
//
// This is an active remediable issue. To remediate, vSphere ESX Agent Manager
// powers off the agent virtual machine.
//
// This structure may be used only with operations rendered under `/eam`.
type VmPoweredOn struct {
	VmIssue
}

func init() {
	types.Add("eam:VmPoweredOn", reflect.TypeOf((*VmPoweredOn)(nil)).Elem())
}

// An agent virtual machine is protected from modifications
// (example: HA recovery).
//
// This is an active remediable issue. To remediate, vSphere ESX Agent Manager
// modifies the virtual machine.
//
// This structure may be used only with operations rendered under `/eam`.
type VmProtected struct {
	VmIssue
}

func init() {
	types.Add("eam:VmProtected", reflect.TypeOf((*VmProtected)(nil)).Elem())
	types.AddMinAPIVersionForType("eam:VmProtected", "9.0")
}

// An agent virtual machine is expected to be deployed on a host, but the agent
// virtual machine cannot be deployed because the host is in Maintenance Mode.
//
// This is an active remediable issue. To remediate, vSphere ESX Agent Manager
// takes the host out of Maintenance Mode and deploys the agent virtual machine.
//
// Resolving this issue in vSphere Lifecycle Manager environment will be no-op.
// In those cases user must take the host out of Maintenance Mode manually or
// wait vSphere Lifecycle Manager cluster remediation to complete (if any).
//
// This structure may be used only with operations rendered under `/eam`.
type VmRequiresHostOutOfMaintenanceMode struct {
	VmNotDeployed
}

func init() {
	types.Add("eam:VmRequiresHostOutOfMaintenanceMode", reflect.TypeOf((*VmRequiresHostOutOfMaintenanceMode)(nil)).Elem())
	types.AddMinAPIVersionForType("eam:VmRequiresHostOutOfMaintenanceMode", "7.2")
}

// An agent virtual machine is expected to be powered on, but the agent virtual machine is suspended.
//
// This is an active remediable issue. To remediate, vSphere ESX Agent Manager powers on the agent virtual machine.
//
// This structure may be used only with operations rendered under `/eam`.
type VmSuspended struct {
	VmIssue
}

func init() {
	types.Add("eam:VmSuspended", reflect.TypeOf((*VmSuspended)(nil)).Elem())
}

// Deprecated eAM does not try to override any action powerful user has taken.
//
// An agent virtual machine is expected to be located in a designated agent
// virtual machine folder, but is found in a different folder.
//
// This is an active remediable issue. To remediate, vSphere ESX Agent Manager
// moves the agent virtual machine back into the designated agent folder.
//
// This structure may be used only with operations rendered under `/eam`.
type VmWrongFolder struct {
	VmIssue

	// The folder in which the virtual machine currently resides.
	//
	// Refers instance of `Folder`.
	CurrentFolder types.ManagedObjectReference `xml:"currentFolder" json:"currentFolder"`
	// The ESX agent folder in which the virtual machine should reside.
	//
	// Refers instance of `Folder`.
	RequiredFolder types.ManagedObjectReference `xml:"requiredFolder" json:"requiredFolder"`
}

func init() {
	types.Add("eam:VmWrongFolder", reflect.TypeOf((*VmWrongFolder)(nil)).Elem())
}

// Deprecated eAM does not try to override any action powerful user has taken.
//
// An agent virtual machine is expected to be located in a designated agent
// virtual machine resource pool, but is found in a different resource pool.
//
// This is an active remediable issue. To remediate, vSphere ESX Agent Manager
// moves the agent virtual machine back into the designated agent resource pool.
//
// This structure may be used only with operations rendered under `/eam`.
type VmWrongResourcePool struct {
	VmIssue

	// The resource pool in which the VM currently resides.
	//
	// Refers instance of `ResourcePool`.
	CurrentResourcePool types.ManagedObjectReference `xml:"currentResourcePool" json:"currentResourcePool"`
	// The ESX agent resource pool in which the VM should reside.
	//
	// Refers instance of `ResourcePool`.
	RequiredResourcePool types.ManagedObjectReference `xml:"requiredResourcePool" json:"requiredResourcePool"`
}

func init() {
	types.Add("eam:VmWrongResourcePool", reflect.TypeOf((*VmWrongResourcePool)(nil)).Elem())
}

type VersionURI string

func init() {
	types.Add("eam:versionURI", reflect.TypeOf((*VersionURI)(nil)).Elem())
}
