// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"reflect"

	"github.com/vmware/govmomi/vim25/types"
)

func (b *AgencyConfigInfo) GetAgencyConfigInfo() *AgencyConfigInfo { return b }

type BaseAgencyConfigInfo interface {
	GetAgencyConfigInfo() *AgencyConfigInfo
}

func init() {
	types.Add("BaseAgencyConfigInfo", reflect.TypeOf((*AgencyConfigInfo)(nil)).Elem())
}

func (b *AgencyIssue) GetAgencyIssue() *AgencyIssue { return b }

type BaseAgencyIssue interface {
	GetAgencyIssue() *AgencyIssue
}

func init() {
	types.Add("BaseAgencyIssue", reflect.TypeOf((*AgencyIssue)(nil)).Elem())
}

func (b *AgencyScope) GetAgencyScope() *AgencyScope { return b }

type BaseAgencyScope interface {
	GetAgencyScope() *AgencyScope
}

func init() {
	types.Add("BaseAgencyScope", reflect.TypeOf((*AgencyScope)(nil)).Elem())
}

func (b *AgentIssue) GetAgentIssue() *AgentIssue { return b }

type BaseAgentIssue interface {
	GetAgentIssue() *AgentIssue
}

func init() {
	types.Add("BaseAgentIssue", reflect.TypeOf((*AgentIssue)(nil)).Elem())
}

func (b *AgentSslTrust) GetAgentSslTrust() *AgentSslTrust { return b }

type BaseAgentSslTrust interface {
	GetAgentSslTrust() *AgentSslTrust
}

func init() {
	types.Add("BaseAgentSslTrust", reflect.TypeOf((*AgentSslTrust)(nil)).Elem())
}

func (b *AgentStoragePolicy) GetAgentStoragePolicy() *AgentStoragePolicy { return b }

type BaseAgentStoragePolicy interface {
	GetAgentStoragePolicy() *AgentStoragePolicy
}

func init() {
	types.Add("BaseAgentStoragePolicy", reflect.TypeOf((*AgentStoragePolicy)(nil)).Elem())
}

func (b *ClusterAgentAgentIssue) GetClusterAgentAgentIssue() *ClusterAgentAgentIssue { return b }

type BaseClusterAgentAgentIssue interface {
	GetClusterAgentAgentIssue() *ClusterAgentAgentIssue
}

func init() {
	types.Add("BaseClusterAgentAgentIssue", reflect.TypeOf((*ClusterAgentAgentIssue)(nil)).Elem())
}

func (b *ClusterAgentVmHookFailed) GetClusterAgentVmHookFailed() *ClusterAgentVmHookFailed { return b }

type BaseClusterAgentVmHookFailed interface {
	GetClusterAgentVmHookFailed() *ClusterAgentVmHookFailed
}

func init() {
	types.Add("BaseClusterAgentVmHookFailed", reflect.TypeOf((*ClusterAgentVmHookFailed)(nil)).Elem())
}

func (b *ClusterAgentVmIssue) GetClusterAgentVmIssue() *ClusterAgentVmIssue { return b }

type BaseClusterAgentVmIssue interface {
	GetClusterAgentVmIssue() *ClusterAgentVmIssue
}

func init() {
	types.Add("BaseClusterAgentVmIssue", reflect.TypeOf((*ClusterAgentVmIssue)(nil)).Elem())
}

func (b *ClusterAgentVmNotDeployed) GetClusterAgentVmNotDeployed() *ClusterAgentVmNotDeployed {
	return b
}

type BaseClusterAgentVmNotDeployed interface {
	GetClusterAgentVmNotDeployed() *ClusterAgentVmNotDeployed
}

func init() {
	types.Add("BaseClusterAgentVmNotDeployed", reflect.TypeOf((*ClusterAgentVmNotDeployed)(nil)).Elem())
}

func (b *ClusterAgentVmPoweredOff) GetClusterAgentVmPoweredOff() *ClusterAgentVmPoweredOff { return b }

type BaseClusterAgentVmPoweredOff interface {
	GetClusterAgentVmPoweredOff() *ClusterAgentVmPoweredOff
}

func init() {
	types.Add("BaseClusterAgentVmPoweredOff", reflect.TypeOf((*ClusterAgentVmPoweredOff)(nil)).Elem())
}

func (b *EamAppFault) GetEamAppFault() *EamAppFault { return b }

type BaseEamAppFault interface {
	GetEamAppFault() *EamAppFault
}

