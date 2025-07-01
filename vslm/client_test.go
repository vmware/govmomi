// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vslm

import (
	"context"
	"os"
	"testing"

	"github.com/dougm/pretty"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/cns"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"

	cnstypes "github.com/vmware/govmomi/cns/types"
)

func TestClient(t *testing.T) {
	url := os.Getenv("VSLM_VC_URL")            // export VC_URL='https://administrator@vsphere.local:Admin!23@10.184.69.227/sdk'
	datacenter := os.Getenv("VSLM_DATACENTER") // export DATACENTER='test-vpx-1614463083-15492-hostpool'
	volumePath := os.Getenv("VOLUME_PATH")     // export VOLUME_PATH='https://10.186.43.166/folder/27b74660-98cd-3fe5-514f-02009d24d7ab/vm-1.vmdk?dcPath=datacenter&dsName=vsanDatastore'
	if url == "" || datacenter == "" || volumePath == "" {
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

	vslmClient, err := NewClient(ctx, c.Client)
	if err != nil {
		t.Fatal(err)
	}

	globalObjectManager := NewGlobalObjectManager(vslmClient)
	vStorageObject, err := globalObjectManager.RegisterDisk(ctx, volumePath, "volume-name1")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully registered disk with path %s as FCD with storage object id %s", volumePath, vStorageObject.Config.Id.Id)

	containerCluster := cnstypes.CnsContainerCluster{
		ClusterType:         string(cnstypes.CnsClusterTypeKubernetes),
		ClusterId:           "demo-cluster-id",
		VSphereUser:         "Administrator@vsphere.local",
		ClusterFlavor:       string(cnstypes.CnsClusterFlavorVanilla),
		ClusterDistribution: "KUBERNETES",
	}

	// Test CreateVolume API
	var cnsVolumeCreateSpecList []cnstypes.CnsVolumeCreateSpec
	cnsVolumeCreateSpec := cnstypes.CnsVolumeCreateSpec{
		Name:       "pvc-abc123",
		VolumeType: string(cnstypes.CnsVolumeTypeBlock),
		Metadata: cnstypes.CnsVolumeMetadata{
			ContainerCluster: containerCluster,
		},
		BackingObjectDetails: &cnstypes.CnsBlockBackingDetails{
			CnsBackingObjectDetails: cnstypes.CnsBackingObjectDetails{
				CapacityInMb: 5120,
			},
			BackingDiskId: vStorageObject.Config.Id.Id,
		},
	}
	c.UseServiceVersion("vsan")
	cnsClient, err := cns.NewClient(ctx, c.Client)
	if err != nil {
		t.Fatal(err)
	}
	cnsVolumeCreateSpecList = append(cnsVolumeCreateSpecList, cnsVolumeCreateSpec)
	t.Logf("Creating volume using the spec: %+v", pretty.Sprint(cnsVolumeCreateSpec))
	createTask, err := cnsClient.CreateVolume(ctx, cnsVolumeCreateSpecList)
	if err != nil {
		t.Errorf("Failed to create volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	createTaskInfo, err := cns.GetTaskInfo(ctx, createTask)
	if err != nil {
		t.Errorf("Failed to create volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	createTaskResult, err := cns.GetTaskResult(ctx, createTaskInfo)
	if err != nil {
		t.Errorf("Failed to create volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	if createTaskResult == nil {
		t.Fatalf("Empty create task results")
		t.FailNow()
	}
	createVolumeOperationRes := createTaskResult.GetCnsVolumeOperationResult()
	if createVolumeOperationRes.Fault != nil {
		t.Fatalf("Failed to create volume: fault=%+v", createVolumeOperationRes.Fault)
	}
	volumeId := createVolumeOperationRes.VolumeId.Id
	volumeCreateResult := (createTaskResult).(*cnstypes.CnsVolumeCreateResult)
	t.Logf("volumeCreateResult %+v", volumeCreateResult)
	t.Logf("Volume created sucessfully. volumeId: %s", volumeId)

	err = globalObjectManager.SetControlFlags(ctx, types.ID{Id: volumeId}, []string{
		string(types.VslmVStorageObjectControlFlagKeepAfterDeleteVm)})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Control flag: FCD_KEEP_AFTER_DELETE_VM set for the volumeId: %s", volumeId)
	err = globalObjectManager.ClearControlFlags(ctx, types.ID{Id: volumeId})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Control flags removed the volumeId: %s", volumeId)
}
