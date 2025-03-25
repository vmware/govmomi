// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"github.com/vmware/govmomi/vim25/types"
)

type VsanClusterConfigInfo types.VsanClusterConfigInfo

func (b *VsanClusterConfigInfo) GetVsanClusterConfigInfo() *VsanClusterConfigInfo { return b }

type BaseVsanClusterConfigInfo interface {
	GetVsanClusterConfigInfo() *VsanClusterConfigInfo
}

func (b *VsanResourceConstraint) GetVsanResourceConstraint() *VsanResourceConstraint { return b }

type BaseVsanResourceConstraint interface {
	GetVsanResourceConstraint() *VsanResourceConstraint
}

func (b *VsanIscsiLUNCommonInfo) GetVsanIscsiLUNCommonInfo() *VsanIscsiLUNCommonInfo { return b }

type BaseVsanIscsiLUNCommonInfo interface {
	GetVsanIscsiLUNCommonInfo() *VsanIscsiLUNCommonInfo
}

func (b *EntityResourceCheckDetails) GetEntityResourceCheckDetails() *EntityResourceCheckDetails {
	return b
}

type BaseEntityResourceCheckDetails interface {
	GetEntityResourceCheckDetails() *EntityResourceCheckDetails
}

func (b *VsanDatastoreConfig) GetVsanDatastoreConfig() *VsanDatastoreConfig { return b }

type BaseVsanDatastoreConfig interface {
	GetVsanDatastoreConfig() *VsanDatastoreConfig
}

func (b *VsanDatastoreSpec) GetVsanDatastoreSpec() *VsanDatastoreSpec { return b }

type BaseVsanDatastoreSpec interface {
	GetVsanDatastoreSpec() *VsanDatastoreSpec
}

func (b *VsanNetworkConfigBaseIssue) GetVsanNetworkConfigBaseIssue() *VsanNetworkConfigBaseIssue {
	return b
}

type BaseVsanNetworkConfigBaseIssue interface {
	GetVsanNetworkConfigBaseIssue() *VsanNetworkConfigBaseIssue
}

func (b *VsanIscsiTargetCommonInfo) GetVsanIscsiTargetCommonInfo() *VsanIscsiTargetCommonInfo {
	return b
}

type BaseVsanIscsiTargetCommonInfo interface {
	GetVsanIscsiTargetCommonInfo() *VsanIscsiTargetCommonInfo
}

func (b *VsanClusterHealthResultBase) GetVsanClusterHealthResultBase() *VsanClusterHealthResultBase {
	return b
}

type BaseVsanClusterHealthResultBase interface {
	GetVsanClusterHealthResultBase() *VsanClusterHealthResultBase
}

func (b *VsanHclCommonDeviceInfo) GetVsanHclCommonDeviceInfo() *VsanHclCommonDeviceInfo { return b }

type BaseVsanHclCommonDeviceInfo interface {
	GetVsanHclCommonDeviceInfo() *VsanHclCommonDeviceInfo
}

func (b *VsanComparator) GetVsanComparator() *VsanComparator { return b }

type BaseVsanComparator interface {
	GetVsanComparator() *VsanComparator
}

func (b *VsanIscsiTargetServiceConfig) GetVsanIscsiTargetServiceConfig() *VsanIscsiTargetServiceConfig {
	return b
}

type BaseVsanIscsiTargetServiceConfig interface {
	GetVsanIscsiTargetServiceConfig() *VsanIscsiTargetServiceConfig
}
