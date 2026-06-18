// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package types

import "reflect"

type ArrayOfPlaceVmsXClusterResultPlacementFaults struct {
	PlaceVmsXClusterResultPlacementFaults []PlaceVmsXClusterResultPlacementFaults `xml:"PlaceVmsXClusterResultPlacementFaults,omitempty"`
}

func init() {
	t["ArrayOfPlaceVmsXClusterResultPlacementFaults"] = reflect.TypeOf((*ArrayOfPlaceVmsXClusterResultPlacementFaults)(nil)).Elem()
}

type ArrayOfPlaceVmsXClusterResultPlacementInfo struct {
	PlaceVmsXClusterResultPlacementInfo []PlaceVmsXClusterResultPlacementInfo `xml:"PlaceVmsXClusterResultPlacementInfo,omitempty"`
}

func init() {
	t["ArrayOfPlaceVmsXClusterResultPlacementInfo"] = reflect.TypeOf((*ArrayOfPlaceVmsXClusterResultPlacementInfo)(nil)).Elem()
}

type ArrayOfPlaceVmsXClusterSpecVmPlacementSpec struct {
	PlaceVmsXClusterSpecVmPlacementSpec []PlaceVmsXClusterSpecVmPlacementSpec `xml:"PlaceVmsXClusterSpecVmPlacementSpec,omitempty"`
}

func init() {
	t["ArrayOfPlaceVmsXClusterSpecVmPlacementSpec"] = reflect.TypeOf((*ArrayOfPlaceVmsXClusterSpecVmPlacementSpec)(nil)).Elem()
}

type ArrayOfPlaceVmsXClusterSpecVmPlacementSpecCandidateNetworks struct {
	PlaceVmsXClusterSpecVmPlacementSpecCandidateNetworks []PlaceVmsXClusterSpecVmPlacementSpecCandidateNetworks `xml:"PlaceVmsXClusterSpecVmPlacementSpecCandidateNetworks,omitempty"`
}

func init() {
	t["ArrayOfPlaceVmsXClusterSpecVmPlacementSpecCandidateNetworks"] = reflect.TypeOf((*ArrayOfPlaceVmsXClusterSpecVmPlacementSpecCandidateNetworks)(nil)).Elem()
}

type PlaceVmsXCluster PlaceVmsXClusterRequestType

func init() {
	t["PlaceVmsXCluster"] = reflect.TypeOf((*PlaceVmsXCluster)(nil)).Elem()
}

type PlaceVmsXClusterRequestType struct {
	This          ManagedObjectReference `xml:"_this"`
	PlacementSpec PlaceVmsXClusterSpec   `xml:"placementSpec"`
}

func init() {
	t["PlaceVmsXClusterRequestType"] = reflect.TypeOf((*PlaceVmsXClusterRequestType)(nil)).Elem()
}

type PlaceVmsXClusterResponse struct {
	Returnval PlaceVmsXClusterResult `xml:"returnval"`
}

type PlaceVmsXClusterResult struct {
	DynamicData

	PlacementInfos []PlaceVmsXClusterResultPlacementInfo   `xml:"placementInfos,omitempty"`
	Faults         []PlaceVmsXClusterResultPlacementFaults `xml:"faults,omitempty"`
}

func init() {
	t["PlaceVmsXClusterResult"] = reflect.TypeOf((*PlaceVmsXClusterResult)(nil)).Elem()
}

type PlaceVmsXClusterResultPlacementFaults struct {
	DynamicData

	ResourcePool ManagedObjectReference  `xml:"resourcePool"`
	VmName       string                  `xml:"vmName"`
	Faults       []LocalizedMethodFault  `xml:"faults,omitempty"`
	Vm           *ManagedObjectReference `xml:"vm,omitempty"`
}

func init() {
	t["PlaceVmsXClusterResultPlacementFaults"] = reflect.TypeOf((*PlaceVmsXClusterResultPlacementFaults)(nil)).Elem()
}

type PlaceVmsXClusterResultPlacementInfo struct {
	DynamicData

	VmName         string                  `xml:"vmName"`
	Recommendation ClusterRecommendation   `xml:"recommendation"`
	Vm             *ManagedObjectReference `xml:"vm,omitempty"`
}

func init() {
	t["PlaceVmsXClusterResultPlacementInfo"] = reflect.TypeOf((*PlaceVmsXClusterResultPlacementInfo)(nil)).Elem()
}

