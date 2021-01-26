/*
Copyright (c) 2020 VMware, Inc. All Rights Reserved.
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
