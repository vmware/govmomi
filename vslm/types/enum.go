/*
Copyright (c) 2014-2023 VMware, Inc. All Rights Reserved.

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

type VslmEventType string

const (
	// Event type used to notify that FCD is going to be relocated.
	VslmEventTypePreFcdMigrateEvent = VslmEventType("preFcdMigrateEvent")
	// Event type used to notify FCD has been relocated.
	VslmEventTypePostFcdMigrateEvent = VslmEventType("postFcdMigrateEvent")
)

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
	// Indicates `vim.vslm.VStorageObject#capacityInMB` as the
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

func init() {
	types.Add("vslm:VslmVsoVStorageObjectQuerySpecQueryOperatorEnum", reflect.TypeOf((*VslmVsoVStorageObjectQuerySpecQueryOperatorEnum)(nil)).Elem())
}
