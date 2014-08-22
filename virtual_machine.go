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
	"errors"

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

func (v VirtualMachine) Clone(c *Client, folder Folder, name string, config types.VirtualMachineCloneSpec) (*VirtualMachine, error) {
	req := types.CloneVM_Task{
		This:   v.Reference(),
		Folder: folder.Reference(),
		Name:   name,
		Spec:   config,
	}

	task, err := tasks.CloneVM(c, &req)
	if err != nil {
		return nil, err
	}

	res, err := c.waitForTask(task)
	if err != nil {
		return nil, err
	}

	return &VirtualMachine{res.(types.ManagedObjectReference)}, err
}

func (v VirtualMachine) Reconfigure(c *Client, config types.VirtualMachineConfigSpec) error {
	req := types.ReconfigVM_Task{
		This: v.Reference(),
		Spec: config,
	}

	t, err := tasks.ReconfigVM(c, &req)
	if err != nil {
		return err
	}

	info, err := t.Wait()
	if err != nil {
		return err
	}

	if info.Error != nil {
		return errors.New(info.Error.LocalizedMessage)
	}

	return nil
}

func (v VirtualMachine) WaitForIP(c *Client) (string, error) {
	p, err := c.NewPropertyCollector()
	if err != nil {
		return "", err
	}

	defer p.Destroy()

	ref := v.Reference()
	req := types.CreateFilter{
		Spec: types.PropertyFilterSpec{
			ObjectSet: []types.ObjectSpec{
				{
					Obj: ref,
				},
			},
			PropSet: []types.PropertySpec{
				{
					PathSet: []string{"guest.ipAddress"},
					Type:    ref.Type,
				},
			},
		},
	}

	err = p.CreateFilter(req)
	if err != nil {
		return "", err
	}

	for version := ""; ; {
		var prop *types.PropertyChange

		res, err := p.WaitForUpdates(version)
		if err != nil {
			return "", err
		}

		version = res.Version

		for _, fs := range res.FilterSet {
			for _, os := range fs.ObjectSet {
				if os.Obj == ref {
					for _, c := range os.ChangeSet {
						if c.Name != "guest.ipAddress" {
							continue
						}

						if c.Op != types.PropertyChangeOpAssign {
							continue
						}

						prop = &c
						break
					}
				}
			}
		}

		if prop == nil {
			panic("expected to receive property change")
		}

		if prop.Val != nil {
			s, ok := prop.Val.(string)
			if !ok {
				panic("expected to receive string")
			}

			if s != "" {
				return s, nil
			}
		}
	}
}
