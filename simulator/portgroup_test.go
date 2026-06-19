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

func TestPortgroupSubnetId(t *testing.T) {
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
			Name:     "subnet-pg",
			NumPorts: 10,
			SubnetId: "subnet-12345",
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

	// Find the network folder from the datacenter to search for the portgroup
	dc := ctx.Map.Any("Datacenter").(*Datacenter)
	netFolderRef := dc.NetworkFolder
	netFolder := ctx.Map.Get(netFolderRef).(*Folder)
	pgRef := ctx.Map.FindByName("subnet-pg", netFolder.ChildEntity)
	if pgRef == nil {
		t.Fatal("subnet-pg portgroup not found")
	}

	pge := ctx.Map.Get(pgRef.Reference()).(*DistributedVirtualPortgroup)
	if pge.Config.SubnetId != "subnet-12345" {
		t.Fatalf("expected pg SubnetId==subnet-12345; got %q", pge.Config.SubnetId)
	}

	// Reconfigure the SubnetId
	pg := object.NewDistributedVirtualPortgroup(c, pgRef.Reference())
	pgspec := types.DVPortgroupConfigSpec{
		Name:     "subnet-pg",
		SubnetId: "subnet-67890",
	}

	task, err = pg.Reconfigure(ctx, pgspec)
	if err != nil {
		t.Fatal(err)
	}

	err = task.Wait(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if pge.Config.SubnetId != "subnet-67890" {
		t.Fatalf("expected reconfigured pg SubnetId==subnet-67890; got %q", pge.Config.SubnetId)
	}

	// Now query config target via environment browser to check SubnetInfo
	vm := ctx.Map.Any("VirtualMachine").(*VirtualMachine)
	browser := object.NewEnvironmentBrowser(c, vm.EnvironmentBrowser)

	target, err := browser.QueryConfigTarget(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}

	foundSubnet := false
	for _, subnet := range target.SubnetInfo {
		if subnet.Id == "subnet-67890" {
			foundSubnet = true
			break
		}
	}

	if !foundSubnet {
		t.Error("expected subnet-67890 in QueryConfigTarget SubnetInfo")
	}

	foundPGSubnet := false
	for _, pgInfo := range target.DistributedVirtualPortgroup {
		if pgInfo.PortgroupKey == pge.Key && pgInfo.SubnetId == "subnet-67890" {
			foundPGSubnet = true
			break
		}
	}

	if !foundPGSubnet {
		t.Error("expected portgroup to have SubnetId==subnet-67890 in QueryConfigTarget")
	}

	// Now test virtual machine ethernet card inheriting SubnetId
	// Add an ethernet card backing to the newly created portgroup
	backing, err := pg.EthernetCardBackingInfo(ctx)
	if err != nil {
		t.Fatal(err)
	}

	clientVM := object.NewVirtualMachine(c, vm.Reference())

	devices, err := clientVM.Device(ctx)
	if err != nil {
		t.Fatal(err)
	}

	card, err := devices.CreateEthernetCard("vmxnet3", backing)
	if err != nil {
		t.Fatal(err)
	}

	err = clientVM.AddDevice(ctx, card)
	if err != nil {
		t.Fatal(err)
	}

	// Retrieve devices again and check if SubnetId was inherited from the portgroup
	devices, err = clientVM.Device(ctx)
	if err != nil {
		t.Fatal(err)
	}

	cards := devices.SelectByType((*types.VirtualEthernetCard)(nil))
	foundInheritedSubnet := false
	for _, dev := range cards {
		nic := dev.(types.BaseVirtualEthernetCard).GetVirtualEthernetCard()
		if nic.SubnetId == "subnet-67890" {
			foundInheritedSubnet = true
			break
		}
	}

	if !foundInheritedSubnet {
		t.Error("expected newly added VirtualEthernetCard to inherit SubnetId from its portgroup backing")
	}

	// Now test creating an ethernet card specifying ONLY SubnetId and NO backing
	devices, err = clientVM.Device(ctx)
	if err != nil {
		t.Fatal(err)
	}

	// Create a raw VirtualVmxnet3 card without specifying a backing, but setting SubnetId
	rawCard := &types.VirtualVmxnet3{}
	rawCard.SubnetId = "subnet-67890"

	err = clientVM.AddDevice(ctx, rawCard)
	if err != nil {
		t.Fatalf("failed to add device with only SubnetId: %s", err)
	}

	// Retrieve devices again to check if vCenter resolved and populated the backing
	devices, err = clientVM.Device(ctx)
	if err != nil {
		t.Fatal(err)
	}

	cards = devices.SelectByType((*types.VirtualEthernetCard)(nil))
	foundAutoResolvedBacking := false
	for _, dev := range cards {
		nic := dev.(types.BaseVirtualEthernetCard).GetVirtualEthernetCard()
		if nic.SubnetId == "subnet-67890" && nic.Backing != nil {
			// Check if the backing points to our portgroup
			if b, ok := nic.Backing.(*types.VirtualEthernetCardDistributedVirtualPortBackingInfo); ok {
				if b.Port.PortgroupKey == pge.Key {
					foundAutoResolvedBacking = true
					break
				}
			}
		}
	}

	if !foundAutoResolvedBacking {
		t.Error("expected VirtualEthernetCard with only SubnetId to auto-resolve its backing to the matching Portgroup")
	}
}
