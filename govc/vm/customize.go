/*
Copyright (c) 2019 VMware, Inc. All Rights Reserved.

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

package vm

import (
	"context"
	"flag"
	"fmt"
	"strconv"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/types"
)

type customize struct {
	*flags.VirtualMachineFlag

	alc       int
	prefix    types.CustomizationPrefixName
	tz        string
	domain    string
	host      types.CustomizationFixedName
	ip        string
	gateway   flags.StringList
	netmask   string
	dnsserver flags.StringList
	kind      string
}

func init() {
	cli.Register("vm.customize", &customize{})
}

func (cmd *customize) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)

	f.IntVar(&cmd.alc, "auto-login", 0, "Number of times the VM should automatically login as an administrator")
	f.StringVar(&cmd.prefix.Base, "prefix", "", "Host name generator prefix")
	f.StringVar(&cmd.tz, "tz", "", "Time zone")
	f.StringVar(&cmd.domain, "domain", "", "Domain name")
	f.StringVar(&cmd.host.Name, "name", "", "Host name")
	f.StringVar(&cmd.ip, "ip", "", "IP address")
	f.Var(&cmd.gateway, "gateway", "Gateway")
	f.StringVar(&cmd.netmask, "netmask", "", "Netmask")
	f.Var(&cmd.dnsserver, "dns-server", "DNS server")
	f.StringVar(&cmd.kind, "type", "Linux", "Customization type if spec NAME is not specified (Linux|Windows)")
}

func (cmd *customize) Usage() string {
	return "[NAME]"
}

func (cmd *customize) Description() string {
	return `Customize VM.

Optionally specify a customization spec NAME.

Windows -tz value requires the Index (hex): https://support.microsoft.com/en-us/help/973627/microsoft-time-zone-index-values

Examples:
  govc vm.customize -vm VM -name my-hostname
  govc vm.customize -vm VM NAME
  govc vm.customize -vm VM -gateway GATEWAY -netmask NETMASK -ip NEWIP -dns-server DNS1 -dns-server DNS2 NAME
  govc vm.customize -vm VM -auto-login 3 NAME
  govc vm.customize -vm VM -prefix demo NAME
  govc vm.customize -vm VM -tz America/New_York NAME`
}

func (cmd *customize) Run(ctx context.Context, f *flag.FlagSet) error {
	vm, err := cmd.VirtualMachineFlag.VirtualMachine()
	if err != nil {
		return err
	}

	if vm == nil {
		return flag.ErrHelp
	}

	var spec *types.CustomizationSpec

	name := f.Arg(0)
	if name == "" {
		spec = &types.CustomizationSpec{
			NicSettingMap: []types.CustomizationAdapterMapping{{}},
		}
		spec.NicSettingMap[0].Adapter.Ip = new(types.CustomizationDhcpIpGenerator)
		switch cmd.kind {
		case "Linux":
			spec.Identity = new(types.CustomizationLinuxPrep)
		case "Windows":
			spec.Identity = new(types.CustomizationSysprep)
		default:
			return flag.ErrHelp
		}
	} else {
		m := object.NewCustomizationSpecManager(vm.Client())

		exists, err := m.DoesCustomizationSpecExist(ctx, name)
		if err != nil {
			return err
		}
		if !exists {
			return fmt.Errorf("specification %q does not exist", name)
		}

		item, err := m.GetCustomizationSpec(ctx, name)
		if err != nil {
			return err
		}

		spec = &item.Spec
	}

	sysprep, isWindows := spec.Identity.(*types.CustomizationSysprep)
	linprep, _ := spec.Identity.(*types.CustomizationLinuxPrep)

	if cmd.domain != "" {
		if isWindows {
			sysprep.Identification.JoinDomain = cmd.domain
		} else {
			linprep.Domain = cmd.domain
		}
	}

	if cmd.prefix.Base != "" {
		if isWindows {
			sysprep.UserData.ComputerName = &cmd.prefix
		} else {
			linprep.HostName = &cmd.prefix
		}
	}

	if cmd.host.Name != "" {
		if isWindows {
			sysprep.UserData.ComputerName = &cmd.host
		} else {
			linprep.HostName = &cmd.host
		}
	}

	if cmd.alc != 0 {
		if !isWindows {
			return fmt.Errorf("option '-auto-login' is Windows only")
		}
		sysprep.GuiUnattended.AutoLogon = true
		sysprep.GuiUnattended.AutoLogonCount = int32(cmd.alc)
	}

	if cmd.tz != "" {
		if isWindows {
			tz, err := strconv.ParseInt(cmd.tz, 16, 32)
			if err != nil {
				return fmt.Errorf("converting -tz=%q: %s", cmd.tz, err)
			}
			sysprep.GuiUnattended.TimeZone = int32(tz)
		} else {
			linprep.TimeZone = cmd.tz
		}
	}

	nic := &spec.NicSettingMap[0]
	if cmd.ip != "" {
		nic.Adapter.Ip = &types.CustomizationFixedIp{IpAddress: cmd.ip}
	}
	if cmd.netmask != "" {
		nic.Adapter.SubnetMask = cmd.netmask
	}
	if len(cmd.gateway) != 0 {
		nic.Adapter.Gateway = cmd.gateway
	}
	if len(cmd.dnsserver) != 0 {
		spec.GlobalIPSettings.DnsServerList = cmd.dnsserver
		nic.Adapter.DnsServerList = cmd.dnsserver
	}

	task, err := vm.Customize(ctx, *spec)
	if err != nil {
		return err
	}

	return task.Wait(ctx)
}
