// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package types_test

import (
	"context"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"testing"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/session"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

func insecureFromEnv() bool {
	v := os.Getenv("GOVC_INSECURE")
	if v == "" {
		// Verify TLS certs.
		return false
	}
	// 1 - Skip TLS validation. 0 - Enforce TLS validation.
	b, _ := strconv.ParseBool(v)
	return b
}

func credsFromEnvOrURL(u *url.URL, t *testing.T) *url.Userinfo {
	t.Helper()
	if u.User != nil {
		return u.User
	}
	user := os.Getenv("GOVC_USERNAME")
	if user == "" {
		t.Fatal("no credentials provided: set user:pass in GOVC_URL or GOVC_USERNAME/GOVC_PASSWORD")
	}
	return url.UserPassword(user, os.Getenv("GOVC_PASSWORD"))
}

// Scoped override for the *types* registry only.
func withTypesOverride(name string, to reflect.Type, fn func() error) error {
	// Lookup the current registered type for that VMODL name.
	prev, _ := types.TypeFunc()(name)
	// Temporarily override mapping (e.g. ResourcePoolRuntimeInfo → ResourcePoolRuntimeInfoEx)
	types.Add(name, to)
	// Restore original mapping at the end.
	defer types.Add(name, prev)
	// Run the block while override is in effect.
	return fn()
}

func TestResourcePoolRuntimeWithVmRp(t *testing.T) {
	rawURL := os.Getenv("GOVC_URL")
	if rawURL == "" {
		t.Skip("GOVC_URL not set")
	}

	ctx := context.Background()

	u, err := soap.ParseURL(rawURL)
	if err != nil {
		t.Fatalf("parse GOVC_URL: %v", err)
	}

	// SOAP + vim25 client (honor GOVC_INSECURE)
	sc := soap.NewClient(u, insecureFromEnv())
	c, err := vim25.NewClient(ctx, sc)
	if err != nil {
		t.Fatalf("vim25.NewClient: %v", err)
	}

	err = c.UseServiceVersion()
	if err != nil {
		t.Fatalf("vim25.NewClient use service version: %v", err)
	}

	// Login
	userinfo := credsFromEnvOrURL(u, t)
	m := session.NewManager(c)
	if err := m.Login(ctx, userinfo); err != nil {
		t.Fatalf("login failed: %v", err)
	}
	defer m.Logout(ctx)

	// Get Datacenter & Resource pool by env ----
	dcName := os.Getenv("GOVC_DATACENTER")
	if dcName == "" {
		t.Fatal("set GOVC_DATACENTER")
	}
	rpPath := os.Getenv("GOVC_RESOURCE_POOL")
	if rpPath == "" {
		t.Fatal("set GOVC_RESOURCE_POOL (full inventory path)")
	}

	f := find.NewFinder(c, true)
	dc, err := f.Datacenter(ctx, dcName)
	if err != nil {
		t.Fatalf("Datacenter(%q): %v", dcName, err)
	}
	f.SetDatacenter(dc)
	t.Logf("using datacenter: %s", dc.InventoryPath)

	rpObj, err := f.ResourcePool(ctx, rpPath)
	if err != nil {
		t.Fatalf("ResourcePool(%q): %v", rpPath, err)
	}
	rpRef := rpObj.Reference()
	t.Logf("using resource pool: %s", rpObj.InventoryPath)

	// ---- Retrieve "runtime.vmRp" as ObjectContent and decode via *types* override ----
	props := []string{"runtime.vmRp"}
	var oc []types.ObjectContent
	pc := property.DefaultCollector(c)

	err = withTypesOverride("ResourcePoolRuntimeInfo", reflect.TypeOf(types.ResourcePoolRuntimeInfoEx{}), func() error {
		return pc.Retrieve(ctx, []types.ManagedObjectReference{rpRef}, props, &oc)
	})
	if err != nil {
		t.Fatalf("Retrieve(runtime.vmRp): %v", err)
	}
	if len(oc) == 0 || len(oc[0].PropSet) == 0 {
		t.Fatalf("no properties returned for %s", rpRef.String())
	}

	for _, p := range oc[0].PropSet {
		t.Logf("Property: %s, Type: %T, Value: %#v", p.Name, p.Val, p.Val)
	}
	// Find "runtime" and assert it's the Ex type; then dump vmRp
	var got any
	for _, p := range oc[0].PropSet {
		if p.Name == "runtime.vmRp" {
			got = p.Val
			break
		}
	}
	if got == nil {
		t.Fatalf("runtime not in PropSet")
	}

	t.Logf("runtime dynamic type: %T", got)

	var items []types.ResourcePoolVmResourceProfileUsage

	switch v := got.(type) {
	case types.ArrayOfResourcePoolVmResourceProfileUsage:
		// Leaf property "runtime.vmRp" often decodes to the wrapper struct.
		items = v.ResourcePoolVmResourceProfileUsage
	case *types.ArrayOfResourcePoolVmResourceProfileUsage:
		items = v.ResourcePoolVmResourceProfileUsage
	case []types.ResourcePoolVmResourceProfileUsage:
		// When decoding the whole "runtime" into ResourcePoolRuntimeInfoEx,
		// the slice field (VmRp []...) is filled directly.
		items = v
	default:
		t.Fatalf("unexpected type for runtime.vmRp: %T", got)
	}

	if len(items) == 0 {
		t.Log("vmRp present but empty")
	} else {
		t.Logf("vmRp items: %d", len(items))
		for i, u := range items {
			t.Logf("vmRp[%d]: id=%q reservedForPool=%d usedForVms=%d usedForChildPools=%d",
				i, u.Id, u.ReservedForPool, u.ReservationUsedForVms, u.ReservationUsedForChildPools)
		}
	}
}
