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
