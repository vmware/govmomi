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
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

const vTaskSuffix = "_Task" // vmomi suffix
const sTaskSuffix = "Task"  // simulator suffix (avoiding golint warning)

// TaskDelay applies to all tasks.
// Names for DelayConfig.MethodDelay will differ for task and api delays. API
// level names often look like PowerOff_Task, whereas the task name is simply
// PowerOff.
var TaskDelay = DelayConfig{}

type Task struct {
	mo.Task

	ctx     *Context
	Execute func(*Task) (types.AnyType, types.BaseMethodFault)
}

func NewTask(runner TaskRunner) *Task {
	ref := runner.Reference()
	name := reflect.TypeOf(runner).Elem().Name()
	name = strings.Replace(name, "VM", "Vm", 1) // "VM" for the type to make go-lint happy, but "Vm" for the vmodl ID
	return CreateTask(ref, name, runner.Run)
}

func CreateTask(e mo.Reference, name string, run func(*Task) (types.AnyType, types.BaseMethodFault)) *Task {
	ref := e.Reference()
	id := name

	if strings.HasSuffix(id, sTaskSuffix) {
		id = id[:len(id)-len(sTaskSuffix)]
		name = id + vTaskSuffix
	}

	task := &Task{
		Execute: run,
	}

	task.Self = Map.newReference(task)
	task.Info.Key = task.Self.Value
	task.Info.Task = task.Self
	task.Info.Name = ucFirst(name)
	task.Info.DescriptionId = fmt.Sprintf("%s.%s", ref.Type, id)
	task.Info.Entity = &ref
	task.Info.EntityName = ref.Value
	task.Info.Reason = &types.TaskReasonUser{UserName: "vcsim"} // TODO: Context.Session.User
	task.Info.QueueTime = time.Now()
	task.Info.State = types.TaskInfoStateQueued

	Map.Put(task)

	return task
}

type TaskRunner interface {
	mo.Reference

	Run(*Task) (types.AnyType, types.BaseMethodFault)
}

// taskReference is a helper struct so we can call AcquireLock in Run()
type taskReference struct {
	Self types.ManagedObjectReference
}

func (tr *taskReference) Reference() types.ManagedObjectReference {
	return tr.Self
}

func (t *Task) Run(ctx *Context) types.ManagedObjectReference {
	t.ctx = ctx
	// alias the global Map to reduce data races in tests that reset the
	// global Map variable.
	vimMap := Map

	vimMap.AtomicUpdate(t.ctx, t, []types.PropertyChange{
		{Name: "info.startTime", Val: time.Now()},
		{Name: "info.state", Val: types.TaskInfoStateRunning},
	})

	tr := &taskReference{
		Self: *t.Info.Entity,
	}

	// in most cases, the caller already holds this lock, and we would like
	// the lock to be held across the "hand off" to the async goroutine.
	// however, with a TaskDelay, PropertyCollector (for example) cannot read
	// any object properties while the lock is held.
	handoff := true
	if v, ok := TaskDelay.MethodDelay["LockHandoff"]; ok {
		handoff = v != 0
	}
	var unlock func()
	if handoff {
		unlock = vimMap.AcquireLock(ctx, tr)
	}
	go func() {
		TaskDelay.delay(t.Info.Name)
		if !handoff {
			unlock = vimMap.AcquireLock(ctx, tr)
		}
		res, err := t.Execute(t)
		unlock()

		state := types.TaskInfoStateSuccess
		var fault interface{}
		if err != nil {
			state = types.TaskInfoStateError
			fault = types.LocalizedMethodFault{
				Fault:            err,
				LocalizedMessage: fmt.Sprintf("%T", err),
			}
		}

		vimMap.AtomicUpdate(t.ctx, t, []types.PropertyChange{
			{Name: "info.completeTime", Val: time.Now()},
			{Name: "info.state", Val: state},
			{Name: "info.result", Val: res},
			{Name: "info.error", Val: fault},
		})
	}()

	return t.Self
}

