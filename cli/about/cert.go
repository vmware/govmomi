// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package about

import (
	"context"
	"encoding/pem"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/soap"
)

type cert struct {
	*flags.ClientFlag
	*flags.OutputFlag

	show       bool
	thumbprint bool
}

func init() {
	cli.Register("about.cert", &cert{})
}

func (cmd *cert) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)

	f.BoolVar(&cmd.show, "show", false, "Show PEM encoded server certificate only")
	f.BoolVar(&cmd.thumbprint, "thumbprint", false, "Output host hash and thumbprint only")
}

func (cmd *cert) Description() string {
	return `Display TLS certificate info for HOST.

If the HOST certificate cannot be verified, about.cert will return with exit code 60 (as curl does).
If the '-k' flag is provided, about.cert will return with exit code 0 in this case.
The SHA1 thumbprint can also be used as '-thumbprint' for the 'host.add' and 'cluster.add' commands.

Examples:
  govc about.cert -k -json | jq -r .thumbprintSHA1
  govc about.cert -k -show | sudo tee /usr/local/share/ca-certificates/host.crt
  govc about.cert -k -thumbprint | tee -a ~/.govmomi/known_hosts`
}

func (cmd *cert) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

type certResult struct {
	cmd  *cert
	info object.HostCertificateInfo
}

func (r *certResult) Write(w io.Writer) error {
	if r.cmd.show {
		return pem.Encode(w, &pem.Block{Type: "CERTIFICATE", Bytes: r.info.Certificate.Raw})
	}

	if r.cmd.thumbprint {
		u := r.cmd.Session.URL
		_, err := fmt.Fprintf(w, "%s %s\n", u.Host, r.info.ThumbprintSHA256)
		return err
	}

	return r.cmd.WriteResult(&r.info)
}

func (cmd *cert) Run(ctx context.Context, f *flag.FlagSet) error {
	u := cmd.Session.URL
	c := soap.NewClient(u, false)
	t := c.Client.Transport.(*http.Transport)
	r := certResult{cmd: cmd}

	if err := cmd.SetRootCAs(c); err != nil {
		return err
	}

	if err := r.info.FromURL(u, t.TLSClientConfig); err != nil {
		return err
	}

	if r.info.Err != nil && !r.cmd.Session.Insecure {
		cmd.Out = os.Stderr
		// using same exit code as curl:
		defer os.Exit(60)
	}

	return r.Write(cmd.Out)
}
