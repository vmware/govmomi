/*
Copyright (c) 2019 VMware, Inc. All Rights Reserved.

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

func (b *CnsFault) GetCnsFault() *CnsFault {
	return b
}

type BaseCnsFault interface {
	GetCnsFault() *CnsFault
}

func init() {
	types.Add("BaseCnsFault", reflect.TypeOf((*CnsFault)(nil)).Elem())
}

func (b *CnsBackingObjectDetails) GetCnsBackingObjectDetails() *CnsBackingObjectDetails { return b }

type BaseCnsBackingObjectDetails interface {
	GetCnsBackingObjectDetails() *CnsBackingObjectDetails
}

func init() {
	types.Add("BaseCnsBackingObjectDetails", reflect.TypeOf((*CnsBackingObjectDetails)(nil)).Elem())
}

func (b *CnsEntityMetadata) GetCnsEntityMetadata() *CnsEntityMetadata { return b }

type BaseCnsEntityMetadata interface {
	GetCnsEntityMetadata() *CnsEntityMetadata
}

func init() {
	types.Add("BaseCnsEntityMetadata", reflect.TypeOf((*CnsEntityMetadata)(nil)).Elem())
}

func (b *CnsVolumeOperationResult) GetCnsVolumeOperationResult() *CnsVolumeOperationResult { return b }

type BaseCnsVolumeOperationResult interface {
	GetCnsVolumeOperationResult() *CnsVolumeOperationResult
}

func init() {
	types.Add("BaseCnsVolumeOperationResult", reflect.TypeOf((*CnsVolumeOperationResult)(nil)).Elem())
}
