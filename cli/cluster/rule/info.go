// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package rule

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/cli"
)

type info struct {
	*InfoFlag
}

func init() {
	cli.Register("cluster.rule.info", &info{})
}

func (cmd *info) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.InfoFlag, ctx = NewInfoFlag(ctx)
	cmd.InfoFlag.Register(ctx, f)
}

func (cmd *info) Process(ctx context.Context) error {
	return cmd.InfoFlag.Process(ctx)
}

func (cmd *info) Description() string {
	return `Provides detailed infos about cluster rules, their types and rule members.

Examples:
  govc cluster.rule.info -cluster my_cluster
  govc cluster.rule.info -cluster my_cluster -name my_rule`
}

func (cmd *info) Run(ctx context.Context, f *flag.FlagSet) error {
	var res ruleResult

	rules, err := cmd.Rules(ctx)
	if err != nil {
		return err
	}

	for _, rule := range rules {
		ruleName := rule.GetClusterRuleInfo().Name
		ruleInfo := GetExtendedClusterRuleInfo(rule)
		if cmd.name == "" || cmd.name == ruleName {
			res = append(res, fmt.Sprintf("Rule: %s", ruleName))
			res = append(res, fmt.Sprintf("  Type: %s", ruleInfo.ruleType))
			switch ruleInfo.ruleType {
			case "ClusterAffinityRuleSpec", "ClusterAntiAffinityRuleSpec":
				names, err := cmd.Names(ctx, *ruleInfo.refs)
				if err != nil {
					cmd.WriteResult(res)
					return err
				}

				for _, ref := range *ruleInfo.refs {
					res = append(res, fmt.Sprintf("  VM: %s", names[ref]))
				}
			case "ClusterVmHostRuleInfo":
				res = append(res, fmt.Sprintf("  vmGroupName: %s", ruleInfo.vmGroupName))
				res = append(res, fmt.Sprintf("  affineHostGroupName %s", ruleInfo.affineHostGroupName))
				res = append(res, fmt.Sprintf("  antiAffineHostGroupName %s", ruleInfo.antiAffineHostGroupName))
			case "ClusterDependencyRuleInfo":
				res = append(res, fmt.Sprintf("  VmGroup %s", ruleInfo.VmGroup))
				res = append(res, fmt.Sprintf("  DependsOnVmGroup %s", ruleInfo.DependsOnVmGroup))
			default:
				res = append(res, "unknown rule type, no further rule details known")
			}
		}

	}

	return cmd.WriteResult(res)
}
