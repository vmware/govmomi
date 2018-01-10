package vapp

import (
	"context"
	"errors"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vim25/types"
)

var allowedIPProtocols = []string{
	string(types.VAppIPAssignmentInfoProtocolsIPv4),
	string(types.VAppIPAssignmentInfoProtocolsIPv6),
}

var allowedAllocationPolicies = []string{
	string(types.VAppIPAssignmentInfoIpAllocationPolicyDhcpPolicy),
	string(types.VAppIPAssignmentInfoIpAllocationPolicyTransientPolicy),
	string(types.VAppIPAssignmentInfoIpAllocationPolicyFixedPolicy),
	string(types.VAppIPAssignmentInfoIpAllocationPolicyFixedAllocatedPolicy),
}

var allowedAllocationSchemes = []string{
	string(types.VAppIPAssignmentInfoAllocationSchemesDhcp),
	string(types.VAppIPAssignmentInfoAllocationSchemesOvfenv),
}

var allowedOVFTransports = []string{
	"iso",
	"com.vmware.guestInfo",
}

const configureCmdDescription = `Configures general vApp options on a virtual machine.

List values are separated by commas.

To remove vApp product properties, specify -product.remove. All relevant
product properties need to be re-added after this operation - any product
properties added at the same time are ignored.

Examples:

govc vm.vapp.configure -vm=foobar -ip.protocols.supported="IPv4,IPv6"
govc vm.vapp.configure -vm=foobar -product.remove=true
govc vm.vapp.configure -vm=foobar -product.name="Foo Bar"
govc vm.vapp.configure -vm=foobar -product.name="Foo Bar" -product.version="1.0.0"

`

type configure struct {
	*flags.VirtualMachineFlag

	types.VmConfigSpec

	removeProduct bool
}

func init() {
	cli.Register("vm.vapp.configure", &configure{})
}

func (cmd *configure) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)

	cmd.VmConfigSpec.Product = make([]types.VAppProductSpec, 1)
	cmd.VmConfigSpec.Product[0].Operation = types.ArrayUpdateOperationEdit
	cmd.VmConfigSpec.Product[0].Info = new(types.VAppProductInfo)
	cmd.VmConfigSpec.IpAssignment = new(types.VAppIPAssignmentInfo)

	f.StringVar(&cmd.VmConfigSpec.Product[0].Info.Name, "product.name", "", "Product name")
	f.StringVar(&cmd.VmConfigSpec.Product[0].Info.Version, "product.version", "", "Product version")
	f.StringVar(&cmd.VmConfigSpec.Product[0].Info.FullVersion, "product.fullversion", "", "Long product version")
	f.StringVar(&cmd.VmConfigSpec.Product[0].Info.ProductUrl, "product.producturl", "", "Product URL")
	f.StringVar(&cmd.VmConfigSpec.Product[0].Info.Vendor, "product.vendor", "", "Product vendor")
	f.StringVar(&cmd.VmConfigSpec.Product[0].Info.VendorUrl, "product.vendorurl", "", "Product vendor URL")
	f.StringVar(&cmd.VmConfigSpec.Product[0].Info.AppUrl, "product.appurl", "", "Product application URL")

	f.BoolVar(&cmd.removeProduct, "product.remove", false, "Remove product properties")

	cmd.VmConfigSpec.IpAssignment = new(types.VAppIPAssignmentInfo)

	f.Var(
		flags.NewEnum(&cmd.VmConfigSpec.IpAssignment.IpAllocationPolicy, allowedAllocationPolicies),
		"ip.allocation.selected",
		"IP allocation policy to use",
	)
	f.Var(
		flags.NewEnumSlice(&cmd.VmConfigSpec.IpAssignment.SupportedAllocationScheme, allowedAllocationSchemes),
		"ip.scheme.supported",
		"Supported IP allocation schemes",
	)
	f.Var(
		flags.NewEnum(&cmd.VmConfigSpec.IpAssignment.IpProtocol, allowedIPProtocols),
		"ip.protocols.selected",
		"IP protocol to use",
	)
	f.Var(
		flags.NewEnumSlice(&cmd.VmConfigSpec.IpAssignment.SupportedIpProtocol, allowedIPProtocols),
		"ip.protocols.supported",
		"Supported IP protocols",
	)

	f.Var(flags.NewEnumSlice(&cmd.VmConfigSpec.OvfEnvironmentTransport, allowedOVFTransports), "ovf.transports", "Supported OVF transports")
	f.Var(flags.NewOptionalBool(&cmd.VmConfigSpec.InstallBootRequired), "ovf.boot.required", "Install boot required")
	f.Var(flags.NewInt32(&cmd.VmConfigSpec.InstallBootStopDelay), "ovf.boot.delay", "Install boot delay")
}

func (cmd *configure) Description() string {
	return configureCmdDescription
}

func (cmd *configure) Process(ctx context.Context) error {
	return cmd.VirtualMachineFlag.Process(ctx)
}

func (cmd *configure) Run(ctx context.Context, f *flag.FlagSet) error {
	vm, err := cmd.VirtualMachine()
	if err != nil {
		return err
	}

	if vm == nil {
		return errors.New("please specify a virtual machine")
	}

	if cmd.removeProduct {
		cmd.flagProductRemove()
	}

	spec := types.VirtualMachineConfigSpec{
		VAppConfig: &cmd.VmConfigSpec,
	}

	task, err := vm.Reconfigure(ctx, spec)
	if err != nil {
		return err
	}

	if err := task.Wait(ctx); err != nil {
		return err
	}

	fmt.Fprintf(cmd, "vApp configuration successfully updated for virtual machine %s.\n", vm.Name())
	return nil
}

func (cmd *configure) flagProductRemove() {
	zero := int32(0)
	op := types.VAppProductSpec{
		ArrayUpdateSpec: types.ArrayUpdateSpec{
			Operation: types.ArrayUpdateOperationRemove,
			RemoveKey: &zero,
		},
	}
	cmd.VmConfigSpec.Product = []types.VAppProductSpec{op}
}
