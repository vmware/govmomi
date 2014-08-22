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

func (b *Action) GetAction() *Action { return b }

type BaseAction interface {
	GetAction() *Action
}

func (b *AlarmAction) GetAlarmAction() *AlarmAction { return b }

type BaseAlarmAction interface {
	GetAlarmAction() *AlarmAction
}

func (b *AlarmExpression) GetAlarmExpression() *AlarmExpression { return b }

type BaseAlarmExpression interface {
	GetAlarmExpression() *AlarmExpression
}

func (b *AlarmSpec) GetAlarmSpec() *AlarmSpec { return b }

type BaseAlarmSpec interface {
	GetAlarmSpec() *AlarmSpec
}

func (b *AnswerFileCreateSpec) GetAnswerFileCreateSpec() *AnswerFileCreateSpec { return b }

type BaseAnswerFileCreateSpec interface {
	GetAnswerFileCreateSpec() *AnswerFileCreateSpec
}

func (b *ApplyProfile) GetApplyProfile() *ApplyProfile { return b }

type BaseApplyProfile interface {
	GetApplyProfile() *ApplyProfile
}

func (b *ClusterAction) GetClusterAction() *ClusterAction { return b }

type BaseClusterAction interface {
	GetClusterAction() *ClusterAction
}

func (b *ClusterDasAdmissionControlInfo) GetClusterDasAdmissionControlInfo() *ClusterDasAdmissionControlInfo {
	return b
}

type BaseClusterDasAdmissionControlInfo interface {
	GetClusterDasAdmissionControlInfo() *ClusterDasAdmissionControlInfo
}

func (b *ClusterDasAdmissionControlPolicy) GetClusterDasAdmissionControlPolicy() *ClusterDasAdmissionControlPolicy {
	return b
}

type BaseClusterDasAdmissionControlPolicy interface {
	GetClusterDasAdmissionControlPolicy() *ClusterDasAdmissionControlPolicy
}

func (b *ClusterDasAdvancedRuntimeInfo) GetClusterDasAdvancedRuntimeInfo() *ClusterDasAdvancedRuntimeInfo {
	return b
}

type BaseClusterDasAdvancedRuntimeInfo interface {
	GetClusterDasAdvancedRuntimeInfo() *ClusterDasAdvancedRuntimeInfo
}

func (b *ClusterDasData) GetClusterDasData() *ClusterDasData { return b }

type BaseClusterDasData interface {
	GetClusterDasData() *ClusterDasData
}

func (b *ClusterDasHostInfo) GetClusterDasHostInfo() *ClusterDasHostInfo { return b }

type BaseClusterDasHostInfo interface {
	GetClusterDasHostInfo() *ClusterDasHostInfo
}

func (b *ClusterDrsFaultsFaultsByVm) GetClusterDrsFaultsFaultsByVm() *ClusterDrsFaultsFaultsByVm {
	return b
}

type BaseClusterDrsFaultsFaultsByVm interface {
	GetClusterDrsFaultsFaultsByVm() *ClusterDrsFaultsFaultsByVm
}

func (b *ClusterGroupInfo) GetClusterGroupInfo() *ClusterGroupInfo { return b }

type BaseClusterGroupInfo interface {
	GetClusterGroupInfo() *ClusterGroupInfo
}

func (b *ClusterRuleInfo) GetClusterRuleInfo() *ClusterRuleInfo { return b }

type BaseClusterRuleInfo interface {
	GetClusterRuleInfo() *ClusterRuleInfo
}

func (b *ClusterSlotPolicy) GetClusterSlotPolicy() *ClusterSlotPolicy { return b }

type BaseClusterSlotPolicy interface {
	GetClusterSlotPolicy() *ClusterSlotPolicy
}

func (b *ComputeResourceConfigInfo) GetComputeResourceConfigInfo() *ComputeResourceConfigInfo {
	return b
}

type BaseComputeResourceConfigInfo interface {
	GetComputeResourceConfigInfo() *ComputeResourceConfigInfo
}

func (b *ComputeResourceSummary) GetComputeResourceSummary() *ComputeResourceSummary { return b }

type BaseComputeResourceSummary interface {
	GetComputeResourceSummary() *ComputeResourceSummary
}

func (b *CustomFieldValue) GetCustomFieldValue() *CustomFieldValue { return b }

type BaseCustomFieldValue interface {
	GetCustomFieldValue() *CustomFieldValue
}

func (b *CustomizationIdentitySettings) GetCustomizationIdentitySettings() *CustomizationIdentitySettings {
	return b
}

