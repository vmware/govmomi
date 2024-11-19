/*
Copyright (c) 2022 VMware, Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0.
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package shell

import (
	"context"
	"flag"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/appliance/access/shell"
)

type get struct {
	*flags.ClientFlag
	*flags.OutputFlag
}

func init() {
	cli.Register("vcsa.access.shell.get", &get{})
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
	return `Get enabled state of BASH, that is, access to BASH from within the controlled CLI.

Note: This command requires vCenter 7.0.2 or higher.

Examples:
govc vcsa.access.shell.get`
}

type access struct {
	shell.Access
}

func (cmd *get) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	m := shell.NewManager(c)

	status, err := m.Get(ctx)
	if err != nil {
		return err
	}

	return cmd.WriteResult(&access{status})
}

func (res access) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 0, 4, 0, ' ', 0)

	fmt.Fprintf(tw, "Enabled:%t\n", res.Enabled)
	fmt.Fprintf(tw, "Timeout:%d\n", res.Timeout)

	return tw.Flush()
}
