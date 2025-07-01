// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package service

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

type ls struct {
	*flags.ClientFlag
	*flags.OutputFlag
	*flags.HostSystemFlag
}

func init() {
	cli.Register("host.service.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)

	cmd.HostSystemFlag, ctx = flags.NewHostSystemFlag(ctx)
	cmd.HostSystemFlag.Register(ctx, f)
}

func (cmd *ls) Process(ctx context.Context) error {
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

func (cmd *ls) Description() string {
	return `List HOST services.`
}

func Status(s types.HostService) string {
	if s.Running {
		return "Running"
	}
	return "Stopped"
}

func Policy(s types.HostService) string {
	switch types.HostServicePolicy(s.Policy) {
	case types.HostServicePolicyOff:
		return "Disabled"
	case types.HostServicePolicyOn:
		return "Enabled"
	case types.HostServicePolicyAutomatic:
		return "Automatic"
	default:
		return s.Policy
	}
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {
	host, err := cmd.HostSystem()
	if err != nil {
		return err
	}

	s, err := host.ConfigManager().ServiceSystem(ctx)
	if err != nil {
		return err
	}

	services, err := s.Service(ctx)
	if err != nil {
		return err
	}

	return cmd.WriteResult(optionResult(services))
}

type optionResult []types.HostService

func (services optionResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)

	fmt.Fprintf(tw, "%s\t%s\t%v\t%s\n", "Key", "Policy", "Status", "Label")

	for _, s := range services {
		fmt.Fprintf(tw, "%s\t%s\t%s\t%s\n", s.Key, s.Policy, Status(s), s.Label)
	}

	return tw.Flush()
}
