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

	"github.com/vmware/govmomi/vim25/types"
)

type CnsClusterFlavor string

const (
	CnsClusterFlavorVANILLA               = CnsClusterFlavor("VANILLA")
	CnsClusterFlavorWORKLOAD              = CnsClusterFlavor("WORKLOAD")
	CnsClusterFlavorClusterFlavor_Unknown = CnsClusterFlavor("ClusterFlavor_Unknown")
	CnsClusterFlavorGUEST_CLUSTER         = CnsClusterFlavor("GUEST_CLUSTER")
)

func init() {
	types.Add("vsan:CnsClusterFlavor", reflect.TypeOf((*CnsClusterFlavor)(nil)).Elem())
}

type CnsClusterType string

const (
	CnsClusterTypeClusterType_Unknown = CnsClusterType("ClusterType_Unknown")
	CnsClusterTypeKUBERNETES          = CnsClusterType("KUBERNETES")
)

func init() {
	types.Add("vsan:CnsClusterType", reflect.TypeOf((*CnsClusterType)(nil)).Elem())
}

type CnsKubernetesEntityType string

const (
	CnsKubernetesEntityTypePERSISTENT_VOLUME            = CnsKubernetesEntityType("PERSISTENT_VOLUME")
	CnsKubernetesEntityTypePERSISTENT_VOLUME_CLAIM      = CnsKubernetesEntityType("PERSISTENT_VOLUME_CLAIM")
	CnsKubernetesEntityTypePOD                          = CnsKubernetesEntityType("POD")
	CnsKubernetesEntityTypeKubernetesEntityType_Unknown = CnsKubernetesEntityType("KubernetesEntityType_Unknown")
)

func init() {
	types.Add("vsan:CnsKubernetesEntityType", reflect.TypeOf((*CnsKubernetesEntityType)(nil)).Elem())
}

type CnsVolumeType string

const (
	CnsVolumeTypeFILE               = CnsVolumeType("FILE")
	CnsVolumeTypeBLOCK              = CnsVolumeType("BLOCK")
	CnsVolumeTypeVolumeType_Unknown = CnsVolumeType("VolumeType_Unknown")
)

func init() {
	types.Add("vsan:CnsVolumeType", reflect.TypeOf((*CnsVolumeType)(nil)).Elem())
}

type QuerySelectionNameType string

const (
	QuerySelectionNameTypeBACKING_OBJECT_DETAILS         = QuerySelectionNameType("BACKING_OBJECT_DETAILS")
	QuerySelectionNameTypeCOMPLIANCE_STATUS              = QuerySelectionNameType("COMPLIANCE_STATUS")
	QuerySelectionNameTypeVOLUME_TYPE                    = QuerySelectionNameType("VOLUME_TYPE")
	QuerySelectionNameTypeHEALTH_STATUS                  = QuerySelectionNameType("HEALTH_STATUS")
	QuerySelectionNameTypeVOLUME_NAME                    = QuerySelectionNameType("VOLUME_NAME")
	QuerySelectionNameTypeDATASTORE_ACCESSIBILITY_STATUS = QuerySelectionNameType("DATASTORE_ACCESSIBILITY_STATUS")
	QuerySelectionNameTypeQuerySelectionNameType_Unknown = QuerySelectionNameType("QuerySelectionNameType_Unknown")
)

func init() {
	types.Add("vsan:QuerySelectionNameType", reflect.TypeOf((*QuerySelectionNameType)(nil)).Elem())
}

type VimClusterVSANStretchedClusterConfigIssueInfoVSANStretchedClusterConfigIssueEnum string

const (
	VimClusterVSANStretchedClusterConfigIssueInfoVSANStretchedClusterConfigIssueEnumClusterWithMultipleUnicastAgents            = VimClusterVSANStretchedClusterConfigIssueInfoVSANStretchedClusterConfigIssueEnum("ClusterWithMultipleUnicastAgents")
	VimClusterVSANStretchedClusterConfigIssueInfoVSANStretchedClusterConfigIssueEnumWitnessFaultDomainInvalid                   = VimClusterVSANStretchedClusterConfigIssueInfoVSANStretchedClusterConfigIssueEnum("WitnessFaultDomainInvalid")
	VimClusterVSANStretchedClusterConfigIssueInfoVSANStretchedClusterConfigIssueEnumClusterWithoutOneWitnessHost                = VimClusterVSANStretchedClusterConfigIssueInfoVSANStretchedClusterConfigIssueEnum("ClusterWithoutOneWitnessHost")
	VimClusterVSANStretchedClusterConfigIssueInfoVSANStretchedClusterConfigIssueEnumWitnessPreferredFaultDomainNotExist         = VimClusterVSANStretchedClusterConfigIssueInfoVSANStretchedClusterConfigIssueEnum("WitnessPreferredFaultDomainNotExist")
	VimClusterVSANStretchedClusterConfigIssueInfoVSANStretchedClusterConfigIssueEnumClusterWithoutTwoDataFaultDomains           = VimClusterVSANStretchedClusterConfigIssueInfoVSANStretchedClusterConfigIssueEnum("ClusterWithoutTwoDataFaultDomains")
	VimClusterVSANStretchedClusterConfigIssueInfoVSANStretchedClusterConfigIssueEnumHostUnicastAgentUnset                       = VimClusterVSANStretchedClusterConfigIssueInfoVSANStretchedClusterConfigIssueEnum("HostUnicastAgentUnset")
	VimClusterVSANStretchedClusterConfigIssueInfoVSANStretchedClusterConfigIssueEnumHostWithNoStretchedClusterSupport           = VimClusterVSANStretchedClusterConfigIssueInfoVSANStretchedClusterConfigIssueEnum("HostWithNoStretchedClusterSupport")
	VimClusterVSANStretchedClusterConfigIssueInfoVSANStretchedClusterConfigIssueEnumHostWithInvalidUnicastAgent                 = VimClusterVSANStretchedClusterConfigIssueInfoVSANStretchedClusterConfigIssueEnum("HostWithInvalidUnicastAgent")
	VimClusterVSANStretchedClusterConfigIssueInfoVSANStretchedClusterConfigIssueEnumWitnessPreferredFaultDomainInvalid          = VimClusterVSANStretchedClusterConfigIssueInfoVSANStretchedClusterConfigIssueEnum("WitnessPreferredFaultDomainInvalid")
	VimClusterVSANStretchedClusterConfigIssueInfoVSANStretchedClusterConfigIssueEnumWitnessInsideVcCluster                      = VimClusterVSANStretchedClusterConfigIssueInfoVSANStretchedClusterConfigIssueEnum("WitnessInsideVcCluster")
	VimClusterVSANStretchedClusterConfigIssueInfoVSANStretchedClusterConfigIssueEnumVSANStretchedClusterConfigIssueEnum_Unknown = VimClusterVSANStretchedClusterConfigIssueInfoVSANStretchedClusterConfigIssueEnum("VSANStretchedClusterConfigIssueEnum_Unknown")
	VimClusterVSANStretchedClusterConfigIssueInfoVSANStretchedClusterConfigIssueEnumWitnessWithNoDiskMapping                    = VimClusterVSANStretchedClusterConfigIssueInfoVSANStretchedClusterConfigIssueEnum("WitnessWithNoDiskMapping")
)

func init() {
	types.Add("vsan:VimClusterVSANStretchedClusterConfigIssueInfoVSANStretchedClusterConfigIssueEnum", reflect.TypeOf((*VimClusterVSANStretchedClusterConfigIssueInfoVSANStretchedClusterConfigIssueEnum)(nil)).Elem())
}

type VimClusterVsanHostDiskMappingVsanDiskGroupCreationType string

const (
	VimClusterVsanHostDiskMappingVsanDiskGroupCreationTypeAllflash                          = VimClusterVsanHostDiskMappingVsanDiskGroupCreationType("allflash")
	VimClusterVsanHostDiskMappingVsanDiskGroupCreationTypeHybrid                            = VimClusterVsanHostDiskMappingVsanDiskGroupCreationType("hybrid")
	VimClusterVsanHostDiskMappingVsanDiskGroupCreationTypeVsanDiskGroupCreationType_Unknown = VimClusterVsanHostDiskMappingVsanDiskGroupCreationType("VsanDiskGroupCreationType_Unknown")
)

func init() {
	types.Add("vsan:VimClusterVsanHostDiskMappingVsanDiskGroupCreationType", reflect.TypeOf((*VimClusterVsanHostDiskMappingVsanDiskGroupCreationType)(nil)).Elem())
}

type VimVsanClusterComplianceResourceCheckStatusType string

const (
	VimVsanClusterComplianceResourceCheckStatusTypeUninitialized                             = VimVsanClusterComplianceResourceCheckStatusType("uninitialized")
	VimVsanClusterComplianceResourceCheckStatusTypeInProgress                                = VimVsanClusterComplianceResourceCheckStatusType("inProgress")
	VimVsanClusterComplianceResourceCheckStatusTypeComplianceResourceCheckStatusType_Unknown = VimVsanClusterComplianceResourceCheckStatusType("ComplianceResourceCheckStatusType_Unknown")
	VimVsanClusterComplianceResourceCheckStatusTypeCompleted                                 = VimVsanClusterComplianceResourceCheckStatusType("completed")
	VimVsanClusterComplianceResourceCheckStatusTypeAborted                                   = VimVsanClusterComplianceResourceCheckStatusType("aborted")
)

