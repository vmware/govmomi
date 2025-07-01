// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

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
	CryptoSpec types.BaseCryptoSpec `xml:"cryptoSpec,omitempty,typeattr" json:"cryptoSpec"`
}

func init() {
	types.Add("CnsBlockCreateSpec", reflect.TypeOf((*CnsBlockCreateSpec)(nil)).Elem())
}

type CnsUpdateVolumeCryptoRequestType struct {
	This        types.ManagedObjectReference `xml:"_this" json:"-"`
	UpdateSpecs []CnsVolumeCryptoUpdateSpec  `xml:"updateSpecs,omitempty" json:"updateSpecs"`
}

func init() {
	types.Add("CnsUpdateVolumeCryptoRequestType", reflect.TypeOf((*CnsUpdateVolumeCryptoRequestType)(nil)).Elem())
}

type CnsUpdateVolumeCrypto CnsUpdateVolumeCryptoRequestType

func init() {
	types.Add("CnsUpdateVolumeCrypto", reflect.TypeOf((*CnsUpdateVolumeCrypto)(nil)).Elem())
}

type CnsUpdateVolumeCryptoResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}

// CnsVolumeCryptoUpdateSpec is the specification for volume crypto update operation.
type CnsVolumeCryptoUpdateSpec struct {
	types.DynamicData

	VolumeId    CnsVolumeId                           `xml:"volumeId" json:"volumeId"`
	Profile     []types.BaseVirtualMachineProfileSpec `xml:"profile,omitempty,typeattr" json:"profile"`
	DisksCrypto *types.DiskCryptoSpec                 `xml:"disksCrypto,omitempty" json:"disksCrypto"`
}

func init() {
	types.Add("CnsVolumeCryptoUpdateSpec", reflect.TypeOf((*CnsVolumeCryptoUpdateSpec)(nil)).Elem())
}
