// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package esxcli

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"strings"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/esx"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/internal"
)

type model struct {
	*flags.HostSystemFlag

	types     bool
	commands  bool
	instances bool
}

func init() {
	cli.Register("host.esxcli.model", &model{}, true)
}

func (cmd *model) Usage() string {
	return "[NAMESPACE]..."
}

func (cmd *model) Description() string {
	return `Print esxcli model for HOST.

Examples:
  govc host.esxcli.model # prints type model supported by vcsim
  govc host.esxcli.model network.vm network.ip.connection # specific namespaces
  govc host.esxcli.model -dump # generate simulator/esx/type_info.go
  govc host.esxcli.model -c # prints commands supported by vcsim
  govc host.esxcli.model -c network.vm`
}

func (cmd *model) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.HostSystemFlag, ctx = flags.NewHostSystemFlag(ctx)
	cmd.HostSystemFlag.Register(ctx, f)

	f.BoolVar(&cmd.types, "t", true, "Print type info")
	f.BoolVar(&cmd.commands, "c", false, "Print command info")
	f.BoolVar(&cmd.instances, "i", false, "Print instances")
}

var namespaces = []string{
	"hardware.clock",
	"hardware.platform",
	"iscsi.software",
	"network.firewall",
	"network.ip.connection",
	"network.nic.ring.current",
	"network.nic.ring.preset",
	"network.vm",
	"software.vib",
	"system.hostname",
	"system.settings.advanced",
	"system.stats.uptime",
	"vm.process",
}

type modelInfo struct {
	CommandInfo []esx.CommandInfo
	TypeInfo    internal.DynamicTypeMgrAllTypeInfo
	Instances   []internal.DynamicTypeMgrMoInstance
}

func (r modelInfo) Dump() any {
	if len(r.CommandInfo) != 0 {
		return r.CommandInfo
	}
	if len(r.Instances) != 0 {
		return r.Instances
	}
	return r.TypeInfo
}

func (r modelInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.Dump())
}

func (r modelInfo) Write(w io.Writer) error {
	for _, item := range r.CommandInfo {
		ns := strings.ReplaceAll(item.Name, ".", " ")

		for _, m := range item.Method {
			tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)
			fmt.Fprintf(tw, "%s %s\n", ns, m.Name)

			for _, p := range m.Param {
				fmt.Fprintf(tw, "  %s\t%s\n", p.Name, p.Aliases)
			}

			_ = tw.Flush()
		}
	}

	for _, obj := range r.TypeInfo.ManagedTypeInfo {
		for _, m := range obj.Method {
			if m.ReturnTypeInfo == nil {
				continue
			}

			var param []string
			for _, p := range m.ParamTypeInfo {
				param = append(param, p.Name+" "+p.Type)
			}

			fmt.Fprintf(w, "func %s.%s(%s) %s\n",
				obj.Name, m.Name,
				strings.Join(param, ", "),
				m.ReturnTypeInfo.Type)
		}
	}

	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)
	for _, obj := range r.TypeInfo.DataTypeInfo {
		fmt.Fprintf(tw, "type %s struct = {\n", obj.Name)
		for _, p := range obj.Property {
			fmt.Fprintf(tw, "    %s\t%s\n", p.Name, p.Type)
		}
		fmt.Fprintln(tw, "}")
	}
	_ = tw.Flush()

	return nil
}

func (cmd *model) typeInfo(ctx context.Context, e *esx.Executor, args []string) (internal.DynamicTypeMgrAllTypeInfo, error) {
	var info internal.DynamicTypeMgrAllTypeInfo

	for _, ns := range args {
		req := internal.DynamicTypeMgrQueryTypeInfoRequest{
			This: e.DynamicTypeManager(),
			FilterSpec: &internal.DynamicTypeMgrTypeFilterSpec{
				TypeSubstr: ns,
			},
		}

		res, err := internal.DynamicTypeMgrQueryTypeInfo(ctx, e.Client(), &req)
		if err != nil {
			return info, err
		}

		info.DataTypeInfo = append(info.DataTypeInfo, res.Returnval.DataTypeInfo...)
		info.EnumTypeInfo = append(info.EnumTypeInfo, res.Returnval.EnumTypeInfo...)
		info.ManagedTypeInfo = append(info.ManagedTypeInfo, res.Returnval.ManagedTypeInfo...)
	}

	return info, nil
}

func (cmd *model) commandInfo(ctx context.Context, e *esx.Executor, args []string) ([]esx.CommandInfo, error) {
	var info []esx.CommandInfo

	for _, ns := range args {
		c, err := e.CommandInfo(ctx, ns)
		if err != nil {
			return nil, err
		}

		info = append(info, *c)
	}

	return info, nil
}

func (cmd *model) instanceInfo(ctx context.Context, e *esx.Executor, args []string) ([]internal.DynamicTypeMgrMoInstance, error) {
	var info []internal.DynamicTypeMgrMoInstance

	req := internal.DynamicTypeMgrQueryMoInstancesRequest{
		This: e.DynamicTypeManager(),
	}

	res, err := internal.DynamicTypeMgrQueryMoInstances(ctx, e.Client(), &req)
	if err != nil {
		return nil, err
	}

	for _, i := range res.Returnval {
		for _, ns := range args {
			if !strings.HasPrefix(ns, "vim.") {
				ns = "vim.EsxCLI." + ns
			}

			if ns == i.MoType {
				info = append(info, i)
				break
			}
		}
	}

	return info, nil
}

func (cmd *model) Run(ctx context.Context, f *flag.FlagSet) error {
	args := f.Args()
	if len(args) == 0 {
		args = namespaces
	}

	c, err := cmd.Client()
	if err != nil {
		return err
	}

	host, err := cmd.HostSystem()
	if err != nil {
		return err
	}

	e, err := esx.NewExecutor(ctx, c, host)
	if err != nil {
		return err
	}

	if cmd.commands || cmd.instances {
		cmd.types = false
	}

	var res modelInfo

	if cmd.commands {
		res.CommandInfo, err = cmd.commandInfo(ctx, e, args)
		if err != nil {
			return err
		}
	}

	if cmd.types {
		res.TypeInfo, err = cmd.typeInfo(ctx, e, args)
		if err != nil {
			return err
		}
	}

	if cmd.instances {
		res.Instances, err = cmd.instanceInfo(ctx, e, args)
		if err != nil {
			return err
		}
		if !cmd.All() {
			cmd.JSON = true
		}
	}

	return cmd.WriteResult(&res)
}
