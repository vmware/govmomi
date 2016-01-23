/*
Copyright (c) 2015 VMware, Inc. All Rights Reserved.

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

package object

import (
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
	"golang.org/x/net/context"
)

// HostConfigManager represents a client for a host config
type HostConfigManager struct {
	Common
}

// NewHostConfigManager creates a new host config client
func NewHostConfigManager(c *vim25.Client, ref types.ManagedObjectReference) *HostConfigManager {
	return &HostConfigManager{
		Common: NewCommon(c, ref),
	}
}

// DatastoreSystem gets the datastore system
func (m HostConfigManager) DatastoreSystem(ctx context.Context) (*HostDatastoreSystem, error) {
	var h mo.HostSystem

	err := m.Properties(ctx, m.Reference(), []string{"configManager.datastoreSystem"}, &h)
	if err != nil {
		return nil, err
	}

	return NewHostDatastoreSystem(m.c, *h.ConfigManager.DatastoreSystem), nil
}

// NetworkSystem gets the config manager network system
func (m HostConfigManager) NetworkSystem(ctx context.Context) (*HostNetworkSystem, error) {
	var h mo.HostSystem

	err := m.Properties(ctx, m.Reference(), []string{"configManager.networkSystem"}, &h)
	if err != nil {
		return nil, err
	}

	return NewHostNetworkSystem(m.c, *h.ConfigManager.NetworkSystem), nil
}

// FirewallSystem gets the firewall system config manager
func (m HostConfigManager) FirewallSystem(ctx context.Context) (*HostFirewallSystem, error) {
	var h mo.HostSystem

	err := m.Properties(ctx, m.Reference(), []string{"configManager.firewallSystem"}, &h)
	if err != nil {
		return nil, err
	}

	return NewHostFirewallSystem(m.c, *h.ConfigManager.FirewallSystem), nil
}

// StorageSystem gets a storage system config manager
func (m HostConfigManager) StorageSystem(ctx context.Context) (*HostStorageSystem, error) {
	var h mo.HostSystem

	err := m.Properties(ctx, m.Reference(), []string{"configManager.storageSystem"}, &h)
	if err != nil {
		return nil, err
	}

	return NewHostStorageSystem(m.c, *h.ConfigManager.StorageSystem), nil
}

// VirtualNicManager gets a virtual nic config manager
func (m HostConfigManager) VirtualNicManager(ctx context.Context) (*HostVirtualNicManager, error) {
	var h mo.HostSystem

	err := m.Properties(ctx, m.Reference(), []string{"configManager.virtualNicManager"}, &h)
	if err != nil {
		return nil, err
	}

	return NewHostVirtualNicManager(m.c, *h.ConfigManager.VirtualNicManager, m.Reference()), nil
}

// VsanSystem gets the VSAN config manager
func (m HostConfigManager) VsanSystem(ctx context.Context) (*HostVsanSystem, error) {
	var h mo.HostSystem

	err := m.Properties(ctx, m.Reference(), []string{"configManager.vsanSystem"}, &h)
	if err != nil {
		return nil, err
	}

	return NewHostVsanSystem(m.c, *h.ConfigManager.VsanSystem), nil
}
