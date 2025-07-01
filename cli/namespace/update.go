// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package namespace

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/vapi/namespace"
)

type update struct {
	*namespaceFlag

	spec namespace.NamespacesInstanceUpdateSpec
}

func init() {
	cli.Register("namespace.update", &update{})
}

func (cmd *update) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.namespaceFlag = &namespaceFlag{}
	cmd.namespaceFlag.Register(ctx, f)
}

func (cmd *update) Process(ctx context.Context) error {
	if err := cmd.namespaceFlag.Process(ctx); err != nil {
		return err
	}

	cmd.spec.StorageSpecs = cmd.storageSpec()
	cmd.spec.VmServiceSpec = cmd.vmServiceSpec()

	return nil
}

func (cmd *update) Usage() string {
	return "NAME"
}

func (cmd *update) Description() string {
	return `Modifies an existing vSphere Namespace on a Supervisor.

Examples:
  govc namespace.update -library vmsvc test-namespace
  govc namespace.update -library vmsvc -library tkgs -storage wcp-policy test-namespace
  govc namespace.update -vmclass best-effort-2xlarge test-namespace
  govc namespace.update -vmclass best-effort-2xlarge -vmclass best-effort-4xlarge test-namespace
  govc namespace.update -library vmsvc -library tkgs -vmclass best-effort-2xlarge -vmclass best-effort-4xlarge test-namespace`
}

func (cmd *update) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}

	rc, err := cmd.RestClient()
	if err != nil {
		return err
	}

	nm := namespace.NewManager(rc)

	return nm.UpdateNamespace(ctx, f.Arg(0), cmd.spec)
}
