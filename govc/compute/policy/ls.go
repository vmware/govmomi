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
	"io"
	"text/tabwriter"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vapi/compute"
	"github.com/vmware/govmomi/vapi/rest"
)

type ls struct {
	*flags.ClientFlag
	*flags.OutputFlag

	cap bool
}

func init() {
	cli.Register("compute.policy.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
	cmd.OutputFlag.Register(ctx, f)

	f.BoolVar(&cmd.cap, "c", false, "List capabilities")
}

func (cmd *ls) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return cmd.OutputFlag.Process(ctx)
}

func (cmd *ls) Usage() string {
	return "[ID]"
}

func (cmd *ls) Description() string {
	return `List compute policies.

Examples:
  govc compute.policy.ls
  govc compute.policy.ls -c
  govc compute.policy.ls ab6b4752-118a-4ac0-867f-4d67d48ce371
  govc compute.policy.ls -c com.vmware.vcenter.compute.policies.capabilities.vm_vm_anti_affinity`
}

type lsResult []compute.Policy

func (res lsResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)

	for _, r := range res {
		_, _ = fmt.Fprintf(tw, "%s\t%s\t%s\n", r.Policy, r.Capability, r.Name)
	}

	return tw.Flush()
}

func (res lsResult) Dump() interface{} {
	return []compute.Policy(res)
}

type lsCapabilityResult []compute.Capability

func (res lsCapabilityResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)

	for _, r := range res {
		_, _ = fmt.Fprintf(tw, "%s\t%s\n", r.Capability, r.Name)
	}

	return tw.Flush()
}

func (res lsCapabilityResult) Dump() interface{} {
	return []compute.Capability(res)
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {
	return cmd.WithRestClient(ctx, func(c *rest.Client) error {
		m := compute.NewPolicyManager(c)

		var err error
		id := f.Arg(0)

		if cmd.cap {
			var res []compute.Capability
			if id == "" {
				res, err = m.ListCapability(ctx)
				if err != nil {
					return err
				}
			} else {
				r, err := m.GetCapability(ctx, id)
				if err != nil {
					return err
				}
				res = append(res, *r)
			}

			return cmd.WriteResult(lsCapabilityResult(res))
		}

		var res []compute.Policy

		if id == "" {
			res, err = m.List(ctx)
			if err != nil {
				return err
			}
		} else {
			r, err := m.Get(ctx, id)
			if err != nil {
				return err
			}
			res = append(res, *r)
		}

		return cmd.WriteResult(lsResult(res))
	})
}
