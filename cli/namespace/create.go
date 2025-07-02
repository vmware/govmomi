// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package namespace

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/namespace"
)

type create struct {
	*flags.ClusterFlag
	*namespaceFlag

	spec namespace.NamespacesInstanceCreateSpec
}

func init() {
	cli.Register("namespace.create", &create{})
}

func (cmd *create) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClusterFlag, ctx = flags.NewClusterFlag(ctx)
	cmd.ClusterFlag.Register(ctx, f)

	cmd.namespaceFlag = &namespaceFlag{}
	cmd.namespaceFlag.Register(ctx, f)
}

func (cmd *create) Process(ctx context.Context) error {
	if err := cmd.ClusterFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.namespaceFlag.Process(ctx); err != nil {
		return err
	}

	cluster, err := cmd.Cluster()
	if err != nil {
		return err
	}

	cmd.spec.Cluster = cluster.Reference().Value
	cmd.spec.StorageSpecs = cmd.storageSpec()
	cmd.spec.VmServiceSpec = cmd.vmServiceSpec()

	return nil
}

func (*create) Usage() string {
	return "NAME"
}

func (cmd *create) Description() string {
	return `Creates a new vSphere Namespace on a Supervisor.

The '-library', '-vmclass' and '-storage' flags can each be specified multiple times.

Examples:
  govc namespace.create -cluster C1 test-namespace
  govc namespace.create -cluster C1 -library vmsvc test-namespace
  govc namespace.create -cluster C1 -library vmsvc -library tkgs test-namespace -storage wcp-policy
  govc namespace.create -cluster C1 -vmclass best-effort-2xlarge test-namespace
  govc namespace.create -cluster C1 -vmclass best-effort-2xlarge -vmclass best-effort-4xlarge test-namespace
  govc namespace.create -cluster C1 -library vmsvc -library tkgs -vmclass best-effort-2xlarge -vmclass best-effort-4xlarge test-namespace`
}

func (cmd *create) Run(ctx context.Context, f *flag.FlagSet) error {
	cmd.spec.Namespace = f.Arg(0)
	if f.NArg() != 1 {
		return flag.ErrHelp
	}

	rc, err := cmd.namespaceFlag.RestClient()
	if err != nil {
		return err
	}

	nm := namespace.NewManager(rc)

	return nm.CreateNamespace(ctx, cmd.spec)
}