func init() {
	types.Add("vsan:VimVsanClusterComplianceResourceCheckStatusType", reflect.TypeOf((*VimVsanClusterComplianceResourceCheckStatusType)(nil)).Elem())
}

type VimVsanHostDiskMappingCreationSpecDiskMappingCreationType string

const (
	VimVsanHostDiskMappingCreationSpecDiskMappingCreationTypeDiskMappingCreationType_Unknown = VimVsanHostDiskMappingCreationSpecDiskMappingCreationType("DiskMappingCreationType_Unknown")
	VimVsanHostDiskMappingCreationSpecDiskMappingCreationTypeAllFlash                        = VimVsanHostDiskMappingCreationSpecDiskMappingCreationType("allFlash")
	VimVsanHostDiskMappingCreationSpecDiskMappingCreationTypeHybrid                          = VimVsanHostDiskMappingCreationSpecDiskMappingCreationType("hybrid")
)

func init() {
	types.Add("vsan:VimVsanHostDiskMappingCreationSpecDiskMappingCreationType", reflect.TypeOf((*VimVsanHostDiskMappingCreationSpecDiskMappingCreationType)(nil)).Elem())
}

type VimVsanVsanScanObjectsIssueVsanScanObjectsIssueType string

const (
	VimVsanVsanScanObjectsIssueVsanScanObjectsIssueTypeUNKNOWN       = VimVsanVsanScanObjectsIssueVsanScanObjectsIssueType("UNKNOWN")
	VimVsanVsanScanObjectsIssueVsanScanObjectsIssueTypeBROKEN_CHAIN  = VimVsanVsanScanObjectsIssueVsanScanObjectsIssueType("BROKEN_CHAIN")
	VimVsanVsanScanObjectsIssueVsanScanObjectsIssueTypeLEAKED_OBJECT = VimVsanVsanScanObjectsIssueVsanScanObjectsIssueType("LEAKED_OBJECT")
)

func init() {
	types.Add("vsan:VimVsanVsanScanObjectsIssueVsanScanObjectsIssueType", reflect.TypeOf((*VimVsanVsanScanObjectsIssueVsanScanObjectsIssueType)(nil)).Elem())
}

type VimVsanVsanVcsaDeploymentPhase string

const (
	VimVsanVsanVcsaDeploymentPhaseFailed                          = VimVsanVsanVcsaDeploymentPhase("failed")
	VimVsanVsanVcsaDeploymentPhaseVcsadeploy                      = VimVsanVsanVcsaDeploymentPhase("vcsadeploy")
	VimVsanVsanVcsaDeploymentPhaseOvaunpack                       = VimVsanVsanVcsaDeploymentPhase("ovaunpack")
	VimVsanVsanVcsaDeploymentPhaseDone                            = VimVsanVsanVcsaDeploymentPhase("done")
	VimVsanVsanVcsaDeploymentPhaseVsanVcsaDeploymentPhase_Unknown = VimVsanVsanVcsaDeploymentPhase("VsanVcsaDeploymentPhase_Unknown")
	VimVsanVsanVcsaDeploymentPhaseInitializing                    = VimVsanVsanVcsaDeploymentPhase("initializing")
	VimVsanVsanVcsaDeploymentPhaseValidation                      = VimVsanVsanVcsaDeploymentPhase("validation")
	VimVsanVsanVcsaDeploymentPhaseVcconfig                        = VimVsanVsanVcsaDeploymentPhase("vcconfig")
	VimVsanVsanVcsaDeploymentPhaseVsanbootstrap                   = VimVsanVsanVcsaDeploymentPhase("vsanbootstrap")
)

func init() {
	types.Add("vsan:VimVsanVsanVcsaDeploymentPhase", reflect.TypeOf((*VimVsanVsanVcsaDeploymentPhase)(nil)).Elem())
}

type VsanBaselinePreferenceType string

const (
	VsanBaselinePreferenceTypeNoRecommendation                   = VsanBaselinePreferenceType("noRecommendation")
	VsanBaselinePreferenceTypeLatestRelease                      = VsanBaselinePreferenceType("latestRelease")
	VsanBaselinePreferenceTypeLatestPatch                        = VsanBaselinePreferenceType("latestPatch")
	VsanBaselinePreferenceTypeVsanBaselinePreferenceType_Unknown = VsanBaselinePreferenceType("VsanBaselinePreferenceType_Unknown")
)

func init() {
	types.Add("vsan:VsanBaselinePreferenceType", reflect.TypeOf((*VsanBaselinePreferenceType)(nil)).Elem())
}

type VsanCapabilityStatus string

const (
	VsanCapabilityStatusUnknown      = VsanCapabilityStatus("unknown")
	VsanCapabilityStatusCalculated   = VsanCapabilityStatus("calculated")
	VsanCapabilityStatusDisconnected = VsanCapabilityStatus("disconnected")
	VsanCapabilityStatusOldversion   = VsanCapabilityStatus("oldversion")
)

func init() {
	types.Add("vsan:VsanCapabilityStatus", reflect.TypeOf((*VsanCapabilityStatus)(nil)).Elem())
}

type VsanCapabilityType string

