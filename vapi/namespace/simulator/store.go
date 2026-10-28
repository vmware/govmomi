// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"path"
	"slices"
	"strings"
	"sync"

	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vapi/namespace"
	"github.com/vmware/govmomi/vim25/mo"
)

// store holds the simulated Namespace Management state. Every read and write of
// that state is a method here; the HTTP handlers in simulator.go own only
// request decoding and response status.
type store struct {
	mu              sync.Mutex
	registry        *simulator.Registry
	clusters        map[string]namespace.ClusterSummary
	supervisors     map[string]supervisor
	namespaces      map[string]namespace.NamespaceInstanceInfoV2
	vmClasses       map[string]namespace.VirtualMachineClassInfo
	services        map[string]namespace.SupervisorServiceInfo
	serviceVersions map[string]map[string]namespace.SupervisorServiceVersionInfo
}

type supervisor struct {
	info     namespace.SupervisorSummaryInfo
	topology []namespace.SupervisorTopologyInfo
	networks []namespace.NetworkInfo
}

func newStore() *store {
	return &store{
		clusters:        make(map[string]namespace.ClusterSummary),
		supervisors:     make(map[string]supervisor),
		namespaces:      make(map[string]namespace.NamespaceInstanceInfoV2),
		vmClasses:       make(map[string]namespace.VirtualMachineClassInfo),
		services:        make(map[string]namespace.SupervisorServiceInfo),
		serviceVersions: make(map[string]map[string]namespace.SupervisorServiceVersionInfo),
	}
}

// -------- clusters and supervisors --------

func (s *store) listClusters() []namespace.ClusterSummary {
	s.mu.Lock()
	defer s.mu.Unlock()

	clusters := make([]namespace.ClusterSummary, 0, len(s.clusters))
	for _, cluster := range s.clusters {
		clusters = append(clusters, cluster)
	}
	return clusters
}

func (s *store) enableCluster(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.clusters[id] = namespace.ClusterSummary{
		ID:               id,
		ConfigStatus:     &namespace.RunningConfigStatus,
		KubernetesStatus: &namespace.ReadyKubernetesStatus,
	}
	s.supervisors[id] = s.newSupervisor(id)
}

func (s *store) disableCluster(id string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.clusters[id]; !exists {
		return false
	}
	delete(s.clusters, id)
	delete(s.supervisors, id)
	return true
}

func (s *store) getSupervisor(id string) (supervisor, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	sup, exists := s.supervisors[id]
	return sup, exists
}

func (s *store) listSupervisors() []namespace.SupervisorsSummaryResultItem {
	s.mu.Lock()
	defer s.mu.Unlock()

	items := make([]namespace.SupervisorsSummaryResultItem, 0, len(s.supervisors))
	for id, sup := range s.supervisors {
		items = append(items, namespace.SupervisorsSummaryResultItem{
			Supervisor: id,
			Info:       sup.info,
		})
	}

	slices.SortFunc(items, func(a, b namespace.SupervisorsSummaryResultItem) int {
		return strings.Compare(a.Supervisor, b.Supervisor)
	})

	return items
}

func (s *store) newSupervisor(id string) supervisor {
	name := s.lookupClusterName(id)

	return supervisor{
		info: namespace.SupervisorSummaryInfo{
			Name:             name,
			ConfigStatus:     namespace.RunningConfigStatus,
			KubernetesStatus: namespace.ReadyKubernetesStatus,
			Stats: namespace.SupervisorSummaryInfoStats{
				CpuUsed:         1000,
				CpuCapacity:     8000,
				MemoryUsed:      4096,
				MemoryCapacity:  32768,
				StorageUsed:     10240,
				StorageCapacity: 102400,
			},
		},
		topology: []namespace.SupervisorTopologyInfo{{
			Zone:     name,
			Clusters: []string{id},
		}},
		networks: s.collectPortgroupNetworks(),
	}
}

func (s *store) lookupClusterName(id string) string {
	for _, entity := range s.registry.All("ClusterComputeResource") {
		if entity.Entity().Self.Value == id {
			return entity.Entity().Name
		}
	}
	return id
}

