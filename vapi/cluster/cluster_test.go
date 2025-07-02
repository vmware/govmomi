// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package cluster_test

import (
	"context"
	"testing"

	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vapi/rest"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"

	"github.com/vmware/govmomi/vapi/cluster"
	"github.com/vmware/govmomi/vapi/cluster/internal"

	_ "github.com/vmware/govmomi/vapi/cluster/simulator"
	_ "github.com/vmware/govmomi/vapi/simulator"
)

var enoent = types.ManagedObjectReference{Value: "enoent"}

func TestClusterModules(t *testing.T) {
	simulator.Test(func(ctx context.Context, vc *vim25.Client) {
		c := rest.NewClient(vc)

		err := c.Login(ctx, simulator.DefaultLogin)
		if err != nil {
			t.Fatal(err)
		}

		m := cluster.NewManager(c)
		modules, err := m.ListModules(ctx)
		if err != nil {
			t.Fatal(err)
		}

		if len(modules) != 0 {
			t.Errorf("expected 0 modules")
		}

		ccr := simulator.Map(ctx).Any("ClusterComputeResource")

		_, err = m.CreateModule(ctx, enoent)
		if err == nil {
			t.Fatal("expected error")
		}

		id, err := m.CreateModule(ctx, ccr)
		if err != nil {
			t.Fatal(err)
		}

		modules, err = m.ListModules(ctx)
		if err != nil {
			t.Fatal(err)
		}

		if len(modules) != 1 {
			t.Errorf("expected 1 module")
		}

		err = m.DeleteModule(ctx, "enoent")
		if err == nil {
			t.Fatal("expected error")
		}

		err = m.DeleteModule(ctx, id)
		if err != nil {
			t.Fatal(err)
		}

		modules, err = m.ListModules(ctx)
		if err != nil {
			t.Fatal(err)
		}

		if len(modules) != 0 {
			t.Errorf("expected 0 modules")
		}
	})
}

func TestClusterModuleMembers(t *testing.T) {
	simulator.Test(func(ctx context.Context, vc *vim25.Client) {
		c := rest.NewClient(vc)

		err := c.Login(ctx, simulator.DefaultLogin)
		if err != nil {
			t.Fatal(err)
		}

		m := cluster.NewManager(c)

		_, err = m.ListModuleMembers(ctx, "enoent")
		if err == nil {
			t.Error("expected error")
		}

		ccr := simulator.Map(ctx).Any("ClusterComputeResource")

		id, err := m.CreateModule(ctx, ccr)
		if err != nil {
			t.Fatal(err)
		}

		vms, err := internal.ClusterVM(vc, ccr)
		if err != nil {
			t.Fatal(err)
		}

		expect := []struct {
			n       int
			success bool
			action  func(context.Context, string, ...mo.Reference) (bool, error)
			ids     []mo.Reference
		}{
			{0, false, m.AddModuleMembers, []mo.Reference{enoent}},
			{0, false, m.RemoveModuleMembers, []mo.Reference{enoent}},
			{len(vms), true, m.AddModuleMembers, vms},
			{len(vms), false, m.AddModuleMembers, vms},
			{0, true, m.RemoveModuleMembers, vms},
			{len(vms), false, m.AddModuleMembers, append(vms, enoent)},
			{len(vms) - 1, false, m.RemoveModuleMembers, []mo.Reference{vms[0], enoent}},
		}

		for i, test := range expect {
			ok, err := test.action(ctx, id, test.ids...)
			if err != nil {
				t.Fatal(err)
			}
			if ok != test.success {
				t.Errorf("%d: success=%t", i, ok)
			}

			members, err := m.ListModuleMembers(ctx, id)
			if err != nil {
				t.Fatal(err)
			}

			if len(members) != test.n {
				t.Errorf("%d: members=%d", i, len(members))
			}
		}
	})
}
