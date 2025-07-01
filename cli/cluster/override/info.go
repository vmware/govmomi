// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package override

import (
	"context"
	"flag"
	"fmt"
	"io"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/types"
)

type info struct {
	*flags.ClusterFlag
}

func init() {
	cli.Register("cluster.override.info", &info{})
}

func (cmd *info) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClusterFlag, ctx = flags.NewClusterFlag(ctx)
	cmd.ClusterFlag.Register(ctx, f)
}

func (cmd *info) Description() string {
	return `Cluster VM overrides info.

Examples:
  govc cluster.override.info
  govc cluster.override.info -json`
}

func (cmd *info) Process(ctx context.Context) error {
	return cmd.ClusterFlag.Process(ctx)
}

type Override struct {
	id            types.ManagedObjectReference
	Name          string                            `json:"name"`
	Host          string                            `json:"host,omitempty"`
	DRS           *types.ClusterDrsVmConfigInfo     `json:"drs,omitempty"`
	DAS           *types.ClusterDasVmConfigInfo     `json:"das,omitempty"`
	Orchestration *types.ClusterVmOrchestrationInfo `json:"orchestration,omitempty"`
}

type infoResult struct {
	Overrides map[string]*Override `json:"overrides"`
}

func (r *infoResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)

	for _, entry := range r.Overrides {
		behavior := fmt.Sprintf("Default (%s)", types.DrsBehaviorFullyAutomated)
		if entry.DRS != nil {
			if *entry.DRS.Enabled {
				behavior = string(entry.DRS.Behavior)
			}
		}

		priority := fmt.Sprintf("Default (%s)", types.DasVmPriorityMedium)
		if entry.DAS != nil {
			priority = entry.DAS.DasSettings.RestartPriority
		}

		ready := "Default (Resources allocated)"
		additionalDelay := 0
		if entry.Orchestration != nil {
			r := entry.Orchestration.VmReadiness
			if r.ReadyCondition != string(types.ClusterVmReadinessReadyConditionUseClusterDefault) {
				ready = strings.Title(r.ReadyCondition)
			}
			additionalDelay = int(r.PostReadyDelay)
		}

		fmt.Fprintf(tw, "Name:\t%s\n", entry.Name)
		fmt.Fprintf(tw, "  DRS Automation Level:\t%s\n", strings.Title(behavior))
		fmt.Fprintf(tw, "  HA Restart Priority:\t%s\n", strings.Title(priority))
		fmt.Fprintf(tw, "  HA Ready Condition:\t%s\n", strings.Title(ready))
		fmt.Fprintf(tw, "  HA Additional Delay:\t%s\n", time.Duration(additionalDelay)*time.Second)
		fmt.Fprintf(tw, "  Host:\t%s\n", entry.Host)
	}

	return tw.Flush()
}

func (r *infoResult) entry(id types.ManagedObjectReference) *Override {
	key := id.String()
	vm, ok := r.Overrides[key]
	if !ok {
		r.Overrides[key] = &Override{id: id}
		vm = r.Overrides[key]
	}
	return vm
}

func (cmd *info) Run(ctx context.Context, f *flag.FlagSet) error {
	cluster, err := cmd.Cluster()
	if err != nil {
		return err
	}

	config, err := cluster.Configuration(ctx)
	if err != nil {
		return err
	}

	res := &infoResult{
		Overrides: make(map[string]*Override),
	}

	for i := range config.DasVmConfig {
		vm := res.entry(config.DasVmConfig[i].Key)

		vm.DAS = &config.DasVmConfig[i]
	}

	for i := range config.DrsVmConfig {
		vm := res.entry(config.DrsVmConfig[i].Key)

		vm.DRS = &config.DrsVmConfig[i]
	}

	for i := range config.VmOrchestration {
		vm := res.entry(config.VmOrchestration[i].Vm)

		vm.Orchestration = &config.VmOrchestration[i]
	}

	for _, o := range res.Overrides {
		// TODO: can optimize to reduce round trips
		vm := object.NewVirtualMachine(cluster.Client(), o.id)
		o.Name, _ = vm.ObjectName(ctx)
		if h, herr := vm.HostSystem(ctx); herr == nil {
			o.Host, _ = h.ObjectName(ctx)
		}
	}

	return cmd.WriteResult(res)
}