// collectPortgroupNetworks maps inventory DVPGs to workload networks, minus the per-switch uplink.
func (s *store) collectPortgroupNetworks() []namespace.NetworkInfo {
	entities := s.registry.All("DistributedVirtualPortgroup")

	slices.SortFunc(entities, func(a, b mo.Entity) int {
		return strings.Compare(a.Entity().Name, b.Entity().Name)
	})

	networks := make([]namespace.NetworkInfo, 0, len(entities))
	for _, entity := range entities {
		if strings.Contains(entity.Entity().Name, "-DVUplinks") {
			continue
		}

		networks = append(networks, namespace.NetworkInfo{
			Network:         entity.Entity().Name,
			NetworkProvider: namespace.VSphereNetworkProvider,
			VsphereNetwork: &namespace.VsphereNetworkInfo{
				Portgroup:        entity.Entity().Self.Value,
				IPAssignmentMode: namespace.IPAssignmentModeStaticRange,
				Gateway:          "10.0.0.1",
				SubnetMask:       "255.255.255.0",
			},
		})
	}

	return networks
}

// -------- supervisor services --------

func (s *store) listServices() []namespace.SupervisorServiceSummary {
	s.mu.Lock()
	defer s.mu.Unlock()

	result := make([]namespace.SupervisorServiceSummary, 0, len(s.services))
	for id, info := range s.services {
		result = append(result, namespace.SupervisorServiceSummary{
			ID:    id,
			Name:  info.Name,
			State: info.State,
		})
	}
	return result
}

func (s *store) getService(id string) (namespace.SupervisorServiceInfo, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	info, exists := s.services[id]
	return info, exists
}

func (s *store) addService(pkg carvelPackage, content string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.services[pkg.id]; exists {
		return false
	}
	s.services[pkg.id] = namespace.SupervisorServiceInfo{
		Name:        pkg.displayName,
		State:       activatedState,
		Description: pkg.description,
	}
	s.serviceVersions[pkg.id] = map[string]namespace.SupervisorServiceVersionInfo{
		pkg.version: createServiceVersionInfo(pkg, content),
	}
	return true
}

func (s *store) setServiceState(id, state string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	info, exists := s.services[id]
	if !exists {
		return
	}
	info.State = state
	s.services[id] = info
}

func (s *store) removeService(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.services, id)
	delete(s.serviceVersions, id)
}

func (s *store) hasServiceVersions(id string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.serviceVersions[id]
	return exists
}

func (s *store) listServiceVersions(id string) []namespace.SupervisorServiceVersionSummary {
	s.mu.Lock()
	defer s.mu.Unlock()

	versions := s.serviceVersions[id]

	result := make([]namespace.SupervisorServiceVersionSummary, 0, len(versions))
	for version, info := range versions {
		result = append(result, namespace.SupervisorServiceVersionSummary{
			SupervisorServiceInfo: info.SupervisorServiceInfo,
			Version:               version,
		})
	}
	return result
}

func (s *store) getServiceVersion(id, version string) (namespace.SupervisorServiceVersionInfo, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	info, exists := s.serviceVersions[id][version]
	return info, exists
}

func (s *store) addServiceVersion(id string, pkg carvelPackage, content string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	versions, exists := s.serviceVersions[id]
	if !exists {
		return false
	}
	if _, exists := versions[pkg.version]; exists {
		return false
	}
	versions[pkg.version] = createServiceVersionInfo(pkg, content)
	return true
}

func (s *store) setServiceVersionState(id, version, state string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	info, exists := s.serviceVersions[id][version]
	if !exists {
		return
	}
	info.State = state
	s.serviceVersions[id][version] = info
}

func (s *store) removeServiceVersion(id, version string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.serviceVersions[id], version)
}

func createServiceVersionInfo(pkg carvelPackage, content string) namespace.SupervisorServiceVersionInfo {
	return namespace.SupervisorServiceVersionInfo{
		SupervisorServiceInfo: namespace.SupervisorServiceInfo{
			Name:        pkg.displayName,
			State:       activatedState,
			Description: pkg.description,
		},
		Content:     content,
		ContentType: "CARVEL_APPS_YAML",
	}
}

