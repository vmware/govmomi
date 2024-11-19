/*
Copyright (c) 2024-2024 VMware, Inc. All Rights Reserved.

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

package tpm

import (
	"context"
	"flag"
	"fmt"
	"io"
	"reflect"
	"strings"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/types"
)

type report struct {
	*flags.HostSystemFlag

	e bool
}

func init() {
	cli.Register("host.tpm.report", &report{})
}

func (cmd *report) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.HostSystemFlag, ctx = flags.NewHostSystemFlag(ctx)
	cmd.HostSystemFlag.Register(ctx, f)

	f.BoolVar(&cmd.e, "e", false, "Print events")
}

func (cmd *report) Description() string {
	return `Trusted Platform Module report.

Examples:
  govc host.tpm.report
  govc host.tpm.report -e
  govc host.tpm.report -json`
}

func (cmd *report) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.Client()
	if err != nil {
		return err
	}

	host, err := cmd.HostSystem()
	if err != nil {
		return err
	}

	query := types.QueryTpmAttestationReport{This: host.Reference()}
	report, err := methods.QueryTpmAttestationReport(ctx, c, &query)
	if err != nil {
		return err
	}

	return cmd.WriteResult(&reportResult{report.Returnval, cmd})
}

type reportResult struct {
	Report *types.HostTpmAttestationReport
	cmd    *report
}

func (r *reportResult) Write(w io.Writer) error {
	if r.Report == nil {
		return nil
	}

	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)

	if r.cmd.e {
		for _, e := range r.Report.TpmEvents {
			pcr := e.PcrIndex
			d := e.EventDetails.GetHostTpmEventDetails()
			meth := d.DataHashMethod
			hash := d.DataHash
			var name string

			switch x := e.EventDetails.(type) {
			case *types.HostTpmBootSecurityOptionEventDetails:
				name = x.BootSecurityOption
			case *types.HostTpmSoftwareComponentEventDetails:
				name = x.ComponentName
			case *types.HostTpmCommandEventDetails:
				name = x.CommandLine
			case *types.HostTpmSignerEventDetails:
				name = x.BootSecurityOption
			case *types.HostTpmVersionEventDetails:
				name = fmt.Sprintf("%x", x.Version)
			case *types.HostTpmOptionEventDetails:
				name = x.OptionsFileName
			case *types.HostTpmBootCompleteEventDetails:
			}

			kind := reflect.ValueOf(e.EventDetails).Elem().Type().Name()
			kind = strings.TrimPrefix(strings.TrimSuffix(kind, "EventDetails"), "HostTpm")

			fmt.Fprintf(tw, "%d\t%s\t%s\t%x\t%s\n", pcr, kind, meth, hash, name)
		}
	} else {
		for _, e := range r.Report.TpmPcrValues {
			fmt.Fprintf(tw, "PCR %d\t%s\t%x\t%s\n", e.PcrNumber, e.DigestMethod, e.DigestValue, e.ObjectName)
		}
	}

	return tw.Flush()
}
