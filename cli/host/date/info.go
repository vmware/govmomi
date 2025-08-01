// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package date

import (
	"context"
	"flag"
	"fmt"
	"io"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/cli/host/service"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type info struct {
	*flags.HostSystemFlag
	*flags.OutputFlag
}

func init() {
	cli.Register("host.date.info", &info{})
}

func (cmd *info) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.HostSystemFlag, ctx = flags.NewHostSystemFlag(ctx)
	cmd.HostSystemFlag.Register(ctx, f)

	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)
}

func (cmd *info) Description() string {
	return `Display date and time info for HOST.`
}

func (cmd *info) Process(ctx context.Context) error {
	if err := cmd.HostSystemFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

type dateInfo struct {
	types.HostDateTimeInfo
	Service *types.HostService `json:"service"`
	Current *time.Time         `json:"current"`
}

func (info *dateInfo) servers() string {
	if len(info.NtpConfig.Server) == 0 {
		return "None"
	}
	return strings.Join(info.NtpConfig.Server, ", ")
}

func (info *dateInfo) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)

	fmt.Fprintf(tw, "Current date and time:\t%s\n", info.Current.Format(time.UnixDate))
	if info.Service != nil {
		fmt.Fprintf(tw, "NTP client status:\t%s\n", service.Policy(*info.Service))
		fmt.Fprintf(tw, "NTP service status:\t%s\n", service.Status(*info.Service))
	}
	fmt.Fprintf(tw, "NTP servers:\t%s\n", info.servers())

	return tw.Flush()
}

func (cmd *info) Run(ctx context.Context, f *flag.FlagSet) error {
	host, err := cmd.HostSystem()
	if err != nil {
		return err
	}

	s, err := host.ConfigManager().DateTimeSystem(ctx)
	if err != nil {
		return err
	}

	var hs mo.HostDateTimeSystem
	if err = s.Properties(ctx, s.Reference(), nil, &hs); err != nil {
		return nil
	}

	ss, err := host.ConfigManager().ServiceSystem(ctx)
	if err != nil {
		return err
	}

	services, err := ss.Service(ctx)
	if err != nil {
		return err
	}

	res := &dateInfo{HostDateTimeInfo: hs.DateTimeInfo}

	for i, service := range services {
		if service.Key == "ntpd" {
			res.Service = &services[i]
			break
		}
	}

	res.Current, err = s.Query(ctx)
	if err != nil {
		return err
	}

	return cmd.WriteResult(res)
}
