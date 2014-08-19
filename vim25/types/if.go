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

func (b *Action) isAction() {}

type BaseAction interface {
	isAction()
}

func (b *AlarmAction) isAlarmAction() {}

type BaseAlarmAction interface {
	isAlarmAction()
}

func (b *AlarmExpression) isAlarmExpression() {}

type BaseAlarmExpression interface {
	isAlarmExpression()
}

func (b *AlarmSpec) isAlarmSpec() {}

type BaseAlarmSpec interface {
	isAlarmSpec()
}

func (b *AnswerFileCreateSpec) isAnswerFileCreateSpec() {}

type BaseAnswerFileCreateSpec interface {
	isAnswerFileCreateSpec()
}

func (b *ApplyProfile) isApplyProfile() {}

type BaseApplyProfile interface {
	isApplyProfile()
}

func (b *ClusterAction) isClusterAction() {}

type BaseClusterAction interface {
	isClusterAction()
}

func (b *ClusterDasAdmissionControlInfo) isClusterDasAdmissionControlInfo() {}

type BaseClusterDasAdmissionControlInfo interface {
	isClusterDasAdmissionControlInfo()
}

func (b *ClusterDasAdmissionControlPolicy) isClusterDasAdmissionControlPolicy() {}

type BaseClusterDasAdmissionControlPolicy interface {
	isClusterDasAdmissionControlPolicy()
}

func (b *ClusterDasAdvancedRuntimeInfo) isClusterDasAdvancedRuntimeInfo() {}

type BaseClusterDasAdvancedRuntimeInfo interface {
	isClusterDasAdvancedRuntimeInfo()
}

func (b *ClusterDasData) isClusterDasData() {}

type BaseClusterDasData interface {
	isClusterDasData()
}

func (b *ClusterDasHostInfo) isClusterDasHostInfo() {}

type BaseClusterDasHostInfo interface {
	isClusterDasHostInfo()
}

func (b *ClusterDrsFaultsFaultsByVm) isClusterDrsFaultsFaultsByVm() {}

type BaseClusterDrsFaultsFaultsByVm interface {
	isClusterDrsFaultsFaultsByVm()
}

func (b *ClusterGroupInfo) isClusterGroupInfo() {}

type BaseClusterGroupInfo interface {
	isClusterGroupInfo()
}

func (b *ClusterRuleInfo) isClusterRuleInfo() {}

type BaseClusterRuleInfo interface {
	isClusterRuleInfo()
}

func (b *ClusterSlotPolicy) isClusterSlotPolicy() {}

type BaseClusterSlotPolicy interface {
	isClusterSlotPolicy()
}

func (b *ComputeResourceConfigInfo) isComputeResourceConfigInfo() {}

type BaseComputeResourceConfigInfo interface {
	isComputeResourceConfigInfo()
}

func (b *ComputeResourceSummary) isComputeResourceSummary() {}

type BaseComputeResourceSummary interface {
	isComputeResourceSummary()
}

func (b *CustomFieldValue) isCustomFieldValue() {}

type BaseCustomFieldValue interface {
	isCustomFieldValue()
}

func (b *CustomizationIdentitySettings) isCustomizationIdentitySettings() {}

type BaseCustomizationIdentitySettings interface {
	isCustomizationIdentitySettings()
}

func (b *CustomizationIpGenerator) isCustomizationIpGenerator() {}

type BaseCustomizationIpGenerator interface {
	isCustomizationIpGenerator()
}

func (b *CustomizationIpV6Generator) isCustomizationIpV6Generator() {}

type BaseCustomizationIpV6Generator interface {
	isCustomizationIpV6Generator()
}

func (b *CustomizationName) isCustomizationName() {}

type BaseCustomizationName interface {
	isCustomizationName()
}

func (b *CustomizationOptions) isCustomizationOptions() {}

type BaseCustomizationOptions interface {
	isCustomizationOptions()
}

func (b *DVPortSetting) isDVPortSetting() {}

