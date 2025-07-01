// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vm

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
	*flags.VirtualMachineFlag
}

func init() {
	cli.Register("gpu.vm.info", &info{})
}

func (cmd *info) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)
}

func (cmd *info) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.VirtualMachineFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *info) Description() string {
	return `Display GPU information for a VM.

Examples:
  govc gpu.vm.info -vm $vm
  govc gpu.vm.info -vm $vm -json | jq -r '.gpus[].summary'
  govc gpu.vm.info -vm $vm -json | jq -r '.gpus[] | select(.summary | contains("nvidia_a40"))'`
}

type gpuInfo struct {
	Index    int    `json:"index"`
	Label    string `json:"label"`
	Summary  string `json:"summary"`
	NumaNode int32  `json:"numaNode"`
	Key      int32  `json:"key"`
}

type infoResult struct {
	GPUs []gpuInfo `json:"gpus"`
}

func (r *infoResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)
	for _, gpu := range r.GPUs {
		fmt.Fprintf(tw, "GPU %d:\n", gpu.Index)
		fmt.Fprintf(tw, "  Label: %s\n", gpu.Label)
		fmt.Fprintf(tw, "  Summary: %s\n", gpu.Summary)
		fmt.Fprintf(tw, "  Numa Node: %d\n", gpu.NumaNode)
		fmt.Fprintf(tw, "  Key: %d\n", gpu.Key)
	}
	return tw.Flush()
}

func (cmd *info) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.Client()
	if err != nil {
		return err
	}

	vm, err := cmd.VirtualMachine()
	if err != nil {
		return err
	}

	if vm == nil {
		return flag.ErrHelp
	}

	var o mo.VirtualMachine
	pc := property.DefaultCollector(c)
	err = pc.RetrieveOne(ctx, vm.Reference(), []string{"name", "config.hardware", "runtime.powerState"}, &o)
	if err != nil {
		return err
	}

	if o.Config == nil {
		return fmt.Errorf("VM configuration not available")
	}

	var res infoResult
	gpuCount := 0
	for _, device := range o.Config.Hardware.Device {
		if pciDevice, ok := device.(*types.VirtualPCIPassthrough); ok {
			gpu := gpuInfo{
				Index:    gpuCount,
				NumaNode: pciDevice.NumaNode,
				Key:      pciDevice.Key,
			}
			if desc := pciDevice.DeviceInfo.GetDescription(); desc != nil {
				gpu.Label = desc.Label
				gpu.Summary = desc.Summary
			}
			res.GPUs = append(res.GPUs, gpu)
			gpuCount++
		}
	}

	return cmd.WriteResult(&res)
}
