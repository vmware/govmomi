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

func (b *CnsAlreadyRegisteredFault) GetCnsAlreadyRegisteredFault() *CnsAlreadyRegisteredFault {
	return b
}

type BaseCnsAlreadyRegisteredFault interface {
	GetCnsAlreadyRegisteredFault() *CnsAlreadyRegisteredFault
}

func init() {
	types.Add("BaseCnsAlreadyRegisteredFault", reflect.TypeOf((*CnsAlreadyRegisteredFault)(nil)).Elem())
}

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

type BaseCnsVolumeRelocateSpec interface {
	GetCnsVolumeRelocateSpec() CnsVolumeRelocateSpec
}

func (s CnsVolumeRelocateSpec) GetCnsVolumeRelocateSpec() CnsVolumeRelocateSpec { return s }

func init() {
	types.Add("BaseCnsVolumeRelocateSpec", reflect.TypeOf((*CnsVolumeRelocateSpec)(nil)).Elem())
}

func (b *CnsEntityMetadata) GetCnsEntityMetadata() *CnsEntityMetadata { return b }

type BaseCnsEntityMetadata interface {
	GetCnsEntityMetadata() *CnsEntityMetadata
}

func init() {
	types.Add("BaseCnsEntityMetadata", reflect.TypeOf((*CnsEntityMetadata)(nil)).Elem())
}

func (b *CnsVolumeInfo) GetCnsVolumeInfo() *CnsVolumeInfo { return b }

type BaseCnsVolumeInfo interface {
	GetCnsVolumeInfo() *CnsVolumeInfo
}

func init() {
	types.Add("BaseCnsVolumeInfo", reflect.TypeOf((*CnsVolumeInfo)(nil)).Elem())
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
