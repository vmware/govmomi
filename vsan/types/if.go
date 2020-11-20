/*
Copyright (c) 2014-2020 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package types

import (
	"reflect"

	"github.com/vmware/govmomi/vim25/types"
)

type (
	BaseCannotMoveVsanEnabledHost            types.BaseCannotMoveVsanEnabledHost
	BaseVsanDiskFault                        types.BaseVsanDiskFault
	BaseVsanFault                            types.BaseVsanFault
	BaseVsanUpgradeSystemPreflightCheckIssue types.BaseVsanUpgradeSystemPreflightCheckIssue
	BaseVsanUpgradeSystemUpgradeHistoryItem  types.BaseVsanUpgradeSystemUpgradeHistoryItem
)

func (b *CnsBackingObjectDetails) GetCnsBackingObjectDetails() *CnsBackingObjectDetails { return b }

type BaseCnsBackingObjectDetails interface {
	GetCnsBackingObjectDetails() *CnsBackingObjectDetails
}

func init() {
	types.Add("BaseCnsBackingObjectDetails", reflect.TypeOf((*CnsBackingObjectDetails)(nil)).Elem())
}

func (b *CnsBaseCreateSpec) GetCnsBaseCreateSpec() *CnsBaseCreateSpec { return b }

type BaseCnsBaseCreateSpec interface {
	GetCnsBaseCreateSpec() *CnsBaseCreateSpec
}

func init() {
	types.Add("BaseCnsBaseCreateSpec", reflect.TypeOf((*CnsBaseCreateSpec)(nil)).Elem())
}

func (b *CnsEntityMetadata) GetCnsEntityMetadata() *CnsEntityMetadata { return b }

type BaseCnsEntityMetadata interface {
	GetCnsEntityMetadata() *CnsEntityMetadata
}

func init() {
	types.Add("BaseCnsEntityMetadata", reflect.TypeOf((*CnsEntityMetadata)(nil)).Elem())
}

func (b *CnsFileBackingDetails) GetCnsFileBackingDetails() *CnsFileBackingDetails { return b }

type BaseCnsFileBackingDetails interface {
	GetCnsFileBackingDetails() *CnsFileBackingDetails
}

func init() {
	types.Add("BaseCnsFileBackingDetails", reflect.TypeOf((*CnsFileBackingDetails)(nil)).Elem())
}

func (b *CnsFileCreateSpec) GetCnsFileCreateSpec() *CnsFileCreateSpec { return b }

type BaseCnsFileCreateSpec interface {
	GetCnsFileCreateSpec() *CnsFileCreateSpec
}

func init() {
	types.Add("BaseCnsFileCreateSpec", reflect.TypeOf((*CnsFileCreateSpec)(nil)).Elem())
}

func (b *CnsQueryFilter) GetCnsQueryFilter() *CnsQueryFilter { return b }

type BaseCnsQueryFilter interface {
	GetCnsQueryFilter() *CnsQueryFilter
}

func init() {
	types.Add("BaseCnsQueryFilter", reflect.TypeOf((*CnsQueryFilter)(nil)).Elem())
}

func (b *CnsVolumeOperationResult) GetCnsVolumeOperationResult() *CnsVolumeOperationResult { return b }

type BaseCnsVolumeOperationResult interface {
	GetCnsVolumeOperationResult() *CnsVolumeOperationResult
}

func init() {
	types.Add("BaseCnsVolumeOperationResult", reflect.TypeOf((*CnsVolumeOperationResult)(nil)).Elem())
}

func (b *CnsVolumeSource) GetCnsVolumeSource() *CnsVolumeSource { return b }

type BaseCnsVolumeSource interface {
	GetCnsVolumeSource() *CnsVolumeSource
}

func init() {
	types.Add("BaseCnsVolumeSource", reflect.TypeOf((*CnsVolumeSource)(nil)).Elem())
}

func (b *EntityResourceCheckDetails) GetEntityResourceCheckDetails() *EntityResourceCheckDetails {
	return b
}

type BaseEntityResourceCheckDetails interface {
	GetEntityResourceCheckDetails() *EntityResourceCheckDetails
}

func init() {
	types.Add("BaseEntityResourceCheckDetails", reflect.TypeOf((*EntityResourceCheckDetails)(nil)).Elem())
}

func (b *VsanClusterHealthResultBase) GetVsanClusterHealthResultBase() *VsanClusterHealthResultBase {
	return b
}

type BaseVsanClusterHealthResultBase interface {
	GetVsanClusterHealthResultBase() *VsanClusterHealthResultBase
}

func init() {
	types.Add("BaseVsanClusterHealthResultBase", reflect.TypeOf((*VsanClusterHealthResultBase)(nil)).Elem())
}

func (b *VsanClusterConfigInfo) GetVsanClusterConfigInfo() *VsanClusterConfigInfo { return b }

type BaseVsanClusterConfigInfo interface {
	GetVsanClusterConfigInfo() *VsanClusterConfigInfo
}

func init() {
	types.Add("BaseVsanClusterConfigInfo", reflect.TypeOf((*VsanClusterConfigInfo)(nil)).Elem())
}

func (b *VsanComparator) GetVsanComparator() *VsanComparator { return b }

type BaseVsanComparator interface {
	GetVsanComparator() *VsanComparator
}

func init() {
	types.Add("BaseVsanComparator", reflect.TypeOf((*VsanComparator)(nil)).Elem())
}

func (b *VsanConfigBaseIssue) GetVsanConfigBaseIssue() *VsanConfigBaseIssue { return b }

type BaseVsanConfigBaseIssue interface {
	GetVsanConfigBaseIssue() *VsanConfigBaseIssue
}

func init() {
	types.Add("BaseVsanConfigBaseIssue", reflect.TypeOf((*VsanConfigBaseIssue)(nil)).Elem())
}

func (b *VsanHclCommonDeviceInfo) GetVsanHclCommonDeviceInfo() *VsanHclCommonDeviceInfo { return b }

type BaseVsanHclCommonDeviceInfo interface {
	GetVsanHclCommonDeviceInfo() *VsanHclCommonDeviceInfo
}

func init() {
	types.Add("BaseVsanHclCommonDeviceInfo", reflect.TypeOf((*VsanHclCommonDeviceInfo)(nil)).Elem())
}

func (b *VsanIscsiLUNCommonInfo) GetVsanIscsiLUNCommonInfo() *VsanIscsiLUNCommonInfo { return b }

type BaseVsanIscsiLUNCommonInfo interface {
	GetVsanIscsiLUNCommonInfo() *VsanIscsiLUNCommonInfo
}

func init() {
	types.Add("BaseVsanIscsiLUNCommonInfo", reflect.TypeOf((*VsanIscsiLUNCommonInfo)(nil)).Elem())
}

func (b *VsanIscsiTargetBasicInfo) GetVsanIscsiTargetBasicInfo() *VsanIscsiTargetBasicInfo { return b }

type BaseVsanIscsiTargetBasicInfo interface {
	GetVsanIscsiTargetBasicInfo() *VsanIscsiTargetBasicInfo
}

func init() {
	types.Add("BaseVsanIscsiTargetBasicInfo", reflect.TypeOf((*VsanIscsiTargetBasicInfo)(nil)).Elem())
}

func (b *VsanIscsiTargetCommonInfo) GetVsanIscsiTargetCommonInfo() *VsanIscsiTargetCommonInfo {
	return b
}

type BaseVsanIscsiTargetCommonInfo interface {
	GetVsanIscsiTargetCommonInfo() *VsanIscsiTargetCommonInfo
}

func init() {
	types.Add("BaseVsanIscsiTargetCommonInfo", reflect.TypeOf((*VsanIscsiTargetCommonInfo)(nil)).Elem())
}

func (b *VsanIscsiTargetServiceConfig) GetVsanIscsiTargetServiceConfig() *VsanIscsiTargetServiceConfig {
	return b
}

type BaseVsanIscsiTargetServiceConfig interface {
	GetVsanIscsiTargetServiceConfig() *VsanIscsiTargetServiceConfig
}

func init() {
	types.Add("BaseVsanIscsiTargetServiceConfig", reflect.TypeOf((*VsanIscsiTargetServiceConfig)(nil)).Elem())
}

func (b *VsanNetworkConfigBaseIssue) GetVsanNetworkConfigBaseIssue() *VsanNetworkConfigBaseIssue {
	return b
}

type BaseVsanNetworkConfigBaseIssue interface {
	GetVsanNetworkConfigBaseIssue() *VsanNetworkConfigBaseIssue
}

func init() {
	types.Add("BaseVsanNetworkConfigBaseIssue", reflect.TypeOf((*VsanNetworkConfigBaseIssue)(nil)).Elem())
}

func (b *VsanResourceConstraint) GetVsanResourceConstraint() *VsanResourceConstraint { return b }

type BaseVsanResourceConstraint interface {
	GetVsanResourceConstraint() *VsanResourceConstraint
}

func init() {
	types.Add("BaseVsanResourceConstraint", reflect.TypeOf((*VsanResourceConstraint)(nil)).Elem())
}
