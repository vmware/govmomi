// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"context"
	"reflect"
	"testing"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/task"
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
	c := m.Service.client()

	finder := find.NewFinder(c, false)
	dc, _ := finder.DatacenterList(ctx, "*")
	finder.SetDatacenter(dc[0])
	folders, _ := dc[0].Folders(ctx)
	hosts, _ := finder.HostSystemList(ctx, "*/*")
	vswitch := m.Map().Any("DistributedVirtualSwitch").(*DistributedVirtualSwitch)
	dvs0 := object.NewDistributedVirtualSwitch(c, vswitch.Reference())

	if len(vswitch.Summary.HostMember) == 0 {
		t.Fatal("no host member")
	}

	for _, ref := range vswitch.Summary.HostMember {
		host := m.Map().Get(ref).(*HostSystem)
		if len(host.Network) == 0 {
			t.Fatalf("%s.Network=%v", ref, host.Network)
		}
		parent := hostParent(m.Service.Context, &host.HostSystem)
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
	c := m.Service.client()

	finder := find.NewFinder(c, false)
	dc, _ := finder.DatacenterList(ctx, "*")
	finder.SetDatacenter(dc[0])
	vswitch := m.Map().Any("DistributedVirtualSwitch").(*DistributedVirtualSwitch)
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
				{PortgroupKey: pgs[1].Value, Key: "0"},
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
				{PortgroupKey: pgs[1].Value, Key: "0"},
			},
		},
		{
			"PortKeys",
			&types.DistributedVirtualSwitchPortCriteria{
				PortKey: []string{"1"},
			},
			[]types.DistributedVirtualPort{},
		},
		{
			"connected",
			&types.DistributedVirtualSwitchPortCriteria{
				Connected: types.NewBool(true),
			},
			[]types.DistributedVirtualPort{},
		},
		{
			"not connected",
			&types.DistributedVirtualSwitchPortCriteria{
				Connected: types.NewBool(false),
			},
			[]types.DistributedVirtualPort{
				{PortgroupKey: pgs[0].Value, Key: "0"},
				{PortgroupKey: pgs[1].Value, Key: "0"},
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
