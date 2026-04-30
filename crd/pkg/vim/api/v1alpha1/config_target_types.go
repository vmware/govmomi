// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:validation:Enum=Fixed;Regex;Glob
type MatchType string

const (
	// MatchTypeFixed matches the exact value.
	MatchTypeFixed MatchType = "Fixed"

	// MatchTypeRegex matches the value using a regular expression.
	MatchTypeRegex MatchType = "Regex"

	// MatchTypeGlob matches the value using a glob pattern.
	MatchTypeGlob MatchType = "Glob"
)

type ConfigTargetExtraConfigKey struct {
	// +optional
	// +kubebuilder:default=Glob

	// Type describes the type of match to use.
	// Defaults to Glob.
	Type MatchType `json:"type"`

	// +required

	// Key is the extra config key to match.
	Key string `json:"value"`
}

type ConfigTargetExtraConfigSpec struct {
	// +optional
	// +listType=map
	// +listMapKey=key

	// Allowed describes the list of allowed extra config keys.
	//
	// If Allowed is non-empty, then a key *must* match one of the values in
	// the list.
	//
	// Denied takes precedent over Allowed.
	Allowed []ConfigTargetExtraConfigKey `json:"allowed,omitempty"`

	// +optional
	// +listType=map
	// +listMapKey=key

	// Denied describes the list of denied extra config keys.
	//
	// If Denied is non-empty, then a key *must not* match one of the values in
	// the list.
	//
	// Denied takes precedent over Allowed.
	Denied []ConfigTargetExtraConfigKey `json:"denied,omitempty"`
}

// ConfigTargetSpec defines the desired state of a ConfigTarget.
type ConfigTargetSpec struct {
	// +optional

	// NumCPUCores describes the number of CPU cores available to run a
	// virtual machine.
	NumCPUCores *IntRange `json:"numCPUCores,omitempty"`

	// +optional

	// NumNUMANodes describes the total number of NUMA nodes available.
	// A non-zero maximum value indicates NUMA alignment is supported.
	NumNUMANodes *IntRange `json:"numNUMANodes,omitempty"`

	// +optional

	// NumSimultaneousThreads describes the number of simultaneous threads
	// available.
	// A non-zero maximum value indicates HT/SMT is supported.
	NumSimultaneousThreads *IntRange `json:"numSimultaneousThreads,omitempty"`

	// +optional

	// Memory describes the amount of memory available to run a virtual machine.
	Memory *ResourceQuantityRange `json:"memory,omitempty"`

	// +optional

	// SMCPresent describes the presence of the System Management
	// Controller (Apple hardware).
	SMCPresent bool `json:"smcPresent,omitempty"`

	// +optional

	// SEVSupported describes whether AMD SEV is supported.
	SEVSupported bool `json:"sevSupported,omitempty"`

	// +optional

	// SEVSNPSupported describes whether AMD SEV-SNP is supported.
	SEVSNPSupported bool `json:"sevSnpSupported,omitempty"`

	// +optional

	// TDXSupported describes whether Intel TDX is supported.
	TDXSupported bool `json:"tdxSupported,omitempty"`

	// +optional

	// ExtraConfig describes the allowed/denied extra config keys.
	ExtraConfig *ConfigTargetExtraConfigSpec `json:"extraConfig,omitempty"`

	// +optional

	// LatencySensitivityLevels describes the supported latency sensitivity
	// levels.
	LatencySensitivityLevels []LatencySensitivityLevel `json:"latencySensitivityLevels,omitempty"`

	// +optional

	// CPULockedToMaxSupported describes whether CPU locking to the max
	// is supported.
	CPULockedToMaxSupported bool `json:"cpuLockedToMaxSupported,omitempty"`

	// +optional

	// MemoryLockedToMaxSupported describes whether memory locking to the max
	// is supported.
	MemoryLockedToMaxSupported bool `json:"memoryLockedToMaxSupported,omitempty"`

	// +optional

	// HugePagesSupported describes whether huge pages are supported.
	HugePagesSupported bool `json:"hugePagesSupported,omitempty"`

	// +optional

	// IOMMUSupported describes whether IOMMU is supported.
	IOMMUSupported bool `json:"iommuSupported,omitempty"`

	// +optional

	// RSSSupported describes whether Receive Side Scaling (RSS) is supported.
	// RSS enables RSS on the vNIC, allowing the guest OS to distribute incoming
	// traffic across multiple vCPU cores rather than relying on a single core,
	// which is a major bottleneck.
	RSSSupported bool `json:"rssSupported,omitempty"`

	// +optional

	// UDPRSSSupported describes whether RSS for UDP traffic is supported.
	UDPRSSSupported bool `json:"udpRSSSupported,omitempty"`

	// +optional

	// LargeReceiveOffloadSupported describes whether Large Receive Offload
	// (LRO) is supported.
	LROSupported bool `json:"lroSupported,omitempty"`

	// +optional

	// TxRxThreadModels describes the supported transmit/receive models.
	TxRxThreadModels []TxRxThreadModel `json:"txRxThreadModels,omitempty"`

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
