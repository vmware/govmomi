/*
Copyright (c) 2014 VMware, Inc. All Rights Reserved.

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

package mo

import (
	"os"
	"testing"

	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vim25/xml"
)

func load(name string) *types.RetrievePropertiesResponse {
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

	return &b
}

func TestNotAuthenticatedFault(t *testing.T) {
	var s SessionManager

	err := LoadRetrievePropertiesResponse(load("fixtures/not_authenticated_fault.xml"), &s)
	if !soap.IsVimFault(err) {
		t.Errorf("Expected IsVimFault")
	}

	fault := soap.ToVimFault(err).(*types.NotAuthenticated)
	if fault.PrivilegeId != "System.View" {
		t.Errorf("Expected first fault to be returned")
	}
}

func TestNestedProperty(t *testing.T) {
	var vm VirtualMachine

	err := LoadRetrievePropertiesResponse(load("fixtures/nested_property.xml"), &vm)
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

	err := LoadRetrievePropertiesResponse(load("fixtures/pointer_property.xml"), &vm)
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
