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

package guest

import (
	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
	"golang.org/x/net/context"
)

type OperationsManager struct {
	c  *govmomi.Client
	vm types.ManagedObjectReference
}

func NewOperationsManager(c *govmomi.Client, vm types.ManagedObjectReference) *OperationsManager {
	return &OperationsManager{c, vm}
}

func (m OperationsManager) AuthManager(ctx context.Context) (*AuthManager, error) {
	var g mo.GuestOperationsManager

	err := m.c.Properties(*m.c.ServiceContent.GuestOperationsManager, []string{"authManager"}, &g)
	if err != nil {
		return nil, err
	}

	return &AuthManager{*g.AuthManager, m.vm, m.c}, nil
}

func (m OperationsManager) FileManager(ctx context.Context) (*FileManager, error) {
	var g mo.GuestOperationsManager

	err := m.c.Properties(*m.c.ServiceContent.GuestOperationsManager, []string{"fileManager"}, &g)
	if err != nil {
		return nil, err
	}

	return &FileManager{*g.FileManager, m.vm, m.c}, nil
}

func (m OperationsManager) ProcessManager(ctx context.Context) (*ProcessManager, error) {
	var g mo.GuestOperationsManager

	err := m.c.Properties(*m.c.ServiceContent.GuestOperationsManager, []string{"processManager"}, &g)
	if err != nil {
		return nil, err
	}

	return &ProcessManager{*g.ProcessManager, m.vm, m.c}, nil
}
