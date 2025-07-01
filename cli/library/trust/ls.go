// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package trust

import (
	"context"
	"crypto/x509"
	"encoding/pem"
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
	cli.Register("library.trust.ls", &ls{})
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
	return `List trusted certificates for content libraries.

Examples:
  govc library.trust.ls
  govc library.trust.ls -json`
}

type lsResultsWriter struct {
	TrustedCertificates []library.TrustedCertificateSummary `json:"certificates,omitempty"`
}

func (r lsResultsWriter) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)

	for _, cert := range r.TrustedCertificates {
		block, _ := pem.Decode([]byte(cert.Text))
		x, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return err
		}

		x.Subject.Names = nil // trim x.Subject.String() output

		fmt.Fprintf(tw, "%s\t%s\n", cert.ID, x.Subject)
	}

	return tw.Flush()
}

func (r lsResultsWriter) Dump() any {
	return r.TrustedCertificates
}

func (cmd *ls) Run(ctx context.Context, _ *flag.FlagSet) error {
	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	certs, err := library.NewManager(c).ListTrustedCertificates(ctx)
	if err != nil {
		return err
	}

	return cmd.WriteResult(&lsResultsWriter{certs})
}
