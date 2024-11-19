/*
Copyright (c) 2021-2023 VMware, Inc. All Rights Reserved.

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

package service

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/namespace"
)

type ls struct {
	*flags.ClientFlag
	*flags.OutputFlag

	long bool
}

func init() {
	cli.Register("namespace.service.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
	cmd.OutputFlag.Register(ctx, f)

	f.BoolVar(&cmd.long, "l", false, "Long listing format")
}

func (cmd *ls) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return cmd.OutputFlag.Process(ctx)
}

func (cmd *ls) Description() string {
	return `List namepace registered supervisor services.

Examples:
  govc namespace.service.ls
  govc namespace.service.ls -l
  govc namespace.service.ls -json | jq .`
}

type lsWriter struct {
	cmd     *ls
	Service []namespace.SupervisorServiceSummary `json:"service"`
}

func (r *lsWriter) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(os.Stdout, 2, 0, 2, ' ', 0)

	for _, svc := range r.Service {
		fmt.Fprintf(tw, "%s", svc.ID)
		if r.cmd.long {
			fmt.Fprintf(tw, "\t%s", svc.State)
			fmt.Fprintf(tw, "\t%s", svc.Name)
		}
		fmt.Fprintf(tw, "\n")
	}
	return tw.Flush()
}

func (r *lsWriter) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.Service)
}

func (r *lsWriter) Dump() interface{} {
	return r.Service
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	m := namespace.NewManager(c)
	supervisorservices, err := m.ListSupervisorServices(ctx)
	if err != nil {
		return err
	}

	return cmd.WriteResult(&lsWriter{cmd, supervisorservices})
}
