// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"encoding/xml"
	"os"
	"path/filepath"
	"testing"

	"github.com/vmware/govmomi/simulator/vpx"
	"github.com/vmware/govmomi/vim25/types"
)

func compareModel(t *testing.T, m *Model) {
	count := m.Count()

	hosts := (m.Host + (m.ClusterHost * m.Cluster)) * m.Datacenter
	vms := ((m.Host + m.Cluster + m.Pool) * m.Datacenter) * m.Machine
	// child pools + root pools
	pools := (m.Pool * m.Cluster * m.Datacenter) + (m.Host+m.Cluster)*m.Datacenter
	// root folder + Datacenter folders {host,vm,datastore,network} + top-level folders
	folders := 1 + (4 * m.Datacenter) + (5 * m.Folder)
	pgs := m.Portgroup + m.PortgroupNSX
	if pgs > 0 {
		pgs++ // uplinks
	}
	pgs += m.OpaqueNetwork // pg for each opaque network

	tests := []struct {
		expect int
		actual int
		kind   string
	}{
		{m.Datacenter, count.Datacenter, "Datacenter"},
		{m.Cluster * m.Datacenter, count.Cluster, "Cluster"},
		{pgs * m.Datacenter, count.Portgroup, "Portgroup"},
		{m.OpaqueNetwork * m.Datacenter, count.OpaqueNetwork, "OpaqueNetwork"},
		{m.Datastore * m.Datacenter, count.Datastore, "Datastore"},
		{hosts, count.Host, "Host"},
		{vms, count.Machine, "VirtualMachine"},
		{pools, count.Pool, "ResourcePool"},
		{folders, count.Folder, "Folder"},
	}

	for _, test := range tests {
		if test.expect != test.actual {
			t.Errorf("expected %d %s, actual: %d", test.expect, test.kind, test.actual)
		}
	}
}

func TestModelESX(t *testing.T) {
	m := ESX()
	defer m.Remove()

	err := m.Create()
	if err != nil {
		t.Fatal(err)
	}

	// Set these for the compareModel math, and for m.Create to fail below
	m.Datacenter = 1
	m.Host = 1

	compareModel(t, m)

	err = m.Create()
	if err == nil {
		t.Error("expected error")
	}
}

func TestModelVPX(t *testing.T) {
	m := VPX()

	defer m.Remove()

	err := m.Create()
	if err != nil {
		t.Fatal(err)
	}

	compareModel(t, m)
}

func TestModelNoSwitchVPX(t *testing.T) {
	m := VPX()
	m.Portgroup = 0 // disabled DVS creation

	defer m.Remove()

	err := m.Create()
	if err != nil {
		t.Fatal(err)
	}

	compareModel(t, m)
}

func TestModelNSX(t *testing.T) {
	m := VPX()
	m.Portgroup = 0
	m.PortgroupNSX = 1
	m.OpaqueNetwork = 1

	defer m.Remove()

	err := m.Create()
	if err != nil {
		t.Fatal(err)
	}

	compareModel(t, m)
}

func TestModelCustomVPX(t *testing.T) {
	m := &Model{
		ServiceContent: vpx.ServiceContent,
		RootFolder:     vpx.RootFolder,
		Datacenter:     2,
		Cluster:        2,
		Host:           2,
		ClusterHost:    3,
		Datastore:      1,
		Machine:        3,
		Pool:           2,
		Portgroup:      2,
	}

	defer m.Remove()

	err := m.Create()
	if err != nil {
		t.Fatal(err)
	}

	compareModel(t, m)
}

func TestModelWithFolders(t *testing.T) {
	m := VPX()
	m.Datacenter = 3
	m.Folder = 2

	defer m.Remove()

	err := m.Create()
	if err != nil {
		t.Fatal(err)
	}

	compareModel(t, m)
}

func TestModelLoadAlignCounter(t *testing.T) {
	dir, err := os.MkdirTemp("", "vcsim-load-test-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	// Create a Datacenter XML file to prevent resolveReferences from creating a default one
	dcContent := types.ObjectContent{
		Obj: types.ManagedObjectReference{Type: "Datacenter", Value: "datacenter-10"},
	}
	dcData, err := xml.Marshal(dcContent)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "datacenter-10.xml"), dcData, 0644); err != nil {
		t.Fatal(err)
	}

	// Create a VirtualMachine XML file with vm-500
	vmContent := types.ObjectContent{
		Obj: types.ManagedObjectReference{Type: "VirtualMachine", Value: "vm-500"},
	}
	vmData, err := xml.Marshal(vmContent)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "vm-500.xml"), vmData, 0644); err != nil {
		t.Fatal(err)
	}

	// Load the model from the directory using the VPX baseline configuration
	m := VPX()
	defer m.Remove()

	if err := m.Load(dir); err != nil {
		t.Fatal(err)
	}

	// Verify the model registry counter was aligned to 500
	reg := m.Map()
	if reg.counter != 500 {
		for k, _ := range reg.objects {
			t.Logf("Registry object: %s:%s", k.Type, k.Value)
		}
		t.Errorf("expected registry counter to be 500, got %d", reg.counter)
	}
}