type BaseCustomizationIdentitySettings interface {
	GetCustomizationIdentitySettings() *CustomizationIdentitySettings
}

func (b *CustomizationIpGenerator) GetCustomizationIpGenerator() *CustomizationIpGenerator { return b }

type BaseCustomizationIpGenerator interface {
	GetCustomizationIpGenerator() *CustomizationIpGenerator
}

func (b *CustomizationIpV6Generator) GetCustomizationIpV6Generator() *CustomizationIpV6Generator {
	return b
}

type BaseCustomizationIpV6Generator interface {
	GetCustomizationIpV6Generator() *CustomizationIpV6Generator
}

func (b *CustomizationName) GetCustomizationName() *CustomizationName { return b }

type BaseCustomizationName interface {
	GetCustomizationName() *CustomizationName
}

func (b *CustomizationOptions) GetCustomizationOptions() *CustomizationOptions { return b }

type BaseCustomizationOptions interface {
	GetCustomizationOptions() *CustomizationOptions
}

func (b *DVPortSetting) GetDVPortSetting() *DVPortSetting { return b }

type BaseDVPortSetting interface {
	GetDVPortSetting() *DVPortSetting
}

func (b *DVPortgroupPolicy) GetDVPortgroupPolicy() *DVPortgroupPolicy { return b }

type BaseDVPortgroupPolicy interface {
	GetDVPortgroupPolicy() *DVPortgroupPolicy
}

func (b *DVSConfigInfo) GetDVSConfigInfo() *DVSConfigInfo { return b }

type BaseDVSConfigInfo interface {
	GetDVSConfigInfo() *DVSConfigInfo
}

func (b *DVSConfigSpec) GetDVSConfigSpec() *DVSConfigSpec { return b }

type BaseDVSConfigSpec interface {
	GetDVSConfigSpec() *DVSConfigSpec
}

func (b *DVSFeatureCapability) GetDVSFeatureCapability() *DVSFeatureCapability { return b }

type BaseDVSFeatureCapability interface {
	GetDVSFeatureCapability() *DVSFeatureCapability
}

func (b *DVSHealthCheckCapability) GetDVSHealthCheckCapability() *DVSHealthCheckCapability { return b }

type BaseDVSHealthCheckCapability interface {
	GetDVSHealthCheckCapability() *DVSHealthCheckCapability
}

func (b *DVSHealthCheckConfig) GetDVSHealthCheckConfig() *DVSHealthCheckConfig { return b }

type BaseDVSHealthCheckConfig interface {
	GetDVSHealthCheckConfig() *DVSHealthCheckConfig
}

func (b *DVSUplinkPortPolicy) GetDVSUplinkPortPolicy() *DVSUplinkPortPolicy { return b }

type BaseDVSUplinkPortPolicy interface {
	GetDVSUplinkPortPolicy() *DVSUplinkPortPolicy
}

func (b *DatastoreInfo) GetDatastoreInfo() *DatastoreInfo { return b }

type BaseDatastoreInfo interface {
	GetDatastoreInfo() *DatastoreInfo
}

func (b *Description) GetDescription() *Description { return b }

type BaseDescription interface {
	GetDescription() *Description
}

func (b *DistributedVirtualSwitchHostMemberBacking) GetDistributedVirtualSwitchHostMemberBacking() *DistributedVirtualSwitchHostMemberBacking {
	return b
}

type BaseDistributedVirtualSwitchHostMemberBacking interface {
	GetDistributedVirtualSwitchHostMemberBacking() *DistributedVirtualSwitchHostMemberBacking
}

func (b *DistributedVirtualSwitchManagerHostDvsFilterSpec) GetDistributedVirtualSwitchManagerHostDvsFilterSpec() *DistributedVirtualSwitchManagerHostDvsFilterSpec {
	return b
}

type BaseDistributedVirtualSwitchManagerHostDvsFilterSpec interface {
	GetDistributedVirtualSwitchManagerHostDvsFilterSpec() *DistributedVirtualSwitchManagerHostDvsFilterSpec
}

func (b *DvsNetworkRuleAction) GetDvsNetworkRuleAction() *DvsNetworkRuleAction { return b }

type BaseDvsNetworkRuleAction interface {
	GetDvsNetworkRuleAction() *DvsNetworkRuleAction
}

func (b *DvsNetworkRuleQualifier) GetDvsNetworkRuleQualifier() *DvsNetworkRuleQualifier { return b }

