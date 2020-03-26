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
package vsan

import (
	"context"
	"os"
	"testing"

	"github.com/kr/pretty"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vsan/types"
)

func TestClient(t *testing.T) {
	url := os.Getenv("VC_URL")            // example: export VC_URL='https://username:password@vc-ip/sdk'
	datacenter := os.Getenv("DATACENTER") // example: export DATACENTER='name-of-datacenter'
	if url == "" || datacenter == "" {
		t.Skip("VC_URL or DATACENTER is not set")
	}

	u, err := soap.ParseURL(url)
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	c, err := govmomi.NewClient(ctx, u, true)
	if err != nil {
		t.Fatal(err)
	}

	vsanHealthClient, err := NewClient(ctx, c.Client)
	if err != nil {
		t.Fatal(err)
	}

	finder := find.NewFinder(vsanHealthClient.vim25Client, false)
	dc, err := finder.Datacenter(ctx, datacenter)
	if err != nil {
		t.Fatal(err)
	}
	finder.SetDatacenter(dc)

	clusterComputeResource, err := finder.ClusterComputeResourceList(ctx, "*")
	if err != nil {
		t.Logf("Error occurred while getting clusterComputeResource %+v", err.Error())
		t.Fatal(err)
	}

	isFileServiceEnabled := false
	var clusterConfigToPrint *types.VsanConfigInfoEx
	for _, cluster := range clusterComputeResource {
		clusterConfig, err := vsanHealthClient.VsanClusterGetConfig(ctx, cluster.Reference())
		if err != nil {
			t.Logf("Error occurred: %+v", err.Error())
			t.Fatal(err)
		}
		if clusterConfig.FileServiceConfig.Enabled {
			clusterConfigToPrint = clusterConfig
			isFileServiceEnabled = true
		}
	}
	if isFileServiceEnabled {
		t.Logf("Printing one of the clusterConfig where file service is enabled:\n %+v", pretty.Sprint(clusterConfigToPrint))
	}
}
