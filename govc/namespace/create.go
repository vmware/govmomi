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

type create struct {
	*flags.ClientFlag

	libraries string
	vmClasses string
	spec      namespace.NamespacesInstanceCreateSpec
}

func init() {
	cli.Register("namespace.create", &create{})
}

func (cmd *create) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.StringVar(&cmd.spec.Cluster, "supervisor", "", "The identifier of the Supervisor.")
	f.StringVar(&cmd.spec.Namespace, "namespace", "", "The name of the vSphere Namespace.")
	f.StringVar(&cmd.libraries, "content-libraries", "", "The identifiers of the content libraries to associate with the vSphere Namespace.")
	f.StringVar(&cmd.vmClasses, "vm-classes", "", "The identifiers of the virtual machine classes to associate with the vSphere Namespace.")
}

func (cmd *create) Process(ctx context.Context) error {
	if len(cmd.libraries) > 0 {
		cmd.spec.VmServiceSpec.ContentLibraries = strings.Split(cmd.libraries, ",")
	}
	if len(cmd.vmClasses) > 0 {
		cmd.spec.VmServiceSpec.VmClasses = strings.Split(cmd.vmClasses, ",")
	}
	return cmd.ClientFlag.Process(ctx)
}

func (cmd *create) Description() string {
	return `Creates a new vSphere Namespace on a Supervisor. 

Examples:
  govc namespace.create -namespace=test-namespace -supervisor=domain-c1
  govc namespace.create -namespace=test-namespace -supervisor=domain-c1 -content-libraries=dca9cc16-9460-4da0-802c-4aa148ac6cf7
  govc namespace.create -namespace=test-namespace -supervisor=domain-c1 -content-libraries=dca9cc16-9460-4da0-802c-4aa148ac6cf7,dca9cc16-9460-4da0-802c-4aa148ac6cf7
  govc namespace.create -namespace=test-namespace -supervisor=domain-c1 -vm-classes=best-effort-2xlarge
  govc namespace.create -namespace=test-namespace -supervisor=domain-c1 -vm-classes=best-effort-2xlarge,best-effort-4xlarge
  govc namespace.create -namespace=test-namespace -supervisor=domain-c1 -content-libraries=dca9cc16-9460-4da0-802c-4aa148ac6cf7,dca9cc16-9460-4da0-802c-4aa148ac6cf7 -vm-classes=best-effort-2xlarge,best-effort-4xlarge`
}

func (cmd *create) Run(ctx context.Context, f *flag.FlagSet) error {
	rc, err := cmd.RestClient()

	if err != nil {
		return err
	}

	nm := namespace.NewManager(rc)

	return nm.CreateNamespace(ctx, cmd.spec)
}
