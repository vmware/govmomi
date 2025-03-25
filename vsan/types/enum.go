// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"reflect"

	"github.com/vmware/govmomi/vim25/types"
)

type VsanPerfDiagnosticQueryType string

const (
	VsanPerfDiagnosticQueryTypeiops                                = VsanPerfDiagnosticQueryType("iops")
	VsanPerfDiagnosticQueryTypelat                                 = VsanPerfDiagnosticQueryType("lat")
	VsanPerfDiagnosticQueryTypetput                                = VsanPerfDiagnosticQueryType("tput")
	VsanPerfDiagnosticQueryTypeVsanPerfDiagnosticQueryType_Unknown = VsanPerfDiagnosticQueryType("VsanPerfDiagnosticQueryType_Unknown")
	VsanPerfDiagnosticQueryTypeeval                                = VsanPerfDiagnosticQueryType("eval")
)

func init() {
	types.Add("vsan:VsanPerfDiagnosticQueryType", reflect.TypeOf((*VsanPerfDiagnosticQueryType)(nil)).Elem())
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

type VsanCapacityReservationState string

const (
	VsanCapacityReservationStateDisabled      = VsanCapacityReservationState("Disabled")
	VsanCapacityReservationStateState_Unknown = VsanCapacityReservationState("State_Unknown")
	VsanCapacityReservationStateEnforced      = VsanCapacityReservationState("Enforced")
	VsanCapacityReservationStateUnsupported   = VsanCapacityReservationState("Unsupported")
	VsanCapacityReservationStateReported      = VsanCapacityReservationState("Reported")
)

func init() {
	types.Add("vsan:VsanCapacityReservationState", reflect.TypeOf((*VsanCapacityReservationState)(nil)).Elem())
}

type VsanFileShareManagingEntity string

const (
	VsanFileShareManagingEntitycns                             = VsanFileShareManagingEntity("cns")
	VsanFileShareManagingEntityFileShareManagingEntity_Unknown = VsanFileShareManagingEntity("FileShareManagingEntity_Unknown")
	VsanFileShareManagingEntityuser                            = VsanFileShareManagingEntity("user")
)

func init() {
	types.Add("vsan:VsanFileShareManagingEntity", reflect.TypeOf((*VsanFileShareManagingEntity)(nil)).Elem())
}

type VsanObjectTypeEnum string

const (
	VsanObjectTypeEnumfileServiceRoot              = VsanObjectTypeEnum("fileServiceRoot")
	VsanObjectTypeEnumvmswap                       = VsanObjectTypeEnum("vmswap")
	VsanObjectTypeEnumchecksumOverhead             = VsanObjectTypeEnum("checksumOverhead")
	VsanObjectTypeEnumhaMetadataObject             = VsanObjectTypeEnum("haMetadataObject")
	VsanObjectTypeEnumslackSpaceCapRequiredForHost = VsanObjectTypeEnum("slackSpaceCapRequiredForHost")
	VsanObjectTypeEnumdedupOverhead                = VsanObjectTypeEnum("dedupOverhead")
	VsanObjectTypeEnumfileSystemOverhead           = VsanObjectTypeEnum("fileSystemOverhead")
	VsanObjectTypeEnumresynPauseThresholdForHost   = VsanObjectTypeEnum("resynPauseThresholdForHost")
	VsanObjectTypeEnumattachedCnsVolBlock          = VsanObjectTypeEnum("attachedCnsVolBlock")
	VsanObjectTypeEnumspaceUnderDedupConsideration = VsanObjectTypeEnum("spaceUnderDedupConsideration")
	VsanObjectTypeEnumdetachedCnsVolBlock          = VsanObjectTypeEnum("detachedCnsVolBlock")
	VsanObjectTypeEnumminSpaceRequiredForVsanOp    = VsanObjectTypeEnum("minSpaceRequiredForVsanOp")
	VsanObjectTypeEnumiscsiLun                     = VsanObjectTypeEnum("iscsiLun")
	VsanObjectTypeEnumhbrPersist                   = VsanObjectTypeEnum("hbrPersist")
	VsanObjectTypeEnumhostRebuildCapacity          = VsanObjectTypeEnum("hostRebuildCapacity")
	VsanObjectTypeEnumcnsVolFile                   = VsanObjectTypeEnum("cnsVolFile")
	VsanObjectTypeEnumhbrDisk                      = VsanObjectTypeEnum("hbrDisk")
	VsanObjectTypeEnumattachedCnsVolFile           = VsanObjectTypeEnum("attachedCnsVolFile")
	VsanObjectTypeEnumfileShare                    = VsanObjectTypeEnum("fileShare")
	VsanObjectTypeEnumimprovedVirtualDisk          = VsanObjectTypeEnum("improvedVirtualDisk")
	VsanObjectTypeEnumvdisk                        = VsanObjectTypeEnum("vdisk")
	VsanObjectTypeEnumVsanObjectTypeEnum_Unknown   = VsanObjectTypeEnum("VsanObjectTypeEnum_Unknown")
	VsanObjectTypeEnumnamespace                    = VsanObjectTypeEnum("namespace")
	VsanObjectTypeEnumstatsdb                      = VsanObjectTypeEnum("statsdb")
	VsanObjectTypeEnumvmem                         = VsanObjectTypeEnum("vmem")
	VsanObjectTypeEnumother                        = VsanObjectTypeEnum("other")
	VsanObjectTypeEnumextension                    = VsanObjectTypeEnum("extension")
	VsanObjectTypeEnumtransientSpace               = VsanObjectTypeEnum("transientSpace")
	VsanObjectTypeEnumhbrCfg                       = VsanObjectTypeEnum("hbrCfg")
	VsanObjectTypeEnumphysicalTransientSpace       = VsanObjectTypeEnum("physicalTransientSpace")
	VsanObjectTypeEnumiscsiTarget                  = VsanObjectTypeEnum("iscsiTarget")
	VsanObjectTypeEnumdetachedCnsVolFile           = VsanObjectTypeEnum("detachedCnsVolFile")
)

func init() {
	types.Add("vsan:VsanObjectTypeEnum", reflect.TypeOf((*VsanObjectTypeEnum)(nil)).Elem())
}

type VsanPerfsvcRemediateAction string

const (
	VsanPerfsvcRemediateActionupdate_profile                 = VsanPerfsvcRemediateAction("update_profile")
	VsanPerfsvcRemediateActionPerfsvcRemediateAction_Unknown = VsanPerfsvcRemediateAction("PerfsvcRemediateAction_Unknown")
	VsanPerfsvcRemediateActionenable                         = VsanPerfsvcRemediateAction("enable")
	VsanPerfsvcRemediateActiondisable                        = VsanPerfsvcRemediateAction("disable")
	VsanPerfsvcRemediateActionno_action                      = VsanPerfsvcRemediateAction("no_action")
)

func init() {
	types.Add("vsan:VsanPerfsvcRemediateAction", reflect.TypeOf((*VsanPerfsvcRemediateAction)(nil)).Elem())
}

type VsanIoInsightInstanceState string

const (
	VsanIoInsightInstanceStatecrashed                            = VsanIoInsightInstanceState("crashed")
	VsanIoInsightInstanceStaterunning                            = VsanIoInsightInstanceState("running")
	VsanIoInsightInstanceStatecompleted                          = VsanIoInsightInstanceState("completed")
	VsanIoInsightInstanceStateVsanIoInsightInstanceState_unknown = VsanIoInsightInstanceState("VsanIoInsightInstanceState_unknown")
)

func init() {
	types.Add("vsan:VsanIoInsightInstanceState", reflect.TypeOf((*VsanIoInsightInstanceState)(nil)).Elem())
}

type VsanUpdateItemImpactType string

const (
	VsanUpdateItemImpactTypeVsanUpdateItemImpactType_Unknown = VsanUpdateItemImpactType("VsanUpdateItemImpactType_Unknown")
	VsanUpdateItemImpactTypereboot                           = VsanUpdateItemImpactType("reboot")
)

func init() {
	types.Add("vsan:VsanUpdateItemImpactType", reflect.TypeOf((*VsanUpdateItemImpactType)(nil)).Elem())
}

type VsanUpdateItemType string

const (
	VsanUpdateItemTypevib                        = VsanUpdateItemType("vib")
	VsanUpdateItemTypeofflinebundle              = VsanUpdateItemType("offlinebundle")
	VsanUpdateItemTypefullStackFirmware          = VsanUpdateItemType("fullStackFirmware")
	VsanUpdateItemTypevmhbaFirmware              = VsanUpdateItemType("vmhbaFirmware")
	VsanUpdateItemTypeVsanUpdateItemType_Unknown = VsanUpdateItemType("VsanUpdateItemType_Unknown")
)

func init() {
	types.Add("vsan:VsanUpdateItemType", reflect.TypeOf((*VsanUpdateItemType)(nil)).Elem())
}

type VsanEncryptionIssue string

const (
	VsanEncryptionIssuekeyencryptionkeyinconsistent    = VsanEncryptionIssue("keyencryptionkeyinconsistent")
	VsanEncryptionIssuecmknotinenabledstate            = VsanEncryptionIssue("cmknotinenabledstate")
	VsanEncryptionIssueclientkeyinconsistent           = VsanEncryptionIssue("clientkeyinconsistent")
	VsanEncryptionIssuekeknotavailable                 = VsanEncryptionIssue("keknotavailable")
	VsanEncryptionIssuehostkeynotavailable             = VsanEncryptionIssue("hostkeynotavailable")
	VsanEncryptionIssueservercertificatesinconsistent  = VsanEncryptionIssue("servercertificatesinconsistent")
	VsanEncryptionIssueVsanEncryptionIssue_Unknown     = VsanEncryptionIssue("VsanEncryptionIssue_Unknown")
	VsanEncryptionIssuedataencryptionkeyinconsistent   = VsanEncryptionIssue("dataencryptionkeyinconsistent")
	VsanEncryptionIssuehostkeyinconsistent             = VsanEncryptionIssue("hostkeyinconsistent")
	VsanEncryptionIssueerasedisksbeforeuseinconsistent = VsanEncryptionIssue("erasedisksbeforeuseinconsistent")
	VsanEncryptionIssueclientcertificateinconsistent   = VsanEncryptionIssue("clientcertificateinconsistent")
	VsanEncryptionIssuecmkcannotretrieve               = VsanEncryptionIssue("cmkcannotretrieve")
	VsanEncryptionIssuekmsinfoinconsistent             = VsanEncryptionIssue("kmsinfoinconsistent")
	VsanEncryptionIssueenabledwhenclusterdisabled      = VsanEncryptionIssue("enabledwhenclusterdisabled")
	VsanEncryptionIssuedisabledwhenclusterenabled      = VsanEncryptionIssue("disabledwhenclusterenabled")
)

func init() {
	types.Add("vsan:VsanEncryptionIssue", reflect.TypeOf((*VsanEncryptionIssue)(nil)).Elem())
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

type VimVsanHostDiskMappingCreationType string

const (
	VimVsanHostDiskMappingCreationTypeDiskMappingCreationType_Unknown = VimVsanHostDiskMappingCreationType("DiskMappingCreationType_Unknown")
	VimVsanHostDiskMappingCreationTypeallFlash                        = VimVsanHostDiskMappingCreationType("allFlash")
	VimVsanHostDiskMappingCreationTypepmem                            = VimVsanHostDiskMappingCreationType("pmem")
	VimVsanHostDiskMappingCreationTypehybrid                          = VimVsanHostDiskMappingCreationType("hybrid")
	VimVsanHostDiskMappingCreationTypevsandirect                      = VimVsanHostDiskMappingCreationType("vsandirect")
)

func init() {
	types.Add("vsan:VimVsanHostDiskMappingCreationType", reflect.TypeOf((*VimVsanHostDiskMappingCreationType)(nil)).Elem())
}

type VsanDiskBalanceState string

const (
	VsanDiskBalanceStatereactiverebalancefailed      = VsanDiskBalanceState("reactiverebalancefailed")
	VsanDiskBalanceStateproactivenotmustdo           = VsanDiskBalanceState("proactivenotmustdo")
	VsanDiskBalanceStaterebalancediskunhealthy       = VsanDiskBalanceState("rebalancediskunhealthy")
	VsanDiskBalanceStateimbalancewithintolerance     = VsanDiskBalanceState("imbalancewithintolerance")
	VsanDiskBalanceStateproactiverebalancefailed     = VsanDiskBalanceState("proactiverebalancefailed")
	VsanDiskBalanceStaterebalanceentitydecom         = VsanDiskBalanceState("rebalanceentitydecom")
	VsanDiskBalanceStateproactiveneededbutdisabled   = VsanDiskBalanceState("proactiveneededbutdisabled")
	VsanDiskBalanceStateproactiverebalanceinprogress = VsanDiskBalanceState("proactiverebalanceinprogress")
	VsanDiskBalanceStaterebalanceoff                 = VsanDiskBalanceState("rebalanceoff")
	VsanDiskBalanceStatereactiverebalanceinprogress  = VsanDiskBalanceState("reactiverebalanceinprogress")
	VsanDiskBalanceStateVsanDiskBalanceState_Unknown = VsanDiskBalanceState("VsanDiskBalanceState_Unknown")
)

func init() {
	types.Add("vsan:VsanDiskBalanceState", reflect.TypeOf((*VsanDiskBalanceState)(nil)).Elem())
}

type VsanFileShareSmbEncryptionType string

const (
	VsanFileShareSmbEncryptionTypedisabled                           = VsanFileShareSmbEncryptionType("disabled")
	VsanFileShareSmbEncryptionTypemandatory                          = VsanFileShareSmbEncryptionType("mandatory")
	VsanFileShareSmbEncryptionTypeFileShareSmbEncryptionType_Unknown = VsanFileShareSmbEncryptionType("FileShareSmbEncryptionType_Unknown")
)

func init() {
	types.Add("vsan:VsanFileShareSmbEncryptionType", reflect.TypeOf((*VsanFileShareSmbEncryptionType)(nil)).Elem())
}

type VsanSiteLocationType string

const (
	VsanSiteLocationTypeNone                         = VsanSiteLocationType("None")
	VsanSiteLocationTypeVsanSiteLocationType_Unknown = VsanSiteLocationType("VsanSiteLocationType_Unknown")
	VsanSiteLocationTypeNonPreferred                 = VsanSiteLocationType("NonPreferred")
	VsanSiteLocationTypePreferred                    = VsanSiteLocationType("Preferred")
)

func init() {
	types.Add("vsan:VsanSiteLocationType", reflect.TypeOf((*VsanSiteLocationType)(nil)).Elem())
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

type VsanHostWipeDiskEligible string

const (
	VsanHostWipeDiskEligibleUnknown                  = VsanHostWipeDiskEligible("Unknown")
	VsanHostWipeDiskEligibleYes                      = VsanHostWipeDiskEligible("Yes")
	VsanHostWipeDiskEligibleWipeDiskEligible_Unknown = VsanHostWipeDiskEligible("WipeDiskEligible_Unknown")
	VsanHostWipeDiskEligibleNo                       = VsanHostWipeDiskEligible("No")
)

func init() {
	types.Add("vsan:VsanHostWipeDiskEligible", reflect.TypeOf((*VsanHostWipeDiskEligible)(nil)).Elem())
}

type VimVsanMountPrecheckType string

const (
	VimVsanMountPrecheckTypelocalVsanDatastore     = VimVsanMountPrecheckType("localVsanDatastore")
	VimVsanMountPrecheckTypenetworkLatency         = VimVsanMountPrecheckType("networkLatency")
	VimVsanMountPrecheckTyperemoteDatastoreLimit   = VimVsanMountPrecheckType("remoteDatastoreLimit")
	VimVsanMountPrecheckTypedatastorePolicy        = VimVsanMountPrecheckType("datastorePolicy")
	VimVsanMountPrecheckTypeconnectivity           = VimVsanMountPrecheckType("connectivity")
	VimVsanMountPrecheckTypeclientClusterLimit     = VimVsanMountPrecheckType("clientClusterLimit")
	VimVsanMountPrecheckTypedatacenter             = VimVsanMountPrecheckType("datacenter")
	VimVsanMountPrecheckTypesupportedConfiguration = VimVsanMountPrecheckType("supportedConfiguration")
	VimVsanMountPrecheckTypeserverClusterHealth    = VimVsanMountPrecheckType("serverClusterHealth")
	VimVsanMountPrecheckTypevsanFormatVersion      = VimVsanMountPrecheckType("vsanFormatVersion")
	VimVsanMountPrecheckTypelicense                = VimVsanMountPrecheckType("license")
	VimVsanMountPrecheckTypedatastoreType          = VimVsanMountPrecheckType("datastoreType")
	VimVsanMountPrecheckTypeserverClusterLimit     = VimVsanMountPrecheckType("serverClusterLimit")
	VimVsanMountPrecheckTypeprecheck_unknown       = VimVsanMountPrecheckType("precheck_unknown")
)

func init() {
	types.Add("vsan:VimVsanMountPrecheckType", reflect.TypeOf((*VimVsanMountPrecheckType)(nil)).Elem())
}

type VsanEncryptionTransitionState string

const (
	VsanEncryptionTransitionStateEncryptionTransitionState_Unknown = VsanEncryptionTransitionState("EncryptionTransitionState_Unknown")
	VsanEncryptionTransitionStatesettled                           = VsanEncryptionTransitionState("settled")
	VsanEncryptionTransitionStatepreparing                         = VsanEncryptionTransitionState("preparing")
	VsanEncryptionTransitionStateprepared                          = VsanEncryptionTransitionState("prepared")
)

func init() {
	types.Add("vsan:VsanEncryptionTransitionState", reflect.TypeOf((*VsanEncryptionTransitionState)(nil)).Elem())
}

type VimVsanClusterComplianceResourceCheckStatusType string

const (
	VimVsanClusterComplianceResourceCheckStatusTypeuninitialized                             = VimVsanClusterComplianceResourceCheckStatusType("uninitialized")
	VimVsanClusterComplianceResourceCheckStatusTypeinProgress                                = VimVsanClusterComplianceResourceCheckStatusType("inProgress")
	VimVsanClusterComplianceResourceCheckStatusTypeComplianceResourceCheckStatusType_Unknown = VimVsanClusterComplianceResourceCheckStatusType("ComplianceResourceCheckStatusType_Unknown")
	VimVsanClusterComplianceResourceCheckStatusTypecompleted                                 = VimVsanClusterComplianceResourceCheckStatusType("completed")
	VimVsanClusterComplianceResourceCheckStatusTypeaborted                                   = VimVsanClusterComplianceResourceCheckStatusType("aborted")
)

func init() {
	types.Add("vsan:VimVsanClusterComplianceResourceCheckStatusType", reflect.TypeOf((*VimVsanClusterComplianceResourceCheckStatusType)(nil)).Elem())
}

type VsanIscsiLUNStatus string

const (
	VsanIscsiLUNStatusOffline                    = VsanIscsiLUNStatus("Offline")
	VsanIscsiLUNStatusVsanIscsiLUNStatus_Unknown = VsanIscsiLUNStatus("VsanIscsiLUNStatus_Unknown")
	VsanIscsiLUNStatusOnline                     = VsanIscsiLUNStatus("Online")
)

func init() {
	types.Add("vsan:VsanIscsiLUNStatus", reflect.TypeOf((*VsanIscsiLUNStatus)(nil)).Elem())
}

type VsanCapabilityType string

const (
	VsanCapabilityTypediagnosticmode                 = VsanCapabilityType("diagnosticmode")
	VsanCapabilityTypeobjectidentities               = VsanCapabilityType("objectidentities")
	VsanCapabilityTypesharedwitness                  = VsanCapabilityType("sharedwitness")
	VsanCapabilityTypevumbaselinerecommendation      = VsanCapabilityType("vumbaselinerecommendation")
	VsanCapabilityTypeupgrade                        = VsanCapabilityType("upgrade")
	VsanCapabilityTypevitstretchedcluster            = VsanCapabilityType("vitstretchedcluster")
	VsanCapabilityTypeenhancedresyncapi              = VsanCapabilityType("enhancedresyncapi")
	VsanCapabilityTypepolicyhostapi                  = VsanCapabilityType("policyhostapi")
	VsanCapabilityTypefileservicecrx                 = VsanCapabilityType("fileservicecrx")
	VsanCapabilityTypecnsvolumes                     = VsanCapabilityType("cnsvolumes")
	VsanCapabilityTypethrottleresync                 = VsanCapabilityType("throttleresync")
	VsanCapabilityTypeverbosemodeconfiguration       = VsanCapabilityType("verbosemodeconfiguration")
	VsanCapabilityTypelargecapacitydrive             = VsanCapabilityType("largecapacitydrive")
	VsanCapabilityTypeiscsitargets                   = VsanCapabilityType("iscsitargets")
	VsanCapabilityTypecapacityoversubscription       = VsanCapabilityType("capacityoversubscription")
	VsanCapabilityTypevsanencrkmx                    = VsanCapabilityType("vsanencrkmx")
	VsanCapabilityTypepurgeinaccessiblevmswapobjects = VsanCapabilityType("purgeinaccessiblevmswapobjects")
	VsanCapabilityTypevsanclient                     = VsanCapabilityType("vsanclient")
	VsanCapabilityTypevsandefaultgatewaysupported    = VsanCapabilityType("vsandefaultgatewaysupported")
	VsanCapabilityTyperesyncetaimprovement           = VsanCapabilityType("resyncetaimprovement")
	VsanCapabilityTypevmlevelcapacity                = VsanCapabilityType("vmlevelcapacity")
	VsanCapabilityTypevitonlineresize                = VsanCapabilityType("vitonlineresize")
	VsanCapabilityTypevsanrdma                       = VsanCapabilityType("vsanrdma")
	VsanCapabilityTypesecurewipe                     = VsanCapabilityType("securewipe")
	VsanCapabilityTypedataefficiency                 = VsanCapabilityType("dataefficiency")
	VsanCapabilityTypemetricsconfig                  = VsanCapabilityType("metricsconfig")
	VsanCapabilityTypehistoricalcapacity             = VsanCapabilityType("historicalcapacity")
	VsanCapabilityTypeallflash                       = VsanCapabilityType("allflash")
	VsanCapabilityTypeioinsight                      = VsanCapabilityType("ioinsight")
	VsanCapabilityTypeunicasttest                    = VsanCapabilityType("unicasttest")
	VsanCapabilityTypewcpappplatform                 = VsanCapabilityType("wcpappplatform")
	VsanCapabilityTypeVsanFileAnalytics              = VsanCapabilityType("VsanFileAnalytics")
	VsanCapabilityTypefileservicesmb                 = VsanCapabilityType("fileservicesmb")
	VsanCapabilityTypenestedfd                       = VsanCapabilityType("nestedfd")
	VsanCapabilityTypepr1741414fixed                 = VsanCapabilityType("pr1741414fixed")
	VsanCapabilityTypedit4sw                         = VsanCapabilityType("dit4sw")
	VsanCapabilityTypegethcllastupdateonvc           = VsanCapabilityType("gethcllastupdateonvc")
	VsanCapabilityTypecapability                     = VsanCapabilityType("capability")
	VsanCapabilityTypedecomwhatif                    = VsanCapabilityType("decomwhatif")
	VsanCapabilityTypeclusterconfig                  = VsanCapabilityType("clusterconfig")
	VsanCapabilityTypevsandiagnostics                = VsanCapabilityType("vsandiagnostics")
	VsanCapabilityTypepolicyassociation              = VsanCapabilityType("policyassociation")
	VsanCapabilityTypesupportinsight                 = VsanCapabilityType("supportinsight")
	VsanCapabilityTypeperfsvcautoconfig              = VsanCapabilityType("perfsvcautoconfig")
	VsanCapabilityTypegenericnestedfd                = VsanCapabilityType("genericnestedfd")
	VsanCapabilityTypeperfsvcverbosemode             = VsanCapabilityType("perfsvcverbosemode")
	VsanCapabilityTypefilevolumes                    = VsanCapabilityType("filevolumes")
	VsanCapabilityTypeupdatevumreleasecatalogoffline = VsanCapabilityType("updatevumreleasecatalogoffline")
	VsanCapabilityTyperesourceprecheck               = VsanCapabilityType("resourceprecheck")
	VsanCapabilityTypeunicastmode                    = VsanCapabilityType("unicastmode")
	VsanCapabilityTypefileservicesc                  = VsanCapabilityType("fileservicesc")
	VsanCapabilityTypehardwaremgmt                   = VsanCapabilityType("hardwaremgmt")
	VsanCapabilityTypehealthcheck2018q2              = VsanCapabilityType("healthcheck2018q2")
	VsanCapabilityTypeperformanceforsupport          = VsanCapabilityType("performanceforsupport")
	VsanCapabilityTypefirmwareupdate                 = VsanCapabilityType("firmwareupdate")
	VsanCapabilityTypeimprovedcapacityscreen         = VsanCapabilityType("improvedcapacityscreen")
	VsanCapabilityTypevalidateconfigspec             = VsanCapabilityType("validateconfigspec")
	VsanCapabilityTypediskresourceprecheck           = VsanCapabilityType("diskresourceprecheck")
	VsanCapabilityTypedevice4ksupport                = VsanCapabilityType("device4ksupport")
	VsanCapabilityTypevsanmanagedvmfs                = VsanCapabilityType("vsanmanagedvmfs")
	VsanCapabilityTypefullStackFw                    = VsanCapabilityType("fullStackFw")
	VsanCapabilityTypemasspropertycollector          = VsanCapabilityType("masspropertycollector")
	VsanCapabilityTypenondatamovementdfc             = VsanCapabilityType("nondatamovementdfc")
	VsanCapabilityTypevumintegration                 = VsanCapabilityType("vumintegration")
	VsanCapabilityTyperemotedatastore                = VsanCapabilityType("remotedatastore")
	VsanCapabilityTypeencryption                     = VsanCapabilityType("encryption")
	VsanCapabilityTypehostreservedcapacity           = VsanCapabilityType("hostreservedcapacity")
	VsanCapabilityTypefileservicenfsv3               = VsanCapabilityType("fileservicenfsv3")
	VsanCapabilityTypenetperftest                    = VsanCapabilityType("netperftest")
	VsanCapabilityTypeslackspacecapacity             = VsanCapabilityType("slackspacecapacity")
	VsanCapabilityTypevsananalyticsevents            = VsanCapabilityType("vsananalyticsevents")
	VsanCapabilityTypewhatifcapacity                 = VsanCapabilityType("whatifcapacity")
	VsanCapabilityTypereadlocalitytodrs              = VsanCapabilityType("readlocalitytodrs")
	VsanCapabilityTypeautomaticrebalance             = VsanCapabilityType("automaticrebalance")
	VsanCapabilityTypecompressiononly                = VsanCapabilityType("compressiononly")
	VsanCapabilityTypeumap                           = VsanCapabilityType("umap")
	VsanCapabilityTypefileservicekerberos            = VsanCapabilityType("fileservicekerberos")
	VsanCapabilityTypedataintransitencryption        = VsanCapabilityType("dataintransitencryption")
	VsanCapabilityTyperecreatediskgroup              = VsanCapabilityType("recreatediskgroup")
	VsanCapabilityTypeconfigassist                   = VsanCapabilityType("configassist")
	VsanCapabilityTypeupgraderesourceprecheck        = VsanCapabilityType("upgraderesourceprecheck")
	VsanCapabilityTypelocaldataprotection            = VsanCapabilityType("localdataprotection")
	VsanCapabilityTypeapidevversionenabled           = VsanCapabilityType("apidevversionenabled")
	VsanCapabilityTypeclusteradvancedoptions         = VsanCapabilityType("clusteradvancedoptions")
	VsanCapabilityTypeensuredurability               = VsanCapabilityType("ensuredurability")
	VsanCapabilityTypefileserviceowe                 = VsanCapabilityType("fileserviceowe")
	VsanCapabilityTypehostaffinity                   = VsanCapabilityType("hostaffinity")
	VsanCapabilityTypepmanintegration                = VsanCapabilityType("pmanintegration")
	VsanCapabilityTypewitnessmanagement              = VsanCapabilityType("witnessmanagement")
	VsanCapabilityTypenativelargeclustersupport      = VsanCapabilityType("nativelargeclustersupport")
	VsanCapabilityTypecapacityreservation            = VsanCapabilityType("capacityreservation")
	VsanCapabilityTypeperfsvctwoyaxisgraph           = VsanCapabilityType("perfsvctwoyaxisgraph")
	VsanCapabilityTypecloudhealth                    = VsanCapabilityType("cloudhealth")
	VsanCapabilityTypeidentitiessupportpolicyid      = VsanCapabilityType("identitiessupportpolicyid")
	VsanCapabilityTypefileservices                   = VsanCapabilityType("fileservices")
	VsanCapabilityTypeVsanCapabilityType_Unknown     = VsanCapabilityType("VsanCapabilityType_Unknown")
	VsanCapabilityTypevsanmetadatanode               = VsanCapabilityType("vsanmetadatanode")
	VsanCapabilityTypediagnosticsfeedback            = VsanCapabilityType("diagnosticsfeedback")
	VsanCapabilityTypefileservicesnapshot            = VsanCapabilityType("fileservicesnapshot")
	VsanCapabilityTypehistoricalhealth               = VsanCapabilityType("historicalhealth")
	VsanCapabilityTypevsanmanagedpmem                = VsanCapabilityType("vsanmanagedpmem")
	VsanCapabilityTyperemotedataprotection           = VsanCapabilityType("remotedataprotection")
	VsanCapabilityTypecapacityevaluationonvc         = VsanCapabilityType("capacityevaluationonvc")
	VsanCapabilityTypestretchedcluster               = VsanCapabilityType("stretchedcluster")
	VsanCapabilityTypepspairgap                      = VsanCapabilityType("pspairgap")
	VsanCapabilityTypearchivaldataprotection         = VsanCapabilityType("archivaldataprotection")
	VsanCapabilityTypecomplianceprecheck             = VsanCapabilityType("complianceprecheck")
	VsanCapabilityTypefcd                            = VsanCapabilityType("fcd")
	VsanCapabilityTypesupportApiVersion              = VsanCapabilityType("supportApiVersion")
	VsanCapabilityTyperepairtimerinresyncstats       = VsanCapabilityType("repairtimerinresyncstats")
	VsanCapabilityTypeperfanalysis                   = VsanCapabilityType("perfanalysis")
)

func init() {
	types.Add("vsan:VsanCapabilityType", reflect.TypeOf((*VsanCapabilityType)(nil)).Elem())
}

type VsanVibType string

const (
	VsanVibTypetool                = VsanVibType("tool")
	VsanVibTypeVsanVibType_Unknown = VsanVibType("VsanVibType_Unknown")
	VsanVibTypedriver              = VsanVibType("driver")
)

func init() {
	types.Add("vsan:VsanVibType", reflect.TypeOf((*VsanVibType)(nil)).Elem())
}

type VsanRelayoutObjectsErrorCode string

const (
	VsanRelayoutObjectsErrorCodeoutOfResources                       = VsanRelayoutObjectsErrorCode("outOfResources")
	VsanRelayoutObjectsErrorCodegeneric                              = VsanRelayoutObjectsErrorCode("generic")
	VsanRelayoutObjectsErrorCodeVsanRelayoutObjectsErrorCode_Unknown = VsanRelayoutObjectsErrorCode("VsanRelayoutObjectsErrorCode_Unknown")
)

func init() {
	types.Add("vsan:VsanRelayoutObjectsErrorCode", reflect.TypeOf((*VsanRelayoutObjectsErrorCode)(nil)).Elem())
}

type VsanBaselinePreferenceType string

const (
	VsanBaselinePreferenceTypenoRecommendation                   = VsanBaselinePreferenceType("noRecommendation")
	VsanBaselinePreferenceTypelatestRelease                      = VsanBaselinePreferenceType("latestRelease")
	VsanBaselinePreferenceTypelatestPatch                        = VsanBaselinePreferenceType("latestPatch")
	VsanBaselinePreferenceTypeVsanBaselinePreferenceType_Unknown = VsanBaselinePreferenceType("VsanBaselinePreferenceType_Unknown")
)

func init() {
	types.Add("vsan:VsanBaselinePreferenceType", reflect.TypeOf((*VsanBaselinePreferenceType)(nil)).Elem())
}

type VsanStorageComplianceStatus string

const (
	VsanStorageComplianceStatusunknown       = VsanStorageComplianceStatus("unknown")
	VsanStorageComplianceStatuscompliant     = VsanStorageComplianceStatus("compliant")
	VsanStorageComplianceStatusnonCompliant  = VsanStorageComplianceStatus("nonCompliant")
	VsanStorageComplianceStatusnotApplicable = VsanStorageComplianceStatus("notApplicable")
)

func init() {
	types.Add("vsan:VsanStorageComplianceStatus", reflect.TypeOf((*VsanStorageComplianceStatus)(nil)).Elem())
}

type VsanHealthStatusType string

const (
	VsanHealthStatusTypeunknown = VsanHealthStatusType("unknown")
	VsanHealthStatusTypegreen   = VsanHealthStatusType("green")
	VsanHealthStatusTypered     = VsanHealthStatusType("red")
	VsanHealthStatusTypeyellow  = VsanHealthStatusType("yellow")
)

func init() {
	types.Add("vsan:VsanHealthStatusType", reflect.TypeOf((*VsanHealthStatusType)(nil)).Elem())
}

type VsanPerfStatsType string

const (
	VsanPerfStatsTypeVsanPerfStatsType_Unknown = VsanPerfStatsType("VsanPerfStatsType_Unknown")
	VsanPerfStatsTyperate                      = VsanPerfStatsType("rate")
	VsanPerfStatsTypedelta                     = VsanPerfStatsType("delta")
	VsanPerfStatsTypeabsolute                  = VsanPerfStatsType("absolute")
)

func init() {
	types.Add("vsan:VsanPerfStatsType", reflect.TypeOf((*VsanPerfStatsType)(nil)).Elem())
}

type VsanFileProtocol string

const (
	VsanFileProtocolNFSv4                     = VsanFileProtocol("NFSv4")
	VsanFileProtocolFileShareProtocol_Unknown = VsanFileProtocol("FileShareProtocol_Unknown")
	VsanFileProtocolSMB                       = VsanFileProtocol("SMB")
	VsanFileProtocolNFSv3                     = VsanFileProtocol("NFSv3")
)

func init() {
	types.Add("vsan:VsanFileProtocol", reflect.TypeOf((*VsanFileProtocol)(nil)).Elem())
}

type VsanResourceCheckStatusType string

const (
	VsanResourceCheckStatusTyperesourceCheckCompleted          = VsanResourceCheckStatusType("resourceCheckCompleted")
	VsanResourceCheckStatusTyperesourceCheckNotSupported       = VsanResourceCheckStatusType("resourceCheckNotSupported")
	VsanResourceCheckStatusTyperesourceCheckCancelled          = VsanResourceCheckStatusType("resourceCheckCancelled")
	VsanResourceCheckStatusTypeResourceCheckStatusType_Unknown = VsanResourceCheckStatusType("ResourceCheckStatusType_Unknown")
	VsanResourceCheckStatusTyperesourceCheckFailed             = VsanResourceCheckStatusType("resourceCheckFailed")
	VsanResourceCheckStatusTyperesourceCheckNoRecentValue      = VsanResourceCheckStatusType("resourceCheckNoRecentValue")
	VsanResourceCheckStatusTyperesourceCheckUninitialized      = VsanResourceCheckStatusType("resourceCheckUninitialized")
	VsanResourceCheckStatusTyperesourceCheckRunning            = VsanResourceCheckStatusType("resourceCheckRunning")
)

func init() {
	types.Add("vsan:VsanResourceCheckStatusType", reflect.TypeOf((*VsanResourceCheckStatusType)(nil)).Elem())
}

type VsanServiceStatus string

const (
	VsanServiceStatusstarted                   = VsanServiceStatus("started")
	VsanServiceStatusstopped                   = VsanServiceStatus("stopped")
	VsanServiceStatusVsanServiceStatus_Unknown = VsanServiceStatus("VsanServiceStatus_Unknown")
)

func init() {
	types.Add("vsan:VsanServiceStatus", reflect.TypeOf((*VsanServiceStatus)(nil)).Elem())
}

type VsanObjectHealthState string

const (
	VsanObjectHealthStateVsanObjectHealthState_Unknown                             = VsanObjectHealthState("VsanObjectHealthState_Unknown")
	VsanObjectHealthStatereducedavailabilitywithnorebuilddelaytimer                = VsanObjectHealthState("reducedavailabilitywithnorebuilddelaytimer")
	VsanObjectHealthStatereducedavailabilitywithpausedrebuild                      = VsanObjectHealthState("reducedavailabilitywithpausedrebuild")
	VsanObjectHealthStatenonavailabilityrelatedincompliancewithpolicypendingfailed = VsanObjectHealthState("nonavailabilityrelatedincompliancewithpolicypendingfailed")
	VsanObjectHealthStatenonavailabilityrelatedincompliancewithpausedrebuild       = VsanObjectHealthState("nonavailabilityrelatedincompliancewithpausedrebuild")
	VsanObjectHealthStatehealthy                                                   = VsanObjectHealthState("healthy")
	VsanObjectHealthStateinaccessible                                              = VsanObjectHealthState("inaccessible")
	VsanObjectHealthStatereducedavailabilitywithactiverebuild                      = VsanObjectHealthState("reducedavailabilitywithactiverebuild")
	VsanObjectHealthStatedatamove                                                  = VsanObjectHealthState("datamove")
	VsanObjectHealthStateremoteAccessible                                          = VsanObjectHealthState("remoteAccessible")
	VsanObjectHealthStatenonavailabilityrelatedincompliancewithpolicypending       = VsanObjectHealthState("nonavailabilityrelatedincompliancewithpolicypending")
	VsanObjectHealthStatereducedavailabilitywithpolicypending                      = VsanObjectHealthState("reducedavailabilitywithpolicypending")
	VsanObjectHealthStatenonavailabilityrelatedincompliance                        = VsanObjectHealthState("nonavailabilityrelatedincompliance")
	VsanObjectHealthStatereducedavailabilitywithpolicypendingfailed                = VsanObjectHealthState("reducedavailabilitywithpolicypendingfailed")
	VsanObjectHealthStatenonavailabilityrelatedreconfig                            = VsanObjectHealthState("nonavailabilityrelatedreconfig")
	VsanObjectHealthStatereducedavailabilitywithnorebuild                          = VsanObjectHealthState("reducedavailabilitywithnorebuild")
)

func init() {
	types.Add("vsan:VsanObjectHealthState", reflect.TypeOf((*VsanObjectHealthState)(nil)).Elem())
}

type VsanHealthPerspective string

const (
	VsanHealthPerspectiveupgradeBeforeExitMM           = VsanHealthPerspective("upgradeBeforeExitMM")
	VsanHealthPerspectiveupgradePreCheck               = VsanHealthPerspective("upgradePreCheck")
	VsanHealthPerspectiveupgradePreCheckPman           = VsanHealthPerspective("upgradePreCheckPman")
	VsanHealthPerspectiveupgradeAfterExitMM            = VsanHealthPerspective("upgradeAfterExitMM")
	VsanHealthPerspectiveupgradeBeforeExitMMPman       = VsanHealthPerspective("upgradeBeforeExitMMPman")
	VsanHealthPerspectivebeforeConfigureHost           = VsanHealthPerspective("beforeConfigureHost")
	VsanHealthPerspectivedefaultView                   = VsanHealthPerspective("defaultView")
	VsanHealthPerspectivevsanUpgradeAfterExitMM        = VsanHealthPerspective("vsanUpgradeAfterExitMM")
	VsanHealthPerspectivedeployAssist                  = VsanHealthPerspective("deployAssist")
	VsanHealthPerspectivevsanUpgradePreCheck           = VsanHealthPerspective("vsanUpgradePreCheck")
	VsanHealthPerspectiveVsanHealthPerspective_Unknown = VsanHealthPerspective("VsanHealthPerspective_Unknown")
	VsanHealthPerspectiveupgradeAfterExitMMPman        = VsanHealthPerspective("upgradeAfterExitMMPman")
	VsanHealthPerspectiveCreateExtendClusterView       = VsanHealthPerspective("CreateExtendClusterView")
	VsanHealthPerspectivevsanUpgradeBeforeExitMM       = VsanHealthPerspective("vsanUpgradeBeforeExitMM")
	VsanHealthPerspectivevmcUpgradePreChecks           = VsanHealthPerspective("vmcUpgradePreChecks")
)

func init() {
	types.Add("vsan:VsanHealthPerspective", reflect.TypeOf((*VsanHealthPerspective)(nil)).Elem())
}

type VsanDatastoreType string

const (
	VsanDatastoreTypevsandirect                = VsanDatastoreType("vsandirect")
	VsanDatastoreTypevsan                      = VsanDatastoreType("vsan")
	VsanDatastoreTypeVsanDatastoreType_Unknown = VsanDatastoreType("VsanDatastoreType_Unknown")
	VsanDatastoreTypepmem                      = VsanDatastoreType("pmem")
)

func init() {
	types.Add("vsan:VsanDatastoreType", reflect.TypeOf((*VsanDatastoreType)(nil)).Elem())
}

type VsanSyncReason string

const (
	VsanSyncReasonobject_format_change   = VsanSyncReason("object_format_change")
	VsanSyncReasonrepair                 = VsanSyncReason("repair")
	VsanSyncReasondying_evacuate         = VsanSyncReason("dying_evacuate")
	VsanSyncReasonreconfigure            = VsanSyncReason("reconfigure")
	VsanSyncReasonVsanSyncReason_Unknown = VsanSyncReason("VsanSyncReason_Unknown")
	VsanSyncReasonstale                  = VsanSyncReason("stale")
	VsanSyncReasonrebalance              = VsanSyncReason("rebalance")
	VsanSyncReasonevacuate               = VsanSyncReason("evacuate")
	VsanSyncReasonmerge_concat           = VsanSyncReason("merge_concat")
)

func init() {
	types.Add("vsan:VsanSyncReason", reflect.TypeOf((*VsanSyncReason)(nil)).Elem())
}

type VsanHealthLogLevelEnum string

const (
	VsanHealthLogLevelEnumINFO                           = VsanHealthLogLevelEnum("INFO")
	VsanHealthLogLevelEnumCRITICAL                       = VsanHealthLogLevelEnum("CRITICAL")
	VsanHealthLogLevelEnumVsanHealthLogLevelEnum_Unknown = VsanHealthLogLevelEnum("VsanHealthLogLevelEnum_Unknown")
	VsanHealthLogLevelEnumWARNING                        = VsanHealthLogLevelEnum("WARNING")
	VsanHealthLogLevelEnumERROR                          = VsanHealthLogLevelEnum("ERROR")
	VsanHealthLogLevelEnumDEBUG                          = VsanHealthLogLevelEnum("DEBUG")
)

func init() {
	types.Add("vsan:VsanHealthLogLevelEnum", reflect.TypeOf((*VsanHealthLogLevelEnum)(nil)).Elem())
}

type VsanPerfSummaryType string

const (
	VsanPerfSummaryTypenone                        = VsanPerfSummaryType("none")
	VsanPerfSummaryTypeaverage                     = VsanPerfSummaryType("average")
	VsanPerfSummaryTypemaximum                     = VsanPerfSummaryType("maximum")
	VsanPerfSummaryTypeVsanPerfSummaryType_Unknown = VsanPerfSummaryType("VsanPerfSummaryType_Unknown")
	VsanPerfSummaryTypeminimum                     = VsanPerfSummaryType("minimum")
	VsanPerfSummaryTypesummation                   = VsanPerfSummaryType("summation")
	VsanPerfSummaryTypelatest                      = VsanPerfSummaryType("latest")
)

func init() {
	types.Add("vsan:VsanPerfSummaryType", reflect.TypeOf((*VsanPerfSummaryType)(nil)).Elem())
}

type VsanPerfStatsUnitType string

const (
	VsanPerfStatsUnitTypesize_bytes                    = VsanPerfStatsUnitType("size_bytes")
	VsanPerfStatsUnitTypepermille                      = VsanPerfStatsUnitType("permille")
	VsanPerfStatsUnitTypetime_ms                       = VsanPerfStatsUnitType("time_ms")
	VsanPerfStatsUnitTypepercentage                    = VsanPerfStatsUnitType("percentage")
	VsanPerfStatsUnitTypetime_s                        = VsanPerfStatsUnitType("time_s")
	VsanPerfStatsUnitTyperate_bytes                    = VsanPerfStatsUnitType("rate_bytes")
	VsanPerfStatsUnitTypenumber                        = VsanPerfStatsUnitType("number")
	VsanPerfStatsUnitTypeVsanPerfStatsUnitType_Unknown = VsanPerfStatsUnitType("VsanPerfStatsUnitType_Unknown")
)

func init() {
	types.Add("vsan:VsanPerfStatsUnitType", reflect.TypeOf((*VsanPerfStatsUnitType)(nil)).Elem())
}

type VsanClusterHealthActionIdEnum string

const (
	VsanClusterHealthActionIdEnumVsanClusterHealthActionIdEnum_Unknown = VsanClusterHealthActionIdEnum("VsanClusterHealthActionIdEnum_Unknown")
	VsanClusterHealthActionIdEnumConfigureVSAN                         = VsanClusterHealthActionIdEnum("ConfigureVSAN")
	VsanClusterHealthActionIdEnumUploadHclDb                           = VsanClusterHealthActionIdEnum("UploadHclDb")
	VsanClusterHealthActionIdEnumRemediateDedup                        = VsanClusterHealthActionIdEnum("RemediateDedup")
	VsanClusterHealthActionIdEnumEnablePerformanceServiceAction        = VsanClusterHealthActionIdEnum("EnablePerformanceServiceAction")
	VsanClusterHealthActionIdEnumEnableCeip                            = VsanClusterHealthActionIdEnum("EnableCeip")
	VsanClusterHealthActionIdEnumLoginVumIsoDepot                      = VsanClusterHealthActionIdEnum("LoginVumIsoDepot")
	VsanClusterHealthActionIdEnumRelayoutVsanObjects                   = VsanClusterHealthActionIdEnum("RelayoutVsanObjects")
	VsanClusterHealthActionIdEnumRemediateFileService                  = VsanClusterHealthActionIdEnum("RemediateFileService")
	VsanClusterHealthActionIdEnumConfigureHA                           = VsanClusterHealthActionIdEnum("ConfigureHA")
	VsanClusterHealthActionIdEnumConfigureAutomaticRebalance           = VsanClusterHealthActionIdEnum("ConfigureAutomaticRebalance")
	VsanClusterHealthActionIdEnumCreateDVS                             = VsanClusterHealthActionIdEnum("CreateDVS")
	VsanClusterHealthActionIdEnumRemediateFileServiceImbalance         = VsanClusterHealthActionIdEnum("RemediateFileServiceImbalance")
	VsanClusterHealthActionIdEnumRunBurnInTest                         = VsanClusterHealthActionIdEnum("RunBurnInTest")
	VsanClusterHealthActionIdEnumUploadReleaseCatalog                  = VsanClusterHealthActionIdEnum("UploadReleaseCatalog")
	VsanClusterHealthActionIdEnumUpgradeVsanDiskFormat                 = VsanClusterHealthActionIdEnum("UpgradeVsanDiskFormat")
	VsanClusterHealthActionIdEnumEnableHealthService                   = VsanClusterHealthActionIdEnum("EnableHealthService")
	VsanClusterHealthActionIdEnumPurgeInaccessSwapObjs                 = VsanClusterHealthActionIdEnum("PurgeInaccessSwapObjs")
	VsanClusterHealthActionIdEnumDiskBalance                           = VsanClusterHealthActionIdEnum("DiskBalance")
	VsanClusterHealthActionIdEnumEnableIscsiTargetService              = VsanClusterHealthActionIdEnum("EnableIscsiTargetService")
	VsanClusterHealthActionIdEnumRepairClusterObjectsAction            = VsanClusterHealthActionIdEnum("RepairClusterObjectsAction")
	VsanClusterHealthActionIdEnumClaimVSANDisks                        = VsanClusterHealthActionIdEnum("ClaimVSANDisks")
	VsanClusterHealthActionIdEnumStopDiskBalance                       = VsanClusterHealthActionIdEnum("StopDiskBalance")
	VsanClusterHealthActionIdEnumConfigureDRS                          = VsanClusterHealthActionIdEnum("ConfigureDRS")
	VsanClusterHealthActionIdEnumClusterUpgrade                        = VsanClusterHealthActionIdEnum("ClusterUpgrade")
	VsanClusterHealthActionIdEnumCreateVMKnic                          = VsanClusterHealthActionIdEnum("CreateVMKnic")
	VsanClusterHealthActionIdEnumUpdateHclDbFromInternet               = VsanClusterHealthActionIdEnum("UpdateHclDbFromInternet")
	VsanClusterHealthActionIdEnumRemediateClusterConfig                = VsanClusterHealthActionIdEnum("RemediateClusterConfig")
	VsanClusterHealthActionIdEnumCreateVMKnicWithVMotion               = VsanClusterHealthActionIdEnum("CreateVMKnicWithVMotion")
)

func init() {
	types.Add("vsan:VsanClusterHealthActionIdEnum", reflect.TypeOf((*VsanClusterHealthActionIdEnum)(nil)).Elem())
}

type VsanSmartParameterType string

const (
	VsanSmartParameterTypesmartdrivetemperature          = VsanSmartParameterType("smartdrivetemperature")
	VsanSmartParameterTypeVsanSmartParameterType_Unknown = VsanSmartParameterType("VsanSmartParameterType_Unknown")
	VsanSmartParameterTypesmartinitialbadblockcount      = VsanSmartParameterType("smartinitialbadblockcount")
	VsanSmartParameterTypesmartdriveratedmaxtemperature  = VsanSmartParameterType("smartdriveratedmaxtemperature")
	VsanSmartParameterTypesmartmediawearoutindicator     = VsanSmartParameterType("smartmediawearoutindicator")
	VsanSmartParameterTypesmartwritesectorstotct         = VsanSmartParameterType("smartwritesectorstotct")
	VsanSmartParameterTypesmartreallocatedsectorct       = VsanSmartParameterType("smartreallocatedsectorct")
	VsanSmartParameterTypesmartreadsectorstotct          = VsanSmartParameterType("smartreadsectorstotct")
	VsanSmartParameterTypesmartpowercyclecount           = VsanSmartParameterType("smartpowercyclecount")
	VsanSmartParameterTypesmarthealthstatus              = VsanSmartParameterType("smarthealthstatus")
	VsanSmartParameterTypesmartpoweronhours              = VsanSmartParameterType("smartpoweronhours")
	VsanSmartParameterTypesmartwriteerrorcount           = VsanSmartParameterType("smartwriteerrorcount")
	VsanSmartParameterTypesmartrawreaderrorrate          = VsanSmartParameterType("smartrawreaderrorrate")
	VsanSmartParameterTypesmartreaderrorcount            = VsanSmartParameterType("smartreaderrorcount")
)

func init() {
	types.Add("vsan:VsanSmartParameterType", reflect.TypeOf((*VsanSmartParameterType)(nil)).Elem())
}

type VsanSyncStatus string

const (
	VsanSyncStatusactive                 = VsanSyncStatus("active")
	VsanSyncStatusVsanSyncStatus_Unknown = VsanSyncStatus("VsanSyncStatus_Unknown")
	VsanSyncStatusqueued                 = VsanSyncStatus("queued")
	VsanSyncStatussuspended              = VsanSyncStatus("suspended")
)

func init() {
	types.Add("vsan:VsanSyncStatus", reflect.TypeOf((*VsanSyncStatus)(nil)).Elem())
}

type VsanFileShareNfsSecType string

const (
	VsanFileShareNfsSecTypeSYS                         = VsanFileShareNfsSecType("SYS")
	VsanFileShareNfsSecTypeKRB5                        = VsanFileShareNfsSecType("KRB5")
	VsanFileShareNfsSecTypeFileShareNfsSecType_Unknown = VsanFileShareNfsSecType("FileShareNfsSecType_Unknown")
	VsanFileShareNfsSecTypeKRB5I                       = VsanFileShareNfsSecType("KRB5I")
	VsanFileShareNfsSecTypeKRB5P                       = VsanFileShareNfsSecType("KRB5P")
)

func init() {
	types.Add("vsan:VsanFileShareNfsSecType", reflect.TypeOf((*VsanFileShareNfsSecType)(nil)).Elem())
}

type VimVsanHostTrafficType string

const (
	VimVsanHostTrafficTypeTrafficType_Unknown = VimVsanHostTrafficType("TrafficType_Unknown")
	VimVsanHostTrafficTypevsan                = VimVsanHostTrafficType("vsan")
	VimVsanHostTrafficTypewitness             = VimVsanHostTrafficType("witness")
)

func init() {
	types.Add("vsan:VimVsanHostTrafficType", reflect.TypeOf((*VimVsanHostTrafficType)(nil)).Elem())
}

type VimClusterVsanDiskGroupCreationType string

const (
	VimClusterVsanDiskGroupCreationTypeallflash                          = VimClusterVsanDiskGroupCreationType("allFlash")
	VimClusterVsanDiskGroupCreationTypepmem                              = VimClusterVsanDiskGroupCreationType("pmem")
	VimClusterVsanDiskGroupCreationTypehybrid                            = VimClusterVsanDiskGroupCreationType("hybrid")
	VimClusterVsanDiskGroupCreationTypeVsanDiskGroupCreationType_Unknown = VimClusterVsanDiskGroupCreationType("VsanDiskGroupCreationType_Unknown")
	VimClusterVsanDiskGroupCreationTypevsandirect                        = VimClusterVsanDiskGroupCreationType("vsandirect")
)

func init() {
	types.Add("vsan:VimClusterVsanDiskGroupCreationType", reflect.TypeOf((*VimClusterVsanDiskGroupCreationType)(nil)).Elem())
}

type VsanHostQueryCheckLimitsOptionType string

const (
	VsanHostQueryCheckLimitsOptionTypelogicalCapacityUsed                        = VsanHostQueryCheckLimitsOptionType("logicalCapacityUsed")
	VsanHostQueryCheckLimitsOptionTypededupMetadata                              = VsanHostQueryCheckLimitsOptionType("dedupMetadata")
	VsanHostQueryCheckLimitsOptionTypeVsanHostQueryCheckLimitsOptionType_Unknown = VsanHostQueryCheckLimitsOptionType("VsanHostQueryCheckLimitsOptionType_Unknown")
	VsanHostQueryCheckLimitsOptionTypelogicalCapacity                            = VsanHostQueryCheckLimitsOptionType("logicalCapacity")
	VsanHostQueryCheckLimitsOptionTypedgTransientCapacityUsed                    = VsanHostQueryCheckLimitsOptionType("dgTransientCapacityUsed")
	VsanHostQueryCheckLimitsOptionTypediskTransientCapacityUsed                  = VsanHostQueryCheckLimitsOptionType("diskTransientCapacityUsed")
)

func init() {
	types.Add("vsan:VsanHostQueryCheckLimitsOptionType", reflect.TypeOf((*VsanHostQueryCheckLimitsOptionType)(nil)).Elem())
}

type VsanHostWipeDiskState string

const (
	VsanHostWipeDiskStateFailure               = VsanHostWipeDiskState("Failure")
	VsanHostWipeDiskStateWiping                = VsanHostWipeDiskState("Wiping")
	VsanHostWipeDiskStateWipeDiskState_Unknown = VsanHostWipeDiskState("WipeDiskState_Unknown")
	VsanHostWipeDiskStateSuccess               = VsanHostWipeDiskState("Success")
)

func init() {
	types.Add("vsan:VsanHostWipeDiskState", reflect.TypeOf((*VsanHostWipeDiskState)(nil)).Elem())
}

type VsanHostStatsType string

const (
	VsanHostStatsTypeconfigGeneration         = VsanHostStatsType("configGeneration")
	VsanHostStatsTyperesyncIopsInfo           = VsanHostStatsType("resyncIopsInfo")
	VsanHostStatsTypecomponentLimitPerCluster = VsanHostStatsType("componentLimitPerCluster")
	VsanHostStatsTypesupportedClusterSize     = VsanHostStatsType("supportedClusterSize")
	VsanHostStatsTyperepairTimerInfo          = VsanHostStatsType("repairTimerInfo")
	VsanHostStatsTypemaxWitnessClusters       = VsanHostStatsType("maxWitnessClusters")
	VsanHostStatsTypeStatsType_Unknown        = VsanHostStatsType("StatsType_Unknown")
)

func init() {
	types.Add("vsan:VsanHostStatsType", reflect.TypeOf((*VsanHostStatsType)(nil)).Elem())
}

type VimVsanVsanVcsaDeploymentPhase string

const (
	VimVsanVsanVcsaDeploymentPhasefailed                          = VimVsanVsanVcsaDeploymentPhase("failed")
	VimVsanVsanVcsaDeploymentPhasevcsadeploy                      = VimVsanVsanVcsaDeploymentPhase("vcsadeploy")
	VimVsanVsanVcsaDeploymentPhaseovaunpack                       = VimVsanVsanVcsaDeploymentPhase("ovaunpack")
	VimVsanVsanVcsaDeploymentPhasedone                            = VimVsanVsanVcsaDeploymentPhase("done")
	VimVsanVsanVcsaDeploymentPhaseVsanVcsaDeploymentPhase_Unknown = VimVsanVsanVcsaDeploymentPhase("VsanVcsaDeploymentPhase_Unknown")
	VimVsanVsanVcsaDeploymentPhaseinitializing                    = VimVsanVsanVcsaDeploymentPhase("initializing")
	VimVsanVsanVcsaDeploymentPhasevalidation                      = VimVsanVsanVcsaDeploymentPhase("validation")
	VimVsanVsanVcsaDeploymentPhasevcconfig                        = VimVsanVsanVcsaDeploymentPhase("vcconfig")
	VimVsanVsanVcsaDeploymentPhasevsanbootstrap                   = VimVsanVsanVcsaDeploymentPhase("vsanbootstrap")
)

func init() {
	types.Add("vsan:VimVsanVsanVcsaDeploymentPhase", reflect.TypeOf((*VimVsanVsanVcsaDeploymentPhase)(nil)).Elem())
}

type VsanCapabilityStatus string

const (
	VsanCapabilityStatusunknown      = VsanCapabilityStatus("unknown")
	VsanCapabilityStatuscalculated   = VsanCapabilityStatus("calculated")
	VsanCapabilityStatusdisconnected = VsanCapabilityStatus("disconnected")
	VsanCapabilityStatusoldversion   = VsanCapabilityStatus("oldversion")
)

func init() {
	types.Add("vsan:VsanCapabilityStatus", reflect.TypeOf((*VsanCapabilityStatus)(nil)).Elem())
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

type VsanIscsiTargetAuthType string

const (
	VsanIscsiTargetAuthTypeCHAP                            = VsanIscsiTargetAuthType("CHAP")
	VsanIscsiTargetAuthTypeNoAuth                          = VsanIscsiTargetAuthType("NoAuth")
	VsanIscsiTargetAuthTypeCHAP_Mutual                     = VsanIscsiTargetAuthType("CHAP_Mutual")
	VsanIscsiTargetAuthTypeVsanIscsiTargetAuthType_Unknown = VsanIscsiTargetAuthType("VsanIscsiTargetAuthType_Unknown")
)

func init() {
	types.Add("vsan:VsanIscsiTargetAuthType", reflect.TypeOf((*VsanIscsiTargetAuthType)(nil)).Elem())
}

type VsanIoInsightState string

const (
	VsanIoInsightStatenotFound                   = VsanIoInsightState("notFound")
	VsanIoInsightStaterunning                    = VsanIoInsightState("running")
	VsanIoInsightStatestopped                    = VsanIoInsightState("stopped")
	VsanIoInsightStateVsanIoInsightState_unknown = VsanIoInsightState("VsanIoInsightState_unknown")
)

func init() {
	types.Add("vsan:VsanIoInsightState", reflect.TypeOf((*VsanIoInsightState)(nil)).Elem())
}

type VsanPerfThresholdDirectionType string

const (
	VsanPerfThresholdDirectionTypeupper                                  = VsanPerfThresholdDirectionType("upper")
	VsanPerfThresholdDirectionTypelower                                  = VsanPerfThresholdDirectionType("lower")
	VsanPerfThresholdDirectionTypeVsanPerfThresholdDirectionType_Unknown = VsanPerfThresholdDirectionType("VsanPerfThresholdDirectionType_Unknown")
)

func init() {
	types.Add("vsan:VsanPerfThresholdDirectionType", reflect.TypeOf((*VsanPerfThresholdDirectionType)(nil)).Elem())
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
