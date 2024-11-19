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

package module

import (
	"bufio"
	"context"
	"flag"
	"os"
	"strings"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/cluster"
)

type rm struct {
	*flags.ClientFlag
	ignoreNotFound bool
}

func init() {
	cli.Register("cluster.module.rm", &rm{})
}

func (cmd *rm) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.BoolVar(&cmd.ignoreNotFound, "ignore-not-found", false, "Treat \"404 Not Found\" as a successful delete.")
}

func (cmd *rm) Usage() string {
	return "ID"
}

func (cmd *rm) Description() string {
	return `Delete cluster module ID.

If ID is "-", read a list from stdin.

Examples:
  govc cluster.module.rm module_id
  govc cluster.module.rm - < input-file.txt`
}

func (cmd *rm) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}

	moduleID := f.Arg(0)

	c, err := cmd.RestClient()
	if err != nil {
		return err
	}
	m := cluster.NewManager(c)

	if moduleID == "-" {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			moduleID := scanner.Text()
			if moduleID == "" {
				continue
			}
			if err := cmd.deleteModule(ctx, m, moduleID); err != nil {
				return err
			}
		}
		return nil
	}

	return cmd.deleteModule(ctx, m, moduleID)
}

func (cmd *rm) deleteModule(ctx context.Context, m *cluster.Manager, moduleID string) error {
	if err := m.DeleteModule(ctx, moduleID); err != nil {
		if cmd.ignoreNotFound && strings.HasSuffix(err.Error(), "404 Not Found") {
			return nil
		}
		return err
	}
	return nil
}
