/*
Copyright (c) 2021 VMware, Inc. All Rights Reserved.

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
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/types"
)

func sortMoRefSlice(a []types.ManagedObjectReference) {
	sort.SliceStable(a, func(i, j int) bool {
		lhs, rhs := a[i], a[j]
		switch strings.Compare(lhs.Type, rhs.Type) {
		case -1:
			return true
		case 1:
			return false
		}
		return lhs.Value < rhs.Value
	})
}

func TestTenantManagerVPX(t *testing.T) {
	ctx := context.Background()
	m := VPX()
	defer m.Remove()

	err := m.Create()
	if err != nil {
		t.Fatal(err)
	}

	s := m.Service.NewServer()
	defer s.Close()

	c, err := govmomi.NewClient(ctx, s.URL, true)
	if err != nil {
		t.Fatal(err)
	}

	tenantManager := object.NewTenantManager(c.Client)
	serviceProviderEntities := []types.ManagedObjectReference{
		{Type: "VirtualMachine", Value: "vm-123"},
		{Type: "HostSystem", Value: "host-1"},
	}
	sortMoRefSlice(serviceProviderEntities)

	// "Read your writes", mark entities and verify they are marked.
	err = tenantManager.MarkServiceProviderEntities(ctx, serviceProviderEntities)
	if err != nil {
		t.Fatal(err)
	}
	markedEntities, err := tenantManager.RetrieveServiceProviderEntities(ctx)
	if err != nil {
		t.Fatal(err)
	}
	sortMoRefSlice(markedEntities)
	if !reflect.DeepEqual(serviceProviderEntities, markedEntities) {
		t.Errorf("Requested-to-be-marked entities mismatch with acutally marked: %+v, %+v", serviceProviderEntities, markedEntities)
	}

	// Repeatedely mark same entities and verify they are deduped.
	err = tenantManager.MarkServiceProviderEntities(ctx, serviceProviderEntities)
	if err != nil {
		t.Fatal(err)
	}
	markedEntities, err = tenantManager.RetrieveServiceProviderEntities(ctx)
	if err != nil {
		t.Fatal(err)
	}
	sortMoRefSlice(markedEntities)
	if !reflect.DeepEqual(serviceProviderEntities, markedEntities) {
		t.Errorf("Requested-to-be-marked entities mismatch with acutally marked: %+v, %+v", serviceProviderEntities, markedEntities)
	}

	// Unmark not-previously-marked entity and verify no-op.
	unknownEntities := []types.ManagedObjectReference{{Type: "Folder", Value: "group-3"}}
	err = tenantManager.UnmarkServiceProviderEntities(ctx, unknownEntities)
	if err != nil {
		t.Fatal(err)
	}
	markedEntities, err = tenantManager.RetrieveServiceProviderEntities(ctx)
	if err != nil {
		t.Fatal(err)
	}
	sortMoRefSlice(markedEntities)
	if !reflect.DeepEqual(serviceProviderEntities, markedEntities) {
		t.Errorf("Requested-to-be-marked entities mismatch with acutally marked: %+v, %+v", serviceProviderEntities, markedEntities)
	}

	// Unmark marked entities and verify no longer marked.
	err = tenantManager.UnmarkServiceProviderEntities(ctx, serviceProviderEntities)
	if err != nil {
		t.Fatal(err)
	}
	markedEntities, err = tenantManager.RetrieveServiceProviderEntities(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(markedEntities) > 0 {
		t.Errorf("Expected all entities to be unmarked but still found marked: %+v", markedEntities)
	}
}
