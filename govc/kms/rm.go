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

	"github.com/vmware/govmomi/crypto"
	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
)

type rm struct {
	*flags.ClientFlag

	server string
}

func init() {
	cli.Register("kms.rm", &rm{})
}

func (cmd *rm) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.StringVar(&cmd.server, "s", "", "Server name")
}

func (cmd *rm) Usage() string {
	return "NAME"
}

func (cmd *rm) Description() string {
	return `Remove KMS server or cluster.

Examples:
  govc kms.rm my-kp
  govc kms.rm -s my-server my-kp`
}

func (cmd *rm) Run(ctx context.Context, f *flag.FlagSet) error {
	id := f.Arg(0)
	if id == "" {
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

	if cmd.server != "" {
		return m.RemoveKmipServer(ctx, id, cmd.server)
	}

	return m.UnregisterKmsCluster(ctx, id)
}
