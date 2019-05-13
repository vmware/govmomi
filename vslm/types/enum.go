/*
Copyright (c) 2014-2018 VMware, Inc. All Rights Reserved.

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

type VslmTaskInfoState string

const (
	VslmTaskInfoStateQueued  = VslmTaskInfoState("queued")
	VslmTaskInfoStateRunning = VslmTaskInfoState("running")
	VslmTaskInfoStateSuccess = VslmTaskInfoState("success")
	VslmTaskInfoStateError   = VslmTaskInfoState("error")
)

func init() {
	types.Add("vslm:VslmTaskInfoState", reflect.TypeOf((*VslmTaskInfoState)(nil)).Elem())
}

type VslmVsoVStorageObjectQuerySpecQueryFieldEnum string

const (
	VslmVsoVStorageObjectQuerySpecQueryFieldEnumId              = VslmVsoVStorageObjectQuerySpecQueryFieldEnum("id")
	VslmVsoVStorageObjectQuerySpecQueryFieldEnumName            = VslmVsoVStorageObjectQuerySpecQueryFieldEnum("name")
	VslmVsoVStorageObjectQuerySpecQueryFieldEnumCapacity        = VslmVsoVStorageObjectQuerySpecQueryFieldEnum("capacity")
	VslmVsoVStorageObjectQuerySpecQueryFieldEnumCreateTime      = VslmVsoVStorageObjectQuerySpecQueryFieldEnum("createTime")
	VslmVsoVStorageObjectQuerySpecQueryFieldEnumBackingObjectId = VslmVsoVStorageObjectQuerySpecQueryFieldEnum("backingObjectId")
	VslmVsoVStorageObjectQuerySpecQueryFieldEnumDatastoreMoId   = VslmVsoVStorageObjectQuerySpecQueryFieldEnum("datastoreMoId")
	VslmVsoVStorageObjectQuerySpecQueryFieldEnumMetadataKey     = VslmVsoVStorageObjectQuerySpecQueryFieldEnum("metadataKey")
	VslmVsoVStorageObjectQuerySpecQueryFieldEnumMetadataValue   = VslmVsoVStorageObjectQuerySpecQueryFieldEnum("metadataValue")
)

func init() {
	types.Add("vslm:VslmVsoVStorageObjectQuerySpecQueryFieldEnum", reflect.TypeOf((*VslmVsoVStorageObjectQuerySpecQueryFieldEnum)(nil)).Elem())
}

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
