// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"reflect"

	"github.com/vmware/govmomi/vim25/types"
)

func (b *DeviceId) GetDeviceId() *DeviceId { return b }

type BaseDeviceId interface {
	GetDeviceId() *DeviceId
}

func init() {
	types.Add("BaseDeviceId", reflect.TypeOf((*DeviceId)(nil)).Elem())
}

func (b *FailoverParam) GetFailoverParam() *FailoverParam { return b }

type BaseFailoverParam interface {
	GetFailoverParam() *FailoverParam
}

func init() {
	types.Add("BaseFailoverParam", reflect.TypeOf((*FailoverParam)(nil)).Elem())
}

func (b *GroupInfo) GetGroupInfo() *GroupInfo { return b }

type BaseGroupInfo interface {
	GetGroupInfo() *GroupInfo
}

func init() {
	types.Add("BaseGroupInfo", reflect.TypeOf((*GroupInfo)(nil)).Elem())
}

func (b *GroupOperationResult) GetGroupOperationResult() *GroupOperationResult { return b }

type BaseGroupOperationResult interface {
	GetGroupOperationResult() *GroupOperationResult
}

func init() {
	types.Add("BaseGroupOperationResult", reflect.TypeOf((*GroupOperationResult)(nil)).Elem())
}

func (b *ProviderRegistrationFault) GetProviderRegistrationFault() *ProviderRegistrationFault {
	return b
}

type BaseProviderRegistrationFault interface {
	GetProviderRegistrationFault() *ProviderRegistrationFault
}

func init() {
	types.Add("BaseProviderRegistrationFault", reflect.TypeOf((*ProviderRegistrationFault)(nil)).Elem())
}

func (b *ProviderSyncFailed) GetProviderSyncFailed() *ProviderSyncFailed { return b }

type BaseProviderSyncFailed interface {
	GetProviderSyncFailed() *ProviderSyncFailed
}

func init() {
	types.Add("BaseProviderSyncFailed", reflect.TypeOf((*ProviderSyncFailed)(nil)).Elem())
}

func (b *QueryExecutionFault) GetQueryExecutionFault() *QueryExecutionFault { return b }

type BaseQueryExecutionFault interface {
	GetQueryExecutionFault() *QueryExecutionFault
}

func init() {
	types.Add("BaseQueryExecutionFault", reflect.TypeOf((*QueryExecutionFault)(nil)).Elem())
}

func (b *SmsProviderInfo) GetSmsProviderInfo() *SmsProviderInfo { return b }

type BaseSmsProviderInfo interface {
	GetSmsProviderInfo() *SmsProviderInfo
}

func init() {
	types.Add("BaseSmsProviderInfo", reflect.TypeOf((*SmsProviderInfo)(nil)).Elem())
}

func (b *SmsProviderSpec) GetSmsProviderSpec() *SmsProviderSpec { return b }

type BaseSmsProviderSpec interface {
	GetSmsProviderSpec() *SmsProviderSpec
}

func init() {
	types.Add("BaseSmsProviderSpec", reflect.TypeOf((*SmsProviderSpec)(nil)).Elem())
}

func (b *SmsReplicationFault) GetSmsReplicationFault() *SmsReplicationFault { return b }

type BaseSmsReplicationFault interface {
	GetSmsReplicationFault() *SmsReplicationFault
}

func init() {
	types.Add("BaseSmsReplicationFault", reflect.TypeOf((*SmsReplicationFault)(nil)).Elem())
}

func (b *StoragePort) GetStoragePort() *StoragePort { return b }

type BaseStoragePort interface {
	GetStoragePort() *StoragePort
}

func init() {
	types.Add("BaseStoragePort", reflect.TypeOf((*StoragePort)(nil)).Elem())
}

func (b *TargetGroupMemberInfo) GetTargetGroupMemberInfo() *TargetGroupMemberInfo { return b }

type BaseTargetGroupMemberInfo interface {
	GetTargetGroupMemberInfo() *TargetGroupMemberInfo
}

func init() {
	types.Add("BaseTargetGroupMemberInfo", reflect.TypeOf((*TargetGroupMemberInfo)(nil)).Elem())
}

func (b *VirtualMachineId) GetVirtualMachineId() *VirtualMachineId { return b }

type BaseVirtualMachineId interface {
	GetVirtualMachineId() *VirtualMachineId
}

func init() {
	types.Add("BaseVirtualMachineId", reflect.TypeOf((*VirtualMachineId)(nil)).Elem())
}
