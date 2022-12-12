/*
Copyright (c) 2020-2022 VMware, Inc. All Rights Reserved.

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

package simulator

import (
	"archive/tar"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/google/uuid"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vapi/namespace"
	vapi "github.com/vmware/govmomi/vapi/simulator"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25/types"

	"github.com/vmware/govmomi/vapi/namespace/internal"
)

func init() {
	simulator.RegisterEndpoint(func(s *simulator.Service, r *simulator.Registry) {
		New(s.Listen).Register(s, r)
	})
}

// Handler implements the Cluster Modules API simulator
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
		s.HandleFunc(internal.NamespaceClusterPath, h.clusters)
		s.HandleFunc(internal.NamespaceClusterPath+"/", h.clustersID)
		s.HandleFunc(internal.NamespaceDistributedSwitchCompatibility+"/", h.listCompatibleDistributedSwitches)
		s.HandleFunc(internal.NamespaceEdgeClusterCompatibility+"/", h.listCompatibleEdgeClusters)

		s.HandleFunc(internal.SupervisorServicesPath, h.listServices)
		s.HandleFunc(internal.SupervisorServicesPath+"/", h.getService)

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

	return v.Find(ctx, kind, property.Filter{"name": "WCP-*"})
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

var supervisorServices []namespace.SupervisorServiceSummary = []namespace.SupervisorServiceSummary{
	{
		ID:    "service1",
		Name:  "mock-service-1",
		State: "ACTIVATED",
	},
	{
		ID:    "service2",
		Name:  "mock-service-2",
		State: "DE-ACTIVATED",
	},
}

func (h *Handler) listServices(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		vapi.StatusOK(w, supervisorServices)
	}
}

func (h *Handler) getService(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	switch r.Method {
	case http.MethodGet:
		for _, svc := range supervisorServices {
			if svc.ID == id {
				svcInfo := namespace.SupervisorServiceInfo{
					Name:        svc.Name,
					State:       svc.State,
					Description: fmt.Sprintf("Description of %s", svc.ID),
				}
				vapi.StatusOK(w, svcInfo)
				return
			}
		}
		w.WriteHeader(http.StatusNotFound)
	}
}
