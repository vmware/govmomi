// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"archive/tar"
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"slices"
	"time"

	"github.com/google/uuid"
	"gopkg.in/yaml.v3"

	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vapi/namespace"
	vapi "github.com/vmware/govmomi/vapi/simulator"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"

	"github.com/vmware/govmomi/vapi/namespace/internal"
)

func init() {
	simulator.RegisterEndpoint(func(s *simulator.Service, r *simulator.Registry) {
		New(s.Listen).Register(s, r)
	})
}

// Handler implements the Namespace Management Modules API simulator
type Handler struct {
	URL   *url.URL
	store *store
}

type store struct {
	clusters        map[string]namespace.ClusterSummary
	namespaces      map[string]*namespace.NamespaceInstanceInfoV2
	vmClasses       map[string]*namespace.VirtualMachineClassInfo
	services        map[string]*namespace.SupervisorServiceInfo
	serviceVersions map[string]map[string]*namespace.SupervisorServiceVersionInfo
}

// New creates a Handler instance
func New(u *url.URL) *Handler {
	return &Handler{
		URL: u,
		store: &store{
			clusters:        make(map[string]namespace.ClusterSummary),
			namespaces:      make(map[string]*namespace.NamespaceInstanceInfoV2),
			vmClasses:       make(map[string]*namespace.VirtualMachineClassInfo),
			services:        make(map[string]*namespace.SupervisorServiceInfo),
			serviceVersions: make(map[string]map[string]*namespace.SupervisorServiceVersionInfo),
		},
	}
}

// Register Namespace Management API paths with the vapi simulator's http.ServeMux
func (h *Handler) Register(s *simulator.Service, r *simulator.Registry) {
	if r.IsVPX() {
		s.HandleFunc(internal.NamespacesPath, h.namespaces)
		s.HandleFunc(internal.NamespacesPath+"/{namespace}", h.namespace)
		s.HandleFunc(internal.NamespacesPath+"/{namespace}/{action}", h.namespaceAction)
		s.HandleFunc(internal.NamespacesPathV2, h.namespacesV2)
		s.HandleFunc(internal.NamespacesPathV2+"/{namespace}", h.namespaceV2)
		s.HandleFunc(internal.NamespaceClusterPath, h.clusters)
		s.HandleFunc(internal.NamespaceClusterPath+"/{id}", h.cluster)
		s.HandleFunc(internal.NamespaceClusterPath+"/{id}/support-bundle", h.clustersSupportBundle)
		s.HandleFunc(internal.NamespaceDistributedSwitchCompatibility, h.listCompatibleDistributedSwitches)
		s.HandleFunc(internal.NamespaceEdgeClusterCompatibility, h.listCompatibleEdgeClusters)

		s.HandleFunc(internal.SupervisorServicesPath, h.services)
		s.HandleFunc(internal.SupervisorServicesPath+"/{id}", h.service)
		s.HandleFunc(internal.SupervisorServicesPath+"/{id}"+internal.SupervisorServicesVersionsPath, h.serviceVersions)
		s.HandleFunc(internal.SupervisorServicesPath+"/{id}"+internal.SupervisorServicesVersionsPath+"/{version}", h.serviceVersion)

		s.HandleFunc(internal.VmClassesPath, h.vmClasses)
		s.HandleFunc(internal.VmClassesPath+"/{vmClass}", h.vmClass)
	}
}

func (h *Handler) clusters(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		clusters := make([]namespace.ClusterSummary, 0, len(h.store.clusters))
		for _, cluster := range h.store.clusters {
			clusters = append(clusters, cluster)
		}
		vapi.StatusOK(w, clusters)
	}
}

