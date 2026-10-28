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

// Handler implements the Namespace Management Modules API simulator. Handler
// methods own HTTP concerns only; all state lives behind store in store.go.
type Handler struct {
	URL   *url.URL
	store *store
}

// New creates a Handler instance
func New(u *url.URL) *Handler {
	return &Handler{
		URL:   u,
		store: newStore(),
	}
}

// Register Namespace Management API paths with the vapi simulator's http.ServeMux
func (h *Handler) Register(s *simulator.Service, r *simulator.Registry) {
	if r.IsVPX() {
		h.store.registry = r

		s.HandleFunc(internal.NamespacesPath, h.namespaces)
		s.HandleFunc(internal.NamespacesPath+"/{namespace}", h.namespace)
		s.HandleFunc(internal.NamespacesPath+"/{namespace}/{action}", h.namespaceAction)
		s.HandleFunc(internal.NamespacesPathV2, h.namespacesV2)
		s.HandleFunc(internal.NamespacesPathV2+"/{namespace}", h.namespaceV2)
		s.HandleFunc(internal.NamespaceClusterPath, h.clusters)
		s.HandleFunc(internal.NamespaceClusterPath+"/{id}", h.cluster)
		s.HandleFunc(internal.NamespaceClusterPath+"/{id}/support-bundle", h.clustersSupportBundle)
		s.HandleFunc(internal.NamespaceClusterPath+"/{id}/networks", h.clusterNetworks)
		s.HandleFunc(internal.NamespaceDistributedSwitchCompatibility, h.compatibleDistributedSwitches)
		s.HandleFunc(internal.NamespaceEdgeClusterCompatibility, h.compatibleEdgeClusters)

		s.HandleFunc(internal.SupervisorsSummariesPath, h.supervisorSummaries)
		s.HandleFunc(internal.SupervisorsPath+"/{id}/summary", h.supervisorSummary)
		s.HandleFunc(internal.SupervisorsPath+"/{id}/topology", h.supervisorTopology)

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
		vapi.StatusOK(w, h.store.listClusters())
	}
}

func (h *Handler) cluster(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	switch r.Method {
	case http.MethodPost:
		switch r.URL.Query().Get("action") {
		case "enable":
			h.store.enableCluster(id)
			vapi.StatusOK(w)
		case "disable":
			if !h.store.disableCluster(id) {
				vapi.ApiErrorNotFound(w)
				return
			}
			vapi.StatusOK(w)
		default:
			vapi.ApiErrorInvalidArgument(w)
		}
	}
}

func (h *Handler) supervisorSummaries(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		vapi.StatusOK(w, namespace.SupervisorsSummaryResult{
			Items: h.store.listSupervisors(),
		})
	}
}

func (h *Handler) supervisorSummary(w http.ResponseWriter, r *http.Request) {
	s, exists := h.store.getSupervisor(r.PathValue("id"))
	if !exists {
		vapi.ApiErrorNotFound(w)
		return
	}

	switch r.Method {
	case http.MethodGet:
		vapi.StatusOK(w, s.info)
	}
}

func (h *Handler) supervisorTopology(w http.ResponseWriter, r *http.Request) {
	s, exists := h.store.getSupervisor(r.PathValue("id"))
	if !exists {
		vapi.ApiErrorNotFound(w)
		return
	}

	switch r.Method {
	case http.MethodGet:
		vapi.StatusOK(w, s.topology)
	}
}

