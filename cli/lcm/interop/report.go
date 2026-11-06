// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package interop

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/lcm"
)

type report struct {
	*flags.ClientFlag
	*flags.OutputFlag
}

func init() {
	cli.Register("lcm.interop.report", &report{}, true)
}

func (cmd *report) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)
}

func (cmd *report) Description() string {
	return `Fetch and display the interoperability report created for a task.

The task must have been created by lcm.interop.check. If the task is not yet
complete, the current status is printed and the command exits with an error.

Examples:
  govc lcm.interop.report <task-id>
  govc lcm.interop.report -json <task-id>`
}

func (cmd *report) Usage() string {
	return "TASK_ID"
}

func (cmd *report) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return cmd.OutputFlag.Process(ctx)
}

func (cmd *report) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}
	taskID := f.Arg(0)

	m := newManager(cmd.ClientFlag)

	info, err := m.GetTask(ctx, taskID)
	if err != nil {
		return err
	}

	if !info.IsDone() {
		return fmt.Errorf("task %s is not yet complete (status: %s, progress: %d/%d)",
			taskID, info.Status, info.Progress.Completed, info.Progress.Total)
	}

	if err := info.Err(); err != nil {
		return fmt.Errorf("task %s failed: %w", taskID, err)
	}

	interopResult, err := lcm.ParseResult(info)
	if err != nil {
		return err
	}

	return cmd.WriteResult(&reportOutput{result: interopResult})
}

// reportOutput implements OutputWriter for JSON/text rendering.
// JSON output is scoped to the report object only.
type reportOutput struct {
	result *lcm.InteropResult
}

func (o *reportOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.result)
}

func (o *reportOutput) Write(w io.Writer) error {
	return printReport(w, o.result)
}

// printReport writes a human-readable interop report table.
func printReport(w io.Writer, r *lcm.InteropResult) error {
	if r == nil {
		fmt.Fprintln(w, "No result payload in task.")
		return nil
	}

	if r.Issues != nil {
		printIssues(w, "ERROR", r.Issues.Errors)
		printIssues(w, "WARNING", r.Issues.Warnings)
		printIssues(w, "INFO", r.Issues.Info)
		if len(r.Issues.Errors)+len(r.Issues.Warnings)+len(r.Issues.Info) > 0 {
			fmt.Fprintln(w)
		}
	}

	fmt.Fprintf(w, "Report date: %s\n\n", r.Report.DateCreated.Format("2006-01-02 15:04:05 UTC"))

	tw := tabwriter.NewWriter(w, 0, 0, 2, ' ', 0)

	for _, row := range r.Report.Rows {
		target := row.TargetComponent
		fmt.Fprintf(tw, "Target:\t%s %s\n", target.Product.DisplayName, target.Version)
		fmt.Fprintf(tw, "  Compatible:\t%d\tIncompatible:\t%d\tUnknown:\t%d\n",
			row.Summary.CompatibleCount,
			row.Summary.IncompatibleCount,
			row.Summary.UnknownCount,
		)

		if len(row.AssociatedComponents) > 0 {
			fmt.Fprintln(tw)
			fmt.Fprintf(tw, "  %-40s\t%-15s\t%s\n", "Component", "Status", "Compatible Releases")
			fmt.Fprintf(tw, "  %-40s\t%-15s\t%s\n", "---------", "------", "-------------------")
			printAssociatedComponents(tw, row.AssociatedComponents, "  ")
		}
		fmt.Fprintln(tw)
	}

	return tw.Flush()
}

func printIssues(w io.Writer, level string, notifications []lcm.InteropNotification) {
	for _, n := range notifications {
		fmt.Fprintf(w, "[%s] %s (id: %s)\n", level, n.Message.DefaultMessage, n.Id)
		if n.Resolution != nil {
			fmt.Fprintf(w, "         Resolution: %s\n", n.Resolution.DefaultMessage)
		}
	}
}

func printAssociatedComponents(w io.Writer, components []lcm.InteropAssociatedComponent, indent string) {
	for _, assoc := range components {
		version := ""
		if assoc.Component.Version != nil {
			version = *assoc.Component.Version
		}

		releases := make([]string, 0, len(assoc.CompatibleReleases))
		for _, r := range assoc.CompatibleReleases {
			releases = append(releases, r.Version)
		}
		compatStr := "-"
		if len(releases) > 0 {
			compatStr = fmt.Sprintf("%v", releases)
		}

		name := assoc.Component.Product.DisplayName
		if version != "" {
			name = fmt.Sprintf("%s %s", name, version)
		}

		fmt.Fprintf(w, "%s%-40s\t%-15s\t%s\n", indent, name, assoc.Status, compatStr)

		if len(assoc.LinkedComponents) > 0 {
			printAssociatedComponents(w, assoc.LinkedComponents, indent+"  ")
		}
	}
}
