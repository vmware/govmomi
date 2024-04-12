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
	*flags.ClientFlag

	libraries flags.StringList
	vmClasses flags.StringList
	spec      namespace.NamespacesInstanceCreateSpec
}

func init() {
	cli.Register("namespace.create", &create{})
}

func (cmd *create) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.StringVar(&cmd.spec.Cluster, "supervisor", "", "The identifier of the Supervisor.")
	f.Var(&cmd.libraries, "library", "Content library IDs to associate with the vSphere Namespace.")
	f.Var(&cmd.vmClasses, "vm-class", "Virtual machine class IDs to associate with the vSphere Namespace.")
}

func (*create) Usage() string {
	return "NAME"
}

func (cmd *create) Process(ctx context.Context) error {
	cmd.spec.VmServiceSpec.ContentLibraries = cmd.libraries
	cmd.spec.VmServiceSpec.VmClasses = cmd.vmClasses

	return cmd.ClientFlag.Process(ctx)
}

func (cmd *create) Description() string {
	return `Creates a new vSphere Namespace on a Supervisor.

The '-library' and '-vm-class' flags can each be specified multiple times.

Examples:
  govc namespace.create -supervisor=domain-c1 test-namespace
  govc namespace.create -supervisor=domain-c1 -library=dca9cc16-9460-4da0-802c-4aa148ac6cf7 test-namespace
  govc namespace.create -supervisor=domain-c1 -library=dca9cc16-9460-4da0-802c-4aa148ac6cf7 -library=dca9cc16-9460-4da0-802c-4aa148ac6cf7 test-namespace
  govc namespace.create -supervisor=domain-c1 -vm-class=best-effort-2xlarge test-namespace
  govc namespace.create -supervisor=domain-c1 -vm-class=best-effort-2xlarge -vm-class best-effort-4xlarge test-namespace
  govc namespace.create -supervisor=domain-c1 -library=dca9cc16-9460-4da0-802c-4aa148ac6cf7 -library=dca9cc16-9460-4da0-802c-4aa148ac6cf7 -vm-class=best-effort-2xlarge -vm-class=best-effort-4xlarge test-namespace`
}

func (cmd *create) Run(ctx context.Context, f *flag.FlagSet) error {
	cmd.spec.Namespace = f.Arg(0)
	if f.NArg() != 1 {
		return flag.ErrHelp
	}

	rc, err := cmd.RestClient()
	if err != nil {
		return err
	}

	nm := namespace.NewManager(rc)

	return nm.CreateNamespace(ctx, cmd.spec)
}
