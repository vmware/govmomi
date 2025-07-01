// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
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

type ArrayOfPbmProfileDetails struct {
	PbmProfileDetails []PbmProfileDetails `xml:"PbmProfileDetails,omitempty" json:"_value"`
}

func init() {
	types.Add("pbm:ArrayOfPbmProfileDetails", reflect.TypeOf((*ArrayOfPbmProfileDetails)(nil)).Elem())
}

type PbmQueryProfileDetails PbmQueryProfileDetailsRequestType

func init() {
	types.Add("pbm:PbmQueryProfileDetails", reflect.TypeOf((*PbmQueryProfileDetails)(nil)).Elem())
}

type PbmQueryProfileDetailsRequestType struct {
	This            types.ManagedObjectReference `xml:"_this" json:"_this"`
	ProfileCategory string                       `xml:"profileCategory" json:"profileCategory"`
	FetchAllFields  bool                         `xml:"fetchAllFields" json:"fetchAllFields"`
}

func init() {
	types.Add("pbm:PbmQueryProfileDetailsRequestType", reflect.TypeOf((*PbmQueryProfileDetailsRequestType)(nil)).Elem())
}

type PbmQueryProfileDetailsResponse struct {
	Returnval []PbmProfileDetails `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

type PbmResolveK8sCompliantNames PbmResolveK8sCompliantNamesRequestType

func init() {
	types.Add("pbm:PbmResolveK8sCompliantNames", reflect.TypeOf((*PbmResolveK8sCompliantNames)(nil)).Elem())
}

type PbmResolveK8sCompliantNamesRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
}

func init() {
	types.Add("pbm:PbmResolveK8sCompliantNamesRequestType", reflect.TypeOf((*PbmResolveK8sCompliantNamesRequestType)(nil)).Elem())
}

type PbmResolveK8sCompliantNamesResponse struct {
}

type PbmUpdateK8sCompliantNames PbmUpdateK8sCompliantNamesRequestType

func init() {
	types.Add("pbm:PbmUpdateK8sCompliantNames", reflect.TypeOf((*PbmUpdateK8sCompliantNames)(nil)).Elem())
}

type PbmUpdateK8sCompliantNamesRequestType struct {
	This                 types.ManagedObjectReference   `xml:"_this" json:"_this"`
	K8sCompliantNameSpec PbmProfileK8sCompliantNameSpec `xml:"k8sCompliantNameSpec" json:"k8sCompliantNameSpec"`
}

func init() {
	types.Add("pbm:PbmUpdateK8sCompliantNamesRequestType", reflect.TypeOf((*PbmUpdateK8sCompliantNamesRequestType)(nil)).Elem())
}

type PbmUpdateK8sCompliantNamesResponse struct {
}

type PbmValidateProfileK8sCompliantName PbmValidateProfileK8sCompliantNameRequestType

func init() {
	types.Add("pbm:PbmValidateProfileK8sCompliantName", reflect.TypeOf((*PbmValidateProfileK8sCompliantName)(nil)).Elem())
}

type PbmValidateProfileK8sCompliantNameRequestType struct {
	This             types.ManagedObjectReference `xml:"_this" json:"_this"`
	K8sCompliantName string                       `xml:"k8sCompliantName" json:"k8sCompliantName"`
}

func init() {
	types.Add("pbm:PbmValidateProfileK8sCompliantNameRequestType", reflect.TypeOf((*PbmValidateProfileK8sCompliantNameRequestType)(nil)).Elem())
}

type PbmK8sCompliantNameValidationResult struct {
	types.DynamicData

	K8sCompliantName string                      `xml:"k8sCompliantName" json:"k8sCompliantName"`
	Valid            bool                        `xml:"valid" json:"valid"`
	Fault            *types.LocalizedMethodFault `xml:"fault,omitempty" json:"fault,omitempty"`
}

func init() {
	types.Add("pbm:PbmK8sCompliantNameValidationResult", reflect.TypeOf((*PbmK8sCompliantNameValidationResult)(nil)).Elem())
}

type PbmValidateProfileK8sCompliantNameResponse struct {
	Returnval PbmK8sCompliantNameValidationResult `xml:"returnval" json:"returnval"`
}
