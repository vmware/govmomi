// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ConfigTargetSpec defines the desired state of a
// ConfigTarget.
type ConfigTargetSpec struct {
	// +optional
	// +listType=map
	// +listMapKey=name

	// Zones describes the desired state for one or more config target zones.
	Zones []ConfigTargetZoneSpec `json:"zones,omitempty"`
}

// ConfigTargetZoneSpec describes the desired state of a
// zone within a ConfigTarget.
type ConfigTargetZoneSpec struct {
	// +required

	// Name describes the desired state for the zone's name.
	Name string `json:"name"`

	// +listType=map
	// +listMapKey=name

	// Clusters describes the desired state for one or more vSphere
	// clusters within the zone.
	Clusters []ConfigTargetClusterSpec `json:"clusters"`
}

// ConfigTargetClusterSpec describes the desired state of a
// vSphere cluster within a ConfigTarget zone.
type ConfigTargetClusterSpec struct {
	// +required

	// ID describes the desired state for the config target's managed
	// object ID.
	ID ManagedObjectID `json:"id"`

	// +required

	// Name describes the name of the config target.
	Name string `json:"name"`

	// +optional

	// NumCPUs describes the number of logical CPUs available to run
	// a virtual machine.
	NumCPUs *IntRange `json:"numCPUs,omitempty"`

	// +optional

	// NumCPUCores describes the number of CPU cores available to run a
	// virtual machine.
	NumCPUCores *IntRange `json:"numCPUCores,omitempty"`

	// +optional

	// NumNUMANodes describes the total number of NUMA nodes available.
	NumNUMANodes *IntRange `json:"numNUMANodes,omitempty"`

	// +optional

	// NumSimultaneousThreads describes the number of simultaneous threads
	// available.
	NumSimultaneousThreads *IntRange `json:"numSimultaneousThreads,omitempty"`

	// +optional

	// Memory describes the amount of memory available to run a virtual machine.
	Memory *ResourceQuantityRange `json:"memory,omitempty"`

	// +optional

	// RecommendedMemory describes the recommended amount of memory to run a
	// virtual machine.
	RecommendedMemory *resource.Quantity `json:"recommendedMemory,omitempty"`

	// +optional

	// AvailablePersistentMemoryReservation describes the maximum available
	// persistent memory reservation.
	AvailablePersistentMemoryReservation *resource.Quantity `json:"availablePersistentMemoryReservation,omitempty"`

	// +optional

	// SmcPresent describes the presence of the System Management
	// Controller (Apple hardware).
	SmcPresent bool `json:"smcPresent,omitempty"`

	// +optional

	// SEVSupported describes whether AMD SEV is supported.
	SEVSupported bool `json:"sevSupported,omitempty"`

	// +optional

	// SEVSNPSupported describes whether AMD SEV-SNP is supported.
	SEVSNPSupported bool `json:"sevSnpSupported,omitempty"`

	// +optional

	// TDXSupported describes whether Intel TDX is supported.
	TDXSupported bool `json:"tdxSupported,omitempty"`

	ConfigTargetDevices `json:",inline"`
}

// ConfigTargetStatus defines the observed state of a ConfigTarget.
type ConfigTargetStatus struct {
	// +optional

	// ObservedGeneration describes the observed state of the
	// metadata.generation field at the time this object was last
	// reconciled by its primary controller.
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// +optional

	// Conditions describes the observed state of any conditions
	// associated with this object.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:storageversion:true
// +kubebuilder:subresource:status

// ConfigTarget is the schema for the ConfigTarget API.
type ConfigTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec describes the desired state of the ConfigTarget.
	Spec ConfigTargetSpec `json:"spec,omitempty"`

	// Status describes the observed state of the ConfigTarget.
	Status ConfigTargetStatus `json:"status,omitempty"`
}

// GetConditions returns the status conditions for the ConfigTarget.
func (p ConfigTarget) GetConditions() []metav1.Condition {
	return p.Status.Conditions
}

// SetConditions sets the status conditions for the ConfigTarget.
func (p *ConfigTarget) SetConditions(conditions []metav1.Condition) {
	p.Status.Conditions = conditions
}

// GetConditions returns the conditions for the ConfigTargetStatus.
func (p ConfigTargetStatus) GetConditions() []metav1.Condition {
	return p.Conditions
}

// SetConditions sets the conditions for the ConfigTargetStatus.
func (p *ConfigTargetStatus) SetConditions(conditions []metav1.Condition) {
	p.Conditions = conditions
}

// +kubebuilder:object:root=true

// ConfigTargetList contains a list of ConfigTarget objects.
type ConfigTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConfigTarget `json:"items"`
}

func init() {
	objectTypes = append(
		objectTypes,
		&ConfigTarget{},
		&ConfigTargetList{})
}
