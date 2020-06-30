/*
Copyright (c) 2020 VMware, Inc. All Rights Reserved.

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
	"context"
	"net/http"
	"net/url"

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
				ConfigStatus:     "RUNNING",
				KubernetesStatus: "READY",
			}
		}
		vapi.StatusOK(w, clusters)
	}
}
