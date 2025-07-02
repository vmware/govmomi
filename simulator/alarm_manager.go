// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"strings"
	"time"

	"github.com/vmware/govmomi/simulator/vpx"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

type AlarmManager struct {
	mo.AlarmManager

	types.GetAlarmResponse
}

func (m *AlarmManager) init(r *Registry) {
	if m.GetAlarmResponse.Returnval != nil {
		return
	}

	m.GetAlarmResponse.Returnval = make([]types.ManagedObjectReference, len(vpx.Alarm))
	for i, alarm := range vpx.Alarm {
		m.GetAlarmResponse.Returnval[i] = alarm.Self
		r.Put(&Alarm{Alarm: alarm})
	}
}

func (*AlarmManager) trimPrefix(s string) string {
	return strings.TrimPrefix(s, "vim.")
}

func (*AlarmManager) key(refs ...types.ManagedObjectReference) string {
	keys := make([]string, len(refs))
	for i := range refs {
		s := strings.Split(refs[i].Value, "-")
		keys[i] = s[len(s)-1]
	}
	return strings.Join(keys, ".")
}

// only handling the common use case of EventEx for now
func (m *AlarmManager) matchAlarm(alarm *Alarm, event *types.EventEx) (*mo.Alarm, types.ManagedEntityStatus) {
	id := event.EventTypeId
	kind := m.trimPrefix(event.ObjectType)

	switch op := alarm.Info.Expression.(type) {
	case *types.OrAlarmExpression:
		for i := range op.Expression {
			switch x := op.Expression[i].(type) {
			case *types.EventAlarmExpression:
				if x.EventTypeId == id && kind == m.trimPrefix(x.ObjectType) {
					return &alarm.Alarm, x.Status
				}
			}
		}
	}
	return nil, ""
}

// update (e.g. triggeredAlarmState) and propagate up the inventory hierarchy
func (*AlarmManager) update(ctx *Context, me mo.Entity, update func(mo.Entity) *types.ManagedObjectReference) {
	for {
		if me == nil {
			break
		}
		ctx.WithLock(me, func() {
			parent := update(me)
			if parent == nil {
				me = nil
			} else {
				me = ctx.Map.Get(*parent).(mo.Entity)
			}
		})
	}
}

// postEvent triggers Alarms based on Events
func (m *AlarmManager) postEvent(ctx *Context, base types.BaseEvent) {
	event, ok := base.(*types.EventEx)
	if !ok {
		return
	}

	entity := types.ManagedObjectReference{Type: event.ObjectType, Value: event.ObjectId}
	me := ctx.Map.Get(entity).(mo.Entity)

	for _, ref := range m.GetAlarmResponse.Returnval {
		alarm := ctx.Map.Get(ref).(*Alarm)
		match, status := m.matchAlarm(alarm, event)
		if match == nil {
			continue
		}

		now := time.Now()
		key := m.key(match.Self, entity)

		update := func(me mo.Entity) *types.ManagedObjectReference {
			obj := me.Entity()

			for i, state := range obj.TriggeredAlarmState {
				if state.Key != key {
					continue
				}

				switch status {
				case state.OverallStatus:
					// no change
					return nil
				case types.ManagedEntityStatusGreen:
					// remove
					obj.TriggeredAlarmState =
						append(obj.TriggeredAlarmState[:i],
							obj.TriggeredAlarmState[i+1:]...)
					return obj.Parent
				default:
					// status change (e.g. yellow -> red)
					obj.TriggeredAlarmState[i].OverallStatus = status
					return obj.Parent
				}
			}

			if status == types.ManagedEntityStatusGreen {
				return nil // green only clears a triggered alarm
			}

			// add
			state := types.AlarmState{
				Key:           key,
				Entity:        entity,
				Alarm:         match.Self,
				OverallStatus: status,
				Time:          now,
				EventKey:      event.Key,
				Acknowledged:  types.NewBool(false),
			}

			obj.TriggeredAlarmState = append(obj.TriggeredAlarmState, state)

			return obj.Parent
		}

		m.update(ctx, me, update)
	}
}

func (m *AlarmManager) GetAlarm(ctx *Context, req *types.GetAlarm) soap.HasFault {
	body := &methods.GetAlarmBody{
		Res: new(types.GetAlarmResponse),
	}

	if req.Entity == nil || *req.Entity == ctx.Map.content().RootFolder {
		body.Res.Returnval = m.GetAlarmResponse.Returnval
	} // else TODO

	return body
}

func (m *AlarmManager) CreateAlarm(ctx *Context, req *types.CreateAlarm) soap.HasFault {
	body := new(methods.CreateAlarmBody)

	name := req.Spec.GetAlarmSpec().Name

	for _, alarm := range ctx.Map.AllReference("Alarm") {
		if alarm.(*Alarm).Info.Name == name {
			body.Fault_ = Fault("", &types.DuplicateName{Name: name})
			return body
		}
	}

	alarm := Alarm{
		Alarm: mo.Alarm{
			Info: types.AlarmInfo{
				AlarmSpec:        *req.Spec.GetAlarmSpec(),
				Entity:           req.Entity,
				LastModifiedTime: time.Now(),
				LastModifiedUser: ctx.Session.UserName,
			},
		},
	}

	ref := ctx.Map.Put(&alarm).Reference()
	alarm.Info.Alarm = ref
	m.GetAlarmResponse.Returnval = append(m.GetAlarmResponse.Returnval, ref)

	body.Res = &types.CreateAlarmResponse{
		Returnval: ref,
	}

	return body
}

func (m *AlarmManager) AcknowledgeAlarm(ctx *Context, req *types.AcknowledgeAlarm) soap.HasFault {
	body := new(methods.AcknowledgeAlarmBody)

	now := types.NewTime(time.Now())
	key := m.key(req.Alarm, req.Entity)
	me := ctx.Map.Get(req.Entity).(mo.Entity)

	update := func(me mo.Entity) *types.ManagedObjectReference {
		obj := me.Entity()

		for i, state := range obj.TriggeredAlarmState {
			if state.Key == key {
				if *obj.TriggeredAlarmState[i].Acknowledged {
					return nil // already ack-ed
				}
				obj.TriggeredAlarmState[i].Acknowledged = types.NewBool(true)
				obj.TriggeredAlarmState[i].AcknowledgedTime = now
				obj.TriggeredAlarmState[i].AcknowledgedByUser = ctx.Session.UserName
				return obj.Parent
			}
		}

		return nil
	}

	m.update(ctx, me, update)

	body.Res = new(types.AcknowledgeAlarmResponse)

	return body
}

type Alarm struct {
	mo.Alarm
}

func (a *Alarm) ReconfigureAlarm(ctx *Context, req *types.ReconfigureAlarm) soap.HasFault {
	body := new(methods.ReconfigureAlarmBody)

	// TODO: spec validation

	a.Info.AlarmSpec = *req.Spec.GetAlarmSpec()

	body.Res = new(types.ReconfigureAlarmResponse)

	return body
}

func (a *Alarm) RemoveAlarm(ctx *Context, req *types.RemoveAlarm) soap.HasFault {
	m := ctx.Map.AlarmManager()

	RemoveReference(&m.GetAlarmResponse.Returnval, req.This)

	ctx.Map.Remove(ctx, req.This)

	return &methods.RemoveAlarmBody{
		Res: new(types.RemoveAlarmResponse),
	}
}
