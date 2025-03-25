// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package guest

import (
	"context"
	"sync"

	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type OperationsManager struct {
	c  *vim25.Client
	vm types.ManagedObjectReference
}

func NewOperationsManager(c *vim25.Client, vm types.ManagedObjectReference) *OperationsManager {
	return &OperationsManager{c, vm}
}

func (m OperationsManager) retrieveOne(ctx context.Context, p string, dst *mo.GuestOperationsManager) error {
	pc := property.DefaultCollector(m.c)
	return pc.RetrieveOne(ctx, *m.c.ServiceContent.GuestOperationsManager, []string{p}, dst)
}

func (m OperationsManager) AuthManager(ctx context.Context) (*AuthManager, error) {
	var g mo.GuestOperationsManager

	err := m.retrieveOne(ctx, "authManager", &g)
	if err != nil {
		return nil, err
	}

	return &AuthManager{*g.AuthManager, m.vm, m.c}, nil
}

func (m OperationsManager) FileManager(ctx context.Context) (*FileManager, error) {
	var g mo.GuestOperationsManager

	err := m.retrieveOne(ctx, "fileManager", &g)
	if err != nil {
		return nil, err
	}

	return &FileManager{
		ManagedObjectReference: *g.FileManager,
		vm:                     m.vm,
		c:                      m.c,
		mu:                     new(sync.Mutex),
		hosts:                  make(map[string]string),
	}, nil
}

func (m OperationsManager) ProcessManager(ctx context.Context) (*ProcessManager, error) {
	var g mo.GuestOperationsManager

	err := m.retrieveOne(ctx, "processManager", &g)
	if err != nil {
		return nil, err
	}

	return &ProcessManager{*g.ProcessManager, m.vm, m.c}, nil
}
