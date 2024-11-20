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

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/appliance/access/shell"
)

type set struct {
	*flags.ClientFlag
	enabled bool
	timeout int
}

func init() {
	cli.Register("vcsa.access.shell.set", &set{})
}

func (cmd *set) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.BoolVar(&cmd.enabled,
		"enabled",
		false,
		"Enable BASH, that is, access to BASH from within the controlled CLI.")
	f.IntVar(&cmd.timeout,
		"timeout",
		0,
		"The timeout (in seconds) specifies how long you enable the Shell access. The maximum timeout is 86400 seconds(1 day).")
}

func (cmd *set) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}

	return nil
}

func (cmd *set) Description() string {
	return `Set enabled state of BASH, that is, access to BASH from within the controlled CLI.

Note: This command requires vCenter 7.0.2 or higher.

Examples:
# Enable Shell
govc vcsa.access.shell.set -enabled=true -timeout=240

# Disable Shell
govc vcsa.access.shell.set -enabled=false`
}

func (cmd *set) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	m := shell.NewManager(c)

	input := shell.Access{
		Enabled: cmd.enabled,
		Timeout: cmd.timeout,
	}
	if err = m.Set(ctx, input); err != nil {
		return err
	}

	return nil
}