const (
	VsanCapabilityTypeDiagnosticmode                 = VsanCapabilityType("diagnosticmode")
	VsanCapabilityTypeObjectidentities               = VsanCapabilityType("objectidentities")
	VsanCapabilityTypeSharedwitness                  = VsanCapabilityType("sharedwitness")
	VsanCapabilityTypeVumbaselinerecommendation      = VsanCapabilityType("vumbaselinerecommendation")
	VsanCapabilityTypeUpgrade                        = VsanCapabilityType("upgrade")
	VsanCapabilityTypeVitstretchedcluster            = VsanCapabilityType("vitstretchedcluster")
	VsanCapabilityTypeEnhancedresyncapi              = VsanCapabilityType("enhancedresyncapi")
	VsanCapabilityTypeCnsvolumes                     = VsanCapabilityType("cnsvolumes")
	VsanCapabilityTypeThrottleresync                 = VsanCapabilityType("throttleresync")
	VsanCapabilityTypeVerbosemodeconfiguration       = VsanCapabilityType("verbosemodeconfiguration")
	VsanCapabilityTypeLargecapacitydrive             = VsanCapabilityType("largecapacitydrive")
	VsanCapabilityTypeIscsitargets                   = VsanCapabilityType("iscsitargets")
	VsanCapabilityTypePurgeinaccessiblevmswapobjects = VsanCapabilityType("purgeinaccessiblevmswapobjects")
	VsanCapabilityTypeResyncetaimprovement           = VsanCapabilityType("resyncetaimprovement")
	VsanCapabilityTypeVmlevelcapacity                = VsanCapabilityType("vmlevelcapacity")
	VsanCapabilityTypeVitonlineresize                = VsanCapabilityType("vitonlineresize")
	VsanCapabilityTypeVsanrdma                       = VsanCapabilityType("vsanrdma")
	VsanCapabilityTypeDataefficiency                 = VsanCapabilityType("dataefficiency")
	VsanCapabilityTypeMetricsconfig                  = VsanCapabilityType("metricsconfig")
	VsanCapabilityTypeHistoricalcapacity             = VsanCapabilityType("historicalcapacity")
	VsanCapabilityTypeAllflash                       = VsanCapabilityType("allflash")
	VsanCapabilityTypeIoinsight                      = VsanCapabilityType("ioinsight")
	VsanCapabilityTypeUnicasttest                    = VsanCapabilityType("unicasttest")
	VsanCapabilityTypeWcpappplatform                 = VsanCapabilityType("wcpappplatform")
	VsanCapabilityTypeFileservicesmb                 = VsanCapabilityType("fileservicesmb")
	VsanCapabilityTypeNestedfd                       = VsanCapabilityType("nestedfd")
	VsanCapabilityTypePr1741414fixed                 = VsanCapabilityType("pr1741414fixed")
	VsanCapabilityTypeGethcllastupdateonvc           = VsanCapabilityType("gethcllastupdateonvc")
	VsanCapabilityTypeCapability                     = VsanCapabilityType("capability")
	VsanCapabilityTypeDecomwhatif                    = VsanCapabilityType("decomwhatif")
	VsanCapabilityTypeClusterconfig                  = VsanCapabilityType("clusterconfig")
	VsanCapabilityTypePolicyassociation              = VsanCapabilityType("policyassociation")
	VsanCapabilityTypeSupportinsight                 = VsanCapabilityType("supportinsight")
	VsanCapabilityTypePerfsvcautoconfig              = VsanCapabilityType("perfsvcautoconfig")
	VsanCapabilityTypeGenericnestedfd                = VsanCapabilityType("genericnestedfd")
	VsanCapabilityTypePerfsvcverbosemode             = VsanCapabilityType("perfsvcverbosemode")
	VsanCapabilityTypeFilevolumes                    = VsanCapabilityType("filevolumes")
	VsanCapabilityTypeUpdatevumreleasecatalogoffline = VsanCapabilityType("updatevumreleasecatalogoffline")
	VsanCapabilityTypeResourceprecheck               = VsanCapabilityType("resourceprecheck")
	VsanCapabilityTypeUnicastmode                    = VsanCapabilityType("unicastmode")
	VsanCapabilityTypeHardwaremgmt                   = VsanCapabilityType("hardwaremgmt")
	VsanCapabilityTypeHealthcheck2018q2              = VsanCapabilityType("healthcheck2018q2")
	VsanCapabilityTypePerformanceforsupport          = VsanCapabilityType("performanceforsupport")
	VsanCapabilityTypeFirmwareupdate                 = VsanCapabilityType("firmwareupdate")
	VsanCapabilityTypeImprovedcapacityscreen         = VsanCapabilityType("improvedcapacityscreen")
	VsanCapabilityTypeDiskresourceprecheck           = VsanCapabilityType("diskresourceprecheck")
	VsanCapabilityTypeDevice4ksupport                = VsanCapabilityType("device4ksupport")
	VsanCapabilityTypeFullStackFw                    = VsanCapabilityType("fullStackFw")
	VsanCapabilityTypeMasspropertycollector          = VsanCapabilityType("masspropertycollector")
	VsanCapabilityTypeNondatamovementdfc             = VsanCapabilityType("nondatamovementdfc")
	VsanCapabilityTypeVumintegration                 = VsanCapabilityType("vumintegration")
	VsanCapabilityTypeRemotedatastore                = VsanCapabilityType("remotedatastore")
	VsanCapabilityTypeEncryption                     = VsanCapabilityType("encryption")
	VsanCapabilityTypeHostreservedcapacity           = VsanCapabilityType("hostreservedcapacity")
	VsanCapabilityTypeFileservicenfsv3               = VsanCapabilityType("fileservicenfsv3")
	VsanCapabilityTypeNetperftest                    = VsanCapabilityType("netperftest")
	VsanCapabilityTypeSlackspacecapacity             = VsanCapabilityType("slackspacecapacity")
	VsanCapabilityTypeWhatifcapacity                 = VsanCapabilityType("whatifcapacity")
	VsanCapabilityTypeAutomaticrebalance             = VsanCapabilityType("automaticrebalance")
	VsanCapabilityTypeUmap                           = VsanCapabilityType("umap")
	VsanCapabilityTypeFileservicekerberos            = VsanCapabilityType("fileservicekerberos")
	VsanCapabilityTypeDataintransitencryption        = VsanCapabilityType("dataintransitencryption")
	VsanCapabilityTypeRecreatediskgroup              = VsanCapabilityType("recreatediskgroup")
	VsanCapabilityTypeConfigassist                   = VsanCapabilityType("configassist")
	VsanCapabilityTypeUpgraderesourceprecheck        = VsanCapabilityType("upgraderesourceprecheck")
	VsanCapabilityTypeLocaldataprotection            = VsanCapabilityType("localdataprotection")
	VsanCapabilityTypeApidevversionenabled           = VsanCapabilityType("apidevversionenabled")
	VsanCapabilityTypeClusteradvancedoptions         = VsanCapabilityType("clusteradvancedoptions")
	VsanCapabilityTypeHostaffinity                   = VsanCapabilityType("hostaffinity")
	VsanCapabilityTypePmanintegration                = VsanCapabilityType("pmanintegration")
	VsanCapabilityTypeWitnessmanagement              = VsanCapabilityType("witnessmanagement")
	VsanCapabilityTypeNativelargeclustersupport      = VsanCapabilityType("nativelargeclustersupport")
	VsanCapabilityTypePerfsvctwoyaxisgraph           = VsanCapabilityType("perfsvctwoyaxisgraph")
	VsanCapabilityTypeCloudhealth                    = VsanCapabilityType("cloudhealth")
	VsanCapabilityTypeIdentitiessupportpolicyid      = VsanCapabilityType("identitiessupportpolicyid")
	VsanCapabilityTypeFileservices                   = VsanCapabilityType("fileservices")
	VsanCapabilityTypeVsanCapabilityType_Unknown     = VsanCapabilityType("VsanCapabilityType_Unknown")
	VsanCapabilityTypeVsanmetadatanode               = VsanCapabilityType("vsanmetadatanode")
	VsanCapabilityTypeDiagnosticsfeedback            = VsanCapabilityType("diagnosticsfeedback")
	VsanCapabilityTypeHistoricalhealth               = VsanCapabilityType("historicalhealth")
	VsanCapabilityTypeRemotedataprotection           = VsanCapabilityType("remotedataprotection")
	VsanCapabilityTypeStretchedcluster               = VsanCapabilityType("stretchedcluster")
	VsanCapabilityTypeArchivaldataprotection         = VsanCapabilityType("archivaldataprotection")
	VsanCapabilityTypeComplianceprecheck             = VsanCapabilityType("complianceprecheck")
	VsanCapabilityTypeFcd                            = VsanCapabilityType("fcd")
	VsanCapabilityTypeSupportApiVersion              = VsanCapabilityType("supportApiVersion")
	VsanCapabilityTypeRepairtimerinresyncstats       = VsanCapabilityType("repairtimerinresyncstats")
	VsanCapabilityTypePerfanalysis                   = VsanCapabilityType("perfanalysis")
)

func init() {
	types.Add("vsan:VsanCapabilityType", reflect.TypeOf((*VsanCapabilityType)(nil)).Elem())
}

type VsanClusterHealthActionVsanClusterHealthActionIdEnum string

const (
	VsanClusterHealthActionVsanClusterHealthActionIdEnumVsanClusterHealthActionIdEnum_Unknown = VsanClusterHealthActionVsanClusterHealthActionIdEnum("VsanClusterHealthActionIdEnum_Unknown")
	VsanClusterHealthActionVsanClusterHealthActionIdEnumConfigureVSAN                         = VsanClusterHealthActionVsanClusterHealthActionIdEnum("ConfigureVSAN")
	VsanClusterHealthActionVsanClusterHealthActionIdEnumUploadHclDb                           = VsanClusterHealthActionVsanClusterHealthActionIdEnum("UploadHclDb")
	VsanClusterHealthActionVsanClusterHealthActionIdEnumRemediateDedup                        = VsanClusterHealthActionVsanClusterHealthActionIdEnum("RemediateDedup")
	VsanClusterHealthActionVsanClusterHealthActionIdEnumEnablePerformanceServiceAction        = VsanClusterHealthActionVsanClusterHealthActionIdEnum("EnablePerformanceServiceAction")
	VsanClusterHealthActionVsanClusterHealthActionIdEnumEnableCeip                            = VsanClusterHealthActionVsanClusterHealthActionIdEnum("EnableCeip")
	VsanClusterHealthActionVsanClusterHealthActionIdEnumLoginVumIsoDepot                      = VsanClusterHealthActionVsanClusterHealthActionIdEnum("LoginVumIsoDepot")
	VsanClusterHealthActionVsanClusterHealthActionIdEnumRelayoutVsanObjects                   = VsanClusterHealthActionVsanClusterHealthActionIdEnum("RelayoutVsanObjects")
	VsanClusterHealthActionVsanClusterHealthActionIdEnumRemediateFileService                  = VsanClusterHealthActionVsanClusterHealthActionIdEnum("RemediateFileService")
	VsanClusterHealthActionVsanClusterHealthActionIdEnumConfigureHA                           = VsanClusterHealthActionVsanClusterHealthActionIdEnum("ConfigureHA")
	VsanClusterHealthActionVsanClusterHealthActionIdEnumConfigureAutomaticRebalance           = VsanClusterHealthActionVsanClusterHealthActionIdEnum("ConfigureAutomaticRebalance")
	VsanClusterHealthActionVsanClusterHealthActionIdEnumCreateDVS                             = VsanClusterHealthActionVsanClusterHealthActionIdEnum("CreateDVS")
	VsanClusterHealthActionVsanClusterHealthActionIdEnumRunBurnInTest                         = VsanClusterHealthActionVsanClusterHealthActionIdEnum("RunBurnInTest")
	VsanClusterHealthActionVsanClusterHealthActionIdEnumUploadReleaseCatalog                  = VsanClusterHealthActionVsanClusterHealthActionIdEnum("UploadReleaseCatalog")
	VsanClusterHealthActionVsanClusterHealthActionIdEnumUpgradeVsanDiskFormat                 = VsanClusterHealthActionVsanClusterHealthActionIdEnum("UpgradeVsanDiskFormat")
	VsanClusterHealthActionVsanClusterHealthActionIdEnumEnableHealthService                   = VsanClusterHealthActionVsanClusterHealthActionIdEnum("EnableHealthService")
	VsanClusterHealthActionVsanClusterHealthActionIdEnumPurgeInaccessSwapObjs                 = VsanClusterHealthActionVsanClusterHealthActionIdEnum("PurgeInaccessSwapObjs")
	VsanClusterHealthActionVsanClusterHealthActionIdEnumDiskBalance                           = VsanClusterHealthActionVsanClusterHealthActionIdEnum("DiskBalance")
	VsanClusterHealthActionVsanClusterHealthActionIdEnumEnableIscsiTargetService              = VsanClusterHealthActionVsanClusterHealthActionIdEnum("EnableIscsiTargetService")
	VsanClusterHealthActionVsanClusterHealthActionIdEnumRepairClusterObjectsAction            = VsanClusterHealthActionVsanClusterHealthActionIdEnum("RepairClusterObjectsAction")
	VsanClusterHealthActionVsanClusterHealthActionIdEnumClaimVSANDisks                        = VsanClusterHealthActionVsanClusterHealthActionIdEnum("ClaimVSANDisks")
	VsanClusterHealthActionVsanClusterHealthActionIdEnumStopDiskBalance                       = VsanClusterHealthActionVsanClusterHealthActionIdEnum("StopDiskBalance")
	VsanClusterHealthActionVsanClusterHealthActionIdEnumConfigureDRS                          = VsanClusterHealthActionVsanClusterHealthActionIdEnum("ConfigureDRS")
	VsanClusterHealthActionVsanClusterHealthActionIdEnumClusterUpgrade                        = VsanClusterHealthActionVsanClusterHealthActionIdEnum("ClusterUpgrade")
	VsanClusterHealthActionVsanClusterHealthActionIdEnumCreateVMKnic                          = VsanClusterHealthActionVsanClusterHealthActionIdEnum("CreateVMKnic")
	VsanClusterHealthActionVsanClusterHealthActionIdEnumUpdateHclDbFromInternet               = VsanClusterHealthActionVsanClusterHealthActionIdEnum("UpdateHclDbFromInternet")
	VsanClusterHealthActionVsanClusterHealthActionIdEnumRemediateClusterConfig                = VsanClusterHealthActionVsanClusterHealthActionIdEnum("RemediateClusterConfig")
	VsanClusterHealthActionVsanClusterHealthActionIdEnumCreateVMKnicWithVMotion               = VsanClusterHealthActionVsanClusterHealthActionIdEnum("CreateVMKnicWithVMotion")
)

