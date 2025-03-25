// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package flags

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/dougm/pretty"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/task"
	"github.com/vmware/govmomi/vim25/progress"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vim25/xml"
)

type OutputWriter interface {
	Write(io.Writer) error
}

type OutputFlag struct {
	common

	JSON bool
	XML  bool
	TTY  bool
	Dump bool
	Out  io.Writer
	Spec bool

	formatError  bool
	formatIndent bool
}

var outputFlagKey = flagKey("output")

func NewOutputFlag(ctx context.Context) (*OutputFlag, context.Context) {
	if v := ctx.Value(outputFlagKey); v != nil {
		return v.(*OutputFlag), ctx
	}

	v := &OutputFlag{Out: os.Stdout}
	ctx = context.WithValue(ctx, outputFlagKey, v)
	return v, ctx
}

func (flag *OutputFlag) Register(ctx context.Context, f *flag.FlagSet) {
	flag.RegisterOnce(func() {
		f.BoolVar(&flag.JSON, "json", false, "Enable JSON output")
		f.BoolVar(&flag.XML, "xml", false, "Enable XML output")
		f.BoolVar(&flag.Dump, "dump", false, "Enable Go output")
		if cli.ShowUnreleased() {
			f.BoolVar(&flag.Spec, "spec", false, "Output spec without sending request")
		}
		// Avoid adding more flags for now..
		flag.formatIndent = os.Getenv("GOVC_INDENT") != "false"      // Default to indented output
		flag.formatError = os.Getenv("GOVC_FORMAT_ERROR") != "false" // Default to formatted errors
	})
}

func (flag *OutputFlag) Process(ctx context.Context) error {
	return flag.ProcessOnce(func() error {
		if !flag.All() {
			// Assume we have a tty if not outputting JSON
			flag.TTY = true
		}

		return nil
	})
}

// Log outputs the specified string, prefixed with the current time.
// A newline is not automatically added. If the specified string
// starts with a '\r', the current line is cleared first.
func (flag *OutputFlag) Log(s string) (int, error) {
	if len(s) > 0 && s[0] == '\r' {
		flag.Write([]byte{'\r', 033, '[', 'K'})
		s = s[1:]
	}

	return flag.WriteString(time.Now().Format("[02-01-06 15:04:05] ") + s)
}

func (flag *OutputFlag) Write(b []byte) (int, error) {
	if !flag.TTY {
		return 0, nil
	}

	w := flag.Out
	if w == nil {
		w = os.Stdout
	}
	n, err := w.Write(b)
	if w == os.Stdout {
		os.Stdout.Sync()
	}
	return n, err
}

func (flag *OutputFlag) WriteString(s string) (int, error) {
	return flag.Write([]byte(s))
}

func (flag *OutputFlag) All() bool {
	return flag.JSON || flag.XML || flag.Dump
}

func dumpValue(val any) any {
	type dumper interface {
		Dump() any
	}

	if d, ok := val.(dumper); ok {
		return d.Dump()
	}

	rval := reflect.ValueOf(val)
	if rval.Type().Kind() != reflect.Ptr {
		return val
	}

	rval = rval.Elem()
	if rval.Type().Kind() == reflect.Struct {
		f := rval.Field(0)
		if f.Type().Kind() == reflect.Slice {
			// common case for the various 'type infoResult'
			if f.Len() == 1 {
				return f.Index(0).Interface()
			}
			return f.Interface()
		}

		if rval.NumField() == 1 && rval.Type().Field(0).Anonymous {
			// common case where govc type wraps govmomi type to implement OutputWriter
			return f.Interface()
		}
	}

	return val
}

type outputAny struct {
	Value any
}

func (*outputAny) Write(io.Writer) error {
	return nil
}

func (a *outputAny) Dump() any {
	return a.Value
}

func (flag *OutputFlag) WriteAny(val any) error {
	if !flag.All() {
		flag.XML = true
	}
	return flag.WriteResult(&outputAny{val})
}

func (flag *OutputFlag) WriteResult(result OutputWriter) error {
	var err error

	switch {
	case flag.Dump:
		format := "%#v\n"
		if flag.formatIndent {
			format = "%# v\n"
		}
		_, err = pretty.Fprintf(flag.Out, format, dumpValue(result))
	case flag.JSON:
		e := json.NewEncoder(flag.Out)
		if flag.formatIndent {
			e.SetIndent("", "  ")
		}
		err = e.Encode(result)
	case flag.XML:
		e := xml.NewEncoder(flag.Out)
		if flag.formatIndent {
			e.Indent("", "  ")
		}
		err = e.Encode(dumpValue(result))
		if err == nil {
			fmt.Fprintln(flag.Out)
		}
	default:
		err = result.Write(flag.Out)
	}

	return err
}

func (flag *OutputFlag) WriteError(err error) bool {
	if flag.formatError {
		flag.Out = os.Stderr
		return flag.WriteResult(&errorOutput{err}) == nil
	}
	return false
}

type errorOutput struct {
	error
}

func (e errorOutput) Write(w io.Writer) error {
	reason := e.Error()
	var messages []string
	var faults []types.LocalizableMessage

	switch err := e.error.(type) {
	case task.Error:
		faults = err.LocalizedMethodFault.Fault.GetMethodFault().FaultMessage
		if err.Description != nil {
			reason = fmt.Sprintf("%s (%s)", reason, err.Description.Message)
		}
	default:
		if soap.IsSoapFault(err) {
			detail := soap.ToSoapFault(err).Detail.Fault
			if f, ok := detail.(types.BaseMethodFault); ok {
				faults = f.GetMethodFault().FaultMessage
			}
		}
	}

	for _, m := range faults {
		if m.Message != "" && !strings.HasPrefix(m.Message, "[context]") {
			messages = append(messages, fmt.Sprintf("%s (%s)", m.Message, m.Key))
		}
	}

	messages = append(messages, reason)

	for _, message := range messages {
		if _, err := fmt.Fprintf(w, "%s: %s\n", os.Args[0], message); err != nil {
			return err
		}
	}

	return nil
}

func (e errorOutput) Dump() any {
	if f, ok := e.error.(task.Error); ok {
		return f.LocalizedMethodFault
	}
	if soap.IsSoapFault(e.error) {
		return soap.ToSoapFault(e.error)
	}
	if soap.IsVimFault(e.error) {
		return soap.ToVimFault(e.error)
	}
	return e
}

func (e errorOutput) canEncode() bool {
	switch e.error.(type) {
	case task.Error:
		return true
	}
	return soap.IsSoapFault(e.error) || soap.IsVimFault(e.error)
}

// errCannotEncode causes cli.Run to output err.Error() as it would without an error format specified
var errCannotEncode = errors.New("cannot encode error")

func (e errorOutput) MarshalJSON() ([]byte, error) {
	_, ok := e.error.(json.Marshaler)
	if ok || e.canEncode() {
		return json.Marshal(e.error)
	}
	return nil, errCannotEncode
}

func (e errorOutput) MarshalXML(encoder *xml.Encoder, start xml.StartElement) error {
	_, ok := e.error.(xml.Marshaler)
	if ok || e.canEncode() {
		return encoder.Encode(e.error)
	}
	return errCannotEncode
}

func (flag *OutputFlag) ProgressLogger(prefix string) *progress.ProgressLogger {
	return progress.NewProgressLogger(flag.Log, prefix)
}
