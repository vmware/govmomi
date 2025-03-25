// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"reflect"

	"github.com/vmware/govmomi/vim25/types"
)

type ArrayOfPbmIofilterInfo struct {
	PbmIofilterInfo []PbmIofilterInfo `xml:"PbmIofilterInfo,omitempty" json:"_value"`
}

func init() {
	types.Add("pbm:ArrayOfPbmIofilterInfo", reflect.TypeOf((*ArrayOfPbmIofilterInfo)(nil)).Elem())
}

type ArrayOfPbmProfileToIofilterMap struct {
	PbmProfileToIofilterMap []PbmProfileToIofilterMap `xml:"PbmProfileToIofilterMap,omitempty" json:"_value"`
}

func init() {
	types.Add("pbm:ArrayOfPbmProfileToIofilterMap", reflect.TypeOf((*ArrayOfPbmProfileToIofilterMap)(nil)).Elem())
}

type PbmQueryIOFiltersFromProfileId PbmQueryIOFiltersFromProfileIdRequestType

func init() {
	types.Add("pbm:PbmQueryIOFiltersFromProfileId", reflect.TypeOf((*PbmQueryIOFiltersFromProfileId)(nil)).Elem())
}

type PbmQueryIOFiltersFromProfileIdRequestType struct {
	This       types.ManagedObjectReference `xml:"_this" json:"_this"`
	ProfileIds []PbmProfileId               `xml:"profileIds,typeattr" json:"profileIds"`
}

func init() {
	types.Add("pbm:PbmQueryIOFiltersFromProfileIdRequestType", reflect.TypeOf((*PbmQueryIOFiltersFromProfileIdRequestType)(nil)).Elem())
}

type PbmQueryIOFiltersFromProfileIdResponse struct {
	Returnval []PbmProfileToIofilterMap `xml:"returnval" json:"returnval"`
}

type PbmIofilterInfo struct {
	types.DynamicData

	VibId      string `xml:"vibId" json:"vibId"`
	FilterType string `xml:"filterType,omitempty" json:"filterType,omitempty"`
}

func init() {
	types.Add("pbm:PbmIofilterInfo", reflect.TypeOf((*PbmIofilterInfo)(nil)).Elem())
}

type PbmProfileToIofilterMap struct {
	types.DynamicData

	Key       PbmProfileId                `xml:"key,typeattr" json:"key"`
	Iofilters []PbmIofilterInfo           `xml:"iofilters,omitempty" json:"iofilters,omitempty"`
	Fault     *types.LocalizedMethodFault `xml:"fault,omitempty" json:"fault,omitempty"`
}

func init() {
	types.Add("pbm:PbmProfileToIofilterMap", reflect.TypeOf((*PbmProfileToIofilterMap)(nil)).Elem())
}

type PbmProfileDetails struct {
	types.DynamicData

	Profile             BasePbmCapabilityProfile    `xml:"profile,typeattr" json:"profile"`
	IofInfos            []PbmIofilterInfo           `xml:"iofInfos,omitempty" json:"iofInfos,omitempty"`
	IofMethodFault      *types.LocalizedMethodFault `xml:"iofMethodFault,omitempty" json:"iofMethodFault,omitempty"`
	IsReplicationPolicy bool                        `xml:"isReplicationPolicy" json:"isReplicationPolicy"`
	RepMethodFault      *types.LocalizedMethodFault `xml:"repMethodFault,omitempty" json:"repMethodFault,omitempty"`
}

func init() {
	types.Add("pbm:PbmProfileDetails", reflect.TypeOf((*PbmProfileDetails)(nil)).Elem())
}
