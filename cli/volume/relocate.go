// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package volume

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	cnstypes "github.com/vmware/govmomi/cns/types"
	"github.com/vmware/govmomi/task"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

// defaultRelocateLogFile is the debug log file that task info and fault
// details are appended to on failure, relative to the current working
// directory, when -log-file is not given.
const defaultRelocateLogFile = "relocate.log"

type relocate struct {
	*flags.ClientFlag
	*flags.DatastoreFlag
	*flags.OutputFlag

	file    string
	logFile string

	// resultsWritten is set once the per-volume result table has been
	// printed, so WriteError can tell an already-reported per-volume
	// failure (in firstErr) apart from an earlier setup error (e.g. a
	// bad -ds, -file, or CNS client failure) that was never shown
	// anywhere.
	resultsWritten bool

	// logDisabled is set once cmd.logFile fails to open, so that failure
	// is only reported once (by logNotice or, if it happens later, by
	// logWrite) instead of on every subsequent logWrite call.
	logDisabled bool
}

func init() {
	cli.Register("volume.relocate", &relocate{})
}

func (cmd *relocate) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.DatastoreFlag, ctx = flags.NewDatastoreFlag(ctx)
	cmd.DatastoreFlag.Register(ctx, f)

	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)

	f.StringVar(&cmd.file, "file", "", "CSV file containing volume IDs and optional destination datastore names, one per row (mutually exclusive with ID arguments)")
	f.StringVar(&cmd.logFile, "log-file", defaultRelocateLogFile, "File that failed relocations' task info and fault details are appended to")
}

func (cmd *relocate) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.DatastoreFlag.Process(ctx); err != nil {
		return err
	}
	return cmd.OutputFlag.Process(ctx)
}

func (cmd *relocate) Usage() string {
	return "ID..."
}

// WriteError suppresses govc's default "<argv0>: <error>" line only once
// the per-volume result table has been printed (see relocateResult.Write):
// in that case the failure is already reported in the command's output and
// logged to the log file, so printing the first error again would be
// redundant. Errors that occur before the table is printed (e.g. a bad
// -ds/-file, or a CNS client failure) are not reported anywhere else, so
// they must still be printed normally.
func (cmd *relocate) WriteError(error) bool {
	return cmd.resultsWritten
}

func (cmd *relocate) Description() string {
	return `Relocate one or more CNS volumes to a target datastore.

IDs can be given as arguments, or read from a CSV file with -file. Each
CSV row has the volume ID as the first field and, optionally, the name of
the destination datastore as the second field. Rows without a datastore
field fall back to the datastore given by -ds. -file and ID arguments are
mutually exclusive.

Per-volume results are printed; the command exits non-zero if any relocation failed.
The per-volume result table is also appended to a log file (relocate.log in
the current working directory by default, or the file given by -log-file).
For any relocation that failed, the CNS task info is appended if available,
otherwise the full fault is appended. The log file's path is printed to
stderr.

Examples:
  govc volume.relocate -ds vsanDatastore f75989dc-95b9-4db7-af96-8583f24bc59d
  govc volume.relocate -ds vsanDatastore id1 id2 id3
  govc volume.relocate -ds vsanDatastore -json id1 id2 | jq .
  govc volume.relocate -ds vsanDatastore -file volumes.csv
  govc volume.relocate -ds vsanDatastore -log-file /var/log/relocate.log id1
  govc volume.relocate -ds vsanDatastore -file volumes.csv -log-file /var/log/relocate.log`
}

// relocateResult holds the per-volume outcomes of a batch relocation and
// implements flags.OutputWriter for both text and JSON rendering.
type relocateResult struct {
	Results []relocateVolumeResult `json:"results"`
}

type relocateVolumeResult struct {
	VolumeID  string `json:"volumeId"`
	Datastore string `json:"datastore"`
	OpID      string `json:"opId,omitempty"`
	Status    string `json:"status"`
	Error     string `json:"error,omitempty"`
}

func (r *relocateResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)
	fmt.Fprintf(tw, "%s\t%s\t%s\t%s\t%s\n", "Volume ID", "Datastore", "OPID", "Status", "Details")
	for _, res := range r.Results {
		if res.Error != "" {
			fmt.Fprintf(tw, "%s\t%s\t%s\tFAILED\t%s\n", res.VolumeID, res.Datastore, res.OpID, res.Error)
		} else {
			fmt.Fprintf(tw, "%s\t%s\t%s\tOK\n", res.VolumeID, res.Datastore, res.OpID)
		}
	}
	if err := tw.Flush(); err != nil {
		return err
	}
	succeeded, failed := r.counts()
	_, err := fmt.Fprintf(w, "\n%d volume(s) relocated successfully, %d failed\n", succeeded, failed)
	return err
}

