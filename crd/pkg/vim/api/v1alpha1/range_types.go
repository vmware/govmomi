// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import "k8s.io/apimachinery/pkg/api/resource"

// IntRange describes a range of 32-bit integer values with min, max,
// and default.
type IntRange struct {
	// +required

	// Min is the minimum value.
	Min int32 `json:"min"`

	// +required

	// Max is the maximum value.
	Max int32 `json:"max"`

	// +required

	// Default is the default value.
	Default int32 `json:"default"`
}

// LongRange describes a range of 64-bit integer values with min, max,
// and default.
type LongRange struct {
	// +required

	// Min is the minimum value.
	Min int64 `json:"min"`

	// +required

	// Max is the maximum value.
	Max int64 `json:"max"`

	// +required

	// Default is the default value.
	Default int64 `json:"default"`
}

// ResourceQuantityRange describes a range of resource.Quantity values with
// min, max, and default.
type ResourceQuantityRange struct {
	// +required

	// Min is the minimum value.
	Min resource.Quantity `json:"min"`

	// +required

	// Max is the maximum value.
	Max resource.Quantity `json:"max"`

	// +required

	// Default is the default value.
	Default resource.Quantity `json:"default"`
}
