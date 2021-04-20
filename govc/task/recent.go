/*
Copyright (c) 2017 VMware, Inc. All Rights Reserved.

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

package task

import (
	"context"
	"flag"
	"fmt"
	"io"
	"strings"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type recent struct {
	*flags.DatacenterFlag

	max    int
	follow bool
	long   bool
}

func init() {
	cli.Register("tasks", &recent{})
}

func (cmd *recent) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatacenterFlag, ctx = flags.NewDatacenterFlag(ctx)
	cmd.DatacenterFlag.Register(ctx, f)
	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)

	f.IntVar(&cmd.max, "n", 25, "Output the last N tasks")
	f.BoolVar(&cmd.follow, "f", false, "Follow recent task updates")
	f.BoolVar(&cmd.long, "l", false, "Use long task description")
}

func (cmd *recent) Description() string {
	return `Display info for recent tasks.

When a task has completed, the result column includes the task duration on success or
error message on failure.  If a task is still in progress, the result column displays
the completion percentage and the task ID.  The task ID can be used as an argument to
the 'task.cancel' command.

By default, all recent tasks are included (via TaskManager), but can be limited by PATH
to a specific inventory object.

Examples:
  govc tasks
  govc tasks -f
  govc tasks -f /dc1/host/cluster1`
}

func (cmd *recent) Usage() string {
	return "[PATH]"
}

func (cmd *recent) Process(ctx context.Context) error {
	if err := cmd.DatacenterFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func chop(s string, n int) string {
	if len(s) < n {
		return s
	}

	return s[:n-1] + "*"
}

// taskName describes the tasks similar to the ESX ui
func taskName(info *types.TaskInfo) string {
	name := strings.TrimSuffix(info.Name, "_Task")
	switch name {
	case "":
		return info.DescriptionId
	case "Destroy", "Rename":
		return info.Entity.Type + "." + name
	default:
		return name
	}
}

func (cmd *recent) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.Client()
	if err != nil {
		return err
	}

	m := c.ServiceContent.TaskManager

	tn := taskName

	if cmd.long {
		var o mo.TaskManager
		err = property.DefaultCollector(c).RetrieveOne(ctx, *m, []string{"description.methodInfo"}, &o)
		if err != nil {
			return err
		}

		desc := make(map[string]string, len(o.Description.MethodInfo))

		for _, entry := range o.Description.MethodInfo {
			info := entry.GetElementDescription()
			desc[info.Key] = info.Label
		}

		tn = func(info *types.TaskInfo) string {
			if name, ok := desc[info.DescriptionId]; ok {
				return name
			}

			return taskName(info)
		}
	}

	watch := *m

	if f.NArg() == 1 {
		refs, merr := cmd.ManagedObjects(ctx, f.Args())
		if merr != nil {
			return merr
		}
		watch = refs[0]
	}

	v, err := view.NewManager(c).CreateTaskView(ctx, &watch)
	if err != nil {
		return nil
	}

	defer func() {
		_ = v.Destroy(context.Background())
	}()

	v.Follow = cmd.follow

	var last string
	updated := false

	return v.Collect(ctx, func(tasks []types.TaskInfo) {
		if !updated && len(tasks) > cmd.max {
			tasks = tasks[len(tasks)-cmd.max:]
		}
		updated = true

		for _, info := range tasks {
			cmd.WriteResult(taskRecentResult{
				TaskInfo: info,
				TaskName: tn(&info),
				lastLine: &last,
			})
		}
	})
}

type taskRecentResult struct {
	TaskName string
	TaskInfo types.TaskInfo
	lastLine *string
}

const (
	timestampFormat = "15:04:05"
	tableTemplate   = "%-40s %-30s %-30s %9s %9s %9s %s\n"
)

func (trr taskRecentResult) Write(w io.Writer) error {
	if *trr.lastLine == "" {
		fmt.Fprintf(w, tableTemplate, "Task", "Target", "Initiator", "Queued", "Started", "Completed", "Result")
	}
	var user string

	switch x := trr.TaskInfo.Reason.(type) {
	case *types.TaskReasonUser:
		user = x.UserName
	}

	if trr.TaskInfo.EntityName == "" || user == "" {
		return nil
	}

	ruser := strings.SplitN(user, "\\", 2)
	if len(ruser) == 2 {
		user = ruser[1] // discard domain
	} else {
		user = strings.TrimPrefix(user, "com.vmware.") // e.g. com.vmware.vsan.health
	}

	queued := trr.TaskInfo.QueueTime.Format(timestampFormat)
	start := "-"
	end := start

	if trr.TaskInfo.StartTime != nil {
		start = trr.TaskInfo.StartTime.Format(timestampFormat)
	}

	msg := fmt.Sprintf("%2d%% %s", trr.TaskInfo.Progress, trr.TaskInfo.Task)

	if trr.TaskInfo.CompleteTime != nil {
		msg = trr.TaskInfo.CompleteTime.Sub(*trr.TaskInfo.StartTime).String()

		if trr.TaskInfo.State == types.TaskInfoStateError {
			msg = strings.TrimSuffix(trr.TaskInfo.Error.LocalizedMessage, ".")
		}

		end = trr.TaskInfo.CompleteTime.Format(timestampFormat)
	}

	result := fmt.Sprintf("%-7s [%s]", trr.TaskInfo.State, msg)

	item := fmt.Sprintf(tableTemplate, chop(trr.TaskName, 40), chop(trr.TaskInfo.EntityName, 30), chop(user, 30), queued, start, end, result)

	if item == *trr.lastLine {
		return nil // task info was updated, but the fields we display were not
	}
	*trr.lastLine = item

	fmt.Fprint(w, item)

	return nil
}
