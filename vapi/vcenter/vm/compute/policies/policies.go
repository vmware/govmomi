// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

// Package policies provides a client for the com.vmware.vcenter.vm.compute.Policies
// service, which reports the compliance of virtual machines with external
// compute policies.
package policies

import (
	"context"
	"encoding/json"
	"fmt"
	"maps"
	"net/http"

	"github.com/vmware/govmomi/vapi/rest"
	"github.com/vmware/govmomi/vim25/types"
)

const basePath = "/api/vcenter/vm/compute/policies"

// ObjectCompliance is the compliance state of a policy on an object.
// com.vmware.vcenter.compute.policies.ObjectCompliance
type ObjectCompliance string

const (
	// ObjectComplianceUnknown means the object has an unknown compliance state.
	ObjectComplianceUnknown ObjectCompliance = "UNKNOWN"
	// ObjectComplianceNotApplicable means the policy does not apply to the object.
	ObjectComplianceNotApplicable ObjectCompliance = "NOT_APPLICABLE"
	// ObjectComplianceCompliant means the object is compliant with the policy.
	ObjectComplianceCompliant ObjectCompliance = "COMPLIANT"
	// ObjectComplianceNotCompliant means the object is not compliant with the policy.
	ObjectComplianceNotCompliant ObjectCompliance = "NOT_COMPLIANT"
)

// PolicyStrictness describes how strictly a policy is enforced.
// com.vmware.vcenter.compute.policies.PolicyStrictness
type PolicyStrictness string

const (
	PolicyStrictnessPreferredDuringPlacementPreferredDuringExecution PolicyStrictness = "PREFERRED_DURING_PLACEMENT_PREFERRED_DURING_EXECUTION"
	PolicyStrictnessRequiredDuringPlacementPreferredDuringExecution  PolicyStrictness = "REQUIRED_DURING_PLACEMENT_PREFERRED_DURING_EXECUTION"
	PolicyStrictnessRequiredDuringPlacementRequiredDuringExecution   PolicyStrictness = "REQUIRED_DURING_PLACEMENT_REQUIRED_DURING_EXECUTION"
)

// Manager extends rest.Client, adding VM compute policy related methods.
type Manager struct {
	*rest.Client
}

// NewManager creates a new Manager instance with the given client.
func NewManager(client *rest.Client) *Manager {
	return &Manager{
		Client: client,
	}
}

// FilterSpec filters the policies returned for each virtual machine.
// Multiple non-empty fields are combined with AND semantics.
// com.vmware.vcenter.vm.compute.Policies.FilterSpec
type FilterSpec struct {
	// Policies must have one of these identifiers.
	Policies []string `json:"policies,omitempty"`
	// Policies must be instances of one of these capabilities.
	Capabilities []string `json:"capabilities,omitempty"`
	// Policies must have one of these compliance states.
	Statuses []ObjectCompliance `json:"statuses,omitempty"`
	// Policies must have one of these strictness values.
	Strictnesses []PolicyStrictness `json:"strictnesses,omitempty"`
	// Policies must have one of these topology values, as described by
	// vim.vm.VmPlacementPolicy.VmPlacementPolicyTopology.
	Topologies []types.VmPlacementPolicyVmPlacementPolicyTopology `json:"topologies,omitempty"`
}

// IterationSpec contains the fields used to page through query results.
// com.vmware.vcenter.vm.compute.Policies.IterationSpec
type IterationSpec struct {
	// Marker is the opaque marker returned by a previous query. The caller
	// must resend the same VMs and Filter as in the request that produced it.
	Marker string `json:"marker,omitempty"`
	// PageSize is the maximum number of VMs returned per call.
	// The valid range is defined by the service. If unset, the service default
	// is used.
	PageSize *int64 `json:"page_size,omitempty"`
}

// QuerySpec selects the virtual machines whose policy compliance is returned.
// com.vmware.vcenter.vm.compute.Policies.QuerySpec
type QuerySpec struct {
	// VMs to query. Must not be empty.
	VMs []string `json:"vms"`
	// Filter applied to the policies returned for each VM.
	Filter *FilterSpec `json:"filter,omitempty"`
	// Iterate holds the pagination parameters.
	Iterate *IterationSpec `json:"iterate,omitempty"`
}