type PlaceVmsXClusterSpec struct {
	DynamicData

	ResourcePools           []ManagedObjectReference              `xml:"resourcePools,omitempty"`
	PlacementType           string                                `xml:"placementType,omitempty"`
	VmPlacementSpecs        []PlaceVmsXClusterSpecVmPlacementSpec `xml:"vmPlacementSpecs,omitempty"`
	HostRecommRequired      *bool                                 `xml:"hostRecommRequired"`
	DatastoreRecommRequired *bool                                 `xml:"datastoreRecommRequired"`
}

func init() {
	t["PlaceVmsXClusterSpec"] = reflect.TypeOf((*PlaceVmsXClusterSpec)(nil)).Elem()
}

type PlaceVmsXClusterSpecVmPlacementSpec struct {
	DynamicData

	Vm                *ManagedObjectReference                                `xml:"vm,omitempty"`
	ConfigSpec        VirtualMachineConfigSpec                               `xml:"configSpec"`
	RelocateSpec      *VirtualMachineRelocateSpec                            `xml:"relocateSpec,omitempty"`
	CandidateNetworks []PlaceVmsXClusterSpecVmPlacementSpecCandidateNetworks `xml:"candidateNetworks,omitempty"`
}

func init() {
	t["PlaceVmsXClusterSpecVmPlacementSpec"] = reflect.TypeOf((*PlaceVmsXClusterSpecVmPlacementSpec)(nil)).Elem()
}

type PlaceVmsXClusterSpecVmPlacementSpecCandidateNetworks struct {
	DynamicData

	Networks []ManagedObjectReference `xml:"networks,omitempty"`
}

func init() {
	t["PlaceVmsXClusterSpecVmPlacementSpecCandidateNetworks"] = reflect.TypeOf((*PlaceVmsXClusterSpecVmPlacementSpecCandidateNetworks)(nil)).Elem()
}

const RecommendationReasonCodeXClusterPlacement = RecommendationReasonCode("xClusterPlacement")

type ClusterClusterReconfigurePlacementAction struct {
	ClusterAction
	TargetHost *ManagedObjectReference   `xml:"targetHost,omitempty"`
	Pool       ManagedObjectReference    `xml:"pool"`
	ConfigSpec *VirtualMachineConfigSpec `xml:"configSpec,omitempty"`
}

func init() {
	t["ClusterClusterReconfigurePlacementAction"] = reflect.TypeOf((*ClusterClusterReconfigurePlacementAction)(nil)).Elem()
}

type ClusterClusterRelocatePlacementAction struct {
	ClusterAction
	TargetHost        *ManagedObjectReference     `xml:"targetHost,omitempty"`
	Pool              ManagedObjectReference      `xml:"pool"`
	RelocateSpec      *VirtualMachineRelocateSpec `xml:"relocateSpec,omitempty"`
	AvailableNetworks []ManagedObjectReference    `xml:"availableNetworks,omitempty"`
}

func init() {
	t["ClusterClusterRelocatePlacementAction"] = reflect.TypeOf((*ClusterClusterRelocatePlacementAction)(nil)).Elem()
}

func init() {
	minAPIVersionForType["HostRuntimeInfoPodVMInfo"] = "9.1.0.0"
	Add("HostRuntimeInfoPodVMInfo", reflect.TypeOf((*HostRuntimeInfoPodVMInfo)(nil)).Elem())
}

type HostRuntimeInfoPodVMInfo struct {
	DynamicData

	HasPageSharingPodVM bool              `xml:"hasPageSharingPodVM"`
	PodVMOverheadInfo   PodVMOverheadInfo `xml:"podVMOverheadInfo"`
}

type UpdatePodVMPropertyRequestType struct {
	This ManagedObjectReference `xml:"_this" json:"-"`
	// Indicates the property within PodVMInfo to update
	PropertyPath string `xml:"propertyPath" json:"propertyPath"`
	// Value of propertyPath requested to be updated
	Property AnyType `xml:"property,omitempty,typeattr" json:"property,omitempty"`
}

func init() {
	t["UpdatePodVMPropertyRequestType"] = reflect.TypeOf((*UpdatePodVMPropertyRequestType)(nil)).Elem()
}

type UpdatePodVMProperty UpdatePodVMPropertyRequestType

func init() {
	minAPIVersionForType["UpdatePodVMProperty"] = "9.1.0.0"
	t["UpdatePodVMProperty"] = reflect.TypeOf((*UpdatePodVMProperty)(nil)).Elem()
}

type UpdatePodVMPropertyResponse struct {
}

