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
	"path"

	"github.com/google/uuid"
	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vapi/rest"
	vapi "github.com/vmware/govmomi/vapi/simulator"
	"github.com/vmware/govmomi/vim25/types"

	"github.com/vmware/govmomi/vapi/cluster"
	"github.com/vmware/govmomi/vapi/cluster/internal"
)

func init() {
	simulator.RegisterEndpoint(func(s *simulator.Service, r *simulator.Registry) {
		New(s.Listen).Register(s, r)
	})
}

type module struct {
	cluster.ModuleSummary
	members map[string]bool
}

// Handler implements the Cluster Modules API simulator
type Handler struct {
	Modules map[string]module
	URL     *url.URL
}

// New creates a Handler instance
func New(u *url.URL) *Handler {
	return &Handler{
		Modules: make(map[string]module),
		URL:     u,
	}
}

// Register Cluster Modules API paths with the vapi simulator's http.ServeMux
func (h *Handler) Register(s *simulator.Service, r *simulator.Registry) {
	if r.IsVPX() {
		s.HandleFunc(rest.Path+internal.ModulesPath, h.modules)
		s.HandleFunc(rest.Path+internal.ModulesPath+"/", h.modules)
		s.HandleFunc(rest.Path+internal.ModulesVMPath+"/", h.modulesVM)
	}
}

func (h *Handler) modules(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		var modules cluster.ModuleSummaryList
		for _, s := range h.Modules {
			modules.Summaries = append(modules.Summaries, s.ModuleSummary)
		}
		vapi.OK(w, modules)
	case http.MethodPost:
		var m internal.CreateModule
		if vapi.Decode(r, w, &m) {
			ref := types.ManagedObjectReference{Type: "ClusterComputeResource", Value: m.Spec.ID}
			if simulator.Map.Get(ref) == nil {
				vapi.BadRequest(w, "com.vmware.vapi.std.errors.invalid_argument")
				return
			}

			id := uuid.New().String()
			h.Modules[id] = module{
				cluster.ModuleSummary{
					Cluster: m.Spec.ID,
					Module:  id,
				},
				make(map[string]bool),
			}
			vapi.OK(w, id)
		}
	case http.MethodDelete:
		id := path.Base(r.URL.Path)
		_, ok := h.Modules[id]
		if !ok {
			http.NotFound(w, r)
			return
		}
		delete(h.Modules, id)
		vapi.OK(w)
	}
}

func (*Handler) action(r *http.Request) string {
	return r.URL.Query().Get("action")
}

func (h *Handler) addMembers(members internal.ModuleMembers, m module) bool {
	cluster := types.ManagedObjectReference{Type: "ClusterComputeResource", Value: m.Cluster}
	c, err := govmomi.NewClient(context.Background(), h.URL, true)
	if err != nil {
		panic(err)
	}
	vms, err := internal.ClusterVM(c.Client, cluster)
	if err != nil {
		panic(err)
	}
	_ = c.Logout(context.Background())

	validVM := func(id string) bool {
		for i := range vms {
			if vms[i].Reference().Value == id {
				return true
			}
		}
		return false
	}

	for _, id := range members.VMs {
		if m.members[id] {
			return false
		}
		if !validVM(id) {
			return false
		}
		m.members[id] = true
	}
	return true
}

func (h *Handler) removeMembers(members internal.ModuleMembers, m module) bool {
	for _, id := range members.VMs {
		if !m.members[id] {
			return false
		}
		delete(m.members, id)
	}
	return true
}

func (h *Handler) modulesVM(w http.ResponseWriter, r *http.Request) {
	p := path.Dir(r.URL.Path)
	id := path.Base(p)

	m, ok := h.Modules[id]
	if !ok {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		var members internal.ModuleMembers
		for member := range m.members {
			members.VMs = append(members.VMs, member)
		}
		vapi.OK(w, members)
	case http.MethodPost:
		action := h.addMembers

		switch h.action(r) {
		case "add":
		case "remove":
			action = h.removeMembers
		default:
			http.NotFound(w, r)
			return
		}

		var status internal.Status
		var members internal.ModuleMembers
		if vapi.Decode(r, w, &members) {
			status.Success = action(members, m)
			vapi.OK(w, status)
		}
	}
}
