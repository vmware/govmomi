// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package trust

import (
	"context"
	"flag"
	"io"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vapi/library"
)

type info struct {
	*flags.ClientFlag
	*flags.OutputFlag
}

func init() {
	cli.Register("library.trust.info", &info{})
}

func (cmd *info) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
	cmd.OutputFlag.Register(ctx, f)
}

func (cmd *info) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *info) Usage() string {
	return "ID"
}

func (cmd *info) Description() string {
	return `Display trusted certificate info.

Examples:
  govc library.trust.info vmware_signed`
}

type infoResultsWriter struct {
	TrustedCertificateInfo *library.TrustedCertificate `json:"info,omitempty"`
}

func (r infoResultsWriter) Write(w io.Writer) error {
	var info object.HostCertificateInfo
	_, err := info.FromPEM([]byte(r.TrustedCertificateInfo.Text))
	if err != nil {
		return err
	}
	return info.Write(w)
}

func (r infoResultsWriter) Dump() any {
	return r.TrustedCertificateInfo
}

func (cmd *info) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}

	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	cert, err := library.NewManager(c).GetTrustedCertificate(ctx, f.Arg(0))
	if err != nil {
		return err
	}
	return cmd.WriteResult(&infoResultsWriter{cert})
}
