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
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/task"
	"github.com/vmware/govmomi/vim25"
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
	c := m.Service.client()

	finder := find.NewFinder(c, false)
	dc, _ := finder.DatacenterList(ctx, "*")
	finder.SetDatacenter(dc[0])
	folders, _ := dc[0].Folders(ctx)
	hosts, _ := finder.HostSystemList(ctx, "*/*")
	vswitch := m.Map().Any("VmwareDistributedVirtualSwitch").(*VmwareDistributedVirtualSwitch)
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
	vswitch := m.Map().Any("VmwareDistributedVirtualSwitch").(*VmwareDistributedVirtualSwitch)
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

// TestDVSConcreteType guards against vcsim regressing to reporting its DVS as
// the abstract DistributedVirtualSwitch type. A real vCenter always reports
// the concrete VmwareDistributedVirtualSwitch -- vRNI's collector (and any
// client that keys off the managed object type, e.g. to pick the SDM/config
// type to publish) derives its behavior from that MOR type, not from
// config content, so an abstractly-typed switch is invisible to it even
// though its portgroups are discovered fine.
func TestDVSConcreteType(t *testing.T) {
	Test(func(ctx context.Context, c *vim25.Client) {
		ref := Map(ctx).Any("VmwareDistributedVirtualSwitch").Reference()
		if ref.Type != "VmwareDistributedVirtualSwitch" {
			t.Fatalf("MOR type = %q, want VmwareDistributedVirtualSwitch", ref.Type)
		}

		// A client may still query generically by the abstract type (as real
		// vCenter clients have always been able to do, since VmwareDVS is a
		// vmodl subtype of it) -- the property collector's type-hierarchy
		// matching (simulator/property_collector.go's use of reflect on the
		// anonymously-embedded mo field) must still resolve it against the
		// concrete object after this rename.
		pc := property.DefaultCollector(c)
		for _, queryType := range []string{"VmwareDistributedVirtualSwitch", "DistributedVirtualSwitch"} {
			res, err := pc.RetrieveProperties(ctx, types.RetrieveProperties{
				SpecSet: []types.PropertyFilterSpec{{
					ObjectSet: []types.ObjectSpec{{Obj: ref}},
					PropSet:   []types.PropertySpec{{Type: queryType, PathSet: []string{"name"}}},
				}},
			})
			if err != nil {
				t.Fatalf("retrieve via type %s: %s", queryType, err)
			}
			if len(res.Returnval) != 1 {
				t.Errorf("query by type %q: got %d objects, want 1 (hierarchy match against the concrete type failed)",
					queryType, len(res.Returnval))
			}
		}
	})
}

