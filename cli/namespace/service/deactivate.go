// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package service

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/namespace"
)

type deactivate struct {
	*flags.ClientFlag
}

func init() {
	cli.Register("namespace.service.deactivate", &deactivate{})
}

func (cmd *deactivate) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
}

func (cmd *deactivate) Description() string {
	return `Deactivates a vSphere Supervisor Service (and all its versions).

Examples:
  govc namespace.service.deactivate my-supervisor-service other-supervisor-service`
}

func (cmd *deactivate) Usage() string {
	return "NAME..."
}

func (cmd *deactivate) Run(ctx context.Context, f *flag.FlagSet) error {

	services := f.Args()
	if len(services) < 1 {
		return flag.ErrHelp
	}

	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	m := namespace.NewManager(c)
	for _, svc := range services {
		if err := m.DeactivateSupervisorServices(ctx, svc); err != nil {
			return err
		}
	}

	return nil
}
