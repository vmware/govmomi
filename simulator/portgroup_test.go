// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"context"
	"strings"
	"testing"

	"github.com/google/uuid"

	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/types"
)

func TestReconfigurePortgroup(t *testing.T) {
	m := VPX()

	err := m.Create()
	if err != nil {
		t.Fatal(err)
	}

	defer m.Remove()

	c := m.Service.client()
	ctx := m.Service.Context

	dvs := object.NewDistributedVirtualSwitch(c,
		ctx.Map.Any("DistributedVirtualSwitch").Reference())

	spec := []types.DVPortgroupConfigSpec{
		{
			Name:     "pg1",
			NumPorts: 10,
		},
	}

	task, err := dvs.AddPortgroup(ctx, spec)
	if err != nil {
		t.Fatal(err)
	}

	err = task.Wait(ctx)
	if err != nil {
		t.Fatal(err)
	}

	pg := object.NewDistributedVirtualPortgroup(c,
		ctx.Map.Any("DistributedVirtualPortgroup").Reference())
	pgspec := types.DVPortgroupConfigSpec{
		NumPorts: 5,
		Name:     "pg1",
	}

	task, err = pg.Reconfigure(ctx, pgspec)
	if err != nil {
		t.Fatal(err)
	}

	err = task.Wait(ctx)
	if err != nil {
		t.Fatal(err)
	}

	pge := ctx.Map.Get(pg.Reference()).(*DistributedVirtualPortgroup)
	if pge.Config.Name != "pg1" || pge.Config.NumPorts != 5 {
		t.Fatalf("expect pg.Name==pg1 && pg.Config.NumPort==5; got %s,%d",
			pge.Config.Name, pge.Config.NumPorts)
	}

	task, err = pg.Destroy(ctx)
	if err != nil {
		t.Fatal(err)
	}

	err = task.Wait(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPortgroupBacking(t *testing.T) {
	ctx := context.Background()

	m := VPX()

	err := m.Create()
	if err != nil {
		t.Fatal(err)
	}

	defer m.Remove()

	c := m.Service.client()

	pg := m.Map().Any("DistributedVirtualPortgroup").(*DistributedVirtualPortgroup)

	net := object.NewDistributedVirtualPortgroup(c, pg.Reference())
	t.Logf("pg=%s", net.Reference())

	_, err = net.EthernetCardBackingInfo(ctx)
	if err != nil {
		t.Fatal(err)
	}

	// "This property should always be set unless the user's setting does not have System.Read privilege on the object referred to by this property."
	// Test that we return an error in this case, rather than panic.
	pg.Config.DistributedVirtualSwitch = nil
	_, err = net.EthernetCardBackingInfo(ctx)
	if err == nil {
		t.Error("expected error")
	}
}

func TestPortgroupBackingWithNSX(t *testing.T) {
	model := VPX()
	model.Portgroup = 0
	model.PortgroupNSX = 1

	Test(func(ctx context.Context, _ *vim25.Client) {
		pgs := Map(ctx).All("DistributedVirtualPortgroup")
		n := len(pgs) - 1
		if model.PortgroupNSX != n {
			t.Errorf("%d pgs", n)
		}

		for _, obj := range pgs {
			pg := obj.(*DistributedVirtualPortgroup)
			if strings.Contains(pg.Name, "DVUplinks") {
				continue
			}

			if pg.Config.BackingType != "nsx" {
				t.Errorf("backing=%q", pg.Config.BackingType)
			}

			_, err := uuid.Parse(pg.Config.LogicalSwitchUuid)
			if err != nil {
				t.Errorf("parsing %q: %s", pg.Config.LogicalSwitchUuid, err)
			}
		}
	}, model)
}
