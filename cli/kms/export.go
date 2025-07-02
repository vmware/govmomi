// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package kms

import (
	"context"
	"flag"
	"fmt"
	"net/url"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/crypto"
)

type export struct {
	*flags.ClientFlag

	spec crypto.KmsProviderExportSpec
	file string
}

func init() {
	cli.Register("kms.export", &export{})
}

func (cmd *export) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.StringVar(&cmd.file, "f", "", "File name")
	f.StringVar(&cmd.spec.Password, "p", "", "Password")
}

func (cmd *export) Usage() string {
	return "NAME"
}

func (cmd *export) Description() string {
	return `Export KMS cluster for backup.

Examples:
  govc kms.export my-kp
  govc kms.export -f my-backup.p12 my-kp`
}

func (cmd *export) Run(ctx context.Context, f *flag.FlagSet) error {
	id := f.Arg(0)
	if id == "" {
		return flag.ErrHelp
	}

	rc, err := cmd.RestClient()
	if err != nil {
		return err
	}

	m := crypto.NewManager(rc)

	cmd.spec.Provider = id
	export, err := m.KmsProviderExport(ctx, cmd.spec)
	if err != nil {
		return err
	}

	if export.Type != "LOCATION" {
		return fmt.Errorf("unsupported export type: %s", export.Type)
	}

	// Rewrite URL to use the host we connected to vCenter with
	u, err := url.Parse(export.Location.URL)
	if err != nil {
		return err
	}

	u.Host = rc.URL().Host
	export.Location.URL = u.String()

	req, err := m.KmsProviderExportRequest(ctx, export.Location)
	if err != nil {
		return err
	}

	return rc.DownloadAttachment(ctx, req, cmd.file)
}
