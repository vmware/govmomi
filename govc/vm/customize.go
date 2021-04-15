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
	"strings"

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
	mac       flags.StringList
	ip        flags.StringList
	ip6       flags.StringList
	gateway   flags.StringList
	netmask   flags.StringList
	dnsserver flags.StringList
	dnssuffix flags.StringList
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
	f.Var(&cmd.mac, "mac", "MAC address")
	cmd.mac = nil
	f.Var(&cmd.ip, "ip", "IPv4 address")
	cmd.ip = nil
	f.Var(&cmd.ip6, "ip6", "IPv6 addresses with optional netmask (defaults to /64), separated by comma")
	cmd.ip6 = nil
	f.Var(&cmd.gateway, "gateway", "Gateway")
	cmd.gateway = nil
	f.Var(&cmd.netmask, "netmask", "Netmask")
	cmd.netmask = nil
	f.Var(&cmd.dnsserver, "dns-server", "DNS server list")
	cmd.dnsserver = nil
	f.Var(&cmd.dnssuffix, "dns-suffix", "DNS suffix list")
	cmd.dnssuffix = nil
	f.StringVar(&cmd.kind, "type", "Linux", "Customization type if spec NAME is not specified (Linux|Windows)")
}

func (cmd *customize) Usage() string {
	return "[NAME]"
}

func (cmd *customize) Description() string {
	return `Customize VM.

Optionally specify a customization spec NAME.

The '-ip', '-netmask' and '-gateway' flags are for static IP configuration.
If the VM has multiple NICs, an '-ip' and '-netmask' must be specified for each.

The '-dns-server' and '-dns-suffix' flags can be specified multiple times.

Windows -tz value requires the Index (hex): https://support.microsoft.com/en-us/help/973627/microsoft-time-zone-index-values

Examples:
  govc vm.customize -vm VM NAME
  govc vm.customize -vm VM -name my-hostname -ip dhcp
  govc vm.customize -vm VM -gateway GATEWAY -ip NEWIP -netmask NETMASK -dns-server DNS1,DNS2 NAME
  # Multiple -ip without -mac are applied by vCenter in the order in which the NICs appear on the bus
  govc vm.customize -vm VM -ip 10.0.0.178 -netmask 255.255.255.0 -ip 10.0.0.162 -netmask 255.255.255.0
  # Multiple -ip with -mac are applied by vCenter to the NIC with the given MAC address
  govc vm.customize -vm VM -mac 00:50:56:be:dd:f8 -ip 10.0.0.178 -netmask 255.255.255.0 -mac 00:50:56:be:60:cf -ip 10.0.0.162 -netmask 255.255.255.0
  # Dual stack IPv4/IPv6, single NIC
  govc vm.customize -vm VM -ip 10.0.0.1 -netmask 255.255.255.0 -ip6 '2001:db8::1/64' -name my-hostname NAME
  # DHCPv6, single NIC
  govc vm.customize -vm VM -ip6 dhcp6 NAME
  # Static IPv6, three NICs, last one with two addresses
  govc vm.customize -vm VM -ip6 2001:db8::1/64 -ip6 2001:db8::2/64 -ip6 2001:db8::3/64,2001:db8::4/64 NAME
  govc vm.customize -vm VM -auto-login 3 NAME
  govc vm.customize -vm VM -prefix demo NAME
  govc vm.customize -vm VM -tz America/New_York NAME`
}

// Parse a string of multiple IPv6 addresses with optional netmask; separated by comma
func parseIPv6Argument(argv string) (ipconf []types.BaseCustomizationIpV6Generator, err error) {
	for _, substring := range strings.Split(argv, ",") {
		// remove leading and trailing white space
		substring = strings.TrimSpace(substring)
		// handle "dhcp6" and lists of static IPv6 addresses
		switch substring {
		case "dhcp6":
			ipconf = append(
				ipconf,
				&types.CustomizationDhcpIpV6Generator{},
			)
		default:
			// check if subnet mask was specified
			switch strings.Count(substring, "/") {
			// no mask, set default
			case 0:
				ipconf = append(ipconf, &types.CustomizationFixedIpV6{
					IpAddress:  substring,
					SubnetMask: 64,
				})
			// a single forward slash was found: parse and use subnet mask
			case 1:
				parts := strings.Split(substring, "/")
				mask, err := strconv.Atoi(parts[1])
				if err != nil {
					return nil, fmt.Errorf("unable to convert subnet mask to int: %w", err)
				}
				ipconf = append(ipconf, &types.CustomizationFixedIpV6{
					IpAddress:  parts[0],
					SubnetMask: int32(mask),
				})
			// too many forward slashes; return error
			default:
				return nil, fmt.Errorf("unable to parse IPv6 address (too many subnet separators): %s", substring)
			}
		}
	}
	return ipconf, nil
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
			NicSettingMap: make([]types.CustomizationAdapterMapping, len(cmd.ip)),
		}

		switch cmd.kind {
		case "Linux":
			spec.Identity = &types.CustomizationLinuxPrep{
				HostName: new(types.CustomizationVirtualMachineName),
			}
		case "Windows":
			spec.Identity = &types.CustomizationSysprep{
				UserData: types.CustomizationUserData{
					ComputerName: new(types.CustomizationVirtualMachineName),
				},
			}
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

	if len(cmd.ip) > len(spec.NicSettingMap) {
		return fmt.Errorf("%d -ip specified, spec %q has %d", len(cmd.ip), name, len(spec.NicSettingMap))
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

	if len(cmd.dnsserver) != 0 {
		if !isWindows {
			for _, s := range cmd.dnsserver {
				spec.GlobalIPSettings.DnsServerList =
					append(spec.GlobalIPSettings.DnsServerList, strings.Split(s, ",")...)
			}
		}
	}

	spec.GlobalIPSettings.DnsSuffixList = cmd.dnssuffix

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

	for i, ip := range cmd.ip {
		nic := &spec.NicSettingMap[i]
		switch ip {
		case "dhcp":
			nic.Adapter.Ip = new(types.CustomizationDhcpIpGenerator)
		default:
			nic.Adapter.Ip = &types.CustomizationFixedIp{IpAddress: ip}
		}

		if i < len(cmd.netmask) {
			nic.Adapter.SubnetMask = cmd.netmask[i]
		}
		if i < len(cmd.mac) {
			nic.MacAddress = cmd.mac[i]
		}
		if i < len(cmd.gateway) {
			nic.Adapter.Gateway = strings.Split(cmd.gateway[i], ",")
		}
		if isWindows {
			if i < len(cmd.dnsserver) {
				nic.Adapter.DnsServerList = strings.Split(cmd.dnsserver[i], ",")
			}
		}
	}

	for i, ip6 := range cmd.ip6 {
		ipconfig, err := parseIPv6Argument(ip6)
		if err != nil {
			return err
		}
		// use the same logic as the ip switch: the first occurrence of the ip6 switch is assigned to the first nic,
		// the second to the second nic and so forth.
		if spec.NicSettingMap == nil || len(spec.NicSettingMap) < i {
			return fmt.Errorf("unable to find a network adapter for IPv6 settings %d (%s)", i, ip6)
		}
		nic := &spec.NicSettingMap[i]
		if nic.Adapter.IpV6Spec == nil {
			nic.Adapter.IpV6Spec = new(types.CustomizationIPSettingsIpV6AddressSpec)
		}
		nic.Adapter.IpV6Spec.Ip = append(nic.Adapter.IpV6Spec.Ip, ipconfig...)
	}

	task, err := vm.Customize(ctx, *spec)
	if err != nil {
		return err
	}

	return task.Wait(ctx)
}
