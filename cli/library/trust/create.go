/*
Copyright (c) 2022-2022 VMware, Inc. All Rights Reserved.

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
