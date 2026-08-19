// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package interop

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/lcm"
)

type check struct {
	*flags.ClientFlag
	*flags.OutputFlag

	product    string
	version    string
	identifier string
	wait       bool
}

func init() {
	cli.Register("lcm.interop.check", &check{}, true)
}

func (cmd *check) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)

	f.StringVar(&cmd.product, "product", "com.vmware.vcenter.supervisor", "Product name to validate")
	f.StringVar(&cmd.version, "version", "", "Target version to validate (required)")
	f.StringVar(&cmd.identifier, "identifier", "", "Component instance identifier (optional)")
	f.BoolVar(&cmd.wait, "wait", false, "Wait for task completion and print the report")
}

func (cmd *check) Description() string {
	return `Submit an LCM interoperability validation task.

Prints the task ID on success. Use lcm.interop.result to retrieve the report,
or pass -wait to block until completion and print it immediately.

Examples:
  govc lcm.interop.check -version 8.0.3
  govc lcm.interop.check -product com.vmware.vcenter.supervisor -version 8.0.3 -identifier my-instance
  govc lcm.interop.check -version 8.0.3 -wait`
}

func (cmd *check) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return cmd.OutputFlag.Process(ctx)
}

func (cmd *check) Run(ctx context.Context, f *flag.FlagSet) error {
	if cmd.version == "" {
		return fmt.Errorf("-version is required")
	}

	m := newManager(cmd.ClientFlag)

	component := lcm.InteropComponent{
		ProductName: cmd.product,
		Version:     cmd.version,
		Identifier:  cmd.identifier,
	}

	taskID, err := m.CreateInteropTask(ctx, component)
	if err != nil {
		return err
	}

	if !cmd.wait {
		fmt.Fprintln(cmd.OutputFlag.Out, taskID)
		return nil
	}

	info, err := m.WaitForCompletion(ctx, taskID)
	if err != nil {
		return err
	}

	interopResult, err := lcm.ParseResult(info)
	if err != nil {
		return err
	}

	return cmd.WriteResult(&reportOutput{result: interopResult})
}
