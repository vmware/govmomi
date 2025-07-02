// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"reflect"

	"github.com/vmware/govmomi/vim25/types"
)

type VsanPerformanceManager struct {
	Self types.ManagedObjectReference
}

func (m VsanPerformanceManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	types.Add("vsan:VsanPerformanceManager", reflect.TypeOf((*VsanPerformanceManager)(nil)).Elem())
}

type VimClusterVsanVcDiskManagementSystem struct {
	Self types.ManagedObjectReference
}

func (m VimClusterVsanVcDiskManagementSystem) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	types.Add("vsan:VimClusterVsanVcDiskManagementSystem", reflect.TypeOf((*VimClusterVsanVcDiskManagementSystem)(nil)).Elem())
}

type VsanHostVdsSystem struct {
	Self types.ManagedObjectReference
}

func (m VsanHostVdsSystem) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	types.Add("vsan:VsanHostVdsSystem", reflect.TypeOf((*VsanHostVdsSystem)(nil)).Elem())
}

type VsanPhoneHomeSystem struct {
	Self types.ManagedObjectReference
}

func (m VsanPhoneHomeSystem) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	types.Add("vsan:VsanPhoneHomeSystem", reflect.TypeOf((*VsanPhoneHomeSystem)(nil)).Elem())
}

type VsanIscsiTargetSystem struct {
	Self types.ManagedObjectReference
}

func (m VsanIscsiTargetSystem) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	types.Add("vsan:VsanIscsiTargetSystem", reflect.TypeOf((*VsanIscsiTargetSystem)(nil)).Elem())
}

type VsanClusterMgmtInternalSystem struct {
	Self types.ManagedObjectReference
}

func (m VsanClusterMgmtInternalSystem) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	types.Add("vsan:VsanClusterMgmtInternalSystem", reflect.TypeOf((*VsanClusterMgmtInternalSystem)(nil)).Elem())
}

type VsanCapabilitySystem struct {
	Self types.ManagedObjectReference
}

func (m VsanCapabilitySystem) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	types.Add("vsan:VsanCapabilitySystem", reflect.TypeOf((*VsanCapabilitySystem)(nil)).Elem())
}

type VsanVumSystem struct {
	Self types.ManagedObjectReference
}

func (m VsanVumSystem) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	types.Add("vsan:VsanVumSystem", reflect.TypeOf((*VsanVumSystem)(nil)).Elem())
}

type VsanResourceCheckSystem struct {
	Self types.ManagedObjectReference
}

func (m VsanResourceCheckSystem) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	types.Add("vsan:VsanResourceCheckSystem", reflect.TypeOf((*VsanResourceCheckSystem)(nil)).Elem())
}

type VsanObjectSystem struct {
	Self types.ManagedObjectReference
}

func (m VsanObjectSystem) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	types.Add("vsan:VsanObjectSystem", reflect.TypeOf((*VsanObjectSystem)(nil)).Elem())
}

type VsanVcClusterConfigSystem struct {
	Self types.ManagedObjectReference
}

func (m VsanVcClusterConfigSystem) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	types.Add("vsan:VsanVcClusterConfigSystem", reflect.TypeOf((*VsanVcClusterConfigSystem)(nil)).Elem())
}

type VsanVcClusterHealthSystem struct {
	Self types.ManagedObjectReference
}

func (m VsanVcClusterHealthSystem) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	types.Add("vsan:VsanVcClusterHealthSystem", reflect.TypeOf((*VsanVcClusterHealthSystem)(nil)).Elem())
}

type VsanVcsaDeployerSystem struct {
	Self types.ManagedObjectReference
}

func (m VsanVcsaDeployerSystem) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	types.Add("vsan:VsanVcsaDeployerSystem", reflect.TypeOf((*VsanVcsaDeployerSystem)(nil)).Elem())
}

type VsanVdsSystem struct {
	Self types.ManagedObjectReference
}

func (m VsanVdsSystem) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	types.Add("vsan:VsanVdsSystem", reflect.TypeOf((*VsanVdsSystem)(nil)).Elem())
}

type VsanRemoteDatastoreSystem struct {
	Self types.ManagedObjectReference
}

func (m VsanRemoteDatastoreSystem) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	types.Add("vsan:VsanRemoteDatastoreSystem", reflect.TypeOf((*VsanRemoteDatastoreSystem)(nil)).Elem())
}

type VsanUpgradeSystemEx struct {
	Self types.ManagedObjectReference
}

func (m VsanUpgradeSystemEx) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	types.Add("vsan:VsanUpgradeSystemEx", reflect.TypeOf((*VsanUpgradeSystemEx)(nil)).Elem())
}

type VsanSpaceReportSystem struct {
	Self types.ManagedObjectReference
}

func (m VsanSpaceReportSystem) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	types.Add("vsan:VsanSpaceReportSystem", reflect.TypeOf((*VsanSpaceReportSystem)(nil)).Elem())
}

type VsanIoInsightManager struct {
	Self types.ManagedObjectReference
}

func (m VsanIoInsightManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	types.Add("vsan:VsanIoInsightManager", reflect.TypeOf((*VsanIoInsightManager)(nil)).Elem())
}

type VsanUpdateManager struct {
	Self types.ManagedObjectReference
}

func (m VsanUpdateManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	types.Add("vsan:VsanUpdateManager", reflect.TypeOf((*VsanUpdateManager)(nil)).Elem())
}

type VsanSystemEx struct {
	Self types.ManagedObjectReference
}

func (m VsanSystemEx) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	types.Add("vsan:VsanSystemEx", reflect.TypeOf((*VsanSystemEx)(nil)).Elem())
}

type VsanFileServiceSystem struct {
	Self types.ManagedObjectReference
}

func (m VsanFileServiceSystem) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	types.Add("vsan:VsanFileServiceSystem", reflect.TypeOf((*VsanFileServiceSystem)(nil)).Elem())
}

type HostVsanHealthSystem struct {
	Self types.ManagedObjectReference
}

func (m HostVsanHealthSystem) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	types.Add("vsan:HostVsanHealthSystem", reflect.TypeOf((*HostVsanHealthSystem)(nil)).Elem())
}

type HostSpbm struct {
	Self types.ManagedObjectReference
}

func (m HostSpbm) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	types.Add("vsan:HostSpbm", reflect.TypeOf((*HostSpbm)(nil)).Elem())
}

type VimClusterVsanVcStretchedClusterSystem struct {
	Self types.ManagedObjectReference
}

func (m VimClusterVsanVcStretchedClusterSystem) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	types.Add("vsan:VimClusterVsanVcStretchedClusterSystem", reflect.TypeOf((*VimClusterVsanVcStretchedClusterSystem)(nil)).Elem())
}

type VsanClusterHealthSystem struct {
	Self types.ManagedObjectReference
}

func (m VsanClusterHealthSystem) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	types.Add("vsan:VsanClusterHealthSystem", reflect.TypeOf((*VsanClusterHealthSystem)(nil)).Elem())
}

type VsanMassCollector struct {
	Self types.ManagedObjectReference
}

func (m VsanMassCollector) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	types.Add("vsan:VsanMassCollector", reflect.TypeOf((*VsanMassCollector)(nil)).Elem())
}
