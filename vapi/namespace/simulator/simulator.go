// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"archive/tar"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vapi/namespace"
	vapi "github.com/vmware/govmomi/vapi/simulator"
	"github.com/vmware/govmomi/view"
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
	URL *url.URL
}

// New creates a Handler instance
func New(u *url.URL) *Handler {
	return &Handler{
		URL: u,
	}
}

// Register Namespace Management API paths with the vapi simulator's http.ServeMux
func (h *Handler) Register(s *simulator.Service, r *simulator.Registry) {
	if r.IsVPX() {
		s.HandleFunc(internal.NamespacesPath, h.namespaces)
		s.HandleFunc(internal.NamespacesPath+"/", h.namespaces)
		s.HandleFunc(internal.NamespaceClusterPath, h.clusters)
		s.HandleFunc(internal.NamespaceClusterPath+"/", h.clustersID)
		s.HandleFunc(internal.NamespaceDistributedSwitchCompatibility+"/", h.listCompatibleDistributedSwitches)
		s.HandleFunc(internal.NamespaceEdgeClusterCompatibility+"/", h.listCompatibleEdgeClusters)

		s.HandleFunc(internal.SupervisorServicesPath, h.listServices)
		s.HandleFunc(internal.SupervisorServicesPath+"/", h.getService)
		s.HandleFunc(internal.SupervisorServicesPath+"/{id}"+internal.SupervisorServicesVersionsPath, h.versionsForService)
		s.HandleFunc(internal.SupervisorServicesPath+"/{id}"+internal.SupervisorServicesVersionsPath+"/{version}", h.getServiceVersion)

		s.HandleFunc(internal.VmClassesPath, h.vmClasses)
		s.HandleFunc(internal.VmClassesPath+"/", h.vmClasses)
	}
}

// enabledClusters returns refs for cluster names with a "WCP-" prefix.
// Using the name as a simple hack until we add support for enabling via the API.
func enabledClusters(c *govmomi.Client) ([]types.ManagedObjectReference, error) {
	ctx := context.Background()
	kind := []string{"ClusterComputeResource"}

	m := view.NewManager(c.Client)
	v, err := m.CreateContainerView(ctx, c.ServiceContent.RootFolder, kind, true)
	if err != nil {
		return nil, err
	}
	defer func() { _ = v.Destroy(ctx) }()

	return v.Find(ctx, kind, property.Match{"name": "WCP-*"})
}

