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
	"github.com/vmware/govmomi/vim25/tasks"
	"github.com/vmware/govmomi/vim25/types"
)

type VirtualMachine struct {
	types.ManagedObjectReference
}

func (v VirtualMachine) Reference() types.ManagedObjectReference {
	return v.ManagedObjectReference
}

func (v VirtualMachine) PowerOn(c *Client) error {
	req := types.PowerOnVM_Task{
		This: v.Reference(),
	}

	task, err := tasks.PowerOnVM(c, &req)
	if err != nil {
		return err
	}

	_, err = c.waitForTask(task)
	return err
}

func (v VirtualMachine) PowerOff(c *Client) error {
	req := types.PowerOffVM_Task{
		This: v.Reference(),
	}

	task, err := tasks.PowerOffVM(c, &req)
	if err != nil {
		return err
	}

	_, err = c.waitForTask(task)
	return err
}

func (v VirtualMachine) Reset(c *Client) error {
	req := types.ResetVM_Task{
		This: v.Reference(),
	}

	task, err := tasks.ResetVM(c, &req)
	if err != nil {
		return err
	}

	_, err = c.waitForTask(task)
	return err
}

func (v VirtualMachine) Destroy(c *Client) error {
	req := types.Destroy_Task{
		This: v.Reference(),
	}

	task, err := tasks.Destroy(c, &req)
	if err != nil {
		return err
	}

	_, err = c.waitForTask(task)
	return err
}
