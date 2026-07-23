// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

// VirtualControllerOption describes the options for a virtual controller.
// It corresponds to vim.vm.device.VirtualControllerOption.
type VirtualControllerOption struct {

	// Devices describes the minimum and maximum number of devices this
	// controller can control at run time.
	Devices IntOption `json:"devices"`

	// +optional

	// SupportedDevice lists the device types supported by this controller.
	SupportedDevice []SupportedDeviceType `json:"supportedDevice,omitempty"`
}
