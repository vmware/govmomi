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
	"errors"

	cnstypes "github.com/vmware/govmomi/cns/types"
	"github.com/vmware/govmomi/object"
	vim25types "github.com/vmware/govmomi/vim25/types"
)

// GetTaskInfo gets the task info given a task
func GetTaskInfo(ctx context.Context, task *object.Task) (*vim25types.TaskInfo, error) {
	taskInfo, err := task.WaitForResult(ctx, nil)
	if err != nil {
		return nil, err
	}
	return taskInfo, nil
}

// GetQuerySnapshotsTaskResult gets the task result of QuerySnapshots given a task info
func GetQuerySnapshotsTaskResult(ctx context.Context, taskInfo *vim25types.TaskInfo) (*cnstypes.CnsSnapshotQueryResult, error) {
	if taskInfo == nil {
		return nil, errors.New("TaskInfo is empty")
	}
	if taskInfo.Result != nil {
		snapshotQueryResult := taskInfo.Result.(cnstypes.CnsSnapshotQueryResult)
		if &snapshotQueryResult == nil {
			return nil, errors.New("Cannot get SnapshotQueryResult")
		}
		return &snapshotQueryResult, nil
	}
	return nil, errors.New("TaskInfo result is empty")
}

// GetTaskResult gets the task result given a task info
func GetTaskResult(ctx context.Context, taskInfo *vim25types.TaskInfo) (cnstypes.BaseCnsVolumeOperationResult, error) {
	if taskInfo == nil {
		return nil, errors.New("TaskInfo is empty")
	}
	if taskInfo.Result != nil {
		volumeOperationBatchResult := taskInfo.Result.(cnstypes.CnsVolumeOperationBatchResult)
		if &volumeOperationBatchResult == nil ||
			volumeOperationBatchResult.VolumeResults == nil ||
			len(volumeOperationBatchResult.VolumeResults) == 0 {
			return nil, errors.New("Cannot get VolumeOperationResult")
		}
		return volumeOperationBatchResult.VolumeResults[0], nil
	}
	return nil, errors.New("TaskInfo result is empty")
}

// GetTaskResultArray gets the task result array for a specified task info
func GetTaskResultArray(ctx context.Context, taskInfo *vim25types.TaskInfo) ([]cnstypes.BaseCnsVolumeOperationResult, error) {
	if taskInfo == nil {
		return nil, errors.New("TaskInfo is empty")
	}
	if taskInfo.Result != nil {
		volumeOperationBatchResult := taskInfo.Result.(cnstypes.CnsVolumeOperationBatchResult)
		if &volumeOperationBatchResult == nil ||
			volumeOperationBatchResult.VolumeResults == nil ||
			len(volumeOperationBatchResult.VolumeResults) == 0 {
			return nil, errors.New("Cannot get VolumeOperationResult")
		}
		return volumeOperationBatchResult.VolumeResults, nil
	}
	return nil, errors.New("TaskInfo result is empty")
}

