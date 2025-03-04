// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"container/list"
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

	history *history
}

func (m *TaskManager) init(r *Registry) {
	if len(m.Description.MethodInfo) == 0 {
		if r.IsVPX() {
			m.Description = vpx.Description
		} else {
			m.Description = esx.Description
		}
	}

	if m.MaxCollector == 0 {
		// In real VC this default can be changed via OptionManager "task.maxCollectors"
		m.MaxCollector = maxCollectors
	}

	m.history = newHistory()

	r.AddHandler(m)
}

func recentTask(recent []types.ManagedObjectReference, ref types.ManagedObjectReference) []types.PropertyChange {
	// TODO: tasks completed > 10m ago should be removed
	recent = append(recent, ref)
	if len(recent) > recentTaskMax {
		recent = recent[1:]
	}
	return []types.PropertyChange{{Name: "recentTask", Val: recent}}
}

func (m *TaskManager) PutObject(ctx *Context, obj mo.Reference) {
	task, ok := obj.(*Task)
	if !ok {
		return
	}

	// propagate new Tasks to:
	// - TaskManager.RecentTask
	// - TaskHistoryCollector instances, if Filter matches
	// - $MO.RecentTask
	m.Lock()
	ctx.Update(m, recentTask(m.RecentTask, task.Self))

	pushHistory(m.history.page, task)

	for _, hc := range m.history.collectors {
		c := hc.(*TaskHistoryCollector)
		ctx.WithLock(c, func() {
			if c.taskMatches(ctx, &task.Info) {
				pushHistory(c.page, task)
				ctx.Update(c, []types.PropertyChange{{Name: "latestPage", Val: c.GetLatestPage()}})
			}
		})
	}
	m.Unlock()

	entity := ctx.Map.Get(*task.Info.Entity)
	if e, ok := entity.(mo.Entity); ok {
		ctx.Update(entity, recentTask(e.Entity().RecentTask, task.Self))
	}
}

func taskStateChanged(pc []types.PropertyChange) bool {
	for i := range pc {
		if pc[i].Name == "info.state" {
			return true
		}
	}
	return false
}

func (m *TaskManager) UpdateObject(ctx *Context, obj mo.Reference, pc []types.PropertyChange) {
	task, ok := obj.(*mo.Task)
	if !ok {
		return
	}

	if !taskStateChanged(pc) {
		// real vCenter only updates latestPage when Tasks are created (see PutObject above) and
		// if Task.Info.State changes.
		// Changes to Task.Info.{Description,Progress} does not update lastestPage.
		return
	}

	m.Lock()
	for _, hc := range m.history.collectors {
		c := hc.(*TaskHistoryCollector)
		ctx.WithLock(c, func() {
			if c.hasTask(ctx, &task.Info) {
				ctx.Update(c, []types.PropertyChange{{Name: "latestPage", Val: c.GetLatestPage()}})
			}
		})
	}
	m.Unlock()
}

func (*TaskManager) RemoveObject(*Context, types.ManagedObjectReference) {}

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

type TaskHistoryCollector struct {
	mo.TaskHistoryCollector

	*HistoryCollector
}

func (m *TaskManager) createCollector(ctx *Context, req *types.CreateCollectorForTasks) (*TaskHistoryCollector, *soap.Fault) {
	if len(m.history.collectors) >= int(m.MaxCollector) {
		return nil, Fault("Too many task collectors to create", new(types.InvalidState))
	}

	collector := &TaskHistoryCollector{
		HistoryCollector: newHistoryCollector(ctx, m.history, defaultPageSize),
	}
	collector.Filter = req.Filter

	return collector, nil
}

// taskFilterChildren returns true if a child of self is the task entity.
func taskFilterChildren(ctx *Context, root types.ManagedObjectReference, task *types.TaskInfo) bool {
	seen := false

	var match func(types.ManagedObjectReference)

	match = func(child types.ManagedObjectReference) {
		if child == *task.Entity {
			seen = true
			return
		}

		walk(ctx.Map.Get(child), match)
	}

	walk(ctx.Map.Get(root), match)

	return seen
}

// entityMatches returns true if the spec Entity filter matches the task entity.
func (c *TaskHistoryCollector) entityMatches(ctx *Context, task *types.TaskInfo, spec types.TaskFilterSpec) bool {
	e := spec.Entity
	if e == nil {
		return true
	}

	isSelf := *task.Entity == e.Entity

	switch e.Recursion {
	case types.TaskFilterSpecRecursionOptionSelf:
		return isSelf
	case types.TaskFilterSpecRecursionOptionChildren:
		return taskFilterChildren(ctx, e.Entity, task)
	case types.TaskFilterSpecRecursionOptionAll:
		return isSelf || taskFilterChildren(ctx, e.Entity, task)
	}

	return false
}

