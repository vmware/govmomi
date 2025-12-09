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
	SupervisorsSummariesPath                = SupervisorsPath + "/summaries"
	SupervisorSummaryPath                   = SupervisorsPath + "/%s/summary"
	SupervisorTopologyPath                  = SupervisorsPath + "/%s/topology"

	NamespacesPath   = "/api/vcenter/namespaces/instances"
	NamespacesPathV2 = "/api/vcenter/namespaces/instances/v2"
	VmClassesPath    = "/api/vcenter/namespace-management/virtual-machine-classes"
)

type SupportBundleToken struct {
	Value string `json:"wcp-support-bundle-token"`
}
