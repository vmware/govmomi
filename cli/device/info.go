// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package device

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"path"
	"strings"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/types"
)

type info struct {
	*flags.VirtualMachineFlag
	*flags.OutputFlag
	*flags.NetworkFlag
}

func init() {
	cli.Register("device.info", &info{})
}

func (cmd *info) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)

	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)

	cmd.NetworkFlag, ctx = flags.NewNetworkFlag(ctx)
	cmd.NetworkFlag.Register(ctx, f)
}

func (cmd *info) Process(ctx context.Context) error {
	if err := cmd.VirtualMachineFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.NetworkFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *info) Usage() string {
	return "[DEVICE]..."
}

func (cmd *info) Description() string {
	return `Device info for VM.

Examples:
  govc device.info -vm $name
  govc device.info -vm $name disk-*
  govc device.info -vm $name -json disk-* | jq -r .devices[].backing.uuid
  govc device.info -vm $name -json 'disk-*' | jq -r .devices[].backing.fileName # vmdk path
  govc device.info -vm $name -json ethernet-0 | jq -r .devices[].macAddress`
}

func match(p string, devices object.VirtualDeviceList) object.VirtualDeviceList {
	var matches object.VirtualDeviceList
	match := func(name string) bool {
		matched, _ := path.Match(p, name)
		return matched
	}

	for _, device := range devices {
		name := devices.Name(device)
		eq := name == p
		if eq || match(name) {
			matches = append(matches, device)
		}
		if eq {
			break
		}
	}

	return matches
}

// Match returns devices where VirtualDeviceList.Name matches any of the strings in given args.
// See also: path.Match, govc device.info, govc device.ls
func Match(devices object.VirtualDeviceList, args []string) (object.VirtualDeviceList, error) {
	var found object.VirtualDeviceList

	for _, name := range args {
		matches := match(name, devices)
		if len(matches) == 0 {
			return nil, fmt.Errorf("device '%s' not found", name)
		}
		found = append(found, matches...)
	}

	return found, nil
}

func (cmd *info) Run(ctx context.Context, f *flag.FlagSet) error {
	vm, err := cmd.VirtualMachine()
	if err != nil {
		return err
	}

	if vm == nil {
		return flag.ErrHelp
	}

	devices, err := vm.Device(ctx)
	if err != nil {
		return err
	}

	res := infoResult{
		list: devices,
	}

	if cmd.NetworkFlag.IsSet() {
		net, err := cmd.Network()
		if err != nil {
			return err
		}

		backing, err := net.EthernetCardBackingInfo(ctx)
		if err != nil {
			return err
		}

		devices = devices.SelectByBackingInfo(backing)
	}

	if f.NArg() == 0 {
		res.Devices = toInfoList(devices)
	} else {
		devices, err = Match(devices, f.Args())
		if err != nil {
			return err
		}

		res.Devices = append(res.Devices, toInfoList(devices)...)
	}

	return cmd.WriteResult(&res)
}

func toInfoList(devices object.VirtualDeviceList) []infoDevice {
	var res []infoDevice

	for _, device := range devices {
		res = append(res, infoDevice{
			Name:              devices.Name(device),
			Type:              devices.TypeName(device),
			BaseVirtualDevice: device,
		})
	}

	return res
}

type infoDevice struct {
	Name string `json:"name"`
	Type string `json:"type"`
	types.BaseVirtualDevice
}

func (d *infoDevice) MarshalJSON() ([]byte, error) {
	b, err := json.Marshal(d.BaseVirtualDevice)
	if err != nil {
		return b, err
	}

	// TODO: make use of "inline" tag if it comes to be: https://github.com/golang/go/issues/6213

	return append([]byte(fmt.Sprintf(`{"name":"%s","type":"%s",`, d.Name, d.Type)), b[1:]...), err
}

type infoResult struct {
	Devices []infoDevice `json:"devices"`
	// need the full list of devices to lookup attached devices and controllers
	list object.VirtualDeviceList
}

func (r *infoResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)

	for i := range r.Devices {
		device := r.Devices[i].BaseVirtualDevice
		d := device.GetVirtualDevice()
		info := d.DeviceInfo.GetDescription()

		fmt.Fprintf(tw, "Name:\t%s\n", r.Devices[i].Name)
		fmt.Fprintf(tw, "  Type:\t%s\n", r.list.TypeName(device))
		fmt.Fprintf(tw, "  Label:\t%s\n", info.Label)
		fmt.Fprintf(tw, "  Summary:\t%s\n", info.Summary)
		fmt.Fprintf(tw, "  Key:\t%d\n", d.Key)

		if c, ok := device.(types.BaseVirtualController); ok {
			var attached []string
			for _, key := range c.GetVirtualController().Device {
				attached = append(attached, r.list.Name(r.list.FindByKey(key)))
			}
			fmt.Fprintf(tw, "  Devices:\t%s\n", strings.Join(attached, ", "))
		} else {
			if c := r.list.FindByKey(d.ControllerKey); c != nil {
				fmt.Fprintf(tw, "  Controller:\t%s\n", r.list.Name(c))
				if d.UnitNumber != nil {
					fmt.Fprintf(tw, "  Unit number:\t%d\n", *d.UnitNumber)
				} else {
					fmt.Fprintf(tw, "  Unit number:\t<nil>\n")
				}
			}
		}

		if ca := d.Connectable; ca != nil {
			fmt.Fprintf(tw, "  Connected:\t%t\n", ca.Connected)
			fmt.Fprintf(tw, "  Start connected:\t%t\n", ca.StartConnected)
			fmt.Fprintf(tw, "  Guest control:\t%t\n", ca.AllowGuestControl)
			fmt.Fprintf(tw, "  Status:\t%s\n", ca.Status)
		}

		switch md := device.(type) {
		case types.BaseVirtualEthernetCard:
			fmt.Fprintf(tw, "  MAC Address:\t%s\n", md.GetVirtualEthernetCard().MacAddress)
			fmt.Fprintf(tw, "  Address type:\t%s\n", md.GetVirtualEthernetCard().AddressType)
		case *types.VirtualDisk:
			if b, ok := md.Backing.(types.BaseVirtualDeviceFileBackingInfo); ok {
				fmt.Fprintf(tw, "  File:\t%s\n", b.GetVirtualDeviceFileBackingInfo().FileName)
			}
			if b, ok := md.Backing.(*types.VirtualDiskFlatVer2BackingInfo); ok && b.Parent != nil {
				fmt.Fprintf(tw, "  Parent:\t%s\n", b.Parent.GetVirtualDeviceFileBackingInfo().FileName)
			}
		case *types.VirtualSerialPort:
			if b, ok := md.Backing.(*types.VirtualSerialPortURIBackingInfo); ok {
				fmt.Fprintf(tw, "  Direction:\t%s\n", b.Direction)
				fmt.Fprintf(tw, "  Service URI:\t%s\n", b.ServiceURI)
				fmt.Fprintf(tw, "  Proxy URI:\t%s\n", b.ProxyURI)
			}
		case *types.VirtualPrecisionClock:
			if b, ok := md.Backing.(*types.VirtualPrecisionClockSystemClockBackingInfo); ok {
				proto := b.Protocol
				if proto == "" {
					proto = string(types.HostDateTimeInfoProtocolPtp)
				}
				fmt.Fprintf(tw, "  Protocol:\t%s\n", proto)
			}
		}
	}

	return tw.Flush()
}
