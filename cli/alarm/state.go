// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package alarm

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"text/tabwriter"
	"time"

	"github.com/vmware/govmomi/alarm"
	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vim25/types"
)

type state struct {
	*flags.DatacenterFlag
	*flags.OutputFlag

	ack  bool
	name string
	alarm.StateInfoOptions
}

func init() {
	cli.Register("alarms", &state{})
}

func (cmd *state) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatacenterFlag, ctx = flags.NewDatacenterFlag(ctx)
	cmd.DatacenterFlag.Register(ctx, f)

	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)

	f.BoolVar(&cmd.ack, "ack", false, "Acknowledge alarms")
	f.StringVar(&cmd.name, "n", "", "Filter by alarm name")
	f.BoolVar(&cmd.Declared, "d", false, "Show declared alarms")
	f.BoolVar(&cmd.InventoryPath, "l", false, "Long listing output")
}

func (cmd *state) Usage() string {
	return "[PATH]"
}

func (cmd *state) Description() string {
	return `Show triggered or declared alarms.

Triggered alarms: alarms triggered by this entity or by its descendants.
Triggered alarms are propagated up the inventory hierarchy so that a user
can readily tell when a descendant has triggered an alarm.

Declared alarms: alarms that apply to this managed entity.
Includes alarms defined on this entity and alarms inherited from the parent
entity, or from any ancestors in the inventory hierarchy.

PATH defaults to the root folder '/'.
When PATH is provided it should be an absolute inventory path or relative
to GOVC_DATACENTER. See also:
  govc find -h
  govc tree -h

Examples:
  govc alarms
  govc alarms vm/folder/vm-name
  govc alarms /dc1/host/cluster1
  govc alarms /dc1/host/cluster1 -d

  govc alarms -n alarm.WCPRegisterVMFailedAlarm
  govc alarms -ack /dc1/host/cluster1
  govc alarms -ack -n alarm.WCPRegisterVMFailedAlarm vm/vm-name`
}

func (cmd *state) Process(ctx context.Context) error {
	if err := cmd.DatacenterFlag.Process(ctx); err != nil {
		return err
	}
	return cmd.OutputFlag.Process(ctx)
}

func (cmd *state) filterAlarms(ctx context.Context, m *alarm.Manager, obj types.ManagedObjectReference) ([]alarm.StateInfo, error) {
	alarms, err := m.GetStateInfo(ctx, obj, cmd.StateInfoOptions)
	if err != nil || cmd.name == "" {
		return alarms, err
	}

	var match []alarm.StateInfo
	for _, alarm := range alarms {
		if alarm.Info.SystemName == cmd.name || alarm.Info.Name == cmd.name {
			match = append(match, alarm)
		}
	}

	return match, nil
}

type stateResult struct {
	info []alarm.StateInfo
	cmd  *state
}

func (r *stateResult) Dump() any {
	return r.info
}

func (r *stateResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.info)
}

func (r *stateResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)

	fmt.Fprintf(tw, "Alarm Name\tObject\tSeverity\tTriggered Time\tAcknowledged Time\tAcknowledged By\n")

	for _, a := range r.info {
		name := a.Info.Name
		tt := a.Time.Format(time.Stamp)
		at := ""
		by := a.AcknowledgedByUser
		if a.AcknowledgedTime != nil {
			at = a.AcknowledgedTime.Format(time.Stamp)
		}
		fmt.Fprintf(tw, "%s\t%s\t%s\t%s\t%s\t%s\n", name, a.Path, alarm.Severity[a.OverallStatus], tt, at, by)
	}

	return tw.Flush()
}

func (cmd *state) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() > 1 {
		return flag.ErrHelp
	}
	if cmd.ack && cmd.Declared {
		return flag.ErrHelp
	}

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

	cmd.StateInfoOptions.Event = cmd.All()

	m, err := alarm.GetManager(c)
	if err != nil {
		return err
	}

	alarms, err := cmd.filterAlarms(ctx, m, obj)
	if err != nil {
		return err
	}

	if cmd.ack {
		for _, alarm := range alarms {
			if alarm.Acknowledged != nil && *alarm.Acknowledged {
				continue
			}
			if err := m.AcknowledgeAlarm(ctx, alarm.Alarm, alarm.Entity); err != nil {
				return err
			}
		}

		alarms, err = cmd.filterAlarms(ctx, m, obj)
		if err != nil {
			return err
		}
	}

	return cmd.WriteResult(&stateResult{alarms, cmd})
}