// RunBlocking() should only be used when an async simulator task needs to wait
// on another async simulator task.
// It polls for task completion to avoid the need to set up a PropertyCollector.
func (t *Task) RunBlocking(ctx *Context) {
	_ = t.Run(ctx)
	t.Wait()
}

// Wait blocks until the task is complete.
func (t *Task) Wait() {
	// we do NOT want to share our lock with the tasks's context, because
	// the goroutine that executes the task will use ctx to update the
	// state (among other things).
	isolatedLockingContext := &Context{}

	isRunning := func() bool {
		var running bool
		Map.WithLock(isolatedLockingContext, t, func() {
			switch t.Info.State {
			case types.TaskInfoStateSuccess, types.TaskInfoStateError:
				running = false
			default:
				running = true
			}
		})
		return running
	}

	for isRunning() {
		time.Sleep(10 * time.Millisecond)
	}
}

func (t *Task) isDone() bool {
	return t.Info.State == types.TaskInfoStateError || t.Info.State == types.TaskInfoStateSuccess
}

func (t *Task) SetTaskState(ctx *Context, req *types.SetTaskState) soap.HasFault {
	body := new(methods.SetTaskStateBody)

	if t.isDone() {
		body.Fault_ = Fault("", new(types.InvalidState))
		return body
	}

	changes := []types.PropertyChange{
		{Name: "info.state", Val: req.State},
	}

	switch req.State {
	case types.TaskInfoStateRunning:
		changes = append(changes, types.PropertyChange{Name: "info.startTime", Val: time.Now()})
	case types.TaskInfoStateError, types.TaskInfoStateSuccess:
		changes = append(changes, types.PropertyChange{Name: "info.completeTime", Val: time.Now()})

		if req.Fault != nil {
			changes = append(changes, types.PropertyChange{Name: "info.error", Val: req.Fault})
		}
		if req.Result != nil {
			changes = append(changes, types.PropertyChange{Name: "info.result", Val: req.Result})
		}
	}

	ctx.Map.Update(t, changes)

	body.Res = new(types.SetTaskStateResponse)
	return body
}

func (t *Task) SetTaskDescription(ctx *Context, req *types.SetTaskDescription) soap.HasFault {
	body := new(methods.SetTaskDescriptionBody)

	if t.isDone() {
		body.Fault_ = Fault("", new(types.InvalidState))
		return body
	}

	ctx.Map.Update(t, []types.PropertyChange{{Name: "info.description", Val: req.Description}})

	body.Res = new(types.SetTaskDescriptionResponse)
	return body
}

func (t *Task) UpdateProgress(ctx *Context, req *types.UpdateProgress) soap.HasFault {
	body := new(methods.UpdateProgressBody)

	if t.Info.State != types.TaskInfoStateRunning {
		body.Fault_ = Fault("", new(types.InvalidState))
		return body
	}

	ctx.Map.Update(t, []types.PropertyChange{{Name: "info.progress", Val: req.PercentDone}})

	body.Res = new(types.UpdateProgressResponse)
	return body
}

func (t *Task) CancelTask(ctx *Context, req *types.CancelTask) soap.HasFault {
	body := new(methods.CancelTaskBody)

	if t.isDone() {
		body.Fault_ = Fault("", new(types.InvalidState))
		return body
	}

	changes := []types.PropertyChange{
		{Name: "info.canceled", Val: true},
		{Name: "info.completeTime", Val: time.Now()},
		{Name: "info.state", Val: types.TaskInfoStateError},
		{Name: "info.error", Val: &types.LocalizedMethodFault{
			Fault:            &types.RequestCanceled{},
			LocalizedMessage: "The task was canceled by a user",
		}},
	}

	ctx.Map.Update(t, changes)

	body.Res = new(types.CancelTaskResponse)
	return body
}