// counts returns the number of successful and failed results.
func (r *relocateResult) counts() (succeeded, failed int) {
	for _, res := range r.Results {
		if res.Error != "" {
			failed++
		} else {
			succeeded++
		}
	}
	return
}

// consoleRowFormat renders the streamed console rows (and their header)
// with fixed column widths: unlike the log file's table, printed via
// tabwriter, these rows are written one at a time as each volume completes,
// so column widths can't be measured across the whole result set first.
const consoleRowFormat = "%-36s  %-20s  %-10s  %-6s  %s\n"

// printHeader writes the console row header, once, before any per-volume
// results are streamed. A no-op for -json/-xml/-dump, which render the full
// relocateResult at the end via WriteResult instead.
func (cmd *relocate) printHeader() {
	if cmd.All() {
		return
	}
	fmt.Fprintf(cmd.Out, consoleRowFormat, "Volume ID", "Datastore", "OPID", "Status", "Details")
}

// printResult writes a single volume's result straight to the console as
// soon as it's known, so progress is visible without waiting for the whole
// batch to finish. Only used for the default text output: -json/-xml/-dump
// still render the full relocateResult at the end via WriteResult, since
// those formats aren't meaningfully streamable. The log file (written via
// logResults/Write) is unaffected either way.
func (cmd *relocate) printResult(res relocateVolumeResult) {
	if cmd.All() {
		return
	}
	fmt.Fprintf(cmd.Out, consoleRowFormat, res.VolumeID, res.Datastore, res.OpID, res.Status, res.Error)
}

// relocateEntry is one unit of work: a volume ID and, optionally, the name
// of the datastore it should be relocated to. An empty Datastore falls
// back to the datastore given by -ds.
type relocateEntry struct {
	VolumeID  string
	Datastore string
}

// volumeIDPattern matches the UUID format used for CNS volume IDs. It is
// used to detect and skip an optional header row in a -file CSV input.
var volumeIDPattern = regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)

// relocateEntriesFromFile reads volume IDs and optional destination
// datastore names from a -file CSV input. warn receives a notice, one line
// per row, whenever the header-detection heuristic below skips a row: a
// row is only ever skipped as a header if it's the first row in the file
// and doesn't look like a CNS volume ID, and since that's a heuristic
// rather than an explicit "-header" flag, the caller needs to be told
// whenever it fires so a real ID that doesn't match the UUID pattern
// isn't silently dropped from the batch.
func relocateEntriesFromFile(path string, warn io.Writer) ([]relocateEntry, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.FieldsPerRecord = -1

	records, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("parsing %s: %w", path, err)
	}

	var entries []relocateEntry
	for i, rec := range records {
		if len(rec) == 0 {
			continue
		}

		id := strings.TrimSpace(rec[0])
		if id == "" {
			continue
		}
		if i == 0 && !volumeIDPattern.MatchString(id) {
			fmt.Fprintf(warn, "%s: row 1 (%s) does not look like a volume ID, treating it as a header row and skipping it\n", path, strings.Join(rec, ","))
			continue
		}

		entry := relocateEntry{VolumeID: id}
		if len(rec) > 1 {
			entry.Datastore = strings.TrimSpace(rec[1])
		}
		entries = append(entries, entry)
	}

	if len(entries) == 0 {
		return nil, fmt.Errorf("no volume IDs found in %s", path)
	}

	return entries, nil
}

// disableLogging is called once cmd.logFile fails to open, so the failure
// is surfaced (instead of the log content being silently discarded) and
// isn't attempted, or reported, again for the rest of the run.
func (cmd *relocate) disableLogging(err error) {
	cmd.logDisabled = true
	fmt.Fprintf(os.Stderr, "warning: could not open -log-file %s: %v; relocation details will not be logged\n", cmd.logFile, err)
}

// logWrite opens cmd.logFile for appending and hands it to fn, which is
// responsible for the actual log content.
func (cmd *relocate) logWrite(fn func(io.Writer)) {
	if cmd.logDisabled {
		return
	}

	f, err := os.OpenFile(cmd.logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		cmd.disableLogging(err)
		return
	}
	defer f.Close()

	fn(f)
}

