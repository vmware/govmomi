/*
Copyright (c) 2014-2020 VMware, Inc. All Rights Reserved.

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

type (
	ArrayOfHostVsanInternalSystemCmmdsQuery                        types.ArrayOfHostVsanInternalSystemCmmdsQuery
	ArrayOfHostVsanInternalSystemDeleteVsanObjectsResult           types.ArrayOfHostVsanInternalSystemDeleteVsanObjectsResult
	ArrayOfHostVsanInternalSystemVsanObjectOperationResult         types.ArrayOfHostVsanInternalSystemVsanObjectOperationResult
	ArrayOfHostVsanInternalSystemVsanPhysicalDiskDiagnosticsResult types.ArrayOfHostVsanInternalSystemVsanPhysicalDiskDiagnosticsResult
	ArrayOfVsanHostConfigInfo                                      types.ArrayOfVsanHostConfigInfo
	ArrayOfVsanHostConfigInfoNetworkInfoPortConfig                 types.ArrayOfVsanHostConfigInfoNetworkInfoPortConfig
	ArrayOfVsanHostDiskMapInfo                                     types.ArrayOfVsanHostDiskMapInfo
	ArrayOfVsanHostDiskMapResult                                   types.ArrayOfVsanHostDiskMapResult
	ArrayOfVsanHostDiskMapping                                     types.ArrayOfVsanHostDiskMapping
	ArrayOfVsanHostDiskResult                                      types.ArrayOfVsanHostDiskResult
	ArrayOfVsanHostMembershipInfo                                  types.ArrayOfVsanHostMembershipInfo
	ArrayOfVsanHostRuntimeInfoDiskIssue                            types.ArrayOfVsanHostRuntimeInfoDiskIssue
	ArrayOfVsanNewPolicyBatch                                      types.ArrayOfVsanNewPolicyBatch
	ArrayOfVsanPolicyChangeBatch                                   types.ArrayOfVsanPolicyChangeBatch
	ArrayOfVsanPolicySatisfiability                                types.ArrayOfVsanPolicySatisfiability
	ArrayOfVsanUpgradeSystemNetworkPartitionInfo                   types.ArrayOfVsanUpgradeSystemNetworkPartitionInfo
	ArrayOfVsanUpgradeSystemPreflightCheckIssue                    types.ArrayOfVsanUpgradeSystemPreflightCheckIssue
	ArrayOfVsanUpgradeSystemUpgradeHistoryItem                     types.ArrayOfVsanUpgradeSystemUpgradeHistoryItem
	CannotChangeVsanClusterUuid                                    types.CannotChangeVsanClusterUuid
	CannotChangeVsanClusterUuidFault                               types.CannotChangeVsanClusterUuidFault
	CannotChangeVsanNodeUuid                                       types.CannotChangeVsanNodeUuid
	CannotChangeVsanNodeUuidFault                                  types.CannotChangeVsanNodeUuidFault
	CannotMoveVsanEnabledHost                                      types.CannotMoveVsanEnabledHost
	CannotMoveVsanEnabledHostFault                                 types.CannotMoveVsanEnabledHostFault
	CannotReconfigureVsanWhenHaEnabled                             types.CannotReconfigureVsanWhenHaEnabled
	CannotReconfigureVsanWhenHaEnabledFault                        types.CannotReconfigureVsanWhenHaEnabledFault
	DeleteVsanObjects                                              types.DeleteVsanObjects
	DeleteVsanObjectsRequestType                                   types.DeleteVsanObjectsRequestType
	DeleteVsanObjectsResponse                                      types.DeleteVsanObjectsResponse
	DestinationVsanDisabled                                        types.DestinationVsanDisabled
	DestinationVsanDisabledFault                                   types.DestinationVsanDisabledFault
	DuplicateVsanNetworkInterface                                  types.DuplicateVsanNetworkInterface
	DuplicateVsanNetworkInterfaceFault                             types.DuplicateVsanNetworkInterfaceFault
	EvacuateVsanNodeRequestType                                    types.EvacuateVsanNodeRequestType
	EvacuateVsanNode_Task                                          types.EvacuateVsanNode_Task
	EvacuateVsanNode_TaskResponse                                  types.EvacuateVsanNode_TaskResponse
	GetVsanObjExtAttrs                                             types.GetVsanObjExtAttrs
	GetVsanObjExtAttrsRequestType                                  types.GetVsanObjExtAttrsRequestType
	GetVsanObjExtAttrsResponse                                     types.GetVsanObjExtAttrsResponse
	HostVsanInternalSystemCmmdsQuery                               types.HostVsanInternalSystemCmmdsQuery
	HostVsanInternalSystemDeleteVsanObjectsResult                  types.HostVsanInternalSystemDeleteVsanObjectsResult
	HostVsanInternalSystemVsanObjectOperationResult                types.HostVsanInternalSystemVsanObjectOperationResult
	HostVsanInternalSystemVsanPhysicalDiskDiagnosticsResult        types.HostVsanInternalSystemVsanPhysicalDiskDiagnosticsResult
	NotSupportedHostForVsan                                        types.NotSupportedHostForVsan
	NotSupportedHostForVsanFault                                   types.NotSupportedHostForVsanFault
	PerformVsanUpgradePreflightCheck                               types.PerformVsanUpgradePreflightCheck
	PerformVsanUpgradePreflightCheckRequestType                    types.PerformVsanUpgradePreflightCheckRequestType
	PerformVsanUpgradePreflightCheckResponse                       types.PerformVsanUpgradePreflightCheckResponse
	PerformVsanUpgradeRequestType                                  types.PerformVsanUpgradeRequestType
	PerformVsanUpgrade_Task                                        types.PerformVsanUpgrade_Task
	PerformVsanUpgrade_TaskResponse                                types.PerformVsanUpgrade_TaskResponse
	QueryDisksForVsan                                              types.QueryDisksForVsan
	QueryDisksForVsanRequestType                                   types.QueryDisksForVsanRequestType
	QueryDisksForVsanResponse                                      types.QueryDisksForVsanResponse
	QueryObjectsOnPhysicalVsanDisk                                 types.QueryObjectsOnPhysicalVsanDisk
	QueryObjectsOnPhysicalVsanDiskRequestType                      types.QueryObjectsOnPhysicalVsanDiskRequestType
	QueryObjectsOnPhysicalVsanDiskResponse                         types.QueryObjectsOnPhysicalVsanDiskResponse
	QueryPhysicalVsanDisks                                         types.QueryPhysicalVsanDisks
	QueryPhysicalVsanDisksRequestType                              types.QueryPhysicalVsanDisksRequestType
	QueryPhysicalVsanDisksResponse                                 types.QueryPhysicalVsanDisksResponse
	QuerySyncingVsanObjects                                        types.QuerySyncingVsanObjects
	QuerySyncingVsanObjectsRequestType                             types.QuerySyncingVsanObjectsRequestType
	QuerySyncingVsanObjectsResponse                                types.QuerySyncingVsanObjectsResponse
	QueryVsanObjectUuidsByFilter                                   types.QueryVsanObjectUuidsByFilter
	QueryVsanObjectUuidsByFilterRequestType                        types.QueryVsanObjectUuidsByFilterRequestType
	QueryVsanObjectUuidsByFilterResponse                           types.QueryVsanObjectUuidsByFilterResponse
	QueryVsanObjects                                               types.QueryVsanObjects
	QueryVsanObjectsRequestType                                    types.QueryVsanObjectsRequestType
	QueryVsanObjectsResponse                                       types.QueryVsanObjectsResponse
	QueryVsanStatistics                                            types.QueryVsanStatistics
	QueryVsanStatisticsRequestType                                 types.QueryVsanStatisticsRequestType
	QueryVsanStatisticsResponse                                    types.QueryVsanStatisticsResponse
	QueryVsanUpgradeStatus                                         types.QueryVsanUpgradeStatus
	QueryVsanUpgradeStatusRequestType                              types.QueryVsanUpgradeStatusRequestType
	QueryVsanUpgradeStatusResponse                                 types.QueryVsanUpgradeStatusResponse
	RecommissionVsanNodeRequestType                                types.RecommissionVsanNodeRequestType
	RecommissionVsanNode_Task                                      types.RecommissionVsanNode_Task
	RecommissionVsanNode_TaskResponse                              types.RecommissionVsanNode_TaskResponse
	RunVsanPhysicalDiskDiagnostics                                 types.RunVsanPhysicalDiskDiagnostics
	RunVsanPhysicalDiskDiagnosticsRequestType                      types.RunVsanPhysicalDiskDiagnosticsRequestType
	RunVsanPhysicalDiskDiagnosticsResponse                         types.RunVsanPhysicalDiskDiagnosticsResponse
	UpdateVsanRequestType                                          types.UpdateVsanRequestType
	UpdateVsan_Task                                                types.UpdateVsan_Task
	UpdateVsan_TaskResponse                                        types.UpdateVsan_TaskResponse
	UpgradeVsanObjects                                             types.UpgradeVsanObjects
	UpgradeVsanObjectsRequestType                                  types.UpgradeVsanObjectsRequestType
	UpgradeVsanObjectsResponse                                     types.UpgradeVsanObjectsResponse
	VsanClusterConfigInfo                                          types.VsanClusterConfigInfo
	VsanClusterConfigInfoHostDefaultInfo                           types.VsanClusterConfigInfoHostDefaultInfo
	VsanClusterUuidMismatch                                        types.VsanClusterUuidMismatch
	VsanClusterUuidMismatchFault                                   types.VsanClusterUuidMismatchFault
	VsanDiskFault                                                  types.VsanDiskFault
	VsanDiskFaultFault                                             types.VsanDiskFaultFault
	VsanDiskIssueType                                              types.VsanDiskIssueType
	VsanFault                                                      types.VsanFault
	VsanFaultFault                                                 types.VsanFaultFault
	VsanHostClusterStatus                                          types.VsanHostClusterStatus
	VsanHostClusterStatusState                                     types.VsanHostClusterStatusState
	VsanHostClusterStatusStateCompletionEstimate                   types.VsanHostClusterStatusStateCompletionEstimate
	VsanHostConfigInfo                                             types.VsanHostConfigInfo
	VsanHostConfigInfoClusterInfo                                  types.VsanHostConfigInfoClusterInfo
	VsanHostConfigInfoNetworkInfo                                  types.VsanHostConfigInfoNetworkInfo
	VsanHostConfigInfoNetworkInfoPortConfig                        types.VsanHostConfigInfoNetworkInfoPortConfig
	VsanHostConfigInfoStorageInfo                                  types.VsanHostConfigInfoStorageInfo
	VsanHostDecommissionMode                                       types.VsanHostDecommissionMode
	VsanHostDecommissionModeObjectAction                           types.VsanHostDecommissionModeObjectAction
	VsanHostDiskMapInfo                                            types.VsanHostDiskMapInfo
	VsanHostDiskMapResult                                          types.VsanHostDiskMapResult
	VsanHostDiskMapping                                            types.VsanHostDiskMapping
	VsanHostDiskResult                                             types.VsanHostDiskResult
	VsanHostDiskResultState                                        types.VsanHostDiskResultState
	VsanHostFaultDomainInfo                                        types.VsanHostFaultDomainInfo
	VsanHostHealthState                                            types.VsanHostHealthState
	VsanHostIpConfig                                               types.VsanHostIpConfig
	VsanHostMembershipInfo                                         types.VsanHostMembershipInfo
	VsanHostNodeState                                              types.VsanHostNodeState
	VsanHostRuntimeInfo                                            types.VsanHostRuntimeInfo
	VsanHostRuntimeInfoDiskIssue                                   types.VsanHostRuntimeInfoDiskIssue
	VsanHostVsanDiskInfo                                           types.VsanHostVsanDiskInfo
	VsanIncompatibleDiskMapping                                    types.VsanIncompatibleDiskMapping
	VsanIncompatibleDiskMappingFault                               types.VsanIncompatibleDiskMappingFault
	VsanNewPolicyBatch                                             types.VsanNewPolicyBatch
	VsanPolicyChangeBatch                                          types.VsanPolicyChangeBatch
	VsanPolicyCost                                                 types.VsanPolicyCost
	VsanPolicySatisfiability                                       types.VsanPolicySatisfiability
	VsanUpgradeSystemAPIBrokenIssue                                types.VsanUpgradeSystemAPIBrokenIssue
	VsanUpgradeSystemAutoClaimEnabledOnHostsIssue                  types.VsanUpgradeSystemAutoClaimEnabledOnHostsIssue
	VsanUpgradeSystemHostsDisconnectedIssue                        types.VsanUpgradeSystemHostsDisconnectedIssue
	VsanUpgradeSystemMissingHostsInClusterIssue                    types.VsanUpgradeSystemMissingHostsInClusterIssue
	VsanUpgradeSystemNetworkPartitionInfo                          types.VsanUpgradeSystemNetworkPartitionInfo
	VsanUpgradeSystemNetworkPartitionIssue                         types.VsanUpgradeSystemNetworkPartitionIssue
	VsanUpgradeSystemNotEnoughFreeCapacityIssue                    types.VsanUpgradeSystemNotEnoughFreeCapacityIssue
	VsanUpgradeSystemPreflightCheckIssue                           types.VsanUpgradeSystemPreflightCheckIssue
	VsanUpgradeSystemPreflightCheckResult                          types.VsanUpgradeSystemPreflightCheckResult
	VsanUpgradeSystemRogueHostsInClusterIssue                      types.VsanUpgradeSystemRogueHostsInClusterIssue
	VsanUpgradeSystemUpgradeHistoryDiskGroupOp                     types.VsanUpgradeSystemUpgradeHistoryDiskGroupOp
	VsanUpgradeSystemUpgradeHistoryDiskGroupOpType                 types.VsanUpgradeSystemUpgradeHistoryDiskGroupOpType
	VsanUpgradeSystemUpgradeHistoryItem                            types.VsanUpgradeSystemUpgradeHistoryItem
	VsanUpgradeSystemUpgradeHistoryPreflightFail                   types.VsanUpgradeSystemUpgradeHistoryPreflightFail
	VsanUpgradeSystemUpgradeStatus                                 types.VsanUpgradeSystemUpgradeStatus
	VsanUpgradeSystemV2ObjectsPresentDuringDowngradeIssue          types.VsanUpgradeSystemV2ObjectsPresentDuringDowngradeIssue
	VsanUpgradeSystemWrongEsxVersionIssue                          types.VsanUpgradeSystemWrongEsxVersionIssue
)

type ArrayOfCnsContainerCluster struct {
	CnsContainerCluster []CnsContainerCluster `xml:"CnsContainerCluster,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfCnsContainerCluster", reflect.TypeOf((*ArrayOfCnsContainerCluster)(nil)).Elem())
}

type ArrayOfCnsEntityMetadata struct {
	CnsEntityMetadata []BaseCnsEntityMetadata `xml:"CnsEntityMetadata,omitempty,typeattr"`
}

func init() {
	types.Add("vsan:ArrayOfCnsEntityMetadata", reflect.TypeOf((*ArrayOfCnsEntityMetadata)(nil)).Elem())
}

type ArrayOfCnsKubernetesEntityReference struct {
	CnsKubernetesEntityReference []CnsKubernetesEntityReference `xml:"CnsKubernetesEntityReference,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfCnsKubernetesEntityReference", reflect.TypeOf((*ArrayOfCnsKubernetesEntityReference)(nil)).Elem())
}

type ArrayOfCnsVolume struct {
	CnsVolume []CnsVolume `xml:"CnsVolume,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfCnsVolume", reflect.TypeOf((*ArrayOfCnsVolume)(nil)).Elem())
}

type ArrayOfCnsVolumeAttachDetachSpec struct {
	CnsVolumeAttachDetachSpec []CnsVolumeAttachDetachSpec `xml:"CnsVolumeAttachDetachSpec,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfCnsVolumeAttachDetachSpec", reflect.TypeOf((*ArrayOfCnsVolumeAttachDetachSpec)(nil)).Elem())
}

type ArrayOfCnsVolumeCreateSpec struct {
	CnsVolumeCreateSpec []CnsVolumeCreateSpec `xml:"CnsVolumeCreateSpec,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfCnsVolumeCreateSpec", reflect.TypeOf((*ArrayOfCnsVolumeCreateSpec)(nil)).Elem())
}

type ArrayOfCnsVolumeId struct {
	CnsVolumeId []CnsVolumeId `xml:"CnsVolumeId,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfCnsVolumeId", reflect.TypeOf((*ArrayOfCnsVolumeId)(nil)).Elem())
}

type ArrayOfCnsVolumeMetadataUpdateSpec struct {
	CnsVolumeMetadataUpdateSpec []CnsVolumeMetadataUpdateSpec `xml:"CnsVolumeMetadataUpdateSpec,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfCnsVolumeMetadataUpdateSpec", reflect.TypeOf((*ArrayOfCnsVolumeMetadataUpdateSpec)(nil)).Elem())
}

type ArrayOfCnsVolumeOperationResult struct {
	CnsVolumeOperationResult []BaseCnsVolumeOperationResult `xml:"CnsVolumeOperationResult,omitempty,typeattr"`
}

func init() {
	types.Add("vsan:ArrayOfCnsVolumeOperationResult", reflect.TypeOf((*ArrayOfCnsVolumeOperationResult)(nil)).Elem())
}

type ArrayOfDynamicData struct {
	DynamicData []types.BaseDynamicData `xml:"DynamicData,omitempty,typeattr"`
}

func init() {
	types.Add("vsan:ArrayOfDynamicData", reflect.TypeOf((*ArrayOfDynamicData)(nil)).Elem())
}

type ArrayOfKmipServerSpec struct {
	KmipServerSpec []types.KmipServerSpec `xml:"KmipServerSpec,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfKmipServerSpec", reflect.TypeOf((*ArrayOfKmipServerSpec)(nil)).Elem())
}

type ArrayOfVimClusterVSANStretchedClusterCapability struct {
	VimClusterVSANStretchedClusterCapability []VimClusterVSANStretchedClusterCapability `xml:"VimClusterVSANStretchedClusterCapability,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVimClusterVSANStretchedClusterCapability", reflect.TypeOf((*ArrayOfVimClusterVSANStretchedClusterCapability)(nil)).Elem())
}

type ArrayOfVimClusterVSANWitnessHostInfo struct {
	VimClusterVSANWitnessHostInfo []VimClusterVSANWitnessHostInfo `xml:"VimClusterVSANWitnessHostInfo,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVimClusterVSANWitnessHostInfo", reflect.TypeOf((*ArrayOfVimClusterVSANWitnessHostInfo)(nil)).Elem())
}

type ArrayOfVimClusterVsanFaultDomainSpec struct {
	VimClusterVsanFaultDomainSpec []VimClusterVsanFaultDomainSpec `xml:"VimClusterVsanFaultDomainSpec,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVimClusterVsanFaultDomainSpec", reflect.TypeOf((*ArrayOfVimClusterVsanFaultDomainSpec)(nil)).Elem())
}

type ArrayOfVimClusterVsanHostDiskMapping struct {
	VimClusterVsanHostDiskMapping []VimClusterVsanHostDiskMapping `xml:"VimClusterVsanHostDiskMapping,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVimClusterVsanHostDiskMapping", reflect.TypeOf((*ArrayOfVimClusterVsanHostDiskMapping)(nil)).Elem())
}

type ArrayOfVimVsanHostDiskMapInfoEx struct {
	VimVsanHostDiskMapInfoEx []VimVsanHostDiskMapInfoEx `xml:"VimVsanHostDiskMapInfoEx,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVimVsanHostDiskMapInfoEx", reflect.TypeOf((*ArrayOfVimVsanHostDiskMapInfoEx)(nil)).Elem())
}

type ArrayOfVimVsanHostVsanHostCapability struct {
	VimVsanHostVsanHostCapability []VimVsanHostVsanHostCapability `xml:"VimVsanHostVsanHostCapability,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVimVsanHostVsanHostCapability", reflect.TypeOf((*ArrayOfVimVsanHostVsanHostCapability)(nil)).Elem())
}

type ArrayOfVsanAttachToSrOperation struct {
	VsanAttachToSrOperation []VsanAttachToSrOperation `xml:"VsanAttachToSrOperation,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanAttachToSrOperation", reflect.TypeOf((*ArrayOfVsanAttachToSrOperation)(nil)).Elem())
}

type ArrayOfVsanBasicDeviceInfo struct {
	VsanBasicDeviceInfo []VsanBasicDeviceInfo `xml:"VsanBasicDeviceInfo,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanBasicDeviceInfo", reflect.TypeOf((*ArrayOfVsanBasicDeviceInfo)(nil)).Elem())
}

type ArrayOfVsanBurnInTest struct {
	VsanBurnInTest []VsanBurnInTest `xml:"VsanBurnInTest,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanBurnInTest", reflect.TypeOf((*ArrayOfVsanBurnInTest)(nil)).Elem())
}

type ArrayOfVsanCapability struct {
	VsanCapability []VsanCapability `xml:"VsanCapability,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanCapability", reflect.TypeOf((*ArrayOfVsanCapability)(nil)).Elem())
}

type ArrayOfVsanClusterAdvCfgSyncHostResult struct {
	VsanClusterAdvCfgSyncHostResult []VsanClusterAdvCfgSyncHostResult `xml:"VsanClusterAdvCfgSyncHostResult,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanClusterAdvCfgSyncHostResult", reflect.TypeOf((*ArrayOfVsanClusterAdvCfgSyncHostResult)(nil)).Elem())
}

type ArrayOfVsanClusterAdvCfgSyncResult struct {
	VsanClusterAdvCfgSyncResult []VsanClusterAdvCfgSyncResult `xml:"VsanClusterAdvCfgSyncResult,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanClusterAdvCfgSyncResult", reflect.TypeOf((*ArrayOfVsanClusterAdvCfgSyncResult)(nil)).Elem())
}

type ArrayOfVsanClusterBalancePerDiskInfo struct {
	VsanClusterBalancePerDiskInfo []VsanClusterBalancePerDiskInfo `xml:"VsanClusterBalancePerDiskInfo,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanClusterBalancePerDiskInfo", reflect.TypeOf((*ArrayOfVsanClusterBalancePerDiskInfo)(nil)).Elem())
}

type ArrayOfVsanClusterCreateVmHealthTestResult struct {
	VsanClusterCreateVmHealthTestResult []VsanClusterCreateVmHealthTestResult `xml:"VsanClusterCreateVmHealthTestResult,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanClusterCreateVmHealthTestResult", reflect.TypeOf((*ArrayOfVsanClusterCreateVmHealthTestResult)(nil)).Elem())
}

type ArrayOfVsanClusterHealthAction struct {
	VsanClusterHealthAction []VsanClusterHealthAction `xml:"VsanClusterHealthAction,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanClusterHealthAction", reflect.TypeOf((*ArrayOfVsanClusterHealthAction)(nil)).Elem())
}

type ArrayOfVsanClusterHealthCheckInfo struct {
	VsanClusterHealthCheckInfo []VsanClusterHealthCheckInfo `xml:"VsanClusterHealthCheckInfo,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanClusterHealthCheckInfo", reflect.TypeOf((*ArrayOfVsanClusterHealthCheckInfo)(nil)).Elem())
}

type ArrayOfVsanClusterHealthGroup struct {
	VsanClusterHealthGroup []VsanClusterHealthGroup `xml:"VsanClusterHealthGroup,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanClusterHealthGroup", reflect.TypeOf((*ArrayOfVsanClusterHealthGroup)(nil)).Elem())
}

type ArrayOfVsanClusterHealthResultBase struct {
	VsanClusterHealthResultBase []BaseVsanClusterHealthResultBase `xml:"VsanClusterHealthResultBase,omitempty,typeattr"`
}

func init() {
	types.Add("vsan:ArrayOfVsanClusterHealthResultBase", reflect.TypeOf((*ArrayOfVsanClusterHealthResultBase)(nil)).Elem())
}

type ArrayOfVsanClusterHealthResultColumnInfo struct {
	VsanClusterHealthResultColumnInfo []VsanClusterHealthResultColumnInfo `xml:"VsanClusterHealthResultColumnInfo,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanClusterHealthResultColumnInfo", reflect.TypeOf((*ArrayOfVsanClusterHealthResultColumnInfo)(nil)).Elem())
}

type ArrayOfVsanClusterHealthResultKeyValuePair struct {
	VsanClusterHealthResultKeyValuePair []VsanClusterHealthResultKeyValuePair `xml:"VsanClusterHealthResultKeyValuePair,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanClusterHealthResultKeyValuePair", reflect.TypeOf((*ArrayOfVsanClusterHealthResultKeyValuePair)(nil)).Elem())
}

type ArrayOfVsanClusterHealthResultRow struct {
	VsanClusterHealthResultRow []VsanClusterHealthResultRow `xml:"VsanClusterHealthResultRow,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanClusterHealthResultRow", reflect.TypeOf((*ArrayOfVsanClusterHealthResultRow)(nil)).Elem())
}

type ArrayOfVsanClusterHealthTest struct {
	VsanClusterHealthTest []VsanClusterHealthTest `xml:"VsanClusterHealthTest,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanClusterHealthTest", reflect.TypeOf((*ArrayOfVsanClusterHealthTest)(nil)).Elem())
}

type ArrayOfVsanClusterHostVmknicMapping struct {
	VsanClusterHostVmknicMapping []VsanClusterHostVmknicMapping `xml:"VsanClusterHostVmknicMapping,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanClusterHostVmknicMapping", reflect.TypeOf((*ArrayOfVsanClusterHostVmknicMapping)(nil)).Elem())
}

type ArrayOfVsanClusterNetworkLoadTestResult struct {
	VsanClusterNetworkLoadTestResult []VsanClusterNetworkLoadTestResult `xml:"VsanClusterNetworkLoadTestResult,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanClusterNetworkLoadTestResult", reflect.TypeOf((*ArrayOfVsanClusterNetworkLoadTestResult)(nil)).Elem())
}

type ArrayOfVsanClusterNetworkPartitionInfo struct {
	VsanClusterNetworkPartitionInfo []VsanClusterNetworkPartitionInfo `xml:"VsanClusterNetworkPartitionInfo,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanClusterNetworkPartitionInfo", reflect.TypeOf((*ArrayOfVsanClusterNetworkPartitionInfo)(nil)).Elem())
}

type ArrayOfVsanClusterObjectExtAttrs struct {
	VsanClusterObjectExtAttrs []VsanClusterObjectExtAttrs `xml:"VsanClusterObjectExtAttrs,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanClusterObjectExtAttrs", reflect.TypeOf((*ArrayOfVsanClusterObjectExtAttrs)(nil)).Elem())
}

type ArrayOfVsanClusterVMsHealthSummaryResult struct {
	VsanClusterVMsHealthSummaryResult []VsanClusterVMsHealthSummaryResult `xml:"VsanClusterVMsHealthSummaryResult,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanClusterVMsHealthSummaryResult", reflect.TypeOf((*ArrayOfVsanClusterVMsHealthSummaryResult)(nil)).Elem())
}

type ArrayOfVsanClusterVmdkLoadTestResult struct {
	VsanClusterVmdkLoadTestResult []VsanClusterVmdkLoadTestResult `xml:"VsanClusterVmdkLoadTestResult,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanClusterVmdkLoadTestResult", reflect.TypeOf((*ArrayOfVsanClusterVmdkLoadTestResult)(nil)).Elem())
}

type ArrayOfVsanClusterWhatifHostFailuresResult struct {
	VsanClusterWhatifHostFailuresResult []VsanClusterWhatifHostFailuresResult `xml:"VsanClusterWhatifHostFailuresResult,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanClusterWhatifHostFailuresResult", reflect.TypeOf((*ArrayOfVsanClusterWhatifHostFailuresResult)(nil)).Elem())
}

type ArrayOfVsanCompliantDriver struct {
	VsanCompliantDriver []VsanCompliantDriver `xml:"VsanCompliantDriver,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanCompliantDriver", reflect.TypeOf((*ArrayOfVsanCompliantDriver)(nil)).Elem())
}

type ArrayOfVsanCompliantFirmware struct {
	VsanCompliantFirmware []VsanCompliantFirmware `xml:"VsanCompliantFirmware,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanCompliantFirmware", reflect.TypeOf((*ArrayOfVsanCompliantFirmware)(nil)).Elem())
}

type ArrayOfVsanConfigBaseIssue struct {
	VsanConfigBaseIssue []BaseVsanConfigBaseIssue `xml:"VsanConfigBaseIssue,omitempty,typeattr"`
}

func init() {
	types.Add("vsan:ArrayOfVsanConfigBaseIssue", reflect.TypeOf((*ArrayOfVsanConfigBaseIssue)(nil)).Elem())
}

type ArrayOfVsanDatastoreSpec struct {
	VsanDatastoreSpec []VsanDatastoreSpec `xml:"VsanDatastoreSpec,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanDatastoreSpec", reflect.TypeOf((*ArrayOfVsanDatastoreSpec)(nil)).Elem())
}

type ArrayOfVsanDiskEncryptionHealth struct {
	VsanDiskEncryptionHealth []VsanDiskEncryptionHealth `xml:"VsanDiskEncryptionHealth,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanDiskEncryptionHealth", reflect.TypeOf((*ArrayOfVsanDiskEncryptionHealth)(nil)).Elem())
}

type ArrayOfVsanDiskGroupResourceCheckResult struct {
	VsanDiskGroupResourceCheckResult []VsanDiskGroupResourceCheckResult `xml:"VsanDiskGroupResourceCheckResult,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanDiskGroupResourceCheckResult", reflect.TypeOf((*ArrayOfVsanDiskGroupResourceCheckResult)(nil)).Elem())
}

type ArrayOfVsanDiskResourceCheckResult struct {
	VsanDiskResourceCheckResult []VsanDiskResourceCheckResult `xml:"VsanDiskResourceCheckResult,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanDiskResourceCheckResult", reflect.TypeOf((*ArrayOfVsanDiskResourceCheckResult)(nil)).Elem())
}

type ArrayOfVsanDownloadItem struct {
	VsanDownloadItem []VsanDownloadItem `xml:"VsanDownloadItem,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanDownloadItem", reflect.TypeOf((*ArrayOfVsanDownloadItem)(nil)).Elem())
}

type ArrayOfVsanEncryptionHealthSummary struct {
	VsanEncryptionHealthSummary []VsanEncryptionHealthSummary `xml:"VsanEncryptionHealthSummary,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanEncryptionHealthSummary", reflect.TypeOf((*ArrayOfVsanEncryptionHealthSummary)(nil)).Elem())
}

type ArrayOfVsanEntitySpaceUsage struct {
	VsanEntitySpaceUsage []VsanEntitySpaceUsage `xml:"VsanEntitySpaceUsage,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanEntitySpaceUsage", reflect.TypeOf((*ArrayOfVsanEntitySpaceUsage)(nil)).Elem())
}

type ArrayOfVsanFailedRepairObjectResult struct {
	VsanFailedRepairObjectResult []VsanFailedRepairObjectResult `xml:"VsanFailedRepairObjectResult,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanFailedRepairObjectResult", reflect.TypeOf((*ArrayOfVsanFailedRepairObjectResult)(nil)).Elem())
}

type ArrayOfVsanFaultDomainResourceCheckResult struct {
	VsanFaultDomainResourceCheckResult []VsanFaultDomainResourceCheckResult `xml:"VsanFaultDomainResourceCheckResult,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanFaultDomainResourceCheckResult", reflect.TypeOf((*ArrayOfVsanFaultDomainResourceCheckResult)(nil)).Elem())
}

type ArrayOfVsanFileServerHealthSummary struct {
	VsanFileServerHealthSummary []VsanFileServerHealthSummary `xml:"VsanFileServerHealthSummary,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanFileServerHealthSummary", reflect.TypeOf((*ArrayOfVsanFileServerHealthSummary)(nil)).Elem())
}

type ArrayOfVsanFileServiceDomain struct {
	VsanFileServiceDomain []VsanFileServiceDomain `xml:"VsanFileServiceDomain,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanFileServiceDomain", reflect.TypeOf((*ArrayOfVsanFileServiceDomain)(nil)).Elem())
}

type ArrayOfVsanFileServiceDomainConfig struct {
	VsanFileServiceDomainConfig []VsanFileServiceDomainConfig `xml:"VsanFileServiceDomainConfig,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanFileServiceDomainConfig", reflect.TypeOf((*ArrayOfVsanFileServiceDomainConfig)(nil)).Elem())
}

type ArrayOfVsanFileServiceHealthSummary struct {
	VsanFileServiceHealthSummary []VsanFileServiceHealthSummary `xml:"VsanFileServiceHealthSummary,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanFileServiceHealthSummary", reflect.TypeOf((*ArrayOfVsanFileServiceHealthSummary)(nil)).Elem())
}

type ArrayOfVsanFileServiceIpConfig struct {
	VsanFileServiceIpConfig []VsanFileServiceIpConfig `xml:"VsanFileServiceIpConfig,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanFileServiceIpConfig", reflect.TypeOf((*ArrayOfVsanFileServiceIpConfig)(nil)).Elem())
}

type ArrayOfVsanFileServiceOvfSpec struct {
	VsanFileServiceOvfSpec []VsanFileServiceOvfSpec `xml:"VsanFileServiceOvfSpec,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanFileServiceOvfSpec", reflect.TypeOf((*ArrayOfVsanFileServiceOvfSpec)(nil)).Elem())
}

type ArrayOfVsanFileServiceShareHealthSummary struct {
	VsanFileServiceShareHealthSummary []VsanFileServiceShareHealthSummary `xml:"VsanFileServiceShareHealthSummary,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanFileServiceShareHealthSummary", reflect.TypeOf((*ArrayOfVsanFileServiceShareHealthSummary)(nil)).Elem())
}

type ArrayOfVsanFileShare struct {
	VsanFileShare []VsanFileShare `xml:"VsanFileShare,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanFileShare", reflect.TypeOf((*ArrayOfVsanFileShare)(nil)).Elem())
}

type ArrayOfVsanFileShareNetPermission struct {
	VsanFileShareNetPermission []VsanFileShareNetPermission `xml:"VsanFileShareNetPermission,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanFileShareNetPermission", reflect.TypeOf((*ArrayOfVsanFileShareNetPermission)(nil)).Elem())
}

type ArrayOfVsanGenericClusterBaseIssue struct {
	VsanGenericClusterBaseIssue []VsanGenericClusterBaseIssue `xml:"VsanGenericClusterBaseIssue,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanGenericClusterBaseIssue", reflect.TypeOf((*ArrayOfVsanGenericClusterBaseIssue)(nil)).Elem())
}

type ArrayOfVsanHclControllerInfo struct {
	VsanHclControllerInfo []VsanHclControllerInfo `xml:"VsanHclControllerInfo,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanHclControllerInfo", reflect.TypeOf((*ArrayOfVsanHclControllerInfo)(nil)).Elem())
}

type ArrayOfVsanHclDeviceConstraint struct {
	VsanHclDeviceConstraint []VsanHclDeviceConstraint `xml:"VsanHclDeviceConstraint,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanHclDeviceConstraint", reflect.TypeOf((*ArrayOfVsanHclDeviceConstraint)(nil)).Elem())
}

type ArrayOfVsanHclDiskInfo struct {
	VsanHclDiskInfo []VsanHclDiskInfo `xml:"VsanHclDiskInfo,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanHclDiskInfo", reflect.TypeOf((*ArrayOfVsanHclDiskInfo)(nil)).Elem())
}

type ArrayOfVsanHclDriverInfo struct {
	VsanHclDriverInfo []VsanHclDriverInfo `xml:"VsanHclDriverInfo,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanHclDriverInfo", reflect.TypeOf((*ArrayOfVsanHclDriverInfo)(nil)).Elem())
}

type ArrayOfVsanHclFirmwareFile struct {
	VsanHclFirmwareFile []VsanHclFirmwareFile `xml:"VsanHclFirmwareFile,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanHclFirmwareFile", reflect.TypeOf((*ArrayOfVsanHclFirmwareFile)(nil)).Elem())
}

type ArrayOfVsanHclFirmwareUpdateSpec struct {
	VsanHclFirmwareUpdateSpec []VsanHclFirmwareUpdateSpec `xml:"VsanHclFirmwareUpdateSpec,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanHclFirmwareUpdateSpec", reflect.TypeOf((*ArrayOfVsanHclFirmwareUpdateSpec)(nil)).Elem())
}

type ArrayOfVsanHclNicInfo struct {
	VsanHclNicInfo []VsanHclNicInfo `xml:"VsanHclNicInfo,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanHclNicInfo", reflect.TypeOf((*ArrayOfVsanHclNicInfo)(nil)).Elem())
}

type ArrayOfVsanHclReleaseConstraint struct {
	VsanHclReleaseConstraint []VsanHclReleaseConstraint `xml:"VsanHclReleaseConstraint,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanHclReleaseConstraint", reflect.TypeOf((*ArrayOfVsanHclReleaseConstraint)(nil)).Elem())
}

type ArrayOfVsanHostAssociatedObjects struct {
	VsanHostAssociatedObjects []VsanHostAssociatedObjects `xml:"VsanHostAssociatedObjects,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanHostAssociatedObjects", reflect.TypeOf((*ArrayOfVsanHostAssociatedObjects)(nil)).Elem())
}

type ArrayOfVsanHostClomdLivenessResult struct {
	VsanHostClomdLivenessResult []VsanHostClomdLivenessResult `xml:"VsanHostClomdLivenessResult,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanHostClomdLivenessResult", reflect.TypeOf((*ArrayOfVsanHostClomdLivenessResult)(nil)).Elem())
}

type ArrayOfVsanHostComponentSyncState struct {
	VsanHostComponentSyncState []VsanHostComponentSyncState `xml:"VsanHostComponentSyncState,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanHostComponentSyncState", reflect.TypeOf((*ArrayOfVsanHostComponentSyncState)(nil)).Elem())
}

type ArrayOfVsanHostCreateVmHealthTestResult struct {
	VsanHostCreateVmHealthTestResult []VsanHostCreateVmHealthTestResult `xml:"VsanHostCreateVmHealthTestResult,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanHostCreateVmHealthTestResult", reflect.TypeOf((*ArrayOfVsanHostCreateVmHealthTestResult)(nil)).Elem())
}

type ArrayOfVsanHostDeviceInfo struct {
	VsanHostDeviceInfo []VsanHostDeviceInfo `xml:"VsanHostDeviceInfo,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanHostDeviceInfo", reflect.TypeOf((*ArrayOfVsanHostDeviceInfo)(nil)).Elem())
}

type ArrayOfVsanHostDrsStats struct {
	VsanHostDrsStats []VsanHostDrsStats `xml:"VsanHostDrsStats,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanHostDrsStats", reflect.TypeOf((*ArrayOfVsanHostDrsStats)(nil)).Elem())
}

type ArrayOfVsanHostFwComponent struct {
	VsanHostFwComponent []VsanHostFwComponent `xml:"VsanHostFwComponent,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanHostFwComponent", reflect.TypeOf((*ArrayOfVsanHostFwComponent)(nil)).Elem())
}

type ArrayOfVsanHostHclInfo struct {
	VsanHostHclInfo []VsanHostHclInfo `xml:"VsanHostHclInfo,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanHostHclInfo", reflect.TypeOf((*ArrayOfVsanHostHclInfo)(nil)).Elem())
}

type ArrayOfVsanHostHealthSystemStatusResult struct {
	VsanHostHealthSystemStatusResult []VsanHostHealthSystemStatusResult `xml:"VsanHostHealthSystemStatusResult,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanHostHealthSystemStatusResult", reflect.TypeOf((*ArrayOfVsanHostHealthSystemStatusResult)(nil)).Elem())
}

type ArrayOfVsanHostHealthSystemVersionResult struct {
	VsanHostHealthSystemVersionResult []VsanHostHealthSystemVersionResult `xml:"VsanHostHealthSystemVersionResult,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanHostHealthSystemVersionResult", reflect.TypeOf((*ArrayOfVsanHostHealthSystemVersionResult)(nil)).Elem())
}

type ArrayOfVsanHostResourceCheckResult struct {
	VsanHostResourceCheckResult []VsanHostResourceCheckResult `xml:"VsanHostResourceCheckResult,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanHostResourceCheckResult", reflect.TypeOf((*ArrayOfVsanHostResourceCheckResult)(nil)).Elem())
}

type ArrayOfVsanHostVirtualApplianceInfo struct {
	VsanHostVirtualApplianceInfo []VsanHostVirtualApplianceInfo `xml:"VsanHostVirtualApplianceInfo,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanHostVirtualApplianceInfo", reflect.TypeOf((*ArrayOfVsanHostVirtualApplianceInfo)(nil)).Elem())
}

type ArrayOfVsanHostVmdkLoadTestResult struct {
	VsanHostVmdkLoadTestResult []VsanHostVmdkLoadTestResult `xml:"VsanHostVmdkLoadTestResult,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanHostVmdkLoadTestResult", reflect.TypeOf((*ArrayOfVsanHostVmdkLoadTestResult)(nil)).Elem())
}

type ArrayOfVsanHostVsanObjectSyncState struct {
	VsanHostVsanObjectSyncState []VsanHostVsanObjectSyncState `xml:"VsanHostVsanObjectSyncState,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanHostVsanObjectSyncState", reflect.TypeOf((*ArrayOfVsanHostVsanObjectSyncState)(nil)).Elem())
}

type ArrayOfVsanIscsiInitiatorGroup struct {
	VsanIscsiInitiatorGroup []VsanIscsiInitiatorGroup `xml:"VsanIscsiInitiatorGroup,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanIscsiInitiatorGroup", reflect.TypeOf((*ArrayOfVsanIscsiInitiatorGroup)(nil)).Elem())
}

type ArrayOfVsanIscsiLUN struct {
	VsanIscsiLUN []VsanIscsiLUN `xml:"VsanIscsiLUN,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanIscsiLUN", reflect.TypeOf((*ArrayOfVsanIscsiLUN)(nil)).Elem())
}

type ArrayOfVsanIscsiTarget struct {
	VsanIscsiTarget []VsanIscsiTarget `xml:"VsanIscsiTarget,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanIscsiTarget", reflect.TypeOf((*ArrayOfVsanIscsiTarget)(nil)).Elem())
}

type ArrayOfVsanIscsiTargetBasicInfo struct {
	VsanIscsiTargetBasicInfo []BaseVsanIscsiTargetBasicInfo `xml:"VsanIscsiTargetBasicInfo,omitempty,typeattr"`
}

func init() {
	types.Add("vsan:ArrayOfVsanIscsiTargetBasicInfo", reflect.TypeOf((*ArrayOfVsanIscsiTargetBasicInfo)(nil)).Elem())
}

type ArrayOfVsanJsonComparator struct {
	VsanJsonComparator []VsanJsonComparator `xml:"VsanJsonComparator,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanJsonComparator", reflect.TypeOf((*ArrayOfVsanJsonComparator)(nil)).Elem())
}

type ArrayOfVsanKmsHealth struct {
	VsanKmsHealth []VsanKmsHealth `xml:"VsanKmsHealth,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanKmsHealth", reflect.TypeOf((*ArrayOfVsanKmsHealth)(nil)).Elem())
}

type ArrayOfVsanLimitHealthResult struct {
	VsanLimitHealthResult []VsanLimitHealthResult `xml:"VsanLimitHealthResult,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanLimitHealthResult", reflect.TypeOf((*ArrayOfVsanLimitHealthResult)(nil)).Elem())
}

type ArrayOfVsanMassCollectorPropertyParams struct {
	VsanMassCollectorPropertyParams []VsanMassCollectorPropertyParams `xml:"VsanMassCollectorPropertyParams,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanMassCollectorPropertyParams", reflect.TypeOf((*ArrayOfVsanMassCollectorPropertyParams)(nil)).Elem())
}

type ArrayOfVsanMassCollectorSpec struct {
	VsanMassCollectorSpec []VsanMassCollectorSpec `xml:"VsanMassCollectorSpec,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanMassCollectorSpec", reflect.TypeOf((*ArrayOfVsanMassCollectorSpec)(nil)).Elem())
}

type ArrayOfVsanMetricProfile struct {
	VsanMetricProfile []VsanMetricProfile `xml:"VsanMetricProfile,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanMetricProfile", reflect.TypeOf((*ArrayOfVsanMetricProfile)(nil)).Elem())
}

type ArrayOfVsanNetworkConfigBaseIssue struct {
	VsanNetworkConfigBaseIssue []BaseVsanNetworkConfigBaseIssue `xml:"VsanNetworkConfigBaseIssue,omitempty,typeattr"`
}

func init() {
	types.Add("vsan:ArrayOfVsanNetworkConfigBaseIssue", reflect.TypeOf((*ArrayOfVsanNetworkConfigBaseIssue)(nil)).Elem())
}

type ArrayOfVsanNetworkHealthResult struct {
	VsanNetworkHealthResult []VsanNetworkHealthResult `xml:"VsanNetworkHealthResult,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanNetworkHealthResult", reflect.TypeOf((*ArrayOfVsanNetworkHealthResult)(nil)).Elem())
}

type ArrayOfVsanNetworkLoadTestResult struct {
	VsanNetworkLoadTestResult []VsanNetworkLoadTestResult `xml:"VsanNetworkLoadTestResult,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanNetworkLoadTestResult", reflect.TypeOf((*ArrayOfVsanNetworkLoadTestResult)(nil)).Elem())
}

type ArrayOfVsanNetworkPeerHealthResult struct {
	VsanNetworkPeerHealthResult []VsanNetworkPeerHealthResult `xml:"VsanNetworkPeerHealthResult,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanNetworkPeerHealthResult", reflect.TypeOf((*ArrayOfVsanNetworkPeerHealthResult)(nil)).Elem())
}

type ArrayOfVsanObjectHealth struct {
	VsanObjectHealth []VsanObjectHealth `xml:"VsanObjectHealth,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanObjectHealth", reflect.TypeOf((*ArrayOfVsanObjectHealth)(nil)).Elem())
}

type ArrayOfVsanObjectIdentity struct {
	VsanObjectIdentity []VsanObjectIdentity `xml:"VsanObjectIdentity,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanObjectIdentity", reflect.TypeOf((*ArrayOfVsanObjectIdentity)(nil)).Elem())
}

type ArrayOfVsanObjectInformation struct {
	VsanObjectInformation []VsanObjectInformation `xml:"VsanObjectInformation,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanObjectInformation", reflect.TypeOf((*ArrayOfVsanObjectInformation)(nil)).Elem())
}

type ArrayOfVsanObjectQuerySpec struct {
	VsanObjectQuerySpec []VsanObjectQuerySpec `xml:"VsanObjectQuerySpec,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanObjectQuerySpec", reflect.TypeOf((*ArrayOfVsanObjectQuerySpec)(nil)).Elem())
}

type ArrayOfVsanObjectSpaceSummary struct {
	VsanObjectSpaceSummary []VsanObjectSpaceSummary `xml:"VsanObjectSpaceSummary,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanObjectSpaceSummary", reflect.TypeOf((*ArrayOfVsanObjectSpaceSummary)(nil)).Elem())
}

type ArrayOfVsanPerfDiagnosticException struct {
	VsanPerfDiagnosticException []VsanPerfDiagnosticException `xml:"VsanPerfDiagnosticException,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanPerfDiagnosticException", reflect.TypeOf((*ArrayOfVsanPerfDiagnosticException)(nil)).Elem())
}

type ArrayOfVsanPerfDiagnosticResult struct {
	VsanPerfDiagnosticResult []VsanPerfDiagnosticResult `xml:"VsanPerfDiagnosticResult,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanPerfDiagnosticResult", reflect.TypeOf((*ArrayOfVsanPerfDiagnosticResult)(nil)).Elem())
}

type ArrayOfVsanPerfEntityMetricCSV struct {
	VsanPerfEntityMetricCSV []VsanPerfEntityMetricCSV `xml:"VsanPerfEntityMetricCSV,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanPerfEntityMetricCSV", reflect.TypeOf((*ArrayOfVsanPerfEntityMetricCSV)(nil)).Elem())
}

type ArrayOfVsanPerfEntityType struct {
	VsanPerfEntityType []VsanPerfEntityType `xml:"VsanPerfEntityType,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanPerfEntityType", reflect.TypeOf((*ArrayOfVsanPerfEntityType)(nil)).Elem())
}

type ArrayOfVsanPerfGraph struct {
	VsanPerfGraph []VsanPerfGraph `xml:"VsanPerfGraph,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanPerfGraph", reflect.TypeOf((*ArrayOfVsanPerfGraph)(nil)).Elem())
}

type ArrayOfVsanPerfMemberInfo struct {
	VsanPerfMemberInfo []VsanPerfMemberInfo `xml:"VsanPerfMemberInfo,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanPerfMemberInfo", reflect.TypeOf((*ArrayOfVsanPerfMemberInfo)(nil)).Elem())
}

type ArrayOfVsanPerfMetricId struct {
	VsanPerfMetricId []VsanPerfMetricId `xml:"VsanPerfMetricId,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanPerfMetricId", reflect.TypeOf((*ArrayOfVsanPerfMetricId)(nil)).Elem())
}

type ArrayOfVsanPerfMetricSeriesCSV struct {
	VsanPerfMetricSeriesCSV []VsanPerfMetricSeriesCSV `xml:"VsanPerfMetricSeriesCSV,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanPerfMetricSeriesCSV", reflect.TypeOf((*ArrayOfVsanPerfMetricSeriesCSV)(nil)).Elem())
}

type ArrayOfVsanPerfNodeInformation struct {
	VsanPerfNodeInformation []VsanPerfNodeInformation `xml:"VsanPerfNodeInformation,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanPerfNodeInformation", reflect.TypeOf((*ArrayOfVsanPerfNodeInformation)(nil)).Elem())
}

type ArrayOfVsanPerfQuerySpec struct {
	VsanPerfQuerySpec []VsanPerfQuerySpec `xml:"VsanPerfQuerySpec,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanPerfQuerySpec", reflect.TypeOf((*ArrayOfVsanPerfQuerySpec)(nil)).Elem())
}

type ArrayOfVsanPerfTimeRange struct {
	VsanPerfTimeRange []VsanPerfTimeRange `xml:"VsanPerfTimeRange,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanPerfTimeRange", reflect.TypeOf((*ArrayOfVsanPerfTimeRange)(nil)).Elem())
}

type ArrayOfVsanPerfTopEntity struct {
	VsanPerfTopEntity []VsanPerfTopEntity `xml:"VsanPerfTopEntity,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanPerfTopEntity", reflect.TypeOf((*ArrayOfVsanPerfTopEntity)(nil)).Elem())
}

type ArrayOfVsanPhysicalDiskHealth struct {
	VsanPhysicalDiskHealth []VsanPhysicalDiskHealth `xml:"VsanPhysicalDiskHealth,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanPhysicalDiskHealth", reflect.TypeOf((*ArrayOfVsanPhysicalDiskHealth)(nil)).Elem())
}

type ArrayOfVsanPhysicalDiskHealthSummary struct {
	VsanPhysicalDiskHealthSummary []VsanPhysicalDiskHealthSummary `xml:"VsanPhysicalDiskHealthSummary,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanPhysicalDiskHealthSummary", reflect.TypeOf((*ArrayOfVsanPhysicalDiskHealthSummary)(nil)).Elem())
}

type ArrayOfVsanQueryResultHostInfo struct {
	VsanQueryResultHostInfo []VsanQueryResultHostInfo `xml:"VsanQueryResultHostInfo,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanQueryResultHostInfo", reflect.TypeOf((*ArrayOfVsanQueryResultHostInfo)(nil)).Elem())
}

type ArrayOfVsanResourceConstraint struct {
	VsanResourceConstraint []BaseVsanResourceConstraint `xml:"VsanResourceConstraint,omitempty,typeattr"`
}

func init() {
	types.Add("vsan:ArrayOfVsanResourceConstraint", reflect.TypeOf((*ArrayOfVsanResourceConstraint)(nil)).Elem())
}

type ArrayOfVsanResourceHealth struct {
	VsanResourceHealth []VsanResourceHealth `xml:"VsanResourceHealth,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanResourceHealth", reflect.TypeOf((*ArrayOfVsanResourceHealth)(nil)).Elem())
}

type ArrayOfVsanRuntimeStatsHostMap struct {
	VsanRuntimeStatsHostMap []VsanRuntimeStatsHostMap `xml:"VsanRuntimeStatsHostMap,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanRuntimeStatsHostMap", reflect.TypeOf((*ArrayOfVsanRuntimeStatsHostMap)(nil)).Elem())
}

type ArrayOfVsanSmartDiskStats struct {
	VsanSmartDiskStats []VsanSmartDiskStats `xml:"VsanSmartDiskStats,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanSmartDiskStats", reflect.TypeOf((*ArrayOfVsanSmartDiskStats)(nil)).Elem())
}

type ArrayOfVsanSmartParameter struct {
	VsanSmartParameter []VsanSmartParameter `xml:"VsanSmartParameter,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanSmartParameter", reflect.TypeOf((*ArrayOfVsanSmartParameter)(nil)).Elem())
}

type ArrayOfVsanSmartStatsHostSummary struct {
	VsanSmartStatsHostSummary []VsanSmartStatsHostSummary `xml:"VsanSmartStatsHostSummary,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanSmartStatsHostSummary", reflect.TypeOf((*ArrayOfVsanSmartStatsHostSummary)(nil)).Elem())
}

type ArrayOfVsanStorageComplianceResult struct {
	VsanStorageComplianceResult []VsanStorageComplianceResult `xml:"VsanStorageComplianceResult,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanStorageComplianceResult", reflect.TypeOf((*ArrayOfVsanStorageComplianceResult)(nil)).Elem())
}

type ArrayOfVsanStoragePolicyStatus struct {
	VsanStoragePolicyStatus []VsanStoragePolicyStatus `xml:"VsanStoragePolicyStatus,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanStoragePolicyStatus", reflect.TypeOf((*ArrayOfVsanStoragePolicyStatus)(nil)).Elem())
}

type ArrayOfVsanStorageWorkloadType struct {
	VsanStorageWorkloadType []VsanStorageWorkloadType `xml:"VsanStorageWorkloadType,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanStorageWorkloadType", reflect.TypeOf((*ArrayOfVsanStorageWorkloadType)(nil)).Elem())
}

type ArrayOfVsanUnicastAddressInfo struct {
	VsanUnicastAddressInfo []VsanUnicastAddressInfo `xml:"VsanUnicastAddressInfo,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanUnicastAddressInfo", reflect.TypeOf((*ArrayOfVsanUnicastAddressInfo)(nil)).Elem())
}

type ArrayOfVsanUpdateItem struct {
	VsanUpdateItem []VsanUpdateItem `xml:"VsanUpdateItem,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanUpdateItem", reflect.TypeOf((*ArrayOfVsanUpdateItem)(nil)).Elem())
}

type ArrayOfVsanVcsaDeploymentProgress struct {
	VsanVcsaDeploymentProgress []VsanVcsaDeploymentProgress `xml:"VsanVcsaDeploymentProgress,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanVcsaDeploymentProgress", reflect.TypeOf((*ArrayOfVsanVcsaDeploymentProgress)(nil)).Elem())
}

type ArrayOfVsanVdsPgMigrationHostInfo struct {
	VsanVdsPgMigrationHostInfo []VsanVdsPgMigrationHostInfo `xml:"VsanVdsPgMigrationHostInfo,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanVdsPgMigrationHostInfo", reflect.TypeOf((*ArrayOfVsanVdsPgMigrationHostInfo)(nil)).Elem())
}

type ArrayOfVsanVdsPgMigrationSpec struct {
	VsanVdsPgMigrationSpec []VsanVdsPgMigrationSpec `xml:"VsanVdsPgMigrationSpec,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanVdsPgMigrationSpec", reflect.TypeOf((*ArrayOfVsanVdsPgMigrationSpec)(nil)).Elem())
}

type ArrayOfVsanVdsPgMigrationVmInfo struct {
	VsanVdsPgMigrationVmInfo []VsanVdsPgMigrationVmInfo `xml:"VsanVdsPgMigrationVmInfo,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanVdsPgMigrationVmInfo", reflect.TypeOf((*ArrayOfVsanVdsPgMigrationVmInfo)(nil)).Elem())
}

type ArrayOfVsanVibScanResult struct {
	VsanVibScanResult []VsanVibScanResult `xml:"VsanVibScanResult,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanVibScanResult", reflect.TypeOf((*ArrayOfVsanVibScanResult)(nil)).Elem())
}

type ArrayOfVsanVibSpec struct {
	VsanVibSpec []VsanVibSpec `xml:"VsanVibSpec,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanVibSpec", reflect.TypeOf((*ArrayOfVsanVibSpec)(nil)).Elem())
}

type ArrayOfVsanVmVdsMigrationSpec struct {
	VsanVmVdsMigrationSpec []VsanVmVdsMigrationSpec `xml:"VsanVmVdsMigrationSpec,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanVmVdsMigrationSpec", reflect.TypeOf((*ArrayOfVsanVmVdsMigrationSpec)(nil)).Elem())
}

type ArrayOfVsanVmdkIOLoadSpec struct {
	VsanVmdkIOLoadSpec []VsanVmdkIOLoadSpec `xml:"VsanVmdkIOLoadSpec,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanVmdkIOLoadSpec", reflect.TypeOf((*ArrayOfVsanVmdkIOLoadSpec)(nil)).Elem())
}

type ArrayOfVsanVmdkLoadTestResult struct {
	VsanVmdkLoadTestResult []VsanVmdkLoadTestResult `xml:"VsanVmdkLoadTestResult,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanVmdkLoadTestResult", reflect.TypeOf((*ArrayOfVsanVmdkLoadTestResult)(nil)).Elem())
}

type ArrayOfVsanVmdkLoadTestSpec struct {
	VsanVmdkLoadTestSpec []VsanVmdkLoadTestSpec `xml:"VsanVmdkLoadTestSpec,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanVmdkLoadTestSpec", reflect.TypeOf((*ArrayOfVsanVmdkLoadTestSpec)(nil)).Elem())
}

type ArrayOfVsanVnicVdsMigrationSpec struct {
	VsanVnicVdsMigrationSpec []VsanVnicVdsMigrationSpec `xml:"VsanVnicVdsMigrationSpec,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanVnicVdsMigrationSpec", reflect.TypeOf((*ArrayOfVsanVnicVdsMigrationSpec)(nil)).Elem())
}

type ArrayOfVsanVsanClusterPcapGroup struct {
	VsanVsanClusterPcapGroup []VsanVsanClusterPcapGroup `xml:"VsanVsanClusterPcapGroup,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanVsanClusterPcapGroup", reflect.TypeOf((*ArrayOfVsanVsanClusterPcapGroup)(nil)).Elem())
}

type ArrayOfVsanVsanPcapResult struct {
	VsanVsanPcapResult []VsanVsanPcapResult `xml:"VsanVsanPcapResult,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanVsanPcapResult", reflect.TypeOf((*ArrayOfVsanVsanPcapResult)(nil)).Elem())
}

type ArrayOfVsanWhatifCapacity struct {
	VsanWhatifCapacity []VsanWhatifCapacity `xml:"VsanWhatifCapacity,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanWhatifCapacity", reflect.TypeOf((*ArrayOfVsanWhatifCapacity)(nil)).Elem())
}

type ArrayOfVsanWitnessHostConfig struct {
	VsanWitnessHostConfig []VsanWitnessHostConfig `xml:"VsanWitnessHostConfig,omitempty"`
}

func init() {
	types.Add("vsan:ArrayOfVsanWitnessHostConfig", reflect.TypeOf((*ArrayOfVsanWitnessHostConfig)(nil)).Elem())
}

type CnsAttachVolume CnsAttachVolumeRequestType

func init() {
	types.Add("vsan:CnsAttachVolume", reflect.TypeOf((*CnsAttachVolume)(nil)).Elem())
}

type CnsAttachVolumeRequestType struct {
	This        types.ManagedObjectReference `xml:"_this"`
	AttachSpecs []CnsVolumeAttachDetachSpec  `xml:"attachSpecs"`
}

func init() {
	types.Add("vsan:CnsAttachVolumeRequestType", reflect.TypeOf((*CnsAttachVolumeRequestType)(nil)).Elem())
}

type CnsAttachVolumeResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type CnsBackingObjectDetails struct {
	types.DynamicData

	CapacityInMb int64 `xml:"capacityInMb,omitempty"`
}

func init() {
	types.Add("vsan:CnsBackingObjectDetails", reflect.TypeOf((*CnsBackingObjectDetails)(nil)).Elem())
}

type CnsBaseCreateSpec struct {
	types.DynamicData
}

func init() {
	types.Add("vsan:CnsBaseCreateSpec", reflect.TypeOf((*CnsBaseCreateSpec)(nil)).Elem())
}

type CnsBlockBackingDetails struct {
	CnsBackingObjectDetails

	BackingDiskId string `xml:"backingDiskId,omitempty"`
}

func init() {
	types.Add("vsan:CnsBlockBackingDetails", reflect.TypeOf((*CnsBlockBackingDetails)(nil)).Elem())
}

type CnsContainerCluster struct {
	types.DynamicData

	ClusterType   string `xml:"clusterType"`
	ClusterId     string `xml:"clusterId"`
	VSphereUser   string `xml:"vSphereUser"`
	ClusterFlavor string `xml:"clusterFlavor,omitempty"`
}

func init() {
	types.Add("vsan:CnsContainerCluster", reflect.TypeOf((*CnsContainerCluster)(nil)).Elem())
}

type CnsCreateVolume CnsCreateVolumeRequestType

func init() {
	types.Add("vsan:CnsCreateVolume", reflect.TypeOf((*CnsCreateVolume)(nil)).Elem())
}

type CnsCreateVolumeRequestType struct {
	This        types.ManagedObjectReference `xml:"_this"`
	CreateSpecs []CnsVolumeCreateSpec        `xml:"createSpecs"`
}

func init() {
	types.Add("vsan:CnsCreateVolumeRequestType", reflect.TypeOf((*CnsCreateVolumeRequestType)(nil)).Elem())
}

type CnsCreateVolumeResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type CnsCursor struct {
	types.DynamicData

	Offset       int64 `xml:"offset"`
	Limit        int64 `xml:"limit"`
	TotalRecords int64 `xml:"totalRecords,omitempty"`
}

func init() {
	types.Add("vsan:CnsCursor", reflect.TypeOf((*CnsCursor)(nil)).Elem())
}

type CnsDeleteVolume CnsDeleteVolumeRequestType

func init() {
	types.Add("vsan:CnsDeleteVolume", reflect.TypeOf((*CnsDeleteVolume)(nil)).Elem())
}

type CnsDeleteVolumeRequestType struct {
	This       types.ManagedObjectReference `xml:"_this"`
	VolumeIds  []CnsVolumeId                `xml:"volumeIds"`
	DeleteDisk bool                         `xml:"deleteDisk"`
}

func init() {
	types.Add("vsan:CnsDeleteVolumeRequestType", reflect.TypeOf((*CnsDeleteVolumeRequestType)(nil)).Elem())
}

type CnsDeleteVolumeResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type CnsDetachVolume CnsDetachVolumeRequestType

func init() {
	types.Add("vsan:CnsDetachVolume", reflect.TypeOf((*CnsDetachVolume)(nil)).Elem())
}

type CnsDetachVolumeRequestType struct {
	This        types.ManagedObjectReference `xml:"_this"`
	DetachSpecs []CnsVolumeAttachDetachSpec  `xml:"detachSpecs"`
}

func init() {
	types.Add("vsan:CnsDetachVolumeRequestType", reflect.TypeOf((*CnsDetachVolumeRequestType)(nil)).Elem())
}

type CnsDetachVolumeResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type CnsEntityMetadata struct {
	types.DynamicData

	EntityName string           `xml:"entityName"`
	Labels     []types.KeyValue `xml:"labels,omitempty"`
	Delete     *bool            `xml:"delete"`
	ClusterId  string           `xml:"clusterId,omitempty"`
}

func init() {
	types.Add("vsan:CnsEntityMetadata", reflect.TypeOf((*CnsEntityMetadata)(nil)).Elem())
}

type CnsFault struct {
	types.VimFault

	Reason string `xml:"reason"`
}

func init() {
	types.Add("vsan:CnsFault", reflect.TypeOf((*CnsFault)(nil)).Elem())
}

type CnsFaultFault CnsFault

func init() {
	types.Add("vsan:CnsFaultFault", reflect.TypeOf((*CnsFaultFault)(nil)).Elem())
}

type CnsFileBackingDetails struct {
	CnsBackingObjectDetails

	BackingFileId string `xml:"backingFileId,omitempty"`
}

func init() {
	types.Add("vsan:CnsFileBackingDetails", reflect.TypeOf((*CnsFileBackingDetails)(nil)).Elem())
}

type CnsFileCreateSpec struct {
	CnsBaseCreateSpec
}

func init() {
	types.Add("vsan:CnsFileCreateSpec", reflect.TypeOf((*CnsFileCreateSpec)(nil)).Elem())
}

type CnsKubernetesEntityMetadata struct {
	CnsEntityMetadata

	EntityType     string                         `xml:"entityType"`
	Namespace      string                         `xml:"namespace,omitempty"`
	ReferredEntity []CnsKubernetesEntityReference `xml:"referredEntity,omitempty"`
}

func init() {
	types.Add("vsan:CnsKubernetesEntityMetadata", reflect.TypeOf((*CnsKubernetesEntityMetadata)(nil)).Elem())
}

type CnsKubernetesEntityReference struct {
	types.DynamicData

	EntityType string `xml:"entityType"`
	EntityName string `xml:"entityName"`
	Namespace  string `xml:"namespace,omitempty"`
	ClusterId  string `xml:"clusterId,omitempty"`
}

func init() {
	types.Add("vsan:CnsKubernetesEntityReference", reflect.TypeOf((*CnsKubernetesEntityReference)(nil)).Elem())
}

type CnsKubernetesQueryFilter struct {
	CnsQueryFilter

	Namespaces []string `xml:"namespaces,omitempty"`
	PodNames   []string `xml:"podNames,omitempty"`
	PvcNames   []string `xml:"pvcNames,omitempty"`
	PvNames    []string `xml:"pvNames,omitempty"`
}

func init() {
	types.Add("vsan:CnsKubernetesQueryFilter", reflect.TypeOf((*CnsKubernetesQueryFilter)(nil)).Elem())
}

type CnsQueryFilter struct {
	types.DynamicData

	VolumeIds                    []CnsVolumeId                  `xml:"volumeIds,omitempty"`
	Names                        []string                       `xml:"names,omitempty"`
	ContainerClusterIds          []string                       `xml:"containerClusterIds,omitempty"`
	StoragePolicyId              string                         `xml:"storagePolicyId,omitempty"`
	Datastores                   []types.ManagedObjectReference `xml:"datastores,omitempty"`
	Labels                       []types.KeyValue               `xml:"labels,omitempty"`
	ComplianceStatus             string                         `xml:"complianceStatus,omitempty"`
	DatastoreAccessibilityStatus string                         `xml:"datastoreAccessibilityStatus,omitempty"`
	Cursor                       *CnsCursor                     `xml:"cursor,omitempty"`
	HealthStatus                 string                         `xml:"healthStatus,omitempty"`
}

func init() {
	types.Add("vsan:CnsQueryFilter", reflect.TypeOf((*CnsQueryFilter)(nil)).Elem())
}

type CnsQueryResult struct {
	types.DynamicData

	Volumes []CnsVolume `xml:"volumes,omitempty"`
	Cursor  CnsCursor   `xml:"cursor"`
}

func init() {
	types.Add("vsan:CnsQueryResult", reflect.TypeOf((*CnsQueryResult)(nil)).Elem())
}

type CnsQueryVolume CnsQueryVolumeRequestType

func init() {
	types.Add("vsan:CnsQueryVolume", reflect.TypeOf((*CnsQueryVolume)(nil)).Elem())
}

type CnsQueryVolumeRequestType struct {
	This   types.ManagedObjectReference `xml:"_this"`
	Filter BaseCnsQueryFilter           `xml:"filter,typeattr"`
}

func init() {
	types.Add("vsan:CnsQueryVolumeRequestType", reflect.TypeOf((*CnsQueryVolumeRequestType)(nil)).Elem())
}

type CnsQueryVolumeResponse struct {
	Returnval CnsQueryResult `xml:"returnval"`
}

type CnsSnapshotId struct {
	types.DynamicData

	Id string `xml:"id"`
}

func init() {
	types.Add("vsan:CnsSnapshotId", reflect.TypeOf((*CnsSnapshotId)(nil)).Elem())
}

type CnsSnapshotVolumeSource struct {
	CnsVolumeSource

	VolumeId   *CnsVolumeId   `xml:"volumeId,omitempty"`
	SnapshotId *CnsSnapshotId `xml:"snapshotId,omitempty"`
}

func init() {
	types.Add("vsan:CnsSnapshotVolumeSource", reflect.TypeOf((*CnsSnapshotVolumeSource)(nil)).Elem())
}

type CnsUpdateVolumeMetadata CnsUpdateVolumeMetadataRequestType

func init() {
	types.Add("vsan:CnsUpdateVolumeMetadata", reflect.TypeOf((*CnsUpdateVolumeMetadata)(nil)).Elem())
}

type CnsUpdateVolumeMetadataRequestType struct {
	This        types.ManagedObjectReference  `xml:"_this"`
	UpdateSpecs []CnsVolumeMetadataUpdateSpec `xml:"updateSpecs"`
}

func init() {
	types.Add("vsan:CnsUpdateVolumeMetadataRequestType", reflect.TypeOf((*CnsUpdateVolumeMetadataRequestType)(nil)).Elem())
}

type CnsUpdateVolumeMetadataResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type CnsVSANFileCreateSpec struct {
	CnsFileCreateSpec

	SoftQuotaInMb int64                        `xml:"softQuotaInMb,omitempty"`
	Permission    []VsanFileShareNetPermission `xml:"permission,omitempty"`
}

func init() {
	types.Add("vsan:CnsVSANFileCreateSpec", reflect.TypeOf((*CnsVSANFileCreateSpec)(nil)).Elem())
}

type CnsVolume struct {
	types.DynamicData

	VolumeId                     CnsVolumeId                 `xml:"volumeId"`
	DatastoreUrl                 string                      `xml:"datastoreUrl,omitempty"`
	Name                         string                      `xml:"name,omitempty"`
	VolumeType                   string                      `xml:"volumeType,omitempty"`
	StoragePolicyId              string                      `xml:"storagePolicyId,omitempty"`
	Metadata                     *CnsVolumeMetadata          `xml:"metadata,omitempty"`
	BackingObjectDetails         BaseCnsBackingObjectDetails `xml:"backingObjectDetails,omitempty,typeattr"`
	ComplianceStatus             string                      `xml:"complianceStatus,omitempty"`
	DatastoreAccessibilityStatus string                      `xml:"datastoreAccessibilityStatus,omitempty"`
	HealthStatus                 string                      `xml:"healthStatus,omitempty"`
}

func init() {
	types.Add("vsan:CnsVolume", reflect.TypeOf((*CnsVolume)(nil)).Elem())
}

type CnsVolumeAttachDetachSpec struct {
	types.DynamicData

	VolumeId CnsVolumeId                  `xml:"volumeId"`
	Vm       types.ManagedObjectReference `xml:"vm"`
}

func init() {
	types.Add("vsan:CnsVolumeAttachDetachSpec", reflect.TypeOf((*CnsVolumeAttachDetachSpec)(nil)).Elem())
}

type CnsVolumeAttachResult struct {
	CnsVolumeOperationResult

	DiskUUID string `xml:"diskUUID,omitempty"`
}

func init() {
	types.Add("vsan:CnsVolumeAttachResult", reflect.TypeOf((*CnsVolumeAttachResult)(nil)).Elem())
}

type CnsVolumeCreateResult struct {
	CnsVolumeOperationResult

	Name string `xml:"name,omitempty"`
}

func init() {
	types.Add("vsan:CnsVolumeCreateResult", reflect.TypeOf((*CnsVolumeCreateResult)(nil)).Elem())
}

type CnsVolumeCreateSpec struct {
	types.DynamicData

	Name                 string                                `xml:"name"`
	VolumeType           string                                `xml:"volumeType"`
	Datastores           []types.ManagedObjectReference        `xml:"datastores,omitempty"`
	Metadata             *CnsVolumeMetadata                    `xml:"metadata,omitempty"`
	BackingObjectDetails BaseCnsBackingObjectDetails           `xml:"backingObjectDetails,typeattr"`
	Profile              []types.BaseVirtualMachineProfileSpec `xml:"profile,omitempty,typeattr"`
	CreateSpec           BaseCnsBaseCreateSpec                 `xml:"createSpec,omitempty,typeattr"`
	VolumeSource         BaseCnsVolumeSource                   `xml:"volumeSource,omitempty,typeattr"`
}

func init() {
	types.Add("vsan:CnsVolumeCreateSpec", reflect.TypeOf((*CnsVolumeCreateSpec)(nil)).Elem())
}

type CnsVolumeId struct {
	types.DynamicData

	Id string `xml:"id"`
}

func init() {
	types.Add("vsan:CnsVolumeId", reflect.TypeOf((*CnsVolumeId)(nil)).Elem())
}

type CnsVolumeMetadata struct {
	types.DynamicData

	ContainerCluster      CnsContainerCluster     `xml:"containerCluster"`
	EntityMetadata        []BaseCnsEntityMetadata `xml:"entityMetadata,omitempty,typeattr"`
	ContainerClusterArray []CnsContainerCluster   `xml:"containerClusterArray,omitempty"`
}

func init() {
	types.Add("vsan:CnsVolumeMetadata", reflect.TypeOf((*CnsVolumeMetadata)(nil)).Elem())
}

type CnsVolumeMetadataUpdateSpec struct {
	types.DynamicData

	VolumeId CnsVolumeId       `xml:"volumeId"`
	Metadata CnsVolumeMetadata `xml:"metadata"`
}

func init() {
	types.Add("vsan:CnsVolumeMetadataUpdateSpec", reflect.TypeOf((*CnsVolumeMetadataUpdateSpec)(nil)).Elem())
}

type CnsVolumeOperationBatchResult struct {
	types.DynamicData

	VolumeResults []BaseCnsVolumeOperationResult `xml:"volumeResults,omitempty,typeattr"`
}

func init() {
	types.Add("vsan:CnsVolumeOperationBatchResult", reflect.TypeOf((*CnsVolumeOperationBatchResult)(nil)).Elem())
}

type CnsVolumeOperationResult struct {
	types.DynamicData

	VolumeId *CnsVolumeId                `xml:"volumeId,omitempty"`
	Fault    *types.LocalizedMethodFault `xml:"fault,omitempty"`
}

func init() {
	types.Add("vsan:CnsVolumeOperationResult", reflect.TypeOf((*CnsVolumeOperationResult)(nil)).Elem())
}

type CnsVolumeSource struct {
	types.DynamicData
}

func init() {
	types.Add("vsan:CnsVolumeSource", reflect.TypeOf((*CnsVolumeSource)(nil)).Elem())
}

type CnsVsanFileShareBackingDetails struct {
	CnsFileBackingDetails

	Name         string           `xml:"name,omitempty"`
	AccessPoints []types.KeyValue `xml:"accessPoints,omitempty"`
}

func init() {
	types.Add("vsan:CnsVsanFileShareBackingDetails", reflect.TypeOf((*CnsVsanFileShareBackingDetails)(nil)).Elem())
}

type EntityResourceCheckDetails struct {
	types.DynamicData

	Name                       string `xml:"name,omitempty"`
	Uuid                       string `xml:"uuid,omitempty"`
	IsNew                      *bool  `xml:"isNew"`
	Capacity                   int64  `xml:"capacity,omitempty"`
	PostOperationCapacity      int64  `xml:"postOperationCapacity,omitempty"`
	UsedCapacity               int64  `xml:"usedCapacity,omitempty"`
	PostOperationUsedCapacity  int64  `xml:"postOperationUsedCapacity,omitempty"`
	AdditionalRequiredCapacity int64  `xml:"additionalRequiredCapacity,omitempty"`
	MaxComponents              int64  `xml:"maxComponents,omitempty"`
	Components                 int64  `xml:"components,omitempty"`
}

func init() {
	types.Add("vsan:EntityResourceCheckDetails", reflect.TypeOf((*EntityResourceCheckDetails)(nil)).Elem())
}

type FetchIsoDepotCookie FetchIsoDepotCookieRequestType

func init() {
	types.Add("vsan:FetchIsoDepotCookie", reflect.TypeOf((*FetchIsoDepotCookie)(nil)).Elem())
}

type FetchIsoDepotCookieRequestType struct {
	This     types.ManagedObjectReference `xml:"_this"`
	Username string                       `xml:"username"`
	Password string                       `xml:"password"`
}

func init() {
	types.Add("vsan:FetchIsoDepotCookieRequestType", reflect.TypeOf((*FetchIsoDepotCookieRequestType)(nil)).Elem())
}

type FetchIsoDepotCookieResponse struct {
}

type FileShareQueryResult struct {
	types.DynamicData

	FileShares      []VsanFileShare `xml:"fileShares,omitempty"`
	NextOffset      string          `xml:"nextOffset,omitempty"`
	TotalShareCount int64           `xml:"totalShareCount,omitempty"`
}

func init() {
	types.Add("vsan:FileShareQueryResult", reflect.TypeOf((*FileShareQueryResult)(nil)).Elem())
}

type GetVsanPerfDiagnosisResult GetVsanPerfDiagnosisResultRequestType

func init() {
	types.Add("vsan:GetVsanPerfDiagnosisResult", reflect.TypeOf((*GetVsanPerfDiagnosisResult)(nil)).Elem())
}

type GetVsanPerfDiagnosisResultRequestType struct {
	This    types.ManagedObjectReference  `xml:"_this"`
	Task    types.ManagedObjectReference  `xml:"task"`
	Cluster *types.ManagedObjectReference `xml:"cluster,omitempty"`
}

func init() {
	types.Add("vsan:GetVsanPerfDiagnosisResultRequestType", reflect.TypeOf((*GetVsanPerfDiagnosisResultRequestType)(nil)).Elem())
}

type GetVsanPerfDiagnosisResultResponse struct {
	Returnval []VsanPerfDiagnosticResult `xml:"returnval,omitempty"`
}

type GetVsanVumConfig GetVsanVumConfigRequestType

func init() {
	types.Add("vsan:GetVsanVumConfig", reflect.TypeOf((*GetVsanVumConfig)(nil)).Elem())
}

type GetVsanVumConfigRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

func init() {
	types.Add("vsan:GetVsanVumConfigRequestType", reflect.TypeOf((*GetVsanVumConfigRequestType)(nil)).Elem())
}

type GetVsanVumConfigResponse struct {
	Returnval VsanVumSystemConfig `xml:"returnval"`
}

type InitializeDiskMappings InitializeDiskMappingsRequestType

func init() {
	types.Add("vsan:InitializeDiskMappings", reflect.TypeOf((*InitializeDiskMappings)(nil)).Elem())
}

type InitializeDiskMappingsRequestType struct {
	This types.ManagedObjectReference       `xml:"_this"`
	Spec VimVsanHostDiskMappingCreationSpec `xml:"spec"`
}

func init() {
	types.Add("vsan:InitializeDiskMappingsRequestType", reflect.TypeOf((*InitializeDiskMappingsRequestType)(nil)).Elem())
}

type InitializeDiskMappingsResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type PerformVsanUpgradeEx PerformVsanUpgradeExRequestType

func init() {
	types.Add("vsan:PerformVsanUpgradeEx", reflect.TypeOf((*PerformVsanUpgradeEx)(nil)).Elem())
}

type PerformVsanUpgradeExRequestType struct {
	This                   types.ManagedObjectReference   `xml:"_this"`
	Cluster                types.ManagedObjectReference   `xml:"cluster"`
	PerformObjectUpgrade   *bool                          `xml:"performObjectUpgrade"`
	DowngradeFormat        *bool                          `xml:"downgradeFormat"`
	AllowReducedRedundancy *bool                          `xml:"allowReducedRedundancy"`
	ExcludeHosts           []types.ManagedObjectReference `xml:"excludeHosts,omitempty"`
	Spec                   *VsanDiskFormatConversionSpec  `xml:"spec,omitempty"`
}

func init() {
	types.Add("vsan:PerformVsanUpgradeExRequestType", reflect.TypeOf((*PerformVsanUpgradeExRequestType)(nil)).Elem())
}

type PerformVsanUpgradeExResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type PerformVsanUpgradePreflightAsyncCheckRequestType struct {
	This            types.ManagedObjectReference  `xml:"_this"`
	Cluster         types.ManagedObjectReference  `xml:"cluster"`
	DowngradeFormat *bool                         `xml:"downgradeFormat"`
	Spec            *VsanDiskFormatConversionSpec `xml:"spec,omitempty"`
}

func init() {
	types.Add("vsan:PerformVsanUpgradePreflightAsyncCheckRequestType", reflect.TypeOf((*PerformVsanUpgradePreflightAsyncCheckRequestType)(nil)).Elem())
}

type PerformVsanUpgradePreflightAsyncCheck_Task PerformVsanUpgradePreflightAsyncCheckRequestType

func init() {
	types.Add("vsan:PerformVsanUpgradePreflightAsyncCheck_Task", reflect.TypeOf((*PerformVsanUpgradePreflightAsyncCheck_Task)(nil)).Elem())
}

type PerformVsanUpgradePreflightAsyncCheck_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type PerformVsanUpgradePreflightCheckEx PerformVsanUpgradePreflightCheckExRequestType

func init() {
	types.Add("vsan:PerformVsanUpgradePreflightCheckEx", reflect.TypeOf((*PerformVsanUpgradePreflightCheckEx)(nil)).Elem())
}

type PerformVsanUpgradePreflightCheckExRequestType struct {
	This            types.ManagedObjectReference  `xml:"_this"`
	Cluster         types.ManagedObjectReference  `xml:"cluster"`
	DowngradeFormat *bool                         `xml:"downgradeFormat"`
	Spec            *VsanDiskFormatConversionSpec `xml:"spec,omitempty"`
}

func init() {
	types.Add("vsan:PerformVsanUpgradePreflightCheckExRequestType", reflect.TypeOf((*PerformVsanUpgradePreflightCheckExRequestType)(nil)).Elem())
}

type PerformVsanUpgradePreflightCheckExResponse struct {
	Returnval VsanDiskFormatConversionCheckResult `xml:"returnval"`
}

type QueryClusterDataEfficiencyCapacityState QueryClusterDataEfficiencyCapacityStateRequestType

func init() {
	types.Add("vsan:QueryClusterDataEfficiencyCapacityState", reflect.TypeOf((*QueryClusterDataEfficiencyCapacityState)(nil)).Elem())
}

type QueryClusterDataEfficiencyCapacityStateRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

func init() {
	types.Add("vsan:QueryClusterDataEfficiencyCapacityStateRequestType", reflect.TypeOf((*QueryClusterDataEfficiencyCapacityStateRequestType)(nil)).Elem())
}

type QueryClusterDataEfficiencyCapacityStateResponse struct {
	Returnval VimVsanDataEfficiencyCapacityState `xml:"returnval"`
}

type QueryDiskMappings QueryDiskMappingsRequestType

func init() {
	types.Add("vsan:QueryDiskMappings", reflect.TypeOf((*QueryDiskMappings)(nil)).Elem())
}

type QueryDiskMappingsRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
	Host types.ManagedObjectReference `xml:"host"`
}

func init() {
	types.Add("vsan:QueryDiskMappingsRequestType", reflect.TypeOf((*QueryDiskMappingsRequestType)(nil)).Elem())
}

type QueryDiskMappingsResponse struct {
	Returnval []VimVsanHostDiskMapInfoEx `xml:"returnval,omitempty"`
}

type QuerySyncingVsanObjectsSummary QuerySyncingVsanObjectsSummaryRequestType

func init() {
	types.Add("vsan:QuerySyncingVsanObjectsSummary", reflect.TypeOf((*QuerySyncingVsanObjectsSummary)(nil)).Elem())
}

type QuerySyncingVsanObjectsSummaryRequestType struct {
	This                types.ManagedObjectReference `xml:"_this"`
	Cluster             types.ManagedObjectReference `xml:"cluster"`
	SyncingObjectFilter *VsanSyncingObjectFilter     `xml:"syncingObjectFilter,omitempty"`
}

func init() {
	types.Add("vsan:QuerySyncingVsanObjectsSummaryRequestType", reflect.TypeOf((*QuerySyncingVsanObjectsSummaryRequestType)(nil)).Elem())
}

type QuerySyncingVsanObjectsSummaryResponse struct {
	Returnval VsanHostVsanObjectSyncQueryResult `xml:"returnval"`
}

type QueryVsanCloudHealthStatus QueryVsanCloudHealthStatusRequestType

func init() {
	types.Add("vsan:QueryVsanCloudHealthStatus", reflect.TypeOf((*QueryVsanCloudHealthStatus)(nil)).Elem())
}

type QueryVsanCloudHealthStatusRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

func init() {
	types.Add("vsan:QueryVsanCloudHealthStatusRequestType", reflect.TypeOf((*QueryVsanCloudHealthStatusRequestType)(nil)).Elem())
}

type QueryVsanCloudHealthStatusResponse struct {
	Returnval *VsanCloudHealthStatus `xml:"returnval,omitempty"`
}

type RebuildDiskMapping RebuildDiskMappingRequestType

func init() {
	types.Add("vsan:RebuildDiskMapping", reflect.TypeOf((*RebuildDiskMapping)(nil)).Elem())
}

type RebuildDiskMappingRequestType struct {
	This            types.ManagedObjectReference `xml:"_this"`
	Host            types.ManagedObjectReference `xml:"host"`
	Mapping         VsanHostDiskMapping          `xml:"mapping"`
	MaintenanceSpec types.HostMaintenanceSpec    `xml:"maintenanceSpec"`
}

func init() {
	types.Add("vsan:RebuildDiskMappingRequestType", reflect.TypeOf((*RebuildDiskMappingRequestType)(nil)).Elem())
}

type RebuildDiskMappingResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type RepairTimerInfo struct {
	types.DynamicData

	MaxTimeToRepair            int32 `xml:"maxTimeToRepair"`
	MinTimeToRepair            int32 `xml:"minTimeToRepair"`
	ObjectCount                int32 `xml:"objectCount"`
	ObjectCountWithRepairTimer int32 `xml:"objectCountWithRepairTimer,omitempty"`
}

func init() {
	types.Add("vsan:RepairTimerInfo", reflect.TypeOf((*RepairTimerInfo)(nil)).Elem())
}

type ResyncIopsInfo struct {
	types.DynamicData

	ResyncIops int32 `xml:"resyncIops"`
}

func init() {
	types.Add("vsan:ResyncIopsInfo", reflect.TypeOf((*ResyncIopsInfo)(nil)).Elem())
}

type RetrieveAllFlashCapabilities RetrieveAllFlashCapabilitiesRequestType

func init() {
	types.Add("vsan:RetrieveAllFlashCapabilities", reflect.TypeOf((*RetrieveAllFlashCapabilities)(nil)).Elem())
}

type RetrieveAllFlashCapabilitiesRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

func init() {
	types.Add("vsan:RetrieveAllFlashCapabilitiesRequestType", reflect.TypeOf((*RetrieveAllFlashCapabilitiesRequestType)(nil)).Elem())
}

type RetrieveAllFlashCapabilitiesResponse struct {
	Returnval []VimVsanHostVsanHostCapability `xml:"returnval,omitempty"`
}

type RetrieveSupportedVsanFormatVersion RetrieveSupportedVsanFormatVersionRequestType

func init() {
	types.Add("vsan:RetrieveSupportedVsanFormatVersion", reflect.TypeOf((*RetrieveSupportedVsanFormatVersion)(nil)).Elem())
}

type RetrieveSupportedVsanFormatVersionRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

func init() {
	types.Add("vsan:RetrieveSupportedVsanFormatVersionRequestType", reflect.TypeOf((*RetrieveSupportedVsanFormatVersionRequestType)(nil)).Elem())
}

type RetrieveSupportedVsanFormatVersionResponse struct {
	Returnval int32 `xml:"returnval"`
}

type VSANIsWitnessVirtualAppliance VSANIsWitnessVirtualApplianceRequestType

func init() {
	types.Add("vsan:VSANIsWitnessVirtualAppliance", reflect.TypeOf((*VSANIsWitnessVirtualAppliance)(nil)).Elem())
}

type VSANIsWitnessVirtualApplianceRequestType struct {
	This  types.ManagedObjectReference   `xml:"_this"`
	Hosts []types.ManagedObjectReference `xml:"hosts"`
}

func init() {
	types.Add("vsan:VSANIsWitnessVirtualApplianceRequestType", reflect.TypeOf((*VSANIsWitnessVirtualApplianceRequestType)(nil)).Elem())
}

type VSANIsWitnessVirtualApplianceResponse struct {
	Returnval []VsanHostVirtualApplianceInfo `xml:"returnval,omitempty"`
}

type VSANStretchedClusterHostVirtualApplianceStatus struct {
	types.DynamicData

	VcCluster    *types.ManagedObjectReference `xml:"vcCluster,omitempty"`
	IsVirtualApp *bool                         `xml:"isVirtualApp"`
}

func init() {
	types.Add("vsan:VSANStretchedClusterHostVirtualApplianceStatus", reflect.TypeOf((*VSANStretchedClusterHostVirtualApplianceStatus)(nil)).Elem())
}

type VSANVcAddWitnessHost VSANVcAddWitnessHostRequestType

func init() {
	types.Add("vsan:VSANVcAddWitnessHost", reflect.TypeOf((*VSANVcAddWitnessHost)(nil)).Elem())
}

type VSANVcAddWitnessHostRequestType struct {
	This         types.ManagedObjectReference `xml:"_this"`
	Cluster      types.ManagedObjectReference `xml:"cluster"`
	WitnessHost  types.ManagedObjectReference `xml:"witnessHost"`
	PreferredFd  string                       `xml:"preferredFd"`
	DiskMapping  *VsanHostDiskMapping         `xml:"diskMapping,omitempty"`
	MetadataMode *bool                        `xml:"metadataMode"`
}

func init() {
	types.Add("vsan:VSANVcAddWitnessHostRequestType", reflect.TypeOf((*VSANVcAddWitnessHostRequestType)(nil)).Elem())
}

type VSANVcAddWitnessHostResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VSANVcConvertToStretchedCluster VSANVcConvertToStretchedClusterRequestType

func init() {
	types.Add("vsan:VSANVcConvertToStretchedCluster", reflect.TypeOf((*VSANVcConvertToStretchedCluster)(nil)).Elem())
}

type VSANVcConvertToStretchedClusterRequestType struct {
	This              types.ManagedObjectReference                    `xml:"_this"`
	Cluster           types.ManagedObjectReference                    `xml:"cluster"`
	FaultDomainConfig VimClusterVSANStretchedClusterFaultDomainConfig `xml:"faultDomainConfig"`
	WitnessHost       types.ManagedObjectReference                    `xml:"witnessHost"`
	PreferredFd       string                                          `xml:"preferredFd"`
	DiskMapping       *VsanHostDiskMapping                            `xml:"diskMapping,omitempty"`
}

func init() {
	types.Add("vsan:VSANVcConvertToStretchedClusterRequestType", reflect.TypeOf((*VSANVcConvertToStretchedClusterRequestType)(nil)).Elem())
}

type VSANVcConvertToStretchedClusterResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VSANVcGetPreferredFaultDomain VSANVcGetPreferredFaultDomainRequestType

func init() {
	types.Add("vsan:VSANVcGetPreferredFaultDomain", reflect.TypeOf((*VSANVcGetPreferredFaultDomain)(nil)).Elem())
}

type VSANVcGetPreferredFaultDomainRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

func init() {
	types.Add("vsan:VSANVcGetPreferredFaultDomainRequestType", reflect.TypeOf((*VSANVcGetPreferredFaultDomainRequestType)(nil)).Elem())
}

type VSANVcGetPreferredFaultDomainResponse struct {
	Returnval *VimClusterVSANPreferredFaultDomainInfo `xml:"returnval,omitempty"`
}

type VSANVcGetWitnessHosts VSANVcGetWitnessHostsRequestType

func init() {
	types.Add("vsan:VSANVcGetWitnessHosts", reflect.TypeOf((*VSANVcGetWitnessHosts)(nil)).Elem())
}

type VSANVcGetWitnessHostsRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

func init() {
	types.Add("vsan:VSANVcGetWitnessHostsRequestType", reflect.TypeOf((*VSANVcGetWitnessHostsRequestType)(nil)).Elem())
}

type VSANVcGetWitnessHostsResponse struct {
	Returnval []VimClusterVSANWitnessHostInfo `xml:"returnval,omitempty"`
}

type VSANVcIsWitnessHost VSANVcIsWitnessHostRequestType

func init() {
	types.Add("vsan:VSANVcIsWitnessHost", reflect.TypeOf((*VSANVcIsWitnessHost)(nil)).Elem())
}

type VSANVcIsWitnessHostRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
	Host types.ManagedObjectReference `xml:"host"`
}

func init() {
	types.Add("vsan:VSANVcIsWitnessHostRequestType", reflect.TypeOf((*VSANVcIsWitnessHostRequestType)(nil)).Elem())
}

type VSANVcIsWitnessHostResponse struct {
	Returnval bool `xml:"returnval"`
}

type VSANVcRemoveWitnessHost VSANVcRemoveWitnessHostRequestType

func init() {
	types.Add("vsan:VSANVcRemoveWitnessHost", reflect.TypeOf((*VSANVcRemoveWitnessHost)(nil)).Elem())
}

type VSANVcRemoveWitnessHostRequestType struct {
	This           types.ManagedObjectReference  `xml:"_this"`
	Cluster        types.ManagedObjectReference  `xml:"cluster"`
	WitnessHost    *types.ManagedObjectReference `xml:"witnessHost,omitempty"`
	WitnessAddress string                        `xml:"witnessAddress,omitempty"`
}

func init() {
	types.Add("vsan:VSANVcRemoveWitnessHostRequestType", reflect.TypeOf((*VSANVcRemoveWitnessHostRequestType)(nil)).Elem())
}

type VSANVcRemoveWitnessHostResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VSANVcRetrieveStretchedClusterVcCapability VSANVcRetrieveStretchedClusterVcCapabilityRequestType

func init() {
	types.Add("vsan:VSANVcRetrieveStretchedClusterVcCapability", reflect.TypeOf((*VSANVcRetrieveStretchedClusterVcCapability)(nil)).Elem())
}

type VSANVcRetrieveStretchedClusterVcCapabilityRequestType struct {
	This               types.ManagedObjectReference `xml:"_this"`
	Cluster            types.ManagedObjectReference `xml:"cluster"`
	VerifyAllConnected *bool                        `xml:"verifyAllConnected"`
}

func init() {
	types.Add("vsan:VSANVcRetrieveStretchedClusterVcCapabilityRequestType", reflect.TypeOf((*VSANVcRetrieveStretchedClusterVcCapabilityRequestType)(nil)).Elem())
}

type VSANVcRetrieveStretchedClusterVcCapabilityResponse struct {
	Returnval []VimClusterVSANStretchedClusterCapability `xml:"returnval,omitempty"`
}

type VSANVcSetPreferredFaultDomain VSANVcSetPreferredFaultDomainRequestType

func init() {
	types.Add("vsan:VSANVcSetPreferredFaultDomain", reflect.TypeOf((*VSANVcSetPreferredFaultDomain)(nil)).Elem())
}

type VSANVcSetPreferredFaultDomainRequestType struct {
	This        types.ManagedObjectReference  `xml:"_this"`
	Cluster     types.ManagedObjectReference  `xml:"cluster"`
	PreferredFd string                        `xml:"preferredFd"`
	WitnessHost *types.ManagedObjectReference `xml:"witnessHost,omitempty"`
}

func init() {
	types.Add("vsan:VSANVcSetPreferredFaultDomainRequestType", reflect.TypeOf((*VSANVcSetPreferredFaultDomainRequestType)(nil)).Elem())
}

type VSANVcSetPreferredFaultDomainResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VimClusterVSANPreferredFaultDomainInfo struct {
	types.DynamicData

	PreferredFaultDomainName string `xml:"preferredFaultDomainName,omitempty"`
	PreferredFaultDomainId   string `xml:"preferredFaultDomainId,omitempty"`
}

func init() {
	types.Add("vsan:VimClusterVSANPreferredFaultDomainInfo", reflect.TypeOf((*VimClusterVSANPreferredFaultDomainInfo)(nil)).Elem())
}

type VimClusterVSANStretchedClusterCapability struct {
	types.DynamicData

	HostMoId       string                                     `xml:"hostMoId"`
	ConnStatus     string                                     `xml:"connStatus,omitempty"`
	IsSupported    *bool                                      `xml:"isSupported"`
	HostCapability *VimHostVSANStretchedClusterHostCapability `xml:"hostCapability,omitempty"`
}

func init() {
	types.Add("vsan:VimClusterVSANStretchedClusterCapability", reflect.TypeOf((*VimClusterVSANStretchedClusterCapability)(nil)).Elem())
}

type VimClusterVSANStretchedClusterFaultDomainConfig struct {
	types.DynamicData

	FirstFdName   string                         `xml:"firstFdName"`
	FirstFdHosts  []types.ManagedObjectReference `xml:"firstFdHosts"`
	SecondFdName  string                         `xml:"secondFdName"`
	SecondFdHosts []types.ManagedObjectReference `xml:"secondFdHosts"`
}

func init() {
	types.Add("vsan:VimClusterVSANStretchedClusterFaultDomainConfig", reflect.TypeOf((*VimClusterVSANStretchedClusterFaultDomainConfig)(nil)).Elem())
}

type VimClusterVSANWitnessHostInfo struct {
	types.DynamicData

	NodeUuid         string                        `xml:"nodeUuid"`
	FaultDomainName  string                        `xml:"faultDomainName,omitempty"`
	PreferredFdName  string                        `xml:"preferredFdName,omitempty"`
	PreferredFdUuid  string                        `xml:"preferredFdUuid,omitempty"`
	UnicastAgentAddr string                        `xml:"unicastAgentAddr,omitempty"`
	Host             *types.ManagedObjectReference `xml:"host,omitempty"`
	MetadataMode     *bool                         `xml:"metadataMode"`
}

func init() {
	types.Add("vsan:VimClusterVSANWitnessHostInfo", reflect.TypeOf((*VimClusterVSANWitnessHostInfo)(nil)).Elem())
}

type VimClusterVsanDiskMappingsConfigSpec struct {
	types.DynamicData

	HostDiskMappings []VimClusterVsanHostDiskMapping `xml:"hostDiskMappings"`
}

func init() {
	types.Add("vsan:VimClusterVsanDiskMappingsConfigSpec", reflect.TypeOf((*VimClusterVsanDiskMappingsConfigSpec)(nil)).Elem())
}

type VimClusterVsanFaultDomainSpec struct {
	types.DynamicData

	Hosts []types.ManagedObjectReference `xml:"hosts"`
	Name  string                         `xml:"name"`
}

func init() {
	types.Add("vsan:VimClusterVsanFaultDomainSpec", reflect.TypeOf((*VimClusterVsanFaultDomainSpec)(nil)).Elem())
}

type VimClusterVsanFaultDomainsConfigSpec struct {
	types.DynamicData

	FaultDomains []VimClusterVsanFaultDomainSpec `xml:"faultDomains"`
	Witness      *VimClusterVsanWitnessSpec      `xml:"witness,omitempty"`
}

func init() {
	types.Add("vsan:VimClusterVsanFaultDomainsConfigSpec", reflect.TypeOf((*VimClusterVsanFaultDomainsConfigSpec)(nil)).Elem())
}

type VimClusterVsanHostDiskMapping struct {
	types.DynamicData

	Host          types.ManagedObjectReference `xml:"host"`
	CacheDisks    []types.HostScsiDisk         `xml:"cacheDisks,omitempty"`
	CapacityDisks []types.HostScsiDisk         `xml:"capacityDisks"`
	Type          string                       `xml:"type"`
}

func init() {
	types.Add("vsan:VimClusterVsanHostDiskMapping", reflect.TypeOf((*VimClusterVsanHostDiskMapping)(nil)).Elem())
}

type VimClusterVsanWitnessSpec struct {
	types.DynamicData

	Host                     types.ManagedObjectReference `xml:"host"`
	PreferredFaultDomainName string                       `xml:"preferredFaultDomainName"`
	DiskMapping              *VsanHostDiskMapping         `xml:"diskMapping,omitempty"`
}

func init() {
	types.Add("vsan:VimClusterVsanWitnessSpec", reflect.TypeOf((*VimClusterVsanWitnessSpec)(nil)).Elem())
}

type VimHostVSANStretchedClusterHostCapability struct {
	types.DynamicData

	FeatureVersion string `xml:"featureVersion"`
}

func init() {
	types.Add("vsan:VimHostVSANStretchedClusterHostCapability", reflect.TypeOf((*VimHostVSANStretchedClusterHostCapability)(nil)).Elem())
}

type VimVsanDataEfficiencyCapacityState struct {
	types.DynamicData

	LogicalCapacity      int64 `xml:"logicalCapacity,omitempty"`
	LogicalCapacityUsed  int64 `xml:"logicalCapacityUsed,omitempty"`
	PhysicalCapacity     int64 `xml:"physicalCapacity,omitempty"`
	PhysicalCapacityUsed int64 `xml:"physicalCapacityUsed,omitempty"`
	DedupMetadataSize    int64 `xml:"dedupMetadataSize,omitempty"`
}

func init() {
	types.Add("vsan:VimVsanDataEfficiencyCapacityState", reflect.TypeOf((*VimVsanDataEfficiencyCapacityState)(nil)).Elem())
}

type VimVsanHostDiskMapInfoEx struct {
	types.DynamicData

	Mapping           VsanHostDiskMapping       `xml:"mapping"`
	IsMounted         bool                      `xml:"isMounted"`
	UnlockedEncrypted *bool                     `xml:"unlockedEncrypted"`
	IsAllFlash        bool                      `xml:"isAllFlash"`
	IsDataEfficiency  *bool                     `xml:"isDataEfficiency"`
	EncryptionInfo    *VsanDataEncryptionConfig `xml:"encryptionInfo,omitempty"`
}

func init() {
	types.Add("vsan:VimVsanHostDiskMapInfoEx", reflect.TypeOf((*VimVsanHostDiskMapInfoEx)(nil)).Elem())
}

type VimVsanHostDiskMappingCreationSpec struct {
	types.DynamicData

	Host          types.ManagedObjectReference `xml:"host"`
	CacheDisks    []types.HostScsiDisk         `xml:"cacheDisks,omitempty"`
	CapacityDisks []types.HostScsiDisk         `xml:"capacityDisks"`
	CreationType  string                       `xml:"creationType"`
}

func init() {
	types.Add("vsan:VimVsanHostDiskMappingCreationSpec", reflect.TypeOf((*VimVsanHostDiskMappingCreationSpec)(nil)).Elem())
}

type VimVsanHostVsanDiskManagementSystemCapability struct {
	types.DynamicData

	Version string `xml:"version"`
}

func init() {
	types.Add("vsan:VimVsanHostVsanDiskManagementSystemCapability", reflect.TypeOf((*VimVsanHostVsanDiskManagementSystemCapability)(nil)).Elem())
}

type VimVsanHostVsanHostCapability struct {
	types.DynamicData

	Host        types.ManagedObjectReference `xml:"host"`
	IsSupported bool                         `xml:"isSupported"`
	IsLicensed  bool                         `xml:"isLicensed"`
}

func init() {
	types.Add("vsan:VimVsanHostVsanHostCapability", reflect.TypeOf((*VimVsanHostVsanHostCapability)(nil)).Elem())
}

type VimVsanReconfigSpec struct {
	types.SDDCBase

	VsanClusterConfig      BaseVsanClusterConfigInfo             `xml:"vsanClusterConfig,omitempty,typeattr"`
	DataEfficiencyConfig   *VsanDataEfficiencyConfig             `xml:"dataEfficiencyConfig,omitempty"`
	DiskMappingSpec        *VimClusterVsanDiskMappingsConfigSpec `xml:"diskMappingSpec,omitempty"`
	FaultDomainsSpec       *VimClusterVsanFaultDomainsConfigSpec `xml:"faultDomainsSpec,omitempty"`
	Modify                 bool                                  `xml:"modify"`
	AllowReducedRedundancy *bool                                 `xml:"allowReducedRedundancy"`
	ResyncIopsLimitConfig  *ResyncIopsInfo                       `xml:"resyncIopsLimitConfig,omitempty"`
	IscsiSpec              *VsanIscsiTargetServiceSpec           `xml:"iscsiSpec,omitempty"`
	DataEncryptionConfig   *VsanDataEncryptionConfig             `xml:"dataEncryptionConfig,omitempty"`
	ExtendedConfig         *VsanExtendedConfig                   `xml:"extendedConfig,omitempty"`
	DatastoreConfig        *VsanDatastoreConfig                  `xml:"datastoreConfig,omitempty"`
	PerfsvcConfig          *VsanPerfsvcConfig                    `xml:"perfsvcConfig,omitempty"`
	UnmapConfig            *VsanUnmapConfig                      `xml:"unmapConfig,omitempty"`
	VumConfig              *VsanVumConfig                        `xml:"vumConfig,omitempty"`
	MetricsConfig          *VsanMetricsConfig                    `xml:"metricsConfig,omitempty"`
	FileServiceConfig      *VsanFileServiceConfig                `xml:"fileServiceConfig,omitempty"`
}

func init() {
	types.Add("vsan:VimVsanReconfigSpec", reflect.TypeOf((*VimVsanReconfigSpec)(nil)).Elem())
}

type VosQueryVsanObjectInformation VosQueryVsanObjectInformationRequestType

func init() {
	types.Add("vsan:VosQueryVsanObjectInformation", reflect.TypeOf((*VosQueryVsanObjectInformation)(nil)).Elem())
}

type VosQueryVsanObjectInformationRequestType struct {
	This                 types.ManagedObjectReference  `xml:"_this"`
	Cluster              *types.ManagedObjectReference `xml:"cluster,omitempty"`
	VsanObjectQuerySpecs []VsanObjectQuerySpec         `xml:"vsanObjectQuerySpecs"`
}

func init() {
	types.Add("vsan:VosQueryVsanObjectInformationRequestType", reflect.TypeOf((*VosQueryVsanObjectInformationRequestType)(nil)).Elem())
}

type VosQueryVsanObjectInformationResponse struct {
	Returnval []VsanObjectInformation `xml:"returnval,omitempty"`
}

type VosSetVsanObjectPolicy VosSetVsanObjectPolicyRequestType

func init() {
	types.Add("vsan:VosSetVsanObjectPolicy", reflect.TypeOf((*VosSetVsanObjectPolicy)(nil)).Elem())
}

type VosSetVsanObjectPolicyRequestType struct {
	This           types.ManagedObjectReference        `xml:"_this"`
	Cluster        *types.ManagedObjectReference       `xml:"cluster,omitempty"`
	VsanObjectUuid string                              `xml:"vsanObjectUuid"`
	Profile        types.BaseVirtualMachineProfileSpec `xml:"profile,omitempty,typeattr"`
}

func init() {
	types.Add("vsan:VosSetVsanObjectPolicyRequestType", reflect.TypeOf((*VosSetVsanObjectPolicyRequestType)(nil)).Elem())
}

type VosSetVsanObjectPolicyResponse struct {
	Returnval bool `xml:"returnval"`
}

type VsanAttachToSrOperation struct {
	types.DynamicData

	Task      *types.ManagedObjectReference `xml:"task,omitempty"`
	Success   *bool                         `xml:"success"`
	Timestamp *time.Time                    `xml:"timestamp"`
	SrNumber  string                        `xml:"srNumber"`
}

func init() {
	types.Add("vsan:VsanAttachToSrOperation", reflect.TypeOf((*VsanAttachToSrOperation)(nil)).Elem())
}

type VsanAttachVsanSupportBundleToSr VsanAttachVsanSupportBundleToSrRequestType

func init() {
	types.Add("vsan:VsanAttachVsanSupportBundleToSr", reflect.TypeOf((*VsanAttachVsanSupportBundleToSr)(nil)).Elem())
}

type VsanAttachVsanSupportBundleToSrRequestType struct {
	This     types.ManagedObjectReference `xml:"_this"`
	Cluster  types.ManagedObjectReference `xml:"cluster"`
	SrNumber string                       `xml:"srNumber"`
}

func init() {
	types.Add("vsan:VsanAttachVsanSupportBundleToSrRequestType", reflect.TypeOf((*VsanAttachVsanSupportBundleToSrRequestType)(nil)).Elem())
}

type VsanAttachVsanSupportBundleToSrResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanBasicDeviceInfo struct {
	types.DynamicData

	DeviceName string `xml:"deviceName"`
	PciId      string `xml:"pciId,omitempty"`
	FwVersion  string `xml:"fwVersion,omitempty"`
}

func init() {
	types.Add("vsan:VsanBasicDeviceInfo", reflect.TypeOf((*VsanBasicDeviceInfo)(nil)).Elem())
}

type VsanBrokenDiskChainIssue struct {
	VsanUpgradeSystemPreflightCheckIssue

	Uuids []string `xml:"uuids"`
}

func init() {
	types.Add("vsan:VsanBrokenDiskChainIssue", reflect.TypeOf((*VsanBrokenDiskChainIssue)(nil)).Elem())
}

type VsanBurnInTest struct {
	types.DynamicData

	Testname string `xml:"testname"`
	Workload string `xml:"workload,omitempty"`
	Duration int64  `xml:"duration"`
	Result   string `xml:"result"`
}

func init() {
	types.Add("vsan:VsanBurnInTest", reflect.TypeOf((*VsanBurnInTest)(nil)).Elem())
}

type VsanBurnInTestCheckResult struct {
	types.DynamicData

	PassedTests       []VsanBurnInTest `xml:"passedTests,omitempty"`
	NotPerformedTests []VsanBurnInTest `xml:"notPerformedTests,omitempty"`
	FailedTests       []VsanBurnInTest `xml:"failedTests,omitempty"`
}

func init() {
	types.Add("vsan:VsanBurnInTestCheckResult", reflect.TypeOf((*VsanBurnInTestCheckResult)(nil)).Elem())
}

type VsanCapability struct {
	types.DynamicData

	Target       *types.ManagedObjectReference `xml:"target,omitempty"`
	Capabilities []string                      `xml:"capabilities,omitempty"`
	Statuses     []string                      `xml:"statuses,omitempty"`
}

func init() {
	types.Add("vsan:VsanCapability", reflect.TypeOf((*VsanCapability)(nil)).Elem())
}

type VsanCheckClusterClomdLiveness VsanCheckClusterClomdLivenessRequestType

func init() {
	types.Add("vsan:VsanCheckClusterClomdLiveness", reflect.TypeOf((*VsanCheckClusterClomdLiveness)(nil)).Elem())
}

type VsanCheckClusterClomdLivenessRequestType struct {
	This            types.ManagedObjectReference `xml:"_this"`
	Hosts           []string                     `xml:"hosts"`
	EsxRootPassword string                       `xml:"esxRootPassword"`
}

func init() {
	types.Add("vsan:VsanCheckClusterClomdLivenessRequestType", reflect.TypeOf((*VsanCheckClusterClomdLivenessRequestType)(nil)).Elem())
}

type VsanCheckClusterClomdLivenessResponse struct {
	Returnval VsanClusterClomdLivenessResult `xml:"returnval"`
}

type VsanCloudHealthStatus struct {
	types.DynamicData

	CollectorRunning     *bool  `xml:"collectorRunning"`
	LastSentTimestamp    string `xml:"lastSentTimestamp,omitempty"`
	InternetConnectivity *bool  `xml:"internetConnectivity"`
}

func init() {
	types.Add("vsan:VsanCloudHealthStatus", reflect.TypeOf((*VsanCloudHealthStatus)(nil)).Elem())
}

type VsanClusterAdvCfgSyncHostResult struct {
	types.DynamicData

	Hostname  string `xml:"hostname"`
	Value     string `xml:"value"`
	IsDefault *bool  `xml:"isDefault"`
}

func init() {
	types.Add("vsan:VsanClusterAdvCfgSyncHostResult", reflect.TypeOf((*VsanClusterAdvCfgSyncHostResult)(nil)).Elem())
}

type VsanClusterAdvCfgSyncResult struct {
	types.DynamicData

	InSync     bool                              `xml:"inSync"`
	Name       string                            `xml:"name"`
	HostValues []VsanClusterAdvCfgSyncHostResult `xml:"hostValues,omitempty"`
}

func init() {
	types.Add("vsan:VsanClusterAdvCfgSyncResult", reflect.TypeOf((*VsanClusterAdvCfgSyncResult)(nil)).Elem())
}

type VsanClusterBalancePerDiskInfo struct {
	types.DynamicData

	Uuid                   string `xml:"uuid,omitempty"`
	Fullness               int64  `xml:"fullness"`
	Variance               int64  `xml:"variance"`
	FullnessAboveThreshold int64  `xml:"fullnessAboveThreshold"`
	DataToMoveB            int64  `xml:"dataToMoveB"`
}

func init() {
	types.Add("vsan:VsanClusterBalancePerDiskInfo", reflect.TypeOf((*VsanClusterBalancePerDiskInfo)(nil)).Elem())
}

type VsanClusterBalanceSummary struct {
	types.DynamicData

	VarianceThreshold int64                           `xml:"varianceThreshold"`
	Disks             []VsanClusterBalancePerDiskInfo `xml:"disks,omitempty"`
}

func init() {
	types.Add("vsan:VsanClusterBalanceSummary", reflect.TypeOf((*VsanClusterBalanceSummary)(nil)).Elem())
}

type VsanClusterBurnInTestResultList struct {
	types.DynamicData

	Items []VsanBurnInTest `xml:"items,omitempty"`
	Hosts []string         `xml:"hosts,omitempty"`
}

func init() {
	types.Add("vsan:VsanClusterBurnInTestResultList", reflect.TypeOf((*VsanClusterBurnInTestResultList)(nil)).Elem())
}

type VsanClusterClomdLivenessResult struct {
	types.DynamicData

	ClomdLivenessResult []VsanHostClomdLivenessResult `xml:"clomdLivenessResult,omitempty"`
	IssueFound          bool                          `xml:"issueFound"`
}

func init() {
	types.Add("vsan:VsanClusterClomdLivenessResult", reflect.TypeOf((*VsanClusterClomdLivenessResult)(nil)).Elem())
}

type VsanClusterConfig struct {
	types.DynamicData

	Config      BaseVsanClusterConfigInfo `xml:"config,typeattr"`
	Name        string                    `xml:"name"`
	Hosts       []string                  `xml:"hosts,omitempty"`
	ToBeDeleted *types.HostApplyProfile   `xml:"toBeDeleted,omitempty"`
}

func init() {
	types.Add("vsan:VsanClusterConfig", reflect.TypeOf((*VsanClusterConfig)(nil)).Elem())
}

type VsanClusterCreateFsDomain VsanClusterCreateFsDomainRequestType

func init() {
	types.Add("vsan:VsanClusterCreateFsDomain", reflect.TypeOf((*VsanClusterCreateFsDomain)(nil)).Elem())
}

type VsanClusterCreateFsDomainRequestType struct {
	This         types.ManagedObjectReference  `xml:"_this"`
	DomainConfig VsanFileServiceDomainConfig   `xml:"domainConfig"`
	Cluster      *types.ManagedObjectReference `xml:"cluster,omitempty"`
}

func init() {
	types.Add("vsan:VsanClusterCreateFsDomainRequestType", reflect.TypeOf((*VsanClusterCreateFsDomainRequestType)(nil)).Elem())
}

type VsanClusterCreateFsDomainResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanClusterCreateVmHealthTestResult struct {
	types.DynamicData

	ClusterResult VsanClusterProactiveTestResult     `xml:"clusterResult"`
	HostResults   []VsanHostCreateVmHealthTestResult `xml:"hostResults,omitempty"`
}

func init() {
	types.Add("vsan:VsanClusterCreateVmHealthTestResult", reflect.TypeOf((*VsanClusterCreateVmHealthTestResult)(nil)).Elem())
}

type VsanClusterEncryptionHealthSummary struct {
	types.DynamicData

	OverallHealth string                        `xml:"overallHealth,omitempty"`
	ConfigHealth  string                        `xml:"configHealth,omitempty"`
	KmsHealth     string                        `xml:"kmsHealth,omitempty"`
	VcKmsResult   *VsanVcKmipServersHealth      `xml:"vcKmsResult,omitempty"`
	HostResults   []VsanEncryptionHealthSummary `xml:"hostResults,omitempty"`
	AesniHealth   string                        `xml:"aesniHealth,omitempty"`
}

func init() {
	types.Add("vsan:VsanClusterEncryptionHealthSummary", reflect.TypeOf((*VsanClusterEncryptionHealthSummary)(nil)).Elem())
}

type VsanClusterFileServiceHealthSummary struct {
	types.DynamicData

	OverallHealth string                         `xml:"overallHealth,omitempty"`
	HostResults   []VsanFileServiceHealthSummary `xml:"hostResults,omitempty"`
}

func init() {
	types.Add("vsan:VsanClusterFileServiceHealthSummary", reflect.TypeOf((*VsanClusterFileServiceHealthSummary)(nil)).Elem())
}

type VsanClusterGetConfig VsanClusterGetConfigRequestType

func init() {
	types.Add("vsan:VsanClusterGetConfig", reflect.TypeOf((*VsanClusterGetConfig)(nil)).Elem())
}

type VsanClusterGetConfigRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

func init() {
	types.Add("vsan:VsanClusterGetConfigRequestType", reflect.TypeOf((*VsanClusterGetConfigRequestType)(nil)).Elem())
}

type VsanClusterGetConfigResponse struct {
	Returnval VsanConfigInfoEx `xml:"returnval"`
}

type VsanClusterGetHclInfo VsanClusterGetHclInfoRequestType

func init() {
	types.Add("vsan:VsanClusterGetHclInfo", reflect.TypeOf((*VsanClusterGetHclInfo)(nil)).Elem())
}

type VsanClusterGetHclInfoRequestType struct {
	This            types.ManagedObjectReference `xml:"_this"`
	Hosts           []string                     `xml:"hosts"`
	EsxRootPassword string                       `xml:"esxRootPassword"`
}

func init() {
	types.Add("vsan:VsanClusterGetHclInfoRequestType", reflect.TypeOf((*VsanClusterGetHclInfoRequestType)(nil)).Elem())
}

type VsanClusterGetHclInfoResponse struct {
	Returnval VsanClusterHclInfo `xml:"returnval"`
}

type VsanClusterGetRuntimeStats VsanClusterGetRuntimeStatsRequestType

func init() {
	types.Add("vsan:VsanClusterGetRuntimeStats", reflect.TypeOf((*VsanClusterGetRuntimeStats)(nil)).Elem())
}

type VsanClusterGetRuntimeStatsRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
	Stats   []string                     `xml:"stats,omitempty"`
}

func init() {
	types.Add("vsan:VsanClusterGetRuntimeStatsRequestType", reflect.TypeOf((*VsanClusterGetRuntimeStatsRequestType)(nil)).Elem())
}

type VsanClusterGetRuntimeStatsResponse struct {
	Returnval []VsanRuntimeStatsHostMap `xml:"returnval,omitempty"`
}

type VsanClusterHclInfo struct {
	types.DynamicData

	HclDbLastUpdate *time.Time        `xml:"hclDbLastUpdate"`
	HclDbAgeHealth  string            `xml:"hclDbAgeHealth,omitempty"`
	HostResults     []VsanHostHclInfo `xml:"hostResults,omitempty"`
	UpdateItems     []VsanUpdateItem  `xml:"updateItems,omitempty"`
}

func init() {
	types.Add("vsan:VsanClusterHclInfo", reflect.TypeOf((*VsanClusterHclInfo)(nil)).Elem())
}

type VsanClusterHealthAction struct {
	types.DynamicData

	ActionId          string                   `xml:"actionId"`
	ActionLabel       types.LocalizableMessage `xml:"actionLabel"`
	ActionDescription types.LocalizableMessage `xml:"actionDescription"`
	Enabled           bool                     `xml:"enabled"`
}

func init() {
	types.Add("vsan:VsanClusterHealthAction", reflect.TypeOf((*VsanClusterHealthAction)(nil)).Elem())
}

type VsanClusterHealthCheckInfo struct {
	types.DynamicData

	TestId    string `xml:"testId"`
	TestName  string `xml:"testName,omitempty"`
	GroupId   string `xml:"groupId"`
	GroupName string `xml:"groupName,omitempty"`
}

func init() {
	types.Add("vsan:VsanClusterHealthCheckInfo", reflect.TypeOf((*VsanClusterHealthCheckInfo)(nil)).Elem())
}

type VsanClusterHealthConfigs struct {
	types.DynamicData

	EnableVsanTelemetry   *bool                                 `xml:"enableVsanTelemetry"`
	VsanTelemetryInterval int32                                 `xml:"vsanTelemetryInterval,omitempty"`
	VsanTelemetryProxy    *VsanClusterTelemetryProxyConfig      `xml:"vsanTelemetryProxy,omitempty"`
	Configs               []VsanClusterHealthResultKeyValuePair `xml:"configs,omitempty"`
}

func init() {
	types.Add("vsan:VsanClusterHealthConfigs", reflect.TypeOf((*VsanClusterHealthConfigs)(nil)).Elem())
}

type VsanClusterHealthGroup struct {
	types.DynamicData

	GroupId      string                            `xml:"groupId"`
	GroupName    string                            `xml:"groupName"`
	GroupHealth  string                            `xml:"groupHealth"`
	GroupTests   []VsanClusterHealthTest           `xml:"groupTests,omitempty"`
	GroupDetails []BaseVsanClusterHealthResultBase `xml:"groupDetails,omitempty,typeattr"`
}

func init() {
	types.Add("vsan:VsanClusterHealthGroup", reflect.TypeOf((*VsanClusterHealthGroup)(nil)).Elem())
}

type VsanClusterHealthResultBase struct {
	types.DynamicData

	Label string `xml:"label,omitempty"`
}

func init() {
	types.Add("vsan:VsanClusterHealthResultBase", reflect.TypeOf((*VsanClusterHealthResultBase)(nil)).Elem())
}

type VsanClusterHealthResultColumnInfo struct {
	types.DynamicData

	Label string `xml:"label"`
	Type  string `xml:"type"`
}

func init() {
	types.Add("vsan:VsanClusterHealthResultColumnInfo", reflect.TypeOf((*VsanClusterHealthResultColumnInfo)(nil)).Elem())
}

type VsanClusterHealthResultKeyValuePair struct {
	types.DynamicData

	Key   string `xml:"key,omitempty"`
	Value string `xml:"value,omitempty"`
}

func init() {
	types.Add("vsan:VsanClusterHealthResultKeyValuePair", reflect.TypeOf((*VsanClusterHealthResultKeyValuePair)(nil)).Elem())
}

type VsanClusterHealthResultRow struct {
	types.DynamicData

	Values     []string                     `xml:"values"`
	NestedRows []VsanClusterHealthResultRow `xml:"nestedRows,omitempty"`
}

func init() {
	types.Add("vsan:VsanClusterHealthResultRow", reflect.TypeOf((*VsanClusterHealthResultRow)(nil)).Elem())
}

type VsanClusterHealthResultTable struct {
	VsanClusterHealthResultBase

	Columns []VsanClusterHealthResultColumnInfo `xml:"columns,omitempty"`
	Rows    []VsanClusterHealthResultRow        `xml:"rows,omitempty"`
}

func init() {
	types.Add("vsan:VsanClusterHealthResultTable", reflect.TypeOf((*VsanClusterHealthResultTable)(nil)).Elem())
}

type VsanClusterHealthSummary struct {
	types.DynamicData

	ClusterStatus            *VsanClusterHealthSystemStatusResult  `xml:"clusterStatus,omitempty"`
	Timestamp                *time.Time                            `xml:"timestamp"`
	ClusterVersions          *VsanClusterHealthSystemVersionResult `xml:"clusterVersions,omitempty"`
	ObjectHealth             *VsanObjectOverallHealth              `xml:"objectHealth,omitempty"`
	VmHealth                 *VsanClusterVMsHealthOverallResult    `xml:"vmHealth,omitempty"`
	NetworkHealth            *VsanClusterNetworkHealthResult       `xml:"networkHealth,omitempty"`
	LimitHealth              *VsanClusterLimitHealthResult         `xml:"limitHealth,omitempty"`
	AdvCfgSync               []VsanClusterAdvCfgSyncResult         `xml:"advCfgSync,omitempty"`
	CreateVmHealth           []VsanHostCreateVmHealthTestResult    `xml:"createVmHealth,omitempty"`
	PhysicalDisksHealth      []VsanPhysicalDiskHealthSummary       `xml:"physicalDisksHealth,omitempty"`
	EncryptionHealth         *VsanClusterEncryptionHealthSummary   `xml:"encryptionHealth,omitempty"`
	HclInfo                  *VsanClusterHclInfo                   `xml:"hclInfo,omitempty"`
	Groups                   []VsanClusterHealthGroup              `xml:"groups,omitempty"`
	OverallHealth            string                                `xml:"overallHealth"`
	OverallHealthDescription string                                `xml:"overallHealthDescription"`
	ClomdLiveness            *VsanClusterClomdLivenessResult       `xml:"clomdLiveness,omitempty"`
	DiskBalance              *VsanClusterBalanceSummary            `xml:"diskBalance,omitempty"`
	GenericCluster           *VsanGenericClusterBestPracticeHealth `xml:"genericCluster,omitempty"`
	NetworkConfig            *VsanNetworkConfigBestPracticeHealth  `xml:"networkConfig,omitempty"`
	VsanConfig               BaseVsanClusterConfigInfo             `xml:"vsanConfig,omitempty,typeattr"`
	BurnInTest               *VsanBurnInTestCheckResult            `xml:"burnInTest,omitempty"`
	PerfsvcHealth            *VsanPerfsvcHealthResult              `xml:"perfsvcHealth,omitempty"`
	Cluster                  *types.ManagedObjectReference         `xml:"cluster,omitempty"`
	FileServiceHealth        *VsanClusterFileServiceHealthSummary  `xml:"fileServiceHealth,omitempty"`
}

func init() {
	types.Add("vsan:VsanClusterHealthSummary", reflect.TypeOf((*VsanClusterHealthSummary)(nil)).Elem())
}

type VsanClusterHealthSystemObjectsRepairResult struct {
	types.DynamicData

	InRepairingQueueObjects []string                       `xml:"inRepairingQueueObjects,omitempty"`
	FailedRepairObjects     []VsanFailedRepairObjectResult `xml:"failedRepairObjects,omitempty"`
	IssueFound              bool                           `xml:"issueFound"`
}

func init() {
	types.Add("vsan:VsanClusterHealthSystemObjectsRepairResult", reflect.TypeOf((*VsanClusterHealthSystemObjectsRepairResult)(nil)).Elem())
}

type VsanClusterHealthSystemStatusResult struct {
	types.DynamicData

	Status             string                             `xml:"status"`
	GoalState          string                             `xml:"goalState"`
	UntrackedHosts     []string                           `xml:"untrackedHosts,omitempty"`
	TrackedHostsStatus []VsanHostHealthSystemStatusResult `xml:"trackedHostsStatus,omitempty"`
}

func init() {
	types.Add("vsan:VsanClusterHealthSystemStatusResult", reflect.TypeOf((*VsanClusterHealthSystemStatusResult)(nil)).Elem())
}

type VsanClusterHealthSystemVersionResult struct {
	types.DynamicData

	HostResults     []VsanHostHealthSystemVersionResult `xml:"hostResults,omitempty"`
	VcVersion       string                              `xml:"vcVersion,omitempty"`
	IssueFound      bool                                `xml:"issueFound"`
	UpgradePossible *bool                               `xml:"upgradePossible"`
}

func init() {
	types.Add("vsan:VsanClusterHealthSystemVersionResult", reflect.TypeOf((*VsanClusterHealthSystemVersionResult)(nil)).Elem())
}

type VsanClusterHealthTest struct {
	types.DynamicData

	TestId               string                            `xml:"testId,omitempty"`
	TestName             string                            `xml:"testName,omitempty"`
	TestDescription      string                            `xml:"testDescription,omitempty"`
	TestShortDescription string                            `xml:"testShortDescription,omitempty"`
	TestHealthyEntities  int32                             `xml:"testHealthyEntities,omitempty"`
	TestAllEntities      int32                             `xml:"testAllEntities,omitempty"`
	TestHealth           string                            `xml:"testHealth,omitempty"`
	TestDetails          []BaseVsanClusterHealthResultBase `xml:"testDetails,omitempty,typeattr"`
	TestActions          []VsanClusterHealthAction         `xml:"testActions,omitempty"`
}

func init() {
	types.Add("vsan:VsanClusterHealthTest", reflect.TypeOf((*VsanClusterHealthTest)(nil)).Elem())
}

type VsanClusterHostVmknicMapping struct {
	types.DynamicData

	Host   string `xml:"host"`
	Vmknic string `xml:"vmknic"`
}

func init() {
	types.Add("vsan:VsanClusterHostVmknicMapping", reflect.TypeOf((*VsanClusterHostVmknicMapping)(nil)).Elem())
}

type VsanClusterLimitHealthResult struct {
	types.DynamicData

	IssueFound              bool                                  `xml:"issueFound"`
	ComponentLimitHealth    string                                `xml:"componentLimitHealth"`
	DiskFreeSpaceHealth     string                                `xml:"diskFreeSpaceHealth"`
	RcFreeReservationHealth string                                `xml:"rcFreeReservationHealth"`
	HostResults             []VsanLimitHealthResult               `xml:"hostResults,omitempty"`
	WhatifHostFailures      []VsanClusterWhatifHostFailuresResult `xml:"whatifHostFailures,omitempty"`
	HostsCommFailure        []string                              `xml:"hostsCommFailure,omitempty"`
}

func init() {
	types.Add("vsan:VsanClusterLimitHealthResult", reflect.TypeOf((*VsanClusterLimitHealthResult)(nil)).Elem())
}

type VsanClusterNetworkHealthResult struct {
	types.DynamicData

	HostResults                []VsanNetworkHealthResult         `xml:"hostResults,omitempty"`
	IssueFound                 *bool                             `xml:"issueFound"`
	VsanVmknicPresent          *bool                             `xml:"vsanVmknicPresent"`
	MatchingMulticastConfig    *bool                             `xml:"matchingMulticastConfig"`
	MatchingIpSubnets          *bool                             `xml:"matchingIpSubnets"`
	PingTestSuccess            *bool                             `xml:"pingTestSuccess"`
	LargePingTestSuccess       *bool                             `xml:"largePingTestSuccess"`
	HostLatencyCheckSuccess    *bool                             `xml:"hostLatencyCheckSuccess"`
	PotentialMulticastIssue    *bool                             `xml:"potentialMulticastIssue"`
	OtherHostsInVsanCluster    []string                          `xml:"otherHostsInVsanCluster,omitempty"`
	Partitions                 []VsanClusterNetworkPartitionInfo `xml:"partitions,omitempty"`
	HostsWithVsanDisabled      []string                          `xml:"hostsWithVsanDisabled,omitempty"`
	HostsDisconnected          []string                          `xml:"hostsDisconnected,omitempty"`
	HostsCommFailure           []string                          `xml:"hostsCommFailure,omitempty"`
	HostsInEsxMaintenanceMode  []string                          `xml:"hostsInEsxMaintenanceMode,omitempty"`
	HostsInVsanMaintenanceMode []string                          `xml:"hostsInVsanMaintenanceMode,omitempty"`
	InfoAboutUnexpectedHosts   []VsanQueryResultHostInfo         `xml:"infoAboutUnexpectedHosts,omitempty"`
	ClusterInUnicastMode       *bool                             `xml:"clusterInUnicastMode"`
}

func init() {
	types.Add("vsan:VsanClusterNetworkHealthResult", reflect.TypeOf((*VsanClusterNetworkHealthResult)(nil)).Elem())
}

type VsanClusterNetworkLoadTestResult struct {
	types.DynamicData

	ClusterResult VsanClusterProactiveTestResult `xml:"clusterResult"`
	HostResults   []VsanNetworkLoadTestResult    `xml:"hostResults,omitempty"`
}

func init() {
	types.Add("vsan:VsanClusterNetworkLoadTestResult", reflect.TypeOf((*VsanClusterNetworkLoadTestResult)(nil)).Elem())
}

type VsanClusterNetworkPartitionInfo struct {
	types.DynamicData

	Hosts            []string `xml:"hosts,omitempty"`
	PartitionUnknown *bool    `xml:"partitionUnknown"`
}

func init() {
	types.Add("vsan:VsanClusterNetworkPartitionInfo", reflect.TypeOf((*VsanClusterNetworkPartitionInfo)(nil)).Elem())
}

type VsanClusterObjectExtAttrs struct {
	types.DynamicData

	Uuid          string `xml:"uuid"`
	ObjectType    string `xml:"objectType,omitempty"`
	ObjectPath    string `xml:"objectPath,omitempty"`
	GroupUuid     string `xml:"groupUuid,omitempty"`
	DirectoryName string `xml:"directoryName,omitempty"`
}

func init() {
	types.Add("vsan:VsanClusterObjectExtAttrs", reflect.TypeOf((*VsanClusterObjectExtAttrs)(nil)).Elem())
}

type VsanClusterProactiveTestResult struct {
	types.DynamicData

	OverallStatus            string                 `xml:"overallStatus"`
	OverallStatusDescription string                 `xml:"overallStatusDescription"`
	Timestamp                time.Time              `xml:"timestamp"`
	HealthTest               *VsanClusterHealthTest `xml:"healthTest,omitempty"`
}

func init() {
	types.Add("vsan:VsanClusterProactiveTestResult", reflect.TypeOf((*VsanClusterProactiveTestResult)(nil)).Elem())
}

type VsanClusterQueryFileServiceHealthSummary VsanClusterQueryFileServiceHealthSummaryRequestType

func init() {
	types.Add("vsan:VsanClusterQueryFileServiceHealthSummary", reflect.TypeOf((*VsanClusterQueryFileServiceHealthSummary)(nil)).Elem())
}

type VsanClusterQueryFileServiceHealthSummaryRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

func init() {
	types.Add("vsan:VsanClusterQueryFileServiceHealthSummaryRequestType", reflect.TypeOf((*VsanClusterQueryFileServiceHealthSummaryRequestType)(nil)).Elem())
}

type VsanClusterQueryFileServiceHealthSummaryResponse struct {
	Returnval *VsanClusterFileServiceHealthSummary `xml:"returnval,omitempty"`
}

type VsanClusterQueryFileShares VsanClusterQueryFileSharesRequestType

func init() {
	types.Add("vsan:VsanClusterQueryFileShares", reflect.TypeOf((*VsanClusterQueryFileShares)(nil)).Elem())
}

type VsanClusterQueryFileSharesRequestType struct {
	This      types.ManagedObjectReference  `xml:"_this"`
	QuerySpec VsanFileShareQuerySpec        `xml:"querySpec"`
	Cluster   *types.ManagedObjectReference `xml:"cluster,omitempty"`
}

func init() {
	types.Add("vsan:VsanClusterQueryFileSharesRequestType", reflect.TypeOf((*VsanClusterQueryFileSharesRequestType)(nil)).Elem())
}

type VsanClusterQueryFileSharesResponse struct {
	Returnval *FileShareQueryResult `xml:"returnval,omitempty"`
}

type VsanClusterQueryFsDomains VsanClusterQueryFsDomainsRequestType

func init() {
	types.Add("vsan:VsanClusterQueryFsDomains", reflect.TypeOf((*VsanClusterQueryFsDomains)(nil)).Elem())
}

type VsanClusterQueryFsDomainsRequestType struct {
	This      types.ManagedObjectReference    `xml:"_this"`
	QuerySpec *VsanFileServiceDomainQuerySpec `xml:"querySpec,omitempty"`
	Cluster   *types.ManagedObjectReference   `xml:"cluster,omitempty"`
}

func init() {
	types.Add("vsan:VsanClusterQueryFsDomainsRequestType", reflect.TypeOf((*VsanClusterQueryFsDomainsRequestType)(nil)).Elem())
}

type VsanClusterQueryFsDomainsResponse struct {
	Returnval []VsanFileServiceDomain `xml:"returnval,omitempty"`
}

type VsanClusterReconfig VsanClusterReconfigRequestType

func init() {
	types.Add("vsan:VsanClusterReconfig", reflect.TypeOf((*VsanClusterReconfig)(nil)).Elem())
}

type VsanClusterReconfigRequestType struct {
	This             types.ManagedObjectReference `xml:"_this"`
	Cluster          types.ManagedObjectReference `xml:"cluster"`
	VsanReconfigSpec VimVsanReconfigSpec          `xml:"vsanReconfigSpec"`
}

func init() {
	types.Add("vsan:VsanClusterReconfigRequestType", reflect.TypeOf((*VsanClusterReconfigRequestType)(nil)).Elem())
}

type VsanClusterReconfigResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanClusterReconfigureFsDomain VsanClusterReconfigureFsDomainRequestType

func init() {
	types.Add("vsan:VsanClusterReconfigureFsDomain", reflect.TypeOf((*VsanClusterReconfigureFsDomain)(nil)).Elem())
}

type VsanClusterReconfigureFsDomainRequestType struct {
	This         types.ManagedObjectReference  `xml:"_this"`
	DomainUuid   string                        `xml:"domainUuid"`
	DomainConfig VsanFileServiceDomainConfig   `xml:"domainConfig"`
	Cluster      *types.ManagedObjectReference `xml:"cluster,omitempty"`
}

func init() {
	types.Add("vsan:VsanClusterReconfigureFsDomainRequestType", reflect.TypeOf((*VsanClusterReconfigureFsDomainRequestType)(nil)).Elem())
}

type VsanClusterReconfigureFsDomainResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanClusterRemoveFsDomain VsanClusterRemoveFsDomainRequestType

func init() {
	types.Add("vsan:VsanClusterRemoveFsDomain", reflect.TypeOf((*VsanClusterRemoveFsDomain)(nil)).Elem())
}

type VsanClusterRemoveFsDomainRequestType struct {
	This       types.ManagedObjectReference  `xml:"_this"`
	DomainUuid string                        `xml:"domainUuid"`
	Cluster    *types.ManagedObjectReference `xml:"cluster,omitempty"`
}

func init() {
	types.Add("vsan:VsanClusterRemoveFsDomainRequestType", reflect.TypeOf((*VsanClusterRemoveFsDomainRequestType)(nil)).Elem())
}

type VsanClusterRemoveFsDomainResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanClusterRemoveShare VsanClusterRemoveShareRequestType

func init() {
	types.Add("vsan:VsanClusterRemoveShare", reflect.TypeOf((*VsanClusterRemoveShare)(nil)).Elem())
}

type VsanClusterRemoveShareRequestType struct {
	This      types.ManagedObjectReference  `xml:"_this"`
	ShareUuid string                        `xml:"shareUuid"`
	Cluster   *types.ManagedObjectReference `xml:"cluster,omitempty"`
	Force     *bool                         `xml:"force"`
}

func init() {
	types.Add("vsan:VsanClusterRemoveShareRequestType", reflect.TypeOf((*VsanClusterRemoveShareRequestType)(nil)).Elem())
}

type VsanClusterRemoveShareResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanClusterTelemetryProxyConfig struct {
	types.DynamicData

	Host           string `xml:"host,omitempty"`
	Port           int32  `xml:"port,omitempty"`
	User           string `xml:"user,omitempty"`
	Password       string `xml:"password,omitempty"`
	AutoDiscovered *bool  `xml:"autoDiscovered"`
}

func init() {
	types.Add("vsan:VsanClusterTelemetryProxyConfig", reflect.TypeOf((*VsanClusterTelemetryProxyConfig)(nil)).Elem())
}

type VsanClusterVMsHealthOverallResult struct {
	types.DynamicData

	HealthStateList    []VsanClusterVMsHealthSummaryResult `xml:"healthStateList,omitempty"`
	OverallHealthState string                              `xml:"overallHealthState,omitempty"`
}

func init() {
	types.Add("vsan:VsanClusterVMsHealthOverallResult", reflect.TypeOf((*VsanClusterVMsHealthOverallResult)(nil)).Elem())
}

type VsanClusterVMsHealthSummaryResult struct {
	types.DynamicData

	NumVMs          int32    `xml:"numVMs"`
	State           string   `xml:"state,omitempty"`
	Health          string   `xml:"health"`
	VmInstanceUuids []string `xml:"vmInstanceUuids,omitempty"`
}

func init() {
	types.Add("vsan:VsanClusterVMsHealthSummaryResult", reflect.TypeOf((*VsanClusterVMsHealthSummaryResult)(nil)).Elem())
}

type VsanClusterVmdkLoadTestResult struct {
	types.DynamicData

	Task          *types.ManagedObjectReference   `xml:"task,omitempty"`
	ClusterResult *VsanClusterProactiveTestResult `xml:"clusterResult,omitempty"`
	HostResults   []VsanHostVmdkLoadTestResult    `xml:"hostResults,omitempty"`
}

func init() {
	types.Add("vsan:VsanClusterVmdkLoadTestResult", reflect.TypeOf((*VsanClusterVmdkLoadTestResult)(nil)).Elem())
}

type VsanClusterWhatifHostFailuresResult struct {
	types.DynamicData

	NumFailures             int64  `xml:"numFailures"`
	TotalUsedCapacityB      int64  `xml:"totalUsedCapacityB"`
	TotalCapacityB          int64  `xml:"totalCapacityB"`
	TotalRcReservationB     int64  `xml:"totalRcReservationB"`
	TotalRcSizeB            int64  `xml:"totalRcSizeB"`
	UsedComponents          int64  `xml:"usedComponents"`
	TotalComponents         int64  `xml:"totalComponents"`
	ComponentLimitHealth    string `xml:"componentLimitHealth,omitempty"`
	DiskFreeSpaceHealth     string `xml:"diskFreeSpaceHealth,omitempty"`
	RcFreeReservationHealth string `xml:"rcFreeReservationHealth,omitempty"`
}

func init() {
	types.Add("vsan:VsanClusterWhatifHostFailuresResult", reflect.TypeOf((*VsanClusterWhatifHostFailuresResult)(nil)).Elem())
}

type VsanComparator struct {
	types.DynamicData
}

func init() {
	types.Add("vsan:VsanComparator", reflect.TypeOf((*VsanComparator)(nil)).Elem())
}

type VsanCompleteMigrateVmsToVds VsanCompleteMigrateVmsToVdsRequestType

func init() {
	types.Add("vsan:VsanCompleteMigrateVmsToVds", reflect.TypeOf((*VsanCompleteMigrateVmsToVds)(nil)).Elem())
}

type VsanCompleteMigrateVmsToVdsRequestType struct {
	This     types.ManagedObjectReference `xml:"_this"`
	JobId    string                       `xml:"jobId"`
	NewState string                       `xml:"newState"`
}

func init() {
	types.Add("vsan:VsanCompleteMigrateVmsToVdsRequestType", reflect.TypeOf((*VsanCompleteMigrateVmsToVdsRequestType)(nil)).Elem())
}

type VsanCompleteMigrateVmsToVdsResponse struct {
}

type VsanCompliantDriver struct {
	types.DynamicData

	DriverName    string `xml:"driverName"`
	DriverVersion string `xml:"driverVersion"`
}

func init() {
	types.Add("vsan:VsanCompliantDriver", reflect.TypeOf((*VsanCompliantDriver)(nil)).Elem())
}

type VsanCompliantFirmware struct {
	types.DynamicData

	FirmwareVersion  string                `xml:"firmwareVersion"`
	CompliantDrivers []VsanCompliantDriver `xml:"compliantDrivers"`
}

func init() {
	types.Add("vsan:VsanCompliantFirmware", reflect.TypeOf((*VsanCompliantFirmware)(nil)).Elem())
}

type VsanCompositeConstraint struct {
	VsanResourceConstraint

	NestedConstraints []BaseVsanResourceConstraint `xml:"nestedConstraints,omitempty,typeattr"`
	Conjoiner         string                       `xml:"conjoiner,omitempty"`
}

func init() {
	types.Add("vsan:VsanCompositeConstraint", reflect.TypeOf((*VsanCompositeConstraint)(nil)).Elem())
}

type VsanConfigBaseIssue struct {
	types.DynamicData
}

func init() {
	types.Add("vsan:VsanConfigBaseIssue", reflect.TypeOf((*VsanConfigBaseIssue)(nil)).Elem())
}

type VsanConfigCheckResult struct {
	types.DynamicData

	VsanEnabled bool                      `xml:"vsanEnabled"`
	Issues      []BaseVsanConfigBaseIssue `xml:"issues,omitempty,typeattr"`
}

func init() {
	types.Add("vsan:VsanConfigCheckResult", reflect.TypeOf((*VsanConfigCheckResult)(nil)).Elem())
}

type VsanConfigGeneration struct {
	types.DynamicData

	VcUuid  string `xml:"vcUuid"`
	GenNum  int64  `xml:"genNum"`
	GenTime int64  `xml:"genTime"`
}

func init() {
	types.Add("vsan:VsanConfigGeneration", reflect.TypeOf((*VsanConfigGeneration)(nil)).Elem())
}

type VsanConfigInfoEx struct {
	VsanClusterConfigInfo

	DataEfficiencyConfig  *VsanDataEfficiencyConfig        `xml:"dataEfficiencyConfig,omitempty"`
	ResyncIopsLimitConfig *ResyncIopsInfo                  `xml:"resyncIopsLimitConfig,omitempty"`
	IscsiConfig           BaseVsanIscsiTargetServiceConfig `xml:"iscsiConfig,omitempty,typeattr"`
	DataEncryptionConfig  *VsanDataEncryptionConfig        `xml:"dataEncryptionConfig,omitempty"`
	ExtendedConfig        *VsanExtendedConfig              `xml:"extendedConfig,omitempty"`
	DatastoreConfig       *VsanDatastoreConfig             `xml:"datastoreConfig,omitempty"`
	PerfsvcConfig         *VsanPerfsvcConfig               `xml:"perfsvcConfig,omitempty"`
	UnmapConfig           *VsanUnmapConfig                 `xml:"unmapConfig,omitempty"`
	VumConfig             *VsanVumConfig                   `xml:"vumConfig,omitempty"`
	FileServiceConfig     *VsanFileServiceConfig           `xml:"fileServiceConfig,omitempty"`
	MetricsConfig         *VsanMetricsConfig               `xml:"metricsConfig,omitempty"`
}

func init() {
	types.Add("vsan:VsanConfigInfoEx", reflect.TypeOf((*VsanConfigInfoEx)(nil)).Elem())
}

type VsanConfigNotAllDisksClaimedIssue struct {
	VsanConfigBaseIssue

	Host  types.ManagedObjectReference `xml:"host"`
	Disks []string                     `xml:"disks"`
}

func init() {
	types.Add("vsan:VsanConfigNotAllDisksClaimedIssue", reflect.TypeOf((*VsanConfigNotAllDisksClaimedIssue)(nil)).Elem())
}

type VsanCreateFileShare VsanCreateFileShareRequestType

func init() {
	types.Add("vsan:VsanCreateFileShare", reflect.TypeOf((*VsanCreateFileShare)(nil)).Elem())
}

type VsanCreateFileShareRequestType struct {
	This    types.ManagedObjectReference  `xml:"_this"`
	Config  VsanFileShareConfig           `xml:"config"`
	Cluster *types.ManagedObjectReference `xml:"cluster,omitempty"`
}

func init() {
	types.Add("vsan:VsanCreateFileShareRequestType", reflect.TypeOf((*VsanCreateFileShareRequestType)(nil)).Elem())
}

type VsanCreateFileShareResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanDaemonHealth struct {
	types.DynamicData

	Name  string                      `xml:"name"`
	Alive bool                        `xml:"alive"`
	Error *types.LocalizedMethodFault `xml:"error,omitempty"`
}

func init() {
	types.Add("vsan:VsanDaemonHealth", reflect.TypeOf((*VsanDaemonHealth)(nil)).Elem())
}

type VsanDataEfficiencyConfig struct {
	types.DynamicData

	DedupEnabled       bool  `xml:"dedupEnabled"`
	CompressionEnabled *bool `xml:"compressionEnabled"`
}

func init() {
	types.Add("vsan:VsanDataEfficiencyConfig", reflect.TypeOf((*VsanDataEfficiencyConfig)(nil)).Elem())
}

type VsanDataEncryptionConfig struct {
	types.DynamicData

	EncryptionEnabled   bool                 `xml:"encryptionEnabled"`
	KmsProviderId       *types.KeyProviderId `xml:"kmsProviderId,omitempty"`
	KekId               string               `xml:"kekId,omitempty"`
	HostKeyId           string               `xml:"hostKeyId,omitempty"`
	DekGenerationId     int64                `xml:"dekGenerationId,omitempty"`
	Changing            *bool                `xml:"changing"`
	EraseDisksBeforeUse *bool                `xml:"eraseDisksBeforeUse"`
}

func init() {
	types.Add("vsan:VsanDataEncryptionConfig", reflect.TypeOf((*VsanDataEncryptionConfig)(nil)).Elem())
}

type VsanDataObfuscationRule struct {
	types.DynamicData
}

func init() {
	types.Add("vsan:VsanDataObfuscationRule", reflect.TypeOf((*VsanDataObfuscationRule)(nil)).Elem())
}

type VsanDatastoreConfig struct {
	types.DynamicData

	Datastores []VsanDatastoreSpec `xml:"datastores,omitempty"`
}

func init() {
	types.Add("vsan:VsanDatastoreConfig", reflect.TypeOf((*VsanDatastoreConfig)(nil)).Elem())
}

type VsanDatastoreSpec struct {
	types.DynamicData

	Uuid string `xml:"uuid"`
	Name string `xml:"name"`
}

func init() {
	types.Add("vsan:VsanDatastoreSpec", reflect.TypeOf((*VsanDatastoreSpec)(nil)).Elem())
}

type VsanDeleteObjectsRequestType struct {
	This     types.ManagedObjectReference  `xml:"_this"`
	Cluster  *types.ManagedObjectReference `xml:"cluster,omitempty"`
	ObjUuids []string                      `xml:"objUuids"`
	Force    *bool                         `xml:"force"`
}

func init() {
	types.Add("vsan:VsanDeleteObjectsRequestType", reflect.TypeOf((*VsanDeleteObjectsRequestType)(nil)).Elem())
}

type VsanDeleteObjects_Task VsanDeleteObjectsRequestType

func init() {
	types.Add("vsan:VsanDeleteObjects_Task", reflect.TypeOf((*VsanDeleteObjects_Task)(nil)).Elem())
}

type VsanDeleteObjects_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanDisallowDataMovementIssue struct {
	VsanUpgradeSystemPreflightCheckIssue
}

func init() {
	types.Add("vsan:VsanDisallowDataMovementIssue", reflect.TypeOf((*VsanDisallowDataMovementIssue)(nil)).Elem())
}

type VsanDiskEncryptionHealth struct {
	types.DynamicData

	DiskHealth       *VsanPhysicalDiskHealth `xml:"diskHealth,omitempty"`
	EncryptionIssues []string                `xml:"encryptionIssues,omitempty"`
}

func init() {
	types.Add("vsan:VsanDiskEncryptionHealth", reflect.TypeOf((*VsanDiskEncryptionHealth)(nil)).Elem())
}

type VsanDiskFormatConversionCheckResult struct {
	VsanUpgradeSystemPreflightCheckResult

	IsSupported            bool  `xml:"isSupported"`
	TargetVersion          int32 `xml:"targetVersion,omitempty"`
	IsDataMovementRequired *bool `xml:"isDataMovementRequired"`
}

func init() {
	types.Add("vsan:VsanDiskFormatConversionCheckResult", reflect.TypeOf((*VsanDiskFormatConversionCheckResult)(nil)).Elem())
}

type VsanDiskFormatConversionSpec struct {
	types.DynamicData

	DataEfficiencyConfig *VsanDataEfficiencyConfig `xml:"dataEfficiencyConfig,omitempty"`
	DataEncryptionConfig *VsanDataEncryptionConfig `xml:"dataEncryptionConfig,omitempty"`
	SkipHostRemediation  *bool                     `xml:"skipHostRemediation"`
	AllowDataMovement    *bool                     `xml:"allowDataMovement"`
}

func init() {
	types.Add("vsan:VsanDiskFormatConversionSpec", reflect.TypeOf((*VsanDiskFormatConversionSpec)(nil)).Elem())
}

type VsanDiskGroupResourceCheckResult struct {
	EntityResourceCheckDetails

	CacheTierDisk     *VsanDiskResourceCheckResult  `xml:"cacheTierDisk,omitempty"`
	CapacityTierDisks []VsanDiskResourceCheckResult `xml:"capacityTierDisks,omitempty"`
}

func init() {
	types.Add("vsan:VsanDiskGroupResourceCheckResult", reflect.TypeOf((*VsanDiskGroupResourceCheckResult)(nil)).Elem())
}

type VsanDiskRebalanceResult struct {
	types.DynamicData

	Status               string  `xml:"status"`
	BytesMoving          int64   `xml:"bytesMoving,omitempty"`
	RemainingBytesToMove int64   `xml:"remainingBytesToMove,omitempty"`
	DiskUsage            float32 `xml:"diskUsage,omitempty"`
	MaxDiskUsage         float32 `xml:"maxDiskUsage,omitempty"`
	MinDiskUsage         float32 `xml:"minDiskUsage,omitempty"`
	AvgDiskUsage         float32 `xml:"avgDiskUsage,omitempty"`
}

func init() {
	types.Add("vsan:VsanDiskRebalanceResult", reflect.TypeOf((*VsanDiskRebalanceResult)(nil)).Elem())
}

type VsanDiskResourceCheckResult struct {
	EntityResourceCheckDetails
}

func init() {
	types.Add("vsan:VsanDiskResourceCheckResult", reflect.TypeOf((*VsanDiskResourceCheckResult)(nil)).Elem())
}

type VsanDiskUnhealthIssue struct {
	VsanUpgradeSystemPreflightCheckIssue

	Uuids []string `xml:"uuids"`
}

func init() {
	types.Add("vsan:VsanDiskUnhealthIssue", reflect.TypeOf((*VsanDiskUnhealthIssue)(nil)).Elem())
}

type VsanDownloadAndInstallVendorToolRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

func init() {
	types.Add("vsan:VsanDownloadAndInstallVendorToolRequestType", reflect.TypeOf((*VsanDownloadAndInstallVendorToolRequestType)(nil)).Elem())
}

type VsanDownloadAndInstallVendorTool_Task VsanDownloadAndInstallVendorToolRequestType

func init() {
	types.Add("vsan:VsanDownloadAndInstallVendorTool_Task", reflect.TypeOf((*VsanDownloadAndInstallVendorTool_Task)(nil)).Elem())
}

type VsanDownloadAndInstallVendorTool_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanDownloadFileServiceOvf VsanDownloadFileServiceOvfRequestType

func init() {
	types.Add("vsan:VsanDownloadFileServiceOvf", reflect.TypeOf((*VsanDownloadFileServiceOvf)(nil)).Elem())
}

type VsanDownloadFileServiceOvfRequestType struct {
	This        types.ManagedObjectReference `xml:"_this"`
	DownloadUrl string                       `xml:"downloadUrl"`
}

func init() {
	types.Add("vsan:VsanDownloadFileServiceOvfRequestType", reflect.TypeOf((*VsanDownloadFileServiceOvfRequestType)(nil)).Elem())
}

type VsanDownloadFileServiceOvfResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanDownloadHclFileRequestType struct {
	This     types.ManagedObjectReference `xml:"_this"`
	Sha1sums []string                     `xml:"sha1sums"`
}

func init() {
	types.Add("vsan:VsanDownloadHclFileRequestType", reflect.TypeOf((*VsanDownloadHclFileRequestType)(nil)).Elem())
}

type VsanDownloadHclFile_Task VsanDownloadHclFileRequestType

func init() {
	types.Add("vsan:VsanDownloadHclFile_Task", reflect.TypeOf((*VsanDownloadHclFile_Task)(nil)).Elem())
}

type VsanDownloadHclFile_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanDownloadItem struct {
	types.DynamicData

	Url        string `xml:"url"`
	Sha1sum    string `xml:"sha1sum"`
	FormatType string `xml:"formatType,omitempty"`
	ItemId     string `xml:"itemId,omitempty"`
}

func init() {
	types.Add("vsan:VsanDownloadItem", reflect.TypeOf((*VsanDownloadItem)(nil)).Elem())
}

type VsanEncryptedClusterRekeyRequestType struct {
	This                   types.ManagedObjectReference `xml:"_this"`
	EncryptedCluster       types.ManagedObjectReference `xml:"encryptedCluster"`
	DeepRekey              *bool                        `xml:"deepRekey"`
	AllowReducedRedundancy *bool                        `xml:"allowReducedRedundancy"`
}

func init() {
	types.Add("vsan:VsanEncryptedClusterRekeyRequestType", reflect.TypeOf((*VsanEncryptedClusterRekeyRequestType)(nil)).Elem())
}

type VsanEncryptedClusterRekey_Task VsanEncryptedClusterRekeyRequestType

func init() {
	types.Add("vsan:VsanEncryptedClusterRekey_Task", reflect.TypeOf((*VsanEncryptedClusterRekey_Task)(nil)).Elem())
}

type VsanEncryptedClusterRekey_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanEncryptionHealthSummary struct {
	types.DynamicData

	Hostname         string                      `xml:"hostname,omitempty"`
	EncryptionInfo   *VsanHostEncryptionInfo     `xml:"encryptionInfo,omitempty"`
	OverallKmsHealth string                      `xml:"overallKmsHealth"`
	KmsHealth        []VsanKmsHealth             `xml:"kmsHealth,omitempty"`
	EncryptionIssues []string                    `xml:"encryptionIssues,omitempty"`
	DiskResults      []VsanDiskEncryptionHealth  `xml:"diskResults,omitempty"`
	Error            *types.LocalizedMethodFault `xml:"error,omitempty"`
	AesniEnabled     *bool                       `xml:"aesniEnabled"`
}

func init() {
	types.Add("vsan:VsanEncryptionHealthSummary", reflect.TypeOf((*VsanEncryptionHealthSummary)(nil)).Elem())
}

type VsanEntitySpaceUsage struct {
	types.DynamicData

	EntityId               string                              `xml:"entityId,omitempty"`
	SpaceUsageByObjectType []VsanObjectSpaceSummary            `xml:"spaceUsageByObjectType,omitempty"`
	TotalCapacityB         int64                               `xml:"totalCapacityB,omitempty"`
	FreeCapacityB          int64                               `xml:"freeCapacityB,omitempty"`
	EfficientCapacity      *VimVsanDataEfficiencyCapacityState `xml:"efficientCapacity,omitempty"`
}

func init() {
	types.Add("vsan:VsanEntitySpaceUsage", reflect.TypeOf((*VsanEntitySpaceUsage)(nil)).Elem())
}

type VsanExtendedConfig struct {
	types.DynamicData

	ObjectRepairTimer          int64                       `xml:"objectRepairTimer,omitempty"`
	DisableSiteReadLocality    *bool                       `xml:"disableSiteReadLocality"`
	EnableCustomizedSwapObject *bool                       `xml:"enableCustomizedSwapObject"`
	LargeScaleClusterSupport   *bool                       `xml:"largeScaleClusterSupport"`
	ProactiveRebalanceInfo     *VsanProactiveRebalanceInfo `xml:"proactiveRebalanceInfo,omitempty"`
}

func init() {
	types.Add("vsan:VsanExtendedConfig", reflect.TypeOf((*VsanExtendedConfig)(nil)).Elem())
}

type VsanFailedRepairObjectResult struct {
	types.DynamicData

	Uuid       string `xml:"uuid"`
	ErrMessage string `xml:"errMessage,omitempty"`
}

func init() {
	types.Add("vsan:VsanFailedRepairObjectResult", reflect.TypeOf((*VsanFailedRepairObjectResult)(nil)).Elem())
}

type VsanFaultDomainResourceCheckResult struct {
	EntityResourceCheckDetails

	Hosts []VsanHostResourceCheckResult `xml:"hosts,omitempty"`
}

func init() {
	types.Add("vsan:VsanFaultDomainResourceCheckResult", reflect.TypeOf((*VsanFaultDomainResourceCheckResult)(nil)).Elem())
}

type VsanFileServerHealthSummary struct {
	types.DynamicData

	DomainName    string `xml:"domainName,omitempty"`
	FileServerIp  string `xml:"fileServerIp,omitempty"`
	NfsdHealth    string `xml:"nfsdHealth,omitempty"`
	NetworkHealth string `xml:"networkHealth,omitempty"`
	RootfsHealth  string `xml:"rootfsHealth,omitempty"`
	Description   string `xml:"description,omitempty"`
}

func init() {
	types.Add("vsan:VsanFileServerHealthSummary", reflect.TypeOf((*VsanFileServerHealthSummary)(nil)).Elem())
}

type VsanFileServiceAdServiceHealthSummary struct {
	types.DynamicData

	DomainName   string `xml:"domainName,omitempty"`
	FileServerIp string `xml:"fileServerIp,omitempty"`
	Health       string `xml:"health,omitempty"`
	Description  string `xml:"description,omitempty"`
}

func init() {
	types.Add("vsan:VsanFileServiceAdServiceHealthSummary", reflect.TypeOf((*VsanFileServiceAdServiceHealthSummary)(nil)).Elem())
}

type VsanFileServiceConfig struct {
	types.DynamicData

	Enabled            bool                          `xml:"enabled"`
	FileServerMemoryMB int64                         `xml:"fileServerMemoryMB,omitempty"`
	FileServerCPUMhz   int64                         `xml:"fileServerCPUMhz,omitempty"`
	FsvmMemoryMB       int64                         `xml:"fsvmMemoryMB,omitempty"`
	FsvmCPU            int64                         `xml:"fsvmCPU,omitempty"`
	Network            *types.ManagedObjectReference `xml:"network,omitempty"`
	Domains            []VsanFileServiceDomainConfig `xml:"domains,omitempty"`
}

func init() {
	types.Add("vsan:VsanFileServiceConfig", reflect.TypeOf((*VsanFileServiceConfig)(nil)).Elem())
}

type VsanFileServiceDomain struct {
	types.DynamicData

	Uuid   string                       `xml:"uuid"`
	Config *VsanFileServiceDomainConfig `xml:"config,omitempty"`
}

func init() {
	types.Add("vsan:VsanFileServiceDomain", reflect.TypeOf((*VsanFileServiceDomain)(nil)).Elem())
}

type VsanFileServiceDomainConfig struct {
	types.DynamicData

	Name               string                    `xml:"name,omitempty"`
	DnsServerAddresses []string                  `xml:"dnsServerAddresses,omitempty"`
	DnsSuffixes        []string                  `xml:"dnsSuffixes,omitempty"`
	FileServerIpConfig []VsanFileServiceIpConfig `xml:"fileServerIpConfig,omitempty"`
}

func init() {
	types.Add("vsan:VsanFileServiceDomainConfig", reflect.TypeOf((*VsanFileServiceDomainConfig)(nil)).Elem())
}

type VsanFileServiceDomainQuerySpec struct {
	types.DynamicData

	Uuids []string `xml:"uuids,omitempty"`
	Names []string `xml:"names,omitempty"`
}

func init() {
	types.Add("vsan:VsanFileServiceDomainQuerySpec", reflect.TypeOf((*VsanFileServiceDomainQuerySpec)(nil)).Elem())
}

type VsanFileServiceHealthSummary struct {
	types.DynamicData

	Hostname         string                              `xml:"hostname,omitempty"`
	OverallHealth    string                              `xml:"overallHealth,omitempty"`
	Enabled          *bool                               `xml:"enabled"`
	VdfsdStatus      *VsanResourceHealth                 `xml:"vdfsdStatus,omitempty"`
	FsvmStatus       *VsanResourceHealth                 `xml:"fsvmStatus,omitempty"`
	RootFsStatus     *VsanFileServiceRootFsHealth        `xml:"rootFsStatus,omitempty"`
	FileServerHealth []VsanFileServerHealthSummary       `xml:"fileServerHealth,omitempty"`
	FileShareHealth  []VsanFileServiceShareHealthSummary `xml:"fileShareHealth,omitempty"`
}

func init() {
	types.Add("vsan:VsanFileServiceHealthSummary", reflect.TypeOf((*VsanFileServiceHealthSummary)(nil)).Elem())
}

type VsanFileServiceIpConfig struct {
	types.HostIpConfig

	Fqdn      string `xml:"fqdn,omitempty"`
	IsPrimary *bool  `xml:"isPrimary"`
	Gateway   string `xml:"gateway"`
}

func init() {
	types.Add("vsan:VsanFileServiceIpConfig", reflect.TypeOf((*VsanFileServiceIpConfig)(nil)).Elem())
}

type VsanFileServiceOvfSpec struct {
	types.DynamicData

	Version    string                        `xml:"version,omitempty"`
	UpdateTime *time.Time                    `xml:"updateTime"`
	Task       *types.ManagedObjectReference `xml:"task,omitempty"`
}

func init() {
	types.Add("vsan:VsanFileServiceOvfSpec", reflect.TypeOf((*VsanFileServiceOvfSpec)(nil)).Elem())
}

type VsanFileServicePreflightCheckResult struct {
	types.DynamicData

	OvfInstalled          string     `xml:"ovfInstalled,omitempty"`
	FsvmVersion           string     `xml:"fsvmVersion,omitempty"`
	LastUpgradeDate       *time.Time `xml:"lastUpgradeDate"`
	OvfMixedModeIssue     string     `xml:"ovfMixedModeIssue,omitempty"`
	HostVersion           string     `xml:"hostVersion,omitempty"`
	MixedModeIssue        string     `xml:"mixedModeIssue,omitempty"`
	NetworkPartitionIssue string     `xml:"networkPartitionIssue,omitempty"`
	VsanDatastoreIssue    string     `xml:"vsanDatastoreIssue,omitempty"`
	DomainConfigIssue     string     `xml:"domainConfigIssue,omitempty"`
	DvsConfigIssue        string     `xml:"dvsConfigIssue,omitempty"`
}

func init() {
	types.Add("vsan:VsanFileServicePreflightCheckResult", reflect.TypeOf((*VsanFileServicePreflightCheckResult)(nil)).Elem())
}

type VsanFileServiceRootFsHealth struct {
	types.DynamicData

	Created     *bool  `xml:"created"`
	Health      string `xml:"health,omitempty"`
	Description string `xml:"description,omitempty"`
}

func init() {
	types.Add("vsan:VsanFileServiceRootFsHealth", reflect.TypeOf((*VsanFileServiceRootFsHealth)(nil)).Elem())
}

type VsanFileServiceShareHealthSummary struct {
	types.DynamicData

	OverallHealth string                   `xml:"overallHealth,omitempty"`
	DomainName    string                   `xml:"domainName,omitempty"`
	ShareUuid     string                   `xml:"shareUuid,omitempty"`
	ShareName     string                   `xml:"shareName,omitempty"`
	ObjectHealth  *VsanObjectOverallHealth `xml:"objectHealth,omitempty"`
	Description   string                   `xml:"description,omitempty"`
}

func init() {
	types.Add("vsan:VsanFileServiceShareHealthSummary", reflect.TypeOf((*VsanFileServiceShareHealthSummary)(nil)).Elem())
}

type VsanFileShare struct {
	types.DynamicData

	Uuid    string                    `xml:"uuid"`
	Config  *VsanFileShareConfig      `xml:"config,omitempty"`
	Runtime *VsanFileShareRuntimeInfo `xml:"runtime,omitempty"`
}

func init() {
	types.Add("vsan:VsanFileShare", reflect.TypeOf((*VsanFileShare)(nil)).Elem())
}

type VsanFileShareConfig struct {
	types.DynamicData

	Name          string                              `xml:"name,omitempty"`
	DomainName    string                              `xml:"domainName,omitempty"`
	Quota         string                              `xml:"quota,omitempty"`
	SoftQuota     string                              `xml:"softQuota,omitempty"`
	Labels        []types.KeyValue                    `xml:"labels,omitempty"`
	StoragePolicy types.BaseVirtualMachineProfileSpec `xml:"storagePolicy,omitempty,typeattr"`
	Permission    []VsanFileShareNetPermission        `xml:"permission,omitempty"`
}

func init() {
	types.Add("vsan:VsanFileShareConfig", reflect.TypeOf((*VsanFileShareConfig)(nil)).Elem())
}

type VsanFileShareNetPermission struct {
	types.DynamicData

	Ips         string `xml:"ips"`
	Permissions string `xml:"permissions,omitempty"`
	AllowRoot   *bool  `xml:"allowRoot"`
}

func init() {
	types.Add("vsan:VsanFileShareNetPermission", reflect.TypeOf((*VsanFileShareNetPermission)(nil)).Elem())
}

type VsanFileShareQuerySpec struct {
	types.DynamicData

	DomainName string   `xml:"domainName,omitempty"`
	Uuids      []string `xml:"uuids,omitempty"`
	Names      []string `xml:"names,omitempty"`
	Offset     string   `xml:"offset,omitempty"`
	Limit      *int64   `xml:"limit"`
}

func init() {
	types.Add("vsan:VsanFileShareQuerySpec", reflect.TypeOf((*VsanFileShareQuerySpec)(nil)).Elem())
}

type VsanFileShareRuntimeInfo struct {
	types.DynamicData

	UsedCapacity    int64            `xml:"usedCapacity,omitempty"`
	Hostname        string           `xml:"hostname,omitempty"`
	Address         string           `xml:"address,omitempty"`
	VsanObjectUuids []string         `xml:"vsanObjectUuids,omitempty"`
	AccessPoints    []types.KeyValue `xml:"accessPoints,omitempty"`
	ManagedBy       string           `xml:"managedBy,omitempty"`
}

func init() {
	types.Add("vsan:VsanFileShareRuntimeInfo", reflect.TypeOf((*VsanFileShareRuntimeInfo)(nil)).Elem())
}

type VsanFindOvfDownloadUrl VsanFindOvfDownloadUrlRequestType

func init() {
	types.Add("vsan:VsanFindOvfDownloadUrl", reflect.TypeOf((*VsanFindOvfDownloadUrl)(nil)).Elem())
}

type VsanFindOvfDownloadUrlRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

func init() {
	types.Add("vsan:VsanFindOvfDownloadUrlRequestType", reflect.TypeOf((*VsanFindOvfDownloadUrlRequestType)(nil)).Elem())
}

type VsanFindOvfDownloadUrlResponse struct {
	Returnval string `xml:"returnval"`
}

type VsanFlashScsiControllerFirmwareRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
	Spec VsanHclFirmwareUpdateSpec    `xml:"spec"`
}

func init() {
	types.Add("vsan:VsanFlashScsiControllerFirmwareRequestType", reflect.TypeOf((*VsanFlashScsiControllerFirmwareRequestType)(nil)).Elem())
}

type VsanFlashScsiControllerFirmware_Task VsanFlashScsiControllerFirmwareRequestType

func init() {
	types.Add("vsan:VsanFlashScsiControllerFirmware_Task", reflect.TypeOf((*VsanFlashScsiControllerFirmware_Task)(nil)).Elem())
}

type VsanFlashScsiControllerFirmware_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanGenericClusterBaseIssue struct {
	types.DynamicData
}

func init() {
	types.Add("vsan:VsanGenericClusterBaseIssue", reflect.TypeOf((*VsanGenericClusterBaseIssue)(nil)).Elem())
}

type VsanGenericClusterBestPracticeHealth struct {
	types.DynamicData

	DrsEnabled bool                          `xml:"drsEnabled"`
	HaEnabled  bool                          `xml:"haEnabled"`
	Issues     []VsanGenericClusterBaseIssue `xml:"issues,omitempty"`
}

func init() {
	types.Add("vsan:VsanGenericClusterBestPracticeHealth", reflect.TypeOf((*VsanGenericClusterBestPracticeHealth)(nil)).Elem())
}

type VsanGetAboutInfoEx VsanGetAboutInfoExRequestType

func init() {
	types.Add("vsan:VsanGetAboutInfoEx", reflect.TypeOf((*VsanGetAboutInfoEx)(nil)).Elem())
}

type VsanGetAboutInfoExRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

func init() {
	types.Add("vsan:VsanGetAboutInfoExRequestType", reflect.TypeOf((*VsanGetAboutInfoExRequestType)(nil)).Elem())
}

type VsanGetAboutInfoExResponse struct {
	Returnval VsanHostAboutInfoEx `xml:"returnval"`
}

type VsanGetCapabilities VsanGetCapabilitiesRequestType

func init() {
	types.Add("vsan:VsanGetCapabilities", reflect.TypeOf((*VsanGetCapabilities)(nil)).Elem())
}

type VsanGetCapabilitiesRequestType struct {
	This    types.ManagedObjectReference   `xml:"_this"`
	Targets []types.ManagedObjectReference `xml:"targets,omitempty"`
}

func init() {
	types.Add("vsan:VsanGetCapabilitiesRequestType", reflect.TypeOf((*VsanGetCapabilitiesRequestType)(nil)).Elem())
}

type VsanGetCapabilitiesResponse struct {
	Returnval []VsanCapability `xml:"returnval"`
}

type VsanGetHclConstraints VsanGetHclConstraintsRequestType

func init() {
	types.Add("vsan:VsanGetHclConstraints", reflect.TypeOf((*VsanGetHclConstraints)(nil)).Elem())
}

type VsanGetHclConstraintsRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
	Release string                       `xml:"release"`
}

func init() {
	types.Add("vsan:VsanGetHclConstraintsRequestType", reflect.TypeOf((*VsanGetHclConstraintsRequestType)(nil)).Elem())
}

type VsanGetHclConstraintsResponse struct {
	Returnval VsanHclReleaseConstraint `xml:"returnval"`
}

type VsanGetHclInfo VsanGetHclInfoRequestType

func init() {
	types.Add("vsan:VsanGetHclInfo", reflect.TypeOf((*VsanGetHclInfo)(nil)).Elem())
}

type VsanGetHclInfoRequestType struct {
	This              types.ManagedObjectReference `xml:"_this"`
	IncludeVendorInfo *bool                        `xml:"includeVendorInfo"`
}

func init() {
	types.Add("vsan:VsanGetHclInfoRequestType", reflect.TypeOf((*VsanGetHclInfoRequestType)(nil)).Elem())
}

type VsanGetHclInfoResponse struct {
	Returnval VsanHostHclInfo `xml:"returnval"`
}

type VsanGetProactiveRebalanceInfo VsanGetProactiveRebalanceInfoRequestType

func init() {
	types.Add("vsan:VsanGetProactiveRebalanceInfo", reflect.TypeOf((*VsanGetProactiveRebalanceInfo)(nil)).Elem())
}

type VsanGetProactiveRebalanceInfoRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

func init() {
	types.Add("vsan:VsanGetProactiveRebalanceInfoRequestType", reflect.TypeOf((*VsanGetProactiveRebalanceInfoRequestType)(nil)).Elem())
}

type VsanGetProactiveRebalanceInfoResponse struct {
	Returnval VsanProactiveRebalanceInfoEx `xml:"returnval"`
}

type VsanGetReleaseRecommendation VsanGetReleaseRecommendationRequestType

func init() {
	types.Add("vsan:VsanGetReleaseRecommendation", reflect.TypeOf((*VsanGetReleaseRecommendation)(nil)).Elem())
}

type VsanGetReleaseRecommendationRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
	Minor   []string                     `xml:"minor"`
	Major   []string                     `xml:"major"`
}

func init() {
	types.Add("vsan:VsanGetReleaseRecommendationRequestType", reflect.TypeOf((*VsanGetReleaseRecommendationRequestType)(nil)).Elem())
}

type VsanGetReleaseRecommendationResponse struct {
	Returnval []VsanHclReleaseConstraint `xml:"returnval,omitempty"`
}

type VsanGetResourceCheckStatus VsanGetResourceCheckStatusRequestType

func init() {
	types.Add("vsan:VsanGetResourceCheckStatus", reflect.TypeOf((*VsanGetResourceCheckStatus)(nil)).Elem())
}

type VsanGetResourceCheckStatusRequestType struct {
	This              types.ManagedObjectReference  `xml:"_this"`
	ResourceCheckSpec *VsanResourceCheckSpec        `xml:"resourceCheckSpec,omitempty"`
	Cluster           *types.ManagedObjectReference `xml:"cluster,omitempty"`
}

func init() {
	types.Add("vsan:VsanGetResourceCheckStatusRequestType", reflect.TypeOf((*VsanGetResourceCheckStatusRequestType)(nil)).Elem())
}

type VsanGetResourceCheckStatusResponse struct {
	Returnval VsanResourceCheckStatus `xml:"returnval"`
}

type VsanHclCommonDeviceInfo struct {
	types.DynamicData

	DeviceName             string              `xml:"deviceName"`
	DisplayName            string              `xml:"displayName,omitempty"`
	DriverName             string              `xml:"driverName,omitempty"`
	DriverVersion          string              `xml:"driverVersion,omitempty"`
	VendorId               int64               `xml:"vendorId,omitempty"`
	DeviceId               int64               `xml:"deviceId,omitempty"`
	SubVendorId            int64               `xml:"subVendorId,omitempty"`
	SubDeviceId            int64               `xml:"subDeviceId,omitempty"`
	ExtraInfo              []types.KeyValue    `xml:"extraInfo,omitempty"`
	DeviceOnHcl            *bool               `xml:"deviceOnHcl"`
	ReleaseSupported       *bool               `xml:"releaseSupported"`
	ReleasesOnHcl          []string            `xml:"releasesOnHcl,omitempty"`
	DriverVersionsOnHcl    []string            `xml:"driverVersionsOnHcl,omitempty"`
	DriverVersionSupported *bool               `xml:"driverVersionSupported"`
	FwVersionSupported     *bool               `xml:"fwVersionSupported"`
	FwVersionOnHcl         []string            `xml:"fwVersionOnHcl,omitempty"`
	FwVersion              string              `xml:"fwVersion,omitempty"`
	DriversOnHcl           []VsanHclDriverInfo `xml:"driversOnHcl,omitempty"`
}

func init() {
	types.Add("vsan:VsanHclCommonDeviceInfo", reflect.TypeOf((*VsanHclCommonDeviceInfo)(nil)).Elem())
}

type VsanHclControllerInfo struct {
	types.DynamicData

	DeviceName             string                   `xml:"deviceName"`
	DeviceDisplayName      string                   `xml:"deviceDisplayName,omitempty"`
	DriverName             string                   `xml:"driverName,omitempty"`
	DriverVersion          string                   `xml:"driverVersion,omitempty"`
	VendorId               int64                    `xml:"vendorId,omitempty"`
	DeviceId               int64                    `xml:"deviceId,omitempty"`
	SubVendorId            int64                    `xml:"subVendorId,omitempty"`
	SubDeviceId            int64                    `xml:"subDeviceId,omitempty"`
	ExtraInfo              []types.KeyValue         `xml:"extraInfo,omitempty"`
	DeviceOnHcl            *bool                    `xml:"deviceOnHcl"`
	ReleaseSupported       *bool                    `xml:"releaseSupported"`
	ReleasesOnHcl          []string                 `xml:"releasesOnHcl,omitempty"`
	DriverVersionsOnHcl    []string                 `xml:"driverVersionsOnHcl,omitempty"`
	DriverVersionSupported *bool                    `xml:"driverVersionSupported"`
	FwVersionSupported     *bool                    `xml:"fwVersionSupported"`
	FwVersionOnHcl         []string                 `xml:"fwVersionOnHcl,omitempty"`
	CacheConfigSupported   *bool                    `xml:"cacheConfigSupported"`
	CacheConfigOnHcl       []string                 `xml:"cacheConfigOnHcl,omitempty"`
	RaidConfigSupported    *bool                    `xml:"raidConfigSupported"`
	RaidConfigOnHcl        []string                 `xml:"raidConfigOnHcl,omitempty"`
	FwVersion              string                   `xml:"fwVersion,omitempty"`
	RaidConfig             string                   `xml:"raidConfig,omitempty"`
	CacheConfig            string                   `xml:"cacheConfig,omitempty"`
	CimProviderInfo        *VsanHostCimProviderInfo `xml:"cimProviderInfo,omitempty"`
	UsedByVsan             *bool                    `xml:"usedByVsan"`
	Disks                  []VsanPhysicalDiskHealth `xml:"disks,omitempty"`
	Issues                 []string                 `xml:"issues,omitempty"`
	RemediableIssues       []string                 `xml:"remediableIssues,omitempty"`
	DriversOnHcl           []VsanHclDriverInfo      `xml:"driversOnHcl,omitempty"`
	FwAuxVersion           string                   `xml:"fwAuxVersion,omitempty"`
	QueueDepth             int32                    `xml:"queueDepth,omitempty"`
	QueueDepthOnHcl        int64                    `xml:"queueDepthOnHcl,omitempty"`
	QueueDepthSupported    *bool                    `xml:"queueDepthSupported"`
	DiskMode               *types.ChoiceOption      `xml:"diskMode,omitempty"`
	DiskModeOnHcl          []string                 `xml:"diskModeOnHcl,omitempty"`
	DiskModeSupported      *bool                    `xml:"diskModeSupported"`
	ToolName               string                   `xml:"toolName,omitempty"`
	ToolVersion            string                   `xml:"toolVersion,omitempty"`
}

func init() {
	types.Add("vsan:VsanHclControllerInfo", reflect.TypeOf((*VsanHclControllerInfo)(nil)).Elem())
}

type VsanHclDeviceConstraint struct {
	types.DynamicData

	PciId              string                  `xml:"pciId"`
	VcgLink            string                  `xml:"vcgLink,omitempty"`
	SimilarVcgLinks    []string                `xml:"similarVcgLinks,omitempty"`
	CompliantFirmwares []VsanCompliantFirmware `xml:"compliantFirmwares,omitempty"`
}

func init() {
	types.Add("vsan:VsanHclDeviceConstraint", reflect.TypeOf((*VsanHclDeviceConstraint)(nil)).Elem())
}

type VsanHclDiskInfo struct {
	types.DynamicData

	DeviceName       string                       `xml:"deviceName"`
	Model            string                       `xml:"model,omitempty"`
	IsSsd            *bool                        `xml:"isSsd"`
	VsanDisk         bool                         `xml:"vsanDisk"`
	Issues           []types.LocalizedMethodFault `xml:"issues,omitempty"`
	RemediableIssues []string                     `xml:"remediableIssues,omitempty"`
}

func init() {
	types.Add("vsan:VsanHclDiskInfo", reflect.TypeOf((*VsanHclDiskInfo)(nil)).Elem())
}

type VsanHclDriverInfo struct {
	types.DynamicData

	DriverVersion string             `xml:"driverVersion,omitempty"`
	DriverLink    *VsanDownloadItem  `xml:"driverLink,omitempty"`
	FwVersion     string             `xml:"fwVersion,omitempty"`
	FwLinks       []VsanDownloadItem `xml:"fwLinks,omitempty"`
	ToolsLinks    []VsanDownloadItem `xml:"toolsLinks,omitempty"`
	Eula          string             `xml:"eula,omitempty"`
	DriverType    string             `xml:"driverType,omitempty"`
	DriverName    string             `xml:"driverName,omitempty"`
	DiskModes     []string           `xml:"diskModes,omitempty"`
}

func init() {
	types.Add("vsan:VsanHclDriverInfo", reflect.TypeOf((*VsanHclDriverInfo)(nil)).Elem())
}

type VsanHclFirmwareFile struct {
	types.DynamicData

	FileType      string `xml:"fileType"`
	FilenameOrUrl string `xml:"filenameOrUrl"`
	Sha1sum       string `xml:"sha1sum"`
}

func init() {
	types.Add("vsan:VsanHclFirmwareFile", reflect.TypeOf((*VsanHclFirmwareFile)(nil)).Elem())
}

type VsanHclFirmwareUpdateSpec struct {
	types.DynamicData

	Host              types.ManagedObjectReference `xml:"host"`
	HbaDevice         string                       `xml:"hbaDevice"`
	FwFiles           []VsanHclFirmwareFile        `xml:"fwFiles"`
	AllowDowngrade    *bool                        `xml:"allowDowngrade"`
	FirmwareComponent []VsanHostFwComponent        `xml:"firmwareComponent,omitempty"`
}

func init() {
	types.Add("vsan:VsanHclFirmwareUpdateSpec", reflect.TypeOf((*VsanHclFirmwareUpdateSpec)(nil)).Elem())
}

type VsanHclNicInfo struct {
	VsanHclCommonDeviceInfo
}

func init() {
	types.Add("vsan:VsanHclNicInfo", reflect.TypeOf((*VsanHclNicInfo)(nil)).Elem())
}

type VsanHclReleaseConstraint struct {
	types.DynamicData

	Cluster     types.ManagedObjectReference `xml:"cluster"`
	Release     string                       `xml:"release"`
	HostDevices []VsanHostDeviceInfo         `xml:"hostDevices,omitempty"`
	Constraints []VsanHclDeviceConstraint    `xml:"constraints,omitempty"`
}

func init() {
	types.Add("vsan:VsanHclReleaseConstraint", reflect.TypeOf((*VsanHclReleaseConstraint)(nil)).Elem())
}

type VsanHealthExtMgmtPreCheckResult struct {
	types.DynamicData

	OverallResult            bool                    `xml:"overallResult"`
	EsxVersionCheckPassed    *bool                   `xml:"esxVersionCheckPassed"`
	DrsCheckPassed           *bool                   `xml:"drsCheckPassed"`
	EamConnectionCheckPassed *bool                   `xml:"eamConnectionCheckPassed"`
	InstallStateCheckPassed  *bool                   `xml:"installStateCheckPassed"`
	Results                  []VsanClusterHealthTest `xml:"results"`
	VumRegistered            *bool                   `xml:"vumRegistered"`
}

func init() {
	types.Add("vsan:VsanHealthExtMgmtPreCheckResult", reflect.TypeOf((*VsanHealthExtMgmtPreCheckResult)(nil)).Elem())
}

type VsanHealthGetVsanClusterSilentChecks VsanHealthGetVsanClusterSilentChecksRequestType

func init() {
	types.Add("vsan:VsanHealthGetVsanClusterSilentChecks", reflect.TypeOf((*VsanHealthGetVsanClusterSilentChecks)(nil)).Elem())
}

type VsanHealthGetVsanClusterSilentChecksRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

func init() {
	types.Add("vsan:VsanHealthGetVsanClusterSilentChecksRequestType", reflect.TypeOf((*VsanHealthGetVsanClusterSilentChecksRequestType)(nil)).Elem())
}

type VsanHealthGetVsanClusterSilentChecksResponse struct {
	Returnval []string `xml:"returnval,omitempty"`
}

type VsanHealthIsRebalanceRunning VsanHealthIsRebalanceRunningRequestType

func init() {
	types.Add("vsan:VsanHealthIsRebalanceRunning", reflect.TypeOf((*VsanHealthIsRebalanceRunning)(nil)).Elem())
}

type VsanHealthIsRebalanceRunningRequestType struct {
	This        types.ManagedObjectReference   `xml:"_this"`
	Cluster     types.ManagedObjectReference   `xml:"cluster"`
	TargetHosts []types.ManagedObjectReference `xml:"targetHosts,omitempty"`
}

func init() {
	types.Add("vsan:VsanHealthIsRebalanceRunningRequestType", reflect.TypeOf((*VsanHealthIsRebalanceRunningRequestType)(nil)).Elem())
}

type VsanHealthIsRebalanceRunningResponse struct {
	Returnval bool `xml:"returnval"`
}

type VsanHealthQueryVsanClusterHealthCheckInterval VsanHealthQueryVsanClusterHealthCheckIntervalRequestType

func init() {
	types.Add("vsan:VsanHealthQueryVsanClusterHealthCheckInterval", reflect.TypeOf((*VsanHealthQueryVsanClusterHealthCheckInterval)(nil)).Elem())
}

type VsanHealthQueryVsanClusterHealthCheckIntervalRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

func init() {
	types.Add("vsan:VsanHealthQueryVsanClusterHealthCheckIntervalRequestType", reflect.TypeOf((*VsanHealthQueryVsanClusterHealthCheckIntervalRequestType)(nil)).Elem())
}

type VsanHealthQueryVsanClusterHealthCheckIntervalResponse struct {
	Returnval int32 `xml:"returnval"`
}

type VsanHealthQueryVsanClusterHealthConfig VsanHealthQueryVsanClusterHealthConfigRequestType

func init() {
	types.Add("vsan:VsanHealthQueryVsanClusterHealthConfig", reflect.TypeOf((*VsanHealthQueryVsanClusterHealthConfig)(nil)).Elem())
}

type VsanHealthQueryVsanClusterHealthConfigRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

func init() {
	types.Add("vsan:VsanHealthQueryVsanClusterHealthConfigRequestType", reflect.TypeOf((*VsanHealthQueryVsanClusterHealthConfigRequestType)(nil)).Elem())
}

type VsanHealthQueryVsanClusterHealthConfigResponse struct {
	Returnval VsanClusterHealthConfigs `xml:"returnval"`
}

type VsanHealthQueryVsanProxyConfig VsanHealthQueryVsanProxyConfigRequestType

func init() {
	types.Add("vsan:VsanHealthQueryVsanProxyConfig", reflect.TypeOf((*VsanHealthQueryVsanProxyConfig)(nil)).Elem())
}

type VsanHealthQueryVsanProxyConfigRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

func init() {
	types.Add("vsan:VsanHealthQueryVsanProxyConfigRequestType", reflect.TypeOf((*VsanHealthQueryVsanProxyConfigRequestType)(nil)).Elem())
}

type VsanHealthQueryVsanProxyConfigResponse struct {
	Returnval VsanClusterTelemetryProxyConfig `xml:"returnval"`
}

type VsanHealthRepairClusterObjectsImmediate VsanHealthRepairClusterObjectsImmediateRequestType

func init() {
	types.Add("vsan:VsanHealthRepairClusterObjectsImmediate", reflect.TypeOf((*VsanHealthRepairClusterObjectsImmediate)(nil)).Elem())
}

type VsanHealthRepairClusterObjectsImmediateRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
	Uuids   []string                     `xml:"uuids,omitempty"`
}

func init() {
	types.Add("vsan:VsanHealthRepairClusterObjectsImmediateRequestType", reflect.TypeOf((*VsanHealthRepairClusterObjectsImmediateRequestType)(nil)).Elem())
}

type VsanHealthRepairClusterObjectsImmediateResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanHealthSendVsanTelemetry VsanHealthSendVsanTelemetryRequestType

func init() {
	types.Add("vsan:VsanHealthSendVsanTelemetry", reflect.TypeOf((*VsanHealthSendVsanTelemetry)(nil)).Elem())
}

type VsanHealthSendVsanTelemetryRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

func init() {
	types.Add("vsan:VsanHealthSendVsanTelemetryRequestType", reflect.TypeOf((*VsanHealthSendVsanTelemetryRequestType)(nil)).Elem())
}

type VsanHealthSendVsanTelemetryResponse struct {
}

type VsanHealthSetLogLevel VsanHealthSetLogLevelRequestType

func init() {
	types.Add("vsan:VsanHealthSetLogLevel", reflect.TypeOf((*VsanHealthSetLogLevel)(nil)).Elem())
}

type VsanHealthSetLogLevelRequestType struct {
	This  types.ManagedObjectReference `xml:"_this"`
	Level string                       `xml:"level,omitempty"`
}

func init() {
	types.Add("vsan:VsanHealthSetLogLevelRequestType", reflect.TypeOf((*VsanHealthSetLogLevelRequestType)(nil)).Elem())
}

type VsanHealthSetLogLevelResponse struct {
}

type VsanHealthSetVsanClusterHealthCheckInterval VsanHealthSetVsanClusterHealthCheckIntervalRequestType

func init() {
	types.Add("vsan:VsanHealthSetVsanClusterHealthCheckInterval", reflect.TypeOf((*VsanHealthSetVsanClusterHealthCheckInterval)(nil)).Elem())
}

type VsanHealthSetVsanClusterHealthCheckIntervalRequestType struct {
	This                           types.ManagedObjectReference `xml:"_this"`
	Cluster                        types.ManagedObjectReference `xml:"cluster"`
	VsanClusterHealthCheckInterval int32                        `xml:"vsanClusterHealthCheckInterval"`
}

func init() {
	types.Add("vsan:VsanHealthSetVsanClusterHealthCheckIntervalRequestType", reflect.TypeOf((*VsanHealthSetVsanClusterHealthCheckIntervalRequestType)(nil)).Elem())
}

type VsanHealthSetVsanClusterHealthCheckIntervalResponse struct {
}

type VsanHealthSetVsanClusterSilentChecks VsanHealthSetVsanClusterSilentChecksRequestType

func init() {
	types.Add("vsan:VsanHealthSetVsanClusterSilentChecks", reflect.TypeOf((*VsanHealthSetVsanClusterSilentChecks)(nil)).Elem())
}

type VsanHealthSetVsanClusterSilentChecksRequestType struct {
	This               types.ManagedObjectReference `xml:"_this"`
	Cluster            types.ManagedObjectReference `xml:"cluster"`
	AddSilentChecks    []string                     `xml:"addSilentChecks,omitempty"`
	RemoveSilentChecks []string                     `xml:"removeSilentChecks,omitempty"`
}

func init() {
	types.Add("vsan:VsanHealthSetVsanClusterSilentChecksRequestType", reflect.TypeOf((*VsanHealthSetVsanClusterSilentChecksRequestType)(nil)).Elem())
}

type VsanHealthSetVsanClusterSilentChecksResponse struct {
	Returnval bool `xml:"returnval"`
}

type VsanHealthSetVsanClusterTelemetryConfig VsanHealthSetVsanClusterTelemetryConfigRequestType

func init() {
	types.Add("vsan:VsanHealthSetVsanClusterTelemetryConfig", reflect.TypeOf((*VsanHealthSetVsanClusterTelemetryConfig)(nil)).Elem())
}

type VsanHealthSetVsanClusterTelemetryConfigRequestType struct {
	This                    types.ManagedObjectReference `xml:"_this"`
	Cluster                 types.ManagedObjectReference `xml:"cluster"`
	VsanClusterHealthConfig VsanClusterHealthConfigs     `xml:"vsanClusterHealthConfig"`
}

func init() {
	types.Add("vsan:VsanHealthSetVsanClusterTelemetryConfigRequestType", reflect.TypeOf((*VsanHealthSetVsanClusterTelemetryConfigRequestType)(nil)).Elem())
}

type VsanHealthSetVsanClusterTelemetryConfigResponse struct {
}

type VsanHealthTestVsanClusterTelemetryProxy VsanHealthTestVsanClusterTelemetryProxyRequestType

func init() {
	types.Add("vsan:VsanHealthTestVsanClusterTelemetryProxy", reflect.TypeOf((*VsanHealthTestVsanClusterTelemetryProxy)(nil)).Elem())
}

type VsanHealthTestVsanClusterTelemetryProxyRequestType struct {
	This        types.ManagedObjectReference    `xml:"_this"`
	ProxyConfig VsanClusterTelemetryProxyConfig `xml:"proxyConfig"`
}

func init() {
	types.Add("vsan:VsanHealthTestVsanClusterTelemetryProxyRequestType", reflect.TypeOf((*VsanHealthTestVsanClusterTelemetryProxyRequestType)(nil)).Elem())
}

type VsanHealthTestVsanClusterTelemetryProxyResponse struct {
	Returnval bool `xml:"returnval"`
}

type VsanHealthThreshold struct {
	types.DynamicData

	YellowValue int64 `xml:"yellowValue"`
	RedValue    int64 `xml:"redValue"`
}

func init() {
	types.Add("vsan:VsanHealthThreshold", reflect.TypeOf((*VsanHealthThreshold)(nil)).Elem())
}

type VsanHigherObjectsPresentDuringDowngradeIssue struct {
	VsanUpgradeSystemPreflightCheckIssue

	Uuids []string `xml:"uuids"`
}

func init() {
	types.Add("vsan:VsanHigherObjectsPresentDuringDowngradeIssue", reflect.TypeOf((*VsanHigherObjectsPresentDuringDowngradeIssue)(nil)).Elem())
}

type VsanHistoryItemQuerySpec struct {
	types.DynamicData

	Clusters []types.ManagedObjectReference `xml:"clusters,omitempty"`
	CleanAll *bool                          `xml:"cleanAll"`
	Start    *time.Time                     `xml:"start"`
	End      *time.Time                     `xml:"end"`
}

func init() {
	types.Add("vsan:VsanHistoryItemQuerySpec", reflect.TypeOf((*VsanHistoryItemQuerySpec)(nil)).Elem())
}

type VsanHostAboutInfoEx struct {
	types.DynamicData

	Name       string `xml:"name,omitempty"`
	Version    string `xml:"version,omitempty"`
	Build      string `xml:"build,omitempty"`
	BuildType  string `xml:"buildType,omitempty"`
	ApiVersion string `xml:"apiVersion,omitempty"`
}

func init() {
	types.Add("vsan:VsanHostAboutInfoEx", reflect.TypeOf((*VsanHostAboutInfoEx)(nil)).Elem())
}

type VsanHostAssociatedObjects struct {
	types.DynamicData

	SpbmProfileId            string   `xml:"spbmProfileId"`
	SpbmProfileGenerationNum int32    `xml:"spbmProfileGenerationNum"`
	VsanObjects              []string `xml:"vsanObjects,omitempty"`
}

func init() {
	types.Add("vsan:VsanHostAssociatedObjects", reflect.TypeOf((*VsanHostAssociatedObjects)(nil)).Elem())
}

type VsanHostAssociatedObjectsResult struct {
	types.DynamicData

	Data   []VsanHostAssociatedObjects `xml:"data"`
	Offset int32                       `xml:"offset"`
	Limit  int32                       `xml:"limit"`
}

func init() {
	types.Add("vsan:VsanHostAssociatedObjectsResult", reflect.TypeOf((*VsanHostAssociatedObjectsResult)(nil)).Elem())
}

type VsanHostCancelResourceCheck VsanHostCancelResourceCheckRequestType

func init() {
	types.Add("vsan:VsanHostCancelResourceCheck", reflect.TypeOf((*VsanHostCancelResourceCheck)(nil)).Elem())
}

type VsanHostCancelResourceCheckRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

func init() {
	types.Add("vsan:VsanHostCancelResourceCheckRequestType", reflect.TypeOf((*VsanHostCancelResourceCheckRequestType)(nil)).Elem())
}

type VsanHostCancelResourceCheckResponse struct {
	Returnval bool `xml:"returnval"`
}

type VsanHostCimProviderInfo struct {
	types.DynamicData

	CimProviderSupported  *bool              `xml:"cimProviderSupported"`
	InstalledCIMProvider  string             `xml:"installedCIMProvider,omitempty"`
	CimProviderOnHcl      []string           `xml:"cimProviderOnHcl,omitempty"`
	CimProviderLinksOnHcl []VsanDownloadItem `xml:"cimProviderLinksOnHcl,omitempty"`
}

func init() {
	types.Add("vsan:VsanHostCimProviderInfo", reflect.TypeOf((*VsanHostCimProviderInfo)(nil)).Elem())
}

type VsanHostCleanupVmdkLoadTest VsanHostCleanupVmdkLoadTestRequestType

func init() {
	types.Add("vsan:VsanHostCleanupVmdkLoadTest", reflect.TypeOf((*VsanHostCleanupVmdkLoadTest)(nil)).Elem())
}

type VsanHostCleanupVmdkLoadTestRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Runname string                       `xml:"runname"`
	Specs   []VsanVmdkLoadTestSpec       `xml:"specs,omitempty"`
}

func init() {
	types.Add("vsan:VsanHostCleanupVmdkLoadTestRequestType", reflect.TypeOf((*VsanHostCleanupVmdkLoadTestRequestType)(nil)).Elem())
}

type VsanHostCleanupVmdkLoadTestResponse struct {
	Returnval string `xml:"returnval"`
}

type VsanHostClomdLiveness VsanHostClomdLivenessRequestType

func init() {
	types.Add("vsan:VsanHostClomdLiveness", reflect.TypeOf((*VsanHostClomdLiveness)(nil)).Elem())
}

type VsanHostClomdLivenessRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

func init() {
	types.Add("vsan:VsanHostClomdLivenessRequestType", reflect.TypeOf((*VsanHostClomdLivenessRequestType)(nil)).Elem())
}

type VsanHostClomdLivenessResponse struct {
	Returnval bool `xml:"returnval"`
}

type VsanHostClomdLivenessResult struct {
	types.DynamicData

	Hostname  string                      `xml:"hostname"`
	ClomdStat string                      `xml:"clomdStat"`
	Error     *types.LocalizedMethodFault `xml:"error,omitempty"`
}

func init() {
	types.Add("vsan:VsanHostClomdLivenessResult", reflect.TypeOf((*VsanHostClomdLivenessResult)(nil)).Elem())
}

type VsanHostComponentSyncState struct {
	types.DynamicData

	Uuid        string   `xml:"uuid"`
	DiskUuid    string   `xml:"diskUuid"`
	HostUuid    string   `xml:"hostUuid"`
	BytesToSync int64    `xml:"bytesToSync"`
	RecoveryETA int64    `xml:"recoveryETA,omitempty"`
	Reasons     []string `xml:"reasons,omitempty"`
}

func init() {
	types.Add("vsan:VsanHostComponentSyncState", reflect.TypeOf((*VsanHostComponentSyncState)(nil)).Elem())
}

type VsanHostConfigInfoEx struct {
	VsanHostConfigInfo

	EncryptionInfo         *VsanHostEncryptionInfo     `xml:"encryptionInfo,omitempty"`
	DataEfficiencyInfo     *VsanDataEfficiencyConfig   `xml:"dataEfficiencyInfo,omitempty"`
	ResyncIopsLimitInfo    *ResyncIopsInfo             `xml:"resyncIopsLimitInfo,omitempty"`
	ExtendedConfig         *VsanExtendedConfig         `xml:"extendedConfig,omitempty"`
	DatastoreInfo          *VsanDatastoreConfig        `xml:"datastoreInfo,omitempty"`
	UnmapConfig            *VsanUnmapConfig            `xml:"unmapConfig,omitempty"`
	WitnessHostConfig      []VsanWitnessHostConfig     `xml:"witnessHostConfig,omitempty"`
	InternalExtendedConfig *VsanInternalExtendedConfig `xml:"internalExtendedConfig,omitempty"`
	MetricsConfig          *VsanMetricsConfig          `xml:"metricsConfig,omitempty"`
}

func init() {
	types.Add("vsan:VsanHostConfigInfoEx", reflect.TypeOf((*VsanHostConfigInfoEx)(nil)).Elem())
}

type VsanHostCreateVmHealthTest VsanHostCreateVmHealthTestRequestType

func init() {
	types.Add("vsan:VsanHostCreateVmHealthTest", reflect.TypeOf((*VsanHostCreateVmHealthTest)(nil)).Elem())
}

type VsanHostCreateVmHealthTestRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Timeout int32                        `xml:"timeout"`
}

func init() {
	types.Add("vsan:VsanHostCreateVmHealthTestRequestType", reflect.TypeOf((*VsanHostCreateVmHealthTestRequestType)(nil)).Elem())
}

type VsanHostCreateVmHealthTestResponse struct {
	Returnval VsanHostCreateVmHealthTestResult `xml:"returnval"`
}

type VsanHostCreateVmHealthTestResult struct {
	types.DynamicData

	Hostname string                      `xml:"hostname"`
	State    string                      `xml:"state"`
	Fault    *types.LocalizedMethodFault `xml:"fault,omitempty"`
}

func init() {
	types.Add("vsan:VsanHostCreateVmHealthTestResult", reflect.TypeOf((*VsanHostCreateVmHealthTestResult)(nil)).Elem())
}

type VsanHostDeviceInfo struct {
	types.DynamicData

	Hostname string                `xml:"hostname"`
	Devices  []VsanBasicDeviceInfo `xml:"devices,omitempty"`
}

func init() {
	types.Add("vsan:VsanHostDeviceInfo", reflect.TypeOf((*VsanHostDeviceInfo)(nil)).Elem())
}

type VsanHostDrsStats struct {
	types.DynamicData

	Host  types.ManagedObjectReference `xml:"host"`
	Stats []byte                       `xml:"stats"`
}

func init() {
	types.Add("vsan:VsanHostDrsStats", reflect.TypeOf((*VsanHostDrsStats)(nil)).Elem())
}

type VsanHostEMMSummary struct {
	types.DynamicData

	Hostname          string `xml:"hostname,omitempty"`
	InMaintenanceMode *bool  `xml:"inMaintenanceMode"`
	InDecomState      *bool  `xml:"inDecomState"`
}

func init() {
	types.Add("vsan:VsanHostEMMSummary", reflect.TypeOf((*VsanHostEMMSummary)(nil)).Elem())
}

type VsanHostEncryptionInfo struct {
	types.DynamicData

	Enabled             *bool                  `xml:"enabled"`
	KekId               string                 `xml:"kekId,omitempty"`
	HostKeyId           string                 `xml:"hostKeyId,omitempty"`
	KmipServers         []types.KmipServerSpec `xml:"kmipServers,omitempty"`
	KmsServerCerts      []string               `xml:"kmsServerCerts,omitempty"`
	ClientKey           string                 `xml:"clientKey,omitempty"`
	ClientCert          string                 `xml:"clientCert,omitempty"`
	DekGenerationId     int64                  `xml:"dekGenerationId,omitempty"`
	Changing            *bool                  `xml:"changing"`
	EraseDisksBeforeUse *bool                  `xml:"eraseDisksBeforeUse"`
}

func init() {
	types.Add("vsan:VsanHostEncryptionInfo", reflect.TypeOf((*VsanHostEncryptionInfo)(nil)).Elem())
}

type VsanHostFwComponent struct {
	types.DynamicData

	Name             string   `xml:"name"`
	Url              string   `xml:"url,omitempty"`
	Sha1sum          string   `xml:"sha1sum,omitempty"`
	CurrentVersion   string   `xml:"currentVersion,omitempty"`
	SuggestedVersion string   `xml:"suggestedVersion,omitempty"`
	ComponentID      []string `xml:"componentID,omitempty"`
}

func init() {
	types.Add("vsan:VsanHostFwComponent", reflect.TypeOf((*VsanHostFwComponent)(nil)).Elem())
}

type VsanHostGetRuntimeStats VsanHostGetRuntimeStatsRequestType

func init() {
	types.Add("vsan:VsanHostGetRuntimeStats", reflect.TypeOf((*VsanHostGetRuntimeStats)(nil)).Elem())
}

type VsanHostGetRuntimeStatsRequestType struct {
	This  types.ManagedObjectReference `xml:"_this"`
	Stats []string                     `xml:"stats,omitempty"`
}

func init() {
	types.Add("vsan:VsanHostGetRuntimeStatsRequestType", reflect.TypeOf((*VsanHostGetRuntimeStatsRequestType)(nil)).Elem())
}

type VsanHostGetRuntimeStatsResponse struct {
	Returnval VsanHostRuntimeStats `xml:"returnval"`
}

type VsanHostHclInfo struct {
	types.DynamicData

	Hostname    string                      `xml:"hostname"`
	HclChecked  bool                        `xml:"hclChecked"`
	ReleaseName string                      `xml:"releaseName,omitempty"`
	Error       *types.LocalizedMethodFault `xml:"error,omitempty"`
	Controllers []VsanHclControllerInfo     `xml:"controllers,omitempty"`
	Pnics       []VsanHclNicInfo            `xml:"pnics,omitempty"`
}

func init() {
	types.Add("vsan:VsanHostHclInfo", reflect.TypeOf((*VsanHostHclInfo)(nil)).Elem())
}

type VsanHostHealthSystemStatusResult struct {
	types.DynamicData

	Hostname string   `xml:"hostname"`
	Status   string   `xml:"status"`
	Issues   []string `xml:"issues,omitempty"`
}

func init() {
	types.Add("vsan:VsanHostHealthSystemStatusResult", reflect.TypeOf((*VsanHostHealthSystemStatusResult)(nil)).Elem())
}

type VsanHostHealthSystemVersionResult struct {
	types.DynamicData

	Hostname string                      `xml:"hostname"`
	Version  string                      `xml:"version,omitempty"`
	Error    *types.LocalizedMethodFault `xml:"error,omitempty"`
}

func init() {
	types.Add("vsan:VsanHostHealthSystemVersionResult", reflect.TypeOf((*VsanHostHealthSystemVersionResult)(nil)).Elem())
}

type VsanHostIpConfigEx struct {
	VsanHostIpConfig

	UpstreamIpV6Address   string `xml:"upstreamIpV6Address,omitempty"`
	DownstreamIpV6Address string `xml:"downstreamIpV6Address,omitempty"`
}

func init() {
	types.Add("vsan:VsanHostIpConfigEx", reflect.TypeOf((*VsanHostIpConfigEx)(nil)).Elem())
}

type VsanHostPerformResourceCheck VsanHostPerformResourceCheckRequestType

func init() {
	types.Add("vsan:VsanHostPerformResourceCheck", reflect.TypeOf((*VsanHostPerformResourceCheck)(nil)).Elem())
}

type VsanHostPerformResourceCheckRequestType struct {
	This              types.ManagedObjectReference `xml:"_this"`
	ResourceCheckSpec VsanResourceCheckSpec        `xml:"resourceCheckSpec"`
}

func init() {
	types.Add("vsan:VsanHostPerformResourceCheckRequestType", reflect.TypeOf((*VsanHostPerformResourceCheckRequestType)(nil)).Elem())
}

type VsanHostPerformResourceCheckResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanHostPortConfigEx struct {
	VsanHostConfigInfoNetworkInfoPortConfig

	TrafficTypes []string `xml:"trafficTypes,omitempty"`
}

func init() {
	types.Add("vsan:VsanHostPortConfigEx", reflect.TypeOf((*VsanHostPortConfigEx)(nil)).Elem())
}

type VsanHostPrepareVmdkLoadTest VsanHostPrepareVmdkLoadTestRequestType

func init() {
	types.Add("vsan:VsanHostPrepareVmdkLoadTest", reflect.TypeOf((*VsanHostPrepareVmdkLoadTest)(nil)).Elem())
}

type VsanHostPrepareVmdkLoadTestRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Runname string                       `xml:"runname"`
	Specs   []VsanVmdkLoadTestSpec       `xml:"specs"`
}

func init() {
	types.Add("vsan:VsanHostPrepareVmdkLoadTestRequestType", reflect.TypeOf((*VsanHostPrepareVmdkLoadTestRequestType)(nil)).Elem())
}

type VsanHostPrepareVmdkLoadTestResponse struct {
	Returnval string `xml:"returnval"`
}

type VsanHostPropertyRetrieveIssue struct {
	VsanUpgradeSystemPreflightCheckIssue

	Hosts []types.ManagedObjectReference `xml:"hosts"`
}

func init() {
	types.Add("vsan:VsanHostPropertyRetrieveIssue", reflect.TypeOf((*VsanHostPropertyRetrieveIssue)(nil)).Elem())
}

type VsanHostQueryAdvCfg VsanHostQueryAdvCfgRequestType

func init() {
	types.Add("vsan:VsanHostQueryAdvCfg", reflect.TypeOf((*VsanHostQueryAdvCfg)(nil)).Elem())
}

type VsanHostQueryAdvCfgRequestType struct {
	This                 types.ManagedObjectReference `xml:"_this"`
	Options              []string                     `xml:"options"`
	IncludeAllAdvOptions *bool                        `xml:"includeAllAdvOptions"`
	NonDefaultOnly       *bool                        `xml:"nonDefaultOnly"`
}

func init() {
	types.Add("vsan:VsanHostQueryAdvCfgRequestType", reflect.TypeOf((*VsanHostQueryAdvCfgRequestType)(nil)).Elem())
}

type VsanHostQueryAdvCfgResponse struct {
	Returnval []types.BaseOptionValue `xml:"returnval,omitempty,typeattr"`
}

type VsanHostQueryCheckLimits VsanHostQueryCheckLimitsRequestType

func init() {
	types.Add("vsan:VsanHostQueryCheckLimits", reflect.TypeOf((*VsanHostQueryCheckLimits)(nil)).Elem())
}

type VsanHostQueryCheckLimitsRequestType struct {
	This types.ManagedObjectReference  `xml:"_this"`
	Spec *VsanHostQueryCheckLimitsSpec `xml:"spec,omitempty"`
}

func init() {
	types.Add("vsan:VsanHostQueryCheckLimitsRequestType", reflect.TypeOf((*VsanHostQueryCheckLimitsRequestType)(nil)).Elem())
}

type VsanHostQueryCheckLimitsResponse struct {
	Returnval VsanLimitHealthResult `xml:"returnval"`
}

type VsanHostQueryCheckLimitsSpec struct {
	types.DynamicData

	OptionTypes []string `xml:"optionTypes,omitempty"`
	FetchAll    bool     `xml:"fetchAll"`
}

func init() {
	types.Add("vsan:VsanHostQueryCheckLimitsSpec", reflect.TypeOf((*VsanHostQueryCheckLimitsSpec)(nil)).Elem())
}

type VsanHostQueryEncryptionHealthSummary VsanHostQueryEncryptionHealthSummaryRequestType

func init() {
	types.Add("vsan:VsanHostQueryEncryptionHealthSummary", reflect.TypeOf((*VsanHostQueryEncryptionHealthSummary)(nil)).Elem())
}

type VsanHostQueryEncryptionHealthSummaryRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

func init() {
	types.Add("vsan:VsanHostQueryEncryptionHealthSummaryRequestType", reflect.TypeOf((*VsanHostQueryEncryptionHealthSummaryRequestType)(nil)).Elem())
}

type VsanHostQueryEncryptionHealthSummaryResponse struct {
	Returnval VsanEncryptionHealthSummary `xml:"returnval"`
}

type VsanHostQueryFileServiceHealthSummary VsanHostQueryFileServiceHealthSummaryRequestType

func init() {
	types.Add("vsan:VsanHostQueryFileServiceHealthSummary", reflect.TypeOf((*VsanHostQueryFileServiceHealthSummary)(nil)).Elem())
}

type VsanHostQueryFileServiceHealthSummaryRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

func init() {
	types.Add("vsan:VsanHostQueryFileServiceHealthSummaryRequestType", reflect.TypeOf((*VsanHostQueryFileServiceHealthSummaryRequestType)(nil)).Elem())
}

type VsanHostQueryFileServiceHealthSummaryResponse struct {
	Returnval VsanFileServiceHealthSummary `xml:"returnval"`
}

type VsanHostQueryHealthSystemVersion VsanHostQueryHealthSystemVersionRequestType

func init() {
	types.Add("vsan:VsanHostQueryHealthSystemVersion", reflect.TypeOf((*VsanHostQueryHealthSystemVersion)(nil)).Elem())
}

type VsanHostQueryHealthSystemVersionRequestType struct {
	This           types.ManagedObjectReference `xml:"_this"`
	DisplayVersion *bool                        `xml:"displayVersion"`
}

func init() {
	types.Add("vsan:VsanHostQueryHealthSystemVersionRequestType", reflect.TypeOf((*VsanHostQueryHealthSystemVersionRequestType)(nil)).Elem())
}

type VsanHostQueryHealthSystemVersionResponse struct {
	Returnval string `xml:"returnval"`
}

type VsanHostQueryHostInfoByUuids VsanHostQueryHostInfoByUuidsRequestType

func init() {
	types.Add("vsan:VsanHostQueryHostInfoByUuids", reflect.TypeOf((*VsanHostQueryHostInfoByUuids)(nil)).Elem())
}

type VsanHostQueryHostInfoByUuidsRequestType struct {
	This  types.ManagedObjectReference `xml:"_this"`
	Uuids []string                     `xml:"uuids"`
}

func init() {
	types.Add("vsan:VsanHostQueryHostInfoByUuidsRequestType", reflect.TypeOf((*VsanHostQueryHostInfoByUuidsRequestType)(nil)).Elem())
}

type VsanHostQueryHostInfoByUuidsResponse struct {
	Returnval []VsanQueryResultHostInfo `xml:"returnval"`
}

type VsanHostQueryObjectHealthSummary VsanHostQueryObjectHealthSummaryRequestType

func init() {
	types.Add("vsan:VsanHostQueryObjectHealthSummary", reflect.TypeOf((*VsanHostQueryObjectHealthSummary)(nil)).Elem())
}

type VsanHostQueryObjectHealthSummaryRequestType struct {
	This                          types.ManagedObjectReference `xml:"_this"`
	ObjUuids                      []string                     `xml:"objUuids,omitempty"`
	IncludeObjUuids               *bool                        `xml:"includeObjUuids"`
	LocalHostOnly                 *bool                        `xml:"localHostOnly"`
	IncludeNonComplianceObjDetail *bool                        `xml:"includeNonComplianceObjDetail"`
}

func init() {
	types.Add("vsan:VsanHostQueryObjectHealthSummaryRequestType", reflect.TypeOf((*VsanHostQueryObjectHealthSummaryRequestType)(nil)).Elem())
}

type VsanHostQueryObjectHealthSummaryResponse struct {
	Returnval VsanObjectOverallHealth `xml:"returnval"`
}

type VsanHostQueryPhysicalDiskHealthSummary VsanHostQueryPhysicalDiskHealthSummaryRequestType

func init() {
	types.Add("vsan:VsanHostQueryPhysicalDiskHealthSummary", reflect.TypeOf((*VsanHostQueryPhysicalDiskHealthSummary)(nil)).Elem())
}

type VsanHostQueryPhysicalDiskHealthSummaryRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

func init() {
	types.Add("vsan:VsanHostQueryPhysicalDiskHealthSummaryRequestType", reflect.TypeOf((*VsanHostQueryPhysicalDiskHealthSummaryRequestType)(nil)).Elem())
}

type VsanHostQueryPhysicalDiskHealthSummaryResponse struct {
	Returnval VsanPhysicalDiskHealthSummary `xml:"returnval"`
}

type VsanHostQueryRunIperfClient VsanHostQueryRunIperfClientRequestType

func init() {
	types.Add("vsan:VsanHostQueryRunIperfClient", reflect.TypeOf((*VsanHostQueryRunIperfClient)(nil)).Elem())
}

type VsanHostQueryRunIperfClientRequestType struct {
	This        types.ManagedObjectReference `xml:"_this"`
	Multicast   bool                         `xml:"multicast"`
	ServerIp    string                       `xml:"serverIp"`
	DurationSec int32                        `xml:"durationSec,omitempty"`
}

func init() {
	types.Add("vsan:VsanHostQueryRunIperfClientRequestType", reflect.TypeOf((*VsanHostQueryRunIperfClientRequestType)(nil)).Elem())
}

type VsanHostQueryRunIperfClientResponse struct {
	Returnval VsanNetworkLoadTestResult `xml:"returnval"`
}

type VsanHostQueryRunIperfServer VsanHostQueryRunIperfServerRequestType

func init() {
	types.Add("vsan:VsanHostQueryRunIperfServer", reflect.TypeOf((*VsanHostQueryRunIperfServer)(nil)).Elem())
}

type VsanHostQueryRunIperfServerRequestType struct {
	This        types.ManagedObjectReference `xml:"_this"`
	Multicast   bool                         `xml:"multicast"`
	ServerIp    string                       `xml:"serverIp,omitempty"`
	DurationSec int32                        `xml:"durationSec,omitempty"`
}

func init() {
	types.Add("vsan:VsanHostQueryRunIperfServerRequestType", reflect.TypeOf((*VsanHostQueryRunIperfServerRequestType)(nil)).Elem())
}

type VsanHostQueryRunIperfServerResponse struct {
	Returnval VsanNetworkLoadTestResult `xml:"returnval"`
}

type VsanHostQuerySmartStats VsanHostQuerySmartStatsRequestType

func init() {
	types.Add("vsan:VsanHostQuerySmartStats", reflect.TypeOf((*VsanHostQuerySmartStats)(nil)).Elem())
}

type VsanHostQuerySmartStatsRequestType struct {
	This            types.ManagedObjectReference `xml:"_this"`
	Disks           []string                     `xml:"disks,omitempty"`
	IncludeAllDisks *bool                        `xml:"includeAllDisks"`
}

func init() {
	types.Add("vsan:VsanHostQuerySmartStatsRequestType", reflect.TypeOf((*VsanHostQuerySmartStatsRequestType)(nil)).Elem())
}

type VsanHostQuerySmartStatsResponse struct {
	Returnval VsanSmartStatsHostSummary `xml:"returnval"`
}

type VsanHostQueryVerifyNetworkSettings VsanHostQueryVerifyNetworkSettingsRequestType

func init() {
	types.Add("vsan:VsanHostQueryVerifyNetworkSettings", reflect.TypeOf((*VsanHostQueryVerifyNetworkSettings)(nil)).Elem())
}

type VsanHostQueryVerifyNetworkSettingsRequestType struct {
	This                          types.ManagedObjectReference `xml:"_this"`
	Peers                         []string                     `xml:"peers,omitempty"`
	ROBOStretchedClusterWitnesses []string                     `xml:"ROBOStretchedClusterWitnesses,omitempty"`
	VMotionPeers                  []string                     `xml:"vMotionPeers,omitempty"`
}

func init() {
	types.Add("vsan:VsanHostQueryVerifyNetworkSettingsRequestType", reflect.TypeOf((*VsanHostQueryVerifyNetworkSettingsRequestType)(nil)).Elem())
}

type VsanHostQueryVerifyNetworkSettingsResponse struct {
	Returnval VsanNetworkHealthResult `xml:"returnval"`
}

type VsanHostReference struct {
	types.DynamicData

	Hostname string `xml:"hostname"`
}

func init() {
	types.Add("vsan:VsanHostReference", reflect.TypeOf((*VsanHostReference)(nil)).Elem())
}

type VsanHostRepairImmediateObjects VsanHostRepairImmediateObjectsRequestType

func init() {
	types.Add("vsan:VsanHostRepairImmediateObjects", reflect.TypeOf((*VsanHostRepairImmediateObjects)(nil)).Elem())
}

type VsanHostRepairImmediateObjectsRequestType struct {
	This       types.ManagedObjectReference `xml:"_this"`
	Uuids      []string                     `xml:"uuids,omitempty"`
	RepairType string                       `xml:"repairType,omitempty"`
}

func init() {
	types.Add("vsan:VsanHostRepairImmediateObjectsRequestType", reflect.TypeOf((*VsanHostRepairImmediateObjectsRequestType)(nil)).Elem())
}

type VsanHostRepairImmediateObjectsResponse struct {
	Returnval VsanRepairObjectsResult `xml:"returnval"`
}

type VsanHostResourceCheckResult struct {
	EntityResourceCheckDetails

	Host       *types.ManagedObjectReference      `xml:"host,omitempty"`
	DiskGroups []VsanDiskGroupResourceCheckResult `xml:"diskGroups,omitempty"`
}

func init() {
	types.Add("vsan:VsanHostResourceCheckResult", reflect.TypeOf((*VsanHostResourceCheckResult)(nil)).Elem())
}

type VsanHostRunVmdkLoadTest VsanHostRunVmdkLoadTestRequestType

func init() {
	types.Add("vsan:VsanHostRunVmdkLoadTest", reflect.TypeOf((*VsanHostRunVmdkLoadTest)(nil)).Elem())
}

type VsanHostRunVmdkLoadTestRequestType struct {
	This        types.ManagedObjectReference `xml:"_this"`
	Runname     string                       `xml:"runname"`
	DurationSec int32                        `xml:"durationSec"`
	Specs       []VsanVmdkLoadTestSpec       `xml:"specs"`
}

func init() {
	types.Add("vsan:VsanHostRunVmdkLoadTestRequestType", reflect.TypeOf((*VsanHostRunVmdkLoadTestRequestType)(nil)).Elem())
}

type VsanHostRunVmdkLoadTestResponse struct {
	Returnval []VsanVmdkLoadTestResult `xml:"returnval"`
}

type VsanHostRuntimeStats struct {
	types.DynamicData

	ResyncIopsInfo       *ResyncIopsInfo       `xml:"resyncIopsInfo,omitempty"`
	ConfigGeneration     *VsanConfigGeneration `xml:"configGeneration,omitempty"`
	SupportedClusterSize int32                 `xml:"supportedClusterSize,omitempty"`
	RepairTimerInfo      *RepairTimerInfo      `xml:"repairTimerInfo,omitempty"`
}

func init() {
	types.Add("vsan:VsanHostRuntimeStats", reflect.TypeOf((*VsanHostRuntimeStats)(nil)).Elem())
}

type VsanHostUpdateFirmware VsanHostUpdateFirmwareRequestType

func init() {
	types.Add("vsan:VsanHostUpdateFirmware", reflect.TypeOf((*VsanHostUpdateFirmware)(nil)).Elem())
}

type VsanHostUpdateFirmwareRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
	Host types.ManagedObjectReference `xml:"host"`
}

func init() {
	types.Add("vsan:VsanHostUpdateFirmwareRequestType", reflect.TypeOf((*VsanHostUpdateFirmwareRequestType)(nil)).Elem())
}

type VsanHostUpdateFirmwareResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanHostVirtualApplianceInfo struct {
	types.DynamicData

	HostKey      types.ManagedObjectReference `xml:"hostKey"`
	IsVirtualApp bool                         `xml:"isVirtualApp"`
}

func init() {
	types.Add("vsan:VsanHostVirtualApplianceInfo", reflect.TypeOf((*VsanHostVirtualApplianceInfo)(nil)).Elem())
}

type VsanHostVmdkLoadTestResult struct {
	types.DynamicData

	Hostname     string                   `xml:"hostname"`
	IssueFound   bool                     `xml:"issueFound"`
	FaultMessage string                   `xml:"faultMessage,omitempty"`
	VmdkResults  []VsanVmdkLoadTestResult `xml:"vmdkResults,omitempty"`
}

func init() {
	types.Add("vsan:VsanHostVmdkLoadTestResult", reflect.TypeOf((*VsanHostVmdkLoadTestResult)(nil)).Elem())
}

type VsanHostVsanObjectSyncQueryResult struct {
	types.DynamicData

	TotalObjectsToSync           int64                             `xml:"totalObjectsToSync,omitempty"`
	TotalBytesToSync             int64                             `xml:"totalBytesToSync,omitempty"`
	TotalRecoveryETA             int64                             `xml:"totalRecoveryETA,omitempty"`
	Objects                      []VsanHostVsanObjectSyncState     `xml:"objects,omitempty"`
	SyncingObjectRecoveryDetails *VsanSyncingObjectRecoveryDetails `xml:"syncingObjectRecoveryDetails,omitempty"`
}

func init() {
	types.Add("vsan:VsanHostVsanObjectSyncQueryResult", reflect.TypeOf((*VsanHostVsanObjectSyncQueryResult)(nil)).Elem())
}

type VsanHostVsanObjectSyncState struct {
	types.DynamicData

	Uuid       string                       `xml:"uuid"`
	Components []VsanHostComponentSyncState `xml:"components"`
}

func init() {
	types.Add("vsan:VsanHostVsanObjectSyncState", reflect.TypeOf((*VsanHostVsanObjectSyncState)(nil)).Elem())
}

type VsanHostWithHybridDiskgroupIssue struct {
	VsanUpgradeSystemPreflightCheckIssue

	Hosts []types.ManagedObjectReference `xml:"hosts"`
}

func init() {
	types.Add("vsan:VsanHostWithHybridDiskgroupIssue", reflect.TypeOf((*VsanHostWithHybridDiskgroupIssue)(nil)).Elem())
}

type VsanInternalExtendedConfig struct {
	types.DynamicData

	VcMaxDiskVersion int32 `xml:"vcMaxDiskVersion,omitempty"`
}

func init() {
	types.Add("vsan:VsanInternalExtendedConfig", reflect.TypeOf((*VsanInternalExtendedConfig)(nil)).Elem())
}

type VsanIscsiHomeObjectSpec struct {
	types.DynamicData

	StoragePolicy types.BaseVirtualMachineProfileSpec      `xml:"storagePolicy,omitempty,typeattr"`
	DefaultConfig *VsanIscsiTargetServiceDefaultConfigSpec `xml:"defaultConfig,omitempty"`
}

func init() {
	types.Add("vsan:VsanIscsiHomeObjectSpec", reflect.TypeOf((*VsanIscsiHomeObjectSpec)(nil)).Elem())
}

type VsanIscsiInitiatorGroup struct {
	types.DynamicData

	Name       string                         `xml:"name"`
	Initiators []string                       `xml:"initiators,omitempty"`
	Targets    []BaseVsanIscsiTargetBasicInfo `xml:"targets,omitempty,typeattr"`
}

func init() {
	types.Add("vsan:VsanIscsiInitiatorGroup", reflect.TypeOf((*VsanIscsiInitiatorGroup)(nil)).Elem())
}

type VsanIscsiLUN struct {
	VsanIscsiLUNCommonInfo

	TargetAlias       string                 `xml:"targetAlias"`
	Uuid              string                 `xml:"uuid"`
	ActualSize        int64                  `xml:"actualSize"`
	ObjectInformation *VsanObjectInformation `xml:"objectInformation,omitempty"`
}

func init() {
	types.Add("vsan:VsanIscsiLUN", reflect.TypeOf((*VsanIscsiLUN)(nil)).Elem())
}

type VsanIscsiLUNCommonInfo struct {
	types.DynamicData

	LunId   int32  `xml:"lunId,omitempty"`
	Alias   string `xml:"alias,omitempty"`
	LunSize int64  `xml:"lunSize"`
	Status  string `xml:"status,omitempty"`
}

func init() {
	types.Add("vsan:VsanIscsiLUNCommonInfo", reflect.TypeOf((*VsanIscsiLUNCommonInfo)(nil)).Elem())
}

type VsanIscsiLUNSpec struct {
	VsanIscsiLUNCommonInfo

	StoragePolicy types.BaseVirtualMachineProfileSpec `xml:"storagePolicy,omitempty,typeattr"`
	NewLunId      int32                               `xml:"newLunId,omitempty"`
}

func init() {
	types.Add("vsan:VsanIscsiLUNSpec", reflect.TypeOf((*VsanIscsiLUNSpec)(nil)).Elem())
}

type VsanIscsiTarget struct {
	VsanIscsiTargetCommonInfo

	LunCount          int32                  `xml:"lunCount,omitempty"`
	ObjectInformation *VsanObjectInformation `xml:"objectInformation,omitempty"`
	IoOwnerHost       string                 `xml:"ioOwnerHost,omitempty"`
	Initiators        []string               `xml:"initiators,omitempty"`
	InitiatorGroups   []string               `xml:"initiatorGroups,omitempty"`
}

func init() {
	types.Add("vsan:VsanIscsiTarget", reflect.TypeOf((*VsanIscsiTarget)(nil)).Elem())
}

type VsanIscsiTargetAuthSpec struct {
	types.DynamicData

	AuthType                    string `xml:"authType,omitempty"`
	UserNameAttachToTarget      string `xml:"userNameAttachToTarget,omitempty"`
	UserSecretAttachToTarget    string `xml:"userSecretAttachToTarget,omitempty"`
	UserNameAttachToInitiator   string `xml:"userNameAttachToInitiator,omitempty"`
	UserSecretAttachToInitiator string `xml:"userSecretAttachToInitiator,omitempty"`
}

func init() {
	types.Add("vsan:VsanIscsiTargetAuthSpec", reflect.TypeOf((*VsanIscsiTargetAuthSpec)(nil)).Elem())
}

type VsanIscsiTargetBasicInfo struct {
	types.DynamicData

	Alias string `xml:"alias"`
	Iqn   string `xml:"iqn,omitempty"`
}

func init() {
	types.Add("vsan:VsanIscsiTargetBasicInfo", reflect.TypeOf((*VsanIscsiTargetBasicInfo)(nil)).Elem())
}

type VsanIscsiTargetCommonInfo struct {
	VsanIscsiTargetBasicInfo

	AuthSpec         *VsanIscsiTargetAuthSpec `xml:"authSpec,omitempty"`
	Port             int32                    `xml:"port,omitempty"`
	NetworkInterface string                   `xml:"networkInterface,omitempty"`
}

func init() {
	types.Add("vsan:VsanIscsiTargetCommonInfo", reflect.TypeOf((*VsanIscsiTargetCommonInfo)(nil)).Elem())
}

type VsanIscsiTargetServiceConfig struct {
	types.DynamicData

	DefaultConfig *VsanIscsiTargetServiceDefaultConfigSpec `xml:"defaultConfig,omitempty"`
	Enabled       *bool                                    `xml:"enabled"`
}

func init() {
	types.Add("vsan:VsanIscsiTargetServiceConfig", reflect.TypeOf((*VsanIscsiTargetServiceConfig)(nil)).Elem())
}

type VsanIscsiTargetServiceDefaultConfigSpec struct {
	types.DynamicData

	NetworkInterface    string                   `xml:"networkInterface,omitempty"`
	Port                int32                    `xml:"port,omitempty"`
	IscsiTargetAuthSpec *VsanIscsiTargetAuthSpec `xml:"iscsiTargetAuthSpec,omitempty"`
}

func init() {
	types.Add("vsan:VsanIscsiTargetServiceDefaultConfigSpec", reflect.TypeOf((*VsanIscsiTargetServiceDefaultConfigSpec)(nil)).Elem())
}

type VsanIscsiTargetServiceSpec struct {
	VsanIscsiTargetServiceConfig

	HomeObjectStoragePolicy types.BaseVirtualMachineProfileSpec `xml:"homeObjectStoragePolicy,omitempty,typeattr"`
}

func init() {
	types.Add("vsan:VsanIscsiTargetServiceSpec", reflect.TypeOf((*VsanIscsiTargetServiceSpec)(nil)).Elem())
}

type VsanIscsiTargetSpec struct {
	VsanIscsiTargetCommonInfo

	StoragePolicy types.BaseVirtualMachineProfileSpec `xml:"storagePolicy,omitempty,typeattr"`
	NewAlias      string                              `xml:"newAlias,omitempty"`
}

func init() {
	types.Add("vsan:VsanIscsiTargetSpec", reflect.TypeOf((*VsanIscsiTargetSpec)(nil)).Elem())
}

type VsanJsonComparator struct {
	VsanComparator

	Comparator      string             `xml:"comparator,omitempty"`
	ComparableValue *types.KeyAnyValue `xml:"comparableValue,omitempty"`
}

func init() {
	types.Add("vsan:VsanJsonComparator", reflect.TypeOf((*VsanJsonComparator)(nil)).Elem())
}

type VsanJsonFilterRule struct {
	types.DynamicData

	FilterComparator BaseVsanComparator `xml:"filterComparator,omitempty,typeattr"`
	ComparablePath   []string           `xml:"comparablePath,omitempty"`
	KeysWithStrVal   []string           `xml:"keysWithStrVal,omitempty"`
	PropertyName     string             `xml:"propertyName,omitempty"`
}

func init() {
	types.Add("vsan:VsanJsonFilterRule", reflect.TypeOf((*VsanJsonFilterRule)(nil)).Elem())
}

type VsanKmsHealth struct {
	types.DynamicData

	ServerName     string                      `xml:"serverName"`
	Health         string                      `xml:"health"`
	Error          *types.LocalizedMethodFault `xml:"error,omitempty"`
	TrustHealth    string                      `xml:"trustHealth,omitempty"`
	CertHealth     string                      `xml:"certHealth,omitempty"`
	CertExpireDate *time.Time                  `xml:"certExpireDate"`
}

func init() {
	types.Add("vsan:VsanKmsHealth", reflect.TypeOf((*VsanKmsHealth)(nil)).Elem())
}

type VsanLimitHealthResult struct {
	types.DynamicData

	Hostname                   string `xml:"hostname,omitempty"`
	IssueFound                 bool   `xml:"issueFound"`
	MaxComponents              int32  `xml:"maxComponents"`
	FreeComponents             int32  `xml:"freeComponents"`
	ComponentLimitHealth       string `xml:"componentLimitHealth"`
	LowestFreeDiskSpacePct     int32  `xml:"lowestFreeDiskSpacePct"`
	UsedDiskSpaceB             int64  `xml:"usedDiskSpaceB"`
	TotalDiskSpaceB            int64  `xml:"totalDiskSpaceB"`
	DiskFreeSpaceHealth        string `xml:"diskFreeSpaceHealth"`
	ReservedRcSizeB            int64  `xml:"reservedRcSizeB"`
	TotalRcSizeB               int64  `xml:"totalRcSizeB"`
	RcFreeReservationHealth    string `xml:"rcFreeReservationHealth"`
	TotalLogicalSpaceB         int64  `xml:"totalLogicalSpaceB,omitempty"`
	LogicalSpaceUsedB          int64  `xml:"logicalSpaceUsedB,omitempty"`
	DedupMetadataSizeB         int64  `xml:"dedupMetadataSizeB,omitempty"`
	DiskTransientCapacityUsedB int64  `xml:"diskTransientCapacityUsedB,omitempty"`
	DgTransientCapacityUsedB   int64  `xml:"dgTransientCapacityUsedB,omitempty"`
}

func init() {
	types.Add("vsan:VsanLimitHealthResult", reflect.TypeOf((*VsanLimitHealthResult)(nil)).Elem())
}

type VsanMassCollectorPropertyParams struct {
	types.DynamicData

	PropertyName   string              `xml:"propertyName,omitempty"`
	PropertyParams []types.KeyAnyValue `xml:"propertyParams,omitempty"`
}

func init() {
	types.Add("vsan:VsanMassCollectorPropertyParams", reflect.TypeOf((*VsanMassCollectorPropertyParams)(nil)).Elem())
}

type VsanMassCollectorSpec struct {
	types.DynamicData

	Objects          []types.ManagedObjectReference    `xml:"objects,omitempty"`
	ObjectCollection string                            `xml:"objectCollection,omitempty"`
	Properties       []string                          `xml:"properties"`
	PropertiesParams []VsanMassCollectorPropertyParams `xml:"propertiesParams,omitempty"`
	Constraint       BaseVsanResourceConstraint        `xml:"constraint,omitempty,typeattr"`
}

func init() {
	types.Add("vsan:VsanMassCollectorSpec", reflect.TypeOf((*VsanMassCollectorSpec)(nil)).Elem())
}

type VsanMetricProfile struct {
	types.DynamicData

	AuthToken string `xml:"authToken"`
}

func init() {
	types.Add("vsan:VsanMetricProfile", reflect.TypeOf((*VsanMetricProfile)(nil)).Elem())
}

type VsanMetricsConfig struct {
	types.DynamicData

	Profiles []VsanMetricProfile `xml:"profiles,omitempty"`
}

func init() {
	types.Add("vsan:VsanMetricsConfig", reflect.TypeOf((*VsanMetricsConfig)(nil)).Elem())
}

type VsanMigrateVmsToVds VsanMigrateVmsToVdsRequestType

func init() {
	types.Add("vsan:VsanMigrateVmsToVds", reflect.TypeOf((*VsanMigrateVmsToVds)(nil)).Elem())
}

type VsanMigrateVmsToVdsRequestType struct {
	This          types.ManagedObjectReference `xml:"_this"`
	VmConfigSpecs []VsanVmVdsMigrationSpec     `xml:"vmConfigSpecs"`
	VdsUuid       string                       `xml:"vdsUuid"`
	TimeoutSec    int64                        `xml:"timeoutSec"`
	Revert        *bool                        `xml:"revert"`
}

func init() {
	types.Add("vsan:VsanMigrateVmsToVdsRequestType", reflect.TypeOf((*VsanMigrateVmsToVdsRequestType)(nil)).Elem())
}

type VsanMigrateVmsToVdsResponse struct {
	Returnval string `xml:"returnval"`
}

type VsanMixedEsxVersionIssue struct {
	VsanUpgradeSystemPreflightCheckIssue
}

func init() {
	types.Add("vsan:VsanMixedEsxVersionIssue", reflect.TypeOf((*VsanMixedEsxVersionIssue)(nil)).Elem())
}

type VsanNestJsonComparator struct {
	VsanComparator

	NestedComparators []VsanJsonComparator `xml:"nestedComparators,omitempty"`
	Conjoiner         string               `xml:"conjoiner,omitempty"`
}

func init() {
	types.Add("vsan:VsanNestJsonComparator", reflect.TypeOf((*VsanNestJsonComparator)(nil)).Elem())
}

type VsanNetworkConfigBaseIssue struct {
	types.DynamicData
}

func init() {
	types.Add("vsan:VsanNetworkConfigBaseIssue", reflect.TypeOf((*VsanNetworkConfigBaseIssue)(nil)).Elem())
}

type VsanNetworkConfigBestPracticeHealth struct {
	types.DynamicData

	VdsPresent bool                             `xml:"vdsPresent"`
	Issues     []BaseVsanNetworkConfigBaseIssue `xml:"issues,omitempty,typeattr"`
}

func init() {
	types.Add("vsan:VsanNetworkConfigBestPracticeHealth", reflect.TypeOf((*VsanNetworkConfigBestPracticeHealth)(nil)).Elem())
}

type VsanNetworkConfigPnicSpeedInconsistencyIssue struct {
	VsanNetworkConfigBaseIssue

	Host        types.ManagedObjectReference  `xml:"host"`
	VswitchName string                        `xml:"vswitchName,omitempty"`
	Vds         *types.ManagedObjectReference `xml:"vds,omitempty"`
	SpeedsMb    []int64                       `xml:"speedsMb"`
}

func init() {
	types.Add("vsan:VsanNetworkConfigPnicSpeedInconsistencyIssue", reflect.TypeOf((*VsanNetworkConfigPnicSpeedInconsistencyIssue)(nil)).Elem())
}

type VsanNetworkConfigPortgroupWithNoRedundancyIssue struct {
	VsanNetworkConfigBaseIssue

	Host          types.ManagedObjectReference  `xml:"host"`
	PortgroupName string                        `xml:"portgroupName,omitempty"`
	Vds           *types.ManagedObjectReference `xml:"vds,omitempty"`
	Pg            *types.ManagedObjectReference `xml:"pg,omitempty"`
	NumPnics      int64                         `xml:"numPnics"`
}

func init() {
	types.Add("vsan:VsanNetworkConfigPortgroupWithNoRedundancyIssue", reflect.TypeOf((*VsanNetworkConfigPortgroupWithNoRedundancyIssue)(nil)).Elem())
}

type VsanNetworkConfigVdsScopeIssue struct {
	VsanNetworkConfigBaseIssue

	Vds            types.ManagedObjectReference   `xml:"vds"`
	MemberHosts    []types.ManagedObjectReference `xml:"memberHosts"`
	NonMemberHosts []types.ManagedObjectReference `xml:"nonMemberHosts"`
}

func init() {
	types.Add("vsan:VsanNetworkConfigVdsScopeIssue", reflect.TypeOf((*VsanNetworkConfigVdsScopeIssue)(nil)).Elem())
}

type VsanNetworkConfigVsanNotOnVdsIssue struct {
	VsanNetworkConfigBaseIssue

	Host   types.ManagedObjectReference `xml:"host"`
	Vmknic string                       `xml:"vmknic"`
}

func init() {
	types.Add("vsan:VsanNetworkConfigVsanNotOnVdsIssue", reflect.TypeOf((*VsanNetworkConfigVsanNotOnVdsIssue)(nil)).Elem())
}

type VsanNetworkConfigVswitchWithNoRedundancyIssue struct {
	VsanNetworkConfigBaseIssue

	Host        types.ManagedObjectReference  `xml:"host"`
	VswitchName string                        `xml:"vswitchName,omitempty"`
	Vds         *types.ManagedObjectReference `xml:"vds,omitempty"`
	NumPnics    int64                         `xml:"numPnics"`
}

func init() {
	types.Add("vsan:VsanNetworkConfigVswitchWithNoRedundancyIssue", reflect.TypeOf((*VsanNetworkConfigVswitchWithNoRedundancyIssue)(nil)).Elem())
}

type VsanNetworkHealthResult struct {
	types.DynamicData

	Host              *types.ManagedObjectReference `xml:"host,omitempty"`
	Hostname          string                        `xml:"hostname,omitempty"`
	VsanVmknicPresent *bool                         `xml:"vsanVmknicPresent"`
	IpSubnets         []string                      `xml:"ipSubnets,omitempty"`
	IssueFound        *bool                         `xml:"issueFound"`
	PeerHealth        []VsanNetworkPeerHealthResult `xml:"peerHealth,omitempty"`
	VMotionHealth     []VsanNetworkPeerHealthResult `xml:"vMotionHealth,omitempty"`
	MulticastConfig   string                        `xml:"multicastConfig,omitempty"`
	UnicastConfig     string                        `xml:"unicastConfig,omitempty"`
	InUnicast         *bool                         `xml:"inUnicast"`
}

func init() {
	types.Add("vsan:VsanNetworkHealthResult", reflect.TypeOf((*VsanNetworkHealthResult)(nil)).Elem())
}

type VsanNetworkLoadTestResult struct {
	types.DynamicData

	Hostname      string  `xml:"hostname"`
	Status        string  `xml:"status,omitempty"`
	Client        bool    `xml:"client"`
	BandwidthBps  int64   `xml:"bandwidthBps"`
	TotalBytes    int64   `xml:"totalBytes"`
	LostDatagrams int64   `xml:"lostDatagrams,omitempty"`
	LossPct       int64   `xml:"lossPct,omitempty"`
	SentDatagrams int64   `xml:"sentDatagrams,omitempty"`
	JitterMs      float32 `xml:"jitterMs,omitempty"`
}

func init() {
	types.Add("vsan:VsanNetworkLoadTestResult", reflect.TypeOf((*VsanNetworkLoadTestResult)(nil)).Elem())
}

type VsanNetworkPeerHealthResult struct {
	types.DynamicData

	Peer                    string `xml:"peer,omitempty"`
	PeerHostname            string `xml:"peerHostname,omitempty"`
	PeerVmknicName          string `xml:"peerVmknicName,omitempty"`
	SmallPingTestSuccessPct int32  `xml:"smallPingTestSuccessPct,omitempty"`
	LargePingTestSuccessPct int32  `xml:"largePingTestSuccessPct,omitempty"`
	MaxLatencyUs            int64  `xml:"maxLatencyUs,omitempty"`
	OnSameIpSubnet          *bool  `xml:"onSameIpSubnet"`
	SourceVmknicName        string `xml:"sourceVmknicName,omitempty"`
}

func init() {
	types.Add("vsan:VsanNetworkPeerHealthResult", reflect.TypeOf((*VsanNetworkPeerHealthResult)(nil)).Elem())
}

type VsanNetworkVMotionVmknicNotFountIssue struct {
	VsanNetworkConfigBaseIssue

	HostWithoutVmotionVmknic types.ManagedObjectReference `xml:"hostWithoutVmotionVmknic"`
}

func init() {
	types.Add("vsan:VsanNetworkVMotionVmknicNotFountIssue", reflect.TypeOf((*VsanNetworkVMotionVmknicNotFountIssue)(nil)).Elem())
}

type VsanNodeNotMaster struct {
	types.VimFault

	VsanMasterUuid               string `xml:"vsanMasterUuid,omitempty"`
	CmmdsMasterButNotStatsMaster *bool  `xml:"cmmdsMasterButNotStatsMaster"`
}

func init() {
	types.Add("vsan:VsanNodeNotMaster", reflect.TypeOf((*VsanNodeNotMaster)(nil)).Elem())
}

type VsanNodeNotMasterFault VsanNodeNotMaster

func init() {
	types.Add("vsan:VsanNodeNotMasterFault", reflect.TypeOf((*VsanNodeNotMasterFault)(nil)).Elem())
}

type VsanObjectExtraAttributes struct {
	types.DynamicData

	Uuid     string `xml:"uuid"`
	ObjPath  string `xml:"objPath"`
	ObjClass int32  `xml:"objClass"`
	Ufn      string `xml:"ufn"`
	IsHbrCfg bool   `xml:"isHbrCfg"`
}

func init() {
	types.Add("vsan:VsanObjectExtraAttributes", reflect.TypeOf((*VsanObjectExtraAttributes)(nil)).Elem())
}

type VsanObjectHealth struct {
	types.DynamicData

	NumObjects int32    `xml:"numObjects"`
	Health     string   `xml:"health"`
	ObjUuids   []string `xml:"objUuids,omitempty"`
}

func init() {
	types.Add("vsan:VsanObjectHealth", reflect.TypeOf((*VsanObjectHealth)(nil)).Elem())
}

type VsanObjectIdentity struct {
	types.DynamicData

	Uuid           string                        `xml:"uuid"`
	Type           string                        `xml:"type"`
	VmInstanceUuid string                        `xml:"vmInstanceUuid,omitempty"`
	VmNsObjectUuid string                        `xml:"vmNsObjectUuid,omitempty"`
	Vm             *types.ManagedObjectReference `xml:"vm,omitempty"`
	Description    string                        `xml:"description,omitempty"`
}

func init() {
	types.Add("vsan:VsanObjectIdentity", reflect.TypeOf((*VsanObjectIdentity)(nil)).Elem())
}

type VsanObjectIdentityAndHealth struct {
	types.DynamicData

	Identities   []VsanObjectIdentity     `xml:"identities,omitempty"`
	Health       *VsanObjectOverallHealth `xml:"health,omitempty"`
	SpaceSummary []VsanObjectSpaceSummary `xml:"spaceSummary,omitempty"`
	RawData      string                   `xml:"rawData,omitempty"`
}

func init() {
	types.Add("vsan:VsanObjectIdentityAndHealth", reflect.TypeOf((*VsanObjectIdentityAndHealth)(nil)).Elem())
}

type VsanObjectInaccessibleIssue struct {
	VsanUpgradeSystemPreflightCheckIssue

	Uuids []string `xml:"uuids"`
}

func init() {
	types.Add("vsan:VsanObjectInaccessibleIssue", reflect.TypeOf((*VsanObjectInaccessibleIssue)(nil)).Elem())
}

type VsanObjectInformation struct {
	types.DynamicData

	DirectoryName           string                       `xml:"directoryName,omitempty"`
	VsanObjectUuid          string                       `xml:"vsanObjectUuid,omitempty"`
	VsanHealth              string                       `xml:"vsanHealth,omitempty"`
	PolicyAttributes        []types.KeyValue             `xml:"policyAttributes,omitempty"`
	SpbmProfileUuid         string                       `xml:"spbmProfileUuid,omitempty"`
	SpbmProfileGenerationId string                       `xml:"spbmProfileGenerationId,omitempty"`
	SpbmComplianceResult    *VsanStorageComplianceResult `xml:"spbmComplianceResult,omitempty"`
}

func init() {
	types.Add("vsan:VsanObjectInformation", reflect.TypeOf((*VsanObjectInformation)(nil)).Elem())
}

type VsanObjectOverallHealth struct {
	types.DynamicData

	ObjectHealthDetail      []VsanObjectHealth            `xml:"objectHealthDetail,omitempty"`
	ObjectsComplianceDetail []VsanStorageComplianceResult `xml:"objectsComplianceDetail,omitempty"`
	ObjectVersionCompliance *bool                         `xml:"objectVersionCompliance"`
}

func init() {
	types.Add("vsan:VsanObjectOverallHealth", reflect.TypeOf((*VsanObjectOverallHealth)(nil)).Elem())
}

type VsanObjectPolicyIssue struct {
	VsanUpgradeSystemPreflightCheckIssue

	Uuids []string `xml:"uuids"`
}

func init() {
	types.Add("vsan:VsanObjectPolicyIssue", reflect.TypeOf((*VsanObjectPolicyIssue)(nil)).Elem())
}

type VsanObjectProfileInfo struct {
	types.DynamicData

	VsanObjectUuid           string `xml:"vsanObjectUuid"`
	SpbmProfileId            string `xml:"spbmProfileId"`
	SpbmProfileGenerationNum int32  `xml:"spbmProfileGenerationNum"`
}

func init() {
	types.Add("vsan:VsanObjectProfileInfo", reflect.TypeOf((*VsanObjectProfileInfo)(nil)).Elem())
}

type VsanObjectQuerySpec struct {
	types.DynamicData

	Uuid                    string `xml:"uuid"`
	SpbmProfileGenerationId string `xml:"spbmProfileGenerationId,omitempty"`
}

func init() {
	types.Add("vsan:VsanObjectQuerySpec", reflect.TypeOf((*VsanObjectQuerySpec)(nil)).Elem())
}

type VsanObjectSpaceSummary struct {
	types.DynamicData

	ObjType            string `xml:"objType,omitempty"`
	OverheadB          int64  `xml:"overheadB,omitempty"`
	TemporaryOverheadB int64  `xml:"temporaryOverheadB,omitempty"`
	PrimaryCapacityB   int64  `xml:"primaryCapacityB,omitempty"`
	ProvisionCapacityB int64  `xml:"provisionCapacityB,omitempty"`
	ReservedCapacityB  int64  `xml:"reservedCapacityB,omitempty"`
	OverReservedB      int64  `xml:"overReservedB,omitempty"`
	PhysicalUsedB      int64  `xml:"physicalUsedB,omitempty"`
	UsedB              int64  `xml:"usedB,omitempty"`
}

func init() {
	types.Add("vsan:VsanObjectSpaceSummary", reflect.TypeOf((*VsanObjectSpaceSummary)(nil)).Elem())
}

type VsanObjectTypeRule struct {
	types.DynamicData

	ObjectType string   `xml:"objectType,omitempty"`
	Attributes []string `xml:"attributes,omitempty"`
}

func init() {
	types.Add("vsan:VsanObjectTypeRule", reflect.TypeOf((*VsanObjectTypeRule)(nil)).Elem())
}

type VsanPerfCreateStatsObject VsanPerfCreateStatsObjectRequestType

func init() {
	types.Add("vsan:VsanPerfCreateStatsObject", reflect.TypeOf((*VsanPerfCreateStatsObject)(nil)).Elem())
}

type VsanPerfCreateStatsObjectRequestType struct {
	This    types.ManagedObjectReference        `xml:"_this"`
	Cluster *types.ManagedObjectReference       `xml:"cluster,omitempty"`
	Profile types.BaseVirtualMachineProfileSpec `xml:"profile,omitempty,typeattr"`
}

func init() {
	types.Add("vsan:VsanPerfCreateStatsObjectRequestType", reflect.TypeOf((*VsanPerfCreateStatsObjectRequestType)(nil)).Elem())
}

type VsanPerfCreateStatsObjectResponse struct {
	Returnval string `xml:"returnval"`
}

type VsanPerfCreateStatsObjectTask VsanPerfCreateStatsObjectTaskRequestType

func init() {
	types.Add("vsan:VsanPerfCreateStatsObjectTask", reflect.TypeOf((*VsanPerfCreateStatsObjectTask)(nil)).Elem())
}

type VsanPerfCreateStatsObjectTaskRequestType struct {
	This    types.ManagedObjectReference        `xml:"_this"`
	Cluster *types.ManagedObjectReference       `xml:"cluster,omitempty"`
	Profile types.BaseVirtualMachineProfileSpec `xml:"profile,omitempty,typeattr"`
}

func init() {
	types.Add("vsan:VsanPerfCreateStatsObjectTaskRequestType", reflect.TypeOf((*VsanPerfCreateStatsObjectTaskRequestType)(nil)).Elem())
}

type VsanPerfCreateStatsObjectTaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanPerfDeleteStatsObject VsanPerfDeleteStatsObjectRequestType

func init() {
	types.Add("vsan:VsanPerfDeleteStatsObject", reflect.TypeOf((*VsanPerfDeleteStatsObject)(nil)).Elem())
}

type VsanPerfDeleteStatsObjectRequestType struct {
	This    types.ManagedObjectReference  `xml:"_this"`
	Cluster *types.ManagedObjectReference `xml:"cluster,omitempty"`
}

func init() {
	types.Add("vsan:VsanPerfDeleteStatsObjectRequestType", reflect.TypeOf((*VsanPerfDeleteStatsObjectRequestType)(nil)).Elem())
}

type VsanPerfDeleteStatsObjectResponse struct {
	Returnval bool `xml:"returnval"`
}

type VsanPerfDeleteStatsObjectTask VsanPerfDeleteStatsObjectTaskRequestType

func init() {
	types.Add("vsan:VsanPerfDeleteStatsObjectTask", reflect.TypeOf((*VsanPerfDeleteStatsObjectTask)(nil)).Elem())
}

type VsanPerfDeleteStatsObjectTaskRequestType struct {
	This    types.ManagedObjectReference  `xml:"_this"`
	Cluster *types.ManagedObjectReference `xml:"cluster,omitempty"`
}

func init() {
	types.Add("vsan:VsanPerfDeleteStatsObjectTaskRequestType", reflect.TypeOf((*VsanPerfDeleteStatsObjectTaskRequestType)(nil)).Elem())
}

type VsanPerfDeleteStatsObjectTaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanPerfDeleteTimeRange VsanPerfDeleteTimeRangeRequestType

func init() {
	types.Add("vsan:VsanPerfDeleteTimeRange", reflect.TypeOf((*VsanPerfDeleteTimeRange)(nil)).Elem())
}

type VsanPerfDeleteTimeRangeRequestType struct {
	This    types.ManagedObjectReference  `xml:"_this"`
	Cluster *types.ManagedObjectReference `xml:"cluster,omitempty"`
	Name    string                        `xml:"name"`
}

func init() {
	types.Add("vsan:VsanPerfDeleteTimeRangeRequestType", reflect.TypeOf((*VsanPerfDeleteTimeRangeRequestType)(nil)).Elem())
}

type VsanPerfDeleteTimeRangeResponse struct {
}

type VsanPerfDiagnose VsanPerfDiagnoseRequestType

func init() {
	types.Add("vsan:VsanPerfDiagnose", reflect.TypeOf((*VsanPerfDiagnose)(nil)).Elem())
}

type VsanPerfDiagnoseQuerySpec struct {
	types.DynamicData

	StartTime time.Time `xml:"startTime"`
	EndTime   time.Time `xml:"endTime"`
	QueryType string    `xml:"queryType"`
	Context   string    `xml:"context,omitempty"`
}

func init() {
	types.Add("vsan:VsanPerfDiagnoseQuerySpec", reflect.TypeOf((*VsanPerfDiagnoseQuerySpec)(nil)).Elem())
}

type VsanPerfDiagnoseRequestType struct {
	This              types.ManagedObjectReference  `xml:"_this"`
	PerfDiagnoseQuery VsanPerfDiagnoseQuerySpec     `xml:"perfDiagnoseQuery"`
	Cluster           *types.ManagedObjectReference `xml:"cluster,omitempty"`
}

func init() {
	types.Add("vsan:VsanPerfDiagnoseRequestType", reflect.TypeOf((*VsanPerfDiagnoseRequestType)(nil)).Elem())
}

type VsanPerfDiagnoseResponse struct {
	Returnval []VsanPerfDiagnosticResult `xml:"returnval,omitempty"`
}

type VsanPerfDiagnoseTask VsanPerfDiagnoseTaskRequestType

func init() {
	types.Add("vsan:VsanPerfDiagnoseTask", reflect.TypeOf((*VsanPerfDiagnoseTask)(nil)).Elem())
}

type VsanPerfDiagnoseTaskRequestType struct {
	This              types.ManagedObjectReference  `xml:"_this"`
	PerfDiagnoseQuery VsanPerfDiagnoseQuerySpec     `xml:"perfDiagnoseQuery"`
	Cluster           *types.ManagedObjectReference `xml:"cluster,omitempty"`
}

func init() {
	types.Add("vsan:VsanPerfDiagnoseTaskRequestType", reflect.TypeOf((*VsanPerfDiagnoseTaskRequestType)(nil)).Elem())
}

type VsanPerfDiagnoseTaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanPerfDiagnosticException struct {
	types.DynamicData

	ExceptionId      string `xml:"exceptionId"`
	ExceptionMessage string `xml:"exceptionMessage"`
	ExceptionDetails string `xml:"exceptionDetails"`
	ExceptionUrl     string `xml:"exceptionUrl"`
}

func init() {
	types.Add("vsan:VsanPerfDiagnosticException", reflect.TypeOf((*VsanPerfDiagnosticException)(nil)).Elem())
}

type VsanPerfDiagnosticResult struct {
	types.DynamicData

	ExceptionId         string                    `xml:"exceptionId"`
	Recommendation      string                    `xml:"recommendation,omitempty"`
	AggregationFunction string                    `xml:"aggregationFunction,omitempty"`
	AggregationData     *VsanPerfEntityMetricCSV  `xml:"aggregationData,omitempty"`
	ExceptionData       []VsanPerfEntityMetricCSV `xml:"exceptionData"`
}

func init() {
	types.Add("vsan:VsanPerfDiagnosticResult", reflect.TypeOf((*VsanPerfDiagnosticResult)(nil)).Elem())
}

type VsanPerfEntityMetricCSV struct {
	types.DynamicData

	EntityRefId string                    `xml:"entityRefId"`
	SampleInfo  string                    `xml:"sampleInfo,omitempty"`
	Value       []VsanPerfMetricSeriesCSV `xml:"value,omitempty"`
}

func init() {
	types.Add("vsan:VsanPerfEntityMetricCSV", reflect.TypeOf((*VsanPerfEntityMetricCSV)(nil)).Elem())
}

type VsanPerfEntityType struct {
	types.DynamicData

	Name        string          `xml:"name"`
	Id          string          `xml:"id"`
	Graphs      []VsanPerfGraph `xml:"graphs"`
	Description string          `xml:"description,omitempty"`
}

func init() {
	types.Add("vsan:VsanPerfEntityType", reflect.TypeOf((*VsanPerfEntityType)(nil)).Elem())
}

type VsanPerfGetAggregatedEntityTypes VsanPerfGetAggregatedEntityTypesRequestType

func init() {
	types.Add("vsan:VsanPerfGetAggregatedEntityTypes", reflect.TypeOf((*VsanPerfGetAggregatedEntityTypes)(nil)).Elem())
}

type VsanPerfGetAggregatedEntityTypesRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

func init() {
	types.Add("vsan:VsanPerfGetAggregatedEntityTypesRequestType", reflect.TypeOf((*VsanPerfGetAggregatedEntityTypesRequestType)(nil)).Elem())
}

type VsanPerfGetAggregatedEntityTypesResponse struct {
	Returnval []VsanPerfEntityType `xml:"returnval,omitempty"`
}

type VsanPerfGetSupportedDiagnosticExceptions VsanPerfGetSupportedDiagnosticExceptionsRequestType

func init() {
	types.Add("vsan:VsanPerfGetSupportedDiagnosticExceptions", reflect.TypeOf((*VsanPerfGetSupportedDiagnosticExceptions)(nil)).Elem())
}

type VsanPerfGetSupportedDiagnosticExceptionsRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

func init() {
	types.Add("vsan:VsanPerfGetSupportedDiagnosticExceptionsRequestType", reflect.TypeOf((*VsanPerfGetSupportedDiagnosticExceptionsRequestType)(nil)).Elem())
}

type VsanPerfGetSupportedDiagnosticExceptionsResponse struct {
	Returnval []VsanPerfDiagnosticException `xml:"returnval,omitempty"`
}

type VsanPerfGetSupportedEntityTypes VsanPerfGetSupportedEntityTypesRequestType

func init() {
	types.Add("vsan:VsanPerfGetSupportedEntityTypes", reflect.TypeOf((*VsanPerfGetSupportedEntityTypes)(nil)).Elem())
}

type VsanPerfGetSupportedEntityTypesRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

func init() {
	types.Add("vsan:VsanPerfGetSupportedEntityTypesRequestType", reflect.TypeOf((*VsanPerfGetSupportedEntityTypesRequestType)(nil)).Elem())
}

type VsanPerfGetSupportedEntityTypesResponse struct {
	Returnval []VsanPerfEntityType `xml:"returnval,omitempty"`
}

type VsanPerfGraph struct {
	types.DynamicData

	Id          string             `xml:"id"`
	Metrics     []VsanPerfMetricId `xml:"metrics"`
	Unit        string             `xml:"unit"`
	Threshold   *VsanPerfThreshold `xml:"threshold,omitempty"`
	Name        string             `xml:"name,omitempty"`
	Description string             `xml:"description,omitempty"`
}

func init() {
	types.Add("vsan:VsanPerfGraph", reflect.TypeOf((*VsanPerfGraph)(nil)).Elem())
}

type VsanPerfMasterInformation struct {
	types.DynamicData

	SecSinceLastStatsWrite     int64      `xml:"secSinceLastStatsWrite,omitempty"`
	SecSinceLastStatsCollect   int64      `xml:"secSinceLastStatsCollect,omitempty"`
	StatsIntervalSec           int64      `xml:"statsIntervalSec"`
	CollectionFailureHostUuids []string   `xml:"collectionFailureHostUuids,omitempty"`
	RenamedStatsDirectories    []string   `xml:"renamedStatsDirectories,omitempty"`
	StatsDirectoryPercentFree  int64      `xml:"statsDirectoryPercentFree,omitempty"`
	VerboseMode                *bool      `xml:"verboseMode"`
	VerboseModeLastUpdate      *time.Time `xml:"verboseModeLastUpdate"`
}

func init() {
	types.Add("vsan:VsanPerfMasterInformation", reflect.TypeOf((*VsanPerfMasterInformation)(nil)).Elem())
}

type VsanPerfMemberInfo struct {
	types.DynamicData

	Thumbprint          string                   `xml:"thumbprint"`
	MemberUuid          string                   `xml:"memberUuid,omitempty"`
	IsSupportUnicast    *bool                    `xml:"isSupportUnicast"`
	UnicastAddressInfos []VsanUnicastAddressInfo `xml:"unicastAddressInfos,omitempty"`
	Hostname            string                   `xml:"hostname,omitempty"`
}

func init() {
	types.Add("vsan:VsanPerfMemberInfo", reflect.TypeOf((*VsanPerfMemberInfo)(nil)).Elem())
}

type VsanPerfMetricId struct {
	types.DynamicData

	Label                  string `xml:"label"`
	Group                  string `xml:"group,omitempty"`
	RollupType             string `xml:"rollupType,omitempty"`
	StatsType              string `xml:"statsType,omitempty"`
	Name                   string `xml:"name,omitempty"`
	Description            string `xml:"description,omitempty"`
	MetricsCollectInterval int32  `xml:"metricsCollectInterval,omitempty"`
}

func init() {
	types.Add("vsan:VsanPerfMetricId", reflect.TypeOf((*VsanPerfMetricId)(nil)).Elem())
}

type VsanPerfMetricSeriesCSV struct {
	types.DynamicData

	MetricId      VsanPerfMetricId   `xml:"metricId"`
	Threshold     *VsanPerfThreshold `xml:"threshold,omitempty"`
	NumExceptions string             `xml:"numExceptions,omitempty"`
	Values        string             `xml:"values,omitempty"`
}

func init() {
	types.Add("vsan:VsanPerfMetricSeriesCSV", reflect.TypeOf((*VsanPerfMetricSeriesCSV)(nil)).Elem())
}

type VsanPerfNodeInformation struct {
	types.DynamicData

	Version        string                      `xml:"version"`
	Hostname       string                      `xml:"hostname,omitempty"`
	Error          *types.LocalizedMethodFault `xml:"error,omitempty"`
	IsCmmdsMaster  bool                        `xml:"isCmmdsMaster"`
	IsStatsMaster  bool                        `xml:"isStatsMaster"`
	VsanMasterUuid string                      `xml:"vsanMasterUuid,omitempty"`
	VsanNodeUuid   string                      `xml:"vsanNodeUuid,omitempty"`
	MasterInfo     *VsanPerfMasterInformation  `xml:"masterInfo,omitempty"`
	DiagnosticMode *bool                       `xml:"diagnosticMode"`
}

func init() {
	types.Add("vsan:VsanPerfNodeInformation", reflect.TypeOf((*VsanPerfNodeInformation)(nil)).Elem())
}

type VsanPerfQueryClusterHealth VsanPerfQueryClusterHealthRequestType

func init() {
	types.Add("vsan:VsanPerfQueryClusterHealth", reflect.TypeOf((*VsanPerfQueryClusterHealth)(nil)).Elem())
}

type VsanPerfQueryClusterHealthRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

func init() {
	types.Add("vsan:VsanPerfQueryClusterHealthRequestType", reflect.TypeOf((*VsanPerfQueryClusterHealthRequestType)(nil)).Elem())
}

type VsanPerfQueryClusterHealthResponse struct {
	Returnval []types.BaseDynamicData `xml:"returnval,typeattr"`
}

type VsanPerfQueryNodeInformation VsanPerfQueryNodeInformationRequestType

func init() {
	types.Add("vsan:VsanPerfQueryNodeInformation", reflect.TypeOf((*VsanPerfQueryNodeInformation)(nil)).Elem())
}

type VsanPerfQueryNodeInformationRequestType struct {
	This    types.ManagedObjectReference  `xml:"_this"`
	Cluster *types.ManagedObjectReference `xml:"cluster,omitempty"`
}

func init() {
	types.Add("vsan:VsanPerfQueryNodeInformationRequestType", reflect.TypeOf((*VsanPerfQueryNodeInformationRequestType)(nil)).Elem())
}

type VsanPerfQueryNodeInformationResponse struct {
	Returnval []VsanPerfNodeInformation `xml:"returnval,omitempty"`
}

type VsanPerfQueryPerf VsanPerfQueryPerfRequestType

func init() {
	types.Add("vsan:VsanPerfQueryPerf", reflect.TypeOf((*VsanPerfQueryPerf)(nil)).Elem())
}

type VsanPerfQueryPerfRequestType struct {
	This       types.ManagedObjectReference  `xml:"_this"`
	QuerySpecs []VsanPerfQuerySpec           `xml:"querySpecs"`
	Cluster    *types.ManagedObjectReference `xml:"cluster,omitempty"`
}

func init() {
	types.Add("vsan:VsanPerfQueryPerfRequestType", reflect.TypeOf((*VsanPerfQueryPerfRequestType)(nil)).Elem())
}

type VsanPerfQueryPerfResponse struct {
	Returnval []VsanPerfEntityMetricCSV `xml:"returnval"`
}

type VsanPerfQuerySpec struct {
	types.DynamicData

	EntityRefId string     `xml:"entityRefId"`
	StartTime   *time.Time `xml:"startTime"`
	EndTime     *time.Time `xml:"endTime"`
	Group       string     `xml:"group,omitempty"`
	Labels      []string   `xml:"labels,omitempty"`
	Interval    int32      `xml:"interval,omitempty"`
}

func init() {
	types.Add("vsan:VsanPerfQuerySpec", reflect.TypeOf((*VsanPerfQuerySpec)(nil)).Elem())
}

type VsanPerfQueryStatsObjectInformation VsanPerfQueryStatsObjectInformationRequestType

func init() {
	types.Add("vsan:VsanPerfQueryStatsObjectInformation", reflect.TypeOf((*VsanPerfQueryStatsObjectInformation)(nil)).Elem())
}

type VsanPerfQueryStatsObjectInformationRequestType struct {
	This    types.ManagedObjectReference  `xml:"_this"`
	Cluster *types.ManagedObjectReference `xml:"cluster,omitempty"`
}

func init() {
	types.Add("vsan:VsanPerfQueryStatsObjectInformationRequestType", reflect.TypeOf((*VsanPerfQueryStatsObjectInformationRequestType)(nil)).Elem())
}

type VsanPerfQueryStatsObjectInformationResponse struct {
	Returnval VsanObjectInformation `xml:"returnval"`
}

type VsanPerfQueryTimeRanges VsanPerfQueryTimeRangesRequestType

func init() {
	types.Add("vsan:VsanPerfQueryTimeRanges", reflect.TypeOf((*VsanPerfQueryTimeRanges)(nil)).Elem())
}

type VsanPerfQueryTimeRangesRequestType struct {
	This      types.ManagedObjectReference  `xml:"_this"`
	Cluster   *types.ManagedObjectReference `xml:"cluster,omitempty"`
	QuerySpec VsanPerfTimeRangeQuerySpec    `xml:"querySpec"`
}

func init() {
	types.Add("vsan:VsanPerfQueryTimeRangesRequestType", reflect.TypeOf((*VsanPerfQueryTimeRangesRequestType)(nil)).Elem())
}

type VsanPerfQueryTimeRangesResponse struct {
	Returnval []VsanPerfTimeRange `xml:"returnval,omitempty"`
}

type VsanPerfSaveTimeRanges VsanPerfSaveTimeRangesRequestType

func init() {
	types.Add("vsan:VsanPerfSaveTimeRanges", reflect.TypeOf((*VsanPerfSaveTimeRanges)(nil)).Elem())
}

type VsanPerfSaveTimeRangesRequestType struct {
	This       types.ManagedObjectReference  `xml:"_this"`
	Cluster    *types.ManagedObjectReference `xml:"cluster,omitempty"`
	TimeRanges []VsanPerfTimeRange           `xml:"timeRanges"`
}

func init() {
	types.Add("vsan:VsanPerfSaveTimeRangesRequestType", reflect.TypeOf((*VsanPerfSaveTimeRangesRequestType)(nil)).Elem())
}

type VsanPerfSaveTimeRangesResponse struct {
}

type VsanPerfSetStatsObjectPolicy VsanPerfSetStatsObjectPolicyRequestType

func init() {
	types.Add("vsan:VsanPerfSetStatsObjectPolicy", reflect.TypeOf((*VsanPerfSetStatsObjectPolicy)(nil)).Elem())
}

type VsanPerfSetStatsObjectPolicyRequestType struct {
	This    types.ManagedObjectReference        `xml:"_this"`
	Cluster *types.ManagedObjectReference       `xml:"cluster,omitempty"`
	Profile types.BaseVirtualMachineProfileSpec `xml:"profile,omitempty,typeattr"`
}

func init() {
	types.Add("vsan:VsanPerfSetStatsObjectPolicyRequestType", reflect.TypeOf((*VsanPerfSetStatsObjectPolicyRequestType)(nil)).Elem())
}

type VsanPerfSetStatsObjectPolicyResponse struct {
	Returnval bool `xml:"returnval"`
}

type VsanPerfThreshold struct {
	types.DynamicData

	Direction string `xml:"direction"`
	Yellow    string `xml:"yellow,omitempty"`
	Red       string `xml:"red,omitempty"`
}

func init() {
	types.Add("vsan:VsanPerfThreshold", reflect.TypeOf((*VsanPerfThreshold)(nil)).Elem())
}

type VsanPerfTimeRange struct {
	types.DynamicData

	Name      string    `xml:"name"`
	StartTime time.Time `xml:"startTime"`
	EndTime   time.Time `xml:"endTime"`
}

func init() {
	types.Add("vsan:VsanPerfTimeRange", reflect.TypeOf((*VsanPerfTimeRange)(nil)).Elem())
}

type VsanPerfTimeRangeQuerySpec struct {
	types.DynamicData

	Name          string     `xml:"name,omitempty"`
	StartTimeFrom *time.Time `xml:"startTimeFrom"`
	StartTimeTo   *time.Time `xml:"startTimeTo"`
	EndTimeFrom   *time.Time `xml:"endTimeFrom"`
	EndTimeTo     *time.Time `xml:"endTimeTo"`
}

func init() {
	types.Add("vsan:VsanPerfTimeRangeQuerySpec", reflect.TypeOf((*VsanPerfTimeRangeQuerySpec)(nil)).Elem())
}

type VsanPerfToggleVerboseMode VsanPerfToggleVerboseModeRequestType

func init() {
	types.Add("vsan:VsanPerfToggleVerboseMode", reflect.TypeOf((*VsanPerfToggleVerboseMode)(nil)).Elem())
}

type VsanPerfToggleVerboseModeRequestType struct {
	This        types.ManagedObjectReference  `xml:"_this"`
	Cluster     *types.ManagedObjectReference `xml:"cluster,omitempty"`
	VerboseMode bool                          `xml:"verboseMode"`
}

func init() {
	types.Add("vsan:VsanPerfToggleVerboseModeRequestType", reflect.TypeOf((*VsanPerfToggleVerboseModeRequestType)(nil)).Elem())
}

type VsanPerfToggleVerboseModeResponse struct {
}

type VsanPerfTopEntities struct {
	types.DynamicData

	MetricId VsanPerfMetricId    `xml:"metricId"`
	Entities []VsanPerfTopEntity `xml:"entities"`
}

func init() {
	types.Add("vsan:VsanPerfTopEntities", reflect.TypeOf((*VsanPerfTopEntities)(nil)).Elem())
}

type VsanPerfTopEntity struct {
	types.DynamicData

	EntityRefId string `xml:"entityRefId"`
	Value       string `xml:"value"`
}

func init() {
	types.Add("vsan:VsanPerfTopEntity", reflect.TypeOf((*VsanPerfTopEntity)(nil)).Elem())
}

type VsanPerformFileServiceEnablePreflightCheck VsanPerformFileServiceEnablePreflightCheckRequestType

func init() {
	types.Add("vsan:VsanPerformFileServiceEnablePreflightCheck", reflect.TypeOf((*VsanPerformFileServiceEnablePreflightCheck)(nil)).Elem())
}

type VsanPerformFileServiceEnablePreflightCheckRequestType struct {
	This         types.ManagedObjectReference  `xml:"_this"`
	Cluster      types.ManagedObjectReference  `xml:"cluster"`
	DomainConfig *VsanFileServiceDomainConfig  `xml:"domainConfig,omitempty"`
	Network      *types.ManagedObjectReference `xml:"network,omitempty"`
}

func init() {
	types.Add("vsan:VsanPerformFileServiceEnablePreflightCheckRequestType", reflect.TypeOf((*VsanPerformFileServiceEnablePreflightCheckRequestType)(nil)).Elem())
}

type VsanPerformFileServiceEnablePreflightCheckResponse struct {
	Returnval VsanFileServicePreflightCheckResult `xml:"returnval"`
}

type VsanPerformOnlineHealthCheck VsanPerformOnlineHealthCheckRequestType

func init() {
	types.Add("vsan:VsanPerformOnlineHealthCheck", reflect.TypeOf((*VsanPerformOnlineHealthCheck)(nil)).Elem())
}

type VsanPerformOnlineHealthCheckRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

func init() {
	types.Add("vsan:VsanPerformOnlineHealthCheckRequestType", reflect.TypeOf((*VsanPerformOnlineHealthCheckRequestType)(nil)).Elem())
}

type VsanPerformOnlineHealthCheckResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanPerformResourceCheck VsanPerformResourceCheckRequestType

func init() {
	types.Add("vsan:VsanPerformResourceCheck", reflect.TypeOf((*VsanPerformResourceCheck)(nil)).Elem())
}

type VsanPerformResourceCheckRequestType struct {
	This              types.ManagedObjectReference  `xml:"_this"`
	ResourceCheckSpec VsanResourceCheckSpec         `xml:"resourceCheckSpec"`
	Cluster           *types.ManagedObjectReference `xml:"cluster,omitempty"`
}

func init() {
	types.Add("vsan:VsanPerformResourceCheckRequestType", reflect.TypeOf((*VsanPerformResourceCheckRequestType)(nil)).Elem())
}

type VsanPerformResourceCheckResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanPerfsvcConfig struct {
	types.DynamicData

	Enabled        bool                                `xml:"enabled"`
	Profile        types.BaseVirtualMachineProfileSpec `xml:"profile,omitempty,typeattr"`
	DiagnosticMode *bool                               `xml:"diagnosticMode"`
	VerboseMode    *bool                               `xml:"verboseMode"`
}

func init() {
	types.Add("vsan:VsanPerfsvcConfig", reflect.TypeOf((*VsanPerfsvcConfig)(nil)).Elem())
}

type VsanPerfsvcHealthResult struct {
	types.DynamicData

	StatsObjectInfo             *VsanObjectInformation    `xml:"statsObjectInfo,omitempty"`
	StatsObjectConsistent       *bool                     `xml:"statsObjectConsistent"`
	StatsObjectPolicyConsistent *bool                     `xml:"statsObjectPolicyConsistent"`
	DatastoreCompatible         *bool                     `xml:"datastoreCompatible"`
	EnoughFreeSpace             *bool                     `xml:"enoughFreeSpace"`
	RemediateAction             string                    `xml:"remediateAction,omitempty"`
	HostResults                 []VsanPerfNodeInformation `xml:"hostResults,omitempty"`
	VerboseModeStatus           *bool                     `xml:"verboseModeStatus"`
}

func init() {
	types.Add("vsan:VsanPerfsvcHealthResult", reflect.TypeOf((*VsanPerfsvcHealthResult)(nil)).Elem())
}

type VsanPhysicalDiskHealth struct {
	types.DynamicData

	Name                         string                   `xml:"name"`
	Uuid                         string                   `xml:"uuid"`
	InCmmds                      bool                     `xml:"inCmmds"`
	InVsi                        bool                     `xml:"inVsi"`
	DedupScope                   int64                    `xml:"dedupScope,omitempty"`
	FormatVersion                int32                    `xml:"formatVersion,omitempty"`
	IsAllFlash                   int32                    `xml:"isAllFlash,omitempty"`
	CongestionValue              int32                    `xml:"congestionValue,omitempty"`
	CongestionArea               string                   `xml:"congestionArea,omitempty"`
	CongestionHealth             string                   `xml:"congestionHealth,omitempty"`
	MetadataHealth               string                   `xml:"metadataHealth,omitempty"`
	OperationalHealthDescription string                   `xml:"operationalHealthDescription,omitempty"`
	OperationalHealth            string                   `xml:"operationalHealth,omitempty"`
	DedupUsageHealth             string                   `xml:"dedupUsageHealth,omitempty"`
	CapacityHealth               string                   `xml:"capacityHealth,omitempty"`
	SummaryHealth                string                   `xml:"summaryHealth"`
	Capacity                     int64                    `xml:"capacity,omitempty"`
	UsedCapacity                 int64                    `xml:"usedCapacity,omitempty"`
	ReservedCapacity             int64                    `xml:"reservedCapacity,omitempty"`
	TotalBytes                   int64                    `xml:"totalBytes,omitempty"`
	FreeBytes                    int64                    `xml:"freeBytes,omitempty"`
	HashedBytes                  int64                    `xml:"hashedBytes,omitempty"`
	DedupedBytes                 int64                    `xml:"dedupedBytes,omitempty"`
	ScsiDisk                     *types.HostScsiDisk      `xml:"scsiDisk,omitempty"`
	UsedComponents               int64                    `xml:"usedComponents,omitempty"`
	MaxComponents                int64                    `xml:"maxComponents,omitempty"`
	CompLimitHealth              string                   `xml:"compLimitHealth,omitempty"`
	EncryptionEnabled            *bool                    `xml:"encryptionEnabled"`
	KmsProviderId                string                   `xml:"kmsProviderId,omitempty"`
	KekId                        string                   `xml:"kekId,omitempty"`
	DekGenerationId              int64                    `xml:"dekGenerationId,omitempty"`
	EncryptedUnlocked            *bool                    `xml:"encryptedUnlocked"`
	RebalanceResult              *VsanDiskRebalanceResult `xml:"rebalanceResult,omitempty"`
}

func init() {
	types.Add("vsan:VsanPhysicalDiskHealth", reflect.TypeOf((*VsanPhysicalDiskHealth)(nil)).Elem())
}

type VsanPhysicalDiskHealthSummary struct {
	types.DynamicData

	OverallHealth        string                      `xml:"overallHealth"`
	HeapsWithIssues      []VsanResourceHealth        `xml:"heapsWithIssues,omitempty"`
	SlabsWithIssues      []VsanResourceHealth        `xml:"slabsWithIssues,omitempty"`
	Disks                []VsanPhysicalDiskHealth    `xml:"disks,omitempty"`
	ComponentsWithIssues []VsanResourceHealth        `xml:"componentsWithIssues,omitempty"`
	Hostname             string                      `xml:"hostname,omitempty"`
	HostDedupScope       int32                       `xml:"hostDedupScope,omitempty"`
	Error                *types.LocalizedMethodFault `xml:"error,omitempty"`
}

func init() {
	types.Add("vsan:VsanPhysicalDiskHealthSummary", reflect.TypeOf((*VsanPhysicalDiskHealthSummary)(nil)).Elem())
}

type VsanPostConfigForVcsa VsanPostConfigForVcsaRequestType

func init() {
	types.Add("vsan:VsanPostConfigForVcsa", reflect.TypeOf((*VsanPostConfigForVcsa)(nil)).Elem())
}

type VsanPostConfigForVcsaRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
	Spec VsanVcPostDeployConfigSpec   `xml:"spec"`
}

func init() {
	types.Add("vsan:VsanPostConfigForVcsaRequestType", reflect.TypeOf((*VsanPostConfigForVcsaRequestType)(nil)).Elem())
}

type VsanPostConfigForVcsaResponse struct {
	Returnval string `xml:"returnval,omitempty"`
}

type VsanPrepareVsanForVcsa VsanPrepareVsanForVcsaRequestType

func init() {
	types.Add("vsan:VsanPrepareVsanForVcsa", reflect.TypeOf((*VsanPrepareVsanForVcsa)(nil)).Elem())
}

type VsanPrepareVsanForVcsaRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
	Spec VsanPrepareVsanForVcsaSpec   `xml:"spec"`
}

func init() {
	types.Add("vsan:VsanPrepareVsanForVcsaRequestType", reflect.TypeOf((*VsanPrepareVsanForVcsaRequestType)(nil)).Elem())
}

type VsanPrepareVsanForVcsaResponse struct {
	Returnval string `xml:"returnval,omitempty"`
}

type VsanPrepareVsanForVcsaSpec struct {
	types.DynamicData

	VsanDiskMappingCreationSpec *VimVsanHostDiskMappingCreationSpec `xml:"vsanDiskMappingCreationSpec,omitempty"`
	VsanDataEfficiencyConfig    *VsanDataEfficiencyConfig           `xml:"vsanDataEfficiencyConfig,omitempty"`
	TaskId                      string                              `xml:"taskId,omitempty"`
	VsanDataEncryptionConfig    *VsanHostEncryptionInfo             `xml:"vsanDataEncryptionConfig,omitempty"`
}

func init() {
	types.Add("vsan:VsanPrepareVsanForVcsaSpec", reflect.TypeOf((*VsanPrepareVsanForVcsaSpec)(nil)).Elem())
}

type VsanProactiveRebalanceInfo struct {
	types.DynamicData

	Enabled   *bool `xml:"enabled"`
	Threshold int32 `xml:"threshold,omitempty"`
}

func init() {
	types.Add("vsan:VsanProactiveRebalanceInfo", reflect.TypeOf((*VsanProactiveRebalanceInfo)(nil)).Elem())
}

type VsanProactiveRebalanceInfoEx struct {
	types.DynamicData

	Running           *bool                       `xml:"running"`
	StartTs           *time.Time                  `xml:"startTs"`
	StopTs            *time.Time                  `xml:"stopTs"`
	VarianceThreshold float32                     `xml:"varianceThreshold,omitempty"`
	TimeThreshold     int32                       `xml:"timeThreshold,omitempty"`
	RateThreshold     int32                       `xml:"rateThreshold,omitempty"`
	Hostname          string                      `xml:"hostname,omitempty"`
	Error             *types.LocalizedMethodFault `xml:"error,omitempty"`
}

func init() {
	types.Add("vsan:VsanProactiveRebalanceInfoEx", reflect.TypeOf((*VsanProactiveRebalanceInfoEx)(nil)).Elem())
}

type VsanPropertyConstraint struct {
	VsanResourceConstraint

	PropertyName    string             `xml:"propertyName,omitempty"`
	Comparator      string             `xml:"comparator,omitempty"`
	ComparableValue *types.KeyAnyValue `xml:"comparableValue,omitempty"`
}

func init() {
	types.Add("vsan:VsanPropertyConstraint", reflect.TypeOf((*VsanPropertyConstraint)(nil)).Elem())
}

type VsanPurgeHclFiles VsanPurgeHclFilesRequestType

func init() {
	types.Add("vsan:VsanPurgeHclFiles", reflect.TypeOf((*VsanPurgeHclFiles)(nil)).Elem())
}

type VsanPurgeHclFilesRequestType struct {
	This     types.ManagedObjectReference `xml:"_this"`
	Sha1sums []string                     `xml:"sha1sums"`
}

func init() {
	types.Add("vsan:VsanPurgeHclFilesRequestType", reflect.TypeOf((*VsanPurgeHclFilesRequestType)(nil)).Elem())
}

type VsanPurgeHclFilesResponse struct {
}

type VsanQueryAllSupportedHealthChecks VsanQueryAllSupportedHealthChecksRequestType

func init() {
	types.Add("vsan:VsanQueryAllSupportedHealthChecks", reflect.TypeOf((*VsanQueryAllSupportedHealthChecks)(nil)).Elem())
}

type VsanQueryAllSupportedHealthChecksRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

func init() {
	types.Add("vsan:VsanQueryAllSupportedHealthChecksRequestType", reflect.TypeOf((*VsanQueryAllSupportedHealthChecksRequestType)(nil)).Elem())
}

type VsanQueryAllSupportedHealthChecksResponse struct {
	Returnval []VsanClusterHealthCheckInfo `xml:"returnval"`
}

type VsanQueryAttachToSrHistory VsanQueryAttachToSrHistoryRequestType

func init() {
	types.Add("vsan:VsanQueryAttachToSrHistory", reflect.TypeOf((*VsanQueryAttachToSrHistory)(nil)).Elem())
}

type VsanQueryAttachToSrHistoryRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
	Count   int32                        `xml:"count,omitempty"`
	TaskId  string                       `xml:"taskId,omitempty"`
}

func init() {
	types.Add("vsan:VsanQueryAttachToSrHistoryRequestType", reflect.TypeOf((*VsanQueryAttachToSrHistoryRequestType)(nil)).Elem())
}

type VsanQueryAttachToSrHistoryResponse struct {
	Returnval []VsanAttachToSrOperation `xml:"returnval,omitempty"`
}

type VsanQueryClusterAdvCfgSync VsanQueryClusterAdvCfgSyncRequestType

func init() {
	types.Add("vsan:VsanQueryClusterAdvCfgSync", reflect.TypeOf((*VsanQueryClusterAdvCfgSync)(nil)).Elem())
}

type VsanQueryClusterAdvCfgSyncRequestType struct {
	This            types.ManagedObjectReference `xml:"_this"`
	Hosts           []string                     `xml:"hosts"`
	EsxRootPassword string                       `xml:"esxRootPassword"`
	Options         []string                     `xml:"options,omitempty"`
}

func init() {
	types.Add("vsan:VsanQueryClusterAdvCfgSyncRequestType", reflect.TypeOf((*VsanQueryClusterAdvCfgSyncRequestType)(nil)).Elem())
}

type VsanQueryClusterAdvCfgSyncResponse struct {
	Returnval []VsanClusterAdvCfgSyncResult `xml:"returnval,omitempty"`
}

type VsanQueryClusterCaptureVsanPcap VsanQueryClusterCaptureVsanPcapRequestType

func init() {
	types.Add("vsan:VsanQueryClusterCaptureVsanPcap", reflect.TypeOf((*VsanQueryClusterCaptureVsanPcap)(nil)).Elem())
}

type VsanQueryClusterCaptureVsanPcapRequestType struct {
	This               types.ManagedObjectReference   `xml:"_this"`
	Hosts              []string                       `xml:"hosts"`
	EsxRootPassword    string                         `xml:"esxRootPassword"`
	Duration           int32                          `xml:"duration"`
	Vmknic             []VsanClusterHostVmknicMapping `xml:"vmknic,omitempty"`
	IncludeRawPcap     *bool                          `xml:"includeRawPcap"`
	IncludeIgmp        *bool                          `xml:"includeIgmp"`
	CmmdsMsgTypeFilter []string                       `xml:"cmmdsMsgTypeFilter,omitempty"`
	CmmdsPorts         []int32                        `xml:"cmmdsPorts,omitempty"`
	ClusterUuid        string                         `xml:"clusterUuid,omitempty"`
}

func init() {
	types.Add("vsan:VsanQueryClusterCaptureVsanPcapRequestType", reflect.TypeOf((*VsanQueryClusterCaptureVsanPcapRequestType)(nil)).Elem())
}

type VsanQueryClusterCaptureVsanPcapResponse struct {
	Returnval VsanVsanClusterPcapResult `xml:"returnval"`
}

type VsanQueryClusterCheckLimits VsanQueryClusterCheckLimitsRequestType

func init() {
	types.Add("vsan:VsanQueryClusterCheckLimits", reflect.TypeOf((*VsanQueryClusterCheckLimits)(nil)).Elem())
}

type VsanQueryClusterCheckLimitsRequestType struct {
	This            types.ManagedObjectReference `xml:"_this"`
	Hosts           []string                     `xml:"hosts"`
	EsxRootPassword string                       `xml:"esxRootPassword"`
}

func init() {
	types.Add("vsan:VsanQueryClusterCheckLimitsRequestType", reflect.TypeOf((*VsanQueryClusterCheckLimitsRequestType)(nil)).Elem())
}

type VsanQueryClusterCheckLimitsResponse struct {
	Returnval VsanClusterLimitHealthResult `xml:"returnval"`
}

type VsanQueryClusterCreateVmHealthTest VsanQueryClusterCreateVmHealthTestRequestType

func init() {
	types.Add("vsan:VsanQueryClusterCreateVmHealthTest", reflect.TypeOf((*VsanQueryClusterCreateVmHealthTest)(nil)).Elem())
}

type VsanQueryClusterCreateVmHealthTestRequestType struct {
	This            types.ManagedObjectReference `xml:"_this"`
	Hosts           []string                     `xml:"hosts"`
	EsxRootPassword string                       `xml:"esxRootPassword"`
	Timeout         int32                        `xml:"timeout"`
}

func init() {
	types.Add("vsan:VsanQueryClusterCreateVmHealthTestRequestType", reflect.TypeOf((*VsanQueryClusterCreateVmHealthTestRequestType)(nil)).Elem())
}

type VsanQueryClusterCreateVmHealthTestResponse struct {
	Returnval VsanClusterCreateVmHealthTestResult `xml:"returnval"`
}

type VsanQueryClusterDrsStats VsanQueryClusterDrsStatsRequestType

func init() {
	types.Add("vsan:VsanQueryClusterDrsStats", reflect.TypeOf((*VsanQueryClusterDrsStats)(nil)).Elem())
}

type VsanQueryClusterDrsStatsRequestType struct {
	This    types.ManagedObjectReference   `xml:"_this"`
	Cluster types.ManagedObjectReference   `xml:"cluster"`
	Vms     []types.ManagedObjectReference `xml:"vms,omitempty"`
}

func init() {
	types.Add("vsan:VsanQueryClusterDrsStatsRequestType", reflect.TypeOf((*VsanQueryClusterDrsStatsRequestType)(nil)).Elem())
}

type VsanQueryClusterDrsStatsResponse struct {
	Returnval []VsanHostDrsStats `xml:"returnval,omitempty"`
}

type VsanQueryClusterHealthSystemVersions VsanQueryClusterHealthSystemVersionsRequestType

func init() {
	types.Add("vsan:VsanQueryClusterHealthSystemVersions", reflect.TypeOf((*VsanQueryClusterHealthSystemVersions)(nil)).Elem())
}

type VsanQueryClusterHealthSystemVersionsRequestType struct {
	This            types.ManagedObjectReference `xml:"_this"`
	Hosts           []string                     `xml:"hosts"`
	EsxRootPassword string                       `xml:"esxRootPassword"`
}

func init() {
	types.Add("vsan:VsanQueryClusterHealthSystemVersionsRequestType", reflect.TypeOf((*VsanQueryClusterHealthSystemVersionsRequestType)(nil)).Elem())
}

type VsanQueryClusterHealthSystemVersionsResponse struct {
	Returnval VsanClusterHealthSystemVersionResult `xml:"returnval"`
}

type VsanQueryClusterNetworkPerfTest VsanQueryClusterNetworkPerfTestRequestType

func init() {
	types.Add("vsan:VsanQueryClusterNetworkPerfTest", reflect.TypeOf((*VsanQueryClusterNetworkPerfTest)(nil)).Elem())
}

type VsanQueryClusterNetworkPerfTestRequestType struct {
	This            types.ManagedObjectReference `xml:"_this"`
	Hosts           []string                     `xml:"hosts"`
	EsxRootPassword string                       `xml:"esxRootPassword"`
	Multicast       bool                         `xml:"multicast"`
	DurationSec     int32                        `xml:"durationSec,omitempty"`
}

func init() {
	types.Add("vsan:VsanQueryClusterNetworkPerfTestRequestType", reflect.TypeOf((*VsanQueryClusterNetworkPerfTestRequestType)(nil)).Elem())
}

type VsanQueryClusterNetworkPerfTestResponse struct {
	Returnval VsanClusterNetworkLoadTestResult `xml:"returnval"`
}

type VsanQueryClusterPhysicalDiskHealthSummary VsanQueryClusterPhysicalDiskHealthSummaryRequestType

func init() {
	types.Add("vsan:VsanQueryClusterPhysicalDiskHealthSummary", reflect.TypeOf((*VsanQueryClusterPhysicalDiskHealthSummary)(nil)).Elem())
}

type VsanQueryClusterPhysicalDiskHealthSummaryRequestType struct {
	This            types.ManagedObjectReference `xml:"_this"`
	Hosts           []string                     `xml:"hosts"`
	EsxRootPassword string                       `xml:"esxRootPassword"`
}

func init() {
	types.Add("vsan:VsanQueryClusterPhysicalDiskHealthSummaryRequestType", reflect.TypeOf((*VsanQueryClusterPhysicalDiskHealthSummaryRequestType)(nil)).Elem())
}

type VsanQueryClusterPhysicalDiskHealthSummaryResponse struct {
	Returnval []VsanPhysicalDiskHealthSummary `xml:"returnval"`
}

type VsanQueryEntitySpaceUsage VsanQueryEntitySpaceUsageRequestType

func init() {
	types.Add("vsan:VsanQueryEntitySpaceUsage", reflect.TypeOf((*VsanQueryEntitySpaceUsage)(nil)).Elem())
}

type VsanQueryEntitySpaceUsageRequestType struct {
	This      types.ManagedObjectReference `xml:"_this"`
	Cluster   types.ManagedObjectReference `xml:"cluster"`
	QuerySpec VsanSpaceQuerySpec           `xml:"querySpec"`
}

func init() {
	types.Add("vsan:VsanQueryEntitySpaceUsageRequestType", reflect.TypeOf((*VsanQueryEntitySpaceUsageRequestType)(nil)).Elem())
}

type VsanQueryEntitySpaceUsageResponse struct {
	Returnval []VsanEntitySpaceUsage `xml:"returnval,omitempty"`
}

type VsanQueryFileServiceOvfs VsanQueryFileServiceOvfsRequestType

func init() {
	types.Add("vsan:VsanQueryFileServiceOvfs", reflect.TypeOf((*VsanQueryFileServiceOvfs)(nil)).Elem())
}

type VsanQueryFileServiceOvfsRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

func init() {
	types.Add("vsan:VsanQueryFileServiceOvfsRequestType", reflect.TypeOf((*VsanQueryFileServiceOvfsRequestType)(nil)).Elem())
}

type VsanQueryFileServiceOvfsResponse struct {
	Returnval []VsanFileServiceOvfSpec `xml:"returnval,omitempty"`
}

type VsanQueryHostDrsStats VsanQueryHostDrsStatsRequestType

func init() {
	types.Add("vsan:VsanQueryHostDrsStats", reflect.TypeOf((*VsanQueryHostDrsStats)(nil)).Elem())
}

type VsanQueryHostDrsStatsRequestType struct {
	This      types.ManagedObjectReference `xml:"_this"`
	HostUuids []string                     `xml:"hostUuids"`
	Vms       []string                     `xml:"vms,omitempty"`
}

func init() {
	types.Add("vsan:VsanQueryHostDrsStatsRequestType", reflect.TypeOf((*VsanQueryHostDrsStatsRequestType)(nil)).Elem())
}

type VsanQueryHostDrsStatsResponse struct {
	Returnval VsanHostDrsStats `xml:"returnval"`
}

type VsanQueryHostEMMState VsanQueryHostEMMStateRequestType

func init() {
	types.Add("vsan:VsanQueryHostEMMState", reflect.TypeOf((*VsanQueryHostEMMState)(nil)).Elem())
}

type VsanQueryHostEMMStateRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

func init() {
	types.Add("vsan:VsanQueryHostEMMStateRequestType", reflect.TypeOf((*VsanQueryHostEMMStateRequestType)(nil)).Elem())
}

type VsanQueryHostEMMStateResponse struct {
	Returnval VsanHostEMMSummary `xml:"returnval"`
}

type VsanQueryInaccessibleVmSwapObjects VsanQueryInaccessibleVmSwapObjectsRequestType

func init() {
	types.Add("vsan:VsanQueryInaccessibleVmSwapObjects", reflect.TypeOf((*VsanQueryInaccessibleVmSwapObjects)(nil)).Elem())
}

type VsanQueryInaccessibleVmSwapObjectsRequestType struct {
	This    types.ManagedObjectReference  `xml:"_this"`
	Cluster *types.ManagedObjectReference `xml:"cluster,omitempty"`
}

func init() {
	types.Add("vsan:VsanQueryInaccessibleVmSwapObjectsRequestType", reflect.TypeOf((*VsanQueryInaccessibleVmSwapObjectsRequestType)(nil)).Elem())
}

type VsanQueryInaccessibleVmSwapObjectsResponse struct {
	Returnval []string `xml:"returnval,omitempty"`
}

type VsanQueryObjectIdentities VsanQueryObjectIdentitiesRequestType

func init() {
	types.Add("vsan:VsanQueryObjectIdentities", reflect.TypeOf((*VsanQueryObjectIdentities)(nil)).Elem())
}

type VsanQueryObjectIdentitiesRequestType struct {
	This                types.ManagedObjectReference  `xml:"_this"`
	Cluster             *types.ManagedObjectReference `xml:"cluster,omitempty"`
	ObjUuids            []string                      `xml:"objUuids,omitempty"`
	ObjTypes            []string                      `xml:"objTypes,omitempty"`
	IncludeHealth       *bool                         `xml:"includeHealth"`
	IncludeObjIdentity  *bool                         `xml:"includeObjIdentity"`
	IncludeSpaceSummary *bool                         `xml:"includeSpaceSummary"`
}

func init() {
	types.Add("vsan:VsanQueryObjectIdentitiesRequestType", reflect.TypeOf((*VsanQueryObjectIdentitiesRequestType)(nil)).Elem())
}

type VsanQueryObjectIdentitiesResponse struct {
	Returnval *VsanObjectIdentityAndHealth `xml:"returnval,omitempty"`
}

type VsanQueryResultHostInfo struct {
	types.DynamicData

	Uuid              string   `xml:"uuid,omitempty"`
	HostnameInCmmds   string   `xml:"hostnameInCmmds,omitempty"`
	VsanIpv4Addresses []string `xml:"vsanIpv4Addresses,omitempty"`
}

func init() {
	types.Add("vsan:VsanQueryResultHostInfo", reflect.TypeOf((*VsanQueryResultHostInfo)(nil)).Elem())
}

type VsanQuerySpaceUsage VsanQuerySpaceUsageRequestType

func init() {
	types.Add("vsan:VsanQuerySpaceUsage", reflect.TypeOf((*VsanQuerySpaceUsage)(nil)).Elem())
}

type VsanQuerySpaceUsageRequestType struct {
	This               types.ManagedObjectReference          `xml:"_this"`
	Cluster            types.ManagedObjectReference          `xml:"cluster"`
	StoragePolicies    []types.BaseVirtualMachineProfileSpec `xml:"storagePolicies,omitempty,typeattr"`
	WhatifCapacityOnly *bool                                 `xml:"whatifCapacityOnly"`
}

func init() {
	types.Add("vsan:VsanQuerySpaceUsageRequestType", reflect.TypeOf((*VsanQuerySpaceUsageRequestType)(nil)).Elem())
}

type VsanQuerySpaceUsageResponse struct {
	Returnval VsanSpaceUsage `xml:"returnval"`
}

type VsanQuerySyncingVsanObjects VsanQuerySyncingVsanObjectsRequestType

func init() {
	types.Add("vsan:VsanQuerySyncingVsanObjects", reflect.TypeOf((*VsanQuerySyncingVsanObjects)(nil)).Elem())
}

type VsanQuerySyncingVsanObjectsRequestType struct {
	This           types.ManagedObjectReference `xml:"_this"`
	Uuids          []string                     `xml:"uuids,omitempty"`
	Start          int32                        `xml:"start,omitempty"`
	Limit          *int32                       `xml:"limit"`
	IncludeSummary *bool                        `xml:"includeSummary"`
}

func init() {
	types.Add("vsan:VsanQuerySyncingVsanObjectsRequestType", reflect.TypeOf((*VsanQuerySyncingVsanObjectsRequestType)(nil)).Elem())
}

type VsanQuerySyncingVsanObjectsResponse struct {
	Returnval VsanHostVsanObjectSyncQueryResult `xml:"returnval"`
}

type VsanQueryUpgradeStatusEx VsanQueryUpgradeStatusExRequestType

func init() {
	types.Add("vsan:VsanQueryUpgradeStatusEx", reflect.TypeOf((*VsanQueryUpgradeStatusEx)(nil)).Elem())
}

type VsanQueryUpgradeStatusExRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

func init() {
	types.Add("vsan:VsanQueryUpgradeStatusExRequestType", reflect.TypeOf((*VsanQueryUpgradeStatusExRequestType)(nil)).Elem())
}

type VsanQueryUpgradeStatusExResponse struct {
	Returnval VsanUpgradeStatusEx `xml:"returnval"`
}

type VsanQueryVcClusterCreateVmHealthHistoryTest VsanQueryVcClusterCreateVmHealthHistoryTestRequestType

func init() {
	types.Add("vsan:VsanQueryVcClusterCreateVmHealthHistoryTest", reflect.TypeOf((*VsanQueryVcClusterCreateVmHealthHistoryTest)(nil)).Elem())
}

type VsanQueryVcClusterCreateVmHealthHistoryTestRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
	Count   int32                        `xml:"count,omitempty"`
}

func init() {
	types.Add("vsan:VsanQueryVcClusterCreateVmHealthHistoryTestRequestType", reflect.TypeOf((*VsanQueryVcClusterCreateVmHealthHistoryTestRequestType)(nil)).Elem())
}

type VsanQueryVcClusterCreateVmHealthHistoryTestResponse struct {
	Returnval []VsanClusterCreateVmHealthTestResult `xml:"returnval,omitempty"`
}

type VsanQueryVcClusterCreateVmHealthTest VsanQueryVcClusterCreateVmHealthTestRequestType

func init() {
	types.Add("vsan:VsanQueryVcClusterCreateVmHealthTest", reflect.TypeOf((*VsanQueryVcClusterCreateVmHealthTest)(nil)).Elem())
}

type VsanQueryVcClusterCreateVmHealthTestRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
	Timeout int32                        `xml:"timeout"`
}

func init() {
	types.Add("vsan:VsanQueryVcClusterCreateVmHealthTestRequestType", reflect.TypeOf((*VsanQueryVcClusterCreateVmHealthTestRequestType)(nil)).Elem())
}

type VsanQueryVcClusterCreateVmHealthTestResponse struct {
	Returnval VsanClusterCreateVmHealthTestResult `xml:"returnval"`
}

type VsanQueryVcClusterHealthSummary VsanQueryVcClusterHealthSummaryRequestType

func init() {
	types.Add("vsan:VsanQueryVcClusterHealthSummary", reflect.TypeOf((*VsanQueryVcClusterHealthSummary)(nil)).Elem())
}

type VsanQueryVcClusterHealthSummaryRequestType struct {
	This            types.ManagedObjectReference   `xml:"_this"`
	Cluster         *types.ManagedObjectReference  `xml:"cluster,omitempty"`
	VmCreateTimeout int32                          `xml:"vmCreateTimeout,omitempty"`
	ObjUuids        []string                       `xml:"objUuids,omitempty"`
	IncludeObjUuids *bool                          `xml:"includeObjUuids"`
	Fields          []string                       `xml:"fields,omitempty"`
	FetchFromCache  *bool                          `xml:"fetchFromCache"`
	Perspective     string                         `xml:"perspective,omitempty"`
	Hosts           []types.ManagedObjectReference `xml:"hosts,omitempty"`
}

func init() {
	types.Add("vsan:VsanQueryVcClusterHealthSummaryRequestType", reflect.TypeOf((*VsanQueryVcClusterHealthSummaryRequestType)(nil)).Elem())
}

type VsanQueryVcClusterHealthSummaryResponse struct {
	Returnval VsanClusterHealthSummary `xml:"returnval"`
}

type VsanQueryVcClusterHealthSummaryTask VsanQueryVcClusterHealthSummaryTaskRequestType

func init() {
	types.Add("vsan:VsanQueryVcClusterHealthSummaryTask", reflect.TypeOf((*VsanQueryVcClusterHealthSummaryTask)(nil)).Elem())
}

type VsanQueryVcClusterHealthSummaryTaskRequestType struct {
	This                        types.ManagedObjectReference   `xml:"_this"`
	Cluster                     types.ManagedObjectReference   `xml:"cluster"`
	Hosts                       []types.ManagedObjectReference `xml:"hosts,omitempty"`
	IncludeDataProtectionHealth *bool                          `xml:"includeDataProtectionHealth"`
}

func init() {
	types.Add("vsan:VsanQueryVcClusterHealthSummaryTaskRequestType", reflect.TypeOf((*VsanQueryVcClusterHealthSummaryTaskRequestType)(nil)).Elem())
}

type VsanQueryVcClusterHealthSummaryTaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanQueryVcClusterNetworkPerfHistoryTest VsanQueryVcClusterNetworkPerfHistoryTestRequestType

func init() {
	types.Add("vsan:VsanQueryVcClusterNetworkPerfHistoryTest", reflect.TypeOf((*VsanQueryVcClusterNetworkPerfHistoryTest)(nil)).Elem())
}

type VsanQueryVcClusterNetworkPerfHistoryTestRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
	Count   int32                        `xml:"count,omitempty"`
}

func init() {
	types.Add("vsan:VsanQueryVcClusterNetworkPerfHistoryTestRequestType", reflect.TypeOf((*VsanQueryVcClusterNetworkPerfHistoryTestRequestType)(nil)).Elem())
}

type VsanQueryVcClusterNetworkPerfHistoryTestResponse struct {
	Returnval []VsanClusterNetworkLoadTestResult `xml:"returnval,omitempty"`
}

type VsanQueryVcClusterNetworkPerfTest VsanQueryVcClusterNetworkPerfTestRequestType

func init() {
	types.Add("vsan:VsanQueryVcClusterNetworkPerfTest", reflect.TypeOf((*VsanQueryVcClusterNetworkPerfTest)(nil)).Elem())
}

type VsanQueryVcClusterNetworkPerfTestRequestType struct {
	This        types.ManagedObjectReference `xml:"_this"`
	Cluster     types.ManagedObjectReference `xml:"cluster"`
	Multicast   bool                         `xml:"multicast"`
	DurationSec int32                        `xml:"durationSec,omitempty"`
}

func init() {
	types.Add("vsan:VsanQueryVcClusterNetworkPerfTestRequestType", reflect.TypeOf((*VsanQueryVcClusterNetworkPerfTestRequestType)(nil)).Elem())
}

type VsanQueryVcClusterNetworkPerfTestResponse struct {
	Returnval VsanClusterNetworkLoadTestResult `xml:"returnval"`
}

type VsanQueryVcClusterObjExtAttrs VsanQueryVcClusterObjExtAttrsRequestType

func init() {
	types.Add("vsan:VsanQueryVcClusterObjExtAttrs", reflect.TypeOf((*VsanQueryVcClusterObjExtAttrs)(nil)).Elem())
}

type VsanQueryVcClusterObjExtAttrsRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
	Uuids   []string                     `xml:"uuids"`
}

func init() {
	types.Add("vsan:VsanQueryVcClusterObjExtAttrsRequestType", reflect.TypeOf((*VsanQueryVcClusterObjExtAttrsRequestType)(nil)).Elem())
}

type VsanQueryVcClusterObjExtAttrsResponse struct {
	Returnval []VsanClusterObjectExtAttrs `xml:"returnval,omitempty"`
}

type VsanQueryVcClusterSmartStatsSummary VsanQueryVcClusterSmartStatsSummaryRequestType

func init() {
	types.Add("vsan:VsanQueryVcClusterSmartStatsSummary", reflect.TypeOf((*VsanQueryVcClusterSmartStatsSummary)(nil)).Elem())
}

type VsanQueryVcClusterSmartStatsSummaryRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

func init() {
	types.Add("vsan:VsanQueryVcClusterSmartStatsSummaryRequestType", reflect.TypeOf((*VsanQueryVcClusterSmartStatsSummaryRequestType)(nil)).Elem())
}

type VsanQueryVcClusterSmartStatsSummaryResponse struct {
	Returnval []VsanSmartStatsHostSummary `xml:"returnval"`
}

type VsanQueryVcClusterVmdkLoadHistoryTest VsanQueryVcClusterVmdkLoadHistoryTestRequestType

func init() {
	types.Add("vsan:VsanQueryVcClusterVmdkLoadHistoryTest", reflect.TypeOf((*VsanQueryVcClusterVmdkLoadHistoryTest)(nil)).Elem())
}

type VsanQueryVcClusterVmdkLoadHistoryTestRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
	Count   int32                        `xml:"count,omitempty"`
	TaskId  string                       `xml:"taskId,omitempty"`
}

func init() {
	types.Add("vsan:VsanQueryVcClusterVmdkLoadHistoryTestRequestType", reflect.TypeOf((*VsanQueryVcClusterVmdkLoadHistoryTestRequestType)(nil)).Elem())
}

type VsanQueryVcClusterVmdkLoadHistoryTestResponse struct {
	Returnval []VsanClusterVmdkLoadTestResult `xml:"returnval,omitempty"`
}

type VsanQueryVcClusterVmdkWorkloadTypes VsanQueryVcClusterVmdkWorkloadTypesRequestType

func init() {
	types.Add("vsan:VsanQueryVcClusterVmdkWorkloadTypes", reflect.TypeOf((*VsanQueryVcClusterVmdkWorkloadTypes)(nil)).Elem())
}

type VsanQueryVcClusterVmdkWorkloadTypesRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

func init() {
	types.Add("vsan:VsanQueryVcClusterVmdkWorkloadTypesRequestType", reflect.TypeOf((*VsanQueryVcClusterVmdkWorkloadTypesRequestType)(nil)).Elem())
}

type VsanQueryVcClusterVmdkWorkloadTypesResponse struct {
	Returnval []VsanStorageWorkloadType `xml:"returnval"`
}

type VsanQueryVerifyClusterNetworkSettings VsanQueryVerifyClusterNetworkSettingsRequestType

func init() {
	types.Add("vsan:VsanQueryVerifyClusterNetworkSettings", reflect.TypeOf((*VsanQueryVerifyClusterNetworkSettings)(nil)).Elem())
}

type VsanQueryVerifyClusterNetworkSettingsRequestType struct {
	This            types.ManagedObjectReference `xml:"_this"`
	Hosts           []string                     `xml:"hosts"`
	EsxRootPassword string                       `xml:"esxRootPassword"`
}

func init() {
	types.Add("vsan:VsanQueryVerifyClusterNetworkSettingsRequestType", reflect.TypeOf((*VsanQueryVerifyClusterNetworkSettingsRequestType)(nil)).Elem())
}

type VsanQueryVerifyClusterNetworkSettingsResponse struct {
	Returnval VsanClusterNetworkHealthResult `xml:"returnval"`
}

type VsanQueryWhatIfEvacuationResult VsanQueryWhatIfEvacuationResultRequestType

func init() {
	types.Add("vsan:VsanQueryWhatIfEvacuationResult", reflect.TypeOf((*VsanQueryWhatIfEvacuationResult)(nil)).Elem())
}

type VsanQueryWhatIfEvacuationResultRequestType struct {
	This           types.ManagedObjectReference `xml:"_this"`
	EvacEntityUuid string                       `xml:"evacEntityUuid"`
}

func init() {
	types.Add("vsan:VsanQueryWhatIfEvacuationResultRequestType", reflect.TypeOf((*VsanQueryWhatIfEvacuationResultRequestType)(nil)).Elem())
}

type VsanQueryWhatIfEvacuationResultResponse struct {
	Returnval VsanWhatIfEvacResult `xml:"returnval"`
}

type VsanRebalanceCluster VsanRebalanceClusterRequestType

func init() {
	types.Add("vsan:VsanRebalanceCluster", reflect.TypeOf((*VsanRebalanceCluster)(nil)).Elem())
}

type VsanRebalanceClusterRequestType struct {
	This        types.ManagedObjectReference   `xml:"_this"`
	Cluster     types.ManagedObjectReference   `xml:"cluster"`
	TargetHosts []types.ManagedObjectReference `xml:"targetHosts,omitempty"`
}

func init() {
	types.Add("vsan:VsanRebalanceClusterRequestType", reflect.TypeOf((*VsanRebalanceClusterRequestType)(nil)).Elem())
}

type VsanRebalanceClusterResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanReconfigureFileShare VsanReconfigureFileShareRequestType

func init() {
	types.Add("vsan:VsanReconfigureFileShare", reflect.TypeOf((*VsanReconfigureFileShare)(nil)).Elem())
}

type VsanReconfigureFileShareRequestType struct {
	This            types.ManagedObjectReference  `xml:"_this"`
	ShareUuid       string                        `xml:"shareUuid"`
	Config          VsanFileShareConfig           `xml:"config"`
	Cluster         *types.ManagedObjectReference `xml:"cluster,omitempty"`
	DeleteLabelKeys []string                      `xml:"deleteLabelKeys,omitempty"`
	Force           *bool                         `xml:"force"`
}

func init() {
	types.Add("vsan:VsanReconfigureFileShareRequestType", reflect.TypeOf((*VsanReconfigureFileShareRequestType)(nil)).Elem())
}

type VsanReconfigureFileShareResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanRegexBasedRule struct {
	types.DynamicData

	Rules []string `xml:"rules,omitempty"`
}

func init() {
	types.Add("vsan:VsanRegexBasedRule", reflect.TypeOf((*VsanRegexBasedRule)(nil)).Elem())
}

type VsanRemediateVsanCluster VsanRemediateVsanClusterRequestType

func init() {
	types.Add("vsan:VsanRemediateVsanCluster", reflect.TypeOf((*VsanRemediateVsanCluster)(nil)).Elem())
}

type VsanRemediateVsanClusterRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

func init() {
	types.Add("vsan:VsanRemediateVsanClusterRequestType", reflect.TypeOf((*VsanRemediateVsanClusterRequestType)(nil)).Elem())
}

type VsanRemediateVsanClusterResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanRemediateVsanHost VsanRemediateVsanHostRequestType

func init() {
	types.Add("vsan:VsanRemediateVsanHost", reflect.TypeOf((*VsanRemediateVsanHost)(nil)).Elem())
}

type VsanRemediateVsanHostRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
	Host types.ManagedObjectReference `xml:"host"`
}

func init() {
	types.Add("vsan:VsanRemediateVsanHostRequestType", reflect.TypeOf((*VsanRemediateVsanHostRequestType)(nil)).Elem())
}

type VsanRemediateVsanHostResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanRepairClusterImmediateObjects VsanRepairClusterImmediateObjectsRequestType

func init() {
	types.Add("vsan:VsanRepairClusterImmediateObjects", reflect.TypeOf((*VsanRepairClusterImmediateObjects)(nil)).Elem())
}

type VsanRepairClusterImmediateObjectsRequestType struct {
	This            types.ManagedObjectReference `xml:"_this"`
	Hosts           []string                     `xml:"hosts"`
	EsxRootPassword string                       `xml:"esxRootPassword"`
	Uuids           []string                     `xml:"uuids,omitempty"`
}

func init() {
	types.Add("vsan:VsanRepairClusterImmediateObjectsRequestType", reflect.TypeOf((*VsanRepairClusterImmediateObjectsRequestType)(nil)).Elem())
}

type VsanRepairClusterImmediateObjectsResponse struct {
	Returnval VsanClusterHealthSystemObjectsRepairResult `xml:"returnval"`
}

type VsanRepairObjectsResult struct {
	types.DynamicData

	InQueueObjects      []string                       `xml:"inQueueObjects,omitempty"`
	FailedRepairObjects []VsanFailedRepairObjectResult `xml:"failedRepairObjects,omitempty"`
	NotInQueueObjects   []string                       `xml:"notInQueueObjects,omitempty"`
}

func init() {
	types.Add("vsan:VsanRepairObjectsResult", reflect.TypeOf((*VsanRepairObjectsResult)(nil)).Elem())
}

type VsanResourceCheckResult struct {
	EntityResourceCheckDetails

	Timestamp           time.Time                            `xml:"timestamp"`
	Status              string                               `xml:"status"`
	Messages            []types.LocalizableMessage           `xml:"messages,omitempty"`
	FaultDomains        []VsanFaultDomainResourceCheckResult `xml:"faultDomains,omitempty"`
	DataToMove          int64                                `xml:"dataToMove,omitempty"`
	NonCompliantObjects []string                             `xml:"nonCompliantObjects,omitempty"`
	InaccessibleObjects []string                             `xml:"inaccessibleObjects,omitempty"`
	CapacityThreshold   *VsanHealthThreshold                 `xml:"capacityThreshold,omitempty"`
	Health              *VsanClusterHealthSummary            `xml:"health,omitempty"`
}

func init() {
	types.Add("vsan:VsanResourceCheckResult", reflect.TypeOf((*VsanResourceCheckResult)(nil)).Elem())
}

type VsanResourceCheckSpec struct {
	types.DynamicData

	Operation       string                        `xml:"operation"`
	Entities        []string                      `xml:"entities,omitempty"`
	MaintenanceSpec *types.HostMaintenanceSpec    `xml:"maintenanceSpec,omitempty"`
	Parent          *types.ManagedObjectReference `xml:"parent,omitempty"`
}

func init() {
	types.Add("vsan:VsanResourceCheckSpec", reflect.TypeOf((*VsanResourceCheckSpec)(nil)).Elem())
}

type VsanResourceCheckStatus struct {
	types.DynamicData

	Status     string                        `xml:"status"`
	Result     *VsanResourceCheckResult      `xml:"result,omitempty"`
	Task       *VsanResourceCheckTaskDetails `xml:"task,omitempty"`
	ParentTask *VsanResourceCheckTaskDetails `xml:"parentTask,omitempty"`
}

func init() {
	types.Add("vsan:VsanResourceCheckStatus", reflect.TypeOf((*VsanResourceCheckStatus)(nil)).Elem())
}

type VsanResourceCheckTaskDetails struct {
	types.DynamicData

	Task            types.ManagedObjectReference  `xml:"task"`
	Host            *types.ManagedObjectReference `xml:"host,omitempty"`
	HostUuid        string                        `xml:"hostUuid,omitempty"`
	MaintenanceSpec *types.HostMaintenanceSpec    `xml:"maintenanceSpec,omitempty"`
}

func init() {
	types.Add("vsan:VsanResourceCheckTaskDetails", reflect.TypeOf((*VsanResourceCheckTaskDetails)(nil)).Elem())
}

type VsanResourceConstraint struct {
	types.DynamicData

	TargetType string `xml:"targetType,omitempty"`
}

func init() {
	types.Add("vsan:VsanResourceConstraint", reflect.TypeOf((*VsanResourceConstraint)(nil)).Elem())
}

type VsanResourceHealth struct {
	types.DynamicData

	Resource    string `xml:"resource"`
	Health      string `xml:"health"`
	Description string `xml:"description,omitempty"`
}

func init() {
	types.Add("vsan:VsanResourceHealth", reflect.TypeOf((*VsanResourceHealth)(nil)).Elem())
}

type VsanRetrieveProperties VsanRetrievePropertiesRequestType

func init() {
	types.Add("vsan:VsanRetrieveProperties", reflect.TypeOf((*VsanRetrieveProperties)(nil)).Elem())
}

type VsanRetrievePropertiesRequestType struct {
	This               types.ManagedObjectReference `xml:"_this"`
	MassCollectorSpecs []VsanMassCollectorSpec      `xml:"massCollectorSpecs"`
}

func init() {
	types.Add("vsan:VsanRetrievePropertiesRequestType", reflect.TypeOf((*VsanRetrievePropertiesRequestType)(nil)).Elem())
}

type VsanRetrievePropertiesResponse struct {
	Returnval []types.ObjectContent `xml:"returnval,omitempty"`
}

type VsanRollbackVdsToVss VsanRollbackVdsToVssRequestType

func init() {
	types.Add("vsan:VsanRollbackVdsToVss", reflect.TypeOf((*VsanRollbackVdsToVss)(nil)).Elem())
}

type VsanRollbackVdsToVssRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
	Task types.ManagedObjectReference `xml:"task"`
}

func init() {
	types.Add("vsan:VsanRollbackVdsToVssRequestType", reflect.TypeOf((*VsanRollbackVdsToVssRequestType)(nil)).Elem())
}

type VsanRollbackVdsToVssResponse struct {
	Returnval bool `xml:"returnval"`
}

type VsanRuntimeStatsHostMap struct {
	types.DynamicData

	Host  types.ManagedObjectReference `xml:"host"`
	Stats *VsanHostRuntimeStats        `xml:"stats,omitempty"`
}

func init() {
	types.Add("vsan:VsanRuntimeStatsHostMap", reflect.TypeOf((*VsanRuntimeStatsHostMap)(nil)).Elem())
}

type VsanSmartDiskStats struct {
	types.DynamicData

	Disk  string                      `xml:"disk"`
	Stats []VsanSmartParameter        `xml:"stats,omitempty"`
	Error *types.LocalizedMethodFault `xml:"error,omitempty"`
}

func init() {
	types.Add("vsan:VsanSmartDiskStats", reflect.TypeOf((*VsanSmartDiskStats)(nil)).Elem())
}

type VsanSmartParameter struct {
	types.DynamicData

	Parameter string `xml:"parameter,omitempty"`
	Value     int32  `xml:"value,omitempty"`
	Threshold int32  `xml:"threshold,omitempty"`
	Worst     int32  `xml:"worst,omitempty"`
}

func init() {
	types.Add("vsan:VsanSmartParameter", reflect.TypeOf((*VsanSmartParameter)(nil)).Elem())
}

type VsanSmartStatsHostSummary struct {
	types.DynamicData

	Hostname   string               `xml:"hostname,omitempty"`
	SmartStats []VsanSmartDiskStats `xml:"smartStats,omitempty"`
}

func init() {
	types.Add("vsan:VsanSmartStatsHostSummary", reflect.TypeOf((*VsanSmartStatsHostSummary)(nil)).Elem())
}

type VsanSpaceQuerySpec struct {
	types.DynamicData

	EntityType string   `xml:"entityType"`
	EntityIds  []string `xml:"entityIds,omitempty"`
}

func init() {
	types.Add("vsan:VsanSpaceQuerySpec", reflect.TypeOf((*VsanSpaceQuerySpec)(nil)).Elem())
}

type VsanSpaceUsage struct {
	types.DynamicData

	TotalCapacityB    int64                               `xml:"totalCapacityB"`
	FreeCapacityB     int64                               `xml:"freeCapacityB,omitempty"`
	SpaceOverview     *VsanObjectSpaceSummary             `xml:"spaceOverview,omitempty"`
	SpaceDetail       *VsanSpaceUsageDetailResult         `xml:"spaceDetail,omitempty"`
	EfficientCapacity *VimVsanDataEfficiencyCapacityState `xml:"efficientCapacity,omitempty"`
	WhatifCapacities  []VsanWhatifCapacity                `xml:"whatifCapacities,omitempty"`
	UncommittedB      int64                               `xml:"uncommittedB,omitempty"`
}

func init() {
	types.Add("vsan:VsanSpaceUsage", reflect.TypeOf((*VsanSpaceUsage)(nil)).Elem())
}

type VsanSpaceUsageDetailResult struct {
	types.DynamicData

	SpaceUsageByObjectType []VsanObjectSpaceSummary `xml:"spaceUsageByObjectType,omitempty"`
}

func init() {
	types.Add("vsan:VsanSpaceUsageDetailResult", reflect.TypeOf((*VsanSpaceUsageDetailResult)(nil)).Elem())
}

type VsanStartProactiveRebalance VsanStartProactiveRebalanceRequestType

func init() {
	types.Add("vsan:VsanStartProactiveRebalance", reflect.TypeOf((*VsanStartProactiveRebalance)(nil)).Elem())
}

type VsanStartProactiveRebalanceRequestType struct {
	This              types.ManagedObjectReference `xml:"_this"`
	TimeSpan          int32                        `xml:"timeSpan,omitempty"`
	VarianceThreshold float32                      `xml:"varianceThreshold,omitempty"`
	TimeThreshold     int32                        `xml:"timeThreshold,omitempty"`
	RateThreshold     int32                        `xml:"rateThreshold,omitempty"`
}

func init() {
	types.Add("vsan:VsanStartProactiveRebalanceRequestType", reflect.TypeOf((*VsanStartProactiveRebalanceRequestType)(nil)).Elem())
}

type VsanStartProactiveRebalanceResponse struct {
	Returnval bool `xml:"returnval"`
}

type VsanStopProactiveRebalance VsanStopProactiveRebalanceRequestType

func init() {
	types.Add("vsan:VsanStopProactiveRebalance", reflect.TypeOf((*VsanStopProactiveRebalance)(nil)).Elem())
}

type VsanStopProactiveRebalanceRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

func init() {
	types.Add("vsan:VsanStopProactiveRebalanceRequestType", reflect.TypeOf((*VsanStopProactiveRebalanceRequestType)(nil)).Elem())
}

type VsanStopProactiveRebalanceResponse struct {
	Returnval bool `xml:"returnval"`
}

type VsanStopRebalanceCluster VsanStopRebalanceClusterRequestType

func init() {
	types.Add("vsan:VsanStopRebalanceCluster", reflect.TypeOf((*VsanStopRebalanceCluster)(nil)).Elem())
}

type VsanStopRebalanceClusterRequestType struct {
	This        types.ManagedObjectReference   `xml:"_this"`
	Cluster     types.ManagedObjectReference   `xml:"cluster"`
	TargetHosts []types.ManagedObjectReference `xml:"targetHosts,omitempty"`
}

func init() {
	types.Add("vsan:VsanStopRebalanceClusterRequestType", reflect.TypeOf((*VsanStopRebalanceClusterRequestType)(nil)).Elem())
}

type VsanStopRebalanceClusterResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanStorageComplianceResult struct {
	types.DynamicData

	CheckTime             *time.Time                    `xml:"checkTime"`
	Profile               string                        `xml:"profile,omitempty"`
	ObjectUUID            string                        `xml:"objectUUID,omitempty"`
	ComplianceStatus      string                        `xml:"complianceStatus"`
	Mismatch              bool                          `xml:"mismatch"`
	ViolatedPolicies      []VsanStoragePolicyStatus     `xml:"violatedPolicies,omitempty"`
	OperationalStatus     *VsanStorageOperationalStatus `xml:"operationalStatus,omitempty"`
	ObjPolicyGenerationId string                        `xml:"objPolicyGenerationId,omitempty"`
}

func init() {
	types.Add("vsan:VsanStorageComplianceResult", reflect.TypeOf((*VsanStorageComplianceResult)(nil)).Elem())
}

type VsanStorageOperationalStatus struct {
	types.DynamicData

	Healthy           *bool      `xml:"healthy"`
	OperationETA      *time.Time `xml:"operationETA"`
	OperationProgress int64      `xml:"operationProgress,omitempty"`
	Transitional      *bool      `xml:"transitional"`
}

func init() {
	types.Add("vsan:VsanStorageOperationalStatus", reflect.TypeOf((*VsanStorageOperationalStatus)(nil)).Elem())
}

type VsanStoragePolicyStatus struct {
	types.DynamicData

	Id            string `xml:"id,omitempty"`
	ExpectedValue string `xml:"expectedValue,omitempty"`
	CurrentValue  string `xml:"currentValue,omitempty"`
}

func init() {
	types.Add("vsan:VsanStoragePolicyStatus", reflect.TypeOf((*VsanStoragePolicyStatus)(nil)).Elem())
}

type VsanStorageWorkloadType struct {
	types.DynamicData

	Specs       []VsanVmdkLoadTestSpec `xml:"specs"`
	TypeId      string                 `xml:"typeId"`
	Name        string                 `xml:"name"`
	Description string                 `xml:"description"`
	Duration    int64                  `xml:"duration,omitempty"`
}

func init() {
	types.Add("vsan:VsanStorageWorkloadType", reflect.TypeOf((*VsanStorageWorkloadType)(nil)).Elem())
}

type VsanSyncingObjectFilter struct {
	types.DynamicData

	ResyncType      string `xml:"resyncType,omitempty"`
	ResyncStatus    string `xml:"resyncStatus,omitempty"`
	NumberOfObjects int64  `xml:"numberOfObjects,omitempty"`
	Offset          int64  `xml:"offset,omitempty"`
}

func init() {
	types.Add("vsan:VsanSyncingObjectFilter", reflect.TypeOf((*VsanSyncingObjectFilter)(nil)).Elem())
}

type VsanSyncingObjectRecoveryDetails struct {
	types.DynamicData

	ActivelySyncingObjectRecoveryETA int64 `xml:"activelySyncingObjectRecoveryETA,omitempty"`
	QueuedForSyncObjectRecoveryETA   int64 `xml:"queuedForSyncObjectRecoveryETA,omitempty"`
	SuspendedObjectRecoveryETA       int64 `xml:"suspendedObjectRecoveryETA,omitempty"`
	ActiveObjectsToSync              int64 `xml:"activeObjectsToSync,omitempty"`
	QueuedObjectsToSync              int64 `xml:"queuedObjectsToSync,omitempty"`
	SuspendedObjectsToSync           int64 `xml:"suspendedObjectsToSync,omitempty"`
	BytesToSyncForActiveObjects      int64 `xml:"bytesToSyncForActiveObjects,omitempty"`
	BytesToSyncForQueuedObjects      int64 `xml:"bytesToSyncForQueuedObjects,omitempty"`
	BytesToSyncForSuspendedObjects   int64 `xml:"bytesToSyncForSuspendedObjects,omitempty"`
}

func init() {
	types.Add("vsan:VsanSyncingObjectRecoveryDetails", reflect.TypeOf((*VsanSyncingObjectRecoveryDetails)(nil)).Elem())
}

type VsanUnicastAddressInfo struct {
	types.DynamicData

	Address string `xml:"address"`
	Port    int32  `xml:"port,omitempty"`
}

func init() {
	types.Add("vsan:VsanUnicastAddressInfo", reflect.TypeOf((*VsanUnicastAddressInfo)(nil)).Elem())
}

type VsanUnknownScanIssue struct {
	VsanUpgradeSystemPreflightCheckIssue

	Uuids []string `xml:"uuids"`
}

func init() {
	types.Add("vsan:VsanUnknownScanIssue", reflect.TypeOf((*VsanUnknownScanIssue)(nil)).Elem())
}

type VsanUnmapConfig struct {
	types.DynamicData

	Enable bool `xml:"enable"`
}

func init() {
	types.Add("vsan:VsanUnmapConfig", reflect.TypeOf((*VsanUnmapConfig)(nil)).Elem())
}

type VsanUnmountDiskMappingEx VsanUnmountDiskMappingExRequestType

func init() {
	types.Add("vsan:VsanUnmountDiskMappingEx", reflect.TypeOf((*VsanUnmountDiskMappingEx)(nil)).Elem())
}

type VsanUnmountDiskMappingExRequestType struct {
	This            types.ManagedObjectReference `xml:"_this"`
	Mappings        []VsanHostDiskMapping        `xml:"mappings"`
	MaintenanceSpec *types.HostMaintenanceSpec   `xml:"maintenanceSpec,omitempty"`
	Timeout         int32                        `xml:"timeout,omitempty"`
}

func init() {
	types.Add("vsan:VsanUnmountDiskMappingExRequestType", reflect.TypeOf((*VsanUnmountDiskMappingExRequestType)(nil)).Elem())
}

type VsanUnmountDiskMappingExResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanUnsupportedHighDiskVersionIssue struct {
	VsanUpgradeSystemPreflightCheckIssue

	Hosts []types.ManagedObjectReference `xml:"hosts"`
}

func init() {
	types.Add("vsan:VsanUnsupportedHighDiskVersionIssue", reflect.TypeOf((*VsanUnsupportedHighDiskVersionIssue)(nil)).Elem())
}

type VsanUpdateItem struct {
	types.DynamicData

	Host            types.ManagedObjectReference `xml:"host"`
	Type            string                       `xml:"type"`
	Name            string                       `xml:"name"`
	Version         string                       `xml:"version"`
	ExistingVersion string                       `xml:"existingVersion,omitempty"`
	Present         bool                         `xml:"present"`
	VibSpec         []VsanVibSpec                `xml:"vibSpec,omitempty"`
	VibType         string                       `xml:"vibType,omitempty"`
	FirmwareSpec    *VsanHclFirmwareUpdateSpec   `xml:"firmwareSpec,omitempty"`
	DownloadInfo    []VsanDownloadItem           `xml:"downloadInfo,omitempty"`
	Eula            string                       `xml:"eula,omitempty"`
	Adapter         string                       `xml:"adapter,omitempty"`
	Key             string                       `xml:"key,omitempty"`
	Impact          string                       `xml:"impact,omitempty"`
	FirmwareUnknown *bool                        `xml:"firmwareUnknown"`
}

func init() {
	types.Add("vsan:VsanUpdateItem", reflect.TypeOf((*VsanUpdateItem)(nil)).Elem())
}

type VsanUpgradeFsvm VsanUpgradeFsvmRequestType

func init() {
	types.Add("vsan:VsanUpgradeFsvm", reflect.TypeOf((*VsanUpgradeFsvm)(nil)).Elem())
}

type VsanUpgradeFsvmRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

func init() {
	types.Add("vsan:VsanUpgradeFsvmRequestType", reflect.TypeOf((*VsanUpgradeFsvmRequestType)(nil)).Elem())
}

type VsanUpgradeFsvmResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanUpgradeStatusEx struct {
	VsanUpgradeSystemUpgradeStatus

	IsPrecheck     *bool                                `xml:"isPrecheck"`
	PrecheckResult *VsanDiskFormatConversionCheckResult `xml:"precheckResult,omitempty"`
}

func init() {
	types.Add("vsan:VsanUpgradeStatusEx", reflect.TypeOf((*VsanUpgradeStatusEx)(nil)).Elem())
}

type VsanVcClusterGetHclInfo VsanVcClusterGetHclInfoRequestType

func init() {
	types.Add("vsan:VsanVcClusterGetHclInfo", reflect.TypeOf((*VsanVcClusterGetHclInfo)(nil)).Elem())
}

type VsanVcClusterGetHclInfoRequestType struct {
	This               types.ManagedObjectReference  `xml:"_this"`
	Cluster            *types.ManagedObjectReference `xml:"cluster,omitempty"`
	IncludeHostsResult *bool                         `xml:"includeHostsResult"`
	IncludeVendorInfo  *bool                         `xml:"includeVendorInfo"`
	EsxRelease         string                        `xml:"esxRelease,omitempty"`
}

func init() {
	types.Add("vsan:VsanVcClusterGetHclInfoRequestType", reflect.TypeOf((*VsanVcClusterGetHclInfoRequestType)(nil)).Elem())
}

type VsanVcClusterGetHclInfoResponse struct {
	Returnval VsanClusterHclInfo `xml:"returnval"`
}

type VsanVcClusterQueryVerifyHealthSystemVersions VsanVcClusterQueryVerifyHealthSystemVersionsRequestType

func init() {
	types.Add("vsan:VsanVcClusterQueryVerifyHealthSystemVersions", reflect.TypeOf((*VsanVcClusterQueryVerifyHealthSystemVersions)(nil)).Elem())
}

type VsanVcClusterQueryVerifyHealthSystemVersionsRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

func init() {
	types.Add("vsan:VsanVcClusterQueryVerifyHealthSystemVersionsRequestType", reflect.TypeOf((*VsanVcClusterQueryVerifyHealthSystemVersionsRequestType)(nil)).Elem())
}

type VsanVcClusterQueryVerifyHealthSystemVersionsResponse struct {
	Returnval VsanClusterHealthSystemVersionResult `xml:"returnval"`
}

type VsanVcClusterRunVmdkLoadTest VsanVcClusterRunVmdkLoadTestRequestType

func init() {
	types.Add("vsan:VsanVcClusterRunVmdkLoadTest", reflect.TypeOf((*VsanVcClusterRunVmdkLoadTest)(nil)).Elem())
}

type VsanVcClusterRunVmdkLoadTestRequestType struct {
	This        types.ManagedObjectReference `xml:"_this"`
	Cluster     types.ManagedObjectReference `xml:"cluster"`
	Runname     string                       `xml:"runname"`
	DurationSec int32                        `xml:"durationSec,omitempty"`
	Specs       []VsanVmdkLoadTestSpec       `xml:"specs,omitempty"`
	Action      string                       `xml:"action,omitempty"`
}

func init() {
	types.Add("vsan:VsanVcClusterRunVmdkLoadTestRequestType", reflect.TypeOf((*VsanVcClusterRunVmdkLoadTestRequestType)(nil)).Elem())
}

type VsanVcClusterRunVmdkLoadTestResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanVcKmipServersHealth struct {
	types.DynamicData

	Health               string                      `xml:"health,omitempty"`
	Error                *types.LocalizedMethodFault `xml:"error,omitempty"`
	KmsProviderId        string                      `xml:"kmsProviderId,omitempty"`
	KmsHealth            []VsanKmsHealth             `xml:"kmsHealth,omitempty"`
	ClientCertHealth     string                      `xml:"clientCertHealth,omitempty"`
	ClientCertExpireDate *time.Time                  `xml:"clientCertExpireDate"`
	IsAwsKms             *bool                       `xml:"isAwsKms"`
	CmkHealth            string                      `xml:"cmkHealth,omitempty"`
}

func init() {
	types.Add("vsan:VsanVcKmipServersHealth", reflect.TypeOf((*VsanVcKmipServersHealth)(nil)).Elem())
}

type VsanVcPostDeployConfigSpec struct {
	types.DynamicData

	DcName                   string                    `xml:"dcName,omitempty"`
	ClusterName              string                    `xml:"clusterName,omitempty"`
	FirstHost                *types.HostConnectSpec    `xml:"firstHost,omitempty"`
	HostsToAdd               []types.HostConnectSpec   `xml:"hostsToAdd,omitempty"`
	VsanDataEfficiencyConfig *VsanDataEfficiencyConfig `xml:"vsanDataEfficiencyConfig,omitempty"`
	VsanLicenseKey           string                    `xml:"vsanLicenseKey,omitempty"`
	HostLicenseKey           string                    `xml:"hostLicenseKey,omitempty"`
	TaskId                   string                    `xml:"taskId,omitempty"`
	VsanDataEncryptionConfig *VsanHostEncryptionInfo   `xml:"vsanDataEncryptionConfig,omitempty"`
}

func init() {
	types.Add("vsan:VsanVcPostDeployConfigSpec", reflect.TypeOf((*VsanVcPostDeployConfigSpec)(nil)).Elem())
}

type VsanVcUpdateHclDbFromWeb VsanVcUpdateHclDbFromWebRequestType

func init() {
	types.Add("vsan:VsanVcUpdateHclDbFromWeb", reflect.TypeOf((*VsanVcUpdateHclDbFromWeb)(nil)).Elem())
}

type VsanVcUpdateHclDbFromWebRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
	Url  string                       `xml:"url,omitempty"`
}

func init() {
	types.Add("vsan:VsanVcUpdateHclDbFromWebRequestType", reflect.TypeOf((*VsanVcUpdateHclDbFromWebRequestType)(nil)).Elem())
}

type VsanVcUpdateHclDbFromWebResponse struct {
	Returnval bool `xml:"returnval"`
}

type VsanVcUploadHclDb VsanVcUploadHclDbRequestType

func init() {
	types.Add("vsan:VsanVcUploadHclDb", reflect.TypeOf((*VsanVcUploadHclDb)(nil)).Elem())
}

type VsanVcUploadHclDbRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
	Db   string                       `xml:"db"`
}

func init() {
	types.Add("vsan:VsanVcUploadHclDbRequestType", reflect.TypeOf((*VsanVcUploadHclDbRequestType)(nil)).Elem())
}

type VsanVcUploadHclDbResponse struct {
	Returnval bool `xml:"returnval"`
}

type VsanVcUploadReleaseDb VsanVcUploadReleaseDbRequestType

func init() {
	types.Add("vsan:VsanVcUploadReleaseDb", reflect.TypeOf((*VsanVcUploadReleaseDb)(nil)).Elem())
}

type VsanVcUploadReleaseDbRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
	Db   string                       `xml:"db"`
}

func init() {
	types.Add("vsan:VsanVcUploadReleaseDbRequestType", reflect.TypeOf((*VsanVcUploadReleaseDbRequestType)(nil)).Elem())
}

type VsanVcUploadReleaseDbResponse struct {
}

type VsanVcsaDeploymentProgress struct {
	types.DynamicData

	Phase         string                        `xml:"phase"`
	ProgressPct   int64                         `xml:"progressPct"`
	Message       string                        `xml:"message"`
	Success       bool                          `xml:"success"`
	Error         *types.LocalizedMethodFault   `xml:"error,omitempty"`
	UpdateCounter int64                         `xml:"updateCounter"`
	TaskId        string                        `xml:"taskId,omitempty"`
	Vm            *types.ManagedObjectReference `xml:"vm,omitempty"`
}

func init() {
	types.Add("vsan:VsanVcsaDeploymentProgress", reflect.TypeOf((*VsanVcsaDeploymentProgress)(nil)).Elem())
}

type VsanVcsaGetBootstrapProgress VsanVcsaGetBootstrapProgressRequestType

func init() {
	types.Add("vsan:VsanVcsaGetBootstrapProgress", reflect.TypeOf((*VsanVcsaGetBootstrapProgress)(nil)).Elem())
}

type VsanVcsaGetBootstrapProgressRequestType struct {
	This   types.ManagedObjectReference `xml:"_this"`
	TaskId []string                     `xml:"taskId"`
}

func init() {
	types.Add("vsan:VsanVcsaGetBootstrapProgressRequestType", reflect.TypeOf((*VsanVcsaGetBootstrapProgressRequestType)(nil)).Elem())
}

type VsanVcsaGetBootstrapProgressResponse struct {
	Returnval []VsanVcsaDeploymentProgress `xml:"returnval"`
}

type VsanVdsGetMigrationPlan VsanVdsGetMigrationPlanRequestType

func init() {
	types.Add("vsan:VsanVdsGetMigrationPlan", reflect.TypeOf((*VsanVdsGetMigrationPlan)(nil)).Elem())
}

type VsanVdsGetMigrationPlanRequestType struct {
	This         types.ManagedObjectReference   `xml:"_this"`
	Cluster      types.ManagedObjectReference   `xml:"cluster"`
	VswitchName  string                         `xml:"vswitchName,omitempty"`
	VdsName      string                         `xml:"vdsName,omitempty"`
	VmnicDevices []string                       `xml:"vmnicDevices,omitempty"`
	InfraVm      []types.ManagedObjectReference `xml:"infraVm,omitempty"`
	Vds          *types.ManagedObjectReference  `xml:"vds,omitempty"`
	Hosts        []types.ManagedObjectReference `xml:"hosts,omitempty"`
}

func init() {
	types.Add("vsan:VsanVdsGetMigrationPlanRequestType", reflect.TypeOf((*VsanVdsGetMigrationPlanRequestType)(nil)).Elem())
}

type VsanVdsGetMigrationPlanResponse struct {
	Returnval VsanVdsMigrationPlan `xml:"returnval"`
}

type VsanVdsMigrateVss VsanVdsMigrateVssRequestType

func init() {
	types.Add("vsan:VsanVdsMigrateVss", reflect.TypeOf((*VsanVdsMigrateVss)(nil)).Elem())
}

type VsanVdsMigrateVssRequestType struct {
	This          types.ManagedObjectReference   `xml:"_this"`
	Cluster       types.ManagedObjectReference   `xml:"cluster"`
	MigrationPlan *VsanVdsMigrationPlan          `xml:"migrationPlan,omitempty"`
	VswitchName   string                         `xml:"vswitchName,omitempty"`
	VdsName       string                         `xml:"vdsName,omitempty"`
	VmnicDevices  []string                       `xml:"vmnicDevices,omitempty"`
	InfraVm       []types.ManagedObjectReference `xml:"infraVm,omitempty"`
	Vds           *types.ManagedObjectReference  `xml:"vds,omitempty"`
	Hosts         []types.ManagedObjectReference `xml:"hosts,omitempty"`
}

func init() {
	types.Add("vsan:VsanVdsMigrateVssRequestType", reflect.TypeOf((*VsanVdsMigrateVssRequestType)(nil)).Elem())
}

type VsanVdsMigrateVssResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanVdsMigrationPlan struct {
	types.DynamicData

	VdsSpec         types.DVSCreateSpec            `xml:"vdsSpec"`
	Pgs             []VsanVdsPgMigrationSpec       `xml:"pgs,omitempty"`
	InaccessibleVms []types.ManagedObjectReference `xml:"inaccessibleVms,omitempty"`
	InfraVms        []types.ManagedObjectReference `xml:"infraVms,omitempty"`
}

func init() {
	types.Add("vsan:VsanVdsMigrationPlan", reflect.TypeOf((*VsanVdsMigrationPlan)(nil)).Elem())
}

type VsanVdsPgMigrationHostInfo struct {
	types.DynamicData

	Host          types.ManagedObjectReference `xml:"host"`
	Hostname      string                       `xml:"hostname"`
	VmknicDevices []string                     `xml:"vmknicDevices,omitempty"`
	VmVnics       []VsanVdsPgMigrationVmInfo   `xml:"vmVnics,omitempty"`
}

func init() {
	types.Add("vsan:VsanVdsPgMigrationHostInfo", reflect.TypeOf((*VsanVdsPgMigrationHostInfo)(nil)).Elem())
}

type VsanVdsPgMigrationSpec struct {
	types.DynamicData

	VssPgName       string                       `xml:"vssPgName"`
	DvPgName        string                       `xml:"dvPgName"`
	VdsPgSetting    types.VMwareDVSPortSetting   `xml:"vdsPgSetting"`
	VdsPgType       string                       `xml:"vdsPgType"`
	Hosts           []VsanVdsPgMigrationHostInfo `xml:"hosts,omitempty"`
	CollisionRename bool                         `xml:"collisionRename"`
}

func init() {
	types.Add("vsan:VsanVdsPgMigrationSpec", reflect.TypeOf((*VsanVdsPgMigrationSpec)(nil)).Elem())
}

type VsanVdsPgMigrationVmInfo struct {
	types.DynamicData

	Vm        types.ManagedObjectReference `xml:"vm"`
	VnicLabel []string                     `xml:"vnicLabel"`
}

func init() {
	types.Add("vsan:VsanVdsPgMigrationVmInfo", reflect.TypeOf((*VsanVdsPgMigrationVmInfo)(nil)).Elem())
}

type VsanVibInstallPreflightCheck VsanVibInstallPreflightCheckRequestType

func init() {
	types.Add("vsan:VsanVibInstallPreflightCheck", reflect.TypeOf((*VsanVibInstallPreflightCheck)(nil)).Elem())
}

type VsanVibInstallPreflightCheckRequestType struct {
	This    types.ManagedObjectReference  `xml:"_this"`
	Cluster *types.ManagedObjectReference `xml:"cluster,omitempty"`
}

func init() {
	types.Add("vsan:VsanVibInstallPreflightCheckRequestType", reflect.TypeOf((*VsanVibInstallPreflightCheckRequestType)(nil)).Elem())
}

type VsanVibInstallPreflightCheckResponse struct {
	Returnval VsanVibInstallPreflightStatus `xml:"returnval"`
}

type VsanVibInstallPreflightStatus struct {
	types.DynamicData

	ManualVmotionRequired bool `xml:"manualVmotionRequired"`
	RollingRequired       bool `xml:"rollingRequired"`
}

func init() {
	types.Add("vsan:VsanVibInstallPreflightStatus", reflect.TypeOf((*VsanVibInstallPreflightStatus)(nil)).Elem())
}

type VsanVibInstallRequestType struct {
	This            types.ManagedObjectReference  `xml:"_this"`
	Cluster         *types.ManagedObjectReference `xml:"cluster,omitempty"`
	VibSpecs        []VsanVibSpec                 `xml:"vibSpecs,omitempty"`
	ScanResults     []VsanVibScanResult           `xml:"scanResults,omitempty"`
	FirmwareSpecs   []VsanHclFirmwareUpdateSpec   `xml:"firmwareSpecs,omitempty"`
	MaintenanceSpec *types.HostMaintenanceSpec    `xml:"maintenanceSpec,omitempty"`
	Rolling         *bool                         `xml:"rolling"`
	NoSigCheck      *bool                         `xml:"noSigCheck"`
}

func init() {
	types.Add("vsan:VsanVibInstallRequestType", reflect.TypeOf((*VsanVibInstallRequestType)(nil)).Elem())
}

type VsanVibInstall_Task VsanVibInstallRequestType

func init() {
	types.Add("vsan:VsanVibInstall_Task", reflect.TypeOf((*VsanVibInstall_Task)(nil)).Elem())
}

type VsanVibInstall_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanVibScan VsanVibScanRequestType

func init() {
	types.Add("vsan:VsanVibScan", reflect.TypeOf((*VsanVibScan)(nil)).Elem())
}

type VsanVibScanRequestType struct {
	This     types.ManagedObjectReference  `xml:"_this"`
	Cluster  *types.ManagedObjectReference `xml:"cluster,omitempty"`
	VibSpecs []VsanVibSpec                 `xml:"vibSpecs"`
}

func init() {
	types.Add("vsan:VsanVibScanRequestType", reflect.TypeOf((*VsanVibScanRequestType)(nil)).Elem())
}

type VsanVibScanResponse struct {
	Returnval []VsanVibScanResult `xml:"returnval,omitempty"`
}

type VsanVibScanResult struct {
	types.DynamicData

	Host                    types.ManagedObjectReference `xml:"host"`
	VibName                 string                       `xml:"vibName"`
	VibVersion              string                       `xml:"vibVersion"`
	ExistingVersion         string                       `xml:"existingVersion,omitempty"`
	MaintenanceModeRequired bool                         `xml:"maintenanceModeRequired"`
	RebootRequired          bool                         `xml:"rebootRequired"`
	MeetsSystemReq          bool                         `xml:"meetsSystemReq"`
	PkgDepsMetByHost        bool                         `xml:"pkgDepsMetByHost"`
}

func init() {
	types.Add("vsan:VsanVibScanResult", reflect.TypeOf((*VsanVibScanResult)(nil)).Elem())
}

type VsanVibSpec struct {
	types.DynamicData

	Host        types.ManagedObjectReference `xml:"host"`
	MetaUrl     string                       `xml:"metaUrl,omitempty"`
	MetaSha1Sum string                       `xml:"metaSha1Sum,omitempty"`
	VibUrl      string                       `xml:"vibUrl"`
	VibSha1Sum  string                       `xml:"vibSha1Sum"`
}

func init() {
	types.Add("vsan:VsanVibSpec", reflect.TypeOf((*VsanVibSpec)(nil)).Elem())
}

type VsanVitAddIscsiInitiatorGroup VsanVitAddIscsiInitiatorGroupRequestType

func init() {
	types.Add("vsan:VsanVitAddIscsiInitiatorGroup", reflect.TypeOf((*VsanVitAddIscsiInitiatorGroup)(nil)).Elem())
}

type VsanVitAddIscsiInitiatorGroupRequestType struct {
	This               types.ManagedObjectReference `xml:"_this"`
	Cluster            types.ManagedObjectReference `xml:"cluster"`
	InitiatorGroupName string                       `xml:"initiatorGroupName"`
}

func init() {
	types.Add("vsan:VsanVitAddIscsiInitiatorGroupRequestType", reflect.TypeOf((*VsanVitAddIscsiInitiatorGroupRequestType)(nil)).Elem())
}

type VsanVitAddIscsiInitiatorGroupResponse struct {
}

type VsanVitAddIscsiInitiatorsToGroup VsanVitAddIscsiInitiatorsToGroupRequestType

func init() {
	types.Add("vsan:VsanVitAddIscsiInitiatorsToGroup", reflect.TypeOf((*VsanVitAddIscsiInitiatorsToGroup)(nil)).Elem())
}

type VsanVitAddIscsiInitiatorsToGroupRequestType struct {
	This               types.ManagedObjectReference `xml:"_this"`
	Cluster            types.ManagedObjectReference `xml:"cluster"`
	InitiatorGroupName string                       `xml:"initiatorGroupName"`
	InitiatorNames     []string                     `xml:"initiatorNames"`
}

func init() {
	types.Add("vsan:VsanVitAddIscsiInitiatorsToGroupRequestType", reflect.TypeOf((*VsanVitAddIscsiInitiatorsToGroupRequestType)(nil)).Elem())
}

type VsanVitAddIscsiInitiatorsToGroupResponse struct {
}

type VsanVitAddIscsiInitiatorsToTarget VsanVitAddIscsiInitiatorsToTargetRequestType

func init() {
	types.Add("vsan:VsanVitAddIscsiInitiatorsToTarget", reflect.TypeOf((*VsanVitAddIscsiInitiatorsToTarget)(nil)).Elem())
}

type VsanVitAddIscsiInitiatorsToTargetRequestType struct {
	This           types.ManagedObjectReference `xml:"_this"`
	Cluster        types.ManagedObjectReference `xml:"cluster"`
	TargetAlias    string                       `xml:"targetAlias"`
	InitiatorNames []string                     `xml:"initiatorNames"`
}

func init() {
	types.Add("vsan:VsanVitAddIscsiInitiatorsToTargetRequestType", reflect.TypeOf((*VsanVitAddIscsiInitiatorsToTargetRequestType)(nil)).Elem())
}

type VsanVitAddIscsiInitiatorsToTargetResponse struct {
}

type VsanVitAddIscsiLUN VsanVitAddIscsiLUNRequestType

func init() {
	types.Add("vsan:VsanVitAddIscsiLUN", reflect.TypeOf((*VsanVitAddIscsiLUN)(nil)).Elem())
}

type VsanVitAddIscsiLUNRequestType struct {
	This        types.ManagedObjectReference `xml:"_this"`
	Cluster     types.ManagedObjectReference `xml:"cluster"`
	TargetAlias string                       `xml:"targetAlias"`
	LunSpec     VsanIscsiLUNSpec             `xml:"lunSpec"`
}

func init() {
	types.Add("vsan:VsanVitAddIscsiLUNRequestType", reflect.TypeOf((*VsanVitAddIscsiLUNRequestType)(nil)).Elem())
}

type VsanVitAddIscsiLUNResponse struct {
	Returnval *types.ManagedObjectReference `xml:"returnval,omitempty"`
}

type VsanVitAddIscsiTarget VsanVitAddIscsiTargetRequestType

func init() {
	types.Add("vsan:VsanVitAddIscsiTarget", reflect.TypeOf((*VsanVitAddIscsiTarget)(nil)).Elem())
}

type VsanVitAddIscsiTargetRequestType struct {
	This       types.ManagedObjectReference `xml:"_this"`
	Cluster    types.ManagedObjectReference `xml:"cluster"`
	TargetSpec VsanIscsiTargetSpec          `xml:"targetSpec"`
}

func init() {
	types.Add("vsan:VsanVitAddIscsiTargetRequestType", reflect.TypeOf((*VsanVitAddIscsiTargetRequestType)(nil)).Elem())
}

type VsanVitAddIscsiTargetResponse struct {
	Returnval *types.ManagedObjectReference `xml:"returnval,omitempty"`
}

type VsanVitAddIscsiTargetToGroup VsanVitAddIscsiTargetToGroupRequestType

func init() {
	types.Add("vsan:VsanVitAddIscsiTargetToGroup", reflect.TypeOf((*VsanVitAddIscsiTargetToGroup)(nil)).Elem())
}

type VsanVitAddIscsiTargetToGroupRequestType struct {
	This               types.ManagedObjectReference `xml:"_this"`
	Cluster            types.ManagedObjectReference `xml:"cluster"`
	InitiatorGroupName string                       `xml:"initiatorGroupName"`
	TargetAlias        string                       `xml:"targetAlias"`
}

func init() {
	types.Add("vsan:VsanVitAddIscsiTargetToGroupRequestType", reflect.TypeOf((*VsanVitAddIscsiTargetToGroupRequestType)(nil)).Elem())
}

type VsanVitAddIscsiTargetToGroupResponse struct {
}

type VsanVitEditIscsiLUN VsanVitEditIscsiLUNRequestType

func init() {
	types.Add("vsan:VsanVitEditIscsiLUN", reflect.TypeOf((*VsanVitEditIscsiLUN)(nil)).Elem())
}

type VsanVitEditIscsiLUNRequestType struct {
	This        types.ManagedObjectReference `xml:"_this"`
	Cluster     types.ManagedObjectReference `xml:"cluster"`
	TargetAlias string                       `xml:"targetAlias"`
	LunSpec     VsanIscsiLUNSpec             `xml:"lunSpec"`
}

func init() {
	types.Add("vsan:VsanVitEditIscsiLUNRequestType", reflect.TypeOf((*VsanVitEditIscsiLUNRequestType)(nil)).Elem())
}

type VsanVitEditIscsiLUNResponse struct {
	Returnval *types.ManagedObjectReference `xml:"returnval,omitempty"`
}

type VsanVitEditIscsiTarget VsanVitEditIscsiTargetRequestType

func init() {
	types.Add("vsan:VsanVitEditIscsiTarget", reflect.TypeOf((*VsanVitEditIscsiTarget)(nil)).Elem())
}

type VsanVitEditIscsiTargetRequestType struct {
	This       types.ManagedObjectReference `xml:"_this"`
	Cluster    types.ManagedObjectReference `xml:"cluster"`
	TargetSpec VsanIscsiTargetSpec          `xml:"targetSpec"`
}

func init() {
	types.Add("vsan:VsanVitEditIscsiTargetRequestType", reflect.TypeOf((*VsanVitEditIscsiTargetRequestType)(nil)).Elem())
}

type VsanVitEditIscsiTargetResponse struct {
	Returnval *types.ManagedObjectReference `xml:"returnval,omitempty"`
}

type VsanVitGetHomeObject VsanVitGetHomeObjectRequestType

func init() {
	types.Add("vsan:VsanVitGetHomeObject", reflect.TypeOf((*VsanVitGetHomeObject)(nil)).Elem())
}

type VsanVitGetHomeObjectRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

func init() {
	types.Add("vsan:VsanVitGetHomeObjectRequestType", reflect.TypeOf((*VsanVitGetHomeObjectRequestType)(nil)).Elem())
}

type VsanVitGetHomeObjectResponse struct {
	Returnval VsanObjectInformation `xml:"returnval"`
}

type VsanVitGetIscsiInitiatorGroup VsanVitGetIscsiInitiatorGroupRequestType

func init() {
	types.Add("vsan:VsanVitGetIscsiInitiatorGroup", reflect.TypeOf((*VsanVitGetIscsiInitiatorGroup)(nil)).Elem())
}

type VsanVitGetIscsiInitiatorGroupRequestType struct {
	This               types.ManagedObjectReference `xml:"_this"`
	Cluster            types.ManagedObjectReference `xml:"cluster"`
	InitiatorGroupName string                       `xml:"initiatorGroupName"`
}

func init() {
	types.Add("vsan:VsanVitGetIscsiInitiatorGroupRequestType", reflect.TypeOf((*VsanVitGetIscsiInitiatorGroupRequestType)(nil)).Elem())
}

type VsanVitGetIscsiInitiatorGroupResponse struct {
	Returnval *VsanIscsiInitiatorGroup `xml:"returnval,omitempty"`
}

type VsanVitGetIscsiInitiatorGroups VsanVitGetIscsiInitiatorGroupsRequestType

func init() {
	types.Add("vsan:VsanVitGetIscsiInitiatorGroups", reflect.TypeOf((*VsanVitGetIscsiInitiatorGroups)(nil)).Elem())
}

type VsanVitGetIscsiInitiatorGroupsRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

func init() {
	types.Add("vsan:VsanVitGetIscsiInitiatorGroupsRequestType", reflect.TypeOf((*VsanVitGetIscsiInitiatorGroupsRequestType)(nil)).Elem())
}

type VsanVitGetIscsiInitiatorGroupsResponse struct {
	Returnval []VsanIscsiInitiatorGroup `xml:"returnval,omitempty"`
}

type VsanVitGetIscsiLUN VsanVitGetIscsiLUNRequestType

func init() {
	types.Add("vsan:VsanVitGetIscsiLUN", reflect.TypeOf((*VsanVitGetIscsiLUN)(nil)).Elem())
}

type VsanVitGetIscsiLUNRequestType struct {
	This        types.ManagedObjectReference `xml:"_this"`
	Cluster     types.ManagedObjectReference `xml:"cluster"`
	TargetAlias string                       `xml:"targetAlias"`
	LunId       int32                        `xml:"lunId"`
}

func init() {
	types.Add("vsan:VsanVitGetIscsiLUNRequestType", reflect.TypeOf((*VsanVitGetIscsiLUNRequestType)(nil)).Elem())
}

type VsanVitGetIscsiLUNResponse struct {
	Returnval *VsanIscsiLUN `xml:"returnval,omitempty"`
}

type VsanVitGetIscsiLUNs VsanVitGetIscsiLUNsRequestType

func init() {
	types.Add("vsan:VsanVitGetIscsiLUNs", reflect.TypeOf((*VsanVitGetIscsiLUNs)(nil)).Elem())
}

type VsanVitGetIscsiLUNsRequestType struct {
	This          types.ManagedObjectReference `xml:"_this"`
	Cluster       types.ManagedObjectReference `xml:"cluster"`
	TargetAliases []string                     `xml:"targetAliases,omitempty"`
}

func init() {
	types.Add("vsan:VsanVitGetIscsiLUNsRequestType", reflect.TypeOf((*VsanVitGetIscsiLUNsRequestType)(nil)).Elem())
}

type VsanVitGetIscsiLUNsResponse struct {
	Returnval []VsanIscsiLUN `xml:"returnval,omitempty"`
}

type VsanVitGetIscsiTarget VsanVitGetIscsiTargetRequestType

func init() {
	types.Add("vsan:VsanVitGetIscsiTarget", reflect.TypeOf((*VsanVitGetIscsiTarget)(nil)).Elem())
}

type VsanVitGetIscsiTargetRequestType struct {
	This        types.ManagedObjectReference `xml:"_this"`
	Cluster     types.ManagedObjectReference `xml:"cluster"`
	TargetAlias string                       `xml:"targetAlias"`
}

func init() {
	types.Add("vsan:VsanVitGetIscsiTargetRequestType", reflect.TypeOf((*VsanVitGetIscsiTargetRequestType)(nil)).Elem())
}

type VsanVitGetIscsiTargetResponse struct {
	Returnval *VsanIscsiTarget `xml:"returnval,omitempty"`
}

type VsanVitGetIscsiTargets VsanVitGetIscsiTargetsRequestType

func init() {
	types.Add("vsan:VsanVitGetIscsiTargets", reflect.TypeOf((*VsanVitGetIscsiTargets)(nil)).Elem())
}

type VsanVitGetIscsiTargetsRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

func init() {
	types.Add("vsan:VsanVitGetIscsiTargetsRequestType", reflect.TypeOf((*VsanVitGetIscsiTargetsRequestType)(nil)).Elem())
}

type VsanVitGetIscsiTargetsResponse struct {
	Returnval []VsanIscsiTarget `xml:"returnval,omitempty"`
}

type VsanVitQueryIscsiTargetServiceVersion VsanVitQueryIscsiTargetServiceVersionRequestType

func init() {
	types.Add("vsan:VsanVitQueryIscsiTargetServiceVersion", reflect.TypeOf((*VsanVitQueryIscsiTargetServiceVersion)(nil)).Elem())
}

type VsanVitQueryIscsiTargetServiceVersionRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

func init() {
	types.Add("vsan:VsanVitQueryIscsiTargetServiceVersionRequestType", reflect.TypeOf((*VsanVitQueryIscsiTargetServiceVersionRequestType)(nil)).Elem())
}

type VsanVitQueryIscsiTargetServiceVersionResponse struct {
	Returnval string `xml:"returnval"`
}

type VsanVitRemoveIscsiInitiatorGroup VsanVitRemoveIscsiInitiatorGroupRequestType

func init() {
	types.Add("vsan:VsanVitRemoveIscsiInitiatorGroup", reflect.TypeOf((*VsanVitRemoveIscsiInitiatorGroup)(nil)).Elem())
}

type VsanVitRemoveIscsiInitiatorGroupRequestType struct {
	This               types.ManagedObjectReference `xml:"_this"`
	Cluster            types.ManagedObjectReference `xml:"cluster"`
	InitiatorGroupName string                       `xml:"initiatorGroupName"`
}

func init() {
	types.Add("vsan:VsanVitRemoveIscsiInitiatorGroupRequestType", reflect.TypeOf((*VsanVitRemoveIscsiInitiatorGroupRequestType)(nil)).Elem())
}

type VsanVitRemoveIscsiInitiatorGroupResponse struct {
}

type VsanVitRemoveIscsiInitiatorsFromGroup VsanVitRemoveIscsiInitiatorsFromGroupRequestType

func init() {
	types.Add("vsan:VsanVitRemoveIscsiInitiatorsFromGroup", reflect.TypeOf((*VsanVitRemoveIscsiInitiatorsFromGroup)(nil)).Elem())
}

type VsanVitRemoveIscsiInitiatorsFromGroupRequestType struct {
	This               types.ManagedObjectReference `xml:"_this"`
	Cluster            types.ManagedObjectReference `xml:"cluster"`
	InitiatorGroupName string                       `xml:"initiatorGroupName"`
	InitiatorNames     []string                     `xml:"initiatorNames"`
}

func init() {
	types.Add("vsan:VsanVitRemoveIscsiInitiatorsFromGroupRequestType", reflect.TypeOf((*VsanVitRemoveIscsiInitiatorsFromGroupRequestType)(nil)).Elem())
}

type VsanVitRemoveIscsiInitiatorsFromGroupResponse struct {
}

type VsanVitRemoveIscsiInitiatorsFromTarget VsanVitRemoveIscsiInitiatorsFromTargetRequestType

func init() {
	types.Add("vsan:VsanVitRemoveIscsiInitiatorsFromTarget", reflect.TypeOf((*VsanVitRemoveIscsiInitiatorsFromTarget)(nil)).Elem())
}

type VsanVitRemoveIscsiInitiatorsFromTargetRequestType struct {
	This           types.ManagedObjectReference `xml:"_this"`
	Cluster        types.ManagedObjectReference `xml:"cluster"`
	TargetAlias    string                       `xml:"targetAlias"`
	InitiatorNames []string                     `xml:"initiatorNames"`
}

func init() {
	types.Add("vsan:VsanVitRemoveIscsiInitiatorsFromTargetRequestType", reflect.TypeOf((*VsanVitRemoveIscsiInitiatorsFromTargetRequestType)(nil)).Elem())
}

type VsanVitRemoveIscsiInitiatorsFromTargetResponse struct {
}

type VsanVitRemoveIscsiLUN VsanVitRemoveIscsiLUNRequestType

func init() {
	types.Add("vsan:VsanVitRemoveIscsiLUN", reflect.TypeOf((*VsanVitRemoveIscsiLUN)(nil)).Elem())
}

type VsanVitRemoveIscsiLUNRequestType struct {
	This        types.ManagedObjectReference `xml:"_this"`
	Cluster     types.ManagedObjectReference `xml:"cluster"`
	TargetAlias string                       `xml:"targetAlias"`
	LunId       int32                        `xml:"lunId"`
}

func init() {
	types.Add("vsan:VsanVitRemoveIscsiLUNRequestType", reflect.TypeOf((*VsanVitRemoveIscsiLUNRequestType)(nil)).Elem())
}

type VsanVitRemoveIscsiLUNResponse struct {
	Returnval *types.ManagedObjectReference `xml:"returnval,omitempty"`
}

type VsanVitRemoveIscsiTarget VsanVitRemoveIscsiTargetRequestType

func init() {
	types.Add("vsan:VsanVitRemoveIscsiTarget", reflect.TypeOf((*VsanVitRemoveIscsiTarget)(nil)).Elem())
}

type VsanVitRemoveIscsiTargetFromGroup VsanVitRemoveIscsiTargetFromGroupRequestType

func init() {
	types.Add("vsan:VsanVitRemoveIscsiTargetFromGroup", reflect.TypeOf((*VsanVitRemoveIscsiTargetFromGroup)(nil)).Elem())
}

type VsanVitRemoveIscsiTargetFromGroupRequestType struct {
	This               types.ManagedObjectReference `xml:"_this"`
	Cluster            types.ManagedObjectReference `xml:"cluster"`
	InitiatorGroupName string                       `xml:"initiatorGroupName"`
	TargetAlias        string                       `xml:"targetAlias"`
}

func init() {
	types.Add("vsan:VsanVitRemoveIscsiTargetFromGroupRequestType", reflect.TypeOf((*VsanVitRemoveIscsiTargetFromGroupRequestType)(nil)).Elem())
}

type VsanVitRemoveIscsiTargetFromGroupResponse struct {
}

type VsanVitRemoveIscsiTargetRequestType struct {
	This        types.ManagedObjectReference `xml:"_this"`
	Cluster     types.ManagedObjectReference `xml:"cluster"`
	TargetAlias string                       `xml:"targetAlias"`
}

func init() {
	types.Add("vsan:VsanVitRemoveIscsiTargetRequestType", reflect.TypeOf((*VsanVitRemoveIscsiTargetRequestType)(nil)).Elem())
}

type VsanVitRemoveIscsiTargetResponse struct {
	Returnval *types.ManagedObjectReference `xml:"returnval,omitempty"`
}

type VsanVmVdsMigrationSpec struct {
	types.DynamicData

	VmInstanceUuid string                     `xml:"vmInstanceUuid"`
	Vnics          []VsanVnicVdsMigrationSpec `xml:"vnics"`
}

func init() {
	types.Add("vsan:VsanVmVdsMigrationSpec", reflect.TypeOf((*VsanVmVdsMigrationSpec)(nil)).Elem())
}

type VsanVmdkIOLoadSpec struct {
	types.DynamicData

	ReadPct      int32 `xml:"readPct"`
	Oio          int32 `xml:"oio"`
	IosizeB      int32 `xml:"iosizeB"`
	DataSizeMb   int64 `xml:"dataSizeMb"`
	Random       bool  `xml:"random"`
	StartOffsetB int64 `xml:"startOffsetB,omitempty"`
}

func init() {
	types.Add("vsan:VsanVmdkIOLoadSpec", reflect.TypeOf((*VsanVmdkIOLoadSpec)(nil)).Elem())
}

type VsanVmdkLoadTestResult struct {
	types.DynamicData

	Success                    bool                 `xml:"success"`
	FaultMessage               string               `xml:"faultMessage,omitempty"`
	Spec                       VsanVmdkLoadTestSpec `xml:"spec"`
	ActualDurationSec          int32                `xml:"actualDurationSec,omitempty"`
	TotalBytes                 int64                `xml:"totalBytes,omitempty"`
	Iops                       int64                `xml:"iops,omitempty"`
	TputBps                    int64                `xml:"tputBps,omitempty"`
	AvgLatencyUs               int64                `xml:"avgLatencyUs,omitempty"`
	MaxLatencyUs               int64                `xml:"maxLatencyUs,omitempty"`
	NumIoAboveLatencyThreshold int64                `xml:"numIoAboveLatencyThreshold,omitempty"`
}

func init() {
	types.Add("vsan:VsanVmdkLoadTestResult", reflect.TypeOf((*VsanVmdkLoadTestResult)(nil)).Elem())
}

type VsanVmdkLoadTestSpec struct {
	types.DynamicData

	VmdkCreateSpec     types.BaseFileBackedVirtualDiskSpec `xml:"vmdkCreateSpec,omitempty,typeattr"`
	VmdkIOSpec         *VsanVmdkIOLoadSpec                 `xml:"vmdkIOSpec,omitempty"`
	VmdkIOSpecSequence []VsanVmdkIOLoadSpec                `xml:"vmdkIOSpecSequence,omitempty"`
	StepDurationSec    int64                               `xml:"stepDurationSec,omitempty"`
}

func init() {
	types.Add("vsan:VsanVmdkLoadTestSpec", reflect.TypeOf((*VsanVmdkLoadTestSpec)(nil)).Elem())
}

type VsanVnicVdsMigrationSpec struct {
	types.DynamicData

	Key        int32                              `xml:"key"`
	VdsBacking types.BaseVirtualDeviceBackingInfo `xml:"vdsBacking,typeattr"`
}

func init() {
	types.Add("vsan:VsanVnicVdsMigrationSpec", reflect.TypeOf((*VsanVnicVdsMigrationSpec)(nil)).Elem())
}

type VsanVsanClusterPcapGroup struct {
	types.DynamicData

	Master  string   `xml:"master"`
	Members []string `xml:"members,omitempty"`
}

func init() {
	types.Add("vsan:VsanVsanClusterPcapGroup", reflect.TypeOf((*VsanVsanClusterPcapGroup)(nil)).Elem())
}

type VsanVsanClusterPcapResult struct {
	types.DynamicData

	Pkts        []string                   `xml:"pkts,omitempty"`
	Groups      []VsanVsanClusterPcapGroup `xml:"groups,omitempty"`
	Issues      []string                   `xml:"issues,omitempty"`
	HostResults []VsanVsanPcapResult       `xml:"hostResults,omitempty"`
}

func init() {
	types.Add("vsan:VsanVsanClusterPcapResult", reflect.TypeOf((*VsanVsanClusterPcapResult)(nil)).Elem())
}

type VsanVsanPcapResult struct {
	types.DynamicData

	Calltime      float32                     `xml:"calltime"`
	Vmknic        string                      `xml:"vmknic"`
	TcpdumpFilter string                      `xml:"tcpdumpFilter"`
	Snaplen       int32                       `xml:"snaplen"`
	Pkts          []string                    `xml:"pkts,omitempty"`
	Pcap          string                      `xml:"pcap,omitempty"`
	Error         *types.LocalizedMethodFault `xml:"error,omitempty"`
	Hostname      string                      `xml:"hostname,omitempty"`
}

func init() {
	types.Add("vsan:VsanVsanPcapResult", reflect.TypeOf((*VsanVsanPcapResult)(nil)).Elem())
}

type VsanVssMigrateVds VsanVssMigrateVdsRequestType

func init() {
	types.Add("vsan:VsanVssMigrateVds", reflect.TypeOf((*VsanVssMigrateVds)(nil)).Elem())
}

type VsanVssMigrateVdsRequestType struct {
	This         types.ManagedObjectReference   `xml:"_this"`
	Cluster      *types.ManagedObjectReference  `xml:"cluster,omitempty"`
	Hosts        []types.ManagedObjectReference `xml:"hosts,omitempty"`
	Vds          types.ManagedObjectReference   `xml:"vds"`
	VswitchName  string                         `xml:"vswitchName,omitempty"`
	VmnicDevices []string                       `xml:"vmnicDevices,omitempty"`
	InfraVm      []types.ManagedObjectReference `xml:"infraVm,omitempty"`
}

func init() {
	types.Add("vsan:VsanVssMigrateVdsRequestType", reflect.TypeOf((*VsanVssMigrateVdsRequestType)(nil)).Elem())
}

type VsanVssMigrateVdsResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanVumConfig struct {
	types.DynamicData

	BaselinePreferenceType string `xml:"baselinePreferenceType"`
}

func init() {
	types.Add("vsan:VsanVumConfig", reflect.TypeOf((*VsanVumConfig)(nil)).Elem())
}

type VsanVumSystemConfig struct {
	types.DynamicData

	Enabled                *bool      `xml:"enabled"`
	AutoCheckInterval      int32      `xml:"autoCheckInterval,omitempty"`
	MetadataUpdateInterval int32      `xml:"metadataUpdateInterval,omitempty"`
	ReleaseDbLastUpdate    *time.Time `xml:"releaseDbLastUpdate"`
}

func init() {
	types.Add("vsan:VsanVumSystemConfig", reflect.TypeOf((*VsanVumSystemConfig)(nil)).Elem())
}

type VsanWaitForVsanHealthGenerationIdChange VsanWaitForVsanHealthGenerationIdChangeRequestType

func init() {
	types.Add("vsan:VsanWaitForVsanHealthGenerationIdChange", reflect.TypeOf((*VsanWaitForVsanHealthGenerationIdChange)(nil)).Elem())
}

type VsanWaitForVsanHealthGenerationIdChangeRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Timeout int32                        `xml:"timeout"`
}

func init() {
	types.Add("vsan:VsanWaitForVsanHealthGenerationIdChangeRequestType", reflect.TypeOf((*VsanWaitForVsanHealthGenerationIdChangeRequestType)(nil)).Elem())
}

type VsanWaitForVsanHealthGenerationIdChangeResponse struct {
	Returnval bool `xml:"returnval"`
}

type VsanWhatIfEvacDetail struct {
	types.DynamicData

	Success                        *bool    `xml:"success"`
	BytesToSync                    int64    `xml:"bytesToSync,omitempty"`
	InaccessibleObjects            []string `xml:"inaccessibleObjects,omitempty"`
	IncompliantObjects             []string `xml:"incompliantObjects,omitempty"`
	ExtraSpaceNeeded               int64    `xml:"extraSpaceNeeded,omitempty"`
	FailedDueToInaccessibleObjects *bool    `xml:"failedDueToInaccessibleObjects"`
}

func init() {
	types.Add("vsan:VsanWhatIfEvacDetail", reflect.TypeOf((*VsanWhatIfEvacDetail)(nil)).Elem())
}

type VsanWhatIfEvacResult struct {
	types.DynamicData

	NoAction     VsanWhatIfEvacDetail `xml:"noAction"`
	EnsureAccess VsanWhatIfEvacDetail `xml:"ensureAccess"`
	EvacAllData  VsanWhatIfEvacDetail `xml:"evacAllData"`
}

func init() {
	types.Add("vsan:VsanWhatIfEvacResult", reflect.TypeOf((*VsanWhatIfEvacResult)(nil)).Elem())
}

type VsanWhatifCapacity struct {
	types.DynamicData

	TotalWhatifCapacityB int64                               `xml:"totalWhatifCapacityB"`
	FreeWhatifCapacityB  int64                               `xml:"freeWhatifCapacityB"`
	StoragePolicy        types.BaseVirtualMachineProfileSpec `xml:"storagePolicy,typeattr"`
	IsSatisfiable        bool                                `xml:"isSatisfiable"`
}

func init() {
	types.Add("vsan:VsanWhatifCapacity", reflect.TypeOf((*VsanWhatifCapacity)(nil)).Elem())
}

type VsanWitnessHostConfig struct {
	types.DynamicData

	SubClusterUuid           string `xml:"subClusterUuid"`
	PreferredFaultDomainName string `xml:"preferredFaultDomainName"`
	MetadataMode             *bool  `xml:"metadataMode"`
}

func init() {
	types.Add("vsan:VsanWitnessHostConfig", reflect.TypeOf((*VsanWitnessHostConfig)(nil)).Elem())
}

type AllowDataMovement bool

func init() {
	types.Add("vsan:allowDataMovement", reflect.TypeOf((*AllowDataMovement)(nil)).Elem())
}

type BurnInTest VsanBurnInTestCheckResult

func init() {
	types.Add("vsan:burnInTest", reflect.TypeOf((*BurnInTest)(nil)).Elem())
}

type CimProviderLinksOnHcl VsanDownloadItem

func init() {
	types.Add("vsan:cimProviderLinksOnHcl", reflect.TypeOf((*CimProviderLinksOnHcl)(nil)).Elem())
}

type Cluster types.ManagedObjectReference

func init() {
	types.Add("vsan:cluster", reflect.TypeOf((*Cluster)(nil)).Elem())
}

type ClusterInUnicastMode bool

func init() {
	types.Add("vsan:clusterInUnicastMode", reflect.TypeOf((*ClusterInUnicastMode)(nil)).Elem())
}

type DataEncryptionConfig VsanDataEncryptionConfig

func init() {
	types.Add("vsan:dataEncryptionConfig", reflect.TypeOf((*DataEncryptionConfig)(nil)).Elem())
}

type DatastoreConfig VsanDatastoreConfig

func init() {
	types.Add("vsan:datastoreConfig", reflect.TypeOf((*DatastoreConfig)(nil)).Elem())
}

type DedupMetadataSize int64

func init() {
	types.Add("vsan:dedupMetadataSize", reflect.TypeOf((*DedupMetadataSize)(nil)).Elem())
}

type DedupMetadataSizeB int64

func init() {
	types.Add("vsan:dedupMetadataSizeB", reflect.TypeOf((*DedupMetadataSizeB)(nil)).Elem())
}

type DekGenerationId int64

func init() {
	types.Add("vsan:dekGenerationId", reflect.TypeOf((*DekGenerationId)(nil)).Elem())
}

type DgTransientCapacityUsedB int64

func init() {
	types.Add("vsan:dgTransientCapacityUsedB", reflect.TypeOf((*DgTransientCapacityUsedB)(nil)).Elem())
}

type DiagnosticMode bool

func init() {
	types.Add("vsan:diagnosticMode", reflect.TypeOf((*DiagnosticMode)(nil)).Elem())
}

type DiskMode string

func init() {
	types.Add("vsan:diskMode", reflect.TypeOf((*DiskMode)(nil)).Elem())
}

type DiskModeOnHcl string

func init() {
	types.Add("vsan:diskModeOnHcl", reflect.TypeOf((*DiskModeOnHcl)(nil)).Elem())
}

type DiskModeSupported bool

func init() {
	types.Add("vsan:diskModeSupported", reflect.TypeOf((*DiskModeSupported)(nil)).Elem())
}

type DiskTransientCapacityUsedB int64

func init() {
	types.Add("vsan:diskTransientCapacityUsedB", reflect.TypeOf((*DiskTransientCapacityUsedB)(nil)).Elem())
}

type Disks VsanHclDiskInfo

func init() {
	types.Add("vsan:disks", reflect.TypeOf((*Disks)(nil)).Elem())
}

type DriversOnHcl VsanHclDriverInfo

func init() {
	types.Add("vsan:driversOnHcl", reflect.TypeOf((*DriversOnHcl)(nil)).Elem())
}

type Duration int64

func init() {
	types.Add("vsan:duration", reflect.TypeOf((*Duration)(nil)).Elem())
}

type EfficientCapacity VimVsanDataEfficiencyCapacityState

func init() {
	types.Add("vsan:efficientCapacity", reflect.TypeOf((*EfficientCapacity)(nil)).Elem())
}

type EncryptedUnlocked bool

func init() {
	types.Add("vsan:encryptedUnlocked", reflect.TypeOf((*EncryptedUnlocked)(nil)).Elem())
}

type EncryptionEnabled bool

func init() {
	types.Add("vsan:encryptionEnabled", reflect.TypeOf((*EncryptionEnabled)(nil)).Elem())
}

type EncryptionHealth VsanClusterEncryptionHealthSummary

func init() {
	types.Add("vsan:encryptionHealth", reflect.TypeOf((*EncryptionHealth)(nil)).Elem())
}

type EncryptionInfo VsanDataEncryptionConfig

func init() {
	types.Add("vsan:encryptionInfo", reflect.TypeOf((*EncryptionInfo)(nil)).Elem())
}

type ExtendedConfig VsanExtendedConfig

func init() {
	types.Add("vsan:extendedConfig", reflect.TypeOf((*ExtendedConfig)(nil)).Elem())
}

type FileServiceConfig VsanFileServiceConfig

func init() {
	types.Add("vsan:fileServiceConfig", reflect.TypeOf((*FileServiceConfig)(nil)).Elem())
}

type FileServiceHealth VsanClusterFileServiceHealthSummary

func init() {
	types.Add("vsan:fileServiceHealth", reflect.TypeOf((*FileServiceHealth)(nil)).Elem())
}

type FwAuxVersion string

func init() {
	types.Add("vsan:fwAuxVersion", reflect.TypeOf((*FwAuxVersion)(nil)).Elem())
}

type GenericCluster VsanGenericClusterBestPracticeHealth

func init() {
	types.Add("vsan:genericCluster", reflect.TypeOf((*GenericCluster)(nil)).Elem())
}

type HostLatencyCheckSuccess bool

func init() {
	types.Add("vsan:hostLatencyCheckSuccess", reflect.TypeOf((*HostLatencyCheckSuccess)(nil)).Elem())
}

type Hostname string

func init() {
	types.Add("vsan:hostname", reflect.TypeOf((*Hostname)(nil)).Elem())
}

type InUnicast bool

func init() {
	types.Add("vsan:inUnicast", reflect.TypeOf((*InUnicast)(nil)).Elem())
}

type IsDataMovementRequired bool

func init() {
	types.Add("vsan:isDataMovementRequired", reflect.TypeOf((*IsDataMovementRequired)(nil)).Elem())
}

type IsDefault bool

func init() {
	types.Add("vsan:isDefault", reflect.TypeOf((*IsDefault)(nil)).Elem())
}

type IsSupportUnicast bool

func init() {
	types.Add("vsan:isSupportUnicast", reflect.TypeOf((*IsSupportUnicast)(nil)).Elem())
}

type Issues types.LocalizedMethodFault

func init() {
	types.Add("vsan:issues", reflect.TypeOf((*Issues)(nil)).Elem())
}

type KekId string

func init() {
	types.Add("vsan:kekId", reflect.TypeOf((*KekId)(nil)).Elem())
}

type KmsProviderId string

func init() {
	types.Add("vsan:kmsProviderId", reflect.TypeOf((*KmsProviderId)(nil)).Elem())
}

type LogicalSpaceUsedB int64

func init() {
	types.Add("vsan:logicalSpaceUsedB", reflect.TypeOf((*LogicalSpaceUsedB)(nil)).Elem())
}

type MemberUuid string

func init() {
	types.Add("vsan:memberUuid", reflect.TypeOf((*MemberUuid)(nil)).Elem())
}

type MetadataMode bool

func init() {
	types.Add("vsan:metadataMode", reflect.TypeOf((*MetadataMode)(nil)).Elem())
}

type MetricsConfig VsanMetricsConfig

func init() {
	types.Add("vsan:metricsConfig", reflect.TypeOf((*MetricsConfig)(nil)).Elem())
}

type NetworkConfig VsanNetworkConfigBestPracticeHealth

func init() {
	types.Add("vsan:networkConfig", reflect.TypeOf((*NetworkConfig)(nil)).Elem())
}

type NumExceptions string

func init() {
	types.Add("vsan:numExceptions", reflect.TypeOf((*NumExceptions)(nil)).Elem())
}

type ObjPolicyGenerationId string

func init() {
	types.Add("vsan:objPolicyGenerationId", reflect.TypeOf((*ObjPolicyGenerationId)(nil)).Elem())
}

type ObjectsComplianceDetail VsanStorageComplianceResult

func init() {
	types.Add("vsan:objectsComplianceDetail", reflect.TypeOf((*ObjectsComplianceDetail)(nil)).Elem())
}

type PartitionUnknown bool

func init() {
	types.Add("vsan:partitionUnknown", reflect.TypeOf((*PartitionUnknown)(nil)).Elem())
}

type PerfsvcConfig VsanPerfsvcConfig

func init() {
	types.Add("vsan:perfsvcConfig", reflect.TypeOf((*PerfsvcConfig)(nil)).Elem())
}

type PerfsvcHealth VsanPerfsvcHealthResult

func init() {
	types.Add("vsan:perfsvcHealth", reflect.TypeOf((*PerfsvcHealth)(nil)).Elem())
}

type Pnics VsanHclNicInfo

func init() {
	types.Add("vsan:pnics", reflect.TypeOf((*Pnics)(nil)).Elem())
}

type QueueDepth int64

func init() {
	types.Add("vsan:queueDepth", reflect.TypeOf((*QueueDepth)(nil)).Elem())
}

type QueueDepthOnHcl int64

func init() {
	types.Add("vsan:queueDepthOnHcl", reflect.TypeOf((*QueueDepthOnHcl)(nil)).Elem())
}

type QueueDepthSupported bool

func init() {
	types.Add("vsan:queueDepthSupported", reflect.TypeOf((*QueueDepthSupported)(nil)).Elem())
}

type RebalanceResult VsanDiskRebalanceResult

func init() {
	types.Add("vsan:rebalanceResult", reflect.TypeOf((*RebalanceResult)(nil)).Elem())
}

type RemediableIssues string

func init() {
	types.Add("vsan:remediableIssues", reflect.TypeOf((*RemediableIssues)(nil)).Elem())
}

type ResyncIopsLimitConfig ResyncIopsInfo

func init() {
	types.Add("vsan:resyncIopsLimitConfig", reflect.TypeOf((*ResyncIopsLimitConfig)(nil)).Elem())
}

type SkipHostRemediation bool

func init() {
	types.Add("vsan:skipHostRemediation", reflect.TypeOf((*SkipHostRemediation)(nil)).Elem())
}

type Statuses string

func init() {
	types.Add("vsan:statuses", reflect.TypeOf((*Statuses)(nil)).Elem())
}

type TestAllEntities int32

func init() {
	types.Add("vsan:testAllEntities", reflect.TypeOf((*TestAllEntities)(nil)).Elem())
}

type TestHealthyEntities int32

func init() {
	types.Add("vsan:testHealthyEntities", reflect.TypeOf((*TestHealthyEntities)(nil)).Elem())
}

type Threshold VsanPerfThreshold

func init() {
	types.Add("vsan:threshold", reflect.TypeOf((*Threshold)(nil)).Elem())
}

type ToBeDeleted bool

func init() {
	types.Add("vsan:toBeDeleted", reflect.TypeOf((*ToBeDeleted)(nil)).Elem())
}

type ToolName string

func init() {
	types.Add("vsan:toolName", reflect.TypeOf((*ToolName)(nil)).Elem())
}

type ToolVersion string

func init() {
	types.Add("vsan:toolVersion", reflect.TypeOf((*ToolVersion)(nil)).Elem())
}

type TotalLogicalSpaceB int64

func init() {
	types.Add("vsan:totalLogicalSpaceB", reflect.TypeOf((*TotalLogicalSpaceB)(nil)).Elem())
}

type UncommittedB int64

func init() {
	types.Add("vsan:uncommittedB", reflect.TypeOf((*UncommittedB)(nil)).Elem())
}

type UnicastAddressInfos VsanUnicastAddressInfo

func init() {
	types.Add("vsan:unicastAddressInfos", reflect.TypeOf((*UnicastAddressInfos)(nil)).Elem())
}

type UnicastConfig string

func init() {
	types.Add("vsan:unicastConfig", reflect.TypeOf((*UnicastConfig)(nil)).Elem())
}

type UnlockedEncrypted bool

func init() {
	types.Add("vsan:unlockedEncrypted", reflect.TypeOf((*UnlockedEncrypted)(nil)).Elem())
}

type UnmapConfig VsanUnmapConfig

func init() {
	types.Add("vsan:unmapConfig", reflect.TypeOf((*UnmapConfig)(nil)).Elem())
}

type UpdateItems VsanUpdateItem

func init() {
	types.Add("vsan:updateItems", reflect.TypeOf((*UpdateItems)(nil)).Elem())
}

type UpgradePossible bool

func init() {
	types.Add("vsan:upgradePossible", reflect.TypeOf((*UpgradePossible)(nil)).Elem())
}

type UsedByVsan bool

func init() {
	types.Add("vsan:usedByVsan", reflect.TypeOf((*UsedByVsan)(nil)).Elem())
}

type VMotionHealth VsanNetworkPeerHealthResult

func init() {
	types.Add("vsan:vMotionHealth", reflect.TypeOf((*VMotionHealth)(nil)).Elem())
}

type VerboseMode bool

func init() {
	types.Add("vsan:verboseMode", reflect.TypeOf((*VerboseMode)(nil)).Elem())
}

type VerboseModeLastUpdate time.Time

func init() {
	types.Add("vsan:verboseModeLastUpdate", reflect.TypeOf((*VerboseModeLastUpdate)(nil)).Elem())
}

type VsanConfig VsanConfigCheckResult

func init() {
	types.Add("vsan:vsanConfig", reflect.TypeOf((*VsanConfig)(nil)).Elem())
}

/*
type VsanDataEncryptionConfig VsanHostEncryptionInfo

func init() {
	types.Add("vsan:vsanDataEncryptionConfig", reflect.TypeOf((*VsanDataEncryptionConfig)(nil)).Elem())
}
*/

type VumConfig VsanVumConfig

func init() {
	types.Add("vsan:vumConfig", reflect.TypeOf((*VumConfig)(nil)).Elem())
}

type WhatifCapacities VsanWhatifCapacity

func init() {
	types.Add("vsan:whatifCapacities", reflect.TypeOf((*WhatifCapacities)(nil)).Elem())
}
