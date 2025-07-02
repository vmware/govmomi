// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vpx

import "github.com/vmware/govmomi/vim25/types"

// TODO: figure out whether this is Setting or AdvancedOptions - see esx/setting.go for the difference

// Setting is captured from VC's ServiceContent.OptionManager.setting
var Setting = []types.BaseOptionValue{
	// This list is currently pruned to include sso options only with sso.enabled set to false
	&types.OptionValue{
		Key:   "config.vpxd.sso.sts.uri",
		Value: "https://127.0.0.1/sts/STSService/vsphere.local",
	},
	&types.OptionValue{
		Key:   "config.vpxd.sso.solutionUser.privateKey",
		Value: "/etc/vmware-vpx/ssl/vcsoluser.key",
	},
	&types.OptionValue{
		Key:   "config.vpxd.sso.solutionUser.name",
		Value: "vpxd-b643d01c-928f-469b-96a5-d571d762a78e@vsphere.local",
	},
	&types.OptionValue{
		Key:   "config.vpxd.sso.solutionUser.certificate",
		Value: "/etc/vmware-vpx/ssl/vcsoluser.crt",
	},
	&types.OptionValue{
		Key:   "config.vpxd.sso.groupcheck.uri",
		Value: "https://127.0.0.1/sso-adminserver/sdk/vsphere.local",
	},
	&types.OptionValue{
		Key:   "config.vpxd.sso.enabled",
		Value: "false",
	},
	&types.OptionValue{
		Key:   "config.vpxd.sso.default.isGroup",
		Value: "false",
	},
	&types.OptionValue{
		Key:   "config.vpxd.sso.default.admin",
		Value: "Administrator@vsphere.local",
	},
	&types.OptionValue{
		Key:   "config.vpxd.sso.admin.uri",
		Value: "https://127.0.0.1/sso-adminserver/sdk/vsphere.local",
	},
	&types.OptionValue{
		Key:   "VirtualCenter.InstanceName",
		Value: "127.0.0.1",
	},
	&types.OptionValue{
		Key:   "event.batchsize",
		Value: int32(2000),
	},
	&types.OptionValue{
		Key:   "event.maxAge",
		Value: int32(30),
	},
	&types.OptionValue{
		Key:   "event.maxAgeEnabled",
		Value: bool(true),
	},
}
