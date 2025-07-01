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
	"github.com/vmware/govmomi/vim25/mo"
)

type info struct {
	*flags.HostSystemFlag
	*flags.OutputFlag

	show bool
}

func init() {
	cli.Register("host.cert.info", &info{})
}

func (cmd *info) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.HostSystemFlag, ctx = flags.NewHostSystemFlag(ctx)
	cmd.HostSystemFlag.Register(ctx, f)

	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)

	f.BoolVar(&cmd.show, "show", false, "Show PEM encoded server certificate only")
}

func (cmd *info) Description() string {
	return `Display SSL certificate info for HOST.`
}

func (cmd *info) Process(ctx context.Context) error {
	if err := cmd.HostSystemFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *info) Run(ctx context.Context, f *flag.FlagSet) error {
	host, err := cmd.HostSystem()
	if err != nil {
		return err
	}

	if cmd.show {
		var props mo.HostSystem
		err = host.Properties(ctx, host.Reference(), []string{"config.certificate"}, &props)
		if err != nil {
			return err
		}
		_, err = fmt.Fprint(cmd.Out, string(props.Config.Certificate))
		return err
	}

	m, err := host.ConfigManager().CertificateManager(ctx)
	if err != nil {
		return err
	}

	info, err := m.CertificateInfo(ctx)
	if err != nil {
		return err
	}

	return cmd.WriteResult(info)
}
