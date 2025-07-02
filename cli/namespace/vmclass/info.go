// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vmclass

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/vmware/govmomi/units"
	"github.com/vmware/govmomi/vapi/namespace"
	"github.com/vmware/govmomi/vim25/types"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
)

type infoResult namespace.VirtualMachineClassInfo

func (r infoResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)

	fmt.Fprintf(tw, "ID:\t%s\n", r.Id)
	fmt.Fprintf(tw, "CPUs:\t%d\n", r.CpuCount)
	fmt.Fprintf(tw, "Memory:\t%s\n", units.ByteSize(r.MemoryMb*1024*1024))
	if len(r.ConfigSpec) != 0 {
		// While ConfigSpec already is JSON, vmodl keys are sorted.
		// Use vim25/types.NewJSONEncoder for consistency, order is as
		// defined in the VirtualMachineConfigSpec struct.
		spec, err := configSpecFromJSON(r.ConfigSpec)
		if err != nil {
			return err
		}
		buf, err := configSpecToJSON(spec)
		if err != nil {
			return err
		}
		fmt.Fprintf(tw, "ConfigSpec:\t%s\n", string(buf))
		fmt.Fprintf(tw, "ConfigSpecHash:\t%s\n", configSpecSHA(buf))
	}

	return tw.Flush()
}

func configSpecFromJSON(m json.RawMessage) (*types.VirtualMachineConfigSpec, error) {
	var spec types.VirtualMachineConfigSpec
	err := types.NewJSONDecoder(bytes.NewReader(m)).Decode(&spec)
	if err != nil {
		return nil, err
	}
	return &spec, nil
}

func configSpecToJSON(spec *types.VirtualMachineConfigSpec) ([]byte, error) {
	var buf bytes.Buffer
	err := types.NewJSONEncoder(&buf).Encode(spec)
	if err != nil {
		return nil, err
	}

	return bytes.TrimSuffix(buf.Bytes(), []byte{'\n'}), nil
}

func configSpecSHA(buf []byte) string {
	h := sha256.New()
	_, _ = h.Write(buf)
	return hex.EncodeToString(h.Sum(nil))[:17]
}

type info struct {
	*flags.ClientFlag
	*flags.OutputFlag
}

func init() {
	cli.Register("namespace.vmclass.info", &info{})
}

func (cmd *info) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)
}

func (cmd *info) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}

	return nil
}

func (cmd *info) Usage() string {
	return "NAME"
}

func (cmd *info) Description() string {
	return `Displays the details of a virtual machine class.

Examples:
  govc namespace.vmclass.info test-class`
}

func (cmd *info) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}

	rc, err := cmd.RestClient()
	if err != nil {
		return err
	}

	nm := namespace.NewManager(rc)

	d, err := nm.GetVmClass(ctx, f.Arg(0))
	if err != nil {
		return err
	}

	return cmd.WriteResult(infoResult(d))
}