type BaseDVPortSetting interface {
	isDVPortSetting()
}

func (b *DVPortgroupPolicy) isDVPortgroupPolicy() {}

type BaseDVPortgroupPolicy interface {
	isDVPortgroupPolicy()
}

func (b *DVSConfigInfo) isDVSConfigInfo() {}

type BaseDVSConfigInfo interface {
	isDVSConfigInfo()
}

func (b *DVSConfigSpec) isDVSConfigSpec() {}

type BaseDVSConfigSpec interface {
	isDVSConfigSpec()
}

func (b *DVSFeatureCapability) isDVSFeatureCapability() {}

type BaseDVSFeatureCapability interface {
	isDVSFeatureCapability()
}

func (b *DVSHealthCheckCapability) isDVSHealthCheckCapability() {}

type BaseDVSHealthCheckCapability interface {
	isDVSHealthCheckCapability()
}

func (b *DVSHealthCheckConfig) isDVSHealthCheckConfig() {}

type BaseDVSHealthCheckConfig interface {
	isDVSHealthCheckConfig()
}

func (b *DVSUplinkPortPolicy) isDVSUplinkPortPolicy() {}

type BaseDVSUplinkPortPolicy interface {
	isDVSUplinkPortPolicy()
}

func (b *DatastoreInfo) isDatastoreInfo() {}

type BaseDatastoreInfo interface {
	isDatastoreInfo()
}

func (b *Description) isDescription() {}

type BaseDescription interface {
	isDescription()
}

func (b *DistributedVirtualSwitchHostMemberBacking) isDistributedVirtualSwitchHostMemberBacking() {}

type BaseDistributedVirtualSwitchHostMemberBacking interface {
	isDistributedVirtualSwitchHostMemberBacking()
}

func (b *DistributedVirtualSwitchManagerHostDvsFilterSpec) isDistributedVirtualSwitchManagerHostDvsFilterSpec() {
}

type BaseDistributedVirtualSwitchManagerHostDvsFilterSpec interface {
	isDistributedVirtualSwitchManagerHostDvsFilterSpec()
}

func (b *DvsNetworkRuleAction) isDvsNetworkRuleAction() {}

type BaseDvsNetworkRuleAction interface {
	isDvsNetworkRuleAction()
}

func (b *DvsNetworkRuleQualifier) isDvsNetworkRuleQualifier() {}

type BaseDvsNetworkRuleQualifier interface {
	isDvsNetworkRuleQualifier()
}

func (b *DynamicData) isDynamicData() {}

type BaseDynamicData interface {
	isDynamicData()
}

func (b *Event) isEvent() {}

type BaseEvent interface {
	isEvent()
}

func (b *FaultToleranceConfigInfo) isFaultToleranceConfigInfo() {}

type BaseFaultToleranceConfigInfo interface {
	isFaultToleranceConfigInfo()
}

func (b *FileInfo) isFileInfo() {}

type BaseFileInfo interface {
	isFileInfo()
}

func (b *FileQuery) isFileQuery() {}

type BaseFileQuery interface {
	isFileQuery()
}

func (b *GuestAuthentication) isGuestAuthentication() {}

type BaseGuestAuthentication interface {
	isGuestAuthentication()
}

func (b *GuestFileAttributes) isGuestFileAttributes() {}

type BaseGuestFileAttributes interface {
	isGuestFileAttributes()
}

func (b *GuestProgramSpec) isGuestProgramSpec() {}

type BaseGuestProgramSpec interface {
	isGuestProgramSpec()
}

func (b *HostAccountSpec) isHostAccountSpec() {}

type BaseHostAccountSpec interface {
	isHostAccountSpec()
}

func (b *HostAuthenticationStoreInfo) isHostAuthenticationStoreInfo() {}

type BaseHostAuthenticationStoreInfo interface {
	isHostAuthenticationStoreInfo()
}

func (b *HostConnectInfoNetworkInfo) isHostConnectInfoNetworkInfo() {}

