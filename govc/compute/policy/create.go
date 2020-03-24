/*
Copyright (c) 2020 VMware, Inc. All Rights Reserved.

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

package policy

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vapi/compute"
	"github.com/vmware/govmomi/vapi/rest"
)

type create struct {
	*flags.ClientFlag
	compute.Policy
}

func init() {
	cli.Register("compute.policy.create", &create{})
}

func (cmd *create) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.StringVar(&cmd.Policy.Name, "n", "", "Policy name")
	f.StringVar(&cmd.Policy.Description, "d", "", "Description")
	f.StringVar(&cmd.Policy.HostTag, "host", "", "Host Tag")
	f.StringVar(&cmd.Policy.VMTag, "vm", "", "VM Tag")
}

func (cmd *create) Usage() string {
	return "CAPABILITY"
}

func (cmd *create) Description() string {
	return `Create compute policy.

Examples:
  capability=com.vmware.vcenter.compute.policies.capabilities.vm_vm_anti_affinity
  govc compute.policy.create -n my-policy -d my-desc -vm my-tag $capability`
}

func (cmd *create) Run(ctx context.Context, f *flag.FlagSet) error {
	return cmd.WithRestClient(ctx, func(c *rest.Client) error {
		m := compute.NewPolicyManager(c)

		cmd.Capability = f.Arg(0)

		id, err := m.Create(ctx, cmd.Policy)
		if err != nil {
			return err
		}

		fmt.Println(id)

		return nil
	})
}
