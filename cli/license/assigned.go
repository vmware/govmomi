// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package license

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/license"
	"github.com/vmware/govmomi/vim25/types"
)

type assigned struct {
	*flags.ClientFlag
	*flags.OutputFlag

	id string
}

func init() {
	cli.Register("license.assigned.ls", &assigned{})
}

func (cmd *assigned) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)

	f.StringVar(&cmd.id, "id", "", "Entity ID")
}

func (cmd *assigned) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *assigned) Run(ctx context.Context, f *flag.FlagSet) error {
	client, err := cmd.Client()
	if err != nil {
		return err
	}

	m, err := license.NewManager(client).AssignmentManager(ctx)
	if err != nil {
		return err
	}

	assigned, err := m.QueryAssigned(ctx, cmd.id)
	if err != nil {
		return err
	}

	return cmd.WriteResult(assignedOutput(assigned))
}

type assignedOutput []types.LicenseAssignmentManagerLicenseAssignment

func (res assignedOutput) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(os.Stdout, 4, 0, 2, ' ', 0)
	fmt.Fprintf(tw, "Id:\tScope:\tName:\tLicense:\n")
	for _, v := range res {
		fmt.Fprintf(tw, "%s\t", v.EntityId)
		fmt.Fprintf(tw, "%s\t", v.Scope)
		fmt.Fprintf(tw, "%s\t", v.EntityDisplayName)
		fmt.Fprintf(tw, "%s\t", v.AssignedLicense.LicenseKey)
		fmt.Fprintf(tw, "\n")
	}
	return tw.Flush()
}
