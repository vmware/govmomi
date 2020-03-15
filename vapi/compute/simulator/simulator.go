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
	"net/http"
	"net/url"
	"path"

	"github.com/google/uuid"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vapi/compute"
	"github.com/vmware/govmomi/vapi/rest"
	vapi "github.com/vmware/govmomi/vapi/simulator"

	"github.com/vmware/govmomi/vapi/compute/internal"
)

func init() {
	simulator.RegisterEndpoint(func(s *simulator.Service, r *simulator.Registry) {
		New(s.Listen).Register(s, r)
	})
}

// captured via govc compute.policy.ls -c -dump
var capabilities = []compute.Capability{
	{
		Capability:  "com.vmware.vcenter.compute.policies.capabilities.vm_host_anti_affinity",
		Name:        "VM-Host anti-affinity",
		Description: "Virtual machines that have the VM tag will not be placed on hosts that have the host tag.",
	},
	{
		Capability:  "com.vmware.vcenter.compute.policies.capabilities.vm_vm_affinity",
		Name:        "VM-VM affinity",
		Description: "All virtual machines that share the tag will be affined to each other.",
	},
	{
		Capability:  "com.vmware.vcenter.compute.policies.capabilities.vm_host_affinity",
		Name:        "VM-Host affinity",
		Description: "Virtual machines that have the VM tag will be placed on hosts that have the host tag.",
	},
	{
		Capability:  "com.vmware.vcenter.compute.policies.capabilities.vm.evacuation.vmotion",
		Name:        "VM evacuation by vMotion",
		Description: "Virtual machines that have the VM tag will be vMotioned when their host is evacuated.",
	},
	{
		Capability:  "com.vmware.vcenter.compute.policies.capabilities.vm_vm_anti_affinity",
		Name:        "VM-VM anti-affinity",
		Description: "All virtual machines that share the tag will be anti-affine to each other.",
	},
	{
		Capability:  "com.vmware.vcenter.compute.policies.capabilities.disable_drs_vmotion",
		Name:        "Disable DRS vMotion",
		Description: "All virtual machines that share the tag will stay on the host on which they are powered-on, except when the host is put into maintenance mode or failed over.",
	},
	{
		Capability:  "com.vmware.vcenter.compute.policies.capabilities.cluster_scale_in_ignore_vm_capabilities",
		Name:        "Scale-in ignore VM capabilities",
		Description: "When considering scaling-in a cluster, policies that have been created with the listed capabilities are ignored for virtual machines that have the tag.",
	},
}

// Handler implements the Cluster Policies API simulator
type Handler struct {
	Policies map[string]compute.Policy
	URL      *url.URL
}

// New creates a Handler instance
func New(u *url.URL) *Handler {
	return &Handler{
		Policies: make(map[string]compute.Policy),
		URL:      u,
	}
}

// Register Cluster Policies API paths with the vapi simulator's http.ServeMux
func (h *Handler) Register(s *simulator.Service, r *simulator.Registry) {
	if r.IsVPX() {
		s.HandleFunc(rest.Path+internal.PolicyPath, h.policies)
		s.HandleFunc(rest.Path+internal.PolicyPath+"/", h.policiesID)
		s.HandleFunc(rest.Path+internal.PolicyCapabilitiesPath, h.capabilities)
		s.HandleFunc(rest.Path+internal.PolicyCapabilitiesPath+"/", h.capabilitiesID)
	}
}

func capability(id string) *compute.Capability {
	for _, c := range capabilities {
		if c.Capability == id {
			return &c
		}
	}
	return nil
}

func (h *Handler) policies(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		var res []compute.Policy
		for _, p := range h.Policies {
			res = append(res, p)
		}
		vapi.OK(w, res)
	case http.MethodPost:
		var spec struct {
			compute.Policy `json:"spec"`
		}

		if !vapi.Decode(r, w, &spec) {
			return
		}

		c := capability(spec.Capability)

		tag := spec.VMTag
		if tag == "" {
			tag = spec.HostTag
		}
		// TODO: validate tag

		if c == nil || tag == "" || spec.Name == "" || spec.Description == "" {
			vapi.BadRequest(w, "com.vmware.vapi.std.errors.invalid_argument")
			return
		}

		id := uuid.New().String()
		h.Policies[id] = compute.Policy{
			Capability:  spec.Capability,
			Description: spec.Description,
			Name:        spec.Name,
			Policy:      id,
		}

		vapi.OK(w, id)
	}
}

func (h *Handler) policiesID(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	p, ok := h.Policies[id]
	if !ok {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		vapi.OK(w, p)
	case http.MethodDelete:
		delete(h.Policies, id)
		vapi.OK(w)
	}
}

func (h *Handler) capabilities(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	vapi.OK(w, capabilities)
}

func (h *Handler) capabilitiesID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	id := path.Base(r.URL.Path)
	c := capability(id)
	if c == nil {
		http.NotFound(w, r)
		return
	}

	vapi.OK(w, c)
}