func (h *Handler) cluster(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	switch r.Method {
	case http.MethodPost:
		switch r.URL.Query().Get("action") {
		case "enable":
			h.store.clusters[id] = namespace.ClusterSummary{
				ID:               id,
				ConfigStatus:     &namespace.RunningConfigStatus,
				KubernetesStatus: &namespace.ReadyKubernetesStatus,
			}
			vapi.StatusOK(w)
		case "disable":
			if _, exists := h.store.clusters[id]; !exists {
				vapi.ApiErrorNotFound(w)
				return
			}
			delete(h.store.clusters, id)
			vapi.StatusOK(w)
		default:
			vapi.ApiErrorInvalidArgument(w)
		}
	}
}

func (h *Handler) clustersSupportBundle(w http.ResponseWriter, r *http.Request) {
	var token internal.SupportBundleToken
	_ = json.NewDecoder(r.Body).Decode(&token)
	_ = r.Body.Close()

	if token.Value == "" {
		u := *h.URL
		u.Path = r.URL.Path
		// Create support bundle request
		location := namespace.SupportBundleLocation{
			Token: namespace.SupportBundleToken{
				Token: uuid.New().String(),
			},
			URL: u.String(),
		}

		vapi.StatusOK(w, &location)
		return
	}

	// Get support bundle
	id := path.Base(path.Dir(r.URL.Path))
	name := fmt.Sprintf("wcp-support-bundle-%s-%s--00-00.tar", id, time.Now().Format("2006Jan02"))

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", name))
	w.Header().Set("Content-Type", "application/octet-stream")

	readme := "vcsim generated support bundle"
	tw := tar.NewWriter(w)
	_ = tw.WriteHeader(&tar.Header{
		Name:    "README",
		Size:    int64(len(readme) + 1),
		Mode:    0444,
		ModTime: time.Now(),
	})
	_, _ = fmt.Fprintln(tw, readme)
	_ = tw.Close()
}

func (h *Handler) listCompatibleDistributedSwitches(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

		// normally expect to get exactly one result back
		switches := []namespace.DistributedSwitchCompatibilitySummary{
			{
				Compatible:        true,
				DistributedSwitch: "Compatible-DVS-1",
			},
		}
		vapi.StatusOK(w, switches)
	}
}

func (h *Handler) listCompatibleEdgeClusters(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

		// CLI is able to filter in case we get multiple results
		switches := []namespace.EdgeClusterCompatibilitySummary{
			{
				Compatible:  true,
				EdgeCluster: "Compat-Edge-ID1",
				DisplayName: "Edge-Cluster-1",
			},
			{
				Compatible:  true,
				EdgeCluster: "Compat-Edge-ID2",
				DisplayName: "Edge-Cluster-2",
			},
		}
		vapi.StatusOK(w, switches)
	}
}

func (h *Handler) services(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		result := make([]namespace.SupervisorServiceSummary, 0, len(h.store.services))
		for id, info := range h.store.services {
			result = append(result, namespace.SupervisorServiceSummary{
				ID:    id,
				Name:  info.Name,
				State: info.State,
			})
		}
		vapi.StatusOK(w, result)
	case http.MethodPost:
		var spec namespace.SupervisorService
		if !vapi.Decode(r, w, &spec) {
			return
		}
		if spec.CarvelService == nil {
			vapi.ApiErrorUnsupported(w)
			return
		}

		content := spec.CarvelService.VersionSpec.Content
		pkg, err := parseCarvelPackage(content)
		if err != nil {
			vapi.ApiErrorInvalidArgument(w)
			return
		}
		if _, exists := h.store.services[pkg.id]; exists {
			vapi.ApiErrorAlreadyExists(w)
			return
		}

		h.store.services[pkg.id] = &namespace.SupervisorServiceInfo{
			Name:        pkg.displayName,
			State:       activatedState,
			Description: pkg.description,
		}
		h.store.serviceVersions[pkg.id] = map[string]*namespace.SupervisorServiceVersionInfo{
			pkg.version: createServiceVersionInfo(pkg, content),
		}

		vapi.StatusOK(w)
	}
}

