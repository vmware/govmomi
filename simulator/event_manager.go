/*
Copyright (c) 2018-2024 VMware, Inc. All Rights Reserved.

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
	"bytes"
	"container/list"
	"log"
	"reflect"
	"text/template"
	"time"

	"github.com/vmware/govmomi/simulator/esx"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

var (
	logEvents = false
)

type EventManager struct {
	mo.EventManager

	history   *history
	key       int32
	templates map[string]*template.Template
}

func (m *EventManager) init(r *Registry) {
	if len(m.Description.EventInfo) == 0 {
		m.Description.EventInfo = esx.EventInfo
	}
	if m.MaxCollector == 0 {
		// In real VC this default can be changed via OptionManager "event.maxCollectors"
		m.MaxCollector = maxCollectors
	}

	m.history = newHistory()
	m.templates = make(map[string]*template.Template)
}

func (m *EventManager) createCollector(ctx *Context, req *types.CreateCollectorForEvents) (*EventHistoryCollector, *soap.Fault) {
	size, err := validatePageSize(req.Filter.MaxCount)
	if err != nil {
		return nil, err
	}

	if len(m.history.collectors) >= int(m.MaxCollector) {
		return nil, Fault("Too many event collectors to create", new(types.InvalidState))
	}

	collector := &EventHistoryCollector{
		HistoryCollector: newHistoryCollector(ctx, m.history, size),
	}
	collector.Filter = req.Filter

	return collector, nil
}

func (m *EventManager) CreateCollectorForEvents(ctx *Context, req *types.CreateCollectorForEvents) soap.HasFault {
	body := new(methods.CreateCollectorForEventsBody)
	collector, err := m.createCollector(ctx, req)
	if err != nil {
		body.Fault_ = err
		return body
	}

	collector.fill = func(x *Context) { m.fillPage(x, collector) }
	collector.fill(ctx)

	body.Res = &types.CreateCollectorForEventsResponse{
		Returnval: m.history.add(ctx, collector),
	}

	return body
}

func (m *EventManager) QueryEvents(ctx *Context, req *types.QueryEvents) soap.HasFault {
	if ctx.Map.IsESX() {
		return &methods.QueryEventsBody{
			Fault_: Fault("", new(types.NotImplemented)),
		}
	}

	body := new(methods.QueryEventsBody)
	collector, err := m.createCollector(ctx, &types.CreateCollectorForEvents{Filter: req.Filter})
	if err != nil {
		body.Fault_ = err
		return body
	}

	m.fillPage(ctx, collector)

	body.Res = &types.QueryEventsResponse{
		Returnval: collector.GetLatestPage(),
	}

	return body
}

// formatMessage applies the EventDescriptionEventDetail.FullFormat template to the given event's FullFormattedMessage field.
func (m *EventManager) formatMessage(event types.BaseEvent) {
	id := reflect.ValueOf(event).Elem().Type().Name()
	e := event.GetEvent()

	t, ok := m.templates[id]
	if !ok {
		for _, info := range m.Description.EventInfo {
			if info.Key == id {
				t = template.Must(template.New(id).Parse(info.FullFormat))
				m.templates[id] = t
				break
			}
		}
	}

	if t != nil {
		var buf bytes.Buffer
		if err := t.Execute(&buf, event); err != nil {
			log.Print(err)
		}
		e.FullFormattedMessage = buf.String()
	}

	if logEvents {
		log.Printf("[%s] %s", id, e.FullFormattedMessage)
	}
}

func (m *EventManager) PostEvent(ctx *Context, req *types.PostEvent) soap.HasFault {
	m.key++
	event := req.EventToPost.GetEvent()
	event.Key = m.key
	event.ChainId = event.Key
	event.CreatedTime = time.Now()
	event.UserName = ctx.Session.UserName

	m.formatMessage(req.EventToPost)

	pushHistory(m.history.page, req.EventToPost)

	for _, hc := range m.history.collectors {
		c := hc.(*EventHistoryCollector)
		ctx.WithLock(c, func() {
			if c.eventMatches(ctx, req.EventToPost) {
				pushHistory(c.page, req.EventToPost)
				ctx.Update(c, []types.PropertyChange{{Name: "latestPage", Val: c.GetLatestPage()}})
			}
		})
	}

	if m := ctx.Map.AlarmManager(); m != nil {
		ctx.WithLock(m, func() { m.postEvent(ctx, req.EventToPost) })
	}

	return &methods.PostEventBody{
		Res: new(types.PostEventResponse),
	}
}

type EventHistoryCollector struct {
	mo.EventHistoryCollector

	*HistoryCollector
}

// doEntityEventArgument calls f for each entity argument in the event.
// If f returns true, the iteration stops.
func doEntityEventArgument(event types.BaseEvent, f func(types.ManagedObjectReference, *types.EntityEventArgument) bool) bool {
	e := event.GetEvent()

	if arg := e.Vm; arg != nil {
		if f(arg.Vm, &arg.EntityEventArgument) {
			return true
		}
	}

	if arg := e.Host; arg != nil {
		if f(arg.Host, &arg.EntityEventArgument) {
			return true
		}
	}

	if arg := e.ComputeResource; arg != nil {
		if f(arg.ComputeResource, &arg.EntityEventArgument) {
			return true
		}
	}

	if arg := e.Ds; arg != nil {
		if f(arg.Datastore, &arg.EntityEventArgument) {
			return true
		}
	}

	if arg := e.Net; arg != nil {
		if f(arg.Network, &arg.EntityEventArgument) {
			return true
		}
	}

	if arg := e.Dvs; arg != nil {
		if f(arg.Dvs, &arg.EntityEventArgument) {
			return true
		}
	}

	if arg := e.Datacenter; arg != nil {
		if f(arg.Datacenter, &arg.EntityEventArgument) {
			return true
		}
	}

	return false
}

// eventFilterSelf returns true if self is one of the entity arguments in the event.
func eventFilterSelf(event types.BaseEvent, self types.ManagedObjectReference) bool {
	if x, ok := event.(*types.EventEx); ok {
		if self.Type == x.ObjectType && self.Value == x.ObjectId {
			return true
		}
	}
	return doEntityEventArgument(event, func(ref types.ManagedObjectReference, _ *types.EntityEventArgument) bool {
		return self == ref
	})
}

// eventFilterChildren returns true if a child of self is one of the entity arguments in the event.
func eventFilterChildren(ctx *Context, root types.ManagedObjectReference, event types.BaseEvent) bool {
	return doEntityEventArgument(event, func(ref types.ManagedObjectReference, _ *types.EntityEventArgument) bool {
		seen := false

		var match func(types.ManagedObjectReference)

		match = func(child types.ManagedObjectReference) {
			if child == ref {
				seen = true
				return
			}

			walk(ctx.Map.Get(child), match)
		}

		walk(ctx.Map.Get(root), match)

		return seen
	})
}

// entityMatches returns true if the spec Entity filter matches the event.
func (c *EventHistoryCollector) entityMatches(ctx *Context, event types.BaseEvent, spec *types.EventFilterSpec) bool {
	e := spec.Entity
	if e == nil {
		return true
	}

	isRootFolder := c.root == e.Entity

	switch e.Recursion {
	case types.EventFilterSpecRecursionOptionSelf:
		return isRootFolder || eventFilterSelf(event, e.Entity)
	case types.EventFilterSpecRecursionOptionChildren:
		return eventFilterChildren(ctx, e.Entity, event)
	case types.EventFilterSpecRecursionOptionAll:
		if isRootFolder || eventFilterSelf(event, e.Entity) {
			return true
		}
		return eventFilterChildren(ctx, e.Entity, event)
	}

	return false
}

// chainMatches returns true if spec.EventChainId matches the event.
func (c *EventHistoryCollector) chainMatches(_ *Context, event types.BaseEvent, spec *types.EventFilterSpec) bool {
	e := event.GetEvent()
	if spec.EventChainId != 0 {
		if e.ChainId != spec.EventChainId {
			return false
		}
	}
	return true
}

// typeMatches returns true if one of the spec EventTypeId types matches the event.
func (c *EventHistoryCollector) typeMatches(_ *Context, event types.BaseEvent, spec *types.EventFilterSpec) bool {
	if len(spec.EventTypeId) == 0 {
		return true
	}

	matches := func(name string) bool {
		for _, id := range spec.EventTypeId {
			if id == name {
				return true
			}
		}
		return false
	}

	if x, ok := event.(*types.EventEx); ok {
		return matches(x.EventTypeId)
	}

	kind := reflect.ValueOf(event).Elem().Type()

	if matches(kind.Name()) {
		return true // concrete type
	}

	field, ok := kind.FieldByNameFunc(matches)
	if ok {
		return field.Anonymous // base type (embedded field)
	}
	return false
}

func (c *EventHistoryCollector) timeMatches(_ *Context, event types.BaseEvent, spec *types.EventFilterSpec) bool {
	if spec.Time == nil {
		return true
	}

	created := event.GetEvent().CreatedTime

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

// eventMatches returns true one of the filters matches the event.
func (c *EventHistoryCollector) eventMatches(ctx *Context, event types.BaseEvent) bool {
	spec := c.Filter.(types.EventFilterSpec)

	matchers := []func(*Context, types.BaseEvent, *types.EventFilterSpec) bool{
		c.chainMatches,
		c.typeMatches,
		c.timeMatches,
		c.entityMatches,
		// TODO: spec.UserName, etc
	}

	for _, match := range matchers {
		if !match(ctx, event, &spec) {
			return false
		}
	}

	return true
}

// fillPage copies the manager's latest events into the collector's page with Filter applied.
func (m *EventManager) fillPage(ctx *Context, c *EventHistoryCollector) {
	m.history.Lock()
	defer m.history.Unlock()

	for e := m.history.page.Front(); e != nil; e = e.Next() {
		event := e.Value.(types.BaseEvent)
		if c.eventMatches(ctx, event) {
			c.page.PushBack(event)
		}
	}
}

func (c *EventHistoryCollector) ReadNextEvents(ctx *Context, req *types.ReadNextEvents) soap.HasFault {
	body := &methods.ReadNextEventsBody{}
	if req.MaxCount <= 0 {
		body.Fault_ = Fault("", errInvalidArgMaxCount)
		return body
	}
	body.Res = new(types.ReadNextEventsResponse)

	c.next(req.MaxCount, func(e *list.Element) {
		body.Res.Returnval = append(body.Res.Returnval, e.Value.(types.BaseEvent))
	})

	return body
}

func (c *EventHistoryCollector) ReadPreviousEvents(ctx *Context, req *types.ReadPreviousEvents) soap.HasFault {
	body := &methods.ReadPreviousEventsBody{}
	if req.MaxCount <= 0 {
		body.Fault_ = Fault("", errInvalidArgMaxCount)
		return body
	}
	body.Res = new(types.ReadPreviousEventsResponse)

	c.prev(req.MaxCount, func(e *list.Element) {
		body.Res.Returnval = append(body.Res.Returnval, e.Value.(types.BaseEvent))
	})

	return body
}

func (c *EventHistoryCollector) GetLatestPage() []types.BaseEvent {
	var latestPage []types.BaseEvent

	e := c.page.Back()
	for i := 0; i < c.size; i++ {
		if e == nil {
			break
		}
		latestPage = append(latestPage, e.Value.(types.BaseEvent))
		e = e.Prev()
	}

	return latestPage
}

func (c *EventHistoryCollector) Get() mo.Reference {
	clone := *c

	clone.LatestPage = clone.GetLatestPage()

	return &clone
}
