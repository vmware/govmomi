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

// version and namespace constants for task client
const (
	// ClientVersion "vSAN 6.7U3" is the stable version for vSphere 6.7u3 release
	taskClientVersion   = "vSAN 6.7U3"
	taskClientNamespace = "urn:vsan"
)

// GetTaskInfo gets the task info given a task
func GetTaskInfo(ctx context.Context, task *object.Task) (*vim25types.TaskInfo, error) {
	task.Client().Version = taskClientVersion
	task.Client().Namespace = taskClientNamespace
	taskInfo, err := task.WaitForResult(ctx, nil)
	if err != nil {
		return nil, err
	}
	return taskInfo, nil
}

// GetTaskResult gets the task result given a task info
func GetTaskResult(ctx context.Context, taskInfo *vim25types.TaskInfo) (cnstypes.BaseCnsVolumeOperationResult, error) {
	if taskInfo == nil {
		return nil, errors.New("TaskInfo is empty")
	}
	volumeOperationBatchResult := taskInfo.Result.(cnstypes.CnsVolumeOperationBatchResult)
	if &volumeOperationBatchResult == nil ||
		volumeOperationBatchResult.VolumeResults == nil ||
		len(volumeOperationBatchResult.VolumeResults) == 0 {
		return nil, errors.New("Cannot get VolumeOperationResult")
	}
	return volumeOperationBatchResult.VolumeResults[0], nil
}
