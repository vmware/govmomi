// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package alarm

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/alarm"
	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type create struct {
	*flags.DatacenterFlag

	types.AlarmSpec

	r    bool
	kind string

	green, yellow, red string
}

func init() {
	cli.Register("alarm.create", &create{}, true)
}

func (cmd *create) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatacenterFlag, ctx = flags.NewDatacenterFlag(ctx)
	cmd.DatacenterFlag.Register(ctx, f)

	f.StringVar(&cmd.AlarmSpec.Name, "n", "", "Alarm name")
	f.StringVar(&cmd.AlarmSpec.Description, "d", "", "Alarm description")
	f.BoolVar(&cmd.Enabled, "enabled", true, "Enabled")

	f.StringVar(&cmd.kind, "type", "VirtualMachine", "Object type")
	f.BoolVar(&cmd.r, "r", false, "Reconfigure existing alarm")

	f.StringVar(&cmd.green, "green", "", "green status event type")
	f.StringVar(&cmd.yellow, "yellow", "", "yellow status event type")
	f.StringVar(&cmd.red, "red", "", "red status event type")
}

func (cmd *create) Usage() string {
	return "[PATH]"
}

func (cmd *create) Description() string {
	return `Create alarm.

Examples:
  govc alarm.create -n "My Alarm" -green my.alarm.success -yellow my.alarm.failure
  govc event.post -i my.alarm.failure $vm
  govc alarms $vm`
}

func (cmd *create) Run(ctx context.Context, f *flag.FlagSet) error {
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

	var or types.OrAlarmExpression

	expressions := []struct {
		status types.ManagedEntityStatus
		typeID string
	}{
		{types.ManagedEntityStatusGreen, cmd.green},
		{types.ManagedEntityStatusYellow, cmd.yellow},
		{types.ManagedEntityStatusRed, cmd.red},
	}

	for _, exp := range expressions {
		if exp.typeID != "" {
			or.Expression = append(or.Expression, &types.EventAlarmExpression{
				EventType:   "vim.event.EventEx",
				EventTypeId: exp.typeID,
				ObjectType:  "vim." + cmd.kind,
				Status:      exp.status,
			})
		}
	}

	cmd.AlarmSpec.Expression = &or

	m, err := alarm.GetManager(c)
	if err != nil {
		return err
	}

	if cmd.r {
		alarms, err := m.GetAlarm(ctx, obj)
		if err != nil {
			return err
		}

		var alarm *mo.Alarm
		for i := range alarms {
			if alarms[i].Info.Name == cmd.AlarmSpec.Name {
				alarm = &alarms[i]
				break
			}
		}
		if alarm == nil {
			return fmt.Errorf("%s not found", cmd.AlarmSpec.Name)
		}

		_, err = methods.ReconfigureAlarm(ctx, c, &types.ReconfigureAlarm{
			This: alarm.Self,
			Spec: &cmd.AlarmSpec,
		})
		return err
	}

	ref, err := m.CreateAlarm(ctx, obj, &cmd.AlarmSpec)
	if err != nil {
		return err
	}

	fmt.Println(ref.Value)

	return nil
}
