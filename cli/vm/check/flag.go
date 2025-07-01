// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package check

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vim25/xml"
)

var checkTestTypesList = []string{
	string(types.CheckTestTypeDatastoreTests),
	string(types.CheckTestTypeHostTests),
	string(types.CheckTestTypeNetworkTests),
	string(types.CheckTestTypeResourcePoolTests),
	string(types.CheckTestTypeSourceTests),
}

type checkTestTypes []types.CheckTestType

func (c *checkTestTypes) String() string {
	return fmt.Sprint(*c)
}

func (c *checkTestTypes) Set(value string) error {
	if !slices.Contains(checkTestTypesList, value) {
		return fmt.Errorf("invalid CheckTestType value %q", value)
	}
	*c = append(*c, types.CheckTestType(value))
	return nil
}

type checkFlag struct {
	*flags.VirtualMachineFlag
	*flags.HostSystemFlag
	*flags.ResourcePoolFlag

	Machine, Host, Pool *types.ManagedObjectReference

	testTypes checkTestTypes
}

func (cmd *checkFlag) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)
	cmd.HostSystemFlag, ctx = flags.NewHostSystemFlag(ctx)
	cmd.HostSystemFlag.Register(ctx, f)
	cmd.ResourcePoolFlag, ctx = flags.NewResourcePoolFlag(ctx)
	cmd.ResourcePoolFlag.Register(ctx, f)

	f.Var(&cmd.testTypes, "test", fmt.Sprintf("The set of tests to run (%s)", strings.Join(checkTestTypesList, ",")))
}

func (cmd *checkFlag) Process(ctx context.Context) error {
	if err := cmd.VirtualMachineFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.HostSystemFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.ResourcePoolFlag.Process(ctx); err != nil {
		return err
	}

	vm, err := cmd.VirtualMachine()
	if err != nil {
		return err
	}
	if vm != nil {
		cmd.Machine = types.NewReference(vm.Reference())
	}

	host, err := cmd.HostSystemIfSpecified()
	if err != nil {
		return err
	}
	if host != nil {
		cmd.Host = types.NewReference(host.Reference())
	}

	pool, err := cmd.ResourcePoolIfSpecified()
	if err != nil {
		return err
	}
	if pool != nil {
		cmd.Pool = types.NewReference(pool.Reference())
	}

	return nil
}

func (cmd *checkFlag) provChecker() (*object.VmProvisioningChecker, error) {
	c, err := cmd.VirtualMachineFlag.Client()
	if err != nil {
		return nil, err
	}

	return object.NewVmProvisioningChecker(c), nil
}

func (cmd *checkFlag) compatChecker() (*object.VmCompatibilityChecker, error) {
	c, err := cmd.VirtualMachineFlag.Client()
	if err != nil {
		return nil, err
	}

	return object.NewVmCompatibilityChecker(c), nil
}

func (cmd *checkFlag) Spec(spec any) error {
	dec := xml.NewDecoder(os.Stdin)
	dec.TypeFunc = types.TypeFunc()
	return dec.Decode(spec)
}

// return cmd.VirtualMachineFlag.WriteResult(&checkResult{res, ctx, cmd.VirtualMachineFlag})
func (cmd *checkFlag) result(ctx context.Context, res []types.CheckResult) error {
	return cmd.VirtualMachineFlag.WriteResult(&checkResult{res, ctx, cmd.VirtualMachineFlag})
}

type checkResult struct {
	Result []types.CheckResult `json:"result"`
	ctx    context.Context
	vm     *flags.VirtualMachineFlag
}

func (res *checkResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)

	c, err := res.vm.Client()
	if err != nil {
		return err
	}

	for _, r := range res.Result {
		fields := []struct {
			name  string
			moid  *types.ManagedObjectReference
			fault []types.LocalizedMethodFault
		}{
			{"VM", r.Vm, nil},
			{"Host", r.Host, nil},
			{"Warning", nil, r.Warning},
			{"Error", nil, r.Error},
		}

		for _, f := range fields {
			var val string
			if f.moid == nil {
				var msgs []string
				for _, m := range f.fault {
					msgs = append(msgs, m.LocalizedMessage)
				}
				val = strings.Join(slices.Compact(msgs), "\n\t")
			} else {
				val, err = find.InventoryPath(res.ctx, c, *f.moid)
				if err != nil {
					return err
				}
			}
			fmt.Fprintf(tw, "%s:\t%s\n", f.name, val)
		}
	}

	return tw.Flush()
}