// logNotice checks that cmd.logFile can be opened and, if so, prints the
// "Details logged to ..." notice to stderr, once, up front: Run always
// appends the full result table to cmd.logFile via logResults regardless
// of outcome, so the reader can be told where to look before any
// per-volume results start streaming to the console. If cmd.logFile can't
// be opened, disableLogging's warning is printed instead so the notice
// isn't printed for a log that will never actually be written.
func (cmd *relocate) logNotice() {
	f, err := os.OpenFile(cmd.logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		cmd.disableLogging(err)
		return
	}
	f.Close()

	path := cmd.logFile
	if abs, err := filepath.Abs(cmd.logFile); err == nil {
		path = abs
	}
	fmt.Fprintf(os.Stderr, "Details logged to %s\n", path)
}

// logDetail appends a timestamped, labeled value to cmd.logFile, for
// debugging that isn't fully captured by the summarized output shown for
// the command (e.g. faults, or CNS task details). It is not printed to the
// console.
func (cmd *relocate) logDetail(label string, v any) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return
	}

	cmd.logWrite(func(w io.Writer) {
		fmt.Fprintf(w, "%s %s:\n%s\n", time.Now().Format(time.RFC3339), label, b)
	})
}

// logResults appends the per-volume result table for the batch relocation
// to cmd.logFile, so a full record of every volume's outcome (not just the
// failures) is kept alongside the task and fault details logged for
// failures.
func (cmd *relocate) logResults(out *relocateResult) {
	cmd.logWrite(func(w io.Writer) {
		fmt.Fprintf(w, "relocate results:\n")
		out.Write(w)
	})
}

// logFault appends the full fault for a volume to cmd.logFile, for
// debugging failures that aren't fully captured by the summarized
// LocalizedMessage shown in command output. It is not printed to the
// console.
func (cmd *relocate) logFault(volumeID string, fault any) {
	cmd.logDetail(fmt.Sprintf("volume %s fault", volumeID), fault)
}

// logTaskInfo appends the CNS task info for a volume to cmd.logFile, so
// the task can be correlated with vpxd logs. It is not printed to the
// console.
func (cmd *relocate) logTaskInfo(volumeID string, v any) {
	cmd.logDetail(fmt.Sprintf("volume %s CNS task", volumeID), v)
}

// errorFault extracts the BaseMethodFault carried by an error returned from
// the CNS/vim25 APIs (e.g. soap.soapFaultError, soap.vimFaultError,
// task.Error), or nil if err doesn't carry one.
func errorFault(err error) types.BaseMethodFault {
	if f, ok := err.(interface{ Fault() types.BaseMethodFault }); ok {
		return f.Fault()
	}
	return nil
}

// errorFaultDetail returns the richest available fault representation for
// logging: the raw SOAP fault (faultcode, faultstring and the typed detail)
// for a synchronous API error, or the task's LocalizedMethodFault (fault
// plus its localized message) for a failed task. Returns nil if err doesn't
// carry a fault.
func errorFaultDetail(err error) any {
	if soap.IsSoapFault(err) {
		return soap.ToSoapFault(err)
	}

	var taskErr task.Error
	if errors.As(err, &taskErr) {
		return taskErr.LocalizedMethodFault
	}

	if soap.IsVimFault(err) {
		return soap.ToVimFault(err)
	}

	return nil
}

// faultOpID returns the value of the "opId" fault message argument, if
// present, which carries the vsanvcmgmtd operation ID for a failed
// CNS volume operation.
func faultOpID(fault types.BaseMethodFault) string {
	if fault == nil {
		return ""
	}
	for _, msg := range fault.GetMethodFault().FaultMessage {
		for _, arg := range msg.Arg {
			if arg.Key == "opId" {
				return fmt.Sprintf("%v", arg.Value)
			}
		}
	}
	return ""
}

