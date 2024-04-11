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
	"strings"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vapi/namespace"
)

type update struct {
	*flags.ClientFlag

	libraries string
	vmClasses string
	spec      namespace.NamespacesInstanceUpdateSpec
}

func init() {
	cli.Register("namespace.update", &update{})
}

func (cmd *update) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.StringVar(&cmd.libraries, "content-libraries", "", "The list of content libraries to associate with the vSphere Namespace.")
	f.StringVar(&cmd.vmClasses, "vm-classes", "", "The list of virtual machine classes to associate with the vSphere Namespace.")
}

func (cmd *update) Process(ctx context.Context) error {
	if len(cmd.libraries) > 0 {
		cmd.spec.VmServiceSpec.ContentLibraries = strings.Split(cmd.libraries, ",")
	}
	if len(cmd.vmClasses) > 0 {
		cmd.spec.VmServiceSpec.VmClasses = strings.Split(cmd.vmClasses, ",")
	}
	return cmd.ClientFlag.Process(ctx)
}

func (cmd *update) Usage() string {
	return "NAME"
}

func (cmd *update) Description() string {
	return `Modifies an existing vSphere Namespace on a Supervisor. 

Examples:
  govc namespace.update -content-libraries=dca9cc16-9460-4da0-802c-4aa148ac6cf7 test-namespace
  govc namespace.update -content-libraries=dca9cc16-9460-4da0-802c-4aa148ac6cf7,617a3ee3-a2ff-4311-9a7c-0016ccf958bd test-namespace
  govc namespace.update -vm-classes=best-effort-2xlarge test-namespace
  govc namespace.update -vm-classes=best-effort-2xlarge,best-effort-4xlarge test-namespace
  govc namespace.update -content-libraries=dca9cc16-9460-4da0-802c-4aa148ac6cf7,617a3ee3-a2ff-4311-9a7c-0016ccf958bd -vm-classes=best-effort-2xlarge,best-effort-4xlarge test-namespace`
}

func (cmd *update) Run(ctx context.Context, f *flag.FlagSet) error {
	rc, err := cmd.RestClient()

	if err != nil {
		return err
	}

	nm := namespace.NewManager(rc)

	return nm.UpdateNamespace(ctx, f.Arg(0), cmd.spec)
}