func (h *Handler) service(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	info, exists := h.store.services[id]
	if !exists {
		vapi.ApiErrorNotFound(w)
		return
	}

	switch r.Method {
	case http.MethodGet:
		vapi.StatusOK(w, info)
	case http.MethodPatch:
		if !updateServiceState(w, r, &info.State) {
			return
		}
		vapi.StatusOK(w)
	case http.MethodDelete:
		if info.State != deactivatedState {
			vapi.ApiErrorNotAllowedInCurrentState(w)
			return
		}
		delete(h.store.services, id)
		delete(h.store.serviceVersions, id)
		vapi.StatusOK(w)
	}
}

func (h *Handler) serviceVersions(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	versions, exists := h.store.serviceVersions[id]
	if !exists {
		vapi.ApiErrorNotFound(w)
		return
	}

	switch r.Method {
	case http.MethodGet:
		result := make([]namespace.SupervisorServiceVersionSummary, 0, len(versions))
		for version, info := range versions {
			result = append(result, namespace.SupervisorServiceVersionSummary{
				SupervisorServiceInfo: info.SupervisorServiceInfo,
				Version:               version,
			})
		}
		vapi.StatusOK(w, result)
	case http.MethodPost:
		var spec namespace.SupervisorServiceVersion
		if !vapi.Decode(r, w, &spec) {
			return
		}
		if spec.CarvelService == nil {
			vapi.ApiErrorUnsupported(w)
			return
		}

		pkg, err := parseCarvelPackage(spec.CarvelService.Content)
		if err != nil || pkg.id != id {
			vapi.ApiErrorInvalidArgument(w)
			return
		}
		if _, exists := versions[pkg.version]; exists {
			vapi.ApiErrorAlreadyExists(w)
			return
		}

		versions[pkg.version] = createServiceVersionInfo(pkg, spec.CarvelService.Content)
		vapi.StatusOK(w)
	}
}

func (h *Handler) serviceVersion(w http.ResponseWriter, r *http.Request) {
	versions := h.store.serviceVersions[r.PathValue("id")]
	version := r.PathValue("version")

	info, exists := versions[version]
	if !exists {
		vapi.ApiErrorNotFound(w)
		return
	}

	switch r.Method {
	case http.MethodGet:
		vapi.StatusOK(w, info)
	case http.MethodPatch:
		if !updateServiceState(w, r, &info.State) {
			return
		}
		vapi.StatusOK(w)
	case http.MethodDelete:
		if info.State != deactivatedState {
			vapi.ApiErrorNotAllowedInCurrentState(w)
			return
		}
		delete(versions, version)
		vapi.StatusOK(w)
	}
}

const (
	activatedState   = "ACTIVATED"
	deactivatedState = "DEACTIVATED"
)

func updateServiceState(w http.ResponseWriter, r *http.Request, state *string) bool {
	switch r.URL.Query().Get("action") {
	case "activate":
		*state = activatedState
	case "deactivate":
		*state = deactivatedState
	default:
		vapi.ApiErrorInvalidArgument(w)
		return false
	}
	return true
}

type carvelPackage struct {
	id          string
	displayName string
	description string
	version     string
}

type carvelResource struct {
	Kind     string `yaml:"kind"`
	Metadata struct {
		Name string `yaml:"name"`
	} `yaml:"metadata"`
	Spec struct {
		DisplayName      string `yaml:"displayName"`
		ShortDescription string `yaml:"shortDescription"`
		Version          string `yaml:"version"`
	} `yaml:"spec"`
}

func parseCarvelPackage(content string) (carvelPackage, error) {
	manifest, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		return carvelPackage{}, err
	}

	var pkg carvelPackage
	decoder := yaml.NewDecoder(bytes.NewReader(manifest))
	for {
		var resource carvelResource
		err := decoder.Decode(&resource)
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return carvelPackage{}, err
		}

		switch resource.Kind {
		case "PackageMetadata":
			pkg.id = resource.Metadata.Name
			pkg.displayName = resource.Spec.DisplayName
			pkg.description = resource.Spec.ShortDescription
		case "Package":
			pkg.version = resource.Spec.Version
		}
	}

	if pkg.id == "" || pkg.version == "" {
		return carvelPackage{}, errors.New("manifest must declare a PackageMetadata and a Package")
	}
	return pkg, nil
}

