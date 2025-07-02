// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package esxcli

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/dougm/pretty"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/esx"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/internal"
)

type esxcli struct {
	*flags.HostSystemFlag

	hints bool
	trace bool
}

func init() {
	cli.Register("host.esxcli", &esxcli{})
}

func (cmd *esxcli) Usage() string {
	return "COMMAND [ARG]..."
}

func (cmd *esxcli) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.HostSystemFlag, ctx = flags.NewHostSystemFlag(ctx)
	cmd.HostSystemFlag.Register(ctx, f)

	f.BoolVar(&cmd.hints, "hints", true, "Use command info hints when formatting output")
	if cli.ShowUnreleased() {
		f.BoolVar(&cmd.trace, "T", false, "Write esxcli nested SOAP traffic to stderr")
	}
}

func (cmd *esxcli) Description() string {
	return `Invoke esxcli command on HOST.

Output is rendered in table form when possible, unless disabled with '-hints=false'.

Examples:
  govc host.esxcli network ip connection list
  govc host.esxcli system settings advanced set -o /Net/GuestIPHack -i 1
  govc host.esxcli network firewall ruleset set -r remoteSerialPort -e true
  govc host.esxcli network firewall set -e false
  govc host.esxcli hardware platform get`
}

func (cmd *esxcli) Process(ctx context.Context) error {
	if err := cmd.HostSystemFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func fmtXML(s string) {
	cmd := exec.Command("xmlstarlet", "fo")
	cmd.Stdout = os.Stderr
	cmd.Stderr = os.Stderr
	cmd.Stdin = strings.NewReader(s)
	if err := cmd.Run(); err != nil {
		panic(err) // yes xmlstarlet is required (your eyes will thank you)
	}
}

func (cmd *esxcli) Trace(req *internal.ExecuteSoapRequest, res *internal.ExecuteSoapResponse) {
	x := res.Returnval

	if req.Moid == "ha-dynamic-type-manager-local-cli-cliinfo" {
		if x.Fault == nil {
			return // TODO: option to trace this
		}
	}

	pretty.Fprintf(os.Stderr, "%# v\n", req)

	if x.Fault == nil {
		fmtXML(res.Returnval.Response)
	} else {
		fmt.Fprintln(os.Stderr, "Message=", x.Fault.FaultMsg)
		if x.Fault.FaultDetail != "" {
			fmt.Fprint(os.Stderr, "Detail=")
			fmtXML(x.Fault.FaultDetail)
		}
	}
}

func (cmd *esxcli) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() == 0 {
		return flag.ErrHelp
	}

	c, err := cmd.Client()
	if err != nil {
		return err
	}

	host, err := cmd.HostSystem()
	if err != nil {
		return err
	}

	e, err := esx.NewExecutor(ctx, c, host)
	if err != nil {
		return err
	}
	if cmd.trace {
		e.Trace = cmd.Trace
	}

	res, err := e.Run(ctx, f.Args())
	if err != nil {
		if f, ok := err.(*esx.Fault); ok {
			return errors.New(f.MessageDetail())
		}
		return err
	}

	if len(res.Values) == 0 {
		if res.String != "" {
			fmt.Print(res.String)
			if !strings.HasSuffix(res.String, "\n") {
				fmt.Println()
			}
		}
		return nil
	}

	return cmd.WriteResult(&result{res, cmd})
}

type result struct {
	*esx.Response
	cmd *esxcli
}

func (r *result) Dump() any {
	return r.Response
}

func (r *result) Write(w io.Writer) error {
	var formatType string
	if r.cmd.hints {
		formatType = r.Info.Hints.Formatter()
	}

	switch formatType {
	case "table":
		r.cmd.formatTable(w, r.Response)
	default:
		r.cmd.formatSimple(w, r.Response)
	}

	return nil
}

func (cmd *esxcli) formatSimple(w io.Writer, res *esx.Response) {
	var keys []string
	for key := range res.Values[0] {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)

	for i, rv := range res.Values {
		if i > 0 {
			fmt.Fprintln(tw)
			_ = tw.Flush()
		}
		for _, key := range keys {
			fmt.Fprintf(tw, "%s:\t%s\n", key, strings.Join(rv[key], ", "))
		}
	}

	_ = tw.Flush()
}

func (cmd *esxcli) formatTable(w io.Writer, res *esx.Response) {
	fields := res.Info.Hints.Fields()
	if len(fields) == 0 {
		cmd.formatSimple(w, res)
		return
	}
	tw := tabwriter.NewWriter(w, len(fields), 0, 2, ' ', 0)

	var hr []string
	for _, name := range fields {
		hr = append(hr, strings.Repeat("-", len(name)))
	}

	fmt.Fprintln(tw, strings.Join(fields, "\t"))
	fmt.Fprintln(tw, strings.Join(hr, "\t"))

	for _, vals := range res.Values {
		var row []string

		for _, name := range fields {
			key := strings.Replace(name, " ", "", -1)
			if val, ok := vals[key]; ok {
				row = append(row, strings.Join(val, ", "))
			} else {
				row = append(row, "")
			}
		}

		fmt.Fprintln(tw, strings.Join(row, "\t"))
	}

	_ = tw.Flush()
}
