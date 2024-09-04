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

package namespace

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
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