func createServiceVersionInfo(pkg carvelPackage, content string) *namespace.SupervisorServiceVersionInfo {
	return &namespace.SupervisorServiceVersionInfo{
		SupervisorServiceInfo: namespace.SupervisorServiceInfo{
			Name:        pkg.displayName,
			State:       activatedState,
			Description: pkg.description,
		},
		Content:     content,
		ContentType: "CARVEL_APPS_YAML",
	}
}

func (h *Handler) namespaces(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		result := make([]namespace.NamespacesInstanceSummary, 0, len(h.store.namespaces))

		for name, info := range h.store.namespaces {
			result = append(result, namespace.NamespacesInstanceSummary{
				ClusterId:            info.Supervisor,
				Namespace:            name,
				ConfigStatus:         string(info.ConfigStatus),
				Description:          info.Description,
				Stats:                convertToInstanceStats(info.Stats),
				SelfServiceNamespace: info.SelfServiceNamespace,
			})
		}

		vapi.StatusOK(w, result)
	case http.MethodPost:
		var spec namespace.NamespacesInstanceCreateSpec
		if !vapi.Decode(r, w, &spec) {
			return
		}

		h.store.namespaces[spec.Namespace] = &namespace.NamespaceInstanceInfoV2{
			Supervisor:    spec.Cluster,
			ConfigStatus:  namespace.NamespaceInstanceConfigStatusRunning,
			VmServiceSpec: spec.VmServiceSpec,
			StorageSpecs:  spec.StorageSpecs,
		}

		vapi.StatusOK(w)
	}
}

func (h *Handler) namespace(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("namespace")

	info, exists := h.store.namespaces[name]
	if !exists {
		vapi.ApiErrorNotFound(w)
		return
	}

	switch r.Method {
	case http.MethodGet:
		vapi.StatusOK(w, namespace.NamespacesInstanceInfo{
			ClusterId:            info.Supervisor,
			ConfigStatus:         string(info.ConfigStatus),
			Description:          info.Description,
			Stats:                convertToInstanceStats(info.Stats),
			SelfServiceNamespace: info.SelfServiceNamespace,
			VmServiceSpec:        info.VmServiceSpec,
			StorageSpecs:         info.StorageSpecs,
		})
	case http.MethodPatch:
		var spec namespace.NamespacesInstanceUpdateSpec
		if !vapi.Decode(r, w, &spec) {
			return
		}

		updateNamespace(info, spec)
		vapi.StatusOK(w)
	case http.MethodDelete:
		delete(h.store.namespaces, name)
		vapi.StatusOK(w)
	}
}

func (h *Handler) namespacesV2(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		result := make([]namespace.NamespaceInstanceSummaryV2, 0, len(h.store.namespaces))

		for name, info := range h.store.namespaces {
			result = append(result, namespace.NamespaceInstanceSummaryV2{
				Supervisor:           info.Supervisor,
				Namespace:            name,
				Description:          info.Description,
				ConfigStatus:         info.ConfigStatus,
				Stats:                info.Stats,
				SelfServiceNamespace: info.SelfServiceNamespace,
			})
		}

		vapi.StatusOK(w, result)
	case http.MethodPost:
		var spec namespace.NamespaceInstanceCreateSpecV2
		if !vapi.Decode(r, w, &spec) {
			return
		}

		h.store.namespaces[spec.Namespace] = createNamespaceInfo(spec)
		vapi.StatusOK(w)
	}
}

func (h *Handler) namespaceV2(w http.ResponseWriter, r *http.Request) {
	info, exists := h.store.namespaces[r.PathValue("namespace")]
	if !exists {
		vapi.ApiErrorNotFound(w)
		return
	}

	switch r.Method {
	case http.MethodGet:
		vapi.StatusOK(w, info)
	}
}

