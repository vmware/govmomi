// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package mo

import (
	"os"
	"testing"
	"time"

	"github.com/vmware/govmomi/fault"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vim25/xml"
)

func load(name string) []types.ObjectContent {
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	var b types.RetrievePropertiesResponse

	dec := xml.NewDecoder(f)
	dec.TypeFunc = types.TypeFunc()
	if err := dec.Decode(&b); err != nil {
		panic(err)
	}

	return b.Returnval
}

func TestNotAuthenticatedFault(t *testing.T) {
	var s SessionManager

	err := LoadObjectContent(load("fixtures/not_authenticated_fault.xml"), &s)
	if !soap.IsVimFault(err) {
		t.Errorf("Expected IsVimFault")
	}

	var not *types.NotAuthenticated
	fault.As(err, &not)
	if not.PrivilegeId != "System.View" {
		t.Errorf("Expected first fault to be returned")
	}
}

func TestNestedProperty(t *testing.T) {
	var vm VirtualMachine

	err := LoadObjectContent(load("fixtures/nested_property.xml"), &vm)
	if err != nil {
		t.Fatalf("Expected no error, got: %s", err)
	}

	self := types.ManagedObjectReference{
		Type:  "VirtualMachine",
		Value: "vm-411",
	}

	if vm.Self != self {
		t.Fatalf("Expected vm.Self to be set")
	}

	if vm.Config == nil {
		t.Fatalf("Expected vm.Config to be set")
	}

	if vm.Config.Name != "kubernetes-master" {
		t.Errorf("Got: %s", vm.Config.Name)
	}

	if vm.Config.Uuid != "422ec880-ab06-06b4-23f3-beb7a052a4c9" {
		t.Errorf("Got: %s", vm.Config.Uuid)
	}
}

func TestPointerProperty(t *testing.T) {
	var vm VirtualMachine

	err := LoadObjectContent(load("fixtures/pointer_property.xml"), &vm)
	if err != nil {
		t.Fatalf("Expected no error, got: %s", err)
	}

	if vm.Config == nil {
		t.Fatalf("Expected vm.Config to be set")
	}

	if vm.Config.BootOptions == nil {
		t.Fatalf("Expected vm.Config.BootOptions to be set")
	}
}

func TestEmbeddedTypeProperty(t *testing.T) {
	// Test that we avoid in this case:
	// panic: reflect.Set: value of type mo.ClusterComputeResource is not assignable to type mo.ComputeResource
	var cr ComputeResource

	err := LoadObjectContent(load("fixtures/cluster_host_property.xml"), &cr)
	if err != nil {
		t.Fatalf("Expected no error, got: %s", err)
	}

	if len(cr.Host) != 4 {
		t.Fatalf("Expected cr.Host to be set")
	}
}

func TestEmbeddedTypePropertySlice(t *testing.T) {
	var me []ManagedEntity

	err := LoadObjectContent(load("fixtures/hostsystem_list_name_property.xml"), &me)
	if err != nil {
		t.Fatalf("Expected no error, got: %s", err)
	}

	if len(me) != 2 {
		t.Fatalf("Expected 2 elements")
	}

	for _, m := range me {
		if m.Name == "" {
			t.Fatal("Expected Name field to be set")
		}
	}

	if me[0].Name == me[1].Name {
		t.Fatal("Name fields should not be the same")
	}
}

func TestReferences(t *testing.T) {
	var cr ComputeResource

	err := LoadObjectContent(load("fixtures/cluster_host_property.xml"), &cr)
	if err != nil {
		t.Fatalf("Expected no error, got: %s", err)
	}

	refs := References(cr)
	n := len(refs)
	if n != 5 {
		t.Errorf("%d refs", n)
	}
}

