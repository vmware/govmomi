/*
Copyright (c) 2023-2023 VMware, Inc. All Rights Reserved.

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

package option

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vim25/types"
)

type ls struct {
	flags.EnvBrowser
}

func init() {
	cli.Register("vm.option.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.EnvBrowser.Register(ctx, f)
}

func (cmd *ls) Description() string {
	return `List VM config option keys for CLUSTER.

Examples:
  govc vm.option.ls -cluster C0`
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {
	b, err := cmd.Browser(ctx)
	if err != nil {
		return err
	}

	opts, err := b.QueryConfigOptionDescriptor(ctx)
	if err != nil {
		return err
	}

	return cmd.VirtualMachineFlag.WriteResult(&lsResult{opts})
}

type lsResult struct {
	opts []types.VirtualMachineConfigOptionDescriptor
}

func (r *lsResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)

	for _, d := range r.opts {
		_, _ = fmt.Fprintf(tw, "%s\t%s\n", d.Key, d.Description)
	}

	return tw.Flush()
}

func (r *lsResult) Dump() interface{} {
	return r.opts
}

func (r *lsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.opts)
}