// -------- namespaces --------

func (s *store) getNamespace(name string) (namespace.NamespaceInstanceInfoV2, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	info, exists := s.namespaces[name]
	return info, exists
}

func (s *store) listNamespaces() []namespace.NamespacesInstanceSummary {
	s.mu.Lock()
	defer s.mu.Unlock()

	result := make([]namespace.NamespacesInstanceSummary, 0, len(s.namespaces))
	for name, info := range s.namespaces {
		result = append(result, namespace.NamespacesInstanceSummary{
			ClusterId:            info.Supervisor,
			Namespace:            name,
			ConfigStatus:         string(info.ConfigStatus),
			Description:          info.Description,
			Stats:                convertToInstanceStats(info.Stats),
			SelfServiceNamespace: info.SelfServiceNamespace,
		})
	}
	return result
}

func (s *store) listNamespacesV2() []namespace.NamespaceInstanceSummaryV2 {
	s.mu.Lock()
	defer s.mu.Unlock()

	result := make([]namespace.NamespaceInstanceSummaryV2, 0, len(s.namespaces))
	for name, info := range s.namespaces {
		result = append(result, namespace.NamespaceInstanceSummaryV2{
			Supervisor:           info.Supervisor,
			Namespace:            name,
			Description:          info.Description,
			ConfigStatus:         info.ConfigStatus,
			Stats:                info.Stats,
			SelfServiceNamespace: info.SelfServiceNamespace,
		})
	}
	return result
}

func (s *store) addNamespace(spec namespace.NamespacesInstanceCreateSpec) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.namespaces[spec.Namespace]; exists {
		return false
	}
	s.namespaces[spec.Namespace] = namespace.NamespaceInstanceInfoV2{
		Supervisor:    spec.Cluster,
		ConfigStatus:  namespace.NamespaceInstanceConfigStatusRunning,
		VmServiceSpec: spec.VmServiceSpec,
		StorageSpecs:  spec.StorageSpecs,
	}
	return true
}

func (s *store) addNamespaceV2(spec namespace.NamespaceInstanceCreateSpecV2) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.namespaces[spec.Namespace]; exists {
		return false
	}
	s.namespaces[spec.Namespace] = createNamespaceInfo(spec)
	return true
}

func (s *store) updateNamespace(name string, spec namespace.NamespacesInstanceUpdateSpec) {
	s.mu.Lock()
	defer s.mu.Unlock()

	info, exists := s.namespaces[name]
	if !exists {
		return
	}

	if spec.VmServiceSpec.ContentLibraries != nil || spec.VmServiceSpec.VmClasses != nil {
		info.VmServiceSpec = spec.VmServiceSpec
	}

	if spec.StorageSpecs != nil {
		info.StorageSpecs = spec.StorageSpecs
	}
	if spec.Zones != nil {
		info.Zones = createZoneInfos(*spec.Zones)
	}
	if spec.AccessList != nil {
		info.AccessList = *spec.AccessList
	}
	if spec.NetworkSpec != nil {
		updateNetworkConfig(&info, *spec.NetworkSpec)
	}

	s.namespaces[name] = info
}

func (s *store) removeNamespace(name string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.namespaces, name)
}

func createNamespaceInfo(spec namespace.NamespaceInstanceCreateSpecV2) namespace.NamespaceInstanceInfoV2 {
	info := namespace.NamespaceInstanceInfoV2{
		Supervisor:   spec.Supervisor,
		ConfigStatus: namespace.NamespaceInstanceConfigStatusRunning,
	}

	if spec.Description != nil {
		info.Description = *spec.Description
	}
	if spec.SelfServiceNamespace != nil {
		info.SelfServiceNamespace = *spec.SelfServiceNamespace
	}
	if spec.StorageSpecs != nil {
		info.StorageSpecs = *spec.StorageSpecs
	}
	if spec.VmServiceSpec != nil {
		info.VmServiceSpec = *spec.VmServiceSpec
	}
	if spec.ContentLibraries != nil {
		info.ContentLibraries = *spec.ContentLibraries
	}
	if spec.Edges != nil {
		info.Edges = *spec.Edges
	}
	if spec.AccessList != nil {
		info.AccessList = *spec.AccessList
	}
	if spec.Zones != nil {
		info.Zones = createZoneInfos(*spec.Zones)
	}
	if spec.NetworkSpec != nil {
		info.NetworkSpec = createNetworkConfigInfo(*spec.NetworkSpec)
	} else if spec.Networks == nil {
		info.NetworkSpec = createAutoVpcNetworkConfigInfo(spec.Namespace)
	}
	if spec.Networks != nil {
		addNetworks(&info, *spec.Networks)
	}

	return info
}

