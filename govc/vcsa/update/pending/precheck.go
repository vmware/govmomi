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

package pending

import (
	"context"
	"flag"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vapi/appliance/update/pending"
)

type precheck struct {
	*flags.ClientFlag
	*flags.OutputFlag
}

func init() {
	cli.Register("vcsa.update.pending.precheck", &precheck{})
}

func (cmd *precheck) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)
}

func (cmd *precheck) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *precheck) Description() string {
	return `Runs update precheck
Examples:
  govc vcsa.update.pending.precheck 7.0.3.00000`
}

func (cmd *precheck) Usage() string {
	return "VERSION"
}

type updatePrecheck struct {
	Values pending.PrecheckResult `json:"values"`
}

func (cmd *precheck) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}
	version := f.Arg(0)

	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	u := pending.NewManager(c)

	info, err := u.Precheck(ctx, version)
	if err != nil {
		return err
	}

	return cmd.WriteResult(&updatePrecheck{Values: info})
}

func (res updatePrecheck) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 10, 4, 0, ' ', 0)

	fmt.Fprintf(w, "Check Time:%s\n", res.Values.CheckTime)
	fmt.Fprintf(w, "Reboot Required:%t\n", res.Values.RebootRequired)
	fmt.Fprintf(w, "Questions:\n")
	for _, question := range res.Values.Questions {
		fmt.Fprintf(w, "\tData Item:%s\n", question.DataItem)
		fmt.Fprintf(w, "\tType:%s\n", question.Type)
		fmt.Fprintf(w, "\tDescription:\n")
		fmt.Fprintf(w, "\t\tID:%s\n", question.Description.ID)
		fmt.Fprintf(w, "\t\tArgs:\n")
		for _, arg := range question.Description.Args {
			fmt.Fprintf(w, "\t\t\t%s\n", arg)
		}
		fmt.Fprintf(w, "\t\tDefault Message:%s\n", question.Description.DefaultMessage)
		fmt.Fprintf(w, "\tText:\n")
		fmt.Fprintf(w, "\t\tID:%s\n", question.Text.ID)
		fmt.Fprintf(w, "\t\tArgs:\n")
		for _, arg := range question.Text.Args {
			fmt.Fprintf(w, "\t\t\t%s\n", arg)
		}
		fmt.Fprintf(w, "\t\tDefault Message:%s\n", question.Text.DefaultMessage)
	}

	return tw.Flush()
}
