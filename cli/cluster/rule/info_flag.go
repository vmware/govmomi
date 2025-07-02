// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package rule

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vim25/types"
)

type InfoFlag struct {
	*flags.ClusterFlag

	rules []types.BaseClusterRuleInfo

	name string
	Long bool
}

func NewInfoFlag(ctx context.Context) (*InfoFlag, context.Context) {
	f := &InfoFlag{}
	f.ClusterFlag, ctx = flags.NewClusterFlag(ctx)
	return f, ctx
}

func (f *InfoFlag) Register(ctx context.Context, fs *flag.FlagSet) {
	f.ClusterFlag.Register(ctx, fs)

	fs.StringVar(&f.name, "name", "", "Cluster rule name")
	fs.BoolVar(&f.Long, "l", false, "Long listing format")
}

func (f *InfoFlag) Process(ctx context.Context) error {
	return f.ClusterFlag.Process(ctx)
}

func (f *InfoFlag) Rules(ctx context.Context) ([]types.BaseClusterRuleInfo, error) {
	if f.rules != nil {
		return f.rules, nil
	}

	cluster, err := f.Cluster()
	if err != nil {
		return nil, err
	}

	config, err := cluster.Configuration(ctx)
	if err != nil {
		return nil, err
	}

	f.rules = config.Rule

	return f.rules, nil
}

type ClusterRuleInfo struct {
	info types.BaseClusterRuleInfo

	ruleType string

	// only ClusterAffinityRuleSpec and ClusterAntiAffinityRuleSpec
	refs *[]types.ManagedObjectReference

	kind string

	// only ClusterVmHostRuleInfo
	vmGroupName             string
	affineHostGroupName     string
	antiAffineHostGroupName string

	// only ClusterDependencyRuleInfo
	VmGroup          string
	DependsOnVmGroup string
}

func (f *InfoFlag) Rule(ctx context.Context) (*ClusterRuleInfo, error) {
	rules, err := f.Rules(ctx)
	if err != nil {
		return nil, err
	}

	for _, rule := range rules {
		if rule.GetClusterRuleInfo().Name != f.name {
			continue
		}

		r := GetExtendedClusterRuleInfo(rule)
		return &r, nil
	}

	return nil, fmt.Errorf("rule %q not found", f.name)
}

func GetExtendedClusterRuleInfo(rule types.BaseClusterRuleInfo) ClusterRuleInfo {
	r := ClusterRuleInfo{info: rule}

	switch info := rule.(type) {
	case *types.ClusterAffinityRuleSpec:
		r.ruleType = "ClusterAffinityRuleSpec"
		r.refs = &info.Vm
		r.kind = "VirtualMachine"
	case *types.ClusterAntiAffinityRuleSpec:
		r.ruleType = "ClusterAntiAffinityRuleSpec"
		r.refs = &info.Vm
		r.kind = "VirtualMachine"
	case *types.ClusterVmHostRuleInfo:
		r.ruleType = "ClusterVmHostRuleInfo"
		r.vmGroupName = info.VmGroupName
		r.affineHostGroupName = info.AffineHostGroupName
		r.antiAffineHostGroupName = info.AntiAffineHostGroupName
	case *types.ClusterDependencyRuleInfo:
		r.ruleType = "ClusterDependencyRuleInfo"
		r.VmGroup = info.VmGroup
		r.DependsOnVmGroup = info.DependsOnVmGroup
	}
	return r
}

func (f *InfoFlag) Apply(ctx context.Context, update types.ArrayUpdateSpec, info types.BaseClusterRuleInfo) error {
	spec := &types.ClusterConfigSpecEx{
		RulesSpec: []types.ClusterRuleSpec{
			{
				ArrayUpdateSpec: update,
				Info:            info,
			},
		},
	}

	return f.Reconfigure(ctx, spec)
}

type SpecFlag struct {
	types.ClusterRuleInfo
	types.ClusterVmHostRuleInfo
	types.ClusterAffinityRuleSpec     //nolint:govet
	types.ClusterAntiAffinityRuleSpec //nolint:govet
}

func (s *SpecFlag) Register(ctx context.Context, f *flag.FlagSet) {
	f.Var(flags.NewOptionalBool(&s.Enabled), "enable", "Enable rule")
	f.Var(flags.NewOptionalBool(&s.Mandatory), "mandatory", "Enforce rule compliance")

	f.StringVar(&s.VmGroupName, "vm-group", "", "VM group name")
	f.StringVar(&s.AffineHostGroupName, "host-affine-group", "", "Host affine group name")
	f.StringVar(&s.AntiAffineHostGroupName, "host-anti-affine-group", "", "Host anti-affine group name")
}
