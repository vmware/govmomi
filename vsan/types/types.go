/*
Copyright (c) 2014-2018 VMware, Inc. All Rights Reserved.

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

type VsanPerfGetSupportedEntityTypes VsanPerfGetSupportedEntityTypesRequestType

func init() {
	t["VsanPerfGetSupportedEntityTypes"] = reflect.TypeOf((*VsanPerfGetSupportedEntityTypes)(nil)).Elem()
}

type VsanPerfGetSupportedEntityTypesRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

func init() {
	t["VsanPerfGetSupportedEntityTypesRequestType"] = reflect.TypeOf((*VsanPerfGetSupportedEntityTypesRequestType)(nil)).Elem()
}

type VsanPerfGetSupportedEntityTypesResponse struct {
	Returnval []VsanPerfEntityType `xml:"returnval,omitempty"`
}

type VsanPerfEntityType struct {
	DynamicData

	Name        string          `xml:"name"`
	Id          string          `xml:"id"`
	Graphs      []VsanPerfGraph `xml:"graphs"`
	Description string          `xml:"description,omitempty"`
}

func init() {
	t["VsanPerfEntityType"] = reflect.TypeOf((*VsanPerfEntityType)(nil)).Elem()
}

type VsanPerfGraph struct {
	DynamicData

	Id          string             `xml:"id"`
	Metrics     []VsanPerfMetricId `xml:"metrics"`
	Unit        string             `xml:"unit"`
	Threshold   *VsanPerfThreshold `xml:"threshold,omitempty"`
	Name        string             `xml:"name,omitempty"`
	Description string             `xml:"description,omitempty"`
}

func init() {
	t["VsanPerfGraph"] = reflect.TypeOf((*VsanPerfGraph)(nil)).Elem()
}

type VsanPerfThreshold struct {
	DynamicData

	Direction string `xml:"direction"`
	Yellow    string `xml:"yellow,omitempty"`
	Red       string `xml:"red,omitempty"`
}

func init() {
	t["VsanPerfThreshold"] = reflect.TypeOf((*VsanPerfThreshold)(nil)).Elem()
}

// Cluster health summary
type VsanQueryVcClusterHealthSummary VsanQueryVcClusterHealthSummaryRequestType

func init() {
	t["VsanQueryVcClusterHealthSummary"] = reflect.TypeOf((*VsanQueryVcClusterHealthSummary)(nil)).Elem()
}

type VsanQueryVcClusterHealthSummaryRequestType struct {
	This            types.ManagedObjectReference `xml:"_this"`
	Cluster         types.ManagedObjectReference `xml:"cluster"`
	VmCreateTimeout int32                        `xml:"vmCreateTimeout,omitempty"`
	ObjUuids        []string                     `xml:"objUuids,omitempty"`
	IncludeObjUuids *bool                        `xml:"includeObjUuids"`
	Fields          []string                     `xml:"fields,omitempty"`
	FetchFromCache  *bool                        `xml:"fetchFromCache"`
	Perspective     string                       `xml:"perspective,omitempty"`
}

func init() {
	t["VsanQueryVcClusterHealthSummaryRequestType"] = reflect.TypeOf((*VsanQueryVcClusterHealthSummaryRequestType)(nil)).Elem()
}

type VsanQueryVcClusterHealthSummaryResponse struct {
	Returnval VsanClusterHealthSummary `xml:"returnval"`
}

type VsanClusterHealthSummary struct {
	DynamicData

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
}

func init() {
	t["VsanClusterHealthSummary"] = reflect.TypeOf((*VsanClusterHealthSummary)(nil)).Elem()
}

type VsanClusterHealthSystemStatusResult struct {
	DynamicData

	Status             string                             `xml:"status"`
	GoalState          string                             `xml:"goalState"`
	UntrackedHosts     []string                           `xml:"untrackedHosts,omitempty"`
	TrackedHostsStatus []VsanHostHealthSystemStatusResult `xml:"trackedHostsStatus,omitempty"`
}

func init() {
	t["VsanClusterHealthSystemStatusResult"] = reflect.TypeOf((*VsanClusterHealthSystemStatusResult)(nil)).Elem()
}

type VsanHostHealthSystemStatusResult struct {
	DynamicData

	Hostname string   `xml:"hostname"`
	Status   string   `xml:"status"`
	Issues   []string `xml:"issues,omitempty"`
}

func init() {
	t["VsanHostHealthSystemStatusResult"] = reflect.TypeOf((*VsanHostHealthSystemStatusResult)(nil)).Elem()
}

type VsanClusterHealthSystemVersionResult struct {
	DynamicData

	HostResults     []VsanHostHealthSystemVersionResult `xml:"hostResults,omitempty"`
	VcVersion       string                              `xml:"vcVersion,omitempty"`
	IssueFound      bool                                `xml:"issueFound"`
	UpgradePossible *bool                               `xml:"upgradePossible"`
}

func init() {
	t["VsanClusterHealthSystemVersionResult"] = reflect.TypeOf((*VsanClusterHealthSystemVersionResult)(nil)).Elem()
}

type VsanHostHealthSystemVersionResult struct {
	DynamicData

	Hostname string                `xml:"hostname"`
	Version  string                `xml:"version,omitempty"`
	Error    *LocalizedMethodFault `xml:"error,omitempty"`
}

func init() {
	t["VsanHostHealthSystemVersionResult"] = reflect.TypeOf((*VsanHostHealthSystemVersionResult)(nil)).Elem()
}

type LocalizedMethodFault struct {
	DynamicData

	Fault            BaseMethodFault `xml:"fault,typeattr"`
	LocalizedMessage string          `xml:"localizedMessage,omitempty"`
}

func init() {
	t["LocalizedMethodFault"] = reflect.TypeOf((*LocalizedMethodFault)(nil)).Elem()
}

type VsanObjectOverallHealth struct {
	DynamicData

	ObjectHealthDetail      []VsanObjectHealth `xml:"objectHealthDetail,omitempty"`
	ObjectVersionCompliance *bool              `xml:"objectVersionCompliance"`
}

func init() {
	t["VsanObjectOverallHealth"] = reflect.TypeOf((*VsanObjectOverallHealth)(nil)).Elem()
}

type VsanObjectHealth struct {
	DynamicData

	NumObjects int32    `xml:"numObjects"`
	Health     string   `xml:"health"`
	ObjUuids   []string `xml:"objUuids,omitempty"`
}

func init() {
	t["VsanObjectHealth"] = reflect.TypeOf((*VsanObjectHealth)(nil)).Elem()
}

type VsanClusterVMsHealthOverallResult struct {
	DynamicData

	HealthStateList    []VsanClusterVMsHealthSummaryResult `xml:"healthStateList,omitempty"`
	OverallHealthState string                              `xml:"overallHealthState,omitempty"`
}

func init() {
	t["VsanClusterVMsHealthOverallResult"] = reflect.TypeOf((*VsanClusterVMsHealthOverallResult)(nil)).Elem()
}

type VsanClusterVMsHealthSummaryResult struct {
	DynamicData

	NumVMs          int32    `xml:"numVMs"`
	State           string   `xml:"state,omitempty"`
	Health          string   `xml:"health"`
	VmInstanceUuids []string `xml:"vmInstanceUuids,omitempty"`
}

func init() {
	t["VsanClusterVMsHealthSummaryResult"] = reflect.TypeOf((*VsanClusterVMsHealthSummaryResult)(nil)).Elem()
}

type VsanClusterNetworkHealthResult struct {
	DynamicData

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
	t["VsanClusterNetworkHealthResult"] = reflect.TypeOf((*VsanClusterNetworkHealthResult)(nil)).Elem()
}

type VsanNetworkHealthResult struct {
	DynamicData

	Host              *types.ManagedObjectReference `xml:"host,omitempty"`
	Hostname          string                        `xml:"hostname,omitempty"`
	VsanVmknicPresent *bool                         `xml:"vsanVmknicPresent"`
	IpSubnets         []string                      `xml:"ipSubnets,omitempty"`
	IssueFound        *bool                         `xml:"issueFound"`
	PeerHealth        []VsanNetworkPeerHealthResult `xml:"peerHealth,omitempty"`
	VMotionHealth     []VsanNetworkPeerHealthResult `xml:"vMotionHealth,omitempty"`
	MulticastConfig   string                        `xml:"multicastConfig,omitempty"`
	InUnicast         *bool                         `xml:"inUnicast"`
}

func init() {
	t["VsanNetworkHealthResult"] = reflect.TypeOf((*VsanNetworkHealthResult)(nil)).Elem()
}

type VsanClusterNetworkPartitionInfo struct {
	DynamicData

	Hosts []string `xml:"hosts,omitempty"`
}

func init() {
	t["VsanClusterNetworkPartitionInfo"] = reflect.TypeOf((*VsanClusterNetworkPartitionInfo)(nil)).Elem()
}

type VsanQueryResultHostInfo struct {
	DynamicData

	Uuid              string   `xml:"uuid,omitempty"`
	HostnameInCmmds   string   `xml:"hostnameInCmmds,omitempty"`
	VsanIpv4Addresses []string `xml:"vsanIpv4Addresses,omitempty"`
}

func init() {
	t["VsanQueryResultHostInfo"] = reflect.TypeOf((*VsanQueryResultHostInfo)(nil)).Elem()
}

type VsanNetworkPeerHealthResult struct {
	DynamicData

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
	t["VsanNetworkPeerHealthResult"] = reflect.TypeOf((*VsanNetworkPeerHealthResult)(nil)).Elem()
}

type VsanClusterLimitHealthResult struct {
	DynamicData

	IssueFound              bool                                  `xml:"issueFound"`
	ComponentLimitHealth    string                                `xml:"componentLimitHealth"`
	DiskFreeSpaceHealth     string                                `xml:"diskFreeSpaceHealth"`
	RcFreeReservationHealth string                                `xml:"rcFreeReservationHealth"`
	HostResults             []VsanLimitHealthResult               `xml:"hostResults,omitempty"`
	WhatifHostFailures      []VsanClusterWhatifHostFailuresResult `xml:"whatifHostFailures,omitempty"`
	HostsCommFailure        []string                              `xml:"hostsCommFailure,omitempty"`
}

func init() {
	t["VsanClusterLimitHealthResult"] = reflect.TypeOf((*VsanClusterLimitHealthResult)(nil)).Elem()
}

type VsanLimitHealthResult struct {
	DynamicData

	Hostname                string `xml:"hostname,omitempty"`
	IssueFound              bool   `xml:"issueFound"`
	MaxComponents           int32  `xml:"maxComponents"`
	FreeComponents          int32  `xml:"freeComponents"`
	ComponentLimitHealth    string `xml:"componentLimitHealth"`
	LowestFreeDiskSpacePct  int32  `xml:"lowestFreeDiskSpacePct"`
	UsedDiskSpaceB          int64  `xml:"usedDiskSpaceB"`
	TotalDiskSpaceB         int64  `xml:"totalDiskSpaceB"`
	DiskFreeSpaceHealth     string `xml:"diskFreeSpaceHealth"`
	ReservedRcSizeB         int64  `xml:"reservedRcSizeB"`
	TotalRcSizeB            int64  `xml:"totalRcSizeB"`
	RcFreeReservationHealth string `xml:"rcFreeReservationHealth"`
}

func init() {
	t["VsanLimitHealthResult"] = reflect.TypeOf((*VsanLimitHealthResult)(nil)).Elem()
}

type VsanClusterWhatifHostFailuresResult struct {
	DynamicData

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
	t["VsanClusterWhatifHostFailuresResult"] = reflect.TypeOf((*VsanClusterWhatifHostFailuresResult)(nil)).Elem()
}

type VsanClusterAdvCfgSyncResult struct {
	DynamicData

	InSync     bool                              `xml:"inSync"`
	Name       string                            `xml:"name"`
	HostValues []VsanClusterAdvCfgSyncHostResult `xml:"hostValues,omitempty"`
}

func init() {
	t["VsanClusterAdvCfgSyncResult"] = reflect.TypeOf((*VsanClusterAdvCfgSyncResult)(nil)).Elem()
}

type VsanClusterAdvCfgSyncHostResult struct {
	DynamicData

	Hostname string `xml:"hostname"`
	Value    string `xml:"value"`
}

func init() {
	t["VsanClusterAdvCfgSyncHostResult"] = reflect.TypeOf((*VsanClusterAdvCfgSyncHostResult)(nil)).Elem()
}

type VsanHostCreateVmHealthTestResult struct {
	DynamicData

	Hostname string                `xml:"hostname"`
	State    string                `xml:"state"`
	Fault    *LocalizedMethodFault `xml:"fault,omitempty"`
}

func init() {
	t["VsanHostCreateVmHealthTestResult"] = reflect.TypeOf((*VsanHostCreateVmHealthTestResult)(nil)).Elem()
}

type VsanPhysicalDiskHealthSummary struct {
	DynamicData

	OverallHealth        string                   `xml:"overallHealth"`
	HeapsWithIssues      []VsanResourceHealth     `xml:"heapsWithIssues,omitempty"`
	SlabsWithIssues      []VsanResourceHealth     `xml:"slabsWithIssues,omitempty"`
	Disks                []VsanPhysicalDiskHealth `xml:"disks,omitempty"`
	ComponentsWithIssues []VsanResourceHealth     `xml:"componentsWithIssues,omitempty"`
	Hostname             string                   `xml:"hostname,omitempty"`
	HostDedupScope       int32                    `xml:"hostDedupScope,omitempty"`
	Error                *LocalizedMethodFault    `xml:"error,omitempty"`
}

func init() {
	t["VsanPhysicalDiskHealthSummary"] = reflect.TypeOf((*VsanPhysicalDiskHealthSummary)(nil)).Elem()
}

type VsanResourceHealth struct {
	DynamicData

	Resource    string `xml:"resource"`
	Health      string `xml:"health"`
	Description string `xml:"description,omitempty"`
}

func init() {
	t["VsanResourceHealth"] = reflect.TypeOf((*VsanResourceHealth)(nil)).Elem()
}

type VsanPhysicalDiskHealth struct {
	DynamicData

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
	ScsiDisk                     *HostScsiDisk            `xml:"scsiDisk,omitempty"`
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
	t["VsanPhysicalDiskHealth"] = reflect.TypeOf((*VsanPhysicalDiskHealth)(nil)).Elem()
}

type HostScsiDisk struct {
	ScsiLun

	Capacity              HostDiskDimensionsLba `xml:"capacity"`
	DevicePath            string                `xml:"devicePath"`
	Ssd                   *bool                 `xml:"ssd"`
	LocalDisk             *bool                 `xml:"localDisk"`
	PhysicalLocation      []string              `xml:"physicalLocation,omitempty"`
	EmulatedDIXDIFEnabled *bool                 `xml:"emulatedDIXDIFEnabled"`
	VsanDiskInfo          *VsanHostVsanDiskInfo `xml:"vsanDiskInfo,omitempty"`
	ScsiDiskType          string                `xml:"scsiDiskType,omitempty"`
}

func init() {
	t["HostScsiDisk"] = reflect.TypeOf((*HostScsiDisk)(nil)).Elem()
}

type ScsiLun struct {
	HostDevice

	Key              string               `xml:"key,omitempty"`
	Uuid             string               `xml:"uuid"`
	Descriptor       []ScsiLunDescriptor  `xml:"descriptor,omitempty"`
	CanonicalName    string               `xml:"canonicalName,omitempty"`
	DisplayName      string               `xml:"displayName,omitempty"`
	LunType          string               `xml:"lunType"`
	Vendor           string               `xml:"vendor,omitempty"`
	Model            string               `xml:"model,omitempty"`
	Revision         string               `xml:"revision,omitempty"`
	ScsiLevel        int32                `xml:"scsiLevel,omitempty"`
	SerialNumber     string               `xml:"serialNumber,omitempty"`
	DurableName      *ScsiLunDurableName  `xml:"durableName,omitempty"`
	AlternateName    []ScsiLunDurableName `xml:"alternateName,omitempty"`
	StandardInquiry  []byte               `xml:"standardInquiry,omitempty"`
	QueueDepth       int32                `xml:"queueDepth,omitempty"`
	OperationalState []string             `xml:"operationalState"`
	Capabilities     *ScsiLunCapabilities `xml:"capabilities,omitempty"`
	VStorageSupport  string               `xml:"vStorageSupport,omitempty"`
	ProtocolEndpoint *bool                `xml:"protocolEndpoint"`
}

func init() {
	t["ScsiLun"] = reflect.TypeOf((*ScsiLun)(nil)).Elem()
}

type HostDiskDimensionsLba struct {
	DynamicData

	BlockSize int32 `xml:"blockSize"`
	Block     int64 `xml:"block"`
}

func init() {
	t["HostDiskDimensionsLba"] = reflect.TypeOf((*HostDiskDimensionsLba)(nil)).Elem()
}

type HostDevice struct {
	DynamicData

	DeviceName string `xml:"deviceName"`
	DeviceType string `xml:"deviceType"`
}

func init() {
	t["HostDevice"] = reflect.TypeOf((*HostDevice)(nil)).Elem()
}

type ScsiLunDescriptor struct {
	DynamicData

	Quality string `xml:"quality"`
	Id      string `xml:"id"`
}

func init() {
	t["ScsiLunDescriptor"] = reflect.TypeOf((*ScsiLunDescriptor)(nil)).Elem()
}

type ScsiLunDurableName struct {
	DynamicData

	Namespace   string `xml:"namespace"`
	NamespaceId byte   `xml:"namespaceId"`
	Data        []byte `xml:"data,omitempty"`
}

func init() {
	t["ScsiLunDurableName"] = reflect.TypeOf((*ScsiLunDurableName)(nil)).Elem()
}

type ScsiLunCapabilities struct {
	DynamicData

	UpdateDisplayNameSupported bool `xml:"updateDisplayNameSupported"`
}

func init() {
	t["ScsiLunCapabilities"] = reflect.TypeOf((*ScsiLunCapabilities)(nil)).Elem()
}

type VsanHostVsanDiskInfo struct {
	DynamicData

	VsanUuid      string `xml:"vsanUuid"`
	FormatVersion int32  `xml:"formatVersion"`
}

func init() {
	t["VsanHostVsanDiskInfo"] = reflect.TypeOf((*VsanHostVsanDiskInfo)(nil)).Elem()
}

type VsanDiskRebalanceResult struct {
	DynamicData

	Status               string  `xml:"status"`
	BytesMoving          int64   `xml:"bytesMoving,omitempty"`
	RemainingBytesToMove int64   `xml:"remainingBytesToMove,omitempty"`
	DiskUsage            float32 `xml:"diskUsage,omitempty"`
	MaxDiskUsage         float32 `xml:"maxDiskUsage,omitempty"`
	MinDiskUsage         float32 `xml:"minDiskUsage,omitempty"`
	AvgDiskUsage         float32 `xml:"avgDiskUsage,omitempty"`
}

func init() {
	t["VsanDiskRebalanceResult"] = reflect.TypeOf((*VsanDiskRebalanceResult)(nil)).Elem()
}

type VsanClusterEncryptionHealthSummary struct {
	DynamicData

	OverallHealth string                        `xml:"overallHealth,omitempty"`
	ConfigHealth  string                        `xml:"configHealth,omitempty"`
	KmsHealth     string                        `xml:"kmsHealth,omitempty"`
	VcKmsResult   *VsanVcKmipServersHealth      `xml:"vcKmsResult,omitempty"`
	HostResults   []VsanEncryptionHealthSummary `xml:"hostResults,omitempty"`
}

func init() {
	t["VsanClusterEncryptionHealthSummary"] = reflect.TypeOf((*VsanClusterEncryptionHealthSummary)(nil)).Elem()
}

type VsanVcKmipServersHealth struct {
	DynamicData

	Health               string                `xml:"health,omitempty"`
	Error                *LocalizedMethodFault `xml:"error,omitempty"`
	KmsProviderId        string                `xml:"kmsProviderId,omitempty"`
	KmsHealth            []VsanKmsHealth       `xml:"kmsHealth,omitempty"`
	ClientCertHealth     string                `xml:"clientCertHealth,omitempty"`
	ClientCertExpireDate *time.Time            `xml:"clientCertExpireDate"`
}

func init() {
	t["VsanVcKmipServersHealth"] = reflect.TypeOf((*VsanVcKmipServersHealth)(nil)).Elem()
}

type VsanKmsHealth struct {
	DynamicData

	ServerName     string                `xml:"serverName"`
	Health         string                `xml:"health"`
	Error          *LocalizedMethodFault `xml:"error,omitempty"`
	TrustHealth    string                `xml:"trustHealth,omitempty"`
	CertHealth     string                `xml:"certHealth,omitempty"`
	CertExpireDate *time.Time            `xml:"certExpireDate"`
}

func init() {
	t["VsanKmsHealth"] = reflect.TypeOf((*VsanKmsHealth)(nil)).Elem()
}

type VsanEncryptionHealthSummary struct {
	DynamicData

	Hostname         string                     `xml:"hostname,omitempty"`
	EncryptionInfo   *VsanHostEncryptionInfo    `xml:"encryptionInfo,omitempty"`
	OverallKmsHealth string                     `xml:"overallKmsHealth"`
	KmsHealth        []VsanKmsHealth            `xml:"kmsHealth,omitempty"`
	EncryptionIssues []string                   `xml:"encryptionIssues,omitempty"`
	DiskResults      []VsanDiskEncryptionHealth `xml:"diskResults,omitempty"`
	Error            *LocalizedMethodFault      `xml:"error,omitempty"`
	AesniEnabled     *bool                      `xml:"aesniEnabled"`
}

func init() {
	t["VsanEncryptionHealthSummary"] = reflect.TypeOf((*VsanEncryptionHealthSummary)(nil)).Elem()
}

type VsanHostEncryptionInfo struct {
	DynamicData

	Enabled             *bool            `xml:"enabled"`
	KekId               string           `xml:"kekId,omitempty"`
	HostKeyId           string           `xml:"hostKeyId,omitempty"`
	KmipServers         []KmipServerSpec `xml:"kmipServers,omitempty"`
	KmsServerCerts      []string         `xml:"kmsServerCerts,omitempty"`
	ClientKey           string           `xml:"clientKey,omitempty"`
	ClientCert          string           `xml:"clientCert,omitempty"`
	DekGenerationId     int64            `xml:"dekGenerationId,omitempty"`
	Changing            *bool            `xml:"changing"`
	EraseDisksBeforeUse *bool            `xml:"eraseDisksBeforeUse"`
}

func init() {
	t["VsanHostEncryptionInfo"] = reflect.TypeOf((*VsanHostEncryptionInfo)(nil)).Elem()
}

type VsanDiskEncryptionHealth struct {
	DynamicData

	DiskHealth       *VsanPhysicalDiskHealth `xml:"diskHealth,omitempty"`
	EncryptionIssues []string                `xml:"encryptionIssues,omitempty"`
}

func init() {
	t["VsanDiskEncryptionHealth"] = reflect.TypeOf((*VsanDiskEncryptionHealth)(nil)).Elem()
}

type KmipServerSpec struct {
	DynamicData

	ClusterId KeyProviderId  `xml:"clusterId"`
	Info      KmipServerInfo `xml:"info"`
	Password  string         `xml:"password,omitempty"`
}

func init() {
	t["KmipServerSpec"] = reflect.TypeOf((*KmipServerSpec)(nil)).Elem()
}

type KeyProviderId struct {
	DynamicData

	Id string `xml:"id"`
}

func init() {
	t["KeyProviderId"] = reflect.TypeOf((*KeyProviderId)(nil)).Elem()
}

type KmipServerInfo struct {
	DynamicData

	Name         string `xml:"name"`
	Address      string `xml:"address"`
	Port         int32  `xml:"port"`
	ProxyAddress string `xml:"proxyAddress,omitempty"`
	ProxyPort    int32  `xml:"proxyPort,omitempty"`
	Reconnect    int32  `xml:"reconnect,omitempty"`
	Protocol     string `xml:"protocol,omitempty"`
	Nbio         int32  `xml:"nbio,omitempty"`
	Timeout      int32  `xml:"timeout,omitempty"`
	UserName     string `xml:"userName,omitempty"`
}

func init() {
	t["KmipServerInfo"] = reflect.TypeOf((*KmipServerInfo)(nil)).Elem()
}

type VsanClusterHclInfo struct {
	DynamicData

	HclDbLastUpdate *time.Time        `xml:"hclDbLastUpdate"`
	HclDbAgeHealth  string            `xml:"hclDbAgeHealth,omitempty"`
	HostResults     []VsanHostHclInfo `xml:"hostResults,omitempty"`
	UpdateItems     []VsanUpdateItem  `xml:"updateItems,omitempty"`
}

func init() {
	t["VsanClusterHclInfo"] = reflect.TypeOf((*VsanClusterHclInfo)(nil)).Elem()
}

type VsanHostHclInfo struct {
	DynamicData

	Hostname    string                  `xml:"hostname"`
	HclChecked  bool                    `xml:"hclChecked"`
	ReleaseName string                  `xml:"releaseName,omitempty"`
	Error       *LocalizedMethodFault   `xml:"error,omitempty"`
	Controllers []VsanHclControllerInfo `xml:"controllers,omitempty"`
}

func init() {
	t["VsanHostHclInfo"] = reflect.TypeOf((*VsanHostHclInfo)(nil)).Elem()
}

type VsanHclControllerInfo struct {
	DynamicData

	DeviceName             string                   `xml:"deviceName"`
	DeviceDisplayName      string                   `xml:"deviceDisplayName,omitempty"`
	DriverName             string                   `xml:"driverName,omitempty"`
	DriverVersion          string                   `xml:"driverVersion,omitempty"`
	VendorId               int64                    `xml:"vendorId,omitempty"`
	DeviceId               int64                    `xml:"deviceId,omitempty"`
	SubVendorId            int64                    `xml:"subVendorId,omitempty"`
	SubDeviceId            int64                    `xml:"subDeviceId,omitempty"`
	ExtraInfo              []KeyValue               `xml:"extraInfo,omitempty"`
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
	DiskMode               *ChoiceOption            `xml:"diskMode,omitempty"`
	DiskModeOnHcl          []string                 `xml:"diskModeOnHcl,omitempty"`
	DiskModeSupported      *bool                    `xml:"diskModeSupported"`
}

func init() {
	t["VsanHclControllerInfo"] = reflect.TypeOf((*VsanHclControllerInfo)(nil)).Elem()
}

type KeyValue struct {
	DynamicData

	Key   string `xml:"key"`
	Value string `xml:"value"`
}

func init() {
	t["KeyValue"] = reflect.TypeOf((*KeyValue)(nil)).Elem()
}

type VsanUpdateItem struct {
	DynamicData

	Host            types.ManagedObjectReference `xml:"host"`
	Type            string                       `xml:"type"`
	Name            string                       `xml:"name"`
	Version         string                       `xml:"version"`
	ExistingVersion string                       `xml:"existingVersion,omitempty"`
	Present         bool                         `xml:"present"`
	VibSpec         []VsanVibSpec                `xml:"vibSpec,omitempty"`
	FirmwareSpec    *VsanHclFirmwareUpdateSpec   `xml:"firmwareSpec,omitempty"`
	DownloadInfo    []VsanDownloadItem           `xml:"downloadInfo,omitempty"`
	Eula            string                       `xml:"eula,omitempty"`
	Adapter         string                       `xml:"adapter,omitempty"`
}

func init() {
	t["VsanUpdateItem"] = reflect.TypeOf((*VsanUpdateItem)(nil)).Elem()
}

type VsanHclFirmwareUpdateSpec struct {
	DynamicData

	Host           types.ManagedObjectReference `xml:"host"`
	HbaDevice      string                       `xml:"hbaDevice"`
	FwFiles        []VsanHclFirmwareFile        `xml:"fwFiles"`
	AllowDowngrade *bool                        `xml:"allowDowngrade"`
}

func init() {
	t["VsanHclFirmwareUpdateSpec"] = reflect.TypeOf((*VsanHclFirmwareUpdateSpec)(nil)).Elem()
}

type VsanHclFirmwareFile struct {
	DynamicData

	FileType      string `xml:"fileType"`
	FilenameOrUrl string `xml:"filenameOrUrl"`
	Sha1sum       string `xml:"sha1sum"`
}

func init() {
	t["VsanHclFirmwareFile"] = reflect.TypeOf((*VsanHclFirmwareFile)(nil)).Elem()
}

type VsanClusterHealthGroup struct {
	DynamicData

	GroupId      string                            `xml:"groupId"`
	GroupName    string                            `xml:"groupName"`
	GroupHealth  string                            `xml:"groupHealth"`
	GroupTests   []VsanClusterHealthTest           `xml:"groupTests,omitempty"`
	GroupDetails []BaseVsanClusterHealthResultBase `xml:"groupDetails,omitempty,typeattr"`
}

func init() {
	t["VsanClusterHealthGroup"] = reflect.TypeOf((*VsanClusterHealthGroup)(nil)).Elem()
}

type VsanHostCimProviderInfo struct {
	DynamicData

	CimProviderSupported  *bool              `xml:"cimProviderSupported"`
	InstalledCIMProvider  string             `xml:"installedCIMProvider,omitempty"`
	CimProviderOnHcl      []string           `xml:"cimProviderOnHcl,omitempty"`
	CimProviderLinksOnHcl []VsanDownloadItem `xml:"cimProviderLinksOnHcl,omitempty"`
}

func init() {
	t["VsanHostCimProviderInfo"] = reflect.TypeOf((*VsanHostCimProviderInfo)(nil)).Elem()
}

type VsanDownloadItem struct {
	DynamicData

	Url        string `xml:"url"`
	Sha1sum    string `xml:"sha1sum"`
	FormatType string `xml:"formatType,omitempty"`
}

func init() {
	t["VsanDownloadItem"] = reflect.TypeOf((*VsanDownloadItem)(nil)).Elem()
}

type VsanClusterHealthTest struct {
	DynamicData

	TestId               string                            `xml:"testId,omitempty"`
	TestName             string                            `xml:"testName,omitempty"`
	TestDescription      string                            `xml:"testDescription,omitempty"`
	TestShortDescription string                            `xml:"testShortDescription,omitempty"`
	TestHealth           string                            `xml:"testHealth,omitempty"`
	TestDetails          []BaseVsanClusterHealthResultBase `xml:"testDetails,omitempty,typeattr"`
	TestActions          []VsanClusterHealthAction         `xml:"testActions,omitempty"`
}

func init() {
	t["VsanClusterHealthTest"] = reflect.TypeOf((*VsanClusterHealthTest)(nil)).Elem()
}

type VsanVibSpec struct {
	DynamicData

	Host        types.ManagedObjectReference `xml:"host"`
	MetaUrl     string                       `xml:"metaUrl,omitempty"`
	MetaSha1Sum string                       `xml:"metaSha1Sum,omitempty"`
	VibUrl      string                       `xml:"vibUrl"`
	VibSha1Sum  string                       `xml:"vibSha1Sum"`
}

func init() {
	t["VsanVibSpec"] = reflect.TypeOf((*VsanVibSpec)(nil)).Elem()
}

type VsanClusterHealthAction struct {
	DynamicData

	ActionId          string             `xml:"actionId"`
	ActionLabel       LocalizableMessage `xml:"actionLabel"`
	ActionDescription LocalizableMessage `xml:"actionDescription"`
	Enabled           bool               `xml:"enabled"`
}

func init() {
	t["VsanClusterHealthAction"] = reflect.TypeOf((*VsanClusterHealthAction)(nil)).Elem()
}

type LocalizableMessage struct {
	DynamicData

	Key     string        `xml:"key"`
	Arg     []KeyAnyValue `xml:"arg,omitempty"`
	Message string        `xml:"message,omitempty"`
}

func init() {
	t["LocalizableMessage"] = reflect.TypeOf((*LocalizableMessage)(nil)).Elem()
}

type KeyAnyValue struct {
	DynamicData

	Key   string  `xml:"key"`
	Value AnyType `xml:"value,typeattr"`
}

func init() {
	t["KeyAnyValue"] = reflect.TypeOf((*KeyAnyValue)(nil)).Elem()
}

type ChoiceOption struct {
	OptionType

	ChoiceInfo   []BaseElementDescription `xml:"choiceInfo,typeattr"`
	DefaultIndex int32                    `xml:"defaultIndex,omitempty"`
}

func init() {
	t["ChoiceOption"] = reflect.TypeOf((*ChoiceOption)(nil)).Elem()
}

type VsanHclDriverInfo struct {
	DynamicData

	DriverVersion string             `xml:"driverVersion,omitempty"`
	DriverLink    *VsanDownloadItem  `xml:"driverLink,omitempty"`
	FwVersion     string             `xml:"fwVersion,omitempty"`
	FwLinks       []VsanDownloadItem `xml:"fwLinks,omitempty"`
	ToolsLinks    []VsanDownloadItem `xml:"toolsLinks,omitempty"`
	Eula          string             `xml:"eula,omitempty"`
}

func init() {
	t["VsanHclDriverInfo"] = reflect.TypeOf((*VsanHclDriverInfo)(nil)).Elem()
}

type OptionType struct {
	DynamicData

	ValueIsReadonly *bool `xml:"valueIsReadonly"`
}

func init() {
	t["OptionType"] = reflect.TypeOf((*OptionType)(nil)).Elem()
}

type VsanClusterClomdLivenessResult struct {
	DynamicData

	ClomdLivenessResult []VsanHostClomdLivenessResult `xml:"clomdLivenessResult,omitempty"`
	IssueFound          bool                          `xml:"issueFound"`
}

func init() {
	t["VsanClusterClomdLivenessResult"] = reflect.TypeOf((*VsanClusterClomdLivenessResult)(nil)).Elem()
}

type VsanHostClomdLivenessResult struct {
	DynamicData

	Hostname  string                `xml:"hostname"`
	ClomdStat string                `xml:"clomdStat"`
	Error     *LocalizedMethodFault `xml:"error,omitempty"`
}

func init() {
	t["VsanHostClomdLivenessResult"] = reflect.TypeOf((*VsanHostClomdLivenessResult)(nil)).Elem()
}

type VsanClusterBalanceSummary struct {
	DynamicData

	VarianceThreshold int64                           `xml:"varianceThreshold"`
	Disks             []VsanClusterBalancePerDiskInfo `xml:"disks,omitempty"`
}

func init() {
	t["VsanClusterBalanceSummary"] = reflect.TypeOf((*VsanClusterBalanceSummary)(nil)).Elem()
}

type VsanClusterBalancePerDiskInfo struct {
	DynamicData

	Uuid                   string `xml:"uuid,omitempty"`
	Fullness               int64  `xml:"fullness"`
	Variance               int64  `xml:"variance"`
	FullnessAboveThreshold int64  `xml:"fullnessAboveThreshold"`
	DataToMoveB            int64  `xml:"dataToMoveB"`
}

func init() {
	t["VsanClusterBalancePerDiskInfo"] = reflect.TypeOf((*VsanClusterBalancePerDiskInfo)(nil)).Elem()
}

type VsanGenericClusterBestPracticeHealth struct {
	DynamicData

	DrsEnabled bool                          `xml:"drsEnabled"`
	HaEnabled  bool                          `xml:"haEnabled"`
	Issues     []VsanGenericClusterBaseIssue `xml:"issues,omitempty"`
}

func init() {
	t["VsanGenericClusterBestPracticeHealth"] = reflect.TypeOf((*VsanGenericClusterBestPracticeHealth)(nil)).Elem()
}

type VsanGenericClusterBaseIssue struct {
	DynamicData
}

func init() {
	t["VsanGenericClusterBaseIssue"] = reflect.TypeOf((*VsanGenericClusterBaseIssue)(nil)).Elem()
}

type VsanNetworkConfigBestPracticeHealth struct {
	DynamicData

	VdsPresent bool                             `xml:"vdsPresent"`
	Issues     []BaseVsanNetworkConfigBaseIssue `xml:"issues,omitempty,typeattr"`
}

func init() {
	t["VsanNetworkConfigBestPracticeHealth"] = reflect.TypeOf((*VsanNetworkConfigBestPracticeHealth)(nil)).Elem()
}

type VsanBurnInTestCheckResult struct {
	DynamicData

	PassedTests       []VsanBurnInTest `xml:"passedTests,omitempty"`
	NotPerformedTests []VsanBurnInTest `xml:"notPerformedTests,omitempty"`
	FailedTests       []VsanBurnInTest `xml:"failedTests,omitempty"`
}

func init() {
	t["VsanBurnInTestCheckResult"] = reflect.TypeOf((*VsanBurnInTestCheckResult)(nil)).Elem()
}

type VsanBurnInTest struct {
	DynamicData

	Testname string `xml:"testname"`
	Workload string `xml:"workload,omitempty"`
	Duration int64  `xml:"duration"`
	Result   string `xml:"result"`
}

func init() {
	t["VsanBurnInTest"] = reflect.TypeOf((*VsanBurnInTest)(nil)).Elem()
}

// Space Usage
type DynamicData struct {
}

type VsanQuerySpaceUsage VsanQuerySpaceUsageRequestType

func init() {
	t["VsanQuerySpaceUsage"] = reflect.TypeOf((*VsanQuerySpaceUsage)(nil)).Elem()
}

type VsanQuerySpaceUsageRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

func init() {
	t["VsanQuerySpaceUsageRequestType"] = reflect.TypeOf((*VsanQuerySpaceUsageRequestType)(nil)).Elem()
}

type VsanQuerySpaceUsageResponse struct {
	Returnval VsanSpaceUsage `xml:"returnval"`
}

type VsanSpaceUsage struct {
	DynamicData

	TotalCapacityB int64                       `xml:"totalCapacityB"`
	FreeCapacityB  int64                       `xml:"freeCapacityB,omitempty"`
	SpaceOverview  *VsanObjectSpaceSummary     `xml:"spaceOverview,omitempty"`
	SpaceDetail    *VsanSpaceUsageDetailResult `xml:"spaceDetail,omitempty"`
}

func init() {
	t["VsanSpaceUsage"] = reflect.TypeOf((*VsanSpaceUsage)(nil)).Elem()
}

type VsanObjectSpaceSummary struct {
	DynamicData

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
	t["VsanObjectSpaceSummary"] = reflect.TypeOf((*VsanObjectSpaceSummary)(nil)).Elem()
}

type VsanSpaceUsageDetailResult struct {
	DynamicData

	SpaceUsageByObjectType []VsanObjectSpaceSummary `xml:"spaceUsageByObjectType,omitempty"`
}

func init() {
	t["VsanSpaceUsageDetailResult"] = reflect.TypeOf((*VsanSpaceUsageDetailResult)(nil)).Elem()
}

// Performance
type VsanPerfQueryPerf VsanPerfQueryPerfRequestType

func init() {
	t["VsanPerfQueryPerf"] = reflect.TypeOf((*VsanPerfQueryPerf)(nil)).Elem()
}

type VsanPerfQueryPerfRequestType struct {
	This       types.ManagedObjectReference  `xml:"_this"`
	QuerySpecs []VsanPerfQuerySpec           `xml:"querySpecs"`
	Cluster    *types.ManagedObjectReference `xml:"cluster,omitempty"`
}

func init() {
	t["VsanPerfQueryPerfRequestType"] = reflect.TypeOf((*VsanPerfQueryPerfRequestType)(nil)).Elem()
}

type VsanPerfQueryPerfResponse struct {
	Returnval []VsanPerfEntityMetricCSV `xml:"returnval"`
}

type VsanPerfQuerySpec struct {
	DynamicData

	EntityRefId string     `xml:"entityRefId"`
	StartTime   *time.Time `xml:"startTime"`
	EndTime     *time.Time `xml:"endTime"`
	Group       string     `xml:"group,omitempty"`
	Labels      []string   `xml:"labels,omitempty"`
	Interval    int32      `xml:"interval,omitempty"`
}

type VsanPerfEntityMetricCSV struct {
	DynamicData

	EntityRefId string                    `xml:"entityRefId"`
	SampleInfo  string                    `xml:"sampleInfo,omitempty"`
	Value       []VsanPerfMetricSeriesCSV `xml:"value,omitempty"`
}

type VsanPerfMetricSeriesCSV struct {
	DynamicData

	MetricId  VsanPerfMetricId   `xml:"metricId"`
	Threshold *VsanPerfThreshold `xml:"threshold,omitempty"`
	Values    string             `xml:"values,omitempty"`
}

func init() {
	t["VsanPerfMetricSeriesCSV"] = reflect.TypeOf((*VsanPerfMetricSeriesCSV)(nil)).Elem()
}

type VsanPerfMetricId struct {
	DynamicData

	Label                  string `xml:"label"`
	Group                  string `xml:"group,omitempty"`
	RollupType             string `xml:"rollupType,omitempty"`
	StatsType              string `xml:"statsType,omitempty"`
	Name                   string `xml:"name,omitempty"`
	Description            string `xml:"description,omitempty"`
	MetricsCollectInterval int32  `xml:"metricsCollectInterval,omitempty"`
}

func init() {
	t["VsanPerfMetricId"] = reflect.TypeOf((*VsanPerfMetricId)(nil)).Elem()
}

// Syncing summary
type VsanQuerySyncingVsanObjects VsanQuerySyncingVsanObjectsRequestType

func init() {
	t["VsanQuerySyncingVsanObjects"] = reflect.TypeOf((*VsanQuerySyncingVsanObjects)(nil)).Elem()
}

type VsanQuerySyncingVsanObjectsRequestType struct {
	This           types.ManagedObjectReference `xml:"_this"`
	Uuids          []string                     `xml:"uuids,omitempty"`
	Start          int32                        `xml:"start,omitempty"`
	Limit          *int32                       `xml:"limit"`
	IncludeSummary *bool                        `xml:"includeSummary"`
}

func init() {
	t["VsanQuerySyncingVsanObjectsRequestType"] = reflect.TypeOf((*VsanQuerySyncingVsanObjectsRequestType)(nil)).Elem()
}

type VsanQuerySyncingVsanObjectsResponse struct {
	Returnval VsanHostVsanObjectSyncQueryResult `xml:"returnval"`
}

type VsanHostVsanObjectSyncQueryResult struct {
	DynamicData

	TotalObjectsToSync int64                         `xml:"totalObjectsToSync,omitempty"`
	TotalBytesToSync   int64                         `xml:"totalBytesToSync,omitempty"`
	TotalRecoveryETA   int64                         `xml:"totalRecoveryETA,omitempty"`
	Objects            []VsanHostVsanObjectSyncState `xml:"objects,omitempty"`
}

func init() {
	t["VsanHostVsanObjectSyncQueryResult"] = reflect.TypeOf((*VsanHostVsanObjectSyncQueryResult)(nil)).Elem()
}

type VsanHostVsanObjectSyncState struct {
	DynamicData

	Uuid       string                       `xml:"uuid"`
	Components []VsanHostComponentSyncState `xml:"components"`
}

func init() {
	t["VsanHostVsanObjectSyncState"] = reflect.TypeOf((*VsanHostVsanObjectSyncState)(nil)).Elem()
}

type VsanHostComponentSyncState struct {
	DynamicData

	Uuid        string   `xml:"uuid"`
	DiskUuid    string   `xml:"diskUuid"`
	HostUuid    string   `xml:"hostUuid"`
	BytesToSync int64    `xml:"bytesToSync"`
	RecoveryETA int64    `xml:"recoveryETA,omitempty"`
	Reasons     []string `xml:"reasons,omitempty"`
}

func init() {
	t["VsanHostComponentSyncState"] = reflect.TypeOf((*VsanHostComponentSyncState)(nil)).Elem()
}

type MethodFault struct {
	FaultCause   *LocalizedMethodFault `xml:"faultCause,omitempty"`
	FaultMessage []LocalizableMessage  `xml:"faultMessage,omitempty"`
}

func init() {
	t["MethodFault"] = reflect.TypeOf((*MethodFault)(nil)).Elem()
}

type VsanClusterHealthResultBase struct {
	DynamicData

	Label string `xml:"label,omitempty"`
}

func init() {
	t["VsanClusterHealthResultBase"] = reflect.TypeOf((*VsanClusterHealthResultBase)(nil)).Elem()
}

type ElementDescription struct {
	Description

	Key string `xml:"key"`
}

func init() {
	t["ElementDescription"] = reflect.TypeOf((*ElementDescription)(nil)).Elem()
}

type Description struct {
	DynamicData

	Label   string `xml:"label"`
	Summary string `xml:"summary"`
}

func init() {
	t["Description"] = reflect.TypeOf((*Description)(nil)).Elem()
}

type VsanNetworkConfigBaseIssue struct {
	DynamicData
}

func init() {
	t["VsanNetworkConfigBaseIssue"] = reflect.TypeOf((*VsanNetworkConfigBaseIssue)(nil)).Elem()
}

type VsanClusterConfigInfo struct {
	DynamicData

	Enabled       *bool                                 `xml:"enabled"`
	DefaultConfig *VsanClusterConfigInfoHostDefaultInfo `xml:"defaultConfig,omitempty"`
}

func init() {
	t["VsanClusterConfigInfo"] = reflect.TypeOf((*VsanClusterConfigInfo)(nil)).Elem()
}

type VsanClusterConfigInfoHostDefaultInfo struct {
	DynamicData

	Uuid             string `xml:"uuid,omitempty"`
	AutoClaimStorage *bool  `xml:"autoClaimStorage"`
	ChecksumEnabled  *bool  `xml:"checksumEnabled"`
}

func init() {
	t["VsanClusterConfigInfoHostDefaultInfo"] = reflect.TypeOf((*VsanClusterConfigInfoHostDefaultInfo)(nil)).Elem()
}
