// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package service

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/namespace"
)

type rm struct {
	*flags.ClientFlag
}

func init() {
	cli.Register("namespace.service.rm", &rm{})
}

func (cmd *rm) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
}

func (cmd *rm) Description() string {
	return `Removes a vSphere Supervisor Service.
Note that a service must be deactivated and all versions must be removed before being deleted.

Examples:
  govc namespace.service.rm my-supervisor-service other-supervisor-service`
}

func (cmd *rm) Usage() string {
	return "NAME..."
}

func (cmd *rm) Run(ctx context.Context, f *flag.FlagSet) error {

	services := f.Args()
	if len(services) < 1 {
		return fmt.Errorf("at least one service must be passed as argument")
	}

	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	m := namespace.NewManager(c)
	for _, svc := range services {

		if err := m.RemoveSupervisorService(ctx, svc); err != nil {
			return err
		}
	}
	return nil
}
