// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package cert

import (
	"bytes"
	"context"
	"flag"
	"io"
	"os"
	"path/filepath"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
)

type install struct {
	*flags.HostSystemFlag
}

func init() {
	cli.Register("host.cert.import", &install{})
}

func (cmd *install) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.HostSystemFlag, ctx = flags.NewHostSystemFlag(ctx)
	cmd.HostSystemFlag.Register(ctx, f)
}

func (cmd *install) Usage() string {
	return "FILE"
}

func (cmd *install) Description() string {
	return `Install SSL certificate FILE on HOST.

If FILE name is "-", read certificate from stdin.`
}

func (cmd *install) Process(ctx context.Context) error {
	if err := cmd.HostSystemFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *install) Run(ctx context.Context, f *flag.FlagSet) error {
	host, err := cmd.HostSystem()
	if err != nil {
		return err
	}

	m, err := host.ConfigManager().CertificateManager(ctx)
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

	return m.InstallServerCertificate(ctx, cert)
}
