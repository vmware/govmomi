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
	"testing"

	"github.com/kr/pretty"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/debug"
	"github.com/vmware/govmomi/vim25/soap"

	"github.com/vmware/govmomi"
	cnstypes "github.com/vmware/govmomi/cns/types"
	vim25types "github.com/vmware/govmomi/vim25/types"
	vsanfstypes "github.com/vmware/govmomi/vsan/vsanfs/types"
)

func TestClient(t *testing.T) {
	// set CNS_DEBUG to true if you need to emit soap traces from these tests
	// soap traces will be emitted in the govmomi/cns/.soap directory
	// example export CNS_DEBUG='true'
	enableDebug := os.Getenv("CNS_DEBUG")
	soapTraceDirectory := ".soap"

	url := os.Getenv("CNS_VC_URL") // example: export CNS_VC_URL='https://username:password@vc-ip/sdk'
	datacenter := os.Getenv("CNS_DATACENTER")
	datastore := os.Getenv("CNS_DATASTORE")

	// set CNS_RUN_FILESHARE_TESTS environment to true, if your setup has vsanfileshare enabled.
	// when CNS_RUN_FILESHARE_TESTS is not set to true, vsan file share related tests are skipped.
	// example: export export CNS_RUN_FILESHARE_TESTS='true'
	run_fileshare_tests := os.Getenv("CNS_RUN_FILESHARE_TESTS")

	// if backingDiskURLPath is not set, test for Creating Volume with setting BackingDiskUrlPath in the BackingObjectDetails of
	// CnsVolumeCreateSpec will be skipped.
	// example: Export BACKING_DISK_URL_PATH='https://vc-ip/folder/vmdkfilePath.vmdk?dcPath=DataCenterPath&dsName=DataStoreName'
	backingDiskURLPath := os.Getenv("BACKING_DISK_URL_PATH")

	if url == "" || datacenter == "" || datastore == "" {
		t.Skip("CNS_VC_URL or CNS_DATACENTER or CNS_DATASTORE is not set")
	}
	resporcePoolPath := os.Getenv("CNS_RESOURCE_POOL_PATH") // example "/datacenter-name/host/host-ip/Resources" or  /datacenter-name/host/cluster-name/Resources
	u, err := soap.ParseURL(url)
	if err != nil {
		t.Fatal(err)
	}

	if enableDebug == "true" {
		if _, err := os.Stat(soapTraceDirectory); os.IsNotExist(err) {
			os.Mkdir(soapTraceDirectory, 0755)
		}
		p := debug.FileProvider{
			Path: soapTraceDirectory,
		}
		debug.SetProvider(&p)
	}

	ctx := context.Background()
	c, err := govmomi.NewClient(ctx, u, true)
	if err != nil {
		t.Fatal(err)
	}
	// UseServiceVersion sets soap.Client.Version to the current version of the service endpoint via /sdk/vsanServiceVersions.xml
	c.UseServiceVersion("vsan")
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

	var containerClusterArray []cnstypes.CnsContainerCluster
	containerCluster := cnstypes.CnsContainerCluster{
		ClusterType:   string(cnstypes.CnsClusterTypeKubernetes),
		ClusterId:     "demo-cluster-id",
		VSphereUser:   "Administrator@vsphere.local",
		ClusterFlavor: string(cnstypes.CnsClusterFlavorVanilla),
	}
	containerClusterArray = append(containerClusterArray, containerCluster)

	// Test CreateVolume API
	var cnsVolumeCreateSpecList []cnstypes.CnsVolumeCreateSpec
	cnsVolumeCreateSpec := cnstypes.CnsVolumeCreateSpec{
		Name:       "pvc-901e87eb-c2bd-11e9-806f-005056a0c9a0",
		VolumeType: string(cnstypes.CnsVolumeTypeBlock),
		Datastores: dsList,
		Metadata: cnstypes.CnsVolumeMetadata{
			ContainerCluster: containerCluster,
		},
		BackingObjectDetails: &cnstypes.CnsBlockBackingDetails{
			CnsBackingObjectDetails: cnstypes.CnsBackingObjectDetails{
				CapacityInMb: 5120,
			},
		},
	}
	cnsVolumeCreateSpecList = append(cnsVolumeCreateSpecList, cnsVolumeCreateSpec)
	t.Logf("Creating volume using the spec: %+v", pretty.Sprint(cnsVolumeCreateSpec))
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

	if cnsClient.serviceClient.Version != ReleaseVSAN67u3 {
		// Test creating static volume using existing CNS volume should fail
		var staticCnsVolumeCreateSpecList []cnstypes.CnsVolumeCreateSpec
		staticCnsVolumeCreateSpec := cnstypes.CnsVolumeCreateSpec{
			Name:       "pvc-901e87eb-c2bd-11e9-806f-005056a0c9a0",
			VolumeType: string(cnstypes.CnsVolumeTypeBlock),
			Metadata: cnstypes.CnsVolumeMetadata{
				ContainerCluster: containerCluster,
			},
			BackingObjectDetails: &cnstypes.CnsBlockBackingDetails{
				CnsBackingObjectDetails: cnstypes.CnsBackingObjectDetails{
					CapacityInMb: 5120,
				},
				BackingDiskId: volumeId,
			},
		}

		staticCnsVolumeCreateSpecList = append(staticCnsVolumeCreateSpecList, staticCnsVolumeCreateSpec)
		t.Logf("Creating volume using the spec: %+v", pretty.Sprint(staticCnsVolumeCreateSpec))
		recreateTask, err := cnsClient.CreateVolume(ctx, staticCnsVolumeCreateSpecList)
		if err != nil {
			t.Errorf("Failed to create volume. Error: %+v \n", err)
			t.Fatal(err)
		}
		reCreateTaskInfo, err := GetTaskInfo(ctx, recreateTask)
		if err != nil {
			t.Errorf("Failed to create volume. Error: %+v \n", err)
			t.Fatal(err)
		}
		reCreateTaskResult, err := GetTaskResult(ctx, reCreateTaskInfo)
		if err != nil {
			t.Errorf("Failed to create volume. Error: %+v \n", err)
			t.Fatal(err)
		}
		if reCreateTaskResult == nil {
			t.Fatalf("Empty create task results")
			t.FailNow()
		}
		reCreateVolumeOperationRes := reCreateTaskResult.GetCnsVolumeOperationResult()
		t.Logf("reCreateVolumeOperationRes.: %+v", pretty.Sprint(reCreateVolumeOperationRes))
		if reCreateVolumeOperationRes.Fault != nil {
			t.Logf("reCreateVolumeOperationRes.Fault: %+v", pretty.Sprint(reCreateVolumeOperationRes.Fault))
			_, ok := reCreateVolumeOperationRes.Fault.Fault.(cnstypes.CnsFault)
			if !ok {
				t.Fatalf("Fault is not CnsFault")
			}
		} else {
			t.Fatalf("re-create same volume should fail with CnsFault")
		}
	}

	// Test QueryVolume API
	var queryFilter cnstypes.CnsQueryFilter
	var volumeIDList []cnstypes.CnsVolumeId
	volumeIDList = append(volumeIDList, cnstypes.CnsVolumeId{Id: volumeId})
	queryFilter.VolumeIds = volumeIDList
	t.Logf("Calling QueryVolume using queryFilter: %+v", pretty.Sprint(queryFilter))
	queryResult, err := cnsClient.QueryVolume(ctx, queryFilter)
	if err != nil {
		t.Errorf("Failed to query volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	t.Logf("Successfully Queried Volumes. queryResult: %+v", pretty.Sprint(queryResult))

	// Test QueryVolumeInfo API
	// QueryVolumeInfo is not supported on ReleaseVSAN67u3 and ReleaseVSAN70
	// This API is available on vSphere 7.0u1 onward
	if cnsClient.serviceClient.Version != ReleaseVSAN67u3 && cnsClient.serviceClient.Version != ReleaseVSAN70 {
		t.Logf("Calling QueryVolumeInfo using: %+v", pretty.Sprint(volumeIDList))
		queryVolumeInfoTask, err := cnsClient.QueryVolumeInfo(ctx, volumeIDList)
		if err != nil {
			t.Errorf("Failed to query volumes with QueryVolumeInfo. Error: %+v \n", err)
			t.Fatal(err)
		}
		queryVolumeInfoTaskInfo, err := GetTaskInfo(ctx, queryVolumeInfoTask)
		if err != nil {
			t.Errorf("Failed to query volumes with QueryVolumeInfo. Error: %+v \n", err)
			t.Fatal(err)
		}
		queryVolumeInfoTaskResults, err := GetTaskResultArray(ctx, queryVolumeInfoTaskInfo)
		if err != nil {
			t.Errorf("Failed to query volumes with QueryVolumeInfo. Error: %+v \n", err)
			t.Fatal(err)
		}
		if queryVolumeInfoTaskResults == nil {
			t.Fatalf("Empty queryVolumeInfoTaskResult")
			t.FailNow()
		}
		for _, queryVolumeInfoTaskResult := range queryVolumeInfoTaskResults {
			queryVolumeInfoOperationRes := queryVolumeInfoTaskResult.GetCnsVolumeOperationResult()
			if queryVolumeInfoOperationRes.Fault != nil {
				t.Fatalf("Failed to query volumes with QueryVolumeInfo. fault=%+v", queryVolumeInfoOperationRes.Fault)
			}
			t.Logf("Successfully Queried Volumes. queryVolumeInfoTaskResult: %+v", pretty.Sprint(queryVolumeInfoTaskResult))
		}
	}
	// Test ExtendVolume API
	var newCapacityInMb int64 = 10240
	var cnsVolumeExtendSpecList []cnstypes.CnsVolumeExtendSpec
	cnsVolumeExtendSpec := cnstypes.CnsVolumeExtendSpec{
		VolumeId: cnstypes.CnsVolumeId{
			Id: volumeId,
		},
		CapacityInMb: newCapacityInMb,
	}
	cnsVolumeExtendSpecList = append(cnsVolumeExtendSpecList, cnsVolumeExtendSpec)
	t.Logf("Extending volume using the spec: %+v", pretty.Sprint(cnsVolumeExtendSpecList))
	extendTask, err := cnsClient.ExtendVolume(ctx, cnsVolumeExtendSpecList)
	if err != nil {
		t.Errorf("Failed to extend volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	extendTaskInfo, err := GetTaskInfo(ctx, extendTask)
	if err != nil {
		t.Errorf("Failed to extend volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	extendTaskResult, err := GetTaskResult(ctx, extendTaskInfo)
	if err != nil {
		t.Errorf("Failed to extend volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	if extendTaskResult == nil {
		t.Fatalf("Empty extend task results")
		t.FailNow()
	}
	extendVolumeOperationRes := extendTaskResult.GetCnsVolumeOperationResult()
	if extendVolumeOperationRes.Fault != nil {
		t.Fatalf("Failed to extend volume: fault=%+v", extendVolumeOperationRes.Fault)
	}
	extendVolumeId := extendVolumeOperationRes.VolumeId.Id
	t.Logf("Volume extended sucessfully. Volume ID: %s", extendVolumeId)

	// Verify volume is extended to the specified size
	t.Logf("Calling QueryVolume after ExtendVolume using queryFilter: %+v", queryFilter)
	queryResult, err = cnsClient.QueryVolume(ctx, queryFilter)
	if err != nil {
		t.Errorf("Failed to query volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	t.Logf("Successfully Queried Volumes after ExtendVolume. queryResult: %+v", pretty.Sprint(queryResult))
	queryCapacity := queryResult.Volumes[0].BackingObjectDetails.(*cnstypes.CnsBlockBackingDetails).CapacityInMb
	if newCapacityInMb != queryCapacity {
		t.Errorf("After extend volume %s, expected new volume size is %d, but actual volume size is %d.", extendVolumeId, newCapacityInMb, queryCapacity)
	} else {
		t.Logf("Volume extended sucessfully to the new size. Volume ID: %s New Size: %d", extendVolumeId, newCapacityInMb)
	}

	// Test UpdateVolumeMetadata
	var updateSpecList []cnstypes.CnsVolumeMetadataUpdateSpec

	var metadataList []cnstypes.BaseCnsEntityMetadata
	newLabels := []vim25types.KeyValue{
		{
			Key:   "testLabel",
			Value: "testValue",
		},
	}
	pvmetadata := &cnstypes.CnsKubernetesEntityMetadata{
		CnsEntityMetadata: cnstypes.CnsEntityMetadata{
			DynamicData: vim25types.DynamicData{},
			EntityName:  "pvc-53465372-5c12-4818-96f8-0ace4f4fd116",
			Labels:      newLabels,
			Delete:      false,
			ClusterID:   "demo-cluster-id",
		},
		EntityType: string(cnstypes.CnsKubernetesEntityTypePV),
		Namespace:  "",
	}
	metadataList = append(metadataList, cnstypes.BaseCnsEntityMetadata(pvmetadata))

	pvcmetadata := &cnstypes.CnsKubernetesEntityMetadata{
		CnsEntityMetadata: cnstypes.CnsEntityMetadata{
			DynamicData: vim25types.DynamicData{},
			EntityName:  "example-vanilla-block-pvc",
			Labels:      newLabels,
			Delete:      false,
			ClusterID:   "demo-cluster-id",
		},
		EntityType: string(cnstypes.CnsKubernetesEntityTypePVC),
		Namespace:  "default",
		ReferredEntity: []cnstypes.CnsKubernetesEntityReference{
			cnstypes.CnsKubernetesEntityReference{
				EntityType: string(cnstypes.CnsKubernetesEntityTypePV),
				EntityName: "pvc-53465372-5c12-4818-96f8-0ace4f4fd116",
				Namespace:  "",
				ClusterID:  "demo-cluster-id",
			},
		},
	}
	metadataList = append(metadataList, cnstypes.BaseCnsEntityMetadata(pvcmetadata))

	podmetadata := &cnstypes.CnsKubernetesEntityMetadata{
		CnsEntityMetadata: cnstypes.CnsEntityMetadata{
			DynamicData: vim25types.DynamicData{},
			EntityName:  "example-pod",
			Delete:      false,
			ClusterID:   "demo-cluster-id",
		},
		EntityType: string(cnstypes.CnsKubernetesEntityTypePOD),
		Namespace:  "default",
		ReferredEntity: []cnstypes.CnsKubernetesEntityReference{
			cnstypes.CnsKubernetesEntityReference{
				EntityType: string(cnstypes.CnsKubernetesEntityTypePVC),
				EntityName: "example-vanilla-block-pvc",
				Namespace:  "default",
				ClusterID:  "demo-cluster-id",
			},
		},
	}
	metadataList = append(metadataList, cnstypes.BaseCnsEntityMetadata(podmetadata))

	cnsVolumeMetadataUpdateSpec := cnstypes.CnsVolumeMetadataUpdateSpec{
		VolumeId: cnstypes.CnsVolumeId{Id: volumeId},
		Metadata: cnstypes.CnsVolumeMetadata{
			DynamicData:           vim25types.DynamicData{},
			ContainerCluster:      containerCluster,
			EntityMetadata:        metadataList,
			ContainerClusterArray: containerClusterArray,
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
		t.Logf("Successfully updated volume metadata")
	}

	t.Logf("Calling QueryVolume using queryFilter: %+v", pretty.Sprint(queryFilter))
	queryResult, err = cnsClient.QueryVolume(ctx, queryFilter)
	if err != nil {
		t.Errorf("Failed to query volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	t.Logf("Successfully Queried Volumes. queryResult: %+v", pretty.Sprint(queryResult))

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
	t.Logf("Successfully Queried all Volumes. queryResult: %+v", pretty.Sprint(queryResult))

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
	var resourcePool *object.ResourcePool
	if resporcePoolPath == "" {
		resourcePool, err = finder.DefaultResourcePool(ctx)
	} else {
		resourcePool, err = finder.ResourcePool(ctx, resporcePoolPath)
	}
	if err != nil {
		t.Errorf("Error occurred while getting DefaultResourcePool. err: %+v", err)
		t.Fatal(err)
	}
	task, err := defaultFolder.CreateVM(ctx, virtualMachineConfigSpec, resourcePool, nil)
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
	diskUUID := interface{}(attachTaskResult).(*cnstypes.CnsVolumeAttachResult).DiskUUID
	t.Logf("Volume attached sucessfully. Disk UUID: %s", diskUUID)

	// Re-Attach same volume to the same node and expect ResourceInUse fault
	t.Logf("Re-Attaching volume using the spec: %+v", cnsVolumeAttachSpec)
	attachTask, err = cnsClient.AttachVolume(ctx, cnsVolumeAttachSpecList)
	if err != nil {
		t.Errorf("Failed to attach volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	attachTaskInfo, err = GetTaskInfo(ctx, attachTask)
	if err != nil {
		t.Errorf("Failed to attach volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	attachTaskResult, err = GetTaskResult(ctx, attachTaskInfo)
	if err != nil {
		t.Errorf("Failed to attach volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	if attachTaskResult == nil {
		t.Fatalf("Empty attach task results")
		t.FailNow()
	}
	reAttachVolumeOperationRes := attachTaskResult.GetCnsVolumeOperationResult()
	if reAttachVolumeOperationRes.Fault != nil {
		t.Logf("reAttachVolumeOperationRes.Fault: %+v", pretty.Sprint(reAttachVolumeOperationRes.Fault))
		_, ok := reAttachVolumeOperationRes.Fault.Fault.(*vim25types.ResourceInUse)
		if !ok {
			t.Fatalf("Fault is not ResourceInUse")
		}
	} else {
		t.Fatalf("re-attach same volume should fail with ResourceInUse fault")
	}

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
		t.Errorf("Failed to delete volume. Error: %+v \n", err)
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
	t.Logf("Volume: %q deleted sucessfully", volumeId)

	if run_fileshare_tests == "true" && cnsClient.serviceClient.Version != ReleaseVSAN67u3 {
		// Test creating vSAN file-share Volume
		var cnsFileVolumeCreateSpecList []cnstypes.CnsVolumeCreateSpec
		vSANFileCreateSpec := &cnstypes.CnsVSANFileCreateSpec{
			SoftQuotaInMb: 5120,
			Permission: []vsanfstypes.VsanFileShareNetPermission{
				{
					Ips:         "*",
					Permissions: vsanfstypes.VsanFileShareAccessTypeREAD_WRITE,
					AllowRoot:   true,
				},
			},
		}

		cnsFileVolumeCreateSpec := cnstypes.CnsVolumeCreateSpec{
			Name:       "pvc-file-share-volume",
			VolumeType: string(cnstypes.CnsVolumeTypeFile),
			Datastores: dsList,
			Metadata: cnstypes.CnsVolumeMetadata{
				ContainerCluster:      containerCluster,
				ContainerClusterArray: containerClusterArray,
			},
			BackingObjectDetails: &cnstypes.CnsVsanFileShareBackingDetails{
				CnsFileBackingDetails: cnstypes.CnsFileBackingDetails{
					CnsBackingObjectDetails: cnstypes.CnsBackingObjectDetails{
						CapacityInMb: 5120,
					},
				},
			},
			CreateSpec: vSANFileCreateSpec,
		}
		cnsFileVolumeCreateSpecList = append(cnsFileVolumeCreateSpecList, cnsFileVolumeCreateSpec)
		t.Logf("Creating CNS file volume using the spec: %+v", cnsFileVolumeCreateSpec)
		createTask, err = cnsClient.CreateVolume(ctx, cnsFileVolumeCreateSpecList)
		if err != nil {
			t.Errorf("Failed to create vsan fileshare volume. Error: %+v \n", err)
			t.Fatal(err)
		}
		createTaskInfo, err = GetTaskInfo(ctx, createTask)
		if err != nil {
			t.Errorf("Failed to create Fileshare volume. Error: %+v \n", err)
			t.Fatal(err)
		}
		createTaskResult, err = GetTaskResult(ctx, createTaskInfo)
		if err != nil {
			t.Errorf("Failed to create Fileshare volume. Error: %+v \n", err)
			t.Fatal(err)
		}
		if createTaskResult == nil {
			t.Fatalf("Empty create task results")
			t.FailNow()
		}
		createVolumeOperationRes = createTaskResult.GetCnsVolumeOperationResult()
		if createVolumeOperationRes.Fault != nil {
			t.Fatalf("Failed to create Fileshare volume: fault=%+v", createVolumeOperationRes.Fault)
		}
		filevolumeId := createVolumeOperationRes.VolumeId.Id
		t.Logf("Fileshare volume created sucessfully. filevolumeId: %s", filevolumeId)

		// Test QueryVolume API
		volumeIDList = []cnstypes.CnsVolumeId{{Id: filevolumeId}}
		queryFilter.VolumeIds = volumeIDList
		t.Logf("Calling QueryVolume using queryFilter: %+v", queryFilter)
		queryResult, err = cnsClient.QueryVolume(ctx, queryFilter)
		if err != nil {
			t.Errorf("Failed to query volume. Error: %+v \n", err)
			t.Fatal(err)
		}
		t.Logf("Successfully Queried Volumes. queryResult: %+v", queryResult)
		fileBackingInfo := queryResult.Volumes[0].BackingObjectDetails.(*cnstypes.CnsVsanFileShareBackingDetails)
		t.Logf("File Share Name: %s with accessPoints: %+v", fileBackingInfo.Name, fileBackingInfo.AccessPoints)

		// Test Deleting vSAN file-share Volume
		var fileVolumeIDList []cnstypes.CnsVolumeId
		fileVolumeIDList = append(fileVolumeIDList, cnstypes.CnsVolumeId{Id: filevolumeId})
		t.Logf("Deleting fileshare volume: %+v", fileVolumeIDList)
		deleteTask, err = cnsClient.DeleteVolume(ctx, fileVolumeIDList, true)
		if err != nil {
			t.Errorf("Failed to delete fileshare volume. Error: %+v \n", err)
			t.Fatal(err)
		}
		deleteTaskInfo, err = GetTaskInfo(ctx, deleteTask)
		if err != nil {
			t.Errorf("Failed to delete fileshare volume. Error: %+v \n", err)
			t.Fatal(err)
		}
		deleteTaskResult, err = GetTaskResult(ctx, deleteTaskInfo)
		if err != nil {
			t.Errorf("Failed to delete fileshare volume. Error: %+v \n", err)
			t.Fatal(err)
		}
		if deleteTaskResult == nil {
			t.Fatalf("Empty delete task results")
			t.FailNow()
		}
		deleteVolumeOperationRes = deleteTaskResult.GetCnsVolumeOperationResult()
		if deleteVolumeOperationRes.Fault != nil {
			t.Fatalf("Failed to delete fileshare volume: fault=%+v", deleteVolumeOperationRes.Fault)
		}
		t.Logf("fileshare volume:%q deleted sucessfully", filevolumeId)
	}
	if backingDiskURLPath != "" && cnsClient.serviceClient.Version != ReleaseVSAN67u3 && cnsClient.serviceClient.Version != ReleaseVSAN70 {
		// Test CreateVolume API with existing VMDK
		var cnsVolumeCreateSpecList []cnstypes.CnsVolumeCreateSpec
		cnsVolumeCreateSpec := cnstypes.CnsVolumeCreateSpec{
			Name:       "pvc-901e87eb-c2bd-11e9-806f-005056a0c9a0",
			VolumeType: string(cnstypes.CnsVolumeTypeBlock),
			Metadata: cnstypes.CnsVolumeMetadata{
				ContainerCluster: containerCluster,
			},
			BackingObjectDetails: &cnstypes.CnsBlockBackingDetails{
				BackingDiskUrlPath: backingDiskURLPath,
			},
		}
		cnsVolumeCreateSpecList = append(cnsVolumeCreateSpecList, cnsVolumeCreateSpec)
		t.Logf("Creating volume using the spec: %+v", pretty.Sprint(cnsVolumeCreateSpec))
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
		var volumeID string
		if createVolumeOperationRes.Fault != nil {
			t.Logf("Failed to create volume: fault=%+v", createVolumeOperationRes.Fault)
			fault, ok := createVolumeOperationRes.Fault.Fault.(cnstypes.CnsAlreadyRegisteredFault)
			if !ok {
				t.Fatalf("Fault is not CnsAlreadyRegisteredFault")
			} else {
				t.Logf("Fault is CnsAlreadyRegisteredFault. backingDiskURLPath: %s is already registered", backingDiskURLPath)
				volumeID = fault.VolumeId.Id
			}
		} else {
			volumeID = createVolumeOperationRes.VolumeId.Id
			t.Logf("Volume created sucessfully with backingDiskURLPath: %s. volumeId: %s", backingDiskURLPath, volumeID)

			// Test re creating volume using BACKING_DISK_URL_PATH
			var reCreateCnsVolumeCreateSpecList []cnstypes.CnsVolumeCreateSpec
			reCreateCnsVolumeCreateSpec := cnstypes.CnsVolumeCreateSpec{
				Name:       "pvc-901e87eb-c2bd-11e9-806f-005056a0c9a0",
				VolumeType: string(cnstypes.CnsVolumeTypeBlock),
				Metadata: cnstypes.CnsVolumeMetadata{
					ContainerCluster: containerCluster,
				},
				BackingObjectDetails: &cnstypes.CnsBlockBackingDetails{
					BackingDiskUrlPath: backingDiskURLPath,
				},
			}

			reCreateCnsVolumeCreateSpecList = append(reCreateCnsVolumeCreateSpecList, reCreateCnsVolumeCreateSpec)
			t.Logf("Creating volume using the spec: %+v", pretty.Sprint(reCreateCnsVolumeCreateSpec))
			recreateTask, err := cnsClient.CreateVolume(ctx, reCreateCnsVolumeCreateSpecList)
			if err != nil {
				t.Errorf("Failed to create volume. Error: %+v \n", err)
				t.Fatal(err)
			}
			reCreateTaskInfo, err := GetTaskInfo(ctx, recreateTask)
			if err != nil {
				t.Errorf("Failed to create volume. Error: %+v \n", err)
				t.Fatal(err)
			}
			reCreateTaskResult, err := GetTaskResult(ctx, reCreateTaskInfo)
			if err != nil {
				t.Errorf("Failed to create volume. Error: %+v \n", err)
				t.Fatal(err)
			}
			if reCreateTaskResult == nil {
				t.Fatalf("Empty create task results")
				t.FailNow()
			}
			reCreateVolumeOperationRes := reCreateTaskResult.GetCnsVolumeOperationResult()
			t.Logf("reCreateVolumeOperationRes.: %+v", pretty.Sprint(reCreateVolumeOperationRes))
			if reCreateVolumeOperationRes.Fault != nil {
				t.Logf("Failed to create volume: fault=%+v", reCreateVolumeOperationRes.Fault)
				_, ok := reCreateVolumeOperationRes.Fault.Fault.(cnstypes.CnsAlreadyRegisteredFault)
				if !ok {
					t.Fatalf("Fault is not CnsAlreadyRegisteredFault")
				} else {
					t.Logf("Fault is CnsAlreadyRegisteredFault. backingDiskURLPath: %q is already registered", backingDiskURLPath)
				}
			}
		}

		// Test QueryVolume API
		var queryFilter cnstypes.CnsQueryFilter
		var volumeIDList []cnstypes.CnsVolumeId
		volumeIDList = append(volumeIDList, cnstypes.CnsVolumeId{Id: volumeID})
		queryFilter.VolumeIds = volumeIDList
		t.Logf("Calling QueryVolume using queryFilter: %+v", pretty.Sprint(queryFilter))
		queryResult, err := cnsClient.QueryVolume(ctx, queryFilter)
		if err != nil {
			t.Errorf("Failed to query volume. Error: %+v \n", err)
			t.Fatal(err)
		}
		t.Logf("Successfully Queried Volumes. queryResult: %+v", pretty.Sprint(queryResult))

		t.Logf("Deleting CNS volume created above using BACKING_DISK_URL_PATH: %s with volume: %+v", backingDiskURLPath, volumeIDList)
		deleteTask, err = cnsClient.DeleteVolume(ctx, volumeIDList, true)
		if err != nil {
			t.Errorf("Failed to delete volume. Error: %+v \n", err)
			t.Fatal(err)
		}
		deleteTaskInfo, err = GetTaskInfo(ctx, deleteTask)
		if err != nil {
			t.Errorf("Failed to delete volume. Error: %+v \n", err)
			t.Fatal(err)
		}
		deleteTaskResult, err = GetTaskResult(ctx, deleteTaskInfo)
		if err != nil {
			t.Errorf("Failed to delete volume. Error: %+v \n", err)
			t.Fatal(err)
		}
		if deleteTaskResult == nil {
			t.Fatalf("Empty delete task results")
			t.FailNow()
		}
		deleteVolumeOperationRes = deleteTaskResult.GetCnsVolumeOperationResult()
		if deleteVolumeOperationRes.Fault != nil {
			t.Fatalf("Failed to delete volume: fault=%+v", deleteVolumeOperationRes.Fault)
		}
		t.Logf("volume:%q deleted sucessfully", volumeID)
	}
}