type BaseHostConnectInfoNetworkInfo interface {
	isHostConnectInfoNetworkInfo()
}

func (b *HostDatastoreConnectInfo) isHostDatastoreConnectInfo() {}

type BaseHostDatastoreConnectInfo interface {
	isHostDatastoreConnectInfo()
}

func (b *HostDnsConfig) isHostDnsConfig() {}

type BaseHostDnsConfig interface {
	isHostDnsConfig()
}

func (b *HostFileSystemVolume) isHostFileSystemVolume() {}

type BaseHostFileSystemVolume interface {
	isHostFileSystemVolume()
}

func (b *HostHostBusAdapter) isHostHostBusAdapter() {}

type BaseHostHostBusAdapter interface {
	isHostHostBusAdapter()
}

func (b *HostIpRouteConfig) isHostIpRouteConfig() {}

type BaseHostIpRouteConfig interface {
	isHostIpRouteConfig()
}

func (b *HostMemberHealthCheckResult) isHostMemberHealthCheckResult() {}

type BaseHostMemberHealthCheckResult interface {
	isHostMemberHealthCheckResult()
}

func (b *HostMultipathInfoLogicalUnitPolicy) isHostMultipathInfoLogicalUnitPolicy() {}

type BaseHostMultipathInfoLogicalUnitPolicy interface {
	isHostMultipathInfoLogicalUnitPolicy()
}

func (b *HostPciPassthruConfig) isHostPciPassthruConfig() {}

type BaseHostPciPassthruConfig interface {
	isHostPciPassthruConfig()
}

func (b *HostPciPassthruInfo) isHostPciPassthruInfo() {}

type BaseHostPciPassthruInfo interface {
	isHostPciPassthruInfo()
}

func (b *HostSystemSwapConfigurationSystemSwapOption) isHostSystemSwapConfigurationSystemSwapOption() {
}

type BaseHostSystemSwapConfigurationSystemSwapOption interface {
	isHostSystemSwapConfigurationSystemSwapOption()
}

func (b *HostTargetTransport) isHostTargetTransport() {}

type BaseHostTargetTransport interface {
	isHostTargetTransport()
}

func (b *HostTpmEventDetails) isHostTpmEventDetails() {}

type BaseHostTpmEventDetails interface {
	isHostTpmEventDetails()
}

func (b *HostVirtualSwitchBridge) isHostVirtualSwitchBridge() {}

type BaseHostVirtualSwitchBridge interface {
	isHostVirtualSwitchBridge()
}

func (b *ImportSpec) isImportSpec() {}

type BaseImportSpec interface {
	isImportSpec()
}

func (b *LicenseSource) isLicenseSource() {}

type BaseLicenseSource interface {
	isLicenseSource()
}

func (b *MethodFault) isMethodFault() {}

type BaseMethodFault interface {
	isMethodFault()
}

func (b *NetBIOSConfigInfo) isNetBIOSConfigInfo() {}

type BaseNetBIOSConfigInfo interface {
	isNetBIOSConfigInfo()
}

func (b *OptionType) isOptionType() {}

type BaseOptionType interface {
	isOptionType()
}

func (b *PerfEntityMetricBase) isPerfEntityMetricBase() {}

type BasePerfEntityMetricBase interface {
	isPerfEntityMetricBase()
}

func (b *PerfMetricSeries) isPerfMetricSeries() {}

type BasePerfMetricSeries interface {
	isPerfMetricSeries()
}

func (b *PolicyOption) isPolicyOption() {}

type BasePolicyOption interface {
	isPolicyOption()
}

func (b *ProfileConfigInfo) isProfileConfigInfo() {}

type BaseProfileConfigInfo interface {
	isProfileConfigInfo()
}

func (b *ProfileCreateSpec) isProfileCreateSpec() {}

type BaseProfileCreateSpec interface {
	isProfileCreateSpec()
}

func (b *ProfileExpression) isProfileExpression() {}

type BaseProfileExpression interface {
	isProfileExpression()
}

