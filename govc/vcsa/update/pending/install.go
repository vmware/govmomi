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
	"strings"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vapi/appliance/update/pending"
)

type install struct {
	*flags.ClientFlag
	*flags.OutputFlag
}

func init() {
	cli.Register("vcsa.update.pending.install", &install{})
}

func (cmd *install) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)
}

func (cmd *install) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *install) Description() string {
	return `Validates the user provided data before the update installation.

Examples:
  govc vcsa.update.pending.install 7.0.3.00000 "key1=val1,key2=val2"`
}

func (cmd *install) Usage() string {
	return "[VERSION] [USERDATA]"
}

func (cmd *install) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 2 {
		return flag.ErrHelp
	}

	version := f.Arg(0)
	userdata := make(map[string]string)

	for _, inputs := range strings.Split(f.Arg(1), ",") {
		input := strings.Split(inputs, "=")
		userdata[input[0]] = input[1]
	}

	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	m := pending.NewManager(c)

	err = m.Install(ctx, version, userdata)
	if err != nil {
		return err
	}

	return nil
}
