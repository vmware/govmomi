// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

// Package simulator implements the com.vmware.vcenter.vm.compute.Policies API.
package simulator

import (
	"encoding/json"
	"net/http"
	"slices"
	"sort"
	"strconv"

	"github.com/vmware/govmomi/simulator"
	vapi "github.com/vmware/govmomi/vapi/simulator"
	"github.com/vmware/govmomi/vapi/vcenter/vm/compute/policies"
	"github.com/vmware/govmomi/vim25/types"
)

const (
	path = "/api/vcenter/vm/compute/policies"

	// DefaultPageSize is the number of VMs returned per page when the request
	// does not specify a page size.
	DefaultPageSize = 100
)

// Policies that the simulator reports for every VM that exists.
var vmPolicies = []policies.PolicyInfo{
	{
		Status:     policies.ObjectComplianceCompliant,
		Policy:     "policy-compliant",
		Name:       "Compliant policy",
		Capability: "com.vmware.vcenter.compute.capability.vm-host-affinity",
		Strictness: policies.PolicyStrictnessRequiredDuringPlacementPreferredDuringExecution,
		Topology:   policies.Topology(types.VmPlacementPolicyVmPlacementPolicyTopologyHost),
		Tag:        "urn:vmomi:InventoryServiceTag:00000000-0000-0000-0000-000000000001:GLOBAL",
	},
	{
		Status:     policies.ObjectComplianceNotCompliant,
		Policy:     "policy-not-compliant",
		Name:       "Non-compliant policy",
		Capability: "com.vmware.vcenter.compute.capability.vm-vm-anti-affinity",
		Strictness: policies.PolicyStrictnessRequiredDuringPlacementRequiredDuringExecution,
		Topology:   policies.Topology(types.VmPlacementPolicyVmPlacementPolicyTopologyHost),
	},
}

func init() {
	simulator.RegisterEndpoint(func(s *simulator.Service, r *simulator.Registry) {
		if r.IsVPX() {
			s.HandleFunc(path, handler{r}.query)
		}
	})
}

type handler struct {
	*simulator.Registry
}

func (h handler) query(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost || r.URL.Query().Get("action") != "query" {
		http.NotFound(w, r)
		return
	}

	var spec policies.QuerySpec
	if !vapi.Decode(r, w, &spec) {
		return
	}
	if len(spec.VMs) == 0 {
		vapi.ApiErrorInvalidArgument(w)
		return
	}

	// The marker is the index of the next VM to return.
	start, size := 0, DefaultPageSize
	if it := spec.Iterate; it != nil {
		if it.Marker != "" {
			n, err := strconv.Atoi(it.Marker)
			if err != nil || n < 0 {
				vapi.ApiErrorInvalidArgument(w)
				return
			}
			start = n
		}
		if it.PageSize != nil {
			if *it.PageSize < 1 {
				vapi.ApiErrorInvalidArgument(w)
				return
			}
			size = int(*it.PageSize)
		}
	}

	res := policies.QueryResult{
		Items:  []policies.VmPoliciesInfo{},
		Errors: map[string]json.RawMessage{},
	}

	vms := spec.VMs
	if start > len(vms) {
		start = len(vms)
	}
	end := min(start+size, len(vms))
	for _, id := range vms[start:end] {
		ref := types.ManagedObjectReference{Type: "VirtualMachine", Value: id}
		if h.Get(ref) == nil {
			res.Errors[id] = json.RawMessage(`{"error_type":"NOT_FOUND","messages":[]}`)
			continue
		}
		res.Items = append(res.Items, policies.VmPoliciesInfo{
			VM:       id,
			Policies: filter(spec.Filter),
		})
	}
	if end < len(vms) {
		res.Marker = strconv.Itoa(end)
	}

	vapi.StatusOK(w, res)
}

// filter returns the simulated policies that match f, ordered by identifier.
func filter(f *policies.FilterSpec) []policies.PolicyInfo {
	out := []policies.PolicyInfo{}
	for _, p := range vmPolicies {
		if f != nil &&
			(!match(f.Policies, p.Policy) ||
				!match(f.Capabilities, p.Capability) ||
				!match(f.Statuses, p.Status) ||
				!match(f.Strictnesses, p.Strictness) ||
				!match(f.Topologies, types.VmPlacementPolicyVmPlacementPolicyTopology(p.Topology))) {
			continue
		}
		out = append(out, p)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Policy < out[j].Policy })
	return out
}

// match reports whether v is in want, or want is empty.
func match[T comparable](want []T, v T) bool {
	return len(want) == 0 || slices.Contains(want, v)
}