// TestDVSDefaultProductInfo guards against a DVS ending up with an empty
// Config.ProductInfo.Vendor. DVSCreateSpec.ProductInfo's own doc comment
// states a real vCenter defaults it when the client doesn't supply one
// ("the Server will use the latest version") -- at least one real collector
// (vRNI) reads config.productInfo.vendor and crashes on a null value.
// DVSConfigInfo.ProductInfo is a required, always-serialized (non-pointer)
// field whose own sub-fields are all `omitempty`, so leaving it zero-valued
// serializes as an empty <productInfo/> with no vendor element at all,
// which deserializes as a null Vendor on the Java side. Neither
// `govc dvs.create` (no -product-version) nor this simulator's own default
// model-driven DVS creation ever supplied one, so this was always empty
// before the fix. Summary.ProductInfo is checked too since it's the same
// concept and real vCenter keeps both consistent.
func TestDVSDefaultProductInfo(t *testing.T) {
	Test(func(ctx context.Context, c *vim25.Client) {
		finder := find.NewFinder(c, false)
		dc, err := finder.DatacenterList(ctx, "*")
		if err != nil {
			t.Fatal(err)
		}
		finder.SetDatacenter(dc[0])
		folders, err := dc[0].Folders(ctx)
		if err != nil {
			t.Fatal(err)
		}

		// No ProductInfo supplied -- must default rather than stay nil.
		var spec types.DVSCreateSpec
		spec.ConfigSpec = &types.VMwareDVSConfigSpec{}
		spec.ConfigSpec.GetDVSConfigSpec().Name = "DVS-default-product-info"

		task, err := folders.NetworkFolder.CreateDVS(ctx, spec)
		if err != nil {
			t.Fatal(err)
		}
		info, err := task.WaitForResult(ctx, nil)
		if err != nil {
			t.Fatal(err)
		}
		dvs := object.NewDistributedVirtualSwitch(c, info.Result.(types.ManagedObjectReference))

		var moDVS mo.DistributedVirtualSwitch
		if err := dvs.Properties(ctx, dvs.Reference(), []string{"summary", "config"}, &moDVS); err != nil {
			t.Fatal(err)
		}
		if moDVS.Summary.ProductInfo == nil {
			t.Fatal("summary.productInfo is nil, want a default value")
		}
		if moDVS.Summary.ProductInfo.Vendor == "" {
			t.Error("summary.productInfo.vendor is empty, want a non-empty default")
		}
		if moDVS.Config.GetDVSConfigInfo().ProductInfo.Vendor == "" {
			t.Error("config.productInfo.vendor is empty, want a non-empty default")
		}

		// An explicitly-supplied ProductInfo must still be honored, not
		// overridden by the default.
		var spec2 types.DVSCreateSpec
		spec2.ConfigSpec = &types.VMwareDVSConfigSpec{}
		spec2.ConfigSpec.GetDVSConfigSpec().Name = "DVS-explicit-product-info"
		spec2.ProductInfo = &types.DistributedVirtualSwitchProductSpec{Vendor: "Acme Corp", Version: "1.2.3"}

		task2, err := folders.NetworkFolder.CreateDVS(ctx, spec2)
		if err != nil {
			t.Fatal(err)
		}
		info2, err := task2.WaitForResult(ctx, nil)
		if err != nil {
			t.Fatal(err)
		}
		dvs2 := object.NewDistributedVirtualSwitch(c, info2.Result.(types.ManagedObjectReference))

		var moDVS2 mo.DistributedVirtualSwitch
		if err := dvs2.Properties(ctx, dvs2.Reference(), []string{"summary", "config"}, &moDVS2); err != nil {
			t.Fatal(err)
		}
		if moDVS2.Summary.ProductInfo == nil || moDVS2.Summary.ProductInfo.Vendor != "Acme Corp" {
			t.Errorf("explicit ProductInfo was not preserved: %+v", moDVS2.Summary.ProductInfo)
		}
		if moDVS2.Config.GetDVSConfigInfo().ProductInfo.Vendor != "Acme Corp" {
			t.Errorf("explicit ProductInfo was not preserved in config: %+v", moDVS2.Config.GetDVSConfigInfo().ProductInfo)
		}

		// `govc dvs.create` (without -product-version) sends a non-nil
		// ProductInfo with an empty Vendor, not a nil ProductInfo -- this
		// must also default rather than leave Vendor empty.
		var spec3 types.DVSCreateSpec
		spec3.ConfigSpec = &types.VMwareDVSConfigSpec{}
		spec3.ConfigSpec.GetDVSConfigSpec().Name = "DVS-empty-product-info"
		spec3.ProductInfo = new(types.DistributedVirtualSwitchProductSpec)

		task3, err := folders.NetworkFolder.CreateDVS(ctx, spec3)
		if err != nil {
			t.Fatal(err)
		}
		info3, err := task3.WaitForResult(ctx, nil)
		if err != nil {
			t.Fatal(err)
		}
		dvs3 := object.NewDistributedVirtualSwitch(c, info3.Result.(types.ManagedObjectReference))

		var moDVS3 mo.DistributedVirtualSwitch
		if err := dvs3.Properties(ctx, dvs3.Reference(), []string{"summary", "config"}, &moDVS3); err != nil {
			t.Fatal(err)
		}
		if moDVS3.Summary.ProductInfo == nil {
			t.Fatal("summary.productInfo is nil, want a default value")
		}
		if moDVS3.Summary.ProductInfo.Vendor == "" {
			t.Error("summary.productInfo.vendor is empty, want a non-empty default")
		}
		if moDVS3.Config.GetDVSConfigInfo().ProductInfo.Vendor == "" {
			t.Error("config.productInfo.vendor is empty, want a non-empty default")
		}
	})
}
