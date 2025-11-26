// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package internal

const (
	// NamespaceClusterPath is the rest endpoint for the namespace cluster management API
	NamespaceClusterPath                    = "/api/vcenter/namespace-management/clusters"
	NamespaceDistributedSwitchCompatibility = "/api/vcenter/namespace-management/distributed-switch-compatibility"
	NamespaceEdgeClusterCompatibility       = "/api/vcenter/namespace-management/edge-cluster-compatibility"
	SupervisorServicesPath                  = "/api/vcenter/namespace-management/supervisor-services"
	SupervisorServicesVersionsPath          = "/versions"
	SupervisorsPath                         = "/api/vcenter/namespace-management/supervisors"

	NamespacesPath = "/api/vcenter/namespaces/instances"
	VmClassesPath  = "/api/vcenter/namespace-management/virtual-machine-classes"
)

type SupportBundleToken struct {
	Value string `json:"wcp-support-bundle-token"`
}
