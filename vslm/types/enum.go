// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"reflect"

	"github.com/vmware/govmomi/vim25/types"
)

type ManagedObjectType string

const (
	ManagedObjectTypeVslmServiceInstance         = ManagedObjectType("VslmServiceInstance")
	ManagedObjectTypeVslmStorageLifecycleManager = ManagedObjectType("VslmStorageLifecycleManager")
	ManagedObjectTypeVslmTask                    = ManagedObjectType("VslmTask")
	ManagedObjectTypeVslmSessionManager          = ManagedObjectType("VslmSessionManager")
	ManagedObjectTypeVslmVStorageObjectManager   = ManagedObjectType("VslmVStorageObjectManager")
)

func (e ManagedObjectType) Values() []ManagedObjectType {
	return []ManagedObjectType{
		ManagedObjectTypeVslmServiceInstance,
		ManagedObjectTypeVslmStorageLifecycleManager,
		ManagedObjectTypeVslmTask,
		ManagedObjectTypeVslmSessionManager,
		ManagedObjectTypeVslmVStorageObjectManager,
	}
}

func (e ManagedObjectType) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

func init() {
	types.Add("vslm:ManagedObjectType", reflect.TypeOf((*ManagedObjectType)(nil)).Elem())
}

type VslmEventType string

const (
	// Event type used to notify that FCD is going to be relocated.
	VslmEventTypePreFcdMigrateEvent = VslmEventType("preFcdMigrateEvent")
	// Event type used to notify FCD has been relocated.
	VslmEventTypePostFcdMigrateEvent = VslmEventType("postFcdMigrateEvent")
)

func (e VslmEventType) Values() []VslmEventType {
	return []VslmEventType{
		VslmEventTypePreFcdMigrateEvent,
		VslmEventTypePostFcdMigrateEvent,
	}
}

func (e VslmEventType) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

func init() {
	types.Add("vslm:VslmEventType", reflect.TypeOf((*VslmEventType)(nil)).Elem())
}

// The possible states of the vlsm event processing.
type VslmEventVslmEventInfoState string

const (
	// When the event has been successfully processed.
	VslmEventVslmEventInfoStateSuccess = VslmEventVslmEventInfoState("success")
	// When there is error while processing the event.
	VslmEventVslmEventInfoStateError = VslmEventVslmEventInfoState("error")
)

func (e VslmEventVslmEventInfoState) Values() []VslmEventVslmEventInfoState {
	return []VslmEventVslmEventInfoState{
		VslmEventVslmEventInfoStateSuccess,
		VslmEventVslmEventInfoStateError,
	}
}

func (e VslmEventVslmEventInfoState) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

func init() {
	types.Add("vslm:VslmEventVslmEventInfoState", reflect.TypeOf((*VslmEventVslmEventInfoState)(nil)).Elem())
}

// List of possible states of a task.
type VslmTaskInfoState string

const (
	// When there are too many tasks for threads to handle.
	VslmTaskInfoStateQueued = VslmTaskInfoState("queued")
	// When the busy thread is freed from its current task by
	// finishing the task, it picks a queued task to run.
	//
	// Then the queued tasks are marked as running.
	VslmTaskInfoStateRunning = VslmTaskInfoState("running")
	// When a running task has completed.
	VslmTaskInfoStateSuccess = VslmTaskInfoState("success")
	// When a running task has encountered an error.
	VslmTaskInfoStateError = VslmTaskInfoState("error")
)

func (e VslmTaskInfoState) Values() []VslmTaskInfoState {
	return []VslmTaskInfoState{
		VslmTaskInfoStateQueued,
		VslmTaskInfoStateRunning,
		VslmTaskInfoStateSuccess,
		VslmTaskInfoStateError,
	}
}

func (e VslmTaskInfoState) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

func init() {
	types.Add("vslm:VslmTaskInfoState", reflect.TypeOf((*VslmTaskInfoState)(nil)).Elem())
}

// The `VslmVsoVStorageObjectQuerySpecQueryFieldEnum_enum` enumerated
// type defines the searchable fields.
type VslmVsoVStorageObjectQuerySpecQueryFieldEnum string

const (
	// Indicates `BaseConfigInfo.id` as the searchable field.
	VslmVsoVStorageObjectQuerySpecQueryFieldEnumId = VslmVsoVStorageObjectQuerySpecQueryFieldEnum("id")
	// Indicates `BaseConfigInfo.name` as the searchable
	// field.
	VslmVsoVStorageObjectQuerySpecQueryFieldEnumName = VslmVsoVStorageObjectQuerySpecQueryFieldEnum("name")
	// Indicates `VStorageObjectConfigInfo.capacityInMB` as the
	// searchable field.
	VslmVsoVStorageObjectQuerySpecQueryFieldEnumCapacity = VslmVsoVStorageObjectQuerySpecQueryFieldEnum("capacity")
	// Indicates `BaseConfigInfo.createTime` as the searchable
	// field.
	VslmVsoVStorageObjectQuerySpecQueryFieldEnumCreateTime = VslmVsoVStorageObjectQuerySpecQueryFieldEnum("createTime")
	// Indicates
	// `BaseConfigInfoFileBackingInfo.backingObjectId` as the
	// searchable field.
	VslmVsoVStorageObjectQuerySpecQueryFieldEnumBackingObjectId = VslmVsoVStorageObjectQuerySpecQueryFieldEnum("backingObjectId")
	// Indicates `BaseConfigInfoBackingInfo.datastore` as the
	// searchable field.
	VslmVsoVStorageObjectQuerySpecQueryFieldEnumDatastoreMoId = VslmVsoVStorageObjectQuerySpecQueryFieldEnum("datastoreMoId")
	// Indicates it as the searchable field.
	VslmVsoVStorageObjectQuerySpecQueryFieldEnumMetadataKey = VslmVsoVStorageObjectQuerySpecQueryFieldEnum("metadataKey")
	// Indicates it as the searchable field.
	VslmVsoVStorageObjectQuerySpecQueryFieldEnumMetadataValue = VslmVsoVStorageObjectQuerySpecQueryFieldEnum("metadataValue")
)