type BaseClusterClusterInitialPlacementAction interface {
	GetClusterClusterInitialPlacementAction() *ClusterClusterInitialPlacementAction
}

func (a ClusterClusterInitialPlacementAction) GetClusterClusterInitialPlacementAction() *ClusterClusterInitialPlacementAction {
	return &a
}

func init() {
	minAPIVersionForType["ClusterClusterInitialPlacementActionEx"] = "9.1.0.0"
	t["ClusterClusterInitialPlacementAction"] = reflect.TypeOf((*ClusterClusterInitialPlacementAction)(nil)).Elem()
	t["BaseClusterClusterInitialPlacementAction"] = reflect.TypeOf((*ClusterClusterInitialPlacementAction)(nil)).Elem()
}

// SharedDiskVmGroupInfoSharedDiskVmInfo is a row in SharedDiskVmGroupInfo (vim.vm.SharedDiskVmGroupInfo.SharedDiskVmInfo).
type SharedDiskVmGroupInfoSharedDiskVmInfo struct {
	DynamicData

	DiskKey       int32           `xml:"diskKey" json:"diskKey"`
	VirtualDiskId []VirtualDiskId `xml:"virtualDiskId,omitempty" json:"virtualDiskId,omitempty"`
}

func init() {
	t["SharedDiskVmGroupInfoSharedDiskVmInfo"] = reflect.TypeOf((*SharedDiskVmGroupInfoSharedDiskVmInfo)(nil)).Elem()
}

// SharedDiskVmGroupInfo describes VMs sharing multi-writer or SCSI bus-sharing disks (vim.vm.SharedDiskVmGroupInfo).
type SharedDiskVmGroupInfo struct {
	DynamicData

	SharedDiskVmInfo []SharedDiskVmGroupInfoSharedDiskVmInfo `xml:"sharedDiskVmInfo,omitempty" json:"sharedDiskVmInfo,omitempty"`
}

func init() {
	t["SharedDiskVmGroupInfo"] = reflect.TypeOf((*SharedDiskVmGroupInfo)(nil)).Elem()
}

type FetchVmGroupForMultiwriterDisksRequestType struct {
	This    ManagedObjectReference `xml:"_this" json:"-"`
	DiskIds *ArrayOfInt            `xml:"diskIds,omitempty" json:"diskIds,omitempty"`
}

func init() {
	t["FetchVmGroupForMultiwriterDisksRequestType"] = reflect.TypeOf((*FetchVmGroupForMultiwriterDisksRequestType)(nil)).Elem()
}

// FetchVmGroupForMultiwriterDisks is VirtualMachine#fetchVmGroupForMultiwriterDisks (MultiwriterDiskVMotion).
type FetchVmGroupForMultiwriterDisks FetchVmGroupForMultiwriterDisksRequestType

func init() {
	t["FetchVmGroupForMultiwriterDisks"] = reflect.TypeOf((*FetchVmGroupForMultiwriterDisks)(nil)).Elem()
}

