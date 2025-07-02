// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package fields

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/object"
)

type add struct {
	*flags.ClientFlag
	kind string
}

func init() {
	cli.Register("fields.add", &add{})
}

func (cmd *add) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.StringVar(&cmd.kind, "type", "", "Managed object type")
}

func (cmd *add) Usage() string {
	return "NAME"
}

func (cmd *add) Description() string {
	return `Add a custom field type with NAME.

Examples:
  govc fields.add my-field-name # adds a field to all managed object types
  govc fields.add -type VirtualMachine my-vm-field-name # adds a field to the VirtualMachine type`
}

func (cmd *add) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}

	c, err := cmd.Client()
	if err != nil {
		return err
	}

	m, err := object.GetCustomFieldsManager(c)
	if err != nil {
		return err
	}

	name := f.Arg(0)

	def, err := m.Add(ctx, name, cmd.kind, nil, nil)
	if err != nil {
		return err
	}

	fmt.Printf("%d\n", def.Key)

	return nil
}
