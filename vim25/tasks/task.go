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

package tasks

import (
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

type Task interface {
	Wait() (*types.TaskInfo, error)
	Cancel() error
}

var serviceInstance = types.ManagedObjectReference{
	Type:  "ServiceInstance",
	Value: "ServiceInstance",
}

func serviceContent(c soap.RoundTripper) (*types.ServiceContent, error) {
	req := types.RetrieveServiceContent{
		This: serviceInstance,
	}

	res, err := methods.RetrieveServiceContent(c, &req)
	if err != nil {
		return nil, err
	}

	return &res.Returnval, nil
}

type task struct {
	r soap.RoundTripper

	// References
	taskRef              types.ManagedObjectReference
	propertyCollectorRef types.ManagedObjectReference
	filterRef            types.ManagedObjectReference

	// Final result or error
	res *types.TaskInfo
	err error
}

func newTask(r soap.RoundTripper, t types.ManagedObjectReference) Task {
	task := task{
		r: r,

		taskRef: t,
	}

	return &task
}

func (t *task) propertyCollector() error {
	sc, err := serviceContent(t.r)
	if err != nil {
		return err
	}

	req := types.CreatePropertyCollector{
		This: sc.PropertyCollector,
	}

	res, err := methods.CreatePropertyCollector(t.r, &req)
	if err != nil {
		return err
	}

	t.propertyCollectorRef = res.Returnval
	return nil
}

func (t *task) destroyPropertyCollector() error {
	if t.propertyCollectorRef.Type == "" {
		return nil
	}

	req := types.DestroyPropertyCollector{
		This: t.propertyCollectorRef,
	}

	_, err := methods.DestroyPropertyCollector(t.r, &req)
	if err != nil {
		return err
	}

	t.propertyCollectorRef = types.ManagedObjectReference{}
	return nil
}

func (t *task) createFilter() error {
	err := t.propertyCollector()
	if err != nil {
		return err
	}

	req := types.CreateFilter{
		This: t.propertyCollectorRef,
		Spec: types.PropertyFilterSpec{
			ObjectSet: []types.ObjectSpec{
				{
					Obj: t.taskRef,
				},
			},
			PropSet: []types.PropertySpec{
				{
					All:  true,
					Type: "Task",
				},
			},
		},
		PartialUpdates: false,
	}

	res, err := methods.CreateFilter(t.r, &req)
	if err != nil {
		return err
	}

	t.filterRef = res.Returnval
	return nil
}

func (t *task) destroyFilter() error {
	if t.filterRef.Type == "" {
		return nil
	}

	req := types.DestroyPropertyFilter{
		This: t.filterRef,
	}

	_, err := methods.DestroyPropertyFilter(t.r, &req)
	if err != nil {
		return err
	}

	t.filterRef = types.ManagedObjectReference{}
	return nil
}

func (t *task) waitLoop() {
	defer t.destroyPropertyCollector()

	err := t.createFilter()
	if err != nil {
		t.err = err
		return
	}

	for version := ""; ; {
		var ti *types.TaskInfo

		req := types.WaitForUpdatesEx{
			This:    t.propertyCollectorRef,
			Version: version,
		}

		res, err := methods.WaitForUpdatesEx(t.r, &req)
		if err != nil {
			panic(err)
		}

		version = res.Returnval.Version

		// Find TaskInfo in response
		for _, f := range res.Returnval.FilterSet {
			if f.Filter == t.filterRef {
				for _, o := range f.ObjectSet {
					if o.Obj == t.taskRef {
						for _, c := range o.ChangeSet {
							if c.Name == "info" {
								tiv := c.Val.(types.TaskInfo)
								ti = &tiv
							}
						}
					}
				}
			}
		}

		if ti == nil {
			continue
		}

		switch ti.State {
		case types.TaskInfoStateQueued, types.TaskInfoStateRunning:
			// Skip
		case types.TaskInfoStateSuccess, types.TaskInfoStateError:
			t.res = ti
			return
		default:
			panic("unknown task state: " + ti.State)
		}
	}
}

func (t *task) Wait() (*types.TaskInfo, error) {
	t.waitLoop()

	return t.res, t.err
}

func (t *task) Cancel() error {
	panic("todo")
}
