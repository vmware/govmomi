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
	"fmt"
	"io"
	"strings"
	"text/tabwriter"

	"github.com/vmware/govmomi/vapi/namespace"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
)

type infoResult namespace.NamespacesInstanceInfo

func (r infoResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)
	fmt.Fprintf(tw, "Cluster:\t%s\n", r.ClusterId)
	fmt.Fprintf(tw, "Status:\t%s\n", r.ConfigStatus)
	fmt.Fprintf(tw, "VM Classes:\t%s\n", strings.Join(r.VmServiceSpec.VmClasses, ","))
	fmt.Fprintf(tw, "VM Libraries:\t%s\n", strings.Join(r.VmServiceSpec.ContentLibraries, ","))
	return tw.Flush()
}

type info struct {
	*flags.ClientFlag
	*flags.OutputFlag
}

func init() {
	cli.Register("namespace.info", &info{})
}

func (cmd *info) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)
}

func (cmd *info) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}

	return nil
}

func (cmd *info) Usage() string {
	return "NAME"
}

func (cmd *info) Description() string {
	return `Displays the details of a vSphere Namespace.

Examples:
  govc namespace.info test-namespace`
}

func (cmd *info) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}

	rc, err := cmd.RestClient()
	if err != nil {
		return err
	}

	nm := namespace.NewManager(rc)

	d, err := nm.GetNamespace(ctx, f.Arg(0))
	if err != nil {
		return err
	}

	return cmd.WriteResult(infoResult(d))
}