// dropUnknownCreateSpecElements helps drop newly added elements in the CnsVolumeCreateSpec, which are not known to the prior vSphere releases
func dropUnknownCreateSpecElements(c *Client, createSpecList []cnstypes.CnsVolumeCreateSpec) []cnstypes.CnsVolumeCreateSpec {
	updatedcreateSpecList := make([]cnstypes.CnsVolumeCreateSpec, 0, len(createSpecList))
	switch c.Version {
	case ReleaseVSAN67u3:
		// Dropping optional fields not known to vSAN 6.7U3
		for _, createSpec := range createSpecList {
			createSpec.Metadata.ContainerCluster.ClusterFlavor = ""
			createSpec.Metadata.ContainerCluster.ClusterDistribution = ""
			createSpec.Metadata.ContainerClusterArray = nil
			var updatedEntityMetadata []cnstypes.BaseCnsEntityMetadata
			for _, entityMetadata := range createSpec.Metadata.EntityMetadata {
				k8sEntityMetadata := interface{}(entityMetadata).(*cnstypes.CnsKubernetesEntityMetadata)
				k8sEntityMetadata.ClusterID = ""
				k8sEntityMetadata.ReferredEntity = nil
				updatedEntityMetadata = append(updatedEntityMetadata, cnstypes.BaseCnsEntityMetadata(k8sEntityMetadata))
			}
			createSpec.Metadata.EntityMetadata = updatedEntityMetadata
			_, ok := createSpec.BackingObjectDetails.(*cnstypes.CnsBlockBackingDetails)
			if ok {
				createSpec.BackingObjectDetails.(*cnstypes.CnsBlockBackingDetails).BackingDiskUrlPath = ""
			}
			updatedcreateSpecList = append(updatedcreateSpecList, createSpec)
		}
		createSpecList = updatedcreateSpecList
	case ReleaseVSAN70:
		// Dropping optional fields not known to vSAN 7.0
		for _, createSpec := range createSpecList {
			createSpec.Metadata.ContainerCluster.ClusterDistribution = ""
			var updatedContainerClusterArray []cnstypes.CnsContainerCluster
			for _, containerCluster := range createSpec.Metadata.ContainerClusterArray {
				containerCluster.ClusterDistribution = ""
				updatedContainerClusterArray = append(updatedContainerClusterArray, containerCluster)
			}
			createSpec.Metadata.ContainerClusterArray = updatedContainerClusterArray
			_, ok := createSpec.BackingObjectDetails.(*cnstypes.CnsBlockBackingDetails)
			if ok {
				createSpec.BackingObjectDetails.(*cnstypes.CnsBlockBackingDetails).BackingDiskUrlPath = ""
			}
			updatedcreateSpecList = append(updatedcreateSpecList, createSpec)
		}
		createSpecList = updatedcreateSpecList
	case ReleaseVSAN70u1:
		// Dropping optional fields not known to vSAN 7.0U1
		for _, createSpec := range createSpecList {
			createSpec.Metadata.ContainerCluster.ClusterDistribution = ""
			var updatedContainerClusterArray []cnstypes.CnsContainerCluster
			for _, containerCluster := range createSpec.Metadata.ContainerClusterArray {
				containerCluster.ClusterDistribution = ""
				updatedContainerClusterArray = append(updatedContainerClusterArray, containerCluster)
			}
			createSpec.Metadata.ContainerClusterArray = updatedContainerClusterArray
			updatedcreateSpecList = append(updatedcreateSpecList, createSpec)
		}
		createSpecList = updatedcreateSpecList
	}
	return createSpecList
}

// dropUnknownVolumeMetadataUpdateSpecElements helps drop newly added elements in the CnsVolumeMetadataUpdateSpec, which are not known to the prior vSphere releases
func dropUnknownVolumeMetadataUpdateSpecElements(c *Client, updateSpecList []cnstypes.CnsVolumeMetadataUpdateSpec) []cnstypes.CnsVolumeMetadataUpdateSpec {
	// Dropping optional fields not known to vSAN 6.7U3
	if c.Version == ReleaseVSAN67u3 {
		updatedUpdateSpecList := make([]cnstypes.CnsVolumeMetadataUpdateSpec, 0, len(updateSpecList))
		for _, updateSpec := range updateSpecList {
			updateSpec.Metadata.ContainerCluster.ClusterFlavor = ""
			updateSpec.Metadata.ContainerCluster.ClusterDistribution = ""
			var updatedEntityMetadata []cnstypes.BaseCnsEntityMetadata
			for _, entityMetadata := range updateSpec.Metadata.EntityMetadata {
				k8sEntityMetadata := interface{}(entityMetadata).(*cnstypes.CnsKubernetesEntityMetadata)
				k8sEntityMetadata.ClusterID = ""
				k8sEntityMetadata.ReferredEntity = nil
				updatedEntityMetadata = append(updatedEntityMetadata, cnstypes.BaseCnsEntityMetadata(k8sEntityMetadata))
			}
			updateSpec.Metadata.ContainerClusterArray = nil
			updateSpec.Metadata.EntityMetadata = updatedEntityMetadata
			updatedUpdateSpecList = append(updatedUpdateSpecList, updateSpec)
		}
		updateSpecList = updatedUpdateSpecList
	} else if c.Version == ReleaseVSAN70 || c.Version == ReleaseVSAN70u1 {
		updatedUpdateSpecList := make([]cnstypes.CnsVolumeMetadataUpdateSpec, 0, len(updateSpecList))
		for _, updateSpec := range updateSpecList {
			updateSpec.Metadata.ContainerCluster.ClusterDistribution = ""
			var updatedContainerClusterArray []cnstypes.CnsContainerCluster
			for _, containerCluster := range updateSpec.Metadata.ContainerClusterArray {
				containerCluster.ClusterDistribution = ""
				updatedContainerClusterArray = append(updatedContainerClusterArray, containerCluster)
			}
			updateSpec.Metadata.ContainerClusterArray = updatedContainerClusterArray
			updatedUpdateSpecList = append(updatedUpdateSpecList, updateSpec)
		}
		updateSpecList = updatedUpdateSpecList
	}
	return updateSpecList
}
