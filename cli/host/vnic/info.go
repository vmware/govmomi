// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vnic

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type info struct {
	*flags.HostSystemFlag
}

func init() {
	cli.Register("host.vnic.info", &info{})
}

func (cmd *info) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.HostSystemFlag, ctx = flags.NewHostSystemFlag(ctx)
	cmd.HostSystemFlag.Register(ctx, f)
}

func (cmd *info) Process(ctx context.Context) error {
	if err := cmd.HostSystemFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

type Info struct {
	Device   string   `json:"device"`
	Network  string   `json:"network"`
	Switch   string   `json:"switch"`
	Address  string   `json:"address"`
	Stack    string   `json:"stack"`
	Services []string `json:"services"`
}

type infoResult struct {
	Info []Info `json:"info"`
}

func (i *infoResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(os.Stdout, 2, 0, 2, ' ', 0)

	for _, info := range i.Info {
		fmt.Fprintf(tw, "Device:\t%s\n", info.Device)
		fmt.Fprintf(tw, "Network label:\t%s\n", info.Network)
		fmt.Fprintf(tw, "Switch:\t%s\n", info.Switch)
		fmt.Fprintf(tw, "IP address:\t%s\n", info.Address)
		fmt.Fprintf(tw, "TCP/IP stack:\t%s\n", info.Stack)
		fmt.Fprintf(tw, "Enabled services:\t%s\n", strings.Join(info.Services, ", "))
	}

	return tw.Flush()
}

func (cmd *info) Run(ctx context.Context, f *flag.FlagSet) error {
	host, err := cmd.HostSystem()
	if err != nil {
		return err
	}

	ns, err := cmd.HostNetworkSystem()
	if err != nil {
		return err
	}

	var mns mo.HostNetworkSystem

	m, err := host.ConfigManager().VirtualNicManager(ctx)
	if err != nil {
		return err
	}

	minfo, err := m.Info(ctx)
	if err != nil {
		return err
	}

	err = ns.Properties(ctx, ns.Reference(), []string{"networkInfo"}, &mns)
	if err != nil {
		return err
	}

	type dnet struct {
		dvp mo.DistributedVirtualPortgroup
		dvs mo.VmwareDistributedVirtualSwitch
	}

	dnets := make(map[string]*dnet)
	var res infoResult

	for _, nic := range mns.NetworkInfo.Vnic {
		info := Info{Device: nic.Device}

		if dvp := nic.Spec.DistributedVirtualPort; dvp != nil {
			dn, ok := dnets[dvp.PortgroupKey]

			if !ok {
				dn = new(dnet)
				o := object.NewDistributedVirtualPortgroup(host.Client(), types.ManagedObjectReference{
					Type:  "DistributedVirtualPortgroup",
					Value: dvp.PortgroupKey,
				})

				err = o.Properties(ctx, o.Reference(), []string{"name", "config.distributedVirtualSwitch"}, &dn.dvp)
				if err != nil {
					return err
				}

				err = o.Properties(ctx, *dn.dvp.Config.DistributedVirtualSwitch, []string{"name"}, &dn.dvs)
				if err != nil {
					return err
				}

				dnets[dvp.PortgroupKey] = dn
			}

			info.Network = dn.dvp.Name
			info.Switch = dn.dvs.Name
		} else {
			info.Network = nic.Portgroup
			for _, pg := range mns.NetworkInfo.Portgroup {
				if pg.Spec.Name == nic.Portgroup {
					info.Switch = pg.Spec.VswitchName
					break
				}
			}
		}

		info.Address = nic.Spec.Ip.IpAddress
		info.Stack = nic.Spec.NetStackInstanceKey

		for _, nc := range minfo.NetConfig {
			for _, dev := range nc.SelectedVnic {
				key := nc.NicType + "." + nic.Key
				if dev == key {
					info.Services = append(info.Services, nc.NicType)
				}
			}

		}

		res.Info = append(res.Info, info)
	}

	return cmd.WriteResult(&res)
}
