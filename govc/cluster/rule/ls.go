/*
Copyright (c) 2017 VMware, Inc. All Rights Reserved.

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

package rule

import (
	"context"
	"flag"
	"fmt"
	"io"

	"github.com/vmware/govmomi/govc/cli"
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
  govc cluster.rule.ls -cluster my_cluster -name my_rule`
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
			res = append(res, g.GetClusterRuleInfo().Name)
		}
	} else {
		rule, err := cmd.Rule(ctx)
		if err != nil {
			return err
		}

		res = append(res, rule.ruleType+":")
		switch rule.ruleType {
		case "ClusterAffinityRuleSpec", "ClusterAntiAffinityRuleSpec":
			names, err := cmd.Names(ctx, *rule.refs)
			if err != nil {
				cmd.WriteResult(res)
				return err
			}

			for _, ref := range *rule.refs {
				res = append(res, names[ref])
			}
		case "ClusterVmHostRuleInfo":
			res = append(res, "VmGroupName="+rule.vmGroupName)
			res = append(res, "AffineHostGroupName="+rule.affineHostGroupName)
			res = append(res, "AntiAffineHostGroupName="+rule.antiAffineHostGroupName)
		default:
			res = append(res, "unknown rule type, no further rule details known")
		}

	}

	return cmd.WriteResult(res)
}
