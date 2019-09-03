/*
Copyright (c) 2019 VMware, Inc. All Rights Reserved.

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

type PropertyCollector struct {
	simulator.PropertyCollector
}

// CreatePropertyCollector overrides the vcsim impl to return this test's PC impl
func (pc *PropertyCollector) CreatePropertyCollector(ctx *simulator.Context, c *types.CreatePropertyCollector) soap.HasFault {
	return &methods.CreatePropertyCollectorBody{
		Res: &types.CreatePropertyCollectorResponse{
			Returnval: ctx.Session.Put(new(PropertyCollector)).Reference(),
		},
	}
}

// WaitForUpdatesEx overrides the vcsim impl to inject a fault via MissingSet
func (pc *PropertyCollector) WaitForUpdatesEx(ctx *simulator.Context, r *types.WaitForUpdatesEx) soap.HasFault {
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

	pc := new(PropertyCollector)
	pc.Self = model.ServiceContent.PropertyCollector
	simulator.Map.Put(pc)

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
