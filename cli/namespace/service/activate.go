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

type activate struct {
	*flags.ClientFlag
}

func init() {
	cli.Register("namespace.service.activate", &activate{})
}

func (cmd *activate) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
}

func (cmd *activate) Description() string {
	return `Activates a vSphere Supervisor Service (and all its versions).

Examples:
  govc namespace.service.activate my-supervisor-service other-supervisor-service`
}

func (cmd *activate) Usage() string {
	return "NAME..."
}

func (cmd *activate) Run(ctx context.Context, f *flag.FlagSet) error {
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
		if err := m.ActivateSupervisorServices(ctx, svc); err != nil {
			return err
		}
	}

	return nil
}
