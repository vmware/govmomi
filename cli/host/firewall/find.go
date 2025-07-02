// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package firewall

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/esx"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/types"
)

type find struct {
	*flags.ClientFlag
	*flags.OutputFlag
	*flags.HostSystemFlag

	enabled bool
	check   bool

	types.HostFirewallRule
}

func init() {
	cli.Register("firewall.ruleset.find", &find{})
}

func (cmd *find) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)
	cmd.HostSystemFlag, ctx = flags.NewHostSystemFlag(ctx)
	cmd.HostSystemFlag.Register(ctx, f)

	f.BoolVar(&cmd.check, "c", true, "Check if esx firewall is enabled")
	f.BoolVar(&cmd.enabled, "enabled", true, "Find enabled rule sets if true, disabled if false")
	f.StringVar((*string)(&cmd.Direction), "direction", string(types.HostFirewallRuleDirectionOutbound), "Direction")
	f.StringVar((*string)(&cmd.PortType), "type", string(types.HostFirewallRulePortTypeDst), "Port type")
	f.StringVar((*string)(&cmd.Protocol), "proto", string(types.HostFirewallRuleProtocolTcp), "Protocol")
	f.Var(flags.NewInt32(&cmd.Port), "port", "Port")
}

func (cmd *find) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.HostSystemFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *find) Description() string {
	return `Find firewall rulesets matching the given rule.

For a complete list of rulesets: govc host.esxcli network firewall ruleset list
For a complete list of rules:    govc host.esxcli network firewall ruleset rule list

Examples:
  govc firewall.ruleset.find -direction inbound -port 22
  govc firewall.ruleset.find -direction outbound -port 2377`
}

func (cmd *find) Run(ctx context.Context, f *flag.FlagSet) error {
	host, err := cmd.HostSystem()
	if err != nil {
		return err
	}

	fs, err := host.ConfigManager().FirewallSystem(ctx)
	if err != nil {
		return err
	}

	if cmd.check {
		x, err := esx.NewExecutor(ctx, host.Client(), host)
		if err != nil {
			return err
		}

		esxfw, err := x.GetFirewallInfo(ctx)
		if err != nil {
			return err
		}

		if !esxfw.Enabled {
			fmt.Fprintln(os.Stderr, "host firewall is disabled")
		}
	}

	info, err := fs.Info(ctx)
	if err != nil {
		return err
	}

	if f.NArg() != 0 {
		// TODO: f.Args() -> types.HostFirewallRulesetIpList
		return flag.ErrHelp
	}

	rs := object.HostFirewallRulesetList(info.Ruleset)
	matched, err := rs.EnabledByRule(cmd.HostFirewallRule, cmd.enabled)

	if err != nil {
		return err
	}

	for _, r := range matched {
		fmt.Println(r.Key)
	}

	return nil
}
