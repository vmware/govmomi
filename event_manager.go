/*
Copyright (c) 2014 VMware, Inc. All Rights Reserved.

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

package govmomi

import (
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/types"
)

type EventManager struct {
	c *Client
}

func (em EventManager) CreateCollectorForEvents(filter types.EventFilterSpec) (*types.ManagedObjectReference, error) {
	req := types.CreateCollectorForEvents{
		This:   *em.c.ServiceContent.EventManager,
		Filter: filter,
	}

	res, err := methods.CreateCollectorForEvents(em.c, &req)
	if err != nil {
		return nil, err
	}
	return &res.Returnval, nil
}

func (em EventManager) LogUserEvent(entity types.ManagedObjectReference, msg string) error {
	req := types.LogUserEvent{
		This:   *em.c.ServiceContent.EventManager,
		Entity: entity,
		Msg:    msg,
	}

	_, err := methods.LogUserEvent(em.c, &req)
	if err != nil {
		return err
	}

	return nil
}

func (em EventManager) PostEvent(eventToPost types.BaseEvent, taskInfo types.TaskInfo) error {
	req := types.PostEvent{
		This:        *em.c.ServiceContent.EventManager,
		EventToPost: eventToPost,
		TaskInfo:    &taskInfo,
	}

	_, err := methods.PostEvent(em.c, &req)
	if err != nil {
		return err
	}

	return nil
}

func (em EventManager) QueryEvents(filter types.EventFilterSpec) ([]types.BaseEvent, error) {
	req := types.QueryEvents{
		This:   *em.c.ServiceContent.EventManager,
		Filter: filter,
	}

	res, err := methods.QueryEvents(em.c, &req)
	if err != nil {
		return nil, err
	}
	return res.Returnval, nil
}

func (em EventManager) RetrieveArgumentDescription(eventTypeID string) ([]types.EventArgDesc, error) {
	req := types.RetrieveArgumentDescription{
		This:        *em.c.ServiceContent.EventManager,
		EventTypeId: eventTypeID,
	}

	res, err := methods.RetrieveArgumentDescription(em.c, &req)
	if err != nil {
		return nil, err
	}
	return res.Returnval, nil

}
