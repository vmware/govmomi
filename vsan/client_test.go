// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vsan

import (
	"context"
	"os"
	"testing"

	"github.com/dougm/pretty"

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

	finder := find.NewFinder(c.Client, false)
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

	hosts, _ := finder.HostSystemList(ctx, "*")
	if err != nil {
		t.Logf("Error occurred while getting hostSystem %+v", err.Error())
		t.Fatal(err)
	}

	for _, host := range hosts {
		vsanSystem, _ := host.ConfigManager().VsanInternalSystem(ctx)
		if err != nil {
			t.Logf("Error occurred: %+v", err.Error())
			t.Fatal(err)
		}
		hostConfig, err := vsanHealthClient.VsanHostGetConfig(ctx, vsanSystem.Reference())
		if err != nil {
			t.Logf("Error occurred: %+v", err.Error())
			t.Fatal(err)
		}
		t.Logf("Printing hostConfig:\n %+s", pretty.Sprint(hostConfig))
	}
}

func TestVsanQueryObjectIdentities(t *testing.T) {
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

	finder := find.NewFinder(c.Client, false)
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

	for _, cluster := range clusterComputeResource {
		clusterConfig, err := vsanHealthClient.VsanQueryObjectIdentities(ctx, cluster.Reference())
		if err != nil {
			t.Logf("Error occurred: %+v", err.Error())
			t.Fatal(err)
		}
		t.Logf("Printing clusterConfig:\n %+s", pretty.Sprint(clusterConfig))
	}

}
