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
	"path"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/simulator/esx"
	"github.com/vmware/govmomi/simulator/vpx"
	"github.com/vmware/govmomi/vim25/types"
)

func TestClusterESX(t *testing.T) {
	content := esx.ServiceContent
	s := New(NewServiceInstance(SpoofContext(), content, esx.RootFolder))

	ts := s.NewServer()
	defer ts.Close()

	ctx := context.Background()
	c, err := govmomi.NewClient(ctx, ts.URL, true)
	if err != nil {
		t.Fatal(err)
	}

	dc := object.NewDatacenter(c.Client, esx.Datacenter.Reference())

	folders, err := dc.Folders(ctx)
	if err != nil {
		t.Fatal(err)
	}

	_, err = folders.HostFolder.CreateCluster(ctx, "cluster1", types.ClusterConfigSpecEx{})
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestClusterVC(t *testing.T) {
	content := vpx.ServiceContent
	s := New(NewServiceInstance(SpoofContext(), content, vpx.RootFolder))

	ts := s.NewServer()
	defer ts.Close()

	ctx := context.Background()
	c, err := govmomi.NewClient(ctx, ts.URL, true)
	if err != nil {
		t.Fatal(err)
	}

	f := object.NewRootFolder(c.Client)

	dc, err := f.CreateDatacenter(ctx, "foo")
	if err != nil {
		t.Error(err)
	}

	folders, err := dc.Folders(ctx)
	if err != nil {
		t.Fatal(err)
	}

	cluster, err := folders.HostFolder.CreateCluster(ctx, "cluster1", types.ClusterConfigSpecEx{})
	if err != nil {
		t.Fatal(err)
	}

	// Enable DRS and HA for the cluster
	clusterSpec := types.ClusterConfigSpecEx{
		DrsConfig: &types.ClusterDrsConfigInfo{
			Enabled:           types.NewBool(true),
			DefaultVmBehavior: types.DrsBehaviorFullyAutomated, // Set DRS to fully automated
		},
		DasConfig: &types.ClusterDasConfigInfo{
			Enabled: types.NewBool(true),
		},
	}

	task, err := cluster.Reconfigure(ctx, &clusterSpec, true)
	require.NoError(t, err)
	err = task.Wait(ctx)
	require.NoError(t, err)

	// Check if DRS and HA is set

	finder := find.NewFinder(c.Client, false).SetDatacenter(dc)
	pathname := path.Join(dc.InventoryPath, "host", "cluster1")
	clusterComputeResource, err := finder.ClusterComputeResource(ctx, pathname)
	require.NoError(t, err)
	clusterComputeInfo, err := clusterComputeResource.Configuration(ctx)
	require.NoError(t, err)
	require.NotNil(t, clusterComputeInfo.DrsConfig.Enabled)
	require.NotNil(t, clusterComputeInfo.DasConfig.Enabled)
	require.True(t, *clusterComputeInfo.DrsConfig.Enabled)
	require.True(t, *clusterComputeInfo.DasConfig.Enabled)
	require.True(t, clusterComputeInfo.DrsConfig.DefaultVmBehavior == types.DrsBehaviorFullyAutomated)
	_, err = folders.HostFolder.CreateCluster(ctx, "cluster1", types.ClusterConfigSpecEx{})
	if err == nil {
		t.Error("expected DuplicateName error")
	}

	spec := types.HostConnectSpec{}

	for _, fail := range []bool{true, false} {
		task, err := cluster.AddHost(ctx, spec, true, nil, nil)
		if err != nil {
			t.Fatal(err)
		}

		_, err = task.WaitForResult(ctx, nil)

		if fail {
			if err == nil {
				t.Error("expected error")
			}
			spec.HostName = "localhost"
		} else {
			if err != nil {
				t.Error(err)
			}
		}
	}
}
