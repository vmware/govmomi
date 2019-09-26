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
	"sync"
	"testing"

	"github.com/vmware/govmomi/vim25"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

func TestContainerViewVPX(t *testing.T) {
	ctx := context.Background()

	m := VPX()
	m.Datacenter = 3
	m.Folder = 2
	m.Pool = 1
	m.App = 1
	m.Pod = 1

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

	v := view.NewManager(c.Client)
	root := c.Client.ServiceContent.RootFolder

	// test container type validation
	_, err = v.CreateContainerView(ctx, v.Reference(), nil, false)
	if err == nil {
		t.Fatal("expected error")
	}

	// test container value validation
	_, err = v.CreateContainerView(ctx, types.ManagedObjectReference{Value: "enoent"}, nil, false)
	if err == nil {
		t.Fatal("expected error")
	}

	// test types validation
	_, err = v.CreateContainerView(ctx, root, []string{"enoent"}, false)
	if err == nil {
		t.Fatal("expected error")
	}

	vapp := object.NewVirtualApp(c.Client, Map.Any("VirtualApp").Reference())

	count := m.Count()

	tests := []struct {
		root    types.ManagedObjectReference
		recurse bool
		kinds   []string
		expect  int
	}{
		{root, false, nil, m.Datacenter - m.Folder + m.Folder},
		{root, true, nil, count.total - 1},                             // not including the root Folder
		{root, true, []string{"ManagedEntity"}, count.total - 1},       // not including the root Folder
		{root, true, []string{"Folder"}, count.Folder + count.Pod - 1}, // not including the root Folder
		{root, false, []string{"HostSystem"}, 0},
		{root, true, []string{"HostSystem"}, count.Host},
		{root, false, []string{"Datacenter"}, m.Datacenter - m.Folder},
		{root, true, []string{"Datacenter"}, count.Datacenter},
		{root, true, []string{"Datastore"}, count.Datastore},
		{root, true, []string{"VirtualMachine"}, count.Machine},
		{root, true, []string{"ResourcePool"}, count.Pool + count.App},
		{root, true, []string{"VirtualApp"}, count.App},
		{vapp.Reference(), true, []string{"VirtualMachine"}, m.Machine},
		{root, true, []string{"ClusterComputeResource"}, count.Cluster},
		{root, true, []string{"ComputeResource"}, (m.Cluster + m.Host) * m.Datacenter},
		{root, true, []string{"DistributedVirtualSwitch"}, count.Datacenter},
		{root, true, []string{"DistributedVirtualPortgroup"}, count.Portgroup},
		{root, true, []string{"Network"}, count.Portgroup + m.Datacenter},
		{root, true, []string{"OpaqueNetwork"}, 0},
		{root, true, []string{"StoragePod"}, m.Pod * m.Datacenter},
	}

	pc := property.DefaultCollector(c.Client)

	for i, test := range tests {
		cv, err := v.CreateContainerView(ctx, test.root, test.kinds, test.recurse)
		if err != nil {
			t.Fatal(err)
		}

		var mcv mo.ContainerView
		err = pc.RetrieveOne(ctx, cv.Reference(), nil, &mcv)
		if err != nil {
			t.Fatal(err)
		}

		n := len(mcv.View)

		if n != test.expect {
			t.Errorf("%d: %d != %d", i, n, test.expect)
		}

		err = cv.Destroy(ctx)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestViewManager_CreateContainerView(t *testing.T) {
	m := VPX()
	m.Datacenter = 10 // smaller numbers than this sometimes fail to trigger the DATA RACE
	m.Datastore = 10
	err := m.Run(func(ctx context.Context, client *vim25.Client) (err error) {
		manager := view.NewManager(client)
		datacenterView, err := manager.CreateContainerView(ctx, client.ServiceContent.RootFolder, []string{"Datacenter"}, true)
		if err != nil {
			return err
		}
		defer datacenterView.Destroy(ctx)
		var datacenterList []mo.Datacenter
		err = datacenterView.Retrieve(ctx, []string{"Datacenter"}, []string{"name", "datastoreFolder"}, &datacenterList)
		if err != nil {
			return err
		}
		var getDataStores = func(dc mo.Datacenter) (dsNames []string, err error) {
			dataStoreView, err := manager.CreateContainerView(ctx, dc.DatastoreFolder, []string{"Datastore"}, true)
			if err != nil {
				return dsNames, err
			}
			defer dataStoreView.Destroy(ctx)
			var dsList []mo.Datastore
			err = dataStoreView.Retrieve(ctx, []string{"Datastore"}, []string{"name"}, &dsList)
			if err != nil {
				return dsNames, err
			}
			for _, ds := range dsList {
				dsNames = append(dsNames, ds.Name)
			}
			return dsNames, err
		}
		wg := &sync.WaitGroup{}
		mtx := sync.Mutex{}
		var datastores [][]string
		wg.Add(len(datacenterList))
		for _, dc := range datacenterList {
			go func(ref mo.Datacenter) {
				defer wg.Done()
				ds, err := getDataStores(ref)
				if err != nil {
					return
				}
				mtx.Lock()
				datastores = append(datastores, ds)
				mtx.Unlock()
			}(dc)
		}
		wg.Wait()
		if len(datastores) != m.Datacenter {
			t.Errorf("Invalid number of datacenters: %d", len(datastores))
		} else {
			if len(datastores[0]) != m.Datastore {
				t.Errorf("Invalid number of datastores per datacenter: %d", len(datastores[0]))
			}
		}
		return nil
	})
	if err != nil {
		t.Errorf("Failed to run simulation: %s", err.Error())
	}
}
