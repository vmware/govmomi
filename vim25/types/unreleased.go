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
	Add("PodVMOverheadInfo", reflect.TypeOf((*PodVMOverheadInfo)(nil)).Elem())
}

type PodVMOverheadInfo struct {
	CrxPageSharingSupported         bool  `xml:"crxPageSharingSupported"`
	PodVMOverheadWithoutPageSharing int32 `xml:"podVMOverheadWithoutPageSharing"`
	PodVMOverheadWithPageSharing    int32 `xml:"podVMOverheadWithPageSharing"`
}

// Describes an action for the initial placement of a virtual machine in a cluster.
//
// This action is used by the cross cluster placement API when a virtual machine
// needs to be placed across a set of given clusters. See `Folder.PlaceVmsXCluster`.
// This action encapsulates details about the chosen cluster (via the resource pool
// inside that cluster), the chosen host and the chosen datastores for the disks of
// the virtual machine.
type ClusterClusterInitialPlacementActionEx struct {
	ClusterAction

	// The host where the virtual machine should be initially placed.
	//
	// This field is optional because the primary use case of
	// `Folder.PlaceVmsXCluster` is to select the best cluster for placing VMs. This
	// `ClusterClusterInitialPlacementAction.targetHost` denotes the best host
	// within the best cluster and it is only returned if the client asks for it,
	// which is determined by `PlaceVmsXClusterSpec.hostRecommRequired`.
	// If `PlaceVmsXClusterSpec.hostRecommRequired` is set to true, then the
	// targetHost is returned with a valid value and if it is either set to false
	// or left unset, then targetHost is also left unset. When this field is unset,
	// then it means that the client did not ask for the target host within the
	// recommended cluster. It does not mean that there is no recommended host
	// for placing this VM in the recommended cluster.
	//
	// Refers instance of `HostSystem`.
	TargetHost *ManagedObjectReference `xml:"targetHost,omitempty" json:"targetHost,omitempty"`

	// The chosen resource pool for placing the virtual machine.
	//
	// This is non-optional because recommending the best cluster (by recommending the
	// resource pool in the best cluster) is the primary use case for the
	// `ClusterClusterInitialPlacementAction`.
	//
	// Refers instance of `ResourcePool`.
	Pool ManagedObjectReference `xml:"pool" json:"pool"`

	// The config spec of the virtual machine to be placed.
	//
	// The `Folder.PlaceVmsXCluster` method takes input of `VirtualMachineConfigSpec`
	// from client and populates the backing for each virtual disk and the VM home
	// path in it unless the input ConfigSpec already provides them. The existing
	// settings in the input ConfigSpec are preserved and not overridden in the
	// returned ConfigSpec in this action as well as the resulting
	// `ClusterRecommendation`. This field is set based on whether the client needs
	// `Folder.PlaceVmsXCluster` to recommend a backing datastore for the disks of
	// the candidate VMs or not, which is specified via
	// `PlaceVmsXClusterSpec.datastoreRecommRequired`. If
	// `PlaceVmsXClusterSpec.datastoreRecommRequired` is set to true, then this
	// `ClusterClusterInitialPlacementAction.configSpec` is also set with the
	// backing of each disk populated. If
	// `PlaceVmsXClusterSpec.datastoreRecommRequired` is either set to false or left
	// unset, then this field is also left unset. When this field is left unset,
	// then it means that the client did not ask to populate the backing datastore
	// for the disks of the candidate VMs.
	ConfigSpec *VirtualMachineConfigSpec `xml:"configSpec,omitempty" json:"configSpec,omitempty"`

	AvailableNetworks []ManagedObjectReference `xml:"availableNetworks,omitempty" json:"availableNetworks,omitempty"`
}

func init() {
	minAPIVersionForType["ClusterClusterInitialPlacementActionEx"] = "9.1.0.0"
	t["ClusterClusterInitialPlacementActionEx"] = reflect.TypeOf((*ClusterClusterInitialPlacementActionEx)(nil)).Elem()
	Add("ClusterClusterInitialPlacementAction", reflect.TypeOf((*ClusterClusterInitialPlacementActionEx)(nil)).Elem())
}
