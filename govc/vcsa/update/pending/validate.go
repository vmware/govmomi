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
	"strings"
	"text/tabwriter"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vapi/appliance/update/pending"
)

type validate struct {
	*flags.ClientFlag
	*flags.OutputFlag
}

func init() {
	cli.Register("vcsa.update.pending.validate", &validate{})
}

func (cmd *validate) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)
}

func (cmd *validate) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *validate) Description() string {
	return `Validates the user provided data before the update installation.

Examples:
  govc vcsa.update.pending.validate 7.0.3.00000 "key1=val1,key2=val2"`
}

func (cmd *validate) Usage() string {
	return "[VERSION] [USERDATA]"
}

type validateResult struct {
	Values pending.Notifications `json:"values"`
}

func (cmd *validate) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 2 {
		return flag.ErrHelp
	}

	version := f.Arg(0)
	userdata := make(map[string]string)

	for _, inputs := range strings.Split(f.Arg(1), ",") {
		input := strings.Split(inputs, "=")
		userdata[input[0]] = input[1]
	}

	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	m := pending.NewManager(c)

	n, err := m.Validate(ctx, version, userdata)
	if err != nil {
		return err
	}

	return cmd.WriteResult(&validateResult{Values: n})
}

func (res validateResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 10, 4, 0, ' ', 0)
	fmt.Fprintf(w, "Errors:\n")
	for _, err := range res.Values.Errors {
		fmt.Fprintf(w, "\tId:%s\n", err.ID)
		fmt.Fprintf(w, "\tMessage:\n")
		fmt.Fprintf(w, "\t\tId:%s\n", err.Message.ID)
		fmt.Fprintf(w, "\t\tArgs:\n")
		for _, arg := range err.Message.Args {
			fmt.Fprintf(w, "\t\t\t%s\n", arg)
		}
		fmt.Fprintf(w, "\t\tDefault Message:%s\n", err.Message.DefaultMessage)
		fmt.Fprintf(w, "\tResolution:\n")
		fmt.Fprintf(w, "\t\tId:%s\n", err.Resolution.ID)
		fmt.Fprintf(w, "\t\tArgs:\n")
		for _, arg := range err.Resolution.Args {
			fmt.Fprintf(w, "\t\t\t%s\n", arg)
		}
		fmt.Fprintf(w, "\t\tDefault Message:%s\n", err.Resolution.DefaultMessage)
	}
	fmt.Fprintf(w, "Warnings:\n")
	for _, warning := range res.Values.Warnings {
		fmt.Fprintf(w, "\tId:%s\n", warning.ID)
		fmt.Fprintf(w, "\tMessage:\n")
		fmt.Fprintf(w, "\t\tId:%s\n", warning.Message.ID)
		fmt.Fprintf(w, "\t\tArgs:\n")
		for _, arg := range warning.Message.Args {
			fmt.Fprintf(w, "\t\t\t%s\n", arg)
		}
		fmt.Fprintf(w, "\t\tDefault Message:%s\n", warning.Message.DefaultMessage)
		fmt.Fprintf(w, "\tResolution:\n")
		fmt.Fprintf(w, "\t\tId:%s\n", warning.Resolution.ID)
		fmt.Fprintf(w, "\t\tArgs:\n")
		for _, arg := range warning.Resolution.Args {
			fmt.Fprintf(w, "\t\t\t%s\n", arg)
		}
		fmt.Fprintf(w, "\t\tDefault Message:%s\n", warning.Resolution.DefaultMessage)
	}
	fmt.Fprintf(w, "Info:\n")
	for _, i := range res.Values.Info {
		fmt.Fprintf(w, "\tId:%s\n", i.ID)
		fmt.Fprintf(w, "\tMessage:\n")
		fmt.Fprintf(w, "\t\tId:%s\n", i.Message.ID)
		fmt.Fprintf(w, "\t\tArgs:\n")
		for _, arg := range i.Message.Args {
			fmt.Fprintf(w, "\t\t\t%s\n", arg)
		}
		fmt.Fprintf(w, "\t\tDefault Message:%s\n", i.Message.DefaultMessage)
	}

	return tw.Flush()
}
