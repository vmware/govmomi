// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package policy

import (
	"context"
	"flag"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/pbm"
	"github.com/vmware/govmomi/pbm/types"
)

type ls struct {
	*flags.ClientFlag
	*flags.OutputFlag

	id bool
}

func init() {
	cli.Register("storage.policy.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)

	f.BoolVar(&cmd.id, "i", false, "List policy ID only")
}

func (cmd *ls) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *ls) Usage() string {
	return "[NAME]"
}

func (cmd *ls) Description() string {
	return `VM Storage Policy listing.

Examples:
  govc storage.policy.ls
  govc storage.policy.ls "vSAN Default Storage Policy"
  govc storage.policy.ls -i "vSAN Default Storage Policy"`
}

func ListProfiles(ctx context.Context, c *pbm.Client, name string) ([]types.BasePbmProfile, error) {
	m, err := c.ProfileMap(ctx)
	if err != nil {
		return nil, err
	}
	if name == "" {
		return m.Profile, nil
	}
	if p, ok := m.Name[name]; ok {
		return []types.BasePbmProfile{p}, nil
	}
	return nil, fmt.Errorf("profile %q not found", name)
}

type lsResult struct {
	Profile []types.BasePbmProfile `json:"profile"`
	cmd     *ls
}

func (r *lsResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(r.cmd.Out, 2, 0, 2, ' ', 0)

	for i := range r.Profile {
		p := r.Profile[i].GetPbmProfile()
		_, _ = fmt.Fprintf(tw, "%s", p.ProfileId.UniqueId)
		if !r.cmd.id {
			_, _ = fmt.Fprintf(tw, "\t%s", p.Name)
		}
		_, _ = fmt.Fprintln(tw)
	}

	return tw.Flush()
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.PbmClient()
	if err != nil {
		return err
	}

	profiles, err := ListProfiles(ctx, c, f.Arg(0))
	if err != nil {
		return err
	}

	return cmd.WriteResult(&lsResult{profiles, cmd})
}