// convertToNamespacesInstanceInfo projects the v2 record onto the v1 API shape.
func convertToNamespacesInstanceInfo(info namespace.NamespaceInstanceInfoV2) namespace.NamespacesInstanceInfo {
	return namespace.NamespacesInstanceInfo{
		ClusterId:            info.Supervisor,
		ConfigStatus:         string(info.ConfigStatus),
		Description:          info.Description,
		Stats:                convertToInstanceStats(info.Stats),
		SelfServiceNamespace: info.SelfServiceNamespace,
		VmServiceSpec:        info.VmServiceSpec,
		StorageSpecs:         info.StorageSpecs,
	}
}

// createAutoVpcNetworkConfigInfo mirrors NSX provisioning a VPC when a namespace requests no network.
func createAutoVpcNetworkConfigInfo(name string) *namespace.NetworkConfigInfo {
	return &namespace.NetworkConfigInfo{
		NetworkProvider: "NSX_VPC",
		VpcNetwork: &namespace.VpcNetworkInfo{
			Vpc:         path.Join("/orgs/default/projects/default/vpcs", name),
			AutoCreated: true,
		},
	}
}

func createNetworkConfigInfo(spec namespace.NetworkConfigCreateSpec) *namespace.NetworkConfigInfo {
	info := &namespace.NetworkConfigInfo{NetworkProvider: spec.NetworkProvider}

	if vpc := spec.VpcNetwork; vpc != nil {
		info.VpcNetwork = &namespace.VpcNetworkInfo{}
		if vpc.VpcConfig != nil {
			info.VpcNetwork.VpcConfig = *vpc.VpcConfig
		}
		if vpc.Vpc != nil {
			info.VpcNetwork.Vpc = *vpc.Vpc
		}
		if vpc.SharedSubnets != nil {
			info.VpcNetwork.SharedSubnets = *vpc.SharedSubnets
		}
		if vpc.DefaultSubnetSize != nil {
			info.VpcNetwork.DefaultSubnetSize = *vpc.DefaultSubnetSize
		}
	}

	if vsphere := spec.VsphereNetwork; vsphere != nil {
		info.VsphereNetwork = &namespace.VsphereNetworkConfigInfo{}
		if vsphere.DefaultWorkloadNetwork != nil {
			info.VsphereNetwork.DefaultWorkloadNetwork = *vsphere.DefaultWorkloadNetwork
		}
	}

	return info
}

func updateNetworkConfig(info *namespace.NamespaceInstanceInfoV2, spec namespace.NetworkConfigUpdateSpec) {
	if info.NetworkSpec == nil {
		info.NetworkSpec = &namespace.NetworkConfigInfo{}
	}
	if spec.NetworkProvider != "" {
		info.NetworkSpec.NetworkProvider = spec.NetworkProvider
	}

	if vpc := spec.VpcConfig; vpc != nil {
		if info.NetworkSpec.VpcNetwork == nil {
			info.NetworkSpec.VpcNetwork = &namespace.VpcNetworkInfo{}
		}
		if vpc.DefaultSubnetSize != nil {
			info.NetworkSpec.VpcNetwork.DefaultSubnetSize = *vpc.DefaultSubnetSize
		}
		if vpc.PrivateCidrs != nil {
			info.NetworkSpec.VpcNetwork.VpcConfig.PrivateCidrs = vpc.PrivateCidrs
		}
		if vpc.SharedSubnets != nil {
			info.NetworkSpec.VpcNetwork.SharedSubnets = *vpc.SharedSubnets
		}
	}

	if update := spec.VsphereNetworkConfig; update != nil {
		removeNetworks(info, update.NetworksToRemove)
		addNetworks(info, update.NetworksToAdd)
	}
}

