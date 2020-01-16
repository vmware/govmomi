/*
Copyright (c) 2017 VMware, Inc. All Rights Reserved.

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
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/types"
)

func TestReconfigurePortgroup(t *testing.T) {
	ctx := context.Background()

	m := VPX()

	err := m.Create()
	if err != nil {
		t.Fatal(err)
	}

	defer m.Remove()

	c := m.Service.client

	dvs := object.NewDistributedVirtualSwitch(c,
		Map.Any("DistributedVirtualSwitch").Reference())

	spec := []types.DVPortgroupConfigSpec{
		types.DVPortgroupConfigSpec{
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
		Map.Any("DistributedVirtualPortgroup").Reference())
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

	pge := Map.Get(pg.Reference()).(*DistributedVirtualPortgroup)
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

	c := m.Service.client

	pg := Map.Any("DistributedVirtualPortgroup").(*DistributedVirtualPortgroup)

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

	Test(func(context.Context, *vim25.Client) {
		pgs := Map.All("DistributedVirtualPortgroup")
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
