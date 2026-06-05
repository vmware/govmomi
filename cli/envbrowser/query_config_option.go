// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package envbrowser

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vim25/xml"
)

type queryConfigOption struct {
	command
	hardwareVersion     string
	guestIDs            string
	allHardwareVersions bool
}

func init() {
	cli.Register("envbrowser.query-config-option", &queryConfigOption{})
}

func (cmd *queryConfigOption) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.command.Register(ctx, f)

	f.StringVar(
		&cmd.hardwareVersion,
		"hardware-version",
		"vmx-19",
		"The hardware version to query.")

	f.StringVar(
		&cmd.guestIDs,
		"guest-ids",
		"",
		"A comma-delimited list of guest IDs.")

	f.BoolVar(
		&cmd.allHardwareVersions,
		"all-hardware-versions",
		false,
		"True to query all hardware versions.")
}

func (cmd *queryConfigOption) Description() string {
	return `Query the environment browser for a config option.

Examples:
  govc envbrowser.query-config-option -cluster my-cluster
  govc envbrowser.query-config-option -cluster my-cluster -host my-host
  govc envbrowser.query-config-option -cluster my-cluster -hardware-version vmx-22
  govc envbrowser.query-config-option -cluster my-cluster -hardware-version vmx-22 -guest-ids otherLinuxGuest,ubuntu64Guest
  govc envbrowser.query-config-option -cluster my-cluster -all-hardware-versions
  govc envbrowser.query-config-option -cluster my-cluster -hardware-version vmx-22 -copy-to-file
  govc envbrowser.query-config-option -cluster my-cluster -all-hardware-versions -copy-to-file`
}

func (cmd *queryConfigOption) Run(ctx context.Context, f *flag.FlagSet) error {
	if err := cmd.command.Run(ctx, f); err != nil {
		return err
	}

	if cmd.allHardwareVersions {
		for _, v := range types.GetHardwareVersions() {
			if v < types.VMX10 {
				continue
			}
			r, err := cmd.eb.QueryConfigOption(
				ctx,
				&types.EnvironmentBrowserConfigOptionQuerySpec{
					Key:     v.String(),
					GuestId: strings.Split(cmd.guestIDs, ","),
				})
			if err != nil {
				return fmt.Errorf(
					"failed to get config options for key=%q, guestId=%q: %w",
					v.String(),
					cmd.guestIDs,
					err)
			}
			if err := cmd.WriteResult(queryConfigOptionResult{
				r:          r,
				copyToFile: cmd.copyToFile,
			}); err != nil {

				return err
			}
		}
		return nil
	}

	var hostRef *types.ManagedObjectReference
	if cmd.host != nil {
		ref := cmd.host.Reference()
		hostRef = &ref
	}

	r, err := cmd.eb.QueryConfigOption(
		ctx,
		&types.EnvironmentBrowserConfigOptionQuerySpec{
			Key:     cmd.hardwareVersion,
			GuestId: strings.Split(cmd.guestIDs, ","),
			Host:    hostRef,
		})
	if err != nil {
		return fmt.Errorf(
			"failed to get config options for key=%q, guestId=%q: %w",
			cmd.hardwareVersion,
			cmd.guestIDs,
			err)
	}

	return cmd.WriteResult(queryConfigOptionResult{
		r:          r,
		copyToFile: cmd.copyToFile,
	})
}

type queryConfigOptionResult struct {
	r          *types.VirtualMachineConfigOption
	copyToFile bool
}

func (r queryConfigOptionResult) Write(w io.Writer) error {
	fmt.Println(r.r.Version)
	if r.copyToFile {

		f, err := os.Create("config-option-" + r.r.Version + ".xml")
		if err != nil {
			return err
		}
		defer f.Close()
		enc := xml.NewEncoder(f)
		enc.Indent("", "  ")
		return enc.Encode(r.r)
	}
	return nil
}
