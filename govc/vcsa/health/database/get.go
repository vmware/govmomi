/*
Copyright (c) 2022 VMware, Inc. All Rights Reserved.

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

package database

import (
	"context"
	"flag"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vapi/appliance/health/database"
)

type get struct {
	*flags.ClientFlag
	*flags.OutputFlag
}

func init() {
	cli.Register("vcsa.health.database.get", &get{})
}

func (cmd *get) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)
}

func (cmd *get) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *get) Description() string {
	return `Returns the health status of the database.
Examples:
  govc vcsa.health.database.get`
}

type databaseHealth struct {
	Values database.Info `json:"values"`
}

func (cmd *get) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	m := database.NewManager(c)

	var status database.Info
	status, err = m.Get(ctx)
	if err != nil {
		return err
	}

	return cmd.WriteResult(&databaseHealth{Values: status})
}

func (res databaseHealth) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 10, 4, 0, ' ', 0)
	fmt.Fprintf(w, "Status:%s\n", res.Values.Status)
	fmt.Fprintf(w, "Messages:\n")
	fmt.Fprintf(w, "\tMessage:\n")
	for _, message := range res.Values.Messages {
		fmt.Fprintf(w, "\t\tArgs:\n")
		for _, arg := range message.Args {
			fmt.Fprintf(w, "\t\t\t%s\n", arg)
		}
		fmt.Fprintf(w, "\t\tID:%s\n", message.ID)
		fmt.Fprintf(w, "\t\tDefault Message:%s\n", message.DefaultMessage)
	}

	return tw.Flush()
}
