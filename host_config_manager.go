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
	"github.com/vmware/govmomi/vim25/mo"
)

type HostConfigManager struct {
	c *Client
	h HostSystem
}

func (m HostConfigManager) NetworkSystem() (*HostNetworkSystem, error) {
	var h mo.HostSystem

	err := m.c.Properties(m.h.Reference(), []string{"configManager.networkSystem"}, &h)
	if err != nil {
		return nil, err
	}

	return &HostNetworkSystem{
		ManagedObjectReference: *h.ConfigManager.NetworkSystem,
		c: m.c,
	}, nil
}
