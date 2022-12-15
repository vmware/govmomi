/*
Copyright (c) 2021 VMware, Inc. All Rights Reserved.

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

type AddIssueRequestType struct {
	This  types.ManagedObjectReference `xml:"_this" json:"_this"`
	Issue BaseIssue                    `xml:"issue,typeattr" json:"issue"`
}

func init() {
	types.Add("eam:AddIssueRequestType", reflect.TypeOf((*AddIssueRequestType)(nil)).Elem())
}

type AddIssueResponse struct {
	Returnval BaseIssue `xml:"returnval,typeattr" json:"returnval"`
}

type AgencyComputeResourceScope struct {
	AgencyScope

	ComputeResource []types.ManagedObjectReference `xml:"computeResource,omitempty" json:"computeResource,omitempty"`
}

func init() {
	types.Add("eam:AgencyComputeResourceScope", reflect.TypeOf((*AgencyComputeResourceScope)(nil)).Elem())
}

type AgencyConfigInfo struct {
	types.DynamicData

	AgentConfig                                   []AgentConfigInfo              `xml:"agentConfig,omitempty" json:"agentConfig,omitempty"`
	Scope                                         BaseAgencyScope                `xml:"scope,omitempty,typeattr" json:"scope,omitempty"`
	ManuallyMarkAgentVmAvailableAfterProvisioning *bool                          `xml:"manuallyMarkAgentVmAvailableAfterProvisioning" json:"manuallyMarkAgentVmAvailableAfterProvisioning,omitempty"`
	ManuallyMarkAgentVmAvailableAfterPowerOn      *bool                          `xml:"manuallyMarkAgentVmAvailableAfterPowerOn" json:"manuallyMarkAgentVmAvailableAfterPowerOn,omitempty"`
	OptimizedDeploymentEnabled                    *bool                          `xml:"optimizedDeploymentEnabled" json:"optimizedDeploymentEnabled,omitempty"`
	AgentName                                     string                         `xml:"agentName,omitempty" json:"agentName,omitempty"`
	AgencyName                                    string                         `xml:"agencyName,omitempty" json:"agencyName,omitempty"`
	UseUuidVmName                                 *bool                          `xml:"useUuidVmName" json:"useUuidVmName,omitempty"`
	ManuallyProvisioned                           *bool                          `xml:"manuallyProvisioned" json:"manuallyProvisioned,omitempty"`
	ManuallyMonitored                             *bool                          `xml:"manuallyMonitored" json:"manuallyMonitored,omitempty"`
	BypassVumEnabled                              *bool                          `xml:"bypassVumEnabled" json:"bypassVumEnabled,omitempty"`
	AgentVmNetwork                                []types.ManagedObjectReference `xml:"agentVmNetwork,omitempty" json:"agentVmNetwork,omitempty"`
	AgentVmDatastore                              []types.ManagedObjectReference `xml:"agentVmDatastore,omitempty" json:"agentVmDatastore,omitempty"`
	PreferHostConfiguration                       *bool                          `xml:"preferHostConfiguration" json:"preferHostConfiguration,omitempty"`
	IpPool                                        *types.IpPool                  `xml:"ipPool,omitempty" json:"ipPool,omitempty"`
	ResourcePools                                 []AgencyVMResourcePool         `xml:"resourcePools,omitempty" json:"resourcePools,omitempty"`
	Folders                                       []AgencyVMFolder               `xml:"folders,omitempty" json:"folders,omitempty"`
}

func init() {
	types.Add("eam:AgencyConfigInfo", reflect.TypeOf((*AgencyConfigInfo)(nil)).Elem())
}

type AgencyDisabled struct {
	AgencyIssue
}

func init() {
	types.Add("eam:AgencyDisabled", reflect.TypeOf((*AgencyDisabled)(nil)).Elem())
}

type AgencyIssue struct {
	Issue

	Agency       types.ManagedObjectReference `xml:"agency" json:"agency"`
	AgencyName   string                       `xml:"agencyName" json:"agencyName"`
	SolutionId   string                       `xml:"solutionId" json:"solutionId"`
	SolutionName string                       `xml:"solutionName" json:"solutionName"`
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

type AgencyScope struct {
	types.DynamicData
}

func init() {
	types.Add("eam:AgencyScope", reflect.TypeOf((*AgencyScope)(nil)).Elem())
}

type AgencyVMFolder struct {
	types.DynamicData

	FolderId     types.ManagedObjectReference `xml:"folderId" json:"folderId"`
	DatacenterId types.ManagedObjectReference `xml:"datacenterId" json:"datacenterId"`
}

func init() {
	types.Add("eam:AgencyVMFolder", reflect.TypeOf((*AgencyVMFolder)(nil)).Elem())
}

type AgencyVMResourcePool struct {
	types.DynamicData

	ResourcePoolId    types.ManagedObjectReference `xml:"resourcePoolId" json:"resourcePoolId"`
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

type AgentConfigInfo struct {
	types.DynamicData

	ProductLineId               string                   `xml:"productLineId,omitempty" json:"productLineId,omitempty"`
	HostVersion                 string                   `xml:"hostVersion,omitempty" json:"hostVersion,omitempty"`
	OvfPackageUrl               string                   `xml:"ovfPackageUrl,omitempty" json:"ovfPackageUrl,omitempty"`
	OvfEnvironment              *AgentOvfEnvironmentInfo `xml:"ovfEnvironment,omitempty" json:"ovfEnvironment,omitempty"`
	VibUrl                      string                   `xml:"vibUrl,omitempty" json:"vibUrl,omitempty"`
	VibMatchingRules            []AgentVibMatchingRule   `xml:"vibMatchingRules,omitempty" json:"vibMatchingRules,omitempty"`
	VibName                     string                   `xml:"vibName,omitempty" json:"vibName,omitempty"`
	DvFilterEnabled             *bool                    `xml:"dvFilterEnabled" json:"dvFilterEnabled,omitempty"`
	RebootHostAfterVibUninstall *bool                    `xml:"rebootHostAfterVibUninstall" json:"rebootHostAfterVibUninstall,omitempty"`
	VmciService                 []string                 `xml:"vmciService,omitempty" json:"vmciService,omitempty"`
	OvfDiskProvisioning         string                   `xml:"ovfDiskProvisioning,omitempty" json:"ovfDiskProvisioning,omitempty"`
	VmStoragePolicies           []BaseAgentStoragePolicy `xml:"vmStoragePolicies,omitempty,typeattr" json:"vmStoragePolicies,omitempty"`
}

func init() {
	types.Add("eam:AgentConfigInfo", reflect.TypeOf((*AgentConfigInfo)(nil)).Elem())
}

type AgentIssue struct {
	AgencyIssue

	Agent     types.ManagedObjectReference `xml:"agent" json:"agent"`
	AgentName string                       `xml:"agentName" json:"agentName"`
	Host      types.ManagedObjectReference `xml:"host" json:"host"`
	HostName  string                       `xml:"hostName" json:"hostName"`
}

func init() {
	types.Add("eam:AgentIssue", reflect.TypeOf((*AgentIssue)(nil)).Elem())
}

type AgentOvfEnvironmentInfo struct {
	types.DynamicData

	OvfProperty []AgentOvfEnvironmentInfoOvfProperty `xml:"ovfProperty,omitempty" json:"ovfProperty,omitempty"`
}

func init() {
	types.Add("eam:AgentOvfEnvironmentInfo", reflect.TypeOf((*AgentOvfEnvironmentInfo)(nil)).Elem())
}

type AgentOvfEnvironmentInfoOvfProperty struct {
	types.DynamicData

	Key   string `xml:"key" json:"key"`
	Value string `xml:"value" json:"value"`
}

func init() {
	types.Add("eam:AgentOvfEnvironmentInfoOvfProperty", reflect.TypeOf((*AgentOvfEnvironmentInfoOvfProperty)(nil)).Elem())
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

type AgentRuntimeInfo struct {
	EamObjectRuntimeInfo

	VmPowerState         types.VirtualMachinePowerState `xml:"vmPowerState" json:"vmPowerState"`
	ReceivingHeartBeat   bool                           `xml:"receivingHeartBeat" json:"receivingHeartBeat"`
	Host                 *types.ManagedObjectReference  `xml:"host,omitempty" json:"host,omitempty"`
	Vm                   *types.ManagedObjectReference  `xml:"vm,omitempty" json:"vm,omitempty"`
	VmIp                 string                         `xml:"vmIp,omitempty" json:"vmIp,omitempty"`
	VmName               string                         `xml:"vmName" json:"vmName"`
	EsxAgentResourcePool *types.ManagedObjectReference  `xml:"esxAgentResourcePool,omitempty" json:"esxAgentResourcePool,omitempty"`
	EsxAgentFolder       *types.ManagedObjectReference  `xml:"esxAgentFolder,omitempty" json:"esxAgentFolder,omitempty"`
	InstalledBulletin    []string                       `xml:"installedBulletin,omitempty" json:"installedBulletin,omitempty"`
	InstalledVibs        []VibVibInfo                   `xml:"installedVibs,omitempty" json:"installedVibs,omitempty"`
	Agency               *types.ManagedObjectReference  `xml:"agency,omitempty" json:"agency,omitempty"`
	VmHook               *AgentVmHook                   `xml:"vmHook,omitempty" json:"vmHook,omitempty"`
}

func init() {
	types.Add("eam:AgentRuntimeInfo", reflect.TypeOf((*AgentRuntimeInfo)(nil)).Elem())
}

type AgentStoragePolicy struct {
	types.DynamicData
}

func init() {
	types.Add("eam:AgentStoragePolicy", reflect.TypeOf((*AgentStoragePolicy)(nil)).Elem())
}

type AgentVibMatchingRule struct {
	types.DynamicData

	VibNameRegex    string `xml:"vibNameRegex" json:"vibNameRegex"`
	VibVersionRegex string `xml:"vibVersionRegex" json:"vibVersionRegex"`
}

func init() {
	types.Add("eam:AgentVibMatchingRule", reflect.TypeOf((*AgentVibMatchingRule)(nil)).Elem())
}

type AgentVmHook struct {
	types.DynamicData

	Vm      types.ManagedObjectReference `xml:"vm" json:"vm"`
	VmState string                       `xml:"vmState" json:"vmState"`
}

func init() {
	types.Add("eam:AgentVmHook", reflect.TypeOf((*AgentVmHook)(nil)).Elem())
}

type AgentVsanStoragePolicy struct {
	AgentStoragePolicy

	ProfileId string `xml:"profileId" json:"profileId"`
}

func init() {
	types.Add("eam:AgentVsanStoragePolicy", reflect.TypeOf((*AgentVsanStoragePolicy)(nil)).Elem())
}

type ArrayOfAgencyVMFolder struct {
	AgencyVMFolder []AgencyVMFolder `xml:"AgencyVMFolder,omitempty" json:"AgencyVMFolder,omitempty"`
}

func init() {
	types.Add("eam:ArrayOfAgencyVMFolder", reflect.TypeOf((*ArrayOfAgencyVMFolder)(nil)).Elem())
}

type ArrayOfAgencyVMResourcePool struct {
	AgencyVMResourcePool []AgencyVMResourcePool `xml:"AgencyVMResourcePool,omitempty" json:"AgencyVMResourcePool,omitempty"`
}

func init() {
	types.Add("eam:ArrayOfAgencyVMResourcePool", reflect.TypeOf((*ArrayOfAgencyVMResourcePool)(nil)).Elem())
}

type ArrayOfAgentConfigInfo struct {
	AgentConfigInfo []AgentConfigInfo `xml:"AgentConfigInfo,omitempty" json:"AgentConfigInfo,omitempty"`
}

func init() {
	types.Add("eam:ArrayOfAgentConfigInfo", reflect.TypeOf((*ArrayOfAgentConfigInfo)(nil)).Elem())
}

type ArrayOfAgentOvfEnvironmentInfoOvfProperty struct {
	AgentOvfEnvironmentInfoOvfProperty []AgentOvfEnvironmentInfoOvfProperty `xml:"AgentOvfEnvironmentInfoOvfProperty,omitempty" json:"AgentOvfEnvironmentInfoOvfProperty,omitempty"`
}

func init() {
	types.Add("eam:ArrayOfAgentOvfEnvironmentInfoOvfProperty", reflect.TypeOf((*ArrayOfAgentOvfEnvironmentInfoOvfProperty)(nil)).Elem())
}

type ArrayOfAgentStoragePolicy struct {
	AgentStoragePolicy []BaseAgentStoragePolicy `xml:"AgentStoragePolicy,omitempty,typeattr" json:"AgentStoragePolicy,omitempty"`
}

func init() {
	types.Add("eam:ArrayOfAgentStoragePolicy", reflect.TypeOf((*ArrayOfAgentStoragePolicy)(nil)).Elem())
}

type ArrayOfAgentVibMatchingRule struct {
	AgentVibMatchingRule []AgentVibMatchingRule `xml:"AgentVibMatchingRule,omitempty" json:"AgentVibMatchingRule,omitempty"`
}

func init() {
	types.Add("eam:ArrayOfAgentVibMatchingRule", reflect.TypeOf((*ArrayOfAgentVibMatchingRule)(nil)).Elem())
}

type ArrayOfIssue struct {
	Issue []BaseIssue `xml:"Issue,omitempty,typeattr" json:"Issue,omitempty"`
}

func init() {
	types.Add("eam:ArrayOfIssue", reflect.TypeOf((*ArrayOfIssue)(nil)).Elem())
}

type ArrayOfVibVibInfo struct {
	VibVibInfo []VibVibInfo `xml:"VibVibInfo,omitempty" json:"VibVibInfo,omitempty"`
}

func init() {
	types.Add("eam:ArrayOfVibVibInfo", reflect.TypeOf((*ArrayOfVibVibInfo)(nil)).Elem())
}

type CannotAccessAgentOVF struct {
	VmNotDeployed

	DownloadUrl string `xml:"downloadUrl" json:"downloadUrl"`
}

func init() {
	types.Add("eam:CannotAccessAgentOVF", reflect.TypeOf((*CannotAccessAgentOVF)(nil)).Elem())
}

type CannotAccessAgentVib struct {
	VibNotInstalled

	DownloadUrl string `xml:"downloadUrl" json:"downloadUrl"`
}

func init() {
	types.Add("eam:CannotAccessAgentVib", reflect.TypeOf((*CannotAccessAgentVib)(nil)).Elem())
}

type ClusterAgentAgentIssue struct {
	AgencyIssue

	Agent   types.ManagedObjectReference  `xml:"agent" json:"agent"`
	Cluster *types.ManagedObjectReference `xml:"cluster,omitempty" json:"cluster,omitempty"`
}

func init() {
	types.Add("eam:ClusterAgentAgentIssue", reflect.TypeOf((*ClusterAgentAgentIssue)(nil)).Elem())
}

type ClusterAgentInsufficientClusterResources struct {
	ClusterAgentVmPoweredOff
}

func init() {
	types.Add("eam:ClusterAgentInsufficientClusterResources", reflect.TypeOf((*ClusterAgentInsufficientClusterResources)(nil)).Elem())
}

type ClusterAgentInsufficientClusterSpace struct {
	ClusterAgentVmNotDeployed
}

func init() {
	types.Add("eam:ClusterAgentInsufficientClusterSpace", reflect.TypeOf((*ClusterAgentInsufficientClusterSpace)(nil)).Elem())
}

type ClusterAgentInvalidConfig struct {
	ClusterAgentVmIssue

	Error types.AnyType `xml:"error,typeattr" json:"error"`
}

func init() {
	types.Add("eam:ClusterAgentInvalidConfig", reflect.TypeOf((*ClusterAgentInvalidConfig)(nil)).Elem())
}

type ClusterAgentMissingClusterVmDatastore struct {
	ClusterAgentVmNotDeployed

	MissingDatastores []types.ManagedObjectReference `xml:"missingDatastores,omitempty" json:"missingDatastores,omitempty"`
}

func init() {
	types.Add("eam:ClusterAgentMissingClusterVmDatastore", reflect.TypeOf((*ClusterAgentMissingClusterVmDatastore)(nil)).Elem())
}

type ClusterAgentMissingClusterVmNetwork struct {
	ClusterAgentVmNotDeployed

	MissingNetworks []types.ManagedObjectReference `xml:"missingNetworks,omitempty" json:"missingNetworks,omitempty"`
	NetworkNames    []string                       `xml:"networkNames,omitempty" json:"networkNames,omitempty"`
}

func init() {
	types.Add("eam:ClusterAgentMissingClusterVmNetwork", reflect.TypeOf((*ClusterAgentMissingClusterVmNetwork)(nil)).Elem())
}

type ClusterAgentOvfInvalidProperty struct {
	ClusterAgentAgentIssue

	Error []types.LocalizedMethodFault `xml:"error,omitempty" json:"error,omitempty"`
}

func init() {
	types.Add("eam:ClusterAgentOvfInvalidProperty", reflect.TypeOf((*ClusterAgentOvfInvalidProperty)(nil)).Elem())
}

type ClusterAgentVmIssue struct {
	ClusterAgentAgentIssue

	Vm types.ManagedObjectReference `xml:"vm" json:"vm"`
}

func init() {
	types.Add("eam:ClusterAgentVmIssue", reflect.TypeOf((*ClusterAgentVmIssue)(nil)).Elem())
}

type ClusterAgentVmNotDeployed struct {
	ClusterAgentAgentIssue
}

func init() {
	types.Add("eam:ClusterAgentVmNotDeployed", reflect.TypeOf((*ClusterAgentVmNotDeployed)(nil)).Elem())
}

type ClusterAgentVmNotRemoved struct {
	ClusterAgentVmIssue
}

func init() {
	types.Add("eam:ClusterAgentVmNotRemoved", reflect.TypeOf((*ClusterAgentVmNotRemoved)(nil)).Elem())
}

type ClusterAgentVmPoweredOff struct {
	ClusterAgentVmIssue
}

func init() {
	types.Add("eam:ClusterAgentVmPoweredOff", reflect.TypeOf((*ClusterAgentVmPoweredOff)(nil)).Elem())
}

type ClusterAgentVmPoweredOn struct {
	ClusterAgentVmIssue
}

func init() {
	types.Add("eam:ClusterAgentVmPoweredOn", reflect.TypeOf((*ClusterAgentVmPoweredOn)(nil)).Elem())
}

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

type CreateAgencyRequestType struct {
	This             types.ManagedObjectReference `xml:"_this" json:"_this"`
	AgencyConfigInfo BaseAgencyConfigInfo         `xml:"agencyConfigInfo,typeattr" json:"agencyConfigInfo"`
	InitialGoalState string                       `xml:"initialGoalState" json:"initialGoalState"`
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

type DisabledClusterFault struct {
	EamAppFault

	DisabledComputeResource []types.ManagedObjectReference `xml:"disabledComputeResource,omitempty" json:"disabledComputeResource,omitempty"`
}

func init() {
	types.Add("eam:DisabledClusterFault", reflect.TypeOf((*DisabledClusterFault)(nil)).Elem())
}

type DisabledClusterFaultFault DisabledClusterFault

func init() {
	types.Add("eam:DisabledClusterFaultFault", reflect.TypeOf((*DisabledClusterFaultFault)(nil)).Elem())
}

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

type EamObjectRuntimeInfo struct {
	types.DynamicData

	Status    string                       `xml:"status" json:"status"`
	Issue     []BaseIssue                  `xml:"issue,omitempty,typeattr" json:"issue,omitempty"`
	GoalState string                       `xml:"goalState" json:"goalState"`
	Entity    types.ManagedObjectReference `xml:"entity" json:"entity"`
}

func init() {
	types.Add("eam:EamObjectRuntimeInfo", reflect.TypeOf((*EamObjectRuntimeInfo)(nil)).Elem())
}

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

type ExtensibleIssue struct {
	Issue

	TypeId   string                        `xml:"typeId" json:"typeId"`
	Argument []types.KeyAnyValue           `xml:"argument,omitempty" json:"argument,omitempty"`
	Target   *types.ManagedObjectReference `xml:"target,omitempty" json:"target,omitempty"`
	Agent    *types.ManagedObjectReference `xml:"agent,omitempty" json:"agent,omitempty"`
	Agency   *types.ManagedObjectReference `xml:"agency,omitempty" json:"agency,omitempty"`
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

type HostInMaintenanceMode struct {
	VmDeployed
}

func init() {
	types.Add("eam:HostInMaintenanceMode", reflect.TypeOf((*HostInMaintenanceMode)(nil)).Elem())
}

type HostInStandbyMode struct {
	VmDeployed
}

func init() {
	types.Add("eam:HostInStandbyMode", reflect.TypeOf((*HostInStandbyMode)(nil)).Elem())
}

type HostIssue struct {
	Issue

	Host types.ManagedObjectReference `xml:"host" json:"host"`
}

func init() {
	types.Add("eam:HostIssue", reflect.TypeOf((*HostIssue)(nil)).Elem())
}

type HostPoweredOff struct {
	VmDeployed
}

func init() {
	types.Add("eam:HostPoweredOff", reflect.TypeOf((*HostPoweredOff)(nil)).Elem())
}

type ImmediateHostRebootRequired struct {
	VibIssue
}

func init() {
	types.Add("eam:ImmediateHostRebootRequired", reflect.TypeOf((*ImmediateHostRebootRequired)(nil)).Elem())
}

type IncompatibleHostVersion struct {
	VmNotDeployed
}

func init() {
	types.Add("eam:IncompatibleHostVersion", reflect.TypeOf((*IncompatibleHostVersion)(nil)).Elem())
}

type InsufficientIpAddresses struct {
	VmPoweredOff

	Network types.ManagedObjectReference `xml:"network" json:"network"`
}

func init() {
	types.Add("eam:InsufficientIpAddresses", reflect.TypeOf((*InsufficientIpAddresses)(nil)).Elem())
}

type InsufficientResources struct {
	VmNotDeployed
}

func init() {
	types.Add("eam:InsufficientResources", reflect.TypeOf((*InsufficientResources)(nil)).Elem())
}

type InsufficientSpace struct {
	VmNotDeployed
}

func init() {
	types.Add("eam:InsufficientSpace", reflect.TypeOf((*InsufficientSpace)(nil)).Elem())
}

type IntegrityAgencyCannotDeleteSoftware struct {
	IntegrityAgencyVUMIssue
}

func init() {
	types.Add("eam:IntegrityAgencyCannotDeleteSoftware", reflect.TypeOf((*IntegrityAgencyCannotDeleteSoftware)(nil)).Elem())
}

type IntegrityAgencyCannotStageSoftware struct {
	IntegrityAgencyVUMIssue
}

func init() {
	types.Add("eam:IntegrityAgencyCannotStageSoftware", reflect.TypeOf((*IntegrityAgencyCannotStageSoftware)(nil)).Elem())
}

type IntegrityAgencyVUMIssue struct {
	AgencyIssue
}

func init() {
	types.Add("eam:IntegrityAgencyVUMIssue", reflect.TypeOf((*IntegrityAgencyVUMIssue)(nil)).Elem())
}

type IntegrityAgencyVUMUnavailable struct {
	IntegrityAgencyVUMIssue
}

func init() {
	types.Add("eam:IntegrityAgencyVUMUnavailable", reflect.TypeOf((*IntegrityAgencyVUMUnavailable)(nil)).Elem())
}

type InvalidAgencyScope struct {
	EamFault

	UnknownComputeResource []types.ManagedObjectReference `xml:"unknownComputeResource,omitempty" json:"unknownComputeResource,omitempty"`
}

func init() {
	types.Add("eam:InvalidAgencyScope", reflect.TypeOf((*InvalidAgencyScope)(nil)).Elem())
}

type InvalidAgencyScopeFault InvalidAgencyScope

func init() {
	types.Add("eam:InvalidAgencyScopeFault", reflect.TypeOf((*InvalidAgencyScopeFault)(nil)).Elem())
}

type InvalidAgentConfiguration struct {
	EamFault

	InvalidAgentConfiguration *AgentConfigInfo `xml:"invalidAgentConfiguration,omitempty" json:"invalidAgentConfiguration,omitempty"`
	InvalidField              string           `xml:"invalidField,omitempty" json:"invalidField,omitempty"`
}

func init() {
	types.Add("eam:InvalidAgentConfiguration", reflect.TypeOf((*InvalidAgentConfiguration)(nil)).Elem())
}

type InvalidAgentConfigurationFault InvalidAgentConfiguration

func init() {
	types.Add("eam:InvalidAgentConfigurationFault", reflect.TypeOf((*InvalidAgentConfigurationFault)(nil)).Elem())
}

type InvalidConfig struct {
	VmIssue

	Error types.AnyType `xml:"error,typeattr" json:"error"`
}

func init() {
	types.Add("eam:InvalidConfig", reflect.TypeOf((*InvalidConfig)(nil)).Elem())
}

type InvalidUrl struct {
	EamFault

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

type Issue struct {
	types.DynamicData

	Key         int32     `xml:"key" json:"key"`
	Description string    `xml:"description" json:"description"`
	Time        time.Time `xml:"time" json:"time"`
}

func init() {
	types.Add("eam:Issue", reflect.TypeOf((*Issue)(nil)).Elem())
}

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

type MissingAgentIpPool struct {
	VmPoweredOff

	Network types.ManagedObjectReference `xml:"network" json:"network"`
}

func init() {
	types.Add("eam:MissingAgentIpPool", reflect.TypeOf((*MissingAgentIpPool)(nil)).Elem())
}

type MissingDvFilterSwitch struct {
	AgentIssue
}

func init() {
	types.Add("eam:MissingDvFilterSwitch", reflect.TypeOf((*MissingDvFilterSwitch)(nil)).Elem())
}

type NoAgentVmDatastore struct {
	VmNotDeployed
}

func init() {
	types.Add("eam:NoAgentVmDatastore", reflect.TypeOf((*NoAgentVmDatastore)(nil)).Elem())
}

type NoAgentVmNetwork struct {
	VmNotDeployed
}

func init() {
	types.Add("eam:NoAgentVmNetwork", reflect.TypeOf((*NoAgentVmNetwork)(nil)).Elem())
}

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

type NoCustomAgentVmDatastore struct {
	NoAgentVmDatastore

	CustomAgentVmDatastore     []types.ManagedObjectReference `xml:"customAgentVmDatastore" json:"customAgentVmDatastore"`
	CustomAgentVmDatastoreName []string                       `xml:"customAgentVmDatastoreName" json:"customAgentVmDatastoreName"`
}

func init() {
	types.Add("eam:NoCustomAgentVmDatastore", reflect.TypeOf((*NoCustomAgentVmDatastore)(nil)).Elem())
}

type NoCustomAgentVmNetwork struct {
	NoAgentVmNetwork

	CustomAgentVmNetwork     []types.ManagedObjectReference `xml:"customAgentVmNetwork" json:"customAgentVmNetwork"`
	CustomAgentVmNetworkName []string                       `xml:"customAgentVmNetworkName" json:"customAgentVmNetworkName"`
}

func init() {
	types.Add("eam:NoCustomAgentVmNetwork", reflect.TypeOf((*NoCustomAgentVmNetwork)(nil)).Elem())
}

type NoDiscoverableAgentVmDatastore struct {
	VmNotDeployed
}

func init() {
	types.Add("eam:NoDiscoverableAgentVmDatastore", reflect.TypeOf((*NoDiscoverableAgentVmDatastore)(nil)).Elem())
}

type NoDiscoverableAgentVmNetwork struct {
	VmNotDeployed
}

func init() {
	types.Add("eam:NoDiscoverableAgentVmNetwork", reflect.TypeOf((*NoDiscoverableAgentVmNetwork)(nil)).Elem())
}

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

type OrphanedAgency struct {
	AgencyIssue
}

func init() {
	types.Add("eam:OrphanedAgency", reflect.TypeOf((*OrphanedAgency)(nil)).Elem())
}

type OrphanedDvFilterSwitch struct {
	HostIssue
}

func init() {
	types.Add("eam:OrphanedDvFilterSwitch", reflect.TypeOf((*OrphanedDvFilterSwitch)(nil)).Elem())
}

type OvfInvalidFormat struct {
	VmNotDeployed

	Error []types.LocalizedMethodFault `xml:"error,omitempty" json:"error,omitempty"`
}

func init() {
	types.Add("eam:OvfInvalidFormat", reflect.TypeOf((*OvfInvalidFormat)(nil)).Elem())
}

type OvfInvalidProperty struct {
	AgentIssue

	Error []types.LocalizedMethodFault `xml:"error,omitempty" json:"error,omitempty"`
}

func init() {
	types.Add("eam:OvfInvalidProperty", reflect.TypeOf((*OvfInvalidProperty)(nil)).Elem())
}

type PersonalityAgencyCannotConfigureSolutions struct {
	PersonalityAgencyPMIssue

	Cr                types.ManagedObjectReference `xml:"cr" json:"cr"`
	SolutionsToModify []string                     `xml:"solutionsToModify,omitempty" json:"solutionsToModify,omitempty"`
	SolutionsToRemove []string                     `xml:"solutionsToRemove,omitempty" json:"solutionsToRemove,omitempty"`
}

func init() {
	types.Add("eam:PersonalityAgencyCannotConfigureSolutions", reflect.TypeOf((*PersonalityAgencyCannotConfigureSolutions)(nil)).Elem())
}

type PersonalityAgencyCannotUploadDepot struct {
	PersonalityAgencyDepotIssue

	LocalDepotUrl string `xml:"localDepotUrl" json:"localDepotUrl"`
}

func init() {
	types.Add("eam:PersonalityAgencyCannotUploadDepot", reflect.TypeOf((*PersonalityAgencyCannotUploadDepot)(nil)).Elem())
}

type PersonalityAgencyDepotIssue struct {
	PersonalityAgencyPMIssue

	RemoteDepotUrl string `xml:"remoteDepotUrl" json:"remoteDepotUrl"`
}

func init() {
	types.Add("eam:PersonalityAgencyDepotIssue", reflect.TypeOf((*PersonalityAgencyDepotIssue)(nil)).Elem())
}

type PersonalityAgencyInaccessibleDepot struct {
	PersonalityAgencyDepotIssue
}

func init() {
	types.Add("eam:PersonalityAgencyInaccessibleDepot", reflect.TypeOf((*PersonalityAgencyInaccessibleDepot)(nil)).Elem())
}

type PersonalityAgencyInvalidDepot struct {
	PersonalityAgencyDepotIssue
}

func init() {
	types.Add("eam:PersonalityAgencyInvalidDepot", reflect.TypeOf((*PersonalityAgencyInvalidDepot)(nil)).Elem())
}

type PersonalityAgencyPMIssue struct {
	AgencyIssue
}

func init() {
	types.Add("eam:PersonalityAgencyPMIssue", reflect.TypeOf((*PersonalityAgencyPMIssue)(nil)).Elem())
}

type PersonalityAgencyPMUnavailable struct {
	PersonalityAgencyPMIssue
}

func init() {
	types.Add("eam:PersonalityAgencyPMUnavailable", reflect.TypeOf((*PersonalityAgencyPMUnavailable)(nil)).Elem())
}

type PersonalityAgentAwaitingPMRemediation struct {
	PersonalityAgentPMIssue
}

func init() {
	types.Add("eam:PersonalityAgentAwaitingPMRemediation", reflect.TypeOf((*PersonalityAgentAwaitingPMRemediation)(nil)).Elem())
}

type PersonalityAgentBlockedByAgencyOperation struct {
	PersonalityAgentPMIssue
}

func init() {
	types.Add("eam:PersonalityAgentBlockedByAgencyOperation", reflect.TypeOf((*PersonalityAgentBlockedByAgencyOperation)(nil)).Elem())
}

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

type QueryIssueRequestType struct {
	This     types.ManagedObjectReference `xml:"_this" json:"_this"`
	IssueKey []int32                      `xml:"issueKey,omitempty" json:"issueKey,omitempty"`
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

type RegisterAgentVmRequestType struct {
	This    types.ManagedObjectReference `xml:"_this" json:"_this"`
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

type ResolveRequestType struct {
	This     types.ManagedObjectReference `xml:"_this" json:"_this"`
	IssueKey []int32                      `xml:"issueKey" json:"issueKey"`
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

type SetMaintenanceModePolicyRequestType struct {
	This   types.ManagedObjectReference `xml:"_this" json:"_this"`
	Policy string                       `xml:"policy" json:"policy"`
}

func init() {
	types.Add("eam:SetMaintenanceModePolicyRequestType", reflect.TypeOf((*SetMaintenanceModePolicyRequestType)(nil)).Elem())
}

type SetMaintenanceModePolicyResponse struct {
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

type UnknownAgentVm struct {
	HostIssue

	Vm types.ManagedObjectReference `xml:"vm" json:"vm"`
}

func init() {
	types.Add("eam:UnknownAgentVm", reflect.TypeOf((*UnknownAgentVm)(nil)).Elem())
}

type UnregisterAgentVm UnregisterAgentVmRequestType

func init() {
	types.Add("eam:UnregisterAgentVm", reflect.TypeOf((*UnregisterAgentVm)(nil)).Elem())
}

type UnregisterAgentVmRequestType struct {
	This    types.ManagedObjectReference `xml:"_this" json:"_this"`
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

type UpdateRequestType struct {
	This   types.ManagedObjectReference `xml:"_this" json:"_this"`
	Config BaseAgencyConfigInfo         `xml:"config,typeattr" json:"config"`
}

func init() {
	types.Add("eam:UpdateRequestType", reflect.TypeOf((*UpdateRequestType)(nil)).Elem())
}

type UpdateResponse struct {
}

type VibCannotPutHostInMaintenanceMode struct {
	VibIssue
}

func init() {
	types.Add("eam:VibCannotPutHostInMaintenanceMode", reflect.TypeOf((*VibCannotPutHostInMaintenanceMode)(nil)).Elem())
}

type VibCannotPutHostOutOfMaintenanceMode struct {
	VibIssue
}

func init() {
	types.Add("eam:VibCannotPutHostOutOfMaintenanceMode", reflect.TypeOf((*VibCannotPutHostOutOfMaintenanceMode)(nil)).Elem())
}

type VibDependenciesNotMetByHost struct {
	VibNotInstalled
}

func init() {
	types.Add("eam:VibDependenciesNotMetByHost", reflect.TypeOf((*VibDependenciesNotMetByHost)(nil)).Elem())
}

type VibInvalidFormat struct {
	VibNotInstalled
}

func init() {
	types.Add("eam:VibInvalidFormat", reflect.TypeOf((*VibInvalidFormat)(nil)).Elem())
}

type VibIssue struct {
	AgentIssue
}

func init() {
	types.Add("eam:VibIssue", reflect.TypeOf((*VibIssue)(nil)).Elem())
}

type VibNotInstalled struct {
	VibIssue
}

func init() {
	types.Add("eam:VibNotInstalled", reflect.TypeOf((*VibNotInstalled)(nil)).Elem())
}

type VibRequirementsNotMetByHost struct {
	VibNotInstalled
}

func init() {
	types.Add("eam:VibRequirementsNotMetByHost", reflect.TypeOf((*VibRequirementsNotMetByHost)(nil)).Elem())
}

type VibRequiresHostInMaintenanceMode struct {
	VibIssue
}

func init() {
	types.Add("eam:VibRequiresHostInMaintenanceMode", reflect.TypeOf((*VibRequiresHostInMaintenanceMode)(nil)).Elem())
}

type VibRequiresHostReboot struct {
	VibIssue
}

func init() {
	types.Add("eam:VibRequiresHostReboot", reflect.TypeOf((*VibRequiresHostReboot)(nil)).Elem())
}

type VibRequiresManualInstallation struct {
	VibIssue

	Bulletin []string `xml:"bulletin" json:"bulletin"`
}

func init() {
	types.Add("eam:VibRequiresManualInstallation", reflect.TypeOf((*VibRequiresManualInstallation)(nil)).Elem())
}

type VibRequiresManualUninstallation struct {
	VibIssue

	Bulletin []string `xml:"bulletin" json:"bulletin"`
}

func init() {
	types.Add("eam:VibRequiresManualUninstallation", reflect.TypeOf((*VibRequiresManualUninstallation)(nil)).Elem())
}

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

type VibVibInfoSoftwareTags struct {
	types.DynamicData

	Tags []string `xml:"tags,omitempty" json:"tags,omitempty"`
}

func init() {
	types.Add("eam:VibVibInfoSoftwareTags", reflect.TypeOf((*VibVibInfoSoftwareTags)(nil)).Elem())
}

type VmCorrupted struct {
	VmIssue

	MissingFile string `xml:"missingFile,omitempty" json:"missingFile,omitempty"`
}

func init() {
	types.Add("eam:VmCorrupted", reflect.TypeOf((*VmCorrupted)(nil)).Elem())
}

type VmDeployed struct {
	VmIssue
}

func init() {
	types.Add("eam:VmDeployed", reflect.TypeOf((*VmDeployed)(nil)).Elem())
}

type VmIssue struct {
	AgentIssue

	Vm types.ManagedObjectReference `xml:"vm" json:"vm"`
}

func init() {
	types.Add("eam:VmIssue", reflect.TypeOf((*VmIssue)(nil)).Elem())
}

type VmMarkedAsTemplate struct {
	VmIssue
}

func init() {
	types.Add("eam:VmMarkedAsTemplate", reflect.TypeOf((*VmMarkedAsTemplate)(nil)).Elem())
}

type VmNotDeployed struct {
	AgentIssue
}

func init() {
	types.Add("eam:VmNotDeployed", reflect.TypeOf((*VmNotDeployed)(nil)).Elem())
}

type VmOrphaned struct {
	VmIssue
}

func init() {
	types.Add("eam:VmOrphaned", reflect.TypeOf((*VmOrphaned)(nil)).Elem())
}

type VmPoweredOff struct {
	VmIssue
}

func init() {
	types.Add("eam:VmPoweredOff", reflect.TypeOf((*VmPoweredOff)(nil)).Elem())
}

type VmPoweredOn struct {
	VmIssue
}

func init() {
	types.Add("eam:VmPoweredOn", reflect.TypeOf((*VmPoweredOn)(nil)).Elem())
}

type VmRequiresHostOutOfMaintenanceMode struct {
	VmNotDeployed
}

func init() {
	types.Add("eam:VmRequiresHostOutOfMaintenanceMode", reflect.TypeOf((*VmRequiresHostOutOfMaintenanceMode)(nil)).Elem())
}

type VmSuspended struct {
	VmIssue
}

func init() {
	types.Add("eam:VmSuspended", reflect.TypeOf((*VmSuspended)(nil)).Elem())
}

type VmWrongFolder struct {
	VmIssue

	CurrentFolder  types.ManagedObjectReference `xml:"currentFolder" json:"currentFolder"`
	RequiredFolder types.ManagedObjectReference `xml:"requiredFolder" json:"requiredFolder"`
}

func init() {
	types.Add("eam:VmWrongFolder", reflect.TypeOf((*VmWrongFolder)(nil)).Elem())
}

type VmWrongResourcePool struct {
	VmIssue

	CurrentResourcePool  types.ManagedObjectReference `xml:"currentResourcePool" json:"currentResourcePool"`
	RequiredResourcePool types.ManagedObjectReference `xml:"requiredResourcePool" json:"requiredResourcePool"`
}

func init() {
	types.Add("eam:VmWrongResourcePool", reflect.TypeOf((*VmWrongResourcePool)(nil)).Elem())
}

type VersionURI string

func init() {
	types.Add("eam:versionURI", reflect.TypeOf((*VersionURI)(nil)).Elem())
}
