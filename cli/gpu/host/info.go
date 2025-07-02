// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package host

import (
	"context"
	"flag"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type info struct {
	*flags.ClientFlag
	*flags.HostSystemFlag
	*flags.OutputFlag
}

func init() {
	cli.Register("gpu.host.info", &info{})
}

func (cmd *info) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.HostSystemFlag, ctx = flags.NewHostSystemFlag(ctx)
	cmd.HostSystemFlag.Register(ctx, f)

	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)
}

func (cmd *info) Description() string {
	return `Display GPU information for a host.

Examples:
  govc gpu.host.info -host hostname
  govc gpu.host.info -host hostname -json | jq .
  govc gpu.host.info -host hostname -json | jq -r '.devices[] | select(.deviceName | contains("NVIDIA"))'
  govc find / -type h | xargs -n1 govc gpu.host.info -host # all hosts`
}

func (cmd *info) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.HostSystemFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

// Support NVIDIA devices, and exclude virtual GPUs, which have a SubDeviceId of 0x0000
func isPhysicalGPU(device types.HostPciDevice) bool {
	return device.VendorId == 0x10de && device.SubDeviceId != 0x0000
}

type gpuInfo struct {
	PciId       string `json:"pciId"`
	DeviceName  string `json:"deviceName"`
	VendorName  string `json:"vendorName"`
	DeviceId    int16  `json:"deviceId"`
	VendorId    int16  `json:"vendorId"`
	SubVendorId int16  `json:"subVendorId"`
	SubDeviceId int16  `json:"subDeviceId"`
	ClassId     int16  `json:"classId"`
	Bus         uint8  `json:"bus"`
	Slot        uint8  `json:"slot"`
	Function    uint8  `json:"function"`
	Bridge      string `json:"bridge,omitempty"`
}

type infoResult struct {
	Devices []gpuInfo `json:"devices"`
}

func (r *infoResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)

	for _, gpu := range r.Devices {
		fmt.Fprintf(tw, "PCI ID: %s\n", gpu.PciId)
		fmt.Fprintf(tw, "  Device Name: %s\n", gpu.DeviceName)
		fmt.Fprintf(tw, "  Vendor Name: %s\n", gpu.VendorName)
		fmt.Fprintf(tw, "  Device ID: 0x%04x\n", gpu.DeviceId)
		fmt.Fprintf(tw, "  Vendor ID: 0x%04x\n", gpu.VendorId)
		fmt.Fprintf(tw, "  SubVendor ID: 0x%04x\n", gpu.SubVendorId)
		fmt.Fprintf(tw, "  SubDevice ID: 0x%04x\n", gpu.SubDeviceId)
		fmt.Fprintf(tw, "  Class ID: 0x%04x\n", gpu.ClassId)
		fmt.Fprintf(tw, "  Bus: 0x%02x, Slot: 0x%02x, Function: 0x%02x\n", gpu.Bus, gpu.Slot, gpu.Function)
		if gpu.Bridge != "" {
			fmt.Fprintf(tw, "  Parent Bridge: %s\n", gpu.Bridge)
		}
	}

	return tw.Flush()
}

func (cmd *info) Run(ctx context.Context, f *flag.FlagSet) error {
	host, err := cmd.HostSystem()
	if err != nil {
		return err
	}

	if host == nil {
		return flag.ErrHelp
	}

	var h mo.HostSystem
	pc := property.DefaultCollector(host.Client())
	err = pc.RetrieveOne(ctx, host.Reference(), []string{"hardware"}, &h)
	if err != nil {
		return err
	}

	var res infoResult
	for _, device := range h.Hardware.PciDevice {
		if isPhysicalGPU(device) {
			res.Devices = append(res.Devices, gpuInfo{
				PciId:       device.Id,
				DeviceName:  device.DeviceName,
				VendorName:  device.VendorName,
				DeviceId:    device.DeviceId,
				VendorId:    device.VendorId,
				SubVendorId: device.SubVendorId,
				SubDeviceId: device.SubDeviceId,
				ClassId:     device.ClassId,
				Bus:         device.Bus,
				Slot:        device.Slot,
				Function:    device.Function,
				Bridge:      device.ParentBridge,
			})
		}
	}

	return cmd.WriteResult(&res)
}