func init() {
	types.Add("BaseEamAppFault", reflect.TypeOf((*EamAppFault)(nil)).Elem())
}

func (b *EamFault) GetEamFault() *EamFault { return b }

type BaseEamFault interface {
	GetEamFault() *EamFault
}

func init() {
	types.Add("BaseEamFault", reflect.TypeOf((*EamFault)(nil)).Elem())
}

func (b *EamObjectRuntimeInfo) GetEamObjectRuntimeInfo() *EamObjectRuntimeInfo { return b }

type BaseEamObjectRuntimeInfo interface {
	GetEamObjectRuntimeInfo() *EamObjectRuntimeInfo
}

func init() {
	types.Add("BaseEamObjectRuntimeInfo", reflect.TypeOf((*EamObjectRuntimeInfo)(nil)).Elem())
}

func (b *EamRuntimeFault) GetEamRuntimeFault() *EamRuntimeFault { return b }

type BaseEamRuntimeFault interface {
	GetEamRuntimeFault() *EamRuntimeFault
}

func init() {
	types.Add("BaseEamRuntimeFault", reflect.TypeOf((*EamRuntimeFault)(nil)).Elem())
}

func (b *HostIssue) GetHostIssue() *HostIssue { return b }

type BaseHostIssue interface {
	GetHostIssue() *HostIssue
}

func init() {
	types.Add("BaseHostIssue", reflect.TypeOf((*HostIssue)(nil)).Elem())
}

func (b *IntegrityAgencyVUMIssue) GetIntegrityAgencyVUMIssue() *IntegrityAgencyVUMIssue { return b }

type BaseIntegrityAgencyVUMIssue interface {
	GetIntegrityAgencyVUMIssue() *IntegrityAgencyVUMIssue
}

func init() {
	types.Add("BaseIntegrityAgencyVUMIssue", reflect.TypeOf((*IntegrityAgencyVUMIssue)(nil)).Elem())
}

func (b *Issue) GetIssue() *Issue { return b }

type BaseIssue interface {
	GetIssue() *Issue
}

func init() {
	types.Add("BaseIssue", reflect.TypeOf((*Issue)(nil)).Elem())
}

func (b *NoAgentVmDatastore) GetNoAgentVmDatastore() *NoAgentVmDatastore { return b }

type BaseNoAgentVmDatastore interface {
	GetNoAgentVmDatastore() *NoAgentVmDatastore
}

func init() {
	types.Add("BaseNoAgentVmDatastore", reflect.TypeOf((*NoAgentVmDatastore)(nil)).Elem())
}

func (b *NoAgentVmNetwork) GetNoAgentVmNetwork() *NoAgentVmNetwork { return b }

type BaseNoAgentVmNetwork interface {
	GetNoAgentVmNetwork() *NoAgentVmNetwork
}

func init() {
	types.Add("BaseNoAgentVmNetwork", reflect.TypeOf((*NoAgentVmNetwork)(nil)).Elem())
}

func (b *PersonalityAgencyDepotIssue) GetPersonalityAgencyDepotIssue() *PersonalityAgencyDepotIssue {
	return b
}

type BasePersonalityAgencyDepotIssue interface {
	GetPersonalityAgencyDepotIssue() *PersonalityAgencyDepotIssue
}

func init() {
	types.Add("BasePersonalityAgencyDepotIssue", reflect.TypeOf((*PersonalityAgencyDepotIssue)(nil)).Elem())
}

func (b *PersonalityAgencyPMIssue) GetPersonalityAgencyPMIssue() *PersonalityAgencyPMIssue { return b }

type BasePersonalityAgencyPMIssue interface {
	GetPersonalityAgencyPMIssue() *PersonalityAgencyPMIssue
}

func init() {
	types.Add("BasePersonalityAgencyPMIssue", reflect.TypeOf((*PersonalityAgencyPMIssue)(nil)).Elem())
}

func (b *PersonalityAgentPMIssue) GetPersonalityAgentPMIssue() *PersonalityAgentPMIssue { return b }

type BasePersonalityAgentPMIssue interface {
	GetPersonalityAgentPMIssue() *PersonalityAgentPMIssue
}

func init() {
	types.Add("BasePersonalityAgentPMIssue", reflect.TypeOf((*PersonalityAgentPMIssue)(nil)).Elem())
}

func (b *SolutionsHookAcknowledgeConfig) GetSolutionsHookAcknowledgeConfig() *SolutionsHookAcknowledgeConfig {
	return b
}

