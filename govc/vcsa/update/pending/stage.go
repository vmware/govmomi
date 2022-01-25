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

package pending

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vapi/appliance/update/pending"
)

type stage struct {
	*flags.ClientFlag
}

func init() {
	cli.Register("vcsa.update.pending.stage", &stage{})
}

func (cmd *stage) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

}

func (cmd *stage) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}

	return nil
}

func (cmd *stage) Description() string {
	return `Starts staging the appliance update. 
The updates are searched for in the following order: staged, CDROM, URL.
Examples:
  govc vcsa.update.pending.stage 7.0.3.00000`
}

func (cmd *stage) Usage() string {
	return "VERSION"
}

func (cmd *stage) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}
	version := f.Arg(0)

	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	m := pending.NewManager(c)

	err = m.Stage(ctx, version)
	if err != nil {
		return err
	}

	return nil
}
