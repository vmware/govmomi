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

package alarm

import (
	"context"
	"flag"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/vmware/govmomi/alarm"
	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vim25/mo"
)

type info struct {
	*flags.DatacenterFlag

	name flags.StringList
}

func init() {
	cli.Register("alarm.info", &info{})
}

func (cmd *info) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatacenterFlag, ctx = flags.NewDatacenterFlag(ctx)
	cmd.DatacenterFlag.Register(ctx, f)

	f.Var(&cmd.name, "n", "Alarm name")
}

func (cmd *info) Usage() string {
	return "PATH"
}

func (cmd *info) Description() string {
	return `Alarm definition info.

Examples:
  govc alarm.info
  govc alarm.info /dc1/host/cluster1
  govc alarm.info -n alarm.WCPRegisterVMFailedAlarm`
}

type infoResult []mo.Alarm

func (r infoResult) Dump() any {
	return []mo.Alarm(r)
}

func (r infoResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)
	for _, a := range r {
		fmt.Fprintf(tw, "Name:\t%s\n", a.Info.Name)
		fmt.Fprintf(tw, "  SystemName:\t%s\n", a.Info.SystemName)
		fmt.Fprintf(tw, "  Description:\t%s\n", a.Info.Description)
		fmt.Fprintf(tw, "  Enabled:\t%t\n", a.Info.Enabled)
	}
	return tw.Flush()
}

func (cmd *info) findAlarm(alarms []mo.Alarm) []mo.Alarm {
	if len(cmd.name) == 0 {
		return alarms
	}
	var match []mo.Alarm
	for _, alarm := range alarms {
		for _, name := range cmd.name {
			if alarm.Info.SystemName == name || alarm.Info.Name == name {
				match = append(match, alarm)
			}
		}
	}
	return match
}

func (cmd *info) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.Client()
	if err != nil {
		return err
	}

	obj := c.ServiceContent.RootFolder
	if f.NArg() == 1 {
		obj, err = cmd.ManagedObject(ctx, f.Arg(0))
		if err != nil {
			return err
		}
	}

	m, err := alarm.GetManager(c)
	if err != nil {
		return err
	}

	alarms, err := m.GetAlarm(ctx, obj)
	if err != nil {
		return err
	}

	return cmd.WriteResult(infoResult(cmd.findAlarm(alarms)))
}
