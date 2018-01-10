package vapp

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vim25/types"
)

const propertyInfoCmdUsage = "KEYS..."
const propertyInfoCmdDescription = "Display detailed info on a virtual machine's vApp configuration properties."

type propertyResult types.VAppPropertyInfo

type propertyResults struct {
	Items []propertyResult
}

func (r *propertyResults) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)
	for _, item := range r.Items {
		if err := writePropertyResult(item, tw); err != nil {
			return err
		}
	}
	return tw.Flush()
}

func writePropertyResult(r propertyResult, tw *tabwriter.Writer) error {
	var configurable bool
	if r.UserConfigurable != nil {
		configurable = *r.UserConfigurable
	}

	fmt.Fprintf(tw, "Property: %s\n", r.Id)
	fmt.Fprintf(tw, "  Category:\t%s\n", r.Category)
	fmt.Fprintf(tw, "  Label:\t%s\n", r.Label)
	fmt.Fprintf(tw, "  Key Class ID:\t%s\n", r.ClassId)
	fmt.Fprintf(tw, "  Key:\t%d\n", r.Key)
	fmt.Fprintf(tw, "  Key Instance ID:\t%s\n", r.InstanceId)
	fmt.Fprintf(tw, "  Description:\t%s\n", r.Description)
	fmt.Fprintf(tw, "  Type:\t%s\n", r.Type)
	fmt.Fprintf(tw, "  Type Reference:\t%s\n", r.TypeReference)
	fmt.Fprintf(tw, "  User Configurable:\t%t\n", configurable)
	fmt.Fprintf(tw, "  Current Value:\t%s\n\n", r.Value)

	return nil
}

type propertyInfo struct {
	*flags.VirtualMachineFlag
}

func init() {
	cli.Register("vm.vapp.property.info", &propertyInfo{})
}

func (cmd *propertyInfo) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)
}

func (cmd *propertyInfo) Usage() string {
	return propertyInfoCmdUsage
}

func (cmd *propertyInfo) Description() string {
	return propertyInfoCmdDescription
}

func (cmd *propertyInfo) Process(ctx context.Context) error {
	return cmd.VirtualMachineFlag.Process(ctx)
}

func (cmd *propertyInfo) Run(ctx context.Context, f *flag.FlagSet) error {
	vm, err := cmd.VirtualMachine()
	if err != nil {
		return err
	}

	keys := f.Args()
	if len(keys) < 1 {
		return errors.New("no keys were specified")
	}

	cfg, err := retrieveVAppConfig(ctx, vm)
	if err != nil {
		return err
	}
	if cfg == nil {
		return fmt.Errorf("%s has no vApp configuration", vm.Name())
	}

	var res propertyResults
	for _, k := range keys {
		pi, err := findProperty(cfg.Property, k)
		if err != nil {
			return err
		}
		res.Items = append(res.Items, propertyResult(pi))
	}

	return cmd.WriteResult(&res)
}

func findProperty(ps []types.VAppPropertyInfo, key string) (types.VAppPropertyInfo, error) {
	for _, pi := range ps {
		if pi.Id == key {
			return pi, nil
		}
	}
	return types.VAppPropertyInfo{}, fmt.Errorf("property not found: %s", key)
}