func (c *TaskHistoryCollector) stateMatches(_ *Context, task *types.TaskInfo, spec types.TaskFilterSpec) bool {
	if len(spec.State) == 0 {
		return true
	}

	for _, state := range spec.State {
		if task.State == state {
			return true

		}
	}

	return false
}

func (c *TaskHistoryCollector) timeMatches(_ *Context, task *types.TaskInfo, spec types.TaskFilterSpec) bool {
	if spec.Time == nil {
		return true
	}

	created := task.QueueTime

	if begin := spec.Time.BeginTime; begin != nil {
		if created.Before(*begin) {
			return false
		}
	}

	if end := spec.Time.EndTime; end != nil {
		if created.After(*end) {
			return false
		}
	}

	return true
}

// taskMatches returns true one of the filters matches the task.
func (c *TaskHistoryCollector) taskMatches(ctx *Context, task *types.TaskInfo) bool {
	spec := c.Filter.(types.TaskFilterSpec)

	matchers := []func(*Context, *types.TaskInfo, types.TaskFilterSpec) bool{
		c.stateMatches,
		c.timeMatches,
		c.entityMatches,
		// TODO: spec.UserName, etc
	}

	for _, match := range matchers {
		if !match(ctx, task, spec) {
			return false
		}
	}

	return true
}

func (c *TaskHistoryCollector) hasTask(_ *Context, task *types.TaskInfo) bool {
	for e := c.page.Front(); e != nil; e = e.Next() {
		if e.Value.(*Task).Info.Key == task.Key {
			return true
		}
	}
	return false
}

// fillPage copies the manager's latest tasks into the collector's page with Filter applied.
func (m *TaskManager) fillPage(ctx *Context, c *TaskHistoryCollector) {
	m.history.Lock()
	defer m.history.Unlock()

	for e := m.history.page.Front(); e != nil; e = e.Next() {
		task := e.Value.(*Task)

		if c.taskMatches(ctx, &task.Info) {
			pushHistory(c.page, task)
		}
	}
}

func (m *TaskManager) CreateCollectorForTasks(ctx *Context, req *types.CreateCollectorForTasks) soap.HasFault {
	body := new(methods.CreateCollectorForTasksBody)

	if ctx.Map.IsESX() {
		body.Fault_ = Fault("", new(types.NotSupported))
		return body
	}

	collector, err := m.createCollector(ctx, req)
	if err != nil {
		body.Fault_ = err
		return body
	}

	collector.fill = func(x *Context) { m.fillPage(x, collector) }
	collector.fill(ctx)

	body.Res = &types.CreateCollectorForTasksResponse{
		Returnval: m.history.add(ctx, collector),
	}

	return body
}

func (c *TaskHistoryCollector) ReadNextTasks(ctx *Context, req *types.ReadNextTasks) soap.HasFault {
	body := new(methods.ReadNextTasksBody)
	if req.MaxCount <= 0 {
		body.Fault_ = Fault("", errInvalidArgMaxCount)
		return body
	}
	body.Res = new(types.ReadNextTasksResponse)

	c.next(req.MaxCount, func(e *list.Element) {
		body.Res.Returnval = append(body.Res.Returnval, e.Value.(*Task).Info)
	})

	return body
}

func (c *TaskHistoryCollector) ReadPreviousTasks(ctx *Context, req *types.ReadPreviousTasks) soap.HasFault {
	body := new(methods.ReadPreviousTasksBody)
	if req.MaxCount <= 0 {
		body.Fault_ = Fault("", errInvalidArgMaxCount)
		return body
	}
	body.Res = new(types.ReadPreviousTasksResponse)

	c.prev(req.MaxCount, func(e *list.Element) {
		body.Res.Returnval = append(body.Res.Returnval, e.Value.(*Task).Info)
	})

	return body
}

func (c *TaskHistoryCollector) GetLatestPage() []types.TaskInfo {
	var latestPage []types.TaskInfo

	e := c.page.Back()
	for i := 0; i < c.size; i++ {
		if e == nil {
			break
		}
		latestPage = append(latestPage, e.Value.(*Task).Info)
		e = e.Prev()
	}

	return latestPage
}

func (c *TaskHistoryCollector) Get() mo.Reference {
	clone := *c

	clone.LatestPage = clone.GetLatestPage()

	return &clone
}
