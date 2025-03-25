// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package esx

import "github.com/vmware/govmomi/vim25/types"

// HardwareVersion is the default VirtualMachine.Config.Version
var HardwareVersion = "vmx-13"

// AdvancedOptions is captured from ESX's HostSystem.configManager.advancedOption
// Capture method:
//
//	govc object.collect -s -dump $(govc object.collect -s HostSystem:ha-host configManager.advancedOption) setting
var AdvancedOptions = []types.BaseOptionValue{
	// This list is currently pruned to include a single option for testing
	&types.OptionValue{
		Key:   "Config.HostAgent.log.level",
		Value: "info",
	},
}

// Setting is captured from ESX's HostSystem.ServiceContent.setting
// Capture method:
//
//	govc object.collect -s -dump OptionManager:HostAgentSettings setting
var Setting = []types.BaseOptionValue{}
