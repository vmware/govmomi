// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"reflect"
	"time"

	"github.com/vmware/govmomi/vim25/types"
)

type VsanPerfDiagnose VsanPerfDiagnoseRequestType

func init() {
	types.Add("vsan:VsanPerfDiagnose", reflect.TypeOf((*VsanPerfDiagnose)(nil)).Elem())
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
	Returnval []types.BaseDynamicData `xml:"returnval,omitempty,typeattr"`
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

type QueryRemoteServerClusters QueryRemoteServerClustersRequestType

func init() {
	types.Add("vsan:QueryRemoteServerClusters", reflect.TypeOf((*QueryRemoteServerClusters)(nil)).Elem())
}

type QueryRemoteServerClustersRequestType struct {
	This      types.ManagedObjectReference  `xml:"_this"`
	Cluster   *types.ManagedObjectReference `xml:"cluster,omitempty"`
	QuerySpec *VsanRemoteClusterQuerySpec   `xml:"querySpec,omitempty"`
}

func init() {
	types.Add("vsan:QueryRemoteServerClustersRequestType", reflect.TypeOf((*QueryRemoteServerClustersRequestType)(nil)).Elem())
}

type QueryRemoteServerClustersResponse struct {
	Returnval []string `xml:"returnval,omitempty"`
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

type VsanPerfSetStatsObjectPolicy VsanPerfSetStatsObjectPolicyRequestType

func init() {
	types.Add("vsan:VsanPerfSetStatsObjectPolicy", reflect.TypeOf((*VsanPerfSetStatsObjectPolicy)(nil)).Elem())
}

type VsanPerfSetStatsObjectPolicyRequestType struct {
	This    types.ManagedObjectReference     `xml:"_this"`
	Cluster *types.ManagedObjectReference    `xml:"cluster,omitempty"`
	Profile *types.VirtualMachineProfileSpec `xml:"profile,omitempty"`
}

func init() {
	types.Add("vsan:VsanPerfSetStatsObjectPolicyRequestType", reflect.TypeOf((*VsanPerfSetStatsObjectPolicyRequestType)(nil)).Elem())
}

type VsanPerfSetStatsObjectPolicyResponse struct {
	Returnval bool `xml:"returnval"`
}

type VsanPerfCreateStatsObject VsanPerfCreateStatsObjectRequestType

func init() {
	types.Add("vsan:VsanPerfCreateStatsObject", reflect.TypeOf((*VsanPerfCreateStatsObject)(nil)).Elem())
}

type VsanPerfCreateStatsObjectRequestType struct {
	This    types.ManagedObjectReference     `xml:"_this"`
	Cluster *types.ManagedObjectReference    `xml:"cluster,omitempty"`
	Profile *types.VirtualMachineProfileSpec `xml:"profile,omitempty"`
}

func init() {
	types.Add("vsan:VsanPerfCreateStatsObjectRequestType", reflect.TypeOf((*VsanPerfCreateStatsObjectRequestType)(nil)).Elem())
}

type VsanPerfCreateStatsObjectResponse struct {
	Returnval string `xml:"returnval"`
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
	Returnval []VsanPerfEntityMetricCSV `xml:"returnval,omitempty"`
}

type VsanPerfCreateStatsObjectTask VsanPerfCreateStatsObjectTaskRequestType

func init() {
	types.Add("vsan:VsanPerfCreateStatsObjectTask", reflect.TypeOf((*VsanPerfCreateStatsObjectTask)(nil)).Elem())
}

type VsanPerfCreateStatsObjectTaskRequestType struct {
	This    types.ManagedObjectReference     `xml:"_this"`
	Cluster *types.ManagedObjectReference    `xml:"cluster,omitempty"`
	Profile *types.VirtualMachineProfileSpec `xml:"profile,omitempty"`
}

func init() {
	types.Add("vsan:VsanPerfCreateStatsObjectTaskRequestType", reflect.TypeOf((*VsanPerfCreateStatsObjectTaskRequestType)(nil)).Elem())
}

type VsanPerfCreateStatsObjectTaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
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

type UnmountDiskMappingEx UnmountDiskMappingExRequestType

func init() {
	types.Add("vsan:UnmountDiskMappingEx", reflect.TypeOf((*UnmountDiskMappingEx)(nil)).Elem())
}

type UnmountDiskMappingExRequestType struct {
	This            types.ManagedObjectReference `xml:"_this"`
	Cluster         types.ManagedObjectReference `xml:"cluster"`
	Mappings        []types.VsanHostDiskMapping  `xml:"mappings"`
	MaintenanceSpec types.HostMaintenanceSpec    `xml:"maintenanceSpec"`
}

func init() {
	types.Add("vsan:UnmountDiskMappingExRequestType", reflect.TypeOf((*UnmountDiskMappingExRequestType)(nil)).Elem())
}

type UnmountDiskMappingExResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
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

type RemoveDiskEx RemoveDiskExRequestType

func init() {
	types.Add("vsan:RemoveDiskEx", reflect.TypeOf((*RemoveDiskEx)(nil)).Elem())
}

type RemoveDiskExRequestType struct {
	This            types.ManagedObjectReference `xml:"_this"`
	Cluster         types.ManagedObjectReference `xml:"cluster"`
	Disks           []types.HostScsiDisk         `xml:"disks"`
	MaintenanceSpec types.HostMaintenanceSpec    `xml:"maintenanceSpec"`
}

func init() {
	types.Add("vsan:RemoveDiskExRequestType", reflect.TypeOf((*RemoveDiskExRequestType)(nil)).Elem())
}

type RemoveDiskExResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
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

type QueryVsanManagedDisks QueryVsanManagedDisksRequestType

func init() {
	types.Add("vsan:QueryVsanManagedDisks", reflect.TypeOf((*QueryVsanManagedDisks)(nil)).Elem())
}

type QueryVsanManagedDisksRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
	Host types.ManagedObjectReference `xml:"host"`
}

func init() {
	types.Add("vsan:QueryVsanManagedDisksRequestType", reflect.TypeOf((*QueryVsanManagedDisksRequestType)(nil)).Elem())
}

type QueryVsanManagedDisksResponse struct {
	Returnval *VimVsanHostVsanManagedDisksInfo `xml:"returnval,omitempty"`
}

type RebuildDiskMapping RebuildDiskMappingRequestType

func init() {
	types.Add("vsan:RebuildDiskMapping", reflect.TypeOf((*RebuildDiskMapping)(nil)).Elem())
}

type RebuildDiskMappingRequestType struct {
	This            types.ManagedObjectReference `xml:"_this"`
	Host            types.ManagedObjectReference `xml:"host"`
	Mapping         types.VsanHostDiskMapping    `xml:"mapping"`
	MaintenanceSpec types.HostMaintenanceSpec    `xml:"maintenanceSpec"`
}

func init() {
	types.Add("vsan:RebuildDiskMappingRequestType", reflect.TypeOf((*RebuildDiskMappingRequestType)(nil)).Elem())
}

type RebuildDiskMappingResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type RemoveDiskMappingEx RemoveDiskMappingExRequestType

func init() {
	types.Add("vsan:RemoveDiskMappingEx", reflect.TypeOf((*RemoveDiskMappingEx)(nil)).Elem())
}

type RemoveDiskMappingExRequestType struct {
	This            types.ManagedObjectReference `xml:"_this"`
	Cluster         types.ManagedObjectReference `xml:"cluster"`
	Mappings        []types.VsanHostDiskMapping  `xml:"mappings"`
	MaintenanceSpec types.HostMaintenanceSpec    `xml:"maintenanceSpec"`
}

func init() {
	types.Add("vsan:RemoveDiskMappingExRequestType", reflect.TypeOf((*RemoveDiskMappingExRequestType)(nil)).Elem())
}

type RemoveDiskMappingExResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
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

type VsanVitRemoveIscsiTarget VsanVitRemoveIscsiTargetRequestType

func init() {
	types.Add("vsan:VsanVitRemoveIscsiTarget", reflect.TypeOf((*VsanVitRemoveIscsiTarget)(nil)).Elem())
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
	Returnval []VsanCapability `xml:"returnval,omitempty"`
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

type VosSetVsanObjectPolicy VosSetVsanObjectPolicyRequestType

func init() {
	types.Add("vsan:VosSetVsanObjectPolicy", reflect.TypeOf((*VosSetVsanObjectPolicy)(nil)).Elem())
}

type VosSetVsanObjectPolicyRequestType struct {
	This           types.ManagedObjectReference     `xml:"_this"`
	Cluster        *types.ManagedObjectReference    `xml:"cluster,omitempty"`
	VsanObjectUuid string                           `xml:"vsanObjectUuid"`
	Profile        *types.VirtualMachineProfileSpec `xml:"profile,omitempty"`
}

func init() {
	types.Add("vsan:VosSetVsanObjectPolicyRequestType", reflect.TypeOf((*VosSetVsanObjectPolicyRequestType)(nil)).Elem())
}

type VosSetVsanObjectPolicyResponse struct {
	Returnval bool `xml:"returnval"`
}

type VsanDeleteObjects_Task VsanDeleteObjects_TaskRequestType

func init() {
	types.Add("vsan:VsanDeleteObjects_Task", reflect.TypeOf((*VsanDeleteObjects_Task)(nil)).Elem())
}

type VsanDeleteObjects_TaskRequestType struct {
	This     types.ManagedObjectReference  `xml:"_this"`
	Cluster  *types.ManagedObjectReference `xml:"cluster,omitempty"`
	ObjUuids []string                      `xml:"objUuids"`
	Force    *bool                         `xml:"force"`
}

func init() {
	types.Add("vsan:VsanDeleteObjects_TaskRequestType", reflect.TypeOf((*VsanDeleteObjects_TaskRequestType)(nil)).Elem())
}

type VsanDeleteObjects_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
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

type RelayoutObjects RelayoutObjectsRequestType

func init() {
	types.Add("vsan:RelayoutObjects", reflect.TypeOf((*RelayoutObjects)(nil)).Elem())
}

type RelayoutObjectsRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

func init() {
	types.Add("vsan:RelayoutObjectsRequestType", reflect.TypeOf((*RelayoutObjectsRequestType)(nil)).Elem())
}

type RelayoutObjectsResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
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

type VsanValidateConfigSpec VsanValidateConfigSpecRequestType

func init() {
	types.Add("vsan:VsanValidateConfigSpec", reflect.TypeOf((*VsanValidateConfigSpec)(nil)).Elem())
}

type VsanValidateConfigSpecRequestType struct {
	This             types.ManagedObjectReference `xml:"_this"`
	Cluster          types.ManagedObjectReference `xml:"cluster"`
	VsanReconfigSpec VimVsanReconfigSpec          `xml:"vsanReconfigSpec"`
}

func init() {
	types.Add("vsan:VsanValidateConfigSpecRequestType", reflect.TypeOf((*VsanValidateConfigSpecRequestType)(nil)).Elem())
}

type VsanValidateConfigSpecResponse struct {
	Returnval []types.ClusterComputeResourceValidationResultBase `xml:"returnval,omitempty"`
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

type VsanEncryptedClusterRekey_Task VsanEncryptedClusterRekey_TaskRequestType

func init() {
	types.Add("vsan:VsanEncryptedClusterRekey_Task", reflect.TypeOf((*VsanEncryptedClusterRekey_Task)(nil)).Elem())
}

type VsanEncryptedClusterRekey_TaskRequestType struct {
	This                   types.ManagedObjectReference `xml:"_this"`
	EncryptedCluster       types.ManagedObjectReference `xml:"encryptedCluster"`
	DeepRekey              *bool                        `xml:"deepRekey"`
	AllowReducedRedundancy *bool                        `xml:"allowReducedRedundancy"`
}

func init() {
	types.Add("vsan:VsanEncryptedClusterRekey_TaskRequestType", reflect.TypeOf((*VsanEncryptedClusterRekey_TaskRequestType)(nil)).Elem())
}

type VsanEncryptedClusterRekey_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
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

type VsanQueryVcClusterCreateVmHealthHistoryTest VsanQueryVcClusterCreateVmHealthHistoryTestRequestType

func init() {
	types.Add("vsan:VsanQueryVcClusterCreateVmHealthHistoryTest", reflect.TypeOf((*VsanQueryVcClusterCreateVmHealthHistoryTest)(nil)).Elem())
}

type VsanQueryVcClusterCreateVmHealthHistoryTestRequestType struct {
	This      types.ManagedObjectReference  `xml:"_this"`
	Cluster   types.ManagedObjectReference  `xml:"cluster"`
	Count     int32                         `xml:"count,omitempty"`
	Datastore *types.ManagedObjectReference `xml:"datastore,omitempty"`
}

func init() {
	types.Add("vsan:VsanQueryVcClusterCreateVmHealthHistoryTestRequestType", reflect.TypeOf((*VsanQueryVcClusterCreateVmHealthHistoryTestRequestType)(nil)).Elem())
}

type VsanQueryVcClusterCreateVmHealthHistoryTestResponse struct {
	Returnval []VsanClusterCreateVmHealthTestResult `xml:"returnval,omitempty"`
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
	Returnval []VsanSmartStatsHostSummary `xml:"returnval,omitempty"`
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

type VsanQueryVcClusterCreateVmHealthTest VsanQueryVcClusterCreateVmHealthTestRequestType

func init() {
	types.Add("vsan:VsanQueryVcClusterCreateVmHealthTest", reflect.TypeOf((*VsanQueryVcClusterCreateVmHealthTest)(nil)).Elem())
}

type VsanQueryVcClusterCreateVmHealthTestRequestType struct {
	This      types.ManagedObjectReference  `xml:"_this"`
	Cluster   types.ManagedObjectReference  `xml:"cluster"`
	Timeout   int32                         `xml:"timeout"`
	Datastore *types.ManagedObjectReference `xml:"datastore,omitempty"`
}

func init() {
	types.Add("vsan:VsanQueryVcClusterCreateVmHealthTestRequestType", reflect.TypeOf((*VsanQueryVcClusterCreateVmHealthTestRequestType)(nil)).Elem())
}

type VsanQueryVcClusterCreateVmHealthTestResponse struct {
	Returnval VsanClusterCreateVmHealthTestResult `xml:"returnval"`
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
	Returnval []VsanClusterHealthCheckInfo `xml:"returnval,omitempty"`
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

type VsanQueryVcClusterNetworkPerfHistoryTest VsanQueryVcClusterNetworkPerfHistoryTestRequestType

func init() {
	types.Add("vsan:VsanQueryVcClusterNetworkPerfHistoryTest", reflect.TypeOf((*VsanQueryVcClusterNetworkPerfHistoryTest)(nil)).Elem())
}

type VsanQueryVcClusterNetworkPerfHistoryTestRequestType struct {
	This    types.ManagedObjectReference    `xml:"_this"`
	Cluster types.ManagedObjectReference    `xml:"cluster"`
	Count   int32                           `xml:"count,omitempty"`
	Spec    *VsanClusterNetworkPerfTaskSpec `xml:"spec,omitempty"`
}

func init() {
	types.Add("vsan:VsanQueryVcClusterNetworkPerfHistoryTestRequestType", reflect.TypeOf((*VsanQueryVcClusterNetworkPerfHistoryTestRequestType)(nil)).Elem())
}

type VsanQueryVcClusterNetworkPerfHistoryTestResponse struct {
	Returnval []VsanClusterNetworkLoadTestResult `xml:"returnval,omitempty"`
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
	Spec            *VsanClusterHealthQuerySpec    `xml:"spec,omitempty"`
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
	IncludeOnlineHealth         *bool                          `xml:"includeOnlineHealth"`
}

func init() {
	types.Add("vsan:VsanQueryVcClusterHealthSummaryTaskRequestType", reflect.TypeOf((*VsanQueryVcClusterHealthSummaryTaskRequestType)(nil)).Elem())
}

type VsanQueryVcClusterHealthSummaryTaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanQueryVcClusterNetworkPerfTask VsanQueryVcClusterNetworkPerfTaskRequestType

func init() {
	types.Add("vsan:VsanQueryVcClusterNetworkPerfTask", reflect.TypeOf((*VsanQueryVcClusterNetworkPerfTask)(nil)).Elem())
}

type VsanQueryVcClusterNetworkPerfTaskRequestType struct {
	This    types.ManagedObjectReference    `xml:"_this"`
	Cluster types.ManagedObjectReference    `xml:"cluster"`
	Spec    *VsanClusterNetworkPerfTaskSpec `xml:"spec,omitempty"`
}

func init() {
	types.Add("vsan:VsanQueryVcClusterNetworkPerfTaskRequestType", reflect.TypeOf((*VsanQueryVcClusterNetworkPerfTaskRequestType)(nil)).Elem())
}

type VsanQueryVcClusterNetworkPerfTaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
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

type VsanDownloadHclFile_Task VsanDownloadHclFile_TaskRequestType

func init() {
	types.Add("vsan:VsanDownloadHclFile_Task", reflect.TypeOf((*VsanDownloadHclFile_Task)(nil)).Elem())
}

type VsanDownloadHclFile_TaskRequestType struct {
	This     types.ManagedObjectReference `xml:"_this"`
	Sha1sums []string                     `xml:"sha1sums"`
}

func init() {
	types.Add("vsan:VsanDownloadHclFile_TaskRequestType", reflect.TypeOf((*VsanDownloadHclFile_TaskRequestType)(nil)).Elem())
}

type VsanDownloadHclFile_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
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
	Returnval []VsanStorageWorkloadType `xml:"returnval,omitempty"`
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

type VsanDownloadAndInstallVendorTool_Task VsanDownloadAndInstallVendorTool_TaskRequestType

func init() {
	types.Add("vsan:VsanDownloadAndInstallVendorTool_Task", reflect.TypeOf((*VsanDownloadAndInstallVendorTool_Task)(nil)).Elem())
}

type VsanDownloadAndInstallVendorTool_TaskRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

func init() {
	types.Add("vsan:VsanDownloadAndInstallVendorTool_TaskRequestType", reflect.TypeOf((*VsanDownloadAndInstallVendorTool_TaskRequestType)(nil)).Elem())
}

type VsanDownloadAndInstallVendorTool_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
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
	Returnval string `xml:"returnval"`
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
	Returnval []VsanVcsaDeploymentProgress `xml:"returnval,omitempty"`
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
	Returnval string `xml:"returnval"`
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

type MountPrecheck MountPrecheckRequestType

func init() {
	types.Add("vsan:MountPrecheck", reflect.TypeOf((*MountPrecheck)(nil)).Elem())
}

type MountPrecheckRequestType struct {
	This      types.ManagedObjectReference `xml:"_this"`
	Cluster   types.ManagedObjectReference `xml:"cluster"`
	Datastore types.ManagedObjectReference `xml:"datastore"`
}

func init() {
	types.Add("vsan:MountPrecheckRequestType", reflect.TypeOf((*MountPrecheckRequestType)(nil)).Elem())
}

type MountPrecheckResponse struct {
	Returnval VsanMountPrecheckResult `xml:"returnval"`
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

type PerformVsanUpgradePreflightAsyncCheck_Task PerformVsanUpgradePreflightAsyncCheck_TaskRequestType

func init() {
	types.Add("vsan:PerformVsanUpgradePreflightAsyncCheck_Task", reflect.TypeOf((*PerformVsanUpgradePreflightAsyncCheck_Task)(nil)).Elem())
}

type PerformVsanUpgradePreflightAsyncCheck_TaskRequestType struct {
	This            types.ManagedObjectReference  `xml:"_this"`
	Cluster         types.ManagedObjectReference  `xml:"cluster"`
	DowngradeFormat *bool                         `xml:"downgradeFormat"`
	Spec            *VsanDiskFormatConversionSpec `xml:"spec,omitempty"`
}

func init() {
	types.Add("vsan:PerformVsanUpgradePreflightAsyncCheck_TaskRequestType", reflect.TypeOf((*PerformVsanUpgradePreflightAsyncCheck_TaskRequestType)(nil)).Elem())
}

type PerformVsanUpgradePreflightAsyncCheck_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanQuerySpaceUsage VsanQuerySpaceUsageRequestType

func init() {
	types.Add("vsan:VsanQuerySpaceUsage", reflect.TypeOf((*VsanQuerySpaceUsage)(nil)).Elem())
}

type VsanQuerySpaceUsageRequestType struct {
	This               types.ManagedObjectReference      `xml:"_this"`
	Cluster            types.ManagedObjectReference      `xml:"cluster"`
	StoragePolicies    []types.VirtualMachineProfileSpec `xml:"storagePolicies,omitempty"`
	WhatifCapacityOnly *bool                             `xml:"whatifCapacityOnly"`
}

func init() {
	types.Add("vsan:VsanQuerySpaceUsageRequestType", reflect.TypeOf((*VsanQuerySpaceUsageRequestType)(nil)).Elem())
}

type VsanQuerySpaceUsageResponse struct {
	Returnval VsanSpaceUsage `xml:"returnval"`
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

type QueryVsanManagedStorageSpaceUsage QueryVsanManagedStorageSpaceUsageRequestType

func init() {
	types.Add("vsan:QueryVsanManagedStorageSpaceUsage", reflect.TypeOf((*QueryVsanManagedStorageSpaceUsage)(nil)).Elem())
}

type QueryVsanManagedStorageSpaceUsageRequestType struct {
	This      types.ManagedObjectReference          `xml:"_this"`
	Cluster   types.ManagedObjectReference          `xml:"cluster"`
	QuerySpec QueryVsanManagedStorageSpaceUsageSpec `xml:"querySpec"`
}

func init() {
	types.Add("vsan:QueryVsanManagedStorageSpaceUsageRequestType", reflect.TypeOf((*QueryVsanManagedStorageSpaceUsageRequestType)(nil)).Elem())
}

type QueryVsanManagedStorageSpaceUsageResponse struct {
	Returnval []VsanSpaceUsageWithDatastoreType `xml:"returnval,omitempty"`
}

type StartIoInsight StartIoInsightRequestType

func init() {
	types.Add("vsan:StartIoInsight", reflect.TypeOf((*StartIoInsight)(nil)).Elem())
}

type StartIoInsightRequestType struct {
	This        types.ManagedObjectReference   `xml:"_this"`
	Cluster     *types.ManagedObjectReference  `xml:"cluster,omitempty"`
	RunName     string                         `xml:"runName,omitempty"`
	DurationSec int64                          `xml:"durationSec,omitempty"`
	TargetHosts []types.ManagedObjectReference `xml:"targetHosts,omitempty"`
	TargetVMs   []types.ManagedObjectReference `xml:"targetVMs,omitempty"`
}

func init() {
	types.Add("vsan:StartIoInsightRequestType", reflect.TypeOf((*StartIoInsightRequestType)(nil)).Elem())
}

type StartIoInsightResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type QueryIoInsightInstances QueryIoInsightInstancesRequestType

func init() {
	types.Add("vsan:QueryIoInsightInstances", reflect.TypeOf((*QueryIoInsightInstances)(nil)).Elem())
}

type QueryIoInsightInstancesRequestType struct {
	This      types.ManagedObjectReference   `xml:"_this"`
	QuerySpec VsanIoInsightInstanceQuerySpec `xml:"querySpec"`
	Cluster   *types.ManagedObjectReference  `xml:"cluster,omitempty"`
}

func init() {
	types.Add("vsan:QueryIoInsightInstancesRequestType", reflect.TypeOf((*QueryIoInsightInstancesRequestType)(nil)).Elem())
}

type QueryIoInsightInstancesResponse struct {
	Returnval []VsanIoInsightInstance `xml:"returnval,omitempty"`
}

type RenameIoInsightInstance RenameIoInsightInstanceRequestType

func init() {
	types.Add("vsan:RenameIoInsightInstance", reflect.TypeOf((*RenameIoInsightInstance)(nil)).Elem())
}

type RenameIoInsightInstanceRequestType struct {
	This       types.ManagedObjectReference  `xml:"_this"`
	OldRunName string                        `xml:"oldRunName"`
	NewRunName string                        `xml:"newRunName"`
	Cluster    *types.ManagedObjectReference `xml:"cluster,omitempty"`
}

func init() {
	types.Add("vsan:RenameIoInsightInstanceRequestType", reflect.TypeOf((*RenameIoInsightInstanceRequestType)(nil)).Elem())
}

type RenameIoInsightInstanceResponse struct {
}

type StopIoInsight StopIoInsightRequestType

func init() {
	types.Add("vsan:StopIoInsight", reflect.TypeOf((*StopIoInsight)(nil)).Elem())
}

type StopIoInsightRequestType struct {
	This                types.ManagedObjectReference  `xml:"_this"`
	Cluster             *types.ManagedObjectReference `xml:"cluster,omitempty"`
	RunName             string                        `xml:"runName,omitempty"`
	HostsIoInsightInfos []VsanHostIoInsightInfo       `xml:"hostsIoInsightInfos,omitempty"`
}

func init() {
	types.Add("vsan:StopIoInsightRequestType", reflect.TypeOf((*StopIoInsightRequestType)(nil)).Elem())
}

type StopIoInsightResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type DeleteIoInsightInstance DeleteIoInsightInstanceRequestType

func init() {
	types.Add("vsan:DeleteIoInsightInstance", reflect.TypeOf((*DeleteIoInsightInstance)(nil)).Elem())
}

type DeleteIoInsightInstanceRequestType struct {
	This    types.ManagedObjectReference  `xml:"_this"`
	RunName string                        `xml:"runName"`
	Cluster *types.ManagedObjectReference `xml:"cluster,omitempty"`
}

func init() {
	types.Add("vsan:DeleteIoInsightInstanceRequestType", reflect.TypeOf((*DeleteIoInsightInstanceRequestType)(nil)).Elem())
}

type DeleteIoInsightInstanceResponse struct {
}

type VsanVibInstall_Task VsanVibInstall_TaskRequestType

func init() {
	types.Add("vsan:VsanVibInstall_Task", reflect.TypeOf((*VsanVibInstall_Task)(nil)).Elem())
}

type VsanVibInstall_TaskRequestType struct {
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
	types.Add("vsan:VsanVibInstall_TaskRequestType", reflect.TypeOf((*VsanVibInstall_TaskRequestType)(nil)).Elem())
}

type VsanVibInstall_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
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

type VsanUnmountDiskMappingEx VsanUnmountDiskMappingExRequestType

func init() {
	types.Add("vsan:VsanUnmountDiskMappingEx", reflect.TypeOf((*VsanUnmountDiskMappingEx)(nil)).Elem())
}

type VsanUnmountDiskMappingExRequestType struct {
	This            types.ManagedObjectReference `xml:"_this"`
	Mappings        []types.VsanHostDiskMapping  `xml:"mappings"`
	MaintenanceSpec *types.HostMaintenanceSpec   `xml:"maintenanceSpec,omitempty"`
	Timeout         int32                        `xml:"timeout,omitempty"`
}

func init() {
	types.Add("vsan:VsanUnmountDiskMappingExRequestType", reflect.TypeOf((*VsanUnmountDiskMappingExRequestType)(nil)).Elem())
}

type VsanUnmountDiskMappingExResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanQuerySyncingVsanObjects VsanQuerySyncingVsanObjectsRequestType

func init() {
	types.Add("vsan:VsanQuerySyncingVsanObjects", reflect.TypeOf((*VsanQuerySyncingVsanObjects)(nil)).Elem())
}

type VsanQuerySyncingVsanObjectsRequestType struct {
	This           types.ManagedObjectReference `xml:"_this"`
	Uuids          []string                     `xml:"uuids,omitempty"`
	Start          int32                        `xml:"start,omitempty"`
	Limit          int32                        `xml:"limit,omitempty"`
	IncludeSummary *bool                        `xml:"includeSummary"`
}

func init() {
	types.Add("vsan:VsanQuerySyncingVsanObjectsRequestType", reflect.TypeOf((*VsanQuerySyncingVsanObjectsRequestType)(nil)).Elem())
}

type VsanQuerySyncingVsanObjectsResponse struct {
	Returnval VsanHostVsanObjectSyncQueryResult `xml:"returnval"`
}

type VsanHostQueryWipeDisk VsanHostQueryWipeDiskRequestType

func init() {
	types.Add("vsan:VsanHostQueryWipeDisk", reflect.TypeOf((*VsanHostQueryWipeDisk)(nil)).Elem())
}

type VsanHostQueryWipeDiskRequestType struct {
	This  types.ManagedObjectReference `xml:"_this"`
	Disks []string                     `xml:"disks"`
}

func init() {
	types.Add("vsan:VsanHostQueryWipeDiskRequestType", reflect.TypeOf((*VsanHostQueryWipeDiskRequestType)(nil)).Elem())
}

type VsanHostQueryWipeDiskResponse struct {
	Returnval []VsanHostWipeDiskStatus `xml:"returnval,omitempty"`
}

type VsanQueryHostStatusEx VsanQueryHostStatusExRequestType

func init() {
	types.Add("vsan:VsanQueryHostStatusEx", reflect.TypeOf((*VsanQueryHostStatusEx)(nil)).Elem())
}

type VsanQueryHostStatusExRequestType struct {
	This         types.ManagedObjectReference `xml:"_this"`
	ClusterUuids []string                     `xml:"clusterUuids,omitempty"`
}

func init() {
	types.Add("vsan:VsanQueryHostStatusExRequestType", reflect.TypeOf((*VsanQueryHostStatusExRequestType)(nil)).Elem())
}

type VsanQueryHostStatusExResponse struct {
	Returnval []types.VsanHostClusterStatus `xml:"returnval,omitempty"`
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

type VsanHostGetRuntimeStats VsanHostGetRuntimeStatsRequestType

func init() {
	types.Add("vsan:VsanHostGetRuntimeStats", reflect.TypeOf((*VsanHostGetRuntimeStats)(nil)).Elem())
}

type VsanHostGetRuntimeStatsRequestType struct {
	This        types.ManagedObjectReference `xml:"_this"`
	Stats       []string                     `xml:"stats,omitempty"`
	ClusterUuid string                       `xml:"clusterUuid,omitempty"`
}

func init() {
	types.Add("vsan:VsanHostGetRuntimeStatsRequestType", reflect.TypeOf((*VsanHostGetRuntimeStatsRequestType)(nil)).Elem())
}

type VsanHostGetRuntimeStatsResponse struct {
	Returnval VsanHostRuntimeStats `xml:"returnval"`
}

type VsanHostAbortWipeDisk VsanHostAbortWipeDiskRequestType

func init() {
	types.Add("vsan:VsanHostAbortWipeDisk", reflect.TypeOf((*VsanHostAbortWipeDisk)(nil)).Elem())
}

type VsanHostAbortWipeDiskRequestType struct {
	This  types.ManagedObjectReference `xml:"_this"`
	Disks []string                     `xml:"disks"`
}

func init() {
	types.Add("vsan:VsanHostAbortWipeDiskRequestType", reflect.TypeOf((*VsanHostAbortWipeDiskRequestType)(nil)).Elem())
}

type VsanHostAbortWipeDiskResponse struct {
	Returnval []VsanHostAbortWipeDiskStatus `xml:"returnval,omitempty"`
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

type VsanHostWipeDisk VsanHostWipeDiskRequestType

func init() {
	types.Add("vsan:VsanHostWipeDisk", reflect.TypeOf((*VsanHostWipeDisk)(nil)).Elem())
}

type VsanHostWipeDiskRequestType struct {
	This  types.ManagedObjectReference `xml:"_this"`
	Disks []string                     `xml:"disks"`
}

func init() {
	types.Add("vsan:VsanHostWipeDiskRequestType", reflect.TypeOf((*VsanHostWipeDiskRequestType)(nil)).Elem())
}

type VsanHostWipeDiskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
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

type VsanRebalanceFileService VsanRebalanceFileServiceRequestType

func init() {
	types.Add("vsan:VsanRebalanceFileService", reflect.TypeOf((*VsanRebalanceFileService)(nil)).Elem())
}

type VsanRebalanceFileServiceRequestType struct {
	This    types.ManagedObjectReference  `xml:"_this"`
	Cluster *types.ManagedObjectReference `xml:"cluster,omitempty"`
}

func init() {
	types.Add("vsan:VsanRebalanceFileServiceRequestType", reflect.TypeOf((*VsanRebalanceFileServiceRequestType)(nil)).Elem())
}

type VsanRebalanceFileServiceResponse struct {
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

type VsanClusterReconfigureFsDomain VsanClusterReconfigureFsDomainRequestType

func init() {
	types.Add("vsan:VsanClusterReconfigureFsDomain", reflect.TypeOf((*VsanClusterReconfigureFsDomain)(nil)).Elem())
}

type VsanClusterReconfigureFsDomainRequestType struct {
	This                     types.ManagedObjectReference  `xml:"_this"`
	DomainUuid               string                        `xml:"domainUuid"`
	DomainConfig             VsanFileServiceDomainConfig   `xml:"domainConfig"`
	Cluster                  *types.ManagedObjectReference `xml:"cluster,omitempty"`
	DeleteDomainConfigFields []string                      `xml:"deleteDomainConfigFields,omitempty"`
}

func init() {
	types.Add("vsan:VsanClusterReconfigureFsDomainRequestType", reflect.TypeOf((*VsanClusterReconfigureFsDomainRequestType)(nil)).Elem())
}

type VsanClusterReconfigureFsDomainResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
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
	Returnval []types.OptionValue `xml:"returnval,omitempty"`
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
	Spec        *VsanIperfClientSpec         `xml:"spec,omitempty"`
}

func init() {
	types.Add("vsan:VsanHostQueryRunIperfClientRequestType", reflect.TypeOf((*VsanHostQueryRunIperfClientRequestType)(nil)).Elem())
}

type VsanHostQueryRunIperfClientResponse struct {
	Returnval VsanNetworkLoadTestResult `xml:"returnval"`
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
	Spec                          *VsanHealthQuerySpec         `xml:"spec,omitempty"`
}

func init() {
	types.Add("vsan:VsanHostQueryObjectHealthSummaryRequestType", reflect.TypeOf((*VsanHostQueryObjectHealthSummaryRequestType)(nil)).Elem())
}

type VsanHostQueryObjectHealthSummaryResponse struct {
	Returnval VsanObjectOverallHealth `xml:"returnval"`
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

type VsanHostQueryVerifyNetworkSettings VsanHostQueryVerifyNetworkSettingsRequestType

func init() {
	types.Add("vsan:VsanHostQueryVerifyNetworkSettings", reflect.TypeOf((*VsanHostQueryVerifyNetworkSettings)(nil)).Elem())
}

type VsanHostQueryVerifyNetworkSettingsRequestType struct {
	This                          types.ManagedObjectReference `xml:"_this"`
	Peers                         []string                     `xml:"peers,omitempty"`
	ROBOStretchedClusterWitnesses []string                     `xml:"ROBOStretchedClusterWitnesses,omitempty"`
	VMotionPeers                  []string                     `xml:"vMotionPeers,omitempty"`
	Spec                          *VsanHealthQuerySpec         `xml:"spec,omitempty"`
}

func init() {
	types.Add("vsan:VsanHostQueryVerifyNetworkSettingsRequestType", reflect.TypeOf((*VsanHostQueryVerifyNetworkSettingsRequestType)(nil)).Elem())
}

type VsanHostQueryVerifyNetworkSettingsResponse struct {
	Returnval VsanNetworkHealthResult `xml:"returnval"`
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

type VsanFlashScsiControllerFirmware_Task VsanFlashScsiControllerFirmware_TaskRequestType

func init() {
	types.Add("vsan:VsanFlashScsiControllerFirmware_Task", reflect.TypeOf((*VsanFlashScsiControllerFirmware_Task)(nil)).Elem())
}

type VsanFlashScsiControllerFirmware_TaskRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
	Spec VsanHclFirmwareUpdateSpec    `xml:"spec"`
}

func init() {
	types.Add("vsan:VsanFlashScsiControllerFirmware_TaskRequestType", reflect.TypeOf((*VsanFlashScsiControllerFirmware_TaskRequestType)(nil)).Elem())
}

type VsanFlashScsiControllerFirmware_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
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
	Returnval []VsanVmdkLoadTestResult `xml:"returnval,omitempty"`
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
	Returnval []VsanQueryResultHostInfo `xml:"returnval,omitempty"`
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

type VsanVcReplaceWitnessHostForClusters VsanVcReplaceWitnessHostForClustersRequestType

func init() {
	types.Add("vsan:VsanVcReplaceWitnessHostForClusters", reflect.TypeOf((*VsanVcReplaceWitnessHostForClusters)(nil)).Elem())
}

type VsanVcReplaceWitnessHostForClustersRequestType struct {
	This       types.ManagedObjectReference     `xml:"_this"`
	ConfigSpec VsanVcStretchedClusterConfigSpec `xml:"configSpec"`
}

func init() {
	types.Add("vsan:VsanVcReplaceWitnessHostForClustersRequestType", reflect.TypeOf((*VsanVcReplaceWitnessHostForClustersRequestType)(nil)).Elem())
}

type VsanVcReplaceWitnessHostForClustersResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanVcAddWitnessHostForClusters VsanVcAddWitnessHostForClustersRequestType

func init() {
	types.Add("vsan:VsanVcAddWitnessHostForClusters", reflect.TypeOf((*VsanVcAddWitnessHostForClusters)(nil)).Elem())
}

type VsanVcAddWitnessHostForClustersRequestType struct {
	This       types.ManagedObjectReference     `xml:"_this"`
	ConfigSpec VsanVcStretchedClusterConfigSpec `xml:"configSpec"`
}

func init() {
	types.Add("vsan:VsanVcAddWitnessHostForClustersRequestType", reflect.TypeOf((*VsanVcAddWitnessHostForClustersRequestType)(nil)).Elem())
}

type VsanVcAddWitnessHostForClustersResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
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

type QuerySharedWitnessClusterInfo QuerySharedWitnessClusterInfoRequestType

func init() {
	types.Add("vsan:QuerySharedWitnessClusterInfo", reflect.TypeOf((*QuerySharedWitnessClusterInfo)(nil)).Elem())
}

type QuerySharedWitnessClusterInfoRequestType struct {
	This        types.ManagedObjectReference `xml:"_this"`
	WitnessHost types.ManagedObjectReference `xml:"witnessHost"`
}

func init() {
	types.Add("vsan:QuerySharedWitnessClusterInfoRequestType", reflect.TypeOf((*QuerySharedWitnessClusterInfoRequestType)(nil)).Elem())
}

type QuerySharedWitnessClusterInfoResponse struct {
	Returnval []ClusterRuntimeInfo `xml:"returnval,omitempty"`
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

type VSANVcAddWitnessHost VSANVcAddWitnessHostRequestType

func init() {
	types.Add("vsan:VSANVcAddWitnessHost", reflect.TypeOf((*VSANVcAddWitnessHost)(nil)).Elem())
}

type VSANVcAddWitnessHostRequestType struct {
	This         types.ManagedObjectReference `xml:"_this"`
	Cluster      types.ManagedObjectReference `xml:"cluster"`
	WitnessHost  types.ManagedObjectReference `xml:"witnessHost"`
	PreferredFd  string                       `xml:"preferredFd"`
	DiskMapping  *types.VsanHostDiskMapping   `xml:"diskMapping,omitempty"`
	MetadataMode *bool                        `xml:"metadataMode"`
}

func init() {
	types.Add("vsan:VSANVcAddWitnessHostRequestType", reflect.TypeOf((*VSANVcAddWitnessHostRequestType)(nil)).Elem())
}

type VSANVcAddWitnessHostResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
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
	DiskMapping       *types.VsanHostDiskMapping                      `xml:"diskMapping,omitempty"`
}

func init() {
	types.Add("vsan:VSANVcConvertToStretchedClusterRequestType", reflect.TypeOf((*VSANVcConvertToStretchedClusterRequestType)(nil)).Elem())
}

type VSANVcConvertToStretchedClusterResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
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

type QuerySharedWitnessCompatibility QuerySharedWitnessCompatibilityRequestType

func init() {
	types.Add("vsan:QuerySharedWitnessCompatibility", reflect.TypeOf((*QuerySharedWitnessCompatibility)(nil)).Elem())
}

type QuerySharedWitnessCompatibilityRequestType struct {
	This              types.ManagedObjectReference   `xml:"_this"`
	SharedWitnessHost types.ManagedObjectReference   `xml:"sharedWitnessHost"`
	RoboClusters      []types.ManagedObjectReference `xml:"roboClusters"`
}

func init() {
	types.Add("vsan:QuerySharedWitnessCompatibilityRequestType", reflect.TypeOf((*QuerySharedWitnessCompatibilityRequestType)(nil)).Elem())
}

type QuerySharedWitnessCompatibilityResponse struct {
	Returnval VSANSharedWitnessCompatibilityResult `xml:"returnval"`
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
	Returnval []VsanPhysicalDiskHealthSummary `xml:"returnval,omitempty"`
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

type VsanEncryptionHealthSummary struct {
	types.DynamicData

	Hostname         string                     `xml:"hostname,omitempty"`
	EncryptionInfo   *VsanHostEncryptionInfo    `xml:"encryptionInfo,omitempty"`
	OverallKmsHealth string                     `xml:"overallKmsHealth"`
	KmsHealth        []VsanKmsHealth            `xml:"kmsHealth,omitempty"`
	EncryptionIssues []string                   `xml:"encryptionIssues,omitempty"`
	DiskResults      []VsanDiskEncryptionHealth `xml:"diskResults,omitempty"`
	Error            types.BaseMethodFault      `xml:"error,omitempty,typeattr"`
	AesniEnabled     *bool                      `xml:"aesniEnabled"`
}

func init() {
	types.Add("vsan:VsanEncryptionHealthSummary", reflect.TypeOf((*VsanEncryptionHealthSummary)(nil)).Elem())
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

type VsanClusterHealthGroup struct {
	types.DynamicData

	GroupId      string                            `xml:"groupId"`
	GroupName    string                            `xml:"groupName"`
	GroupHealth  string                            `xml:"groupHealth"`
	GroupTests   []VsanClusterHealthTest           `xml:"groupTests,omitempty"`
	GroupDetails []BaseVsanClusterHealthResultBase `xml:"groupDetails,omitempty,typeattr"`
	InProgress   *bool                             `xml:"inProgress"`
}

func init() {
	types.Add("vsan:VsanClusterHealthGroup", reflect.TypeOf((*VsanClusterHealthGroup)(nil)).Elem())
}

type VsanDiskGroupResourceCheckResult struct {
	EntityResourceCheckDetails

	CacheTierDisk     *VsanDiskResourceCheckResult  `xml:"cacheTierDisk,omitempty"`
	CapacityTierDisks []VsanDiskResourceCheckResult `xml:"capacityTierDisks,omitempty"`
}

func init() {
	types.Add("vsan:VsanDiskGroupResourceCheckResult", reflect.TypeOf((*VsanDiskGroupResourceCheckResult)(nil)).Elem())
}

type VsanSpaceUsageDetailResult struct {
	types.DynamicData

	SpaceUsageByObjectType []VsanObjectSpaceSummary `xml:"spaceUsageByObjectType,omitempty"`
}

func init() {
	types.Add("vsan:VsanSpaceUsageDetailResult", reflect.TypeOf((*VsanSpaceUsageDetailResult)(nil)).Elem())
}

type VsanSmartDiskStats struct {
	types.DynamicData

	Disk  string                `xml:"disk"`
	Stats []VsanSmartParameter  `xml:"stats,omitempty"`
	Error types.BaseMethodFault `xml:"error,omitempty,typeattr"`
}

func init() {
	types.Add("vsan:VsanSmartDiskStats", reflect.TypeOf((*VsanSmartDiskStats)(nil)).Elem())
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

type VsanPerfGraph struct {
	types.DynamicData

	Id          string             `xml:"id"`
	Metrics     []VsanPerfMetricId `xml:"metrics"`
	Unit        string             `xml:"unit"`
	Threshold   *VsanPerfThreshold `xml:"threshold,omitempty"`
	Name        string             `xml:"name,omitempty"`
	Description string             `xml:"description,omitempty"`
	SecondGraph *VsanPerfGraph     `xml:"secondGraph,omitempty"`
}

func init() {
	types.Add("vsan:VsanPerfGraph", reflect.TypeOf((*VsanPerfGraph)(nil)).Elem())
}

type VsanDiskDataEvacuationResourceCheckTaskDetails struct {
	VsanResourceCheckTaskDetails

	DiskUuid       string `xml:"diskUuid,omitempty"`
	IsCapacityTier *bool  `xml:"isCapacityTier"`
}

func init() {
	types.Add("vsan:VsanDiskDataEvacuationResourceCheckTaskDetails", reflect.TypeOf((*VsanDiskDataEvacuationResourceCheckTaskDetails)(nil)).Elem())
}

type VsanFileShareRuntimeInfo struct {
	types.DynamicData

	UsedCapacity    int64            `xml:"usedCapacity,omitempty"`
	Hostname        string           `xml:"hostname,omitempty"`
	Address         string           `xml:"address,omitempty"`
	VsanObjectUuids []string         `xml:"vsanObjectUuids,omitempty"`
	AccessPoints    []types.KeyValue `xml:"accessPoints,omitempty"`
	ManagedBy       string           `xml:"managedBy,omitempty"`
	FileServerFQDN  string           `xml:"fileServerFQDN,omitempty"`
}

func init() {
	types.Add("vsan:VsanFileShareRuntimeInfo", reflect.TypeOf((*VsanFileShareRuntimeInfo)(nil)).Elem())
}

type VsanResourceConstraint struct {
	types.DynamicData

	TargetType string `xml:"targetType,omitempty"`
}

func init() {
	types.Add("vsan:VsanResourceConstraint", reflect.TypeOf((*VsanResourceConstraint)(nil)).Elem())
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
	ServerClusters    []VsanServerClusterInfo       `xml:"serverClusters,omitempty"`
}

func init() {
	types.Add("vsan:VsanNetworkHealthResult", reflect.TypeOf((*VsanNetworkHealthResult)(nil)).Elem())
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

type HostSpbmDatastoreInfo struct {
	types.DynamicData

	DatastoreUrl     string `xml:"datastoreUrl"`
	Namespace        string `xml:"namespace"`
	DefaultProfileId string `xml:"defaultProfileId"`
}

func init() {
	types.Add("vsan:HostSpbmDatastoreInfo", reflect.TypeOf((*HostSpbmDatastoreInfo)(nil)).Elem())
}

type VsanBrokenDiskChainIssue struct {
	types.VsanUpgradeSystemPreflightCheckIssue

	Uuids []string `xml:"uuids"`
}

func init() {
	types.Add("vsan:VsanBrokenDiskChainIssue", reflect.TypeOf((*VsanBrokenDiskChainIssue)(nil)).Elem())
}

type VsanClusterNetworkLoadTestResult struct {
	types.DynamicData

	ClusterResult VsanClusterProactiveTestResult `xml:"clusterResult"`
	HostResults   []VsanNetworkLoadTestResult    `xml:"hostResults,omitempty"`
}

func init() {
	types.Add("vsan:VsanClusterNetworkLoadTestResult", reflect.TypeOf((*VsanClusterNetworkLoadTestResult)(nil)).Elem())
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

type VsanClusterNetworkPartitionInfo struct {
	types.DynamicData

	Hosts            []string `xml:"hosts,omitempty"`
	PartitionUnknown *bool    `xml:"partitionUnknown"`
}

func init() {
	types.Add("vsan:VsanClusterNetworkPartitionInfo", reflect.TypeOf((*VsanClusterNetworkPartitionInfo)(nil)).Elem())
}

type VsanFileServiceDomainConfig struct {
	types.DynamicData

	Name                  string                     `xml:"name,omitempty"`
	DnsServerAddresses    []string                   `xml:"dnsServerAddresses,omitempty"`
	DnsSuffixes           []string                   `xml:"dnsSuffixes,omitempty"`
	FileServerIpConfig    []VsanFileServiceIpConfig  `xml:"fileServerIpConfig,omitempty"`
	DirectoryServerConfig *VsanDirectoryServerConfig `xml:"directoryServerConfig,omitempty"`
	Version               string                     `xml:"version,omitempty"`
}

func init() {
	types.Add("vsan:VsanFileServiceDomainConfig", reflect.TypeOf((*VsanFileServiceDomainConfig)(nil)).Elem())
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

type VsanMixedEsxVersionIssue struct {
	types.VsanUpgradeSystemPreflightCheckIssue
}

func init() {
	types.Add("vsan:VsanMixedEsxVersionIssue", reflect.TypeOf((*VsanMixedEsxVersionIssue)(nil)).Elem())
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

type VsanNetworkConfigPortgroupWithNoRedundancyIssue struct {
	VsanNetworkConfigBaseIssue

	Host          types.ManagedObjectReference  `xml:"host"`
	PortgroupName string                        `xml:"portgroupName,omitempty"`
	Vds           *types.ManagedObjectReference `xml:"vds,omitempty,typeattr"`
	Pg            *types.ManagedObjectReference `xml:"pg,omitempty"`
	NumPnics      int64                         `xml:"numPnics"`
}

func init() {
	types.Add("vsan:VsanNetworkConfigPortgroupWithNoRedundancyIssue", reflect.TypeOf((*VsanNetworkConfigPortgroupWithNoRedundancyIssue)(nil)).Elem())
}

type VsanNetworkVMotionVmknicNotFountIssue struct {
	VsanNetworkConfigBaseIssue

	HostWithoutVmotionVmknic types.ManagedObjectReference `xml:"hostWithoutVmotionVmknic"`
}

func init() {
	types.Add("vsan:VsanNetworkVMotionVmknicNotFountIssue", reflect.TypeOf((*VsanNetworkVMotionVmknicNotFountIssue)(nil)).Elem())
}

type ResyncIopsInfo struct {
	types.DynamicData

	ResyncIops int32 `xml:"resyncIops"`
}

func init() {
	types.Add("vsan:ResyncIopsInfo", reflect.TypeOf((*ResyncIopsInfo)(nil)).Elem())
}

type VsanUnmapConfig struct {
	types.DynamicData

	Enable bool `xml:"enable"`
}

func init() {
	types.Add("vsan:VsanUnmapConfig", reflect.TypeOf((*VsanUnmapConfig)(nil)).Elem())
}

type VimVsanReconfigSpec struct {
	types.SDDCBase

	VsanClusterConfig             BaseVsanClusterConfigInfo             `xml:"vsanClusterConfig,omitempty,typeattr"`
	DataEfficiencyConfig          *VsanDataEfficiencyConfig             `xml:"dataEfficiencyConfig,omitempty"`
	DiskMappingSpec               *VimClusterVsanDiskMappingsConfigSpec `xml:"diskMappingSpec,omitempty"`
	FaultDomainsSpec              *VimClusterVsanFaultDomainsConfigSpec `xml:"faultDomainsSpec,omitempty"`
	Modify                        bool                                  `xml:"modify"`
	AllowReducedRedundancy        *bool                                 `xml:"allowReducedRedundancy"`
	ResyncIopsLimitConfig         *ResyncIopsInfo                       `xml:"resyncIopsLimitConfig,omitempty"`
	IscsiSpec                     *VsanIscsiTargetServiceSpec           `xml:"iscsiSpec,omitempty"`
	DataEncryptionConfig          *VsanDataEncryptionConfig             `xml:"dataEncryptionConfig,omitempty"`
	ExtendedConfig                *VsanExtendedConfig                   `xml:"extendedConfig,omitempty"`
	DatastoreConfig               BaseVsanDatastoreConfig               `xml:"datastoreConfig,omitempty,typeattr"`
	PerfsvcConfig                 *VsanPerfsvcConfig                    `xml:"perfsvcConfig,omitempty"`
	UnmapConfig                   *VsanUnmapConfig                      `xml:"unmapConfig,omitempty"`
	VumConfig                     *VsanVumConfig                        `xml:"vumConfig,omitempty"`
	MetricsConfig                 *VsanMetricsConfig                    `xml:"metricsConfig,omitempty"`
	FileServiceConfig             *VsanFileServiceConfig                `xml:"fileServiceConfig,omitempty"`
	DataInTransitEncryptionConfig *VsanDataInTransitEncryptionConfig    `xml:"dataInTransitEncryptionConfig,omitempty"`
}

func init() {
	types.Add("vsan:VimVsanReconfigSpec", reflect.TypeOf((*VimVsanReconfigSpec)(nil)).Elem())
}

type VsanObjectTypeRule struct {
	types.DynamicData

	ObjectType string   `xml:"objectType,omitempty"`
	Attributes []string `xml:"attributes,omitempty"`
}

func init() {
	types.Add("vsan:VsanObjectTypeRule", reflect.TypeOf((*VsanObjectTypeRule)(nil)).Elem())
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

type VsanLimitHealthResult struct {
	types.DynamicData

	Hostname                     string                           `xml:"hostname,omitempty"`
	IssueFound                   bool                             `xml:"issueFound"`
	MaxComponents                int32                            `xml:"maxComponents"`
	FreeComponents               int32                            `xml:"freeComponents"`
	ComponentLimitHealth         string                           `xml:"componentLimitHealth"`
	LowestFreeDiskSpacePct       int32                            `xml:"lowestFreeDiskSpacePct"`
	UsedDiskSpaceB               int64                            `xml:"usedDiskSpaceB"`
	TotalDiskSpaceB              int64                            `xml:"totalDiskSpaceB"`
	DiskFreeSpaceHealth          string                           `xml:"diskFreeSpaceHealth"`
	ReservedRcSizeB              int64                            `xml:"reservedRcSizeB"`
	TotalRcSizeB                 int64                            `xml:"totalRcSizeB"`
	RcFreeReservationHealth      string                           `xml:"rcFreeReservationHealth"`
	TotalLogicalSpaceB           int64                            `xml:"totalLogicalSpaceB,omitempty"`
	LogicalSpaceUsedB            int64                            `xml:"logicalSpaceUsedB,omitempty"`
	DedupMetadataSizeB           int64                            `xml:"dedupMetadataSizeB,omitempty"`
	DiskTransientCapacityUsedB   int64                            `xml:"diskTransientCapacityUsedB,omitempty"`
	DgTransientCapacityUsedB     int64                            `xml:"dgTransientCapacityUsedB,omitempty"`
	SlackSpaceCapRequired        int64                            `xml:"slackSpaceCapRequired,omitempty"`
	ResyncPauseThreshold         int64                            `xml:"resyncPauseThreshold,omitempty"`
	SpaceEfficiencyMetadataSizeB *VsanSpaceEfficiencyMetadataSize `xml:"spaceEfficiencyMetadataSizeB,omitempty"`
	HostRebuildCapacity          int64                            `xml:"hostRebuildCapacity,omitempty"`
	MinSpaceRequiredForVsanOp    int64                            `xml:"minSpaceRequiredForVsanOp,omitempty"`
	EnforceCapResrvSpace         int64                            `xml:"enforceCapResrvSpace,omitempty"`
}

func init() {
	types.Add("vsan:VsanLimitHealthResult", reflect.TypeOf((*VsanLimitHealthResult)(nil)).Elem())
}

type VsanIscsiInitiatorGroup struct {
	types.DynamicData

	Name       string                     `xml:"name"`
	Initiators []string                   `xml:"initiators,omitempty"`
	Targets    []VsanIscsiTargetBasicInfo `xml:"targets,omitempty"`
}

func init() {
	types.Add("vsan:VsanIscsiInitiatorGroup", reflect.TypeOf((*VsanIscsiInitiatorGroup)(nil)).Elem())
}

type VsanDiskUnhealthIssue struct {
	types.VsanUpgradeSystemPreflightCheckIssue

	Uuids []string `xml:"uuids"`
}

func init() {
	types.Add("vsan:VsanDiskUnhealthIssue", reflect.TypeOf((*VsanDiskUnhealthIssue)(nil)).Elem())
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

type ClusterRuntimeInfo struct {
	types.DynamicData

	ClusterUuid          string                        `xml:"clusterUuid"`
	TotalComponentsCount int32                         `xml:"totalComponentsCount"`
	Cluster              *types.ManagedObjectReference `xml:"cluster,omitempty"`
}

func init() {
	types.Add("vsan:ClusterRuntimeInfo", reflect.TypeOf((*ClusterRuntimeInfo)(nil)).Elem())
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

type VsanResourceHealth struct {
	types.DynamicData

	Resource    string `xml:"resource"`
	Health      string `xml:"health"`
	Description string `xml:"description,omitempty"`
}

func init() {
	types.Add("vsan:VsanResourceHealth", reflect.TypeOf((*VsanResourceHealth)(nil)).Elem())
}

type VsanDatastoreConfig struct {
	types.DynamicData

	Datastores []BaseVsanDatastoreSpec `xml:"datastores,omitempty,typeattr"`
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

type VsanSpaceEfficiencyMetadataSize struct {
	types.DynamicData

	DedupMetadataSize       int64 `xml:"dedupMetadataSize,omitempty"`
	CompressionMetadataSize int64 `xml:"compressionMetadataSize,omitempty"`
}

func init() {
	types.Add("vsan:VsanSpaceEfficiencyMetadataSize", reflect.TypeOf((*VsanSpaceEfficiencyMetadataSize)(nil)).Elem())
}

type VsanKmsHealth struct {
	types.DynamicData

	ServerName     string                `xml:"serverName"`
	Health         string                `xml:"health"`
	Error          types.BaseMethodFault `xml:"error,omitempty,typeattr"`
	TrustHealth    string                `xml:"trustHealth,omitempty"`
	CertHealth     string                `xml:"certHealth,omitempty"`
	CertExpireDate *time.Time            `xml:"certExpireDate"`
}

func init() {
	types.Add("vsan:VsanKmsHealth", reflect.TypeOf((*VsanKmsHealth)(nil)).Elem())
}

type VsanMountPrecheckNetworkConnectivityResult struct {
	VsanMountPrecheckItem

	Details []VsanMountPrecheckNetworkConnectivityDetail `xml:"details,omitempty"`
}

func init() {
	types.Add("vsan:VsanMountPrecheckNetworkConnectivityResult", reflect.TypeOf((*VsanMountPrecheckNetworkConnectivityResult)(nil)).Elem())
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

type VsanWhatIfEvacResult struct {
	types.DynamicData

	NoAction     VsanWhatIfEvacDetail `xml:"noAction"`
	EnsureAccess VsanWhatIfEvacDetail `xml:"ensureAccess"`
	EvacAllData  VsanWhatIfEvacDetail `xml:"evacAllData"`
}

func init() {
	types.Add("vsan:VsanWhatIfEvacResult", reflect.TypeOf((*VsanWhatIfEvacResult)(nil)).Elem())
}

type VsanHealthThreshold struct {
	types.DynamicData

	YellowValue int64 `xml:"yellowValue"`
	RedValue    int64 `xml:"redValue"`
}

func init() {
	types.Add("vsan:VsanHealthThreshold", reflect.TypeOf((*VsanHealthThreshold)(nil)).Elem())
}

type VsanDiskEncryptionHealth struct {
	types.DynamicData

	DiskHealth       *VsanPhysicalDiskHealth `xml:"diskHealth,omitempty"`
	EncryptionIssues []string                `xml:"encryptionIssues,omitempty"`
}

func init() {
	types.Add("vsan:VsanDiskEncryptionHealth", reflect.TypeOf((*VsanDiskEncryptionHealth)(nil)).Elem())
}

type VsanRuntimeStatsHostMap struct {
	types.DynamicData

	Host  types.ManagedObjectReference `xml:"host"`
	Stats *VsanHostRuntimeStats        `xml:"stats,omitempty"`
}

func init() {
	types.Add("vsan:VsanRuntimeStatsHostMap", reflect.TypeOf((*VsanRuntimeStatsHostMap)(nil)).Elem())
}

type VimVsanHostVsanDirectStorage struct {
	types.DynamicData

	ScsiDisks []VimVsanHostVsanScsiDisk `xml:"scsiDisks,omitempty"`
	Tier      string                    `xml:"tier,omitempty"`
}

func init() {
	types.Add("vsan:VimVsanHostVsanDirectStorage", reflect.TypeOf((*VimVsanHostVsanDirectStorage)(nil)).Elem())
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

type VsanSpaceUsage struct {
	types.DynamicData

	TotalCapacityB          int64                               `xml:"totalCapacityB"`
	FreeCapacityB           int64                               `xml:"freeCapacityB,omitempty"`
	SpaceOverview           *VsanObjectSpaceSummary             `xml:"spaceOverview,omitempty"`
	SpaceDetail             *VsanSpaceUsageDetailResult         `xml:"spaceDetail,omitempty"`
	EfficientCapacity       *VimVsanDataEfficiencyCapacityState `xml:"efficientCapacity,omitempty"`
	WhatifCapacities        []VsanWhatifCapacity                `xml:"whatifCapacities,omitempty"`
	UncommittedB            int64                               `xml:"uncommittedB,omitempty"`
	CapacityHealthThreshold *VsanHealthThreshold                `xml:"capacityHealthThreshold,omitempty"`
}

func init() {
	types.Add("vsan:VsanSpaceUsage", reflect.TypeOf((*VsanSpaceUsage)(nil)).Elem())
}

type VsanHclDiskInfo struct {
	types.DynamicData

	DeviceName       string                  `xml:"deviceName"`
	Model            string                  `xml:"model,omitempty"`
	IsSsd            *bool                   `xml:"isSsd"`
	VsanDisk         bool                    `xml:"vsanDisk"`
	Issues           []types.BaseMethodFault `xml:"issues,omitempty,typeattr"`
	RemediableIssues []string                `xml:"remediableIssues,omitempty"`
}

func init() {
	types.Add("vsan:VsanHclDiskInfo", reflect.TypeOf((*VsanHclDiskInfo)(nil)).Elem())
}

type VsanFileServiceDomainQuerySpec struct {
	types.DynamicData

	Uuids []string `xml:"uuids,omitempty"`
	Names []string `xml:"names,omitempty"`
}

func init() {
	types.Add("vsan:VsanFileServiceDomainQuerySpec", reflect.TypeOf((*VsanFileServiceDomainQuerySpec)(nil)).Elem())
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

type VsanQueryResultHostInfo struct {
	types.DynamicData

	Uuid              string   `xml:"uuid,omitempty"`
	HostnameInCmmds   string   `xml:"hostnameInCmmds,omitempty"`
	VsanIpv4Addresses []string `xml:"vsanIpv4Addresses,omitempty"`
}

func init() {
	types.Add("vsan:VsanQueryResultHostInfo", reflect.TypeOf((*VsanQueryResultHostInfo)(nil)).Elem())
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

type VsanNetworkConfigBaseIssue struct {
	types.DynamicData
}

func init() {
	types.Add("vsan:VsanNetworkConfigBaseIssue", reflect.TypeOf((*VsanNetworkConfigBaseIssue)(nil)).Elem())
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

type VsanServerClusterInfo struct {
	types.DynamicData

	Cluster    *types.ManagedObjectReference `xml:"cluster,omitempty"`
	PeerHealth []VsanNetworkPeerHealthResult `xml:"peerHealth,omitempty"`
	Membership *VsanClusterMembershipInfo    `xml:"membership,omitempty"`
}

func init() {
	types.Add("vsan:VsanServerClusterInfo", reflect.TypeOf((*VsanServerClusterInfo)(nil)).Elem())
}

type VsanRemoteClusterQuerySpec struct {
	types.DynamicData

	StartTime *time.Time `xml:"startTime"`
	EndTime   *time.Time `xml:"endTime"`
}

func init() {
	types.Add("vsan:VsanRemoteClusterQuerySpec", reflect.TypeOf((*VsanRemoteClusterQuerySpec)(nil)).Elem())
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

type VsanProactiveRebalanceInfoEx struct {
	types.DynamicData

	Running           *bool                 `xml:"running"`
	StartTs           *time.Time            `xml:"startTs"`
	StopTs            *time.Time            `xml:"stopTs"`
	VarianceThreshold float32               `xml:"varianceThreshold,omitempty"`
	TimeThreshold     int32                 `xml:"timeThreshold,omitempty"`
	RateThreshold     int32                 `xml:"rateThreshold,omitempty"`
	Hostname          string                `xml:"hostname,omitempty"`
	Error             types.BaseMethodFault `xml:"error,omitempty,typeattr"`
}

func init() {
	types.Add("vsan:VsanProactiveRebalanceInfoEx", reflect.TypeOf((*VsanProactiveRebalanceInfoEx)(nil)).Elem())
}

type VsanHostIoInsightInfo struct {
	types.DynamicData

	Host             types.ManagedObjectReference `xml:"host"`
	IoinsightWorldId int64                        `xml:"ioinsightWorldId,omitempty"`
	FaultMessage     string                       `xml:"faultMessage,omitempty"`
	IoinsightInfo    *VsanIoInsightInfo           `xml:"ioinsightInfo,omitempty"`
}

func init() {
	types.Add("vsan:VsanHostIoInsightInfo", reflect.TypeOf((*VsanHostIoInsightInfo)(nil)).Elem())
}

type VsanDataEfficiencyConfig struct {
	types.DynamicData

	DedupEnabled       bool  `xml:"dedupEnabled"`
	CompressionEnabled *bool `xml:"compressionEnabled"`
}

func init() {
	types.Add("vsan:VsanDataEfficiencyConfig", reflect.TypeOf((*VsanDataEfficiencyConfig)(nil)).Elem())
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

type VsanUnknownScanIssue struct {
	types.VsanUpgradeSystemPreflightCheckIssue

	Uuids []string `xml:"uuids"`
}

func init() {
	types.Add("vsan:VsanUnknownScanIssue", reflect.TypeOf((*VsanUnknownScanIssue)(nil)).Elem())
}

type VsanClusterHostVmknicMapping struct {
	types.DynamicData

	Host   string `xml:"host"`
	Vmknic string `xml:"vmknic"`
}

func init() {
	types.Add("vsan:VsanClusterHostVmknicMapping", reflect.TypeOf((*VsanClusterHostVmknicMapping)(nil)).Elem())
}

type VsanClusterFileServiceHealthSummary struct {
	types.DynamicData

	OverallHealth string                         `xml:"overallHealth,omitempty"`
	HostResults   []VsanFileServiceHealthSummary `xml:"hostResults,omitempty"`
}

func init() {
	types.Add("vsan:VsanClusterFileServiceHealthSummary", reflect.TypeOf((*VsanClusterFileServiceHealthSummary)(nil)).Elem())
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

type VsanClusterBalanceSummary struct {
	types.DynamicData

	VarianceThreshold int64                           `xml:"varianceThreshold"`
	Disks             []VsanClusterBalancePerDiskInfo `xml:"disks,omitempty"`
}

func init() {
	types.Add("vsan:VsanClusterBalanceSummary", reflect.TypeOf((*VsanClusterBalanceSummary)(nil)).Elem())
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

type VsanConfigGeneration struct {
	types.DynamicData

	VcUuid  string `xml:"vcUuid"`
	GenNum  int64  `xml:"genNum"`
	GenTime int64  `xml:"genTime"`
}

func init() {
	types.Add("vsan:VsanConfigGeneration", reflect.TypeOf((*VsanConfigGeneration)(nil)).Elem())
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

type VsanFileShareNetPermission struct {
	types.DynamicData

	Ips         string `xml:"ips"`
	Permissions string `xml:"permissions,omitempty"`
	AllowRoot   *bool  `xml:"allowRoot"`
}

func init() {
	types.Add("vsan:VsanFileShareNetPermission", reflect.TypeOf((*VsanFileShareNetPermission)(nil)).Elem())
}

type VsanIscsiTargetCommonInfo struct {
	VsanIscsiTargetBasicInfo

	AuthSpec         *VsanIscsiTargetAuthSpec `xml:"authSpec,omitempty"`
	Port             int32                    `xml:"port,omitempty"`
	NetworkInterface string                   `xml:"networkInterface,omitempty"`
	AffinityLocation string                   `xml:"affinityLocation,omitempty"`
}

func init() {
	types.Add("vsan:VsanIscsiTargetCommonInfo", reflect.TypeOf((*VsanIscsiTargetCommonInfo)(nil)).Elem())
}

type ActiveVsanDirectoryServerConfig struct {
	VsanDirectoryServerConfig

	ActiveDirectoryDomainName string `xml:"activeDirectoryDomainName,omitempty"`
	Username                  string `xml:"username,omitempty"`
	Password                  string `xml:"password,omitempty"`
	OrganizationalUnit        string `xml:"organizationalUnit,omitempty"`
}

func init() {
	types.Add("vsan:ActiveVsanDirectoryServerConfig", reflect.TypeOf((*ActiveVsanDirectoryServerConfig)(nil)).Elem())
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

type VsanClusterHealthResultBase struct {
	types.DynamicData

	Label string `xml:"label,omitempty"`
}

func init() {
	types.Add("vsan:VsanClusterHealthResultBase", reflect.TypeOf((*VsanClusterHealthResultBase)(nil)).Elem())
}

type VsanNetworkConfigVdsScopeIssue struct {
	VsanNetworkConfigBaseIssue

	Vds            types.ManagedObjectReference   `xml:"vds,typeattr"`
	MemberHosts    []types.ManagedObjectReference `xml:"memberHosts"`
	NonMemberHosts []types.ManagedObjectReference `xml:"nonMemberHosts"`
}

func init() {
	types.Add("vsan:VsanNetworkConfigVdsScopeIssue", reflect.TypeOf((*VsanNetworkConfigVdsScopeIssue)(nil)).Elem())
}

type VimClusterVsanFaultDomainsConfigSpec struct {
	types.DynamicData

	FaultDomains []VimClusterVsanFaultDomainSpec `xml:"faultDomains"`
	Witness      *VimClusterVsanWitnessSpec      `xml:"witness,omitempty"`
}

func init() {
	types.Add("vsan:VimClusterVsanFaultDomainsConfigSpec", reflect.TypeOf((*VimClusterVsanFaultDomainsConfigSpec)(nil)).Elem())
}

type VsanMassCollectorSpec struct {
	types.DynamicData

	Objects          []types.ManagedObjectReference    `xml:"objects,omitempty,typeattr"`
	ObjectCollection string                            `xml:"objectCollection,omitempty"`
	Properties       []string                          `xml:"properties"`
	PropertiesParams []VsanMassCollectorPropertyParams `xml:"propertiesParams,omitempty"`
	Constraint       BaseVsanResourceConstraint        `xml:"constraint,omitempty,typeattr"`
}

func init() {
	types.Add("vsan:VsanMassCollectorSpec", reflect.TypeOf((*VsanMassCollectorSpec)(nil)).Elem())
}

type VsanConfigNotAllDisksClaimedIssue struct {
	VsanConfigBaseIssue

	Host  types.ManagedObjectReference `xml:"host"`
	Disks []string                     `xml:"disks"`
}

func init() {
	types.Add("vsan:VsanConfigNotAllDisksClaimedIssue", reflect.TypeOf((*VsanConfigNotAllDisksClaimedIssue)(nil)).Elem())
}

type VsanHostIpConfigEx struct {
	types.VsanHostIpConfig

	UpstreamIpV6Address   string `xml:"upstreamIpV6Address,omitempty"`
	DownstreamIpV6Address string `xml:"downstreamIpV6Address,omitempty"`
}

func init() {
	types.Add("vsan:VsanHostIpConfigEx", reflect.TypeOf((*VsanHostIpConfigEx)(nil)).Elem())
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

type VsanClusterClomdLivenessResult struct {
	types.DynamicData

	ClomdLivenessResult []VsanHostClomdLivenessResult `xml:"clomdLivenessResult,omitempty"`
	IssueFound          bool                          `xml:"issueFound"`
}

func init() {
	types.Add("vsan:VsanClusterClomdLivenessResult", reflect.TypeOf((*VsanClusterClomdLivenessResult)(nil)).Elem())
}

type VsanHostWithHybridDiskgroupIssue struct {
	types.VsanUpgradeSystemPreflightCheckIssue

	Hosts []types.ManagedObjectReference `xml:"hosts"`
}

func init() {
	types.Add("vsan:VsanHostWithHybridDiskgroupIssue", reflect.TypeOf((*VsanHostWithHybridDiskgroupIssue)(nil)).Elem())
}

type VimClusterVsanHostDiskMapping struct {
	types.DynamicData

	Host          types.ManagedObjectReference `xml:"host"`
	CacheDisks    []types.HostScsiDisk         `xml:"cacheDisks,omitempty"`
	CapacityDisks []types.HostScsiDisk         `xml:"capacityDisks,omitempty"`
	Type          string                       `xml:"type"`
}

func init() {
	types.Add("vsan:VimClusterVsanHostDiskMapping", reflect.TypeOf((*VimClusterVsanHostDiskMapping)(nil)).Elem())
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

type VsanMountPrecheckNetworkLatencyResult struct {
	VsanMountPrecheckItem

	Details []VsanMountPrecheckNetworkLatencyDetail `xml:"details"`
}

func init() {
	types.Add("vsan:VsanMountPrecheckNetworkLatencyResult", reflect.TypeOf((*VsanMountPrecheckNetworkLatencyResult)(nil)).Elem())
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

type VsanHostCreateVmHealthTestResult struct {
	types.DynamicData

	Hostname string                `xml:"hostname"`
	State    string                `xml:"state"`
	Fault    types.BaseMethodFault `xml:"fault,omitempty,typeattr"`
}

func init() {
	types.Add("vsan:VsanHostCreateVmHealthTestResult", reflect.TypeOf((*VsanHostCreateVmHealthTestResult)(nil)).Elem())
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

type VsanHostHclInfo struct {
	types.DynamicData

	Hostname    string                  `xml:"hostname"`
	HclChecked  bool                    `xml:"hclChecked"`
	ReleaseName string                  `xml:"releaseName,omitempty"`
	Error       types.BaseMethodFault   `xml:"error,omitempty,typeattr"`
	Controllers []VsanHclControllerInfo `xml:"controllers,omitempty"`
	Pnics       []VsanHclNicInfo        `xml:"pnics,omitempty"`
}

func init() {
	types.Add("vsan:VsanHostHclInfo", reflect.TypeOf((*VsanHostHclInfo)(nil)).Elem())
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

type VimClusterVsanDiskMappingsConfigSpec struct {
	types.DynamicData

	HostDiskMappings []VimClusterVsanHostDiskMapping `xml:"hostDiskMappings"`
}

func init() {
	types.Add("vsan:VimClusterVsanDiskMappingsConfigSpec", reflect.TypeOf((*VimClusterVsanDiskMappingsConfigSpec)(nil)).Elem())
}

type VsanJsonComparator struct {
	VsanComparator

	Comparator      string             `xml:"comparator,omitempty"`
	ComparableValue *types.KeyAnyValue `xml:"comparableValue,omitempty"`
}

func init() {
	types.Add("vsan:VsanJsonComparator", reflect.TypeOf((*VsanJsonComparator)(nil)).Elem())
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

type VsanCompliantDriver struct {
	types.DynamicData

	DriverName    string `xml:"driverName"`
	DriverVersion string `xml:"driverVersion"`
}

func init() {
	types.Add("vsan:VsanCompliantDriver", reflect.TypeOf((*VsanCompliantDriver)(nil)).Elem())
}

type VsanDiskFormatConversionCheckResult struct {
	types.VsanUpgradeSystemPreflightCheckResult

	IsSupported            bool  `xml:"isSupported"`
	TargetVersion          int32 `xml:"targetVersion,omitempty"`
	IsDataMovementRequired *bool `xml:"isDataMovementRequired"`
}

func init() {
	types.Add("vsan:VsanDiskFormatConversionCheckResult", reflect.TypeOf((*VsanDiskFormatConversionCheckResult)(nil)).Elem())
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

type VsanPerfsvcConfig struct {
	types.DynamicData

	Enabled        bool                             `xml:"enabled"`
	Profile        *types.VirtualMachineProfileSpec `xml:"profile,omitempty"`
	DiagnosticMode *bool                            `xml:"diagnosticMode"`
	VerboseMode    *bool                            `xml:"verboseMode"`
}

func init() {
	types.Add("vsan:VsanPerfsvcConfig", reflect.TypeOf((*VsanPerfsvcConfig)(nil)).Elem())
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

type VsanNetworkConfigVswitchWithNoRedundancyIssue struct {
	VsanNetworkConfigBaseIssue

	Host        types.ManagedObjectReference  `xml:"host"`
	VswitchName string                        `xml:"vswitchName,omitempty"`
	Vds         *types.ManagedObjectReference `xml:"vds,omitempty,typeattr"`
	NumPnics    int64                         `xml:"numPnics"`
}

func init() {
	types.Add("vsan:VsanNetworkConfigVswitchWithNoRedundancyIssue", reflect.TypeOf((*VsanNetworkConfigVswitchWithNoRedundancyIssue)(nil)).Elem())
}

type VimClusterVsanWitnessSpec struct {
	types.DynamicData

	Host                     types.ManagedObjectReference `xml:"host"`
	PreferredFaultDomainName string                       `xml:"preferredFaultDomainName"`
	DiskMapping              *types.VsanHostDiskMapping   `xml:"diskMapping,omitempty"`
}

func init() {
	types.Add("vsan:VimClusterVsanWitnessSpec", reflect.TypeOf((*VimClusterVsanWitnessSpec)(nil)).Elem())
}

type VsanConfigBaseIssue struct {
	types.DynamicData
}

func init() {
	types.Add("vsan:VsanConfigBaseIssue", reflect.TypeOf((*VsanConfigBaseIssue)(nil)).Elem())
}

type VsanVsanClusterPcapGroup struct {
	types.DynamicData

	Master  string   `xml:"master"`
	Members []string `xml:"members,omitempty"`
}

func init() {
	types.Add("vsan:VsanVsanClusterPcapGroup", reflect.TypeOf((*VsanVsanClusterPcapGroup)(nil)).Elem())
}

type VsanPerfNodeInformation struct {
	types.DynamicData

	Version        string                     `xml:"version"`
	Hostname       string                     `xml:"hostname,omitempty"`
	Error          types.BaseMethodFault      `xml:"error,omitempty,typeattr"`
	IsCmmdsMaster  bool                       `xml:"isCmmdsMaster"`
	IsStatsMaster  bool                       `xml:"isStatsMaster"`
	VsanMasterUuid string                     `xml:"vsanMasterUuid,omitempty"`
	VsanNodeUuid   string                     `xml:"vsanNodeUuid,omitempty"`
	MasterInfo     *VsanPerfMasterInformation `xml:"masterInfo,omitempty"`
	DiagnosticMode *bool                      `xml:"diagnosticMode"`
}

func init() {
	types.Add("vsan:VsanPerfNodeInformation", reflect.TypeOf((*VsanPerfNodeInformation)(nil)).Elem())
}

type VsanIoInsightInfo struct {
	types.DynamicData

	State        string                         `xml:"state,omitempty"`
	MonitoredVMs []types.ManagedObjectReference `xml:"monitoredVMs,omitempty"`
}

func init() {
	types.Add("vsan:VsanIoInsightInfo", reflect.TypeOf((*VsanIoInsightInfo)(nil)).Elem())
}

type VsanInternalExtendedConfig struct {
	types.DynamicData

	VcMaxDiskVersion int32 `xml:"vcMaxDiskVersion,omitempty"`
}

func init() {
	types.Add("vsan:VsanInternalExtendedConfig", reflect.TypeOf((*VsanInternalExtendedConfig)(nil)).Elem())
}

type VsanVdsPgMigrationVmInfo struct {
	types.DynamicData

	Vm        types.ManagedObjectReference `xml:"vm"`
	VnicLabel []string                     `xml:"vnicLabel"`
}

func init() {
	types.Add("vsan:VsanVdsPgMigrationVmInfo", reflect.TypeOf((*VsanVdsPgMigrationVmInfo)(nil)).Elem())
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

type VsanClusterHealthResultRow struct {
	types.DynamicData

	Values     []string                     `xml:"values"`
	NestedRows []VsanClusterHealthResultRow `xml:"nestedRows,omitempty"`
}

func init() {
	types.Add("vsan:VsanClusterHealthResultRow", reflect.TypeOf((*VsanClusterHealthResultRow)(nil)).Elem())
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

type VsanHostHealthSystemVersionResult struct {
	types.DynamicData

	Hostname string                `xml:"hostname"`
	Version  string                `xml:"version,omitempty"`
	Error    types.BaseMethodFault `xml:"error,omitempty,typeattr"`
}

func init() {
	types.Add("vsan:VsanHostHealthSystemVersionResult", reflect.TypeOf((*VsanHostHealthSystemVersionResult)(nil)).Elem())
}

type VimHostVSANStretchedClusterHostCapability struct {
	types.DynamicData

	FeatureVersion string `xml:"featureVersion"`
}

func init() {
	types.Add("vsan:VimHostVSANStretchedClusterHostCapability", reflect.TypeOf((*VimHostVSANStretchedClusterHostCapability)(nil)).Elem())
}

type VsanNodeNotMaster struct {
	types.VimFault

	VsanMasterUuid               string `xml:"vsanMasterUuid,omitempty"`
	CmmdsMasterButNotStatsMaster *bool  `xml:"cmmdsMasterButNotStatsMaster"`
}

func init() {
	types.Add("vsan:VsanNodeNotMaster", reflect.TypeOf((*VsanNodeNotMaster)(nil)).Elem())
}

type HostSpbmHashInfo struct {
	types.DynamicData

	PolicyInfoHash    string `xml:"policyInfoHash"`
	DatastoreInfoHash string `xml:"datastoreInfoHash"`
}

func init() {
	types.Add("vsan:HostSpbmHashInfo", reflect.TypeOf((*HostSpbmHashInfo)(nil)).Elem())
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

type VsanFileServiceBalanceHealth struct {
	types.DynamicData

	Health      string `xml:"health,omitempty"`
	Description string `xml:"description,omitempty"`
}

func init() {
	types.Add("vsan:VsanFileServiceBalanceHealth", reflect.TypeOf((*VsanFileServiceBalanceHealth)(nil)).Elem())
}

type VsanClusterHealthQuerySpec struct {
	types.DynamicData

	Task *types.ManagedObjectReference `xml:"task,omitempty"`
}

func init() {
	types.Add("vsan:VsanClusterHealthQuerySpec", reflect.TypeOf((*VsanClusterHealthQuerySpec)(nil)).Elem())
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

type VsanHostQueryCheckLimitsSpec struct {
	types.DynamicData

	OptionTypes []string `xml:"optionTypes,omitempty"`
	FetchAll    bool     `xml:"fetchAll"`
}

func init() {
	types.Add("vsan:VsanHostQueryCheckLimitsSpec", reflect.TypeOf((*VsanHostQueryCheckLimitsSpec)(nil)).Elem())
}

type VsanCapability struct {
	types.DynamicData

	Target       *types.ManagedObjectReference `xml:"target,omitempty,typeattr"`
	Capabilities []string                      `xml:"capabilities,omitempty"`
	Statuses     []string                      `xml:"statuses,omitempty"`
}

func init() {
	types.Add("vsan:VsanCapability", reflect.TypeOf((*VsanCapability)(nil)).Elem())
}

type VsanFileShareConfig struct {
	types.DynamicData

	Name          string                           `xml:"name,omitempty"`
	DomainName    string                           `xml:"domainName,omitempty"`
	Quota         string                           `xml:"quota,omitempty"`
	SoftQuota     string                           `xml:"softQuota,omitempty"`
	Labels        []types.KeyValue                 `xml:"labels,omitempty"`
	StoragePolicy *types.VirtualMachineProfileSpec `xml:"storagePolicy,omitempty"`
	Permission    []VsanFileShareNetPermission     `xml:"permission,omitempty"`
	Protocols     []string                         `xml:"protocols,omitempty"`
	SmbOptions    *VsanFileShareSmbOptions         `xml:"smbOptions,omitempty"`
	NfsSecType    string                           `xml:"nfsSecType,omitempty"`
}

func init() {
	types.Add("vsan:VsanFileShareConfig", reflect.TypeOf((*VsanFileShareConfig)(nil)).Elem())
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

type VsanClusterDitEncryptionHealthSummary struct {
	types.DynamicData

	OverallHealth string                           `xml:"overallHealth"`
	Enabled       *bool                            `xml:"enabled"`
	HostResults   []VsanDitEncryptionHealthSummary `xml:"hostResults,omitempty"`
}

func init() {
	types.Add("vsan:VsanClusterDitEncryptionHealthSummary", reflect.TypeOf((*VsanClusterDitEncryptionHealthSummary)(nil)).Elem())
}

type VsanClientDatastoreConfig struct {
	VsanDatastoreSpec

	Clusters []types.ManagedObjectReference `xml:"clusters"`
}

func init() {
	types.Add("vsan:VsanClientDatastoreConfig", reflect.TypeOf((*VsanClientDatastoreConfig)(nil)).Elem())
}

type VsanHigherObjectsPresentDuringDowngradeIssue struct {
	types.VsanUpgradeSystemPreflightCheckIssue

	Uuids []string `xml:"uuids"`
}

func init() {
	types.Add("vsan:VsanHigherObjectsPresentDuringDowngradeIssue", reflect.TypeOf((*VsanHigherObjectsPresentDuringDowngradeIssue)(nil)).Elem())
}

type VsanUpgradeStatusEx struct {
	types.VsanUpgradeSystemUpgradeStatus

	IsPrecheck     *bool                                `xml:"isPrecheck"`
	PrecheckResult *VsanDiskFormatConversionCheckResult `xml:"precheckResult,omitempty"`
}

func init() {
	types.Add("vsan:VsanUpgradeStatusEx", reflect.TypeOf((*VsanUpgradeStatusEx)(nil)).Elem())
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

type VsanComparator struct {
	types.DynamicData
}

func init() {
	types.Add("vsan:VsanComparator", reflect.TypeOf((*VsanComparator)(nil)).Elem())
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

type VsanHostsCompressionOnlyNotSupported struct {
	types.VsanUpgradeSystemPreflightCheckIssue

	Hosts []types.ManagedObjectReference `xml:"hosts"`
}

func init() {
	types.Add("vsan:VsanHostsCompressionOnlyNotSupported", reflect.TypeOf((*VsanHostsCompressionOnlyNotSupported)(nil)).Elem())
}

type VsanGenericClusterBaseIssue struct {
	types.DynamicData
}

func init() {
	types.Add("vsan:VsanGenericClusterBaseIssue", reflect.TypeOf((*VsanGenericClusterBaseIssue)(nil)).Elem())
}

type VsanObjectQuerySpec struct {
	types.DynamicData

	Uuid                    string `xml:"uuid"`
	SpbmProfileGenerationId string `xml:"spbmProfileGenerationId,omitempty"`
}

func init() {
	types.Add("vsan:VsanObjectQuerySpec", reflect.TypeOf((*VsanObjectQuerySpec)(nil)).Elem())
}

type VsanObjectPolicyIssue struct {
	types.VsanUpgradeSystemPreflightCheckIssue

	Uuids []string `xml:"uuids"`
}

func init() {
	types.Add("vsan:VsanObjectPolicyIssue", reflect.TypeOf((*VsanObjectPolicyIssue)(nil)).Elem())
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

type VsanHostPortConfigEx struct {
	types.VsanHostConfigInfoNetworkInfoPortConfig

	TrafficTypes []string `xml:"trafficTypes,omitempty"`
}

func init() {
	types.Add("vsan:VsanHostPortConfigEx", reflect.TypeOf((*VsanHostPortConfigEx)(nil)).Elem())
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

type VsanIscsiHomeObjectSpec struct {
	types.DynamicData

	StoragePolicy *types.VirtualMachineProfileSpec         `xml:"storagePolicy,omitempty"`
	DefaultConfig *VsanIscsiTargetServiceDefaultConfigSpec `xml:"defaultConfig,omitempty"`
}

func init() {
	types.Add("vsan:VsanIscsiHomeObjectSpec", reflect.TypeOf((*VsanIscsiHomeObjectSpec)(nil)).Elem())
}

type VsanMountPrecheckNetworkConnectivityDetail struct {
	types.DynamicData

	Host                types.ManagedObjectReference           `xml:"host"`
	NetworkConnectivity []VsanMountPrecheckNetworkConnectivity `xml:"networkConnectivity,omitempty"`
}

func init() {
	types.Add("vsan:VsanMountPrecheckNetworkConnectivityDetail", reflect.TypeOf((*VsanMountPrecheckNetworkConnectivityDetail)(nil)).Elem())
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

type VsanHostConfigInfoEx struct {
	types.VsanHostConfigInfo

	EncryptionInfo              *VsanHostEncryptionInfo             `xml:"encryptionInfo,omitempty"`
	DataEfficiencyInfo          *VsanDataEfficiencyConfig           `xml:"dataEfficiencyInfo,omitempty"`
	ResyncIopsLimitInfo         *ResyncIopsInfo                     `xml:"resyncIopsLimitInfo,omitempty"`
	ExtendedConfig              *VsanExtendedConfig                 `xml:"extendedConfig,omitempty"`
	DatastoreInfo               BaseVsanDatastoreConfig             `xml:"datastoreInfo,omitempty,typeattr"`
	UnmapConfig                 *VsanUnmapConfig                    `xml:"unmapConfig,omitempty"`
	WitnessHostConfig           []VsanWitnessHostConfig             `xml:"witnessHostConfig,omitempty"`
	InternalExtendedConfig      *VsanInternalExtendedConfig         `xml:"internalExtendedConfig,omitempty"`
	MetricsConfig               *VsanMetricsConfig                  `xml:"metricsConfig,omitempty"`
	UnicastConfig               *VsanHostServerClusterUnicastConfig `xml:"unicastConfig,omitempty"`
	DataInTransitEncryptionInfo *VsanInTransitEncryptionInfo        `xml:"dataInTransitEncryptionInfo,omitempty"`
}

func init() {
	types.Add("vsan:VsanHostConfigInfoEx", reflect.TypeOf((*VsanHostConfigInfoEx)(nil)).Elem())
}

type VimClusterVSANPreferredFaultDomainInfo struct {
	types.DynamicData

	PreferredFaultDomainName string `xml:"preferredFaultDomainName,omitempty"`
	PreferredFaultDomainId   string `xml:"preferredFaultDomainId,omitempty"`
}

func init() {
	types.Add("vsan:VimClusterVSANPreferredFaultDomainInfo", reflect.TypeOf((*VimClusterVSANPreferredFaultDomainInfo)(nil)).Elem())
}

type VsanVsanPcapResult struct {
	types.DynamicData

	Calltime      float32               `xml:"calltime"`
	Vmknic        string                `xml:"vmknic"`
	TcpdumpFilter string                `xml:"tcpdumpFilter"`
	Snaplen       int32                 `xml:"snaplen"`
	Pkts          []string              `xml:"pkts,omitempty"`
	Pcap          string                `xml:"pcap,omitempty"`
	Error         types.BaseMethodFault `xml:"error,omitempty,typeattr"`
	Hostname      string                `xml:"hostname,omitempty"`
}

func init() {
	types.Add("vsan:VsanVsanPcapResult", reflect.TypeOf((*VsanVsanPcapResult)(nil)).Elem())
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

type VsanHostPropertyRetrieveIssue struct {
	types.VsanUpgradeSystemPreflightCheckIssue

	Hosts []types.ManagedObjectReference `xml:"hosts"`
}

func init() {
	types.Add("vsan:VsanHostPropertyRetrieveIssue", reflect.TypeOf((*VsanHostPropertyRetrieveIssue)(nil)).Elem())
}

type VimVsanDataEfficiencyCapacityState struct {
	types.DynamicData

	LogicalCapacity             int64                            `xml:"logicalCapacity,omitempty"`
	LogicalCapacityUsed         int64                            `xml:"logicalCapacityUsed,omitempty"`
	PhysicalCapacity            int64                            `xml:"physicalCapacity,omitempty"`
	PhysicalCapacityUsed        int64                            `xml:"physicalCapacityUsed,omitempty"`
	DedupMetadataSize           int64                            `xml:"dedupMetadataSize,omitempty"`
	SpaceEfficiencyMetadataSize *VsanSpaceEfficiencyMetadataSize `xml:"spaceEfficiencyMetadataSize,omitempty"`
}

func init() {
	types.Add("vsan:VimVsanDataEfficiencyCapacityState", reflect.TypeOf((*VimVsanDataEfficiencyCapacityState)(nil)).Elem())
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

type VimClusterVsanFaultDomainSpec struct {
	types.DynamicData

	Hosts []types.ManagedObjectReference `xml:"hosts"`
	Name  string                         `xml:"name"`
}

func init() {
	types.Add("vsan:VimClusterVsanFaultDomainSpec", reflect.TypeOf((*VimClusterVsanFaultDomainSpec)(nil)).Elem())
}

type VsanExtendedConfig struct {
	types.DynamicData

	ObjectRepairTimer          int64                        `xml:"objectRepairTimer,omitempty"`
	DisableSiteReadLocality    *bool                        `xml:"disableSiteReadLocality"`
	EnableCustomizedSwapObject *bool                        `xml:"enableCustomizedSwapObject"`
	LargeScaleClusterSupport   *bool                        `xml:"largeScaleClusterSupport"`
	ProactiveRebalanceInfo     *VsanProactiveRebalanceInfo  `xml:"proactiveRebalanceInfo,omitempty"`
	CapacityReservationInfo    *VsanCapacityReservationInfo `xml:"capacityReservationInfo,omitempty"`
}

func init() {
	types.Add("vsan:VsanExtendedConfig", reflect.TypeOf((*VsanExtendedConfig)(nil)).Elem())
}

type VsanUnsupportedHighDiskVersionIssue struct {
	types.VsanUpgradeSystemPreflightCheckIssue

	Hosts []types.ManagedObjectReference `xml:"hosts"`
}

func init() {
	types.Add("vsan:VsanUnsupportedHighDiskVersionIssue", reflect.TypeOf((*VsanUnsupportedHighDiskVersionIssue)(nil)).Elem())
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
	ObjTypeExt         string `xml:"objTypeExt,omitempty"`
	ObjTypeExtDesc     string `xml:"objTypeExtDesc,omitempty"`
}

func init() {
	types.Add("vsan:VsanObjectSpaceSummary", reflect.TypeOf((*VsanObjectSpaceSummary)(nil)).Elem())
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

type VsanHostServerClusterUnicastInfo struct {
	types.DynamicData

	ClusterUuid string                      `xml:"clusterUuid"`
	UnicastInfo []VsanServerHostUnicastInfo `xml:"unicastInfo,omitempty"`
}

func init() {
	types.Add("vsan:VsanHostServerClusterUnicastInfo", reflect.TypeOf((*VsanHostServerClusterUnicastInfo)(nil)).Elem())
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

type VsanUnicastAddressInfo struct {
	types.DynamicData

	Address string `xml:"address"`
	Port    int32  `xml:"port,omitempty"`
}

func init() {
	types.Add("vsan:VsanUnicastAddressInfo", reflect.TypeOf((*VsanUnicastAddressInfo)(nil)).Elem())
}

type VsanHealthQuerySpec struct {
	types.DynamicData

	IncludeAllRemoteClusters *bool    `xml:"includeAllRemoteClusters"`
	RemoteClusterUuids       []string `xml:"remoteClusterUuids,omitempty"`
	LatencyOnly              *bool    `xml:"latencyOnly"`
}

func init() {
	types.Add("vsan:VsanHealthQuerySpec", reflect.TypeOf((*VsanHealthQuerySpec)(nil)).Elem())
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

type VsanStoragePolicyStatus struct {
	types.DynamicData

	Id            string `xml:"id,omitempty"`
	ExpectedValue string `xml:"expectedValue,omitempty"`
	CurrentValue  string `xml:"currentValue,omitempty"`
}

func init() {
	types.Add("vsan:VsanStoragePolicyStatus", reflect.TypeOf((*VsanStoragePolicyStatus)(nil)).Elem())
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

type VsanFileServiceIpConfig struct {
	types.HostIpConfig

	Fqdn      string `xml:"fqdn,omitempty"`
	IsPrimary *bool  `xml:"isPrimary"`
	Gateway   string `xml:"gateway"`
}

func init() {
	types.Add("vsan:VsanFileServiceIpConfig", reflect.TypeOf((*VsanFileServiceIpConfig)(nil)).Elem())
}

type VsanMetricProfile struct {
	types.DynamicData

	AuthToken string `xml:"authToken"`
}

func init() {
	types.Add("vsan:VsanMetricProfile", reflect.TypeOf((*VsanMetricProfile)(nil)).Elem())
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

type VsanVnicVdsMigrationSpec struct {
	types.DynamicData

	Key        int32                          `xml:"key"`
	VdsBacking types.VirtualDeviceBackingInfo `xml:"vdsBacking"`
}

func init() {
	types.Add("vsan:VsanVnicVdsMigrationSpec", reflect.TypeOf((*VsanVnicVdsMigrationSpec)(nil)).Elem())
}

type VsanFaultDomainResourceCheckResult struct {
	EntityResourceCheckDetails

	Hosts []VsanHostResourceCheckResult `xml:"hosts,omitempty"`
}

func init() {
	types.Add("vsan:VsanFaultDomainResourceCheckResult", reflect.TypeOf((*VsanFaultDomainResourceCheckResult)(nil)).Elem())
}

type VsanVcKmipServersHealth struct {
	types.DynamicData

	Health               string                `xml:"health,omitempty"`
	Error                types.BaseMethodFault `xml:"error,omitempty,typeattr"`
	KmsProviderId        string                `xml:"kmsProviderId,omitempty"`
	KmsHealth            []VsanKmsHealth       `xml:"kmsHealth,omitempty"`
	ClientCertHealth     string                `xml:"clientCertHealth,omitempty"`
	ClientCertExpireDate *time.Time            `xml:"clientCertExpireDate"`
	IsAwsKms             *bool                 `xml:"isAwsKms"`
	CmkHealth            string                `xml:"cmkHealth,omitempty"`
}

func init() {
	types.Add("vsan:VsanVcKmipServersHealth", reflect.TypeOf((*VsanVcKmipServersHealth)(nil)).Elem())
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
	FileServiceVersion    string     `xml:"fileServiceVersion,omitempty"`
	DvsConfigIssue        string     `xml:"dvsConfigIssue,omitempty"`
}

func init() {
	types.Add("vsan:VsanFileServicePreflightCheckResult", reflect.TypeOf((*VsanFileServicePreflightCheckResult)(nil)).Elem())
}

type VimVsanHostDiskMappingCreationSpec struct {
	types.DynamicData

	Host          types.ManagedObjectReference `xml:"host"`
	CacheDisks    []types.HostScsiDisk         `xml:"cacheDisks,omitempty"`
	CapacityDisks []types.HostScsiDisk         `xml:"capacityDisks,omitempty"`
	CreationType  string                       `xml:"creationType"`
}

func init() {
	types.Add("vsan:VimVsanHostDiskMappingCreationSpec", reflect.TypeOf((*VimVsanHostDiskMappingCreationSpec)(nil)).Elem())
}

type VsanHostWipeDiskStatus struct {
	types.DynamicData

	Disk                string                     `xml:"disk"`
	Eligible            string                     `xml:"eligible"`
	IneligibleReason    []types.LocalizableMessage `xml:"ineligibleReason,omitempty"`
	WipeState           string                     `xml:"wipeState,omitempty"`
	PercentageCompleted int32                      `xml:"percentageCompleted,omitempty"`
	EstimatedTime       int64                      `xml:"estimatedTime,omitempty"`
	WipeStartTime       *time.Time                 `xml:"wipeStartTime"`
	WipeCompleteTime    *time.Time                 `xml:"wipeCompleteTime"`
}

func init() {
	types.Add("vsan:VsanHostWipeDiskStatus", reflect.TypeOf((*VsanHostWipeDiskStatus)(nil)).Elem())
}

type VsanIscsiTargetServiceConfig struct {
	types.DynamicData

	DefaultConfig *VsanIscsiTargetServiceDefaultConfigSpec `xml:"defaultConfig,omitempty"`
	Enabled       *bool                                    `xml:"enabled"`
}

func init() {
	types.Add("vsan:VsanIscsiTargetServiceConfig", reflect.TypeOf((*VsanIscsiTargetServiceConfig)(nil)).Elem())
}

type VsanHostAbortWipeDiskStatus struct {
	types.DynamicData

	Disk    string                     `xml:"disk"`
	Success bool                       `xml:"success"`
	Reason  []types.LocalizableMessage `xml:"reason,omitempty"`
}

func init() {
	types.Add("vsan:VsanHostAbortWipeDiskStatus", reflect.TypeOf((*VsanHostAbortWipeDiskStatus)(nil)).Elem())
}

type VsanHostVsanObjectSyncState struct {
	types.DynamicData

	Uuid       string                       `xml:"uuid"`
	Components []VsanHostComponentSyncState `xml:"components"`
}

func init() {
	types.Add("vsan:VsanHostVsanObjectSyncState", reflect.TypeOf((*VsanHostVsanObjectSyncState)(nil)).Elem())
}

type VimVsanHostVsanDiskManagementSystemCapability struct {
	types.DynamicData

	Version string `xml:"version"`
}

func init() {
	types.Add("vsan:VimVsanHostVsanDiskManagementSystemCapability", reflect.TypeOf((*VimVsanHostVsanDiskManagementSystemCapability)(nil)).Elem())
}

type VimVsanHostVsanScsiDisk struct {
	types.DynamicData

	Capacity         types.HostDiskDimensionsLba `xml:"capacity"`
	UsedCapacity     int64                       `xml:"usedCapacity,omitempty"`
	DevicePath       string                      `xml:"devicePath"`
	Ssd              *bool                       `xml:"ssd"`
	LocalDisk        *bool                       `xml:"localDisk"`
	ScsiDiskType     string                      `xml:"scsiDiskType,omitempty"`
	Uuid             string                      `xml:"uuid"`
	OperationalState []string                    `xml:"operationalState,omitempty"`
	CanonicalName    string                      `xml:"canonicalName,omitempty"`
	DisplayName      string                      `xml:"displayName,omitempty"`
	LunType          string                      `xml:"lunType"`
	Vendor           string                      `xml:"vendor,omitempty"`
	Model            string                      `xml:"model,omitempty"`
	MountInfo        *types.HostMountInfo        `xml:"mountInfo,omitempty"`
}

func init() {
	types.Add("vsan:VimVsanHostVsanScsiDisk", reflect.TypeOf((*VimVsanHostVsanScsiDisk)(nil)).Elem())
}

type VsanObjectOverallHealth struct {
	types.DynamicData

	ObjectHealthDetail              []VsanObjectHealth            `xml:"objectHealthDetail,omitempty"`
	ObjectsComplianceDetail         []VsanStorageComplianceResult `xml:"objectsComplianceDetail,omitempty"`
	ObjectVersionCompliance         *bool                         `xml:"objectVersionCompliance"`
	ObjectFormatChangeRequiredUuids []string                      `xml:"objectFormatChangeRequiredUuids,omitempty"`
	ObjectsRelayoutBytes            int64                         `xml:"objectsRelayoutBytes,omitempty"`
}

func init() {
	types.Add("vsan:VsanObjectOverallHealth", reflect.TypeOf((*VsanObjectOverallHealth)(nil)).Elem())
}

type VsanClusterHealthResultColumnInfo struct {
	types.DynamicData

	Label string `xml:"label"`
	Type  string `xml:"type"`
}

func init() {
	types.Add("vsan:VsanClusterHealthResultColumnInfo", reflect.TypeOf((*VsanClusterHealthResultColumnInfo)(nil)).Elem())
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

type HostSpbmPolicyInfo struct {
	types.DynamicData

	ProfileId      string                   `xml:"profileId"`
	Name           string                   `xml:"name"`
	Description    string                   `xml:"description,omitempty"`
	GenerationId   int64                    `xml:"generationId"`
	PolicyBlobInfo []HostSpbmPolicyBlobInfo `xml:"policyBlobInfo"`
}

func init() {
	types.Add("vsan:HostSpbmPolicyInfo", reflect.TypeOf((*HostSpbmPolicyInfo)(nil)).Elem())
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

type VsanVumConfig struct {
	types.DynamicData

	BaselinePreferenceType string `xml:"baselinePreferenceType"`
}

func init() {
	types.Add("vsan:VsanVumConfig", reflect.TypeOf((*VsanVumConfig)(nil)).Elem())
}

type VsanFileServiceDomain struct {
	types.DynamicData

	Uuid   string                       `xml:"uuid"`
	Config *VsanFileServiceDomainConfig `xml:"config,omitempty"`
}

func init() {
	types.Add("vsan:VsanFileServiceDomain", reflect.TypeOf((*VsanFileServiceDomain)(nil)).Elem())
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

type VsanCloudHealthStatus struct {
	types.DynamicData

	CollectorRunning     *bool  `xml:"collectorRunning"`
	LastSentTimestamp    string `xml:"lastSentTimestamp,omitempty"`
	InternetConnectivity *bool  `xml:"internetConnectivity"`
}

func init() {
	types.Add("vsan:VsanCloudHealthStatus", reflect.TypeOf((*VsanCloudHealthStatus)(nil)).Elem())
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

type VsanIoInsightInstance struct {
	types.DynamicData

	RunName            string                  `xml:"runName"`
	State              string                  `xml:"state,omitempty"`
	StartTime          *time.Time              `xml:"startTime"`
	EndTime            *time.Time              `xml:"endTime"`
	HostsIoInsightInfo []VsanHostIoInsightInfo `xml:"hostsIoInsightInfo,omitempty"`
	HostUuids          []string                `xml:"hostUuids,omitempty"`
	VmUuids            []string                `xml:"vmUuids,omitempty"`
}

func init() {
	types.Add("vsan:VsanIoInsightInstance", reflect.TypeOf((*VsanIoInsightInstance)(nil)).Elem())
}

type VsanNestJsonComparator struct {
	VsanComparator

	NestedComparators []VsanJsonComparator `xml:"nestedComparators,omitempty"`
	Conjoiner         string               `xml:"conjoiner,omitempty"`
}

func init() {
	types.Add("vsan:VsanNestJsonComparator", reflect.TypeOf((*VsanNestJsonComparator)(nil)).Elem())
}

type VsanIscsiTargetSpec struct {
	VsanIscsiTargetCommonInfo

	StoragePolicy *types.VirtualMachineProfileSpec `xml:"storagePolicy,omitempty"`
	NewAlias      string                           `xml:"newAlias,omitempty"`
}

func init() {
	types.Add("vsan:VsanIscsiTargetSpec", reflect.TypeOf((*VsanIscsiTargetSpec)(nil)).Elem())
}

type VsanFileShareQueryProperties struct {
	types.DynamicData

	IncludeBasic           *bool    `xml:"includeBasic"`
	IncludeUsedCapacity    *bool    `xml:"includeUsedCapacity"`
	IncludeVsanObjectUuids *bool    `xml:"includeVsanObjectUuids"`
	IncludeAllLabels       *bool    `xml:"includeAllLabels"`
	LabelKeys              []string `xml:"labelKeys,omitempty"`
}

func init() {
	types.Add("vsan:VsanFileShareQueryProperties", reflect.TypeOf((*VsanFileShareQueryProperties)(nil)).Elem())
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

type VSANSharedWitnessCompatibilityResult struct {
	types.DynamicData

	WitnessHostCompatibility VSANEntityCompatibilityResult   `xml:"witnessHostCompatibility"`
	RoboClusterCompatibility []VSANEntityCompatibilityResult `xml:"roboClusterCompatibility,omitempty"`
}

func init() {
	types.Add("vsan:VSANSharedWitnessCompatibilityResult", reflect.TypeOf((*VSANSharedWitnessCompatibilityResult)(nil)).Elem())
}

type VsanVcStretchedClusterConfigSpec struct {
	types.DynamicData

	WitnessHost         types.ManagedObjectReference `xml:"witnessHost"`
	Clusters            []VsanStretchedClusterConfig `xml:"clusters"`
	WitnessDiskMappings []types.VsanHostDiskMapping  `xml:"witnessDiskMappings,omitempty"`
}

func init() {
	types.Add("vsan:VsanVcStretchedClusterConfigSpec", reflect.TypeOf((*VsanVcStretchedClusterConfigSpec)(nil)).Elem())
}

type VsanFileServiceShareHealthSummary struct {
	types.DynamicData

	OverallHealth string                   `xml:"overallHealth,omitempty"`
	DomainName    string                   `xml:"domainName,omitempty"`
	ShareUuid     string                   `xml:"shareUuid,omitempty"`
	ShareName     string                   `xml:"shareName,omitempty"`
	ObjectHealth  *VsanObjectOverallHealth `xml:"objectHealth,omitempty"`
	Description   string                   `xml:"description,omitempty"`
	Extensible    *bool                    `xml:"extensible"`
}

func init() {
	types.Add("vsan:VsanFileServiceShareHealthSummary", reflect.TypeOf((*VsanFileServiceShareHealthSummary)(nil)).Elem())
}

type VsanDataInTransitEncryptionConfig struct {
	types.DynamicData

	Enabled       *bool `xml:"enabled"`
	RekeyInterval int32 `xml:"rekeyInterval,omitempty"`
}

func init() {
	types.Add("vsan:VsanDataInTransitEncryptionConfig", reflect.TypeOf((*VsanDataInTransitEncryptionConfig)(nil)).Elem())
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

type VsanDitEncryptionHealthSummary struct {
	types.DynamicData

	Hostname          string                       `xml:"hostname,omitempty"`
	Health            string                       `xml:"health,omitempty"`
	Reason            *types.LocalizableMessage    `xml:"reason,omitempty"`
	DitEncryptionInfo *VsanInTransitEncryptionInfo `xml:"ditEncryptionInfo,omitempty"`
}

func init() {
	types.Add("vsan:VsanDitEncryptionHealthSummary", reflect.TypeOf((*VsanDitEncryptionHealthSummary)(nil)).Elem())
}

type VsanMountPrecheckNetworkConnectivity struct {
	types.DynamicData

	Host                    types.ManagedObjectReference `xml:"host"`
	SmallPingTestSuccessPct int32                        `xml:"smallPingTestSuccessPct"`
	LargePingTestSuccessPct int32                        `xml:"largePingTestSuccessPct"`
	Status                  string                       `xml:"status"`
}

func init() {
	types.Add("vsan:VsanMountPrecheckNetworkConnectivity", reflect.TypeOf((*VsanMountPrecheckNetworkConnectivity)(nil)).Elem())
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

type VsanClusterWhatifHostFailuresResult struct {
	types.DynamicData

	NumFailures             int64                        `xml:"numFailures"`
	TotalUsedCapacityB      int64                        `xml:"totalUsedCapacityB"`
	TotalCapacityB          int64                        `xml:"totalCapacityB"`
	TotalRcReservationB     int64                        `xml:"totalRcReservationB"`
	TotalRcSizeB            int64                        `xml:"totalRcSizeB"`
	UsedComponents          int64                        `xml:"usedComponents"`
	TotalComponents         int64                        `xml:"totalComponents"`
	ComponentLimitHealth    string                       `xml:"componentLimitHealth,omitempty"`
	DiskFreeSpaceHealth     string                       `xml:"diskFreeSpaceHealth,omitempty"`
	RcFreeReservationHealth string                       `xml:"rcFreeReservationHealth,omitempty"`
	SlackSpaceCapRequired   int64                        `xml:"slackSpaceCapRequired,omitempty"`
	DiskSpaceThreshold      *VsanHealthThreshold         `xml:"diskSpaceThreshold,omitempty"`
	CapacityReservationInfo *VsanCapacityReservationInfo `xml:"capacityReservationInfo,omitempty"`
}

func init() {
	types.Add("vsan:VsanClusterWhatifHostFailuresResult", reflect.TypeOf((*VsanClusterWhatifHostFailuresResult)(nil)).Elem())
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

type VsanClusterHealthResultTable struct {
	VsanClusterHealthResultBase

	Columns []VsanClusterHealthResultColumnInfo `xml:"columns,omitempty"`
	Rows    []VsanClusterHealthResultRow        `xml:"rows,omitempty"`
}

func init() {
	types.Add("vsan:VsanClusterHealthResultTable", reflect.TypeOf((*VsanClusterHealthResultTable)(nil)).Elem())
}

type VsanClusterNetworkPerfTaskSpec struct {
	types.DynamicData

	Cluster     *types.ManagedObjectReference `xml:"Cluster,omitempty"`
	DurationSec int32                         `xml:"DurationSec,omitempty"`
}

func init() {
	types.Add("vsan:VsanClusterNetworkPerfTaskSpec", reflect.TypeOf((*VsanClusterNetworkPerfTaskSpec)(nil)).Elem())
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
	BalanceStatus    *VsanFileServiceBalanceHealth       `xml:"balanceStatus,omitempty"`
}

func init() {
	types.Add("vsan:VsanFileServiceHealthSummary", reflect.TypeOf((*VsanFileServiceHealthSummary)(nil)).Elem())
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

type VimVsanHostVsanHostCapability struct {
	types.DynamicData

	Host        types.ManagedObjectReference `xml:"host"`
	IsSupported bool                         `xml:"isSupported"`
	IsLicensed  bool                         `xml:"isLicensed"`
}

func init() {
	types.Add("vsan:VimVsanHostVsanHostCapability", reflect.TypeOf((*VimVsanHostVsanHostCapability)(nil)).Elem())
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

type VsanStretchedClusterConfig struct {
	types.DynamicData

	Cluster           types.ManagedObjectReference                     `xml:"cluster"`
	PreferredFdName   string                                           `xml:"preferredFdName,omitempty"`
	FaultDomainConfig *VimClusterVSANStretchedClusterFaultDomainConfig `xml:"faultDomainConfig,omitempty"`
}

func init() {
	types.Add("vsan:VsanStretchedClusterConfig", reflect.TypeOf((*VsanStretchedClusterConfig)(nil)).Elem())
}

type VsanClusterVMsHealthOverallResult struct {
	types.DynamicData

	HealthStateList    []VsanClusterVMsHealthSummaryResult `xml:"healthStateList,omitempty"`
	OverallHealthState string                              `xml:"overallHealthState,omitempty"`
}

func init() {
	types.Add("vsan:VsanClusterVMsHealthOverallResult", reflect.TypeOf((*VsanClusterVMsHealthOverallResult)(nil)).Elem())
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

type VimVsanHostDiskMapInfoEx struct {
	types.DynamicData

	Mapping              types.VsanHostDiskMapping `xml:"mapping"`
	IsMounted            bool                      `xml:"isMounted"`
	UnlockedEncrypted    *bool                     `xml:"unlockedEncrypted"`
	IsAllFlash           bool                      `xml:"isAllFlash"`
	IsDataEfficiency     *bool                     `xml:"isDataEfficiency"`
	EncryptionInfo       *VsanDataEncryptionConfig `xml:"encryptionInfo,omitempty"`
	DataEfficiencyConfig *VsanDataEfficiencyConfig `xml:"dataEfficiencyConfig,omitempty"`
}

func init() {
	types.Add("vsan:VimVsanHostDiskMapInfoEx", reflect.TypeOf((*VimVsanHostDiskMapInfoEx)(nil)).Elem())
}

type VsanClusterHealthSummary struct {
	types.DynamicData

	ClusterStatus            *VsanClusterHealthSystemStatusResult   `xml:"clusterStatus,omitempty"`
	Timestamp                *time.Time                             `xml:"timestamp"`
	ClusterVersions          *VsanClusterHealthSystemVersionResult  `xml:"clusterVersions,omitempty"`
	ObjectHealth             *VsanObjectOverallHealth               `xml:"objectHealth,omitempty"`
	VmHealth                 *VsanClusterVMsHealthOverallResult     `xml:"vmHealth,omitempty"`
	NetworkHealth            *VsanClusterNetworkHealthResult        `xml:"networkHealth,omitempty"`
	LimitHealth              *VsanClusterLimitHealthResult          `xml:"limitHealth,omitempty"`
	AdvCfgSync               []VsanClusterAdvCfgSyncResult          `xml:"advCfgSync,omitempty"`
	CreateVmHealth           []VsanHostCreateVmHealthTestResult     `xml:"createVmHealth,omitempty"`
	PhysicalDisksHealth      []VsanPhysicalDiskHealthSummary        `xml:"physicalDisksHealth,omitempty"`
	EncryptionHealth         *VsanClusterEncryptionHealthSummary    `xml:"encryptionHealth,omitempty"`
	HclInfo                  *VsanClusterHclInfo                    `xml:"hclInfo,omitempty"`
	Groups                   []VsanClusterHealthGroup               `xml:"groups,omitempty"`
	OverallHealth            string                                 `xml:"overallHealth"`
	OverallHealthDescription string                                 `xml:"overallHealthDescription"`
	ClomdLiveness            *VsanClusterClomdLivenessResult        `xml:"clomdLiveness,omitempty"`
	DiskBalance              *VsanClusterBalanceSummary             `xml:"diskBalance,omitempty"`
	GenericCluster           *VsanGenericClusterBestPracticeHealth  `xml:"genericCluster,omitempty"`
	NetworkConfig            *VsanNetworkConfigBestPracticeHealth   `xml:"networkConfig,omitempty"`
	VsanConfig               *VsanConfigCheckResult                 `xml:"vsanConfig,omitempty"`
	BurnInTest               *VsanBurnInTestCheckResult             `xml:"burnInTest,omitempty"`
	PerfsvcHealth            *VsanPerfsvcHealthResult               `xml:"perfsvcHealth,omitempty"`
	Cluster                  *types.ManagedObjectReference          `xml:"cluster,omitempty"`
	FileServiceHealth        *VsanClusterFileServiceHealthSummary   `xml:"fileServiceHealth,omitempty"`
	DitEncryptionHealth      *VsanClusterDitEncryptionHealthSummary `xml:"ditEncryptionHealth,omitempty"`
}

func init() {
	types.Add("vsan:VsanClusterHealthSummary", reflect.TypeOf((*VsanClusterHealthSummary)(nil)).Elem())
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

type VsanPhysicalDiskHealthSummary struct {
	types.DynamicData

	OverallHealth        string                   `xml:"overallHealth"`
	HeapsWithIssues      []VsanResourceHealth     `xml:"heapsWithIssues,omitempty"`
	SlabsWithIssues      []VsanResourceHealth     `xml:"slabsWithIssues,omitempty"`
	Disks                []VsanPhysicalDiskHealth `xml:"disks,omitempty"`
	ComponentsWithIssues []VsanResourceHealth     `xml:"componentsWithIssues,omitempty"`
	Hostname             string                   `xml:"hostname,omitempty"`
	HostDedupScope       int32                    `xml:"hostDedupScope,omitempty"`
	Error                types.BaseMethodFault    `xml:"error,omitempty,typeattr"`
}

func init() {
	types.Add("vsan:VsanPhysicalDiskHealthSummary", reflect.TypeOf((*VsanPhysicalDiskHealthSummary)(nil)).Elem())
}

type VsanNetworkConfigVsanNotOnVdsIssue struct {
	VsanNetworkConfigBaseIssue

	Host   types.ManagedObjectReference `xml:"host"`
	Vmknic string                       `xml:"vmknic"`
}

func init() {
	types.Add("vsan:VsanNetworkConfigVsanNotOnVdsIssue", reflect.TypeOf((*VsanNetworkConfigVsanNotOnVdsIssue)(nil)).Elem())
}

type VsanHostResourceCheckResult struct {
	EntityResourceCheckDetails

	Host       *types.ManagedObjectReference      `xml:"host,omitempty"`
	DiskGroups []VsanDiskGroupResourceCheckResult `xml:"diskGroups,omitempty"`
}

func init() {
	types.Add("vsan:VsanHostResourceCheckResult", reflect.TypeOf((*VsanHostResourceCheckResult)(nil)).Elem())
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

type VsanPerfTopEntity struct {
	types.DynamicData

	EntityRefId string `xml:"entityRefId"`
	Value       string `xml:"value"`
}

func init() {
	types.Add("vsan:VsanPerfTopEntity", reflect.TypeOf((*VsanPerfTopEntity)(nil)).Elem())
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

type VsanHostServerClusterUnicastConfig struct {
	types.DynamicData

	RemoteUnicastConfig []VsanHostServerClusterUnicastInfo `xml:"remoteUnicastConfig,omitempty"`
}

func init() {
	types.Add("vsan:VsanHostServerClusterUnicastConfig", reflect.TypeOf((*VsanHostServerClusterUnicastConfig)(nil)).Elem())
}

type VimVsanHostVsanManagedDisksInfo struct {
	types.DynamicData

	VSANDirectDisks []VimVsanHostVsanDirectStorage `xml:"vSANDirectDisks,omitempty"`
	VSANDiskMapInfo []VimVsanHostDiskMapInfoEx     `xml:"vSANDiskMapInfo,omitempty"`
}

func init() {
	types.Add("vsan:VimVsanHostVsanManagedDisksInfo", reflect.TypeOf((*VimVsanHostVsanManagedDisksInfo)(nil)).Elem())
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

type VsanFailedRepairObjectResult struct {
	types.DynamicData

	Uuid       string `xml:"uuid"`
	ErrMessage string `xml:"errMessage,omitempty"`
}

func init() {
	types.Add("vsan:VsanFailedRepairObjectResult", reflect.TypeOf((*VsanFailedRepairObjectResult)(nil)).Elem())
}

type VsanClusterCreateVmHealthTestResult struct {
	types.DynamicData

	ClusterResult VsanClusterProactiveTestResult     `xml:"clusterResult"`
	HostResults   []VsanHostCreateVmHealthTestResult `xml:"hostResults,omitempty"`
}

func init() {
	types.Add("vsan:VsanClusterCreateVmHealthTestResult", reflect.TypeOf((*VsanClusterCreateVmHealthTestResult)(nil)).Elem())
}

type VsanNetworkConfigBestPracticeHealth struct {
	types.DynamicData

	VdsPresent bool                             `xml:"vdsPresent"`
	Issues     []BaseVsanNetworkConfigBaseIssue `xml:"issues,omitempty,typeattr"`
}

func init() {
	types.Add("vsan:VsanNetworkConfigBestPracticeHealth", reflect.TypeOf((*VsanNetworkConfigBestPracticeHealth)(nil)).Elem())
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

type VsanMountPrecheckNetworkLatencyDetail struct {
	types.DynamicData

	Host             types.ManagedObjectReference      `xml:"host"`
	NetworkLatencies []VsanMountPrecheckNetworkLatency `xml:"networkLatencies,omitempty"`
}

func init() {
	types.Add("vsan:VsanMountPrecheckNetworkLatencyDetail", reflect.TypeOf((*VsanMountPrecheckNetworkLatencyDetail)(nil)).Elem())
}

type VsanSpaceQuerySpec struct {
	types.DynamicData

	EntityType string   `xml:"entityType"`
	EntityIds  []string `xml:"entityIds,omitempty"`
}

func init() {
	types.Add("vsan:VsanSpaceQuerySpec", reflect.TypeOf((*VsanSpaceQuerySpec)(nil)).Elem())
}

type VsanCompliantFirmware struct {
	types.DynamicData

	FirmwareVersion  string                `xml:"firmwareVersion"`
	CompliantDrivers []VsanCompliantDriver `xml:"compliantDrivers"`
}

func init() {
	types.Add("vsan:VsanCompliantFirmware", reflect.TypeOf((*VsanCompliantFirmware)(nil)).Elem())
}

type VSANEntityCompatibilityResult struct {
	types.DynamicData

	Entity              types.ManagedObjectReference `xml:"entity,typeattr"`
	Compatible          bool                         `xml:"compatible"`
	IncompatibleReasons []types.LocalizableMessage   `xml:"incompatibleReasons,omitempty"`
	ExtendedAttributes  []types.KeyAnyValue          `xml:"extendedAttributes,omitempty"`
}

func init() {
	types.Add("vsan:VSANEntityCompatibilityResult", reflect.TypeOf((*VSANEntityCompatibilityResult)(nil)).Elem())
}

type VsanHostDeviceInfo struct {
	types.DynamicData

	Hostname string                `xml:"hostname"`
	Devices  []VsanBasicDeviceInfo `xml:"devices,omitempty"`
}

func init() {
	types.Add("vsan:VsanHostDeviceInfo", reflect.TypeOf((*VsanHostDeviceInfo)(nil)).Elem())
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

type VsanMountPrecheckResult struct {
	types.DynamicData

	Result []VsanMountPrecheckItem `xml:"result"`
}

func init() {
	types.Add("vsan:VsanMountPrecheckResult", reflect.TypeOf((*VsanMountPrecheckResult)(nil)).Elem())
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

type VsanSpaceUsageWithDatastoreType struct {
	types.DynamicData

	SpaceUsage    *VsanSpaceUsage `xml:"spaceUsage,omitempty"`
	DatastoreType string          `xml:"datastoreType,omitempty"`
}

func init() {
	types.Add("vsan:VsanSpaceUsageWithDatastoreType", reflect.TypeOf((*VsanSpaceUsageWithDatastoreType)(nil)).Elem())
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

type VsanHclNicInfo struct {
	VsanHclCommonDeviceInfo
}

func init() {
	types.Add("vsan:VsanHclNicInfo", reflect.TypeOf((*VsanHclNicInfo)(nil)).Elem())
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

type VsanRepairObjectsResult struct {
	types.DynamicData

	InQueueObjects      []string                       `xml:"inQueueObjects,omitempty"`
	FailedRepairObjects []VsanFailedRepairObjectResult `xml:"failedRepairObjects,omitempty"`
	NotInQueueObjects   []string                       `xml:"notInQueueObjects,omitempty"`
}

func init() {
	types.Add("vsan:VsanRepairObjectsResult", reflect.TypeOf((*VsanRepairObjectsResult)(nil)).Elem())
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

type VsanVmVdsMigrationSpec struct {
	types.DynamicData

	VmInstanceUuid string                     `xml:"vmInstanceUuid"`
	Vnics          []VsanVnicVdsMigrationSpec `xml:"vnics"`
}

func init() {
	types.Add("vsan:VsanVmVdsMigrationSpec", reflect.TypeOf((*VsanVmVdsMigrationSpec)(nil)).Elem())
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

type VsanBasicDeviceInfo struct {
	types.DynamicData

	DeviceName string `xml:"deviceName"`
	PciId      string `xml:"pciId,omitempty"`
	FwVersion  string `xml:"fwVersion,omitempty"`
}

func init() {
	types.Add("vsan:VsanBasicDeviceInfo", reflect.TypeOf((*VsanBasicDeviceInfo)(nil)).Elem())
}

type VsanClusterMembershipInfo struct {
	types.DynamicData

	ClusterUuid    string   `xml:"clusterUuid,omitempty"`
	Health         string   `xml:"health,omitempty"`
	MembershipUuid string   `xml:"membershipUuid,omitempty"`
	MemberUuid     []string `xml:"memberUuid,omitempty"`
}

func init() {
	types.Add("vsan:VsanClusterMembershipInfo", reflect.TypeOf((*VsanClusterMembershipInfo)(nil)).Elem())
}

type VsanVcsaDeploymentProgress struct {
	types.DynamicData

	Phase         string                        `xml:"phase"`
	ProgressPct   int64                         `xml:"progressPct"`
	Message       string                        `xml:"message"`
	Success       bool                          `xml:"success"`
	Error         types.BaseMethodFault         `xml:"error,omitempty,typeattr"`
	UpdateCounter int64                         `xml:"updateCounter"`
	TaskId        string                        `xml:"taskId,omitempty"`
	Vm            *types.ManagedObjectReference `xml:"vm,omitempty"`
}

func init() {
	types.Add("vsan:VsanVcsaDeploymentProgress", reflect.TypeOf((*VsanVcsaDeploymentProgress)(nil)).Elem())
}

type VsanFileServerHealthSummary struct {
	types.DynamicData

	DomainName       string `xml:"domainName,omitempty"`
	FileServerIp     string `xml:"fileServerIp,omitempty"`
	NfsdHealth       string `xml:"nfsdHealth,omitempty"`
	NetworkHealth    string `xml:"networkHealth,omitempty"`
	RootfsHealth     string `xml:"rootfsHealth,omitempty"`
	Description      string `xml:"description,omitempty"`
	SmbConnections   int32  `xml:"smbConnections,omitempty"`
	SmbDaemonHealth  string `xml:"smbDaemonHealth,omitempty"`
	AdTestJoinHealth string `xml:"adTestJoinHealth,omitempty"`
	DnsLookupHealth  string `xml:"dnsLookupHealth,omitempty"`
}

func init() {
	types.Add("vsan:VsanFileServerHealthSummary", reflect.TypeOf((*VsanFileServerHealthSummary)(nil)).Elem())
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
	Disks                  []VsanHclDiskInfo        `xml:"disks,omitempty"`
	Issues                 []types.BaseMethodFault  `xml:"issues,omitempty,typeattr"`
	RemediableIssues       []string                 `xml:"remediableIssues,omitempty"`
	DriversOnHcl           []VsanHclDriverInfo      `xml:"driversOnHcl,omitempty"`
	FwAuxVersion           string                   `xml:"fwAuxVersion,omitempty"`
	QueueDepth             int64                    `xml:"queueDepth,omitempty"`
	QueueDepthOnHcl        int64                    `xml:"queueDepthOnHcl,omitempty"`
	QueueDepthSupported    *bool                    `xml:"queueDepthSupported"`
	DiskMode               string                   `xml:"diskMode,omitempty"`
	DiskModeOnHcl          []string                 `xml:"diskModeOnHcl,omitempty"`
	DiskModeSupported      *bool                    `xml:"diskModeSupported"`
	ToolName               string                   `xml:"toolName,omitempty"`
	ToolVersion            string                   `xml:"toolVersion,omitempty"`
}

func init() {
	types.Add("vsan:VsanHclControllerInfo", reflect.TypeOf((*VsanHclControllerInfo)(nil)).Elem())
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

type VsanIscsiTargetBasicInfo struct {
	types.DynamicData

	Alias string `xml:"alias"`
	Iqn   string `xml:"iqn,omitempty"`
}

func init() {
	types.Add("vsan:VsanIscsiTargetBasicInfo", reflect.TypeOf((*VsanIscsiTargetBasicInfo)(nil)).Elem())
}

type VsanFileShareSmbOptions struct {
	types.DynamicData

	Encryption string `xml:"encryption,omitempty"`
}

func init() {
	types.Add("vsan:VsanFileShareSmbOptions", reflect.TypeOf((*VsanFileShareSmbOptions)(nil)).Elem())
}

type VsanCapacityReservationInfo struct {
	types.DynamicData

	HostRebuildThreshold string `xml:"hostRebuildThreshold,omitempty"`
	VsanOpSpaceThreshold string `xml:"vsanOpSpaceThreshold,omitempty"`
}

func init() {
	types.Add("vsan:VsanCapacityReservationInfo", reflect.TypeOf((*VsanCapacityReservationInfo)(nil)).Elem())
}

type VsanSmartStatsHostSummary struct {
	types.DynamicData

	Hostname   string               `xml:"hostname,omitempty"`
	SmartStats []VsanSmartDiskStats `xml:"smartStats,omitempty"`
}

func init() {
	types.Add("vsan:VsanSmartStatsHostSummary", reflect.TypeOf((*VsanSmartStatsHostSummary)(nil)).Elem())
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

type VsanClusterBurnInTestResultList struct {
	types.DynamicData

	Items []VsanBurnInTest `xml:"items,omitempty"`
	Hosts []string         `xml:"hosts,omitempty"`
}

func init() {
	types.Add("vsan:VsanClusterBurnInTestResultList", reflect.TypeOf((*VsanClusterBurnInTestResultList)(nil)).Elem())
}

type VsanFileShareQuerySpec struct {
	types.DynamicData

	DomainName string                        `xml:"domainName,omitempty"`
	Uuids      []string                      `xml:"uuids,omitempty"`
	Names      []string                      `xml:"names,omitempty"`
	Offset     string                        `xml:"offset,omitempty"`
	Limit      int64                         `xml:"limit,omitempty"`
	ManagedBy  []string                      `xml:"managedBy,omitempty"`
	Protocols  []string                      `xml:"protocols,omitempty"`
	PageNumber int64                         `xml:"pageNumber,omitempty"`
	Properties *VsanFileShareQueryProperties `xml:"properties,omitempty"`
}

func init() {
	types.Add("vsan:VsanFileShareQuerySpec", reflect.TypeOf((*VsanFileShareQuerySpec)(nil)).Elem())
}

type VsanMassCollectorPropertyParams struct {
	types.DynamicData

	PropertyName   string              `xml:"propertyName,omitempty"`
	PropertyParams []types.KeyAnyValue `xml:"propertyParams,omitempty"`
}

func init() {
	types.Add("vsan:VsanMassCollectorPropertyParams", reflect.TypeOf((*VsanMassCollectorPropertyParams)(nil)).Elem())
}

type VsanHostVirtualApplianceInfo struct {
	types.DynamicData

	HostKey      types.ManagedObjectReference `xml:"hostKey"`
	IsVirtualApp bool                         `xml:"isVirtualApp"`
}

func init() {
	types.Add("vsan:VsanHostVirtualApplianceInfo", reflect.TypeOf((*VsanHostVirtualApplianceInfo)(nil)).Elem())
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

type VsanDaemonHealth struct {
	types.DynamicData

	Name  string                `xml:"name"`
	Alive bool                  `xml:"alive"`
	Error types.BaseMethodFault `xml:"error,omitempty,typeattr"`
}

func init() {
	types.Add("vsan:VsanDaemonHealth", reflect.TypeOf((*VsanDaemonHealth)(nil)).Elem())
}

type VsanIoInsightInstanceQuerySpec struct {
	types.DynamicData

	State       string `xml:"state,omitempty"`
	EntityRefId string `xml:"entityRefId,omitempty"`
}

func init() {
	types.Add("vsan:VsanIoInsightInstanceQuerySpec", reflect.TypeOf((*VsanIoInsightInstanceQuerySpec)(nil)).Elem())
}

type VsanPerfTopEntities struct {
	types.DynamicData

	MetricId VsanPerfMetricId    `xml:"metricId"`
	Entities []VsanPerfTopEntity `xml:"entities"`
}

func init() {
	types.Add("vsan:VsanPerfTopEntities", reflect.TypeOf((*VsanPerfTopEntities)(nil)).Elem())
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

type VsanNetworkConfigPnicSpeedInconsistencyIssue struct {
	VsanNetworkConfigBaseIssue

	Host        types.ManagedObjectReference  `xml:"host"`
	VswitchName string                        `xml:"vswitchName,omitempty"`
	Vds         *types.ManagedObjectReference `xml:"vds,omitempty,typeattr"`
	SpeedsMb    []int64                       `xml:"speedsMb"`
}

func init() {
	types.Add("vsan:VsanNetworkConfigPnicSpeedInconsistencyIssue", reflect.TypeOf((*VsanNetworkConfigPnicSpeedInconsistencyIssue)(nil)).Elem())
}

type VsanIscsiLUNSpec struct {
	VsanIscsiLUNCommonInfo

	StoragePolicy *types.VirtualMachineProfileSpec `xml:"storagePolicy,omitempty"`
	NewLunId      int32                            `xml:"newLunId,omitempty"`
}

func init() {
	types.Add("vsan:VsanIscsiLUNSpec", reflect.TypeOf((*VsanIscsiLUNSpec)(nil)).Elem())
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

type VsanMetricsConfig struct {
	types.DynamicData

	Profiles []VsanMetricProfile `xml:"profiles,omitempty"`
}

func init() {
	types.Add("vsan:VsanMetricsConfig", reflect.TypeOf((*VsanMetricsConfig)(nil)).Elem())
}

type VsanObjectInaccessibleIssue struct {
	types.VsanUpgradeSystemPreflightCheckIssue

	Uuids []string `xml:"uuids"`
}

func init() {
	types.Add("vsan:VsanObjectInaccessibleIssue", reflect.TypeOf((*VsanObjectInaccessibleIssue)(nil)).Elem())
}

type VsanHostReference struct {
	types.DynamicData

	Hostname string `xml:"hostname"`
}

func init() {
	types.Add("vsan:VsanHostReference", reflect.TypeOf((*VsanHostReference)(nil)).Elem())
}

type VsanRemoteClusterNotCompatible struct {
	types.VsanUpgradeSystemPreflightCheckIssue

	CompatibilityInfo []types.KeyAnyValue `xml:"compatibilityInfo"`
}

func init() {
	types.Add("vsan:VsanRemoteClusterNotCompatible", reflect.TypeOf((*VsanRemoteClusterNotCompatible)(nil)).Elem())
}

type HostSpbmPolicyBlobInfo struct {
	types.DynamicData

	PolicyBlob string `xml:"policyBlob"`
	Namespace  string `xml:"namespace"`
}

func init() {
	types.Add("vsan:HostSpbmPolicyBlobInfo", reflect.TypeOf((*HostSpbmPolicyBlobInfo)(nil)).Elem())
}

type VsanCompositeConstraint struct {
	VsanResourceConstraint

	NestedConstraints []BaseVsanResourceConstraint `xml:"nestedConstraints,omitempty,typeattr"`
	Conjoiner         string                       `xml:"conjoiner,omitempty"`
}

func init() {
	types.Add("vsan:VsanCompositeConstraint", reflect.TypeOf((*VsanCompositeConstraint)(nil)).Elem())
}

type VsanInTransitEncryptionInfo struct {
	types.DynamicData

	Enabled         *bool  `xml:"enabled"`
	RekeyInterval   int32  `xml:"rekeyInterval,omitempty"`
	TransitionState string `xml:"transitionState,omitempty"`
}

func init() {
	types.Add("vsan:VsanInTransitEncryptionInfo", reflect.TypeOf((*VsanInTransitEncryptionInfo)(nil)).Elem())
}

type VsanConfigCheckResult struct {
	types.DynamicData

	VsanEnabled bool                  `xml:"vsanEnabled"`
	Issues      []VsanConfigBaseIssue `xml:"issues,omitempty"`
}

func init() {
	types.Add("vsan:VsanConfigCheckResult", reflect.TypeOf((*VsanConfigCheckResult)(nil)).Elem())
}

type VsanIscsiTargetServiceSpec struct {
	VsanIscsiTargetServiceConfig

	HomeObjectStoragePolicy *types.VirtualMachineProfileSpec `xml:"homeObjectStoragePolicy,omitempty"`
}

func init() {
	types.Add("vsan:VsanIscsiTargetServiceSpec", reflect.TypeOf((*VsanIscsiTargetServiceSpec)(nil)).Elem())
}

type VsanProactiveRebalanceInfo struct {
	types.DynamicData

	Enabled   *bool `xml:"enabled"`
	Threshold int32 `xml:"threshold,omitempty"`
}

func init() {
	types.Add("vsan:VsanProactiveRebalanceInfo", reflect.TypeOf((*VsanProactiveRebalanceInfo)(nil)).Elem())
}

type VsanMountPrecheckItem struct {
	types.DynamicData

	Type        string                     `xml:"type"`
	Description types.LocalizableMessage   `xml:"description"`
	Status      string                     `xml:"status"`
	Reason      []types.LocalizableMessage `xml:"reason,omitempty"`
}

func init() {
	types.Add("vsan:VsanMountPrecheckItem", reflect.TypeOf((*VsanMountPrecheckItem)(nil)).Elem())
}

type VSANStretchedClusterHostVirtualApplianceStatus struct {
	types.DynamicData

	VcCluster    *types.ManagedObjectReference  `xml:"vcCluster,omitempty"`
	IsVirtualApp *bool                          `xml:"isVirtualApp"`
	VcClusters   []types.ManagedObjectReference `xml:"vcClusters,omitempty"`
}

func init() {
	types.Add("vsan:VSANStretchedClusterHostVirtualApplianceStatus", reflect.TypeOf((*VSANStretchedClusterHostVirtualApplianceStatus)(nil)).Elem())
}

type VsanDisallowDataMovementIssue struct {
	types.VsanUpgradeSystemPreflightCheckIssue
}

func init() {
	types.Add("vsan:VsanDisallowDataMovementIssue", reflect.TypeOf((*VsanDisallowDataMovementIssue)(nil)).Elem())
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

type VsanHostDrsStats struct {
	types.DynamicData

	Host  types.ManagedObjectReference `xml:"host"`
	Stats []byte                       `xml:"stats"`
}

func init() {
	types.Add("vsan:VsanHostDrsStats", reflect.TypeOf((*VsanHostDrsStats)(nil)).Elem())
}

type VsanMountPrecheckNetworkLatency struct {
	types.DynamicData

	Host           types.ManagedObjectReference `xml:"host"`
	NetworkLatency int64                        `xml:"networkLatency"`
	Status         string                       `xml:"status"`
}

func init() {
	types.Add("vsan:VsanMountPrecheckNetworkLatency", reflect.TypeOf((*VsanMountPrecheckNetworkLatency)(nil)).Elem())
}

type VsanClusterHealthResultKeyValuePair struct {
	types.DynamicData

	Key   string `xml:"key,omitempty"`
	Value string `xml:"value,omitempty"`
}

func init() {
	types.Add("vsan:VsanClusterHealthResultKeyValuePair", reflect.TypeOf((*VsanClusterHealthResultKeyValuePair)(nil)).Elem())
}

type VsanClusterConfig struct {
	types.DynamicData

	Config      BaseVsanClusterConfigInfo `xml:"config,typeattr"`
	Name        string                    `xml:"name"`
	Hosts       []string                  `xml:"hosts,omitempty"`
	ToBeDeleted *bool                     `xml:"toBeDeleted"`
}

func init() {
	types.Add("vsan:VsanClusterConfig", reflect.TypeOf((*VsanClusterConfig)(nil)).Elem())
}

type VsanAdvancedDatastoreConfig struct {
	VsanDatastoreConfig

	RemoteDatastores []types.ManagedObjectReference `xml:"remoteDatastores,omitempty"`
}

func init() {
	types.Add("vsan:VsanAdvancedDatastoreConfig", reflect.TypeOf((*VsanAdvancedDatastoreConfig)(nil)).Elem())
}

type VsanConfigInfoEx struct {
	VsanClusterConfigInfo

	DataEfficiencyConfig          *VsanDataEfficiencyConfig          `xml:"dataEfficiencyConfig,omitempty"`
	ResyncIopsLimitConfig         *ResyncIopsInfo                    `xml:"resyncIopsLimitConfig,omitempty"`
	IscsiConfig                   BaseVsanIscsiTargetServiceConfig   `xml:"iscsiConfig,omitempty,typeattr"`
	DataEncryptionConfig          *VsanDataEncryptionConfig          `xml:"dataEncryptionConfig,omitempty"`
	ExtendedConfig                *VsanExtendedConfig                `xml:"extendedConfig,omitempty"`
	DatastoreConfig               BaseVsanDatastoreConfig            `xml:"datastoreConfig,omitempty,typeattr"`
	PerfsvcConfig                 *VsanPerfsvcConfig                 `xml:"perfsvcConfig,omitempty"`
	UnmapConfig                   *VsanUnmapConfig                   `xml:"unmapConfig,omitempty"`
	VumConfig                     *VsanVumConfig                     `xml:"vumConfig,omitempty"`
	FileServiceConfig             *VsanFileServiceConfig             `xml:"fileServiceConfig,omitempty"`
	MetricsConfig                 *VsanMetricsConfig                 `xml:"metricsConfig,omitempty"`
	DataInTransitEncryptionConfig *VsanDataInTransitEncryptionConfig `xml:"dataInTransitEncryptionConfig,omitempty"`
}

func init() {
	types.Add("vsan:VsanConfigInfoEx", reflect.TypeOf((*VsanConfigInfoEx)(nil)).Elem())
}

type VsanHostRuntimeStats struct {
	types.DynamicData

	ResyncIopsInfo           *ResyncIopsInfo       `xml:"resyncIopsInfo,omitempty"`
	ConfigGeneration         *VsanConfigGeneration `xml:"configGeneration,omitempty"`
	SupportedClusterSize     int32                 `xml:"supportedClusterSize,omitempty"`
	RepairTimerInfo          *RepairTimerInfo      `xml:"repairTimerInfo,omitempty"`
	ComponentLimitPerCluster int32                 `xml:"componentLimitPerCluster,omitempty"`
	MaxWitnessClusters       int32                 `xml:"maxWitnessClusters,omitempty"`
}

func init() {
	types.Add("vsan:VsanHostRuntimeStats", reflect.TypeOf((*VsanHostRuntimeStats)(nil)).Elem())
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

type VsanObjectIdentity struct {
	types.DynamicData

	Uuid            string                        `xml:"uuid"`
	Type            string                        `xml:"type"`
	VmInstanceUuid  string                        `xml:"vmInstanceUuid,omitempty"`
	VmNsObjectUuid  string                        `xml:"vmNsObjectUuid,omitempty"`
	Vm              *types.ManagedObjectReference `xml:"vm,omitempty"`
	Description     string                        `xml:"description,omitempty"`
	SpbmProfileUuid string                        `xml:"spbmProfileUuid,omitempty"`
	Metadatas       []types.KeyValue              `xml:"metadatas,omitempty"`
	TypeExtId       string                        `xml:"typeExtId,omitempty"`
}

func init() {
	types.Add("vsan:VsanObjectIdentity", reflect.TypeOf((*VsanObjectIdentity)(nil)).Elem())
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

type VsanIperfClientSpec struct {
	types.DynamicData

	Reverse bool `xml:"Reverse"`
}

func init() {
	types.Add("vsan:VsanIperfClientSpec", reflect.TypeOf((*VsanIperfClientSpec)(nil)).Elem())
}

type VsanHostClomdLivenessResult struct {
	types.DynamicData

	Hostname  string                `xml:"hostname"`
	ClomdStat string                `xml:"clomdStat"`
	Error     types.BaseMethodFault `xml:"error,omitempty,typeattr"`
}

func init() {
	types.Add("vsan:VsanHostClomdLivenessResult", reflect.TypeOf((*VsanHostClomdLivenessResult)(nil)).Elem())
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

type VsanDiskResourceCheckResult struct {
	EntityResourceCheckDetails
}

func init() {
	types.Add("vsan:VsanDiskResourceCheckResult", reflect.TypeOf((*VsanDiskResourceCheckResult)(nil)).Elem())
}

type VsanDirectoryServerConfig struct {
	types.DynamicData
}

func init() {
	types.Add("vsan:VsanDirectoryServerConfig", reflect.TypeOf((*VsanDirectoryServerConfig)(nil)).Elem())
}

type VsanWhatifCapacity struct {
	types.DynamicData

	TotalWhatifCapacityB int64                           `xml:"totalWhatifCapacityB"`
	FreeWhatifCapacityB  int64                           `xml:"freeWhatifCapacityB"`
	StoragePolicy        types.VirtualMachineProfileSpec `xml:"storagePolicy"`
	IsSatisfiable        bool                            `xml:"isSatisfiable"`
}

func init() {
	types.Add("vsan:VsanWhatifCapacity", reflect.TypeOf((*VsanWhatifCapacity)(nil)).Elem())
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

type VsanPropertyConstraint struct {
	VsanResourceConstraint

	PropertyName    string             `xml:"propertyName,omitempty"`
	Comparator      string             `xml:"comparator,omitempty"`
	ComparableValue *types.KeyAnyValue `xml:"comparableValue,omitempty"`
}

func init() {
	types.Add("vsan:VsanPropertyConstraint", reflect.TypeOf((*VsanPropertyConstraint)(nil)).Elem())
}

type VsanDataObfuscationRule struct {
	types.DynamicData
}

func init() {
	types.Add("vsan:VsanDataObfuscationRule", reflect.TypeOf((*VsanDataObfuscationRule)(nil)).Elem())
}

type VsanObjectHealth struct {
	types.DynamicData

	NumObjects      int32    `xml:"numObjects"`
	Health          string   `xml:"health"`
	ObjUuids        []string `xml:"objUuids,omitempty"`
	VsanClusterUuid string   `xml:"vsanClusterUuid,omitempty"`
}

func init() {
	types.Add("vsan:VsanObjectHealth", reflect.TypeOf((*VsanObjectHealth)(nil)).Elem())
}

type VsanVmdkLoadTestSpec struct {
	types.DynamicData

	VmdkCreateSpec     *types.FileBackedVirtualDiskSpec `xml:"vmdkCreateSpec,omitempty"`
	VmdkIOSpec         *VsanVmdkIOLoadSpec              `xml:"vmdkIOSpec,omitempty"`
	VmdkIOSpecSequence []VsanVmdkIOLoadSpec             `xml:"vmdkIOSpecSequence,omitempty"`
	StepDurationSec    int64                            `xml:"stepDurationSec,omitempty"`
}

func init() {
	types.Add("vsan:VsanVmdkLoadTestSpec", reflect.TypeOf((*VsanVmdkLoadTestSpec)(nil)).Elem())
}

type VsanRegexBasedRule struct {
	types.DynamicData

	Rules []string `xml:"rules,omitempty"`
}

func init() {
	types.Add("vsan:VsanRegexBasedRule", reflect.TypeOf((*VsanRegexBasedRule)(nil)).Elem())
}

type QueryVsanManagedStorageSpaceUsageSpec struct {
	types.DynamicData

	DatastoreTypes []string `xml:"datastoreTypes"`
}

func init() {
	types.Add("vsan:QueryVsanManagedStorageSpaceUsageSpec", reflect.TypeOf((*QueryVsanManagedStorageSpaceUsageSpec)(nil)).Elem())
}

type VsanServerHostUnicastInfo struct {
	types.DynamicData

	HostUuid    string                   `xml:"hostUuid"`
	UnicastSpec []VsanUnicastAddressInfo `xml:"unicastSpec,omitempty"`
}

func init() {
	types.Add("vsan:VsanServerHostUnicastInfo", reflect.TypeOf((*VsanServerHostUnicastInfo)(nil)).Elem())
}

type VsanVibInstallPreflightStatus struct {
	types.DynamicData

	ManualVmotionRequired bool `xml:"manualVmotionRequired"`
	RollingRequired       bool `xml:"rollingRequired"`
}

func init() {
	types.Add("vsan:VsanVibInstallPreflightStatus", reflect.TypeOf((*VsanVibInstallPreflightStatus)(nil)).Elem())
}
