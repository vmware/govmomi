// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package policies_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"slices"
	"testing"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vapi/rest"
	"github.com/vmware/govmomi/vapi/vcenter/vm/compute/policies"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"

	_ "github.com/vmware/govmomi/vapi/simulator"
	_ "github.com/vmware/govmomi/vapi/vcenter/vm/compute/policies/simulator"
)

func TestTopologyJSON(t *testing.T) {
	var info policies.PolicyInfo
	data := `{"status":"COMPLIANT","policy":"p1","name":"n","capability":"c","topology":{"@enum":"Host"}}`
	if err := json.Unmarshal([]byte(data), &info); err != nil {
		t.Fatal(err)
	}
	if got := types.VmPlacementPolicyVmPlacementPolicyTopology(info.Topology); got != types.VmPlacementPolicyVmPlacementPolicyTopologyHost {
		t.Errorf("topology=%q", got)
	}

	out, err := json.Marshal(info.Topology)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != `{"@enum":"Host"}` {
		t.Errorf("marshal=%s", out)
	}
}

// vmIDs returns the managed object IDs of the VMs in the simulator inventory.
func vmIDs(ctx context.Context, t *testing.T, vc *vim25.Client) []string {
	t.Helper()
	vms, err := find.NewFinder(vc).VirtualMachineList(ctx, "*")
	if err != nil {
		t.Fatal(err)
	}
	if len(vms) < 2 {
		t.Fatalf("need at least 2 VMs, have %d", len(vms))
	}
	ids := make([]string, len(vms))
	for i, vm := range vms {
		ids[i] = vm.Reference().Value
	}
	return ids
}

func login(ctx context.Context, t *testing.T, vc *vim25.Client) *policies.Manager {
	t.Helper()
	rc := rest.NewClient(vc)
	if err := rc.Login(ctx, simulator.DefaultLogin); err != nil {
		t.Fatal(err)
	}
	return policies.NewManager(rc)
}

func TestQuery(t *testing.T) {
	simulator.Test(func(ctx context.Context, vc *vim25.Client) {
		m := login(ctx, t, vc)
		ids := vmIDs(ctx, t, vc)

		res, err := m.Query(ctx, policies.QuerySpec{VMs: []string{ids[0], "vm-none"}})
		if err != nil {
			t.Fatal(err)
		}
		if len(res.Items) != 1 || res.Items[0].VM != ids[0] || len(res.Items[0].Policies) != 2 {
			t.Fatalf("items=%+v", res.Items)
		}
		if _, ok := res.Errors["vm-none"]; !ok || len(res.Errors) != 1 {
			t.Errorf("errors=%v", res.Errors)
		}
		if res.Marker != "" {
			t.Errorf("marker=%q", res.Marker)
		}
		for _, p := range res.Items[0].Policies {
			if p.Topology != policies.Topology(types.VmPlacementPolicyVmPlacementPolicyTopologyHost) {
				t.Errorf("topology=%q", p.Topology)
			}
		}

		// Filtering by status
		res, err = m.Query(ctx, policies.QuerySpec{
			VMs:    ids[:1],
			Filter: &policies.FilterSpec{Statuses: []policies.ObjectCompliance{policies.ObjectComplianceNotCompliant}},
		})
		if err != nil {
			t.Fatal(err)
		}
		ps := res.Items[0].Policies
		if len(ps) != 1 || ps[0].Status != policies.ObjectComplianceNotCompliant {
			t.Errorf("policies=%+v", ps)
		}

		// No VMs is invalid
		if _, err = m.Query(ctx, policies.QuerySpec{}); err == nil {
			t.Error("expected error for empty VM list")
		}
	})
}

func TestQueryAll(t *testing.T) {
	simulator.Test(func(ctx context.Context, vc *vim25.Client) {
		m := login(ctx, t, vc)
		ids := vmIDs(ctx, t, vc)
		vms := append(slices.Clone(ids), "vm-none")

		size := int64(1)
		spec := policies.QuerySpec{VMs: vms, Iterate: &policies.IterationSpec{PageSize: &size}}

		// A single page is limited by the page size
		page, err := m.Query(ctx, spec)
		if err != nil {
			t.Fatal(err)
		}
		if len(page.Items) != 1 || page.Marker == "" {
			t.Fatalf("items=%d marker=%q", len(page.Items), page.Marker)
		}

		// QueryAll follows the markers and combines the pages
		res, err := m.QueryAll(ctx, spec)
		if err != nil {
			t.Fatal(err)
		}
		if len(res.Items) != len(ids) {
			t.Errorf("items=%d, want %d", len(res.Items), len(ids))
		}
		for i, item := range res.Items {
			if item.VM != ids[i] {
				t.Errorf("item %d: vm=%s, want %s", i, item.VM, ids[i])
			}
		}
		if _, ok := res.Errors["vm-none"]; !ok || len(res.Errors) != 1 {
			t.Errorf("errors=%v", res.Errors)
		}
		if res.Marker != "" {
			t.Errorf("marker=%q", res.Marker)
		}
	})
}

// The simulator never repeats a marker, so use a misbehaving server to
// verify that QueryAll gives up instead of looping forever.
func TestQueryAllRepeatedMarker(t *testing.T) {
	calls := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		_, _ = w.Write([]byte(`{"items":[],"marker":"stuck"}`))
	}))
	defer srv.Close()

	u, err := url.Parse(srv.URL)
	if err != nil {
		t.Fatal(err)
	}
	m := policies.NewManager(rest.NewClient(&vim25.Client{Client: soap.NewClient(u, true)}))

	if _, err := m.QueryAll(context.Background(), policies.QuerySpec{VMs: []string{"vm-1"}}); err == nil {
		t.Fatal("expected error for repeated marker")
	}
	if calls != 2 {
		t.Errorf("calls=%d, want 2", calls)
	}
}
