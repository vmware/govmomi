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

type info struct {
	*flags.ClientFlag
	*flags.OutputFlag
}

func init() {
	cli.Register("namespace.service.info", &info{})
}

func (cmd *info) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
	cmd.OutputFlag.Register(ctx, f)

}

func (cmd *info) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return cmd.OutputFlag.Process(ctx)
}

func (cmd *info) Description() string {
	return `Gets information of a specific supervisor service.

Examples:
  govc namespace.service.info my-supervisor-service
  govc namespace.service.info -json my-supervisor-service | jq .`
}

type infoWriter struct {
	cmd     *info
	Service namespace.SupervisorServiceInfo `json:"service"`
}

func (r *infoWriter) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(os.Stdout, 2, 0, 2, ' ', 0)

	fmt.Fprintf(tw, "%s", r.Service.Name)
	fmt.Fprintf(tw, "\t%s", r.Service.State)
	fmt.Fprintf(tw, "\t%s", r.Service.Description)

	fmt.Fprintf(tw, "\n")

	return tw.Flush()
}

func (r *infoWriter) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.Service)
}

func (r *infoWriter) Dump() interface{} {
	return r.Service
}

func (cmd *info) Usage() string {
	return "NAME"
}

func (cmd *info) Run(ctx context.Context, f *flag.FlagSet) error {
	service := f.Args()
	if len(service) != 1 {
		return flag.ErrHelp
	}

	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	m := namespace.NewManager(c)
	supervisorservice, err := m.GetSupervisorService(ctx, service[0])
	if err != nil {
		return err
	}

	return cmd.WriteResult(&infoWriter{cmd, supervisorservice})
}
