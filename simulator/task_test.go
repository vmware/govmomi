// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"testing"

	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type addWaterTask struct {
	*mo.Folder

	fault types.BaseMethodFault
}

func (a *addWaterTask) Run(task *Task) (types.AnyType, types.BaseMethodFault) {
	return nil, a.fault
}

func TestNewTask(t *testing.T) {
	ctx := NewContext()
	f := &mo.Folder{}
	ctx.Map.NewEntity(f)

	add := &addWaterTask{f, nil}
	task := NewTask(add)
	info := &task.Info

	if info.Name != "AddWater_Task" {
		t.Errorf("name=%s", info.Name)
	}

	if info.DescriptionId != "Folder.addWater" {
		t.Errorf("descriptionId=%s", info.DescriptionId)
	}

	task.RunBlocking(ctx)

	if info.State != types.TaskInfoStateSuccess {
		t.Fail()
	}

	add.fault = &types.ManagedObjectNotFound{}

	task.Run(ctx)
	task.Wait()

	if info.State != types.TaskInfoStateError {
		t.Fail()
	}

	if info.Key == "" {
		t.Error("empty info.key")
	}

	if info.Task.Type == "" {
		t.Error("empty info.task.type")
	}
}
