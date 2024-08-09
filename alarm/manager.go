/*
Copyright (c) 2024-2024 VMware, Inc. All Rights Reserved.

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

package alarm

import (
	"context"

	"github.com/vmware/govmomi/event"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type Manager struct {
	object.Common

	pc *property.Collector
}

var Severity = map[types.ManagedEntityStatus]string{
	types.ManagedEntityStatusGray:   "Unknown",
	types.ManagedEntityStatusGreen:  "Normal",
	types.ManagedEntityStatusYellow: "Warning",
	types.ManagedEntityStatusRed:    "Alert",
}

// GetManager wraps NewManager, returning ErrNotSupported
// when the client is not connected to a vCenter instance.
func GetManager(c *vim25.Client) (*Manager, error) {
	if c.ServiceContent.AlarmManager == nil {
		return nil, object.ErrNotSupported
	}
	return NewManager(c), nil
}

func NewManager(c *vim25.Client) *Manager {
	m := Manager{
		Common: object.NewCommon(c, *c.ServiceContent.AlarmManager),
		pc:     property.DefaultCollector(c),
	}

	return &m
}

func (m Manager) CreateAlarm(ctx context.Context, entity object.Reference, spec types.BaseAlarmSpec) (*types.ManagedObjectReference, error) {
	req := types.CreateAlarm{
		This:   m.Reference(),
		Entity: entity.Reference(),
		Spec:   spec,
	}

	res, err := methods.CreateAlarm(ctx, m.Client(), &req)
	if err != nil {
		return nil, err
	}
	return &res.Returnval, err
}

func (m Manager) AcknowledgeAlarm(ctx context.Context, alarm types.ManagedObjectReference, entity object.Reference) error {
	req := types.AcknowledgeAlarm{
		This:   m.Reference(),
		Alarm:  alarm,
		Entity: entity.Reference(),
	}

	_, err := methods.AcknowledgeAlarm(ctx, m.Client(), &req)

	return err
}

// GetAlarm returns available alarms defined on the entity.
func (m Manager) GetAlarm(ctx context.Context, entity object.Reference) ([]mo.Alarm, error) {
	req := types.GetAlarm{
		This: m.Reference(),
	}

	if entity != nil {
		req.Entity = types.NewReference(entity.Reference())
	}

	res, err := methods.GetAlarm(ctx, m.Client(), &req)
	if err != nil {
		return nil, err
	}

	if len(res.Returnval) == 0 {
		return nil, nil
	}

	alarms := make([]mo.Alarm, 0, len(res.Returnval))

	err = m.pc.Retrieve(ctx, res.Returnval, []string{"info"}, &alarms)
	if err != nil {
		return nil, err
	}

	return alarms, nil
}

// StateInfo combines AlarmState with Alarm.Info
type StateInfo struct {
	types.AlarmState
	Info  *types.AlarmInfo `json:"name,omitempty"`
	Path  string           `json:"path,omitempty"`
	Event types.BaseEvent  `json:"event,omitempty"`
}

// StateInfoOptions for the GetStateInfo method
type StateInfoOptions struct {
	Declared      bool
	InventoryPath bool
	Event         bool
}

// GetStateInfo combines AlarmState with Alarm.Info
func (m Manager) GetStateInfo(ctx context.Context, entity object.Reference, opts StateInfoOptions) ([]StateInfo, error) {
	prop := "triggeredAlarmState"
	if opts.Declared {
		prop = "declaredAlarmState"
		opts.Event = false
	}

	var e mo.ManagedEntity

	err := m.pc.RetrieveOne(ctx, entity.Reference(), []string{prop}, &e)
	if err != nil {
		return nil, err
	}

	var objs []types.ManagedObjectReference
	alarms := append(e.DeclaredAlarmState, e.TriggeredAlarmState...)
	if len(alarms) == 0 {
		return nil, nil
	}
	for i := range alarms {
		objs = append(objs, alarms[i].Alarm)
	}

	var info []mo.Alarm
	err = m.pc.Retrieve(ctx, objs, []string{"info"}, &info)
	if err != nil {
		return nil, err
	}

	state := make([]StateInfo, len(alarms))
	paths := make(map[types.ManagedObjectReference]string)

	em := event.NewManager(m.Client())

	for i, a := range alarms {
		path := paths[a.Entity]
		if opts.InventoryPath {
			if path == "" {
				path, err = find.InventoryPath(ctx, m.Client(), a.Entity)
				if err != nil {
					return nil, err
				}
				paths[a.Entity] = path
			}
		} else {
			path = a.Entity.String()
		}

		state[i] = StateInfo{
			AlarmState: a,
			Path:       path,
		}

		for j := range info {
			if info[j].Self == a.Alarm {
				state[i].Info = &info[j].Info
				break
			}
		}

		if !opts.Event || a.EventKey == 0 {
			continue
		}

		spec := types.EventFilterSpec{
			EventChainId: a.EventKey,
			Entity: &types.EventFilterSpecByEntity{
				Entity:    a.Entity,
				Recursion: types.EventFilterSpecRecursionOptionSelf,
			},
		}

		events, err := em.QueryEvents(ctx, spec)
		if err != nil {
			return nil, err
		}

		for j := range events {
			if events[j].GetEvent().Key == a.EventKey {
				state[i].Event = events[j]
				break
			}
		}
	}

	return state, nil
}
