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

package policy

import (
	"context"
	"flag"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vapi/appliance/update/policy"
)

type get struct {
	*flags.ClientFlag
	*flags.OutputFlag
}

func init() {
	cli.Register("vcsa.update.policy.get", &get{})
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
	return `Gets the automatic update checking and staging policy.
Examples:
  govc vcsa.update.policy.get`
}

type result struct {
	Value policy.Info `json:"values"`
}

func (cmd *get) Run(ctx context.Context, f *flag.FlagSet) error {

	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	u := policy.NewManager(c)

	info, err := u.Get(ctx)
	if err != nil {
		return err
	}

	return cmd.WriteResult(&result{Value: info})
}

func (res result) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 10, 4, 0, ' ', 0)
	fmt.Fprintf(w, "Auto Stage:%t\n", res.Value.AutoStage)
	fmt.Fprintf(w, "Auto Update:%t\n", res.Value.AutoUpdate)
	fmt.Fprintf(w, "Certificate Check:%t\n", res.Value.CertificateCheck)
	fmt.Fprintf(w, "Custom URL:%s\n", res.Value.CustomURL)
	fmt.Fprintf(w, "Default URL:%s\n", res.Value.DefaultURL)
	fmt.Fprintf(w, "Manual Control:%t\n", res.Value.ManualControl)
	fmt.Fprintf(w, "Username:%s\n", res.Value.Username)
	fmt.Fprintf(w, "Check Schedule:\n")
	for _, item := range res.Value.CheckSchedule {
		fmt.Fprintf(w, "\tDay:%s\n", item.Day)
		fmt.Fprintf(w, "\tHour:%d\n", item.Hour)
		fmt.Fprintf(w, "\tMinute:%d\n", item.Minute)
	}

	return tw.Flush()
}
