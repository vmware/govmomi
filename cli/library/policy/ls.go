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
	"github.com/vmware/govmomi/vapi/library"
)

type ls struct {
	*flags.ClientFlag
	*flags.OutputFlag
}

func init() {
	cli.Register("library.policy.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
	cmd.OutputFlag.Register(ctx, f)
}

func (cmd *ls) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *ls) Description() string {
	return `List security policies for content libraries.

Examples:
  govc library.policy.ls
`
}

type lsResultsWriter struct {
	Policies []library.ContentSecurityPoliciesInfo `json:"policies,omitempty"`
}

func (r lsResultsWriter) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)

	if len(r.Policies) == 0 {
		_, _ = fmt.Fprintln(tw, "No Policies found")
		return tw.Flush()
	}

	for _, p := range r.Policies {
		if _, err := fmt.Fprintf(tw, "Name:\t%s\n", p.Name); err != nil {
			return err
		}
		if _, err := fmt.Fprintf(tw, "Policy ID:\t%s\n", p.Policy); err != nil {
			return err
		}
		if err := writeItemRules(tw, p); err != nil {
			return err
		}
	}

	return tw.Flush()
}

func (r lsResultsWriter) Dump() any {
	return r.Policies
}

func writeItemRules(w io.Writer, policy library.ContentSecurityPoliciesInfo) error {
	tw := tabwriter.NewWriter(w, 2, 0, 4, ' ', 0)
	if _, err := fmt.Fprintln(w, "Rules:"); err != nil {
		return err
	}
	for k, v := range policy.ItemTypeRules {
		if _, err := fmt.Fprintf(tw, "\tItem: %s\n", k); err != nil {
			return err
		}
		if _, err := fmt.Fprintf(tw, "\tRule: %s\n", v); err != nil {
			return err
		}
	}
	return tw.Flush()
}

func (cmd *ls) Run(ctx context.Context, _ *flag.FlagSet) error {
	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	m := library.NewManager(c)
	policies, err := m.ListSecurityPolicies(ctx)
	if err != nil {
		return err
	}
	return cmd.WriteResult(&lsResultsWriter{policies})
}
