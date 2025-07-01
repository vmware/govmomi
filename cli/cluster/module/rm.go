// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

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
