/*
Copyright (c) 2014 VMware, Inc. All Rights Reserved.

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

import "reflect"

func (b *Action) GetAction() *Action { return b }

type BaseAction interface {
	GetAction() *Action
}

func init() {
	t["BaseAction"] = reflect.TypeOf((*Action)(nil)).Elem()
}

func (b *AlarmAction) GetAlarmAction() *AlarmAction { return b }

type BaseAlarmAction interface {
	GetAlarmAction() *AlarmAction
}

func init() {
	t["BaseAlarmAction"] = reflect.TypeOf((*AlarmAction)(nil)).Elem()
}

func (b *AlarmExpression) GetAlarmExpression() *AlarmExpression { return b }

type BaseAlarmExpression interface {
	GetAlarmExpression() *AlarmExpression
}

func init() {
	t["BaseAlarmExpression"] = reflect.TypeOf((*AlarmExpression)(nil)).Elem()
}

func (b *AlarmSpec) GetAlarmSpec() *AlarmSpec { return b }

type BaseAlarmSpec interface {
	GetAlarmSpec() *AlarmSpec
}

func init() {
	t["BaseAlarmSpec"] = reflect.TypeOf((*AlarmSpec)(nil)).Elem()
}

func (b *AnswerFileCreateSpec) GetAnswerFileCreateSpec() *AnswerFileCreateSpec { return b }

type BaseAnswerFileCreateSpec interface {
	GetAnswerFileCreateSpec() *AnswerFileCreateSpec
}

func init() {
	t["BaseAnswerFileCreateSpec"] = reflect.TypeOf((*AnswerFileCreateSpec)(nil)).Elem()
}

func (b *ApplyProfile) GetApplyProfile() *ApplyProfile { return b }

type BaseApplyProfile interface {
	GetApplyProfile() *ApplyProfile
}

func init() {
	t["BaseApplyProfile"] = reflect.TypeOf((*ApplyProfile)(nil)).Elem()
}

func (b *ClusterAction) GetClusterAction() *ClusterAction { return b }

type BaseClusterAction interface {
	GetClusterAction() *ClusterAction
}

func init() {
	t["BaseClusterAction"] = reflect.TypeOf((*ClusterAction)(nil)).Elem()
}

func (b *ClusterDasAdmissionControlInfo) GetClusterDasAdmissionControlInfo() *ClusterDasAdmissionControlInfo {
	return b
}

type BaseClusterDasAdmissionControlInfo interface {
	GetClusterDasAdmissionControlInfo() *ClusterDasAdmissionControlInfo
}

func init() {
	t["BaseClusterDasAdmissionControlInfo"] = reflect.TypeOf((*ClusterDasAdmissionControlInfo)(nil)).Elem()
}

func (b *ClusterDasAdmissionControlPolicy) GetClusterDasAdmissionControlPolicy() *ClusterDasAdmissionControlPolicy {
	return b
}

type BaseClusterDasAdmissionControlPolicy interface {
	GetClusterDasAdmissionControlPolicy() *ClusterDasAdmissionControlPolicy
}

func init() {
	t["BaseClusterDasAdmissionControlPolicy"] = reflect.TypeOf((*ClusterDasAdmissionControlPolicy)(nil)).Elem()
}

func (b *ClusterDasAdvancedRuntimeInfo) GetClusterDasAdvancedRuntimeInfo() *ClusterDasAdvancedRuntimeInfo {
	return b
}

type BaseClusterDasAdvancedRuntimeInfo interface {
	GetClusterDasAdvancedRuntimeInfo() *ClusterDasAdvancedRuntimeInfo
}

func init() {
	t["BaseClusterDasAdvancedRuntimeInfo"] = reflect.TypeOf((*ClusterDasAdvancedRuntimeInfo)(nil)).Elem()
}

func (b *ClusterDasData) GetClusterDasData() *ClusterDasData { return b }

type BaseClusterDasData interface {
	GetClusterDasData() *ClusterDasData
}

func init() {
	t["BaseClusterDasData"] = reflect.TypeOf((*ClusterDasData)(nil)).Elem()
}

func (b *ClusterDasHostInfo) GetClusterDasHostInfo() *ClusterDasHostInfo { return b }

type BaseClusterDasHostInfo interface {
	GetClusterDasHostInfo() *ClusterDasHostInfo
}

func init() {
	t["BaseClusterDasHostInfo"] = reflect.TypeOf((*ClusterDasHostInfo)(nil)).Elem()
}

func (b *ClusterDrsFaultsFaultsByVm) GetClusterDrsFaultsFaultsByVm() *ClusterDrsFaultsFaultsByVm {
	return b
}

type BaseClusterDrsFaultsFaultsByVm interface {
	GetClusterDrsFaultsFaultsByVm() *ClusterDrsFaultsFaultsByVm
}

func init() {
	t["BaseClusterDrsFaultsFaultsByVm"] = reflect.TypeOf((*ClusterDrsFaultsFaultsByVm)(nil)).Elem()
}

func (b *ClusterGroupInfo) GetClusterGroupInfo() *ClusterGroupInfo { return b }

type BaseClusterGroupInfo interface {
	GetClusterGroupInfo() *ClusterGroupInfo
}

func init() {
	t["BaseClusterGroupInfo"] = reflect.TypeOf((*ClusterGroupInfo)(nil)).Elem()
}

func (b *ClusterRuleInfo) GetClusterRuleInfo() *ClusterRuleInfo { return b }

type BaseClusterRuleInfo interface {
	GetClusterRuleInfo() *ClusterRuleInfo
}

func init() {
	t["BaseClusterRuleInfo"] = reflect.TypeOf((*ClusterRuleInfo)(nil)).Elem()
}

func (b *ClusterSlotPolicy) GetClusterSlotPolicy() *ClusterSlotPolicy { return b }

type BaseClusterSlotPolicy interface {
	GetClusterSlotPolicy() *ClusterSlotPolicy
}

func init() {
	t["BaseClusterSlotPolicy"] = reflect.TypeOf((*ClusterSlotPolicy)(nil)).Elem()
}

func (b *ComputeResourceConfigInfo) GetComputeResourceConfigInfo() *ComputeResourceConfigInfo {
	return b
}

type BaseComputeResourceConfigInfo interface {
	GetComputeResourceConfigInfo() *ComputeResourceConfigInfo
}

func init() {
	t["BaseComputeResourceConfigInfo"] = reflect.TypeOf((*ComputeResourceConfigInfo)(nil)).Elem()
}

func (b *ComputeResourceSummary) GetComputeResourceSummary() *ComputeResourceSummary { return b }

type BaseComputeResourceSummary interface {
	GetComputeResourceSummary() *ComputeResourceSummary
}

func init() {
	t["BaseComputeResourceSummary"] = reflect.TypeOf((*ComputeResourceSummary)(nil)).Elem()
}

func (b *CustomFieldValue) GetCustomFieldValue() *CustomFieldValue { return b }

type BaseCustomFieldValue interface {
	GetCustomFieldValue() *CustomFieldValue
}

func init() {
	t["BaseCustomFieldValue"] = reflect.TypeOf((*CustomFieldValue)(nil)).Elem()
}

func (b *CustomizationIdentitySettings) GetCustomizationIdentitySettings() *CustomizationIdentitySettings {
	return b
}

type BaseCustomizationIdentitySettings interface {
	GetCustomizationIdentitySettings() *CustomizationIdentitySettings
}

func init() {
	t["BaseCustomizationIdentitySettings"] = reflect.TypeOf((*CustomizationIdentitySettings)(nil)).Elem()
}

func (b *CustomizationIpGenerator) GetCustomizationIpGenerator() *CustomizationIpGenerator { return b }

type BaseCustomizationIpGenerator interface {
	GetCustomizationIpGenerator() *CustomizationIpGenerator
}

func init() {
	t["BaseCustomizationIpGenerator"] = reflect.TypeOf((*CustomizationIpGenerator)(nil)).Elem()
}

func (b *CustomizationIpV6Generator) GetCustomizationIpV6Generator() *CustomizationIpV6Generator {
	return b
}

type BaseCustomizationIpV6Generator interface {
	GetCustomizationIpV6Generator() *CustomizationIpV6Generator
}

func init() {
	t["BaseCustomizationIpV6Generator"] = reflect.TypeOf((*CustomizationIpV6Generator)(nil)).Elem()
}

func (b *CustomizationName) GetCustomizationName() *CustomizationName { return b }

type BaseCustomizationName interface {
	GetCustomizationName() *CustomizationName
}

func init() {
	t["BaseCustomizationName"] = reflect.TypeOf((*CustomizationName)(nil)).Elem()
}

func (b *CustomizationOptions) GetCustomizationOptions() *CustomizationOptions { return b }

type BaseCustomizationOptions interface {
	GetCustomizationOptions() *CustomizationOptions
}

func init() {
	t["BaseCustomizationOptions"] = reflect.TypeOf((*CustomizationOptions)(nil)).Elem()
}

func (b *DVPortSetting) GetDVPortSetting() *DVPortSetting { return b }

type BaseDVPortSetting interface {
	GetDVPortSetting() *DVPortSetting
}

func init() {
	t["BaseDVPortSetting"] = reflect.TypeOf((*DVPortSetting)(nil)).Elem()
}

func (b *DVPortgroupPolicy) GetDVPortgroupPolicy() *DVPortgroupPolicy { return b }

type BaseDVPortgroupPolicy interface {
	GetDVPortgroupPolicy() *DVPortgroupPolicy
}

func init() {
	t["BaseDVPortgroupPolicy"] = reflect.TypeOf((*DVPortgroupPolicy)(nil)).Elem()
}

func (b *DVSConfigInfo) GetDVSConfigInfo() *DVSConfigInfo { return b }

type BaseDVSConfigInfo interface {
	GetDVSConfigInfo() *DVSConfigInfo
}

func init() {
	t["BaseDVSConfigInfo"] = reflect.TypeOf((*DVSConfigInfo)(nil)).Elem()
}

func (b *DVSConfigSpec) GetDVSConfigSpec() *DVSConfigSpec { return b }

type BaseDVSConfigSpec interface {
	GetDVSConfigSpec() *DVSConfigSpec
}

func init() {
	t["BaseDVSConfigSpec"] = reflect.TypeOf((*DVSConfigSpec)(nil)).Elem()
}

func (b *DVSFeatureCapability) GetDVSFeatureCapability() *DVSFeatureCapability { return b }

type BaseDVSFeatureCapability interface {
	GetDVSFeatureCapability() *DVSFeatureCapability
}

func init() {
	t["BaseDVSFeatureCapability"] = reflect.TypeOf((*DVSFeatureCapability)(nil)).Elem()
}

func (b *DVSHealthCheckCapability) GetDVSHealthCheckCapability() *DVSHealthCheckCapability { return b }

type BaseDVSHealthCheckCapability interface {
	GetDVSHealthCheckCapability() *DVSHealthCheckCapability
}

func init() {
	t["BaseDVSHealthCheckCapability"] = reflect.TypeOf((*DVSHealthCheckCapability)(nil)).Elem()
}

func (b *DVSHealthCheckConfig) GetDVSHealthCheckConfig() *DVSHealthCheckConfig { return b }

type BaseDVSHealthCheckConfig interface {
	GetDVSHealthCheckConfig() *DVSHealthCheckConfig
}

func init() {
	t["BaseDVSHealthCheckConfig"] = reflect.TypeOf((*DVSHealthCheckConfig)(nil)).Elem()
}

func (b *DVSUplinkPortPolicy) GetDVSUplinkPortPolicy() *DVSUplinkPortPolicy { return b }

type BaseDVSUplinkPortPolicy interface {
	GetDVSUplinkPortPolicy() *DVSUplinkPortPolicy
}

func init() {
	t["BaseDVSUplinkPortPolicy"] = reflect.TypeOf((*DVSUplinkPortPolicy)(nil)).Elem()
}

func (b *DatastoreInfo) GetDatastoreInfo() *DatastoreInfo { return b }

type BaseDatastoreInfo interface {
	GetDatastoreInfo() *DatastoreInfo
}

func init() {
	t["BaseDatastoreInfo"] = reflect.TypeOf((*DatastoreInfo)(nil)).Elem()
}

func (b *Description) GetDescription() *Description { return b }

type BaseDescription interface {
	GetDescription() *Description
}

func init() {
	t["BaseDescription"] = reflect.TypeOf((*Description)(nil)).Elem()
}

func (b *DistributedVirtualSwitchHostMemberBacking) GetDistributedVirtualSwitchHostMemberBacking() *DistributedVirtualSwitchHostMemberBacking {
	return b
}

type BaseDistributedVirtualSwitchHostMemberBacking interface {
	GetDistributedVirtualSwitchHostMemberBacking() *DistributedVirtualSwitchHostMemberBacking
}

func init() {
	t["BaseDistributedVirtualSwitchHostMemberBacking"] = reflect.TypeOf((*DistributedVirtualSwitchHostMemberBacking)(nil)).Elem()
}

func (b *DistributedVirtualSwitchManagerHostDvsFilterSpec) GetDistributedVirtualSwitchManagerHostDvsFilterSpec() *DistributedVirtualSwitchManagerHostDvsFilterSpec {
	return b
}

type BaseDistributedVirtualSwitchManagerHostDvsFilterSpec interface {
	GetDistributedVirtualSwitchManagerHostDvsFilterSpec() *DistributedVirtualSwitchManagerHostDvsFilterSpec
}

func init() {
	t["BaseDistributedVirtualSwitchManagerHostDvsFilterSpec"] = reflect.TypeOf((*DistributedVirtualSwitchManagerHostDvsFilterSpec)(nil)).Elem()
}

func (b *DvsNetworkRuleAction) GetDvsNetworkRuleAction() *DvsNetworkRuleAction { return b }

type BaseDvsNetworkRuleAction interface {
	GetDvsNetworkRuleAction() *DvsNetworkRuleAction
}

func init() {
	t["BaseDvsNetworkRuleAction"] = reflect.TypeOf((*DvsNetworkRuleAction)(nil)).Elem()
}

func (b *DvsNetworkRuleQualifier) GetDvsNetworkRuleQualifier() *DvsNetworkRuleQualifier { return b }

type BaseDvsNetworkRuleQualifier interface {
	GetDvsNetworkRuleQualifier() *DvsNetworkRuleQualifier
}

func init() {
	t["BaseDvsNetworkRuleQualifier"] = reflect.TypeOf((*DvsNetworkRuleQualifier)(nil)).Elem()
}

func (b *DynamicData) GetDynamicData() *DynamicData { return b }

type BaseDynamicData interface {
	GetDynamicData() *DynamicData
}

func init() {
	t["BaseDynamicData"] = reflect.TypeOf((*DynamicData)(nil)).Elem()
}

func (b *Event) GetEvent() *Event { return b }

type BaseEvent interface {
	GetEvent() *Event
}

func init() {
	t["BaseEvent"] = reflect.TypeOf((*Event)(nil)).Elem()
}

func (b *FaultToleranceConfigInfo) GetFaultToleranceConfigInfo() *FaultToleranceConfigInfo { return b }

type BaseFaultToleranceConfigInfo interface {
	GetFaultToleranceConfigInfo() *FaultToleranceConfigInfo
}

func init() {
	t["BaseFaultToleranceConfigInfo"] = reflect.TypeOf((*FaultToleranceConfigInfo)(nil)).Elem()
}

func (b *FileInfo) GetFileInfo() *FileInfo { return b }

type BaseFileInfo interface {
	GetFileInfo() *FileInfo
}

func init() {
	t["BaseFileInfo"] = reflect.TypeOf((*FileInfo)(nil)).Elem()
}

func (b *FileQuery) GetFileQuery() *FileQuery { return b }

type BaseFileQuery interface {
	GetFileQuery() *FileQuery
}

func init() {
	t["BaseFileQuery"] = reflect.TypeOf((*FileQuery)(nil)).Elem()
}

func (b *GuestAuthentication) GetGuestAuthentication() *GuestAuthentication { return b }

type BaseGuestAuthentication interface {
	GetGuestAuthentication() *GuestAuthentication
}

func init() {
	t["BaseGuestAuthentication"] = reflect.TypeOf((*GuestAuthentication)(nil)).Elem()
}

func (b *GuestFileAttributes) GetGuestFileAttributes() *GuestFileAttributes { return b }

type BaseGuestFileAttributes interface {
	GetGuestFileAttributes() *GuestFileAttributes
}

func init() {
	t["BaseGuestFileAttributes"] = reflect.TypeOf((*GuestFileAttributes)(nil)).Elem()
}

func (b *GuestProgramSpec) GetGuestProgramSpec() *GuestProgramSpec { return b }

type BaseGuestProgramSpec interface {
	GetGuestProgramSpec() *GuestProgramSpec
}

func init() {
	t["BaseGuestProgramSpec"] = reflect.TypeOf((*GuestProgramSpec)(nil)).Elem()
}

func (b *HostAccountSpec) GetHostAccountSpec() *HostAccountSpec { return b }

type BaseHostAccountSpec interface {
	GetHostAccountSpec() *HostAccountSpec
}

func init() {
	t["BaseHostAccountSpec"] = reflect.TypeOf((*HostAccountSpec)(nil)).Elem()
}

func (b *HostAuthenticationStoreInfo) GetHostAuthenticationStoreInfo() *HostAuthenticationStoreInfo {
	return b
}

type BaseHostAuthenticationStoreInfo interface {
	GetHostAuthenticationStoreInfo() *HostAuthenticationStoreInfo
}

func init() {
	t["BaseHostAuthenticationStoreInfo"] = reflect.TypeOf((*HostAuthenticationStoreInfo)(nil)).Elem()
}

func (b *HostConnectInfoNetworkInfo) GetHostConnectInfoNetworkInfo() *HostConnectInfoNetworkInfo {
	return b
}

type BaseHostConnectInfoNetworkInfo interface {
	GetHostConnectInfoNetworkInfo() *HostConnectInfoNetworkInfo
}

func init() {
	t["BaseHostConnectInfoNetworkInfo"] = reflect.TypeOf((*HostConnectInfoNetworkInfo)(nil)).Elem()
}

func (b *HostDatastoreConnectInfo) GetHostDatastoreConnectInfo() *HostDatastoreConnectInfo { return b }

type BaseHostDatastoreConnectInfo interface {
	GetHostDatastoreConnectInfo() *HostDatastoreConnectInfo
}

func init() {
	t["BaseHostDatastoreConnectInfo"] = reflect.TypeOf((*HostDatastoreConnectInfo)(nil)).Elem()
}

func (b *HostDnsConfig) GetHostDnsConfig() *HostDnsConfig { return b }

type BaseHostDnsConfig interface {
	GetHostDnsConfig() *HostDnsConfig
}

func init() {
	t["BaseHostDnsConfig"] = reflect.TypeOf((*HostDnsConfig)(nil)).Elem()
}

func (b *HostFileSystemVolume) GetHostFileSystemVolume() *HostFileSystemVolume { return b }

type BaseHostFileSystemVolume interface {
	GetHostFileSystemVolume() *HostFileSystemVolume
}

func init() {
	t["BaseHostFileSystemVolume"] = reflect.TypeOf((*HostFileSystemVolume)(nil)).Elem()
}

func (b *HostHostBusAdapter) GetHostHostBusAdapter() *HostHostBusAdapter { return b }

type BaseHostHostBusAdapter interface {
	GetHostHostBusAdapter() *HostHostBusAdapter
}

func init() {
	t["BaseHostHostBusAdapter"] = reflect.TypeOf((*HostHostBusAdapter)(nil)).Elem()
}

func (b *HostIpRouteConfig) GetHostIpRouteConfig() *HostIpRouteConfig { return b }

type BaseHostIpRouteConfig interface {
	GetHostIpRouteConfig() *HostIpRouteConfig
}

func init() {
	t["BaseHostIpRouteConfig"] = reflect.TypeOf((*HostIpRouteConfig)(nil)).Elem()
}

func (b *HostMemberHealthCheckResult) GetHostMemberHealthCheckResult() *HostMemberHealthCheckResult {
	return b
}

type BaseHostMemberHealthCheckResult interface {
	GetHostMemberHealthCheckResult() *HostMemberHealthCheckResult
}

func init() {
	t["BaseHostMemberHealthCheckResult"] = reflect.TypeOf((*HostMemberHealthCheckResult)(nil)).Elem()
}

func (b *HostMultipathInfoLogicalUnitPolicy) GetHostMultipathInfoLogicalUnitPolicy() *HostMultipathInfoLogicalUnitPolicy {
	return b
}

type BaseHostMultipathInfoLogicalUnitPolicy interface {
	GetHostMultipathInfoLogicalUnitPolicy() *HostMultipathInfoLogicalUnitPolicy
}

func init() {
	t["BaseHostMultipathInfoLogicalUnitPolicy"] = reflect.TypeOf((*HostMultipathInfoLogicalUnitPolicy)(nil)).Elem()
}

func (b *HostPciPassthruConfig) GetHostPciPassthruConfig() *HostPciPassthruConfig { return b }

type BaseHostPciPassthruConfig interface {
	GetHostPciPassthruConfig() *HostPciPassthruConfig
}

func init() {
	t["BaseHostPciPassthruConfig"] = reflect.TypeOf((*HostPciPassthruConfig)(nil)).Elem()
}

func (b *HostPciPassthruInfo) GetHostPciPassthruInfo() *HostPciPassthruInfo { return b }

type BaseHostPciPassthruInfo interface {
	GetHostPciPassthruInfo() *HostPciPassthruInfo
}

func init() {
	t["BaseHostPciPassthruInfo"] = reflect.TypeOf((*HostPciPassthruInfo)(nil)).Elem()
}

func (b *HostSystemSwapConfigurationSystemSwapOption) GetHostSystemSwapConfigurationSystemSwapOption() *HostSystemSwapConfigurationSystemSwapOption {
	return b
}

type BaseHostSystemSwapConfigurationSystemSwapOption interface {
	GetHostSystemSwapConfigurationSystemSwapOption() *HostSystemSwapConfigurationSystemSwapOption
}

func init() {
	t["BaseHostSystemSwapConfigurationSystemSwapOption"] = reflect.TypeOf((*HostSystemSwapConfigurationSystemSwapOption)(nil)).Elem()
}

func (b *HostTargetTransport) GetHostTargetTransport() *HostTargetTransport { return b }

type BaseHostTargetTransport interface {
	GetHostTargetTransport() *HostTargetTransport
}

func init() {
	t["BaseHostTargetTransport"] = reflect.TypeOf((*HostTargetTransport)(nil)).Elem()
}

func (b *HostTpmEventDetails) GetHostTpmEventDetails() *HostTpmEventDetails { return b }

type BaseHostTpmEventDetails interface {
	GetHostTpmEventDetails() *HostTpmEventDetails
}

func init() {
	t["BaseHostTpmEventDetails"] = reflect.TypeOf((*HostTpmEventDetails)(nil)).Elem()
}

func (b *HostVirtualSwitchBridge) GetHostVirtualSwitchBridge() *HostVirtualSwitchBridge { return b }

type BaseHostVirtualSwitchBridge interface {
	GetHostVirtualSwitchBridge() *HostVirtualSwitchBridge
}

func init() {
	t["BaseHostVirtualSwitchBridge"] = reflect.TypeOf((*HostVirtualSwitchBridge)(nil)).Elem()
}

func (b *ImportSpec) GetImportSpec() *ImportSpec { return b }

type BaseImportSpec interface {
	GetImportSpec() *ImportSpec
}

func init() {
	t["BaseImportSpec"] = reflect.TypeOf((*ImportSpec)(nil)).Elem()
}

func (b *LicenseSource) GetLicenseSource() *LicenseSource { return b }

type BaseLicenseSource interface {
	GetLicenseSource() *LicenseSource
}

func init() {
	t["BaseLicenseSource"] = reflect.TypeOf((*LicenseSource)(nil)).Elem()
}

func (b *MethodFault) GetMethodFault() *MethodFault { return b }

type BaseMethodFault interface {
	GetMethodFault() *MethodFault
}

func init() {
	t["BaseMethodFault"] = reflect.TypeOf((*MethodFault)(nil)).Elem()
}

func (b *NetBIOSConfigInfo) GetNetBIOSConfigInfo() *NetBIOSConfigInfo { return b }

type BaseNetBIOSConfigInfo interface {
	GetNetBIOSConfigInfo() *NetBIOSConfigInfo
}

func init() {
	t["BaseNetBIOSConfigInfo"] = reflect.TypeOf((*NetBIOSConfigInfo)(nil)).Elem()
}

func (b *OptionType) GetOptionType() *OptionType { return b }

type BaseOptionType interface {
	GetOptionType() *OptionType
}

func init() {
	t["BaseOptionType"] = reflect.TypeOf((*OptionType)(nil)).Elem()
}

func (b *PerfEntityMetricBase) GetPerfEntityMetricBase() *PerfEntityMetricBase { return b }

type BasePerfEntityMetricBase interface {
	GetPerfEntityMetricBase() *PerfEntityMetricBase
}

func init() {
	t["BasePerfEntityMetricBase"] = reflect.TypeOf((*PerfEntityMetricBase)(nil)).Elem()
}

func (b *PerfMetricSeries) GetPerfMetricSeries() *PerfMetricSeries { return b }

type BasePerfMetricSeries interface {
	GetPerfMetricSeries() *PerfMetricSeries
}

func init() {
	t["BasePerfMetricSeries"] = reflect.TypeOf((*PerfMetricSeries)(nil)).Elem()
}

func (b *PolicyOption) GetPolicyOption() *PolicyOption { return b }

type BasePolicyOption interface {
	GetPolicyOption() *PolicyOption
}

func init() {
	t["BasePolicyOption"] = reflect.TypeOf((*PolicyOption)(nil)).Elem()
}

func (b *ProfileConfigInfo) GetProfileConfigInfo() *ProfileConfigInfo { return b }

type BaseProfileConfigInfo interface {
	GetProfileConfigInfo() *ProfileConfigInfo
}

func init() {
	t["BaseProfileConfigInfo"] = reflect.TypeOf((*ProfileConfigInfo)(nil)).Elem()
}

func (b *ProfileCreateSpec) GetProfileCreateSpec() *ProfileCreateSpec { return b }

type BaseProfileCreateSpec interface {
	GetProfileCreateSpec() *ProfileCreateSpec
}

func init() {
	t["BaseProfileCreateSpec"] = reflect.TypeOf((*ProfileCreateSpec)(nil)).Elem()
}

func (b *ProfileExpression) GetProfileExpression() *ProfileExpression { return b }

type BaseProfileExpression interface {
	GetProfileExpression() *ProfileExpression
}

func init() {
	t["BaseProfileExpression"] = reflect.TypeOf((*ProfileExpression)(nil)).Elem()
}

func (b *ProfilePolicyOptionMetadata) GetProfilePolicyOptionMetadata() *ProfilePolicyOptionMetadata {
	return b
}

type BaseProfilePolicyOptionMetadata interface {
	GetProfilePolicyOptionMetadata() *ProfilePolicyOptionMetadata
}

func init() {
	t["BaseProfilePolicyOptionMetadata"] = reflect.TypeOf((*ProfilePolicyOptionMetadata)(nil)).Elem()
}

func (b *ResourcePoolSummary) GetResourcePoolSummary() *ResourcePoolSummary { return b }

type BaseResourcePoolSummary interface {
	GetResourcePoolSummary() *ResourcePoolSummary
}

func init() {
	t["BaseResourcePoolSummary"] = reflect.TypeOf((*ResourcePoolSummary)(nil)).Elem()
}

func (b *ScheduledTaskSpec) GetScheduledTaskSpec() *ScheduledTaskSpec { return b }

type BaseScheduledTaskSpec interface {
	GetScheduledTaskSpec() *ScheduledTaskSpec
}

func init() {
	t["BaseScheduledTaskSpec"] = reflect.TypeOf((*ScheduledTaskSpec)(nil)).Elem()
}

func (b *SelectionSet) GetSelectionSet() *SelectionSet { return b }

type BaseSelectionSet interface {
	GetSelectionSet() *SelectionSet
}

func init() {
	t["BaseSelectionSet"] = reflect.TypeOf((*SelectionSet)(nil)).Elem()
}

func (b *SelectionSpec) GetSelectionSpec() *SelectionSpec { return b }

type BaseSelectionSpec interface {
	GetSelectionSpec() *SelectionSpec
}

func init() {
	t["BaseSelectionSpec"] = reflect.TypeOf((*SelectionSpec)(nil)).Elem()
}

func (b *SessionManagerServiceRequestSpec) GetSessionManagerServiceRequestSpec() *SessionManagerServiceRequestSpec {
	return b
}

type BaseSessionManagerServiceRequestSpec interface {
	GetSessionManagerServiceRequestSpec() *SessionManagerServiceRequestSpec
}

func init() {
	t["BaseSessionManagerServiceRequestSpec"] = reflect.TypeOf((*SessionManagerServiceRequestSpec)(nil)).Elem()
}

func (b *TaskReason) GetTaskReason() *TaskReason { return b }

type BaseTaskReason interface {
	GetTaskReason() *TaskReason
}

func init() {
	t["BaseTaskReason"] = reflect.TypeOf((*TaskReason)(nil)).Elem()
}

func (b *TaskScheduler) GetTaskScheduler() *TaskScheduler { return b }

type BaseTaskScheduler interface {
	GetTaskScheduler() *TaskScheduler
}

func init() {
	t["BaseTaskScheduler"] = reflect.TypeOf((*TaskScheduler)(nil)).Elem()
}

func (b *UserSearchResult) GetUserSearchResult() *UserSearchResult { return b }

type BaseUserSearchResult interface {
	GetUserSearchResult() *UserSearchResult
}

func init() {
	t["BaseUserSearchResult"] = reflect.TypeOf((*UserSearchResult)(nil)).Elem()
}

func (b *VirtualDevice) GetVirtualDevice() *VirtualDevice { return b }

type BaseVirtualDevice interface {
	GetVirtualDevice() *VirtualDevice
}

func init() {
	t["BaseVirtualDevice"] = reflect.TypeOf((*VirtualDevice)(nil)).Elem()
}

func (b *VirtualDeviceBackingInfo) GetVirtualDeviceBackingInfo() *VirtualDeviceBackingInfo { return b }

type BaseVirtualDeviceBackingInfo interface {
	GetVirtualDeviceBackingInfo() *VirtualDeviceBackingInfo
}

func init() {
	t["BaseVirtualDeviceBackingInfo"] = reflect.TypeOf((*VirtualDeviceBackingInfo)(nil)).Elem()
}

func (b *VirtualDeviceBackingOption) GetVirtualDeviceBackingOption() *VirtualDeviceBackingOption {
	return b
}

type BaseVirtualDeviceBackingOption interface {
	GetVirtualDeviceBackingOption() *VirtualDeviceBackingOption
}

func init() {
	t["BaseVirtualDeviceBackingOption"] = reflect.TypeOf((*VirtualDeviceBackingOption)(nil)).Elem()
}

func (b *VirtualDeviceBusSlotInfo) GetVirtualDeviceBusSlotInfo() *VirtualDeviceBusSlotInfo { return b }

type BaseVirtualDeviceBusSlotInfo interface {
	GetVirtualDeviceBusSlotInfo() *VirtualDeviceBusSlotInfo
}

func init() {
	t["BaseVirtualDeviceBusSlotInfo"] = reflect.TypeOf((*VirtualDeviceBusSlotInfo)(nil)).Elem()
}

func (b *VirtualDeviceConfigSpec) GetVirtualDeviceConfigSpec() *VirtualDeviceConfigSpec { return b }

type BaseVirtualDeviceConfigSpec interface {
	GetVirtualDeviceConfigSpec() *VirtualDeviceConfigSpec
}

func init() {
	t["BaseVirtualDeviceConfigSpec"] = reflect.TypeOf((*VirtualDeviceConfigSpec)(nil)).Elem()
}

func (b *VirtualDeviceOption) GetVirtualDeviceOption() *VirtualDeviceOption { return b }

type BaseVirtualDeviceOption interface {
	GetVirtualDeviceOption() *VirtualDeviceOption
}

func init() {
	t["BaseVirtualDeviceOption"] = reflect.TypeOf((*VirtualDeviceOption)(nil)).Elem()
}

func (b *VirtualDiskSpec) GetVirtualDiskSpec() *VirtualDiskSpec { return b }

type BaseVirtualDiskSpec interface {
	GetVirtualDiskSpec() *VirtualDiskSpec
}

func init() {
	t["BaseVirtualDiskSpec"] = reflect.TypeOf((*VirtualDiskSpec)(nil)).Elem()
}

func (b *VirtualMachineBootOptionsBootableDevice) GetVirtualMachineBootOptionsBootableDevice() *VirtualMachineBootOptionsBootableDevice {
	return b
}

type BaseVirtualMachineBootOptionsBootableDevice interface {
	GetVirtualMachineBootOptionsBootableDevice() *VirtualMachineBootOptionsBootableDevice
}

func init() {
	t["BaseVirtualMachineBootOptionsBootableDevice"] = reflect.TypeOf((*VirtualMachineBootOptionsBootableDevice)(nil)).Elem()
}

func (b *VirtualMachineDeviceRuntimeInfoDeviceRuntimeState) GetVirtualMachineDeviceRuntimeInfoDeviceRuntimeState() *VirtualMachineDeviceRuntimeInfoDeviceRuntimeState {
	return b
}

type BaseVirtualMachineDeviceRuntimeInfoDeviceRuntimeState interface {
	GetVirtualMachineDeviceRuntimeInfoDeviceRuntimeState() *VirtualMachineDeviceRuntimeInfoDeviceRuntimeState
}

func init() {
	t["BaseVirtualMachineDeviceRuntimeInfoDeviceRuntimeState"] = reflect.TypeOf((*VirtualMachineDeviceRuntimeInfoDeviceRuntimeState)(nil)).Elem()
}

func (b *VirtualMachineProfileSpec) GetVirtualMachineProfileSpec() *VirtualMachineProfileSpec {
	return b
}

type BaseVirtualMachineProfileSpec interface {
	GetVirtualMachineProfileSpec() *VirtualMachineProfileSpec
}

func init() {
	t["BaseVirtualMachineProfileSpec"] = reflect.TypeOf((*VirtualMachineProfileSpec)(nil)).Elem()
}

func (b *VirtualMachineTargetInfo) GetVirtualMachineTargetInfo() *VirtualMachineTargetInfo { return b }

type BaseVirtualMachineTargetInfo interface {
	GetVirtualMachineTargetInfo() *VirtualMachineTargetInfo
}

func init() {
	t["BaseVirtualMachineTargetInfo"] = reflect.TypeOf((*VirtualMachineTargetInfo)(nil)).Elem()
}

func (b *VmConfigInfo) GetVmConfigInfo() *VmConfigInfo { return b }

type BaseVmConfigInfo interface {
	GetVmConfigInfo() *VmConfigInfo
}

func init() {
	t["BaseVmConfigInfo"] = reflect.TypeOf((*VmConfigInfo)(nil)).Elem()
}

func (b *VmfsDatastoreBaseOption) GetVmfsDatastoreBaseOption() *VmfsDatastoreBaseOption { return b }

type BaseVmfsDatastoreBaseOption interface {
	GetVmfsDatastoreBaseOption() *VmfsDatastoreBaseOption
}

func init() {
	t["BaseVmfsDatastoreBaseOption"] = reflect.TypeOf((*VmfsDatastoreBaseOption)(nil)).Elem()
}
