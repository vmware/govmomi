// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"testing"

	"github.com/vmware/govmomi/simulator/vpx"
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
