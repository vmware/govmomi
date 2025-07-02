// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package portgroup

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/types"
)

type add struct {
	*flags.DatacenterFlag

	DVPortgroupConfigSpec

	path string
}

func init() {
	cli.Register("dvs.portgroup.add", &add{})
}

func (cmd *add) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatacenterFlag, ctx = flags.NewDatacenterFlag(ctx)
	cmd.DatacenterFlag.Register(ctx, f)

	f.StringVar(&cmd.path, "dvs", "", "DVS path")

	cmd.DVPortgroupConfigSpec.NumPorts = 128 // default

	cmd.DVPortgroupConfigSpec.Register(ctx, f)
}

func (cmd *add) Description() string {
	return `Add portgroup to DVS.

The '-type' options are defined by the dvs.DistributedVirtualPortgroup.PortgroupType API.
The UI labels '-type' as "Port binding" with the following choices:
    "Static binding":  earlyBinding
    "Dynanic binding": lateBinding
    "No binding":      ephemeral

The '-auto-expand' option is labeled in the UI as "Port allocation".
The default value is false, behaves as the UI labeled "Fixed" choice.
When given '-auto-expand=true', behaves as the UI labeled "Elastic" choice.

Examples:
  govc dvs.create DSwitch
  govc dvs.portgroup.add -dvs DSwitch -type earlyBinding -nports 16 ExternalNetwork
  govc dvs.portgroup.add -dvs DSwitch -type ephemeral InternalNetwork
  govc object.destroy network/InternalNetwork # remove the portgroup`
}

func (cmd *add) Process(ctx context.Context) error {
	if err := cmd.DatacenterFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *add) Usage() string {
	return "NAME"
}

func (cmd *add) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() == 0 {
		return flag.ErrHelp
	}

	name := f.Arg(0)

	finder, err := cmd.Finder()
	if err != nil {
		return err
	}

	net, err := finder.Network(ctx, cmd.path)
	if err != nil {
		return err
	}

	dvs, ok := net.(*object.DistributedVirtualSwitch)
	if !ok {
		return fmt.Errorf("%s (%T) is not of type %T", cmd.path, net, dvs)
	}

	cmd.DVPortgroupConfigSpec.Name = name

	task, err := dvs.AddPortgroup(ctx, []types.DVPortgroupConfigSpec{cmd.Spec()})
	if err != nil {
		return err
	}

	logger := cmd.ProgressLogger(fmt.Sprintf("adding %s portgroup to dvs %s... ", name, dvs.InventoryPath))
	defer logger.Wait()

	_, err = task.WaitForResult(ctx, logger)
	return err
}
