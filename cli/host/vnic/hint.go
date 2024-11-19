/*
Copyright (c) 2022-2023 VMware, Inc. All Rights Reserved.

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

package vnic

import (
	"context"
	"errors"
	"flag"
	"io"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vim25/types"
)

type hint struct {
	*flags.HostSystemFlag
}

func init() {
	cli.Register("host.vnic.hint", &hint{})
}

func (cmd *hint) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.HostSystemFlag, ctx = flags.NewHostSystemFlag(ctx)
	cmd.HostSystemFlag.Register(ctx, f)
}

func (cmd *hint) Usage() string {
	return "[DEVICE]..."
}

func (cmd *hint) Description() string {
	return `Query virtual nic DEVICE hints.
Examples:
  govc host.vnic.hint -host hostname
  govc host.vnic.hint -host hostname vmnic1`
}

type hintResult struct {
	Hint []types.PhysicalNicHintInfo `json:"hint"`
}

func (i *hintResult) Write(w io.Writer) error {
	// TODO: human friendly output
	return errors.New("-xml, -json or -dump flag required")
}

func (cmd *hint) Run(ctx context.Context, f *flag.FlagSet) error {
	ns, err := cmd.HostNetworkSystem()
	if err != nil {
		return err
	}

	hints, err := ns.QueryNetworkHint(ctx, f.Args())
	if err != nil {
		return err
	}

	return cmd.WriteResult(&hintResult{Hint: hints})
}