type BaseDvsNetworkRuleQualifier interface {
	GetDvsNetworkRuleQualifier() *DvsNetworkRuleQualifier
}

func (b *DynamicData) GetDynamicData() *DynamicData { return b }

type BaseDynamicData interface {
	GetDynamicData() *DynamicData
}

func (b *Event) GetEvent() *Event { return b }

type BaseEvent interface {
	GetEvent() *Event
}

func (b *FaultToleranceConfigInfo) GetFaultToleranceConfigInfo() *FaultToleranceConfigInfo { return b }

type BaseFaultToleranceConfigInfo interface {
	GetFaultToleranceConfigInfo() *FaultToleranceConfigInfo
}

func (b *FileInfo) GetFileInfo() *FileInfo { return b }

type BaseFileInfo interface {
	GetFileInfo() *FileInfo
}

func (b *FileQuery) GetFileQuery() *FileQuery { return b }

type BaseFileQuery interface {
	GetFileQuery() *FileQuery
}

func (b *GuestAuthentication) GetGuestAuthentication() *GuestAuthentication { return b }

type BaseGuestAuthentication interface {
	GetGuestAuthentication() *GuestAuthentication
}

func (b *GuestFileAttributes) GetGuestFileAttributes() *GuestFileAttributes { return b }

type BaseGuestFileAttributes interface {
	GetGuestFileAttributes() *GuestFileAttributes
}

func (b *GuestProgramSpec) GetGuestProgramSpec() *GuestProgramSpec { return b }

type BaseGuestProgramSpec interface {
	GetGuestProgramSpec() *GuestProgramSpec
}

func (b *HostAccountSpec) GetHostAccountSpec() *HostAccountSpec { return b }

type BaseHostAccountSpec interface {
	GetHostAccountSpec() *HostAccountSpec
}

func (b *HostAuthenticationStoreInfo) GetHostAuthenticationStoreInfo() *HostAuthenticationStoreInfo {
	return b
}

type BaseHostAuthenticationStoreInfo interface {
	GetHostAuthenticationStoreInfo() *HostAuthenticationStoreInfo
}

func (b *HostConnectInfoNetworkInfo) GetHostConnectInfoNetworkInfo() *HostConnectInfoNetworkInfo {
	return b
}

type BaseHostConnectInfoNetworkInfo interface {
	GetHostConnectInfoNetworkInfo() *HostConnectInfoNetworkInfo
}

func (b *HostDatastoreConnectInfo) GetHostDatastoreConnectInfo() *HostDatastoreConnectInfo { return b }

type BaseHostDatastoreConnectInfo interface {
	GetHostDatastoreConnectInfo() *HostDatastoreConnectInfo
}

func (b *HostDnsConfig) GetHostDnsConfig() *HostDnsConfig { return b }

type BaseHostDnsConfig interface {
	GetHostDnsConfig() *HostDnsConfig
}

func (b *HostFileSystemVolume) GetHostFileSystemVolume() *HostFileSystemVolume { return b }

type BaseHostFileSystemVolume interface {
	GetHostFileSystemVolume() *HostFileSystemVolume
}

func (b *HostHostBusAdapter) GetHostHostBusAdapter() *HostHostBusAdapter { return b }

type BaseHostHostBusAdapter interface {
	GetHostHostBusAdapter() *HostHostBusAdapter
}

func (b *HostIpRouteConfig) GetHostIpRouteConfig() *HostIpRouteConfig { return b }

type BaseHostIpRouteConfig interface {
	GetHostIpRouteConfig() *HostIpRouteConfig
}

func (b *HostMemberHealthCheckResult) GetHostMemberHealthCheckResult() *HostMemberHealthCheckResult {
	return b
}

type BaseHostMemberHealthCheckResult interface {
	GetHostMemberHealthCheckResult() *HostMemberHealthCheckResult
}

func (b *HostMultipathInfoLogicalUnitPolicy) GetHostMultipathInfoLogicalUnitPolicy() *HostMultipathInfoLogicalUnitPolicy {
	return b
}

type BaseHostMultipathInfoLogicalUnitPolicy interface {
	GetHostMultipathInfoLogicalUnitPolicy() *HostMultipathInfoLogicalUnitPolicy
}

func (b *HostPciPassthruConfig) GetHostPciPassthruConfig() *HostPciPassthruConfig { return b }

type BaseHostPciPassthruConfig interface {
	GetHostPciPassthruConfig() *HostPciPassthruConfig
}

func (b *HostPciPassthruInfo) GetHostPciPassthruInfo() *HostPciPassthruInfo { return b }

