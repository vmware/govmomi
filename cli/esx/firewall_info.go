// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package esx

import "context"

type FirewallInfo struct {
	Loaded        bool   `json:"loaded"`
	Enabled       bool   `json:"enabled"`
	DefaultAction string `json:"defaultAction"`
}

// GetFirewallInfo via 'esxcli network firewall get'
// The HostFirewallSystem type does not expose this data.
// This helper can be useful in particular to determine if the firewall is enabled or disabled.
func (x *Executor) GetFirewallInfo(ctx context.Context) (*FirewallInfo, error) {
	res, err := x.Run(ctx, []string{"network", "firewall", "get"})
	if err != nil {
		return nil, err
	}

	info := &FirewallInfo{
		Loaded:        res.Values[0]["Loaded"][0] == "true",
		Enabled:       res.Values[0]["Enabled"][0] == "true",
		DefaultAction: res.Values[0]["DefaultAction"][0],
	}

	return info, nil
}
