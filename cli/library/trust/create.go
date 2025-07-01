// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package trust

import (
	"bytes"
	"context"
	"flag"
	"io"
	"os"
	"path/filepath"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/library"
)

type create struct {
	*flags.ClientFlag
}

func init() {
	cli.Register("library.trust.create", &create{})
}

func (cmd *create) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
}

func (cmd *create) Usage() string {
	return "FILE"
}

func (cmd *create) Description() string {
	return `Add a certificate to content library trust store.

If FILE name is "-", read certificate from stdin.

Examples:
  govc library.trust.create cert.pem
  govc about.cert -show -u wp-content-int.vmware.com | govc library.trust.create -`
}

func (cmd *create) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	var cert string

	name := f.Arg(0)
	if name == "-" || name == "" {
		var buf bytes.Buffer
		if _, err := io.Copy(&buf, os.Stdin); err != nil {
			return err
		}
		cert = buf.String()
	} else {
		b, err := os.ReadFile(filepath.Clean(name))
		if err != nil {
			return err
		}
		cert = string(b)
	}

	return library.NewManager(c).CreateTrustedCertificate(ctx, cert)
}