func init() {
	types.Add("vsan:VsanClusterHealthActionVsanClusterHealthActionIdEnum", reflect.TypeOf((*VsanClusterHealthActionVsanClusterHealthActionIdEnum)(nil)).Elem())
}

type VsanCompositeConstraintConjoinerEnum string

const (
	VsanCompositeConstraintConjoinerEnumAND                                          = VsanCompositeConstraintConjoinerEnum("AND")
	VsanCompositeConstraintConjoinerEnumVsanCompositeConstraintConjoinerEnum_Unknown = VsanCompositeConstraintConjoinerEnum("VsanCompositeConstraintConjoinerEnum_Unknown")
	VsanCompositeConstraintConjoinerEnumOR                                           = VsanCompositeConstraintConjoinerEnum("OR")
	VsanCompositeConstraintConjoinerEnumEXCEPT                                       = VsanCompositeConstraintConjoinerEnum("EXCEPT")
)

func init() {
	types.Add("vsan:VsanCompositeConstraintConjoinerEnum", reflect.TypeOf((*VsanCompositeConstraintConjoinerEnum)(nil)).Elem())
}

type VsanDiskBalanceState string

const (
	VsanDiskBalanceStateReactiverebalancefailed      = VsanDiskBalanceState("reactiverebalancefailed")
	VsanDiskBalanceStateProactivenotmustdo           = VsanDiskBalanceState("proactivenotmustdo")
	VsanDiskBalanceStateRebalancediskunhealthy       = VsanDiskBalanceState("rebalancediskunhealthy")
	VsanDiskBalanceStateImbalancewithintolerance     = VsanDiskBalanceState("imbalancewithintolerance")
	VsanDiskBalanceStateProactiverebalancefailed     = VsanDiskBalanceState("proactiverebalancefailed")
	VsanDiskBalanceStateRebalanceentitydecom         = VsanDiskBalanceState("rebalanceentitydecom")
	VsanDiskBalanceStateProactiveneededbutdisabled   = VsanDiskBalanceState("proactiveneededbutdisabled")
	VsanDiskBalanceStateProactiverebalanceinprogress = VsanDiskBalanceState("proactiverebalanceinprogress")
	VsanDiskBalanceStateRebalanceoff                 = VsanDiskBalanceState("rebalanceoff")
	VsanDiskBalanceStateReactiverebalanceinprogress  = VsanDiskBalanceState("reactiverebalanceinprogress")
	VsanDiskBalanceStateVsanDiskBalanceState_Unknown = VsanDiskBalanceState("VsanDiskBalanceState_Unknown")
)

func init() {
	types.Add("vsan:VsanDiskBalanceState", reflect.TypeOf((*VsanDiskBalanceState)(nil)).Elem())
}

type VsanEncryptionIssue string

const (
	VsanEncryptionIssueKeyencryptionkeyinconsistent    = VsanEncryptionIssue("keyencryptionkeyinconsistent")
	VsanEncryptionIssueCmknotinenabledstate            = VsanEncryptionIssue("cmknotinenabledstate")
	VsanEncryptionIssueClientkeyinconsistent           = VsanEncryptionIssue("clientkeyinconsistent")
	VsanEncryptionIssueKeknotavailable                 = VsanEncryptionIssue("keknotavailable")
	VsanEncryptionIssueHostkeynotavailable             = VsanEncryptionIssue("hostkeynotavailable")
	VsanEncryptionIssueServercertificatesinconsistent  = VsanEncryptionIssue("servercertificatesinconsistent")
	VsanEncryptionIssueVsanEncryptionIssue_Unknown     = VsanEncryptionIssue("VsanEncryptionIssue_Unknown")
	VsanEncryptionIssueDataencryptionkeyinconsistent   = VsanEncryptionIssue("dataencryptionkeyinconsistent")
	VsanEncryptionIssueHostkeyinconsistent             = VsanEncryptionIssue("hostkeyinconsistent")
	VsanEncryptionIssueErasedisksbeforeuseinconsistent = VsanEncryptionIssue("erasedisksbeforeuseinconsistent")
	VsanEncryptionIssueClientcertificateinconsistent   = VsanEncryptionIssue("clientcertificateinconsistent")
	VsanEncryptionIssueCmkcannotretrieve               = VsanEncryptionIssue("cmkcannotretrieve")
	VsanEncryptionIssueKmsinfoinconsistent             = VsanEncryptionIssue("kmsinfoinconsistent")
	VsanEncryptionIssueEnabledwhenclusterdisabled      = VsanEncryptionIssue("enabledwhenclusterdisabled")
	VsanEncryptionIssueDisabledwhenclusterenabled      = VsanEncryptionIssue("disabledwhenclusterenabled")
)

func init() {
	types.Add("vsan:VsanEncryptionIssue", reflect.TypeOf((*VsanEncryptionIssue)(nil)).Elem())
}

type VsanFileServiceVMStatus string

const (
	VsanFileServiceVMStatusRunning                     = VsanFileServiceVMStatus("running")
	VsanFileServiceVMStatusUpgrading                   = VsanFileServiceVMStatus("upgrading")
	VsanFileServiceVMStatusFileServiceVMStatus_Unknown = VsanFileServiceVMStatus("FileServiceVMStatus_Unknown")
)

func init() {
	types.Add("vsan:VsanFileServiceVMStatus", reflect.TypeOf((*VsanFileServiceVMStatus)(nil)).Elem())
}

type VsanFileShareAccessType string

const (
	VsanFileShareAccessTypeREAD_ONLY                   = VsanFileShareAccessType("READ_ONLY")
	VsanFileShareAccessTypeFileShareAccessType_Unknown = VsanFileShareAccessType("FileShareAccessType_Unknown")
	VsanFileShareAccessTypeREAD_WRITE                  = VsanFileShareAccessType("READ_WRITE")
	VsanFileShareAccessTypeNO_ACCESS                   = VsanFileShareAccessType("NO_ACCESS")
)

func init() {
	types.Add("vsan:VsanFileShareAccessType", reflect.TypeOf((*VsanFileShareAccessType)(nil)).Elem())
}

type VsanFileShareManagingEntity string

const (
	VsanFileShareManagingEntityCns                             = VsanFileShareManagingEntity("cns")
	VsanFileShareManagingEntityFileShareManagingEntity_Unknown = VsanFileShareManagingEntity("FileShareManagingEntity_Unknown")
	VsanFileShareManagingEntityUser                            = VsanFileShareManagingEntity("user")
)

func init() {
	types.Add("vsan:VsanFileShareManagingEntity", reflect.TypeOf((*VsanFileShareManagingEntity)(nil)).Elem())
}

type VsanHealthPerspective string