func createNamespaceInfo(spec namespace.NamespaceInstanceCreateSpecV2) *namespace.NamespaceInstanceInfoV2 {
	info := &namespace.NamespaceInstanceInfoV2{
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
	}
	if spec.Networks != nil {
		addNetworks(info, *spec.Networks)
	}

	return info
}

func updateNamespace(info *namespace.NamespaceInstanceInfoV2, spec namespace.NamespacesInstanceUpdateSpec) {
	info.VmServiceSpec = spec.VmServiceSpec

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
		updateNetworkConfig(info, *spec.NetworkSpec)
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
		zones[i] = namespace.ZoneInfo{Zone: spec.Name}
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

func (h *Handler) namespaceAction(w http.ResponseWriter, r *http.Request) {
	switch r.PathValue("action") {
	case "registervm":
		h.registerNamespaceVM(w, r)
	default:
		vapi.ApiErrorNotFound(w)
	}
}

func (h *Handler) registerNamespaceVM(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}

	var spec namespace.RegisterVMSpec
	if !vapi.Decode(r, w, &spec) {
		return
	}

	ref := types.ManagedObjectReference{Type: "VirtualMachine", Value: spec.VM}
	task := types.CreateTask{Obj: ref}
	key := &mo.Field{Path: "config.extraConfig", Key: "vmservice.virtualmachine.resource.yaml"}

	vapi.StatusOK(w, vapi.RunTask(*h.URL, task, func(ctx context.Context, c *vim25.Client) error {
		var vm mo.VirtualMachine
		_ = property.DefaultCollector(c).RetrieveOne(ctx, task.Obj, []string{key.String()}, &vm)
		if vm.Config == nil || len(vm.Config.ExtraConfig) == 0 {
			return fmt.Errorf("%s %s not found", task.Obj, key)
		}
		return nil
	}))
}

func (h *Handler) vmClasses(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		result := make([]*namespace.VirtualMachineClassInfo, 0, len(h.store.vmClasses))

		for _, v := range h.store.vmClasses {
			result = append(result, v)
		}

		vapi.StatusOK(w, result)
	case http.MethodPost:
		var spec namespace.VirtualMachineClassCreateSpec
		if !vapi.Decode(r, w, &spec) {
			return
		}

		newClass := namespace.VirtualMachineClassInfo{
			Id:                spec.Id,
			CpuCount:          spec.CpuCount,
			MemoryMb:          spec.MemoryMb,
			MemoryReservation: spec.MemoryReservation,
			CpuReservation:    spec.CpuReservation,
			Devices:           spec.Devices,
			ConfigSpec:        spec.ConfigSpec,
		}

		h.store.vmClasses[spec.Id] = &newClass

		vapi.StatusOK(w)
	}
}

func (h *Handler) vmClass(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("vmClass")

	switch r.Method {
	case http.MethodGet:
		if result, contains := h.store.vmClasses[id]; contains {
			vapi.StatusOK(w, result)
		} else {
			vapi.ApiErrorNotFound(w)
		}
	case http.MethodPatch:
		entry, contains := h.store.vmClasses[id]
		if !contains {
			vapi.ApiErrorNotFound(w)
			return
		}

		var spec namespace.VirtualMachineClassUpdateSpec
		if !vapi.Decode(r, w, &spec) {
			return
		}

		entry.CpuCount = spec.CpuCount
		entry.MemoryMb = spec.MemoryMb
		entry.CpuReservation = spec.CpuReservation
		entry.MemoryReservation = spec.MemoryReservation
		entry.Devices = spec.Devices
		entry.ConfigSpec = spec.ConfigSpec

		vapi.StatusOK(w)
	case http.MethodDelete:
		if _, contains := h.store.vmClasses[id]; contains {
			delete(h.store.vmClasses, id)
			vapi.StatusOK(w)
		} else {
			vapi.ApiErrorNotFound(w)
		}
	}
}