type BaseHostPciPassthruInfo interface {
	GetHostPciPassthruInfo() *HostPciPassthruInfo
}

func (b *HostSystemSwapConfigurationSystemSwapOption) GetHostSystemSwapConfigurationSystemSwapOption() *HostSystemSwapConfigurationSystemSwapOption {
	return b
}

type BaseHostSystemSwapConfigurationSystemSwapOption interface {
	GetHostSystemSwapConfigurationSystemSwapOption() *HostSystemSwapConfigurationSystemSwapOption
}

func (b *HostTargetTransport) GetHostTargetTransport() *HostTargetTransport { return b }

type BaseHostTargetTransport interface {
	GetHostTargetTransport() *HostTargetTransport
}

func (b *HostTpmEventDetails) GetHostTpmEventDetails() *HostTpmEventDetails { return b }

type BaseHostTpmEventDetails interface {
	GetHostTpmEventDetails() *HostTpmEventDetails
}

func (b *HostVirtualSwitchBridge) GetHostVirtualSwitchBridge() *HostVirtualSwitchBridge { return b }

type BaseHostVirtualSwitchBridge interface {
	GetHostVirtualSwitchBridge() *HostVirtualSwitchBridge
}

func (b *ImportSpec) GetImportSpec() *ImportSpec { return b }

type BaseImportSpec interface {
	GetImportSpec() *ImportSpec
}

func (b *LicenseSource) GetLicenseSource() *LicenseSource { return b }

type BaseLicenseSource interface {
	GetLicenseSource() *LicenseSource
}

func (b *MethodFault) GetMethodFault() *MethodFault { return b }

type BaseMethodFault interface {
	GetMethodFault() *MethodFault
}

func (b *NetBIOSConfigInfo) GetNetBIOSConfigInfo() *NetBIOSConfigInfo { return b }

type BaseNetBIOSConfigInfo interface {
	GetNetBIOSConfigInfo() *NetBIOSConfigInfo
}

func (b *OptionType) GetOptionType() *OptionType { return b }

type BaseOptionType interface {
	GetOptionType() *OptionType
}

func (b *PerfEntityMetricBase) GetPerfEntityMetricBase() *PerfEntityMetricBase { return b }

type BasePerfEntityMetricBase interface {
	GetPerfEntityMetricBase() *PerfEntityMetricBase
}

func (b *PerfMetricSeries) GetPerfMetricSeries() *PerfMetricSeries { return b }

type BasePerfMetricSeries interface {
	GetPerfMetricSeries() *PerfMetricSeries
}

func (b *PolicyOption) GetPolicyOption() *PolicyOption { return b }

type BasePolicyOption interface {
	GetPolicyOption() *PolicyOption
}

func (b *ProfileConfigInfo) GetProfileConfigInfo() *ProfileConfigInfo { return b }

type BaseProfileConfigInfo interface {
	GetProfileConfigInfo() *ProfileConfigInfo
}

func (b *ProfileCreateSpec) GetProfileCreateSpec() *ProfileCreateSpec { return b }

type BaseProfileCreateSpec interface {
	GetProfileCreateSpec() *ProfileCreateSpec
}

func (b *ProfileExpression) GetProfileExpression() *ProfileExpression { return b }

type BaseProfileExpression interface {
	GetProfileExpression() *ProfileExpression
}

func (b *ProfilePolicyOptionMetadata) GetProfilePolicyOptionMetadata() *ProfilePolicyOptionMetadata {
	return b
}

type BaseProfilePolicyOptionMetadata interface {
	GetProfilePolicyOptionMetadata() *ProfilePolicyOptionMetadata
}

func (b *ResourcePoolSummary) GetResourcePoolSummary() *ResourcePoolSummary { return b }

type BaseResourcePoolSummary interface {
	GetResourcePoolSummary() *ResourcePoolSummary
}

func (b *ScheduledTaskSpec) GetScheduledTaskSpec() *ScheduledTaskSpec { return b }

type BaseScheduledTaskSpec interface {
	GetScheduledTaskSpec() *ScheduledTaskSpec
}

func (b *SelectionSet) GetSelectionSet() *SelectionSet { return b }

type BaseSelectionSet interface {
	GetSelectionSet() *SelectionSet
}

func (b *SelectionSpec) GetSelectionSpec() *SelectionSpec { return b }

type BaseSelectionSpec interface {
	GetSelectionSpec() *SelectionSpec
}

