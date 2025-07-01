// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package task

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/types"
)

type set struct {
	*flags.ClientFlag

	desc     types.LocalizableMessage
	state    string
	err      string
	progress int
}

func init() {
	cli.Register("task.set", &set{}, true)
}

func (cmd *set) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.StringVar(&cmd.desc.Key, "d", "", "Task description key")
	f.StringVar(&cmd.desc.Message, "m", "", "Task description message")
	f.StringVar(&cmd.state, "s", "", "Task state")
	f.StringVar(&cmd.err, "e", "", "Task error")
	f.IntVar(&cmd.progress, "p", 0, "Task progress")
}

func (cmd *set) Description() string {
	return `Set task state.

Examples:
  id=$(govc task.create com.vmware.govmomi.simulator.test)
  govc task.set $id -s error
  govc task.set $id -p 100 -s success`
}

func (cmd *set) Usage() string {
	return "ID"
}

func (cmd *set) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}

	c, err := cmd.Client()
	if err != nil {
		return err
	}

	ref := types.ManagedObjectReference{Type: "Task"}
	if !ref.FromString(f.Arg(0)) {
		ref.Value = f.Arg(0)
	}

	task := object.NewTask(c, ref)

	if cmd.progress != 0 {
		err := task.UpdateProgress(ctx, cmd.progress)
		if err != nil {
			return err
		}
	}

	var fault *types.LocalizedMethodFault

	if cmd.err != "" {
		fault = &types.LocalizedMethodFault{
			Fault:            &types.SystemError{Reason: cmd.err},
			LocalizedMessage: cmd.err,
		}
		cmd.state = string(types.TaskInfoStateError)
	}

	if cmd.state != "" {
		err := task.SetState(ctx, types.TaskInfoState(cmd.state), nil, fault)
		if err != nil {
			return err
		}
	}

	if cmd.desc.Key != "" {
		err := task.SetDescription(ctx, cmd.desc)
		if err != nil {
			return err
		}
	}

	return nil
}
