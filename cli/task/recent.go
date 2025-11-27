// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package task

import (
	"context"
	"flag"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/task"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type recent struct {
	*flags.DatacenterFlag

	max    int
	follow bool
	long   bool

	state flags.StringList
	begin time.Duration
	end   time.Duration
	r     bool

	plain bool
}

func init() {
	cli.Register("tasks", &recent{})
}

func (cmd *recent) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatacenterFlag, ctx = flags.NewDatacenterFlag(ctx)
	cmd.DatacenterFlag.Register(ctx, f)

	f.IntVar(&cmd.max, "n", 25, "Output the last N tasks")
	f.BoolVar(&cmd.follow, "f", false, "Follow recent task updates")
	f.BoolVar(&cmd.long, "l", false, "Use long task description")
	f.Var(&cmd.state, "s", "Task states")
	f.DurationVar(&cmd.begin, "b", 0, "Begin time of task history")
	f.DurationVar(&cmd.end, "e", 0, "End time of task history")
	f.BoolVar(&cmd.r, "r", false, "Include child entities when PATH is specified")
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
  govc tasks        # tasks completed within the past 10 minutes
  govc tasks -b 24h # tasks completed within the past 24 hours
  govc tasks -s queued -s running # incomplete tasks
  govc tasks -s error -s success  # completed tasks
  govc tasks -r /dc1/vm/Namespaces # tasks for VMs in this Folder only
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
	return nil
}

// chop middle of s if len(s) > n
func chop(s string, n int) string {
	diff := len(s) - n
	if diff <= 0 {
		return s
	}
	diff /= 2
	m := len(s) / 2

	return s[:m-diff] + "*" + s[1+m+diff:]
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

type history struct {
	*task.HistoryCollector

	cmd *recent
}

func (h *history) Collect(ctx context.Context, f func([]types.TaskInfo)) error {
	for {
		tasks, err := h.ReadNextTasks(ctx, 10)
		if err != nil {
			return err
		}

		if len(tasks) == 0 {
			if h.cmd.follow {
				// TODO: this only follows new events.
				// need to watch TaskHistoryCollector.LatestPage for updates to existing Tasks
				time.Sleep(time.Second)
				continue
			}
			break
		}

		f(tasks)
	}
	return nil
}

type collector interface {
	Collect(context.Context, func([]types.TaskInfo)) error
	Destroy(context.Context) error
}

// useRecent returns true if any options are specified that require use of TaskHistoryCollector
func (cmd *recent) useRecent() bool {
	return cmd.begin == 0 && cmd.end == 0 && !cmd.r && len(cmd.state) == 0
}

func (cmd *recent) newCollector(ctx context.Context, c *vim25.Client, ref *types.ManagedObjectReference) (collector, error) {
	if cmd.useRecent() {
		// original flavor of this command that uses `RecentTask` instead of `TaskHistoryCollector`
		if ref == nil {
			ref = c.ServiceContent.TaskManager
		}

		v, err := view.NewManager(c).CreateTaskView(ctx, ref)
		if err != nil {
			return nil, err
		}

		v.Follow = cmd.follow && cmd.plain
		return v, nil
	}

	m := task.NewManager(c)
	r := types.TaskFilterSpecRecursionOptionSelf
	if ref == nil {
		ref = &c.ServiceContent.RootFolder
		cmd.r = true
	}

	now, err := methods.GetCurrentTime(ctx, c) // vCenter server time (UTC)
	if err != nil {
		return nil, err
	}

	if cmd.r {
		r = types.TaskFilterSpecRecursionOptionAll
	}

	if cmd.begin == 0 {
		cmd.begin = 10 * time.Minute
	}

	filter := types.TaskFilterSpec{
		Entity: &types.TaskFilterSpecByEntity{
			Entity:    *ref,
			Recursion: r,
		},
		Time: &types.TaskFilterSpecByTime{
			TimeType:  types.TaskFilterSpecTimeOptionStartedTime,
			BeginTime: types.NewTime(now.Add(-cmd.begin)),
		},
	}

	for _, state := range cmd.state {
		filter.State = append(filter.State, types.TaskInfoState(state))
	}

	if cmd.end != 0 {
		filter.Time.EndTime = types.NewTime(now.Add(-cmd.end))
	}

	collector, err := m.CreateCollectorForTasks(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &history{collector, cmd}, nil
}

func (cmd *recent) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() > 1 {
		return flag.ErrHelp
	}

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

	var watch *types.ManagedObjectReference

	if f.NArg() == 1 {
		refs, merr := cmd.ManagedObjects(ctx, f.Args())
		if merr != nil {
			return merr
		}
		if len(refs) != 1 {
			return fmt.Errorf("%s matches %d objects", f.Arg(0), len(refs))
		}
		watch = &refs[0]
	}

	// writes dump/json/xml once even if follow is specified, otherwise syntax error occurs
	cmd.plain = !(cmd.Dump || cmd.JSON || cmd.XML)

	v, err := cmd.newCollector(ctx, c, watch)
	if err != nil {
		return err
	}

	defer func() {
		_ = v.Destroy(context.Background())
	}()

	res := &taskResult{name: tn}
	if cmd.plain {
		res.WriteHeader(cmd.Out)
	}

	updated := false

	return v.Collect(ctx, func(tasks []types.TaskInfo) {
		if !updated && len(tasks) > cmd.max {
			tasks = tasks[len(tasks)-cmd.max:]
		}
		updated = true

		res.Tasks = tasks
		cmd.WriteResult(res)
	})
}

type taskResult struct {
	Tasks []types.TaskInfo `json:"tasks"`
	last  string
	name  func(info *types.TaskInfo) string
}

func (t *taskResult) WriteHeader(w io.Writer) {
	fmt.Fprint(w, t.format("Task", "Target", "Initiator", "Queued", "Started", "Completed", "Result"))
}

func (t *taskResult) Write(w io.Writer) error {
	stamp := "15:04:05"

	for _, info := range t.Tasks {
		var user string

		switch x := info.Reason.(type) {
		case *types.TaskReasonUser:
			user = x.UserName
		}

		if info.EntityName == "" || user == "" {
			continue
		}

		ruser := strings.SplitN(user, "\\", 2)
		if len(ruser) == 2 {
			user = ruser[1] // discard domain
		} else {
			user = strings.TrimPrefix(user, "com.vmware.") // e.g. com.vmware.vsan.health
		}

		queued := "-"
		start := "-"
		end := start

		if info.StartTime != nil {
			start = info.StartTime.Format(stamp)
			queued = info.StartTime.Sub(info.QueueTime).Round(time.Millisecond).String()
		}

		msg := fmt.Sprintf("%2d%% %s", info.Progress, info.Task)

		if info.CompleteTime != nil && info.StartTime != nil {
			msg = info.CompleteTime.Sub(*info.StartTime).String()

			if info.State == types.TaskInfoStateError {
				msg = strings.TrimSuffix(info.Error.LocalizedMessage, ".")
			}

			end = info.CompleteTime.Format(stamp)
		}

		result := fmt.Sprintf("%-7s [%s]", info.State, msg)

		item := t.format(chop(t.name(&info), 40), chop(info.EntityName, 30), chop(user, 30), queued, start, end, result)

		if item == t.last {
			continue // task info was updated, but the fields we display were not
		}
		t.last = item

		fmt.Fprint(w, item)
	}

	return nil
}

func (t *taskResult) format(task, target, initiator, queued, started, completed, result string) string {
	return fmt.Sprintf("%-40s %-30s %-30s %9s %9s %9s %s\n",
		task, target, initiator, queued, started, completed, result)
}
