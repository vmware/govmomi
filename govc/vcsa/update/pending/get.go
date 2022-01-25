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

type get struct {
	*flags.ClientFlag
	*flags.OutputFlag
}

func init() {
	cli.Register("vcsa.update.pending.get", &get{})
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
	return `Gets update information
Examples:
  govc vcsa.update.pending.get 7.0.3.00000`
}

func (cmd *get) Usage() string {
	return "VERSION"
}

type updateInfo struct {
	Value pending.Info `json:"values"`
}

func (cmd *get) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}
	version := f.Arg(0)

	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	u := pending.NewManager(c)

	info, err := u.Get(ctx, version)
	if err != nil {
		return err
	}

	return cmd.WriteResult(&updateInfo{Value: info})
}

func (res updateInfo) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 10, 4, 0, ' ', 0)

	fmt.Fprintf(w, "Name:%s\n", res.Value.Name)
	fmt.Fprintf(w, "Priority:%s\n", res.Value.Priority)
	fmt.Fprintf(w, "Reboot Required:%t\n", res.Value.RebootRequired)
	fmt.Fprintf(w, "Release Date:%s\n", res.Value.ReleaseDate)
	fmt.Fprintf(w, "Severity:%s\n", res.Value.Severity)
	fmt.Fprintf(w, "Size:%d\n", res.Value.Size)
	fmt.Fprintf(w, "Staged:%t\n", res.Value.Staged)
	fmt.Fprintf(w, "Update Type:%s\n", res.Value.UpdateType)
	fmt.Fprintf(w, "Services will be stopped:\n")
	for _, service := range res.Value.ServicesInfo {
		fmt.Fprintf(w, "\tService:%s\n", service.Service)
		fmt.Fprintf(w, "\tDescription:\n")
		fmt.Fprintf(w, "\t\tDefault Message:%s\n", service.Description.DefaultMessage)
		fmt.Fprintf(w, "\t\tID:%s\n", service.Description.ID)
		fmt.Fprintf(w, "\t\tArgs:\n")
		for _, arg := range service.Description.Args {
			fmt.Fprintf(w, "\t\t\t%s\n", arg)
		}
	}
	fmt.Fprintf(w, "Contents:\n")
	for _, c := range res.Value.Contents {
		fmt.Fprintf(w, "\tDefault Message:%s\n", c.DefaultMessage)
		fmt.Fprintf(w, "\tID:%s\n", c.ID)
		fmt.Fprintf(w, "\tArgs:\n")
		for _, arg := range c.Args {
			fmt.Fprintf(w, "\t\t%s\n", arg)
		}
	}
	fmt.Fprintf(w, "Eulas:\n")
	for _, eula := range res.Value.Eulas {
		fmt.Fprintf(w, "\tID:%s\n", eula.ID)
		fmt.Fprintf(w, "\tArgs:\n")
		for _, arg := range eula.Args {
			fmt.Fprintf(w, "\t\t%s\n", arg)
		}
		fmt.Fprintf(w, "\tDefault Message:%s\n", eula.DefaultMessage)
	}
	return tw.Flush()
}
