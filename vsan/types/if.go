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