func (e VslmVsoVStorageObjectQuerySpecQueryFieldEnum) Values() []VslmVsoVStorageObjectQuerySpecQueryFieldEnum {
	return []VslmVsoVStorageObjectQuerySpecQueryFieldEnum{
		VslmVsoVStorageObjectQuerySpecQueryFieldEnumId,
		VslmVsoVStorageObjectQuerySpecQueryFieldEnumName,
		VslmVsoVStorageObjectQuerySpecQueryFieldEnumCapacity,
		VslmVsoVStorageObjectQuerySpecQueryFieldEnumCreateTime,
		VslmVsoVStorageObjectQuerySpecQueryFieldEnumBackingObjectId,
		VslmVsoVStorageObjectQuerySpecQueryFieldEnumDatastoreMoId,
		VslmVsoVStorageObjectQuerySpecQueryFieldEnumMetadataKey,
		VslmVsoVStorageObjectQuerySpecQueryFieldEnumMetadataValue,
	}
}

func (e VslmVsoVStorageObjectQuerySpecQueryFieldEnum) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

func init() {
	types.Add("vslm:VslmVsoVStorageObjectQuerySpecQueryFieldEnum", reflect.TypeOf((*VslmVsoVStorageObjectQuerySpecQueryFieldEnum)(nil)).Elem())
}

// The `VslmVsoVStorageObjectQuerySpecQueryOperatorEnum_enum` enumerated
// type defines the operators to use for constructing the query criteria.
type VslmVsoVStorageObjectQuerySpecQueryOperatorEnum string

const (
	VslmVsoVStorageObjectQuerySpecQueryOperatorEnumEquals             = VslmVsoVStorageObjectQuerySpecQueryOperatorEnum("equals")
	VslmVsoVStorageObjectQuerySpecQueryOperatorEnumNotEquals          = VslmVsoVStorageObjectQuerySpecQueryOperatorEnum("notEquals")
	VslmVsoVStorageObjectQuerySpecQueryOperatorEnumLessThan           = VslmVsoVStorageObjectQuerySpecQueryOperatorEnum("lessThan")
	VslmVsoVStorageObjectQuerySpecQueryOperatorEnumGreaterThan        = VslmVsoVStorageObjectQuerySpecQueryOperatorEnum("greaterThan")
	VslmVsoVStorageObjectQuerySpecQueryOperatorEnumLessThanOrEqual    = VslmVsoVStorageObjectQuerySpecQueryOperatorEnum("lessThanOrEqual")
	VslmVsoVStorageObjectQuerySpecQueryOperatorEnumGreaterThanOrEqual = VslmVsoVStorageObjectQuerySpecQueryOperatorEnum("greaterThanOrEqual")
	VslmVsoVStorageObjectQuerySpecQueryOperatorEnumContains           = VslmVsoVStorageObjectQuerySpecQueryOperatorEnum("contains")
	VslmVsoVStorageObjectQuerySpecQueryOperatorEnumStartsWith         = VslmVsoVStorageObjectQuerySpecQueryOperatorEnum("startsWith")
	VslmVsoVStorageObjectQuerySpecQueryOperatorEnumEndsWith           = VslmVsoVStorageObjectQuerySpecQueryOperatorEnum("endsWith")
)

func (e VslmVsoVStorageObjectQuerySpecQueryOperatorEnum) Values() []VslmVsoVStorageObjectQuerySpecQueryOperatorEnum {
	return []VslmVsoVStorageObjectQuerySpecQueryOperatorEnum{
		VslmVsoVStorageObjectQuerySpecQueryOperatorEnumEquals,
		VslmVsoVStorageObjectQuerySpecQueryOperatorEnumNotEquals,
		VslmVsoVStorageObjectQuerySpecQueryOperatorEnumLessThan,
		VslmVsoVStorageObjectQuerySpecQueryOperatorEnumGreaterThan,
		VslmVsoVStorageObjectQuerySpecQueryOperatorEnumLessThanOrEqual,
		VslmVsoVStorageObjectQuerySpecQueryOperatorEnumGreaterThanOrEqual,
		VslmVsoVStorageObjectQuerySpecQueryOperatorEnumContains,
		VslmVsoVStorageObjectQuerySpecQueryOperatorEnumStartsWith,
		VslmVsoVStorageObjectQuerySpecQueryOperatorEnumEndsWith,
	}
}

func (e VslmVsoVStorageObjectQuerySpecQueryOperatorEnum) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

func init() {
	types.Add("vslm:VslmVsoVStorageObjectQuerySpecQueryOperatorEnum", reflect.TypeOf((*VslmVsoVStorageObjectQuerySpecQueryOperatorEnum)(nil)).Elem())
}
