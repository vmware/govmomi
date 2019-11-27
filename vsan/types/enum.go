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

type VsanIscsiTargetAuthType string

const (
	VsanIscsiTargetAuthTypeNoAuth                          = VsanIscsiTargetAuthType("NoAuth")
	VsanIscsiTargetAuthTypeCHAP                            = VsanIscsiTargetAuthType("CHAP")
	VsanIscsiTargetAuthTypeCHAP_Mutual                     = VsanIscsiTargetAuthType("CHAP_Mutual")
	VsanIscsiTargetAuthTypeVsanIscsiTargetAuthType_Unknown = VsanIscsiTargetAuthType("AuthType_Unknown")
)

func init() {
	types.Add("VsanIscsiTargetAuthType", reflect.TypeOf((*VsanIscsiTargetAuthType)(nil)).Elem())
}

type VsanBaselinePreferenceType string

const (
	VsanBaselinePreferenceTypelatestRelease                      = VsanBaselinePreferenceType("latestRelease")
	VsanBaselinePreferenceTypelatestPatch                        = VsanBaselinePreferenceType("latestPatch")
	VsanBaselinePreferenceTypenoRecommendation                   = VsanBaselinePreferenceType("noRecommendation")
	VsanBaselinePreferenceTypeVsanBaselinePreferenceType_Unknown = VsanBaselinePreferenceType("VsanBaselinePreferenceType_Unknown")
)

func init() {
	types.Add("VsanBaselinePreferenceType", reflect.TypeOf((*VsanBaselinePreferenceType)(nil)).Elem())
}