func (h *Handler) clusters(w http.ResponseWriter, r *http.Request) {
	c, err := govmomi.NewClient(context.Background(), h.URL, true)
	if err != nil {
		panic(err)
	}

	switch r.Method {
	case http.MethodGet:
		refs, err := enabledClusters(c)
		if err != nil {
			panic(err)
		}

		clusters := make([]namespace.ClusterSummary, len(refs))
		for i, ref := range refs {
			clusters[i] = namespace.ClusterSummary{
				ID:               ref.Value,
				ConfigStatus:     &namespace.RunningConfigStatus,
				KubernetesStatus: &namespace.ReadyKubernetesStatus,
			}
		}
		vapi.StatusOK(w, clusters)
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

func (h *Handler) clustersID(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	route := map[string]func(http.ResponseWriter, *http.Request){
		"support-bundle": h.clustersSupportBundle,
	}[id]

	if route != nil {
		route(w, r)
		return
	}

	// TODO:
	// https://vmware.github.io/vsphere-automation-sdk-rest/vsphere/index.html#SVC_com.vmware.vcenter.namespace_management.clusters
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

// Some fake services, service 1 has 2 versions, service 2 has 1 version
// Summary for services
var supervisorServicesMap = map[string]namespace.SupervisorServiceSummary{
	"service1": {
		ID:    "service1",
		Name:  "mock-service-1",
		State: "ACTIVATED",
	},
	"service2": {
		ID:    "service2",
		Name:  "mock-service-2",
		State: "DEACTIVATED",
	},
}

// Summary for service versions
var supervisorServiceVersionsMap = map[string][]namespace.SupervisorServiceVersionSummary{
	"service1": {
		{
			SupervisorServiceInfo: namespace.SupervisorServiceInfo{
				Name:        "mock-service-1 v1 display name",
				State:       "ACTIVATED",
				Description: "This is service 1 version 1.0.0",
			},
			Version: "1.0.0",
		},
		{
			SupervisorServiceInfo: namespace.SupervisorServiceInfo{
				Name:        "mock-service-1 v2 display name",
				State:       "DEACTIVATED",
				Description: "This is service 1 version 2.0.0",
			},
			Version: "2.0.0",
		}},
	"service2": {
		{
			SupervisorServiceInfo: namespace.SupervisorServiceInfo{
				Name:        "mock-service-2 v1 display name",
				State:       "ACTIVATED",
				Description: "This is service 2 version 1.1.0",
			},
			Version: "1.1.0",
		},
	},
}

func (h *Handler) listServices(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		supervisorServices := make([]namespace.SupervisorServiceSummary, len(supervisorServicesMap))
		i := 0
		for _, service := range supervisorServicesMap {
			supervisorServices[i] = service
			i++
		}
		vapi.StatusOK(w, supervisorServices)
	}
}

func (h *Handler) getService(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	switch r.Method {
	case http.MethodGet:
		if result, contains := supervisorServicesMap[id]; contains {
			vapi.StatusOK(w, result)
		} else {
			vapi.ApiErrorNotFound(w)
		}
		return
	}
}

// versionsForService returns the list of versions for a particular service
func (h *Handler) versionsForService(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("In versionsForService: %v\n", r.URL.Path)
	id := r.PathValue("id")
	switch r.Method {
	case http.MethodGet:
		if result, contains := supervisorServiceVersionsMap[id]; contains {
			vapi.StatusOK(w, result)
		} else {
			vapi.ApiErrorNotFound(w)
		}
		return
	}
}

// getServiceVersion returns info on a particular service version
func (h *Handler) getServiceVersion(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	version := r.PathValue("version")

	switch r.Method {
	case http.MethodGet:
		if versions, contains := supervisorServiceVersionsMap[id]; contains {
			for _, v := range versions {
				if v.Version == version {
					info := namespace.SupervisorServiceVersionInfo{
						SupervisorServiceInfo: namespace.SupervisorServiceInfo{
							Description: v.Description,
							State:       v.State,
							Name:        v.Name,
						},
						ContentType: "CARVEL_APPS_YAML", // return Carvel by default
						Content:     "abc",              // in reality this is base 64 encoded of content
					}
					vapi.StatusOK(w, info)
					return
				}
			}
			vapi.ApiErrorNotFound(w)
		} else {
			vapi.ApiErrorNotFound(w)
		}
		return
	}
}

var namespacesMap = make(map[string]*namespace.NamespacesInstanceInfo)

func (h *Handler) namespaces(w http.ResponseWriter, r *http.Request) {
	subpath := r.URL.Path[len(internal.NamespacesPath):]
	subpath = strings.TrimPrefix(subpath, "/")
	// TODO: move to 1.22's https://go.dev/blog/routing-enhancements
	route := strings.Split(subpath, "/")
	subpath = route[0]
	action := ""
	if len(route) > 1 {
		action = route[1]
	}

	switch r.Method {
	case http.MethodGet:
		if len(subpath) > 0 {
			if result, contains := namespacesMap[subpath]; contains {
				vapi.StatusOK(w, result)
			} else {
				vapi.ApiErrorNotFound(w)
			}
			return
		} else {
			result := make([]namespace.NamespacesInstanceSummary, 0, len(namespacesMap))

			for k, v := range namespacesMap {
				entry := namespace.NamespacesInstanceSummary{
					ClusterId:    v.ClusterId,
					Namespace:    k,
					ConfigStatus: v.ConfigStatus,
					Description:  v.Description,
					Stats:        v.Stats,
				}
				result = append(result, entry)
			}

			vapi.StatusOK(w, result)
		}
	case http.MethodPatch:
		if len(subpath) > 0 {
			if entry, contains := namespacesMap[subpath]; contains {
				var spec namespace.NamespacesInstanceUpdateSpec
				if vapi.Decode(r, w, &spec) {
					entry.VmServiceSpec = spec.VmServiceSpec
					vapi.StatusOK(w)
				}
			}
		}

		vapi.ApiErrorNotFound(w)
	case http.MethodPost:
		if action == "registervm" {
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
			return
		}

		var spec namespace.NamespacesInstanceCreateSpec
		if !vapi.Decode(r, w, &spec) {
			return
		}

		newNamespace := namespace.NamespacesInstanceInfo{
			ClusterId:     spec.Cluster,
			ConfigStatus:  namespace.RunningConfigStatus.String(),
			VmServiceSpec: spec.VmServiceSpec,
		}

		namespacesMap[spec.Namespace] = &newNamespace

		vapi.StatusOK(w)
	case http.MethodDelete:
		if len(subpath) > 0 {
			if _, contains := namespacesMap[subpath]; contains {
				delete(namespacesMap, subpath)
				vapi.StatusOK(w)
				return
			}
		}
		vapi.ApiErrorNotFound(w)
	}
}

var vmClassesMap = make(map[string]*namespace.VirtualMachineClassInfo)

func (h *Handler) vmClasses(w http.ResponseWriter, r *http.Request) {
	subpath := r.URL.Path[len(internal.VmClassesPath):]
	subpath = strings.Replace(subpath, "/", "", -1)

	switch r.Method {
	case http.MethodGet:
		if len(subpath) > 0 {
			if result, contains := vmClassesMap[subpath]; contains {
				vapi.StatusOK(w, result)
			} else {
				vapi.ApiErrorNotFound(w)
			}
			return
		} else {
			result := make([]*namespace.VirtualMachineClassInfo, 0, len(vmClassesMap))

			for _, v := range vmClassesMap {
				result = append(result, v)
			}

			vapi.StatusOK(w, result)
		}
	case http.MethodPatch:
		if len(subpath) > 0 {
			if entry, contains := vmClassesMap[subpath]; contains {
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
				return
			}
		}

		vapi.ApiErrorNotFound(w)
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

		vmClassesMap[spec.Id] = &newClass

		vapi.StatusOK(w)
	case http.MethodDelete:
		if len(subpath) > 0 {
			if _, contains := vmClassesMap[subpath]; contains {
				delete(vmClassesMap, subpath)
				vapi.StatusOK(w)
				return
			}
		}
		vapi.ApiErrorNotFound(w)
	}
}
