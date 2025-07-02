// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package target

import (
	"context"
	"flag"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/units"
	"github.com/vmware/govmomi/vim25/types"
)

type info struct {
	flags.EnvBrowser

	datastore bool
	network   bool
	disk      bool
	device    bool
}

func init() {
	cli.Register("vm.target.info", &info{})
}

func (cmd *info) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.EnvBrowser.Register(ctx, f)

	f.BoolVar(&cmd.datastore, "datastore", true, "Include Datastores")
	f.BoolVar(&cmd.network, "network", true, "Include Networks")
	f.BoolVar(&cmd.disk, "disk", false, "Include Disks")
	f.BoolVar(&cmd.device, "device", true, "Include Devices")
}

func (cmd *info) Description() string {
	return `VM config target info.

The config target data contains information about the execution environment for a VM
in the given CLUSTER, and optionally for a specific HOST.

Examples:
  govc vm.target.info -cluster C0
  govc vm.target.info -host my_hostname
  govc vm.target.info -vm my_vm`
}

func (cmd *info) Run(ctx context.Context, f *flag.FlagSet) error {
	b, err := cmd.Browser(ctx)
	if err != nil {
		return err
	}

	host, err := cmd.HostSystemIfSpecified()
	if err != nil {
		return err
	}

	target, err := b.QueryConfigTarget(ctx, host)
	if err != nil {
		return err
	}

	if cmd.network == false {
		target.Network = nil
		target.DistributedVirtualPortgroup = nil
		target.DistributedVirtualSwitch = nil
		target.OpaqueNetwork = nil
		target.LegacyNetworkInfo = nil
	}

	if cmd.datastore == false {
		target.Datastore = nil
	}

	if cmd.disk == false {
		target.ScsiDisk = nil
		target.IdeDisk = nil
	}

	return cmd.VirtualMachineFlag.WriteResult(&infoResult{target})
}

type infoResult struct {
	*types.ConfigTarget
}

func (r *infoResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)

	fmt.Fprintf(tw, "CPUs:\t%d\n", r.ConfigTarget.NumCpus)
	fmt.Fprintf(tw, "CPU cores:\t%d\n", r.ConfigTarget.NumCpuCores)

	for _, ds := range r.ConfigTarget.Datastore {
		fmt.Fprintf(tw, "Datastore:\t%s\n", ds.Name)
		fmt.Fprintf(tw, "  Capacity:\t%s\n", units.ByteSize(ds.Datastore.Capacity))
		fmt.Fprintf(tw, "  Free:\t%s\n", units.ByteSize(ds.Datastore.FreeSpace))
		fmt.Fprintf(tw, "  Uncommitted:\t%s\n", units.ByteSize(ds.Datastore.Uncommitted))
	}

	for _, net := range r.ConfigTarget.Network {
		fmt.Fprintf(tw, "Network:\t%s\n", net.Name)
	}

	for _, net := range r.ConfigTarget.DistributedVirtualPortgroup {
		if net.UplinkPortgroup {
			continue
		}
		fmt.Fprintf(tw, "Network:\t%s\n", net.PortgroupName)
		fmt.Fprintf(tw, "  Switch:\t%s\n", net.SwitchName)
	}

	for _, disk := range r.ConfigTarget.ScsiPassthrough {
		fmt.Fprintf(tw, "SCSI passthrough:\t%s\n", disk.Name)
		fmt.Fprintf(tw, "  Unit:\t%d\n", disk.PhysicalUnitNumber)
	}

	for _, clock := range r.ConfigTarget.PrecisionClockInfo {
		fmt.Fprintf(tw, "PrecisionClock:\t%s\n", clock.SystemClockProtocol)
	}

	return tw.Flush()
}

func (r *infoResult) Dump() any {
	return r.ConfigTarget
}