const (
	VsanHealthPerspectiveUpgradeBeforeExitMM           = VsanHealthPerspective("upgradeBeforeExitMM")
	VsanHealthPerspectiveUpgradePreCheck               = VsanHealthPerspective("upgradePreCheck")
	VsanHealthPerspectiveUpgradePreCheckPman           = VsanHealthPerspective("upgradePreCheckPman")
	VsanHealthPerspectiveUpgradeAfterExitMM            = VsanHealthPerspective("upgradeAfterExitMM")
	VsanHealthPerspectiveUpgradeBeforeExitMMPman       = VsanHealthPerspective("upgradeBeforeExitMMPman")
	VsanHealthPerspectiveBeforeConfigureHost           = VsanHealthPerspective("beforeConfigureHost")
	VsanHealthPerspectiveDefaultView                   = VsanHealthPerspective("defaultView")
	VsanHealthPerspectiveVsanUpgradeAfterExitMM        = VsanHealthPerspective("vsanUpgradeAfterExitMM")
	VsanHealthPerspectiveDeployAssist                  = VsanHealthPerspective("deployAssist")
	VsanHealthPerspectiveVsanUpgradePreCheck           = VsanHealthPerspective("vsanUpgradePreCheck")
	VsanHealthPerspectiveVsanHealthPerspective_Unknown = VsanHealthPerspective("VsanHealthPerspective_Unknown")
	VsanHealthPerspectiveUpgradeAfterExitMMPman        = VsanHealthPerspective("upgradeAfterExitMMPman")
	VsanHealthPerspectiveCreateExtendClusterView       = VsanHealthPerspective("CreateExtendClusterView")
	VsanHealthPerspectiveVsanUpgradeBeforeExitMM       = VsanHealthPerspective("vsanUpgradeBeforeExitMM")
	VsanHealthPerspectiveVmcUpgradePreChecks           = VsanHealthPerspective("vmcUpgradePreChecks")
)

func init() {
	types.Add("vsan:VsanHealthPerspective", reflect.TypeOf((*VsanHealthPerspective)(nil)).Elem())
}

type VsanHealthStatusType string

const (
	VsanHealthStatusTypeUnknown = VsanHealthStatusType("unknown")
	VsanHealthStatusTypeGreen   = VsanHealthStatusType("green")
	VsanHealthStatusTypeRed     = VsanHealthStatusType("red")
	VsanHealthStatusTypeYellow  = VsanHealthStatusType("yellow")
)

func init() {
	types.Add("vsan:VsanHealthStatusType", reflect.TypeOf((*VsanHealthStatusType)(nil)).Elem())
}

type VsanHostPortConfigExTrafficType string

const (
	VsanHostPortConfigExTrafficTypeTrafficType_Unknown = VsanHostPortConfigExTrafficType("TrafficType_Unknown")
	VsanHostPortConfigExTrafficTypeVsan                = VsanHostPortConfigExTrafficType("vsan")
	VsanHostPortConfigExTrafficTypeWitness             = VsanHostPortConfigExTrafficType("witness")
)

func init() {
	types.Add("vsan:VsanHostPortConfigExTrafficType", reflect.TypeOf((*VsanHostPortConfigExTrafficType)(nil)).Elem())
}

type VsanHostQueryCheckLimitsOptionType string

const (
	VsanHostQueryCheckLimitsOptionTypeLogicalCapacityUsed                        = VsanHostQueryCheckLimitsOptionType("logicalCapacityUsed")
	VsanHostQueryCheckLimitsOptionTypeDedupMetadata                              = VsanHostQueryCheckLimitsOptionType("dedupMetadata")
	VsanHostQueryCheckLimitsOptionTypeVsanHostQueryCheckLimitsOptionType_Unknown = VsanHostQueryCheckLimitsOptionType("VsanHostQueryCheckLimitsOptionType_Unknown")
	VsanHostQueryCheckLimitsOptionTypeLogicalCapacity                            = VsanHostQueryCheckLimitsOptionType("logicalCapacity")
	VsanHostQueryCheckLimitsOptionTypeDgTransientCapacityUsed                    = VsanHostQueryCheckLimitsOptionType("dgTransientCapacityUsed")
	VsanHostQueryCheckLimitsOptionTypeDiskTransientCapacityUsed                  = VsanHostQueryCheckLimitsOptionType("diskTransientCapacityUsed")
)

func init() {
	types.Add("vsan:VsanHostQueryCheckLimitsOptionType", reflect.TypeOf((*VsanHostQueryCheckLimitsOptionType)(nil)).Elem())
}

type VsanHostStatsType string

const (
	VsanHostStatsTypeConfigGeneration     = VsanHostStatsType("configGeneration")
	VsanHostStatsTypeRepairTimerInfo      = VsanHostStatsType("repairTimerInfo")
	VsanHostStatsTypeResyncIopsInfo       = VsanHostStatsType("resyncIopsInfo")
	VsanHostStatsTypeStatsType_Unknown    = VsanHostStatsType("StatsType_Unknown")
	VsanHostStatsTypeSupportedClusterSize = VsanHostStatsType("supportedClusterSize")
)

func init() {
	types.Add("vsan:VsanHostStatsType", reflect.TypeOf((*VsanHostStatsType)(nil)).Elem())
}

type VsanIscsiLUNCommonInfoVsanIscsiLUNStatus string

const (
	VsanIscsiLUNCommonInfoVsanIscsiLUNStatusOffline                    = VsanIscsiLUNCommonInfoVsanIscsiLUNStatus("Offline")
	VsanIscsiLUNCommonInfoVsanIscsiLUNStatusVsanIscsiLUNStatus_Unknown = VsanIscsiLUNCommonInfoVsanIscsiLUNStatus("VsanIscsiLUNStatus_Unknown")
	VsanIscsiLUNCommonInfoVsanIscsiLUNStatusOnline                     = VsanIscsiLUNCommonInfoVsanIscsiLUNStatus("Online")
)

func init() {
	types.Add("vsan:VsanIscsiLUNCommonInfoVsanIscsiLUNStatus", reflect.TypeOf((*VsanIscsiLUNCommonInfoVsanIscsiLUNStatus)(nil)).Elem())
}

type VsanIscsiTargetAuthSpecVsanIscsiTargetAuthType string

const (
	VsanIscsiTargetAuthSpecVsanIscsiTargetAuthTypeCHAP                            = VsanIscsiTargetAuthSpecVsanIscsiTargetAuthType("CHAP")
	VsanIscsiTargetAuthSpecVsanIscsiTargetAuthTypeNoAuth                          = VsanIscsiTargetAuthSpecVsanIscsiTargetAuthType("NoAuth")
	VsanIscsiTargetAuthSpecVsanIscsiTargetAuthTypeCHAP_Mutual                     = VsanIscsiTargetAuthSpecVsanIscsiTargetAuthType("CHAP_Mutual")
	VsanIscsiTargetAuthSpecVsanIscsiTargetAuthTypeVsanIscsiTargetAuthType_Unknown = VsanIscsiTargetAuthSpecVsanIscsiTargetAuthType("VsanIscsiTargetAuthType_Unknown")
)

func init() {
	types.Add("vsan:VsanIscsiTargetAuthSpecVsanIscsiTargetAuthType", reflect.TypeOf((*VsanIscsiTargetAuthSpecVsanIscsiTargetAuthType)(nil)).Elem())
}

type VsanIscsiTargetServiceProcessStatus string

const (
	VsanIscsiTargetServiceProcessStatusRunning                                     = VsanIscsiTargetServiceProcessStatus("Running")
	VsanIscsiTargetServiceProcessStatusStopped                                     = VsanIscsiTargetServiceProcessStatus("Stopped")
	VsanIscsiTargetServiceProcessStatusVsanIscsiTargetServiceProcessStatus_Unknown = VsanIscsiTargetServiceProcessStatus("VsanIscsiTargetServiceProcessStatus_Unknown")
)

func init() {
	types.Add("vsan:VsanIscsiTargetServiceProcessStatus", reflect.TypeOf((*VsanIscsiTargetServiceProcessStatus)(nil)).Elem())
}

type VsanMassCollectorObjectCollectionEnum string

const (
	VsanMassCollectorObjectCollectionEnumVsanMassCollectorObjectCollectionEnum_Unknown = VsanMassCollectorObjectCollectionEnum("VsanMassCollectorObjectCollectionEnum_Unknown")
	VsanMassCollectorObjectCollectionEnumALL_HOSTS                                     = VsanMassCollectorObjectCollectionEnum("ALL_HOSTS")
	VsanMassCollectorObjectCollectionEnumALL_CLUSTERS                                  = VsanMassCollectorObjectCollectionEnum("ALL_CLUSTERS")
	VsanMassCollectorObjectCollectionEnumALL_VSAN_DATASTORES                           = VsanMassCollectorObjectCollectionEnum("ALL_VSAN_DATASTORES")
	VsanMassCollectorObjectCollectionEnumVCENTER                                       = VsanMassCollectorObjectCollectionEnum("VCENTER")
	VsanMassCollectorObjectCollectionEnumALL_DATASTORES                                = VsanMassCollectorObjectCollectionEnum("ALL_DATASTORES")
	VsanMassCollectorObjectCollectionEnumALL_VSAN_ENABLED_HOSTS                        = VsanMassCollectorObjectCollectionEnum("ALL_VSAN_ENABLED_HOSTS")
	VsanMassCollectorObjectCollectionEnumSERVICE_INSTANCE                              = VsanMassCollectorObjectCollectionEnum("SERVICE_INSTANCE")
	VsanMassCollectorObjectCollectionEnumALL_VMFS_DATASTORES                           = VsanMassCollectorObjectCollectionEnum("ALL_VMFS_DATASTORES")
	VsanMassCollectorObjectCollectionEnumALL_VSAN_ENABLED_HOSTS_EXCEPT_WITNESS         = VsanMassCollectorObjectCollectionEnum("ALL_VSAN_ENABLED_HOSTS_EXCEPT_WITNESS")
	VsanMassCollectorObjectCollectionEnumALL_VSAN_ENABLED_CLUSTERS                     = VsanMassCollectorObjectCollectionEnum("ALL_VSAN_ENABLED_CLUSTERS")
)

