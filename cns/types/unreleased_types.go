/*
Copyright (c) 2024-2024 VMware, Inc. All Rights Reserved.

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

// CnsBlockCreateSpec is the specification for creating block volumes.
type CnsBlockCreateSpec struct {
	CnsBaseCreateSpec

	// Crypto specifies the encryption settings for the volume to be created.
	// Works with block volumes only.
	CryptoSpec types.BaseCryptoSpec `xml:"cryptoSpec,omitempty,typeattr"`
}

func init() {
	types.Add("CnsBlockCreateSpec", reflect.TypeOf((*CnsBlockCreateSpec)(nil)).Elem())
}

type CnsUpdateVolumeCryptoRequestType struct {
	This        types.ManagedObjectReference `xml:"_this"`
	UpdateSpecs []CnsVolumeCryptoUpdateSpec  `xml:"updateSpecs,omitempty"`
}

func init() {
	types.Add("CnsUpdateVolumeCryptoRequestType", reflect.TypeOf((*CnsUpdateVolumeCryptoRequestType)(nil)).Elem())
}

type CnsUpdateVolumeCrypto CnsUpdateVolumeCryptoRequestType

func init() {
	types.Add("CnsUpdateVolumeCrypto", reflect.TypeOf((*CnsUpdateVolumeCrypto)(nil)).Elem())
}

type CnsUpdateVolumeCryptoResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

// CnsVolumeCryptoUpdateSpec is the specification for volume crypto update operation.
type CnsVolumeCryptoUpdateSpec struct {
	types.DynamicData

	VolumeId    CnsVolumeId                           `xml:"volumeId"`
	Profile     []types.BaseVirtualMachineProfileSpec `xml:"profile,omitempty,typeattr"`
	DisksCrypto *types.DiskCryptoSpec                 `xml:"disksCrypto,omitempty"`
}

func init() {
	types.Add("CnsVolumeCryptoUpdateSpec", reflect.TypeOf((*CnsVolumeCryptoUpdateSpec)(nil)).Elem())
}
