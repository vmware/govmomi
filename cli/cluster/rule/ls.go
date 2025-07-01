// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package rule

import (
	"context"
	"flag"
	"fmt"
	"io"

	"github.com/vmware/govmomi/cli"
)

type ls struct {
	*InfoFlag
}

func init() {
	cli.Register("cluster.rule.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.InfoFlag, ctx = NewInfoFlag(ctx)
	cmd.InfoFlag.Register(ctx, f)
}

func (cmd *ls) Process(ctx context.Context) error {
	return cmd.InfoFlag.Process(ctx)
}

func (cmd *ls) Description() string {
	return `List cluster rules and rule members.

Examples:
  govc cluster.rule.ls -cluster my_cluster
  govc cluster.rule.ls -cluster my_cluster -name my_rule
  govc cluster.rule.ls -cluster my_cluster -l
  govc cluster.rule.ls -cluster my_cluster -name my_rule -l`
}

type ruleResult []string

func (r ruleResult) Write(w io.Writer) error {
	for i := range r {
		fmt.Fprintln(w, r[i])
	}

	return nil
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {
	var res ruleResult

	if cmd.name == "" {
		rules, err := cmd.Rules(ctx)
		if err != nil {
			return err
		}

		for _, g := range rules {
			ruleName := g.GetClusterRuleInfo().Name
			if cmd.Long {
				ruleTypeInfo := GetExtendedClusterRuleInfo(g).ruleType
				res = append(res, fmt.Sprintf("%s (%s)", ruleName, ruleTypeInfo))
			} else {
				res = append(res, ruleName)
			}
		}
	} else {
		rule, err := cmd.Rule(ctx)
		if err != nil {
			return err
		}

		switch rule.ruleType {
		case "ClusterAffinityRuleSpec", "ClusterAntiAffinityRuleSpec":
			names, err := cmd.Names(ctx, *rule.refs)
			if err != nil {
				cmd.WriteResult(res)
				return err
			}

			for _, ref := range *rule.refs {
				res = extendedAppend(res, cmd.Long, names[ref], "VM")
			}
		case "ClusterVmHostRuleInfo":
			res = extendedAppend(res, cmd.Long, rule.vmGroupName, "vmGroupName")
			res = extendedAppend(res, cmd.Long, rule.affineHostGroupName, "affineHostGroupName")
			res = extendedAppend(res, cmd.Long, rule.antiAffineHostGroupName, "antiAffineHostGroupName")
		case "ClusterDependencyRuleInfo":
			res = extendedAppend(res, cmd.Long, rule.VmGroup, "VmGroup")
			res = extendedAppend(res, cmd.Long, rule.DependsOnVmGroup, "DependsOnVmGroup")
		default:
			res = append(res, "unknown rule type, no further rule details known")
		}

	}

	return cmd.WriteResult(res)
}

func extendedAppend(res ruleResult, long bool, entryValue string, entryType string) ruleResult {
	var newres ruleResult
	if long {
		newres = append(res, fmt.Sprintf("%s (%s)", entryValue, entryType))
	} else {
		newres = append(res, entryValue)
	}
	return newres
}
