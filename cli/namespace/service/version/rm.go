// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package version

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/namespace"
)

type rm struct {
	*flags.ClientFlag
}

func init() {
	cli.Register("namespace.service.version.rm", &rm{})
}

func (cmd *rm) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
}

func (cmd *rm) Description() string {
	return `Removes a vSphere Supervisor Service version.
Note that a service version must be deactivated before being deleted.

Examples:
  govc namespace.service.version.rm my-supervisor-service 1.0.0`
}

func (cmd *rm) Usage() string {
	return "NAME VERSION"
}

func (cmd *rm) Run(ctx context.Context, f *flag.FlagSet) error {
	service := f.Arg(0)
	if len(service) == 0 {
		return flag.ErrHelp
	}
	version := f.Arg(1)
	if len(version) == 0 {
		return flag.ErrHelp
	}

	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	m := namespace.NewManager(c)
	return m.RemoveSupervisorServiceVersion(ctx, service, version)
}
