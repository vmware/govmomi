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

package staged

import (
	"context"
	"flag"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vapi/appliance/update/staged"
)

type get struct {
	*flags.ClientFlag
	*flags.OutputFlag
}

func init() {
	cli.Register("vcsa.update.staged.get", &get{})
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
	return `Gets the current status of the staged update.
Examples:
  govc vcsa.update.staged.get`
}

type result struct {
	Value staged.Info `json:"values"`
}

func (cmd *get) Run(ctx context.Context, f *flag.FlagSet) error {

	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	m := staged.NewManager(c)

	info, err := m.Get(ctx)
	if err != nil {
		return err
	}

	return cmd.WriteResult(&result{Value: info})
}

func (res result) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 10, 4, 0, ' ', 0)
	fmt.Fprintf(w, "Name:%s\n", res.Value.Name)
	fmt.Fprintf(w, "Priority:%s\n", res.Value.Priority)
	fmt.Fprintf(w, "Reboot Required:%t\n", res.Value.RebootRequired)
	fmt.Fprintf(w, "Release Date:%s\n", res.Value.ReleaseDate)
	fmt.Fprintf(w, "Severity:%s\n", res.Value.Severity)
	fmt.Fprintf(w, "Size :%d\n", res.Value.Size)
	fmt.Fprintf(w, "Staging complete:%t\n", res.Value.StagingComplete)
	fmt.Fprintf(w, "Update type:%s\n", res.Value.UpdateType)
	fmt.Fprintf(w, "Version:%s\n", res.Value.Version)
	fmt.Fprintf(w, "Description:\n")
	fmt.Fprintf(w, "\tID:%s\n", res.Value.Description.ID)
	fmt.Fprintf(w, "\tArgs:")
	for _, arg := range res.Value.Description.Args {
		fmt.Fprintf(w, "\t\t%s\n", arg)
	}
	fmt.Fprintf(w, "\tDefault Message:%s\n", res.Value.Description.DefaultMessage)
	return tw.Flush()
}
