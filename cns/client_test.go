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

package cns

import (
	"context"
	"os"

	"github.com/vmware/govmomi/object"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/vim25/soap"

	"testing"

	"github.com/vmware/govmomi"
	cnstypes "github.com/vmware/govmomi/cns/types"
	vim25types "github.com/vmware/govmomi/vim25/types"
)

func TestClient(t *testing.T) {
	url := os.Getenv("CNS_VC_URL")
	datacenter := os.Getenv("CNS_DATACENTER")
	datastore := os.Getenv("CNS_DATASTORE")
	if url == "" || datacenter == "" || datastore == "" {
		t.Skipf("CNS_VC_URL or CNS_DATACENTER or CNS_DATASTORE is not set")
		t.SkipNow()
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

	cnsClient, err := NewClient(ctx, c.Client)
	if err != nil {
		t.Fatal(err)
	}

	finder := find.NewFinder(cnsClient.vim25Client, false)
	dc, err := finder.Datacenter(ctx, datacenter)
	if err != nil {
		t.Fatal(err)
	}
	finder.SetDatacenter(dc)
	ds, err := finder.Datastore(ctx, datastore)
	if err != nil {
		t.Fatal(err)
	}

	var dsList []vim25types.ManagedObjectReference
	dsList = append(dsList, ds.Reference())

	// Test CreateVolume API
	var cnsVolumeCreateSpecList []cnstypes.CnsVolumeCreateSpec
	cnsVolumeCreateSpec := cnstypes.CnsVolumeCreateSpec{
		Name:       "pvc-901e87eb-c2bd-11e9-806f-005056a0c9a0",
		VolumeType: string(cnstypes.CnsVolumeTypeBlock),
		Datastores: dsList,
		Metadata: cnstypes.CnsVolumeMetadata{
			ContainerCluster: cnstypes.CnsContainerCluster{
				ClusterType: string(cnstypes.CnsClusterTypeKubernetes),
				ClusterId:   "demo-cluster-id",
				VSphereUser: "Administrator@vsphere.local",
			},
		},
		BackingObjectDetails: &cnstypes.CnsBackingObjectDetails{
			CapacityInMb: 5120,
		},
	}
	cnsVolumeCreateSpecList = append(cnsVolumeCreateSpecList, cnsVolumeCreateSpec)
	t.Logf("Creating volume using the spec: %+v", cnsVolumeCreateSpec)
	createTask, err := cnsClient.CreateVolume(ctx, cnsVolumeCreateSpecList)
	if err != nil {
		t.Errorf("Failed to create volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	createTaskInfo, err := GetTaskInfo(ctx, createTask)
	if err != nil {
		t.Errorf("Failed to create volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	createTaskResult, err := GetTaskResult(ctx, createTaskInfo)
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
	t.Logf("Volume created sucessfully. volumeId: %s", volumeId)

	// Test QueryVolume API
	var queryFilter cnstypes.CnsQueryFilter
	var volumeIDList []cnstypes.CnsVolumeId
	volumeIDList = append(volumeIDList, cnstypes.CnsVolumeId{Id: volumeId})
	queryFilter.VolumeIds = volumeIDList
	t.Logf("Calling QueryVolume using queryFilter: %+v", queryFilter)
	queryResult, err := cnsClient.QueryVolume(ctx, queryFilter)
	if err != nil {
		t.Errorf("Failed to query volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	t.Logf("Sucessfully Queried Volumes. queryResult: %+v", queryResult)

	// Test UpdateVolumeMetadata
	var updateSpecList []cnstypes.CnsVolumeMetadataUpdateSpec

	var metadataList []cnstypes.BaseCnsEntityMetadata
	newLabels := []vim25types.KeyValue{
		{
			Key:   "testLabel",
			Value: "testValue",
		},
	}
	metadata := &cnstypes.CnsKubernetesEntityMetadata{
		CnsEntityMetadata: cnstypes.CnsEntityMetadata{
			DynamicData: vim25types.DynamicData{},
			EntityName:  "PV NAME",
			Labels:      newLabels,
			Delete:      false,
		},
		EntityType: string(cnstypes.CnsKubernetesEntityTypePV),
		Namespace:  "",
	}
	metadataList = append(metadataList, cnstypes.BaseCnsEntityMetadata(metadata))
	cnsVolumeMetadataUpdateSpec := cnstypes.CnsVolumeMetadataUpdateSpec{
		VolumeId: cnstypes.CnsVolumeId{Id: volumeId},
		Metadata: cnstypes.CnsVolumeMetadata{
			DynamicData: vim25types.DynamicData{},
			ContainerCluster: cnstypes.CnsContainerCluster{
				ClusterType: string(cnstypes.CnsClusterTypeKubernetes),
				ClusterId:   "demo-cluster-id",
				VSphereUser: "Administrator@vsphere.local",
			},
			EntityMetadata: metadataList,
		},
	}
	t.Logf("Updating volume using the spec: %+v", cnsVolumeMetadataUpdateSpec)
	updateSpecList = append(updateSpecList, cnsVolumeMetadataUpdateSpec)
	updateTask, err := cnsClient.UpdateVolumeMetadata(ctx, updateSpecList)
	if err != nil {
		t.Errorf("Failed to update volume metadata. Error: %+v \n", err)
		t.Fatal(err)
	}
	updateTaskInfo, err := GetTaskInfo(ctx, updateTask)
	if err != nil {
		t.Errorf("Failed to update volume metadata. Error: %+v \n", err)
		t.Fatal(err)
	}
	updateTaskResult, err := GetTaskResult(ctx, updateTaskInfo)
	if err != nil {
		t.Errorf("Failed to update volume metadata. Error: %+v \n", err)
		t.Fatal(err)
	}
	if updateTaskResult == nil {
		t.Fatalf("Empty update task results")
		t.FailNow()
	}
	updateVolumeOperationRes := updateTaskResult.GetCnsVolumeOperationResult()
	if updateVolumeOperationRes.Fault != nil {
		t.Fatalf("Failed to update volume metadata: fault=%+v", updateVolumeOperationRes.Fault)
	} else {
		t.Logf("Sucessfully updated volume metadata")
	}

	// Test QueryAll
	querySelection := cnstypes.CnsQuerySelection{
		Names: []string{
			string(cnstypes.CnsQuerySelectionName_VOLUME_NAME),
			string(cnstypes.CnsQuerySelectionName_VOLUME_TYPE),
			string(cnstypes.CnsQuerySelectionName_BACKING_OBJECT_DETAILS),
			string(cnstypes.CnsQuerySelectionName_COMPLIANCE_STATUS),
			string(cnstypes.CnsQuerySelectionName_DATASTORE_ACCESSIBILITY_STATUS),
		},
	}
	queryResult, err = cnsClient.QueryAllVolume(ctx, cnstypes.CnsQueryFilter{}, querySelection)
	if err != nil {
		t.Errorf("Failed to query all volumes. Error: %+v \n", err)
		t.Fatal(err)
	}
	t.Logf("Sucessfully Queried all Volumes. queryResult: %+v", queryResult)

	// Create a VM to test Attach Volume API.
	virtualMachineConfigSpec := vim25types.VirtualMachineConfigSpec{
		Name: "test-node-vm",
		Files: &vim25types.VirtualMachineFileInfo{
			VmPathName: "[" + datastore + "]",
		},
		NumCPUs:  1,
		MemoryMB: 4,
		DeviceChange: []vim25types.BaseVirtualDeviceConfigSpec{
			&vim25types.VirtualDeviceConfigSpec{
				Operation: vim25types.VirtualDeviceConfigSpecOperationAdd,
				Device: &vim25types.ParaVirtualSCSIController{
					VirtualSCSIController: vim25types.VirtualSCSIController{
						SharedBus: vim25types.VirtualSCSISharingNoSharing,
						VirtualController: vim25types.VirtualController{
							BusNumber: 0,
							VirtualDevice: vim25types.VirtualDevice{
								Key: 1000,
							},
						},
					},
				},
			},
		},
	}
	defaultFolder, err := finder.DefaultFolder(ctx)
	defaultResourcePool, err := finder.DefaultResourcePool(ctx)
	task, err := defaultFolder.CreateVM(ctx, virtualMachineConfigSpec, defaultResourcePool, nil)
	if err != nil {
		t.Errorf("Failed to create VM. Error: %+v \n", err)
		t.Fatal(err)
	}

	vmTaskInfo, err := task.WaitForResult(ctx, nil)
	if err != nil {
		t.Errorf("Error occurred while waiting for create VM task result. err: %+v", err)
		t.Fatal(err)
	}

	vmRef := vmTaskInfo.Result.(object.Reference)
	t.Logf("Node VM created sucessfully. vmRef: %+v", vmRef.Reference())

	nodeVM := object.NewVirtualMachine(cnsClient.vim25Client, vmRef.Reference())
	defer nodeVM.Destroy(ctx)

	// Test AttachVolume API
	var cnsVolumeAttachSpecList []cnstypes.CnsVolumeAttachDetachSpec
	cnsVolumeAttachSpec := cnstypes.CnsVolumeAttachDetachSpec{
		VolumeId: cnstypes.CnsVolumeId{
			Id: volumeId,
		},
		Vm: nodeVM.Reference(),
	}
	cnsVolumeAttachSpecList = append(cnsVolumeAttachSpecList, cnsVolumeAttachSpec)
	t.Logf("Attaching volume using the spec: %+v", cnsVolumeAttachSpec)
	attachTask, err := cnsClient.AttachVolume(ctx, cnsVolumeAttachSpecList)
	if err != nil {
		t.Errorf("Failed to attach volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	attachTaskInfo, err := GetTaskInfo(ctx, attachTask)
	if err != nil {
		t.Errorf("Failed to attach volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	attachTaskResult, err := GetTaskResult(ctx, attachTaskInfo)
	if err != nil {
		t.Errorf("Failed to attach volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	if attachTaskResult == nil {
		t.Fatalf("Empty attach task results")
		t.FailNow()
	}
	attachVolumeOperationRes := attachTaskResult.GetCnsVolumeOperationResult()
	if attachVolumeOperationRes.Fault != nil {
		t.Fatalf("Failed to attach volume: fault=%+v", attachVolumeOperationRes.Fault)
	}
	diskUUID := attachVolumeOperationRes.VolumeId.Id
	t.Logf("Volume attached sucessfully. Disk UUID: %s", diskUUID)

	// Test DetachVolume API
	var cnsVolumeDetachSpecList []cnstypes.CnsVolumeAttachDetachSpec
	cnsVolumeDetachSpec := cnstypes.CnsVolumeAttachDetachSpec{
		VolumeId: cnstypes.CnsVolumeId{
			Id: volumeId,
		},
		Vm: nodeVM.Reference(),
	}
	cnsVolumeDetachSpecList = append(cnsVolumeDetachSpecList, cnsVolumeDetachSpec)
	t.Logf("Detaching volume using the spec: %+v", cnsVolumeDetachSpec)
	detachTask, err := cnsClient.DetachVolume(ctx, cnsVolumeDetachSpecList)
	if err != nil {
		t.Errorf("Failed to detach volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	detachTaskInfo, err := GetTaskInfo(ctx, detachTask)
	if err != nil {
		t.Errorf("Failed to detach volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	detachTaskResult, err := GetTaskResult(ctx, detachTaskInfo)
	if err != nil {
		t.Errorf("Failed to detach volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	if detachTaskResult == nil {
		t.Fatalf("Empty detach task results")
		t.FailNow()
	}
	detachVolumeOperationRes := detachTaskResult.GetCnsVolumeOperationResult()
	if detachVolumeOperationRes.Fault != nil {
		t.Fatalf("Failed to detach volume: fault=%+v", detachVolumeOperationRes.Fault)
	}
	t.Logf("Volume detached sucessfully")

	// Test DeleteVolume API
	t.Logf("Deleting volume: %+v", volumeIDList)
	deleteTask, err := cnsClient.DeleteVolume(ctx, volumeIDList, true)
	if err != nil {
		t.Errorf("Failed to delete volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	deleteTaskInfo, err := GetTaskInfo(ctx, deleteTask)
	if err != nil {
		t.Errorf("Failed to delete volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	deleteTaskResult, err := GetTaskResult(ctx, deleteTaskInfo)
	if err != nil {
		t.Errorf("Failed to detach volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	if deleteTaskResult == nil {
		t.Fatalf("Empty delete task results")
		t.FailNow()
	}
	deleteVolumeOperationRes := deleteTaskResult.GetCnsVolumeOperationResult()
	if deleteVolumeOperationRes.Fault != nil {
		t.Fatalf("Failed to delete volume: fault=%+v", deleteVolumeOperationRes.Fault)
	}
	t.Logf("Volume deleted sucessfully")
}