func (cmd *relocate) Run(ctx context.Context, f *flag.FlagSet) error {
	if cmd.file != "" && f.NArg() > 0 {
		return errors.New("-file and ID arguments are mutually exclusive")
	}

	var entries []relocateEntry

	if cmd.file != "" {
		var err error
		entries, err = relocateEntriesFromFile(cmd.file, os.Stderr)
		if err != nil {
			return err
		}
	} else {
		if f.NArg() == 0 {
			return flag.ErrHelp
		}
		for _, id := range f.Args() {
			entries = append(entries, relocateEntry{VolumeID: id})
		}
	}

	var defaultDS types.ManagedObjectReference
	var defaultDSName string
	for _, e := range entries {
		if e.Datastore == "" {
			ds, err := cmd.Datastore()
			if err != nil {
				return err
			}
			defaultDS = ds.Reference()
			defaultDSName = ds.Name()
			break
		}
	}

	finder, err := cmd.Finder()
	if err != nil {
		return err
	}

	c, err := cmd.CnsClient()
	if err != nil {
		return err
	}

	out := new(relocateResult)
	var firstErr error
	// dsInfo caches both outcomes of resolving a -file row's datastore
	// name via finder.Datastore: a resolved reference, or the lookup
	// error. Caching the error too matters as much as caching the
	// reference: a mistyped datastore name repeated across many CSV rows
	// would otherwise trigger a redundant remote lookup for every row
	// instead of once.
	type dsInfo struct {
		ref  types.ManagedObjectReference
		name string
		err  error
	}
	dsCache := make(map[string]*dsInfo)

	cmd.logNotice()
	cmd.printHeader()

	for _, e := range entries {
		dsRef := defaultDS
		dsName := defaultDSName
		if e.Datastore != "" {
			info, ok := dsCache[e.Datastore]
			if !ok {
				info = new(dsInfo)
				if ds, err := finder.Datastore(ctx, e.Datastore); err != nil {
					info.err = err
				} else {
					info.ref = ds.Reference()
					info.name = ds.Name()
				}
				dsCache[e.Datastore] = info
			}
			if info.err != nil {
				res := relocateVolumeResult{
					VolumeID:  e.VolumeID,
					Datastore: e.Datastore,
					Status:    "FAILED",
					Error:     info.err.Error(),
				}
				out.Results = append(out.Results, res)
				cmd.printResult(res)
				if firstErr == nil {
					firstErr = info.err
				}
				continue
			}
			dsRef = info.ref
			dsName = info.name
		}

		spec := cnstypes.CnsBlockVolumeRelocateSpec{
			CnsVolumeRelocateSpec: cnstypes.CnsVolumeRelocateSpec{
				VolumeId:  cnstypes.CnsVolumeId{Id: e.VolumeID},
				Datastore: dsRef,
			},
		}

		entry := relocateVolumeResult{VolumeID: e.VolumeID, Datastore: dsName, Status: "OK"}

		task, err := c.RelocateVolume(ctx, spec)
		if err != nil {
			entry.Status = "FAILED"
			entry.Error = err.Error()
			if detail := errorFaultDetail(err); detail != nil {
				cmd.logFault(e.VolumeID, detail)
			}
			if fault := errorFault(err); fault != nil {
				if id := faultOpID(fault); id != "" {
					entry.OpID = id
				}
			}
			if firstErr == nil {
				firstErr = err
			}
			out.Results = append(out.Results, entry)
			cmd.printResult(entry)
			continue
		}
		info, err := task.WaitForResult(ctx, nil)
		if info != nil {
			entry.OpID = info.ActivationId
		}
		if err != nil {
			entry.Status = "FAILED"
			entry.Error = err.Error()
			if info != nil {
				cmd.logTaskInfo(e.VolumeID, info)
			} else if detail := errorFaultDetail(err); detail != nil {
				cmd.logFault(e.VolumeID, detail)
			}
			if fault := errorFault(err); fault != nil {
				if id := faultOpID(fault); id != "" {
					entry.OpID = id
				}
			}
			if firstErr == nil {
				firstErr = err
			}
			out.Results = append(out.Results, entry)
			cmd.printResult(entry)
			continue
		}

		if batchRes, ok := info.Result.(cnstypes.CnsVolumeOperationBatchResult); ok {
			for _, r := range batchRes.VolumeResults {
				opRes := r.GetCnsVolumeOperationResult()
				if opRes.Fault != nil {
					entry.Status = "FAILED"
					entry.Error = opRes.Fault.LocalizedMessage
					if info != nil {
						cmd.logTaskInfo(e.VolumeID, info)
					} else {
						cmd.logFault(e.VolumeID, opRes.Fault)
					}
					if id := faultOpID(opRes.Fault.Fault); id != "" {
						entry.OpID = id
					}
					if firstErr == nil {
						if opRes.Fault.Fault != nil {
							firstErr = soap.WrapVimFault(opRes.Fault.Fault)
						} else {
							firstErr = errors.New(opRes.Fault.LocalizedMessage)
						}
					}
				}
			}
		}

		out.Results = append(out.Results, entry)
		cmd.printResult(entry)
	}

	cmd.logResults(out)

	if cmd.All() {
		if wErr := cmd.WriteResult(out); wErr != nil {
			return wErr
		}
	} else {
		succeeded, failed := out.counts()
		fmt.Fprintf(cmd.Out, "\n%d volume(s) relocated successfully, %d failed\n", succeeded, failed)
	}
	cmd.resultsWritten = true

	return firstErr
}
