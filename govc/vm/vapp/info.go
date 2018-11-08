package vapp

import (
	"context"
	"flag"
	"fmt"
	"io"
	"strings"
	"text/tabwriter"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vim25/types"
)

const infoCmdUsage = "VM..."
const infoCmdDescription = "Displays the vApp configuration for specified virtual machines."

type infoResult struct {
	Name string
	Info *types.VmConfigInfo
}

type infoResults struct {
	Items []infoResult
}

func (r *infoResults) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)
	for _, item := range r.Items {
		if err := writeInfoResult(item, tw); err != nil {
			return err
		}
	}
	return tw.Flush()
}

func writeInfoResult(r infoResult, tw *tabwriter.Writer) error {
	fmt.Fprintf(tw, "VM Name: %s\n", r.Name)
	if r.Info != nil {
		return writeInfoResultDetail(r, tw)
	}
	fmt.Fprintf(tw, "  <No vApp Configuration>\n\n")
	return nil
}

func writeInfoResultDetail(r infoResult, tw *tabwriter.Writer) error {
	// Not sure under what instances r.Info.VAppProductInfo would have more than
	// one element, but we only display the first one for now.
	product := r.Info.Product[0]
	fmt.Fprintf(tw, "  Product Name:\t%s\n", product.Name)
	fmt.Fprintf(tw, "  Product Version:\t%s\n", product.Version)
	fmt.Fprintf(tw, "  Product Full Version:\t%s\n", product.FullVersion)
	fmt.Fprintf(tw, "  Product URL:\t%s\n", product.ProductUrl)
	fmt.Fprintf(tw, "  Product Vendor:\t%s\n", product.Vendor)
	fmt.Fprintf(tw, "  Product Vendor URL:\t%s\n", product.VendorUrl)
	fmt.Fprintf(tw, "  Product Application URL:\t%s\n", product.AppUrl)
	fmt.Fprintln(tw, "")

	assignment := r.Info.IpAssignment
	fmt.Fprintf(tw, "  Current IP Allocation Policy:\t%s\n", assignment.IpAllocationPolicy)
	fmt.Fprintf(tw, "  Current IP Protocol:\t%s\n", assignment.IpProtocol)
	fmt.Fprintf(tw, "  Supported IP Allocation Schemes:\t%s\n", strings.Join(assignment.SupportedAllocationScheme, ","))
	fmt.Fprintf(tw, "  Supported IP Protocols:\t%s\n", strings.Join(assignment.SupportedIpProtocol, ","))
	fmt.Fprintln(tw, "")

	fmt.Fprintf(tw, "  Supported OVF Transports:\t%s\n", strings.Join(r.Info.OvfEnvironmentTransport, ","))
	fmt.Fprintf(tw, "  Installation Boot Required:\t%t\n", r.Info.InstallBootRequired)
	fmt.Fprintf(tw, "  Installation Boot Delay:\t%d\n", r.Info.InstallBootStopDelay)
	fmt.Fprintln(tw, "")

	var propertyIDs []string
	for _, v := range r.Info.Property {
		propertyIDs = append(propertyIDs, v.Id)
	}
	fmt.Fprintf(tw, "  Application Properties:\t%s\n\n", strings.Join(propertyIDs, "\n\t"))
	return nil
}

type info struct {
	*flags.VirtualMachineFlag
}

func init() {
	cli.Register("vm.vapp.info", &info{})
}

func (cmd *info) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)
}

func (cmd *info) Usage() string {
	return infoCmdUsage
}

func (cmd *info) Description() string {
	return infoCmdDescription
}

func (cmd *info) Process(ctx context.Context) error {
	return cmd.VirtualMachineFlag.Process(ctx)
}

func (cmd *info) Run(ctx context.Context, f *flag.FlagSet) error {
	vms, err := cmd.VirtualMachines(f.Args())
	if err != nil {
		return err
	}

	var res infoResults
	for _, vm := range vms {
		info, err := retrieveVAppConfig(ctx, vm)
		if err != nil {
			return err
		}
		r := infoResult{
			Name: vm.Name(),
			Info: info,
		}
		res.Items = append(res.Items, r)
	}

	return cmd.WriteResult(&res)
}
