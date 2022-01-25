/*
Copyright (c) 2022 VMware, Inc. All Rights Reserved.

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

package staged

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vapi/appliance/update/staged"
)

type delete struct {
	*flags.ClientFlag
	*flags.OutputFlag
}

func init() {
	cli.Register("vcsa.update.staged.delete", &delete{})
}

func (cmd *delete) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)
}

func (cmd *delete) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *delete) Description() string {
	return `Deletes the staged update
Examples:
  govc vcsa.update.staged.delete`
}

func (cmd *delete) Run(ctx context.Context, f *flag.FlagSet) error {

	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	m := staged.NewManager(c)

	err = m.Delete(ctx)
	if err != nil {
		return err
	}

	return nil
}
