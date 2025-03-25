// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package esx

import (
	"context"
	"strings"

	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type hostInfo struct {
	*Executor
	wids map[string]string
}

type GuestInfo struct {
	c     *vim25.Client
	hosts map[string]*hostInfo
}

func NewGuestInfo(c *vim25.Client) *GuestInfo {
	return &GuestInfo{
		c:     c,
		hosts: make(map[string]*hostInfo),
	}
}

func (g *GuestInfo) hostInfo(ctx context.Context, ref *types.ManagedObjectReference) (*hostInfo, error) {
	// cache exectuor and uuid -> worldid map
	if h, ok := g.hosts[ref.Value]; ok {
		return h, nil
	}

	e, err := NewExecutor(ctx, g.c, ref)
	if err != nil {
		return nil, err
	}

	res, err := e.Run(ctx, []string{"vm", "process", "list"})
	if err != nil {
		return nil, err
	}

	ids := make(map[string]string, len(res.Values))

	for _, process := range res.Values {
		// Normalize uuid, esxcli and mo.VirtualMachine have different formats
		uuid := strings.Replace(process["UUID"][0], " ", "", -1)
		uuid = strings.Replace(uuid, "-", "", -1)

		ids[uuid] = process["WorldID"][0]
	}

	h := &hostInfo{e, ids}
	g.hosts[ref.Value] = h

	return h, nil
}

// IpAddress attempts to find the guest IP address using esxcli.
// ESX hosts must be configured with the /Net/GuestIPHack enabled.
// For example:
// $ govc host.esxcli -- system settings advanced set -o /Net/GuestIPHack -i 1
func (g *GuestInfo) IpAddress(ctx context.Context, vm mo.Reference) (string, error) {
	const any = "0.0.0.0"
	var mvm mo.VirtualMachine

	pc := property.DefaultCollector(g.c)
	err := pc.RetrieveOne(ctx, vm.Reference(), []string{"runtime.host", "config.uuid"}, &mvm)
	if err != nil {
		return "", err
	}

	h, err := g.hostInfo(ctx, mvm.Runtime.Host)
	if err != nil {
		return "", err
	}

	// Normalize uuid, esxcli and mo.VirtualMachine have different formats
	uuid := strings.Replace(mvm.Config.Uuid, "-", "", -1)

	if wid, ok := h.wids[uuid]; ok {
		res, err := h.Run(ctx, []string{"network", "vm", "port", "list", "--world-id", wid})
		if err != nil {
			return "", err
		}

		for _, val := range res.Values {
			if ip, ok := val["IPAddress"]; ok {
				if ip[0] != any {
					return ip[0], nil
				}
			}
		}
	}

	return any, nil
}
