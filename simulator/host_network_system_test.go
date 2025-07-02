// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"context"
	"testing"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/simulator/esx"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

func TestHostNetworkSystem(t *testing.T) {
	ctx := context.Background()

	s := New(NewServiceInstance(NewContext(), esx.ServiceContent, esx.RootFolder))

	c := s.client()
	host := object.NewHostSystem(c, esx.HostSystem.Reference())

	ns, err := host.ConfigManager().NetworkSystem(ctx)
	if err != nil {
		t.Fatal(err)
	}

	finder := find.NewFinder(c, false)
	finder.SetDatacenter(object.NewDatacenter(c, esx.Datacenter.Reference()))

	// created by default
	_, err = finder.Network(ctx, "VM Network")
	if err != nil {
		t.Fatal(err)
	}
	var mns mo.HostNetworkSystem
	err = ns.Properties(ctx, ns.Reference(), []string{"networkInfo.portgroup"}, &mns)
	if err != nil {
		t.Fatal(err)
	}
	if len(mns.NetworkInfo.Portgroup) != 2 {
		t.Fatal("expected networkInfo.portgroup to have length of 2")
	}
	if mns.NetworkInfo.Portgroup[0].Key != "key-vim.host.PortGroup-VM Network" {
		t.Fatal("expected networkInfo.portgroup[0] to be VM Network")
	}
	if mns.NetworkInfo.Portgroup[1].Key != "key-vim.host.PortGroup-Management Network" {
		t.Fatal("expected networkInfo.portgroup[1] to be Management Network")
	}

	// not created yet
	_, err = finder.Network(ctx, "bridge")
	if err == nil {
		t.Fatal("expected error")
	}

	err = ns.AddVirtualSwitch(ctx, "vSwitch0", nil)
	if err == nil {
		t.Fatal("expected error") // DuplicateName
	}

	err = ns.AddVirtualSwitch(ctx, "vSwitch1", nil)
	if err != nil {
		t.Fatal(err)
	}

	spec := types.HostPortGroupSpec{}
	err = ns.AddPortGroup(ctx, spec)
	if err == nil {
		t.Fatal("expected error") // InvalidArgument "name"
	}

	spec.Name = "bridge"
	err = ns.AddPortGroup(ctx, spec)
	if err == nil {
		t.Fatal("expected error") // NotFound
	}

	spec.VswitchName = "vSwitch1"
	err = ns.AddPortGroup(ctx, spec)
	if err != nil {
		t.Fatal(err)
	}

	_, err = finder.Network(ctx, "bridge")
	if err != nil {
		t.Fatal(err)
	}

	err = ns.Properties(ctx, ns.Reference(), []string{"networkInfo.portgroup"}, &mns)
	if err != nil {
		t.Fatal(err)
	}
	if len(mns.NetworkInfo.Portgroup) != 3 {
		t.Fatal("expected networkInfo.portgroup to have length of 3")
	}
	if mns.NetworkInfo.Portgroup[2].Spec != spec {
		t.Fatal("expected last networkInfo.portgroup to have an equal spec")
	}
	if mns.NetworkInfo.Portgroup[0].Key != "key-vim.host.PortGroup-VM Network" {
		t.Fatal("expected networkInfo.portgroup[0] to be VM Network")
	}
	if mns.NetworkInfo.Portgroup[1].Key != "key-vim.host.PortGroup-Management Network" {
		t.Fatal("expected networkInfo.portgroup[1] to be Management Network")
	}
	if mns.NetworkInfo.Portgroup[2].Key != "key-vim.host.PortGroup-bridge" {
		t.Fatal("expected networkInfo.portgroup[2] to be bridge")
	}

	err = ns.AddPortGroup(ctx, spec)
	if err == nil {
		t.Error("expected error") // DuplicateName
	}

	err = ns.RemovePortGroup(ctx, "bridge")
	if err != nil {
		t.Fatal(err)
	}

	_, err = finder.Network(ctx, "bridge")
	if err == nil {
		t.Fatal("expected error")
	}

	err = ns.Properties(ctx, ns.Reference(), []string{"networkInfo.portgroup"}, &mns)
	if err != nil {
		t.Fatal(err)
	}
	if len(mns.NetworkInfo.Portgroup) != 2 {
		t.Fatal("expected networkInfo.portgroup to have length of 2")
	}
	if mns.NetworkInfo.Portgroup[0].Key != "key-vim.host.PortGroup-VM Network" {
		t.Fatal("expected networkInfo.portgroup[0] to be VM Network")
	}
	if mns.NetworkInfo.Portgroup[1].Key != "key-vim.host.PortGroup-Management Network" {
		t.Fatal("expected networkInfo.portgroup[1] to be Management Network")
	}

	err = ns.RemovePortGroup(ctx, "bridge")
	if err == nil {
		t.Error("expected error")
	}

	err = ns.RemoveVirtualSwitch(ctx, "vSwitch1")
	if err != nil {
		t.Fatal(err)
	}

	err = ns.RemoveVirtualSwitch(ctx, "vSwitch1")
	if err == nil {
		t.Error("expected error")
	}

	info, err := ns.QueryNetworkHint(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}

	if len(info) != 0 { // TODO: data is only returned when Model.Load is used
		t.Errorf("len=%d", len(info))
	}
}
