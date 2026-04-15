// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// VirtualMachineSpec defines the desired state of a VirtualMachine.
type VirtualMachineSpec struct {
	// +optional

	Config *VirtualMachineConfigInfoSpec `json:"config,omitempty"`
}

// VirtualMachineStatus defines the observed state of a
// VirtualMachine.
type VirtualMachineStatus struct {
	// +optional

	// ObservedGeneration describes the value of the metadata.generation field
	// the last time this object was reconciled by its primary controller.
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// +optional

	// Conditions describes any conditions associated with this object.
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// +optional

	Config *VirtualMachineConfigInfoStatus `json:"config,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=vm
// +kubebuilder:storageversion:true
// +kubebuilder:subresource:status

// VirtualMachine is the schema for the
// VirtualMachine API and
// represents the desired state and observed status of a
// VirtualMachine resource.
type VirtualMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualMachineSpec   `json:"spec,omitempty"`
	Status VirtualMachineStatus `json:"status,omitempty"`
}

// GetConditions returns the status conditions for the
// VirtualMachine.
func (p VirtualMachine) GetConditions() []metav1.Condition {
	return p.Status.Conditions
}

// SetConditions sets the status conditions for the
// VirtualMachine.
func (p *VirtualMachine) SetConditions(conditions []metav1.Condition) {
	p.Status.Conditions = conditions
}

// GetConditions returns the conditions for the
// VirtualMachineStatus.
func (p VirtualMachineStatus) GetConditions() []metav1.Condition {
	return p.Conditions
}

// SetConditions sets the conditions for the
// VirtualMachineStatus.
func (p *VirtualMachineStatus) SetConditions(conditions []metav1.Condition) {
	p.Conditions = conditions
}

// +kubebuilder:object:root=true

// VirtualMachineList contains a list of
// VirtualMachine objects.
type VirtualMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualMachine `json:"items"`
}

func init() {
	objectTypes = append(
		objectTypes,
		&VirtualMachine{},
		&VirtualMachineList{})
}
