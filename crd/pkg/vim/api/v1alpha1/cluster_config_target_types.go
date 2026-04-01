// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ClusterConfigTargetSpec defines the desired state of a
// ClusterConfigTarget.
type ClusterConfigTargetSpec struct {
	// +required

	// ID is the managed object ID of the vSphere cluster.
	ID ManagedObjectID `json:"id"`
}

// ClusterConfigTargetStatus defines the observed state of a
// ClusterConfigTarget.
type ClusterConfigTargetStatus struct {
	ConfigTargetStatus `json:",inline"`

	// +optional

	// NumCPUs describes the number of logical CPUs available to run
	// virtual machines.
	NumCPUs int32 `json:"numCPUs,omitempty"`

	// +optional

	// NumCPUCores describes the number of physical CPU cores available.
	NumCPUCores int32 `json:"numCPUCores,omitempty"`

	// +optional

	// NumNumaNodes describes the total number of NUMA nodes available.
	NumNumaNodes int32 `json:"numNumaNodes,omitempty"`

	// +optional

	// MaxCPUsPerVM describes the maximum number of CPUs a single VM may
	// be assigned.
	MaxCPUsPerVM int32 `json:"maxCPUsPerVM,omitempty"`

	// +optional

	// MaxSimultaneousThreads describes the maximum SMT threads available.
	MaxSimultaneousThreads int32 `json:"maxSimultaneousThreads,omitempty"`

	// +optional

	// MaxMemOptimalPerf describes the maximum recommended memory size.
	MaxMemOptimalPerf *resource.Quantity `json:"maxMemOptimalPerf,omitempty"`

	// +optional

	// SupportedMaxMem describes the maximum supported memory size for
	// creating a new VM.
	SupportedMaxMem *resource.Quantity `json:"supportedMaxMem,omitempty"`

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

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:storageversion:true
// +kubebuilder:subresource:status

// ClusterConfigTarget is the schema for the
// ClusterConfigTarget API and
// represents the desired state and observed status of a
// ClusterConfigTarget resource.
type ClusterConfigTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterConfigTargetSpec   `json:"spec,omitempty"`
	Status ClusterConfigTargetStatus `json:"status,omitempty"`
}

// GetConditions returns the status conditions for the
// ClusterConfigTarget.
func (p ClusterConfigTarget) GetConditions() []metav1.Condition {
	return p.Status.Conditions
}

// SetConditions sets the status conditions for the
// ClusterConfigTarget.
func (p *ClusterConfigTarget) SetConditions(conditions []metav1.Condition) {
	p.Status.Conditions = conditions
}

// +kubebuilder:object:root=true

// ClusterConfigTargetList contains a list of
// ClusterConfigTarget objects.
type ClusterConfigTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterConfigTarget `json:"items"`
}

func init() {
	objectTypes = append(
		objectTypes,
		&ClusterConfigTarget{},
		&ClusterConfigTargetList{})
}