func (b *SessionManagerServiceRequestSpec) GetSessionManagerServiceRequestSpec() *SessionManagerServiceRequestSpec {
	return b
}

type BaseSessionManagerServiceRequestSpec interface {
	GetSessionManagerServiceRequestSpec() *SessionManagerServiceRequestSpec
}

func (b *TaskReason) GetTaskReason() *TaskReason { return b }

type BaseTaskReason interface {
	GetTaskReason() *TaskReason
}

func (b *TaskScheduler) GetTaskScheduler() *TaskScheduler { return b }

type BaseTaskScheduler interface {
	GetTaskScheduler() *TaskScheduler
}

func (b *UserSearchResult) GetUserSearchResult() *UserSearchResult { return b }

type BaseUserSearchResult interface {
	GetUserSearchResult() *UserSearchResult
}

func (b *VirtualDevice) GetVirtualDevice() *VirtualDevice { return b }

type BaseVirtualDevice interface {
	GetVirtualDevice() *VirtualDevice
}

func (b *VirtualDeviceBackingInfo) GetVirtualDeviceBackingInfo() *VirtualDeviceBackingInfo { return b }

type BaseVirtualDeviceBackingInfo interface {
	GetVirtualDeviceBackingInfo() *VirtualDeviceBackingInfo
}

func (b *VirtualDeviceBackingOption) GetVirtualDeviceBackingOption() *VirtualDeviceBackingOption {
	return b
}

type BaseVirtualDeviceBackingOption interface {
	GetVirtualDeviceBackingOption() *VirtualDeviceBackingOption
}

func (b *VirtualDeviceBusSlotInfo) GetVirtualDeviceBusSlotInfo() *VirtualDeviceBusSlotInfo { return b }

type BaseVirtualDeviceBusSlotInfo interface {
	GetVirtualDeviceBusSlotInfo() *VirtualDeviceBusSlotInfo
}

func (b *VirtualDeviceConfigSpec) GetVirtualDeviceConfigSpec() *VirtualDeviceConfigSpec { return b }

type BaseVirtualDeviceConfigSpec interface {
	GetVirtualDeviceConfigSpec() *VirtualDeviceConfigSpec
}

func (b *VirtualDeviceOption) GetVirtualDeviceOption() *VirtualDeviceOption { return b }

type BaseVirtualDeviceOption interface {
	GetVirtualDeviceOption() *VirtualDeviceOption
}

func (b *VirtualDiskSpec) GetVirtualDiskSpec() *VirtualDiskSpec { return b }

type BaseVirtualDiskSpec interface {
	GetVirtualDiskSpec() *VirtualDiskSpec
}

func (b *VirtualMachineBootOptionsBootableDevice) GetVirtualMachineBootOptionsBootableDevice() *VirtualMachineBootOptionsBootableDevice {
	return b
}

type BaseVirtualMachineBootOptionsBootableDevice interface {
	GetVirtualMachineBootOptionsBootableDevice() *VirtualMachineBootOptionsBootableDevice
}

func (b *VirtualMachineDeviceRuntimeInfoDeviceRuntimeState) GetVirtualMachineDeviceRuntimeInfoDeviceRuntimeState() *VirtualMachineDeviceRuntimeInfoDeviceRuntimeState {
	return b
}

type BaseVirtualMachineDeviceRuntimeInfoDeviceRuntimeState interface {
	GetVirtualMachineDeviceRuntimeInfoDeviceRuntimeState() *VirtualMachineDeviceRuntimeInfoDeviceRuntimeState
}

func (b *VirtualMachineProfileSpec) GetVirtualMachineProfileSpec() *VirtualMachineProfileSpec {
	return b
}

type BaseVirtualMachineProfileSpec interface {
	GetVirtualMachineProfileSpec() *VirtualMachineProfileSpec
}

func (b *VirtualMachineTargetInfo) GetVirtualMachineTargetInfo() *VirtualMachineTargetInfo { return b }

type BaseVirtualMachineTargetInfo interface {
	GetVirtualMachineTargetInfo() *VirtualMachineTargetInfo
}

func (b *VmConfigInfo) GetVmConfigInfo() *VmConfigInfo { return b }

type BaseVmConfigInfo interface {
	GetVmConfigInfo() *VmConfigInfo
}

func (b *VmfsDatastoreBaseOption) GetVmfsDatastoreBaseOption() *VmfsDatastoreBaseOption { return b }

type BaseVmfsDatastoreBaseOption interface {
	GetVmfsDatastoreBaseOption() *VmfsDatastoreBaseOption
}