func init() {
	types.Add("vsan:VsanMassCollectorObjectCollectionEnum", reflect.TypeOf((*VsanMassCollectorObjectCollectionEnum)(nil)).Elem())
}

type VsanObjectHealthVsanObjectHealthState string

const (
	VsanObjectHealthVsanObjectHealthStateVsanObjectHealthState_Unknown                             = VsanObjectHealthVsanObjectHealthState("VsanObjectHealthState_Unknown")
	VsanObjectHealthVsanObjectHealthStateInaccessible                                              = VsanObjectHealthVsanObjectHealthState("inaccessible")
	VsanObjectHealthVsanObjectHealthStateDatamove                                                  = VsanObjectHealthVsanObjectHealthState("datamove")
	VsanObjectHealthVsanObjectHealthStateNonavailabilityrelatedincompliancewithpolicypending       = VsanObjectHealthVsanObjectHealthState("nonavailabilityrelatedincompliancewithpolicypending")
	VsanObjectHealthVsanObjectHealthStateNonavailabilityrelatedincompliancewithpolicypendingfailed = VsanObjectHealthVsanObjectHealthState("nonavailabilityrelatedincompliancewithpolicypendingfailed")
	VsanObjectHealthVsanObjectHealthStateNonavailabilityrelatedincompliance                        = VsanObjectHealthVsanObjectHealthState("nonavailabilityrelatedincompliance")
	VsanObjectHealthVsanObjectHealthStateReducedavailabilitywithpolicypending                      = VsanObjectHealthVsanObjectHealthState("reducedavailabilitywithpolicypending")
	VsanObjectHealthVsanObjectHealthStateReducedavailabilitywithpolicypendingfailed                = VsanObjectHealthVsanObjectHealthState("reducedavailabilitywithpolicypendingfailed")
	VsanObjectHealthVsanObjectHealthStateReducedavailabilitywithnorebuilddelaytimer                = VsanObjectHealthVsanObjectHealthState("reducedavailabilitywithnorebuilddelaytimer")
	VsanObjectHealthVsanObjectHealthStateReducedavailabilitywithpausedrebuild                      = VsanObjectHealthVsanObjectHealthState("reducedavailabilitywithpausedrebuild")
	VsanObjectHealthVsanObjectHealthStateNonavailabilityrelatedincompliancewithpausedrebuild       = VsanObjectHealthVsanObjectHealthState("nonavailabilityrelatedincompliancewithpausedrebuild")
	VsanObjectHealthVsanObjectHealthStateHealthy                                                   = VsanObjectHealthVsanObjectHealthState("healthy")
	VsanObjectHealthVsanObjectHealthStateReducedavailabilitywithactiverebuild                      = VsanObjectHealthVsanObjectHealthState("reducedavailabilitywithactiverebuild")
	VsanObjectHealthVsanObjectHealthStateNonavailabilityrelatedreconfig                            = VsanObjectHealthVsanObjectHealthState("nonavailabilityrelatedreconfig")
	VsanObjectHealthVsanObjectHealthStateReducedavailabilitywithnorebuild                          = VsanObjectHealthVsanObjectHealthState("reducedavailabilitywithnorebuild")
)

func init() {
	types.Add("vsan:VsanObjectHealthVsanObjectHealthState", reflect.TypeOf((*VsanObjectHealthVsanObjectHealthState)(nil)).Elem())
}

type VsanObjectSpaceSummaryVsanObjectTypeEnum string

const (
	VsanObjectSpaceSummaryVsanObjectTypeEnumVmswap                       = VsanObjectSpaceSummaryVsanObjectTypeEnum("vmswap")
	VsanObjectSpaceSummaryVsanObjectTypeEnumSlackSpaceCapRequiredForHost = VsanObjectSpaceSummaryVsanObjectTypeEnum("slackSpaceCapRequiredForHost")
	VsanObjectSpaceSummaryVsanObjectTypeEnumDedupOverhead                = VsanObjectSpaceSummaryVsanObjectTypeEnum("dedupOverhead")
	VsanObjectSpaceSummaryVsanObjectTypeEnumResynPauseThresholdForHost   = VsanObjectSpaceSummaryVsanObjectTypeEnum("resynPauseThresholdForHost")
	VsanObjectSpaceSummaryVsanObjectTypeEnumDetachedCnsVolBlock          = VsanObjectSpaceSummaryVsanObjectTypeEnum("detachedCnsVolBlock")
	VsanObjectSpaceSummaryVsanObjectTypeEnumHbrPersist                   = VsanObjectSpaceSummaryVsanObjectTypeEnum("hbrPersist")
	VsanObjectSpaceSummaryVsanObjectTypeEnumAttachedCnsVolFile           = VsanObjectSpaceSummaryVsanObjectTypeEnum("attachedCnsVolFile")
	VsanObjectSpaceSummaryVsanObjectTypeEnumFileShare                    = VsanObjectSpaceSummaryVsanObjectTypeEnum("fileShare")
	VsanObjectSpaceSummaryVsanObjectTypeEnumVdisk                        = VsanObjectSpaceSummaryVsanObjectTypeEnum("vdisk")
	VsanObjectSpaceSummaryVsanObjectTypeEnumVsanObjectTypeEnum_Unknown   = VsanObjectSpaceSummaryVsanObjectTypeEnum("VsanObjectTypeEnum_Unknown")
	VsanObjectSpaceSummaryVsanObjectTypeEnumNamespace                    = VsanObjectSpaceSummaryVsanObjectTypeEnum("namespace")
	VsanObjectSpaceSummaryVsanObjectTypeEnumImprovedVirtualDisk          = VsanObjectSpaceSummaryVsanObjectTypeEnum("improvedVirtualDisk")
	VsanObjectSpaceSummaryVsanObjectTypeEnumOther                        = VsanObjectSpaceSummaryVsanObjectTypeEnum("other")
	VsanObjectSpaceSummaryVsanObjectTypeEnumPhysicalTransientSpace       = VsanObjectSpaceSummaryVsanObjectTypeEnum("physicalTransientSpace")
	VsanObjectSpaceSummaryVsanObjectTypeEnumDetachedCnsVolFile           = VsanObjectSpaceSummaryVsanObjectTypeEnum("detachedCnsVolFile")
	VsanObjectSpaceSummaryVsanObjectTypeEnumFileServiceRoot              = VsanObjectSpaceSummaryVsanObjectTypeEnum("fileServiceRoot")
	VsanObjectSpaceSummaryVsanObjectTypeEnumIscsiLun                     = VsanObjectSpaceSummaryVsanObjectTypeEnum("iscsiLun")
	VsanObjectSpaceSummaryVsanObjectTypeEnumChecksumOverhead             = VsanObjectSpaceSummaryVsanObjectTypeEnum("checksumOverhead")
	VsanObjectSpaceSummaryVsanObjectTypeEnumFileSystemOverhead           = VsanObjectSpaceSummaryVsanObjectTypeEnum("fileSystemOverhead")
	VsanObjectSpaceSummaryVsanObjectTypeEnumAttachedCnsVolBlock          = VsanObjectSpaceSummaryVsanObjectTypeEnum("attachedCnsVolBlock")
	VsanObjectSpaceSummaryVsanObjectTypeEnumSpaceUnderDedupConsideration = VsanObjectSpaceSummaryVsanObjectTypeEnum("spaceUnderDedupConsideration")
	VsanObjectSpaceSummaryVsanObjectTypeEnumMinSpaceRequiredForVsanOp    = VsanObjectSpaceSummaryVsanObjectTypeEnum("minSpaceRequiredForVsanOp")
	VsanObjectSpaceSummaryVsanObjectTypeEnumHostRebuildCapacity          = VsanObjectSpaceSummaryVsanObjectTypeEnum("hostRebuildCapacity")
	VsanObjectSpaceSummaryVsanObjectTypeEnumCnsVolFile                   = VsanObjectSpaceSummaryVsanObjectTypeEnum("cnsVolFile")
	VsanObjectSpaceSummaryVsanObjectTypeEnumHbrDisk                      = VsanObjectSpaceSummaryVsanObjectTypeEnum("hbrDisk")
	VsanObjectSpaceSummaryVsanObjectTypeEnumExtension                    = VsanObjectSpaceSummaryVsanObjectTypeEnum("extension")
	VsanObjectSpaceSummaryVsanObjectTypeEnumStatsdb                      = VsanObjectSpaceSummaryVsanObjectTypeEnum("statsdb")
	VsanObjectSpaceSummaryVsanObjectTypeEnumVmem                         = VsanObjectSpaceSummaryVsanObjectTypeEnum("vmem")
	VsanObjectSpaceSummaryVsanObjectTypeEnumTransientSpace               = VsanObjectSpaceSummaryVsanObjectTypeEnum("transientSpace")
	VsanObjectSpaceSummaryVsanObjectTypeEnumHbrCfg                       = VsanObjectSpaceSummaryVsanObjectTypeEnum("hbrCfg")
	VsanObjectSpaceSummaryVsanObjectTypeEnumIscsiTarget                  = VsanObjectSpaceSummaryVsanObjectTypeEnum("iscsiTarget")
)