// PolicyInfo describes the compliance of a VM with one external policy.
// com.vmware.vcenter.vm.compute.Policies.PolicyInfo
type PolicyInfo struct {
	// Status is the compliance of the VM with the policy.
	Status ObjectCompliance `json:"status"`
	// Policy is the identifier of the policy.
	Policy string `json:"policy"`
	// Name of the policy.
	Name string `json:"name"`
	// Capability is the identifier of the capability the policy is based on.
	Capability string `json:"capability"`
	// Strictness of the policy, unset if the capability does not support it.
	Strictness PolicyStrictness `json:"strictness,omitempty"`
	// Topology of the policy, unset if the capability does not support it.
	// Same values as the topology of a VM placement policy in a config spec.
	Topology Topology `json:"topology,omitempty"`
	// Tag is the identifier of the tag used by the policy, if any.
	Tag string `json:"tag,omitempty"`
}

// Topology is a vim.vm.VmPlacementPolicy.VmPlacementPolicyTopology value.
// vCenter returns it as an enum object, {"@enum": "Host"}, rather than as a
// plain string.
type Topology types.VmPlacementPolicyVmPlacementPolicyTopology

// MarshalJSON encodes the value as an {"@enum": value} object.
func (t Topology) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Enum string `json:"@enum"`
	}{string(t)})
}

// UnmarshalJSON decodes an {"@enum": value} object.
func (t *Topology) UnmarshalJSON(data []byte) error {
	var v struct {
		Enum string `json:"@enum"`
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	*t = Topology(v.Enum)
	return nil
}

// VmPoliciesInfo holds the policy information of one successfully queried VM.
// com.vmware.vcenter.vm.compute.Policies.VmPoliciesInfo
type VmPoliciesInfo struct {
	// VM is the identifier of the virtual machine.
	VM string `json:"vm"`
	// Policies lists all associated external policies. Empty if there are none.
	Policies []PolicyInfo `json:"policies"`
}

// QueryResult is a page of query results.
// com.vmware.vcenter.vm.compute.Policies.QueryResult
type QueryResult struct {
	// Items are the successfully queried VMs in this page.
	Items []VmPoliciesInfo `json:"items"`
	// Errors maps a requested VM identifier to the error for that VM. A VM
	// appears either in Items or in Errors, never both. Values are vAPI error
	// structures, see rest.Error.
	Errors map[string]json.RawMessage `json:"errors"`
	// Marker for the next page, empty if there are no more results.
	Marker string `json:"marker,omitempty"`
}

// Query returns a single page of policy compliance for the VMs in spec.
// Only associated external policies are returned. The operation reads stored
// status and does not trigger compliance evaluation.
// POST /vcenter/vm/compute/policies?action=query
func (c *Manager) Query(ctx context.Context, spec QuerySpec) (*QueryResult, error) {
	path := c.Resource(basePath).WithParam("action", "query")
	req := path.Request(http.MethodPost, spec)
	var res QueryResult
	if err := c.Do(ctx, req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// QueryAll calls Query repeatedly, following the returned marker, and returns
// the combined results of all pages. spec.Iterate.PageSize is honored if set;
// spec.Iterate.Marker is used as the starting marker.
func (c *Manager) QueryAll(ctx context.Context, spec QuerySpec) (*QueryResult, error) {
	var iter IterationSpec
	if spec.Iterate != nil {
		iter = *spec.Iterate
	}
	all := &QueryResult{Errors: map[string]json.RawMessage{}}
	for {
		spec.Iterate = &iter
		page, err := c.Query(ctx, spec)
		if err != nil {
			return nil, err
		}
		all.Items = append(all.Items, page.Items...)
		maps.Copy(all.Errors, page.Errors)
		if page.Marker == "" {
			return all, nil
		}
		if page.Marker == iter.Marker {
			return nil, fmt.Errorf("policies: service returned the same marker %q twice", page.Marker)
		}
		iter.Marker = page.Marker
	}
}