func TestEventReferences(t *testing.T) {
	event := &types.VmPoweredOnEvent{
		VmEvent: types.VmEvent{
			Event: types.Event{
				Key:         0,
				ChainId:     0,
				CreatedTime: time.Now(),
				UserName:    "",
				Datacenter: &types.DatacenterEventArgument{
					EntityEventArgument: types.EntityEventArgument{
						EventArgument: types.EventArgument{},
						Name:          "DC0",
					},
					Datacenter: types.ManagedObjectReference{Type: "Datacenter", Value: "datacenter-2"},
				},
				ComputeResource: &types.ComputeResourceEventArgument{
					EntityEventArgument: types.EntityEventArgument{
						EventArgument: types.EventArgument{},
						Name:          "DC0_C0",
					},
					ComputeResource: types.ManagedObjectReference{Type: "ClusterComputeResource", Value: "clustercomputeresource-26"},
				},
				Host: &types.HostEventArgument{
					EntityEventArgument: types.EntityEventArgument{
						EventArgument: types.EventArgument{},
						Name:          "DC0_C0_H0",
					},
					Host: types.ManagedObjectReference{Type: "HostSystem", Value: "host-32"},
				},
				Vm: &types.VmEventArgument{
					EntityEventArgument: types.EntityEventArgument{
						EventArgument: types.EventArgument{},
						Name:          "DC0_C0_RP0_VM1",
					},
					Vm: types.ManagedObjectReference{Type: "VirtualMachine", Value: "vm-62"},
				},
				Ds:                   (*types.DatastoreEventArgument)(nil),
				Net:                  (*types.NetworkEventArgument)(nil),
				Dvs:                  (*types.DvsEventArgument)(nil),
				FullFormattedMessage: "",
				ChangeTag:            "",
			},
			Template: false,
		},
	}

	refs := References(event, true)
	n := len(refs)
	if n != 4 {
		t.Errorf("%d refs", n)
	}
}

func TestDatastoreInfoURL(t *testing.T) {
	// Datastore.Info is types.BaseDatastoreInfo
	// LoadObjectContent() should populate Info with the base type (*types.DatastoreInfo) in this case
	content := []types.ObjectContent{
		{
			Obj: types.ManagedObjectReference{Type: "Datastore", Value: "datastore-48", ServerGUID: ""},
			PropSet: []types.DynamicProperty{
				{
					Name: "info.url",
					Val:  "ds:///vmfs/volumes/666d7a79-cb0d28b2-57c8-0645602e1b58/",
				},
				{
					Name: "info.name",
					Val:  "foo",
				},
			},
			MissingSet: nil,
		},
	}

	var ds Datastore

	if err := LoadObjectContent(content, &ds); err != nil {
		t.Fatal(err)
	}

	info := ds.Info.GetDatastoreInfo()

	if info.Url != content[0].PropSet[0].Val.(string) {
		t.Errorf("info.url=%s", info.Url)
	}

	if info.Name != content[0].PropSet[1].Val.(string) {
		t.Errorf("info.name=%s", info.Name)
	}
}

func TestIgnoreMissingProperty(t *testing.T) {
	// subset of RetrievePropertiesEx response, see ignoreMissingProperty()
	content := []types.ObjectContent{{
		Obj: types.ManagedObjectReference{Type: "ResourcePool", Value: "ha-root-pool"},
		PropSet: []types.DynamicProperty{{
			Name: "config",
			Val: types.ResourceConfigSpec{
				CpuAllocation: types.ResourceAllocationInfo{
					Reservation: types.NewInt64(5585),
					Limit:       types.NewInt64(5585),
				},
				MemoryAllocation: types.ResourceAllocationInfo{
					Reservation: types.NewInt64(13652),
					Limit:       types.NewInt64(13652),
				},
			},
		}},
		MissingSet: []types.MissingProperty{
			{
				Path: "resourceConfigSpecDetailed",
				Fault: types.LocalizedMethodFault{
					Fault: &types.SystemError{
						Reason: "unexpected error reading property",
					},
				},
			},
		},
	}}

	var pool ResourcePool

	if err := LoadObjectContent(content, &pool); err != nil {
		t.Fatal(err)
	}
}
