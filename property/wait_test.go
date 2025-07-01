// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package property_test

import (
	"context"
	"log"
	"testing"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

// Test that task.Wait() propagates MissingSet errors
func TestWaitPermissionFault(t *testing.T) {
	ctx := context.Background()

	model := simulator.ESX()

	defer model.Remove()
	err := model.Create()
	if err != nil {
		log.Fatal(err)
	}

	s := model.Service.NewServer()
	defer s.Close()

	c, _ := govmomi.NewClient(ctx, s.URL, true)

	pc := new(propCollForWaitForPermsTest)
	pc.Self = model.ServiceContent.PropertyCollector
	model.Map().Put(pc)

	dm := object.NewVirtualDiskManager(c.Client)

	spec := &types.FileBackedVirtualDiskSpec{
		VirtualDiskSpec: types.VirtualDiskSpec{
			AdapterType: string(types.VirtualDiskAdapterTypeLsiLogic),
			DiskType:    string(types.VirtualDiskTypeThin),
		},
		CapacityKb: 1024 * 1024,
	}

	name := "[LocalDS_0] disk1.vmdk"

	task, err := dm.CreateVirtualDisk(ctx, name, nil, spec)
	if err != nil {
		t.Fatal(err)
	}

	err = task.Wait(ctx)
	if err == nil {
		t.Fatal("expected error")
	}

	if !soap.IsVimFault(err) {
		t.Fatal("expected vim fault")
	}

	fault, ok := soap.ToVimFault(err).(*types.NoPermission)
	if !ok {
		t.Fatalf("unexpected vim fault: %T", fault)
	}
}

// Test that task.WaitEx() propagates MissingSet errors
func TestWaitExPermissionFault(t *testing.T) {
	ctx := context.Background()

	model := simulator.ESX()

	defer model.Remove()
	err := model.Create()
	if err != nil {
		log.Fatal(err)
	}

	s := model.Service.NewServer()
	defer s.Close()

	c, _ := govmomi.NewClient(ctx, s.URL, true)

	pc := new(propCollForWaitForPermsTest)
	pc.Self = model.ServiceContent.PropertyCollector
	model.Map().Put(pc)

	dm := object.NewVirtualDiskManager(c.Client)

	spec := &types.FileBackedVirtualDiskSpec{
		VirtualDiskSpec: types.VirtualDiskSpec{
			AdapterType: string(types.VirtualDiskAdapterTypeLsiLogic),
			DiskType:    string(types.VirtualDiskTypeThin),
		},
		CapacityKb: 1024 * 1024,
	}

	name := "[LocalDS_0] disk1.vmdk"

	task, err := dm.CreateVirtualDisk(ctx, name, nil, spec)
	if err != nil {
		t.Fatal(err)
	}

	err = task.WaitEx(ctx)
	if err == nil {
		t.Fatal("expected error")
	}

	if !soap.IsVimFault(err) {
		t.Fatal("expected vim fault")
	}

	fault, ok := soap.ToVimFault(err).(*types.NoPermission)
	if !ok {
		t.Fatalf("unexpected vim fault: %T", fault)
	}
}

type propCollForWaitForPermsTest struct {
	simulator.PropertyCollector
}

// CreatePropertyCollector overrides the vcsim impl to return this test's PC impl
func (pc *propCollForWaitForPermsTest) CreatePropertyCollector(
	ctx *simulator.Context,
	c *types.CreatePropertyCollector) soap.HasFault {

	return &methods.CreatePropertyCollectorBody{
		Res: &types.CreatePropertyCollectorResponse{
			Returnval: ctx.Session.Put(new(propCollForWaitForPermsTest)).Reference(),
		},
	}
}

// WaitForUpdatesEx overrides the vcsim impl to inject a fault via MissingSet
func (pc *propCollForWaitForPermsTest) WaitForUpdatesEx(
	ctx *simulator.Context,
	r *types.WaitForUpdatesEx) soap.HasFault {

	filter := ctx.Session.Get(pc.Filter[0]).(*simulator.PropertyFilter)

	if r.Version != "" {
		// Client should fail on the first response w/ MissingSet.
		// This ensures we don't get into a tight loop if that doesn't happen.
		select {}
	}

	return &methods.WaitForUpdatesExBody{
		Res: &types.WaitForUpdatesExResponse{
			Returnval: &types.UpdateSet{
				Version: "-",
				FilterSet: []types.PropertyFilterUpdate{{
					Filter: filter.Reference(),
					ObjectSet: []types.ObjectUpdate{{
						Kind: types.ObjectUpdateKindEnter,
						Obj:  filter.Spec.ObjectSet[0].Obj,
						MissingSet: []types.MissingProperty{{
							Path: "info",
							Fault: types.LocalizedMethodFault{
								Fault: new(types.NoPermission),
							},
						}},
					}},
				}},
			},
		},
	}
}
