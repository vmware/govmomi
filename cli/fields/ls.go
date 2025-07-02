// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package fields

import (
	"context"
	"flag"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/object"
)

type ls struct {
	*flags.ClientFlag
	kind string
}

func init() {
	cli.Register("fields.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.StringVar(&cmd.kind, "type", "", "Filter by a Managed Object Type")
}

func (cmd *ls) Description() string {
	return `List custom field definitions.

Examples:
  govc fields.ls
  govc fields.ls -type VirtualMachine`
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.Client()
	if err != nil {
		return err
	}

	m, err := object.GetCustomFieldsManager(c)
	if err != nil {
		return err
	}

	field, err := m.Field(ctx)
	if err != nil {
		return err
	}

	tw := tabwriter.NewWriter(os.Stdout, 3, 0, 2, ' ', 0)

	for _, def := range field {
		if cmd.kind == "" || cmd.kind == def.ManagedObjectType {
			fmt.Fprintf(tw, "%d\t%s\t%s\n", def.Key, def.Name, def.ManagedObjectType)
		}
	}

	return tw.Flush()
}
