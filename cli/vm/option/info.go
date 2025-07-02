// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package option

import (
	"context"
	"flag"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vim25/types"
)

type info struct {
	flags.EnvBrowser

	key string
}

func init() {
	cli.Register("vm.option.info", &info{})
}

func (cmd *info) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.EnvBrowser.Register(ctx, f)

	f.StringVar(&cmd.key, "id", "", "Option descriptor key")
}

func (cmd *info) Usage() string {
	return "[GUEST_ID]..."
}

func (cmd *info) Description() string {
	return `VM config options for CLUSTER.

The config option data contains information about the execution environment for a VM
in the given CLUSTER, and optionally for a specific HOST.

By default, supported guest OS IDs and full name are listed.

Examples:
  govc vm.option.info -cluster C0
  govc vm.option.info -cluster C0 -dump ubuntu64Guest
  govc vm.option.info -cluster C0 -json | jq .guestOSDescriptor[].id
  govc vm.option.info -host my_hostname
  govc vm.option.info -vm my_vm`
}

func (cmd *info) Run(ctx context.Context, f *flag.FlagSet) error {
	b, err := cmd.Browser(ctx)
	if err != nil {
		return err
	}

	host, err := cmd.HostSystemIfSpecified()
	if err != nil {
		return err
	}

	var req *types.EnvironmentBrowserConfigOptionQuerySpec

	spec := func() *types.EnvironmentBrowserConfigOptionQuerySpec {
		if req == nil {
			req = new(types.EnvironmentBrowserConfigOptionQuerySpec)
		}
		return req
	}

	if f.NArg() != 0 {
		spec().GuestId = f.Args()
	}

	if host != nil {
		spec().Host = types.NewReference(host.Reference())
	}

	if cmd.key != "" {
		spec().Key = cmd.key
	}

	opt, err := b.QueryConfigOption(ctx, req)
	if err != nil {
		return err
	}

	return cmd.VirtualMachineFlag.WriteResult(&infoResult{opt})
}

type infoResult struct {
	*types.VirtualMachineConfigOption
}

func (r *infoResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)

	for _, d := range r.GuestOSDescriptor {
		_, _ = fmt.Fprintf(tw, "%s\t%s\n", d.Id, d.FullName)
	}

	return tw.Flush()
}

func (r *infoResult) Dump() any {
	return r.VirtualMachineConfigOption
}
