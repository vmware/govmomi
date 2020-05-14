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

package simulator

import (
	"context"
	"testing"

	"github.com/google/uuid"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/cns"
	cnstypes "github.com/vmware/govmomi/cns/types"
	"github.com/vmware/govmomi/simulator"
	vim25types "github.com/vmware/govmomi/vim25/types"
)

const (
	testLabel = "testLabel"
	testValue = "testValue"
)

func TestSimulator(t *testing.T) {
	ctx := context.Background()

	model := simulator.VPX()
	defer model.Remove()

	var err error

	if err = model.Create(); err != nil {
		t.Fatal(err)
	}

	s := model.Service.NewServer()
	defer s.Close()

	model.Service.RegisterSDK(New())

	c, err := govmomi.NewClient(ctx, s.URL, true)
	if err != nil {
		t.Fatal(err)
	}

	cnsClient, err := cns.NewClient(ctx, c.Client)
	if err != nil {
		t.Fatal(err)
	}
	// Query
	queryFilter := cnstypes.CnsQueryFilter{}
	queryResult, err := cnsClient.QueryVolume(ctx, queryFilter)
	if err != nil {
		t.Fatal(err)
	}
	existingNumDisks := len(queryResult.Volumes)

	// Get a simulator DS
	datastore := simulator.Map.Any("Datastore").(*simulator.Datastore)

	// Create
	var capacityInMb int64 = 1024
	createSpecList := []cnstypes.CnsVolumeCreateSpec{
		{
			Name:       "test",
			VolumeType: "TestVolumeType",
			Datastores: []vim25types.ManagedObjectReference{
				datastore.Self,
			},
			BackingObjectDetails: &cnstypes.CnsBackingObjectDetails{
				CapacityInMb: capacityInMb,
			},
			Profile: []vim25types.BaseVirtualMachineProfileSpec{
				&vim25types.VirtualMachineDefinedProfileSpec{
					ProfileId: uuid.New().String(),
				},
			},
		},
	}
	createTask, err := cnsClient.CreateVolume(ctx, createSpecList)
	if err != nil {
		t.Fatal(err)
	}

	createTaskInfo, err := cns.GetTaskInfo(ctx, createTask)
	if err != nil {
		t.Fatal(err)
	}

	createTaskResult, err := cns.GetTaskResult(ctx, createTaskInfo)
	if err != nil {
		t.Fatal(err)
	}
	if createTaskResult == nil {
		t.Fatalf("Empty create task results")
	}
	createVolumeOperationRes := createTaskResult.GetCnsVolumeOperationResult()
	if createVolumeOperationRes.Fault != nil {
		t.Fatalf("Failed to create volume: fault=%+v", createVolumeOperationRes.Fault)
	}
	volumeId := createVolumeOperationRes.VolumeId.Id

	// Extend
	var newCapacityInMb int64 = 2048
	extendSpecList := []cnstypes.CnsVolumeExtendSpec{
		{
			VolumeId:     createVolumeOperationRes.VolumeId,
			CapacityInMb: newCapacityInMb,
		},
	}
	extendTask, err := cnsClient.ExtendVolume(ctx, extendSpecList)
	if err != nil {
		t.Fatal(err)
	}

	extendTaskInfo, err := cns.GetTaskInfo(ctx, extendTask)
	if err != nil {
		t.Fatal(err)
	}

	extendTaskResult, err := cns.GetTaskResult(ctx, extendTaskInfo)
	if err != nil {
		t.Fatal(err)
	}
	if extendTaskResult == nil {
		t.Fatalf("Empty extend task results")
	}

	extendVolumeOperationRes := extendTaskResult.GetCnsVolumeOperationResult()
	if extendVolumeOperationRes.Fault != nil {
		t.Fatalf("Failed to extend: fault=%+v", extendVolumeOperationRes.Fault)
	}

	// Attach
	nodeVM := simulator.Map.Any("VirtualMachine").(*simulator.VirtualMachine)
	attachSpecList := []cnstypes.CnsVolumeAttachDetachSpec{
		{
			VolumeId: createVolumeOperationRes.VolumeId,
			Vm:       nodeVM.Self,
		},
	}
	attachTask, err := cnsClient.AttachVolume(ctx, attachSpecList)
	if err != nil {
		t.Fatal(err)
	}

	attachTaskInfo, err := cns.GetTaskInfo(ctx, attachTask)
	if err != nil {
		t.Fatal(err)
	}

	attachTaskResult, err := cns.GetTaskResult(ctx, attachTaskInfo)
	if err != nil {
		t.Fatal(err)
	}
	if attachTaskResult == nil {
		t.Fatalf("Empty attach task results")
	}

	attachVolumeOperationRes := attachTaskResult.GetCnsVolumeOperationResult()
	if attachVolumeOperationRes.Fault != nil {
		t.Fatalf("Failed to attach: fault=%+v", attachVolumeOperationRes.Fault)
	}

	// Detach
	detachVolumeList := []cnstypes.CnsVolumeAttachDetachSpec{
		{
			VolumeId: createVolumeOperationRes.VolumeId,
		},
	}
	detachTask, err := cnsClient.DetachVolume(ctx, detachVolumeList)

	detachTaskInfo, err := cns.GetTaskInfo(ctx, detachTask)
	if err != nil {
		t.Fatal(err)
	}

	detachTaskResult, err := cns.GetTaskResult(ctx, detachTaskInfo)
	if err != nil {
		t.Fatal(err)
	}
	if detachTaskResult == nil {
		t.Fatalf("Empty detach task results")
	}

	detachVolumeOperationRes := detachTaskResult.GetCnsVolumeOperationResult()
	if detachVolumeOperationRes.Fault != nil {
		t.Fatalf("Failed to detach volume: fault=%+v", detachVolumeOperationRes.Fault)
	}

	// Query
	queryFilter = cnstypes.CnsQueryFilter{}
	queryResult, err = cnsClient.QueryVolume(ctx, queryFilter)
	if err != nil {
		t.Fatal(err)
	}

	if len(queryResult.Volumes) != existingNumDisks+1 {
		t.Fatal("Number of volumes mismatches after creating a single volume")
	}

	// QueryVolumeInfo
	queryVolumeInfoTask, err := cnsClient.QueryVolumeInfo(ctx, []cnstypes.CnsVolumeId{{Id: createVolumeOperationRes.VolumeId.Id}})
	if err != nil {
		t.Fatal(err)
	}
	queryVolumeInfoTaskInfo, err := cns.GetTaskInfo(ctx, queryVolumeInfoTask)
	if err != nil {
		t.Fatal(err)
	}
	queryVolumeInfoTaskResult, err := cns.GetTaskResult(ctx, queryVolumeInfoTaskInfo)
	if err != nil {
		t.Fatal(err)
	}
	if queryVolumeInfoTaskResult == nil {
		t.Fatalf("Empty query VolumeInfo Task Result")
	}
	queryVolumeInfoOperationRes := queryVolumeInfoTaskResult.GetCnsVolumeOperationResult()
	if queryVolumeInfoOperationRes.Fault != nil {
		t.Fatalf("Failed to query volume detail using QueryVolumeInfo: fault=%+v", queryVolumeInfoOperationRes.Fault)
	}

	// QueryAll
	queryFilter = cnstypes.CnsQueryFilter{}
	querySelection := cnstypes.CnsQuerySelection{}
	queryResult, err = cnsClient.QueryAllVolume(ctx, queryFilter, querySelection)

	if len(queryResult.Volumes) != existingNumDisks+1 {
		t.Fatal("Number of volumes mismatches after creating a single volume")
	}

	// Update
	var metadataList []cnstypes.BaseCnsEntityMetadata
	newLabels := []vim25types.KeyValue{
		{
			Key:   testLabel,
			Value: testValue,
		},
	}
	metadata := &cnstypes.CnsKubernetesEntityMetadata{

		CnsEntityMetadata: cnstypes.CnsEntityMetadata{
			DynamicData: vim25types.DynamicData{},
			EntityName:  queryResult.Volumes[0].Name,
			Labels:      newLabels,
			Delete:      false,
		},
		EntityType: string(cnstypes.CnsKubernetesEntityTypePV),
		Namespace:  "",
	}
	metadataList = append(metadataList, cnstypes.BaseCnsEntityMetadata(metadata))
	updateSpecList := []cnstypes.CnsVolumeMetadataUpdateSpec{
		{
			DynamicData: vim25types.DynamicData{},
			VolumeId:    createVolumeOperationRes.VolumeId,
			Metadata: cnstypes.CnsVolumeMetadata{
				DynamicData:      vim25types.DynamicData{},
				ContainerCluster: queryResult.Volumes[0].Metadata.ContainerCluster,
				EntityMetadata:   metadataList,
			},
		},
	}
	updateTask, err := cnsClient.UpdateVolumeMetadata(ctx, updateSpecList)
	if err != nil {
		t.Fatal(err)
	}
	updateTaskInfo, err := cns.GetTaskInfo(ctx, updateTask)
	if err != nil {
		t.Fatal(err)
	}
	updateTaskResult, err := cns.GetTaskResult(ctx, updateTaskInfo)
	if err != nil {
		t.Fatal(err)
	}
	if updateTaskResult == nil {
		t.Fatalf("Empty create task results")
	}

	updateVolumeOperationRes := updateTaskResult.GetCnsVolumeOperationResult()
	if updateVolumeOperationRes.Fault != nil {
		t.Fatalf("Failed to create volume: fault=%+v", updateVolumeOperationRes.Fault)
	}

	// Delete
	deleteVolumeList := []cnstypes.CnsVolumeId{
		{
			Id: volumeId,
		},
	}
	deleteTask, err := cnsClient.DeleteVolume(ctx, deleteVolumeList, true)

	deleteTaskInfo, err := cns.GetTaskInfo(ctx, deleteTask)
	if err != nil {
		t.Fatal(err)
	}

	deleteTaskResult, err := cns.GetTaskResult(ctx, deleteTaskInfo)
	if err != nil {
		t.Fatal(err)
	}
	if deleteTaskResult == nil {
		t.Fatalf("Empty delete task results")
	}

	deleteVolumeOperationRes := deleteTaskResult.GetCnsVolumeOperationResult()
	if deleteVolumeOperationRes.Fault != nil {
		t.Fatalf("Failed to delete volume: fault=%+v", deleteVolumeOperationRes.Fault)
	}

	queryFilter = cnstypes.CnsQueryFilter{}
	queryResult, err = cnsClient.QueryVolume(ctx, queryFilter)
	if err != nil {
		t.Fatalf("Failed to query volume with QueryFilter: err=%+v", err)
	}
	if len(queryResult.Volumes) != existingNumDisks {
		t.Fatal("Number of volumes mismatches after deleting a single volume")
	}

}
