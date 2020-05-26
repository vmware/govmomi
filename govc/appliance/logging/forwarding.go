/*
Copyright (c) 2020 VMware, Inc. All Rights Reserved.

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

package logging

import (
	"context"
	"flag"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	vlogging "github.com/vmware/govmomi/vapi/appliance/logging"
)

func main() {
	fmt.Println("vim-go")
}

type forwarding struct {
	*flags.ClientFlag
	*flags.OutputFlag
}

func init() {
	cli.Register("appliance.logging.forwarding", &forwarding{})
}

func (cmd *forwarding) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)
}

func (cmd *forwarding) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *forwarding) Usage() string {
	return ""
}

func (cmd *forwarding) Description() string {
	return `Retrieve the VC Appliance log forwarding configuration

Examples:
  govc appliance.logging.forwarding`
}

func (cmd *forwarding) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	fwd := vlogging.NewForwarding(c)

	res, err := fwd.Config(ctx)
	if err != nil {
		return nil
	}

	return cmd.WriteResult(forwardingConfigResult(res))
}

type forwardingConfigResult []vlogging.Config

func (res forwardingConfigResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)

	for _, c := range res {
		fmt.Fprintf(tw, "Hostname:\t%s\n", c.Hostname)
		fmt.Fprintf(tw, "Port:\t%d\n", c.Port)
		fmt.Fprintf(tw, "Protocol:\t%s\n", c.Protocol)
	}

	return tw.Flush()
}
