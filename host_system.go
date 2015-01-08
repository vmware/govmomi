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
	"fmt"
	"net"

	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type HostSystem struct {
	types.ManagedObjectReference

	InventoryPath string

	c *Client
}

func NewHostSystem(c *Client, ref types.ManagedObjectReference) *HostSystem {
	return &HostSystem{
		ManagedObjectReference: ref,
		c: c,
	}
}

func (h HostSystem) Reference() types.ManagedObjectReference {
	return h.ManagedObjectReference
}

func (h HostSystem) ConfigManager() *HostConfigManager {
	return &HostConfigManager{h.c, h}
}

func (h HostSystem) ResourcePool() (*ResourcePool, error) {
	var mh mo.HostSystem
	err := h.c.Properties(h.Reference(), []string{"parent"}, &mh)
	if err != nil {
		return nil, err
	}

	var mcr *mo.ComputeResource
	var parent interface{}

	switch mh.Parent.Type {
	case "ComputeResource":
		mcr = new(mo.ComputeResource)
		parent = mcr
	case "ClusterComputeResource":
		mcc := new(mo.ClusterComputeResource)
		mcr = &mcc.ComputeResource
		parent = mcc
	default:
		return nil, fmt.Errorf("unknown host parent type: %s", mh.Parent.Type)
	}

	err = h.c.Properties(*mh.Parent, []string{"resourcePool"}, parent)
	if err != nil {
		return nil, err
	}

	pool := NewResourcePool(h.c, *mcr.ResourcePool)
	return pool, nil
}

func (h HostSystem) ManagementIPs() ([]net.IP, error) {
	var mh mo.HostSystem

	err := h.c.Properties(h.Reference(), []string{"config.virtualNicManagerInfo.netConfig"}, &mh)
	if err != nil {
		return nil, err
	}

	var ips []net.IP
	for _, nc := range mh.Config.VirtualNicManagerInfo.NetConfig {
		if nc.NicType == "management" && len(nc.CandidateVnic) > 0 {
			ip := net.ParseIP(nc.CandidateVnic[0].Spec.Ip.IpAddress)
			if ip != nil {
				ips = append(ips, ip)
			}
		}
	}

	return ips, nil
}