func (b *ProfilePolicyOptionMetadata) isProfilePolicyOptionMetadata() {}

type BaseProfilePolicyOptionMetadata interface {
	isProfilePolicyOptionMetadata()
}

func (b *ResourcePoolSummary) isResourcePoolSummary() {}

type BaseResourcePoolSummary interface {
	isResourcePoolSummary()
}

func (b *ScheduledTaskSpec) isScheduledTaskSpec() {}

type BaseScheduledTaskSpec interface {
	isScheduledTaskSpec()
}

func (b *SelectionSet) isSelectionSet() {}

type BaseSelectionSet interface {
	isSelectionSet()
}

func (b *SelectionSpec) isSelectionSpec() {}

type BaseSelectionSpec interface {
	isSelectionSpec()
}

func (b *SessionManagerServiceRequestSpec) isSessionManagerServiceRequestSpec() {}

type BaseSessionManagerServiceRequestSpec interface {
	isSessionManagerServiceRequestSpec()
}

func (b *TaskReason) isTaskReason() {}

type BaseTaskReason interface {
	isTaskReason()
}

func (b *TaskScheduler) isTaskScheduler() {}

type BaseTaskScheduler interface {
	isTaskScheduler()
}

func (b *UserSearchResult) isUserSearchResult() {}

type BaseUserSearchResult interface {
	isUserSearchResult()
}

func (b *VirtualDevice) isVirtualDevice() {}

type BaseVirtualDevice interface {
	isVirtualDevice()
}

func (b *VirtualDeviceBackingInfo) isVirtualDeviceBackingInfo() {}

type BaseVirtualDeviceBackingInfo interface {
	isVirtualDeviceBackingInfo()
}

func (b *VirtualDeviceBackingOption) isVirtualDeviceBackingOption() {}

type BaseVirtualDeviceBackingOption interface {
	isVirtualDeviceBackingOption()
}

func (b *VirtualDeviceBusSlotInfo) isVirtualDeviceBusSlotInfo() {}

type BaseVirtualDeviceBusSlotInfo interface {
	isVirtualDeviceBusSlotInfo()
}

func (b *VirtualDeviceConfigSpec) isVirtualDeviceConfigSpec() {}

type BaseVirtualDeviceConfigSpec interface {
	isVirtualDeviceConfigSpec()
}

func (b *VirtualDeviceOption) isVirtualDeviceOption() {}

type BaseVirtualDeviceOption interface {
	isVirtualDeviceOption()
}

func (b *VirtualDiskSpec) isVirtualDiskSpec() {}

type BaseVirtualDiskSpec interface {
	isVirtualDiskSpec()
}

func (b *VirtualMachineBootOptionsBootableDevice) isVirtualMachineBootOptionsBootableDevice() {}

type BaseVirtualMachineBootOptionsBootableDevice interface {
	isVirtualMachineBootOptionsBootableDevice()
}

func (b *VirtualMachineDeviceRuntimeInfoDeviceRuntimeState) isVirtualMachineDeviceRuntimeInfoDeviceRuntimeState() {
}

type BaseVirtualMachineDeviceRuntimeInfoDeviceRuntimeState interface {
	isVirtualMachineDeviceRuntimeInfoDeviceRuntimeState()
}

func (b *VirtualMachineProfileSpec) isVirtualMachineProfileSpec() {}

type BaseVirtualMachineProfileSpec interface {
	isVirtualMachineProfileSpec()
}

func (b *VirtualMachineTargetInfo) isVirtualMachineTargetInfo() {}

type BaseVirtualMachineTargetInfo interface {
	isVirtualMachineTargetInfo()
}

func (b *VmConfigInfo) isVmConfigInfo() {}

type BaseVmConfigInfo interface {
	isVmConfigInfo()
}

func (b *VmfsDatastoreBaseOption) isVmfsDatastoreBaseOption() {}

type BaseVmfsDatastoreBaseOption interface {
	isVmfsDatastoreBaseOption()
}
