// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

const (
	// ReadyConditionType is the Ready condition type that summarizes the
	// operational state of an API.
	ReadyConditionType = "Ready"
)

// LocalObjectRef describes a reference to another object in the same
// namespace as the referrer.
type LocalObjectRef struct {
	// APIVersion defines the versioned schema of this representation of an
	// object. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	APIVersion string `json:"apiVersion"`

	// Kind is a string value representing the REST resource this object
	// represents.
	// Servers may infer this from the endpoint the client submits requests to.
	// Cannot be updated.
	// In CamelCase.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind string `json:"kind"`

	// Name refers to a unique resource in the current namespace.
	// More info: http://kubernetes.io/docs/user-guide/identifiers#names
	Name string `json:"name"`
}

// ManagedObjectID is a unique ID used to identify a managed object on a given
// vSphere instance
type ManagedObjectID struct {
	// +required

	// ID is the object's ID.
	ID string `json:"id"`

	// +optional

	// ServerID is the ID of the server to which the object belongs.
	ServerID string `json:"serverID,omitempty"`
}

// ManagedObjectReference is a referenced to a managed object on a given
// vSphere instance.
type ManagedObjectReference struct {
	ManagedObjectID `json:",inline"`

	// Type is the object's type.
	Type string `json:"type"`
}