func init() {
	types.Add("vsan:VsanObjectSpaceSummaryVsanObjectTypeEnum", reflect.TypeOf((*VsanObjectSpaceSummaryVsanObjectTypeEnum)(nil)).Elem())
}

type VsanPerfDiagnosticQueryType string

const (
	VsanPerfDiagnosticQueryTypeIops                                = VsanPerfDiagnosticQueryType("iops")
	VsanPerfDiagnosticQueryTypeLat                                 = VsanPerfDiagnosticQueryType("lat")
	VsanPerfDiagnosticQueryTypeTput                                = VsanPerfDiagnosticQueryType("tput")
	VsanPerfDiagnosticQueryTypeVsanPerfDiagnosticQueryType_Unknown = VsanPerfDiagnosticQueryType("VsanPerfDiagnosticQueryType_Unknown")
	VsanPerfDiagnosticQueryTypeEval                                = VsanPerfDiagnosticQueryType("eval")
)

func init() {
	types.Add("vsan:VsanPerfDiagnosticQueryType", reflect.TypeOf((*VsanPerfDiagnosticQueryType)(nil)).Elem())
}

type VsanPerfGraphVsanPerfStatsUnitType string

const (
	VsanPerfGraphVsanPerfStatsUnitTypeSize_bytes                    = VsanPerfGraphVsanPerfStatsUnitType("size_bytes")
	VsanPerfGraphVsanPerfStatsUnitTypePermille                      = VsanPerfGraphVsanPerfStatsUnitType("permille")
	VsanPerfGraphVsanPerfStatsUnitTypeTime_ms                       = VsanPerfGraphVsanPerfStatsUnitType("time_ms")
	VsanPerfGraphVsanPerfStatsUnitTypePercentage                    = VsanPerfGraphVsanPerfStatsUnitType("percentage")
	VsanPerfGraphVsanPerfStatsUnitTypeTime_s                        = VsanPerfGraphVsanPerfStatsUnitType("time_s")
	VsanPerfGraphVsanPerfStatsUnitTypeRate_bytes                    = VsanPerfGraphVsanPerfStatsUnitType("rate_bytes")
	VsanPerfGraphVsanPerfStatsUnitTypeNumber                        = VsanPerfGraphVsanPerfStatsUnitType("number")
	VsanPerfGraphVsanPerfStatsUnitTypeVsanPerfStatsUnitType_Unknown = VsanPerfGraphVsanPerfStatsUnitType("VsanPerfStatsUnitType_Unknown")
)

func init() {
	types.Add("vsan:VsanPerfGraphVsanPerfStatsUnitType", reflect.TypeOf((*VsanPerfGraphVsanPerfStatsUnitType)(nil)).Elem())
}

type VsanPerfMetricIdVsanPerfStatsType string

const (
	VsanPerfMetricIdVsanPerfStatsTypeVsanPerfStatsType_Unknown = VsanPerfMetricIdVsanPerfStatsType("VsanPerfStatsType_Unknown")
	VsanPerfMetricIdVsanPerfStatsTypeRate                      = VsanPerfMetricIdVsanPerfStatsType("rate")
	VsanPerfMetricIdVsanPerfStatsTypeDelta                     = VsanPerfMetricIdVsanPerfStatsType("delta")
	VsanPerfMetricIdVsanPerfStatsTypeAbsolute                  = VsanPerfMetricIdVsanPerfStatsType("absolute")
)

func init() {
	types.Add("vsan:VsanPerfMetricIdVsanPerfStatsType", reflect.TypeOf((*VsanPerfMetricIdVsanPerfStatsType)(nil)).Elem())
}

type VsanPerfMetricIdVsanPerfSummaryType string

const (
	VsanPerfMetricIdVsanPerfSummaryTypeNone                        = VsanPerfMetricIdVsanPerfSummaryType("none")
	VsanPerfMetricIdVsanPerfSummaryTypeAverage                     = VsanPerfMetricIdVsanPerfSummaryType("average")
	VsanPerfMetricIdVsanPerfSummaryTypeMaximum                     = VsanPerfMetricIdVsanPerfSummaryType("maximum")
	VsanPerfMetricIdVsanPerfSummaryTypeVsanPerfSummaryType_Unknown = VsanPerfMetricIdVsanPerfSummaryType("VsanPerfSummaryType_Unknown")
	VsanPerfMetricIdVsanPerfSummaryTypeMinimum                     = VsanPerfMetricIdVsanPerfSummaryType("minimum")
	VsanPerfMetricIdVsanPerfSummaryTypeSummation                   = VsanPerfMetricIdVsanPerfSummaryType("summation")
	VsanPerfMetricIdVsanPerfSummaryTypeLatest                      = VsanPerfMetricIdVsanPerfSummaryType("latest")
)

func init() {
	types.Add("vsan:VsanPerfMetricIdVsanPerfSummaryType", reflect.TypeOf((*VsanPerfMetricIdVsanPerfSummaryType)(nil)).Elem())
}

type VsanPerfThresholdVsanPerfThresholdDirectionType string

const (
	VsanPerfThresholdVsanPerfThresholdDirectionTypeUpper                                  = VsanPerfThresholdVsanPerfThresholdDirectionType("upper")
	VsanPerfThresholdVsanPerfThresholdDirectionTypeLower                                  = VsanPerfThresholdVsanPerfThresholdDirectionType("lower")
	VsanPerfThresholdVsanPerfThresholdDirectionTypeVsanPerfThresholdDirectionType_Unknown = VsanPerfThresholdVsanPerfThresholdDirectionType("VsanPerfThresholdDirectionType_Unknown")
)

func init() {
	types.Add("vsan:VsanPerfThresholdVsanPerfThresholdDirectionType", reflect.TypeOf((*VsanPerfThresholdVsanPerfThresholdDirectionType)(nil)).Elem())
}

type VsanPerfsvcRemediateAction string

const (
	VsanPerfsvcRemediateActionUpdate_profile                 = VsanPerfsvcRemediateAction("update_profile")
	VsanPerfsvcRemediateActionPerfsvcRemediateAction_Unknown = VsanPerfsvcRemediateAction("PerfsvcRemediateAction_Unknown")
	VsanPerfsvcRemediateActionEnable                         = VsanPerfsvcRemediateAction("enable")
	VsanPerfsvcRemediateActionDisable                        = VsanPerfsvcRemediateAction("disable")
	VsanPerfsvcRemediateActionNo_action                      = VsanPerfsvcRemediateAction("no_action")
)

func init() {
	types.Add("vsan:VsanPerfsvcRemediateAction", reflect.TypeOf((*VsanPerfsvcRemediateAction)(nil)).Elem())
}

type VsanPropertyConstraintComparatorEnum string

const (
	VsanPropertyConstraintComparatorEnumSMALLER                                      = VsanPropertyConstraintComparatorEnum("SMALLER")
	VsanPropertyConstraintComparatorEnumGREATER                                      = VsanPropertyConstraintComparatorEnum("GREATER")
	VsanPropertyConstraintComparatorEnumCONTAINS                                     = VsanPropertyConstraintComparatorEnum("CONTAINS")
	VsanPropertyConstraintComparatorEnumEQUALS                                       = VsanPropertyConstraintComparatorEnum("EQUALS")
	VsanPropertyConstraintComparatorEnumPOP                                          = VsanPropertyConstraintComparatorEnum("POP")
	VsanPropertyConstraintComparatorEnumVsanPropertyConstraintComparatorEnum_Unknown = VsanPropertyConstraintComparatorEnum("VsanPropertyConstraintComparatorEnum_Unknown")
	VsanPropertyConstraintComparatorEnumTEXTUALLY_MATCHES                            = VsanPropertyConstraintComparatorEnum("TEXTUALLY_MATCHES")
)

func init() {
	types.Add("vsan:VsanPropertyConstraintComparatorEnum", reflect.TypeOf((*VsanPropertyConstraintComparatorEnum)(nil)).Elem())
}

type VsanResourceCheckStatusType string

const (
	VsanResourceCheckStatusTypeResourceCheckCompleted          = VsanResourceCheckStatusType("resourceCheckCompleted")
	VsanResourceCheckStatusTypeResourceCheckNotSupported       = VsanResourceCheckStatusType("resourceCheckNotSupported")
	VsanResourceCheckStatusTypeResourceCheckCancelled          = VsanResourceCheckStatusType("resourceCheckCancelled")
	VsanResourceCheckStatusTypeResourceCheckStatusType_Unknown = VsanResourceCheckStatusType("ResourceCheckStatusType_Unknown")
	VsanResourceCheckStatusTypeResourceCheckFailed             = VsanResourceCheckStatusType("resourceCheckFailed")
	VsanResourceCheckStatusTypeResourceCheckNoRecentValue      = VsanResourceCheckStatusType("resourceCheckNoRecentValue")
	VsanResourceCheckStatusTypeResourceCheckUninitialized      = VsanResourceCheckStatusType("resourceCheckUninitialized")
	VsanResourceCheckStatusTypeResourceCheckRunning            = VsanResourceCheckStatusType("resourceCheckRunning")
)

