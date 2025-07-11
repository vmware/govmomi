// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package event

import (
	"context"
	"fmt"
	"reflect"
	"sync"

	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type Manager struct {
	r types.ManagedObjectReference
	c *vim25.Client

	eventCategory   map[string]string
	eventCategoryMu *sync.Mutex
	maxObjects      int
}

func NewManager(c *vim25.Client) *Manager {
	m := Manager{
		r:               c.ServiceContent.EventManager.Reference(),
		c:               c,
		eventCategory:   make(map[string]string),
		eventCategoryMu: new(sync.Mutex),
		maxObjects:      10,
	}

	return &m
}

// Reference returns the event.Manager MOID
func (m Manager) Reference() types.ManagedObjectReference {
	return m.r
}

func (m Manager) Client() *vim25.Client {
	return m.c
}

func (m Manager) CreateCollectorForEvents(ctx context.Context, filter types.EventFilterSpec) (*HistoryCollector, error) {
	req := types.CreateCollectorForEvents{
		This:   m.r,
		Filter: filter,
	}

	res, err := methods.CreateCollectorForEvents(ctx, m.c, &req)
	if err != nil {
		return nil, err
	}

	return newHistoryCollector(m.c, res.Returnval), nil
}

func (m Manager) LogUserEvent(ctx context.Context, entity types.ManagedObjectReference, msg string) error {
	req := types.LogUserEvent{
		This:   m.r,
		Entity: entity,
		Msg:    msg,
	}

	_, err := methods.LogUserEvent(ctx, m.c, &req)
	if err != nil {
		return err
	}

	return nil
}

func (m Manager) PostEvent(ctx context.Context, eventToPost types.BaseEvent, taskInfo ...types.TaskInfo) error {
	req := types.PostEvent{
		This:        m.r,
		EventToPost: eventToPost,
	}

	if len(taskInfo) == 1 {
		req.TaskInfo = &taskInfo[0]
	}

	_, err := methods.PostEvent(ctx, m.c, &req)
	if err != nil {
		return err
	}

	return nil
}

func (m Manager) QueryEvents(ctx context.Context, filter types.EventFilterSpec) ([]types.BaseEvent, error) {
	req := types.QueryEvents{
		This:   m.r,
		Filter: filter,
	}

	res, err := methods.QueryEvents(ctx, m.c, &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}

func (m Manager) RetrieveArgumentDescription(ctx context.Context, eventTypeID string) ([]types.EventArgDesc, error) {
	req := types.RetrieveArgumentDescription{
		This:        m.r,
		EventTypeId: eventTypeID,
	}

	res, err := methods.RetrieveArgumentDescription(ctx, m.c, &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}

func (m Manager) eventCategoryMap(ctx context.Context) (map[string]string, error) {
	m.eventCategoryMu.Lock()
	defer m.eventCategoryMu.Unlock()

	if len(m.eventCategory) != 0 {
		return m.eventCategory, nil
	}

	var o mo.EventManager

	ps := []string{"description.eventInfo"}
	err := property.DefaultCollector(m.c).RetrieveOne(ctx, m.r, ps, &o)
	if err != nil {
		return nil, err
	}

	for _, info := range o.Description.EventInfo {
		m.eventCategory[info.Key] = info.Category
	}

	return m.eventCategory, nil
}

// EventCategory returns the category for an event, such as "info" or "error" for example.
func (m Manager) EventCategory(ctx context.Context, event types.BaseEvent) (string, error) {
	// Most of the event details are included in the Event.FullFormattedMessage, but the category
	// is only available via the EventManager description.eventInfo property.  The value of this
	// property is static, so we fetch and once and cache.
	eventCategory, err := m.eventCategoryMap(ctx)
	if err != nil {
		return "", err
	}

	switch e := event.(type) {
	case *types.EventEx:
		if e.Severity == "" {
			return "info", nil
		}
		return e.Severity, nil
	}

	class := reflect.TypeOf(event).Elem().Name()

	return eventCategory[class], nil
}

// Events gets the events from the specified object(s) and optionanlly tail the
// event stream
func (m Manager) Events(ctx context.Context, objects []types.ManagedObjectReference, pageSize int32, tail bool, force bool, f func(types.ManagedObjectReference, []types.BaseEvent) error, kind ...string) error {
	// TODO: deprecated this method and add one that uses a single config struct, so we can extend further without breaking the method signature.
	if len(objects) >= m.maxObjects && !force {
		return fmt.Errorf("maximum number of objects to monitor (%d) exceeded, refine search", m.maxObjects)
	}

	proc := newEventProcessor(m, pageSize, f, kind)
	for _, o := range objects {
		proc.addObject(ctx, o)
	}

	defer proc.destroy()

	return proc.run(ctx, tail)
}
