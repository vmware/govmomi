/*
Copyright (c) 2017-2024 VMware, Inc. All Rights Reserved.

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
	"sync"
	"time"

	"github.com/vmware/govmomi/simulator/esx"
	"github.com/vmware/govmomi/simulator/vpx"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

var recentTaskMax = 200 // the VC limit

type TaskManager struct {
	mo.TaskManager
	sync.Mutex
}

func (m *TaskManager) init(r *Registry) {
	if len(m.Description.MethodInfo) == 0 {
		if r.IsVPX() {
			m.Description = vpx.Description
		} else {
			m.Description = esx.Description
		}
	}
	r.AddHandler(m)
}

func (m *TaskManager) PutObject(obj mo.Reference) {
	ref := obj.Reference()
	if ref.Type != "Task" {
		return
	}

	m.Lock()
	recent := append(m.RecentTask, ref)
	if len(recent) > recentTaskMax {
		recent = recent[1:]
	}

	Map.Update(m, []types.PropertyChange{{Name: "recentTask", Val: recent}})
	m.Unlock()
}

func (*TaskManager) RemoveObject(*Context, types.ManagedObjectReference) {}

func (*TaskManager) UpdateObject(mo.Reference, []types.PropertyChange) {}

func validTaskID(ctx *Context, taskID string) bool {
	m := ctx.Map.ExtensionManager()

	for _, x := range m.ExtensionList {
		for _, task := range x.TaskList {
			if task.TaskID == taskID {
				return true
			}
		}
	}

	return false
}

func (m *TaskManager) CreateTask(ctx *Context, req *types.CreateTask) soap.HasFault {
	body := &methods.CreateTaskBody{}

	if !validTaskID(ctx, req.TaskTypeId) {
		body.Fault_ = Fault("", &types.InvalidArgument{
			InvalidProperty: "taskType",
		})
		return body
	}

	task := &Task{}

	task.Self = ctx.Map.newReference(task)
	task.Info.Key = task.Self.Value
	task.Info.Task = task.Self
	task.Info.DescriptionId = req.TaskTypeId
	task.Info.Cancelable = req.Cancelable
	task.Info.Entity = &req.Obj
	task.Info.EntityName = req.Obj.Value
	task.Info.Reason = &types.TaskReasonUser{UserName: ctx.Session.UserName}
	task.Info.QueueTime = time.Now()
	task.Info.State = types.TaskInfoStateQueued

	body.Res = &types.CreateTaskResponse{Returnval: task.Info}

	go ctx.Map.Put(task)

	return body
}
