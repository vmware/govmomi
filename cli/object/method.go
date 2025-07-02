// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package object

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/object"
)

type method struct {
	*flags.DatacenterFlag

	name   string
	reason string
	source string
	enable bool
}

func init() {
	cli.Register("object.method", &method{})
}

func (cmd *method) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatacenterFlag, ctx = flags.NewDatacenterFlag(ctx)
	cmd.DatacenterFlag.Register(ctx, f)

	f.StringVar(&cmd.name, "name", "", "Method name")
	f.StringVar(&cmd.reason, "reason", "", "Reason for disabling method")
	f.StringVar(&cmd.source, "source", "govc", "Source ID")
	f.BoolVar(&cmd.enable, "enable", true, "Enable method")
}

func (cmd *method) Usage() string {
	return "PATH..."
}

func (cmd *method) Description() string {
	return `Enable or disable methods for managed objects.

Examples:
  govc object.method -name Destroy_Task -enable=false /dc1/vm/foo
  govc object.collect /dc1/vm/foo disabledMethod | grep --color Destroy_Task
  govc object.method -name Destroy_Task -enable /dc1/vm/foo`
}

func (cmd *method) Process(ctx context.Context) error {
	if err := cmd.DatacenterFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *method) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() == 0 {
		return flag.ErrHelp
	}

	if cmd.name == "" {
		return flag.ErrHelp
	}

	c, err := cmd.Client()
	if err != nil {
		return err
	}

	objs, err := cmd.ManagedObjects(ctx, f.Args())
	if err != nil {
		return err
	}

	m := object.NewAuthorizationManager(c)

	if cmd.enable {
		return m.EnableMethods(ctx, objs, []string{cmd.name}, cmd.source)
	}

	method := []object.DisabledMethodRequest{
		{
			Method: cmd.name,
			Reason: cmd.reason,
		},
	}

	return m.DisableMethods(ctx, objs, method, cmd.source)
}
