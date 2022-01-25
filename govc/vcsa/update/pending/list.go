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

type list struct {
	*flags.ClientFlag
	*flags.OutputFlag
}

func init() {
	cli.Register("vcsa.update.pending.list", &list{})
}

func (cmd *list) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)
}

func (cmd *list) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *list) Description() string {
	return `Checks if new updates are available.
Examples:
  1. govc vcsa.update.pending.list LOCAL_AND_ONLINE
  2. govc vcsa.update.pending.list LAST_CHECK 
  3. govc vcsa.update.pending.list LOCAL`
}

func (cmd *list) Usage() string {
	return "SOURCE_TYPE"
}

type pendingUpdates struct {
	Value []pending.Summary `json:"values"`
}

func (cmd *list) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}
	sourceType := f.Arg(0)

	if !(sourceType == pending.LAST_CHECK || sourceType == pending.LOCAL || sourceType == pending.LOCAL_AND_ONLINE) {
		return flag.ErrHelp
	}

	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	m := pending.NewManager(c)

	summary, err := m.List(ctx, sourceType)
	if err != nil {
		return err
	}

	return cmd.WriteResult(&pendingUpdates{Value: summary})
}

func (res pendingUpdates) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 10, 4, 0, ' ', 0)
	if len(res.Value) == 0 {
		fmt.Fprintf(w, "No Updates Available")
		return tw.Flush()
	}

	for _, s := range res.Value {
		fmt.Fprintf(w, "==========================================================\n")
		fmt.Fprintf(w, "Name:%s\n", s.Name)
		fmt.Fprintf(w, "Priority:%s\n", s.Priority)
		fmt.Fprintf(w, "Reboot Required:%t\n", s.RebootRequired)
		fmt.Fprintf(w, "Release Date:%s\n", s.ReleaseDate)
		fmt.Fprintf(w, "Severity:%s\n", s.Severity)
		fmt.Fprintf(w, "Size:%d\n", s.Size)
		fmt.Fprintf(w, "Update Type:%s\n", s.UpdateType)
		fmt.Fprintf(w, "Version:%s\n", s.Version)
		fmt.Fprintf(w, "Description:\n")
		fmt.Fprintf(w, "\tArgs:\n")
		for _, arg := range s.Description.Args {
			fmt.Fprintf(w, "\t\t%s\n", arg)
		}
		fmt.Fprintf(w, "\tDefault Message:%s\n", s.Description.DefaultMessage)
		fmt.Fprintf(w, "\tId:%s\n", s.Description.ID)
		fmt.Fprintf(w, "==========================================================\n")
	}

	return tw.Flush()
}