func init() {
	types.Add("vsan:VsanResourceCheckStatusType", reflect.TypeOf((*VsanResourceCheckStatusType)(nil)).Elem())
}

type VsanServiceStatus string

const (
	VsanServiceStatusStarted                   = VsanServiceStatus("started")
	VsanServiceStatusStopped                   = VsanServiceStatus("stopped")
	VsanServiceStatusVsanServiceStatus_Unknown = VsanServiceStatus("VsanServiceStatus_Unknown")
)

func init() {
	types.Add("vsan:VsanServiceStatus", reflect.TypeOf((*VsanServiceStatus)(nil)).Elem())
}

type VsanSmartParameterType string

const (
	VsanSmartParameterTypeSmartdrivetemperature          = VsanSmartParameterType("smartdrivetemperature")
	VsanSmartParameterTypeVsanSmartParameterType_Unknown = VsanSmartParameterType("VsanSmartParameterType_Unknown")
	VsanSmartParameterTypeSmartinitialbadblockcount      = VsanSmartParameterType("smartinitialbadblockcount")
	VsanSmartParameterTypeSmartdriveratedmaxtemperature  = VsanSmartParameterType("smartdriveratedmaxtemperature")
	VsanSmartParameterTypeSmartmediawearoutindicator     = VsanSmartParameterType("smartmediawearoutindicator")
	VsanSmartParameterTypeSmartwritesectorstotct         = VsanSmartParameterType("smartwritesectorstotct")
	VsanSmartParameterTypeSmartreallocatedsectorct       = VsanSmartParameterType("smartreallocatedsectorct")
	VsanSmartParameterTypeSmartreadsectorstotct          = VsanSmartParameterType("smartreadsectorstotct")
	VsanSmartParameterTypeSmartpowercyclecount           = VsanSmartParameterType("smartpowercyclecount")
	VsanSmartParameterTypeSmarthealthstatus              = VsanSmartParameterType("smarthealthstatus")
	VsanSmartParameterTypeSmartpoweronhours              = VsanSmartParameterType("smartpoweronhours")
	VsanSmartParameterTypeSmartwriteerrorcount           = VsanSmartParameterType("smartwriteerrorcount")
	VsanSmartParameterTypeSmartrawreaderrorrate          = VsanSmartParameterType("smartrawreaderrorrate")
	VsanSmartParameterTypeSmartreaderrorcount            = VsanSmartParameterType("smartreaderrorcount")
)

func init() {
	types.Add("vsan:VsanSmartParameterType", reflect.TypeOf((*VsanSmartParameterType)(nil)).Elem())
}

type VsanSpaceReportingEntityType string

const (
	VsanSpaceReportingEntityTypeHost                                 = VsanSpaceReportingEntityType("Host")
	VsanSpaceReportingEntityTypeFaultDomain                          = VsanSpaceReportingEntityType("FaultDomain")
	VsanSpaceReportingEntityTypeVsanSpaceReportingEntityType_Unknown = VsanSpaceReportingEntityType("VsanSpaceReportingEntityType_Unknown")
	VsanSpaceReportingEntityTypeVM                                   = VsanSpaceReportingEntityType("VM")
	VsanSpaceReportingEntityTypeFileShare                            = VsanSpaceReportingEntityType("FileShare")
)

func init() {
	types.Add("vsan:VsanSpaceReportingEntityType", reflect.TypeOf((*VsanSpaceReportingEntityType)(nil)).Elem())
}

type VsanStorageComplianceResultStorageComplianceStatus string

const (
	VsanStorageComplianceResultStorageComplianceStatusUnknown       = VsanStorageComplianceResultStorageComplianceStatus("unknown")
	VsanStorageComplianceResultStorageComplianceStatusCompliant     = VsanStorageComplianceResultStorageComplianceStatus("compliant")
	VsanStorageComplianceResultStorageComplianceStatusNonCompliant  = VsanStorageComplianceResultStorageComplianceStatus("nonCompliant")
	VsanStorageComplianceResultStorageComplianceStatusNotApplicable = VsanStorageComplianceResultStorageComplianceStatus("notApplicable")
)

func init() {
	types.Add("vsan:VsanStorageComplianceResultStorageComplianceStatus", reflect.TypeOf((*VsanStorageComplianceResultStorageComplianceStatus)(nil)).Elem())
}

type VsanSyncReason string

const (
	VsanSyncReasonObject_format_change   = VsanSyncReason("object_format_change")
	VsanSyncReasonRepair                 = VsanSyncReason("repair")
	VsanSyncReasonDying_evacuate         = VsanSyncReason("dying_evacuate")
	VsanSyncReasonReconfigure            = VsanSyncReason("reconfigure")
	VsanSyncReasonVsanSyncReason_Unknown = VsanSyncReason("VsanSyncReason_Unknown")
	VsanSyncReasonStale                  = VsanSyncReason("stale")
	VsanSyncReasonRebalance              = VsanSyncReason("rebalance")
	VsanSyncReasonEvacuate               = VsanSyncReason("evacuate")
	VsanSyncReasonMerge_concat           = VsanSyncReason("merge_concat")
)

func init() {
	types.Add("vsan:VsanSyncReason", reflect.TypeOf((*VsanSyncReason)(nil)).Elem())
}

type VsanSyncStatus string

const (
	VsanSyncStatusActive                 = VsanSyncStatus("active")
	VsanSyncStatusVsanSyncStatus_Unknown = VsanSyncStatus("VsanSyncStatus_Unknown")
	VsanSyncStatusQueued                 = VsanSyncStatus("queued")
	VsanSyncStatusSuspended              = VsanSyncStatus("suspended")
)

func init() {
	types.Add("vsan:VsanSyncStatus", reflect.TypeOf((*VsanSyncStatus)(nil)).Elem())
}

type VsanUpdateItemImpactType string

const (
	VsanUpdateItemImpactTypeVsanUpdateItemImpactType_Unknown = VsanUpdateItemImpactType("VsanUpdateItemImpactType_Unknown")
	VsanUpdateItemImpactTypeReboot                           = VsanUpdateItemImpactType("reboot")
)

func init() {
	types.Add("vsan:VsanUpdateItemImpactType", reflect.TypeOf((*VsanUpdateItemImpactType)(nil)).Elem())
}

type VsanUpdateItemType string

const (
	VsanUpdateItemTypeVib                        = VsanUpdateItemType("vib")
	VsanUpdateItemTypeOfflinebundle              = VsanUpdateItemType("offlinebundle")
	VsanUpdateItemTypeFullStackFirmware          = VsanUpdateItemType("fullStackFirmware")
	VsanUpdateItemTypeVmhbaFirmware              = VsanUpdateItemType("vmhbaFirmware")
	VsanUpdateItemTypeVsanUpdateItemType_Unknown = VsanUpdateItemType("VsanUpdateItemType_Unknown")
)

func init() {
	types.Add("vsan:VsanUpdateItemType", reflect.TypeOf((*VsanUpdateItemType)(nil)).Elem())
}

type VsanVcClusterHealthSystemVsanHealthLogLevelEnum string

const (
	VsanVcClusterHealthSystemVsanHealthLogLevelEnumINFO                           = VsanVcClusterHealthSystemVsanHealthLogLevelEnum("INFO")
	VsanVcClusterHealthSystemVsanHealthLogLevelEnumCRITICAL                       = VsanVcClusterHealthSystemVsanHealthLogLevelEnum("CRITICAL")
	VsanVcClusterHealthSystemVsanHealthLogLevelEnumVsanHealthLogLevelEnum_Unknown = VsanVcClusterHealthSystemVsanHealthLogLevelEnum("VsanHealthLogLevelEnum_Unknown")
	VsanVcClusterHealthSystemVsanHealthLogLevelEnumWARNING                        = VsanVcClusterHealthSystemVsanHealthLogLevelEnum("WARNING")
	VsanVcClusterHealthSystemVsanHealthLogLevelEnumERROR                          = VsanVcClusterHealthSystemVsanHealthLogLevelEnum("ERROR")
	VsanVcClusterHealthSystemVsanHealthLogLevelEnumDEBUG                          = VsanVcClusterHealthSystemVsanHealthLogLevelEnum("DEBUG")
)

func init() {
	types.Add("vsan:VsanVcClusterHealthSystemVsanHealthLogLevelEnum", reflect.TypeOf((*VsanVcClusterHealthSystemVsanHealthLogLevelEnum)(nil)).Elem())
}

type VsanVibType string

const (
	VsanVibTypeTool                = VsanVibType("tool")
	VsanVibTypeVsanVibType_Unknown = VsanVibType("VsanVibType_Unknown")
	VsanVibTypeDriver              = VsanVibType("driver")
)

func init() {
	types.Add("vsan:VsanVibType", reflect.TypeOf((*VsanVibType)(nil)).Elem())
}