type BaseSolutionsHookAcknowledgeConfig interface {
	GetSolutionsHookAcknowledgeConfig() *SolutionsHookAcknowledgeConfig
}

func init() {
	types.Add("BaseSolutionsHookAcknowledgeConfig", reflect.TypeOf((*SolutionsHookAcknowledgeConfig)(nil)).Elem())
}

func (b *SolutionsRemediationPolicy) GetSolutionsRemediationPolicy() *SolutionsRemediationPolicy {
	return b
}

type BaseSolutionsRemediationPolicy interface {
	GetSolutionsRemediationPolicy() *SolutionsRemediationPolicy
}

func init() {
	types.Add("BaseSolutionsRemediationPolicy", reflect.TypeOf((*SolutionsRemediationPolicy)(nil)).Elem())
}

func (b *SolutionsStoragePolicy) GetSolutionsStoragePolicy() *SolutionsStoragePolicy { return b }

type BaseSolutionsStoragePolicy interface {
	GetSolutionsStoragePolicy() *SolutionsStoragePolicy
}

func init() {
	types.Add("BaseSolutionsStoragePolicy", reflect.TypeOf((*SolutionsStoragePolicy)(nil)).Elem())
}

func (b *SolutionsTypeSpecificSolutionConfig) GetSolutionsTypeSpecificSolutionConfig() *SolutionsTypeSpecificSolutionConfig {
	return b
}

type BaseSolutionsTypeSpecificSolutionConfig interface {
	GetSolutionsTypeSpecificSolutionConfig() *SolutionsTypeSpecificSolutionConfig
}

func init() {
	types.Add("BaseSolutionsTypeSpecificSolutionConfig", reflect.TypeOf((*SolutionsTypeSpecificSolutionConfig)(nil)).Elem())
}

func (b *SolutionsVMSource) GetSolutionsVMSource() *SolutionsVMSource { return b }

type BaseSolutionsVMSource interface {
	GetSolutionsVMSource() *SolutionsVMSource
}

func init() {
	types.Add("BaseSolutionsVMSource", reflect.TypeOf((*SolutionsVMSource)(nil)).Elem())
}

func (b *VibIssue) GetVibIssue() *VibIssue { return b }

type BaseVibIssue interface {
	GetVibIssue() *VibIssue
}

func init() {
	types.Add("BaseVibIssue", reflect.TypeOf((*VibIssue)(nil)).Elem())
}

func (b *VibNotInstalled) GetVibNotInstalled() *VibNotInstalled { return b }

type BaseVibNotInstalled interface {
	GetVibNotInstalled() *VibNotInstalled
}

func init() {
	types.Add("BaseVibNotInstalled", reflect.TypeOf((*VibNotInstalled)(nil)).Elem())
}

func (b *VibVibServicesSslTrust) GetVibVibServicesSslTrust() *VibVibServicesSslTrust { return b }

type BaseVibVibServicesSslTrust interface {
	GetVibVibServicesSslTrust() *VibVibServicesSslTrust
}

func init() {
	types.Add("BaseVibVibServicesSslTrust", reflect.TypeOf((*VibVibServicesSslTrust)(nil)).Elem())
}

func (b *VmDeployed) GetVmDeployed() *VmDeployed { return b }

type BaseVmDeployed interface {
	GetVmDeployed() *VmDeployed
}

func init() {
	types.Add("BaseVmDeployed", reflect.TypeOf((*VmDeployed)(nil)).Elem())
}

func (b *VmIssue) GetVmIssue() *VmIssue { return b }

type BaseVmIssue interface {
	GetVmIssue() *VmIssue
}

func init() {
	types.Add("BaseVmIssue", reflect.TypeOf((*VmIssue)(nil)).Elem())
}

func (b *VmNotDeployed) GetVmNotDeployed() *VmNotDeployed { return b }

type BaseVmNotDeployed interface {
	GetVmNotDeployed() *VmNotDeployed
}

func init() {
	types.Add("BaseVmNotDeployed", reflect.TypeOf((*VmNotDeployed)(nil)).Elem())
}

func (b *VmPoweredOff) GetVmPoweredOff() *VmPoweredOff { return b }

type BaseVmPoweredOff interface {
	GetVmPoweredOff() *VmPoweredOff
}

func init() {
	types.Add("BaseVmPoweredOff", reflect.TypeOf((*VmPoweredOff)(nil)).Elem())
}
