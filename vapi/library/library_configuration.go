// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package library

// Configuration represents the configuration at individual Content Library level and applies to all types of libraries.
type Configuration struct {
	ApplyLibraryUsageToItems *bool `json:"apply_library_usage_to_items,omitempty"`
}
