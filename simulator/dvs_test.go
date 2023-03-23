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
	"reflect"
	"testing"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/task"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

func TestDVS(t *testing.T) {
	m := VPX()

	defer m.Remove()

	err := m.Create()
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	c := m.Service.client

	finder := find.NewFinder(c, false)
	dc, _ := finder.DatacenterList(ctx, "*")
	finder.SetDatacenter(dc[0])
	folders, _ := dc[0].Folders(ctx)
	hosts, _ := finder.HostSystemList(ctx, "*/*")
	vswitch := Map.Any("DistributedVirtualSwitch").(*DistributedVirtualSwitch)
	dvs0 := object.NewDistributedVirtualSwitch(c, vswitch.Reference())

	if len(vswitch.Summary.HostMember) == 0 {
		t.Fatal("no host member")
	}

	for _, ref := range vswitch.Summary.HostMember {
		host := Map.Get(ref).(*HostSystem)
		if len(host.Network) == 0 {
			t.Fatalf("%s.Network=%v", ref, host.Network)
		}
		parent := hostParent(&host.HostSystem)
		if len(parent.Network) != len(host.Network) {
			t.Fatalf("%s.Network=%v", parent.Reference(), parent.Network)
		}
	}

	var spec types.DVSCreateSpec
	spec.ConfigSpec = &types.VMwareDVSConfigSpec{}
	spec.ConfigSpec.GetDVSConfigSpec().Name = "DVS1"

	dtask, err := folders.NetworkFolder.CreateDVS(ctx, spec)
	if err != nil {
		t.Fatal(err)
	}

	info, err := dtask.WaitForResult(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}

	dvs := object.NewDistributedVirtualSwitch(c, info.Result.(types.ManagedObjectReference))

	config := &types.DVSConfigSpec{}

	for _, host := range hosts {
		config.Host = append(config.Host, types.DistributedVirtualSwitchHostMemberConfigSpec{
			Host: host.Reference(),
		})
	}

	tests := []struct {
		op  types.ConfigSpecOperation
		pg  string
		err types.BaseMethodFault
	}{
		{types.ConfigSpecOperationAdd, "", nil},                               // Add == OK
		{types.ConfigSpecOperationAdd, "", &types.AlreadyExists{}},            // Add == fail (AlreadyExists)
		{types.ConfigSpecOperationEdit, "", &types.NotSupported{}},            // Edit == fail (NotSupported)
		{types.ConfigSpecOperationRemove, "", nil},                            // Remove == OK
		{types.ConfigSpecOperationAdd, "", nil},                               // Add == OK
		{types.ConfigSpecOperationAdd, "DVPG0", nil},                          // Add PG == OK
		{types.ConfigSpecOperationRemove, "", &types.ResourceInUse{}},         // Remove dvs0 == fail (ResourceInUse)
		{types.ConfigSpecOperationRemove, "", nil},                            // Remove dvs1 == OK (no VMs attached)
		{types.ConfigSpecOperationRemove, "", &types.ManagedObjectNotFound{}}, // Remove == fail (ManagedObjectNotFound)
	}

	for x, test := range tests {
		dswitch := dvs

		switch test.err.(type) {
		case *types.ManagedObjectNotFound:
			for i := range config.Host {
				config.Host[i].Host.Value = "enoent"
			}
		case *types.ResourceInUse:
			dswitch = dvs0
		}

		if test.pg == "" {
			for i := range config.Host {
				config.Host[i].Operation = string(test.op)
			}

			dtask, err = dswitch.Reconfigure(ctx, config)
		} else {
			switch test.op {
			case types.ConfigSpecOperationAdd:
				dtask, err = dswitch.AddPortgroup(ctx, []types.DVPortgroupConfigSpec{
					{Name: test.pg, NumPorts: 1},
				})
			}
		}

		if err != nil {
			t.Fatal(err)
		}

		err = dtask.Wait(ctx)

		if test.err == nil {
			if err != nil {
				t.Fatalf("%d: %s", x, err)
			}
			continue
		}

		if err == nil {
			t.Errorf("expected error in test %d", x)
		}

		if reflect.TypeOf(test.err) != reflect.TypeOf(err.(task.Error).Fault()) {
			t.Errorf("expected %T fault in test %d", test.err, x)
		}
	}

	ports, err := dvs.FetchDVPorts(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(ports) != 2 {
		t.Fatalf("expected 2 ports in DVPorts; got %d", len(ports))
	}

	dtask, err = dvs.Destroy(ctx)
	if err != nil {
		t.Fatal(err)
	}

	err = dtask.Wait(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFetchDVPortsCriteria(t *testing.T) {
	m := VPX()

	defer m.Remove()

	err := m.Create()
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	c := m.Service.client

	finder := find.NewFinder(c, false)
	dc, _ := finder.DatacenterList(ctx, "*")
	finder.SetDatacenter(dc[0])
	vswitch := Map.Any("DistributedVirtualSwitch").(*DistributedVirtualSwitch)
	dvs0 := object.NewDistributedVirtualSwitch(c, vswitch.Reference())
	pgs := vswitch.Portgroup
	if len(pgs) != 2 {
		t.Fatalf("expected 2 portgroups in DVS; got %d", len(pgs))
	}

	tests := []struct {
		name     string
		criteria *types.DistributedVirtualSwitchPortCriteria
		expected []types.DistributedVirtualPort
	}{
		{
			"empty criteria",
			&types.DistributedVirtualSwitchPortCriteria{},
			[]types.DistributedVirtualPort{
				{PortgroupKey: pgs[0].Value, Key: "0"},
				{PortgroupKey: pgs[1].Value, Key: "1"},
			},
		},
		{
			"inside PortgroupKeys",
			&types.DistributedVirtualSwitchPortCriteria{
				PortgroupKey: []string{pgs[0].Value},
				Inside:       types.NewBool(true),
			},
			[]types.DistributedVirtualPort{
				{PortgroupKey: pgs[0].Value, Key: "0"},
			},
		},
		{
			"outside PortgroupKeys",
			&types.DistributedVirtualSwitchPortCriteria{
				PortgroupKey: []string{pgs[0].Value},
				Inside:       types.NewBool(false),
			},
			[]types.DistributedVirtualPort{
				{PortgroupKey: pgs[1].Value, Key: "1"},
			},
		},
		{
			"PortKeys",
			&types.DistributedVirtualSwitchPortCriteria{
				PortKey: []string{"1"},
			},
			[]types.DistributedVirtualPort{
				{PortgroupKey: pgs[1].Value, Key: "1"},
			},
		},
		{
			"connected",
			&types.DistributedVirtualSwitchPortCriteria{
				Connected: types.NewBool(true),
			},
			[]types.DistributedVirtualPort{
				{PortgroupKey: pgs[1].Value, Key: "1"},
			},
		},
		{
			"not connected",
			&types.DistributedVirtualSwitchPortCriteria{
				Connected: types.NewBool(false),
			},
			[]types.DistributedVirtualPort{
				{PortgroupKey: pgs[0].Value, Key: "0"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := dvs0.FetchDVPorts(context.TODO(), test.criteria)

			if err != nil {
				t.Fatal(err)
			}

			if len(actual) != len(test.expected) {
				t.Fatalf("expected %d ports; got %d", len(test.expected), len(actual))
			}

			for i, p := range actual {
				if p.Key != test.expected[i].Key {
					t.Errorf("ports[%d]: expected Key `%s`; got `%s`",
						i, test.expected[i].Key, p.Key)
				}

				if p.PortgroupKey != test.expected[i].PortgroupKey {
					t.Errorf("ports[%d]: expected PortgroupKey `%s`; got `%s`",
						i, test.expected[i].PortgroupKey, p.PortgroupKey)
				}
			}
		})
	}
}

func TestDVSAddHostToSpecificPortgroup(t *testing.T) {
	m := VPX()
	defer m.Remove()

	err := m.Create()
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	c := m.Service.client

	finder := find.NewFinder(c, false)
	dc, _ := finder.DatacenterList(ctx, "*")
	finder.SetDatacenter(dc[0])
	folders, _ := dc[0].Folders(ctx)
	hosts, _ := finder.HostSystemList(ctx, "*/*")

	var spec types.DVSCreateSpec
	spec.ConfigSpec = &types.VMwareDVSConfigSpec{}
	spec.ConfigSpec.GetDVSConfigSpec().Name = "DVS1"

	dtask, err := folders.NetworkFolder.CreateDVS(ctx, spec)
	if err != nil {
		t.Fatal(err)
	}

	info, err := dtask.WaitForResult(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}

	dvs := object.NewDistributedVirtualSwitch(c, info.Result.(types.ManagedObjectReference))

	dswitch := dvs
	pgs := []struct {
		portGroupName string
		uplink        bool
	}{
		{"upg0", true}, //Add hosts to specified portgroup
		{"", false},    //Remove hosts from DVS
		{"upg1", true}, //Add hosts to specified portgroup
	}
	for _, pgDet := range pgs {
		var (
			pgKey     string
			pg        mo.DistributedVirtualPortgroup
			portGroup *object.DistributedVirtualPortgroup
			backing   *types.DistributedVirtualSwitchHostMemberPnicBacking
		)

		//Create Configuration
		config := &types.DVSConfigSpec{}
		if pgDet.portGroupName != "" {
			dtask, err = dswitch.AddPortgroup(ctx, []types.DVPortgroupConfigSpec{
				{Name: pgDet.portGroupName, NumPorts: 1, Uplink: &pgDet.uplink}})
			if err != nil {
				t.Fatal(err)
			}

			err = dtask.Wait(ctx)
			if err != nil {
				t.Fatalf("%s", err)
			}

			prtgrp, err := finder.Network(ctx, pgDet.portGroupName)
			if err != nil {
				t.Fatalf("%s", err)
			}
			var ok bool
			portGroup, ok = prtgrp.(*object.DistributedVirtualPortgroup)
			if !ok {
				t.Fatalf("failed to convert %T to %T", prtgrp, portGroup)
			}
			err = portGroup.Properties(ctx, portGroup.Reference(), []string{"config"}, &pg)
			if err != nil {
				t.Fatalf("%s", err)
			}
			pgKey = pg.Config.Key

			backing = new(types.DistributedVirtualSwitchHostMemberPnicBacking)
			backing.PnicSpec = append(backing.PnicSpec, types.DistributedVirtualSwitchHostMemberPnicSpec{
				PnicDevice:         "vmnic0",
				UplinkPortgroupKey: pgKey,
			})
		}

		//Apply Configuration
		for _, host := range hosts {
			config.Host = append(config.Host,
				types.DistributedVirtualSwitchHostMemberConfigSpec{Host: host.Reference()})
		}
		operation := types.ConfigSpecOperationAdd
		if pgDet.portGroupName == "" {
			operation = types.ConfigSpecOperationRemove
		}

		for i := range config.Host {
			config.Host[i].Operation = string(operation)
			config.Host[i].Backing = backing
		}

		dtask, err = dswitch.Reconfigure(ctx, config)
		if err != nil {
			t.Fatalf("%s", err)
		}
		err = dtask.Wait(ctx)
		if err != nil {
			t.Fatalf("%s", err)
		}

		//Validate Configuration
		if pgDet.portGroupName != "" {
			prtgrps, err := finder.NetworkList(ctx, "upg*")
			if err != nil {
				t.Fatalf("%s", err)
			}

			fetchedHosts := make(map[string]map[string]struct{})
			for _, prtgrp := range prtgrps {
				portGroup, ok := prtgrp.(*object.DistributedVirtualPortgroup)
				if !ok {
					t.Fatalf("failed to convert %T to %T", prtgrp, portGroup)
				}
				err = portGroup.Properties(ctx, portGroup.Reference(), []string{"config", "host"}, &pg)
				if err != nil {
					t.Fatalf("%s", err)
				}
				fetchedHosts[pg.Config.Name] = make(map[string]struct{})
				for _, hobj := range pg.Host {
					var h mo.HostSystem
					err := object.NewHostSystem(c, hobj).Properties(ctx, hobj.Reference(),
						[]string{"managedEntity"}, &h)
					if err != nil {
						t.Fatalf("%s", err)
					}
					hostKey := h.ExtensibleManagedObject.Self.Value
					fetchedHosts[pg.Config.Name][hostKey] = struct{}{}
				}
				for pgName, fHosts := range fetchedHosts {
					if pgName == pgDet.portGroupName { //All hosts should get added in this pg
						if len(fHosts) != len(hosts) {
							t.Fatalf("fetched hosts %v does not match with %v\n",
								len(fHosts), len(hosts))
						}
					} else { //No host should get added to other pgs
						if len(fHosts) != 0 {
							t.Fatalf("fetched hosts %v found in portgroup %v\n",
								fHosts, pgName)
						}
					}
				}
			}
		}
	}

	dtask, err = dvs.Destroy(ctx)
	if err != nil {
		t.Fatal(err)
	}

	err = dtask.Wait(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDVPortsConnecteeDetails(t *testing.T) {
	m := VPX()
	defer m.Remove()

	err := m.Create()
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	c := m.Service.client

	finder := find.NewFinder(c, false)
	dc, _ := finder.DatacenterList(ctx, "*")
	finder.SetDatacenter(dc[0])
	folders, _ := dc[0].Folders(ctx)
	hosts, _ := finder.HostSystemList(ctx, "*/*")

	var spec types.DVSCreateSpec
	spec.ConfigSpec = &types.VMwareDVSConfigSpec{}
	spec.ConfigSpec.GetDVSConfigSpec().Name = "DVS1"

	dtask, err := folders.NetworkFolder.CreateDVS(ctx, spec)
	if err != nil {
		t.Fatal(err)
	}

	info, err := dtask.WaitForResult(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}

	dvs := object.NewDistributedVirtualSwitch(c, info.Result.(types.ManagedObjectReference))

	dswitch := dvs
	var pg mo.DistributedVirtualPortgroup

	uplink := true
	pgName := "upg0"

	config := &types.DVSConfigSpec{}
	dtask, err = dswitch.AddPortgroup(ctx, []types.DVPortgroupConfigSpec{
		{Name: pgName, NumPorts: int32(len(hosts)), Uplink: &uplink}})
	if err != nil {
		t.Fatal(err)
	}
	err = dtask.Wait(ctx)
	if err != nil {
		t.Fatalf("%s", err)
	}

	prtgrp, err := finder.Network(ctx, pgName)
	if err != nil {
		t.Fatalf("%s", err)
	}
	portGroup, ok := prtgrp.(*object.DistributedVirtualPortgroup)
	if !ok {
		t.Fatalf("failed to convert %T to %T", prtgrp, portGroup)
	}
	err = portGroup.Properties(ctx, portGroup.Reference(), []string{"config"}, &pg)
	if err != nil {
		t.Fatalf("%s", err)
	}
	backing := new(types.DistributedVirtualSwitchHostMemberPnicBacking)
	backing.PnicSpec = append(backing.PnicSpec, types.DistributedVirtualSwitchHostMemberPnicSpec{
		PnicDevice:         "vmnic0",
		UplinkPortgroupKey: pg.Config.Key,
	})

	expHosts := make(map[types.ManagedObjectReference]struct{})
	for _, host := range hosts {
		expHosts[host.Reference()] = struct{}{}
		config.Host = append(config.Host,
			types.DistributedVirtualSwitchHostMemberConfigSpec{Host: host.Reference()})
	}
	operation := types.ConfigSpecOperationAdd

	for i := range config.Host {
		config.Host[i].Operation = string(operation)
		config.Host[i].Backing = backing
	}

	dtask, err = dswitch.Reconfigure(ctx, config)
	if err != nil {
		t.Fatalf("%s", err)
	}
	err = dtask.Wait(ctx)
	if err != nil {
		t.Fatalf("%s", err)
	}

	dvps, err := dswitch.FetchDVPorts(ctx, nil)
	if err != nil {
		t.Fatalf("%s", err)
	}
	for _, dvp := range dvps {
		if dvp.Connectee != nil {
			delete(expHosts, *dvp.Connectee.ConnectedEntity)
		}
	}
	if len(expHosts) != 0 {
		t.Fatalf("some hosts %v are not present in DVP connected entity", expHosts)
	}

}