func removeNetworks(info *namespace.NamespaceInstanceInfoV2, networks []string) {
	if info.NetworkSpec == nil || info.NetworkSpec.VsphereNetwork == nil {
		return
	}

	vsphere := info.NetworkSpec.VsphereNetwork
	vsphere.Networks = slices.DeleteFunc(vsphere.Networks, func(network namespace.NamespaceNetworkInfo) bool {
		return slices.Contains(networks, network.Network)
	})
}

func addNetworks(info *namespace.NamespaceInstanceInfoV2, networks []string) {
	if len(networks) == 0 {
		return
	}
	if info.NetworkSpec == nil {
		info.NetworkSpec = &namespace.NetworkConfigInfo{}
	}
	if info.NetworkSpec.VsphereNetwork == nil {
		info.NetworkSpec.VsphereNetwork = &namespace.VsphereNetworkConfigInfo{}
	}

	vsphere := info.NetworkSpec.VsphereNetwork
	for _, name := range networks {
		attached := slices.ContainsFunc(vsphere.Networks, func(network namespace.NamespaceNetworkInfo) bool {
			return network.Network == name
		})
		if !attached {
			vsphere.Networks = append(vsphere.Networks, namespace.NamespaceNetworkInfo{
				Network:      name,
				ConfigStatus: namespace.NetworkConfigStatusActive,
			})
		}
	}
}

func createZoneInfos(specs []namespace.ZoneSpec) []namespace.ZoneInfo {
	zones := make([]namespace.ZoneInfo, len(specs))
	for i, spec := range specs {
		zones[i] = namespace.ZoneInfo{Zone: spec.Name, Type: "WORKLOAD"}
	}
	return zones
}

func convertToInstanceStats(stats namespace.Stats) namespace.NamespacesInstanceStats {
	return namespace.NamespacesInstanceStats{
		CpuUsed:     int64(stats.CPUUsed),
		MemoryUsed:  int64(stats.MemoryUsed),
		StorageUsed: int64(stats.StorageUsed),
	}
}

// -------- virtual machine classes --------

func (s *store) listVmClasses() []namespace.VirtualMachineClassInfo {
	s.mu.Lock()
	defer s.mu.Unlock()

	result := make([]namespace.VirtualMachineClassInfo, 0, len(s.vmClasses))
	for _, v := range s.vmClasses {
		result = append(result, v)
	}
	return result
}

func (s *store) getVmClass(id string) (namespace.VirtualMachineClassInfo, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	info, exists := s.vmClasses[id]
	return info, exists
}

func (s *store) addVmClass(spec namespace.VirtualMachineClassCreateSpec) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.vmClasses[spec.Id]; exists {
		return false
	}
	s.vmClasses[spec.Id] = namespace.VirtualMachineClassInfo{
		Id:                spec.Id,
		CpuCount:          spec.CpuCount,
		MemoryMb:          spec.MemoryMb,
		MemoryReservation: spec.MemoryReservation,
		CpuReservation:    spec.CpuReservation,
		Devices:           spec.Devices,
		ConfigSpec:        spec.ConfigSpec,
	}
	return true
}

func (s *store) updateVmClass(id string, spec namespace.VirtualMachineClassUpdateSpec) {
	s.mu.Lock()
	defer s.mu.Unlock()

	entry, exists := s.vmClasses[id]
	if !exists {
		return
	}

	entry.CpuCount = spec.CpuCount
	entry.MemoryMb = spec.MemoryMb
	entry.CpuReservation = spec.CpuReservation
	entry.MemoryReservation = spec.MemoryReservation
	entry.Devices = spec.Devices
	entry.ConfigSpec = spec.ConfigSpec

	s.vmClasses[id] = entry
}

func (s *store) removeVmClass(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.vmClasses, id)
}
