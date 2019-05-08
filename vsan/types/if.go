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

import "reflect"

type BaseMethodFault interface {
	GetMethodFault() *MethodFault
}

func init() {
	t["BaseMethodFault"] = reflect.TypeOf((*MethodFault)(nil)).Elem()
}

type BaseVsanClusterHealthResultBase interface {
	GetVsanClusterHealthResultBase() *VsanClusterHealthResultBase
}

func init() {
	t["BaseVsanClusterHealthResultBase"] = reflect.TypeOf((*VsanClusterHealthResultBase)(nil)).Elem()
}

type BaseElementDescription interface {
	GetElementDescription() *ElementDescription
}

func init() {
	t["BaseElementDescription"] = reflect.TypeOf((*ElementDescription)(nil)).Elem()
}

type BaseVsanNetworkConfigBaseIssue interface {
	GetVsanNetworkConfigBaseIssue() *VsanNetworkConfigBaseIssue
}

func init() {
	t["BaseVsanNetworkConfigBaseIssue"] = reflect.TypeOf((*VsanNetworkConfigBaseIssue)(nil)).Elem()
}

type BaseVsanClusterConfigInfo interface {
	GetVsanClusterConfigInfo() *VsanClusterConfigInfo
}

func init() {
	t["BaseVsanClusterConfigInfo"] = reflect.TypeOf((*VsanClusterConfigInfo)(nil)).Elem()
}
