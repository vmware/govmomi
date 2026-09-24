// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"context"
	"reflect"
	"regexp"
	"strconv"
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

	// dvs (DVS1) had all its hosts removed by the last test case above ("Remove
	// dvs1 == OK"), so its uplink portgroup now correctly contributes 0 ports
	// (uplink ports are host-scoped -- see uplinkPorts()); only the DVPG0
	// portgroup added earlier still has its one port.
	ports, err := dvs.FetchDVPorts(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(ports) != 1 {
		t.Fatalf("expected 1 port in DVPorts; got %d", len(ports))
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

	uplinkPorts := make([]types.DistributedVirtualPort, len(vswitch.Summary.HostMember))
	for i := range uplinkPorts {
		uplinkPorts[i] = types.DistributedVirtualPort{PortgroupKey: pgs[0].Value, Key: "0"}
	}

	regularPg := m.Map().Get(pgs[1]).(*DistributedVirtualPortgroup)
	regularPorts := make([]types.DistributedVirtualPort, len(regularPg.Vm))
	for i := range regularPorts {
		regularPorts[i] = types.DistributedVirtualPort{PortgroupKey: pgs[1].Value, Key: strconv.Itoa(i)}
	}

	tests := []struct {
		name     string
		criteria *types.DistributedVirtualSwitchPortCriteria
		expected []types.DistributedVirtualPort
	}{
		{
			"empty criteria",
			&types.DistributedVirtualSwitchPortCriteria{},
			append(append([]types.DistributedVirtualPort{}, uplinkPorts...), regularPorts...),
		},
		{
			"inside PortgroupKeys",
			&types.DistributedVirtualSwitchPortCriteria{
				PortgroupKey: []string{pgs[0].Value},
				Inside:       types.NewBool(true),
			},
			uplinkPorts,
		},
		{
			"outside PortgroupKeys",
			&types.DistributedVirtualSwitchPortCriteria{
				PortgroupKey: []string{pgs[0].Value},
				Inside:       types.NewBool(false),
			},
			regularPorts,
		},
		{
			"PortKeys",
			&types.DistributedVirtualSwitchPortCriteria{
				PortKey: []string{"1"},
			},
			regularPorts[1:2],
		},
		{
			// both the uplink ports (Connectee set to each host's pnic) and
			// the regular portgroup's port (Connectee set to the default
			// VM's vNIC) are connected.
			"connected",
			&types.DistributedVirtualSwitchPortCriteria{
				Connected: types.NewBool(true),
			},
			append(append([]types.DistributedVirtualPort{}, uplinkPorts...), regularPorts...),
		},
		{
			"not connected",
			&types.DistributedVirtualSwitchPortCriteria{
				Connected: types.NewBool(false),
			},
			[]types.DistributedVirtualPort{},
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

func TestUplinkPortsPerHostScaling(t *testing.T) {
	m := VPX()

	defer m.Remove()

	if err := m.Create(); err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	c := m.Service.client()
	simCtx := m.Service.Context

	vswitch := m.Map().Any("VmwareDistributedVirtualSwitch").(*VmwareDistributedVirtualSwitch)
	dvs := object.NewDistributedVirtualSwitch(c, vswitch.Reference())

	var hostRef types.ManagedObjectReference
	for _, h := range vswitch.Summary.HostMember {
		if len(simCtx.Map.Get(h).(*HostSystem).Vm) == 0 {
			hostRef = h
			break
		}
	}
	if hostRef.Value == "" {
		t.Fatal("expected at least one host with no VMs")
	}

	task, err := dvs.Reconfigure(ctx, &types.DVSConfigSpec{
		Host: []types.DistributedVirtualSwitchHostMemberConfigSpec{{
			Operation: string(types.ConfigSpecOperationRemove),
			Host:      hostRef,
		}},
	})
	if err != nil {
		t.Fatal(err)
	}
	if err = task.Wait(ctx); err != nil {
		t.Fatal(err)
	}

	task, err = dvs.Reconfigure(ctx, &types.DVSConfigSpec{
		Host: []types.DistributedVirtualSwitchHostMemberConfigSpec{{
			Operation: string(types.ConfigSpecOperationAdd),
			Host:      hostRef,
			Backing: &types.DistributedVirtualSwitchHostMemberPnicBacking{
				PnicSpec: []types.DistributedVirtualSwitchHostMemberPnicSpec{
					{PnicDevice: "vmnic0"},
					{PnicDevice: "vmnic1"},
				},
			},
		}},
	})
	if err != nil {
		t.Fatal(err)
	}
	if err = task.Wait(ctx); err != nil {
		t.Fatal(err)
	}

	uplinkPg := vswitch.uplinkPortgroup(simCtx)
	if uplinkPg == nil {
		t.Fatal("expected uplink portgroup")
	}

	allPorts, err := dvs.FetchDVPorts(ctx, &types.DistributedVirtualSwitchPortCriteria{
		PortgroupKey: []string{uplinkPg.Key},
		Inside:       types.NewBool(true),
	})
	if err != nil {
		t.Fatal(err)
	}

	var ports []types.DistributedVirtualPort
	for _, p := range allPorts {
		if p.Connectee != nil && p.Connectee.ConnectedEntity != nil && *p.Connectee.ConnectedEntity == hostRef {
			ports = append(ports, p)
		}
	}

	if len(ports) != 2 {
		t.Fatalf("expected 2 uplink ports for host with 2 pnics; got %d: %v", len(ports), ports)
	}

	seen := make(map[string]bool)
	for _, p := range ports {
		if p.Connectee.Type != string(types.DistributedVirtualSwitchPortConnecteeConnecteeTypePnic) {
			t.Errorf("port %s: expected Connectee.Type pnic; got %q", p.Key, p.Connectee.Type)
		}
		seen[p.Key] = true
	}

	if !seen["0"] || !seen["1"] {
		t.Errorf("expected port keys \"0\" and \"1\"; got %v", ports)
	}
}

func TestRegularPortsExceedPreallocatedKeys(t *testing.T) {
	m := VPX()

	defer m.Remove()

	if err := m.Create(); err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	c := m.Service.client()
	simCtx := m.Service.Context

	vswitch := m.Map().Any("VmwareDistributedVirtualSwitch").(*VmwareDistributedVirtualSwitch)
	dvs := object.NewDistributedVirtualSwitch(c, vswitch.Reference())

	task, err := dvs.AddPortgroup(ctx, []types.DVPortgroupConfigSpec{{Name: "scale-pg", NumPorts: 1}})
	if err != nil {
		t.Fatal(err)
	}
	if err = task.Wait(ctx); err != nil {
		t.Fatal(err)
	}

	dvsMo := simCtx.Map.Get(vswitch.Reference()).(*VmwareDistributedVirtualSwitch)
	pg := simCtx.Map.FindByName("scale-pg", dvsMo.Portgroup).(*DistributedVirtualPortgroup)

	backing := &types.VirtualEthernetCardDistributedVirtualPortBackingInfo{
		Port: types.DistributedVirtualSwitchPortConnection{
			PortgroupKey: pg.Key,
			SwitchUuid:   dvsMo.Uuid,
		},
	}

	allVMs := simCtx.Map.All("VirtualMachine")
	if len(allVMs) < 3 {
		t.Fatalf("expected at least 3 VMs in the default model; got %d", len(allVMs))
	}

	vmRefs := make([]types.ManagedObjectReference, 0, 3)
	for _, v := range allVMs[:3] {
		vmRef := v.Reference()
		clientVM := object.NewVirtualMachine(c, vmRef)

		devices, err := clientVM.Device(ctx)
		if err != nil {
			t.Fatal(err)
		}
		cards := devices.SelectByType((*types.VirtualEthernetCard)(nil))
		if len(cards) == 0 {
			t.Fatalf("expected at least one ethernet card on VM %s", vmRef)
		}
		card := cards[0]
		card.(types.BaseVirtualEthernetCard).GetVirtualEthernetCard().Backing = backing

		if err := clientVM.EditDevice(ctx, card); err != nil {
			t.Fatal(err)
		}
		vmRefs = append(vmRefs, vmRef)
	}

	ports, err := dvs.FetchDVPorts(ctx, &types.DistributedVirtualSwitchPortCriteria{
		PortgroupKey: []string{pg.Key},
		Inside:       types.NewBool(true),
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(ports) != len(vmRefs) {
		t.Fatalf("expected %d ports (one per connected VM); got %d: %v", len(vmRefs), len(ports), ports)
	}

	connected := make(map[types.ManagedObjectReference]bool)
	keys := make(map[string]bool)
	for _, p := range ports {
		if p.Connectee == nil || p.Connectee.Type != string(types.DistributedVirtualSwitchPortConnecteeConnecteeTypeVmVnic) {
			t.Errorf("port %s: expected Connectee.Type vmVnic; got %+v", p.Key, p.Connectee)
			continue
		}
		if p.Connectee.ConnectedEntity == nil {
			t.Errorf("port %s: expected a non-nil ConnectedEntity", p.Key)
			continue
		}
		connected[*p.Connectee.ConnectedEntity] = true
		keys[p.Key] = true
	}

	for _, vmRef := range vmRefs {
		if !connected[vmRef] {
			t.Errorf("expected a port connected to VM %s; ports=%v", vmRef, ports)
		}
	}

	if len(keys) != len(vmRefs) {
		t.Errorf("expected %d unique port keys; got %v", len(vmRefs), keys)
	}

	if int(pg.Config.NumPorts) != len(vmRefs) {
		t.Errorf("expected config.numPorts to grow to %d; got %d", len(vmRefs), pg.Config.NumPorts)
	}
	if len(pg.PortKeys) != len(vmRefs) {
		t.Errorf("expected %d PortKeys after auto-expand; got %v", len(vmRefs), pg.PortKeys)
	}
}

// TestDVSUuidFormat verifies newDVSUuid() produces the same wire shape a
// real vCenter/ESXi uses for a DVS UUID -- 16 hex byte pairs, space
// separated, with a dash between the 8th and 9th pair -- rather than a plain
// dashed UUID, and that a live DVS's Summary.Uuid actually uses it.
func TestDVSUuidFormat(t *testing.T) {
	re := regexp.MustCompile(`^[0-9a-f]{2}( [0-9a-f]{2}){7}-([0-9a-f]{2} ){7}[0-9a-f]{2}$`)

	uuid := newDVSUuid("some-dvs-name")
	if !re.MatchString(uuid) {
		t.Fatalf("newDVSUuid() = %q, want the vCenter wire shape (e.g. %q)",
			uuid, "50 13 a2 63 0a a6 77 65-37 e2 20 e6 2b 8f a2 f6")
	}

	// Stable per input name, matching newUUID()'s own contract.
	if again := newDVSUuid("some-dvs-name"); again != uuid {
		t.Fatalf("newDVSUuid() = %q then %q, want stable output for the same input", uuid, again)
	}

	Test(func(ctx context.Context, c *vim25.Client) {
		vswitch := Map(ctx).Any("VmwareDistributedVirtualSwitch").(*VmwareDistributedVirtualSwitch)
		if !re.MatchString(vswitch.Uuid) {
			t.Fatalf("DVS Summary.Uuid = %q, want the vCenter wire shape", vswitch.Uuid)
		}
	})
}

// TestDVSHostProxySwitch verifies that HostSystem.Config.Network.ProxySwitch
// gains a HostProxySwitch entry for a DVS when the host joins it, and loses
// that entry when the host leaves -- matching real vCenter, which always
// keeps this host-side membership record in sync with the DVS's own
// Summary.HostMember.
func TestDVSHostProxySwitch(t *testing.T) {
	m := VPX()

	defer m.Remove()

	if err := m.Create(); err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	c := m.Service.client()
	simCtx := m.Service.Context

	finder := find.NewFinder(c, false)
	dc, err := finder.DatacenterList(ctx, "*")
	if err != nil {
		t.Fatal(err)
	}
	finder.SetDatacenter(dc[0])

	hosts, err := finder.HostSystemList(ctx, "*/*")
	if err != nil {
		t.Fatal(err)
	}

	vswitch := m.Map().Any("VmwareDistributedVirtualSwitch").(*VmwareDistributedVirtualSwitch)
	dvs := object.NewDistributedVirtualSwitch(c, vswitch.Reference())

	// Every host created by the model already joined this DVS -- pick one
	// with no VMs on it (removing a host with a VM connected to one of the
	// DVS's portgroups is rejected with ResourceInUse) and remove it first,
	// so this test controls the join/leave transition rather than only
	// observing the model's own initial state.
	var hostRef types.ManagedObjectReference
	for _, h := range hosts {
		ref := h.Reference()
		if len(simCtx.Map.Get(ref).(*HostSystem).Vm) == 0 {
			hostRef = ref
			break
		}
	}
	if hostRef.Value == "" {
		t.Fatal("expected at least one host with no VMs")
	}

	hasProxySwitch := func() bool {
		h := simCtx.Map.Get(hostRef).(*HostSystem)
		for _, ps := range h.Config.Network.ProxySwitch {
			if ps.DvsUuid == vswitch.Uuid {
				return true
			}
		}
		return false
	}

	if !hasProxySwitch() {
		t.Fatal("expected host to already have a HostProxySwitch entry for this DVS from model creation")
	}

	config := &types.DVSConfigSpec{
		Host: []types.DistributedVirtualSwitchHostMemberConfigSpec{{
			Operation: string(types.ConfigSpecOperationRemove),
			Host:      hostRef,
		}},
	}
	task, err := dvs.Reconfigure(ctx, config)
	if err != nil {
		t.Fatal(err)
	}
	if err = task.Wait(ctx); err != nil {
		t.Fatal(err)
	}

	if hasProxySwitch() {
		t.Fatal("expected HostProxySwitch entry to be removed after leaving the DVS")
	}

	config.Host[0].Operation = string(types.ConfigSpecOperationAdd)
	task, err = dvs.Reconfigure(ctx, config)
	if err != nil {
		t.Fatal(err)
	}
	if err = task.Wait(ctx); err != nil {
		t.Fatal(err)
	}

	if !hasProxySwitch() {
		t.Fatal("expected HostProxySwitch entry to be recreated after rejoining the DVS")
	}
}

// TestDVSConfigHost verifies that DVSConfigInfo.Host (config.host) and
// Summary.NumHosts both stay in sync with Summary.HostMember when a host
// joins or leaves a DVS -- matching real vCenter, which always keeps all
// three consistent.
func TestDVSConfigHost(t *testing.T) {
	m := VPX()

	defer m.Remove()

	if err := m.Create(); err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	c := m.Service.client()
	simCtx := m.Service.Context

	finder := find.NewFinder(c, false)
	dc, err := finder.DatacenterList(ctx, "*")
	if err != nil {
		t.Fatal(err)
	}
	finder.SetDatacenter(dc[0])

	hosts, err := finder.HostSystemList(ctx, "*/*")
	if err != nil {
		t.Fatal(err)
	}

	vswitch := m.Map().Any("VmwareDistributedVirtualSwitch").(*VmwareDistributedVirtualSwitch)
	dvs := object.NewDistributedVirtualSwitch(c, vswitch.Reference())

	// Same host-selection reasoning as TestDVSHostProxySwitch: pick a host
	// with no VMs, since removing one that has a VM connected to one of
	// this DVS's portgroups is rejected with ResourceInUse.
	var hostRef types.ManagedObjectReference
	for _, h := range hosts {
		ref := h.Reference()
		if len(simCtx.Map.Get(ref).(*HostSystem).Vm) == 0 {
			hostRef = ref
			break
		}
	}
	if hostRef.Value == "" {
		t.Fatal("expected at least one host with no VMs")
	}

	assertConsistent := func() {
		s := simCtx.Map.Get(vswitch.Self).(*VmwareDistributedVirtualSwitch)
		configHost := s.Config.GetDVSConfigInfo().Host

		if int(s.Summary.NumHosts) != len(s.Summary.HostMember) {
			t.Fatalf("Summary.NumHosts=%d, want %d (len(Summary.HostMember))",
				s.Summary.NumHosts, len(s.Summary.HostMember))
		}
		if len(configHost) != len(s.Summary.HostMember) {
			t.Fatalf("len(Config.Host)=%d, want %d (len(Summary.HostMember))",
				len(configHost), len(s.Summary.HostMember))
		}
		for _, ref := range s.Summary.HostMember {
			found := false
			for _, ch := range configHost {
				if ch.Config.Host != nil && *ch.Config.Host == ref {
					found = true
					break
				}
			}
			if !found {
				t.Fatalf("Config.Host=%v missing host %s present in Summary.HostMember", configHost, ref)
			}
		}
	}

	assertConsistent()

	config := &types.DVSConfigSpec{
		Host: []types.DistributedVirtualSwitchHostMemberConfigSpec{{
			Operation: string(types.ConfigSpecOperationRemove),
			Host:      hostRef,
		}},
	}
	task, err := dvs.Reconfigure(ctx, config)
	if err != nil {
		t.Fatal(err)
	}
	if err = task.Wait(ctx); err != nil {
		t.Fatal(err)
	}

	assertConsistent()

	config.Host[0].Operation = string(types.ConfigSpecOperationAdd)
	task, err = dvs.Reconfigure(ctx, config)
	if err != nil {
		t.Fatal(err)
	}
	if err = task.Wait(ctx); err != nil {
		t.Fatal(err)
	}

	assertConsistent()
}

// TestDVSConcreteType guards against vcsim regressing to reporting its DVS as
// the abstract DistributedVirtualSwitch type. A real vCenter always reports
// the concrete VmwareDistributedVirtualSwitch -- any client that keys off
// the managed object type (e.g. to pick the SDM/config type to publish)
// derives its behavior from that MOR type, not from config content, so an
// abstractly-typed switch is invisible to it even though its portgroups are
// discovered fine.
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
// ("the Server will use the latest version") -- a null Vendor is not a
// shape a real vCenter's DVS config ever has, and a collector reading
// config.productInfo.vendor unconditionally can crash on it.
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