type FetchVmGroupForMultiwriterDisksResponse struct {
	Returnval *SharedDiskVmGroupInfo `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

// InsufficientResourcesQuotaResourceType identifies the type of resource whose quota was exceeded.
type InsufficientResourcesQuotaResourceType string

const (
	// InsufficientResourcesQuotaResourceTypeCpu indicates CPU quota was exceeded (unit: MHz).
	InsufficientResourcesQuotaResourceTypeCpu = InsufficientResourcesQuotaResourceType("cpu")
	// InsufficientResourcesQuotaResourceTypeMemory indicates memory quota was exceeded (unit: bytes).
	InsufficientResourcesQuotaResourceTypeMemory = InsufficientResourcesQuotaResourceType("memory")
)

func (e InsufficientResourcesQuotaResourceType) Values() []InsufficientResourcesQuotaResourceType {
	return []InsufficientResourcesQuotaResourceType{
		InsufficientResourcesQuotaResourceTypeCpu,
		InsufficientResourcesQuotaResourceTypeMemory,
	}
}

func (e InsufficientResourcesQuotaResourceType) Strings() []string {
	return EnumValuesAsStrings(e.Values())
}

func init() {
	t["InsufficientResourcesQuotaResourceType"] = reflect.TypeOf((*InsufficientResourcesQuotaResourceType)(nil)).Elem()
}

// InsufficientResourcesQuotaResourceQuotaType identifies the type of quota that was violated.
type InsufficientResourcesQuotaResourceQuotaType string

const (
	// InsufficientResourcesQuotaResourceQuotaTypeReservation indicates a violation against reservation capacity.
	InsufficientResourcesQuotaResourceQuotaTypeReservation = InsufficientResourcesQuotaResourceQuotaType("reservation")
	// InsufficientResourcesQuotaResourceQuotaTypeConfiguredSize indicates a violation against configured size capacity.
	InsufficientResourcesQuotaResourceQuotaTypeConfiguredSize = InsufficientResourcesQuotaResourceQuotaType("configuredSize")
)

func (e InsufficientResourcesQuotaResourceQuotaType) Values() []InsufficientResourcesQuotaResourceQuotaType {
	return []InsufficientResourcesQuotaResourceQuotaType{
		InsufficientResourcesQuotaResourceQuotaTypeReservation,
		InsufficientResourcesQuotaResourceQuotaTypeConfiguredSize,
	}
}

func (e InsufficientResourcesQuotaResourceQuotaType) Strings() []string {
	return EnumValuesAsStrings(e.Values())
}

func init() {
	t["InsufficientResourcesQuotaResourceQuotaType"] = reflect.TypeOf((*InsufficientResourcesQuotaResourceQuotaType)(nil)).Elem()
}

// InsufficientResourcesQuotaFailureNotificationMode indicates whether the failure is transient or permanent.
type InsufficientResourcesQuotaFailureNotificationMode string

const (
	// InsufficientResourcesQuotaFailureNotificationModeFailureModeTransient indicates the failure is temporary and may be resolved through reallocation.
	InsufficientResourcesQuotaFailureNotificationModeFailureModeTransient = InsufficientResourcesQuotaFailureNotificationMode("failureModeTransient")
	// InsufficientResourcesQuotaFailureNotificationModeFailureModePermanent indicates the failure cannot be resolved.
	InsufficientResourcesQuotaFailureNotificationModeFailureModePermanent = InsufficientResourcesQuotaFailureNotificationMode("failureModePermanent")
)

func (e InsufficientResourcesQuotaFailureNotificationMode) Values() []InsufficientResourcesQuotaFailureNotificationMode {
	return []InsufficientResourcesQuotaFailureNotificationMode{
		InsufficientResourcesQuotaFailureNotificationModeFailureModeTransient,
		InsufficientResourcesQuotaFailureNotificationModeFailureModePermanent,
	}
}

func (e InsufficientResourcesQuotaFailureNotificationMode) Strings() []string {
	return EnumValuesAsStrings(e.Values())
}

func init() {
	t["InsufficientResourcesQuotaFailureNotificationMode"] = reflect.TypeOf((*InsufficientResourcesQuotaFailureNotificationMode)(nil)).Elem()
}

// InsufficientResourcesQuota is thrown when an operation exceeds the quota in an associated resource envelope.
type InsufficientResourcesQuota struct {
	InsufficientResourcesFault

	// ResourceType identifies the type of resource whose quota was exceeded.
	// See InsufficientResourcesQuotaResourceType for supported values.
	ResourceType string `xml:"resourceType" json:"resourceType"`
	// ResourceQuotaType identifies the type of quota that was violated.
	// See InsufficientResourcesQuotaResourceQuotaType for supported values.
	ResourceQuotaType string `xml:"resourceQuotaType" json:"resourceQuotaType"`
	// Available is the amount of the resource still available under the quota
	// (MHz for CPU, bytes for memory).
	Available int64 `xml:"available" json:"available"`
	// Requested is the amount of the resource requested in the failed operation.
	Requested int64 `xml:"requested" json:"requested"`
	// Entity is the identifier of the entity that violated the quota (e.g., a VM or resource envelope).
	Entity string `xml:"entity" json:"entity"`
	// ResourceEnvelope is the identifier of the resource envelope whose quota was exceeded.
	ResourceEnvelope string `xml:"resourceEnvelope" json:"resourceEnvelope"`
	// FailureNotificationMode indicates whether the failure is transient or permanent.
	// See InsufficientResourcesQuotaFailureNotificationMode for supported values.
	FailureNotificationMode string `xml:"failureNotificationMode" json:"failureNotificationMode"`
}

func init() {
	t["InsufficientResourcesQuota"] = reflect.TypeOf((*InsufficientResourcesQuota)(nil)).Elem()
}

type InsufficientResourcesQuotaFault InsufficientResourcesQuota

func init() {
	t["InsufficientResourcesQuotaFault"] = reflect.TypeOf((*InsufficientResourcesQuotaFault)(nil)).Elem()
}
