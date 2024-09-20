/*
Copyright (c) 2024-2024 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package kms

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
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

	req, err := m.KmsProviderExportRequest(ctx, export.Location)
	if err != nil {
		return err
	}

	return rc.DownloadAttachment(ctx, req, cmd.file)
}