func (h *Handler) clusterNetworks(w http.ResponseWriter, r *http.Request) {
	s, exists := h.store.getSupervisor(r.PathValue("id"))
	if !exists {
		vapi.ApiErrorNotFound(w)
		return
	}

	switch r.Method {
	case http.MethodGet:
		vapi.StatusOK(w, s.networks)
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

func (h *Handler) compatibleDistributedSwitches(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) compatibleEdgeClusters(w http.ResponseWriter, r *http.Request) {
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
		vapi.StatusOK(w, h.store.listServices())
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
		if !h.store.addService(pkg, content) {
			vapi.ApiErrorAlreadyExists(w)
			return
		}
		vapi.StatusOK(w)
	}
}

func (h *Handler) service(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	info, exists := h.store.getService(id)
	if !exists {
		vapi.ApiErrorNotFound(w)
		return
	}

	switch r.Method {
	case http.MethodGet:
		vapi.StatusOK(w, info)
	case http.MethodPatch:
		state, ok := resolveServiceState(r.URL.Query().Get("action"))
		if !ok {
			vapi.ApiErrorInvalidArgument(w)
			return
		}
		h.store.setServiceState(id, state)
		vapi.StatusOK(w)
	case http.MethodDelete:
		if info.State != deactivatedState {
			vapi.ApiErrorNotAllowedInCurrentState(w)
			return
		}
		h.store.removeService(id)
		vapi.StatusOK(w)
	}
}

func (h *Handler) serviceVersions(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if !h.store.hasServiceVersions(id) {
		vapi.ApiErrorNotFound(w)
		return
	}

	switch r.Method {
	case http.MethodGet:
		vapi.StatusOK(w, h.store.listServiceVersions(id))
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
		if !h.store.addServiceVersion(id, pkg, spec.CarvelService.Content) {
			vapi.ApiErrorAlreadyExists(w)
			return
		}
		vapi.StatusOK(w)
	}
}

func (h *Handler) serviceVersion(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	version := r.PathValue("version")

	info, exists := h.store.getServiceVersion(id, version)
	if !exists {
		vapi.ApiErrorNotFound(w)
		return
	}

	switch r.Method {
	case http.MethodGet:
		vapi.StatusOK(w, info)
	case http.MethodPatch:
		state, ok := resolveServiceState(r.URL.Query().Get("action"))
		if !ok {
			vapi.ApiErrorInvalidArgument(w)
			return
		}
		h.store.setServiceVersionState(id, version, state)
		vapi.StatusOK(w)
	case http.MethodDelete:
		if info.State != deactivatedState {
			vapi.ApiErrorNotAllowedInCurrentState(w)
			return
		}
		h.store.removeServiceVersion(id, version)
		vapi.StatusOK(w)
	}
}

const (
	activatedState   = "ACTIVATED"
	deactivatedState = "DEACTIVATED"
)

func resolveServiceState(action string) (string, bool) {
	switch action {
	case "activate":
		return activatedState, true
	case "deactivate":
		return deactivatedState, true
	default:
		return "", false
	}
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

func (h *Handler) namespaces(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		vapi.StatusOK(w, h.store.listNamespaces())
	case http.MethodPost:
		var spec namespace.NamespacesInstanceCreateSpec
		if !vapi.Decode(r, w, &spec) {
			return
		}

		if !h.store.addNamespace(spec) {
			vapi.ApiErrorAlreadyExists(w)
			return
		}
		vapi.StatusOK(w)
	}
}

func (h *Handler) namespace(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("namespace")

	info, exists := h.store.getNamespace(name)
	if !exists {
		vapi.ApiErrorNotFound(w)
		return
	}

	switch r.Method {
	case http.MethodGet:
		vapi.StatusOK(w, convertToNamespacesInstanceInfo(info))
	case http.MethodPatch:
		var spec namespace.NamespacesInstanceUpdateSpec
		if !vapi.Decode(r, w, &spec) {
			return
		}

		h.store.updateNamespace(name, spec)
		vapi.StatusOK(w)
	case http.MethodDelete:
		h.store.removeNamespace(name)
		vapi.StatusOK(w)
	}
}

func (h *Handler) namespacesV2(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		vapi.StatusOK(w, h.store.listNamespacesV2())
	case http.MethodPost:
		var spec namespace.NamespaceInstanceCreateSpecV2
		if !vapi.Decode(r, w, &spec) {
			return
		}

		if !h.store.addNamespaceV2(spec) {
			vapi.ApiErrorAlreadyExists(w)
			return
		}
		vapi.StatusOK(w)
	}
}

func (h *Handler) namespaceV2(w http.ResponseWriter, r *http.Request) {
	info, exists := h.store.getNamespace(r.PathValue("namespace"))
	if !exists {
		vapi.ApiErrorNotFound(w)
		return
	}

	switch r.Method {
	case http.MethodGet:
		vapi.StatusOK(w, info)
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
		vapi.StatusOK(w, h.store.listVmClasses())
	case http.MethodPost:
		var spec namespace.VirtualMachineClassCreateSpec
		if !vapi.Decode(r, w, &spec) {
			return
		}

		if !h.store.addVmClass(spec) {
			vapi.ApiErrorAlreadyExists(w)
			return
		}
		vapi.StatusOK(w)
	}
}

func (h *Handler) vmClass(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("vmClass")

	switch r.Method {
	case http.MethodGet:
		result, exists := h.store.getVmClass(id)
		if !exists {
			vapi.ApiErrorNotFound(w)
			return
		}
		vapi.StatusOK(w, result)
	case http.MethodPatch:
		if _, exists := h.store.getVmClass(id); !exists {
			vapi.ApiErrorNotFound(w)
			return
		}

		var spec namespace.VirtualMachineClassUpdateSpec
		if !vapi.Decode(r, w, &spec) {
			return
		}

		h.store.updateVmClass(id, spec)
		vapi.StatusOK(w)
	case http.MethodDelete:
		if _, exists := h.store.getVmClass(id); !exists {
			vapi.ApiErrorNotFound(w)
			return
		}
		h.store.removeVmClass(id)
		vapi.StatusOK(w)
	}
}
