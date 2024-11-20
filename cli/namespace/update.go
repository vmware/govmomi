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
