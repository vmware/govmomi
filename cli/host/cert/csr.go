// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package cert

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
)

type csr struct {
	*flags.HostSystemFlag

	ip bool
}

func init() {
	cli.Register("host.cert.csr", &csr{})
}

func (cmd *csr) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.HostSystemFlag, ctx = flags.NewHostSystemFlag(ctx)
	cmd.HostSystemFlag.Register(ctx, f)

	f.BoolVar(&cmd.ip, "ip", false, "Use IP address as CN")
}

func (cmd *csr) Description() string {
	return `Generate a certificate-signing request (CSR) for HOST.`
}

func (cmd *csr) Process(ctx context.Context) error {
	if err := cmd.HostSystemFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *csr) Run(ctx context.Context, f *flag.FlagSet) error {
	host, err := cmd.HostSystem()
	if err != nil {
		return err
	}

	m, err := host.ConfigManager().CertificateManager(ctx)
	if err != nil {
		return err
	}

	output, err := m.GenerateCertificateSigningRequest(ctx, cmd.ip)
	if err != nil {
		return err
	}

	_, err = fmt.Println(output)
	return err
}
