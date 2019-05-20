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
	"testing"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/simulator/esx"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

func TestHostNetworkSystem(t *testing.T) {
	ctx := context.Background()

	s := New(NewServiceInstance(esx.ServiceContent, esx.RootFolder))

	host := object.NewHostSystem(s.client, esx.HostSystem.Reference())

	ns, err := host.ConfigManager().NetworkSystem(ctx)
	if err != nil {
		t.Fatal(err)
	}

	finder := find.NewFinder(s.client, false)
	finder.SetDatacenter(object.NewDatacenter(s.client, esx.Datacenter.Reference()))

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
}
