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

package key

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/crypto"
)

type rm struct {
	*flags.ClientFlag

	provider string
	force    bool
}

func init() {
	cli.Register("kms.key.rm", &rm{}, true)
}

func (cmd *rm) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.StringVar(&cmd.provider, "p", "", "Provider ID")
	f.BoolVar(&cmd.force, "f", false, "Force")
}

func (cmd *rm) Usage() string {
	return "ID..."
}

func (cmd *rm) Description() string {
	return `Remove crypto keys.

Examples:
  govc kms.key.rm -p my-kp ID`
}

func (cmd *rm) Run(ctx context.Context, f *flag.FlagSet) error {
	n := f.NArg()
	if n == 0 {
		return flag.ErrHelp
	}

	c, err := cmd.Client()
	if err != nil {
		return err
	}

	m, err := crypto.GetManagerKmip(c)
	if err != nil {
		return err
	}

	ids := argsToKeys(cmd.provider, f.Args())

	return m.RemoveKeys(ctx, ids, cmd.force)
}
